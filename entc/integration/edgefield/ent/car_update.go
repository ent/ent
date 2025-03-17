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
	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/schema/field"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	hooks    []Hook
	mutation *CarMutation
}

// Where appends a list predicates to the CarUpdate builder.
func (u *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetNumber sets the "number" field.
func (m *CarUpdate) SetNumber(v string) *CarUpdate {
	m.mutation.SetNumber(v)
	return m
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (m *CarUpdate) SetNillableNumber(v *string) *CarUpdate {
	if v != nil {
		m.SetNumber(*v)
	}
	return m
}

// ClearNumber clears the value of the "number" field.
func (m *CarUpdate) ClearNumber() *CarUpdate {
	m.mutation.ClearNumber()
	return m
}

// AddRentalIDs adds the "rentals" edge to the Rental entity by IDs.
func (m *CarUpdate) AddRentalIDs(ids ...int) *CarUpdate {
	m.mutation.AddRentalIDs(ids...)
	return m
}

// AddRentals adds the "rentals" edges to the Rental entity.
func (m *CarUpdate) AddRentals(v ...*Rental) *CarUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddRentalIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (m *CarUpdate) Mutation() *CarMutation {
	return m.mutation
}

// ClearRentals clears all "rentals" edges to the Rental entity.
func (u *CarUpdate) ClearRentals() *CarUpdate {
	u.mutation.ClearRentals()
	return u
}

// RemoveRentalIDs removes the "rentals" edge to Rental entities by IDs.
func (u *CarUpdate) RemoveRentalIDs(ids ...int) *CarUpdate {
	u.mutation.RemoveRentalIDs(ids...)
	return u
}

// RemoveRentals removes "rentals" edges to Rental entities.
func (u *CarUpdate) RemoveRentals(v ...*Rental) *CarUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveRentalIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *CarUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *CarUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CarUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Number(); ok {
		_spec.SetField(car.FieldNumber, field.TypeString, value)
	}
	if u.mutation.NumberCleared() {
		_spec.ClearField(car.FieldNumber, field.TypeString)
	}
	if u.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedRentalsIDs(); len(nodes) > 0 && !u.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RentalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CarMutation
}

// SetNumber sets the "number" field.
func (m *CarUpdateOne) SetNumber(v string) *CarUpdateOne {
	m.mutation.SetNumber(v)
	return m
}

// SetNillableNumber sets the "number" field if the given value is not nil.
func (m *CarUpdateOne) SetNillableNumber(v *string) *CarUpdateOne {
	if v != nil {
		m.SetNumber(*v)
	}
	return m
}

// ClearNumber clears the value of the "number" field.
func (m *CarUpdateOne) ClearNumber() *CarUpdateOne {
	m.mutation.ClearNumber()
	return m
}

// AddRentalIDs adds the "rentals" edge to the Rental entity by IDs.
func (m *CarUpdateOne) AddRentalIDs(ids ...int) *CarUpdateOne {
	m.mutation.AddRentalIDs(ids...)
	return m
}

// AddRentals adds the "rentals" edges to the Rental entity.
func (m *CarUpdateOne) AddRentals(v ...*Rental) *CarUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddRentalIDs(ids...)
}

// Mutation returns the CarMutation object of the builder.
func (m *CarUpdateOne) Mutation() *CarMutation {
	return m.mutation
}

// ClearRentals clears all "rentals" edges to the Rental entity.
func (u *CarUpdateOne) ClearRentals() *CarUpdateOne {
	u.mutation.ClearRentals()
	return u
}

// RemoveRentalIDs removes the "rentals" edge to Rental entities by IDs.
func (u *CarUpdateOne) RemoveRentalIDs(ids ...int) *CarUpdateOne {
	u.mutation.RemoveRentalIDs(ids...)
	return u
}

// RemoveRentals removes "rentals" edges to Rental entities.
func (u *CarUpdateOne) RemoveRentals(v ...*Rental) *CarUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveRentalIDs(ids...)
}

// Where appends a list predicates to the CarUpdate builder.
func (u *CarUpdateOne) Where(ps ...predicate.Car) *CarUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *CarUpdateOne) Select(field string, fields ...string) *CarUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Car entity.
func (u *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CarUpdateOne) SaveX(ctx context.Context) *Car {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CarUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CarUpdateOne) sqlSave(ctx context.Context) (_n *Car, err error) {
	_spec := sqlgraph.NewUpdateSpec(car.Table, car.Columns, sqlgraph.NewFieldSpec(car.FieldID, field.TypeUUID))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Car.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, car.FieldID)
		for _, f := range fields {
			if !car.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != car.FieldID {
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
	if value, ok := u.mutation.Number(); ok {
		_spec.SetField(car.FieldNumber, field.TypeString, value)
	}
	if u.mutation.NumberCleared() {
		_spec.ClearField(car.FieldNumber, field.TypeString)
	}
	if u.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedRentalsIDs(); len(nodes) > 0 && !u.mutation.RentalsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RentalsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   car.RentalsTable,
			Columns: []string{car.RentalsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(rental.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &Car{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{car.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
