// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

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
	"entgo.io/ent/entc/integration/gremlin/ent/filetype"
)

// FileTypeCreate is the builder for creating a FileType entity.
type FileTypeCreate struct {
	config
	mutation *FileTypeMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (_c *FileTypeCreate) SetName(s string) *FileTypeCreate {
	_c.mutation.SetName(s)
	return _c
}

// SetType sets the "type" field.
func (_c *FileTypeCreate) SetType(f filetype.Type) *FileTypeCreate {
	_c.mutation.SetType(f)
	return _c
}

// SetNillableType sets the "type" field if the given value is not nil.
func (_c *FileTypeCreate) SetNillableType(f *filetype.Type) *FileTypeCreate {
	if f != nil {
		_c.SetType(*f)
	}
	return _c
}

// SetState sets the "state" field.
func (_c *FileTypeCreate) SetState(f filetype.State) *FileTypeCreate {
	_c.mutation.SetState(f)
	return _c
}

// SetNillableState sets the "state" field if the given value is not nil.
func (_c *FileTypeCreate) SetNillableState(f *filetype.State) *FileTypeCreate {
	if f != nil {
		_c.SetState(*f)
	}
	return _c
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (_c *FileTypeCreate) AddFileIDs(ids ...string) *FileTypeCreate {
	_c.mutation.AddFileIDs(ids...)
	return _c
}

// AddFiles adds the "files" edges to the File entity.
func (_c *FileTypeCreate) AddFiles(f ...*File) *FileTypeCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return _c.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (_c *FileTypeCreate) Mutation() *FileTypeMutation {
	return _c.mutation
}

// Save creates the FileType in the database.
func (_c *FileTypeCreate) Save(ctx context.Context) (*FileType, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *FileTypeCreate) SaveX(ctx context.Context) *FileType {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *FileTypeCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *FileTypeCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *FileTypeCreate) defaults() {
	if _, ok := _c.mutation.GetType(); !ok {
		v := filetype.DefaultType
		_c.mutation.SetType(v)
	}
	if _, ok := _c.mutation.State(); !ok {
		v := filetype.DefaultState
		_c.mutation.SetState(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *FileTypeCreate) check() error {
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "FileType.name"`)}
	}
	if _, ok := _c.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "FileType.type"`)}
	}
	if v, ok := _c.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if _, ok := _c.mutation.State(); !ok {
		return &ValidationError{Name: "state", err: errors.New(`ent: missing required field "FileType.state"`)}
	}
	if v, ok := _c.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

func (_c *FileTypeCreate) gremlinSave(ctx context.Context) (*FileType, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &FileType{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *FileTypeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(filetype.Label)
	if value, ok := _c.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	if value, ok := _c.mutation.GetType(); ok {
		v.Property(dsl.Single, filetype.FieldType, value)
	}
	if value, ok := _c.mutation.State(); ok {
		v.Property(dsl.Single, filetype.FieldState, value)
	}
	for _, id := range _c.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
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

// FileTypeCreateBulk is the builder for creating many FileType entities in bulk.
type FileTypeCreateBulk struct {
	config
	err      error
	builders []*FileTypeCreate
}
