// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sqlgraph

import (
	"context"
	"database/sql/driver"
	"regexp"
	"strings"
	"testing"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/schema/field"

	"github.com/DATA-DOG/go-sqlmock"
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
			// Since the relation is on the same sql.Table,
			// V used as a reference value.
			input: NewStep(
				From("users", "id", 1),
				To("users", "id"),
				Edge(O2O, false, "users", "spouse_id"),
			),
			wantQuery: "SELECT * FROM `users` WHERE `spouse_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2O/1type/inverse",
			input: NewStep(
				From("nodes", "id", 1),
				To("nodes", "id"),
				Edge(O2O, true, "nodes", "prev_id"),
			),
			wantQuery: "SELECT * FROM `nodes` JOIN (SELECT `prev_id` FROM `nodes` WHERE `id` = ?) AS `t1` ON `nodes`.`id` = `t1`.`prev_id`",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2M/1type",
			input: NewStep(
				From("users", "id", 1),
				To("users", "id"),
				Edge(O2M, false, "users", "parent_id"),
			),
			wantQuery: "SELECT * FROM `users` WHERE `parent_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "O2O/2types",
			input: NewStep(
				From("users", "id", 2),
				To("card", "id"),
				Edge(O2O, false, "cards", "owner_id"),
			),
			wantQuery: "SELECT * FROM `card` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{2},
		},
		{
			name: "O2O/2types/inverse",
			input: NewStep(
				From("cards", "id", 2),
				To("users", "id"),
				Edge(O2O, true, "cards", "owner_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `cards` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "O2M/2types",
			input: NewStep(
				From("users", "id", 1),
				To("pets", "id"),
				Edge(O2M, false, "pets", "owner_id"),
			),
			wantQuery: "SELECT * FROM `pets` WHERE `owner_id` = ?",
			wantArgs:  []interface{}{1},
		},
		{
			name: "M2O/2types/inverse",
			input: NewStep(
				From("pets", "id", 2),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `owner_id` FROM `pets` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`owner_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2O/1type/inverse",
			input: NewStep(
				From("users", "id", 2),
				To("users", "id"),
				Edge(M2O, true, "users", "parent_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `parent_id` FROM `users` WHERE `id` = ?) AS `t1` ON `users`.`id` = `t1`.`parent_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2M/2type",
			input: NewStep(
				From("groups", "id", 2),
				To("users", "id"),
				Edge(M2M, false, "user_groups", "group_id", "user_id"),
			),
			wantQuery: "SELECT * FROM `users` JOIN (SELECT `user_groups`.`user_id` FROM `user_groups` WHERE `user_groups`.`group_id` = ?) AS `t1` ON `users`.`id` = `t1`.`user_id`",
			wantArgs:  []interface{}{2},
		},
		{
			name: "M2M/2type/inverse",
			input: NewStep(
				From("users", "id", 2),
				To("groups", "id"),
				Edge(M2M, true, "user_groups", "group_id", "user_id"),
			),
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
			input: NewStep(
				From("users", "id", sql.Select().From(sql.Table("users")).Where(sql.EQ("name", "a8m"))),
				To("pets", "id"),
				Edge(O2M, false, "users", "owner_id"),
			),
			wantQuery: `SELECT * FROM "pets" JOIN (SELECT "users"."id" FROM "users" WHERE "name" = $1) AS "t1" ON "pets"."owner_id" = "t1"."id"`,
			wantArgs:  []interface{}{"a8m"},
		},
		{
			name: "M2O/2types",
			input: NewStep(
				From("pets", "id", sql.Select().From(sql.Table("pets")).Where(sql.EQ("name", "pedro"))),
				To("users", "id"),
				Edge(M2O, true, "pets", "owner_id"),
			),
			wantQuery: `SELECT * FROM "users" JOIN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $1) AS "t1" ON "users"."id" = "t1"."owner_id"`,
			wantArgs:  []interface{}{"pedro"},
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
			wantArgs: []interface{}{"a8m"},
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
			wantQuery: "SELECT * FROM `nodes` WHERE `nodes`.`id` IN (SELECT `nodes`.`prev_id` FROM `nodes` WHERE `nodes`.`prev_id` IS NOT NULL)",
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
			wantQuery: "SELECT * FROM `users` WHERE `users`.`id` IN (SELECT `pets`.`owner_id` FROM `pets` WHERE `pets`.`owner_id` IS NOT NULL)",
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
		selector  *sql.Selector
		predicate func(*sql.Selector)
		wantQuery string
		wantArgs  []interface{}
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
			wantQuery: `SELECT * FROM "users" WHERE "users"."id" IN (SELECT "cards"."owner_id" FROM "cards" WHERE "expired" = $1)`,
			wantArgs:  []interface{}{false},
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
			wantQuery: `SELECT * FROM "cards" WHERE "cards"."owner_id" IN (SELECT "users"."id" FROM "users" WHERE "name" = $1)`,
			wantArgs:  []interface{}{"a8m"},
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
			wantQuery: `SELECT * FROM "users" WHERE "last_name" = $1 AND "users"."id" IN (SELECT "pets"."owner_id" FROM "pets" WHERE "name" = $2)`,
			wantArgs:  []interface{}{"mashraki", "pedro"},
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
			wantQuery: `SELECT * FROM "pets" WHERE "name" = $1 AND "pets"."owner_id" IN (SELECT "users"."id" FROM "users" WHERE "last_name" = $2)`,
			wantArgs:  []interface{}{"pedro", "mashraki"},
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
  JOIN "groups" AS "t0" ON "user_groups"."group_id" = "t0"."id" WHERE "name" = $1)`,
			wantArgs: []interface{}{"GitHub"},
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
  JOIN "users" AS "t0" ON "user_groups"."user_id" = "t0"."id" WHERE "name" = $1)`,
			wantArgs: []interface{}{"a8m"},
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
  JOIN "users" AS "t0" ON "user_groups"."user_id" = "t0"."id" WHERE "name" IS NOT NULL AND "name" = $1)`,
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
				ID:    &FieldSpec{Column: "id"},
				Fields: []*FieldSpec{
					{Column: "age", Type: field.TypeInt, Value: 30},
					{Column: "name", Type: field.TypeString, Value: "a8m"},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`age`, `name`) VALUES (?, ?)")).
					WithArgs(30, "a8m").
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
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
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`age`, `name`, `id`) VALUES (?, ?, ?)")).
					WithArgs(30, "a8m", 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "fields/json",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id"},
				Fields: []*FieldSpec{
					{Column: "json", Type: field.TypeJSON, Value: struct{}{}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `users` (`json`) VALUES (?)")).
					WithArgs([]byte("{}")).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2o",
			spec: &CreateSpec{
				Table: "pets",
				ID:    &FieldSpec{Column: "id"},
				Fields: []*FieldSpec{
					{Column: "name", Type: field.TypeString, Value: "pedro"},
				},
				Edges: []*EdgeSpec{
					{Rel: M2O, Columns: []string{"owner_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `pets` (`name`, `owner_id`) VALUES (?, ?)")).
					WithArgs("pedro", 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/o2o/inverse",
			spec: &CreateSpec{
				Table: "cards",
				ID:    &FieldSpec{Column: "id"},
				Fields: []*FieldSpec{
					{Column: "number", Type: field.TypeString, Value: "0001"},
				},
				Edges: []*EdgeSpec{
					{Rel: O2O, Columns: []string{"owner_id"}, Inverse: true, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectExec(escape("INSERT INTO `cards` (`number`, `owner_id`) VALUES (?, ?)")).
					WithArgs("0001", 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/o2m",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id"},
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
				ID:    &FieldSpec{Column: "id"},
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
				ID:    &FieldSpec{Column: "id"},
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
				ID:    &FieldSpec{Column: "id"},
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
				ID:    &FieldSpec{Column: "id"},
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
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?)")).
					WithArgs(1, 2).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/inverse",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id"},
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
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?)")).
					WithArgs(2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/bidi",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id"},
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
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(1, 2, 2, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
		{
			name: "edges/m2m/bidi/batch",
			spec: &CreateSpec{
				Table: "users",
				ID:    &FieldSpec{Column: "id"},
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
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(4, 1, 5, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?)")).
					WithArgs(1, 2, 2, 1, 1, 3, 3, 1).
					WillReturnResult(sqlmock.NewResult(1, 1))
				m.ExpectCommit()
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.expect(mock)
			err = CreateNode(context.Background(), sql.OpenDB("", db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
		})
	}
}

func TestBatchCreate(t *testing.T) {
	tests := []struct {
		name    string
		nodes   []*CreateSpec
		expect  func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "empty",
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				m.ExpectCommit()
			},
		},
		{
			name: "multiple",
			nodes: []*CreateSpec{
				{
					Table: "users",
					ID:    &FieldSpec{Column: "id"},
					Fields: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 32},
						{Column: "name", Type: field.TypeString, Value: "a8m"},
						{Column: "active", Type: field.TypeBool, Value: false},
					},
					Edges: []*EdgeSpec{
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
						{Rel: M2M, Table: "user_products", Columns: []string{"user_id", "product_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{2}}},
						{Rel: M2O, Table: "company", Columns: []string{"workplace_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}}},
						{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
					},
				},
				{
					Table: "users",
					ID:    &FieldSpec{Column: "id"},
					Fields: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 30},
						{Column: "name", Type: field.TypeString, Value: "nati"},
					},
					Edges: []*EdgeSpec{
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
						{Rel: M2M, Table: "user_products", Columns: []string{"user_id", "product_id"}, Target: &EdgeTarget{Nodes: []driver.Value{2}, IDSpec: &FieldSpec{Column: "id"}}},
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{2}}},
						{Rel: O2M, Table: "pets", Columns: []string{"owner_id"}, Target: &EdgeTarget{Nodes: []driver.Value{3}, IDSpec: &FieldSpec{Column: "id"}}},
					},
				},
			},
			expect: func(m sqlmock.Sqlmock) {
				m.ExpectBegin()
				// Insert nodes with FKs.
				m.ExpectExec(escape("INSERT INTO `users` (`active`, `age`, `name`, `workplace_id`) VALUES (?, ?, ?, ?), (?, ?, ?, ?)")).
					WithArgs(false, 32, "a8m", 2, nil, 30, "nati", nil).
					WillReturnResult(sqlmock.NewResult(10, 2))
				// Insert M2M inverse-edges.
				m.ExpectExec(escape("INSERT INTO `group_users` (`group_id`, `user_id`) VALUES (?, ?), (?, ?)")).
					WithArgs(2, 10, 2, 11).
					WillReturnResult(sqlmock.NewResult(2, 2))
				// Insert M2M bidirectional edges.
				m.ExpectExec(escape("INSERT INTO `user_friends` (`user_id`, `friend_id`) VALUES (?, ?), (?, ?), (?, ?), (?, ?)")).
					WithArgs(10, 2, 2, 10, 11, 2, 2, 11).
					WillReturnResult(sqlmock.NewResult(2, 2))
				// Insert M2M edges.
				m.ExpectExec(escape("INSERT INTO `user_products` (`user_id`, `product_id`) VALUES (?, ?), (?, ?)")).
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
			err = BatchCreate(context.Background(), sql.OpenDB("mysql", db), &BatchCreateSpec{Nodes: tt.nodes})
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

func (*user) values() []interface{} {
	return []interface{}{&sql.NullInt64{}, &sql.NullInt64{}, &sql.NullString{}}
}

func (u *user) assign(values ...interface{}) error {
	u.id = int(values[0].(*sql.NullInt64).Int64)
	u.age = int(values[1].(*sql.NullInt64).Int64)
	u.name = values[2].(*sql.NullString).String
	// loaded with foreign-keys.
	if len(values) > 3 {
		u.edges.fk1 = int(values[3].(*sql.NullInt64).Int64)
		u.edges.fk2 = int(values[4].(*sql.NullInt64).Int64)
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
			name: "fields/add_clear",
			spec: &UpdateSpec{
				Node: &NodeSpec{
					Table:   "users",
					Columns: []string{"id", "name", "age"},
					ID:      &FieldSpec{Column: "id", Type: field.TypeInt, Value: 1},
				},
				Fields: FieldMut{
					Add: []*FieldSpec{
						{Column: "age", Type: field.TypeInt, Value: 1},
					},
					Clear: []*FieldSpec{
						{Column: "name", Type: field.TypeString},
					},
				},
			},
			prepare: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(escape("UPDATE `users` SET `name` = NULL, `age` = COALESCE(`age`, ?) + ? WHERE `id` = ?")).
					WithArgs(0, 1, 1).
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
						{Rel: O2O, Table: "users", Bidi: true, Columns: []string{"spouse_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{2}}},
					},
					Add: []*EdgeSpec{
						{Rel: O2O, Table: "users", Bidi: true, Columns: []string{"spouse_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{3}}},
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
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{2}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{3, 7}}},
						// Clear all "following" edges (and their inverse).
						{Rel: M2M, Table: "user_following", Bidi: true, Columns: []string{"following_id", "follower_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
						// Clear all "user_blocked" edges.
						{Rel: M2M, Table: "user_blocked", Columns: []string{"user_id", "blocked_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
						// Clear all "comments" edges.
						{Rel: M2M, Inverse: true, Table: "comment_responders", Columns: []string{"comment_id", "responder_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}}},
					},
					Add: []*EdgeSpec{
						{Rel: M2M, Table: "user_friends", Bidi: true, Columns: []string{"user_id", "friend_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{4}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{5}}},
						{Rel: M2M, Inverse: true, Table: "group_users", Columns: []string{"group_id", "user_id"}, Target: &EdgeTarget{IDSpec: &FieldSpec{Column: "id"}, Nodes: []driver.Value{6, 8}}},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			require.NoError(t, err)
			tt.prepare(mock)
			usr := &user{}
			tt.spec.Assign = usr.assign
			tt.spec.ScanValues = usr.values()
			err = UpdateNode(context.Background(), sql.OpenDB("", db), tt.spec)
			require.Equal(t, tt.wantErr, err != nil, err)
			require.Equal(t, tt.wantUser, usr)
		})
	}
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
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users`")).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1).
						AddRow(2))
				// Apply field changes.
				mock.ExpectExec(escape("UPDATE `users` SET `age` = ?, `name` = ? WHERE `id` IN (?, ?)")).
					WithArgs(30, "Ariel", 1, 2).
					WillReturnResult(sqlmock.NewResult(0, 2))
				mock.ExpectCommit()
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
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users` WHERE `name` = ?")).
					WithArgs("a8m").
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
				// Clear fields.
				mock.ExpectExec(escape("UPDATE `users` SET `age` = NULL, `name` = NULL WHERE `id` = ?")).
					WithArgs(1).
					WillReturnResult(sqlmock.NewResult(0, 1))
				mock.ExpectCommit()
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
				mock.ExpectBegin()
				// Get all node ids first.
				mock.ExpectQuery(escape("SELECT `id` FROM `users`")).
					WillReturnRows(sqlmock.NewRows([]string{"id"}).
						AddRow(1))
				// Clear "car" and "workplace" foreign_keys and add "card" and a "parent".
				mock.ExpectExec(escape("UPDATE `users` SET `workplace_id` = NULL, `car_id` = NULL, `parent_id` = ?, `card_id` = ? WHERE `id` = ?")).
					WithArgs(4, 3, 1).
					WillReturnResult(sqlmock.NewResult(0, 1))
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
	mock.ExpectBegin()
	mock.ExpectExec(escape("DELETE FROM `users`")).
		WillReturnResult(sqlmock.NewResult(0, 2))
	mock.ExpectCommit()
	affected, err := DeleteNodes(context.Background(), sql.OpenDB("", db), &DeleteSpec{
		Node: &NodeSpec{
			Table: "users",
			ID:    &FieldSpec{Column: "id", Type: field.TypeInt},
		},
	})
	require.NoError(t, err)
	require.Equal(t, 2, affected)
}

func TestQueryNodes(t *testing.T) {
	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	mock.ExpectQuery(escape("SELECT DISTINCT `users`.`id`, `users`.`age`, `users`.`name`, `users`.`fk1`, `users`.`fk2` FROM `users` WHERE `age` < ? ORDER BY `id` LIMIT 3 OFFSET 4")).
		WithArgs(40).
		WillReturnRows(sqlmock.NewRows([]string{"id", "age", "name", "fk1", "fk2"}).
			AddRow(1, 10, nil, nil, nil).
			AddRow(2, 20, "", 0, 0).
			AddRow(3, 30, "a8m", 1, 1))
	mock.ExpectQuery(escape("SELECT COUNT(DISTINCT `users`.`id`) FROM `users` WHERE `age` < ? ORDER BY `id` LIMIT 3 OFFSET 4")).
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
			ScanValues: func() []interface{} {
				u := &user{}
				users = append(users, u)
				return append(u.values(), &sql.NullInt64{}, &sql.NullInt64{}) // extra values for fks.
			},
			Assign: func(values ...interface{}) error {
				return users[len(users)-1].assign(values...)
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
	n, err := CountNodes(context.Background(), sql.OpenDB("", db), spec)
	require.NoError(t, err)
	require.Equal(t, 3, n)
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
			ScanValues: func() [2]interface{} {
				return [2]interface{}{&sql.NullInt64{}, &sql.NullInt64{}}
			},
			Assign: func(out, in interface{}) error {
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

func escape(query string) string {
	rows := strings.Split(query, "\n")
	for i := range rows {
		rows[i] = strings.TrimPrefix(rows[i], " ")
	}
	query = strings.Join(rows, " ")
	return regexp.QuoteMeta(query)
}
