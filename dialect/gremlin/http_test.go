// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	jsoniter "github.com/json-iterator/go"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHTTPTransportRoundTripper(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Header.Get("Content-Type"), "application/json")
		got, err := io.ReadAll(r.Body)
		require.NoError(t, err)
		assert.JSONEq(t, `{"gremlin": "g.V(1)", "language": "gremlin-groovy"}`, string(got))

		_, err = io.WriteString(w, `{
			"requestId": "f679127f-8701-425c-af55-049a44720db6",
			"result": {
				"data": {
					"@type": "g:List",
					"@value": [
						{
							"@type": "g:Vertex",
							"@value": {
								"id": {
									"@type": "g:Int64",
									"@value": 1
								},
								"label": "person"
							}
						}
					]
				},
				"meta": {
					"@type": "g:Map",
					"@value": []
				}
			},
			"status": {
				"attributes": {
					"@type": "g:Map",
					"@value": []
				},
				"code": 200,
				"message": ""
			}
		}`)
		require.NoError(t, err)
	}))
	defer srv.Close()

	transport, err := NewHTTPTransport(srv.URL, nil)
	require.NoError(t, err)

	rsp, err := transport.RoundTrip(context.Background(), NewEvalRequest("g.V(1)"))
	require.NoError(t, err)

	assert.Equal(t, "f679127f-8701-425c-af55-049a44720db6", rsp.RequestID)
	assert.Equal(t, 200, rsp.Status.Code)
	assert.Empty(t, rsp.Status.Message)

	v := jsoniter.Get(rsp.Result.Data, graphson.ValueKey, 0, graphson.ValueKey)
	require.NoError(t, v.LastError())
	assert.Equal(t, 1, v.Get("id", graphson.ValueKey).ToInt())
	assert.Equal(t, "person", v.Get("label").ToString())
}

func TestNewHTTPTransportBadURL(t *testing.T) {
	transport, err := NewHTTPTransport(":", nil)
	assert.Nil(t, transport)
	assert.Error(t, err)
}

func TestHTTPTransportBadRequest(t *testing.T) {
	transport, err := NewHTTPTransport("example.com", nil)
	require.NoError(t, err)

	req := NewEvalRequest("g.V()")
	req.Operation = ""
	rsp, err := transport.RoundTrip(context.Background(), req)
	assert.EqualError(t, err, `gremlin/http: unsupported operation: ""`)
	assert.Nil(t, rsp)

	req = NewEvalRequest("g.V()")
	delete(req.Arguments, ArgsGremlin)
	rsp, err = transport.RoundTrip(context.Background(), req)
	assert.EqualError(t, err, "gremlin/http: missing query expression")
	assert.Nil(t, rsp)
}

func TestHTTPTransportBadResponseStatus(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer srv.Close()

	transport, err := NewHTTPTransport(srv.URL, nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(context.Background(), NewEvalRequest("g.E()."))
	require.Error(t, err)
	assert.Contains(t, err.Error(), http.StatusText(http.StatusInternalServerError))
}

func TestHTTPTransportBadResponseBody(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		_, err := io.WriteString(w, "{{{")
		require.NoError(t, err)
	}))
	defer srv.Close()

	transport, err := NewHTTPTransport(srv.URL, nil)
	require.NoError(t, err)

	_, err = transport.RoundTrip(context.Background(), NewEvalRequest("g.E()."))
	require.Error(t, err)
	assert.Contains(t, err.Error(), "decoding response")
}
