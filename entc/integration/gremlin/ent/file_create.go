// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

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

// SetSize sets the "size" field.
func (fc *FileCreate) SetSize(i int) *FileCreate {
	fc.mutation.SetSize(i)
	return fc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fc *FileCreate) SetNillableSize(i *int) *FileCreate {
	if i != nil {
		fc.SetSize(*i)
	}
	return fc
}

// SetName sets the "name" field.
func (fc *FileCreate) SetName(s string) *FileCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetUser sets the "user" field.
func (fc *FileCreate) SetUser(s string) *FileCreate {
	fc.mutation.SetUser(s)
	return fc
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fc *FileCreate) SetNillableUser(s *string) *FileCreate {
	if s != nil {
		fc.SetUser(*s)
	}
	return fc
}

// SetGroup sets the "group" field.
func (fc *FileCreate) SetGroup(s string) *FileCreate {
	fc.mutation.SetGroup(s)
	return fc
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fc *FileCreate) SetNillableGroup(s *string) *FileCreate {
	if s != nil {
		fc.SetGroup(*s)
	}
	return fc
}

// SetOp sets the "op" field.
func (fc *FileCreate) SetOp(b bool) *FileCreate {
	fc.mutation.SetOp(b)
	return fc
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fc *FileCreate) SetNillableOp(b *bool) *FileCreate {
	if b != nil {
		fc.SetOp(*b)
	}
	return fc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fc *FileCreate) SetOwnerID(id string) *FileCreate {
	fc.mutation.SetOwnerID(id)
	return fc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableOwnerID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetOwnerID(*id)
	}
	return fc
}

// SetOwner sets the "owner" edge to the User entity.
func (fc *FileCreate) SetOwner(u *User) *FileCreate {
	return fc.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fc *FileCreate) SetTypeID(id string) *FileCreate {
	fc.mutation.SetTypeID(id)
	return fc
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fc *FileCreate) SetNillableTypeID(id *string) *FileCreate {
	if id != nil {
		fc = fc.SetTypeID(*id)
	}
	return fc
}

// SetType sets the "type" edge to the FileType entity.
func (fc *FileCreate) SetType(f *FileType) *FileCreate {
	return fc.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fc *FileCreate) AddFieldIDs(ids ...string) *FileCreate {
	fc.mutation.AddFieldIDs(ids...)
	return fc
}

// AddField adds the "field" edges to the FieldType entity.
func (fc *FileCreate) AddField(f ...*FieldType) *FileCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fc *FileCreate) Mutation() *FileMutation {
	return fc.mutation
}

// Save creates the File in the database.
func (fc *FileCreate) Save(ctx context.Context) (*File, error) {
	var (
		err  error
		node *File
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			node, err = fc.gremlinSave(ctx)
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			mut = fc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, fc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FileCreate) SaveX(ctx context.Context) *File {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// defaults sets the default values of the builder before save.
func (fc *FileCreate) defaults() {
	if _, ok := fc.mutation.Size(); !ok {
		v := file.DefaultSize
		fc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FileCreate) check() error {
	if _, ok := fc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New("ent: missing required field \"size\"")}
	}
	if v, ok := fc.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf("ent: validator failed for field \"size\": %w", err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	return nil
}

func (fc *FileCreate) gremlinSave(ctx context.Context) (*File, error) {
	res := &gremlin.Response{}
	query, bindings := fc.gremlin().Query()
	if err := fc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	f := &File{config: fc.config}
	if err := f.FromResponse(res); err != nil {
		return nil, err
	}
	return f, nil
}

func (fc *FileCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.AddV(file.Label)
	if value, ok := fc.mutation.Size(); ok {
		v.Property(dsl.Single, file.FieldSize, value)
	}
	if value, ok := fc.mutation.Name(); ok {
		v.Property(dsl.Single, file.FieldName, value)
	}
	if value, ok := fc.mutation.User(); ok {
		v.Property(dsl.Single, file.FieldUser, value)
	}
	if value, ok := fc.mutation.Group(); ok {
		v.Property(dsl.Single, file.FieldGroup, value)
	}
	if value, ok := fc.mutation.GetOp(); ok {
		v.Property(dsl.Single, file.FieldOp, value)
	}
	for _, id := range fc.mutation.OwnerIDs() {
		v.AddE(user.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fc.mutation.TypeIDs() {
		v.AddE(filetype.FilesLabel).From(g.V(id)).InV()
	}
	for _, id := range fc.mutation.FieldIDs() {
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
	builders []*FileCreate
}
