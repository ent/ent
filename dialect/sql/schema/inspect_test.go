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
		tables  func(drv string) []*Table
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
			tables: func(drv string) []*Table {
				return nil
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
							AddRow("users").
							AddRow("pets").
							AddRow("groups").
							AddRow("user_groups"))
					mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?")).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
							AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
							AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
							AddRow("text", "longtext", "YES", "YES", "NULL", "", "", "", nil, nil).
							AddRow("uuid", "char(36)", "YES", "YES", "NULL", "", "", "utf8mb4_bin", nil, nil).
							AddRow("price", "decimal(6, 4)", "NO", "YES", "NULL", "", "", "", "6", "4").
							AddRow("bank_id", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
					mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
							AddRow("PRIMARY", "id", nil, "0", "1"))
					mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?")).
						WithArgs("public", "pets").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
							AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
							AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
							AddRow("user_pets", "bigint(20)", "YES", "YES", "NULL", "", "", "", nil, nil))
					mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
						WithArgs("public", "pets").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
							AddRow("PRIMARY", "id", nil, "0", "1"))
					mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?")).
						WithArgs("public", "groups").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
							AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
							AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
					mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
						WithArgs("public", "groups").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
							AddRow("PRIMARY", "id", nil, "0", "1"))
					mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ?")).
						WithArgs("public", "user_groups").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
							AddRow("user_id", "bigint(20)", "NO", "YES", "NULL", "", "", "", nil, nil).
							AddRow("group_id", "bigint(20)", "NO", "YES", "NULL", "", "", "", nil, nil))
					mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = ? AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
						WithArgs("public", "user_groups").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}))
				},
				dialect.SQLite: func(mock mysqlMock) {
					mock.ExpectQuery(escape("SELECT `name` FROM `sqlite_schema` WHERE `type` = ?")).
						WithArgs("table").
						WillReturnRows(sqlmock.NewRows([]string{"name"}).
							AddRow("users").
							AddRow("pets").
							AddRow("groups").
							AddRow("user_groups"))
					mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
							AddRow("id", "integer", 1, "NULL", 1).
							AddRow("name", "varchar(255)", 0, "NULL", 0).
							AddRow("text", "text", 0, "NULL", 0).
							AddRow("uuid", "uuid", 0, "NULL", 0).
							AddRow("price", "real", 1, "NULL", 0).
							AddRow("bank_id", "varchar(255)", 1, "NULL", 0))
					mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
						WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
					mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('pets') ORDER BY `pk`")).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
							AddRow("id", "integer", 1, "NULL", 1).
							AddRow("name", "varchar(255)", 0, "NULL", 0).
							AddRow("user_pets", "integer", 0, "NULL", 0))
					mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('pets')")).
						WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
					mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('groups') ORDER BY `pk`")).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
							AddRow("id", "integer", 1, "NULL", 1).
							AddRow("name", "varchar(255)", 1, "NULL", 0))
					mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('groups')")).
						WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
					mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('user_groups') ORDER BY `pk`")).
						WithArgs().
						WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
							AddRow("user_id", "integer", 1, "NULL", 0).
							AddRow("group_id", "integer", 1, "NULL", 0))
					mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('user_groups')")).
						WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
				},
				dialect.Postgres: func(mock mysqlMock) {
					mock.ExpectQuery(escape(`SELECT "table_name" FROM "information_schema"."tables" WHERE "table_schema" = $1`)).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"name"}).
							AddRow("users").
							AddRow("pets").
							AddRow("groups").
							AddRow("user_groups"))
					mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = $1 AND "table_name" = $2`)).
						WithArgs("public", "users").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
							AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
							AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil).
							AddRow("text", "text", "YES", "NULL", "text", nil, nil, nil).
							AddRow("uuid", "uuid", "YES", "NULL", "uuid", nil, nil, nil).
							AddRow("price", "numeric", "NO", "NULL", "numeric", "6", "4", nil).
							AddRow("bank_id", "character", "NO", "NULL", "bpchar", nil, nil, 20))
					mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "$1", "users"))).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
							AddRow("users_pkey", "id", "t", "t", 0))
					mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = $1 AND "table_name" = $2`)).
						WithArgs("public", "pets").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
							AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
							AddRow("name", "character", "YES", "NULL", "bpchar", nil, nil, nil).
							AddRow("user_pets", "bigint", "YES", "NULL", "int8", nil, nil, nil))
					mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "$1", "pets"))).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
							AddRow("pets_pkey", "id", "t", "t", 0))
					mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = $1 AND "table_name" = $2`)).
						WithArgs("public", "groups").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
							AddRow("id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
							AddRow("name", "character", "NO", "NULL", "bpchar", nil, nil, nil))
					mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "$1", "groups"))).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}).
							AddRow("groups_pkey", "id", "t", "t", 0))
					mock.ExpectQuery(escape(`SELECT "column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length" FROM "information_schema"."columns" WHERE "table_schema" = $1 AND "table_name" = $2`)).
						WithArgs("public", "user_groups").
						WillReturnRows(sqlmock.NewRows([]string{"column_name", "data_type", "is_nullable", "column_default", "udt_name", "numeric_precision", "numeric_scale", "character_maximum_length"}).
							AddRow("user_id", "bigint", "NO", "NULL", "int8", nil, nil, nil).
							AddRow("group_id", "bigint", "NO", "NULL", "int8", nil, nil, nil))
					mock.ExpectQuery(escape(fmt.Sprintf(indexesQuery, "$1", "user_groups"))).
						WithArgs("public").
						WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "primary", "unique", "seq_in_index"}))
					mock.ExpectQuery(escape(fmt.Sprintf(fkQuery, "users"))).
						WillReturnRows(sqlmock.NewRows([]string{"table_schema", "constraint_name", "table_name", "column_name", "foreign_table_schema", "foreign_table_name", "foreign_column_name"}))
					mock.ExpectQuery(escape(fmt.Sprintf(fkQuery, "pets"))).
						WillReturnRows(sqlmock.NewRows([]string{"table_schema", "constraint_name", "table_name", "column_name", "foreign_table_schema", "foreign_table_name", "foreign_column_name"}).
							AddRow("public", "pet_users_pets", "pets", "user_pets", "public", "users", "id"))
					mock.ExpectQuery(escape(fmt.Sprintf(fkQuery, "groups"))).
						WillReturnRows(sqlmock.NewRows([]string{"table_schema", "constraint_name", "table_name", "column_name", "foreign_table_schema", "foreign_table_name", "foreign_column_name"}))
					mock.ExpectQuery(escape(fmt.Sprintf(fkQuery, "user_groups"))).
						WillReturnRows(sqlmock.NewRows([]string{"table_schema", "constraint_name", "table_name", "column_name", "foreign_table_schema", "foreign_table_name", "foreign_column_name"}).
							AddRow("public", "user_groups_group_id", "user_groups", "group_id", "public", "groups", "id").
							AddRow("public", "user_groups_user_id", "user_groups", "user_id", "public", "users", "id"))
				},
			},
			tables: func(drv string) []*Table {
				var (
					c1 = []*Column{
						{Name: "id", Type: field.TypeInt64, Increment: true},
						{Name: "name", Type: field.TypeString, Size: 255, Nullable: true},
						{Name: "text", Type: field.TypeString, Size: math.MaxInt32, Nullable: true},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{
							dialect.MySQL:    "decimal(6,4)",
							dialect.Postgres: "numeric(6,4)",
						}},
						{Name: "bank_id", Type: field.TypeString, SchemaType: map[string]string{
							dialect.Postgres: "varchar(20)",
						}},
					}
					t1 = &Table{
						Name:       "users",
						Columns:    c1,
						PrimaryKey: c1[0:1],
					}
					c2 = []*Column{
						{Name: "id", Type: field.TypeInt64, Increment: true},
						{Name: "name", Type: field.TypeString, Size: 255, Nullable: true},
						{Name: "user_pets", Type: field.TypeInt64, Nullable: true},
					}
					t2 = &Table{
						Name:       "pets",
						Columns:    c2,
						PrimaryKey: c2[0:1],
					}
					c3 = []*Column{
						{Name: "id", Type: field.TypeInt64, Increment: true},
						{Name: "name", Type: field.TypeString},
					}
					t3 = &Table{
						Name:       "groups",
						Columns:    c3,
						PrimaryKey: c3[0:1],
					}
					c4 = []*Column{
						{Name: "user_id", Type: field.TypeInt64},
						{Name: "group_id", Type: field.TypeInt64},
					}
					t4 = &Table{
						Name:    "user_groups",
						Columns: c4,
					}
				)

				// Only postgres currently supports foreign key inspection
				if drv == dialect.Postgres {
					t2.ForeignKeys = []*ForeignKey{
						{
							Symbol:     "pet_users_pets",
							Columns:    []*Column{c2[2]},
							RefTable:   t1,
							RefColumns: []*Column{c1[0]},
						},
					}
					t4.ForeignKeys = []*ForeignKey{
						{
							Symbol:     "user_groups_group_id",
							Columns:    []*Column{c4[1]},
							RefTable:   t3,
							RefColumns: []*Column{c3[0]},
						},
						{
							Symbol:     "user_groups_user_id",
							Columns:    []*Column{c4[0]},
							RefTable:   t1,
							RefColumns: []*Column{c1[0]},
						},
					}
				}

				return []*Table{t1, t2, t3, t4}
			},
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
				tablesMatch(t, drv, tables, tt.tables(drv))
			})
		}
	}
}

func tablesMatch(t *testing.T, drv string, got, expected []*Table) {
	require.Equal(t, len(expected), len(got))
	for i := range got {
		columnsMatch(t, drv, got[i].Columns, expected[i].Columns)
		columnsMatch(t, drv, got[i].PrimaryKey, expected[i].PrimaryKey)
		foreignKeysMatch(t, drv, got[i].ForeignKeys, expected[i].ForeignKeys)
	}
}

func columnsMatch(t *testing.T, drv string, got, expected []*Column) {
	require.Equal(t, len(expected), len(got))
	for i := range got {
		c1, c2 := got[i], expected[i]
		require.Equal(t, c2.Name, c1.Name)
		require.Equal(t, c2.Nullable, c1.Nullable)
		require.True(t, c1.Type == c2.Type || c1.ConvertibleTo(c2), "mismatched types: %s - %s", c1.Type, c2.Type)
		if c2.SchemaType[drv] != "" {
			require.Equal(t, c2.SchemaType[drv], c1.SchemaType[drv])
		}
	}
}

func foreignKeysMatch(t *testing.T, drv string, expected []*ForeignKey, got []*ForeignKey) {
	require.Equal(t, len(expected), len(got))
	for i := range got {
		fk1, fk2 := got[i], expected[i]
		require.Equal(t, fk2.Symbol, fk1.Symbol)
		require.Equal(t, fk2.RefTable.Name, fk1.RefTable.Name)
		columnsMatch(t, drv, fk1.Columns, fk2.Columns)
		columnsMatch(t, drv, fk1.RefColumns, fk2.RefColumns)
	}
}
