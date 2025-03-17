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
	"entgo.io/ent/examples/o2o2types/ent/card"
	"entgo.io/ent/examples/o2o2types/ent/predicate"
	"entgo.io/ent/examples/o2o2types/ent/user"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (u *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetAge sets the "age" field.
func (m *UserUpdate) SetAge(v int) *UserUpdate {
	m.mutation.ResetAge()
	m.mutation.SetAge(v)
	return m
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (m *UserUpdate) SetNillableAge(v *int) *UserUpdate {
	if v != nil {
		m.SetAge(*v)
	}
	return m
}

// AddAge adds value to the "age" field.
func (m *UserUpdate) AddAge(v int) *UserUpdate {
	m.mutation.AddAge(v)
	return m
}

// SetName sets the "name" field.
func (m *UserUpdate) SetName(v string) *UserUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *UserUpdate) SetNillableName(v *string) *UserUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetCardID sets the "card" edge to the Card entity by ID.
func (m *UserUpdate) SetCardID(id int) *UserUpdate {
	m.mutation.SetCardID(id)
	return m
}

// SetNillableCardID sets the "card" edge to the Card entity by ID if the given value is not nil.
func (m *UserUpdate) SetNillableCardID(id *int) *UserUpdate {
	if id != nil {
		m = m.SetCardID(*id)
	}
	return m
}

// SetCard sets the "card" edge to the Card entity.
func (m *UserUpdate) SetCard(v *Card) *UserUpdate {
	return m.SetCardID(v.ID)
}

// Mutation returns the UserMutation object of the builder.
func (m *UserUpdate) Mutation() *UserMutation {
	return m.mutation
}

// ClearCard clears the "card" edge to the Card entity.
func (u *UserUpdate) ClearCard() *UserUpdate {
	u.mutation.ClearCard()
	return u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *UserUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *UserUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if u.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (m *UserUpdateOne) SetAge(v int) *UserUpdateOne {
	m.mutation.ResetAge()
	m.mutation.SetAge(v)
	return m
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (m *UserUpdateOne) SetNillableAge(v *int) *UserUpdateOne {
	if v != nil {
		m.SetAge(*v)
	}
	return m
}

// AddAge adds value to the "age" field.
func (m *UserUpdateOne) AddAge(v int) *UserUpdateOne {
	m.mutation.AddAge(v)
	return m
}

// SetName sets the "name" field.
func (m *UserUpdateOne) SetName(v string) *UserUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *UserUpdateOne) SetNillableName(v *string) *UserUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetCardID sets the "card" edge to the Card entity by ID.
func (m *UserUpdateOne) SetCardID(id int) *UserUpdateOne {
	m.mutation.SetCardID(id)
	return m
}

// SetNillableCardID sets the "card" edge to the Card entity by ID if the given value is not nil.
func (m *UserUpdateOne) SetNillableCardID(id *int) *UserUpdateOne {
	if id != nil {
		m = m.SetCardID(*id)
	}
	return m
}

// SetCard sets the "card" edge to the Card entity.
func (m *UserUpdateOne) SetCard(v *Card) *UserUpdateOne {
	return m.SetCardID(v.ID)
}

// Mutation returns the UserMutation object of the builder.
func (m *UserUpdateOne) Mutation() *UserMutation {
	return m.mutation
}

// ClearCard clears the "card" edge to the Card entity.
func (u *UserUpdateOne) ClearCard() *UserUpdateOne {
	u.mutation.ClearCard()
	return u
}

// Where appends a list predicates to the UserUpdate builder.
func (u *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated User entity.
func (u *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *UserUpdateOne) sqlSave(ctx context.Context) (_n *User, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
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
	if value, ok := u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeInt, value)
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if u.mutation.CardCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.CardIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.CardTable,
			Columns: []string{user.CardColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &User{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
