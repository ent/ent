// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeArray(t *testing.T) {
	t.Parallel()
	got, err := MarshalToString([...]string{"a", "b", "c"})
	require.NoError(t, err)
	want := `{ "@type": "g:List", "@value": ["a", "b", "c"]}`
	assert.JSONEq(t, want, got)
}

func TestEncodeSlice(t *testing.T) {
	tests := []struct {
		in   any
		want string
	}{
		{
			in: []int32{5, 6, 7, 8},
			want: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int32",
						"@value": 5
					},
					{
						"@type": "g:Int32",
						"@value": 6
					},
					{
						"@type": "g:Int32",
						"@value": 7
					},
					{
						"@type": "g:Int32",
						"@value": 8
					}
				]
			}`,
		},
		{
			in: []byte{1, 2, 3, 4, 5},
			want: `{
				"@type": "gx:ByteBuffer",
				"@value": "AQIDBAU="
			}`,
		},
		{
			in: [...]byte{4, 5},
			want: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "gx:Byte",
						"@value": 4
					},
					{
						"@type": "gx:Byte",
						"@value": 5
					}
				]
			}`,
		},
		{
			in:   []uint64(nil),
			want: "null",
		},
	}
	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.in), func(t *testing.T) {
			t.Parallel()
			var got bytes.Buffer
			err := NewEncoder(&got).Encode(tc.in)
			assert.NoError(t, err)
			assert.JSONEq(t, tc.want, got.String())
		})
	}
}

func TestDecodeSlice(t *testing.T) {
	tests := []struct {
		in   string
		want any
	}{
		{
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int32",
						"@value": 3
					},
					{
						"@type": "g:Int32",
						"@value": -2
					},
					{
						"@type": "g:Int32",
						"@value": 1
					}
				]
			}`,
			want: []int32{3, -2, 1},
		},
		{
			in: `{
				"@type": "g:List",
				"@value": ["a", "b", "c"]
			}`,
			want: []string{"a", "b", "c"},
		},
		{
			in: `{
				"@type": "gx:ByteBuffer",
				"@value": "AQIDBAU="
			}`,
			want: []byte{1, 2, 3, 4, 5},
		},
		{
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "gx:Byte",
						"@value": 42
					},
					{
						"@type": "gx:Byte",
						"@value": 55
					},
					{
						"@type": "gx:Byte",
						"@value": 94
					}
				]
			}`,
			want: [...]byte{42, 55},
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.want), func(t *testing.T) {
			t.Parallel()
			typ := reflect2.TypeOf(tc.want)
			got := typ.New()
			err := NewDecoder(strings.NewReader(tc.in)).Decode(got)
			require.NoError(t, err)
			assert.Equal(t, tc.want, typ.Indirect(got))
		})
	}
}

func TestDecodeBadSlice(t *testing.T) {
	tests := []struct {
		name string
		in   string
		new  func() any
	}{
		{
			name: "TypeMismatch",
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int32",
						"@value": 3
					},
					{
						"@type": "g:Int64",
						"@value": 2
					}
				]
			}`,
			new: func() any { return &[]int{} },
		},
		{
			name: "BadValue",
			in: `{
				"@type": "g:List",
				"@value": [
					{
						"@type": "g:Int32",
						"@value": 3
					},
					{
						"@type": "g:Int32",
						"@value": "2"
					}
				]
			}`,
			new: func() any { return &[2]int{} },
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			err := NewDecoder(strings.NewReader(tc.in)).Decode(tc.new())
			assert.Error(t, err)
		})
	}
}
