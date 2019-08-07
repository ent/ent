package sql

import (
	"strconv"
	"testing"

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
			input: AlterTable("users").
				AddColumn(Column("group_id").Type("int").Attr("UNIQUE")).
				AddForeignKey(ForeignKey().Columns("group_id").
					Reference(Reference().Table("groups").Columns("id")).
					OnDelete("CASCADE"),
				),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `group_id` int UNIQUE, ADD CONSTRAINT FOREIGN KEY(`group_id`) REFERENCES `groups`(`id`) ON DELETE CASCADE",
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
			input: AlterTable("users").
				AddColumn(Column("age").Type("int")).
				AddColumn(Column("name").Type("varchar(255)")),
			wantQuery: "ALTER TABLE `users` ADD COLUMN `age` int, ADD COLUMN `name` varchar(255)",
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
			input:     Insert("users").Columns("age").Values(1),
			wantQuery: "INSERT INTO `users` (`age`) VALUES (?)",
			wantArgs:  []interface{}{1},
		},
		{
			input:     Insert("users").Columns("name", "age").Values("a8m", 10),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?)",
			wantArgs:  []interface{}{"a8m", 10},
		},
		{
			input:     Insert("users").Columns("name", "age").Values("a8m", 10).Values("foo", 20),
			wantQuery: "INSERT INTO `users` (`name`, `age`) VALUES (?, ?), (?, ?)",
			wantArgs:  []interface{}{"a8m", 10, "foo", 20},
		},
		{
			input:     Update("users").Set("name", "foo"),
			wantQuery: "UPDATE `users` SET `name` = ?",
			wantArgs:  []interface{}{"foo"},
		},
		{
			input:     Update("users").Set("name", "foo").Set("age", 10),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ?",
			wantArgs:  []interface{}{"foo", 10},
		},
		{
			input:     Update("users").Set("name", "foo").Where(EQ("name", "bar")),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` = ?",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input:     Update("users").Set("name", "foo").SetNull("spouse_id"),
			wantQuery: "UPDATE `users` SET `spouse_id` = NULL, `name` = ?",
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
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(EQ("name", "bar").Or().EQ("name", "baz")),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? OR `name` = ?",
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
			input: Update("users").
				Set("name", "foo").
				Where(In("name", "bar", "baz").And().NotIn("age", 1, 2)),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `name` IN (?, ?) AND `age` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar", "baz", 1, 2},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Where(HasPrefix("nickname", "a8m").And().Contains("lastname", "mash")),
			wantQuery: "UPDATE `users` SET `name` = ? WHERE `nickname` LIKE ? AND `lastname` LIKE ?",
			wantArgs:  []interface{}{"foo", "a8m%", "%mash%"},
		},
		{
			input: Update("users").
				Set("name", "foo").
				Set("age", 10).
				Where(P().EQ("name", "foo").And().EQ("age", 20)),
			wantQuery: "UPDATE `users` SET `name` = ?, `age` = ? WHERE `name` = ? AND `age` = ?",
			wantArgs:  []interface{}{"foo", 10, "foo", 20},
		},
		{
			input: Delete("users").
				Where(NotNull("parent_id")),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL",
		},
		{
			input: Delete("users").
				Where(IsNull("parent_id")),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NULL",
		},
		{
			input: Delete("users").
				Where(IsNull("parent_id").And().NotIn("name", "foo", "bar")),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NULL AND `name` NOT IN (?, ?)",
			wantArgs:  []interface{}{"foo", "bar"},
		},
		{
			input: Delete("users").
				Where(False().And().False()),
			wantQuery: "DELETE FROM `users` WHERE FALSE AND FALSE",
		},
		{
			input: Delete("users").
				Where(NotNull("parent_id").Or().EQ("parent_id", 10)),
			wantQuery: "DELETE FROM `users` WHERE `parent_id` IS NOT NULL OR `parent_id` = ?",
			wantArgs:  []interface{}{10},
		},
		{
			input: Delete("users").
				Where(
					Or(
						EQ("name", "foo").And().EQ("age", 10),
						EQ("name", "bar").And().EQ("age", 20),
						And(
							EQ("name", "qux"),
							EQ("age", 1).Or().EQ("age", 2),
						),
					),
				),
			wantQuery: "DELETE FROM `users` WHERE (`name` = ? AND `age` = ?) OR (`name` = ? AND `age` = ?) OR ((`name` = ?) AND (`age` = ? OR `age` = ?))",
			wantArgs:  []interface{}{"foo", 10, "bar", 20, "qux", 1, 2},
		},
		{
			input:     Select().From(Table("users")),
			wantQuery: "SELECT * FROM `users`",
		},
		{
			input:     Select().From(Table("users").Unquote()),
			wantQuery: "SELECT * FROM users",
		},
		{
			input:     Select().From(Table("users").As("u")),
			wantQuery: "SELECT * FROM `users` AS `u`",
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
				return Select(t1.C("id"), t2.C("name")).
					From(t1).
					Join(t2).
					On(t1.C("id"), t2.C("user_id")).
					Where(EQ(t1.C("name"), "bar").And().NotNull(t2.C("name")))
			}(),
			wantQuery: "SELECT `u`.`id`, `g`.`name` FROM `users` AS `u` JOIN `groups` AS `g` ON `u`.`id` = `g`.`user_id` WHERE `u`.`name` = ? AND `g`.`name` IS NOT NULL",
			wantArgs:  []interface{}{"bar"},
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
				selector := Select().Where(EQ("name", "foo").Or().EQ("name", "bar"))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `users` WHERE `name` = ? OR `name` = ?",
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
				selector := Select().From(Table("groups")).Where(EQ("name", "foo"))
				return Delete("users").FromSelect(selector)
			}(),
			wantQuery: "DELETE FROM `groups` WHERE `name` = ?",
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
			input:     Select().From(Table("users")).Where(Not(EQ("name", "foo").And().EQ("age", "bar"))),
			wantQuery: "SELECT * FROM `users` WHERE NOT (`name` = ? AND `age` = ?)",
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
				return Select().
					From(t1).
					Where(Not(In(t1.C("id"), Select("owner_id").From(Table("pets")).Where(EQ("name", "pedro")))))
			}(),
			wantQuery: "SELECT * FROM `users` WHERE NOT (`users`.`id` IN (SELECT `owner_id` FROM `pets` WHERE `name` = ?))",
			wantArgs:  []interface{}{"pedro"},
		},
		{
			input:     Select().Count().From(Table("users")),
			wantQuery: "SELECT COUNT(*) FROM `users`",
		},
		{
			input:     Select().Count(Distinct("id")).From(Table("users")),
			wantQuery: "SELECT COUNT(DISTINCT `id`) FROM `users`",
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
			input:     Select(Sum("age"), Min("age")).From(Table("users")),
			wantQuery: "SELECT SUM(`age`), MIN(`age`) FROM `users`",
		},
		{
			input: func() Querier {
				t1 := Table("users").As("u")
				return Select(As(Max(t1.C("age")), "max_age")).From(t1)
			}(),
			wantQuery: "SELECT MAX(`u`.`age`) AS `max_age` FROM `users` AS `u`",
		},
		{
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name`",
		},
		{
			input: Select("name", Count("*")).
				From(Table("users")).
				GroupBy("name").
				OrderBy("name"),
			wantQuery: "SELECT `name`, COUNT(*) FROM `users` GROUP BY `name` ORDER BY `name`",
		},
		{
			input: Select("name", "age", Count("*")).
				From(Table("users")).
				GroupBy("name", "age").
				OrderBy(Desc("name"), "age"),
			wantQuery: "SELECT `name`, `age`, COUNT(*) FROM `users` GROUP BY `name`, `age` ORDER BY `name` DESC, `age`",
		},
		{
			input:     Select("*").From(Table("users")).Limit(1),
			wantQuery: "SELECT * FROM `users` LIMIT ?",
			wantArgs:  []interface{}{1},
		},
		{
			input:     Select("age").Distinct().From(Table("users")),
			wantQuery: "SELECT DISTINCT `age` FROM `users`",
		},
		{
			input:     Select("age", "name").From(Table("users")).Distinct().OrderBy("name"),
			wantQuery: "SELECT DISTINCT `age`, `name` FROM `users` ORDER BY `name`",
		},
		{
			input:     Select("age").From(Table("users")).Where(EQ("name", "foo")).Or().Where(EQ("name", "bar")),
			wantQuery: "SELECT `age` FROM `users` WHERE (`name` = ?) OR (`name` = ?)",
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
	}
	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			query, args := tt.input.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
