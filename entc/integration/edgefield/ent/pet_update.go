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
	"entgo.io/ent/entc/integration/edgefield/ent/pet"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
	"entgo.io/ent/schema/field"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks    []Hook
	mutation *PetMutation
}

// Where appends a list predicates to the PetUpdate builder.
func (_u *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetOwnerID sets the "owner_id" field.
func (_u *PetUpdate) SetOwnerID(v int) *PetUpdate {
	_u.mutation.SetOwnerID(v)
	return _u
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (_u *PetUpdate) SetNillableOwnerID(v *int) *PetUpdate {
	if v != nil {
		_u.SetOwnerID(*v)
	}
	return _u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (_u *PetUpdate) ClearOwnerID() *PetUpdate {
	_u.mutation.ClearOwnerID()
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *PetUpdate) SetOwner(v *User) *PetUpdate {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the PetMutation object of the builder.
func (_u *PetUpdate) Mutation() *PetMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *PetUpdate) ClearOwner() *PetUpdate {
	_u.mutation.ClearOwner()
	return _u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *PetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *PetUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *PetUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *PetUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
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
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PetMutation
}

// SetOwnerID sets the "owner_id" field.
func (_u *PetUpdateOne) SetOwnerID(v int) *PetUpdateOne {
	_u.mutation.SetOwnerID(v)
	return _u
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (_u *PetUpdateOne) SetNillableOwnerID(v *int) *PetUpdateOne {
	if v != nil {
		_u.SetOwnerID(*v)
	}
	return _u
}

// ClearOwnerID clears the value of the "owner_id" field.
func (_u *PetUpdateOne) ClearOwnerID() *PetUpdateOne {
	_u.mutation.ClearOwnerID()
	return _u
}

// SetOwner sets the "owner" edge to the User entity.
func (_u *PetUpdateOne) SetOwner(v *User) *PetUpdateOne {
	return _u.SetOwnerID(v.ID)
}

// Mutation returns the PetMutation object of the builder.
func (_u *PetUpdateOne) Mutation() *PetMutation {
	return _u.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (_u *PetUpdateOne) ClearOwner() *PetUpdateOne {
	_u.mutation.ClearOwner()
	return _u
}

// Where appends a list predicates to the PetUpdate builder.
func (_u *PetUpdateOne) Where(ps ...predicate.Pet) *PetUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Pet entity.
func (_u *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *PetUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *PetUpdateOne) sqlSave(ctx context.Context) (_node *Pet, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pet.FieldID)
		for _, f := range fields {
			if !pet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pet.FieldID {
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
	if _u.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   pet.OwnerTable,
			Columns: []string{pet.OwnerColumn},
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
	_node = &Pet{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
