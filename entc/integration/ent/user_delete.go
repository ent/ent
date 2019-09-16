// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// UserDelete is the builder for deleting a User entity.
type UserDelete struct {
	config
	predicates []predicate.User
}

// Where adds a new predicate to the delete builder.
func (ud *UserDelete) Where(ps ...predicate.User) *UserDelete {
	ud.predicates = append(ud.predicates, ps...)
	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserDelete) Exec(ctx context.Context) (int, error) {
	switch ud.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ud.sqlExec(ctx)
	case dialect.Gremlin:
		return ud.gremlinExec(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ud *UserDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(user.Table))
	for _, p := range ud.predicates {
		p(selector)
	}
	query, args := sql.Delete(user.Table).FromSelect(selector).Query()
	if err := ud.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (ud *UserDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ud.gremlin().Query()
	if err := ud.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ud *UserDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(user.Label)
	for _, p := range ud.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// UserDeleteOne is the builder for deleting a single User entity.
type UserDeleteOne struct {
	ud *UserDelete
}

// Exec executes the deletion query.
func (udo *UserDeleteOne) Exec(ctx context.Context) error {
	n, err := udo.ud.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{user.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
