// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

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
func (ud *UserDelete) Where(ps ...predicate.User) *UserDelete {
	ud.mutation.Where(ps...)
	return ud
}

// When runs the provided builder(s) if and only if condition is true.
func (ud *UserDelete) When(condition bool, action func(builder *UserDelete)) *UserDelete {
	if condition {
		action(ud)
	}

	return ud
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ud *UserDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ud.hooks) == 0 {
		affected, err = ud.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ud.mutation = mutation
			affected, err = ud.gremlinExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ud.hooks) - 1; i >= 0; i-- {
			if ud.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ud.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ud.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserDelete) ExecX(ctx context.Context) int {
	n, err := ud.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
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
	for _, p := range ud.mutation.predicates {
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
		return &NotFoundError{user.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
