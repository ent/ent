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
	"entgo.io/ent/entc/integration/ent/spec"
	"entgo.io/ent/schema/field"
)

// SpecDelete is the builder for deleting a Spec entity.
type SpecDelete struct {
	config
	hooks    []Hook
	mutation *SpecMutation
}

// Where appends a list predicates to the SpecDelete builder.
func (sd *SpecDelete) Where(ps ...predicate.Spec) *SpecDelete {
	sd.mutation.Where(ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *SpecDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sd.sqlExec, sd.mutation, sd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *SpecDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Mutation returns the SpecMutation object of the builder.
func (sd *SpecDelete) Mutation() *SpecMutation {
	return sd.mutation
}

func (sd *SpecDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(spec.Table, sqlgraph.NewFieldSpec(spec.FieldID, field.TypeInt))
	if ps := sd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sd.mutation.done = true
	return affected, err
}

// SpecDeleteOne is the builder for deleting a single Spec entity.
type SpecDeleteOne struct {
	sd *SpecDelete
}

// Where appends a list predicates to the SpecDelete builder.
func (sdo *SpecDeleteOne) Where(ps ...predicate.Spec) *SpecDeleteOne {
	sdo.sd.mutation.Where(ps...)
	return sdo
}

// Exec executes the deletion query.
func (sdo *SpecDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{spec.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *SpecDeleteOne) ExecX(ctx context.Context) {
	if err := sdo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Mutation returns the SpecMutation object of the builder.
func (sdo *SpecDeleteOne) Mutation() *SpecMutation {
	return sdo.sd.mutation
}
