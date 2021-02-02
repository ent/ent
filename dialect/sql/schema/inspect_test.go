// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"path"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestInspector_Tables(t *testing.T) {
	tests := []struct {
		name    string
		options []InspectOption
		before  map[string]func(mysqlMock)
		tables  []*Table
		wantErr bool
	}{
		{
			name: "default schema",
			before: map[string]func(mysqlMock){
				dialect.MySQL: func(mock mysqlMock) {
					mock.ExpectQuery(escape("SELECT `TABLE_NAME` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = (SELECT DATABASE())")).
						WillReturnRows(sqlmock.NewRows([]string{"TABLE_NAME"}))
				},
				dialect.SQLite: func(mock mysqlMock) {
					mock.ExpectQuery(escape("SELECT `name` FROM `sqlite_schema` WHERE `type` = ?")).
						WithArgs("table").
						WillReturnRows(sqlmock.NewRows([]string{"name"}))
				},
				dialect.Postgres: func(mock mysqlMock) {
					mock.ExpectQuery(escape(`SELECT "table_name" FROM "information_schema"."tables" WHERE "table_schema" = CURRENT_SCHEMA()`)).
						WillReturnRows(sqlmock.NewRows([]string{"name"}))
				},
			},
		},
		{
			name:    "custom schema",
			options: []InspectOption{WithSchema("public")},
			before: map[string]func(mysqlMock){
				dialect.MySQL: func(mock mysqlMock) {
					mock.ExpectQuery(escape("SELECT `TABLE_NAME` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = ?")).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"TABLE_NAME"}).
							AddRow("users"))
					mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?")).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
							AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
							AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "").
							AddRow("text", "longtext", "YES", "YES", "NULL", "", "", "").
							AddRow("uuid", "char(36)", "YES", "YES", "NULL", "", "", "utf8mb4_bin"))
					mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
							AddRow("PRIMARY", "id", "0", "1"))
				},
				dialect.SQLite: func(mock mysqlMock) {
					mock.ExpectQuery(escape("SELECT `name` FROM `sqlite_schema` WHERE `type` = ?")).
						WithArgs("table").
						WillReturnRows(sqlmock.NewRows([]string{"name"}).
							AddRow("users"))
					mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
							AddRow("id", "integer", 1, "NULL", 1).
							AddRow("name", "varchar(255)", 0, "NULL", 0).
							AddRow("text", "text", 0, "NULL", 0).
							AddRow("uuid", "uuid", 0, "NULL", 0))
					mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
						WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
				},
				dialect.Postgres: func(mock mysqlMock) {
					mock.ExpectQuery(escape(`SELECT "table_name" FROM "information_schema"."tables" WHERE "table_schema" = $1`)).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"name"}).
							AddRow("users"))
					mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name" FROM "information_schema"."columns" WHERE "table_schema" = $1 AND "table_name" = $2`)).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name"}).
							AddRow("id", "bigint", "NO", "NULL", "int8").
							AddRow("name", "character", "YES", "NULL", "bpchar").
							AddRow("text", "text", "YES", "NULL", "text").
							AddRow("uuid", "uuid", "YES", "NULL", "uuid"))
					mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "$1", "users"))).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
							AddRow("users_pkey", "id", "t", "t", 0))
				},
			},
			tables: func() []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt64, Increment: true},
						{Name: "name", Type: field.TypeString, Size: 255, Nullable: true},
						{Name: "text", Type: field.TypeString, Size: math.MaxInt32, Nullable: true},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
					}
				)
				return []*Table{t1}
			}(),
		},
	}
	for _, tt := range tests {
		for drv := range tt.before {
			t.Run(path.Join(drv, tt.name), func(t *testing.T) {
				db, mock, err := sqlmock.New()
				require.NoError(t, err)
				tt.before[drv](mysqlMock{mock})
				inspect, err := NewInspect(sql.OpenDB(drv, db), tt.options...)
				require.NoError(t, err)
				tables, err := inspect.Tables(context.Background())
				require.Equal(t, tt.wantErr, err != nil, err)
				tablesMatch(t, tables, tt.tables)
			})
		}
	}
}

func tablesMatch(t *testing.T, got, expected []*Table) {
	require.Equal(t, len(expected), len(got))
	for i := range got {
		columnsMatch(t, got[i].Columns, expected[i].Columns)
		columnsMatch(t, got[i].PrimaryKey, expected[i].PrimaryKey)
	}
}

func columnsMatch(t *testing.T, got, expected []*Column) {
	require.Equal(t, len(expected), len(got))
	for i := range got {
		c1, c2 := got[i], expected[i]
		require.Equal(t, c1.Name, c2.Name)
		require.Equal(t, c1.Nullable, c2.Nullable)
		require.True(t, c1.Type == c2.Type || c1.ConvertibleTo(c2))
	}
}
