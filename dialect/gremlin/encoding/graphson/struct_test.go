// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"testing"

	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeStruct(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "Simple",
			in: struct {
				S string
				I int
			}{
				S: "string",
				I: 1000,
			},
			want: `{
				"S":"string",
				"I": {
					"@type": "g:Int64",
					"@value": 1000
				}
			}`,
		},
		{
			name: "Tagged",
			in: struct {
				ID   string            `json:"requestId" graphson:"g:UUID"`
				Seq  int               `json:"seq" graphson:"g:Int32"`
				Op   string            `json:"op" graphson:","`
				Args map[string]string `json:"args"`
			}{
				ID:  "cb682578-9d92-4499-9ebc-5c6aa73c5397",
				Seq: 42,
				Op:  "authentication",
				Args: map[string]string{
					"sasl": "AHN0ZXBocGhlbgBwYXNzd29yZA==",
				},
			},
			want: `{
				"requestId": {
					"@type": "g:UUID",
					"@value": "cb682578-9d92-4499-9ebc-5c6aa73c5397"
				},
				"seq": {
					"@type": "g:Int32",
					"@value": 42
				},
				"op": "authentication",
				"args": {
					"@type": "g:Map",
					"@value": ["sasl", "AHN0ZXBocGhlbgBwYXNzd29yZA=="]
				}
			}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			data, err := MarshalToString(tc.in)
			require.NoError(t, err)
			assert.JSONEq(t, tc.want, data)
		})
	}
}

func TestEncodeNestedStruct(t *testing.T) {
	type S struct {
		Parent *S  `json:"parent,omitempty"`
		ID     int `json:"id" graphson:"g:Int32"`
	}

	v := S{Parent: &S{ID: 1}, ID: 2}
	want := `{
		"id": {
			"@type": "g:Int32",
			"@value": 2
		},
		"parent": {
			"id": {
				"@type": "g:Int32",
				"@value": 1
			}
		}
	}`

	got, err := MarshalToString(&v)
	require.NoError(t, err)
	assert.JSONEq(t, want, got)
}

func TestDecodeStruct(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want any
	}{
		{
			name: "Simple",
			in: `{
				"S":"str",
				"I": {
					"@type": "g:Int32",
					"@value": 9999
				}
			}`,
			want: struct {
				S string
				I int32
			}{
				S: "str",
				I: 9999,
			},
		},
		{
			name: "Tagged",
			in: `{
				"requestId": {
					"@type": "g:UUID",
					"@value": "cb682578-9d92-4499-9ebc-5c6aa73c5397"
				},
				"seq": {
					"@type": "g:Int32",
					"@value": 42
				},
				"op": "authentication",
				"args": {
					"@type": "g:Map",
					"@value": ["sasl", "AHN0ZXBocGhlbgBwYXNzd29yZA=="]
				}
			}`,
			want: struct {
				ID   string            `json:"requestId" graphson:"g:UUID"`
				Seq  int               `json:"seq" graphson:"g:Int32"`
				Op   string            `json:"op" graphson:","`
				Args map[string]string `json:"args"`
			}{
				ID:  "cb682578-9d92-4499-9ebc-5c6aa73c5397",
				Seq: 42,
				Op:  "authentication",
				Args: map[string]string{
					"sasl": "AHN0ZXBocGhlbgBwYXNzd29yZA==",
				},
			},
		},
		{
			name: "Empty",
			in:   `{}`,
			want: struct{}{},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			typ := reflect2.TypeOf(tc.want)
			got := typ.New()
			err := UnmarshalFromString(tc.in, got)
			require.NoError(t, err)
			assert.Equal(t, tc.want, typ.Indirect(got))
		})
	}
}

func TestDecodeNestedStruct(t *testing.T) {
	type S struct {
		Parent *S  `json:"parent,omitempty"`
		ID     int `json:"id" graphson:"g:Int32"`
	}

	in := `{
		"id": {
			"@type": "g:Int32",
			"@value": 37
		},
		"parent": {
			"id": {
				"@type": "g:Int32",
				"@value": 65
			}
		}
	}`
	var got S
	err := UnmarshalFromString(in, &got)
	require.NoError(t, err)

	want := S{Parent: &S{ID: 65}, ID: 37}
	assert.Equal(t, want, got)
}
