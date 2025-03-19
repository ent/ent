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
	"entgo.io/ent/entc/integration/edgefield/ent/card"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	hooks    []Hook
	mutation *CardMutation
}

// Where appends a list predicates to the CardUpdate builder.
func (_u *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetNumber sets the "number" field.
func (_u *CardUpdate) SetNumber(v string) *CardUpdate {
	_u.mutation.SetNumber(v)
	return _u
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (_u *CardUpdate) SetNillableNumber(v *string) *CardUpdate {
	if v != nil {
		_u.SetNumber(*v)
	}
	return _u
}

// ClearNumber clears the value of the "number" field.
func (_u *CardUpdate) ClearNumber() *CardUpdate {
	_u.mutation.ClearNumber()
	return _u
}

// SetOwnerID sets the "owner_id" field.
func (_u *CardUpdate) SetOwnerID(v int) *CardUpdate {
	_u.mutation.SetOwnerID(v)
	return _u
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (_u *CardUpdate) SetNillableOwnerID(v *int) *CardUpdate {
	if v != nil {
		_u.SetOwnerID(*v)
	}
	return _u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (_u *CardUpdate) ClearOwnerID() *CardUpdate {
	_u.mutation.ClearOwnerID()
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *CardUpdate) SetOwner(v *User) *CardUpdate {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the CardMutation object of the builder.
func (_u *CardUpdate) Mutation() *CardMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *CardUpdate) ClearOwner() *CardUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *CardUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *CardUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CardUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *CardUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Number(); ok {
		_spec.SetField(card.FieldNumber, field.TypeString, value)
	}
	if _u.mutation.NumberCleared() {
		_spec.ClearField(card.FieldNumber, field.TypeString)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
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
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
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
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CardMutation
}

// SetNumber sets the "number" field.
func (_u *CardUpdateOne) SetNumber(v string) *CardUpdateOne {
	_u.mutation.SetNumber(v)
	return _u
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (_u *CardUpdateOne) SetNillableNumber(v *string) *CardUpdateOne {
	if v != nil {
		_u.SetNumber(*v)
	}
	return _u
}

// ClearNumber clears the value of the "number" field.
func (_u *CardUpdateOne) ClearNumber() *CardUpdateOne {
	_u.mutation.ClearNumber()
	return _u
}

// SetOwnerID sets the "owner_id" field.
func (_u *CardUpdateOne) SetOwnerID(v int) *CardUpdateOne {
	_u.mutation.SetOwnerID(v)
	return _u
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (_u *CardUpdateOne) SetNillableOwnerID(v *int) *CardUpdateOne {
	if v != nil {
		_u.SetOwnerID(*v)
	}
	return _u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (_u *CardUpdateOne) ClearOwnerID() *CardUpdateOne {
	_u.mutation.ClearOwnerID()
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *CardUpdateOne) SetOwner(v *User) *CardUpdateOne {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the CardMutation object of the builder.
func (_u *CardUpdateOne) Mutation() *CardMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *CardUpdateOne) ClearOwner() *CardUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// Where appends a list predicates to the CardUpdate builder.
func (_u *CardUpdateOne) Where(ps ...predicate.Card) *CardUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *CardUpdateOne) Select(field string, fields ...string) *CardUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Card entity.
func (_u *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CardUpdateOne) SaveX(ctx context.Context) *Card {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CardUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *CardUpdateOne) sqlSave(ctx context.Context) (_node *Card, err error) {
	_spec := sqlgraph.NewUpdateSpec(card.Table, card.Columns, sqlgraph.NewFieldSpec(card.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Card.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
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
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Number(); ok {
		_spec.SetField(card.FieldNumber, field.TypeString, value)
	}
	if _u.mutation.NumberCleared() {
		_spec.ClearField(card.FieldNumber, field.TypeString)
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
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
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
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
	_node = &Card{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{card.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
