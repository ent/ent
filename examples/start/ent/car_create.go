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
	"entgo.io/ent/examples/start/ent/car"
	"entgo.io/ent/examples/start/ent/user"
	"entgo.io/ent/schema/field"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	mutation *CarMutation
	hooks    []Hook
}

// SetModel sets the "model" field.
func (_c *CarCreate) SetModel(s string) *CarCreate {
	_c.mutation.SetModel(s)
	return _c
}

// SetRegisteredAt sets the "registered_at" field.
func (_c *CarCreate) SetRegisteredAt(t time.Time) *CarCreate {
	_c.mutation.SetRegisteredAt(t)
	return _c
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_c *CarCreate) SetOwnerID(id int) *CarCreate {
	_c.mutation.SetOwnerID(id)
	return _c
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_c *CarCreate) SetNillableOwnerID(id *int) *CarCreate {
	if id != nil {
		_c = _c.SetOwnerID(*id)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *CarCreate) SetOwner(u *User) *CarCreate {
	return _c.SetOwnerID(u.ID)
}

// Mutation returns the CarMutation object of the builder.
func (_c *CarCreate) Mutation() *CarMutation {
	return _c.mutation
}

// Save creates the Car in the database.
func (_c *CarCreate) Save(ctx context.Context) (*Car, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CarCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CarCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *CarCreate) check() error {
	if _, ok := _c.mutation.Model(); !ok {
		return &ValidationError{Name: "model", err: errors.New(`ent: missing required field "Car.model"`)}
	}
	if _, ok := _c.mutation.RegisteredAt(); !ok {
		return &ValidationError{Name: "registered_at", err: errors.New(`ent: missing required field "Car.registered_at"`)}
	}
	return nil
}

func (_c *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
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

func (_c *CarCreate) createSpec() (*Car, *sqlgraph.CreateSpec) {
	var (
		_node = &Car{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(car.Table, sqlgraph.NewFieldSpec(car.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Model(); ok {
		_spec.SetField(car.FieldModel, field.TypeString, value)
		_node.Model = value
	}
	if value, ok := _c.mutation.RegisteredAt(); ok {
		_spec.SetField(car.FieldRegisteredAt, field.TypeTime, value)
		_node.RegisteredAt = value
	}
	if nodes := _c.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   car.OwnerTable,
			Columns: []string{car.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_cars = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CarCreateBulk is the builder for creating many Car entities in bulk.
type CarCreateBulk struct {
	config
	err      error
	builders []*CarCreate
}

// Save creates the Car entities in the database.
func (_c *CarCreateBulk) Save(ctx context.Context) ([]*Car, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Car, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CarMutation)
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
func (_c *CarCreateBulk) SaveX(ctx context.Context) []*Car {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *CarCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *CarCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
