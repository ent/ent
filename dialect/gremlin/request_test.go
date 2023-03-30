// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"encoding/json"
	"testing"
	"time"

	"entgo.io/ent/dialect/gremlin/encoding/graphson"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluateRequestEncode(t *testing.T) {
	req := NewEvalRequest("g.V(x)",
		WithBindings(map[string]any{"x": 1}),
		WithEvalTimeout(time.Second),
	)
	data, err := graphson.Marshal(req)
	require.NoError(t, err)

	var got map[string]any
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)

	assert.Equal(t, map[string]any{
		"@type":  "g:UUID",
		"@value": req.RequestID,
	}, got["requestId"])
	assert.Equal(t, req.Operation, got["op"])
	assert.Equal(t, req.Processor, got["processor"])

	args := got["args"].(map[string]any)
	assert.Equal(t, "g:Map", args["@type"])
	assert.ElementsMatch(t, args["@value"], []any{
		"gremlin", "g.V(x)", "language", "gremlin-groovy",
		"scriptEvaluationTimeout", map[string]any{
			"@type":  "g:Int64",
			"@value": float64(1000),
		},
		"bindings", map[string]any{
			"@type": "g:Map",
			"@value": []any{
				"x",
				map[string]any{
					"@type":  "g:Int64",
					"@value": float64(1),
				},
			},
		},
	})
}

func TestEvaluateRequestWithoutBindingsEncode(t *testing.T) {
	req := NewEvalRequest("g.E()")
	got, err := graphson.MarshalToString(req)
	require.NoError(t, err)
	assert.NotContains(t, got, "bindings")
}

func TestAuthenticateRequestEncode(t *testing.T) {
	req := NewAuthRequest("41d2e28a-20a4-4ab0-b379-d810dede3786", "user", "pass")
	data, err := graphson.Marshal(req)
	require.NoError(t, err)

	var got map[string]any
	err = json.Unmarshal(data, &got)
	require.NoError(t, err)

	assert.Equal(t, map[string]any{
		"@type":  "g:UUID",
		"@value": req.RequestID,
	}, got["requestId"])
	assert.Equal(t, req.Operation, got["op"])
	assert.Equal(t, req.Processor, got["processor"])

	args := got["args"].(map[string]any)
	assert.Equal(t, "g:Map", args["@type"])
	assert.ElementsMatch(t, args["@value"], []any{
		"sasl", "AHVzZXIAcGFzcw==", "saslMechanism", "PLAIN",
	})
}

func TestCredentialsMarshaling(t *testing.T) {
	want := Credentials{
		Username: "username",
		Password: "password",
	}

	text, err := want.MarshalText()
	assert.NoError(t, err)
	assert.Equal(t, "AHVzZXJuYW1lAHBhc3N3b3Jk", string(text))

	var got Credentials
	err = got.UnmarshalText(text)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestCredentialsBadEncodingMarshaling(t *testing.T) {
	tests := []struct {
		name string
		text []byte
	}{
		{
			name: "BadBase64",
			text: []byte{0x12},
		},
		{
			name: "Empty",
			text: []byte{},
		},
		{
			name: "BadPrefix",
			text: []byte("Kg=="),
		},
		{
			name: "NoSeparator",
			text: []byte("AHVzZXI="),
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			var creds Credentials
			err := creds.UnmarshalText(tc.text)
			assert.Error(t, err)
		})
	}
}
