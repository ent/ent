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
	"entgo.io/ent/entc/integration/customid/ent/intsid"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/sid"
	"entgo.io/ent/schema/field"
)

// IntSIDUpdate is the builder for updating IntSID entities.
type IntSIDUpdate struct {
	config
	hooks    []Hook
	mutation *IntSIDMutation
}

// Where appends a list predicates to the IntSIDUpdate builder.
func (isu *IntSIDUpdate) Where(ps ...predicate.IntSID) *IntSIDUpdate {
	isu.mutation.Where(ps...)
	return isu
}

// SetParentID sets the "parent" edge to the IntSID entity by ID.
func (isu *IntSIDUpdate) SetParentID(id sid.ID) *IntSIDUpdate {
	isu.mutation.SetParentID(id)
	return isu
}

// SetNillableParentID sets the "parent" edge to the IntSID entity by ID if the given value is not nil.
func (isu *IntSIDUpdate) SetNillableParentID(id *sid.ID) *IntSIDUpdate {
	if id != nil {
		isu = isu.SetParentID(*id)
	}
	return isu
}

// SetParent sets the "parent" edge to the IntSID entity.
func (isu *IntSIDUpdate) SetParent(i *IntSID) *IntSIDUpdate {
	return isu.SetParentID(i.ID)
}

// AddChildIDs adds the "children" edge to the IntSID entity by IDs.
func (isu *IntSIDUpdate) AddChildIDs(ids ...sid.ID) *IntSIDUpdate {
	isu.mutation.AddChildIDs(ids...)
	return isu
}

// AddChildren adds the "children" edges to the IntSID entity.
func (isu *IntSIDUpdate) AddChildren(i ...*IntSID) *IntSIDUpdate {
	ids := make([]sid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return isu.AddChildIDs(ids...)
}

// Mutation returns the IntSIDMutation object of the builder.
func (isu *IntSIDUpdate) Mutation() *IntSIDMutation {
	return isu.mutation
}

// ClearParent clears the "parent" edge to the IntSID entity.
func (isu *IntSIDUpdate) ClearParent() *IntSIDUpdate {
	isu.mutation.ClearParent()
	return isu
}

// ClearChildren clears all "children" edges to the IntSID entity.
func (isu *IntSIDUpdate) ClearChildren() *IntSIDUpdate {
	isu.mutation.ClearChildren()
	return isu
}

// RemoveChildIDs removes the "children" edge to IntSID entities by IDs.
func (isu *IntSIDUpdate) RemoveChildIDs(ids ...sid.ID) *IntSIDUpdate {
	isu.mutation.RemoveChildIDs(ids...)
	return isu
}

// RemoveChildren removes "children" edges to IntSID entities.
func (isu *IntSIDUpdate) RemoveChildren(i ...*IntSID) *IntSIDUpdate {
	ids := make([]sid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return isu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (isu *IntSIDUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(isu.hooks) == 0 {
		affected, err = isu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IntSIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			isu.mutation = mutation
			affected, err = isu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(isu.hooks) - 1; i >= 0; i-- {
			if isu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = isu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, isu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (isu *IntSIDUpdate) SaveX(ctx context.Context) int {
	affected, err := isu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (isu *IntSIDUpdate) Exec(ctx context.Context) error {
	_, err := isu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isu *IntSIDUpdate) ExecX(ctx context.Context) {
	if err := isu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (isu *IntSIDUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   intsid.Table,
			Columns: intsid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: intsid.FieldID,
			},
		},
	}
	if ps := isu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if isu.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   intsid.ParentTable,
			Columns: []string{intsid.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isu.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   intsid.ParentTable,
			Columns: []string{intsid.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if isu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isu.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !isu.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isu.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, isu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{intsid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// IntSIDUpdateOne is the builder for updating a single IntSID entity.
type IntSIDUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *IntSIDMutation
}

// SetParentID sets the "parent" edge to the IntSID entity by ID.
func (isuo *IntSIDUpdateOne) SetParentID(id sid.ID) *IntSIDUpdateOne {
	isuo.mutation.SetParentID(id)
	return isuo
}

// SetNillableParentID sets the "parent" edge to the IntSID entity by ID if the given value is not nil.
func (isuo *IntSIDUpdateOne) SetNillableParentID(id *sid.ID) *IntSIDUpdateOne {
	if id != nil {
		isuo = isuo.SetParentID(*id)
	}
	return isuo
}

// SetParent sets the "parent" edge to the IntSID entity.
func (isuo *IntSIDUpdateOne) SetParent(i *IntSID) *IntSIDUpdateOne {
	return isuo.SetParentID(i.ID)
}

// AddChildIDs adds the "children" edge to the IntSID entity by IDs.
func (isuo *IntSIDUpdateOne) AddChildIDs(ids ...sid.ID) *IntSIDUpdateOne {
	isuo.mutation.AddChildIDs(ids...)
	return isuo
}

// AddChildren adds the "children" edges to the IntSID entity.
func (isuo *IntSIDUpdateOne) AddChildren(i ...*IntSID) *IntSIDUpdateOne {
	ids := make([]sid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return isuo.AddChildIDs(ids...)
}

// Mutation returns the IntSIDMutation object of the builder.
func (isuo *IntSIDUpdateOne) Mutation() *IntSIDMutation {
	return isuo.mutation
}

// ClearParent clears the "parent" edge to the IntSID entity.
func (isuo *IntSIDUpdateOne) ClearParent() *IntSIDUpdateOne {
	isuo.mutation.ClearParent()
	return isuo
}

// ClearChildren clears all "children" edges to the IntSID entity.
func (isuo *IntSIDUpdateOne) ClearChildren() *IntSIDUpdateOne {
	isuo.mutation.ClearChildren()
	return isuo
}

// RemoveChildIDs removes the "children" edge to IntSID entities by IDs.
func (isuo *IntSIDUpdateOne) RemoveChildIDs(ids ...sid.ID) *IntSIDUpdateOne {
	isuo.mutation.RemoveChildIDs(ids...)
	return isuo
}

// RemoveChildren removes "children" edges to IntSID entities.
func (isuo *IntSIDUpdateOne) RemoveChildren(i ...*IntSID) *IntSIDUpdateOne {
	ids := make([]sid.ID, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return isuo.RemoveChildIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (isuo *IntSIDUpdateOne) Select(field string, fields ...string) *IntSIDUpdateOne {
	isuo.fields = append([]string{field}, fields...)
	return isuo
}

// Save executes the query and returns the updated IntSID entity.
func (isuo *IntSIDUpdateOne) Save(ctx context.Context) (*IntSID, error) {
	var (
		err  error
		node *IntSID
	)
	if len(isuo.hooks) == 0 {
		node, err = isuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*IntSIDMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			isuo.mutation = mutation
			node, err = isuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(isuo.hooks) - 1; i >= 0; i-- {
			if isuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = isuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, isuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*IntSID)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from IntSIDMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (isuo *IntSIDUpdateOne) SaveX(ctx context.Context) *IntSID {
	node, err := isuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (isuo *IntSIDUpdateOne) Exec(ctx context.Context) error {
	_, err := isuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (isuo *IntSIDUpdateOne) ExecX(ctx context.Context) {
	if err := isuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (isuo *IntSIDUpdateOne) sqlSave(ctx context.Context) (_node *IntSID, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   intsid.Table,
			Columns: intsid.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: intsid.FieldID,
			},
		},
	}
	id, ok := isuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "IntSID.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := isuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, intsid.FieldID)
		for _, f := range fields {
			if !intsid.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != intsid.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := isuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if isuo.mutation.ParentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   intsid.ParentTable,
			Columns: []string{intsid.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isuo.mutation.ParentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   intsid.ParentTable,
			Columns: []string{intsid.ParentColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if isuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isuo.mutation.RemovedChildrenIDs(); len(nodes) > 0 && !isuo.mutation.ChildrenCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := isuo.mutation.ChildrenIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   intsid.ChildrenTable,
			Columns: []string{intsid.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: intsid.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &IntSID{config: isuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, isuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{intsid.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
