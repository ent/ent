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
	"entgo.io/ent/entc/integration/ent/fieldtype"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/filetype"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// FileUpdate is the builder for updating File entities.
type FileUpdate struct {
	config
	hooks     []Hook
	mutation  *FileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the FileUpdate builder.
func (fu *FileUpdate) Where(ps ...predicate.File) *FileUpdate {
	fu.mutation.Where(ps...)
	return fu
}

// SetSize sets the "size" field.
func (fu *FileUpdate) SetSize(i int) *FileUpdate {
	fu.mutation.ResetSize()
	fu.mutation.SetSize(i)
	return fu
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fu *FileUpdate) SetNillableSize(i *int) *FileUpdate {
	if i != nil {
		fu.SetSize(*i)
	}
	return fu
}

// AddSize adds i to the "size" field.
func (fu *FileUpdate) AddSize(i int) *FileUpdate {
	fu.mutation.AddSize(i)
	return fu
}

// SetName sets the "name" field.
func (fu *FileUpdate) SetName(s string) *FileUpdate {
	fu.mutation.SetName(s)
	return fu
}

// SetUser sets the "user" field.
func (fu *FileUpdate) SetUser(s string) *FileUpdate {
	fu.mutation.SetUser(s)
	return fu
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fu *FileUpdate) SetNillableUser(s *string) *FileUpdate {
	if s != nil {
		fu.SetUser(*s)
	}
	return fu
}

// ClearUser clears the value of the "user" field.
func (fu *FileUpdate) ClearUser() *FileUpdate {
	fu.mutation.ClearUser()
	return fu
}

// SetGroup sets the "group" field.
func (fu *FileUpdate) SetGroup(s string) *FileUpdate {
	fu.mutation.SetGroup(s)
	return fu
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fu *FileUpdate) SetNillableGroup(s *string) *FileUpdate {
	if s != nil {
		fu.SetGroup(*s)
	}
	return fu
}

// ClearGroup clears the value of the "group" field.
func (fu *FileUpdate) ClearGroup() *FileUpdate {
	fu.mutation.ClearGroup()
	return fu
}

// SetOp sets the "op" field.
func (fu *FileUpdate) SetOp(b bool) *FileUpdate {
	fu.mutation.SetOpField(b)
	return fu
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fu *FileUpdate) SetNillableOp(b *bool) *FileUpdate {
	if b != nil {
		fu.SetOp(*b)
	}
	return fu
}

// ClearOp clears the value of the "op" field.
func (fu *FileUpdate) ClearOp() *FileUpdate {
	fu.mutation.ClearOp()
	return fu
}

// SetFieldID sets the "field_id" field.
func (fu *FileUpdate) SetFieldID(i int) *FileUpdate {
	fu.mutation.ResetFieldID()
	fu.mutation.SetFieldID(i)
	return fu
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (fu *FileUpdate) SetNillableFieldID(i *int) *FileUpdate {
	if i != nil {
		fu.SetFieldID(*i)
	}
	return fu
}

// AddFieldID adds i to the "field_id" field.
func (fu *FileUpdate) AddFieldID(i int) *FileUpdate {
	fu.mutation.AddFieldID(i)
	return fu
}

// ClearFieldID clears the value of the "field_id" field.
func (fu *FileUpdate) ClearFieldID() *FileUpdate {
	fu.mutation.ClearFieldID()
	return fu
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fu *FileUpdate) SetOwnerID(id int) *FileUpdate {
	fu.mutation.SetOwnerID(id)
	return fu
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fu *FileUpdate) SetNillableOwnerID(id *int) *FileUpdate {
	if id != nil {
		fu = fu.SetOwnerID(*id)
	}
	return fu
}

// SetOwner sets the "owner" edge to the User entity.
func (fu *FileUpdate) SetOwner(u *User) *FileUpdate {
	return fu.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fu *FileUpdate) SetTypeID(id int) *FileUpdate {
	fu.mutation.SetTypeID(id)
	return fu
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fu *FileUpdate) SetNillableTypeID(id *int) *FileUpdate {
	if id != nil {
		fu = fu.SetTypeID(*id)
	}
	return fu
}

// SetType sets the "type" edge to the FileType entity.
func (fu *FileUpdate) SetType(f *FileType) *FileUpdate {
	return fu.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fu *FileUpdate) AddFieldIDs(ids ...int) *FileUpdate {
	fu.mutation.AddFieldIDs(ids...)
	return fu
}

// AddField adds the "field" edges to the FieldType entity.
func (fu *FileUpdate) AddField(f ...*FieldType) *FileUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fu *FileUpdate) Mutation() *FileMutation {
	return fu.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fu *FileUpdate) ClearOwner() *FileUpdate {
	fu.mutation.ClearOwner()
	return fu
}

// ClearType clears the "type" edge to the FileType entity.
func (fu *FileUpdate) ClearType() *FileUpdate {
	fu.mutation.ClearType()
	return fu
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (fu *FileUpdate) ClearFieldEdge() *FileUpdate {
	fu.mutation.ClearFieldEdge()
	return fu
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (fu *FileUpdate) RemoveFieldIDs(ids ...int) *FileUpdate {
	fu.mutation.RemoveFieldIDs(ids...)
	return fu
}

// RemoveField removes "field" edges to FieldType entities.
func (fu *FileUpdate) RemoveField(f ...*FieldType) *FileUpdate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fu.RemoveFieldIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fu *FileUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, fu.sqlSave, fu.mutation, fu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fu *FileUpdate) SaveX(ctx context.Context) int {
	affected, err := fu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fu *FileUpdate) Exec(ctx context.Context) error {
	_, err := fu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fu *FileUpdate) ExecX(ctx context.Context) {
	if err := fu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fu *FileUpdate) check() error {
	if v, ok := fu.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fu *FileUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FileUpdate {
	fu.modifiers = append(fu.modifiers, modifiers...)
	return fu
}

func (fu *FileUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	if ps := fu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fu.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fu.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fu.mutation.User(); ok {
		_spec.SetField(file.FieldUser, field.TypeString, value)
	}
	if fu.mutation.UserCleared() {
		_spec.ClearField(file.FieldUser, field.TypeString)
	}
	if value, ok := fu.mutation.Group(); ok {
		_spec.SetField(file.FieldGroup, field.TypeString, value)
	}
	if fu.mutation.GroupCleared() {
		_spec.ClearField(file.FieldGroup, field.TypeString)
	}
	if value, ok := fu.mutation.GetOp(); ok {
		_spec.SetField(file.FieldOp, field.TypeBool, value)
	}
	if fu.mutation.OpCleared() {
		_spec.ClearField(file.FieldOp, field.TypeBool)
	}
	if value, ok := fu.mutation.FieldID(); ok {
		_spec.SetField(file.FieldFieldID, field.TypeInt, value)
	}
	if value, ok := fu.mutation.AddedFieldID(); ok {
		_spec.AddField(file.FieldFieldID, field.TypeInt, value)
	}
	if fu.mutation.FieldIDCleared() {
		_spec.ClearField(file.FieldFieldID, field.TypeInt)
	}
	if fu.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fu.mutation.FieldEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.RemovedFieldIDs(); len(nodes) > 0 && !fu.mutation.FieldEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fu.mutation.FieldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, fu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fu.mutation.done = true
	return n, nil
}

// FileUpdateOne is the builder for updating a single File entity.
type FileUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *FileMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetSize sets the "size" field.
func (fuo *FileUpdateOne) SetSize(i int) *FileUpdateOne {
	fuo.mutation.ResetSize()
	fuo.mutation.SetSize(i)
	return fuo
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableSize(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetSize(*i)
	}
	return fuo
}

// AddSize adds i to the "size" field.
func (fuo *FileUpdateOne) AddSize(i int) *FileUpdateOne {
	fuo.mutation.AddSize(i)
	return fuo
}

// SetName sets the "name" field.
func (fuo *FileUpdateOne) SetName(s string) *FileUpdateOne {
	fuo.mutation.SetName(s)
	return fuo
}

// SetUser sets the "user" field.
func (fuo *FileUpdateOne) SetUser(s string) *FileUpdateOne {
	fuo.mutation.SetUser(s)
	return fuo
}

// SetNillableUser sets the "user" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableUser(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetUser(*s)
	}
	return fuo
}

// ClearUser clears the value of the "user" field.
func (fuo *FileUpdateOne) ClearUser() *FileUpdateOne {
	fuo.mutation.ClearUser()
	return fuo
}

// SetGroup sets the "group" field.
func (fuo *FileUpdateOne) SetGroup(s string) *FileUpdateOne {
	fuo.mutation.SetGroup(s)
	return fuo
}

// SetNillableGroup sets the "group" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableGroup(s *string) *FileUpdateOne {
	if s != nil {
		fuo.SetGroup(*s)
	}
	return fuo
}

// ClearGroup clears the value of the "group" field.
func (fuo *FileUpdateOne) ClearGroup() *FileUpdateOne {
	fuo.mutation.ClearGroup()
	return fuo
}

// SetOp sets the "op" field.
func (fuo *FileUpdateOne) SetOp(b bool) *FileUpdateOne {
	fuo.mutation.SetOpField(b)
	return fuo
}

// SetNillableOp sets the "op" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOp(b *bool) *FileUpdateOne {
	if b != nil {
		fuo.SetOp(*b)
	}
	return fuo
}

// ClearOp clears the value of the "op" field.
func (fuo *FileUpdateOne) ClearOp() *FileUpdateOne {
	fuo.mutation.ClearOp()
	return fuo
}

// SetFieldID sets the "field_id" field.
func (fuo *FileUpdateOne) SetFieldID(i int) *FileUpdateOne {
	fuo.mutation.ResetFieldID()
	fuo.mutation.SetFieldID(i)
	return fuo
}

// SetNillableFieldID sets the "field_id" field if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableFieldID(i *int) *FileUpdateOne {
	if i != nil {
		fuo.SetFieldID(*i)
	}
	return fuo
}

// AddFieldID adds i to the "field_id" field.
func (fuo *FileUpdateOne) AddFieldID(i int) *FileUpdateOne {
	fuo.mutation.AddFieldID(i)
	return fuo
}

// ClearFieldID clears the value of the "field_id" field.
func (fuo *FileUpdateOne) ClearFieldID() *FileUpdateOne {
	fuo.mutation.ClearFieldID()
	return fuo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (fuo *FileUpdateOne) SetOwnerID(id int) *FileUpdateOne {
	fuo.mutation.SetOwnerID(id)
	return fuo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableOwnerID(id *int) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetOwnerID(*id)
	}
	return fuo
}

// SetOwner sets the "owner" edge to the User entity.
func (fuo *FileUpdateOne) SetOwner(u *User) *FileUpdateOne {
	return fuo.SetOwnerID(u.ID)
}

// SetTypeID sets the "type" edge to the FileType entity by ID.
func (fuo *FileUpdateOne) SetTypeID(id int) *FileUpdateOne {
	fuo.mutation.SetTypeID(id)
	return fuo
}

// SetNillableTypeID sets the "type" edge to the FileType entity by ID if the given value is not nil.
func (fuo *FileUpdateOne) SetNillableTypeID(id *int) *FileUpdateOne {
	if id != nil {
		fuo = fuo.SetTypeID(*id)
	}
	return fuo
}

// SetType sets the "type" edge to the FileType entity.
func (fuo *FileUpdateOne) SetType(f *FileType) *FileUpdateOne {
	return fuo.SetTypeID(f.ID)
}

// AddFieldIDs adds the "field" edge to the FieldType entity by IDs.
func (fuo *FileUpdateOne) AddFieldIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.AddFieldIDs(ids...)
	return fuo
}

// AddField adds the "field" edges to the FieldType entity.
func (fuo *FileUpdateOne) AddField(f ...*FieldType) *FileUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.AddFieldIDs(ids...)
}

// Mutation returns the FileMutation object of the builder.
func (fuo *FileUpdateOne) Mutation() *FileMutation {
	return fuo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (fuo *FileUpdateOne) ClearOwner() *FileUpdateOne {
	fuo.mutation.ClearOwner()
	return fuo
}

// ClearType clears the "type" edge to the FileType entity.
func (fuo *FileUpdateOne) ClearType() *FileUpdateOne {
	fuo.mutation.ClearType()
	return fuo
}

// ClearFieldEdge clears all "field" edges to the FieldType entity.
func (fuo *FileUpdateOne) ClearFieldEdge() *FileUpdateOne {
	fuo.mutation.ClearFieldEdge()
	return fuo
}

// RemoveFieldIDs removes the "field" edge to FieldType entities by IDs.
func (fuo *FileUpdateOne) RemoveFieldIDs(ids ...int) *FileUpdateOne {
	fuo.mutation.RemoveFieldIDs(ids...)
	return fuo
}

// RemoveField removes "field" edges to FieldType entities.
func (fuo *FileUpdateOne) RemoveField(f ...*FieldType) *FileUpdateOne {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fuo.RemoveFieldIDs(ids...)
}

// Where appends a list predicates to the FileUpdate builder.
func (fuo *FileUpdateOne) Where(ps ...predicate.File) *FileUpdateOne {
	fuo.mutation.Where(ps...)
	return fuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fuo *FileUpdateOne) Select(field string, fields ...string) *FileUpdateOne {
	fuo.fields = append([]string{field}, fields...)
	return fuo
}

// Save executes the query and returns the updated File entity.
func (fuo *FileUpdateOne) Save(ctx context.Context) (*File, error) {
	return withHooks(ctx, fuo.sqlSave, fuo.mutation, fuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fuo *FileUpdateOne) SaveX(ctx context.Context) *File {
	node, err := fuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fuo *FileUpdateOne) Exec(ctx context.Context) error {
	_, err := fuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fuo *FileUpdateOne) ExecX(ctx context.Context) {
	if err := fuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fuo *FileUpdateOne) check() error {
	if v, ok := fuo.mutation.Size(); ok {
		if err := file.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "File.size": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (fuo *FileUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *FileUpdateOne {
	fuo.modifiers = append(fuo.modifiers, modifiers...)
	return fuo
}

func (fuo *FileUpdateOne) sqlSave(ctx context.Context) (_node *File, err error) {
	if err := fuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(file.Table, file.Columns, sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt))
	id, ok := fuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "File.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, file.FieldID)
		for _, f := range fields {
			if !file.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != file.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fuo.mutation.Size(); ok {
		_spec.SetField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedSize(); ok {
		_spec.AddField(file.FieldSize, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.Name(); ok {
		_spec.SetField(file.FieldName, field.TypeString, value)
	}
	if value, ok := fuo.mutation.User(); ok {
		_spec.SetField(file.FieldUser, field.TypeString, value)
	}
	if fuo.mutation.UserCleared() {
		_spec.ClearField(file.FieldUser, field.TypeString)
	}
	if value, ok := fuo.mutation.Group(); ok {
		_spec.SetField(file.FieldGroup, field.TypeString, value)
	}
	if fuo.mutation.GroupCleared() {
		_spec.ClearField(file.FieldGroup, field.TypeString)
	}
	if value, ok := fuo.mutation.GetOp(); ok {
		_spec.SetField(file.FieldOp, field.TypeBool, value)
	}
	if fuo.mutation.OpCleared() {
		_spec.ClearField(file.FieldOp, field.TypeBool)
	}
	if value, ok := fuo.mutation.FieldID(); ok {
		_spec.SetField(file.FieldFieldID, field.TypeInt, value)
	}
	if value, ok := fuo.mutation.AddedFieldID(); ok {
		_spec.AddField(file.FieldFieldID, field.TypeInt, value)
	}
	if fuo.mutation.FieldIDCleared() {
		_spec.ClearField(file.FieldFieldID, field.TypeInt)
	}
	if fuo.mutation.OwnerCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.OwnerTable,
			Columns: []string{file.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.TypeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.TypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   file.TypeTable,
			Columns: []string{file.TypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(filetype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fuo.mutation.FieldEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.RemovedFieldIDs(); len(nodes) > 0 && !fuo.mutation.FieldEdgeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fuo.mutation.FieldIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   file.FieldTable,
			Columns: []string{file.FieldColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fieldtype.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(fuo.modifiers...)
	_node = &File{config: fuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{file.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fuo.mutation.done = true
	return _node, nil
}
