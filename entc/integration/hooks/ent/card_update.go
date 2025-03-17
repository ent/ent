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
	"entgo.io/ent/entc/integration/hooks/ent/card"
	"entgo.io/ent/entc/integration/hooks/ent/predicate"
	"entgo.io/ent/entc/integration/hooks/ent/user"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (u *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetName sets the "name" field.
func (m *CardUpdate) SetName(v string) *CardUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CardUpdate) SetNillableName(v *string) *CardUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// ClearName clears the value of the "name" field.
func (m *CardUpdate) ClearName() *CardUpdate {
	m.mutation.ClearName()
	return m
}

// SetCreatedAt sets the "created_at" field.
func (m *CardUpdate) SetCreatedAt(v time.Time) *CardUpdate {
	m.mutation.SetCreatedAt(v)
	return m
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (m *CardUpdate) SetNillableCreatedAt(v *time.Time) *CardUpdate {
	if v != nil {
		m.SetCreatedAt(*v)
	}
	return m
}

// SetInHook sets the "in_hook" field.
func (m *CardUpdate) SetInHook(v string) *CardUpdate {
	m.mutation.SetInHook(v)
	return m
}

// SetNillableInHook sets the "in_hook" field if the given value is not nil.
func (m *CardUpdate) SetNillableInHook(v *string) *CardUpdate {
	if v != nil {
		m.SetInHook(*v)
	}
	return m
}

// SetExpiredAt sets the "expired_at" field.
func (m *CardUpdate) SetExpiredAt(v time.Time) *CardUpdate {
	m.mutation.SetExpiredAt(v)
	return m
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (m *CardUpdate) SetNillableExpiredAt(v *time.Time) *CardUpdate {
	if v != nil {
		m.SetExpiredAt(*v)
	}
	return m
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (m *CardUpdate) ClearExpiredAt() *CardUpdate {
	m.mutation.ClearExpiredAt()
	return m
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (m *CardUpdate) SetOwnerID(id int) *CardUpdate {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (m *CardUpdate) SetNillableOwnerID(id *int) *CardUpdate {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *CardUpdate) SetOwner(v *User) *CardUpdate {
	return m.SetOwnerID(v.ID)
}

// Mutation returns the CardMutation object of the builder.
func (m *CardUpdate) Mutation() *CardMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *CardUpdate) ClearOwner() *CardUpdate {
	u.mutation.ClearOwner()
	return u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *CardUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CardUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if u.mutation.NameCleared() {
		_spec.ClearField(card.FieldName, field.TypeString)
	}
	if value, ok := u.mutation.CreatedAt(); ok {
		_spec.SetField(card.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := u.mutation.InHook(); ok {
		_spec.SetField(card.FieldInHook, field.TypeString, value)
	}
	if value, ok := u.mutation.ExpiredAt(); ok {
		_spec.SetField(card.FieldExpiredAt, field.TypeTime, value)
	}
	if u.mutation.ExpiredAtCleared() {
		_spec.ClearField(card.FieldExpiredAt, field.TypeTime)
	}
	if u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetName sets the "name" field.
func (m *CardUpdateOne) SetName(v string) *CardUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableName(v *string) *CardUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// ClearName clears the value of the "name" field.
func (m *CardUpdateOne) ClearName() *CardUpdateOne {
	m.mutation.ClearName()
	return m
}

// SetCreatedAt sets the "created_at" field.
func (m *CardUpdateOne) SetCreatedAt(v time.Time) *CardUpdateOne {
	m.mutation.SetCreatedAt(v)
	return m
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableCreatedAt(v *time.Time) *CardUpdateOne {
	if v != nil {
		m.SetCreatedAt(*v)
	}
	return m
}

// SetInHook sets the "in_hook" field.
func (m *CardUpdateOne) SetInHook(v string) *CardUpdateOne {
	m.mutation.SetInHook(v)
	return m
}

// SetNillableInHook sets the "in_hook" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableInHook(v *string) *CardUpdateOne {
	if v != nil {
		m.SetInHook(*v)
	}
	return m
}

// SetExpiredAt sets the "expired_at" field.
func (m *CardUpdateOne) SetExpiredAt(v time.Time) *CardUpdateOne {
	m.mutation.SetExpiredAt(v)
	return m
}

// SetNillableExpiredAt sets the "expired_at" field if the given value is not nil.
func (m *CardUpdateOne) SetNillableExpiredAt(v *time.Time) *CardUpdateOne {
	if v != nil {
		m.SetExpiredAt(*v)
	}
	return m
}

// ClearExpiredAt clears the value of the "expired_at" field.
func (m *CardUpdateOne) ClearExpiredAt() *CardUpdateOne {
	m.mutation.ClearExpiredAt()
	return m
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (m *CardUpdateOne) SetOwnerID(id int) *CardUpdateOne {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (m *CardUpdateOne) SetNillableOwnerID(id *int) *CardUpdateOne {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *CardUpdateOne) SetOwner(v *User) *CardUpdateOne {
	return m.SetOwnerID(v.ID)
}

// Mutation returns the CardMutation object of the builder.
func (m *CardUpdateOne) Mutation() *CardMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *CardUpdateOne) ClearOwner() *CardUpdateOne {
	u.mutation.ClearOwner()
	return u
}

// Where appends a list predicates to the CardUpdate builder.
func (u *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Card entity.
func (u *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CardUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CardUpdateOne) sqlSave(ctx context.Context) (_n *Card, err error) {
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, card.FieldID)
		for _, f := range fields {
			if !card.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != card.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(card.FieldName, field.TypeString, value)
	}
	if u.mutation.NameCleared() {
		_spec.ClearField(card.FieldName, field.TypeString)
	}
	if value, ok := u.mutation.CreatedAt(); ok {
		_spec.SetField(card.FieldCreatedAt, field.TypeTime, value)
	}
	if value, ok := u.mutation.InHook(); ok {
		_spec.SetField(card.FieldInHook, field.TypeString, value)
	}
	if value, ok := u.mutation.ExpiredAt(); ok {
		_spec.SetField(card.FieldExpiredAt, field.TypeTime, value)
	}
	if u.mutation.ExpiredAtCleared() {
		_spec.ClearField(card.FieldExpiredAt, field.TypeTime)
	}
	if u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   card.OwnerTable,
			Columns: []string{card.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &Card{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
