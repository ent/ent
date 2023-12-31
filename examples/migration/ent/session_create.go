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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/migration/ent/session"
	"entgo.io/ent/examples/migration/ent/sessiondevice"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// SessionCreate is the builder for creating a Session entity.
type SessionCreate struct {
	config
	mutation *SessionMutation
	hooks    []Hook
}

// SetActive sets the "active" field.
func (sc *SessionCreate) SetActive(b bool) *SessionCreate {
	sc.mutation.SetActive(b)
	return sc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (sc *SessionCreate) SetNillableActive(b *bool) *SessionCreate {
	if b != nil {
		sc.SetActive(*b)
	}
	return sc
}

// SetIssuedAt sets the "issued_at" field.
func (sc *SessionCreate) SetIssuedAt(t time.Time) *SessionCreate {
	sc.mutation.SetIssuedAt(t)
	return sc
}

// SetExpiresAt sets the "expires_at" field.
func (sc *SessionCreate) SetExpiresAt(t time.Time) *SessionCreate {
	sc.mutation.SetExpiresAt(t)
	return sc
}

// SetNillableExpiresAt sets the "expires_at" field if the given value is not nil.
func (sc *SessionCreate) SetNillableExpiresAt(t *time.Time) *SessionCreate {
	if t != nil {
		sc.SetExpiresAt(*t)
	}
	return sc
}

// SetToken sets the "token" field.
func (sc *SessionCreate) SetToken(s string) *SessionCreate {
	sc.mutation.SetToken(s)
	return sc
}

// SetNillableToken sets the "token" field if the given value is not nil.
func (sc *SessionCreate) SetNillableToken(s *string) *SessionCreate {
	if s != nil {
		sc.SetToken(*s)
	}
	return sc
}

// SetMethod sets the "method" field.
func (sc *SessionCreate) SetMethod(m map[string]interface{}) *SessionCreate {
	sc.mutation.SetMethod(m)
	return sc
}

// SetDeviceID sets the "device_id" field.
func (sc *SessionCreate) SetDeviceID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetDeviceID(u)
	return sc
}

// SetNillableDeviceID sets the "device_id" field if the given value is not nil.
func (sc *SessionCreate) SetNillableDeviceID(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetDeviceID(*u)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SessionCreate) SetID(u uuid.UUID) *SessionCreate {
	sc.mutation.SetID(u)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SessionCreate) SetNillableID(u *uuid.UUID) *SessionCreate {
	if u != nil {
		sc.SetID(*u)
	}
	return sc
}

// SetDevice sets the "device" edge to the SessionDevice entity.
func (sc *SessionCreate) SetDevice(s *SessionDevice) *SessionCreate {
	return sc.SetDeviceID(s.ID)
}

// Mutation returns the SessionMutation object of the builder.
func (sc *SessionCreate) Mutation() *SessionMutation {
	return sc.mutation
}

// Save creates the Session in the database.
func (sc *SessionCreate) Save(ctx context.Context) (*Session, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SessionCreate) SaveX(ctx context.Context) *Session {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SessionCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SessionCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SessionCreate) defaults() {
	if _, ok := sc.mutation.Active(); !ok {
		v := session.DefaultActive
		sc.mutation.SetActive(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := session.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SessionCreate) check() error {
	if _, ok := sc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Session.active"`)}
	}
	if _, ok := sc.mutation.IssuedAt(); !ok {
		return &ValidationError{Name: "issued_at", err: errors.New(`ent: missing required field "Session.issued_at"`)}
	}
	return nil
}

func (sc *SessionCreate) sqlSave(ctx context.Context) (*Session, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SessionCreate) createSpec() (*Session, *sqlgraph.CreateSpec) {
	var (
		_node = &Session{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(session.Table, sqlgraph.NewFieldSpec(session.FieldID, field.TypeUUID))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.Active(); ok {
		_spec.SetField(session.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := sc.mutation.IssuedAt(); ok {
		_spec.SetField(session.FieldIssuedAt, field.TypeTime, value)
		_node.IssuedAt = value
	}
	if value, ok := sc.mutation.ExpiresAt(); ok {
		_spec.SetField(session.FieldExpiresAt, field.TypeTime, value)
		_node.ExpiresAt = value
	}
	if value, ok := sc.mutation.Token(); ok {
		_spec.SetField(session.FieldToken, field.TypeString, value)
		_node.Token = value
	}
	if value, ok := sc.mutation.Method(); ok {
		_spec.SetField(session.FieldMethod, field.TypeJSON, value)
		_node.Method = value
	}
	if nodes := sc.mutation.DeviceIDs(); len(nodes) > 0 {
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
		_node.DeviceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SessionCreateBulk is the builder for creating many Session entities in bulk.
type SessionCreateBulk struct {
	config
	err      error
	builders []*SessionCreate
}

// Save creates the Session entities in the database.
func (scb *SessionCreateBulk) Save(ctx context.Context) ([]*Session, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Session, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SessionMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SessionCreateBulk) SaveX(ctx context.Context) []*Session {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SessionCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SessionCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
