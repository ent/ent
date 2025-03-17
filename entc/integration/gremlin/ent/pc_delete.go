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
	"entgo.io/ent/entc/integration/gremlin/ent/pc"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// PCDelete is the builder for deleting a PC entity.
type PCDelete struct {
	config
	hooks    []Hook
	mutation *PCMutation
}

// Where appends a list predicates to the PCDelete builder.
func (d *PCDelete) Where(ps ...predicate.PC) *PCDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *PCDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.gremlinExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *PCDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *PCDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := d.gremlin().Query()
	if err := d.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	d.mutation.done = true
	return res.ReadInt()
}

func (d *PCDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(pc.Label)
	for _, p := range d.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// PCDeleteOne is the builder for deleting a single PC entity.
type PCDeleteOne struct {
	d *PCDelete
}

// Where appends a list predicates to the PCDelete builder.
func (d *PCDeleteOne) Where(ps ...predicate.PC) *PCDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *PCDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{pc.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *PCDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
