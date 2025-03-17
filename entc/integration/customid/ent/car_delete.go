// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/car"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// CarDelete is the builder for deleting a Car entity.
type CarDelete struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarDelete builder.
func (d *CarDelete) Where(ps ...predicate.Car) *CarDelete {
	d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (d *CarDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, d.sqlExec, d.mutation, d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (d *CarDelete) ExecX(ctx context.Context) int {
	n, err := d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (d *CarDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
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

// CarDeleteOne is the builder for deleting a single Car entity.
type CarDeleteOne struct {
	d *CarDelete
}

// Where appends a list predicates to the CarDelete builder.
func (d *CarDeleteOne) Where(ps ...predicate.Car) *CarDeleteOne {
	d.d.mutation.Where(ps...)
	return d
}

// Exec executes the deletion query.
func (d *CarDeleteOne) Exec(ctx context.Context) error {
	n, err := d.d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{car.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (d *CarDeleteOne) ExecX(ctx context.Context) {
	if err := d.Exec(ctx); err != nil {
		panic(err)
	}
}
