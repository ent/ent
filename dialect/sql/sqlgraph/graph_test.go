// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"context"
	"database/sql/driver"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/require"
)

func TestNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		input     *Step
		wantQuery string
		wantArgs  []any
	}{
		{
			name: "O2O/1type",
			// Since the relation is on the same sql.Table,
			// V used as a reference value.
			input: NewStep(
				From("users", "id", 1),
				To("users", "id"),
				Edge(O2O, false, "users", "spouse_id"),
			),
			wantQuery: "SELECT * FROM `users` WHERE `spouse_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "O2O/1type/inverse",
			input: NewStep(
				From("nodes", "id", 1),
				To("nodes", "id"),
				Edge(O2O, true, "nodes", "prev_id"),
			),
			wantQuery: "SELECT * FROM `nodes` JOIN (SELECT `prev_id` FROM `nodes` WHERE `id` = ?) AS `t1` ON `nodes`.`id` = `t1`.`prev_id`",
			wantArgs:  []any{1},
		},
		{
			name: "O2M/1type",
			input: NewStep(
				From("users", "id", 1),
				To("users", "id"),
				Edge(O2M, false, "users", "parent_id"),
			),
			wantQuery: "SELECT * FROM `users` WHERE `parent_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "O2O/2types",
			input: NewStep(
				From("users", "id", 2),
				To("card", "id"),
				Edge(O2O, false, "cards", "owner_id"),
			),
			wantQuery: "SELECT * FROM `card` WHERE `owner_id` = ?",
			wantArgs:  []any{2},
		},
		{
			name: "O2O/2types/inverse",
			input: NewStep(
				From("cards", "id", 2),
				To("users", "id"),
				Edge(O2O, true, "cards", "owner_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `cards` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []any{2},
		},
		{
			name: "O2M/2types",
			input: NewStep(
				From("users", "id", 1),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			wantQuery: "SELECT * FROM `pets` WHERE `owner_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "M2O/2types/inverse",
			input: NewStep(
				From("pets", "id", 2),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `pets` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []any{2},
		},
		{
			name: "M2O/1type/inverse",
			input: NewStep(
				From("users", "id", 2),
				To("users", "id"),
				Edge(M2O, true, "users", "parent_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `parent_id` FROM `users` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`parent_id`",
			wantArgs:  []any{2},
		},
		{
			name: "M2M/2type",
			input: NewStep(
				From("groups", "id", 2),
				To("users", "id"),
				Edge(M2M, false, "user_groups", "group_id", "user_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `user_groups`.`user_id` FROM `user_groups` WHERE `user_groups`.`group_id` = ?) AS `t1` ON `users`.`id` = `t1`.`user_id`",
			wantArgs:  []any{2},
		},
		{
			name: "M2M/2type/inverse",
			input: NewStep(
				From("users", "id", 2),
				To("groups", "id"),
				Edge(M2M, true, "user_groups", "group_id", "user_id"),
			),
			wantQuery: "SELECT * FROM `groups` JOIN (SELECT `user_groups`.`group_id` FROM `user_groups` WHERE `user_groups`.`user_id` = ?) AS `t1` ON `groups`.`id` = `t1`.`group_id`",
			wantArgs:  []any{2},
		},
		{
			name: "schema/O2O/1type",
			// Since the relation is on the same sql.Table,
			// V used as a reference value.
			input: func() *Step {
				step := NewStep(
					From("users", "id", 1),
					To("users", "id"),
					Edge(O2O, false, "users", "spouse_id"),
				)
				step.To.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`users` WHERE `spouse_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "schema/O2O/1type/inverse",
			input: func() *Step {
				step := NewStep(
					From("nodes", "id", 1),
					To("nodes", "id"),
					Edge(O2O, true, "nodes", "prev_id"),
				)
				step.To.Schema = "mydb"
				step.Edge.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`nodes` JOIN (SELECT `prev_id` FROM `mydb`.`nodes` WHERE `id` = ?) AS `t1` ON `mydb`.`nodes`.`id` = `t1`.`prev_id`",
			wantArgs:  []any{1},
		},
		{
			name: "schema/O2M/1type",
			input: func() *Step {
				step := NewStep(
					From("users", "id", 1),
					To("users", "id"),
					Edge(O2M, false, "users", "parent_id"),
				)
				step.To.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`users` WHERE `parent_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "schema/O2O/2types",
			input: func() *Step {
				step := NewStep(
					From("users", "id", 2),
					To("card", "id"),
					Edge(O2O, false, "cards", "owner_id"),
				)
				step.To.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`card` WHERE `owner_id` = ?",
			wantArgs:  []any{2},
		},
		{
			name: "schema/O2O/2types/inverse",
			input: func() *Step {
				step := NewStep(
					From("cards", "id", 2),
					To("users", "id"),
					Edge(O2O, true, "cards", "owner_id"),
				)
				step.To.Schema = "mydb"
				step.Edge.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`users` JOIN (SELECT `owner_id` FROM `mydb`.`cards` WHERE `id` = ?) AS `t1` ON `mydb`.`users`.`id` = `t1`.`owner_id`",
			wantArgs:  []any{2},
		},
		{
			name: "schema/O2M/2types",
			input: func() *Step {
				step := NewStep(
					From("users", "id", 1),
					To("pets", "id"),
					Edge(O2M, false, "pets", "owner_id"),
				)
				step.To.Schema = "mydb"
				return step
			}(),
			wantQuery: "SELECT * FROM `mydb`.`pets` WHERE `owner_id` = ?",
			wantArgs:  []any{1},
		},
		{
			name: "schema/M2O/2types/inverse",
			input: func() *Step {
				step := NewStep(
					From("pets", "id", 2),
					To("users", "id"),
					Edge(M2O, true, "pets", "owner_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s2"
				return step
			}(),
			wantQuery: "SELECT * FROM `s1`.`users` JOIN (SELECT `owner_id` FROM `s2`.`pets` WHERE `id` = ?) AS `t1` ON `s1`.`users`.`id` = `t1`.`owner_id`",
			wantArgs:  []any{2},
		},
		{
			name: "schema/M2O/1type/inverse",
			input: func() *Step {
				step := NewStep(
					From("users", "id", 2),
					To("users", "id"),
					Edge(M2O, true, "users", "parent_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s1"
				return step
			}(),
			wantQuery: "SELECT * FROM `s1`.`users` JOIN (SELECT `parent_id` FROM `s1`.`users` WHERE `id` = ?) AS `t1` ON `s1`.`users`.`id` = `t1`.`parent_id`",
			wantArgs:  []any{2},
		},
		{
			name: "schema/M2M/2type",
			input: func() *Step {
				step := NewStep(
					From("groups", "id", 2),
					To("users", "id"),
					Edge(M2M, false, "user_groups", "group_id", "user_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s2"
				return step
			}(),
			wantQuery: "SELECT * FROM `s1`.`users` JOIN (SELECT `s2`.`user_groups`.`user_id` FROM `s2`.`user_groups` WHERE `s2`.`user_groups`.`group_id` = ?) AS `t1` ON `s1`.`users`.`id` = `t1`.`user_id`",
			wantArgs:  []any{2},
		},
		{
			name: "schema/M2M/2type/inverse",
			input: func() *Step {
				step := NewStep(
					From("users", "id", 2),
					To("groups", "id"),
					Edge(M2M, true, "user_groups", "group_id", "user_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s2"
				return step
			}(),
			wantQuery: "SELECT * FROM `s1`.`groups` JOIN (SELECT `s2`.`user_groups`.`group_id` FROM `s2`.`user_groups` WHERE `s2`.`user_groups`.`user_id` = ?) AS `t1` ON `s1`.`groups`.`id` = `t1`.`group_id`",
			wantArgs:  []any{2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector := Neighbors("", tt.input)
			query, args := selector.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}

func TestSetNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		input     *Step
		wantQuery string
		wantArgs  []any
	}{
		{
			name: "O2M/2types",
			input: NewStep(
				From("users", "id", sql.Select().From(sql.Table("users")).Where(sql.EQ("name", "a8m"))),
				To("pets", "id"),
				Edge(O2M, false, "users", "owner_id"),
			),
			wantQuery: `SELECT * FROM "pets" JOIN (SELECT "users"."id" FROM "users" WHERE "name" = $1) AS "t1" ON "pets"."owner_id" = "t1"."id"`,
			wantArgs:  []any{"a8m"},
		},
		{
			name: "M2O/2types",
			input: NewStep(
				From("pets", "id", sql.Select().From(sql.Table("pets")).Where(sql.EQ("name", "pedro"))),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			wantQuery: `SELECT * FROM "users" JOIN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $1) AS "t1" ON "users"."id" = "t1"."owner_id"`,
			wantArgs:  []any{"pedro"},
		},
		{
			name: "M2M/2types",
			input: NewStep(
				From("users", "id", sql.Select().From(sql.Table("users")).Where(sql.EQ("name", "a8m"))),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			wantQuery: `
SELECT *
FROM "groups"
JOIN
  (SELECT "user_groups"."group_id"
   FROM "user_groups"
   JOIN
     (SELECT "users"."id"
      FROM "users"
      WHERE "name" = $1) AS "t1" ON "user_groups"."user_id" = "t1"."id") AS "t1" ON "groups"."id" = "t1"."group_id"`,
			wantArgs: []any{"a8m"},
		},
		{
			name: "M2M/2types/inverse",
			input: NewStep(
				From("groups", "id", sql.Select().From(sql.Table("groups")).Where(sql.EQ("name", "GitHub"))),
				To("users", "id"),
				Edge(M2M, true, "user_groups", "user_id", "group_id"),
			),
			wantQuery: `
SELECT *
FROM "users"
JOIN
  (SELECT "user_groups"."user_id"
   FROM "user_groups"
   JOIN
     (SELECT "groups"."id"
      FROM "groups"
      WHERE "name" = $1) AS "t1" ON "user_groups"."group_id" = "t1"."id") AS "t1" ON "users"."id" = "t1"."user_id"`,
			wantArgs: []any{"GitHub"},
		},
		{
			name: "schema/O2M/2types",
			input: func() *Step {
				step := NewStep(
					From("users", "id", sql.Select().From(sql.Table("users").Schema("s2")).Where(sql.EQ("name", "a8m"))),
					To("pets", "id"),
					Edge(O2M, false, "users", "owner_id"),
				)
				step.To.Schema = "s1"
				return step
			}(),
			wantQuery: `SELECT * FROM "s1"."pets" JOIN (SELECT "s2"."users"."id" FROM "s2"."users" WHERE "name" = $1) AS "t1" ON "s1"."pets"."owner_id" = "t1"."id"`,
			wantArgs:  []any{"a8m"},
		},
		{
			name: "schema/M2O/2types",
			input: func() *Step {
				step := NewStep(
					From("pets", "id", sql.Select().From(sql.Table("pets").Schema("s2")).Where(sql.EQ("name", "pedro"))),
					To("users", "id"),
					Edge(M2O, true, "pets", "owner_id"),
				)
				step.To.Schema = "s1"
				return step
			}(),
			wantQuery: `SELECT * FROM "s1"."users" JOIN (SELECT "s2"."pets"."owner_id" FROM "s2"."pets" WHERE "name" = $1) AS "t1" ON "s1"."users"."id" = "t1"."owner_id"`,
			wantArgs:  []any{"pedro"},
		},
		{
			name: "schema/M2M/2types",
			input: func() *Step {
				step := NewStep(
					From("users", "id", sql.Select().From(sql.Table("users").Schema("s2")).Where(sql.EQ("name", "a8m"))),
					To("groups", "id"),
					Edge(M2M, false, "user_groups", "user_id", "group_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s3"
				return step
			}(),
			wantQuery: `
SELECT *
FROM "s1"."groups"
JOIN
  (SELECT "s3"."user_groups"."group_id"
   FROM "s3"."user_groups"
   JOIN
     (SELECT "s2"."users"."id"
      FROM "s2"."users"
      WHERE "name" = $1) AS "t1" ON "s3"."user_groups"."user_id" = "t1"."id") AS "t1" ON "s1"."groups"."id" = "t1"."group_id"`,
			wantArgs: []any{"a8m"},
		},
		{
			name: "schema/M2M/2types/inverse",
			input: func() *Step {
				step := NewStep(
					From("groups", "id", sql.Select().From(sql.Table("groups").Schema("s2")).Where(sql.EQ("name", "GitHub"))),
					To("users", "id"),
					Edge(M2M, true, "user_groups", "user_id", "group_id"),
				)
				step.To.Schema = "s1"
				step.Edge.Schema = "s3"
				return step
			}(),
			wantQuery: `
SELECT *
FROM "s1"."users"
JOIN
  (SELECT "s3"."user_groups"."user_id"
   FROM "s3"."user_groups"
   JOIN
     (SELECT "s2"."groups"."id"
      FROM "s2"."groups"
      WHERE "name" = $1) AS "t1" ON "s3"."user_groups"."group_id" = "t1"."id") AS "t1" ON "s1"."users"."id" = "t1"."user_id"`,
			wantArgs: []any{"GitHub"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			selector := SetNeighbors("postgres", tt.input)
			query, args := selector.Query()
			tt.wantQuery = strings.Join(strings.Fields(tt.wantQuery), " ")
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}

func TestHasNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		step      *Step
		selector  *sql.Selector
		wantQuery string
	}{
		{
			name: "O2O/1type",
			// A nodes sql.Table; linked-list (next->prev). The "prev"
			// node holds association pointer. The neighbors query
			// here checks if a node "has-next".
			step: NewStep(
				From("nodes", "id"),
				To("nodes", "id"),
				Edge(O2O, false, "nodes", "prev_id"),
			),
			selector:  sql.Select("*").From(sql.Table("nodes")),
			wantQuery: "SELECT * FROM `nodes` WHERE EXISTS (SELECT `nodes_edge`.`prev_id` FROM `nodes` AS `nodes_edge` WHERE `nodes`.`id` = `nodes_edge`.`prev_id`)",
		},
		{
			name: "O2O/1type/inverse",
			// Same example as above, but the neighbors
			// query checks if a node "has-previous".
			step: NewStep(
				From("nodes", "id"),
				To("nodes", "id"),
				Edge(O2O, true, "nodes", "prev_id"),
			),
			selector:  sql.Select("*").From(sql.Table("nodes")),
			wantQuery: "SELECT * FROM `nodes` WHERE `nodes`.`prev_id` IS NOT NULL",
		},
		{
			name: "O2M/2type2",
			step: NewStep(
				From("users", "id"),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			selector:  sql.Select("*").From(sql.Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE EXISTS (SELECT `pets`.`owner_id` FROM `pets` WHERE `users`.`id` = `pets`.`owner_id`)",
		},
		{
			name: "M2O/2type2",
			step: NewStep(
				From("pets", "id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			selector:  sql.Select("*").From(sql.Table("pets")),
			wantQuery: "SELECT * FROM `pets` WHERE `pets`.`owner_id` IS NOT NULL",
		},
		{
			name: "M2M/2types",
			step: NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			selector:  sql.Select("*").From(sql.Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `user_groups`.`user_id` FROM `user_groups`)",
		},
		{
			name: "M2M/2types/inverse",
			step: NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, true, "group_users", "group_id", "user_id"),
			),
			selector:  sql.Select("*").From(sql.Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `group_users`.`user_id` FROM `group_users`)",
		},
		{
			name: "schema/O2O/1type",
			step: func() *Step {
				step := NewStep(
					From("nodes", "id"),
					To("nodes", "id"),
					Edge(O2O, false, "nodes", "prev_id"),
				)
				step.Edge.Schema = "s1"
				return step
			}(),
			selector:  sql.Select("*").From(sql.Table("nodes").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`nodes` WHERE EXISTS (SELECT `nodes_edge`.`prev_id` FROM `s1`.`nodes` AS `nodes_edge` WHERE `s1`.`nodes`.`id` = `nodes_edge`.`prev_id`)",
		},
		{
			name: "schema/O2O/1type/inverse",
			// Same example as above, but the neighbors
			// query checks if a node "has-previous".
			step: NewStep(
				From("nodes", "id"),
				To("nodes", "id"),
				Edge(O2O, true, "nodes", "prev_id"),
			),
			selector:  sql.Select("*").From(sql.Table("nodes").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`nodes` WHERE `s1`.`nodes`.`prev_id` IS NOT NULL",
		},
		{
			name: "schema/O2M/2type2",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("pets", "id"),
					Edge(O2M, false, "pets", "owner_id"),
				)
				step.Edge.Schema = "s2"
				return step
			}(),
			selector:  sql.Select("*").From(sql.Table("users").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`users` WHERE EXISTS (SELECT `s2`.`pets`.`owner_id` FROM `s2`.`pets` WHERE `s1`.`users`.`id` = `s2`.`pets`.`owner_id`)",
		},
		{
			name: "schema/M2O/2type2",
			step: NewStep(
				From("pets", "id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			selector:  sql.Select("*").From(sql.Table("pets").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`pets` WHERE `s1`.`pets`.`owner_id` IS NOT NULL",
		},
		{
			name: "schema/M2M/2types",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("groups", "id"),
					Edge(M2M, false, "user_groups", "user_id", "group_id"),
				)
				step.Edge.Schema = "s2"
				return step
			}(),
			selector:  sql.Select("*").From(sql.Table("users").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`users` WHERE `s1`.`users`.`id` IN (SELECT `s2`.`user_groups`.`user_id` FROM `s2`.`user_groups`)",
		},
		{
			name: "schema/M2M/2types/inverse",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("groups", "id"),
					Edge(M2M, true, "group_users", "group_id", "user_id"),
				)
				step.Edge.Schema = "s2"
				return step
			}(),
			selector:  sql.Select("*").From(sql.Table("users").Schema("s1")),
			wantQuery: "SELECT * FROM `s1`.`users` WHERE `s1`.`users`.`id` IN (SELECT `s2`.`group_users`.`user_id` FROM `s2`.`group_users`)",
		},
		{
			name: "O2M/2type2/selector",
			step: NewStep(
				From("users", "id"),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			selector:  sql.Select("*").From(sql.Select("*").From(sql.Table("users")).As("users")).As("users"),
			wantQuery: "SELECT * FROM (SELECT * FROM `users`) AS `users` WHERE EXISTS (SELECT `pets`.`owner_id` FROM `pets` WHERE `users`.`id` = `pets`.`owner_id`)",
		},
		{
			name: "M2O/2type2/selector",
			step: NewStep(
				From("pets", "id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			selector:  sql.Select("*").From(sql.Select("*").From(sql.Table("pets")).As("pets")).As("pets"),
			wantQuery: "SELECT * FROM (SELECT * FROM `pets`) AS `pets` WHERE `pets`.`owner_id` IS NOT NULL",
		},
		{
			name: "M2M/2types/selector",
			step: NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			selector:  sql.Select("*").From(sql.Select("*").From(sql.Table("users")).As("users")).As("users"),
			wantQuery: "SELECT * FROM (SELECT * FROM `users`) AS `users` WHERE `users`.`id` IN (SELECT `user_groups`.`user_id` FROM `user_groups`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range []*sql.Selector{tt.selector, tt.selector.Clone()} {
				HasNeighbors(s, tt.step)
				query, args := s.Query()
				require.Equal(t, tt.wantQuery, query)
				require.Empty(t, args)
			}
		})
	}
}

func TestHasNeighborsWith(t *testing.T) {
	tests := []struct {
		name      string
		step      *Step
		selector  *sql.Selector
		predicate func(*sql.Selector)
		wantQuery string
		wantArgs  []any
	}{
		{
			name: "O2O",
			step: NewStep(
				From("users", "id"),
				To("cards", "id"),
				Edge(O2O, false, "cards", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("users")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("expired", false))
			},
			wantQuery: `SELECT * FROM "users" WHERE EXISTS (SELECT "cards"."owner_id" FROM "cards" WHERE "users"."id" = "cards"."owner_id" AND NOT "expired")`,
		},
		{
			name: "O2O/inverse",
			step: NewStep(
				From("cards", "id"),
				To("users", "id"),
				Edge(O2O, true, "cards", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("cards")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "a8m"))
			},
			wantQuery: `SELECT * FROM "cards" WHERE EXISTS (SELECT "users"."id" FROM "users" WHERE "cards"."owner_id" = "users"."id" AND "name" = $1)`,
			wantArgs:  []any{"a8m"},
		},
		{
			name: "O2M",
			step: NewStep(
				From("users", "id"),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").
				From(sql.Table("users")).
				Where(sql.EQ("last_name", "mashraki")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "pedro"))
			},
			wantQuery: `SELECT * FROM "users" WHERE "last_name" = $1 AND EXISTS (SELECT "pets"."owner_id" FROM "pets" WHERE "users"."id" = "pets"."owner_id" AND "name" = $2)`,
			wantArgs:  []any{"mashraki", "pedro"},
		},
		{
			name: "M2O",
			step: NewStep(
				From("pets", "id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").
				From(sql.Table("pets")).
				Where(sql.EQ("name", "pedro")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("last_name", "mashraki"))
			},
			wantQuery: `SELECT * FROM "pets" WHERE "name" = $1 AND EXISTS (SELECT "users"."id" FROM "users" WHERE "pets"."owner_id" = "users"."id" AND "last_name" = $2)`,
			wantArgs:  []any{"pedro", "mashraki"},
		},
		{
			name: "M2M",
			step: NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("users")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "GitHub"))
			},
			wantQuery: `
SELECT *
FROM "users"
WHERE "users"."id" IN
  (SELECT "user_groups"."user_id"
  FROM "user_groups"
  JOIN "groups" AS "t1" ON "user_groups"."group_id" = "t1"."id" WHERE "name" = $1)`,
			wantArgs: []any{"GitHub"},
		},
		{
			name: "M2M/inverse",
			step: NewStep(
				From("groups", "id"),
				To("users", "id"),
				Edge(M2M, true, "user_groups", "user_id", "group_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("groups")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "a8m"))
			},
			wantQuery: `
SELECT *
FROM "groups"
WHERE "groups"."id" IN
  (SELECT "user_groups"."group_id"
  FROM "user_groups"
  JOIN "users" AS "t1" ON "user_groups"."user_id" = "t1"."id" WHERE "name" = $1)`,
			wantArgs: []any{"a8m"},
		},
		{
			name: "M2M/inverse",
			step: NewStep(
				From("groups", "id"),
				To("users", "id"),
				Edge(M2M, true, "user_groups", "user_id", "group_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("groups")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.And(sql.NotNull("name"), sql.EQ("name", "a8m")))
			},
			wantQuery: `
SELECT *
FROM "groups"
WHERE "groups"."id" IN
  (SELECT "user_groups"."group_id"
  FROM "user_groups"
  JOIN "users" AS "t1" ON "user_groups"."user_id" = "t1"."id" WHERE "name" IS NOT NULL AND "name" = $1)`,
			wantArgs: []any{"a8m"},
		},
		{
			name: "schema/O2O",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("cards", "id"),
					Edge(O2O, false, "cards", "owner_id"),
				)
				step.Edge.Schema = "s2"
				return step
			}(),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("users").Schema("s1")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("expired", false))
			},
			wantQuery: `SELECT * FROM "s1"."users" WHERE EXISTS (SELECT "s2"."cards"."owner_id" FROM "s2"."cards" WHERE "s1"."users"."id" = "s2"."cards"."owner_id" AND NOT "expired")`,
		},
		{
			name: "schema/O2M",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("pets", "id"),
					Edge(O2M, false, "pets", "owner_id"),
				)
				step.Edge.Schema = "s2"
				return step
			}(),
			selector: sql.Dialect("postgres").Select("*").
				From(sql.Table("users").Schema("s1")).
				Where(sql.EQ("last_name", "mashraki")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "pedro"))
			},
			wantQuery: `SELECT * FROM "s1"."users" WHERE "last_name" = $1 AND EXISTS (SELECT "s2"."pets"."owner_id" FROM "s2"."pets" WHERE "s1"."users"."id" = "s2"."pets"."owner_id" AND "name" = $2)`,
			wantArgs:  []any{"mashraki", "pedro"},
		},
		{
			name: "schema/M2M",
			step: func() *Step {
				step := NewStep(
					From("users", "id"),
					To("groups", "id"),
					Edge(M2M, false, "user_groups", "user_id", "group_id"),
				)
				step.To.Schema = "s3"
				step.Edge.Schema = "s2"
				return step
			}(),
			selector: sql.Dialect("postgres").Select("*").From(sql.Table("users").Schema("s1")),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "GitHub"))
			},
			wantQuery: `
SELECT *
FROM "s1"."users"
WHERE "s1"."users"."id" IN
  (SELECT "s2"."user_groups"."user_id"
  FROM "s2"."user_groups"
  JOIN "s3"."groups" AS "t1" ON "s2"."user_groups"."group_id" = "t1"."id" WHERE "name" = $1)`,
			wantArgs: []any{"GitHub"},
		},
		{
			name: "O2M/selector",
			step: NewStep(
				From("users", "id"),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").
				From(sql.Select("*").From(sql.Table("users")).As("users")).
				Where(sql.EQ("last_name", "mashraki")).As("users"),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "pedro"))
			},
			wantQuery: `SELECT * FROM (SELECT * FROM "users") AS "users" WHERE "last_name" = $1 AND EXISTS (SELECT "pets"."owner_id" FROM "pets" WHERE "users"."id" = "pets"."owner_id" AND "name" = $2)`,
			wantArgs:  []any{"mashraki", "pedro"},
		},
		{
			name: "M2O/selector",
			step: NewStep(
				From("pets", "id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			selector: sql.Dialect("postgres").Select("*").
				From(sql.Select("*").From(sql.Table("pets")).As("pets")).
				Where(sql.EQ("name", "pedro")).As("pets"),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("last_name", "mashraki"))
			},
			wantQuery: `SELECT * FROM (SELECT * FROM "pets") AS "pets" WHERE "name" = $1 AND EXISTS (SELECT "users"."id" FROM "users" WHERE "pets"."owner_id" = "users"."id" AND "last_name" = $2)`,
			wantArgs:  []any{"pedro", "mashraki"},
		},
		{
			name: "M2M/selector",
			step: NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			selector: sql.Dialect("postgres").Select("*").From(sql.Select("*").From(sql.Table("users")).As("users")).As("users"),
			predicate: func(s *sql.Selector) {
				s.Where(sql.EQ("name", "GitHub"))
			},
			wantQuery: `SELECT * FROM (SELECT * FROM "users") AS "users" WHERE "users"."id" IN (SELECT "user_groups"."user_id" FROM "user_groups" JOIN "groups" AS "t1" ON "user_groups"."group_id" = "t1"."id" WHERE "name" = $1)`,
			wantArgs:  []any{"GitHub"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for _, s := range []*sql.Selector{tt.selector, tt.selector.Clone()} {
				HasNeighborsWith(s, tt.step, tt.predicate)
				query, args := s.Query()
				tt.wantQuery = strings.Join(strings.Fields(tt.wantQuery), " ")
				require.Equal(t, tt.wantQuery, query)
				require.Equal(t, tt.wantArgs, args)
			}
		})
	}
}

func TestHasNeighborsWithContext(t *testing.T) {
	type key string
	ctx := context.WithValue(context.Background(), key("mykey"), "myval")
	for _, rel := range [...]Rel{M2M, O2M, O2O} {
		t.Run(rel.String(), func(t *testing.T) {
			sel := sql.Dialect(dialect.Postgres).
				Select("*").
				From(sql.Table("users")).
				WithContext(ctx)
			step := NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(rel, false, "user_groups", "user_id", "group_id"),
			)
			var called bool
			pred := func(s *sql.Selector) {
				called = true
				got := s.Context().Value(key("mykey")).(string)
				require.Equal(t, "myval", got)
			}
			HasNeighborsWith(sel, step, pred)
			require.True(t, called, "expected predicate function to be called")
		})
	}
}

func TestOrderByNeighborsCount(t *testing.T) {
	build := sql.Dialect(dialect.Postgres)
	t1 := build.Table("users")
	s := build.Select(t1.C("name")).
		From(t1)
	t.Run("O2M", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborsCount(s,
			NewStep(
				From("users", "id"),
				To("pets", "owner_id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			sql.OrderDesc(),
			sql.OrderAs("count_pets"),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" LEFT JOIN (SELECT "pets"."owner_id", COUNT(*) AS "count_pets" FROM "pets" GROUP BY "pets"."owner_id") AS "t1" ON "users"."id" = "t1"."owner_id" ORDER BY "t1"."count_pets" DESC NULLS LAST`, query)
	})
	t.Run("O2M/Selected", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborsCount(s,
			NewStep(
				From("users", "id"),
				To("pets", "owner_id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			sql.OrderDesc(),
			sql.OrderSelectAs("count_pets"),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name", "t1"."count_pets" FROM "users" LEFT JOIN (SELECT "pets"."owner_id", COUNT(*) AS "count_pets" FROM "pets" GROUP BY "pets"."owner_id") AS "t1" ON "users"."id" = "t1"."owner_id" ORDER BY "t1"."count_pets" DESC NULLS LAST`, query)
	})
	t.Run("M2M", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborsCount(s,
			NewStep(
				From("users", "id"),
				To("groups", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" LEFT JOIN (SELECT "user_groups"."user_id", COUNT(*) AS "count_groups" FROM "user_groups" GROUP BY "user_groups"."user_id") AS "t1" ON "users"."id" = "t1"."user_id" ORDER BY "t1"."count_groups" NULLS FIRST`, query)
	})
	// Zero or one.
	t.Run("M2O", func(t *testing.T) {
		s1, s2 := s.Clone(), s.Clone()
		OrderByNeighborsCount(s1,
			NewStep(
				From("pets", "owner_id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
		)
		query, args := s1.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" ORDER BY "owner_id" IS NULL`, query)

		OrderByNeighborsCount(s2,
			NewStep(
				From("pets", "owner_id"),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			sql.OrderDesc(),
		)
		query, args = s2.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" ORDER BY "owner_id" IS NOT NULL`, query)
	})
}

func TestOrderByNeighborTerms(t *testing.T) {
	build := sql.Dialect(dialect.Postgres)
	t1 := build.Table("users")
	s := build.Select(t1.C("name")).
		From(t1)
	t.Run("M2O", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("workplace", "id"),
				Edge(M2O, true, "users", "workplace_id"),
			),
			sql.OrderByField("name"),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" LEFT JOIN (SELECT "workplace"."id", "workplace"."name" FROM "workplace") AS "t1" ON "users"."workplace_id" = "t1"."id" ORDER BY "t1"."name" NULLS FIRST`, query)
	})
	t.Run("M2O/SelectedAs", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("workplace", "id"),
				Edge(M2O, true, "users", "workplace_id"),
			),
			sql.OrderByField(
				"name",
				sql.OrderSelectAs("workplace_name"),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name", "t1"."workplace_name" FROM "users" LEFT JOIN (SELECT "workplace"."id", "workplace"."name" AS "workplace_name" FROM "workplace") AS "t1" ON "users"."workplace_id" = "t1"."id" ORDER BY "t1"."workplace_name" NULLS FIRST`, query)
	})
	t.Run("M2O/NullsLast", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("workplace", "id"),
				Edge(M2O, true, "users", "workplace_id"),
			),
			sql.OrderByField(
				"name",
				sql.OrderNullsLast(),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" LEFT JOIN (SELECT "workplace"."id", "workplace"."name" FROM "workplace") AS "t1" ON "users"."workplace_id" = "t1"."id" ORDER BY "t1"."name" NULLS LAST`, query)
	})
	t.Run("O2M", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("repos", "id"),
				Edge(O2M, false, "repo", "owner_id"),
			),
			sql.OrderBySum(
				"num_stars",
				sql.OrderSelectAs("total_stars"),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name", "t1"."total_stars" FROM "users" LEFT JOIN (SELECT "repo"."owner_id", SUM("repo"."num_stars") AS "total_stars" FROM "repo" GROUP BY "repo"."owner_id") AS "t1" ON "users"."id" = "t1"."owner_id" ORDER BY "t1"."total_stars" NULLS FIRST`, query)
	})
	t.Run("M2M", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("group", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			sql.OrderBySum(
				"num_users",
				sql.OrderSelectAs("total_users"),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name", "t1"."total_users" FROM "users" LEFT JOIN (SELECT "user_id", SUM("group"."num_users") AS "total_users" FROM "group" JOIN "user_groups" AS "t1" ON "group"."id" = "t1"."group_id" GROUP BY "user_id") AS "t1" ON "users"."id" = "t1"."user_id" ORDER BY "t1"."total_users" NULLS FIRST`, query)
	})
	t.Run("M2M/NullsLast", func(t *testing.T) {
		s := s.Clone()
		OrderByNeighborTerms(s,
			NewStep(
				From("users", "id"),
				To("group", "id"),
				Edge(M2M, false, "user_groups", "user_id", "group_id"),
			),
			sql.OrderBySum(
				"num_users",
				sql.OrderAs("total_users"),
				sql.OrderNullsLast(),
			),
		)
		query, args := s.Query()
		require.Empty(t, args)
		require.Equal(t, `SELECT "users"."name" FROM "users" LEFT JOIN (SELECT "user_id", SUM("group"."num_users") AS "total_users" FROM "group" JOIN "user_groups" AS "t1" ON "group"."id" = "t1"."group_id" GROUP BY "user_id") AS "t1" ON "users"."id" = "t1"."user_id" ORDER BY "t1"."total_users" NULLS LAST`, query)
	})
}

func TestCreateNode(t *testing.T) {
	tests := []struct {
		name    string
		spec    *CreateSpec
		expect  func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "fields",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "age", Type: field.TypeInt, Value: 30},
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `users` (`age`, `name`) VALUES (?, ?)")).
					WithArgs(30, "a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "modifiers",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "age", Type: field.TypeInt, Value: 30},
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
				OnConflict: []sql.ConflictOption{
					sql.ResolveWithNewValues(),
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `users` (`age`, `name`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `age` = VALUES(`age`), `name` = VALUES(`name`), `id` = LAST_INSERT_ID(`users`.`id`)")).
					WithArgs(30, "a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "fields/user-defined-id",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Value: 1},
				Fields: []*FieldSpec{
					{Column: "age", Type: field.TypeInt, Value: 30},
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `users` (`age`, `name`, `id`) VALUES (?, ?, ?)")).
					WithArgs(30, "a8m", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "fields/json",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "json", Type: field.TypeJSON, Value: struct{}{}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `users` (`json`) VALUES (?)")).
					WithArgs([]byte("{}")).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "edges/m2o",
			spec: &CreateSpec{
				Table: "pets",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "pedro"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2O, Columns: []string{"owner_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `pets` (`name`, `owner_id`) VALUES (?, ?)")).
					WithArgs("pedro", 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "edges/o2o/inverse",
			spec: &CreateSpec{
				Table: "cards",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "number", Type: field.TypeString, Value: "0001"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2O, Columns: []string{"owner_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `cards` (`number`, `owner_id`) VALUES (?, ?)")).
					WithArgs("0001", 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
		{
			name: "edges/o2m",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("UPDATE `pets` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/o2m",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2, 3, 4}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("UPDATE `pets` SET `owner_id` = ? WHERE `id` IN (?, ?, ?) AND `owner_id` IS NULL")).
					WithArgs(1, 2, 3, 4).
					WillReturnResult(sqlmock.NewResult(1, 3))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/o2o",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2O, Table: "cards", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("UPDATE `cards` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/o2o/bidi",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2O, Bidi: true, Table: "users", Columns: []string{"spouse_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`, `spouse_id`) VALUES (?, ?)")).
					WithArgs("a8m", 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("UPDATE `users` SET `spouse_id` = ? WHERE `id` = ? AND `spouse_id` IS NULL")).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m",
			spec: &CreateSpec{
				Table: "groups",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "GitHub"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `groups` (`name`) VALUES (?)")).
					WithArgs("GitHub").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `group_id` = `group_users`.`group_id`, `user_id` = `group_users`.`user_id`")).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/fields",
			spec: &CreateSpec{
				Table: "groups",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "GitHub"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}, Fields: []*FieldSpec{{Column: "ts", Type: field.TypeInt, Value: 3}}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `groups` (`name`) VALUES (?)")).
					WithArgs("GitHub").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`, `ts`) VALUES (?, ?, ?)")).
					WithArgs(1, 2, 3).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/inverse",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "mashraki"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("mashraki").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?) ON DUPLICATE KEY UPDATE `group_id` = `group_users`.`group_id`, `user_id` = `group_users`.`user_id`")).
					WithArgs(2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/bidi",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "mashraki"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Bidi: true, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("mashraki").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `user_id` = `user_friends`.`user_id`, `friend_id` = `user_friends`.`friend_id`")).
					WithArgs(1, 2, 2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/bidi/fields",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "mashraki"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Bidi: true, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}, Fields: []*FieldSpec{{Column: "ts", Type: field.TypeInt, Value: 3}}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("mashraki").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`, `ts`) VALUES (?, ?, ?), (?, ?, ?)")).
					WithArgs(1, 2, 3, 2, 1, 3).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/bidi/batch",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "mashraki"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2M, Bidi: true, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
					{Rel: M2M, Bidi: true, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{3}, IDSpec: &FieldSpec{Column: "id"}}},
					{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{4}, IDSpec: &FieldSpec{Column: "id"}}},
					{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{5}, IDSpec: &FieldSpec{Column: "id"}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?)")).
					WithArgs("mashraki").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `group_id` = `group_users`.`group_id`, `user_id` = `group_users`.`user_id`")).
					WithArgs(4, 1, 5, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `user_id` = `user_friends`.`user_id`, `friend_id` = `user_friends`.`friend_id`")).
					WithArgs(1, 2, 2, 1, 1, 3, 3, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "schema",
			spec: &CreateSpec{
				Table:  "users",
				Schema: "test",
				ID:     &FieldSpec{Column: "id", Type: field.TypeInt},
				Fields: []*FieldSpec{
					{Column: "age", Type: field.TypeInt, Value: 30},
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `test`.`users` (`age`, `name`) VALUES (?, ?)")).
					WithArgs(30, "a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.expect(mock)
			err = CreateNode(context.Background(), sql.OpenDB(dialect.MySQL, db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

func TestBatchCreate(t *testing.T) {
	tests := []struct {
		name    string
		spec    *BatchCreateSpec
		expect  func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "empty",
			spec: &BatchCreateSpec{},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectCommit()
			},
		},
		{
			name: "fields with modifiers",
			spec: &BatchCreateSpec{
				Nodes: []*CreateSpec{
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 32},
							{Column: "name", Type: field.TypeString, Value: "a8m"},
							{Column: "active", Type: field.TypeBool, Value: false},
						},
					},
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 30},
							{Column: "name", Type: field.TypeString, Value: "nati"},
							{Column: "active", Type: field.TypeBool, Value: true},
						},
					},
				},
				OnConflict: []sql.ConflictOption{
					sql.ResolveWithIgnore(),
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectExec(escape("INSERT INTO `users` (`active`, `age`, `name`) VALUES (?, ?, ?), (?, ?, ?) ON DUPLICATE KEY UPDATE `active` = `users`.`active`, `age` = `users`.`age`, `name` = `users`.`name`")).
					WithArgs(false, 32, "a8m", true, 30, "nati").
					WillReturnResult(sqlmock.NewResult(10, 2))
			},
		},
		{
			name: "no tx",
			spec: &BatchCreateSpec{
				Nodes: []*CreateSpec{
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 32},
							{Column: "name", Type: field.TypeString, Value: "a8m"},
							{Column: "active", Type: field.TypeBool, Value: false},
						},
						Edges: []*EdgeSpec{
							{Rel: M2O, Table: "company", Columns: []string{"workplace_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
							{Rel: O2O, Inverse: true, Table: "users", Columns: []string{"best_friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{3}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 30},
							{Column: "name", Type: field.TypeString, Value: "nati"},
						},
						Edges: []*EdgeSpec{
							{Rel: M2O, Table: "company", Columns: []string{"workplace_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
							{Rel: O2O, Inverse: true, Table: "users", Columns: []string{"best_friend_id"}, Target: &EdgeTarget{Nodes: []driver.Value{4}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				// Insert nodes with FKs.
				m.ExpectExec(escape("INSERT INTO `users` (`active`, `age`, `best_friend_id`, `name`, `workplace_id`) VALUES (?, ?, ?, ?, ?), (NULL, ?, ?, ?, ?)")).
					WithArgs(false, 32, 3, "a8m", 2, 30, 4, "nati", 2).
					WillReturnResult(sqlmock.NewResult(10, 2))
			},
		},
		{
			name: "with tx",
			spec: &BatchCreateSpec{
				Nodes: []*CreateSpec{
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "name", Type: field.TypeString, Value: "a8m"},
						},
						Edges: []*EdgeSpec{
							{Rel: O2O, Table: "cards", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{3}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "name", Type: field.TypeString, Value: "nati"},
						},
						Edges: []*EdgeSpec{
							{Rel: O2O, Table: "cards", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{4}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`name`) VALUES (?), (?)")).
					WithArgs("a8m", "nati").
					WillReturnResult(sqlmock.NewResult(10, 2))
				m.ExpectExec(escape("UPDATE `cards` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(10 /* LAST_INSERT_ID() */, 3).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("UPDATE `cards` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(11 /* LAST_INSERT_ID() + 1 */, 4).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "multiple",
			spec: &BatchCreateSpec{
				Nodes: []*CreateSpec{
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 32},
							{Column: "name", Type: field.TypeString, Value: "a8m"},
							{Column: "active", Type: field.TypeBool, Value: false},
						},
						Edges: []*EdgeSpec{
							{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
							{Rel: M2M, Table: "user_products", Columns: []string{"user_id", "product_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
							{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{2}}},
							{Rel: M2O, Table: "company", Columns: []string{"workplace_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
							{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
					{
						Table: "users",
						ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
						Fields: []*FieldSpec{
							{Column: "age", Type: field.TypeInt, Value: 30},
							{Column: "name", Type: field.TypeString, Value: "nati"},
						},
						Edges: []*EdgeSpec{
							{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
							{Rel: M2M, Table: "user_products", Columns: []string{"user_id", "product_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
							{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{2}}},
							{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{3}, IDSpec: &FieldSpec{Column: "id"}}},
						},
					},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				// Insert nodes with FKs.
				m.ExpectExec(escape("INSERT INTO `users` (`active`, `age`, `name`, `workplace_id`) VALUES (?, ?, ?, ?), (NULL, ?, ?, NULL)")).
					WithArgs(false, 32, "a8m", 2, 30, "nati").
					WillReturnResult(sqlmock.NewResult(10, 2))
				// Insert M2M inverse-edges.
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `group_id` = `group_users`.`group_id`, `user_id` = `group_users`.`user_id`")).
					WithArgs(2, 10, 2, 11).
					WillReturnResult(sqlmock.NewResult(2, 2))
				// Insert M2M bidirectional edges.
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `user_id` = `user_friends`.`user_id`, `friend_id` = `user_friends`.`friend_id`")).
					WithArgs(10, 2, 2, 10, 11, 2, 2, 11).
					WillReturnResult(sqlmock.NewResult(2, 2))
				// Insert M2M edges.
				m.ExpectExec(escape("INSERT INTO `user_products` (`user_id`, `product_id`) VALUES (?, ?), (?, ?) ON DUPLICATE KEY UPDATE `user_id` = `user_products`.`user_id`, `product_id` = `user_products`.`product_id`")).
					WithArgs(10, 2, 11, 2).
					WillReturnResult(sqlmock.NewResult(2, 2))
				// Update FKs exist in different tables.
				m.ExpectExec(escape("UPDATE `pets` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(10 /* id of the 1st new node */, 2 /* pet id */).
					WillReturnResult(sqlmock.NewResult(2, 2))
				m.ExpectExec(escape("UPDATE `pets` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(11 /* id of the 2nd new node */, 3 /* pet id */).
					WillReturnResult(sqlmock.NewResult(2, 2))
				m.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.expect(mock)
			err = BatchCreate(context.Background(), sql.OpenDB("mysql", db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

type user struct {
	id    int
	age   int
	name  string
	edges struct {
		fk1 int
		fk2 int
	}
}

func (*user) values(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch c := columns[i]; c {
		case "id", "age", "fk1", "fk2":
			values[i] = &sql.NullInt64{}
		case "name":
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q", c)
		}
	}
	return values, nil
}

func (u *user) assign(columns []string, values []any) error {
	if len(columns) != len(values) {
		return fmt.Errorf("mismatch number of values")
	}
	for i, c := range columns {
		switch c {
		case "id":
			u.id = int(values[i].(*sql.NullInt64).Int64)
		case "age":
			u.age = int(values[i].(*sql.NullInt64).Int64)
		case "name":
			u.name = values[i].(*sql.NullString).String
		case "fk1":
			u.edges.fk1 = int(values[i].(*sql.NullInt64).Int64)
		case "fk2":
			u.edges.fk2 = int(values[i].(*sql.NullInt64).Int64)
		default:
			return fmt.Errorf("unknown column %q", c)
		}
	}
	return nil
}

func TestUpdateNode(t *testing.T) {
	tests := []struct {
		name     string
		spec     *UpdateSpec
		prepare  func(sqlmock.Sqlmock)
		wantErr  bool
		wantUser *user
	}{
		{
			name: "fields/set",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Fields: FieldMut{
					Set: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 30},
						{Column: "name", Type: field.TypeString, Value: "Ariel"},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `age` = ?, `name` = ? WHERE `id` = ?")).
					WithArgs(30, "Ariel", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 30, "Ariel"))
				mock.ExpectCommit()
			},
			wantUser: &user{name: "Ariel", age: 30, id: 1},
		},
		{
			name: "fields/set_modifier",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Modifiers: []func(*sql.UpdateBuilder){
					func(u *sql.UpdateBuilder) {
						u.Set("name", sql.Expr(sql.Lower("name")))
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `name` = LOWER(`name`) WHERE `id` = ?")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 30, "Ariel"))
				mock.ExpectCommit()
			},
			wantUser: &user{name: "Ariel", age: 30, id: 1},
		},
		{
			name: "fields/add_set_clear",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Predicate: func(s *sql.Selector) {
					s.Where(sql.EQ("deleted", false))
				},
				Fields: FieldMut{
					Add: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 1},
					},
					Set: []*FieldSpec{
						{Column: "deleted", Type: field.TypeBool, Value: true},
					},
					Clear: []*FieldSpec{
						{Column: "name", Type: field.TypeString},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `name` = NULL, `deleted` = ?, `age` = COALESCE(`users`.`age`, 0) + ? WHERE `id` = ? AND NOT `deleted`")).
					WithArgs(true, 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 31, nil))
				mock.ExpectCommit()
			},
			wantUser: &user{age: 31, id: 1},
		},
		{
			name: "fields/ensure_exists",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Predicate: func(s *sql.Selector) {
					s.Where(sql.EQ("deleted", false))
				},
				Fields: FieldMut{
					Add: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 1},
					},
					Set: []*FieldSpec{
						{Column: "deleted", Type: field.TypeBool, Value: true},
					},
					Clear: []*FieldSpec{
						{Column: "name", Type: field.TypeString},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `name` = NULL, `deleted` = ?, `age` = COALESCE(`users`.`age`, 0) + ? WHERE `id` = ? AND NOT `deleted`")).
					WithArgs(true, 1, 1).
					WillReturnResult(sqlmock.NewResult(0, 0))
				mock.ExpectQuery(escape("SELECT EXISTS (SELECT * FROM `users` WHERE `id` = ? AND NOT `deleted`)")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"exists"}).
						AddRow(false))
				mock.ExpectRollback()
			},
			wantErr:  true,
			wantUser: &user{},
		},
		{
			name: "edges/o2o_non_inverse and m2o",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: O2O, Columns: []string{"car_id"}, Inverse: true},
						{Rel: M2O, Columns: []string{"workplace_id"}, Inverse: true},
					},
					Add: []*EdgeSpec{
						{Rel: O2O, Columns: []string{"card_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
						{Rel: M2O, Columns: []string{"parent_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `workplace_id` = NULL, `car_id` = NULL, `parent_id` = ?, `card_id` = ? WHERE `id` = ?")).
					WithArgs(2, 2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 31, nil))
				mock.ExpectCommit()
			},
			wantUser: &user{age: 31, id: 1},
		},
		{
			name: "edges/o2o_bidi",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: O2O, Table: "users", Bidi: true, Columns: []string{"partner_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
						{Rel: O2O, Table: "users", Bidi: true, Columns: []string{"spouse_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{2}}},
					},
					Add: []*EdgeSpec{
						{Rel: O2O, Table: "users", Bidi: true, Columns: []string{"spouse_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{3}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				// Clear the "partner" from 1's column, and set "spouse 3".
				// "spouse 2" is implicitly removed when setting a different foreign-key.
				mock.ExpectExec(escape("UPDATE `users` SET `partner_id` = NULL, `spouse_id` = ? WHERE `id` = ?")).
					WithArgs(3, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Clear the "partner_id" column from previous 1's partner.
				mock.ExpectExec(escape("UPDATE `users` SET `partner_id` = NULL WHERE `partner_id` = ?")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Clear "spouse 1" from 3's column.
				mock.ExpectExec(escape("UPDATE `users` SET `spouse_id` = NULL WHERE `id` = ? AND `spouse_id` = ?")).
					WithArgs(2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Set 3's column to point "spouse 1".
				mock.ExpectExec(escape("UPDATE `users` SET `spouse_id` = ? WHERE `id` = ? AND `spouse_id` IS NULL")).
					WithArgs(1, 3).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 31, nil))
				mock.ExpectCommit()
			},
			wantUser: &user{age: 31, id: 1},
		},
		{
			name: "edges/clear_add_m2m",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{2}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{3, 7}}},
						// Clear all "following" edges (and their inverse).
						{Rel: M2M, Table: "user_following", Bidi: true, Columns: []string{"following_id", "follower_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
						// Clear all "user_blocked" edges.
						{Rel: M2M, Table: "user_blocked", Columns: []string{"user_id", "blocked_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
						// Clear all "comments" edges.
						{Rel: M2M, Inverse: true, Table: "comment_responders", Columns: []string{"comment_id", "responder_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
					},
					Add: []*EdgeSpec{
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{4}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{5}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id", Type: field.TypeInt}, Nodes: []driver.Value{6, 8}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				// Clear comment responders.
				mock.ExpectExec(escape("DELETE FROM `comment_responders` WHERE `responder_id` = ?")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Remove user groups.
				mock.ExpectExec(escape("DELETE FROM `group_users` WHERE `user_id` = ? AND `group_id` IN (?, ?)")).
					WithArgs(1, 3, 7).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Clear all blocked users.
				mock.ExpectExec(escape("DELETE FROM `user_blocked` WHERE `user_id` = ?")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Clear all user following.
				mock.ExpectExec(escape("DELETE FROM `user_following` WHERE `following_id` = ? OR `follower_id` = ?")).
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(1, 2))
				// Clear user friends.
				mock.ExpectExec(escape("DELETE FROM `user_friends` WHERE (`user_id` = ? AND `friend_id` = ?) OR (`friend_id` = ? AND `user_id` = ?)")).
					WithArgs(1, 2, 1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Add new groups.
				mock.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?), (?, ?)")).
					WithArgs(5, 1, 6, 1, 8, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				// Add new friends.
				mock.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(1, 4, 4, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 31, nil))
				mock.ExpectCommit()
			},
			wantUser: &user{age: 31, id: 1},
		},
		{
			name: "schema/fields/set",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Schema:  "mydb",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Fields: FieldMut{
					Set: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 30},
						{Column: "name", Type: field.TypeString, Value: "Ariel"},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `mydb`.`users` SET `age` = ?, `name` = ? WHERE `id` = ?")).
					WithArgs(30, "Ariel", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectQuery(escape("SELECT `id`, `name`, `age` FROM `mydb`.`users` WHERE `id` = ?")).
					WithArgs(1).
					WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name"}).
						AddRow(1, 30, "Ariel"))
				mock.ExpectCommit()
			},
			wantUser: &user{name: "Ariel", age: 30, id: 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.prepare(mock)
			usr := &user{}
			tt.spec.Assign = usr.assign
			tt.spec.ScanValues = usr.values
			err = UpdateNode(context.Background(), sql.OpenDB("", db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
			require.Equal(t, tt.wantUser, usr)
		})
	}
}

func TestExecUpdateNode(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectBegin()
	mock.ExpectExec(escape("UPDATE `users` SET `age` = ?, `name` = ? WHERE `id` = ?")).
		WithArgs(30, "Ariel", 1).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = UpdateNode(context.Background(), sql.OpenDB("", db), &UpdateSpec{
		Node: &NodeSpec{
			Table:   "users",
			Columns: []string{"id", "name", "age"},
			ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
		},
		Fields: FieldMut{
			Set: []*FieldSpec{
				{Column: "age", Type: field.TypeInt, Value: 30},
				{Column: "name", Type: field.TypeString, Value: "Ariel"},
			},
		},
	})
	require.NoError(t, err)
}

func TestUpdateNodes(t *testing.T) {
	tests := []struct {
		name         string
		spec         *UpdateSpec
		prepare      func(sqlmock.Sqlmock)
		wantErr      bool
		wantAffected int
	}{
		{
			name: "without predicate",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Fields: FieldMut{
					Set: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 30},
						{Column: "name", Type: field.TypeString, Value: "Ariel"},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// Apply field changes.
				mock.ExpectExec(escape("UPDATE `users` SET `age` = ?, `name` = ?")).
					WithArgs(30, "Ariel").
					WillReturnResult(sqlmock.NewResult(0, 2))
			},
			wantAffected: 2,
		},
		{
			name: "with predicate",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Fields: FieldMut{
					Clear: []*FieldSpec{
						{Column: "age", Type: field.TypeInt},
						{Column: "name", Type: field.TypeString},
					},
				},
				Predicate: func(s *sql.Selector) {
					s.Where(sql.EQ("name", "a8m"))
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// Clear fields.
				mock.ExpectExec(escape("UPDATE `users` SET `age` = NULL, `name` = NULL WHERE `name` = ?")).
					WithArgs("a8m").
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantAffected: 1,
		},
		{
			name: "with modifier",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Modifiers: []func(*sql.UpdateBuilder){
					func(u *sql.UpdateBuilder) {
						u.Set("id", sql.Expr("id + 1")).OrderBy("id")
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(escape("UPDATE `users` SET `id` = id + 1 ORDER BY `id`")).
					WillReturnResult(sqlmock.NewResult(0, 1))
			},
			wantAffected: 1,
		},
		{
			name: "own_fks/m2o_o2o_inverse",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: O2O, Columns: []string{"car_id"}, Inverse: true},
						{Rel: M2O, Columns: []string{"workplace_id"}, Inverse: true},
					},
					Add: []*EdgeSpec{
						{Rel: O2O, Columns: []string{"card_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{3}}},
						{Rel: M2O, Columns: []string{"parent_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{4}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				// Clear "car" and "workplace" foreign_keys and add "card" and a "parent".
				mock.ExpectExec(escape("UPDATE `users` SET `workplace_id` = NULL, `car_id` = NULL, `parent_id` = ?, `card_id` = ?")).
					WithArgs(4, 3).
					WillReturnResult(sqlmock.NewResult(0, 3))
			},
			wantAffected: 3,
		},
		{
			name: "o2m",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Fields: FieldMut{
					Add: []*FieldSpec{
						{Column: "version", Type: field.TypeInt, Value: 1},
					},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: O2M, Table: "cards", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{20, 30}, IDSpec: &FieldSpec{Column: "id"}}},
					},
					Add: []*EdgeSpec{
						{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{40}, IDSpec: &FieldSpec{Column: "id"}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users`")).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(10))
				mock.ExpectExec(escape("UPDATE `users` SET `version` = COALESCE(`users`.`version`, 0) + ? WHERE `id` = ?")).
					WithArgs(1, 10).
					WillReturnResult(sqlmock.NewResult(0, 1))
				// Clear "owner_id" column in the "cards" table.
				mock.ExpectExec(escape("UPDATE `cards` SET `owner_id` = NULL WHERE `id` IN (?, ?) AND `owner_id` = ?")).
					WithArgs(20, 30, 10).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Set "owner_id" column in the "pets" table.
				mock.ExpectExec(escape("UPDATE `pets` SET `owner_id` = ? WHERE `id` = ? AND `owner_id` IS NULL")).
					WithArgs(10, 40).
					WillReturnResult(sqlmock.NewResult(0, 2))
				mock.ExpectCommit()
			},
			wantAffected: 1,
		},
		{
			name: "m2m_one",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2, 3}}},
						{Rel: M2M, Table: "user_followers", Columns: []string{"user_id", "follower_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{5, 6}}},
						{Rel: M2M, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{4}}},
					},
					Add: []*EdgeSpec{
						{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{7, 8}}},
						{Rel: M2M, Table: "user_followers", Columns: []string{"user_id", "follower_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{9}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users`")).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
				// Clear user's groups.
				mock.ExpectExec(escape("DELETE FROM `group_users` WHERE `user_id` = ? AND `group_id` IN (?, ?)")).
					WithArgs(1, 2, 3).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Clear user's followers.
				mock.ExpectExec(escape("DELETE FROM `user_followers` WHERE (`user_id` = ? AND `follower_id` IN (?, ?)) OR (`follower_id` = ? AND `user_id` IN (?, ?))")).
					WithArgs(1, 5, 6, 1, 5, 6).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Clear user's friends.
				mock.ExpectExec(escape("DELETE FROM `user_friends` WHERE (`user_id` = ? AND `friend_id` = ?) OR (`friend_id` = ? AND `user_id` = ?)")).
					WithArgs(1, 4, 1, 4).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Attach new groups to user.
				mock.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(7, 1, 8, 1).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Attach new friends to user.
				mock.ExpectExec(escape("INSERT INTO `user_followers` (`user_id`, `follower_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(1, 9, 9, 1).
					WillReturnResult(sqlmock.NewResult(0, 2))
				mock.ExpectCommit()
			},
			wantAffected: 1,
		},
		{
			name: "m2m_many",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table: "users",
					ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
				},
				Edges: EdgeMut{
					Clear: []*EdgeSpec{
						{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2, 3}}},
						{Rel: M2M, Table: "user_followers", Columns: []string{"user_id", "follower_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{5, 6}}},
						{Rel: M2M, Table: "user_friends", Columns: []string{"user_id", "friend_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{4}}},
					},
					Add: []*EdgeSpec{
						{Rel: M2M, Table: "group_users", Columns: []string{"group_id", "user_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{7, 8}}},
						{Rel: M2M, Table: "user_followers", Columns: []string{"user_id", "follower_id"}, Bidi: true, Target: &EdgeTarget{Nodes: []driver.Value{9}}},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users`")).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(10).
						AddRow(20))
				// Clear user's groups.
				mock.ExpectExec(escape("DELETE FROM `group_users` WHERE `user_id` IN (?, ?) AND `group_id` IN (?, ?)")).
					WithArgs(10, 20, 2, 3).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Clear user's followers.
				mock.ExpectExec(escape("DELETE FROM `user_followers` WHERE (`user_id` IN (?, ?) AND `follower_id` IN (?, ?)) OR (`follower_id` IN (?, ?) AND `user_id` IN (?, ?))")).
					WithArgs(10, 20, 5, 6, 10, 20, 5, 6).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Clear user's friends.
				mock.ExpectExec(escape("DELETE FROM `user_friends` WHERE (`user_id` IN (?, ?) AND `friend_id` = ?) OR (`friend_id` IN (?, ?) AND `user_id` = ?)")).
					WithArgs(10, 20, 4, 10, 20, 4).
					WillReturnResult(sqlmock.NewResult(0, 2))
				// Attach new groups to user.
				mock.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?)")).
					WithArgs(7, 10, 7, 20, 8, 10, 8, 20).
					WillReturnResult(sqlmock.NewResult(0, 4))
				// Attach new friends to user.
				mock.ExpectExec(escape("INSERT INTO `user_followers` (`user_id`, `follower_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?)")).
					WithArgs(10, 9, 9, 10, 20, 9, 9, 20).
					WillReturnResult(sqlmock.NewResult(0, 4))
				mock.ExpectCommit()
			},
			wantAffected: 2,
		},
		{
			name: "m2m_edge_schema",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:       "users",
					CompositeID: []*FieldSpec{{Column: "user_id", Type: field.TypeInt}, {Column: "group_id", Type: field.TypeInt}},
				},
				Predicate: func(s *sql.Selector) {
					s.Where(sql.EQ("version", 1))
				},
				Fields: FieldMut{
					Add: []*FieldSpec{
						{Column: "version", Type: field.TypeInt, Value: 1},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectExec(escape("UPDATE `users` SET `version` = COALESCE(`users`.`version`, 0) + ? WHERE `version` = ?")).
					WithArgs(1, 1).
					WillReturnResult(sqlmock.NewResult(0, 4))
			},
			wantAffected: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.prepare(mock)
			affected, err := UpdateNodes(context.Background(), sql.OpenDB("", db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
			require.Equal(t, tt.wantAffected, affected)
		})
	}
}

func TestDeleteNodes(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectExec(escape("DELETE FROM `users`")).
		WillReturnResult(sqlmock.NewResult(0, 2))
	affected, err := DeleteNodes(context.Background(), sql.OpenDB("", db), &DeleteSpec{
		Node: &NodeSpec{
			Table: "users",
			ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 2, affected)
}

func TestDeleteNodesSchema(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectExec(escape("DELETE FROM `mydb`.`users`")).
		WillReturnResult(sqlmock.NewResult(0, 2))
	affected, err := DeleteNodes(context.Background(), sql.OpenDB("", db), &DeleteSpec{
		Node: &NodeSpec{
			Table:  "users",
			Schema: "mydb",
			ID:     &FieldSpec{Column: "id", Type: field.TypeInt},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 2, affected)
}

func TestQueryNodes(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectQuery(escape("SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name`, `users`.`fk1`, `users`.`fk2` FROM `users` WHERE `age` < ? ORDER BY `id` LIMIT 3 OFFSET 4 FOR UPDATE NOWAIT")).
		WithArgs(40).
		WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name", "fk1", "fk2"}).
			AddRow(1, 10, nil, nil, nil).
			AddRow(2, 20, "", 0, 0).
			AddRow(3, 30, "a8m", 1, 1))
	mock.ExpectQuery(escape("SELECT COUNT(DISTINCT `users`.`id`) FROM `users` WHERE `age` < ? LIMIT 3 OFFSET 4 FOR UPDATE NOWAIT")).
		WithArgs(40).
		WillReturnRows(sqlmock.NewRows([]string{"COUNT"}).
			AddRow(3))
	mock.ExpectQuery(escape("SELECT COUNT(DISTINCT `users`.`name`) FROM `users` WHERE `age` < ? LIMIT 3 OFFSET 4 FOR UPDATE NOWAIT")).
		WithArgs(40).
		WillReturnRows(sqlmock.NewRows([]string{"COUNT"}).
			AddRow(3))

	var (
		users []*user
		spec  = &QuerySpec{
			Node: &NodeSpec{
				Table:   "users",
				Columns: []string{"id", "age", "name", "fk1", "fk2"},
				ID:      &FieldSpec{Column: "id", Type: field.TypeInt},
			},
			Limit:  3,
			Offset: 4,
			Unique: true,
			Order: func(s *sql.Selector) {
				s.OrderBy("id")
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.LT("age", 40))
			},
			Modifiers: []func(*sql.Selector){
				func(s *sql.Selector) { s.ForUpdate(sql.WithLockAction(sql.NoWait)) },
			},
			ScanValues: func(columns []string) ([]any, error) {
				u := &user{}
				users = append(users, u)
				return u.values(columns)
			},
			Assign: func(columns []string, values []any) error {
				return users[len(users)-1].assign(columns, values)
			},
		}
	)

	// Query and scan.
	err = QueryNodes(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, &user{id: 1, age: 10, name: ""}, users[0])
	require.Equal(t, &user{id: 2, age: 20, name: ""}, users[1])
	require.Equal(t, &user{id: 3, age: 30, name: "a8m", edges: struct{ fk1, fk2 int }{1, 1}}, users[2])

	// Count nodes.
	spec.Node.Columns = nil
	n, err := CountNodes(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, 3, n)

	// Count nodes.
	spec.Node.Columns = []string{"name"}
	n, err = CountNodes(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, 3, n)
}

func TestQueryNodesSchema(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectQuery(escape("SELECT DISTINCT `mydb`.`users`.`id`, `mydb`.`users`.`age`, `mydb`.`users`.`name`, `mydb`.`users`.`fk1`, `mydb`.`users`.`fk2` FROM `mydb`.`users` WHERE `age` < ? ORDER BY `id` LIMIT 3 OFFSET 4")).
		WithArgs(40).
		WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name", "fk1", "fk2"}).
			AddRow(1, 10, nil, nil, nil).
			AddRow(2, 20, "", 0, 0).
			AddRow(3, 30, "a8m", 1, 1))
	var (
		users []*user
		spec  = &QuerySpec{
			Node: &NodeSpec{
				Table:   "users",
				Schema:  "mydb",
				Columns: []string{"id", "age", "name", "fk1", "fk2"},
				ID:      &FieldSpec{Column: "id", Type: field.TypeInt},
			},
			Limit:  3,
			Offset: 4,
			Unique: true,
			Order: func(s *sql.Selector) {
				s.OrderBy("id")
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.LT("age", 40))
			},
			ScanValues: func(columns []string) ([]any, error) {
				u := &user{}
				users = append(users, u)
				return u.values(columns)
			},
			Assign: func(columns []string, values []any) error {
				return users[len(users)-1].assign(columns, values)
			},
		}
	)

	// Query and scan.
	err = QueryNodes(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, &user{id: 1, age: 10, name: ""}, users[0])
	require.Equal(t, &user{id: 2, age: 20, name: ""}, users[1])
	require.Equal(t, &user{id: 3, age: 30, name: "a8m", edges: struct{ fk1, fk2 int }{1, 1}}, users[2])
}

func TestQueryEdges(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectQuery(escape("SELECT `group_id`, `user_id` FROM `user_groups` WHERE `user_id` IN (?, ?, ?)")).
		WithArgs(1, 2, 3).
		WillReturnRows(sqlmock.NewRows([]string{"group_id", "user_id"}).
			AddRow(4, 5).
			AddRow(4, 6))

	var (
		edges [][]int64
		spec  = &EdgeQuerySpec{
			Edge: &EdgeSpec{
				Inverse: true,
				Table:   "user_groups",
				Columns: []string{"user_id", "group_id"},
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues("user_id", 1, 2, 3))
			},
			ScanValues: func() [2]any {
				return [2]any{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in any) error {
				o, i := out.(*sql.NullInt64), in.(*sql.NullInt64)
				edges = append(edges, []int64{o.Int64, i.Int64})
				return nil
			},
		}
	)

	// Query and scan.
	err = QueryEdges(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, [][]int64{{4, 5}, {4, 6}}, edges)
}

func TestQueryEdgesSchema(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectQuery(escape("SELECT `group_id`, `user_id` FROM `mydb`.`user_groups` WHERE `user_id` IN (?, ?, ?)")).
		WithArgs(1, 2, 3).
		WillReturnRows(sqlmock.NewRows([]string{"group_id", "user_id"}).
			AddRow(4, 5).
			AddRow(4, 6))

	var (
		edges [][]int64
		spec  = &EdgeQuerySpec{
			Edge: &EdgeSpec{
				Inverse: true,
				Table:   "user_groups",
				Schema:  "mydb",
				Columns: []string{"user_id", "group_id"},
			},
			Predicate: func(s *sql.Selector) {
				s.Where(sql.InValues("user_id", 1, 2, 3))
			},
			ScanValues: func() [2]any {
				return [2]any{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in any) error {
				o, i := out.(*sql.NullInt64), in.(*sql.NullInt64)
				edges = append(edges, []int64{o.Int64, i.Int64})
				return nil
			},
		}
	)

	// Query and scan.
	err = QueryEdges(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, [][]int64{{4, 5}, {4, 6}}, edges)
}

func TestIsConstraintError(t *testing.T) {
	tests := []struct {
		name               string
		errMessage         string
		expectedConstraint bool
		expectedFK         bool
		expectedUnique     bool
		expectedCheck      bool
	}{
		{
			name: "MySQL FK",
			errMessage: `insert node to table "pets": Error 1452: Cannot add or update a child row: a foreign key` +
				" constraint fails (`test`.`pets`, CONSTRAINT `pets_users_pets` FOREIGN KEY (`user_pets`) REFERENCES " +
				"`users` (`id`) ON DELETE SET NULL)",
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name:               "SQLite FK",
			errMessage:         `insert node to table "pets": FOREIGN KEY constraint failed`,
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name:               "Postgres FK",
			errMessage:         `insert node to table "pets": pq: insert or update on table "pets" violates foreign key constraint "pets_users_pets"`,
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name: "MySQL FK",
			errMessage: "Error 1451: Cannot delete or update a parent row: a foreign key constraint " +
				"fails (`test`.`groups`, CONSTRAINT `groups_group_infos_info` FOREIGN KEY (`group_info`) REFERENCES `group_infos` (`id`))",
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name:               "SQLite FK",
			errMessage:         `FOREIGN KEY constraint failed`,
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name:               "Postgres FK",
			errMessage:         `pq: update or delete on table "group_infos" violates foreign key constraint "groups_group_infos_info" on table "groups"`,
			expectedConstraint: true,
			expectedFK:         true,
			expectedUnique:     false,
			expectedCheck:      false,
		},
		{
			name:               "MySQL Unique",
			errMessage:         `insert node to table "file_types": UNIQUE constraint failed: file_types.name ent: constraint failed: insert node to table "file_types": UNIQUE constraint failed: file_types.name`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     true,
			expectedCheck:      false,
		},
		{
			name:               "SQLite Unique",
			errMessage:         `insert node to table "file_types": UNIQUE constraint failed: file_types.name ent: constraint failed: insert node to table "file_types": UNIQUE constraint failed: file_types.name`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     true,
			expectedCheck:      false,
		},
		{
			name:               "Postgres Unique",
			errMessage:         `insert node to table "file_types": pq: duplicate key value violates unique constraint "file_types_name_key" ent: constraint failed: insert node to table "file_types": pq: duplicate key value violates unique constraint "file_types_name_key"`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     true,
			expectedCheck:      false,
		},
		{
			name:               "MySQL Check",
			errMessage:         `insert node to table "users": Error 3819: Check constraint 'users_age_check' is violated`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     false,
			expectedCheck:      true,
		},
		{
			name:               "SQLite Check",
			errMessage:         `insert node to table "users": CHECK constraint failed: age >= 18`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     false,
			expectedCheck:      true,
		},
		{
			name:               "Postgres Check",
			errMessage:         `insert node to table "users": pq: new row for relation "users" violates check constraint "users_age_check"`,
			expectedConstraint: true,
			expectedFK:         false,
			expectedUnique:     false,
			expectedCheck:      true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := errors.New(tt.errMessage)
			require.EqualValues(t, tt.expectedConstraint, IsConstraintError(err))
			require.EqualValues(t, tt.expectedFK, IsForeignKeyConstraintError(err))
			require.EqualValues(t, tt.expectedUnique, IsUniqueConstraintError(err))
			require.EqualValues(t, tt.expectedCheck, IsCheckConstraintError(err))
		})
	}
}

func TestLimitNeighbors(t *testing.T) {
	t.Run("O2M", func(t *testing.T) {
		const fk = "author_id"
		// Authors load their posts.
		s := sql.Select(fk, "id").From(sql.Table("posts"))
		LimitNeighbors(fk, 2)(s)
		query, args := s.Query()
		require.Equal(t,
			"WITH `src_query` AS (SELECT `author_id`, `id` FROM `posts`), `limited_query` AS (SELECT *, (ROW_NUMBER() OVER (PARTITION BY `author_id` ORDER BY `id`)) AS `row_number` FROM `src_query`) SELECT `author_id`, `id` FROM `limited_query` AS `posts` WHERE `posts`.`row_number` <= ?",
			query,
		)
		require.Equal(t, []any{2}, args)
	})
	t.Run("M2M", func(t *testing.T) {
		const fk = "user_id"
		edgeT, neighborsT := sql.Table("user_groups"), sql.Table("groups")
		s := sql.Select(fk, "id", "name").From(neighborsT).Join(edgeT).On(neighborsT.C("id"), edgeT.C("group_id"))
		LimitNeighbors(fk, 1, sql.ExprFunc(func(b *sql.Builder) { b.Ident("updated_at") }))(s)
		query, args := s.Query()
		require.Equal(t,
			"WITH `src_query` AS (SELECT `user_id`, `id`, `name` FROM `groups` JOIN `user_groups` AS `t1` ON `groups`.`id` = `t1`.`group_id`), `limited_query` AS (SELECT *, (ROW_NUMBER() OVER (PARTITION BY `user_id` ORDER BY `updated_at`)) AS `row_number` FROM `src_query`) SELECT `user_id`, `id`, `name` FROM `limited_query` AS `groups` WHERE `groups`.`row_number` <= ?",
			query,
		)
		require.Equal(t, []any{1}, args)
	})
}

func escape(query string) string {
	rows := strings.Split(query, "\n")
	for i := range rows {
		rows[i] = strings.TrimPrefix(rows[i], " ")
	}
	query = strings.Join(rows, " ")
	return strings.TrimSpace(regexp.QuoteMeta(query)) + "$"
}
