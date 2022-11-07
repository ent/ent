// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package graphson

import (
	"fmt"
	"math"
	"testing"

	jsoniter "github.com/json-iterator/go"
	"github.com/modern-go/reflect2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEncodeNative(t *testing.T) {
	tests := []struct {
		in      any
		want    string
		wantErr bool
	}{
		{
			in:   true,
			want: "true",
		},
		{
			in:   "hello",
			want: `"hello"`,
		},
		{
			in: int8(120),
			want: `{
				"@type": "g:Int32",
				"@value": 120
			}`,
		},
		{
			in: int16(-16),
			want: `{
				"@type": "gx:Int16",
				"@value": -16
			}`,
		},
		{
			in: int32(3232),
			want: `{
				"@type": "g:Int32",
				"@value": 3232
			}`,
		},
		{
			in: int64(646464),
			want: `{
				"@type": "g:Int64",
				"@value": 646464
			}`,
		},
		{
			in: int(127001),
			want: `{
				"@type": "g:Int64",
				"@value": 127001
			}`,
		},
		{
			in: uint8(81),
			want: `{
				"@type": "gx:Byte",
				"@value": 81
			}`,
		},
		{
			in: uint16(12345),
			want: `{
				"@type": "g:Int32",
				"@value": 12345
			}`,
		},
		{
			in: uint32(123454321),
			want: `{
				"@type": "g:Int64",
				"@value": 123454321
			}`,
		},
		{
			in: uint64(1234567890),
			want: `{
				"@type": "gx:BigInteger",
				"@value": 1234567890
			}`,
		},
		{
			in: uint(9876543210),
			want: `{
				"@type" :"gx:BigInteger",
				"@value": 9876543210
			}`,
		},
		{
			in: float32(math.Pi),
			want: `{
				"@type": "g:Float",
				"@value": 3.1415927
			}`,
		},
		{
			in: float64(math.E),
			want: `{
				"@type": "g:Double",
				"@value": 2.718281828459045
			}`,
		},
		{
			in: math.NaN(),
			want: `{
				"@type": "g:Double",
				"@value": "NaN"
			}`,
		},
		{
			in: math.Inf(1),
			want: `{
				"@type": "g:Double",
				"@value": "Infinity"
			}`,
		},
		{
			in: math.Inf(-1),
			want: `{
				"@type": "g:Double",
				"@value": "-Infinity"
			}`,
		},
		{
			in: func() *int { v := 7142; return &v }(),
			want: `{
				"@type": "g:Int64",
				"@value": 7142
			}`,
		},
		{
			in: func() any { v := int16(6116); return &v }(),
			want: `{
				"@type": "gx:Int16",
				"@value": 6116
			}`,
		},
		{
			in:   nil,
			want: "null",
		},
		{
			in:      make(chan int),
			wantErr: true,
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.in), func(t *testing.T) {
			t.Parallel()
			got, err := MarshalToString(tc.in)
			if !tc.wantErr {
				assert.NoError(t, err)
				assert.JSONEq(t, tc.want, got)
			} else {
				assert.Error(t, err)
				assert.Empty(t, got)
			}
		})
	}
}

func TestDecodeNative(t *testing.T) {
	tests := []struct {
		in   string
		want any
	}{
		{
			in:   `{"@type": "g:Float", "@value": 3.14}`,
			want: float32(3.14),
		},
		{
			in: `{"@type": "g:Float", "@value": "Float"}`,
		},
		{
			in:   `{"@type": "g:Double", "@value": 2.71}`,
			want: float64(2.71),
		},
		{
			in:   `{"@type": "gx:BigDecimal", "@value": 3.142}`,
			want: float32(3.142),
		},
		{
			in:   `{"@type": "gx:BigDecimal", "@value": 55512.5176}`,
			want: float64(55512.5176),
		},
		{
			in:   `{"@type": "g:T", "@value": "world"}`,
			want: "world",
		},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(fmt.Sprintf("%T", tc.want), func(t *testing.T) {
			t.Parallel()
			if tc.want != nil {
				typ := reflect2.TypeOf(tc.want)
				got := typ.New()
				err := UnmarshalFromString(tc.in, got)
				require.NoError(t, err)
				assert.Equal(t, tc.want, typ.Indirect(got))
			} else {
				var msg jsoniter.RawMessage
				err := UnmarshalFromString(tc.in, &msg)
				assert.Error(t, err)
			}
		})
	}
}

func TestDecodeTypeMismatch(t *testing.T) {
	t.Run("FloatToInt", func(t *testing.T) {
		var v int
		err := UnmarshalFromString(`{"@type": "g:Float", "@value": 3.14}`, &v)
		assert.Error(t, err)
	})
	t.Run("DoubleToFloat", func(t *testing.T) {
		var v float32
		err := UnmarshalFromString(`{"@type": "g:Double", "@value": 5.51}`, &v)
		assert.Error(t, err)
	})
	t.Run("BigDecimalToUint64", func(t *testing.T) {
		var v uint64
		err := UnmarshalFromString(`{"@type": "gx:BigDecimal", "@value": 5645.51834}`, &v)
		assert.Error(t, err)
	})
}

func TestDecodeNaNInfinity(t *testing.T) {
	tests := []struct {
		data   []byte
		expect func(*testing.T, float64, error)
	}{
		{
			data: []byte(`{"@type": "g:Double", "@value": "NaN"}`),
			expect: func(t *testing.T, f float64, err error) {
				assert.NoError(t, err)
				assert.True(t, math.IsNaN(f))
			},
		},
		{
			data: []byte(`{"@type": "g:Double", "@value": "Infinity"}`),
			expect: func(t *testing.T, f float64, err error) {
				assert.NoError(t, err)
				assert.True(t, math.IsInf(f, 1))
			},
		},
		{
			data: []byte(`{"@type": "g:Double", "@value": "-Infinity"}`),
			expect: func(t *testing.T, f float64, err error) {
				assert.NoError(t, err)
				assert.True(t, math.IsInf(f, -1))
			},
		},
		{
			data: []byte(`{"@type": "g:Double", "@value": "Junk"}`),
			expect: func(t *testing.T, _ float64, err error) {
				assert.Error(t, err)
			},
		},
		{
			data: []byte(`{"@type": "g:Double", "@value": [42]}`),
			expect: func(t *testing.T, _ float64, err error) {
				assert.Error(t, err)
			},
		},
	}

	for _, tc := range tests {
		var f float64
		err := Unmarshal(tc.data, &f)
		tc.expect(t, f, err)
	}
}

func TestDecodeTypeDefinition(t *testing.T) {
	type Status int
	const StatusOk Status = 42

	var status Status
	err := UnmarshalFromString(`{"@type": "g:Int64", "@value": 42}`, &status)
	assert.NoError(t, err)
	assert.Equal(t, StatusOk, status)
}
