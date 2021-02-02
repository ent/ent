// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"context"

	"entgo.io/ent/dialect/gremlin"

	"go.opencensus.io/trace"
)

// Transport is an gremlin.RoundTripper that instruments all outgoing requests with
// OpenCensus stats and tracing.
type Transport struct {
	// Base is a wrapped gremlin.RoundTripper that does the actual requests.
	Base gremlin.RoundTripper

	// StartOptions are applied to the span started by this Transport around each
	// request.
	//
	// StartOptions.SpanKind will always be set to trace.SpanKindClient
	// for spans started by this transport.
	StartOptions trace.StartOptions

	// GetStartOptions allows to set start options per request. If set,
	// StartOptions is going to be ignored.
	GetStartOptions func(context.Context, *gremlin.Request) trace.StartOptions

	// NameFromRequest holds the function to use for generating the span name
	// from the information found in the outgoing Gremlin Request. By default the
	// name equals the URL Path.
	FormatSpanName func(context.Context, *gremlin.Request) string

	// WithQuery, if set to true, will enable recording of gremlin queries in spans.
	// Only allow this if it is safe to have queries recorded with respect to
	// security.
	WithQuery bool
}

// RoundTrip implements gremlin.RoundTripper, delegating to Base and recording stats and traces for the request.
func (t *Transport) RoundTrip(ctx context.Context, req *gremlin.Request) (*gremlin.Response, error) {
	spanNameFormatter := t.FormatSpanName
	if spanNameFormatter == nil {
		spanNameFormatter = func(context.Context, *gremlin.Request) string {
			return "gremlin:traversal"
		}
	}
	startOpts := t.StartOptions
	if t.GetStartOptions != nil {
		startOpts = t.GetStartOptions(ctx, req)
	}

	var rt gremlin.RoundTripper = &traceTransport{
		base:           t.Base,
		formatSpanName: spanNameFormatter,
		startOptions:   startOpts,
		withQuery:      t.WithQuery,
	}
	rt = statsTransport{rt}
	return rt.RoundTrip(ctx, req)
}
