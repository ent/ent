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
	"entgo.io/ent/entc/integration/customid/ent/other"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/schema/field"
)

// OtherUpdate is the builder for updating Other entities.
type OtherUpdate struct {
	config
	hooks    []Hook
	mutation *OtherMutation
}

// Where appends a list predicates to the OtherUpdate builder.
func (ou *OtherUpdate) Where(ps ...predicate.Other) *OtherUpdate {
	ou.mutation.Where(ps...)
	return ou
}

// Mutation returns the OtherMutation object of the builder.
func (ou *OtherUpdate) Mutation() *OtherMutation {
	return ou.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ou *OtherUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, OtherMutation](ctx, ou.sqlSave, ou.mutation, ou.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ou *OtherUpdate) SaveX(ctx context.Context) int {
	affected, err := ou.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ou *OtherUpdate) Exec(ctx context.Context) error {
	_, err := ou.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ou *OtherUpdate) ExecX(ctx context.Context) {
	if err := ou.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ou *OtherUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   other.Table,
			Columns: other.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: other.FieldID,
			},
		},
	}
	if ps := ou.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ou.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{other.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ou.mutation.done = true
	return n, nil
}

// OtherUpdateOne is the builder for updating a single Other entity.
type OtherUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *OtherMutation
}

// Mutation returns the OtherMutation object of the builder.
func (ouo *OtherUpdateOne) Mutation() *OtherMutation {
	return ouo.mutation
}

// Where appends a list predicates to the OtherUpdate builder.
func (ouo *OtherUpdateOne) Where(ps ...predicate.Other) *OtherUpdateOne {
	ouo.mutation.Where(ps...)
	return ouo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ouo *OtherUpdateOne) Select(field string, fields ...string) *OtherUpdateOne {
	ouo.fields = append([]string{field}, fields...)
	return ouo
}

// Save executes the query and returns the updated Other entity.
func (ouo *OtherUpdateOne) Save(ctx context.Context) (*Other, error) {
	return withHooks[*Other, OtherMutation](ctx, ouo.sqlSave, ouo.mutation, ouo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ouo *OtherUpdateOne) SaveX(ctx context.Context) *Other {
	node, err := ouo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ouo *OtherUpdateOne) Exec(ctx context.Context) error {
	_, err := ouo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ouo *OtherUpdateOne) ExecX(ctx context.Context) {
	if err := ouo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ouo *OtherUpdateOne) sqlSave(ctx context.Context) (_node *Other, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   other.Table,
			Columns: other.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: other.FieldID,
			},
		},
	}
	id, ok := ouo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Other.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ouo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, other.FieldID)
		for _, f := range fields {
			if !other.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != other.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ouo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_node = &Other{config: ouo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ouo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{other.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ouo.mutation.done = true
	return _node, nil
}
