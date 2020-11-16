// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"math"
	"testing"

	"github.com/facebook/ent/schema/field"

	"github.com/stretchr/testify/require"
)

func TestColumn_ConvertibleTo(t *testing.T) {
	c1 := &Column{Type: field.TypeString, Size: 10}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 10}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 255}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 9}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))

	c1 = &Column{Type: field.TypeFloat32}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))

	c1 = &Column{Type: field.TypeFloat64}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeFloat64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))

	c1 = &Column{Type: field.TypeUint}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt64}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeUint64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint16}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 1}))

	c1 = &Column{Type: field.TypeInt8}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 3}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 4}))

	c1 = &Column{Type: field.TypeUint8}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 2}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 3}))

	c1 = &Column{Type: field.TypeInt16}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 5}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 6}))

	c1 = &Column{Type: field.TypeUint16}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 4}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 5}))

	c1 = &Column{Type: field.TypeInt32}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 10}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 11}))

	c1 = &Column{Type: field.TypeUint32}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 9}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 10}))

	c1 = &Column{Type: field.TypeInt64}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 19}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 20}))

	c1 = &Column{Type: field.TypeUint64}
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 18}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 19}))

	c1 = &Column{Type: field.TypeInt}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint16}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint32}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 1}))
}

func TestColumn_ScanDefault(t *testing.T) {
	c1 := &Column{Type: field.TypeString, Size: 10}
	require.NoError(t, c1.ScanDefault("Hello World"))
	require.Equal(t, "Hello World", c1.Default)
	require.NoError(t, c1.ScanDefault("1"))
	require.Equal(t, "1", c1.Default)

	c1 = &Column{Type: field.TypeInt64}
	require.NoError(t, c1.ScanDefault("128"))
	require.Equal(t, int64(128), c1.Default)
	require.NoError(t, c1.ScanDefault("1"))
	require.Equal(t, int64(1), c1.Default)
	require.Error(t, c1.ScanDefault("foo"))

	c1 = &Column{Type: field.TypeUint64}
	require.NoError(t, c1.ScanDefault("128"))
	require.Equal(t, uint64(128), c1.Default)
	require.NoError(t, c1.ScanDefault("1"))
	require.Equal(t, uint64(1), c1.Default)
	require.Error(t, c1.ScanDefault("foo"))

	c1 = &Column{Type: field.TypeFloat64}
	require.NoError(t, c1.ScanDefault("128.1"))
	require.Equal(t, 128.1, c1.Default)
	require.NoError(t, c1.ScanDefault("1"))
	require.Equal(t, float64(1), c1.Default)
	require.Error(t, c1.ScanDefault("foo"))

	c1 = &Column{Type: field.TypeBool}
	require.NoError(t, c1.ScanDefault("1"))
	require.Equal(t, true, c1.Default)
	require.NoError(t, c1.ScanDefault("true"))
	require.Equal(t, true, c1.Default)
	require.NoError(t, c1.ScanDefault("0"))
	require.Equal(t, false, c1.Default)
	require.NoError(t, c1.ScanDefault("false"))
	require.Equal(t, false, c1.Default)
	require.Error(t, c1.ScanDefault("foo"))
}

func Test_digitsRequired(t *testing.T) {
	require.Equal(t, int64(3), digitsRequired(math.MaxInt8))
	require.Equal(t, int64(4), digitsRequired(math.MinInt8))
	require.Equal(t, int64(3), digitsRequired(math.MaxUint8))
	require.Equal(t, int64(5), digitsRequired(math.MaxInt16))
	require.Equal(t, int64(6), digitsRequired(math.MinInt16))
}
