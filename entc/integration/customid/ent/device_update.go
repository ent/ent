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

// DeviceUpdate is the builder for updating Device entities.
type DeviceUpdate struct {
	config
	hooks    []Hook
	mutation *DeviceMutation
}

// Where appends a list predicates to the DeviceUpdate builder.
func (du *DeviceUpdate) Where(ps ...predicate.Device) *DeviceUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetActiveSessionID sets the "active_session" edge to the Session entity by ID.
func (du *DeviceUpdate) SetActiveSessionID(id schema.ID) *DeviceUpdate {
	du.mutation.SetActiveSessionID(id)
	return du
}

// SetNillableActiveSessionID sets the "active_session" edge to the Session entity by ID if the given value is not nil.
func (du *DeviceUpdate) SetNillableActiveSessionID(id *schema.ID) *DeviceUpdate {
	if id != nil {
		du = du.SetActiveSessionID(*id)
	}
	return du
}

// SetActiveSession sets the "active_session" edge to the Session entity.
func (du *DeviceUpdate) SetActiveSession(s *Session) *DeviceUpdate {
	return du.SetActiveSessionID(s.ID)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (du *DeviceUpdate) AddSessionIDs(ids ...schema.ID) *DeviceUpdate {
	du.mutation.AddSessionIDs(ids...)
	return du
}

// AddSessions adds the "sessions" edges to the Session entity.
func (du *DeviceUpdate) AddSessions(s ...*Session) *DeviceUpdate {
	ids := make([]schema.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.AddSessionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (du *DeviceUpdate) Mutation() *DeviceMutation {
	return du.mutation
}

// ClearActiveSession clears the "active_session" edge to the Session entity.
func (du *DeviceUpdate) ClearActiveSession() *DeviceUpdate {
	du.mutation.ClearActiveSession()
	return du
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (du *DeviceUpdate) ClearSessions() *DeviceUpdate {
	du.mutation.ClearSessions()
	return du
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (du *DeviceUpdate) RemoveSessionIDs(ids ...schema.ID) *DeviceUpdate {
	du.mutation.RemoveSessionIDs(ids...)
	return du
}

// RemoveSessions removes "sessions" edges to Session entities.
func (du *DeviceUpdate) RemoveSessions(s ...*Session) *DeviceUpdate {
	ids := make([]schema.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return du.RemoveSessionIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DeviceUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(du.hooks) == 0 {
		affected, err = du.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			du.mutation = mutation
			affected, err = du.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(du.hooks) - 1; i >= 0; i-- {
			if du.hooks[i] == nil {
				return 0, errors.New("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = du.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, du.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (du *DeviceUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DeviceUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DeviceUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

func (du *DeviceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeBytes,
				Column: device.FieldID,
			},
		},
	}
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if du.mutation.ActiveSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.ActiveSessionTable,
			Columns: []string{device.ActiveSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.ActiveSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.ActiveSessionTable,
			Columns: []string{device.ActiveSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if du.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !du.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// DeviceUpdateOne is the builder for updating a single Device entity.
type DeviceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *DeviceMutation
}

// SetActiveSessionID sets the "active_session" edge to the Session entity by ID.
func (duo *DeviceUpdateOne) SetActiveSessionID(id schema.ID) *DeviceUpdateOne {
	duo.mutation.SetActiveSessionID(id)
	return duo
}

// SetNillableActiveSessionID sets the "active_session" edge to the Session entity by ID if the given value is not nil.
func (duo *DeviceUpdateOne) SetNillableActiveSessionID(id *schema.ID) *DeviceUpdateOne {
	if id != nil {
		duo = duo.SetActiveSessionID(*id)
	}
	return duo
}

// SetActiveSession sets the "active_session" edge to the Session entity.
func (duo *DeviceUpdateOne) SetActiveSession(s *Session) *DeviceUpdateOne {
	return duo.SetActiveSessionID(s.ID)
}

// AddSessionIDs adds the "sessions" edge to the Session entity by IDs.
func (duo *DeviceUpdateOne) AddSessionIDs(ids ...schema.ID) *DeviceUpdateOne {
	duo.mutation.AddSessionIDs(ids...)
	return duo
}

// AddSessions adds the "sessions" edges to the Session entity.
func (duo *DeviceUpdateOne) AddSessions(s ...*Session) *DeviceUpdateOne {
	ids := make([]schema.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.AddSessionIDs(ids...)
}

// Mutation returns the DeviceMutation object of the builder.
func (duo *DeviceUpdateOne) Mutation() *DeviceMutation {
	return duo.mutation
}

// ClearActiveSession clears the "active_session" edge to the Session entity.
func (duo *DeviceUpdateOne) ClearActiveSession() *DeviceUpdateOne {
	duo.mutation.ClearActiveSession()
	return duo
}

// ClearSessions clears all "sessions" edges to the Session entity.
func (duo *DeviceUpdateOne) ClearSessions() *DeviceUpdateOne {
	duo.mutation.ClearSessions()
	return duo
}

// RemoveSessionIDs removes the "sessions" edge to Session entities by IDs.
func (duo *DeviceUpdateOne) RemoveSessionIDs(ids ...schema.ID) *DeviceUpdateOne {
	duo.mutation.RemoveSessionIDs(ids...)
	return duo
}

// RemoveSessions removes "sessions" edges to Session entities.
func (duo *DeviceUpdateOne) RemoveSessions(s ...*Session) *DeviceUpdateOne {
	ids := make([]schema.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return duo.RemoveSessionIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DeviceUpdateOne) Select(field string, fields ...string) *DeviceUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Device entity.
func (duo *DeviceUpdateOne) Save(ctx context.Context) (*Device, error) {
	var (
		err  error
		node *Device
	)
	if len(duo.hooks) == 0 {
		node, err = duo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DeviceMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			duo.mutation = mutation
			node, err = duo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(duo.hooks) - 1; i >= 0; i-- {
			if duo.hooks[i] == nil {
				return nil, errors.New("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = duo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, duo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Device)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from DeviceMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DeviceUpdateOne) SaveX(ctx context.Context) *Device {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DeviceUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DeviceUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (duo *DeviceUpdateOne) sqlSave(ctx context.Context) (_node *Device, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   device.Table,
			Columns: device.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeBytes,
				Column: device.FieldID,
			},
		},
	}
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Device.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, device.FieldID)
		for _, f := range fields {
			if !device.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != device.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if duo.mutation.ActiveSessionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.ActiveSessionTable,
			Columns: []string{device.ActiveSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.ActiveSessionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   device.ActiveSessionTable,
			Columns: []string{device.ActiveSessionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if duo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedSessionsIDs(); len(nodes) > 0 && !duo.mutation.SessionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.SessionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   device.SessionsTable,
			Columns: []string{device.SessionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeBytes,
					Column: session.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Device{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{device.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
