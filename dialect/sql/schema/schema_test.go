// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"testing"

	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/field"

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
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 1}))

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
	require.True(t, c1.ConvertibleTo(&Column{Type: field.TypeString, Size: 1}))
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

	c1 = &Column{Type: field.TypeUUID}
	require.NoError(t, c1.ScanDefault("gen_random_uuid()"))
	require.Equal(t, nil, c1.Default)
	require.NoError(t, c1.ScanDefault("00000000-0000-0000-0000-000000000000"))
	require.Equal(t, "00000000-0000-0000-0000-000000000000", c1.Default)
}

func TestCopyTables(t *testing.T) {
	users := &Table{
		Name: "users",
		Columns: []*Column{
			{Name: "id", Type: field.TypeInt},
			{Name: "name", Type: field.TypeString},
			{Name: "spouse_id", Type: field.TypeInt},
		},
	}
	users.PrimaryKey = users.Columns[:1]
	users.Indexes = append(users.Indexes, &Index{
		Name:    "name",
		Columns: users.Columns[1:2],
	})
	users.AddForeignKey(&ForeignKey{
		Columns:    users.Columns[2:],
		RefTable:   users,
		RefColumns: users.Columns[:1],
		OnUpdate:   SetNull,
	})
	users.SetAnnotation(&entsql.Annotation{Table: "Users"})
	pets := &Table{
		Name: "pets",
		Columns: []*Column{
			{Name: "id", Type: field.TypeInt},
			{Name: "name", Type: field.TypeString},
			{Name: "owner_id", Type: field.TypeInt},
		},
	}
	pets.Indexes = append(pets.Indexes, &Index{
		Name:       "name",
		Unique:     true,
		Columns:    pets.Columns[1:2],
		Annotation: entsql.Desc(),
	})
	pets.AddForeignKey(&ForeignKey{
		Columns:    pets.Columns[2:],
		RefTable:   users,
		RefColumns: users.Columns[:1],
		OnDelete:   SetDefault,
	})
	tables := []*Table{users, pets}
	copyT, err := CopyTables(tables)
	require.NoError(t, err)
	require.Equal(t, tables, copyT)
}
