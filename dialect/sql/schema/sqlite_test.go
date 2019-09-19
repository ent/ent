// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"testing"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

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
				mock.ExpectExec(escape("CREATE TABLE `users`(`id` integer PRIMARY KEY AUTOINCREMENT NOT NULL, `name` varchar(255) NULL, `age` integer NOT NULL, `doc` json NULL)")).
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
