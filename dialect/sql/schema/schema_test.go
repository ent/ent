// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"testing"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect"
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

func TestDump(t *testing.T) {
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
		OnUpdate:   SetDefault,
	})
	users.SetAnnotation(&entsql.Annotation{Table: "Users"})
	pets := &Table{
		Name: "pets",
		Columns: []*Column{
			{Name: "id", Type: field.TypeInt},
			{Name: "name", Type: field.TypeString},
			{Name: "fur_color", Type: field.TypeEnum, Enums: []string{"black", "white"}},
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
		Columns:    pets.Columns[3:],
		RefTable:   users,
		RefColumns: users.Columns[:1],
		OnDelete:   SetDefault,
	})
	tables = []*Table{users, pets}

	my := func(length int) string {
		return fmt.Sprintf("-- Create \"users\" table\nCREATE TABLE `users` (`id` bigint NOT NULL, `name` varchar(%d) NOT NULL, `spouse_id` bigint NOT NULL, PRIMARY KEY (`id`), INDEX `name` (`name`), FOREIGN KEY (`spouse_id`) REFERENCES `users` (`id`) ON UPDATE SET DEFAULT) CHARSET utf8mb4 COLLATE utf8mb4_bin;\n-- Create \"pets\" table\nCREATE TABLE `pets` (`id` bigint NOT NULL, `name` varchar(%d) NOT NULL, `fur_color` enum('black','white') NOT NULL, `owner_id` bigint NOT NULL, UNIQUE INDEX `name` (`name` DESC), FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET DEFAULT) CHARSET utf8mb4 COLLATE utf8mb4_bin;\n", length, length)
	}

	pg := "-- Create \"users\" table\nCREATE TABLE \"users\" (\"id\" bigint NOT NULL, \"name\" character varying NOT NULL, \"spouse_id\" bigint NOT NULL, PRIMARY KEY (\"id\"), FOREIGN KEY (\"spouse_id\") REFERENCES \"users\" (\"id\") ON UPDATE SET DEFAULT);\n-- Create index \"name\" to table: \"users\"\nCREATE INDEX \"name\" ON \"users\" (\"name\");\n-- Create \"pets\" table\nCREATE TABLE \"pets\" (\"id\" bigint NOT NULL, \"name\" character varying NOT NULL, \"fur_color\" character varying NOT NULL, \"owner_id\" bigint NOT NULL, FOREIGN KEY (\"owner_id\") REFERENCES \"users\" (\"id\") ON DELETE SET DEFAULT);\n-- Create index \"name\" to table: \"pets\"\nCREATE UNIQUE INDEX \"name\" ON \"pets\" (\"name\" DESC);\n"

	for _, tt := range []struct{ dialect, version, expected string }{
		{
			dialect.SQLite, "",
			"-- Create \"users\" table\nCREATE TABLE `users` (`id` integer NOT NULL, `name` text NOT NULL, `spouse_id` integer NOT NULL, PRIMARY KEY (`id`), FOREIGN KEY (`spouse_id`) REFERENCES `users` (`id`) ON UPDATE SET DEFAULT);\n-- Create index \"name\" to table: \"users\"\nCREATE INDEX `name` ON `users` (`name`);\n-- Create \"pets\" table\nCREATE TABLE `pets` (`id` integer NOT NULL, `name` text NOT NULL, `fur_color` text NOT NULL, `owner_id` integer NOT NULL, FOREIGN KEY (`owner_id`) REFERENCES `users` (`id`) ON DELETE SET DEFAULT);\n-- Create index \"name\" to table: \"pets\"\nCREATE UNIQUE INDEX `name` ON `pets` (`name` DESC);\n",
		},
		{dialect.MySQL, "5.6", my(191)},
		{dialect.MySQL, "5.7", my(255)},
		{dialect.MySQL, "8", my(255)},
		{dialect.Postgres, "12", pg},
		{dialect.Postgres, "13", pg},
		{dialect.Postgres, "14", pg},
		{dialect.Postgres, "15", pg},
	} {
		t.Run(fmt.Sprintf("%s:%s", tt.dialect, tt.version), func(t *testing.T) {
			ac, err := Dump(context.Background(), tt.dialect, tt.version, tables, func(o *migrate.PlanOptions) {
				o.Indent = ""
			})
			require.NoError(t, err)
			require.Equal(t, tt.expected, ac)
		})
	}
}
