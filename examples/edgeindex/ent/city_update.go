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
	"entgo.io/ent/examples/edgeindex/ent/city"
	"entgo.io/ent/examples/edgeindex/ent/predicate"
	"entgo.io/ent/examples/edgeindex/ent/street"
	"entgo.io/ent/schema/field"
)

// CityUpdate is the builder for updating City entities.
type CityUpdate struct {
	config
	hooks    []Hook
	mutation *CityMutation
}

// Where appends a list predicates to the CityUpdate builder.
func (u *CityUpdate) Where(ps ...predicate.City) *CityUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetName sets the "name" field.
func (m *CityUpdate) SetName(v string) *CityUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CityUpdate) SetNillableName(v *string) *CityUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (m *CityUpdate) AddStreetIDs(ids ...int) *CityUpdate {
	m.mutation.AddStreetIDs(ids...)
	return m
}

// AddStreets adds the "streets" edges to the Street entity.
func (m *CityUpdate) AddStreets(v ...*Street) *CityUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (m *CityUpdate) Mutation() *CityMutation {
	return m.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (u *CityUpdate) ClearStreets() *CityUpdate {
	u.mutation.ClearStreets()
	return u
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (u *CityUpdate) RemoveStreetIDs(ids ...int) *CityUpdate {
	u.mutation.RemoveStreetIDs(ids...)
	return u
}

// RemoveStreets removes "streets" edges to Street entities.
func (u *CityUpdate) RemoveStreets(v ...*Street) *CityUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *CityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CityUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *CityUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CityUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CityUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(city.Table, city.Columns, sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(city.FieldName, field.TypeString, value)
	}
	if u.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !u.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// CityUpdateOne is the builder for updating a single City entity.
type CityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CityMutation
}

// SetName sets the "name" field.
func (m *CityUpdateOne) SetName(v string) *CityUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *CityUpdateOne) SetNillableName(v *string) *CityUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// AddStreetIDs adds the "streets" edge to the Street entity by IDs.
func (m *CityUpdateOne) AddStreetIDs(ids ...int) *CityUpdateOne {
	m.mutation.AddStreetIDs(ids...)
	return m
}

// AddStreets adds the "streets" edges to the Street entity.
func (m *CityUpdateOne) AddStreets(v ...*Street) *CityUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (m *CityUpdateOne) Mutation() *CityMutation {
	return m.mutation
}

// ClearStreets clears all "streets" edges to the Street entity.
func (u *CityUpdateOne) ClearStreets() *CityUpdateOne {
	u.mutation.ClearStreets()
	return u
}

// RemoveStreetIDs removes the "streets" edge to Street entities by IDs.
func (u *CityUpdateOne) RemoveStreetIDs(ids ...int) *CityUpdateOne {
	u.mutation.RemoveStreetIDs(ids...)
	return u
}

// RemoveStreets removes "streets" edges to Street entities.
func (u *CityUpdateOne) RemoveStreets(v ...*Street) *CityUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveStreetIDs(ids...)
}

// Where appends a list predicates to the CityUpdate builder.
func (u *CityUpdateOne) Where(ps ...predicate.City) *CityUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *CityUpdateOne) Select(field string, fields ...string) *CityUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated City entity.
func (u *CityUpdateOne) Save(ctx context.Context) (*City, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CityUpdateOne) SaveX(ctx context.Context) *City {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *CityUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CityUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *CityUpdateOne) sqlSave(ctx context.Context) (_n *City, err error) {
	_spec := sqlgraph.NewUpdateSpec(city.Table, city.Columns, sqlgraph.NewFieldSpec(city.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "City.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, city.FieldID)
		for _, f := range fields {
			if !city.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != city.FieldID {
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
		_spec.SetField(city.FieldName, field.TypeString, value)
	}
	if u.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !u.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(street.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &City{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
