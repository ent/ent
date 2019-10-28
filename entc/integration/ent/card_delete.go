// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// CardDelete is the builder for deleting a Card entity.
type CardDelete struct {
	config
	predicates []predicate.Card
}

// Where adds a new predicate to the delete builder.
func (cd *CardDelete) Where(ps ...predicate.Card) *CardDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CardDelete) Exec(ctx context.Context) (int, error) {
	switch cd.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return cd.sqlExec(ctx)
	case dialect.Gremlin:
		return cd.gremlinExec(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CardDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CardDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(cd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(card.Table))
	for _, p := range cd.predicates {
		p(selector)
	}
	query, args := builder.Delete(card.Table).FromSelect(selector).Query()
	if err := cd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (cd *CardDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cd.gremlin().Query()
	if err := cd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (cd *CardDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(card.Label)
	for _, p := range cd.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// CardDeleteOne is the builder for deleting a single Card entity.
type CardDeleteOne struct {
	cd *CardDelete
}

// Exec executes the deletion query.
func (cdo *CardDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{card.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CardDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
