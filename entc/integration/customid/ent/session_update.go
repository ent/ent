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
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
	"entgo.io/ent/schema/field"
)

// SessionUpdate is the builder for updating Session entities.
type SessionUpdate struct {
	config
	hooks    []Hook
	mutation *SessionMutation
}

// Where appends a list predicates to the SessionUpdate builder.
func (_u *SessionUpdate) Where(ps ...predicate.Session) *SessionUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (_u *SessionUpdate) SetDeviceID(id schema.ID) *SessionUpdate {
	_u.mutation.SetDeviceID(id)
	return _u
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (_u *SessionUpdate) SetNillableDeviceID(id *schema.ID) *SessionUpdate {
	if id != nil {
		_u = _u.SetDeviceID(*id)
	}
	return _u
}

// SetDevice sets the "device" edge to the Device entity.
func (_u *SessionUpdate) SetDevice(d *Device) *SessionUpdate {
	return _u.SetDeviceID(d.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (_u *SessionUpdate) Mutation() *SessionMutation {
	return _u.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (_u *SessionUpdate) ClearDevice() *SessionUpdate {
	_u.mutation.ClearDevice()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *SessionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *SessionUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *SessionUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *SessionUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *SessionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeBytes))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if _u.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeBytes),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeBytes),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetDeviceID sets the "device" edge to the Device entity by ID.
func (_u *SessionUpdateOne) SetDeviceID(id schema.ID) *SessionUpdateOne {
	_u.mutation.SetDeviceID(id)
	return _u
}

// SetNillableDeviceID sets the "device" edge to the Device entity by ID if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableDeviceID(id *schema.ID) *SessionUpdateOne {
	if id != nil {
		_u = _u.SetDeviceID(*id)
	}
	return _u
}

// SetDevice sets the "device" edge to the Device entity.
func (_u *SessionUpdateOne) SetDevice(d *Device) *SessionUpdateOne {
	return _u.SetDeviceID(d.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (_u *SessionUpdateOne) Mutation() *SessionMutation {
	return _u.mutation
}

// ClearDevice clears the "device" edge to the Device entity.
func (_u *SessionUpdateOne) ClearDevice() *SessionUpdateOne {
	_u.mutation.ClearDevice()
	return _u
}

// Where appends a list predicates to the SessionUpdate builder.
func (_u *SessionUpdateOne) Where(ps ...predicate.Session) *SessionUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *SessionUpdateOne) Select(field string, fields ...string) *SessionUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Session entity.
func (_u *SessionUpdateOne) Save(ctx context.Context) (*Session, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *SessionUpdateOne) SaveX(ctx context.Context) *Session {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *SessionUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *SessionUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *SessionUpdateOne) sqlSave(ctx context.Context) (_node *Session, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeBytes))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Session.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for _, f := range fields {
			if !session.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != session.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if _u.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeBytes),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.DeviceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(device.FieldID, field.TypeBytes),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Session{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
