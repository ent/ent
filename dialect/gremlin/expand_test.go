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
			req:       NewEvalRequest("g.V($0)", WithBindings(map[string]interface{}{"$0": 1})),
			wantQuery: "g.V(1)",
		},
		{
			req:       NewEvalRequest("g.V().has($1, $2)", WithBindings(map[string]interface{}{"$1": "name", "$2": "a8m"})),
			wantQuery: "g.V().has(\"name\", \"a8m\")",
		},
		{
			req:       NewEvalRequest("g.V().limit(n)", WithBindings(map[string]interface{}{"n": 10})),
			wantQuery: "g.V().limit(10)",
		},
		{
			req:     NewEvalRequest("g.V()", WithBindings(map[string]interface{}{"$0": func() {}})),
			wantErr: true,
		},
		{
			req:       NewEvalRequest("g.V().has($0, $1)", WithBindings(map[string]interface{}{"$0": "active", "$1": true})),
			wantQuery: "g.V().has(\"active\", true)",
		},
		{
			req:       NewEvalRequest("g.V().has($1, $11)", WithBindings(map[string]interface{}{"$1": "active", "$11": true})),
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
	_, err := rt.RoundTrip(context.Background(), &Request{Arguments: map[string]interface{}{
		ArgsBindings: map[string]interface{}{},
	}})
	assert.NoError(t, err)
}
