// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/item"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// ItemDelete is the builder for deleting a Item entity.
type ItemDelete struct {
	config
	predicates []predicate.Item
}

// Where adds a new predicate to the delete builder.
func (id *ItemDelete) Where(ps ...predicate.Item) *ItemDelete {
	id.predicates = append(id.predicates, ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *ItemDelete) Exec(ctx context.Context) (int, error) {
	return id.gremlinExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *ItemDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *ItemDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := id.gremlin().Query()
	if err := id.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (id *ItemDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(item.Label)
	for _, p := range id.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// ItemDeleteOne is the builder for deleting a single Item entity.
type ItemDeleteOne struct {
	id *ItemDelete
}

// Exec executes the deletion query.
func (ido *ItemDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{item.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *ItemDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}
