// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNeighbors(t *testing.T) {
	tests := []struct {
		name      string
		input     *Step
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name: "O2O/1type",
			input: func() *Step {
				step := &Step{}
				// Since the relation is on the same table,
				// V used as a reference value.
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"spouse_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` WHERE `spouse_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2O/1type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "nodes"
				step.From.Column = "id"
				step.To.Table = "nodes"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "nodes"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"prev_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `nodes` JOIN (SELECT `prev_id` FROM `nodes` WHERE `id` = ?) AS `t1` ON `nodes`.`id` = `t1`.`prev_id`",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2M/1type",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"parent_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` WHERE `parent_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2O/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "card"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `card` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{2},
		},
		{
			name: "O2O/2types/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "card"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `cards` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "O2M/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = 1
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "pets"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `pets` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "M2O/2types/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "pets"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Inverse = true
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `pets` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2O/1type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Inverse = true
				step.Edge.Table = "users"
				step.Edge.Columns = []string{"parent_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `parent_id` FROM `users` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`parent_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2M/2type",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "groups"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"group_id", "user_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `user_groups`.`user_id` FROM `user_groups` WHERE `user_groups`.`group_id` = ?) AS `t1` ON `users`.`id` = `t1`.`user_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2M/2type/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = 2
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Inverse = true
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"group_id", "user_id"}
				return step
			}(),
			wantQuery: "SELECT * FROM `groups` JOIN (SELECT `user_groups`.`group_id` FROM `user_groups` WHERE `user_groups`.`user_id` = ?) AS `t1` ON `groups`.`id` = `t1`.`group_id`",
			wantArgs:  []interface{}{2},
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
		wantArgs  []interface{}
	}{
		{
			name: "O2M/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = Select().From(Table("users")).Where(EQ("name", "a8m"))
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "pets"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: `SELECT * FROM "pets" JOIN (SELECT "users"."id" FROM "users" WHERE "name" = $1) AS "t1" ON "pets"."owner_id" = "t1"."id"`,
			wantArgs:  []interface{}{"a8m"},
		},
		{
			name: "M2O/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = Select().From(Table("pets")).Where(EQ("name", "pedro"))
				step.From.Table = "pets"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			wantQuery: `SELECT * FROM "users" JOIN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $1) AS "t1" ON "users"."id" = "t1"."owner_id"`,
			wantArgs:  []interface{}{"pedro"},
		},
		{
			name: "M2M/2types",
			input: func() *Step {
				step := &Step{}
				step.From.V = Select().From(Table("users")).Where(EQ("name", "a8m"))
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"user_id", "group_id"}
				return step
			}(),
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
			wantArgs: []interface{}{"a8m"},
		},
		{
			name: "M2M/2types/inverse",
			input: func() *Step {
				step := &Step{}
				step.From.V = Select().From(Table("groups")).Where(EQ("name", "GitHub"))
				step.From.Table = "groups"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Inverse = true
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"user_id", "group_id"}
				return step
			}(),
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
			wantArgs: []interface{}{"GitHub"},
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
		selector  *Selector
		wantQuery string
	}{
		{
			name: "O2O/1type",
			step: func() *Step {
				// A nodes table; linked-list (next->prev). The "prev"
				// node holds association pointer. The neighbors query
				// here checks if a node "has-next".
				step := &Step{}
				step.From.Table = "nodes"
				step.From.Column = "id"
				step.To.Table = "nodes"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "nodes"
				step.Edge.Columns = []string{"prev_id"}
				return step
			}(),
			selector:  Select("*").From(Table("nodes")),
			wantQuery: "SELECT * FROM `nodes` WHERE `nodes`.`id` IN (SELECT `nodes`.`prev_id` FROM `nodes` WHERE `nodes`.`prev_id` IS NOT NULL)",
		},
		{
			name: "O2O/1type/inverse",
			step: func() *Step {
				// Same example as above, but the neighbors
				// query checks if a node "has-previous".
				step := &Step{}
				step.From.Table = "nodes"
				step.From.Column = "id"
				step.To.Table = "nodes"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Inverse = true
				step.Edge.Table = "nodes"
				step.Edge.Columns = []string{"prev_id"}
				return step
			}(),
			selector:  Select("*").From(Table("nodes")),
			wantQuery: "SELECT * FROM `nodes` WHERE `nodes`.`prev_id` IS NOT NULL",
		},
		{
			name: "O2M/2type2",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "pets"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector:  Select("*").From(Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `pets`.`owner_id` FROM `pets` WHERE `pets`.`owner_id` IS NOT NULL)",
		},
		{
			name: "M2O/2type2",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "pets"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Inverse = true
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector:  Select("*").From(Table("pets")),
			wantQuery: "SELECT * FROM `pets` WHERE `pets`.`owner_id` IS NOT NULL",
		},
		{
			name: "M2M/2types",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"user_id", "group_id"}
				return step
			}(),
			selector:  Select("*").From(Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `user_groups`.`user_id` FROM `user_groups`)",
		},
		{
			name: "M2M/2types/inverse",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Inverse = true
				step.Edge.Table = "group_users"
				step.Edge.Columns = []string{"group_id", "user_id"}
				return step
			}(),
			selector:  Select("*").From(Table("users")),
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `group_users`.`user_id` FROM `group_users`)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HasNeighbors(tt.selector, tt.step)
			query, args := tt.selector.Query()
			require.Equal(t, tt.wantQuery, query)
			require.Empty(t, args)
		})
	}
}

func TestHasNeighborsWith(t *testing.T) {
	tests := []struct {
		name      string
		step      *Step
		selector  *Selector
		predicate func(*Selector)
		wantQuery string
		wantArgs  []interface{}
	}{
		{
			name: "O2O",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "cards"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").From(Table("users")),
			predicate: func(s *Selector) {
				s.Where(EQ("expired", false))
			},
			wantQuery: `SELECT * FROM "users" WHERE "users"."id" IN (SELECT "cards"."owner_id" FROM "cards" WHERE "expired" = $1)`,
			wantArgs:  []interface{}{false},
		},
		{
			name: "O2O/inverse",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "cards"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = O2O
				step.Edge.Table = "cards"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").From(Table("cards")),
			predicate: func(s *Selector) {
				s.Where(EQ("name", "a8m"))
			},
			wantQuery: `SELECT * FROM "cards" WHERE "cards"."owner_id" IN (SELECT "users"."id" FROM "users" WHERE "name" = $1)`,
			wantArgs:  []interface{}{"a8m"},
		},
		{
			name: "O2M",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "pets"
				step.To.Column = "id"
				step.Edge.Rel = O2M
				step.Edge.Table = "pets"
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").
				From(Table("users")).
				Where(EQ("last_name", "mashraki")),
			predicate: func(s *Selector) {
				s.Where(EQ("name", "pedro"))
			},
			wantQuery: `SELECT * FROM "users" WHERE "last_name" = $1 AND "users"."id" IN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $2)`,
			wantArgs:  []interface{}{"mashraki", "pedro"},
		},
		{
			name: "M2O",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "pets"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2O
				step.Edge.Table = "pets"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"owner_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").
				From(Table("pets")).
				Where(EQ("name", "pedro")),
			predicate: func(s *Selector) {
				s.Where(EQ("last_name", "mashraki"))
			},
			wantQuery: `SELECT * FROM "pets" WHERE "name" = $1 AND "pets"."owner_id" IN (SELECT "users"."id" FROM "users" WHERE "last_name" = $2)`,
			wantArgs:  []interface{}{"pedro", "mashraki"},
		},
		{
			name: "M2M",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "users"
				step.From.Column = "id"
				step.To.Table = "groups"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Columns = []string{"user_id", "group_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").From(Table("users")),
			predicate: func(s *Selector) {
				s.Where(EQ("name", "GitHub"))
			},
			wantQuery: `
SELECT *
FROM "users"
WHERE "users"."id" IN
  (SELECT "user_groups"."user_id"
  FROM "user_groups"
  JOIN "groups" AS "t0" ON "user_groups"."group_id" = "t0"."id" WHERE "name" = $1)`,
			wantArgs: []interface{}{"GitHub"},
		},
		{
			name: "M2M/inverse",
			step: func() *Step {
				step := &Step{}
				step.From.Table = "groups"
				step.From.Column = "id"
				step.To.Table = "users"
				step.To.Column = "id"
				step.Edge.Rel = M2M
				step.Edge.Table = "user_groups"
				step.Edge.Inverse = true
				step.Edge.Columns = []string{"user_id", "group_id"}
				return step
			}(),
			selector: Dialect("postgres").Select("*").From(Table("groups")),
			predicate: func(s *Selector) {
				s.Where(EQ("name", "a8m"))
			},
			wantQuery: `
SELECT *
FROM "groups"
WHERE "groups"."id" IN
  (SELECT "user_groups"."group_id"
  FROM "user_groups"
  JOIN "users" AS "t0" ON "user_groups"."user_id" = "t0"."id" WHERE "name" = $1)`,
			wantArgs: []interface{}{"a8m"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			HasNeighborsWith(tt.selector, tt.step, tt.predicate)
			query, args := tt.selector.Query()
			tt.wantQuery = strings.Join(strings.Fields(tt.wantQuery), " ")
			require.Equal(t, tt.wantQuery, query)
			require.Equal(t, tt.wantArgs, args)
		})
	}
}
