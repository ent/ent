// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	hooks    []Hook
	mutation *GroupInfoMutation
}

// Where appends a list predicates to the GroupInfoDelete builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.mutation.Where(ps...)
	return gid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gid *GroupInfoDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(gid.hooks) == 0 {
		affected, err = gid.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			gid.mutation = mutation
			affected, err = gid.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(gid.hooks) - 1; i >= 0; i-- {
			if gid.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gid.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gid.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) int {
	n, err := gid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gid *GroupInfoDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: groupinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		},
	}
	if ps := gid.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, gid.driver, _spec)
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := gido.gid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{groupinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	gido.gid.ExecX(ctx)
}
