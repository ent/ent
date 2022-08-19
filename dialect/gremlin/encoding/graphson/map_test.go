// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"strings"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeMap(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{
			name: "simple",
			in: map[int32]string{
				3: "Mar",
				1: "Jan",
				2: "Feb",
			},
			want: `[
				{
					"@type": "g:Int32",
					"@value": 1
				},
				"Jan",
				{
					"@type": "g:Int32",
					"@value": 2
				},
				"Feb",
				{
					"@type": "g:Int32",
					"@value": 3
				},
				"Mar"
			]`,
		},
		{
			name: "mixed",
			in: map[string]any{
				"byte":   byte('a'),
				"string": "str",
				"slice":  []int{1, 2, 3},
				"map":    map[string]int{},
			},
			want: `[
				"byte",
				{
					"@type": "gx:Byte",
					"@value": 97
				},
				"string",
				"str",
				"slice",
				{
					"@type": "g:List",
					"@value": [
						{
							"@type": "g:Int64",
							"@value": 1
						},
						{
							"@type": "g:Int64",
							"@value": 2
						},
						{
							"@type": "g:Int64",
							"@value": 3
						}
					]
				},
				"map",
				{
					"@type": "g:Map",
					"@value": []
				}
			]`,
		},
		{
			name: "struct-key",
			in: map[struct {
				K string `json:"key"`
			}]int32{

				{"result"}: 42,
			},
			want: `[
				{
					"key": "result"
				},
				{
					"@type": "g:Int32",
					"@value": 42
				}
			]`,
		},
		{
			name: "nil",
			in:   map[string]uint8(nil),
			want: "null",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			data, err := Marshal(tc.in)
			require.NoError(t, err)

			assert.Equal(t, "g:Map", jsoniter.Get(data, "@type").ToString())
			var want []any
			err = jsoniter.UnmarshalFromString(tc.want, &want)
			require.NoError(t, err)

			got, ok := jsoniter.Get(data, "@value").GetInterface().([]any)
			require.True(t, ok)
			assert.ElementsMatch(t, want, got)
		})
	}
}

func TestDecodeMap(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want any
	}{
		{
			name: "empty",
			in: `{
				"@type": "g:Map",
				"@value": []
			}`,
			want: map[int]int{},
		},
		{
			name: "simple",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int32",
						"@value": 6
					},
					"Jun",
					{
						"@type": "g:Int32",
						"@value": 7
					},
					"Jul",
					{
						"@type": "g:Int32",
						"@value": 8
					},
					"Aug"
				]
			}`,
			want: map[int]string{
				6: "Jun",
				7: "Jul",
				8: "Aug",
			},
		},
		{
			name: "duplicate",
			in: `{
				"@type": "g:Map",
				"@value": [
					"Sep",
					{
						"@type": "g:Int32",
						"@value": 9
					},
					"Oct",
					{
						"@type": "g:Int32",
						"@value": 65
					},
					"Oct",
					{
						"@type": "g:Int32",
						"@value": 10
					},
					"Nov",
					null
				]
			}`,
			want: map[string]*int{
				"Sep": func() *int { v := 9; return &v }(),
				"Oct": func() *int { v := 10; return &v }(),
				"Nov": nil,
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			typ := reflect2.TypeOf(tc.want).(reflect2.MapType)
			got := typ.MakeMap(0)
			err := UnmarshalFromString(tc.in, got)
			require.NoError(t, err)
			assert.Equal(t, tc.want, typ.Indirect(got))
		})
	}
}

func TestDecodeMapIntoNil(t *testing.T) {
	var got map[int64]int32
	err := UnmarshalFromString(`{
		"@type": "g:Map",
		"@value": [
			{
				"@type": "g:Int64",
				"@value": 9
			},
			{
				"@type": "g:Int32",
				"@value": -9
			},
			{
				"@type": "g:Int64",
				"@value": 99
			},
			{
				"@type": "g:Int32",
				"@value": -99
			},
			{
				"@type": "g:Int64",
				"@value": 999
			},
			{
				"@type": "g:Int32",
				"@value": -999
			}
		]
	}`, &got)
	require.NoError(t, err)
	assert.Equal(t, map[int64]int32{9: -9, 99: -99, 999: -999}, got)
}

func TestDecodeBadMap(t *testing.T) {
	tests := []struct {
		name string
		in   string
	}{
		{
			name: "BadValue",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 9
					},
					{
						"@type": "g:Int32",
						"@value": "55"
					}
				]
			}`,
		},
		{
			name: "NoValue",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 9
					},
					{
						"@type": "g:Int32",
						"@value": 9
					},
					{
						"@type": "g:Int64",
						"@value": 42
					}
				]
			}`,
		},
		{
			name: "AlterKeyType",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 9
					},
					{
						"@type": "g:Int32",
						"@value": 9
					},
					{
						"@type": "g:Int32",
						"@value": 42
					},
					{
						"@type": "g:Int32",
						"@value": 42
					}
				]
			}`,
		},
		{
			name: "AlterValType",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 9
					},
					{
						"@type": "g:Int32",
						"@value": 9
					},
					{
						"@type": "g:Int64",
						"@value": 42
					},
					{
						"@type": "g:Int64",
						"@value": 42
					}
				]
			}`,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var v map[int]int
			err := NewDecoder(strings.NewReader(tc.in)).Decode(&v)
			assert.Error(t, err)
		})
	}
}
