// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"fmt"
	"net/http"
)

// RoundTripper is an interface representing the ability to execute a
// single gremlin transaction, obtaining the Response for a given Request.
type RoundTripper interface {
	RoundTrip(context.Context, *Request) (*Response, error)
}

// The RoundTripperFunc type is an adapter to allow the use of ordinary functions as Gremlin RoundTripper.
type RoundTripperFunc func(context.Context, *Request) (*Response, error)

// RoundTrip calls f(ctx, r).
func (f RoundTripperFunc) RoundTrip(ctx context.Context, r *Request) (*Response, error) {
	return f(ctx, r)
}

// Interceptor provides a hook to intercept the execution of a Gremlin Request.
type Interceptor func(RoundTripper) RoundTripper

// A Client is a gremlin client.
type Client struct {
	// Transport specifies the mechanism by which individual
	// Gremlin requests are made.
	Transport RoundTripper
}

// MaxResponseSize defines the maximum response size allowed.
const MaxResponseSize = 2 << 20

// NewClient creates a gremlin client from config and options.
func NewClient(cfg Config, opt ...Option) (*Client, error) {
	return cfg.Build(opt...)
}

// NewHTTPClient creates an http based gremlin client.
func NewHTTPClient(url string, client *http.Client) (*Client, error) {
	transport, err := NewHTTPTransport(url, client)
	if err != nil {
		return nil, err
	}
	return &Client{transport}, nil
}

// Do sends a gremlin request and returns a gremlin response.
func (c Client) Do(ctx context.Context, req *Request) (*Response, error) {
	rsp, err := c.Transport.RoundTrip(ctx, req)
	if err == nil {
		err = rsp.Err()
	}

	// If we got an error, and the context has been canceled,
	// the context's error is probably more useful.
	if err != nil && ctx.Err() != nil {
		err = ctx.Err()
	}
	return rsp, err
}

// Query issues an eval request via the Do function.
func (c Client) Query(ctx context.Context, query string) (*Response, error) {
	return c.Do(ctx, NewEvalRequest(query))
}

// Queryf formats a query string and invokes Query.
func (c Client) Queryf(ctx context.Context, format string, args ...any) (*Response, error) {
	return c.Query(ctx, fmt.Sprintf(format, args...))
}
