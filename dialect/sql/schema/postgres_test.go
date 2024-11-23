// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestPostgres_scanPostgresColumnDefault(t *testing.T) {
	for _, test := range []struct {
		c   Column
		def string
	}{
		{Column{Type: field.TypeInt, Default: int64(1)}, "1::bigint"},
		{Column{Type: field.TypeString, Default: "val"}, "'val'::character varying"},
		{Column{Type: field.TypeString}, "val"}, // must quote
		{Column{Type: field.TypeString}, "current_setting('block_size')::bigint"},
		{Column{Type: field.TypeString}, "(current_setting('block_size'::text))::bigint"},
		{Column{Type: field.TypeInt, Default: int64(1)}, "1"},
		{Column{Type: field.TypeTime}, "CURRENT_TIMESTAMP"},
		{Column{Type: field.TypeUUID}, "gen_random_uuid()"},
		{Column{Type: field.TypeInt}, "('1')::bigint"}, // unable to detect
	} {
		c := &test.c
		e := c.Default
		c.Default = nil
		require.NoError(t, scanPostgresColumnDefault(c, sql.NullString{
			String: test.def,
			Valid:  true,
		}))
		require.Equalf(t, e, c.Default, "expect %v default to %v", test.def, e)
	}
}

func TestPostgres_Create(t *testing.T) {
	tests := []struct {
		name    string
		tables  []*Table
		options []MigrateOption
		before  func(pgMock)
		wantErr bool
	}{
		{
			name: "tx failed",
			before: func(mock pgMock) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
		{
			name: "unsupported version",
			before: func(mock pgMock) {
				mock.start("90000")
			},
			wantErr: true,
		},
		{
			name: "no tables",
			before: func(mock pgMock) {
				mock.start("120000")
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table",
			tables: []*Table{
				{
					Name: "users",
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Columns: []*Column{
						{Name: "id", Type: field.TypeUUID, Default: "uuid_generate_v4()"},
						{Name: "block_size", Type: field.TypeInt, Default: "current_setting('block_size')::bigint"},
						{Name: "name", Type: field.TypeString, Nullable: true, Collation: "he_IL"},
						{Name: "age", Type: field.TypeInt},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
						{Name: "enums", Type: field.TypeEnum, Enums: []string{"a", "b"}, Default: "a"},
						{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.Postgres: "numeric(5,2)"}},
						{Name: "strings", Type: field.TypeOther, SchemaType: map[string]string{dialect.Postgres: "text[]"}, Nullable: true},
						{Name: "fixed_string", Type: field.TypeString, SchemaType: map[string]string{dialect.Postgres: "varchar(100)"}},
					},
					Annotation: &entsql.Annotation{
						Check: "price > 0",
						Checks: map[string]string{
							"valid_age":  "age > 0",
							"valid_name": "name <> ''",
						},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "users"("id" uuid NOT NULL DEFAULT uuid_generate_v4(), "block_size" bigint NOT NULL DEFAULT current_setting('block_size')::bigint, "name" varchar NULL COLLATE "he_IL", "age" bigint NOT NULL, "doc" jsonb NULL, "enums" varchar NOT NULL DEFAULT 'a', "price" numeric(5,2) NOT NULL, "strings" text[] NULL, "fixed_string" varchar(100) NOT NULL, PRIMARY KEY("id"), CHECK (price > 0), CONSTRAINT "valid_age" CHECK (age > 0), CONSTRAINT "valid_name" CHECK (name <> ''))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table with foreign key",
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "created_at", Type: field.TypeTime},
						{Name: "inet", Type: field.TypeString, Unique: true, SchemaType: map[string]string{dialect.Postgres: "inet"}},
					}
					c2 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString},
						{Name: "owner_id", Type: field.TypeInt, Nullable: true},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
					}
					t2 = &Table{
						Name:       "pets",
						Columns:    c2,
						PrimaryKey: c2[0:1],
						ForeignKeys: []*ForeignKey{
							{
								Symbol:     "pets_owner",
								Columns:    c2[2:],
								RefTable:   t1,
								RefColumns: c1[0:1],
								OnDelete:   Cascade,
							},
						},
					}
				)
				return []*Table{t1, t2}
			}(),
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "users"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "name" varchar NULL, "created_at" timestamp with time zone NOT NULL, "inet" inet UNIQUE NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "pets"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "name" varchar NOT NULL, "owner_id" bigint NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.fkExists("pets_owner", false)
				mock.ExpectExec(escape(`ALTER TABLE "pets" ADD CONSTRAINT "pets_owner" FOREIGN KEY("owner_id") REFERENCES "users"("id") ON DELETE CASCADE`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table with foreign key disabled",
			options: []MigrateOption{
				WithForeignKeys(false),
			},
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "created_at", Type: field.TypeTime},
					}
					c2 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString},
						{Name: "owner_id", Type: field.TypeInt, Nullable: true},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
					}
					t2 = &Table{
						Name:       "pets",
						Columns:    c2,
						PrimaryKey: c2[0:1],
						ForeignKeys: []*ForeignKey{
							{
								Symbol:     "pets_owner",
								Columns:    c2[2:],
								RefTable:   t1,
								RefColumns: c1[0:1],
								OnDelete:   Cascade,
							},
						},
					}
				)
				return []*Table{t1, t2}
			}(),
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "users"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "name" varchar NULL, "created_at" timestamp with time zone NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "pets"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "name" varchar NOT NULL, "owner_id" bigint NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "scan table with default",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "block_size", Type: field.TypeInt, Default: "current_setting('block_size')::bigint"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "nextval('users_colname_seq'::regclass)", "int4", nil, nil, nil).
						AddRow("block_size", "bigint", "NO", "current_setting('block_size')::bigint", "int4", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ALTER COLUMN "block_size" TYPE bigint, ALTER COLUMN "block_size" SET NOT NULL, ALTER COLUMN "block_size" SET DEFAULT current_setting('block_size')::bigint`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "scan table with custom type",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "custom", Type: field.TypeOther, SchemaType: map[string]string{dialect.Postgres: "customtype"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "nextval('users_colname_seq'::regclass)", "NULL", nil, nil, nil).
						AddRow("custom", "USER-DEFINED", "NO", "NULL", "customtype", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "add column to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "text", Type: field.TypeString, Nullable: true, Size: math.MaxInt32},
						{Name: "age", Type: field.TypeInt},
						{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{dialect.Postgres: "date"}, Default: "CURRENT_DATE"},
						{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "date"}, Nullable: true},
						{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
						{Name: "cidr", Type: field.TypeString, SchemaType: map[string]string{dialect.Postgres: "cidr"}},
						{Name: "point", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "point"}},
						{Name: "line", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "line"}},
						{Name: "lseg", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "lseg"}},
						{Name: "box", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "box"}},
						{Name: "path", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "path"}},
						{Name: "polygon", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "polygon"}},
						{Name: "circle", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "circle"}},
						{Name: "macaddr", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "macaddr"}},
						{Name: "macaddr8", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.Postgres: "macaddr8"}},
						{Name: "strings", Type: field.TypeOther, SchemaType: map[string]string{dialect.Postgres: "text[]"}, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character varying", "YES", "NULL", "varchar", nil, nil, nil).
						AddRow("uuid", "uuid", "YES", "NULL", "uuid", nil, nil, nil).
						AddRow("created_at", "date", "NO", "CURRENT_DATE", "date", nil, nil, nil).
						AddRow("updated_at", "timestamp with time zone", "YES", "NULL", "timestamptz", nil, nil, nil).
						AddRow("deleted_at", "date", "YES", "NULL", "date", nil, nil, nil).
						AddRow("text", "text", "YES", "NULL", "text", nil, nil, nil).
						AddRow("cidr", "cidr", "NO", "NULL", "cidr", nil, nil, nil).
						AddRow("inet", "inet", "YES", "NULL", "inet", nil, nil, nil).
						AddRow("point", "point", "YES", "NULL", "point", nil, nil, nil).
						AddRow("line", "line", "YES", "NULL", "line", nil, nil, nil).
						AddRow("lseg", "lseg", "YES", "NULL", "lseg", nil, nil, nil).
						AddRow("box", "box", "YES", "NULL", "box", nil, nil, nil).
						AddRow("path", "path", "YES", "NULL", "path", nil, nil, nil).
						AddRow("polygon", "polygon", "YES", "NULL", "polygon", nil, nil, nil).
						AddRow("circle", "circle", "YES", "NULL", "circle", nil, nil, nil).
						AddRow("macaddr", "macaddr", "YES", "NULL", "macaddr", nil, nil, nil).
						AddRow("macaddr8", "macaddr8", "YES", "NULL", "macaddr8", nil, nil, nil).
						AddRow("strings", "ARRAY", "YES", "NULL", "_text", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "age" bigint NOT NULL, ALTER COLUMN "created_at" TYPE date, ALTER COLUMN "created_at" SET NOT NULL, ALTER COLUMN "created_at" SET DEFAULT CURRENT_DATE, ALTER COLUMN "deleted_at" TYPE timestamp with time zone, ALTER COLUMN "deleted_at" DROP NOT NULL`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add int column with default value to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "age", Type: field.TypeInt, Default: 10},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil).
						AddRow("doc", "jsonb", "YES", "NULL", "jsonb", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "age" bigint NOT NULL DEFAULT 10`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add blob columns",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "blob", Type: field.TypeBytes, Size: 1e3},
						{Name: "longblob", Type: field.TypeBytes, Size: 1e6},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil).
						AddRow("doc", "jsonb", "YES", "NULL", "jsonb", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "blob" bytea NOT NULL, ADD COLUMN "longblob" bytea NOT NULL`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add float column with default value to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "age", Type: field.TypeFloat64, Default: 10.1},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "age" double precision NOT NULL DEFAULT 10.1`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add bool column with default value to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "age", Type: field.TypeBool, Default: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "age" boolean NOT NULL DEFAULT true`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add string column with default value to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "nick", Type: field.TypeString, Default: "unknown"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "nick" varchar NOT NULL DEFAULT 'unknown'`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "drop column to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			options: []MigrateOption{WithDropColumn(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" DROP COLUMN "name"`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify column to nullable",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "NO", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ALTER COLUMN "name" TYPE varchar, ALTER COLUMN "name" DROP NOT NULL`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify column default value",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Default: "unknown"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "NO", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ALTER COLUMN "name" TYPE varchar, ALTER COLUMN "name" SET NOT NULL, ALTER COLUMN "name" SET DEFAULT 'unknown'`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "apply uniqueness on column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt, Unique: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("age", "bigint", "NO", "NULL", "int8", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`CREATE UNIQUE INDEX IF NOT EXISTS "users_age" ON "users"("age")`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "remove uniqueness from column without option",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("age", "bigint", "NO", "NULL", "int8", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0).
						AddRow("users_age_key", "age", "f", "t", 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "remove uniqueness from column with option",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			options: []MigrateOption{WithDropIndex(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("age", "bigint", "NO", "NULL", "int8", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0).
						AddRow("users_age_key", "age", "f", "t", 0))
				mock.ExpectQuery(escape(`SELECT COUNT(*) FROM "information_schema"."table_constraints" WHERE "table_schema" = CURRENT_SCHEMA() AND "constraint_type" = $1 AND "constraint_name" = $2`)).
					WithArgs("UNIQUE", "users_age_key").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectExec(escape(`ALTER TABLE "users" DROP CONSTRAINT "users_age_key"`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add and remove indexes",
			tables: func() []*Table {
				c1 := []*Column{
					{Name: "id", Type: field.TypeInt, Increment: true},
					// Add implicit index.
					{Name: "age", Type: field.TypeInt, Unique: true},
					{Name: "score", Type: field.TypeInt},
				}
				c2 := []*Column{
					{Name: "id", Type: field.TypeInt, Increment: true},
					{Name: "score", Type: field.TypeInt},
					{Name: "email", Type: field.TypeString},
				}
				return []*Table{
					{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
						Indexes: Indexes{
							// Change non-unique index to unique.
							{Name: "user_score", Columns: c1[2:3], Unique: true},
						},
					},
					{
						Name:       "equipment",
						Columns:    c2,
						PrimaryKey: c2[0:1],
						Indexes: Indexes{
							{Name: "equipment_score", Columns: c2[1:2]},
							// Index should not be changed.
							{Name: "equipment_email", Unique: true, Columns: c2[2:]},
						},
					},
				}
			}(),
			options: []MigrateOption{WithDropIndex(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("age", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("score", "bigint", "NO", "NULL", "int8", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0).
						AddRow("user_score", "score", "f", "f", 0))
				mock.ExpectQuery(escape(`SELECT COUNT(*) FROM "information_schema"."table_constraints" WHERE "table_schema" = CURRENT_SCHEMA() AND "constraint_type" = $1 AND "constraint_name" = $2`)).
					WithArgs("UNIQUE", "user_score").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).
						AddRow(0))
				mock.ExpectExec(escape(`DROP INDEX "user_score"`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape(`CREATE UNIQUE INDEX IF NOT EXISTS "users_age" ON "users"("age")`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape(`CREATE UNIQUE INDEX IF NOT EXISTS "user_score" ON "users"("score")`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("equipment", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("equipment").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("score", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("email", "character varying", "YES", "NULL", "varchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "equipment"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0).
						AddRow("equipment_score", "score", "f", "f", 0).
						AddRow("equipment_email", "email", "f", "t", 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "add edge to table",
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "spouse_id", Type: field.TypeInt, Nullable: true},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
						ForeignKeys: []*ForeignKey{
							{
								Symbol:     "user_spouse" + strings.Repeat("_", 64), // super long fk.
								Columns:    c1[2:],
								RefColumns: c1[0:1],
								OnDelete:   Cascade,
							},
						},
					}
				)
				t1.ForeignKeys[0].RefTable = t1
				return []*Table{t1}
			}(),
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "YES", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ADD COLUMN "spouse_id" bigint NULL`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.fkExists("user_spouse____________________390ed76f91d3c57cd3516e7690f621dc", false)
				mock.ExpectExec(`ALTER TABLE "users" ADD CONSTRAINT ".{63}" FOREIGN KEY\("spouse_id"\) REFERENCES "users"\("id"\) ON DELETE CASCADE`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "universal id for all tables",
			tables: []*Table{
				NewTable("users").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
				NewTable("groups").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
			},
			options: []MigrateOption{WithGlobalUniqueID(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("ent_types", false)
				// create ent_types table.
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "ent_types"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, "type" varchar UNIQUE NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("users", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "users"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range.
				mock.ExpectExec(escape(`INSERT INTO "ent_types" ("type") VALUES ($1)`)).
					WithArgs("users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(`ALTER TABLE "users" ALTER COLUMN "id" RESTART WITH 1`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("groups", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "groups"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape(`INSERT INTO "ent_types" ("type") VALUES ($1)`)).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(`ALTER TABLE "groups" ALTER COLUMN "id" RESTART WITH 4294967296`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "universal id for new tables",
			tables: []*Table{
				NewTable("users").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
				NewTable("groups").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
			},
			options: []MigrateOption{WithGlobalUniqueID(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("ent_types", true)
				// query ent_types table.
				mock.ExpectQuery(`SELECT "type" FROM "ent_types" ORDER BY "id" ASC`).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				// query users table.
				mock.tableExists("users", true)
				// users table has no changes.
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "YES", "NULL", "int8", nil, nil, nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				// query groups table.
				mock.tableExists("groups", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "groups"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape(`INSERT INTO "ent_types" ("type") VALUES ($1)`)).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(`ALTER TABLE "groups" ALTER COLUMN "id" RESTART WITH 4294967296`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "universal id for restored tables",
			tables: []*Table{
				NewTable("users").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
				NewTable("groups").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
			},
			options: []MigrateOption{WithGlobalUniqueID(true)},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("ent_types", true)
				// query ent_types table.
				mock.ExpectQuery(`SELECT "type" FROM "ent_types" ORDER BY "id" ASC`).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				// query and create users (restored table).
				mock.tableExists("users", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "users"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range (without inserting to ent_types).
				mock.ExpectExec(`ALTER TABLE "users" ALTER COLUMN "id" RESTART WITH 1`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// query groups table.
				mock.tableExists("groups", false)
				mock.ExpectExec(escape(`CREATE TABLE IF NOT EXISTS "groups"("id" bigint GENERATED BY DEFAULT AS IDENTITY NOT NULL, PRIMARY KEY("id"))`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape(`INSERT INTO "ent_types" ("type") VALUES ($1)`)).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(`ALTER TABLE "groups" ALTER COLUMN "id" RESTART WITH 4294967296`).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "no modify numeric column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.Postgres: "numeric(6,4)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("price", "numeric", "NO", "NULL", "numeric", "6", "4", nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify numeric column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "price", Type: field.TypeFloat64, Nullable: false, SchemaType: map[string]string{dialect.Postgres: "numeric(6,4)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("price", "numeric", "NO", "NULL", "numeric", "5", "4", nil))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ALTER COLUMN "price" TYPE numeric(6,4), ALTER COLUMN "price" SET NOT NULL`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "no modify fixed size varchar column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, SchemaType: map[string]string{dialect.Postgres: "varchar(20)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character varying", "NO", "NULL", "varchar", nil, nil, 20))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify fixed size varchar column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, SchemaType: map[string]string{dialect.Postgres: "varchar(20)"}, Default: "unknown"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock pgMock) {
				mock.start("120000")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
						AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
						AddRow("name", "character varying", "NO", "NULL", "varchar", nil, nil, 10))
				mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "CURRENT_SCHEMA()", "users"))).
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
						AddRow("users_pkey", "id", "t", "t", 0))
				mock.ExpectExec(escape(`ALTER TABLE "users" ALTER COLUMN "name" TYPE varchar(20), ALTER COLUMN "name" SET NOT NULL, ALTER COLUMN "name" SET DEFAULT 'unknown'`)).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.before(pgMock{mock})
			migrate, err := NewMigrate(sql.OpenDB("postgres", db), append(tt.options, WithAtlas(false))...)
			require.NoError(t, err)
			err = migrate.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

type pgMock struct {
	sqlmock.Sqlmock
}

func (m pgMock) start(version string) {
	m.ExpectQuery(escape("SHOW server_version_num")).
		WillReturnRows(sqlmock.NewRows([]string{"server_version_num"}).AddRow(version))
	m.ExpectBegin()
}

func (m pgMock) tableExists(table string, exists bool) {
	count := 0
	if exists {
		count = 1
	}
	m.ExpectQuery(escape(`SELECT COUNT(*) FROM "information_schema"."tables" WHERE "table_schema" = CURRENT_SCHEMA() AND "table_name" = $1`)).
		WithArgs(table).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func (m pgMock) fkExists(fk string, exists bool) {
	count := 0
	if exists {
		count = 1
	}
	m.ExpectQuery(escape(`SELECT COUNT(*) FROM "information_schema"."table_constraints" WHERE "table_schema" = CURRENT_SCHEMA() AND "constraint_type" = $1 AND "constraint_name" = $2`)).
		WithArgs("FOREIGN KEY", fk).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}
