// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"

	enttask "entgo.io/ent/entc/integration/ent/task"
)

// TaskDelete is the builder for deleting a Task entity.
type TaskDelete struct {
	config
	hooks    []Hook
	mutation *TaskMutation
}

// Where appends a list predicates to the TaskDelete builder.
func (d *TaskDelete) Where(ps ...predicate.Task) *TaskDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *TaskDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.sqlExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *TaskDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *TaskDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(enttask.Table, sqlgraph.NewFieldSpec(enttask.FieldID, field.TypeInt))
	if ps := d.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, d.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	d.mutation.done = true
	return affected, err
}

// TaskDeleteOne is the builder for deleting a single Task entity.
type TaskDeleteOne struct {
	d *TaskDelete
}

// Where appends a list predicates to the TaskDelete builder.
func (d *TaskDeleteOne) Where(ps ...predicate.Task) *TaskDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *TaskDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{enttask.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *TaskDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
