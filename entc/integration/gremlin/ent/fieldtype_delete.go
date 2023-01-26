// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/fieldtype"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
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
	return withHooks[int, FieldTypeMutation](ctx, ftd.gremlinExec, ftd.mutation, ftd.hooks)
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
	ftd.mutation.done = true
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

// Where appends a list predicates to the FieldTypeDelete builder.
func (ftdo *FieldTypeDeleteOne) Where(ps ...predicate.FieldType) *FieldTypeDeleteOne {
	ftdo.ftd.mutation.Where(ps...)
	return ftdo
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
	if err := ftdo.Exec(ctx); err != nil {
		panic(err)
	}
}
