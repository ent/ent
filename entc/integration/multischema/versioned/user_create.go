// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package versioned

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/multischema/versioned/friendship"
	"entgo.io/ent/entc/integration/multischema/versioned/group"
	"entgo.io/ent/entc/integration/multischema/versioned/pet"
	"entgo.io/ent/entc/integration/multischema/versioned/user"
	"entgo.io/ent/schema/field"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	mutation *UserMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (m *UserCreate) SetName(v string) *UserCreate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *UserCreate) SetNillableName(v *string) *UserCreate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// AddPetIDs adds the "pets" edge to the Pet entity by IDs.
func (m *UserCreate) AddPetIDs(ids ...int) *UserCreate {
	m.mutation.AddPetIDs(ids...)
	return m
}

// AddPets adds the "pets" edges to the Pet entity.
func (m *UserCreate) AddPets(v ...*Pet) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddPetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (m *UserCreate) AddGroupIDs(ids ...int) *UserCreate {
	m.mutation.AddGroupIDs(ids...)
	return m
}

// AddGroups adds the "groups" edges to the Group entity.
func (m *UserCreate) AddGroups(v ...*Group) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddGroupIDs(ids...)
}

// AddFriendIDs adds the "friends" edge to the User entity by IDs.
func (m *UserCreate) AddFriendIDs(ids ...int) *UserCreate {
	m.mutation.AddFriendIDs(ids...)
	return m
}

// AddFriends adds the "friends" edges to the User entity.
func (m *UserCreate) AddFriends(v ...*User) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFriendIDs(ids...)
}

// AddFollowerIDs adds the "followers" edge to the User entity by IDs.
func (m *UserCreate) AddFollowerIDs(ids ...int) *UserCreate {
	m.mutation.AddFollowerIDs(ids...)
	return m
}

// AddFollowers adds the "followers" edges to the User entity.
func (m *UserCreate) AddFollowers(v ...*User) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the "following" edge to the User entity by IDs.
func (m *UserCreate) AddFollowingIDs(ids ...int) *UserCreate {
	m.mutation.AddFollowingIDs(ids...)
	return m
}

// AddFollowing adds the "following" edges to the User entity.
func (m *UserCreate) AddFollowing(v ...*User) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFollowingIDs(ids...)
}

// AddFriendshipIDs adds the "friendships" edge to the Friendship entity by IDs.
func (m *UserCreate) AddFriendshipIDs(ids ...int) *UserCreate {
	m.mutation.AddFriendshipIDs(ids...)
	return m
}

// AddFriendships adds the "friendships" edges to the Friendship entity.
func (m *UserCreate) AddFriendships(v ...*Friendship) *UserCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFriendshipIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (m *UserCreate) Mutation() *UserMutation {
	return m.mutation
}

// Save creates the User in the database.
func (c *UserCreate) Save(ctx context.Context) (*User, error) {
	c.defaults()
	return withHooks(ctx, c.sqlSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *UserCreate) SaveX(ctx context.Context) *User {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *UserCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *UserCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (c *UserCreate) defaults() {
	if _, ok := c.mutation.Name(); !ok {
		v := user.DefaultName
		c.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *UserCreate) check() error {
	if _, ok := c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`versioned: missing required field "User.name"`)}
	}
	return nil
}

func (c *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	_node, _spec := c.createSpec()
	if err := sqlgraph.CreateNode(ctx, c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	c.mutation.id = &_node.ID
	c.mutation.done = true
	return _node, nil
}

func (c *UserCreate) createSpec() (*User, *sqlgraph.CreateSpec) {
	var (
		_node = &User{config: c.config}
		_spec = sqlgraph.NewCreateSpec(user.Table, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	)
	_spec.Schema = c.schemaConfig.User
	if value, ok := c.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := c.mutation.PetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.PetsTable,
			Columns: []string{user.PetsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(pet.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.Pet
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.FriendsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FriendsTable,
			Columns: user.FriendsPrimaryKey,
			Bidi:    true,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.Friendship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &FriendshipCreate{config: c.config, mutation: newFriendshipMutation(c.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.FollowersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.UserFollowing
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.FollowingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.UserFollowing
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.FriendshipsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   user.FriendshipsTable,
			Columns: []string{user.FriendshipsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(friendship.FieldID, field.TypeInt),
			},
		}
		edge.Schema = c.schemaConfig.Friendship
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
func (c *UserCreateBulk) Save(ctx context.Context) ([]*User, error) {
	if c.err != nil {
		return nil, c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(c.builders))
	nodes := make([]*User, len(c.builders))
	mutators := make([]Mutator, len(c.builders))
	for i := range c.builders {
		func(i int, root context.Context) {
			builder := c.builders[i]
			builder.defaults()
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
					_, err = mutators[i+1].Mutate(root, c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, c.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (c *UserCreateBulk) SaveX(ctx context.Context) []*User {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *UserCreateBulk) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *UserCreateBulk) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}
