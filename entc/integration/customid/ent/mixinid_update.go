// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/integration/customid/ent/mixinid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
)

// MixinIDUpdate is the builder for updating MixinID entities.
type MixinIDUpdate struct {
	config
	hooks    []Hook
	mutation *MixinIDMutation
}

// Where appends a list predicates to the MixinIDUpdate builder.
func (miu *MixinIDUpdate) Where(ps ...predicate.MixinID) *MixinIDUpdate {
	miu.mutation.Where(ps...)
	return miu
}

// SetSomeField sets the "some_field" field.
func (miu *MixinIDUpdate) SetSomeField(s string) *MixinIDUpdate {
	miu.mutation.SetSomeField(s)
	return miu
}

// SetMixinField sets the "mixin_field" field.
func (miu *MixinIDUpdate) SetMixinField(s string) *MixinIDUpdate {
	miu.mutation.SetMixinField(s)
	return miu
}

// Mutation returns the MixinIDMutation object of the builder.
func (miu *MixinIDUpdate) Mutation() *MixinIDMutation {
	return miu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (miu *MixinIDUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(miu.hooks) == 0 {
		affected, err = miu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixinIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			miu.mutation = mutation
			affected, err = miu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(miu.hooks) - 1; i >= 0; i-- {
			if miu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = miu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, miu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (miu *MixinIDUpdate) SaveX(ctx context.Context) int {
	affected, err := miu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (miu *MixinIDUpdate) Exec(ctx context.Context) error {
	_, err := miu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miu *MixinIDUpdate) ExecX(ctx context.Context) {
	if err := miu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (miu *MixinIDUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mixinid.Table,
			Columns: mixinid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		},
	}
	if ps := miu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miu.mutation.SomeField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldSomeField,
		})
	}
	if value, ok := miu.mutation.MixinField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldMixinField,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, miu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mixinid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// MixinIDUpdateOne is the builder for updating a single MixinID entity.
type MixinIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MixinIDMutation
}

// SetSomeField sets the "some_field" field.
func (miuo *MixinIDUpdateOne) SetSomeField(s string) *MixinIDUpdateOne {
	miuo.mutation.SetSomeField(s)
	return miuo
}

// SetMixinField sets the "mixin_field" field.
func (miuo *MixinIDUpdateOne) SetMixinField(s string) *MixinIDUpdateOne {
	miuo.mutation.SetMixinField(s)
	return miuo
}

// Mutation returns the MixinIDMutation object of the builder.
func (miuo *MixinIDUpdateOne) Mutation() *MixinIDMutation {
	return miuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (miuo *MixinIDUpdateOne) Select(field string, fields ...string) *MixinIDUpdateOne {
	miuo.fields = append([]string{field}, fields...)
	return miuo
}

// Save executes the query and returns the updated MixinID entity.
func (miuo *MixinIDUpdateOne) Save(ctx context.Context) (*MixinID, error) {
	var (
		err  error
		node *MixinID
	)
	if len(miuo.hooks) == 0 {
		node, err = miuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MixinIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			miuo.mutation = mutation
			node, err = miuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(miuo.hooks) - 1; i >= 0; i-- {
			if miuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = miuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, miuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (miuo *MixinIDUpdateOne) SaveX(ctx context.Context) *MixinID {
	node, err := miuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (miuo *MixinIDUpdateOne) Exec(ctx context.Context) error {
	_, err := miuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (miuo *MixinIDUpdateOne) ExecX(ctx context.Context) {
	if err := miuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (miuo *MixinIDUpdateOne) sqlSave(ctx context.Context) (_node *MixinID, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   mixinid.Table,
			Columns: mixinid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: mixinid.FieldID,
			},
		},
	}
	id, ok := miuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "MixinID.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := miuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, mixinid.FieldID)
		for _, f := range fields {
			if !mixinid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != mixinid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := miuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := miuo.mutation.SomeField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldSomeField,
		})
	}
	if value, ok := miuo.mutation.MixinField(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: mixinid.FieldMixinField,
		})
	}
	_node = &MixinID{config: miuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, miuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{mixinid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
