// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/predicate"
	"entgo.io/ent/examples/migration/ent/session"
	"entgo.io/ent/examples/migration/ent/sessiondevice"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
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

// SetActive sets the "active" field.
func (_u *SessionUpdate) SetActive(v bool) *SessionUpdate {
	_u.mutation.SetActive(v)
	return _u
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (_u *SessionUpdate) SetNillableActive(v *bool) *SessionUpdate {
	if v != nil {
		_u.SetActive(*v)
	}
	return _u
}

// SetIssuedAt sets the "issued_at" field.
func (_u *SessionUpdate) SetIssuedAt(v time.Time) *SessionUpdate {
	_u.mutation.SetIssuedAt(v)
	return _u
}

// SetNillableIssuedAt sets the "issued_at" field if the given value is not nil.
func (_u *SessionUpdate) SetNillableIssuedAt(v *time.Time) *SessionUpdate {
	if v != nil {
		_u.SetIssuedAt(*v)
	}
	return _u
}

// SetExpiresAt sets the "expires_at" field.
func (_u *SessionUpdate) SetExpiresAt(v time.Time) *SessionUpdate {
	_u.mutation.SetExpiresAt(v)
	return _u
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (_u *SessionUpdate) SetNillableExpiresAt(v *time.Time) *SessionUpdate {
	if v != nil {
		_u.SetExpiresAt(*v)
	}
	return _u
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (_u *SessionUpdate) ClearExpiresAt() *SessionUpdate {
	_u.mutation.ClearExpiresAt()
	return _u
}

// SetToken sets the "token" field.
func (_u *SessionUpdate) SetToken(v string) *SessionUpdate {
	_u.mutation.SetToken(v)
	return _u
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (_u *SessionUpdate) SetNillableToken(v *string) *SessionUpdate {
	if v != nil {
		_u.SetToken(*v)
	}
	return _u
}

// ClearToken clears the value of the "token" field.
func (_u *SessionUpdate) ClearToken() *SessionUpdate {
	_u.mutation.ClearToken()
	return _u
}

// SetMethod sets the "method" field.
func (_u *SessionUpdate) SetMethod(v map[string]interface{}) *SessionUpdate {
	_u.mutation.SetMethod(v)
	return _u
}

// ClearMethod clears the value of the "method" field.
func (_u *SessionUpdate) ClearMethod() *SessionUpdate {
	_u.mutation.ClearMethod()
	return _u
}

// SetDeviceID sets the "device_id" field.
func (_u *SessionUpdate) SetDeviceID(v uuid.UUID) *SessionUpdate {
	_u.mutation.SetDeviceID(v)
	return _u
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (_u *SessionUpdate) SetNillableDeviceID(v *uuid.UUID) *SessionUpdate {
	if v != nil {
		_u.SetDeviceID(*v)
	}
	return _u
}

// ClearDeviceID clears the value of the "device_id" field.
func (_u *SessionUpdate) ClearDeviceID() *SessionUpdate {
	_u.mutation.ClearDeviceID()
	return _u
}

// SetDevice sets the "device" edge to the SessionDevice entity.
func (_u *SessionUpdate) SetDevice(v *SessionDevice) *SessionUpdate {
	return _u.SetDeviceID(v.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (_u *SessionUpdate) Mutation() *SessionMutation {
	return _u.mutation
}

// ClearDevice clears the "device" edge to the SessionDevice entity.
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

func (_u *SessionUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Active(); ok {
		_spec.SetField(session.FieldActive, field.TypeBool, value)
	}
	if value, ok := _u.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
	}
	if value, ok := _u.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if _u.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := _u.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if _u.mutation.TokenCleared() {
		_spec.ClearField(session.FieldToken, field.TypeString)
	}
	if value, ok := _u.mutation.Method(); ok {
		_spec.SetField(session.FieldMethod, field.TypeJSON, value)
	}
	if _u.mutation.MethodCleared() {
		_spec.ClearField(session.FieldMethod, field.TypeJSON)
	}
	if _u.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID),
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
				IDSpec: sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{session.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// SessionUpdateOne is the builder for updating a single Session entity.
type SessionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SessionMutation
}

// SetActive sets the "active" field.
func (_u *SessionUpdateOne) SetActive(v bool) *SessionUpdateOne {
	_u.mutation.SetActive(v)
	return _u
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableActive(v *bool) *SessionUpdateOne {
	if v != nil {
		_u.SetActive(*v)
	}
	return _u
}

// SetIssuedAt sets the "issued_at" field.
func (_u *SessionUpdateOne) SetIssuedAt(v time.Time) *SessionUpdateOne {
	_u.mutation.SetIssuedAt(v)
	return _u
}

// SetNillableIssuedAt sets the "issued_at" field if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableIssuedAt(v *time.Time) *SessionUpdateOne {
	if v != nil {
		_u.SetIssuedAt(*v)
	}
	return _u
}

// SetExpiresAt sets the "expires_at" field.
func (_u *SessionUpdateOne) SetExpiresAt(v time.Time) *SessionUpdateOne {
	_u.mutation.SetExpiresAt(v)
	return _u
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableExpiresAt(v *time.Time) *SessionUpdateOne {
	if v != nil {
		_u.SetExpiresAt(*v)
	}
	return _u
}

// ClearExpiresAt clears the value of the "expires_at" field.
func (_u *SessionUpdateOne) ClearExpiresAt() *SessionUpdateOne {
	_u.mutation.ClearExpiresAt()
	return _u
}

// SetToken sets the "token" field.
func (_u *SessionUpdateOne) SetToken(v string) *SessionUpdateOne {
	_u.mutation.SetToken(v)
	return _u
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableToken(v *string) *SessionUpdateOne {
	if v != nil {
		_u.SetToken(*v)
	}
	return _u
}

// ClearToken clears the value of the "token" field.
func (_u *SessionUpdateOne) ClearToken() *SessionUpdateOne {
	_u.mutation.ClearToken()
	return _u
}

// SetMethod sets the "method" field.
func (_u *SessionUpdateOne) SetMethod(v map[string]interface{}) *SessionUpdateOne {
	_u.mutation.SetMethod(v)
	return _u
}

// ClearMethod clears the value of the "method" field.
func (_u *SessionUpdateOne) ClearMethod() *SessionUpdateOne {
	_u.mutation.ClearMethod()
	return _u
}

// SetDeviceID sets the "device_id" field.
func (_u *SessionUpdateOne) SetDeviceID(v uuid.UUID) *SessionUpdateOne {
	_u.mutation.SetDeviceID(v)
	return _u
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (_u *SessionUpdateOne) SetNillableDeviceID(v *uuid.UUID) *SessionUpdateOne {
	if v != nil {
		_u.SetDeviceID(*v)
	}
	return _u
}

// ClearDeviceID clears the value of the "device_id" field.
func (_u *SessionUpdateOne) ClearDeviceID() *SessionUpdateOne {
	_u.mutation.ClearDeviceID()
	return _u
}

// SetDevice sets the "device" edge to the SessionDevice entity.
func (_u *SessionUpdateOne) SetDevice(v *SessionDevice) *SessionUpdateOne {
	return _u.SetDeviceID(v.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (_u *SessionUpdateOne) Mutation() *SessionMutation {
	return _u.mutation
}

// ClearDevice clears the "device" edge to the SessionDevice entity.
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
	_spec := sqlgraph.NewUpdateSpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
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
	if value, ok := _u.mutation.Active(); ok {
		_spec.SetField(session.FieldActive, field.TypeBool, value)
	}
	if value, ok := _u.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
	}
	if value, ok := _u.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
	}
	if _u.mutation.ExpiresAtCleared() {
		_spec.ClearField(session.FieldExpiresAt, field.TypeTime)
	}
	if value, ok := _u.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
	}
	if _u.mutation.TokenCleared() {
		_spec.ClearField(session.FieldToken, field.TypeString)
	}
	if value, ok := _u.mutation.Method(); ok {
		_spec.SetField(session.FieldMethod, field.TypeJSON, value)
	}
	if _u.mutation.MethodCleared() {
		_spec.ClearField(session.FieldMethod, field.TypeJSON)
	}
	if _u.mutation.DeviceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   session.DeviceTable,
			Columns: []string{session.DeviceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID),
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
				IDSpec: sqlgraph.NewFieldSpec(sessiondevice.FieldID, field.TypeUUID),
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
