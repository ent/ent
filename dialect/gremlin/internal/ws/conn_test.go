// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ws

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type conn struct{ *websocket.Conn }

func (c conn) ReadRequest() (*gremlin.Request, error) {
	_, data, err := c.ReadMessage()
	if err != nil {
		return nil, err
	}
	var req gremlin.Request
	if err := graphson.Unmarshal(data[data[0]+1:], &req); err != nil {
		return nil, err
	}
	return &req, nil
}

func (c conn) WriteResponse(rsp *gremlin.Response) error {
	data, err := graphson.Marshal(rsp)
	if err != nil {
		return err
	}
	return c.WriteMessage(websocket.BinaryMessage, data)
}

func serve(handler func(conn)) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
		c, _ := upgrader.Upgrade(w, r, nil)
		defer c.Close()
		handler(conn{c})
		for {
			_, _, err := c.ReadMessage()
			if err != nil {
				break
			}
		}
	}))
}

func TestConnectClosure(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()

	srv := serve(func(conn conn) {
		defer wg.Done()
		_, _, err := conn.ReadMessage()
		assert.True(t, websocket.IsCloseError(err, websocket.CloseNormalClosure))
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)

	err = conn.Close()
	assert.NoError(t, err)

	_, err = conn.Execute(context.Background(), gremlin.NewEvalRequest("g.V()"))
	assert.EqualError(t, err, ErrConnClosed.Error())
}

func TestSimpleQuery(t *testing.T) {
	srv := serve(func(conn conn) {
		typ, data, err := conn.ReadMessage()
		require.NoError(t, err)
		assert.Equal(t, websocket.BinaryMessage, typ)

		var req gremlin.Request
		err = graphson.Unmarshal(data[data[0]+1:], &req)
		require.NoError(t, err)
		assert.Equal(t, "g.V()", req.Arguments["gremlin"])

		rsp := gremlin.Response{RequestID: req.RequestID}
		rsp.Status.Code = gremlin.StatusNoContent
		err = conn.WriteResponse(&rsp)
		require.NoError(t, err)
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer assert.Condition(t, func() bool { return assert.NoError(t, conn.Close()) })

	rsp, err := conn.Execute(context.Background(), gremlin.NewEvalRequest("g.V()"))
	assert.NoError(t, err)
	require.NotNil(t, rsp)
	assert.Equal(t, gremlin.StatusNoContent, rsp.Status.Code)
}

func TestDuplicateRequest(t *testing.T) {
	// skip until flakiness will be fixed.
	t.SkipNow()
	srv := serve(func(conn conn) {
		req, err := conn.ReadRequest()
		require.NoError(t, err)

		rsp := gremlin.Response{RequestID: req.RequestID}
		rsp.Status.Code = gremlin.StatusNoContent
		err = conn.WriteResponse(&rsp)
		require.NoError(t, err)
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	var errors [2]error
	req := gremlin.NewEvalRequest("g.V()")

	var wg sync.WaitGroup
	wg.Add(len(errors))

	for i := range errors {
		go func(i int) {
			_, errors[i] = conn.Execute(context.Background(), req)
			wg.Done()
		}(i)
	}
	wg.Wait()

	err = errors[0]
	if err == nil {
		err = errors[1]
	}
	assert.EqualError(t, err, ErrDuplicateRequest.Error())
}

func TestConnectCancellation(t *testing.T) {
	srv := serve(func(conn) {})
	defer srv.Close()

	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	conn, err := DefaultDialer.DialContext(ctx, "ws://"+srv.Listener.Addr().String())
	assert.Error(t, err)
	assert.Nil(t, conn)
}

func TestQueryCancellation(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	srv := serve(func(conn conn) {
		if _, _, err := conn.ReadMessage(); err == nil {
			cancel()
		}
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	_, err = conn.Execute(ctx, gremlin.NewEvalRequest("g.E()"))
	assert.EqualError(t, err, context.Canceled.Error())
}

func TestBadResponse(t *testing.T) {
	tests := []struct {
		name   string
		mangle func(*gremlin.Response) *gremlin.Response
	}{
		{
			name: "NoStatus",
			mangle: func(rsp *gremlin.Response) *gremlin.Response {
				return rsp
			},
		},
		{
			name: "Malformed",
			mangle: func(rsp *gremlin.Response) *gremlin.Response {
				rsp.Status.Code = gremlin.StatusMalformedRequest
				rsp.Status.Message = "bad request"
				return rsp
			},
		},
		{
			name: "Unknown",
			mangle: func(rsp *gremlin.Response) *gremlin.Response {
				rsp.Status.Code = 424242
				return rsp
			},
		},
	}

	srv := serve(func(conn conn) {
		for {
			req, err := conn.ReadRequest()
			if err != nil {
				break
			}

			idx, err := strconv.ParseInt(req.Arguments["gremlin"].(string), 10, 0)
			require.NoError(t, err)

			err = conn.WriteResponse(tests[idx].mangle(&gremlin.Response{RequestID: req.RequestID}))
			require.NoError(t, err)
		}
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(len(tests))

	ctx := context.Background()
	for i, tc := range tests {
		i, tc := i, tc
		t.Run(tc.name, func(t *testing.T) {
			defer wg.Done()
			rsp, err := conn.Execute(ctx, gremlin.NewEvalRequest(strconv.FormatInt(int64(i), 10)))
			assert.NoError(t, err)
			assert.True(t, rsp.IsErr())
		})
	}
	wg.Wait()
}

func TestServerHangup(t *testing.T) {
	// skip until flakiness will be fixed.
	t.SkipNow()
	srv := serve(func(conn conn) { _ = conn.Close() })
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	_, err = conn.Execute(context.Background(), gremlin.NewEvalRequest("g.V()"))
	assert.EqualError(t, err, ErrConnClosed.Error())
	assert.Error(t, conn.ctx.Err())
}

func TestCanceledLongRequest(t *testing.T) {
	// skip until flakiness will be fixed.
	t.SkipNow()
	ctx, cancel := context.WithCancel(context.Background())
	srv := serve(func(conn conn) {
		var responses [3]*gremlin.Response
		for i := 0; i < len(responses); i++ {
			req, err := conn.ReadRequest()
			require.NoError(t, err)

			rsp := gremlin.Response{RequestID: req.RequestID}
			rsp.Status.Code = gremlin.StatusSuccess
			rsp.Result.Data = graphson.RawMessage(`"ok"`)
			responses[i] = &rsp
		}

		cancel()

		responses[0], responses[2] = responses[2], responses[0]
		for i := 0; i < len(responses); i++ {
			err := conn.WriteResponse(responses[i])
			require.NoError(t, err)
		}
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	var wg sync.WaitGroup
	wg.Add(3)
	defer wg.Wait()

	for i := 0; i < 3; i++ {
		go func(ctx context.Context, idx int) {
			defer wg.Done()
			rsp, err := conn.Execute(ctx, gremlin.NewEvalRequest("g.V()"))
			if idx > 0 {
				assert.NoError(t, err)
				assert.EqualValues(t, []byte(`"ok"`), rsp.Result.Data)
			} else {
				assert.EqualError(t, err, context.Canceled.Error())
			}
		}(ctx, i)
		ctx = context.Background()
	}
}

func TestPartialResponse(t *testing.T) {
	type kv struct {
		Key   string
		Value int
	}
	kvs := []kv{
		{"one", 1},
		{"two", 2},
		{"three", 3},
	}
	srv := serve(func(conn conn) {
		req, err := conn.ReadRequest()
		require.NoError(t, err)

		for i := range kvs {
			data, err := graphson.Marshal([]kv{kvs[i]})
			require.NoError(t, err)

			rsp := gremlin.Response{RequestID: req.RequestID}
			rsp.Result.Data = graphson.RawMessage(data)

			if i != len(kvs)-1 {
				rsp.Status.Code = gremlin.StatusPartialContent
			} else {
				rsp.Status.Code = gremlin.StatusSuccess
			}

			err = conn.WriteResponse(&rsp)
			require.NoError(t, err)
		}
	})
	defer srv.Close()

	conn, err := DefaultDialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer conn.Close()

	rsp, err := conn.Execute(context.Background(), gremlin.NewEvalRequest("g.E()"))
	assert.NoError(t, err)

	var result []kv
	err = graphson.Unmarshal(rsp.Result.Data, &result)
	require.NoError(t, err)
	assert.Equal(t, kvs, result)
}

func TestAuthentication(t *testing.T) {
	user, pass := "username", "password"
	srv := serve(func(conn conn) {
		req, err := conn.ReadRequest()
		require.NoError(t, err)

		rsp := gremlin.Response{RequestID: req.RequestID}
		rsp.Status.Code = gremlin.StatusAuthenticate
		err = conn.WriteResponse(&rsp)
		require.NoError(t, err)

		areq, err := conn.ReadRequest()
		require.NoError(t, err)

		var acreds gremlin.Credentials
		err = acreds.UnmarshalText([]byte(areq.Arguments["sasl"].(string)))
		assert.NoError(t, err)
		areq.Arguments["sasl"] = acreds
		assert.Equal(t, gremlin.NewAuthRequest(req.RequestID, user, pass), areq)

		rsp = gremlin.Response{RequestID: req.RequestID}
		rsp.Status.Code = gremlin.StatusNoContent
		err = conn.WriteResponse(&rsp)
		require.NoError(t, err)
	})
	defer srv.Close()

	dialer := *DefaultDialer
	dialer.user = user
	dialer.pass = pass

	client, err := dialer.Dial("ws://" + srv.Listener.Addr().String())
	require.NoError(t, err)
	defer client.Close()

	_, err = client.Execute(context.Background(), gremlin.NewEvalRequest("g.E().drop()"))
	assert.NoError(t, err)
}
