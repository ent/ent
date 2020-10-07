// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/examples/edgeindex/ent/city"
	"github.com/facebook/ent/examples/edgeindex/ent/predicate"
	"github.com/facebook/ent/examples/edgeindex/ent/street"
	"github.com/facebook/ent/schema/field"
)

// CityUpdate is the builder for updating City entities.
type CityUpdate struct {
	config
	hooks    []Hook
	mutation *CityMutation
}

// Where adds a new predicate for the builder.
func (cu *CityUpdate) Where(ps ...predicate.City) *CityUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CityUpdate) SetName(s string) *CityUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cu *CityUpdate) AddStreetIDs(ids ...int) *CityUpdate {
	cu.mutation.AddStreetIDs(ids...)
	return cu
}

// AddStreets adds the streets edges to Street.
func (cu *CityUpdate) AddStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cu *CityUpdate) Mutation() *CityMutation {
	return cu.mutation
}

// ClearStreets clears all "streets" edges to type Street.
func (cu *CityUpdate) ClearStreets() *CityUpdate {
	cu.mutation.ClearStreets()
	return cu
}

// RemoveStreetIDs removes the streets edge to Street by ids.
func (cu *CityUpdate) RemoveStreetIDs(ids ...int) *CityUpdate {
	cu.mutation.RemoveStreetIDs(ids...)
	return cu
}

// RemoveStreets removes streets edges to Street.
func (cu *CityUpdate) RemoveStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   city.Table,
			Columns: city.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
	}
	if cu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !cu.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CityUpdateOne is the builder for updating a single City entity.
type CityUpdateOne struct {
	config
	hooks    []Hook
	mutation *CityMutation
}

// SetName sets the name field.
func (cuo *CityUpdateOne) SetName(s string) *CityUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cuo *CityUpdateOne) AddStreetIDs(ids ...int) *CityUpdateOne {
	cuo.mutation.AddStreetIDs(ids...)
	return cuo
}

// AddStreets adds the streets edges to Street.
func (cuo *CityUpdateOne) AddStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStreetIDs(ids...)
}

// Mutation returns the CityMutation object of the builder.
func (cuo *CityUpdateOne) Mutation() *CityMutation {
	return cuo.mutation
}

// ClearStreets clears all "streets" edges to type Street.
func (cuo *CityUpdateOne) ClearStreets() *CityUpdateOne {
	cuo.mutation.ClearStreets()
	return cuo
}

// RemoveStreetIDs removes the streets edge to Street by ids.
func (cuo *CityUpdateOne) RemoveStreetIDs(ids ...int) *CityUpdateOne {
	cuo.mutation.RemoveStreetIDs(ids...)
	return cuo
}

// RemoveStreets removes streets edges to Street.
func (cuo *CityUpdateOne) RemoveStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CityUpdateOne) Save(ctx context.Context) (*City, error) {
	var (
		err  error
		node *City
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CityUpdateOne) SaveX(ctx context.Context) *City {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CityUpdateOne) sqlSave(ctx context.Context) (_node *City, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   city.Table,
			Columns: city.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: city.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing City.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: city.FieldName,
		})
	}
	if cuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStreetsIDs(); len(nodes) > 0 && !cuo.mutation.StreetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StreetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   city.StreetsTable,
			Columns: []string{city.StreetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: street.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &City{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{city.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
