// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/idtype/ent/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (_c *UserCreate) SetName(v string) *UserCreate {
	_c.mutation.SetName(v)
	return _c
}

// SetSpouseID sets the "spouse" edge to the User entity by ID.
func (_c *UserCreate) SetSpouseID(id uint64) *UserCreate {
	_c.mutation.SetSpouseID(id)
	return _c
}

// SetNillableSpouseID sets the "spouse" edge to the User entity by ID if the given value is not nil.
func (_c *UserCreate) SetNillableSpouseID(id *uint64) *UserCreate {
	if id != nil {
		_c = _c.SetSpouseID(*id)
	}
	return _c
}

// SetSpouse sets the "spouse" edge to the User entity.
func (_c *UserCreate) SetSpouse(v *User) *UserCreate {
	return _c.SetSpouseID(v.ID)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (_c *UserCreate) AddFollowerIDs(ids ...uint64) *UserCreate {
	_c.mutation.AddFollowerIDs(ids...)
	return _c
}

// AddFollowers adds the "followers" edges to the User entity.
func (_c *UserCreate) AddFollowers(v ...*User) *UserCreate {
	ids := make([]uint64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (_c *UserCreate) AddFollowingIDs(ids ...uint64) *UserCreate {
	_c.mutation.AddFollowingIDs(ids...)
	return _c
}

// AddFollowing adds the "following" edges to the User entity.
func (_c *UserCreate) AddFollowing(v ...*User) *UserCreate {
	ids := make([]uint64, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddFollowingIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_c *UserCreate) Mutation() *UserMutation {
	return _c.mutation
}

// Save creates the User in the database.
func (_c *UserCreate) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *UserCreate) SaveX(ctx context.Context) *User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserCreate) check() error {
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "User.name"`)}
	}
	return nil
}

func (_c *UserCreate) sqlSave(ctx context.Context) (*User, error) {
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
	_node.ID = uint64(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64))
	)
	if value, ok := _c.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := _c.mutation.SpouseIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   user.SpouseTable,
			Columns: []string{user.SpouseColumn},
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_spouse = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// UserCreateBulk is the builder for creating many User entities in bulk.
type UserCreateBulk struct {
	config
	err      error
	builders []*UserCreate
}

// Save creates the User entities in the database.
func (_c *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*User, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserMutation)
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
					nodes[i].ID = uint64(id)
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
func (_c *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
