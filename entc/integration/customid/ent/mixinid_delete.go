// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// MixinIDDelete is the builder for deleting a MixinID entity.
type MixinIDDelete struct {
	config
	hooks    []Hook
	mutation *MixinIDMutation
}

// Where appends a list predicates to the MixinIDDelete builder.
func (mid *MixinIDDelete) Where(ps ...predicate.MixinID) *MixinIDDelete {
	mid.mutation.Where(ps...)
	return mid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (mid *MixinIDDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mid.hooks) == 0 {
		affected, err = mid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixinIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mid.mutation = mutation
			affected, err = mid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mid.hooks) - 1; i >= 0; i-- {
			if mid.hooks[i] == nil {
				return 0, errors.New("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = mid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (mid *MixinIDDelete) ExecX(ctx context.Context) int {
	n, err := mid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (mid *MixinIDDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mixinid.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		},
	}
	if ps := mid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, mid.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MixinIDDeleteOne is the builder for deleting a single MixinID entity.
type MixinIDDeleteOne struct {
	mid *MixinIDDelete
}

// Exec executes the deletion query.
func (mido *MixinIDDeleteOne) Exec(ctx context.Context) error {
	n, err := mido.mid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mixinid.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (mido *MixinIDDeleteOne) ExecX(ctx context.Context) {
	mido.mid.ExecX(ctx)
}
