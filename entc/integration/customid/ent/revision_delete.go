// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/revision"
	"entgo.io/ent/schema/field"
)

// RevisionDelete is the builder for deleting a Revision entity.
type RevisionDelete struct {
	config
	hooks    []Hook
	mutation *RevisionMutation
}

// Where appends a list predicates to the RevisionDelete builder.
func (d *RevisionDelete) Where(ps ...predicate.Revision) *RevisionDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *RevisionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.sqlExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *RevisionDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *RevisionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(revision.Table, sqlgraph.NewFieldSpec(revision.FieldID, field.TypeString))
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

// RevisionDeleteOne is the builder for deleting a single Revision entity.
type RevisionDeleteOne struct {
	d *RevisionDelete
}

// Where appends a list predicates to the RevisionDelete builder.
func (d *RevisionDeleteOne) Where(ps ...predicate.Revision) *RevisionDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *RevisionDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{revision.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *RevisionDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
