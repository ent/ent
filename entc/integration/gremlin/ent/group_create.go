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

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/group"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
}

// SetActive sets the "active" field.
func (m *GroupCreate) SetActive(v bool) *GroupCreate {
	m.mutation.SetActive(v)
	return m
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (m *GroupCreate) SetNillableActive(v *bool) *GroupCreate {
	if v != nil {
		m.SetActive(*v)
	}
	return m
}

// SetExpire sets the "expire" field.
func (m *GroupCreate) SetExpire(v time.Time) *GroupCreate {
	m.mutation.SetExpire(v)
	return m
}

// SetType sets the "type" field.
func (m *GroupCreate) SetType(v string) *GroupCreate {
	m.mutation.SetType(v)
	return m
}

// SetNillableType sets the "type" field if the given value is not nil.
func (m *GroupCreate) SetNillableType(v *string) *GroupCreate {
	if v != nil {
		m.SetType(*v)
	}
	return m
}

// SetMaxUsers sets the "max_users" field.
func (m *GroupCreate) SetMaxUsers(v int) *GroupCreate {
	m.mutation.SetMaxUsers(v)
	return m
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (m *GroupCreate) SetNillableMaxUsers(v *int) *GroupCreate {
	if v != nil {
		m.SetMaxUsers(*v)
	}
	return m
}

// SetName sets the "name" field.
func (m *GroupCreate) SetName(v string) *GroupCreate {
	m.mutation.SetName(v)
	return m
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (m *GroupCreate) AddFileIDs(ids ...string) *GroupCreate {
	m.mutation.AddFileIDs(ids...)
	return m
}

// AddFiles adds the "files" edges to the File entity.
func (m *GroupCreate) AddFiles(v ...*File) *GroupCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFileIDs(ids...)
}

// AddBlockedIDs adds the "blocked" edge to the User entity by IDs.
func (m *GroupCreate) AddBlockedIDs(ids ...string) *GroupCreate {
	m.mutation.AddBlockedIDs(ids...)
	return m
}

// AddBlocked adds the "blocked" edges to the User entity.
func (m *GroupCreate) AddBlocked(v ...*User) *GroupCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddBlockedIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (m *GroupCreate) AddUserIDs(ids ...string) *GroupCreate {
	m.mutation.AddUserIDs(ids...)
	return m
}

// AddUsers adds the "users" edges to the User entity.
func (m *GroupCreate) AddUsers(v ...*User) *GroupCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddUserIDs(ids...)
}

// SetInfoID sets the "info" edge to the GroupInfo entity by ID.
func (m *GroupCreate) SetInfoID(id string) *GroupCreate {
	m.mutation.SetInfoID(id)
	return m
}

// SetInfo sets the "info" edge to the GroupInfo entity.
func (m *GroupCreate) SetInfo(v *GroupInfo) *GroupCreate {
	return m.SetInfoID(v.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (m *GroupCreate) Mutation() *GroupMutation {
	return m.mutation
}

// Save creates the Group in the database.
func (c *GroupCreate) Save(ctx context.Context) (*Group, error) {
	c.defaults()
	return withHooks(ctx, c.gremlinSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *GroupCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *GroupCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (c *GroupCreate) defaults() {
	if _, ok := c.mutation.Active(); !ok {
		v := group.DefaultActive
		c.mutation.SetActive(v)
	}
	if _, ok := c.mutation.MaxUsers(); !ok {
		v := group.DefaultMaxUsers
		c.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *GroupCreate) check() error {
	if _, ok := c.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Group.active"`)}
	}
	if _, ok := c.mutation.Expire(); !ok {
		return &ValidationError{Name: "expire", err: errors.New(`ent: missing required field "Group.expire"`)}
	}
	if v, ok := c.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Group.type": %w`, err)}
		}
	}
	if v, ok := c.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return &ValidationError{Name: "max_users", err: fmt.Errorf(`ent: validator failed for field "Group.max_users": %w`, err)}
		}
	}
	if _, ok := c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if v, ok := c.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
		}
	}
	if len(c.mutation.InfoIDs()) == 0 {
		return &ValidationError{Name: "info", err: errors.New(`ent: missing required edge "Group.info"`)}
	}
	return nil
}

func (c *GroupCreate) gremlinSave(ctx context.Context) (*Group, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := c.gremlin().Query()
	if err := c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Group{config: c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	c.mutation.id = &rnode.ID
	c.mutation.done = true
	return rnode, nil
}

func (c *GroupCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(group.Label)
	if value, ok := c.mutation.Active(); ok {
		v.Property(dsl.Single, group.FieldActive, value)
	}
	if value, ok := c.mutation.Expire(); ok {
		v.Property(dsl.Single, group.FieldExpire, value)
	}
	if value, ok := c.mutation.GetType(); ok {
		v.Property(dsl.Single, group.FieldType, value)
	}
	if value, ok := c.mutation.MaxUsers(); ok {
		v.Property(dsl.Single, group.FieldMaxUsers, value)
	}
	if value, ok := c.mutation.Name(); ok {
		v.Property(dsl.Single, group.FieldName, value)
	}
	for _, id := range c.mutation.FilesIDs() {
		v.AddE(group.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.FilesLabel, id)),
		})
	}
	for _, id := range c.mutation.BlockedIDs() {
		v.AddE(group.BlockedLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.BlockedLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(group.Label, group.BlockedLabel, id)),
		})
	}
	for _, id := range c.mutation.UsersIDs() {
		v.AddE(user.GroupsLabel).From(g.V(id)).InV()
	}
	for _, id := range c.mutation.InfoIDs() {
		v.AddE(group.InfoLabel).To(g.V(id)).OutV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
}
