// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/goods"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// GoodsDelete is the builder for deleting a Goods entity.
type GoodsDelete struct {
	config
	hooks    []Hook
	mutation *GoodsMutation
}

// Where appends a list predicates to the GoodsDelete builder.
func (d *GoodsDelete) Where(ps ...predicate.Goods) *GoodsDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *GoodsDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.gremlinExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *GoodsDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *GoodsDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := d.gremlin().Query()
	if err := d.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	d.mutation.done = true
	return res.ReadInt()
}

func (d *GoodsDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(goods.Label)
	for _, p := range d.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// GoodsDeleteOne is the builder for deleting a single Goods entity.
type GoodsDeleteOne struct {
	d *GoodsDelete
}

// Where appends a list predicates to the GoodsDelete builder.
func (d *GoodsDeleteOne) Where(ps ...predicate.Goods) *GoodsDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *GoodsDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{goods.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *GoodsDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
