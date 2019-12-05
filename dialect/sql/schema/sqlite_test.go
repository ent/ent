// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"math"
	"strings"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	stdsql "database/sql"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSQLite_DoubleCreate(t *testing.T) {
	tables := []*Table{
		{
			Name: "users",
			PrimaryKey: []*Column{
				{Name: "id", Type: field.TypeInt, Increment: true},
			},
			Columns: []*Column{
				{Name: "id", Type: field.TypeInt, Increment: true},
				{Name: "name", Type: field.TypeString, Nullable: true},
				{Name: "age", Type: field.TypeInt},
				{Name: "doc", Type: field.TypeJSON, Nullable: true},
				{Name: "uuid", Type: field.TypeUUID, Nullable: true},
			},
		},
	}

	for i := 0; i < 2; i++ {
		db, err := stdsql.Open("sqlite3", "file:test?mode=memory&cache=shared&_fk=1")
		require.NoError(t, err)
		migrate, err := NewMigrate(sql.OpenDB("sqlite3", db))
		require.NoError(t, err)
		err = migrate.Create(context.Background(), tables...)
		require.NoError(t, err)
	}
}

func TestSQLite_Create(t *testing.T) {
	tests := []struct {
		name    string
		tables  []*Table
		options []MigrateOption
		before  func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "tx failed",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
		{
			name: "fk disabled",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(0))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "no tables",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
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
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "age", Type: field.TypeInt},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `age` integer NOT NULL, `doc` json NULL, `uuid` uuid NULL)")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` datetime NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "pets").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `pets`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` integer NULL, FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("pragma foreign_key_list('pets')")).
					WillReturnRows(sqlmock.NewRows([]string{"id", "seq", "table", "from", "to", "on_update", "on_delete", "match"}).
						AddRow(0, 0, "users", "owner_id", "id", "NO ACTION", "CASCADE", "NONE"))
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
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0).
						AddRow(2, "uuid", "uuid", 0, nil, 0).
						AddRow(3, "text", "varchar(2147483647)", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` integer NOT NULL")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0).
						AddRow(2, "doc", "json", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` integer NOT NULL DEFAULT 10")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0).
						AddRow(2, "doc", "json", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `blob` blob NOT NULL, ADD COLUMN `longblob` blob NOT NULL")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` real NOT NULL DEFAULT 10.1")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` bool NOT NULL DEFAULT true")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `nick` varchar(255) NOT NULL DEFAULT 'unknown'")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` DROP COLUMN `name`")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 1, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `name` varchar(255) NULL")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 1, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` integer UNIQUE NOT NULL")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "age", "bigint", 1, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0).
						AddRow(1, "sqlite_autoindex_users_2", 1, "u", 0))
				mock.ExpectQuery(escape("pragma index_info('sqlite_autoindex_users_2')")).
					WillReturnRows(sqlmock.NewRows([]string{"seqno", "cid", "name"}).
						AddRow(0, 1, "age"))
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "age", "bigint", 1, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0).
						AddRow(1, "sqlite_autoindex_users_2", 1, "u", 0))
				mock.ExpectQuery(escape("pragma index_info('sqlite_autoindex_users_2')")).
					WillReturnRows(sqlmock.NewRows([]string{"seqno", "cid", "name"}).
						AddRow(0, 1, "age"))
				mock.ExpectExec(escape("DROP INDEX `sqlite_autoindex_users_2`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape(`pragma table_info('users')`)).
					WillReturnRows(sqlmock.NewRows([]string{"id", "name", "type", "not_null", "dflt", "pk"}).
						AddRow(0, "id", "bigint", 1, nil, 1).
						AddRow(1, "name", "varchar", 0, nil, 0))
				mock.ExpectQuery(escape("pragma index_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"seq", "name", "unique", "origin", "partial"}).
						AddRow(0, "sqlite_autoindex_users_1", 1, "pk", 0))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `spouse_id` integer NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("pragma foreign_key_list('users')")).
					WillReturnRows(sqlmock.NewRows([]string{"id", "seq", "table", "from", "to", "on_update", "on_delete", "match"}))
				mock.ExpectExec("ALTER TABLE `users` ADD COLUMN `spouse_id` REFERENCES `users`\\(`id`\\) ON DELETE CASCADE").
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				// creating ent_types table.
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "ent_types").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `ent_types`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `type` varchar(255) UNIQUE NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_sequence` WHERE `name` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("INSERT INTO `sqlite_sequence` (`name`, `seq`) VALUES (?, ?)")).
					WithArgs("users", 0).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `groups`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_sequence` WHERE `name` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("INSERT INTO `sqlite_sequence` (`name`, `seq`) VALUES (?, ?)")).
					WithArgs("groups", 1<<32).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "ent_types").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range (without inserting to ent_types).
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_sequence` WHERE `name` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectExec(escape("UPDATE `sqlite_sequence` SET `seq` = ? WHERE `name` = ?")).
					WithArgs(0, "users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
					WithArgs("table", "groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE `groups`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_sequence` WHERE `name` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("INSERT INTO `sqlite_sequence` (`name`, `seq`) VALUES (?, ?)")).
					WithArgs("groups", 1<<32).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.before(mock)
			migrate, err := NewMigrate(sql.OpenDB("sqlite3", db), tt.options...)
			require.NoError(t, err)
			err = migrate.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}
