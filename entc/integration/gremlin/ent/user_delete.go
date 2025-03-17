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
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// UserDelete is the builder for deleting a User entity.
type UserDelete struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserDelete builder.
func (d *UserDelete) Where(ps ...predicate.User) *UserDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *UserDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.gremlinExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *UserDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *UserDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := d.gremlin().Query()
	if err := d.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	d.mutation.done = true
	return res.ReadInt()
}

func (d *UserDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(user.Label)
	for _, p := range d.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// UserDeleteOne is the builder for deleting a single User entity.
type UserDeleteOne struct {
	d *UserDelete
}

// Where appends a list predicates to the UserDelete builder.
func (d *UserDeleteOne) Where(ps ...predicate.User) *UserDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *UserDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{user.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *UserDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
