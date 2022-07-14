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
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks    []Hook
	mutation *FileTypeMutation
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (ftu *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	ftu.mutation.Where(ps...)
	return ftu
}

// SetName sets the "name" field.
func (ftu *FileTypeUpdate) SetName(s string) *FileTypeUpdate {
	ftu.mutation.SetName(s)
	return ftu
}

// SetType sets the "type" field.
func (ftu *FileTypeUpdate) SetType(f filetype.Type) *FileTypeUpdate {
	ftu.mutation.SetType(f)
	return ftu
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableType(f *filetype.Type) *FileTypeUpdate {
	if f != nil {
		ftu.SetType(*f)
	}
	return ftu
}

// SetState sets the "state" field.
func (ftu *FileTypeUpdate) SetState(f filetype.State) *FileTypeUpdate {
	ftu.mutation.SetState(f)
	return ftu
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftu *FileTypeUpdate) SetNillableState(f *filetype.State) *FileTypeUpdate {
	if f != nil {
		ftu.SetState(*f)
	}
	return ftu
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftu *FileTypeUpdate) AddFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.AddFileIDs(ids...)
	return ftu
}

// AddFiles adds the "files" edges to the File entity.
func (ftu *FileTypeUpdate) AddFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftu *FileTypeUpdate) Mutation() *FileTypeMutation {
	return ftu.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftu *FileTypeUpdate) ClearFiles() *FileTypeUpdate {
	ftu.mutation.ClearFiles()
	return ftu
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftu *FileTypeUpdate) RemoveFileIDs(ids ...string) *FileTypeUpdate {
	ftu.mutation.RemoveFileIDs(ids...)
	return ftu
}

// RemoveFiles removes "files" edges to File entities.
func (ftu *FileTypeUpdate) RemoveFiles(f ...*File) *FileTypeUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftu.RemoveFileIDs(ids...)
}

// When runs the provided builder(s) if and only if condition is true.
func (ftu *FileTypeUpdate) When(condition bool, action func(builder *FileTypeUpdate)) *FileTypeUpdate {
	if condition {
		action(ftu)
	}

	return ftu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ftu *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ftu.hooks) == 0 {
		if err = ftu.check(); err != nil {
			return 0, err
		}
		affected, err = ftu.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftu.check(); err != nil {
				return 0, err
			}
			ftu.mutation = mutation
			affected, err = ftu.gremlinSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ftu.hooks) - 1; i >= 0; i-- {
			if ftu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ftu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftu *FileTypeUpdate) check() error {
	if v, ok := ftu.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := ftu.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

func (ftu *FileTypeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftu.gremlin().Query()
	if err := ftu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (ftu *FileTypeUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V().HasLabel(filetype.Label)
	for _, p := range ftu.mutation.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ftu.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	if value, ok := ftu.mutation.GetType(); ok {
		v.Property(dsl.Single, filetype.FieldType, value)
	}
	if value, ok := ftu.mutation.State(); ok {
		v.Property(dsl.Single, filetype.FieldState, value)
	}
	for _, id := range ftu.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ftu.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ConstraintError{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FileTypeMutation
}

// SetName sets the "name" field.
func (ftuo *FileTypeUpdateOne) SetName(s string) *FileTypeUpdateOne {
	ftuo.mutation.SetName(s)
	return ftuo
}

// SetType sets the "type" field.
func (ftuo *FileTypeUpdateOne) SetType(f filetype.Type) *FileTypeUpdateOne {
	ftuo.mutation.SetType(f)
	return ftuo
}

// SetNillableType sets the "type" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableType(f *filetype.Type) *FileTypeUpdateOne {
	if f != nil {
		ftuo.SetType(*f)
	}
	return ftuo
}

// SetState sets the "state" field.
func (ftuo *FileTypeUpdateOne) SetState(f filetype.State) *FileTypeUpdateOne {
	ftuo.mutation.SetState(f)
	return ftuo
}

// SetNillableState sets the "state" field if the given value is not nil.
func (ftuo *FileTypeUpdateOne) SetNillableState(f *filetype.State) *FileTypeUpdateOne {
	if f != nil {
		ftuo.SetState(*f)
	}
	return ftuo
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (ftuo *FileTypeUpdateOne) AddFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.AddFileIDs(ids...)
	return ftuo
}

// AddFiles adds the "files" edges to the File entity.
func (ftuo *FileTypeUpdateOne) AddFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (ftuo *FileTypeUpdateOne) Mutation() *FileTypeMutation {
	return ftuo.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (ftuo *FileTypeUpdateOne) ClearFiles() *FileTypeUpdateOne {
	ftuo.mutation.ClearFiles()
	return ftuo
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (ftuo *FileTypeUpdateOne) RemoveFileIDs(ids ...string) *FileTypeUpdateOne {
	ftuo.mutation.RemoveFileIDs(ids...)
	return ftuo
}

// RemoveFiles removes "files" edges to File entities.
func (ftuo *FileTypeUpdateOne) RemoveFiles(f ...*File) *FileTypeUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return ftuo.RemoveFileIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ftuo *FileTypeUpdateOne) Select(field string, fields ...string) *FileTypeUpdateOne {
	ftuo.fields = append([]string{field}, fields...)
	return ftuo
}

// Save executes the query and returns the updated FileType entity.
func (ftuo *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	var (
		err  error
		node *FileType
	)
	if len(ftuo.hooks) == 0 {
		if err = ftuo.check(); err != nil {
			return nil, err
		}
		node, err = ftuo.gremlinSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FileTypeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ftuo.check(); err != nil {
				return nil, err
			}
			ftuo.mutation = mutation
			node, err = ftuo.gremlinSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ftuo.hooks) - 1; i >= 0; i-- {
			if ftuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ftuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ftuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*FileType)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FileTypeMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	node, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ftuo *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ftuo *FileTypeUpdateOne) check() error {
	if v, ok := ftuo.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := ftuo.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

func (ftuo *FileTypeUpdateOne) gremlinSave(ctx context.Context) (*FileType, error) {
	res := &gremlin.Response{}
	id, ok := ftuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileType.id" for update`)}
	}
	query, bindings := ftuo.gremlin(id).Query()
	if err := ftuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FileType{config: ftuo.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FileTypeUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value, ok := ftuo.mutation.Name(); ok {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(filetype.Label, filetype.FieldName, value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(filetype.Label, filetype.FieldName, value)),
		})
		v.Property(dsl.Single, filetype.FieldName, value)
	}
	if value, ok := ftuo.mutation.GetType(); ok {
		v.Property(dsl.Single, filetype.FieldType, value)
	}
	if value, ok := ftuo.mutation.State(); ok {
		v.Property(dsl.Single, filetype.FieldState, value)
	}
	for _, id := range ftuo.mutation.RemovedFilesIDs() {
		tr := rv.Clone().OutE(filetype.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for _, id := range ftuo.mutation.FilesIDs() {
		v.AddE(filetype.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(filetype.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(filetype.Label, filetype.FilesLabel, id)),
		})
	}
	if len(ftuo.fields) > 0 {
		fields := make([]interface{}, 0, len(ftuo.fields)+1)
		fields = append(fields, true)
		for _, f := range ftuo.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
