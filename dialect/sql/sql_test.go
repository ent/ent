// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"testing"

	"entgo.io/ent/dialect"

	"github.com/stretchr/testify/require"
)

func TestFieldIsNull(t *testing.T) {
	p := FieldIsNull("name")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` IS NULL", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" IS NULL`, query)
		require.Empty(t, args)
	})
}

func TestFieldNotNull(t *testing.T) {
	p := FieldNotNull("name")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` IS NOT NULL", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" IS NOT NULL`, query)
		require.Empty(t, args)
	})
}

func TestFieldEQ(t *testing.T) {
	p := FieldEQ("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` = ?", query)
		require.Equal(t, []any{"a8m"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" = $1`, query)
		require.Equal(t, []any{"a8m"}, args)
	})
}

func TestFieldsEQ(t *testing.T) {
	p := FieldsEQ("create_time", "update_time")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`create_time` = `users`.`update_time`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."create_time" = "users"."update_time"`, query)
		require.Empty(t, args)
	})
}

func TestFieldsNEQ(t *testing.T) {
	p := FieldsNEQ("create_time", "update_time")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`create_time` <> `users`.`update_time`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."create_time" <> "users"."update_time"`, query)
		require.Empty(t, args)
	})
}

func TestFieldNEQ(t *testing.T) {
	p := FieldNEQ("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` <> ?", query)
		require.Equal(t, []any{"a8m"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" <> $1`, query)
		require.Equal(t, []any{"a8m"}, args)
	})
}

func TestFieldGT(t *testing.T) {
	p := FieldGT("stars", 1000)
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`stars` > ?", query)
		require.Equal(t, []any{1000}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."stars" > $1`, query)
		require.Equal(t, []any{1000}, args)
	})
}

func TestFieldsGT(t *testing.T) {
	p := FieldsGT("a", "b")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`a` > `users`.`b`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."a" > "users"."b"`, query)
		require.Empty(t, args)
	})
}

func TestFieldGTE(t *testing.T) {
	p := FieldGTE("stars", 1000)
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`stars` >= ?", query)
		require.Equal(t, []any{1000}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."stars" >= $1`, query)
		require.Equal(t, []any{1000}, args)
	})
}

func TestFieldsGTE(t *testing.T) {
	p := FieldsGTE("a", "b")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`a` >= `users`.`b`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."a" >= "users"."b"`, query)
		require.Empty(t, args)
	})
}

func TestFieldLT(t *testing.T) {
	p := FieldLT("stars", 1000)
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`stars` < ?", query)
		require.Equal(t, []any{1000}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."stars" < $1`, query)
		require.Equal(t, []any{1000}, args)
	})
}

func TestFieldsLT(t *testing.T) {
	p := FieldsLT("a", "b")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`a` < `users`.`b`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."a" < "users"."b"`, query)
		require.Empty(t, args)
	})
}

func TestFieldLTE(t *testing.T) {
	p := FieldLTE("stars", 1000)
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`stars` <= ?", query)
		require.Equal(t, []any{1000}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."stars" <= $1`, query)
		require.Equal(t, []any{1000}, args)
	})
}

func TestFieldsLTE(t *testing.T) {
	p := FieldsLTE("a", "b")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`a` <= `users`.`b`", query)
		require.Empty(t, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."a" <= "users"."b"`, query)
		require.Empty(t, args)
	})
}

func TestFieldIn(t *testing.T) {
	p := FieldIn("name", "a8m", "foo", "bar")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` IN (?, ?, ?)", query)
		require.Equal(t, []any{"a8m", "foo", "bar"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" IN ($1, $2, $3)`, query)
		require.Equal(t, []any{"a8m", "foo", "bar"}, args)
	})
}

func TestFieldNotIn(t *testing.T) {
	p := FieldNotIn("id", 1, 2, 3)
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`id` NOT IN (?, ?, ?)", query)
		require.Equal(t, []any{1, 2, 3}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."id" NOT IN ($1, $2, $3)`, query)
		require.Equal(t, []any{1, 2, 3}, args)
	})
}

func TestFieldEqualFold(t *testing.T) {
	p := FieldEqualFold("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` COLLATE utf8mb4_general_ci = ?", query)
		require.Equal(t, []any{"a8m"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" ILIKE $1`, query)
		require.Equal(t, []any{"a8m"}, args)
	})
}

func TestFieldHasPrefix(t *testing.T) {
	p := FieldHasPrefix("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` LIKE ?", query)
		require.Equal(t, []any{"a8m%"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" LIKE $1`, query)
		require.Equal(t, []any{"a8m%"}, args)
	})
}

func TestFieldHasPrefixFold(t *testing.T) {
	p := FieldHasPrefixFold("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` COLLATE utf8mb4_general_ci LIKE ?", query)
		require.Equal(t, []any{"a8m%"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" ILIKE $1`, query)
		require.Equal(t, []any{"a8m%"}, args)
	})
}

func TestFieldHasSuffix(t *testing.T) {
	p := FieldHasSuffix("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` LIKE ?", query)
		require.Equal(t, []any{"%a8m"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" LIKE $1`, query)
		require.Equal(t, []any{"%a8m"}, args)
	})
}

func TestFieldHasSuffixFold(t *testing.T) {
	p := FieldHasSuffixFold("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` COLLATE utf8mb4_general_ci LIKE ?", query)
		require.Equal(t, []any{"%a8m"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" ILIKE $1`, query)
		require.Equal(t, []any{"%a8m"}, args)
	})
}

func TestFieldContains(t *testing.T) {
	p := FieldContains("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` LIKE ?", query)
		require.Equal(t, []any{"%a8m%"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" LIKE $1`, query)
		require.Equal(t, []any{"%a8m%"}, args)
	})
}

func TestFieldContainsFold(t *testing.T) {
	p := FieldContainsFold("name", "a8m")
	t.Run("MySQL", func(t *testing.T) {
		s := Dialect(dialect.MySQL).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, "SELECT * FROM `users` WHERE `users`.`name` COLLATE utf8mb4_general_ci LIKE ?", query)
		require.Equal(t, []any{"%a8m%"}, args)
	})
	t.Run("PostgreSQL", func(t *testing.T) {
		s := Dialect(dialect.Postgres).Select("*").From(Table("users"))
		p(s)
		query, args := s.Query()
		require.Equal(t, `SELECT * FROM "users" WHERE "users"."name" ILIKE $1`, query)
		require.Equal(t, []any{"%a8m%"}, args)
	})
}

func TestAndPredicates(t *testing.T) {
	s := Select("*").From(Table("users")).Where(EQ("name", "a8m"))
	p := AndPredicates(
		FieldEQ("a", "foo"),
		FieldEQ("b", 1),
		func(s *Selector) {
			petT := Table("pets").As("p")
			s.Join(petT).On(petT.C("owner_id"), s.C("id"))
		},
	)
	p(s)
	query, args := s.Query()
	require.Equal(t, "SELECT * FROM `users` JOIN `pets` AS `p` ON `p`.`owner_id` = `users`.`id` WHERE `name` = ? AND (`users`.`a` = ? AND `users`.`b` = ?)", query)
	require.Equal(t, []any{"a8m", "foo", 1}, args)
}

func TestOrPredicates(t *testing.T) {
	s := Select("*").From(Table("users")).Where(EQ("name", "a8m"))
	p := OrPredicates(
		AndPredicates(
			FieldEQ("a", "foo"),
			FieldEQ("b", 1),
		),
		func(s *Selector) {
			petT := Table("pets").As("p")
			s.Join(petT).On(petT.C("owner_id"), s.C("id"))
			s.Where(EQ(petT.C("name"), "c"))
		},
	)
	p(s)
	query, args := s.Query()
	require.Equal(t, "SELECT * FROM `users` JOIN `pets` AS `p` ON `p`.`owner_id` = `users`.`id` WHERE `name` = ? AND ((`users`.`a` = ? AND `users`.`b` = ?) OR `p`.`name` = ?)", query)
	require.Equal(t, []any{"a8m", "foo", 1, "c"}, args)
}

func TestNotPredicates(t *testing.T) {
	s := Select("*").From(Table("users")).Where(EQ("name", "a8m"))
	NotPredicates(FieldEQ("a", "a"), FieldEQ("b", "b"))(s)
	NotPredicates(FieldEQ("c", "c"))(s)
	query, args := s.Query()
	require.Equal(t, "SELECT * FROM `users` WHERE (`name` = ? AND (NOT (`users`.`a` = ? AND `users`.`b` = ?))) AND (NOT (`users`.`c` = ?))", query)
	require.Equal(t, []any{"a8m", "a", "b", "c"}, args)
}
