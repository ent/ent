// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/group"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
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
	return withHooks(ctx, gd.sqlExec, gd.mutation, gd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GroupDelete) ExecX(ctx context.Context) int {
	n, err := gd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Mutation returns the GroupMutation object of the builder.
func (gd *GroupDelete) Mutation() *GroupMutation {
	return gd.mutation
}

func (gd *GroupDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := gd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, gd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	gd.mutation.done = true
	return affected, err
}

// GroupDeleteOne is the builder for deleting a single Group entity.
type GroupDeleteOne struct {
	gd *GroupDelete
}

// Where appends a list predicates to the GroupDelete builder.
func (gdo *GroupDeleteOne) Where(ps ...predicate.Group) *GroupDeleteOne {
	gdo.gd.mutation.Where(ps...)
	return gdo
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
	if err := gdo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Mutation returns the GroupMutation object of the builder.
func (gdo *GroupDeleteOne) Mutation() *GroupMutation {
	return gdo.gd.mutation
}
