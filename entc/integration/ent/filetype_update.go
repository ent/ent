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
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// FileTypeUpdate is the builder for updating FileType entities.
type FileTypeUpdate struct {
	config
	hooks     []Hook
	mutation  *FileTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (u *FileTypeUpdate) Where(ps ...predicate.FileType) *FileTypeUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetName sets the "name" field.
func (m *FileTypeUpdate) SetName(v string) *FileTypeUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *FileTypeUpdate) SetNillableName(v *string) *FileTypeUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetType sets the "type" field.
func (m *FileTypeUpdate) SetType(v filetype.Type) *FileTypeUpdate {
	m.mutation.SetType(v)
	return m
}

// SetNillableType sets the "type" field if the given value is not nil.
func (m *FileTypeUpdate) SetNillableType(v *filetype.Type) *FileTypeUpdate {
	if v != nil {
		m.SetType(*v)
	}
	return m
}

// SetState sets the "state" field.
func (m *FileTypeUpdate) SetState(v filetype.State) *FileTypeUpdate {
	m.mutation.SetState(v)
	return m
}

// SetNillableState sets the "state" field if the given value is not nil.
func (m *FileTypeUpdate) SetNillableState(v *filetype.State) *FileTypeUpdate {
	if v != nil {
		m.SetState(*v)
	}
	return m
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (m *FileTypeUpdate) AddFileIDs(ids ...int) *FileTypeUpdate {
	m.mutation.AddFileIDs(ids...)
	return m
}

// AddFiles adds the "files" edges to the File entity.
func (m *FileTypeUpdate) AddFiles(v ...*File) *FileTypeUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (m *FileTypeUpdate) Mutation() *FileTypeMutation {
	return m.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (u *FileTypeUpdate) ClearFiles() *FileTypeUpdate {
	u.mutation.ClearFiles()
	return u
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (u *FileTypeUpdate) RemoveFileIDs(ids ...int) *FileTypeUpdate {
	u.mutation.RemoveFileIDs(ids...)
	return u
}

// RemoveFiles removes "files" edges to File entities.
func (u *FileTypeUpdate) RemoveFiles(v ...*File) *FileTypeUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveFileIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *FileTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *FileTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *FileTypeUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileTypeUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *FileTypeUpdate) check() error {
	if v, ok := u.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := u.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *FileTypeUpdate) Modify(modifiers ...func(*sql.UpdateBuilder)) *FileTypeUpdate {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *FileTypeUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(filetype.FieldName, field.TypeString, value)
	}
	if value, ok := u.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeEnum, value)
	}
	if value, ok := u.mutation.State(); ok {
		_spec.SetField(filetype.FieldState, field.TypeEnum, value)
	}
	if u.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedFilesIDs(); len(nodes) > 0 && !u.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(u.modifiers...)
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// FileTypeUpdateOne is the builder for updating a single FileType entity.
type FileTypeUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FileTypeMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (m *FileTypeUpdateOne) SetName(v string) *FileTypeUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *FileTypeUpdateOne) SetNillableName(v *string) *FileTypeUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// SetType sets the "type" field.
func (m *FileTypeUpdateOne) SetType(v filetype.Type) *FileTypeUpdateOne {
	m.mutation.SetType(v)
	return m
}

// SetNillableType sets the "type" field if the given value is not nil.
func (m *FileTypeUpdateOne) SetNillableType(v *filetype.Type) *FileTypeUpdateOne {
	if v != nil {
		m.SetType(*v)
	}
	return m
}

// SetState sets the "state" field.
func (m *FileTypeUpdateOne) SetState(v filetype.State) *FileTypeUpdateOne {
	m.mutation.SetState(v)
	return m
}

// SetNillableState sets the "state" field if the given value is not nil.
func (m *FileTypeUpdateOne) SetNillableState(v *filetype.State) *FileTypeUpdateOne {
	if v != nil {
		m.SetState(*v)
	}
	return m
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (m *FileTypeUpdateOne) AddFileIDs(ids ...int) *FileTypeUpdateOne {
	m.mutation.AddFileIDs(ids...)
	return m
}

// AddFiles adds the "files" edges to the File entity.
func (m *FileTypeUpdateOne) AddFiles(v ...*File) *FileTypeUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFileIDs(ids...)
}

// Mutation returns the FileTypeMutation object of the builder.
func (m *FileTypeUpdateOne) Mutation() *FileTypeMutation {
	return m.mutation
}

// ClearFiles clears all "files" edges to the File entity.
func (u *FileTypeUpdateOne) ClearFiles() *FileTypeUpdateOne {
	u.mutation.ClearFiles()
	return u
}

// RemoveFileIDs removes the "files" edge to File entities by IDs.
func (u *FileTypeUpdateOne) RemoveFileIDs(ids ...int) *FileTypeUpdateOne {
	u.mutation.RemoveFileIDs(ids...)
	return u
}

// RemoveFiles removes "files" edges to File entities.
func (u *FileTypeUpdateOne) RemoveFiles(v ...*File) *FileTypeUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveFileIDs(ids...)
}

// Where appends a list predicates to the FileTypeUpdate builder.
func (u *FileTypeUpdateOne) Where(ps ...predicate.FileType) *FileTypeUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *FileTypeUpdateOne) Select(field string, fields ...string) *FileTypeUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated FileType entity.
func (u *FileTypeUpdateOne) Save(ctx context.Context) (*FileType, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *FileTypeUpdateOne) SaveX(ctx context.Context) *FileType {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *FileTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *FileTypeUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *FileTypeUpdateOne) check() error {
	if v, ok := u.mutation.GetType(); ok {
		if err := filetype.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "FileType.type": %w`, err)}
		}
	}
	if v, ok := u.mutation.State(); ok {
		if err := filetype.StateValidator(v); err != nil {
			return &ValidationError{Name: "state", err: fmt.Errorf(`ent: validator failed for field "FileType.state": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (u *FileTypeUpdateOne) Modify(modifiers ...func(*sql.UpdateBuilder)) *FileTypeUpdateOne {
	u.modifiers = append(u.modifiers, modifiers...)
	return u
}

func (u *FileTypeUpdateOne) sqlSave(ctx context.Context) (_n *FileType, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(filetype.Table, filetype.Columns, sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FileType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, filetype.FieldID)
		for _, f := range fields {
			if !filetype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != filetype.FieldID {
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
		_spec.SetField(filetype.FieldName, field.TypeString, value)
	}
	if value, ok := u.mutation.GetType(); ok {
		_spec.SetField(filetype.FieldType, field.TypeEnum, value)
	}
	if value, ok := u.mutation.State(); ok {
		_spec.SetField(filetype.FieldState, field.TypeEnum, value)
	}
	if u.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedFilesIDs(); len(nodes) > 0 && !u.mutation.FilesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   filetype.FilesTable,
			Columns: []string{filetype.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(u.modifiers...)
	_n = &FileType{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{filetype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
