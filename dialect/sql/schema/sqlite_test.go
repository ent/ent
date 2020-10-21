// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"math"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestSQLite_Create(t *testing.T) {
	tests := []struct {
		name    string
		tables  []*Table
		options []MigrateOption
		before  func(sqliteMock)
		wantErr bool
	}{
		{
			name: "tx failed",
			before: func(mock sqliteMock) {
				mock.ExpectBegin().WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
		{
			name: "fk disabled",
			before: func(mock sqliteMock) {
				mock.ExpectBegin()
				mock.ExpectQuery("PRAGMA foreign_keys").
					WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(0))
				mock.ExpectRollback()
			},
			wantErr: true,
		},
		{
			name: "no tables",
			before: func(mock sqliteMock) {
				mock.start()
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
						{Name: "decimal", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.SQLite: "decimal(6,2)"}},
					},
				},
			},
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `age` integer NOT NULL, `doc` json NULL, `uuid` uuid NULL, `decimal` decimal(6,2) NOT NULL)")).
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
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` datetime NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape("CREATE TABLE `pets`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` integer NULL, FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE)")).
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
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` datetime NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape("CREATE TABLE `pets`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` integer NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
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
						{Name: "text", Type: field.TypeString, Nullable: true, Size: math.MaxInt32},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "age", Type: field.TypeInt, Default: 0},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
						AddRow("name", "varchar(255)", 0, nil, 0).
						AddRow("text", "text", 0, "NULL", 0).
						AddRow("uuid", "uuid", 0, "Null", 0).
						AddRow("id", "integer", 1, "NULL", 1))
				mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "origin"}))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` integer NOT NULL DEFAULT 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "datetime and timestamp",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "created_at", Type: field.TypeTime, Nullable: true},
						{Name: "updated_at", Type: field.TypeTime, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
						AddRow("created_at", "datetime", 0, nil, 0).
						AddRow("id", "integer", 1, "NULL", 1))
				mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "origin"}))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `updated_at` datetime NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add blob columns",
			tables: []*Table{
				{
					Name: "blobs",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "old_tiny", Type: field.TypeBytes, Size: 100},
						{Name: "old_blob", Type: field.TypeBytes, Size: 1e3},
						{Name: "old_medium", Type: field.TypeBytes, Size: 1e5},
						{Name: "old_long", Type: field.TypeBytes, Size: 1e8},
						{Name: "new_tiny", Type: field.TypeBytes, Size: 100},
						{Name: "new_blob", Type: field.TypeBytes, Size: 1e3},
						{Name: "new_medium", Type: field.TypeBytes, Size: 1e5},
						{Name: "new_long", Type: field.TypeBytes, Size: 1e8},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("blobs", true)
				mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('blobs') ORDER BY `pk`")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
						AddRow("old_tiny", "blob", 1, nil, 0).
						AddRow("old_blob", "blob", 1, nil, 0).
						AddRow("old_medium", "blob", 1, nil, 0).
						AddRow("old_long", "blob", 1, nil, 0).
						AddRow("id", "integer", 1, "NULL", 1))
				mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('blobs')")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "unique"}))
				for _, c := range []string{"tiny", "blob", "medium", "long"} {
					mock.ExpectExec(escape(fmt.Sprintf("ALTER TABLE `blobs` ADD COLUMN `new_%s` blob NOT NULL", c))).
						WillReturnResult(sqlmock.NewResult(0, 1))
				}
				mock.ExpectCommit()
			},
		},
		{
			name: "add columns with default values",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Default: "unknown"},
						{Name: "active", Type: field.TypeBool, Default: false},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
						AddRow("id", "integer", 1, "NULL", 1))
				mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "origin"}))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `name` varchar(255) NOT NULL DEFAULT 'unknown'")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `active` bool NOT NULL DEFAULT false")).
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
								Symbol:     "user_spouse",
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
			before: func(mock sqliteMock) {
				mock.start()
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `name`, `type`, `notnull`, `dflt_value`, `pk` FROM pragma_table_info('users') ORDER BY `pk`")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "type", "notnull", "dflt_value", "pk"}).
						AddRow("name", "varchar(255)", 1, "NULL", 0).
						AddRow("id", "integer", 1, "NULL", 1))
				mock.ExpectQuery(escape("SELECT `name`, `unique`, `origin` FROM pragma_index_list('users')")).
					WithArgs().
					WillReturnRows(sqlmock.NewRows([]string{"name", "unique", "origin"}))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `spouse_id` integer NULL CONSTRAINT user_spouse REFERENCES `users`(`id`) ON DELETE CASCADE")).
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
			before: func(mock sqliteMock) {
				mock.start()
				// creating ent_types table.
				mock.tableExists("ent_types", false)
				mock.ExpectExec(escape("CREATE TABLE `ent_types`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `type` varchar(255) UNIQUE NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("users", false)
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
				mock.tableExists("groups", false)
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
			before: func(mock sqliteMock) {
				mock.start()
				// query ent_types table.
				mock.tableExists("ent_types", true)
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range (without inserting to ent_types).
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_sequence` WHERE `name` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectExec(escape("UPDATE `sqlite_sequence` SET `seq` = ? WHERE `name` = ?")).
					WithArgs(0, "users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("groups", false)
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
			tt.before(sqliteMock{mock})
			migrate, err := NewMigrate(sql.OpenDB("sqlite3", db), tt.options...)
			require.NoError(t, err)
			err = migrate.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

type sqliteMock struct {
	sqlmock.Sqlmock
}

func (m sqliteMock) start() {
	m.ExpectBegin()
	m.ExpectQuery("PRAGMA foreign_keys").
		WillReturnRows(sqlmock.NewRows([]string{"foreign_keys"}).AddRow(1))
}

func (m sqliteMock) tableExists(table string, exists bool) {
	count := 0
	if exists {
		count = 1
	}
	m.ExpectQuery(escape("SELECT COUNT(*) FROM `sqlite_master` WHERE `type` = ? AND `name` = ?")).
		WithArgs("table", table).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}
