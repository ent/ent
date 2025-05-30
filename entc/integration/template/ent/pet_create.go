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

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/template/ent/pet"
	"entgo.io/ent/entc/integration/template/ent/user"
	"entgo.io/ent/schema/field"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	mutation *PetMutation
	hooks    []Hook
}

// SetAge sets the "age" field.
func (_c *PetCreate) SetAge(v int) *PetCreate {
	_c.mutation.SetAge(v)
	return _c
}

// SetLicensedAt sets the "licensed_at" field.
func (_c *PetCreate) SetLicensedAt(v time.Time) *PetCreate {
	_c.mutation.SetLicensedAt(v)
	return _c
}

// SetNillableLicensedAt sets the "licensed_at" field if the given value is not nil.
func (_c *PetCreate) SetNillableLicensedAt(v *time.Time) *PetCreate {
	if v != nil {
		_c.SetLicensedAt(*v)
	}
	return _c
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_c *PetCreate) SetOwnerID(id int) *PetCreate {
	_c.mutation.SetOwnerID(id)
	return _c
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_c *PetCreate) SetNillableOwnerID(id *int) *PetCreate {
	if id != nil {
		_c = _c.SetOwnerID(*id)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *PetCreate) SetOwner(v *User) *PetCreate {
	return _c.SetOwnerID(v.ID)
}

// Mutation returns the PetMutation object of the builder.
func (_c *PetCreate) Mutation() *PetMutation {
	return _c.mutation
}

// Save creates the Pet in the database.
func (_c *PetCreate) Save(ctx context.Context) (*Pet, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *PetCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *PetCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *PetCreate) check() error {
	if _, ok := _c.mutation.Age(); !ok {
		return &ValidationError{Name: "age", err: errors.New(`ent: missing required field "Pet.age"`)}
	}
	return nil
}

func (_c *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *PetCreate) createSpec() (*Pet, *sqlgraph.CreateSpec) {
	var (
		_node = &Pet{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(pet.Table, sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Age(); ok {
		_spec.SetField(pet.FieldAge, field.TypeInt, value)
		_node.Age = value
	}
	if value, ok := _c.mutation.LicensedAt(); ok {
		_spec.SetField(pet.FieldLicensedAt, field.TypeTime, value)
		_node.LicensedAt = &value
	}
	if nodes := _c.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_node.user_pets = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PetCreateBulk is the builder for creating many Pet entities in bulk.
type PetCreateBulk struct {
	config
	err      error
	builders []*PetCreate
}

// Save creates the Pet entities in the database.
func (_c *PetCreateBulk) Save(ctx context.Context) ([]*Pet, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Pet, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PetMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *PetCreateBulk) SaveX(ctx context.Context) []*Pet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *PetCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *PetCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
