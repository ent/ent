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
	"entgo.io/ent/entc/integration/multischema/ent/internal"
	"entgo.io/ent/entc/integration/multischema/ent/pet"
	"entgo.io/ent/entc/integration/multischema/ent/predicate"
	"entgo.io/ent/entc/integration/multischema/ent/user"
	"entgo.io/ent/schema/field"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	hooks     []Hook
	mutation  *PetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the PetUpdate builder.
func (u *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetName sets the "name" field.
func (m *PetUpdate) SetName(v string) *PetUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *PetUpdate) SetNillableName(v *string) *PetUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetOwnerID sets the "owner_id" field.
func (m *PetUpdate) SetOwnerID(v int) *PetUpdate {
	m.mutation.SetOwnerID(v)
	return m
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (m *PetUpdate) SetNillableOwnerID(v *int) *PetUpdate {
	if v != nil {
		m.SetOwnerID(*v)
	}
	return m
}

// ClearOwnerID clears the value of the "owner_id" field.
func (m *PetUpdate) ClearOwnerID() *PetUpdate {
	m.mutation.ClearOwnerID()
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *PetUpdate) SetOwner(v *User) *PetUpdate {
	return m.SetOwnerID(v.ID)
}

// Mutation returns the PetMutation object of the builder.
func (m *PetUpdate) Mutation() *PetMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *PetUpdate) ClearOwner() *PetUpdate {
	u.mutation.ClearOwner()
	return u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *PetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *PetUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *PetUpdate) Modify(modifiers ...func(*sql.UpdateBuilder)) *PetUpdate {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *PetUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
	}
	if u.mutation.OwnerCleared() {
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
		edge.Schema = u.schemaConfig.Pet
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
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
		edge.Schema = u.schemaConfig.Pet
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = u.schemaConfig.Pet
	ctx = internal.NewSchemaConfigContext(ctx, u.schemaConfig)
	_spec.AddModifiers(u.modifiers...)
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *PetMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (m *PetUpdateOne) SetName(v string) *PetUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *PetUpdateOne) SetNillableName(v *string) *PetUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetOwnerID sets the "owner_id" field.
func (m *PetUpdateOne) SetOwnerID(v int) *PetUpdateOne {
	m.mutation.SetOwnerID(v)
	return m
}

// SetNillableOwnerID sets the "owner_id" field if the given value is not nil.
func (m *PetUpdateOne) SetNillableOwnerID(v *int) *PetUpdateOne {
	if v != nil {
		m.SetOwnerID(*v)
	}
	return m
}

// ClearOwnerID clears the value of the "owner_id" field.
func (m *PetUpdateOne) ClearOwnerID() *PetUpdateOne {
	m.mutation.ClearOwnerID()
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *PetUpdateOne) SetOwner(v *User) *PetUpdateOne {
	return m.SetOwnerID(v.ID)
}

// Mutation returns the PetMutation object of the builder.
func (m *PetUpdateOne) Mutation() *PetMutation {
	return m.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (u *PetUpdateOne) ClearOwner() *PetUpdateOne {
	u.mutation.ClearOwner()
	return u
}

// Where appends a list predicates to the PetUpdate builder.
func (u *PetUpdateOne) Where(ps ...predicate.Pet) *PetUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *PetUpdateOne) Select(field string, fields ...string) *PetUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Pet entity.
func (u *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PetUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *PetUpdateOne) Modify(modifiers ...func(*sql.UpdateBuilder)) *PetUpdateOne {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *PetUpdateOne) sqlSave(ctx context.Context) (_n *Pet, err error) {
	_spec := sqlgraph.NewUpdateSpec(pet.Table, pet.Columns, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
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
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(pet.FieldName, field.TypeString, value)
	}
	if u.mutation.OwnerCleared() {
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
		edge.Schema = u.schemaConfig.Pet
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.OwnerIDs(); len(nodes) > 0 {
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
		edge.Schema = u.schemaConfig.Pet
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = u.schemaConfig.Pet
	ctx = internal.NewSchemaConfigContext(ctx, u.schemaConfig)
	_spec.AddModifiers(u.modifiers...)
	_n = &Pet{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
