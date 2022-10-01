// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExpandBindings(t *testing.T) {
	tests := []struct {
		req       *Request
		wantErr   bool
		wantQuery string
	}{
		{
			req:       NewEvalRequest("no bindings"),
			wantQuery: "no bindings",
		},
		{
			req:       NewEvalRequest("g.V($0)", WithBindings(map[string]any{"$0": 1})),
			wantQuery: "g.V(1)",
		},
		{
			req:       NewEvalRequest("g.V().has($1, $2)", WithBindings(map[string]any{"$1": "name", "$2": "a8m"})),
			wantQuery: "g.V().has(\"name\", \"a8m\")",
		},
		{
			req:       NewEvalRequest("g.V().limit(n)", WithBindings(map[string]any{"n": 10})),
			wantQuery: "g.V().limit(10)",
		},
		{
			req:     NewEvalRequest("g.V()", WithBindings(map[string]any{"$0": func() {}})),
			wantErr: true,
		},
		{
			req:       NewEvalRequest("g.V().has($0, $1)", WithBindings(map[string]any{"$0": "active", "$1": true})),
			wantQuery: "g.V().has(\"active\", true)",
		},
		{
			req:       NewEvalRequest("g.V().has($1, $11)", WithBindings(map[string]any{"$1": "active", "$11": true})),
			wantQuery: "g.V().has(\"active\", true)",
		},
	}
	for i, tt := range tests {
		tt := tt
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			rt := ExpandBindings(RoundTripperFunc(func(ctx context.Context, r *Request) (*Response, error) {
				assert.Equal(t, tt.wantQuery, r.Arguments[ArgsGremlin])
				return nil, nil
			}))
			_, err := rt.RoundTrip(context.Background(), tt.req)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}

func TestExpandBindingsNoQuery(t *testing.T) {
	rt := ExpandBindings(RoundTripperFunc(func(ctx context.Context, r *Request) (*Response, error) {
		return nil, nil
	}))
	_, err := rt.RoundTrip(context.Background(), &Request{Arguments: map[string]any{
		ArgsBindings: map[string]any{},
	}})
	assert.NoError(t, err)
}
