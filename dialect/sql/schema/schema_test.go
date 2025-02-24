// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
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
	petsWithoutFur := &Table{
		Name:       "pets_without_fur",
		View:       true,
		Columns:    append(pets.Columns[:2], pets.Columns[3]),
		Annotation: entsql.View("SELECT id, name, owner_id FROM pets"),
	}
	petNames := &Table{
		Name:    "pet_names",
		View:    true,
		Columns: pets.Columns[1:1],
		Annotation: entsql.ViewFor(dialect.Postgres, func(s *sql.Selector) {
			s.Select("name").From(sql.Table("pets"))
		}),
	}
	tables = []*Table{users, pets, petsWithoutFur, petNames}

	my := func(length int) string {
		return fmt.Sprintf(strings.ReplaceAll(`-- Add new schema named "s1"
CREATE DATABASE $s1$;
-- Add new schema named "s2"
CREATE DATABASE $s2$;
-- Add new schema named "s3"
CREATE DATABASE $s3$;
-- Create "users" table
CREATE TABLE $s1$.$users$ (
  $id$ bigint NOT NULL,
  $name$ varchar(%d) NOT NULL,
  $spouse_id$ bigint NOT NULL,
  PRIMARY KEY ($id$),
  INDEX $name$ ($name$),
  FOREIGN KEY ($spouse_id$) REFERENCES $s1$.$users$ ($id$) ON UPDATE SET DEFAULT
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Create "pets" table
CREATE TABLE $s2$.$pets$ (
  $id$ bigint NOT NULL,
  $name$ varchar(%d) NOT NULL,
  $owner_id$ bigint NOT NULL,
  $owner_id$ bigint NOT NULL,
  UNIQUE INDEX $name$ ($name$ DESC),
  FOREIGN KEY ($owner_id$) REFERENCES $s1$.$users$ ($id$) ON DELETE SET DEFAULT
) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- Add "pets_without_fur" view
CREATE VIEW $s3$.$pets_without_fur$ ($id$, $name$, $owner_id$) AS SELECT id, name, owner_id FROM pets;
`, "$", "`"), length, length)
	}

	pg := `-- Add new schema named "s1"
CREATE SCHEMA "s1";
-- Add new schema named "s2"
CREATE SCHEMA "s2";
-- Add new schema named "s3"
CREATE SCHEMA "s3";
-- Create "users" table
CREATE TABLE "s1"."users" (
  "id" bigint NOT NULL,
  "name" character varying NOT NULL,
  "spouse_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  FOREIGN KEY ("spouse_id") REFERENCES "s1"."users" ("id") ON UPDATE SET DEFAULT
);
-- Create index "name" to table: "users"
CREATE INDEX "name" ON "s1"."users" ("name");
-- Create "pets" table
CREATE TABLE "s2"."pets" (
  "id" bigint NOT NULL,
  "name" character varying NOT NULL,
  "owner_id" bigint NOT NULL,
  "owner_id" bigint NOT NULL,
  FOREIGN KEY ("owner_id") REFERENCES "s1"."users" ("id") ON DELETE SET DEFAULT
);
-- Create index "name" to table: "pets"
CREATE UNIQUE INDEX "name" ON "s2"."pets" ("name" DESC);
-- Add "pets_without_fur" view
CREATE VIEW "s3"."pets_without_fur" ("id", "name", "owner_id") AS SELECT id, name, owner_id FROM pets;
-- Add "pet_names" view
CREATE VIEW "s3"."pet_names" AS SELECT "name" FROM "pets";
`

	for _, tt := range []struct{ dialect, version, expected string }{
		{
			dialect.SQLite, "",
			strings.ReplaceAll(`-- Create "users" table
CREATE TABLE $users$ (
  $id$ integer NOT NULL,
  $name$ text NOT NULL,
  $spouse_id$ integer NOT NULL,
  PRIMARY KEY ($id$),
  FOREIGN KEY ($spouse_id$) REFERENCES $users$ ($id$) ON UPDATE SET DEFAULT
);
-- Create index "name" to table: "users"
CREATE INDEX $name$ ON $users$ ($name$);
-- Create "pets" table
CREATE TABLE $pets$ (
  $id$ integer NOT NULL,
  $name$ text NOT NULL,
  $owner_id$ integer NOT NULL,
  $owner_id$ integer NOT NULL,
  FOREIGN KEY ($owner_id$) REFERENCES $users$ ($id$) ON DELETE SET DEFAULT
);
-- Create index "name" to table: "pets"
CREATE UNIQUE INDEX $name$ ON $pets$ ($name$ DESC);
-- Add "pets_without_fur" view
CREATE VIEW $pets_without_fur$ ($id$, $name$, $owner_id$) AS SELECT id, name, owner_id FROM pets;
`, "$", "`"),
		},
		{dialect.MySQL, "5.6", my(191)},
		{dialect.MySQL, "5.7", my(255)},
		{dialect.MySQL, "8", my(255)},
		{dialect.Postgres, "12", pg},
		{dialect.Postgres, "13", pg},
		{dialect.Postgres, "14", pg},
		{dialect.Postgres, "15", pg},
	} {
		n := tt.dialect
		if tt.version != "" {
			n += ":" + tt.version
		}
		if tt.dialect != dialect.SQLite {
			tables[0].Schema = "s1"
			tables[1].Schema = "s2"
			tables[2].Schema = "s3"
			tables[3].Schema = "s3"
		}
		t.Run(n, func(t *testing.T) {
			ac, err := Dump(context.Background(), tt.dialect, tt.version, tables)
			require.NoError(t, err)
			require.Equal(t, tt.expected, ac)
		})
	}
}
