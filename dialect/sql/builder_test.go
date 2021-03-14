// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"github.com/stretchr/testify/require"
)

func TestBuilder(t *testing.T) {
	tests := []struct {
		id        string
		input     Querier
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			id:        "example-1",
			input:     Describe("users"),
			wantQuery: "DESCRIBE `users`",
		},
		{
			id: "example-2",
			input: CreateTable("users").
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("name").Type("varchar(255)"),
				).
				PrimaryKey("id"),
			wantQuery: "CREATE TABLE `users`(`id` int auto_increment, `name` varchar(255), PRIMARY KEY(`id`))",
		},
		{
			id: "example-3",
			input: Dialect(dialect.Postgres).CreateTable("users").
				Columns(
					Column("id").Type("serial").Attr("PRIMARY KEY"),
					Column("name").Type("varchar"),
				),
			wantQuery: `CREATE TABLE "users"("id" serial PRIMARY KEY, "name" varchar)`,
		},
		{
			id: "example-4",
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
			id: "example-5",
			input: CreateTable("users").
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("name").Type("varchar(255)"),
				).
				PrimaryKey("id").
				Charset("utf8mb4").
				Collate("utf8mb4_general_ci").
				Options("ENGINE=InnoDB"),
			wantQuery: "CREATE TABLE `users`(`id` int auto_increment, `name` varchar(255), PRIMARY KEY(`id`)) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci ENGINE=InnoDB",
		},
		{
			id: "example-6",
			input: CreateTable("users").
				IfNotExists().
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
				).
				PrimaryKey("id", "name"),
			wantQuery: "CREATE TABLE IF NOT EXISTS `users`(`id` int auto_increment, PRIMARY KEY(`id`, `name`))",
		},
		{
			id: "example-7",
			input: CreateTable("users").
				IfNotExists().
				Columns(
					Column("id").Type("int").Attr("auto_increment"),
					Column("card_id").Type("int"),
					Column("doc").Type("longtext").Check(func(b *Builder) {
						b.WriteString("JSON_VALID(").Ident("doc").WriteByte(')')
					}),
				).
				PrimaryKey("id", "name").
				ForeignKeys(ForeignKey().Columns("card_id").
					Reference(Reference().Table("cards").Columns("id")).OnDelete("SET NULL")),
			wantQuery: "CREATE TABLE IF NOT EXISTS `users`(`id` int auto_increment, `card_id` int, `doc` longtext CHECK (JSON_VALID(`doc`)), PRIMARY KEY(`id`, `name`), FOREIGN KEY(`card_id`) REFERENCES `cards`(`id`) ON DELETE SET NULL)",
		},
		{
			id: "example-8",
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
			id: "example-9",
			input: AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")).
					OnDelete("CASCADE"),
				),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `group_id` int UNIQUE, ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`) ON DELETE CASCADE",
		},
		{
			id: "example-10",
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey("constraint").Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")).
					OnDelete("CASCADE"),
				),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "group_id" int UNIQUE, ADD CONSTRAINT "constraint" FOREIGN KEY("group_id") REFERENCES "groups"("id") ON DELETE CASCADE`,
		},
		{
			id: "example-11",
			input: AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")),
				),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `group_id` int UNIQUE, ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`)",
		},
		{
			id: "example-12",
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")),
				),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "group_id" int UNIQUE, ADD CONSTRAINT FOREIGN KEY("group_id") REFERENCES "groups"("id")`,
		},
		{
			id: "example-13",
			input: AlterTable("users").
				AddColumn(Column("age").Type("int")).
				AddColumn(Column("name").Type("varchar(255)")),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `age` int, ADD COLUMN `name` varchar(255)",
		},
		{
			id: "example-14",
			input: AlterTable("users").
				DropForeignKey("users_parent_id"),
			wantQuery: "ALTER TABLE `users` DROP FOREIGN KEY `users_parent_id`",
		},
		{
			id: "example-15",
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("age").Type("int")).
				AddColumn(Column("name").Type("varchar(255)")).
				DropConstraint("users_nickname_key"),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "age" int, ADD COLUMN "name" varchar(255), DROP CONSTRAINT "users_nickname_key"`,
		},
		{
			id: "example-16",
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
			id: "example-17",
			input: AlterTable("users").
				ModifyColumn(Column("age").Type("int")),
			wantQuery: "ALTER TABLE `users` MODIFY COLUMN `age` int",
		},
		{
			id: "example-18",
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int`,
		},
		{
			id: "example-19",
			input: AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: "ALTER TABLE `users` MODIFY COLUMN `age` int, DROP COLUMN `name`",
		},
		{
			id: "example-20",
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int, DROP COLUMN "name"`,
		},
		{
			id: "example-21",
			input: Dialect(dialect.Postgres).AlterTable("users").
				ModifyColumn(Column("age").Type("int")).
				ModifyColumn(Column("age").Attr("SET NOT NULL")).
				ModifyColumn(Column("name").Attr("DROP NOT NULL")),
			wantQuery: `ALTER TABLE "users" ALTER COLUMN "age" TYPE int, ALTER COLUMN "age" SET NOT NULL, ALTER COLUMN "name" DROP NOT NULL`,
		},
		{
			id: "example-22",
			input: AlterTable("users").
				ChangeColumn("old_age", Column("age").Type("int")),
			wantQuery: "ALTER TABLE `users` CHANGE COLUMN `old_age` `age` int",
		},
		{
			id: "example-23",
			input: Dialect(dialect.Postgres).AlterTable("users").
				AddColumn(Column("boring").Type("varchar")).
				ModifyColumn(Column("age").Type("int")).
				DropColumn(Column("name")),
			wantQuery: `ALTER TABLE "users" ADD COLUMN "boring" varchar, ALTER COLUMN "age" TYPE int, DROP COLUMN "name"`,
		},
		{
			id:        "example-24",
			input:     AlterTable("users").RenameIndex("old", "new"),
			wantQuery: "ALTER TABLE `users` RENAME INDEX `old` TO `new`",
		},
		{
			id: "example-25",
			input: AlterTable("users").
				DropIndex("old").
				AddIndex(CreateIndex("new1").Columns("c1", "c2")).
				AddIndex(CreateIndex("new2").Columns("c1", "c2").Unique()),
			wantQuery: "ALTER TABLE `users` DROP INDEX `old`, ADD INDEX `new1`(`c1`, `c2`), ADD UNIQUE INDEX `new2`(`c1`, `c2`)",
		},
		{
			id: "example-26",
			input: Dialect(dialect.Postgres).AlterIndex("old").
				Rename("new"),
			wantQuery: `ALTER INDEX "old" RENAME TO "new"`,
		},
		{
			id:        "example-27",
			input:     Insert("users").Columns("age").Values(1),
			wantQuery: "INSERT INTO `users` (`age`) VALUES (?)",
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-28",
			input:     Insert("users").Columns("age").Values(1).Schema("mydb"),
			wantQuery: "INSERT INTO `mydb`.`users` (`age`) VALUES (?)",
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-29",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1)`,
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-30",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1).Schema("mydb"),
			wantQuery: `INSERT INTO "mydb"."users" ("age") VALUES ($1)`,
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-31",
			input:     Dialect(dialect.SQLite).Insert("users").Columns("age").Values(1).Schema("mydb"),
			wantQuery: "INSERT INTO `users` (`age`) VALUES (?)",
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-32",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1).Returning("id"),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1) RETURNING "id"`,
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-33",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("age").Values(1).Returning("id").Returning("name"),
			wantQuery: `INSERT INTO "users" ("age") VALUES ($1) RETURNING "name"`,
			wantArgs:  []interface{}{1},
		},
		{
			id:        "example-34",
			input:     Insert("users").Columns("name", "age").Values("a8m", 10),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)",
			wantArgs:  []interface{}{"a8m", 10},
		},
		{
			id:        "example-35",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("name", "age").Values("a8m", 10),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2)`,
			wantArgs:  []interface{}{"a8m", 10},
		},
		{
			id:        "example-36",
			input:     Insert("users").Columns("name", "age").Values("a8m", 10).Values("foo", 20),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?), (?, ?)",
			wantArgs:  []interface{}{"a8m", 10, "foo", 20},
		},
		{
			id:        "example-37",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("name", "age").Values("a8m", 10).Values("foo", 20),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2), ($3, $4)`,
			wantArgs:  []interface{}{"a8m", 10, "foo", 20},
		},
		{
			id: "example-38",
			input: Dialect(dialect.Postgres).Insert("users").
				Columns("name", "age").
				Values("a8m", 10).
				Values("foo", 20).
				Values("bar", 30),
			wantQuery: `INSERT INTO "users" ("name", "age") VALUES ($1, $2), ($3, $4), ($5, $6)`,
			wantArgs:  []interface{}{"a8m", 10, "foo", 20, "bar", 30},
		},
		{
			id:        "example-39",
			input:     Update("users").Set("name", "foo"),
			wantQuery: "UPDATE `users` SET `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-40",
			input:     Update("users").Set("name", "foo").Schema("mydb"),
			wantQuery: "UPDATE `mydb`.`users` SET `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-41",
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo"),
			wantQuery: `UPDATE "users" SET "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-42",
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").Schema("mydb"),
			wantQuery: `UPDATE "mydb"."users" SET "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-43",
			input:     Dialect(dialect.SQLite).Update("users").Set("name", "foo").Schema("mydb"),
			wantQuery: "UPDATE `users` SET `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-44",
			input:     Update("users").Set("name", "foo").Set("age", 10),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ?",
			wantArgs:  []interface{}{"foo", 10},
		},
		{
			id:        "example-45",
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").Set("age", 10),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2`,
			wantArgs:  []interface{}{"foo", 10},
		},
		{
			id:        "example-46",
			input:     Update("users").Set("name", "foo").Where(EQ("name", "bar")),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id:        "example-47",
			input:     Update("users").Set("name", "foo").Where(EQ("name", Expr("?", "bar"))),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id:        "example-48",
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").Where(EQ("name", "bar")),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-49",
			input: func() Querier {
				p1, p2 := EQ("name", "bar"), Or(EQ("age", 10), EQ("age", 20))
				return Dialect(dialect.Postgres).
					Update("users").
					Set("name", "foo").
					Where(p1).
					Where(p2).
					Where(p1).
					Where(p2)
			}(),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE (("name" = $2 AND ("age" = $3 OR "age" = $4)) AND "name" = $5) AND ("age" = $6 OR "age" = $7)`,
			wantArgs:  []interface{}{"foo", "bar", 10, 20, "bar", 10, 20},
		},
		{
			id:        "example-50",
			input:     Update("users").Set("name", "foo").SetNull("spouse_id"),
			wantQuery: "UPDATE `users` SET `spouse_id` = NULL, `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			id:        "example-51",
			input:     Dialect(dialect.Postgres).Update("users").Set("name", "foo").SetNull("spouse_id"),
			wantQuery: `UPDATE "users" SET "spouse_id" = NULL, "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			id: "example-52",
			input: Update("users").Set("name", "foo").
				Where(EQ("name", "bar")).
				Where(EQ("age", 20)),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ? AND `age` = ?",
			wantArgs:  []interface{}{"foo", "bar", 20},
		},
		{
			id: "example-53",
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(EQ("name", "bar")).
				Where(EQ("age", 20)),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" = $2 AND "age" = $3`,
			wantArgs:  []interface{}{"foo", "bar", 20},
		},
		{
			id: "example-54",
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(Or(EQ("name", "bar"), EQ("name", "baz"))),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", 10, "bar", "baz"},
		},
		{
			id: "example-55",
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(Or(EQ("name", "bar"), EQ("name", "baz"))),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2 WHERE "name" = $3 OR "name" = $4`,
			wantArgs:  []interface{}{"foo", 10, "bar", "baz"},
		},
		{
			id: "example-56",
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(P().EQ("name", "foo")),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", 10, "foo"},
		},
		{
			id: "example-57",
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(P().EQ("name", "foo")),
			wantQuery: `UPDATE "users" SET "name" = $1, "age" = $2 WHERE "name" = $3`,
			wantArgs:  []interface{}{"foo", 10, "foo"},
		},
		{
			id: "example-58",
			input: Update("users").
				Set("name", "foo").
				Where(And(In("name", "bar", "baz"), NotIn("age", 1, 2))),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` IN (?, ?) AND `age` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar", "baz", 1, 2},
		},
		{
			id: "example-59",
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(And(In("name", "bar", "baz"), NotIn("age", 1, 2))),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "name" IN ($2, $3) AND "age" NOT IN ($4, $5)`,
			wantArgs:  []interface{}{"foo", "bar", "baz", 1, 2},
		},
		{
			id: "example-60",
			input: Update("users").
				Set("name", "foo").
				Where(And(HasPrefix("nickname", "a8m"), Contains("lastname", "mash"))),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `nickname` LIKE ? AND `lastname` LIKE ?",
			wantArgs:  []interface{}{"foo", "a8m%", "%mash%"},
		},
		{
			id: "example-61",
			input: Dialect(dialect.Postgres).
				Update("users").
				Set("name", "foo").
				Where(And(HasPrefix("nickname", "a8m"), Contains("lastname", "mash"))),
			wantQuery: `UPDATE "users" SET "name" = $1 WHERE "nickname" LIKE $2 AND "lastname" LIKE $3`,
			wantArgs:  []interface{}{"foo", "a8m%", "%mash%"},
		},
		{
			id: "example-62",
			input: Update("users").
				Add("age", 1).
				Where(HasPrefix("nickname", "a8m")),
			wantQuery: "UPDATE `users` SET `age` = COALESCE(`age`, ?) + ? WHERE `nickname` LIKE ?",
			wantArgs:  []interface{}{0, 1, "a8m%"},
		},
		{
			id: "example-63",
			input: Dialect(dialect.Postgres).
				Update("users").
				Add("age", 1).
				Where(HasPrefix("nickname", "a8m")),
			wantQuery: `UPDATE "users" SET "age" = COALESCE("age", $1) + $2 WHERE "nickname" LIKE $3`,
			wantArgs:  []interface{}{0, 1, "a8m%"},
		},
		{
			id: "example-64",
			input: Update("users").
				Add("age", 1).
				Set("nickname", "a8m").
				Add("version", 10).
				Set("name", "mashraki"),
			wantQuery: "UPDATE `users` SET `age` = COALESCE(`age`, ?) + ?, `nickname` = ?, `version` = COALESCE(`version`, ?) + ?, `name` = ?",
			wantArgs:  []interface{}{0, 1, "a8m", 0, 10, "mashraki"},
		},
		{
			id: "example-65",
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
			id: "example-66",
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
			id:        "example-67",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("id", "email").Values("1", "user@example.com").ConflictColumns("id").UpdateSet("email", "user-1@example.com"),
			wantQuery: `INSERT INTO "users" ("id", "email") VALUES ($1, $2) ON CONFLICT ("id") DO UPDATE SET "id" = "excluded"."id", "email" = "excluded"."email"`,
			wantArgs:  []interface{}{"1", "user@example.com"},
		},

		{
			id:        "example-68",
			input:     Dialect(dialect.Postgres).Insert("users").Columns("id", "email").Values("1", "user@example.com").OnConflict(OpResolveWithIgnore).ConflictColumns("id"),
			wantQuery: `INSERT INTO "users" ("id", "email") VALUES ($1, $2) ON CONFLICT ("id") DO UPDATE SET "id" = "id", "email" = "email"`,
			wantArgs:  []interface{}{"1", "user@example.com"},
		},

		{
			id: "example-69",
			input: Select().
				From(Table("users")).
				Where(EQ("name", "Alex")),
			wantQuery: "SELECT * FROM `users` WHERE `name` = ?",
			wantArgs:  []interface{}{"Alex"},
		},
		{
			id: "example-70",
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")),
			wantQuery: `SELECT * FROM "users"`,
		},
		{
			id: "example-71",
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(EQ("name", "Ariel")),
			wantQuery: `SELECT * FROM "users" WHERE "name" = $1`,
			wantArgs:  []interface{}{"Ariel"},
		},
		{
			id: "example-72",
			input: Select().
				From(Table("users")).
				Where(Or(EQ("name", "BAR"), EQ("name", "BAZ"))),
			wantQuery: "SELECT * FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"BAR", "BAZ"},
		},
		{
			id: "example-73",
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(And(EQ("name", "foo"), EQ("age", 20))),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? AND `age` = ?",
			wantArgs:  []interface{}{"foo", 10, "foo", 20},
		},
		{
			id: "example-74",
			input: Delete("users").
				Where(NotNull("parent_id")),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL",
		},
		{
			id: "example-75",
			input: Delete("users").
				Where(NotNull("parent_id")).
				Schema("mydb"),
			wantQuery: "DELETE FROM `mydb`.`users` WHERE `parent_id` IS NOT NULL",
		},
		{
			id: "example-76",
			input: Dialect(dialect.SQLite).
				Delete("users").
				Where(NotNull("parent_id")).
				Schema("mydb"),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL",
		},
		{
			id: "example-77",
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(IsNull("parent_id")),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NULL`,
		},
		{
			id: "example-78",
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(IsNull("parent_id")).
				Schema("mydb"),
			wantQuery: `DELETE FROM "mydb"."users" WHERE "parent_id" IS NULL`,
		},
		{
			id: "example-79",
			input: Delete("users").
				Where(And(IsNull("parent_id"), NotIn("name", "foo", "bar"))),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NULL AND `name` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-80",
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(And(IsNull("parent_id"), NotIn("name", "foo", "bar"))),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NULL AND "name" NOT IN ($1, $2)`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-81",
			input: Delete("users").
				Where(And(False(), False())),
			wantQuery: "DELETE FROM `users` WHERE FALSE AND FALSE",
		},
		{
			id: "example-82",
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(And(False(), False())),
			wantQuery: `DELETE FROM "users" WHERE FALSE AND FALSE`,
		},
		{
			id: "example-83",
			input: Delete("users").
				Where(Or(NotNull("parent_id"), EQ("parent_id", 10))),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL OR `parent_id` = ?",
			wantArgs:  []interface{}{10},
		},
		{
			id: "example-84",
			input: Dialect(dialect.Postgres).
				Delete("users").
				Where(Or(NotNull("parent_id"), EQ("parent_id", 10))),
			wantQuery: `DELETE FROM "users" WHERE "parent_id" IS NOT NULL OR "parent_id" = $1`,
			wantArgs:  []interface{}{10},
		},
		{
			id: "example-85",
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
			id: "example-86",
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
			id: "example-87",
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
			id: "example-88",
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
			id:        "example-89",
			input:     Select().From(Table("users")),
			wantQuery: "SELECT * FROM `users`",
		},
		{
			id:        "example-90",
			input:     Dialect(dialect.Postgres).Select().From(Table("users")),
			wantQuery: `SELECT * FROM "users"`,
		},
		{
			id:        "example-91",
			input:     Select().From(Table("users").Unquote()),
			wantQuery: "SELECT * FROM users",
		},
		{
			id:        "example-92",
			input:     Dialect(dialect.Postgres).Select().From(Table("users").Unquote()),
			wantQuery: "SELECT * FROM users",
		},
		{
			id:        "example-93",
			input:     Select().From(Table("users").As("u")),
			wantQuery: "SELECT * FROM `users` AS `u`",
		},
		{
			id:        "example-94",
			input:     Dialect(dialect.Postgres).Select().From(Table("users").As("u")),
			wantQuery: `SELECT * FROM "users" AS "u"`,
		},
		{
			id: "example-95",
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Select(t1.C("id"), t2.C("name")).From(t1).Join(t2)
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN `groups` AS `g`",
		},
		{
			id: "example-96",
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("groups").As("g")
				return Dialect(dialect.Postgres).Select(t1.C("id"), t2.C("name")).From(t1).Join(t2)
			}(),
			wantQuery: `SELECT "u"."id", "g"."name" FROM "users" AS "u" JOIN "groups" AS "g"`,
		},
		{
			id: "example-97",
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
			id: "example-98",
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
			id: "example-99",
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
			id: "example-100",
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
			id: "example-101",
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
			id: "example-102",
			input: func() Querier {
				t1 := Table("users").As("u")
				t2 := Table("user_groups").As("ug")
				return Select(t1.C("id"), As(Count("`*`"), "group_count")).
					From(t1).
					LeftJoin(t2).
					OnP(P(func(b *Builder) {
						b.Ident(t1.C("id")).WriteOp(OpEQ).Ident(t2.C("user_id"))
					})).
					GroupBy(t1.C("id")).Clone()
			}(),
			wantQuery: "SELECT `u`.`id`, COUNT(`*`) AS `group_count` FROM `users` AS `u` LEFT JOIN `user_groups` AS `ug` ON `u`.`id` = `ug`.`user_id` GROUP BY `u`.`id`",
		},
		{
			id: "example-103",
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
			id: "example-104",
			input: func() Querier {
				t1 := Table("users").As("u")
				return Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: "SELECT `u`.`name`, `u`.`age` FROM `users` AS `u`",
		},
		{
			id: "example-105",
			input: func() Querier {
				t1 := Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: `SELECT "u"."name", "u"."age" FROM "users" AS "u"`,
		},
		{
			id: "example-106",
			input: func() Querier {
				t1 := Dialect(dialect.Postgres).
					Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(t1.Columns("name", "age")...).From(t1)
			}(),
			wantQuery: `SELECT "u"."name", "u"."age" FROM "users" AS "u"`,
		},
		{
			id: "example-107",
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
			id: "example-108",
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
			id: "example-109",
			input: func() Querier {
				selector := Select().Where(Or(EQ("name", "foo"), EQ("name", "bar")))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-110",
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select().Where(Or(EQ("name", "foo"), EQ("name", "bar")))
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "users" WHERE "name" = $1 OR "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-111",
			input: func() Querier {
				selector := Select().From(Table("users")).As("t")
				return selector.Select(selector.C("name"))
			}(),
			wantQuery: "SELECT `t`.`name` FROM `users`",
		},
		{
			id: "example-112",
			input: func() Querier {
				selector := Dialect(dialect.Postgres).
					Select().From(Table("users")).As("t")
				return selector.Select(selector.C("name"))
			}(),
			wantQuery: `SELECT "t"."name" FROM "users"`,
		},
		{
			id: "example-113",
			input: func() Querier {
				selector := Select().From(Table("groups")).Where(EQ("name", "foo"))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `groups` WHERE `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			id: "example-114",
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select().From(Table("groups")).Where(EQ("name", "foo"))
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "groups" WHERE "name" = $1`,
			wantArgs:  []interface{}{"foo"},
		},
		{
			id: "example-115",
			input: func() Querier {
				selector := Select()
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `users`",
		},
		{
			id: "example-116",
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				selector := d.Select()
				return d.Delete("users").FromSelect(selector)
			}(),
			wantQuery: `DELETE FROM "users"`,
		},
		{
			id: "example-117",
			input: Select().
				From(Table("users")).
				Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))),
			wantQuery: "SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-118",
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))),
			wantQuery: `SELECT * FROM "users" WHERE NOT ("name" = $1 AND "age" = $2)`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-119",
			input: Select().
				From(Table("users")).
				Where(Or(EqualFold("name", "BAR"), EqualFold("name", "BAZ"))),
			wantQuery: "SELECT * FROM `users` WHERE LOWER(`name`) = ? OR LOWER(`name`) = ?",
			wantArgs:  []interface{}{"bar", "baz"},
		},
		{
			id: "example-120",
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(Or(EqualFold("name", "BAR"), EqualFold("name", "BAZ"))),
			wantQuery: `SELECT * FROM "users" WHERE LOWER("name") = $1 OR LOWER("name") = $2`,
			wantArgs:  []interface{}{"bar", "baz"},
		},
		{
			id: "example-121",
			input: Dialect(dialect.SQLite).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: "SELECT * FROM `users` WHERE LOWER(`name`) LIKE ? AND LOWER(`nick`) LIKE ?",
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			id: "example-122",
			input: Dialect(dialect.Postgres).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: `SELECT * FROM "users" WHERE "name" ILIKE $1 AND "nick" ILIKE $2`,
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			id: "example-123",
			input: Dialect(dialect.MySQL).
				Select().
				From(Table("users")).
				Where(And(ContainsFold("name", "Ariel"), ContainsFold("nick", "Bar"))),
			wantQuery: "SELECT * FROM `users` WHERE `name` COLLATE utf8mb4_general_ci LIKE ? AND `nick` COLLATE utf8mb4_general_ci LIKE ?",
			wantArgs:  []interface{}{"%ariel%", "%bar%"},
		},
		{
			id: "example-124",
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
			id: "example-125",
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
			id: "example-126",
			input: func() Querier {
				s1 := Select().From(Table("users")).Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))).As("users_view")
				return Select("name").From(s1)
			}(),
			wantQuery: "SELECT `name` FROM (SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)) AS `users_view`",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-127",
			input: func() Querier {
				d := Dialect(dialect.Postgres)
				s1 := d.Select().From(Table("users")).Where(Not(And(EQ("name", "foo"), EQ("age", "bar")))).As("users_view")
				return d.Select("name").From(s1)
			}(),
			wantQuery: `SELECT "name" FROM (SELECT * FROM "users" WHERE NOT ("name" = $1 AND "age" = $2)) AS "users_view"`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-128",
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
			id: "example-129",
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
			id: "example-130",
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
			id: "example-131",
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
			id:        "example-132",
			input:     Select().Count().From(Table("users")),
			wantQuery: "SELECT COUNT(*) FROM `users`",
		},
		{
			id: "example-133",
			input: Dialect(dialect.Postgres).
				Select().Count().From(Table("users")),
			wantQuery: `SELECT COUNT(*) FROM "users"`,
		},
		{
			id:        "example-134",
			input:     Select().Count(Distinct("id")).From(Table("users")),
			wantQuery: "SELECT COUNT(DISTINCT `id`) FROM `users`",
		},
		{
			id: "example-135",
			input: Dialect(dialect.Postgres).
				Select().Count(Distinct("id")).From(Table("users")),
			wantQuery: `SELECT COUNT(DISTINCT "id") FROM "users"`,
		},
		{
			id: "example-136",
			input: func() Querier {
				t1 := Table("users")
				t2 := Select().From(Table("groups"))
				t3 := Select().Count().From(t1).Join(t1).On(t2.C("id"), t1.C("blocked_id"))
				return t3.Count(Distinct(t3.Columns("id", "name")...))
			}(),
			wantQuery: "SELECT COUNT(DISTINCT `t0`.`id`, `t0`.`name`) FROM `users` AS `t0` JOIN `users` AS `t0` ON `groups`.`id` = `t0`.`blocked_id`",
		},
		{
			id: "example-137",
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
			id:        "example-138",
			input:     Select(Sum("age"), Min("age")).From(Table("users")),
			wantQuery: "SELECT SUM(`age`), MIN(`age`) FROM `users`",
		},
		{
			id: "example-139",
			input: Dialect(dialect.Postgres).
				Select(Sum("age"), Min("age")).
				From(Table("users")),
			wantQuery: `SELECT SUM("age"), MIN("age") FROM "users"`,
		},
		{
			id: "example-140",
			input: func() Querier {
				t1 := Table("users").As("u")
				return Select(As(Max(t1.C("age")), "max_age")).From(t1)
			}(),
			wantQuery: "SELECT MAX(`u`.`age`) AS `max_age` FROM `users` AS `u`",
		},
		{
			id: "example-141",
			input: func() Querier {
				t1 := Table("users").As("u")
				return Dialect(dialect.Postgres).
					Select(As(Max(t1.C("age")), "max_age")).
					From(t1)
			}(),
			wantQuery: `SELECT MAX("u"."age") AS "max_age" FROM "users" AS "u"`,
		},
		{
			id: "example-142",
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name`",
		},
		{
			id: "example-143",
			input: Dialect(dialect.Postgres).
				Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name"),
			wantQuery: `SELECT "name", COUNT(*) FROM "users" GROUP BY "name"`,
		},
		{
			id: "example-144",
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name").
				OrderBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name` ORDER BY `name`",
		},
		{
			id: "example-145",
			input: Dialect(dialect.Postgres).
				Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name").
				OrderBy("name"),
			wantQuery: `SELECT "name", COUNT(*) FROM "users" GROUP BY "name" ORDER BY "name"`,
		},
		{
			id: "example-146",
			input: Select("name", "age", Count("*")).
				From(Table("users")).
				GroupBy("name", "age").
				OrderBy(Desc("name"), "age"),
			wantQuery: "SELECT `name`, `age`, COUNT(*) FROM `users` GROUP BY `name`, `age` ORDER BY `name` DESC, `age`",
		},
		{
			id: "example-147",
			input: Dialect(dialect.Postgres).
				Select("name", "age", Count("*")).
				From(Table("users")).
				GroupBy("name", "age").
				OrderBy(Desc("name"), "age"),
			wantQuery: `SELECT "name", "age", COUNT(*) FROM "users" GROUP BY "name", "age" ORDER BY "name" DESC, "age"`,
		},
		{
			id: "example-148",
			input: Select("*").
				From(Table("users")).
				Limit(1),
			wantQuery: "SELECT * FROM `users` LIMIT 1",
		},
		{
			id: "example-149",
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("users")).
				Limit(1),
			wantQuery: `SELECT * FROM "users" LIMIT 1`,
		},
		{
			id:        "example-150",
			input:     Select("age").Distinct().From(Table("users")),
			wantQuery: "SELECT DISTINCT `age` FROM `users`",
		},
		{
			id: "example-151",
			input: Dialect(dialect.Postgres).
				Select("age").
				Distinct().
				From(Table("users")),
			wantQuery: `SELECT DISTINCT "age" FROM "users"`,
		},
		{
			id:        "example-152",
			input:     Select("age", "name").From(Table("users")).Distinct().OrderBy("name"),
			wantQuery: "SELECT DISTINCT `age`, `name` FROM `users` ORDER BY `name`",
		},
		{
			id: "example-153",
			input: Dialect(dialect.Postgres).
				Select("age", "name").
				From(Table("users")).
				Distinct().
				OrderBy("name"),
			wantQuery: `SELECT DISTINCT "age", "name" FROM "users" ORDER BY "name"`,
		},
		{
			id:        "example-154",
			input:     Select("age").From(Table("users")).Where(EQ("name", "foo")).Or().Where(EQ("name", "bar")),
			wantQuery: "SELECT `age` FROM `users` WHERE `name` = ? OR `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id: "example-155",
			input: Dialect(dialect.Postgres).
				Select("age").
				From(Table("users")).
				Where(EQ("name", "foo")).Or().Where(EQ("name", "bar")),
			wantQuery: `SELECT "age" FROM "users" WHERE "name" = $1 OR "name" = $2`,
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			id:        "example-156",
			input:     Queries{With("users_view").As(Select().From(Table("users"))), Select().From(Table("users_view"))},
			wantQuery: "WITH users_view AS (SELECT * FROM `users`) SELECT * FROM `users_view`",
		},
		{
			id: "example-157",
			input: func() Querier {
				base := Select("*").From(Table("groups"))
				return Queries{With("groups").As(base.Clone().Where(EQ("name", "bar"))), base.Select("age")}
			}(),
			wantQuery: "WITH groups AS (SELECT * FROM `groups` WHERE `name` = ?) SELECT `age` FROM `groups`",
			wantArgs:  []interface{}{"bar"},
		},
		{
			id: "example-158",
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
			id: "example-159",
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
			id: "example-160",
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
			id: "example-161",
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
			id: "example-162",
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
			id:        "example-163",
			input:     CreateIndex("name_index").Table("users").Column("name"),
			wantQuery: "CREATE INDEX `name_index` ON `users`(`name`)",
		},
		{
			id: "example-164",
			input: Dialect(dialect.Postgres).
				CreateIndex("name_index").
				Table("users").
				Column("name"),
			wantQuery: `CREATE INDEX "name_index" ON "users"("name")`,
		},
		{
			id:        "example-165",
			input:     CreateIndex("unique_name").Unique().Table("users").Columns("first", "last"),
			wantQuery: "CREATE UNIQUE INDEX `unique_name` ON `users`(`first`, `last`)",
		},
		{
			id: "example-166",
			input: Dialect(dialect.Postgres).
				CreateIndex("unique_name").
				Unique().
				Table("users").
				Columns("first", "last"),
			wantQuery: `CREATE UNIQUE INDEX "unique_name" ON "users"("first", "last")`,
		},
		{
			id:        "example-167",
			input:     DropIndex("name_index"),
			wantQuery: "DROP INDEX `name_index`",
		},
		{
			id: "example-168",
			input: Dialect(dialect.Postgres).
				DropIndex("name_index"),
			wantQuery: `DROP INDEX "name_index"`,
		},
		{
			id:        "example-169",
			input:     DropIndex("name_index").Table("users"),
			wantQuery: "DROP INDEX `name_index` ON `users`",
		},
		{
			id: "example-170",
			input: Select().
				From(Table("pragma_table_info('t1')").Unquote()).
				OrderBy("pk"),
			wantQuery: "SELECT * FROM pragma_table_info('t1') ORDER BY `pk`",
		},
		{
			id: "example-171",
			input: AlterTable("users").
				AddColumn(Column("spouse").Type("integer").
					Constraint(ForeignKey("user_spouse").
						Reference(Reference().Table("users").Columns("id")).
						OnDelete("SET NULL"))),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `spouse` integer CONSTRAINT user_spouse REFERENCES `users`(`id`) ON DELETE SET NULL",
		},
		{
			id: "example-172",
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
			id: "example-173",
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
			id: "example-174",
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
			id: "example-175",
			input: func() Querier {
				t1, t2 := Table("users").Schema("s1"), Table("pets").Schema("s2")
				return Select("*").
					From(t1).Join(t2).
					OnP(P(func(b *Builder) {
						b.Ident(t1.C("id")).WriteOp(OpEQ).Ident(t2.C("owner_id"))
					})).
					Where(EQ(t2.C("name"), "pedro"))
			}(),
			wantQuery: "SELECT * FROM `s1`.`users` JOIN `s2`.`pets` AS `t0` ON `s1`.`users`.`id` = `t0`.`owner_id` WHERE `t0`.`name` = ?",
			wantArgs:  []interface{}{"pedro"},
		},
		{
			id: "example-176",
			input: func() Querier {
				t1, t2 := Table("users").Schema("s1"), Table("pets").Schema("s2")
				sel := Select("*").
					From(t1).Join(t2).
					OnP(P(func(b *Builder) {
						b.Ident(t1.C("id")).WriteOp(OpEQ).Ident(t2.C("owner_id"))
					})).
					Where(EQ(t2.C("name"), "pedro"))
				sel.SetDialect(dialect.SQLite)
				return sel
			}(),
			wantQuery: "SELECT * FROM `users` JOIN `pets` AS `t0` ON `users`.`id` = `t0`.`owner_id` WHERE `t0`.`name` = ?",
			wantArgs:  []interface{}{"pedro"},
		},
		{
			id: "example-177",
			input: Dialect(dialect.Postgres).
				Select("*").
				From(Table("users")).
				Where(ExprP("name = $1", "pedro")).
				Where(P(func(b *Builder) {
					b.Join(Expr("name = $2", "pedro"))
				})).
				Where(EQ("name", "pedro")).
				Where(
					And(
						In(
							"id",
							Select("owner_id").
								From(Table("pets")).
								Where(EQ("name", "luna")),
						),
						EQ("active", true),
					),
				),
			wantQuery: `SELECT * FROM "users" WHERE ((name = $1 AND name = $2) AND "name" = $3) AND ("id" IN (SELECT "owner_id" FROM "pets" WHERE "name" = $4) AND "active" = $5)`,
			wantArgs:  []interface{}{"pedro", "pedro", "pedro", "luna", true},
		},
		{
			id:        "example-178",
			input:     Dialect(dialect.MySQL).Insert("users").Set("email", "user@example.com").OnConflict(OpResolveWithAlternateValues).UpdateSet("email", "user-1@example.com").ConflictColumns("email"),
			wantQuery: "INSERT INTO `users` (`email`) VALUES (?) ON DUPLICATE KEY UPDATE `email` = ?",
			wantArgs:  []interface{}{"user@example.com", "user-1@example.com"},
		},
		{
			id:        "example-179",
			input:     Dialect(dialect.Postgres).Insert("users").Set("email", "user@example.com").OnConflict(OpResolveWithAlternateValues).UpdateSet("email", "user-1@example.com").ConflictColumns("email"),
			wantQuery: `INSERT INTO "users" ("email") VALUES ($1) ON CONFLICT ("email") DO UPDATE SET "email" = $2`,
			wantArgs:  []interface{}{"user@example.com", "user-1@example.com"},
		},
		{
			id:        "example-180",
			input:     Dialect(dialect.Postgres).Insert("users").Set("email", "user@example.com").OnConflict(OpResolveWithIgnore).ConflictColumns("email"),
			wantQuery: `INSERT INTO "users" ("email") VALUES ($1) ON CONFLICT ("email") DO UPDATE SET "email" = "email"`,
			wantArgs:  []interface{}{"user@example.com"},
		},
		{
			id:        "example-181",
			input:     Dialect(dialect.MySQL).Insert("users").Set("email", "user@example.com").OnConflict(OpResolveWithIgnore).ConflictColumns("email"),
			wantQuery: "INSERT INTO `users` (`email`) VALUES (?) ON DUPLICATE KEY UPDATE `email` = `email`",
			wantArgs:  []interface{}{"user@example.com"},
		},
		{
			id:        "example-182",
			input:     Dialect(dialect.MySQL).Insert("users").Set("email", "user@example.com").OnConflict(OpResolveWithNewValues).ConflictColumns("email"),
			wantQuery: "INSERT INTO `users` (`email`) VALUES (?) ON DUPLICATE KEY UPDATE `email` = VALUES(`email`)",
			wantArgs:  []interface{}{"user@example.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {
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

func TestSelector_OrderByExpr(t *testing.T) {
	query, args := Select("*").
		From(Table("users")).
		Where(GT("age", 28)).
		OrderBy("name").
		OrderExpr(Expr("CASE WHEN id=? THEN id WHEN id=? THEN name END DESC", 1, 2)).
		Query()
	require.Equal(t, "SELECT * FROM `users` WHERE `age` > ? ORDER BY `name`, CASE WHEN id=? THEN id WHEN id=? THEN name END DESC", query)
	require.Equal(t, []interface{}{28, 1, 2}, args)
}

func TestBuilderContext(t *testing.T) {
	type key string
	want := "myval"
	ctx := context.WithValue(context.Background(), key("mykey"), want)
	sel := Dialect(dialect.Postgres).Select().WithContext(ctx)
	if got := sel.Context().Value(key("mykey")).(string); got != want {
		t.Fatalf("expected selector context key to be %q but got %q", want, got)
	}
	if got := sel.Clone().Context().Value(key("mykey")).(string); got != want {
		t.Fatalf("expected cloned selector context key to be %q but got %q", want, got)
	}
}

type point struct {
	xy []float64
	*testing.T
}

func (p point) FormatParam(placeholder string, info *StmtInfo) string {
	require.Equal(p.T, dialect.MySQL, info.Dialect)
	return "ST_GeomFromWKB(" + placeholder + ")"
}

func TestParamFormatter(t *testing.T) {
	p := point{xy: []float64{1, 2}, T: t}
	query, args := Dialect(dialect.MySQL).
		Select().
		From(Table("users")).
		Where(EQ("point", p)).
		Query()
	require.Equal(t, "SELECT * FROM `users` WHERE `point` = ST_GeomFromWKB(?)", query)
	require.Equal(t, p, args[0])
}
