// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"io"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewClient(t *testing.T) {
	var cfg Config
	cfg.Endpoint.URL, _ = url.Parse("http://gremlin-server/gremlin")
	c, err := NewClient(cfg)
	assert.NotNil(t, c)
	assert.NoError(t, err)
}

type mockRoundTripper struct{ mock.Mock }

func (m *mockRoundTripper) RoundTrip(ctx context.Context, req *Request) (*Response, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*Response), args.Error(1)
}

func TestClientRequest(t *testing.T) {
	ctx := context.Background()
	req, rsp := &Request{}, &Response{}

	var m mockRoundTripper
	m.On("RoundTrip", ctx, req).
		Run(func(mock.Arguments) { rsp.Status.Code = StatusSuccess }).
		Return(rsp, nil).
		Once()
	defer m.AssertExpectations(t)

	response, err := Client{&m}.Do(context.Background(), req)
	assert.NoError(t, err)
	assert.Equal(t, rsp, response)
}

func TestClientResponseError(t *testing.T) {
	rsp := &Response{}

	var m mockRoundTripper
	m.On("RoundTrip", mock.Anything, mock.Anything).
		Run(func(mock.Arguments) { rsp.Status.Code = StatusServerError }).
		Return(rsp, nil).
		Once()
	defer m.AssertExpectations(t)

	_, err := Client{&m}.Do(context.Background(), nil)
	assert.Error(t, err)
}

func TestClientCanceledContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	var m mockRoundTripper
	m.On("RoundTrip", ctx, mock.Anything).
		Run(func(mock.Arguments) { cancel() }).
		Return(&Response{}, io.ErrUnexpectedEOF).
		Once()
	defer m.AssertExpectations(t)

	_, err := Client{&m}.Query(ctx, "g.E()")
	assert.EqualError(t, err, context.Canceled.Error())
}

func TestClientQuery(t *testing.T) {
	rsp := &Response{}
	rsp.Status.Code = StatusNoContent

	var m mockRoundTripper
	m.On("RoundTrip", mock.Anything, mock.Anything).
		Run(func(args mock.Arguments) {
			req := args.Get(1).(*Request)
			assert.Equal(t, "g.V(1)", req.Arguments[ArgsGremlin])
		}).
		Return(rsp, nil).
		Once()
	defer m.AssertExpectations(t)

	rsp, err := Client{&m}.Queryf(context.Background(), "g.V(%d)", 1)
	assert.NotNil(t, rsp)
	assert.NoError(t, err)
}
