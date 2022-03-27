// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/conversion"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// ConversionDelete is the builder for deleting a Conversion entity.
type ConversionDelete struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// Where appends a list predicates to the ConversionDelete builder.
func (cd *ConversionDelete) Where(ps ...predicate.Conversion) *ConversionDelete {
	cd.mutation.Where(ps...)
	return cd
}

// WhereIf appends a list predicates to the ConversionDelete builder if b is true.
func (cd *ConversionDelete) WhereIf(b bool, ps ...predicate.Conversion) *ConversionDelete {
	if b {
		cd.mutation.Where(ps...)
	}
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *ConversionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cd.hooks) == 0 {
		affected, err = cd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ConversionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cd.mutation = mutation
			affected, err = cd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cd.hooks) - 1; i >= 0; i-- {
			if cd.hooks[i] == nil {
				return 0, fmt.Errorf("entv2: uninitialized hook (forgotten import entv2/runtime?)")
			}
			mut = cd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *ConversionDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *ConversionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: conversion.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: conversion.FieldID,
			},
		},
	}
	if ps := cd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, cd.driver, _spec)
}

// ConversionDeleteOne is the builder for deleting a single Conversion entity.
type ConversionDeleteOne struct {
	cd *ConversionDelete
}

// Exec executes the deletion query.
func (cdo *ConversionDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{conversion.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *ConversionDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
