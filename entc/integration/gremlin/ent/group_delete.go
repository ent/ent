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
	"entgo.io/ent/entc/integration/gremlin/ent/group"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// GroupDelete is the builder for deleting a Group entity.
type GroupDelete struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupDelete builder.
func (gd *GroupDelete) Where(ps ...predicate.Group) *GroupDelete {
	gd.mutation.Where(ps...)
	return gd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gd *GroupDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, GroupMutation](ctx, gd.exec, gd.mutation, gd.hooks)
}

func (gd *GroupDelete) exec(ctx context.Context) (int, error) {
	return gd.gremlinExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GroupDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gd *GroupDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gd.gremlin().Query()
	if err := gd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	gd.mutation.done = true
	return res.ReadInt()
}

func (gd *GroupDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(group.Label)
	for _, p := range gd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// GroupDeleteOne is the builder for deleting a single Group entity.
type GroupDeleteOne struct {
	gd *GroupDelete
}

// Exec executes the deletion query.
func (gdo *GroupDeleteOne) Exec(ctx context.Context) error {
	n, err := gdo.gd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{group.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GroupDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}
