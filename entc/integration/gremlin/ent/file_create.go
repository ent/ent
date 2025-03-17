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
	"entgo.io/ent/entc/integration/gremlin/ent/file"
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// FileCreate is the builder for creating a File entity.
type FileCreate struct {
	config
	mutation *FileMutation
	hooks    []Hook
}

// SetSetID sets the "set_id" field.
func (m *FileCreate) SetSetID(v int) *FileCreate {
	m.mutation.SetSetID(v)
	return m
}

// SetNillableSetID sets the "set_id" field if the given value is not nil.
func (m *FileCreate) SetNillableSetID(v *int) *FileCreate {
	if v != nil {
		m.SetSetID(*v)
	}
	return m
}

// SetSize sets the "size" field.
func (m *FileCreate) SetSize(v int) *FileCreate {
	m.mutation.SetSize(v)
	return m
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (m *FileCreate) SetNillableSize(v *int) *FileCreate {
	if v != nil {
		m.SetSize(*v)
	}
	return m
}

// SetName sets the "name" field.
func (m *FileCreate) SetName(v string) *FileCreate {
	m.mutation.SetName(v)
	return m
}

// SetUser sets the "user" field.
func (m *FileCreate) SetUser(v string) *FileCreate {
	m.mutation.SetUser(v)
	return m
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (m *FileCreate) SetNillableUser(v *string) *FileCreate {
	if v != nil {
		m.SetUser(*v)
	}
	return m
}

// SetGroup sets the "group" field.
func (m *FileCreate) SetGroup(v string) *FileCreate {
	m.mutation.SetGroup(v)
	return m
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (m *FileCreate) SetNillableGroup(v *string) *FileCreate {
	if v != nil {
		m.SetGroup(*v)
	}
	return m
}

// SetOp sets the "op" field.
func (m *FileCreate) SetOp(v bool) *FileCreate {
	m.mutation.SetOpField(v)
	return m
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (m *FileCreate) SetNillableOp(v *bool) *FileCreate {
	if v != nil {
		m.SetOp(*v)
	}
	return m
}

// SetFieldID sets the "field_id" field.
func (m *FileCreate) SetFieldID(v int) *FileCreate {
	m.mutation.SetFieldID(v)
	return m
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (m *FileCreate) SetNillableFieldID(v *int) *FileCreate {
	if v != nil {
		m.SetFieldID(*v)
	}
	return m
}

// SetCreateTime sets the "create_time" field.
func (m *FileCreate) SetCreateTime(v time.Time) *FileCreate {
	m.mutation.SetCreateTime(v)
	return m
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (m *FileCreate) SetNillableCreateTime(v *time.Time) *FileCreate {
	if v != nil {
		m.SetCreateTime(*v)
	}
	return m
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (m *FileCreate) SetOwnerID(id string) *FileCreate {
	m.mutation.SetOwnerID(id)
	return m
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (m *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		m = m.SetOwnerID(*id)
	}
	return m
}

// SetOwner sets the "owner" edge to the User entity.
func (m *FileCreate) SetOwner(v *User) *FileCreate {
	return m.SetOwnerID(v.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (m *FileCreate) SetTypeID(id string) *FileCreate {
	m.mutation.SetTypeID(id)
	return m
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (m *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		m = m.SetTypeID(*id)
	}
	return m
}

// SetType sets the "type" edge to the FileType entity.
func (m *FileCreate) SetType(v *FileType) *FileCreate {
	return m.SetTypeID(v.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (m *FileCreate) AddFieldIDs(ids ...string) *FileCreate {
	m.mutation.AddFieldIDs(ids...)
	return m
}

// AddField adds the "field" edges to the FieldType entity.
func (m *FileCreate) AddField(v ...*FieldType) *FileCreate {
	ids := make([]string, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (m *FileCreate) Mutation() *FileMutation {
	return m.mutation
}

// Save creates the File in the database.
func (c *FileCreate) Save(ctx context.Context) (*File, error) {
	c.defaults()
	return withHooks(ctx, c.gremlinSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *FileCreate) SaveX(ctx context.Context) *File {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *FileCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *FileCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (c *FileCreate) defaults() {
	if _, ok := c.mutation.Size(); !ok {
		v := file.DefaultSize
		c.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *FileCreate) check() error {
	if v, ok := c.mutation.SetID(); ok {
		if err := file.SetIDValidator(v); err != nil {
			return &ValidationError{Name: "set_id", err: fmt.Errorf(`ent: validator failed for field "File.set_id": %w`, err)}
		}
	}
	if _, ok := c.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "File.size"`)}
	}
	if v, ok := c.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	if _, ok := c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "File.name"`)}
	}
	return nil
}

func (c *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
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
	rnode := &File{config: c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	c.mutation.id = &rnode.ID
	c.mutation.done = true
	return rnode, nil
}

func (c *FileCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(file.Label)
	if value, ok := c.mutation.SetID(); ok {
		v.Property(dsl.Single, file.FieldSetID, value)
	}
	if value, ok := c.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := c.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := c.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := c.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := c.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	if value, ok := c.mutation.FieldID(); ok {
		v.Property(dsl.Single, file.FieldFieldID, value)
	}
	if value, ok := c.mutation.CreateTime(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(file.Label, file.FieldCreateTime, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(file.Label, file.FieldCreateTime, value)),
		})
		v.Property(dsl.Single, file.FieldCreateTime, value)
	}
	for _, id := range c.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range c.mutation.TypeIDs() {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range c.mutation.FieldIDs() {
		v.AddE(file.FieldLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(file.FieldLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(file.Label, file.FieldLabel, id)),
		})
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

// FileCreateBulk is the builder for creating many File entities in bulk.
type FileCreateBulk struct {
	config
	err      error
	builders []*FileCreate
}
