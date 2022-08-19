// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestDecodeInterface(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    any
		wantErr bool
	}{
		{
			name: "Boolean",
			in:   "false",
			want: false,
		},
		{
			name: "String",
			in:   `"str"`,
			want: "str",
		},
		{
			name: "Double",
			in: `{
				"@type": "g:Double",
				"@value": 3.14
			}`,
			want: float64(3.14),
		},
		{
			name: "Float",
			in: `{
				"@type": "g:Float",
				"@value": -22.567
			}`,
			want: float32(-22.567),
		},
		{
			name: "Int32",
			in: `{
				"@type": "g:Int32",
				"@value": 9000
			}`,
			want: int32(9000),
		},
		{
			name: "Int64",
			in: `{
				"@type": "g:Int64",
				"@value": 188786
			}`,
			want: int64(188786),
		},
		{
			name: "BigInteger",
			in: `{
				"@type": "gx:BigInteger",
				"@value": 352353463712
			}`,
			want: int64(352353463712),
		},
		{
			name: "Byte",
			in: `{
				"@type": "gx:Byte",
				"@value": 100
			}`,
			want: uint8(100),
		},
		{
			name: "Int16",
			in: `{
				"@type": "gx:Int16",
				"@value": 2000
			}`,
			want: int16(2000),
		},
		{
			name: "UnknownType",
			in: `{
				"@type": "g:T",
				"@value": "label"
			}`,
			want: "label",
		},
		{
			name:    "UntypedArray",
			in:      "[]",
			wantErr: true,
		},
		{
			name: "NoType",
			in: `{
				"@typ": "g:Int32",
				"@value": 345
			}`,
			wantErr: true,
		},
		{
			name: "BadObject",
			in: `{
				"@type": "g:Int32",
				"@value": 345
			`,
			wantErr: true,
		},
		{
			name: "BadList",
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int64",
						"@val": 123457990
					}
				]
			}`,
			wantErr: true,
		},
		{
			name: "BadMap",
			in: `{
				"@type": "g:Map",
				"@value": [
					{
						"@type": "g:Int64",
						"@val": 123457990
					},
					"First"
				]
			}`,
			wantErr: true,
		},
		{
			name: "KeyOnlyMap",
			in: `{
				"@type": "g:Map",
				"@value": ["Key"]
			}`,
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var got any
			err := UnmarshalFromString(tc.in, &got)
			if !tc.wantErr {
				require.NoError(t, err)
				assert.Equal(t, tc.want, got)
			} else {
				assert.Error(t, err)
			}
		})
	}
}

func TestDecodeInterfaceSlice(t *testing.T) {
	tests := []struct {
		in   string
		want any
	}{
		{
			in: `{
				"@type": "g:List",
				"@value": []
			}`,
			want: []any{},
		},
		{
			in: `{
				"@type": "g:List",
				"@value": ["x", "y", "z"]
			}`,
			want: []string{"x", "y", "z"},
		},
		{
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int64",
						"@value": 123457990
					},
					{
						"@type": "g:Int64",
						"@value": 23456111
					},
					{
						"@type": "g:Int64",
						"@value": -687450
					}
				]
			}`,
			want: []int64{123457990, 23456111, -687450},
		},
		{
			in: `{
				"@type": "gx:ByteBuffer",
				"@value": "AQIDBAU="
			}`,
			want: []byte{1, 2, 3, 4, 5},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.want), func(t *testing.T) {
			t.Parallel()
			var got any
			err := UnmarshalFromString(tc.in, &got)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDecodeInterfaceMap(t *testing.T) {
	tests := []struct {
		in   string
		want any
	}{
		{
			in: `{
				"@type": "g:Map",
				"@value": []
			}`,
			want: map[any]any{},
		},
		{
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
						"@value": 10
					},
					"Nov",
					{
						"@type": "g:Int32",
						"@value": 11
					}
				]
			}`,
			want: map[string]int32{
				"Sep": int32(9),
				"Oct": int32(10),
				"Nov": int32(11),
			},
		},
		{
			in: `{
				"@type": "g:Map",
				"@value": [
					"One",
					{
						"@type": "g:List",
						"@value": [
							{
								"@type": "g:Int32",
								"@value": 1
							}
						]
					},
					"Two",
					{
						"@type": "g:List",
						"@value": [
							{
								"@type": "g:Int32",
								"@value": 2
							}
						]
					},
					"Three",
					{
						"@type": "g:List",
						"@value": [
							{
								"@type": "g:Int32",
								"@value": 3
							}
						]
					}
				]
			}`,
			want: map[string][]int32{
				"One":   {1},
				"Two":   {2},
				"Three": {3},
			},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.want), func(t *testing.T) {
			t.Parallel()
			var got any
			err := UnmarshalFromString(tc.in, &got)
			require.NoError(t, err)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestDecodeInterfaceObject(t *testing.T) {
	book := struct {
		ID       string   `json:"id" graphson:"g:UUID"`
		Title    string   `json:"title"`
		Author   string   `json:"author"`
		Pages    int      `json:"num_pages"`
		Chapters []string `json:"chapters"`
	}{
		ID:       "21d5dcbf-1fd4-493e-9b74-d6c429f9e4a5",
		Title:    "The Art of Computer Programming, Vol. 2",
		Author:   "Donald E. Knuth",
		Pages:    784,
		Chapters: []string{"Random numbers", "Arithmetic"},
	}
	data, err := Marshal(book)
	require.NoError(t, err)

	var v any
	err = Unmarshal(data, &v)
	require.NoError(t, err)

	obj := v.(map[string]any)
	assert.Equal(t, book.ID, obj["id"])
	assert.Equal(t, book.Title, obj["title"])
	assert.Equal(t, book.Author, obj["author"])
	assert.EqualValues(t, book.Pages, obj["num_pages"])
	assert.ElementsMatch(t, book.Chapters, obj["chapters"])
}
