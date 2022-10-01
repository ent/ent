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

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestMySQL_Create(t *testing.T) {
	tests := []struct {
		name    string
		tables  []*Table
		options []MigrateOption
		before  func(mysqlMock)
		wantErr bool
	}{
		{
			name: "tx failed",
			before: func(mock mysqlMock) {
				mock.ExpectBegin().
					WillReturnError(sqlmock.ErrCancelled)
			},
			wantErr: true,
		},
		{
			name: "no tables",
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
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
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "ts", Type: field.TypeTime},
						{Name: "ts_default", Type: field.TypeTime, Default: "CURRENT_TIMESTAMP"},
						{Name: "datetime", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "datetime"}, Default: "CURRENT_TIMESTAMP"},
						{Name: "decimal", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.MySQL: "decimal(6,2)"}},
						{Name: "unsigned decimal", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.MySQL: "decimal(6,2) unsigned"}},
						{Name: "float", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.MySQL: "float"}, Default: "0"},
					},
					Annotation: &entsql.Annotation{
						Charset:   "utf8",
						Collation: "utf8_general_ci",
						Options:   "ENGINE = INNODB",
						Check:     "price > 0",
						Checks: map[string]string{
							"valid_age":  "age > 0",
							"valid_name": "name <> ''",
						},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.8")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `age` bigint NOT NULL, `doc` json NULL, `enums` enum('a', 'b') NOT NULL, `uuid` char(36) binary NULL, `ts` timestamp NULL, `ts_default` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP, `datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, `decimal` decimal(6,2) NOT NULL, `unsigned decimal` decimal(6,2) unsigned NOT NULL, `float` float NOT NULL DEFAULT '0', PRIMARY KEY(`id`), CHECK (price > 0), CONSTRAINT `valid_age` CHECK (age > 0), CONSTRAINT `valid_name` CHECK (name <> '')) CHARACTER SET utf8 COLLATE utf8_general_ci ENGINE = INNODB")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create new table with specific field collation",
			tables: []*Table{
				{
					Name: "users",
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "address", Type: field.TypeString, Nullable: true, Collation: "utf8_unicode_ci"},
						{Name: "age", Type: field.TypeInt},
						{Name: "doc", Type: field.TypeJSON, Nullable: true},
						{Name: "enums", Type: field.TypeEnum, Enums: []string{"a", "b"}},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "datetime", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "datetime"}, Default: "CURRENT_TIMESTAMP"},
						{Name: "decimal", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.MySQL: "decimal(6,2)"}},
					},
					Annotation: &entsql.Annotation{
						Charset:   "utf8",
						Collation: "utf8_general_ci",
						Options:   "ENGINE = INNODB",
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.33")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `address` varchar(255) NULL COLLATE utf8_unicode_ci, `age` bigint NOT NULL, `doc` json NULL, `enums` enum('a', 'b') NOT NULL, `uuid` char(36) binary NULL, `datetime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP, `decimal` decimal(6,2) NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8 COLLATE utf8_general_ci ENGINE = INNODB")).
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
			before: func(mock mysqlMock) {
				mock.start("5.6.35")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `age` bigint NOT NULL, `name` varchar(191) UNIQUE NOT NULL, `doc` longblob NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` timestamp NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` bigint NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.fkExists("pets_owner", false)
				mock.ExpectExec(escape("ALTER TABLE `pets` ADD CONSTRAINT `pets_owner` FOREIGN KEY(`owner_id`) REFERENCES `users`(`id`) ON DELETE CASCADE")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NULL, `created_at` timestamp NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("pets", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `pets`(`id` bigint AUTO_INCREMENT NOT NULL, `name` varchar(255) NOT NULL, `owner_id` bigint NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add columns to table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "text", Type: field.TypeString, Nullable: true, Size: math.MaxInt32},
						{Name: "mediumtext", Type: field.TypeString, Nullable: true, SchemaType: map[string]string{dialect.MySQL: "mediumtext"}},
						{Name: "uuid", Type: field.TypeUUID, Nullable: true},
						{Name: "date", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{dialect.MySQL: "date"}},
						{Name: "age", Type: field.TypeInt},
						{Name: "tiny", Type: field.TypeInt8},
						{Name: "tiny_unsigned", Type: field.TypeUint8},
						{Name: "small", Type: field.TypeInt16},
						{Name: "small_unsigned", Type: field.TypeUint16},
						{Name: "big", Type: field.TypeInt64},
						{Name: "big_unsigned", Type: field.TypeUint64},
						{Name: "decimal", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.MySQL: "decimal(6,2)"}},
						{Name: "unsigned_decimal", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.MySQL: "decimal(6,2) unsigned"}},
						{Name: "ts", Type: field.TypeTime},
						{Name: "timestamp", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "TIMESTAMP"}, Default: "CURRENT_TIMESTAMP"},
						{Name: "float", Type: field.TypeFloat32, SchemaType: map[string]string{dialect.MySQL: "float"}, Default: "0"},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("8.0.19")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("text", "longtext", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("mediumtext", "mediumtext", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("uuid", "char(36)", "YES", "YES", "NULL", "", "", "utf8mb4_bin", nil, nil).
						AddRow("date", "date", "YES", "YES", "NULL", "", "", "", nil, nil).
						// 8.0.19: new int column type formats
						AddRow("tiny", "tinyint", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("tiny_unsigned", "tinyint unsigned", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("small", "smallint", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("small_unsigned", "smallint unsigned", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("big", "bigint", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("big_unsigned", "bigint unsigned", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("decimal", "decimal(6,2)", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("unsigned_decimal", "decimal(6,2) unsigned", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("timestamp", "timestamp", "NO", "NO", "CURRENT_TIMESTAMP", "DEFAULT_GENERATED on update CURRENT_TIMESTAMP", "", "", nil, nil).
						AddRow("float", "float", "NO", "NO", "0", "0", "", "", nil, nil))

				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `age` bigint NOT NULL, ADD COLUMN `ts` timestamp NOT NULL, MODIFY COLUMN `timestamp` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP")).
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
						{Name: "enums1", Type: field.TypeEnum, Enums: []string{"a", "b"}},   // add enum.
						{Name: "enums2", Type: field.TypeEnum, Enums: []string{"a"}},        // remove enum.
						{Name: "enums3", Type: field.TypeEnum, Enums: []string{"a", "b c"}}, // no changes.
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("enums1", "enum('a')", "YES", "NO", "NULL", "", "", "", nil, nil).
						AddRow("enums2", "enum('b', 'a')", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("enums3", "enum('a', 'b c')", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
						{Name: "created_at", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "datetime"}, Nullable: true},
						{Name: "updated_at", Type: field.TypeTime, SchemaType: map[string]string{dialect.MySQL: "datetime"}, Nullable: true},
						{Name: "deleted_at", Type: field.TypeTime, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("created_at", "datetime", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("updated_at", "timestamp", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("deleted_at", "datetime", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `updated_at` datetime NULL, MODIFY COLUMN `deleted_at` timestamp NULL")).
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
			before: func(mock mysqlMock) {
				mock.start("5.6.0")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("doc", "longblob", "YES", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `tiny` tinyblob NOT NULL, ADD COLUMN `blob` blob NOT NULL, ADD COLUMN `medium` mediumblob NOT NULL, ADD COLUMN `long` longblob NOT NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add binary column",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "binary", Type: field.TypeBytes, Size: 20, SchemaType: map[string]string{dialect.MySQL: "binary(20)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("8.0.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `binary` binary(20) NOT NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "accept varbinary columns",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "tiny", Type: field.TypeBytes, Size: 100},
						{Name: "medium", Type: field.TypeBytes, Size: math.MaxUint32},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("tiny", "varbinary(255)", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("medium", "varbinary(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `medium` longblob NOT NULL")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec("ALTER TABLE `users` ADD COLUMN `age` double NOT NULL DEFAULT 10.1").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add bool column with default value",
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec("ALTER TABLE `users` ADD COLUMN `age` boolean NOT NULL DEFAULT true").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add string column with default value",
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `nick` varchar(255) NOT NULL DEFAULT 'unknown'")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "add column with unsupported default value",
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `nick` longtext NOT NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "drop columns",
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
						{Name: "age", Type: field.TypeInt},
						{Name: "name", Type: field.TypeString, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "NO", "YES", "NULL", "", "", "", nil, nil).
						AddRow("age", "bigint(20)", "NO", "NO", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("age", "bigint(20)", "NO", "", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("age", "bigint(20)", "NO", "UNI", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("age", "age", nil, "0", "1"))
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("age", "bigint(20)", "NO", "UNI", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`, `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("age", "age", nil, "0", "1"))
				// check if a foreign-key needs to be dropped.
				mock.ExpectQuery(escape("SELECT `CONSTRAINT_NAME` FROM `INFORMATION_SCHEMA`.`KEY_COLUMN_USAGE` WHERE `TABLE_NAME` = ? AND `COLUMN_NAME` = ? AND `POSITION_IN_UNIQUE_CONSTRAINT` IS NOT NULL AND `TABLE_SCHEMA` = (SELECT DATABASE())")).
					WithArgs("users", "age").
					WillReturnRows(sqlmock.NewRows([]string{"CONSTRAINT_NAME"}))
				// drop the unique index.
				mock.ExpectExec(escape("DROP INDEX `age` ON `users`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "increase index sub_part",
			tables: func() []*Table {
				t := &Table{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "text", Type: field.TypeString, Size: math.MaxInt32, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					Indexes: []*Index{
						{Name: "prefix_text", Annotation: &entsql.IndexAnnotation{Prefix: 100}},
					},
				}
				t.Indexes[0].Columns = t.Columns[1:]
				return []*Table{t}
			}(),
			options: []MigrateOption{WithDropIndex(true)},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("text", "longtext", "YES", "NO", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("prefix_text", "text", "50", "0", "1"))
				mock.ExpectQuery(escape("SELECT `CONSTRAINT_NAME` FROM `INFORMATION_SCHEMA`.`KEY_COLUMN_USAGE` WHERE `TABLE_NAME` = ? AND `COLUMN_NAME` = ? AND `POSITION_IN_UNIQUE_CONSTRAINT` IS NOT NULL AND `TABLE_SCHEMA` = (SELECT DATABASE())")).
					WithArgs("users", "text").
					WillReturnRows(sqlmock.NewRows([]string{"CONSTRAINT_NAME"}))
				// modify index by dropping and creating it.
				mock.ExpectExec(escape("DROP INDEX `prefix_text` ON `users`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("CREATE INDEX `prefix_text` ON `users`(`text`(100))")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "ignore foreign keys on index dropping",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "parent_id", Type: field.TypeInt, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
					ForeignKeys: []*ForeignKey{
						{
							Symbol: "parent_id",
							Columns: []*Column{
								{Name: "parent_id", Type: field.TypeInt, Nullable: true},
							},
						},
					},
				},
			},
			options: []MigrateOption{WithDropIndex(true)},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("parent_id", "bigint(20)", "YES", "NULL", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("old_index", "old", nil, "0", "1").
						AddRow("parent_id", "parent_id", nil, "0", "1"))
				// drop the unique index.
				mock.ExpectExec(escape("DROP INDEX `old_index` ON `users`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// foreign key already exist.
				mock.fkExists("parent_id", true)
				mock.ExpectCommit()
			},
		},
		{
			name: "drop foreign key with column and index",
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
			options: []MigrateOption{WithDropIndex(true), WithDropColumn(true)},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("parent_id", "bigint(20)", "YES", "NULL", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("parent_id", "parent_id", nil, "0", "1"))
				// check if a foreign-key needs to be dropped.
				mock.ExpectQuery(escape("SELECT `CONSTRAINT_NAME` FROM `INFORMATION_SCHEMA`.`KEY_COLUMN_USAGE` WHERE `TABLE_NAME` = ? AND `COLUMN_NAME` = ? AND `POSITION_IN_UNIQUE_CONSTRAINT` IS NOT NULL AND `TABLE_SCHEMA` = (SELECT DATABASE())")).
					WithArgs("users", "parent_id").
					WillReturnRows(sqlmock.NewRows([]string{"CONSTRAINT_NAME"}).AddRow("users_parent_id"))
				mock.ExpectExec(escape("ALTER TABLE `users` DROP FOREIGN KEY `users_parent_id`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// drop the unique index.
				mock.ExpectExec(escape("DROP INDEX `parent_id` ON `users`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// drop the unique index.
				mock.ExpectExec(escape("ALTER TABLE `users` DROP COLUMN `parent_id`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "create a new simple-index for the foreign-key",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "parent_id", Type: field.TypeInt, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			options: []MigrateOption{WithDropIndex(true), WithDropColumn(true)},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("parent_id", "bigint(20)", "YES", "NULL", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1").
						AddRow("parent_id", "parent_id", nil, "0", "1"))
				// check if there's a foreign-key that is associated with this index.
				mock.ExpectQuery(escape("SELECT `CONSTRAINT_NAME` FROM `INFORMATION_SCHEMA`.`KEY_COLUMN_USAGE` WHERE `TABLE_NAME` = ? AND `COLUMN_NAME` = ? AND `POSITION_IN_UNIQUE_CONSTRAINT` IS NOT NULL AND `TABLE_SCHEMA` = (SELECT DATABASE())")).
					WithArgs("users", "parent_id").
					WillReturnRows(sqlmock.NewRows([]string{"CONSTRAINT_NAME"}).AddRow("users_parent_id"))
				// create a new index, to replace the old one (that needs to be dropped).
				mock.ExpectExec(escape("CREATE INDEX `users_parent_id` ON `users`(`parent_id`)")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// drop the unique index.
				mock.ExpectExec(escape("DROP INDEX `parent_id` ON `users`")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` ADD COLUMN `spouse_id` bigint NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.fkExists("user_spouse_____________________390ed76f91d3c57cd3516e7690f621dc", false)
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("ent_types", false)
				// create ent_types table.
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `ent_types`(`id` bigint unsigned AUTO_INCREMENT NOT NULL, `type` varchar(255) UNIQUE NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("ent_types", true)
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.tableExists("users", true)
				// users table has no changes.
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				// query groups table.
				mock.ExpectQuery(escape("SELECT COUNT(*) FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("groups").
					WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
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
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("ent_types", true)
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).AddRow("users"))
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range (without inserting to ent_types).
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("groups", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `groups`(`id` bigint AUTO_INCREMENT NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
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
			name: "universal id mismatch with ent_types",
			tables: []*Table{
				NewTable("users").AddPrimary(&Column{Name: "id", Type: field.TypeInt, Increment: true}),
			},
			options: []MigrateOption{WithGlobalUniqueID(true)},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("ent_types", true)
				// query ent_types table.
				mock.ExpectQuery(escape("SELECT `type` FROM `ent_types` ORDER BY `id` ASC")).
					WillReturnRows(sqlmock.NewRows([]string{"type"}).
						AddRow("deleted").
						AddRow("users"))
				mock.tableExists("users", true)
				// users table has no changes.
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				// query the auto-increment value.
				mock.ExpectQuery(escape("SELECT `AUTO_INCREMENT` FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"AUTO_INCREMENT"}).
						AddRow(1))
				// restore the auto-increment counter.
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 4294967296")).
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
						{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.MySQL: "decimal(6,4)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("price", "decimal(6,4)", "NO", "YES", "NULL", "", "", "", "6", "4"))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
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
						{Name: "price", Type: field.TypeFloat64, SchemaType: map[string]string{dialect.MySQL: "decimal(6,4)"}},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("5.7.23")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("price", "decimal(6,4)", "NO", "YES", "NULL", "", "", "", "5", "4"))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectExec(escape("ALTER TABLE `users` MODIFY COLUMN `price` decimal(6,4) NOT NULL")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		// MariaDB specific tests.
		{
			name: "mariadb/10.2.32/create table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "json", Type: field.TypeJSON, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("10.2.32-MariaDB")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `json` json NULL CHECK (JSON_VALID(`json`)), PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "mariadb/10.3.13/create table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "json", Type: field.TypeJSON, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("10.3.13-MariaDB-1:10.3.13+maria~bionic")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `json` json NULL CHECK (JSON_VALID(`json`)), PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "mariadb/10.5.8/create table",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "json", Type: field.TypeJSON, Nullable: true},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("10.5.8-MariaDB")
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `json` json NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
		{
			name: "mariadb/10.5.8/table exists",
			tables: []*Table{
				{
					Name: "users",
					Columns: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
						{Name: "name", Type: field.TypeString, Nullable: true},
						{Name: "json", Type: field.TypeJSON, Nullable: true},
						{Name: "longtext", Type: field.TypeString, Nullable: true, Size: math.MaxInt32},
					},
					PrimaryKey: []*Column{
						{Name: "id", Type: field.TypeInt, Increment: true},
					},
				},
			},
			before: func(mock mysqlMock) {
				mock.start("10.5.8-MariaDB-1:10.5.8+maria~focal")
				mock.tableExists("users", true)
				mock.ExpectQuery(escape("SELECT `column_name`, `column_type`, `is_nullable`, `column_key`, `column_default`, `extra`, `character_set_name`, `collation_name`, `numeric_precision`, `numeric_scale` FROM `INFORMATION_SCHEMA`.`COLUMNS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"column_name", "column_type", "is_nullable", "column_key", "column_default", "extra", "character_set_name", "collation_name", "numeric_precision", "numeric_scale"}).
						AddRow("id", "bigint(20)", "NO", "PRI", "NULL", "auto_increment", "", "", nil, nil).
						AddRow("name", "varchar(255)", "YES", "YES", "NULL", "", "", "", nil, nil).
						AddRow("json", "longtext", "YES", "YES", "NULL", "", "utf8mb4", "utf8mb4_bin", nil, nil).
						AddRow("longtext", "longtext", "YES", "YES", "NULL", "", "utf8mb4", "utf8mb4_bin", nil, nil))
				mock.ExpectQuery(escape("SELECT `index_name`, `column_name`, `sub_part`,  `non_unique`, `seq_in_index` FROM `INFORMATION_SCHEMA`.`STATISTICS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? ORDER BY `index_name`, `seq_in_index`")).
					WithArgs("users").
					WillReturnRows(sqlmock.NewRows([]string{"index_name", "column_name", "sub_part", "non_unique", "seq_in_index"}).
						AddRow("PRIMARY", "id", nil, "0", "1"))
				mock.ExpectQuery(escape("SELECT `CONSTRAINT_NAME` FROM `INFORMATION_SCHEMA`.`CHECK_CONSTRAINTS` WHERE `CONSTRAINT_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ? AND `CHECK_CLAUSE` LIKE ?")).
					WithArgs("users", "json_valid(%)").
					WillReturnRows(sqlmock.NewRows([]string{"CONSTRAINT_NAME"}).
						AddRow("json"))
				mock.ExpectCommit()
			},
		},
		{
			name: "mariadb/10.1.37/create table",
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
					},
				},
			},
			options: []MigrateOption{WithGlobalUniqueID(true)},
			before: func(mock mysqlMock) {
				mock.start("10.1.48-MariaDB-1~bionic")
				mock.tableExists("ent_types", false)
				// create ent_types table.
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `ent_types`(`id` bigint unsigned AUTO_INCREMENT NOT NULL, `type` varchar(191) UNIQUE NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.tableExists("users", false)
				mock.ExpectExec(escape("CREATE TABLE IF NOT EXISTS `users`(`id` bigint AUTO_INCREMENT NOT NULL, `age` bigint NOT NULL, `name` varchar(191) UNIQUE NOT NULL, PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// set users id range.
				mock.ExpectExec(escape("INSERT INTO `ent_types` (`type`) VALUES (?)")).
					WithArgs("users").
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectExec(escape("ALTER TABLE `users` AUTO_INCREMENT = 0")).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.before(mysqlMock{mock})
			migrate, err := NewMigrate(sql.OpenDB("mysql", db), append(tt.options, WithAtlas(false))...)
			require.NoError(t, err)
			err = migrate.Create(context.Background(), tt.tables...)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

type mysqlMock struct {
	sqlmock.Sqlmock
}

func (m mysqlMock) start(version string) {
	m.ExpectQuery(escape("SHOW VARIABLES LIKE 'version'")).
		WillReturnRows(sqlmock.NewRows([]string{"Variable_name", "Value"}).AddRow("version", version))
	m.ExpectBegin()
}

func (m mysqlMock) tableExists(table string, exists bool) {
	count := 0
	if exists {
		count = 1
	}
	m.ExpectQuery(escape("SELECT COUNT(*) FROM `INFORMATION_SCHEMA`.`TABLES` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `TABLE_NAME` = ?")).
		WithArgs(table).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func (m mysqlMock) fkExists(fk string, exists bool) {
	count := 0
	if exists {
		count = 1
	}
	m.ExpectQuery(escape("SELECT COUNT(*) FROM `INFORMATION_SCHEMA`.`TABLE_CONSTRAINTS` WHERE `TABLE_SCHEMA` = (SELECT DATABASE()) AND `CONSTRAINT_TYPE` = ? AND `CONSTRAINT_NAME` = ?")).
		WithArgs("FOREIGN KEY", fk).
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(count))
}

func escape(query string) string {
	rows := strings.Split(query, "\n")
	for i := range rows {
		rows[i] = strings.TrimPrefix(rows[i], " ")
	}
	query = strings.Join(rows, " ")
	return strings.TrimSpace(regexp.QuoteMeta(query)) + "$"
}
