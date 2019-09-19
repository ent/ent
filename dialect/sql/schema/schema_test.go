// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"testing"

	"github.com/facebookincubator/ent/schema/field"

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
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))

	c1 = &Column{Type: field.TypeInt}
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt}))
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeInt64}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeInt32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint8}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint16}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeUint32}))
	require.False(t, c1.ConvertibleTo(&Column{Type: field.TypeString}))
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

func TestColumn_MySQLType(t *testing.T) {
	c1 := &Column{Type: field.TypeString, Unique: true}
	require.Equal(t, "varchar(191)", c1.MySQLType("5.5"))
	require.Equal(t, "varchar(191)", c1.MySQLType("5.6.1"))
	require.Equal(t, "varchar(191)", c1.MySQLType("5.6.8"))
	require.Equal(t, "varchar(255)", c1.MySQLType("5.7"))
	require.Equal(t, "varchar(255)", c1.MySQLType("5.7.0"))
	require.Equal(t, "varchar(255)", c1.MySQLType("5.7.26-log"))
	require.Equal(t, "varchar(255)", c1.MySQLType("8-log"))

	c1 = &Column{Type: field.TypeJSON}
	require.Equal(t, "json", c1.MySQLType("5.7.8"))
	require.Equal(t, "json", c1.MySQLType("5.7.8-log"))
	require.Equal(t, "longblob", c1.MySQLType("5.5"))
	require.Equal(t, "longblob", c1.MySQLType("5.7"))
}
