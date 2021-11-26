// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"

	"entgo.io/ent/entc/integration/gremlin/ent/fieldtype"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	hooks    []Hook
	mutation *FieldTypeMutation
}

// Where appends a list predicates to the FieldTypeDelete builder.
func (ftd *FieldTypeDelete) Where(ps ...predicate.FieldType) *FieldTypeDelete {
	ftd.mutation.Where(ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FieldTypeDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftd.hooks) == 0 {
		affected, err = ftd.gremlinExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FieldTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ftd.mutation = mutation
			affected, err = ftd.gremlinExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftd.hooks) - 1; i >= 0; i-- {
			if ftd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FieldTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FieldTypeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftd.gremlin().Query()
	if err := ftd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftd *FieldTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(fieldtype.Label)
	for _, p := range ftd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// FieldTypeDeleteOne is the builder for deleting a single FieldType entity.
type FieldTypeDeleteOne struct {
	ftd *FieldTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FieldTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fieldtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FieldTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
