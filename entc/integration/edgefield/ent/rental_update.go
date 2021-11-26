// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"

	"entgo.io/ent/entc/integration/edgefield/ent/car"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/rental"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
)

// RentalUpdate is the builder for updating Rental entities.
type RentalUpdate struct {
	config
	hooks    []Hook
	mutation *RentalMutation
}

// Where appends a list predicates to the RentalUpdate builder.
func (ru *RentalUpdate) Where(ps ...predicate.Rental) *RentalUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetDate sets the "date" field.
func (ru *RentalUpdate) SetDate(t time.Time) *RentalUpdate {
	ru.mutation.SetDate(t)
	return ru
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ru *RentalUpdate) SetNillableDate(t *time.Time) *RentalUpdate {
	if t != nil {
		ru.SetDate(*t)
	}
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RentalUpdate) SetUserID(i int) *RentalUpdate {
	ru.mutation.SetUserID(i)
	return ru
}

// SetCarID sets the "car_id" field.
func (ru *RentalUpdate) SetCarID(u uuid.UUID) *RentalUpdate {
	ru.mutation.SetCarID(u)
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *RentalUpdate) SetUser(u *User) *RentalUpdate {
	return ru.SetUserID(u.ID)
}

// SetCar sets the "car" edge to the Car entity.
func (ru *RentalUpdate) SetCar(c *Car) *RentalUpdate {
	return ru.SetCarID(c.ID)
}

// Mutation returns the RentalMutation object of the builder.
func (ru *RentalUpdate) Mutation() *RentalMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *RentalUpdate) ClearUser() *RentalUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearCar clears the "car" edge to the Car entity.
func (ru *RentalUpdate) ClearCar() *RentalUpdate {
	ru.mutation.ClearCar()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RentalUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ru.hooks) == 0 {
		if err = ru.check(); err != nil {
			return 0, err
		}
		affected, err = ru.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ru.check(); err != nil {
				return 0, err
			}
			ru.mutation = mutation
			affected, err = ru.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ru.hooks) - 1; i >= 0; i-- {
			if ru.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ru.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ru.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RentalUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RentalUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RentalUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RentalUpdate) check() error {
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Rental.user"`)
	}
	if _, ok := ru.mutation.CarID(); ru.mutation.CarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Rental.car"`)
	}
	return nil
}

func (ru *RentalUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rental.Table,
			Columns: rental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rental.FieldID,
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rental.FieldDate,
		})
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// RentalUpdateOne is the builder for updating a single Rental entity.
type RentalUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RentalMutation
}

// SetDate sets the "date" field.
func (ruo *RentalUpdateOne) SetDate(t time.Time) *RentalUpdateOne {
	ruo.mutation.SetDate(t)
	return ruo
}

// SetNillableDate sets the "date" field if the given value is not nil.
func (ruo *RentalUpdateOne) SetNillableDate(t *time.Time) *RentalUpdateOne {
	if t != nil {
		ruo.SetDate(*t)
	}
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RentalUpdateOne) SetUserID(i int) *RentalUpdateOne {
	ruo.mutation.SetUserID(i)
	return ruo
}

// SetCarID sets the "car_id" field.
func (ruo *RentalUpdateOne) SetCarID(u uuid.UUID) *RentalUpdateOne {
	ruo.mutation.SetCarID(u)
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *RentalUpdateOne) SetUser(u *User) *RentalUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetCar sets the "car" edge to the Car entity.
func (ruo *RentalUpdateOne) SetCar(c *Car) *RentalUpdateOne {
	return ruo.SetCarID(c.ID)
}

// Mutation returns the RentalMutation object of the builder.
func (ruo *RentalUpdateOne) Mutation() *RentalMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *RentalUpdateOne) ClearUser() *RentalUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearCar clears the "car" edge to the Car entity.
func (ruo *RentalUpdateOne) ClearCar() *RentalUpdateOne {
	ruo.mutation.ClearCar()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RentalUpdateOne) Select(field string, fields ...string) *RentalUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Rental entity.
func (ruo *RentalUpdateOne) Save(ctx context.Context) (*Rental, error) {
	var (
		err  error
		node *Rental
	)
	if len(ruo.hooks) == 0 {
		if err = ruo.check(); err != nil {
			return nil, err
		}
		node, err = ruo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RentalMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruo.check(); err != nil {
				return nil, err
			}
			ruo.mutation = mutation
			node, err = ruo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruo.hooks) - 1; i >= 0; i-- {
			if ruo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RentalUpdateOne) SaveX(ctx context.Context) *Rental {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RentalUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RentalUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RentalUpdateOne) check() error {
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Rental.user"`)
	}
	if _, ok := ruo.mutation.CarID(); ruo.mutation.CarCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Rental.car"`)
	}
	return nil
}

func (ruo *RentalUpdateOne) sqlSave(ctx context.Context) (_node *Rental, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   rental.Table,
			Columns: rental.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: rental.FieldID,
			},
		},
	}
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Rental.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, rental.FieldID)
		for _, f := range fields {
			if !rental.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != rental.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Date(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: rental.FieldDate,
		})
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.UserTable,
			Columns: []string{rental.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.CarCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: car.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.CarIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   rental.CarTable,
			Columns: []string{rental.CarColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: car.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Rental{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{rental.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
