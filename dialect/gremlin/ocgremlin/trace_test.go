// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"testing"

	"entgo.io/ent/dialect/gremlin"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"go.opencensus.io/trace"
)

type mockTransport struct {
	mock.Mock
}

func (t *mockTransport) RoundTrip(ctx context.Context, req *gremlin.Request) (*gremlin.Response, error) {
	args := t.Called(ctx, req)
	rsp, _ := args.Get(0).(*gremlin.Response)
	return rsp, args.Error(1)
}

func TestTraceTransportRoundTrip(t *testing.T) {
	_, parent := trace.StartSpan(context.Background(), "parent")
	tests := []struct {
		name   string
		parent *trace.Span
	}{
		{
			name:   "no parent",
			parent: nil,
		},
		{
			name:   "parent",
			parent: parent,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			transport := &mockTransport{}
			transport.On("RoundTrip", mock.Anything, mock.Anything).
				Run(func(args mock.Arguments) {
					span := trace.FromContext(args.Get(0).(context.Context))
					require.NotNil(t, span)
					if tt.parent != nil {
						assert.Equal(t, tt.parent.SpanContext().TraceID, span.SpanContext().TraceID)
					}
				}).
				Return(nil, errors.New("noop")).
				Once()
			defer transport.AssertExpectations(t)

			ctx, req := context.Background(), gremlin.NewEvalRequest("g.V()")
			if tt.parent != nil {
				ctx = trace.NewContext(ctx, tt.parent)
			}
			rt := &Transport{Base: transport}
			_, _ = rt.RoundTrip(ctx, req)
		})
	}
}

type testExporter struct {
	spans []*trace.SpanData
}

func (t *testExporter) ExportSpan(s *trace.SpanData) {
	t.spans = append(t.spans, s)
}

func TestEndToEnd(t *testing.T) {
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	var exporter testExporter
	trace.RegisterExporter(&exporter)
	defer trace.UnregisterExporter(&exporter)

	req := gremlin.NewEvalRequest("g.V()")
	rsp := &gremlin.Response{
		RequestID: req.RequestID,
	}
	rsp.Status.Code = 200
	rsp.Status.Message = "OK"

	var transport mockTransport
	transport.On("RoundTrip", mock.Anything, mock.Anything).
		Return(rsp, nil).
		Once()
	defer transport.AssertExpectations(t)

	rt := &Transport{Base: &transport, WithQuery: true}
	_, err := rt.RoundTrip(context.Background(), req)
	require.NoError(t, err)

	require.Len(t, exporter.spans, 1)
	attrs := exporter.spans[0].Attributes
	assert.Len(t, attrs, 5)
	assert.Equal(t, req.RequestID, attrs["gremlin.request_id"])
	assert.Equal(t, req.Operation, attrs["gremlin.operation"])
	assert.Equal(t, req.Arguments[gremlin.ArgsGremlin], attrs["gremlin.query"])
	assert.Equal(t, int64(200), attrs["gremlin.code"])
	assert.Equal(t, "OK", attrs["gremlin.message"])
}

func TestRequestAttributes(t *testing.T) {
	tests := []struct {
		name      string
		makeReq   func() *gremlin.Request
		wantAttrs []trace.Attribute
	}{
		{
			name: "Query without bindings",
			makeReq: func() *gremlin.Request {
				req := gremlin.NewEvalRequest("g.E().count()")
				req.RequestID = "a8b5c664-03ca-4175-a9e7-569b46f3551c"
				return req
			},
			wantAttrs: []trace.Attribute{
				trace.StringAttribute("gremlin.request_id", "a8b5c664-03ca-4175-a9e7-569b46f3551c"),
				trace.StringAttribute("gremlin.operation", "eval"),
				trace.StringAttribute("gremlin.query", "g.E().count()"),
			},
		},
		{
			name: "Query with bindings",
			makeReq: func() *gremlin.Request {
				bindings := map[string]any{
					"$1": "user", "$2": int64(42),
					"$3": 3.14, "$4": bytes.Repeat([]byte{0xff}, 257),
					"$5": true, "$6": nil,
				}
				req := gremlin.NewEvalRequest(
					`g.V().hasLabel($1).has("age",$2).has("v",$3).limit($4).valueMap($5)`,
					gremlin.WithBindings(bindings),
				)
				req.RequestID = "d3d986fa-bd22-41bd-b2f7-ef2f1f639260"
				return req
			},
			wantAttrs: []trace.Attribute{
				trace.StringAttribute("gremlin.request_id", "d3d986fa-bd22-41bd-b2f7-ef2f1f639260"),
				trace.StringAttribute("gremlin.operation", "eval"),
				trace.StringAttribute("gremlin.query", `g.V().hasLabel($1).has("age",$2).has("v",$3).limit($4).valueMap($5)`),
				trace.StringAttribute("gremlin.binding.$1", "user"),
				trace.Int64Attribute("gremlin.binding.$2", 42),
				trace.Float64Attribute("gremlin.binding.$3", 3.14),
				trace.StringAttribute("gremlin.binding.$4", func() string {
					str := fmt.Sprintf("%v", bytes.Repeat([]byte{0xff}, 256))
					return str[:256]
				}()),
				trace.BoolAttribute("gremlin.binding.$5", true),
				trace.StringAttribute("gremlin.binding.$6", ""),
			},
		},
		{
			name: "Authentication",
			makeReq: func() *gremlin.Request {
				return gremlin.NewAuthRequest(
					"d239d950-59a1-41a7-a103-908f976ebd89",
					"user", "pass",
				)
			},
			wantAttrs: []trace.Attribute{
				trace.StringAttribute("gremlin.request_id", "d239d950-59a1-41a7-a103-908f976ebd89"),
				trace.StringAttribute("gremlin.operation", "authentication"),
				trace.StringAttribute("gremlin.query", ""),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			req := tt.makeReq()
			attrs := requestAttrs(req, true)
			for _, attr := range attrs {
				assert.Contains(t, tt.wantAttrs, attr)
			}
			assert.Len(t, attrs, len(tt.wantAttrs))
		})
	}
}

func TestResponseAttributes(t *testing.T) {
	tests := []struct {
		name      string
		makeRsp   func() *gremlin.Response
		wantAttrs []trace.Attribute
	}{
		{
			name: "Success no message",
			makeRsp: func() *gremlin.Response {
				var rsp gremlin.Response
				rsp.Status.Code = 204
				return &rsp
			},
			wantAttrs: []trace.Attribute{
				trace.Int64Attribute("gremlin.code", 204),
			},
		},
		{
			name: "Authenticate with message",
			makeRsp: func() *gremlin.Response {
				var rsp gremlin.Response
				rsp.Status.Code = 407
				rsp.Status.Message = "login required"
				return &rsp
			},
			wantAttrs: []trace.Attribute{
				trace.Int64Attribute("gremlin.code", 407),
				trace.StringAttribute("gremlin.message", "login required"),
			},
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			rsp := tt.makeRsp()
			attrs := responseAttrs(rsp)
			assert.Equal(t, tt.wantAttrs, attrs)
		})
	}
}

func TestTraceStatus(t *testing.T) {
	tests := []struct {
		in   int
		want trace.Status
	}{
		{200, trace.Status{Code: trace.StatusCodeOK, Message: "Success"}},
		{204, trace.Status{Code: trace.StatusCodeOK, Message: "No Content"}},
		{206, trace.Status{Code: trace.StatusCodeOK, Message: "Partial Content"}},
		{401, trace.Status{Code: trace.StatusCodePermissionDenied, Message: "Unauthorized"}},
		{407, trace.Status{Code: trace.StatusCodeUnauthenticated, Message: "Authenticate"}},
		{498, trace.Status{Code: trace.StatusCodeInvalidArgument, Message: "Malformed Request"}},
		{499, trace.Status{Code: trace.StatusCodeInvalidArgument, Message: "Invalid Request Arguments"}},
		{500, trace.Status{Code: trace.StatusCodeInternal, Message: "Server Error"}},
		{597, trace.Status{Code: trace.StatusCodeInvalidArgument, Message: "Script Evaluation Error"}},
		{598, trace.Status{Code: trace.StatusCodeDeadlineExceeded, Message: "Server Timeout"}},
		{599, trace.Status{Code: trace.StatusCodeInternal, Message: "Server Serialization Error"}},
		{600, trace.Status{Code: trace.StatusCodeUnknown, Message: ""}},
	}
	for _, tt := range tests {
		assert.Equal(t, tt.want, TraceStatus(tt.in))
	}
}
