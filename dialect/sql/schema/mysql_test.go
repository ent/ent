// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"math"
	"regexp"
	"strings"
	"testing"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestMySQL_Create(t *testing.T) {
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
			name: "no tables",
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
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
						{Name: "enums", Type: field.TypeEnum, Enums: []string{"a", "b"}},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.8"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `age` bigint NOT NULL, `doc` json NULL, `enums` enum('a', 'b') NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table 5.6",
			tables: []*Table{
				{
					Name: "users",
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
						{Name: "name", Type: field.TypeString, Unique: true},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.6.35"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `age` bigint NOT NULL, `name` varchar(191) UNIQUE NOT NULL, `doc` longblob NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` timestamp NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("pets").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` bigint NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `CONSTRAINT_TYPE` = ? AND `CONSTRAINT_NAME` = ?")).
					WithArgs("FOREIGN KEY", "pets_owner").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("ALTER TABLE `pets` ADD CONSTRAINT `pets_owner` FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE")).
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
						{Name: "age", Type: field.TypeInt},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "").
						AddRow("text", "longtext", "YES", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` bigint NOT NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "enums",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "enums1", Type: field.TypeEnum, Enums: []string{"a", "b"}}, // add enum.
						{Name: "enums2", Type: field.TypeEnum, Enums: []string{"a"}},      // remove enum.
						{Name: "enums3", Type: field.TypeEnum, Enums: []string{"a", "b"}}, // order does not effect.
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "").
						AddRow("enums1", "enum('a')", "YES", "NO", "NULL", "", "", "").
						AddRow("enums2", "enum('b', 'a')", "NO", "YES", "NULL", "", "", "").
						AddRow("enums3", "enum('b', 'a')", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `enums1` enum('a', 'b') NOT NULL, MODIFY COLUMN `enums2` enum('a') NOT NULL")).
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
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("created_at", "datetime", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.6.0"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "").
						AddRow("doc", "longblob", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` bigint NOT NULL DEFAULT 10")).
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
						{Name: "tiny", Type: field.TypeBytes, Size: 100},
						{Name: "blob", Type: field.TypeBytes, Size: 1e3},
						{Name: "medium", Type: field.TypeBytes, Size: 1e5},
						{Name: "long", Type: field.TypeBytes, Size: 1e8},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `tiny` tinyblob NOT NULL, ADD COLUMN `blob` blob NOT NULL, ADD COLUMN `medium` mediumblob NOT NULL, ADD COLUMN `long` longblob NOT NULL")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec("ALTER TABLE `users` ADD COLUMN `age` double NOT NULL DEFAULT 10.1").
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec("ALTER TABLE `users` ADD COLUMN `age` boolean NOT NULL DEFAULT 1").
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `nick` varchar(255) NOT NULL DEFAULT \"unknown\"")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add column with unsupported default value to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "nick", Type: field.TypeString, Size: 1 << 17, Default: "unknown"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `nick` longtext NOT NULL")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` DROP COLUMN `name`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "modify column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "age", Type: field.TypeInt},
						{Name: "name", Type: field.TypeString, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "").
						AddRow("age", "bigint(20)", "NO", "NO", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("age", "bigint(20)", "NO", "", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				// create the unique index.
				mock.ExpectExec(escape("CREATE UNIQUE INDEX `age` ON `users`(`age`)")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("age", "bigint(20)", "NO", "UNI", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1").
						AddRow("age", "age", "0", "1"))
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("age", "bigint(20)", "NO", "UNI", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1").
						AddRow("age", "age", "0", "1"))
				// drop the unique index.
				mock.ExpectExec(escape("DROP INDEX `age` ON `users`")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "").
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `spouse_id` bigint")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLE_CONSTRAINTS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `CONSTRAINT_TYPE` = ? AND `CONSTRAINT_NAME` = ?")).
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec("ALTER TABLE `users` ADD CONSTRAINT `.{64}` FOREIGN KEY\\(`spouse_id`\\) REFERENCES `users`\\(`id`\\) ON DELETE CASCADE").
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("ent_types").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				// create ent_types table.
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `ent_types`(`id` bigint AUTO_INCREMENT NOT NULL, `type` varchar(255) UNIQUE NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `groups` AUTO_INCREMENT = 4294967296")).
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
			before: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("ent_types").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				// users table has no changes.
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name` FROM INFORMATION_SCHEMA.COLUMNS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", ""))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `non_unique`, `seq_in_index` FROM INFORMATION_SCHEMA.STATISTICS WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", "0", "1"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `groups` AUTO_INCREMENT = 4294967296")).
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
				mock.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
					WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", "5.7.23"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("ent_types").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(1))
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range (without inserting to ent_types).
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM INFORMATION_SCHEMA.TABLES WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set groups id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("groups").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `groups` AUTO_INCREMENT = 4294967296")).
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
			migrate, err := NewMigrate(sql.OpenDB("mysql", db), tt.options...)
			require.NoError(t, err)
			err = migrate.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

func escape(query string) string {
	rows := strings.Split(query, "\n")
	for i := range rows {
		rows[i] = strings.TrimPrefix(rows[i], " ")
	}
	query = strings.Join(rows, " ")
	return regexp.QuoteMeta(query)
}
