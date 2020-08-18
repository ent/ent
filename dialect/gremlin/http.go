// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/facebook/ent/dialect/gremlin/encoding/graphson"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

type httpTransport struct {
	client *http.Client
	url    string
}

// NewHTTPTransport returns a new http transport.
func NewHTTPTransport(urlStr string, client *http.Client) (RoundTripper, error) {
	u, err := url.Parse(urlStr)
	if err != nil {
		return nil, errors.Wrap(err, "gremlin/http: parsing url")
	}
	if client == nil {
		client = http.DefaultClient
	}
	return &httpTransport{client, u.String()}, nil
}

// RoundTrip implements RouterTripper interface.
func (t *httpTransport) RoundTrip(ctx context.Context, req *Request) (*Response, error) {
	if req.Operation != OpsEval {
		return nil, errors.Errorf("gremlin/http: unsupported operation: %q", req.Operation)
	}
	if _, ok := req.Arguments[ArgsGremlin]; !ok {
		return nil, errors.New("gremlin/http: missing query expression")
	}

	pr, pw := io.Pipe()
	defer pr.Close()
	go func() {
		err := jsoniter.NewEncoder(pw).Encode(req.Arguments)
		_ = pw.CloseWithError(errors.Wrap(err, "gremlin/http: encoding request"))
	}()

	var br io.Reader
	{
		req, err := http.NewRequest(http.MethodPost, t.url, pr)
		if err != nil {
			return nil, errors.Wrap(err, "gremlin/http: creating http request")
		}
		req.Header.Set("Content-Type", "application/json")

		rsp, err := t.client.Do(req.WithContext(ctx))
		if err != nil {
			return nil, errors.Wrap(err, "gremlin/http: posting http request")
		}
		defer rsp.Body.Close()

		if rsp.StatusCode < http.StatusOK || rsp.StatusCode > http.StatusPartialContent {
			body, _ := ioutil.ReadAll(rsp.Body)
			return nil, errors.Errorf("gremlin/http: status=%q, body=%q", rsp.Status, body)
		}
		if rsp.ContentLength > MaxResponseSize {
			return nil, errors.New("gremlin/http: context length exceeds limit")
		}
		br = rsp.Body
	}

	var rsp Response
	if err := graphson.NewDecoder(io.LimitReader(br, MaxResponseSize)).Decode(&rsp); err != nil {
		return nil, errors.Wrap(err, "gremlin/http: decoding response")
	}
	return &rsp, nil
}
