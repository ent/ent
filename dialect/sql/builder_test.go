// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/facebook/ent/dialect"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		input     Querier
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			input:     Describe("users"),
			wantQuery: "DESCRIBE `users`",
		},
		{
			input: CreateTable("users").
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("name").Type("varchar(255)"),
				).
				PrimaryKey("id"),
			wantQuery: "CREATE TABLE `users`(`id` int auto_increment, `name` varchar(255), PRIMARY KEY(`id`))",
		},
		{
			input: Dialect(dialect.Postgres).CreateTable("users").
				Columns(
					Column("id").Type("serial").Attr("PRIMARY KEY"),
					Column("name").Type("varchar"),
				),
			wantQuery: `CREATE TABLE "users"("id" serial PRIMARY KEY, "name" varchar)`,
		},
		{
			input: CreateTable("users").
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("name").Type("varchar(255)"),
				).
				PrimaryKey("id").
				Charset("utf8mb4"),
			wantQuery: "CREATE TABLE `users`(`id` int auto_increment, `name` varchar(255), PRIMARY KEY(`id`)) CHARACTER SET utf8mb4",
		},
		{
			input: CreateTable("users").
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("name").Type("varchar(255)"),
				).
				PrimaryKey("id").
				Charset("utf8mb4").
				Collate("utf8mb4_general_ci"),
			wantQuery: "CREATE TABLE `users`(`id` int auto_increment, `name` varchar(255), PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci",
		},
		{
			input: CreateTable("users").
				IfNotExists().
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
				).
				PrimaryKey("id", "name"),
			wantQuery: "CREATE TABLE IF NOT EXISTS `users`(`id` int auto_increment, PRIMARY KEY(`id`, `name`))",
		},
		{
			input: CreateTable("users").
				IfNotExists().
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("card_id").Type("int"),
				).
				PrimaryKey("id", "name").
				ForeignKeys(ForeignKey().Columns("card_id").
					Reference(Reference().Table("cards").Columns("id")).OnDelete("SET NULL")),
			wantQuery: "CREATE TABLE IF NOT EXISTS `users`(`id` int auto_increment, `card_id` int, PRIMARY KEY(`id`, `name`), FOREIGN KEY(`card_id`) REFERENCES `cards`(`id`) ON DELETE SET NULL)",
		},
		{
			input: Dialect(dialect.Postgres).CreateTable("users").
				IfNotExists().
				Columns(
					Column("id").Type("serial"),
					Column("card_id").Type("int"),
				).
				PrimaryKey("id", "name").
				ForeignKeys(ForeignKey().Columns("card_id").
					Reference(Reference().Table("cards").Columns("id")).OnDelete("SET NULL")),
			wantQuery: `CREATE TABLE IF NOT EXISTS "users"("id" serial, "card_id" int, PRIMARY KEY("id", "name"), FOREIGN KEY("card_id") REFERENCES "cards"("id") ON DELETE SET NULL)`,
		},
		{
			input: AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")).
					OnDelete("CASCADE"),
				),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `group_id` int UNIQUE, ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`) ON DELETE CASCADE",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey("constraint").Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")).
					OnDelete("CASCADE"),
				),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "group_id" int UNIQUE, ADD CONSTRAINT "constraint" FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE`,
		},
		{
			input: AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")),
				),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `group_id` int UNIQUE, ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`)",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")),
				),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "group_id" int UNIQUE, ADD CONSTRAINT FOREIGN KEY("group_id") REFERENCES "groups"("id")`,
		},
		{
			input: AlterTable("users").
				AddColumn(Column("age").Type("int")).
				AddColumn(Column("name").Type("varchar(255)")),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `age` int, ADD COLUMN `name` varchar(255)",
		},
		{
			input: AlterTable("users").
				DropForeignKey("users_parent_id"),
			wantQuery: "ALTER TABLE `users` DROP FOREIGN KEY `users_parent_id`",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("age").Type("int")).
				AddColumn(Column("name").Type("varchar(255)")).
				DropConstraint("users_nickname_key"),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "age" int, ADD COLUMN "name" varchar(255), DROP CONSTRAINT "users_nickname_key"`,
		},
		{
			input: AlterTable("users").
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")),
				).
				AddForeignKey(ForeignKey().Columns("location_id").
					Reference(Reference().Table("locations").Columns("id")),
				),
			wantQuery: "ALTER TABLE `users` ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`), ADD CONSTRAINT FOREIGN KEY(`location_id`) REFERENCES `locations`(`id`)",
		},
		{
			input: AlterTable("users").
				ModifyColumn(Column("age").Type("int")),
			wantQuery: "ALTER TABLE `users` MODIFY COLUMN `age` int",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int`,
		},
		{
			input: AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: "ALTER TABLE `users` MODIFY COLUMN `age` int, DROP COLUMN `name`",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int, DROP COLUMN "name"`,
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				ModifyColumn(Column("age").Attr("SET NOT NULL")).
				ModifyColumn(Column("name").Attr("DROP NOT NULL")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int, ALTER COLUMN "age" SET NOT NULL, ALTER COLUMN "name" DROP NOT NULL`,
		},
		{
			input: AlterTable("users").
				ChangeColumn("old_age", Column("age").Type("int")),
			wantQuery: "ALTER TABLE `users` CHANGE COLUMN `old_age` `age` int",
		},
		{
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("boring").Type("varchar")).
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "boring" varchar, ALTER COLUMN "age" TYPE int, DROP COLUMN "name"`,
		},
		{
			input:     AlterTable("users").RenameIndex("old", "new"),
			wantQuery: "ALTER TABLE `users` RENAME INDEX `old` TO `new`",
		},
		{
			input: AlterTable("users").
				DropIndex("old").
				AddIndex(CreateIndex("new1").Columns("c1", "c2")).
				AddIndex(CreateIndex("new2").Columns("c1", "c2").Unique()),
			wantQuery: "ALTER TABLE `users` DROP INDEX `old`, ADD INDEX `new1`(`c1`, `c2`), ADD UNIQUE INDEX `new2`(`c1`, `c2`)",
		},
		{
			input: Dialect(dialect.Postgres).AlterIndex("old").
				Rename("new"),
			wantQuery: `ALTER INDEX "old" RENAME TO "new"`,
		},
		{
			input:     Insert("users").Columns("age").Values(1),
			wantQuery: "INSERT INTO `users` (`age`) VALUES (?)",
			wantArgs:  []interface{}{1},
		},
		{
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1)`,
			wantArgs:  []interface{}{1},
		},
		{
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1).Returning("id"),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1) RETURNING "id"`,
			wantArgs:  []interface{}{1},
		},
		{
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1).Returning("id").Returning("name"),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1) RETURNING "name"`,
			wantArgs:  []interface{}{1},
		},
		{
			input:     Insert("users").Columns("name", "age").Values("a8m", 10),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)",
			wantArgs:  []interface{}{"a8m", 10},
		},
		{
			input:     Dialect(dialect.Postgres).Insert("users").Columns("name", "age").Values("a8m", 10),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2)`,
			wantArgs:  []interface{}{"a8m", 10},
		},
		{
			input:     Insert("users").Columns("name", "age").Values("a8m", 10).Values("foo", 20),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?), (?, ?)",
			wantArgs:  []interface{}{"a8m", 10, "foo", 20},
		},
		{
			input:     Dialect(dialect.Postgres).Insert("users").Columns("name", "age").Values("a8m", 10).Values("foo", 20),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2), ($3, $4)`,
			wantArgs:  []interface{}{"a8m", 10, "foo", 20},
		},
		{
			input: Dialect(dialect.Postgres).Insert("users").
				Columns("name", "age").
				Values("a8m", 10).
				Values("foo", 20).
				Values("bar", 30),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2), ($3, $4), ($5, $6)`,
			wantArgs:  []interface{}{"a8m", 10, "foo", 20, "bar", 30},
		},
		{
			input:     Update("users").Set("name", "foo"),
			wantQuery: "UPDATE `users` SET `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo"),
			wantQuery: `UPDATE "users" SET "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			input:     Update("users").Set("name", "foo").Set("age", 10),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ?",
			wantArgs:  []interface{}{"foo", 10},
		},
		{
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").Set("age", 10),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2`,
			wantArgs:  []interface{}{"foo", 10},
		},
		{
			input:     Update("users").Set("name", "foo").Where(EQ("name", "bar")),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").Where(EQ("name", "bar")),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input:     Update("users").Set("name", "foo").SetNull("spouse_id"),
			wantQuery: "UPDATE `users` SET `spouse_id` = NULL, `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").SetNull("spouse_id"),
			wantQuery: `UPDATE "users" SET "spouse_id" = NULL, "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			input: Update("users").Set("name", "foo").
				Where(EQ("name", "bar")).
				Where(EQ("age", 20)),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ? AND `age` = ?",
			wantArgs:  []interface{}{"foo", "bar", 20},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(EQ("name", "bar")).
				Where(EQ("age", 20)),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" = $2 AND "age" = $3`,
			wantArgs:  []interface{}{"foo", "bar", 20},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(Or(EQ("name", "bar"), EQ("name", "baz"))),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", 10, "bar", "baz"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(Or(EQ("name", "bar"), EQ("name", "baz"))),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2 WHERE "name" = $3 OR "name" = $4`,
			wantArgs:  []interface{}{"foo", 10, "bar", "baz"},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(P().EQ("name", "foo")),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", 10, "foo"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(P().EQ("name", "foo")),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2 WHERE "name" = $3`,
			wantArgs:  []interface{}{"foo", 10, "foo"},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Where(And(In("name", "bar", "baz"), NotIn("age", 1, 2))),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` IN (?, ?) AND `age` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar", "baz", 1, 2},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(And(In("name", "bar", "baz"), NotIn("age", 1, 2))),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" IN ($2, $3) AND "age" NOT IN ($4, $5)`,
			wantArgs:  []interface{}{"foo", "bar", "baz", 1, 2},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Where(And(HasPrefix("nickname", "a8m"), Contains("lastname", "mash"))),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `nickname` LIKE ? AND `lastname` LIKE ?",
			wantArgs:  []interface{}{"foo", "a8m%", "%mash%"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(And(HasPrefix("nickname", "a8m"), Contains("lastname", "mash"))),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "nickname" LIKE $2 AND "lastname" LIKE $3`,
			wantArgs:  []interface{}{"foo", "a8m%", "%mash%"},
		},
		{
			input: Update("users").
				Add("age", 1).
				Where(HasPrefix("nickname", "a8m")),
			wantQuery: "UPDATE `users` SET `age` = COALESCE(`age`, ?) + ? WHERE `nickname` LIKE ?",
			wantArgs:  []interface{}{0, 1, "a8m%"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Add("age", 1).
				Where(HasPrefix("nickname", "a8m")),
			wantQuery: `UPDATE "users" SET "age" = COALESCE("age", $1) + $2 WHERE "nickname" LIKE $3`,
			wantArgs:  []interface{}{0, 1, "a8m%"},
		},
		{
			input: Update("users").
				Add("age", 1).
				Set("nickname", "a8m").
				Add("version", 10).
				Set("name", "mashraki"),
			wantQuery: "UPDATE `users` SET `age` = COALESCE(`age`, ?) + ?, `nickname` = ?, `version` = COALESCE(`version`, ?) + ?, `name` = ?",
			wantArgs:  []interface{}{0, 1, "a8m", 0, 10, "mashraki"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Add("age", 1).
				Set("nickname", "a8m").
				Add("version", 10).
				Set("name", "mashraki"),
			wantQuery: `UPDATE "users" SET "age" = COALESCE("age", $1) + $2, "nickname" = $3, "version" = COALESCE("version", $4) + $5, "name" = $6`,
			wantArgs:  []interface{}{0, 1, "a8m", 0, 10, "mashraki"},
		},
		{
			input: Dialect(dialect.Postgres).
				Update("users").
				Add("age", 1).
				Set("nickname", "a8m").
				Add("version", 10).
				Set("name", "mashraki").
				Set("first", "ariel").
				Add("score", 1e5).
				Where(Or(EQ("age", 1), EQ("age", 2))),
			wantQuery: `UPDATE "users" SET "age" = COALESCE("age", $1) + $2, "nickname" = $3, "version" = COALESCE("version", $4) + $5, "name" = $6, "first" = $7, "score" = COALESCE("score", $8) + $9 WHERE "age" = $10 OR "age" = $11`,
			wantArgs:  []interface{}{0, 1, "a8m", 0, 10, "mashraki", "ariel", 0, 1e5, 1, 2},
		},
		{
			input: Select().
				From(Table("users")).
				Where(EQ("name", "Alex")),
			wantQuery: "SELECT * FROM `users` WHERE `name` = ?",
			wantArgs:  []interface{}{"Alex"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")),
			wantQuery: `SELECT * FROM "users"`,
		},
		{
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(EQ("name", "Ariel")),
			wantQuery: `SELECT * FROM "users" WHERE "name" = $1`,
			wantArgs:  []interface{}{"Ariel"},
		},
		{
			input: Select().
				From(Table("users")).
				Where(Or(EQ("name", "BAR"), EQ("name", "BAZ"))),
			wantQuery: "SELECT * FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"BAR", "BAZ"},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(And(EQ("name", "foo"), EQ("age", 20))),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? AND `age` = ?",
			wantArgs:  []interface{}{"foo", 10, "foo", 20},
		},
		{
			input: Delete("users").
				Where(NotNull("parent_id")),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL",
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(IsNull("parent_id")),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NULL`,
		},
		{
			input: Delete("users").
				Where(And(IsNull("parent_id"), NotIn("name", "foo", "bar"))),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NULL AND `name` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(And(IsNull("parent_id"), NotIn("name", "foo", "bar"))),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NULL AND "name" NOT IN ($1, $2)`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Delete("users").
				Where(And(False(), False())),
			wantQuery: "DELETE FROM `users` WHERE FALSE AND FALSE",
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(And(False(), False())),
			wantQuery: `DELETE FROM "users" WHERE FALSE AND FALSE`,
		},
		{
			input: Delete("users").
				Where(Or(NotNull("parent_id"), EQ("parent_id", 10))),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL OR `parent_id` = ?",
			wantArgs:  []interface{}{10},
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(Or(NotNull("parent_id"), EQ("parent_id", 10))),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NOT NULL OR "parent_id" = $1`,
			wantArgs:  []interface{}{10},
		},
		{
			input: Delete("users").
				Where(
					Or(
						And(EQ("name", "foo"), EQ("age", 10)),
						And(EQ("name", "bar"), EQ("age", 20)),
						And(
							EQ("name", "qux"),
							Or(EQ("age", 1), EQ("age", 2)),
						),
					),
				),
			wantQuery: "DELETE FROM `users` WHERE (`name` = ? AND `age` = ?) OR (`name` = ? AND `age` = ?) OR (`name` = ? AND (`age` = ? OR `age` = ?))",
			wantArgs:  []interface{}{"foo", 10, "bar", 20, "qux", 1, 2},
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(
					Or(
						And(EQ("name", "foo"), EQ("age", 10)),
						And(EQ("name", "bar"), EQ("age", 20)),
						And(
							EQ("name", "qux"),
							Or(EQ("age", 1), EQ("age", 2)),
						),
					),
				),
			wantQuery: `DELETE FROM "users" WHERE ("name" = $1 AND "age" = $2) OR ("name" = $3 AND "age" = $4) OR ("name" = $5 AND ("age" = $6 OR "age" = $7))`,
			wantArgs:  []interface{}{"foo", 10, "bar", 20, "qux", 1, 2},
		},
		{
			input: Delete("users").
				Where(
					Or(
						And(EQ("name", "foo"), EQ("age", 10)),
						And(EQ("name", "bar"), EQ("age", 20)),
					),
				).
				Where(EQ("role", "admin")),
			wantQuery: "DELETE FROM `users` WHERE ((`name` = ? AND `age` = ?) OR (`name` = ? AND `age` = ?)) AND `role` = ?",
			wantArgs:  []interface{}{"foo", 10, "bar", 20, "admin"},
		},
		{
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(
					Or(
						And(EQ("name", "foo"), EQ("age", 10)),
						And(EQ("name", "bar"), EQ("age", 20)),
					),
				).
				Where(EQ("role", "admin")),
			wantQuery: `DELETE FROM "users" WHERE (("name" = $1 AND "age" = $2) OR ("name" = $3 AND "age" = $4)) AND "role" = $5`,
			wantArgs:  []interface{}{"foo", 10, "bar", 20, "admin"},
		},
		{
			input:     Select().From(Table("users")),
			wantQuery: "SELECT * FROM `users`",
		},
		{
			input:     Dialect(dialect.Postgres).Select().From(Table("users")),
			wantQuery: `SELECT * FROM "users"`,
		},
		{
			input:     Select().From(Table("users").Unquote()),
			wantQuery: "SELECT * FROM users",
		},
		{
			input:     Dialect(dialect.Postgres).Select().From(Table("users").Unquote()),
			wantQuery: "SELECT * FROM users",
		},
		{
			input:     Select().From(Table("users").As("u")),
			wantQuery: "SELECT * FROM `users` AS `u`",
		},
		{
			input:     Dialect(dialect.Postgres).Select().From(Table("users").As("u")),
			wantQuery: `SELECT * FROM "users" AS "u"`,
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Select(t1.C("id"), t2.C("name")).From(t1).Join(t2)
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN `groups` AS `g`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Dialect(dialect.Postgres).Select(t1.C("id"), t2.C("name")).From(t1).Join(t2)
			}(),
			wantQuery: `SELECT "u"."id", "g"."name" FROM "users" AS "u" JOIN "groups" AS "g"`,
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id"))
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN `groups` AS `g` ON `u`.`id` = `g`.`user_id`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Dialect(dialect.Postgres).
					Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id"))
			}(),
			wantQuery: `SELECT "u"."id", "g"."name" FROM "users" AS "u" JOIN "groups" AS "g" ON "u"."id" = "g"."user_id"`,
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id")).
					Where(And(EQ(t1.C("name"), "bar"), NotNull(t2.C("name"))))
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN `groups` AS `g` ON `u`.`id` = `g`.`user_id` WHERE `u`.`name` = ? AND `g`.`name` IS NOT NULL",
			wantArgs:  []interface{}{"bar"},
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Dialect(dialect.Postgres).
					Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id")).
					Where(And(EQ(t1.C("name"), "bar"), NotNull(t2.C("name"))))
			}(),
			wantQuery: `SELECT "u"."id", "g"."name" FROM "users" AS "u" JOIN "groups" AS "g" ON "u"."id" = "g"."user_id" WHERE "u"."name" = $1 AND "g"."name" IS NOT NULL`,
			wantArgs:  []interface{}{"bar"},
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("user_groups").As("ug")
				return Select(t1.C("id"), As(Count("`*`"), "group_count")).
					From(t1).
					LeftJoin(t2).
					On(t1.C("id"), t2.C("user_id")).
					GroupBy(t1.C("id"))
			}(),
			wantQuery: "SELECT `u`.`id`, COUNT(`*`) AS `group_count` FROM `users` AS `u` LEFT JOIN `user_groups` AS `ug` ON `u`.`id` = `ug`.`user_id` GROUP BY `u`.`id`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("user_groups").As("ug")
				return Select(t1.C("id"), As(Count("`*`"), "group_count")).
					From(t1).
					LeftJoin(t2).
					OnP(P(func(builder *Builder) {
						builder.Ident(t1.C("id")).WriteOp(OpEQ).Ident(t2.C("user_id"))
					})).
					GroupBy(t1.C("id")).Clone()
			}(),
			wantQuery: "SELECT `u`.`id`, COUNT(`*`) AS `group_count` FROM `users` AS `u` LEFT JOIN `user_groups` AS `ug` ON `u`.`id` = `ug`.`user_id` GROUP BY `u`.`id`",
		},
		{
			input: func() Querier {
				t1 := Table("groups").As("g")
				t2 := Table("user_groups").As("ug")
				return Select(t1.C("id"), As(Count("`*`"), "user_count")).
					From(t1).
					RightJoin(t2).
					On(t1.C("id"), t2.C("group_id")).
					GroupBy(t1.C("id"))
			}(),
			wantQuery: "SELECT `g`.`id`, COUNT(`*`) AS `user_count` FROM `groups` AS `g` RIGHT JOIN `user_groups` AS `ug` ON `g`.`id` = `ug`.`group_id` GROUP BY `g`.`id`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				return Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: "SELECT `u`.`name`, `u`.`age` FROM `users` AS `u`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: `SELECT "u"."name", "u"."age" FROM "users" AS "u"`,
		},
		{
			input: func() Querier {
				t1 := Dialect(dialect.Postgres).
					Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: `SELECT "u"."name", "u"."age" FROM "users" AS "u"`,
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Select().From(Table("groups")).Where(EQ("user_id", 10)).As("g")
				return Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id"))
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN (SELECT * FROM `groups` WHERE `user_id` = ?) AS `g` ON `u`.`id` = `g`.`user_id`",
			wantArgs:  []interface{}{10},
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				t1 := d.Table("users").As("u")
				t2 := d.Select().From(Table("groups")).Where(EQ("user_id", 10)).As("g")
				return d.Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id"))
			}(),
			wantQuery: `SELECT "u"."id", "g"."name" FROM "users" AS "u" JOIN (SELECT * FROM "groups" WHERE "user_id" = $1) AS "g" ON "u"."id" = "g"."user_id"`,
			wantArgs:  []interface{}{10},
		},
		{
			input: func() Querier {
				selector := Select().Where(Or(EQ("name", "foo"), EQ("name", "bar")))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select().Where(Or(EQ("name", "foo"), EQ("name", "bar")))
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "users" WHERE "name" = $1 OR "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				selector := Select().From(Table("users")).As("t")
				return selector.Select(selector.C("name"))
			}(),
			wantQuery: "SELECT `t`.`name` FROM `users`",
		},
		{
			input: func() Querier {
				selector := Dialect(dialect.Postgres).
					Select().From(Table("users")).As("t")
				return selector.Select(selector.C("name"))
			}(),
			wantQuery: `SELECT "t"."name" FROM "users"`,
		},
		{
			input: func() Querier {
				selector := Select().From(Table("groups")).Where(EQ("name", "foo"))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `groups` WHERE `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select().From(Table("groups")).Where(EQ("name", "foo"))
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "groups" WHERE "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			input: func() Querier {
				selector := Select()
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `users`",
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select()
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "users"`,
		},
		{
			input: Select().
				From(Table("users")).
				Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))),
			wantQuery: "SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))),
			wantQuery: `SELECT * FROM "users" WHERE NOT ("name" = $1 AND "age" = $2)`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Select().
				From(Table("users")).
				Where(Or(EqualFold("name", "BAR"), EqualFold("name", "BAZ"))),
			wantQuery: "SELECT * FROM `users` WHERE LOWER(`name`) = ? OR LOWER(`name`) = ?",
			wantArgs:  []interface{}{"bar", "baz"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(Or(EqualFold("name", "BAR"), EqualFold("name", "BAZ"))),
			wantQuery: `SELECT * FROM "users" WHERE LOWER("name") = $1 OR LOWER("name") = $2`,
			wantArgs:  []interface{}{"bar", "baz"},
		},
		{
			input: Dialect(dialect.SQLite).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: "SELECT * FROM `users` WHERE LOWER(`name`) LIKE ? AND LOWER(`nick`) LIKE ?",
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: `SELECT * FROM "users" WHERE "name" ILIKE $1 AND "nick" ILIKE $2`,
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			input: Dialect(dialect.MySQL).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: "SELECT * FROM `users` WHERE `name` COLLATE utf8mb4_general_ci LIKE ? AND `nick` COLLATE utf8mb4_general_ci LIKE ?",
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			input: func() Querier {
				s1 := Select().
					From(Table("users")).
					Where(Not(And(EQ("name", "foo"), EQ("age", "bar"))))
				return Queries{With("users_view").As(s1), Select("name").From(Table("users_view"))}
			}(),
			wantQuery: "WITH users_view AS (SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)) SELECT `name` FROM `users_view`",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				s1 := d.Select().
					From(Table("users")).
					Where(Not(And(EQ("name", "foo"), EQ("age", "bar"))))
				return Queries{d.With("users_view").As(s1), d.Select("name").From(Table("users_view"))}
			}(),
			wantQuery: `WITH users_view AS (SELECT * FROM "users" WHERE NOT ("name" = $1 AND "age" = $2)) SELECT "name" FROM "users_view"`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				s1 := Select().From(Table("users")).Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))).As("users_view")
				return Select("name").From(s1)
			}(),
			wantQuery: "SELECT `name` FROM (SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)) AS `users_view`",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				s1 := d.Select().From(Table("users")).Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))).As("users_view")
				return d.Select("name").From(s1)
			}(),
			wantQuery: `SELECT "name" FROM (SELECT * FROM "users" WHERE NOT ("name" = $1 AND "age" = $2)) AS "users_view"`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Select().
					From(t1).
					Where(In(t1.C("id"), Select("owner_id").From(Table("pets")).Where(EQ("name", "pedro"))))
			}(),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `owner_id` FROM `pets` WHERE `name` = ?)",
			wantArgs:  []interface{}{"pedro"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(t1).
					Where(In(t1.C("id"), Select("owner_id").From(Table("pets")).Where(EQ("name", "pedro"))))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE "users"."id" IN (SELECT "owner_id" FROM "pets" WHERE "name" = $1)`,
			wantArgs:  []interface{}{"pedro"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Select().
					From(t1).
					Where(Not(In(t1.C("id"), Select("owner_id").From(Table("pets")).Where(EQ("name", "pedro")))))
			}(),
			wantQuery: "SELECT * FROM `users` WHERE NOT (`users`.`id` IN (SELECT `owner_id` FROM `pets` WHERE `name` = ?))",
			wantArgs:  []interface{}{"pedro"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(t1).
					Where(Not(In(t1.C("id"), Select("owner_id").From(Table("pets")).Where(EQ("name", "pedro")))))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE NOT ("users"."id" IN (SELECT "owner_id" FROM "pets" WHERE "name" = $1))`,
			wantArgs:  []interface{}{"pedro"},
		},
		{
			input:     Select().Count().From(Table("users")),
			wantQuery: "SELECT COUNT(*) FROM `users`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select().Count().From(Table("users")),
			wantQuery: `SELECT COUNT(*) FROM "users"`,
		},
		{
			input:     Select().Count(Distinct("id")).From(Table("users")),
			wantQuery: "SELECT COUNT(DISTINCT `id`) FROM `users`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select().Count(Distinct("id")).From(Table("users")),
			wantQuery: `SELECT COUNT(DISTINCT "id") FROM "users"`,
		},
		{
			input: func() Querier {
				t1 := Table("users")
				t2 := Select().From(Table("groups"))
				t3 := Select().Count().From(t1).Join(t1).On(t2.C("id"), t1.C("blocked_id"))
				return t3.Count(Distinct(t3.Columns("id", "name")...))
			}(),
			wantQuery: "SELECT COUNT(DISTINCT `t0`.`id`, `t0`.`name`) FROM `users` AS `t0` JOIN `users` AS `t0` ON `groups`.`id` = `t0`.`blocked_id`",
		},
		{
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				t1 := d.Table("users")
				t2 := d.Select().From(Table("groups"))
				t3 := d.Select().Count().From(t1).Join(t1).On(t2.C("id"), t1.C("blocked_id"))
				return t3.Count(Distinct(t3.Columns("id", "name")...))
			}(),
			wantQuery: `SELECT COUNT(DISTINCT "t0"."id", "t0"."name") FROM "users" AS "t0" JOIN "users" AS "t0" ON "groups"."id" = "t0"."blocked_id"`,
		},
		{
			input:     Select(Sum("age"), Min("age")).From(Table("users")),
			wantQuery: "SELECT SUM(`age`), MIN(`age`) FROM `users`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select(Sum("age"), Min("age")).
				From(Table("users")),
			wantQuery: `SELECT SUM("age"), MIN("age") FROM "users"`,
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				return Select(As(Max(t1.C("age")), "max_age")).From(t1)
			}(),
			wantQuery: "SELECT MAX(`u`.`age`) AS `max_age` FROM `users` AS `u`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(As(Max(t1.C("age")), "max_age")).
					From(t1)
			}(),
			wantQuery: `SELECT MAX("u"."age") AS "max_age" FROM "users" AS "u"`,
		},
		{
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name"),
			wantQuery: `SELECT "name", COUNT(*) FROM "users" GROUP BY "name"`,
		},
		{
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name").
				OrderBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name` ORDER BY `name`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name").
				OrderBy("name"),
			wantQuery: `SELECT "name", COUNT(*) FROM "users" GROUP BY "name" ORDER BY "name"`,
		},
		{
			input: Select("name", "age", Count("*")).
				From(Table("users")).
				GroupBy("name", "age").
				OrderBy(Desc("name"), "age"),
			wantQuery: "SELECT `name`, `age`, COUNT(*) FROM `users` GROUP BY `name`, `age` ORDER BY `name` DESC, `age`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("name", "age", Count("*")).
				From(Table("users")).
				GroupBy("name", "age").
				OrderBy(Desc("name"), "age"),
			wantQuery: `SELECT "name", "age", COUNT(*) FROM "users" GROUP BY "name", "age" ORDER BY "name" DESC, "age"`,
		},
		{
			input: Select("*").
				From(Table("users")).
				Limit(1),
			wantQuery: "SELECT * FROM `users` LIMIT 1",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("users")).
				Limit(1),
			wantQuery: `SELECT * FROM "users" LIMIT 1`,
		},
		{
			input:     Select("age").Distinct().From(Table("users")),
			wantQuery: "SELECT DISTINCT `age` FROM `users`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("age").
				Distinct().
				From(Table("users")),
			wantQuery: `SELECT DISTINCT "age" FROM "users"`,
		},
		{
			input:     Select("age", "name").From(Table("users")).Distinct().OrderBy("name"),
			wantQuery: "SELECT DISTINCT `age`, `name` FROM `users` ORDER BY `name`",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("age", "name").
				From(Table("users")).
				Distinct().
				OrderBy("name"),
			wantQuery: `SELECT DISTINCT "age", "name" FROM "users" ORDER BY "name"`,
		},
		{
			input:     Select("age").From(Table("users")).Where(EQ("name", "foo")).Or().Where(EQ("name", "bar")),
			wantQuery: "SELECT `age` FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select("age").
				From(Table("users")).
				Where(EQ("name", "foo")).Or().Where(EQ("name", "bar")),
			wantQuery: `SELECT "age" FROM "users" WHERE "name" = $1 OR "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input:     Queries{With("users_view").As(Select().From(Table("users"))), Select().From(Table("users_view"))},
			wantQuery: "WITH users_view AS (SELECT * FROM `users`) SELECT * FROM `users_view`",
		},
		{
			input: func() Querier {
				base := Select("*").From(Table("groups"))
				return Queries{With("groups").As(base.Clone().Where(EQ("name", "bar"))), base.Select("age")}
			}(),
			wantQuery: "WITH groups AS (SELECT * FROM `groups` WHERE `name` = ?) SELECT `age` FROM `groups`",
			wantArgs:  []interface{}{"bar"},
		},
		{
			input: func() Querier {
				builder := Dialect(dialect.Postgres)
				t1 := builder.Table("groups")
				t2 := builder.Table("users")
				t3 := builder.Table("user_groups")
				t4 := builder.Select(t3.C("id")).
					From(t3).
					Join(t2).
					On(t3.C("id"), t2.C("id2")).
					Where(EQ(t2.C("id"), "baz"))
				return builder.Select().
					From(t1).
					Join(t4).
					On(t1.C("id"), t4.C("id")).Limit(1)
			}(),
			wantQuery: `SELECT * FROM "groups" JOIN (SELECT "user_groups"."id" FROM "user_groups" JOIN "users" AS "t0" ON "user_groups"."id" = "t0"."id2" WHERE "t0"."id" = $1) AS "t1" ON "groups"."id" = "t1"."id" LIMIT 1`,
			wantArgs:  []interface{}{"baz"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(t1).
					Where(CompositeGT(t1.Columns("id", "name"), 1, "Ariel"))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE ("users"."id", "users"."name") > ($1, $2)`,
			wantArgs:  []interface{}{1, "Ariel"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(t1).
					Where(And(EQ("name", "Ariel"), CompositeGT(t1.Columns("id", "name"), 1, "Ariel")))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE "name" = $1 AND ("users"."id", "users"."name") > ($2, $3)`,
			wantArgs:  []interface{}{"Ariel", 1, "Ariel"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(t1).
					Where(And(EQ("name", "Ariel"), Or(EQ("surname", "Doe"), CompositeGT(t1.Columns("id", "name"), 1, "Ariel"))))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE "name" = $1 AND ("surname" = $2 OR ("users"."id", "users"."name") > ($3, $4))`,
			wantArgs:  []interface{}{"Ariel", "Doe", 1, "Ariel"},
		},
		{
			input: func() Querier {
				t1 := Table("users")
				return Dialect(dialect.Postgres).
					Select().
					From(Table("users")).
					Where(And(EQ("name", "Ariel"), CompositeLT(t1.Columns("id", "name"), 1, "Ariel")))
			}(),
			wantQuery: `SELECT * FROM "users" WHERE "name" = $1 AND ("users"."id", "users"."name") < ($2, $3)`,
			wantArgs:  []interface{}{"Ariel", 1, "Ariel"},
		},
		{
			input:     CreateIndex("name_index").Table("users").Column("name"),
			wantQuery: "CREATE INDEX `name_index` ON `users`(`name`)",
		},
		{
			input: Dialect(dialect.Postgres).
				CreateIndex("name_index").
				Table("users").
				Column("name"),
			wantQuery: `CREATE INDEX "name_index" ON "users"("name")`,
		},
		{
			input:     CreateIndex("unique_name").Unique().Table("users").Columns("first", "last"),
			wantQuery: "CREATE UNIQUE INDEX `unique_name` ON `users`(`first`, `last`)",
		},
		{
			input: Dialect(dialect.Postgres).
				CreateIndex("unique_name").
				Unique().
				Table("users").
				Columns("first", "last"),
			wantQuery: `CREATE UNIQUE INDEX "unique_name" ON "users"("first", "last")`,
		},
		{
			input:     DropIndex("name_index"),
			wantQuery: "DROP INDEX `name_index`",
		},
		{
			input: Dialect(dialect.Postgres).
				DropIndex("name_index"),
			wantQuery: `DROP INDEX "name_index"`,
		},
		{
			input:     DropIndex("name_index").Table("users"),
			wantQuery: "DROP INDEX `name_index` ON `users`",
		},
		{
			input: Select().
				From(Table("pragma_table_info('t1')").Unquote()).
				OrderBy("pk"),
			wantQuery: "SELECT * FROM pragma_table_info('t1') ORDER BY `pk`",
		},
		{
			input: AlterTable("users").
				AddColumn(Column("spouse").Type("integer").
					Constraint(ForeignKey("user_spouse").
						Reference(Reference().Table("users").Columns("id")).
						OnDelete("SET NULL"))),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `spouse` integer CONSTRAINT user_spouse REFERENCES `users`(`id`) ON DELETE SET NULL",
		},
		{
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("users")).
				Where(Or(
					And(EQ("id", 1), InInts("group_id", 2, 3)),
					And(EQ("id", 2), InValues("group_id", 4, 5)),
				)).
				Where(And(
					Or(EQ("a", "a"), And(EQ("b", "b"), EQ("c", "c"))),
					Not(Or(IsNull("d"), NotNull("e"))),
				)).
				Or().
				Where(And(NEQ("f", "f"), NEQ("g", "g"))),
			wantQuery: strings.NewReplacer("\n", "", "\t", "").Replace(`
SELECT * FROM "users" 
WHERE 
	(
		(("id" = $1 AND "group_id" IN ($2, $3)) OR ("id" = $4 AND "group_id" IN ($5, $6))) 
		AND 
		(("a" = $7 OR ("b" = $8 AND "c" = $9)) AND (NOT ("d" IS NULL OR "e" IS NOT NULL)))
	) 
	OR ("f" <> $10 AND "g" <> $11)`),
			wantArgs: []interface{}{1, 2, 3, 2, 4, 5, "a", "b", "c", "f", "g"},
		},
		{
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("test")).
				Where(P(func(b *Builder) {
					b.WriteString("nlevel(").Ident("path").WriteByte(')').WriteOp(OpGT).Arg(1)
				})),
			wantQuery: `SELECT * FROM "test" WHERE nlevel("path") > $1`,
			wantArgs:  []interface{}{1},
		},
		{
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("test")).
				Where(P(func(b *Builder) {
					b.WriteString("nlevel(").Ident("path").WriteByte(')').WriteOp(OpGT).Arg(1)
				})),
			wantQuery: `SELECT * FROM "test" WHERE nlevel("path") > $1`,
			wantArgs:  []interface{}{1},
		},
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}

func TestBuilder_Err(t *testing.T) {
	b := Select("i-")
	require.NoError(t, b.Err())
	b.AddError(fmt.Errorf("invalid"))
	require.EqualError(t, b.Err(), "invalid")
	b.AddError(fmt.Errorf("unexpected"))
	require.EqualError(t, b.Err(), "invalid; unexpected")
}
