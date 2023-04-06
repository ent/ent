// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package ocgremlin

import (
	"context"
	"errors"
	"testing"

	"entgo.io/ent/dialect/gremlin"

	"github.com/stretchr/testify/mock"
	"go.opencensus.io/trace"
)

type mockExporter struct {
	mock.Mock
}

func (e *mockExporter) ExportSpan(s *trace.SpanData) {
	e.Called(s)
}

func TestTransportOptions(t *testing.T) {
	tests := []struct {
		name     string
		spanName string
		wantName string
	}{
		{
			name:     "Default formatter",
			wantName: "gremlin:traversal",
		},
		{
			name:     "Custom formatter",
			spanName: "tester",
			wantName: "tester",
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var exporter mockExporter
			exporter.On(
				"ExportSpan",
				mock.MatchedBy(func(s *trace.SpanData) bool { return s.Name == tt.wantName })).
				Once()
			defer exporter.AssertExpectations(t)
			trace.RegisterExporter(&exporter)
			defer trace.UnregisterExporter(&exporter)

			transport := &mockTransport{}
			transport.On("RoundTrip", mock.Anything, mock.Anything).
				Return(nil, errors.New("noop")).
				Once()
			defer transport.AssertExpectations(t)

			rt := &Transport{
				Base: transport,
				GetStartOptions: func(context.Context, *gremlin.Request) trace.StartOptions {
					return trace.StartOptions{Sampler: trace.AlwaysSample()}
				},
			}
			if tt.spanName != "" {
				rt.FormatSpanName = func(context.Context, *gremlin.Request) string {
					return tt.spanName
				}
			}
			_, _ = rt.RoundTrip(context.Background(), gremlin.NewEvalRequest("g.E()"))
		})
	}
}
