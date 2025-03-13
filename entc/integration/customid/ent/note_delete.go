// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/note"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// NoteDelete is the builder for deleting a Note entity.
type NoteDelete struct {
	config
	hooks    []Hook
	mutation *NoteMutation
}

// Where appends a list predicates to the NoteDelete builder.
func (nd *NoteDelete) Where(ps ...predicate.Note) *NoteDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NoteDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, nd.sqlExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NoteDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Mutation returns the NoteMutation object of the builder.
func (nd *NoteDelete) Mutation() *NoteMutation {
	return nd.mutation
}

func (nd *NoteDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(note.Table, sqlgraph.NewFieldSpec(note.FieldID, field.TypeString))
	if ps := nd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, nd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	nd.mutation.done = true
	return affected, err
}

// NoteDeleteOne is the builder for deleting a single Note entity.
type NoteDeleteOne struct {
	nd *NoteDelete
}

// Where appends a list predicates to the NoteDelete builder.
func (ndo *NoteDeleteOne) Where(ps ...predicate.Note) *NoteDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NoteDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{note.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NoteDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Mutation returns the NoteMutation object of the builder.
func (ndo *NoteDeleteOne) Mutation() *NoteMutation {
	return ndo.nd.mutation
}
