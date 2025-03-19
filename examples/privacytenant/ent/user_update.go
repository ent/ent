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
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/examples/privacytenant/ent/group"
	"entgo.io/ent/examples/privacytenant/ent/predicate"
	"entgo.io/ent/examples/privacytenant/ent/user"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (_u *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetName sets the "name" field.
func (_u *UserUpdate) SetName(v string) *UserUpdate {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdate) SetNillableName(v *string) *UserUpdate {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetFoods sets the "foods" field.
func (_u *UserUpdate) SetFoods(v []string) *UserUpdate {
	_u.mutation.SetFoods(v)
	return _u
}

// AppendFoods appends value to the "foods" field.
func (_u *UserUpdate) AppendFoods(v []string) *UserUpdate {
	_u.mutation.AppendFoods(v)
	return _u
}

// ClearFoods clears the value of the "foods" field.
func (_u *UserUpdate) ClearFoods() *UserUpdate {
	_u.mutation.ClearFoods()
	return _u
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *UserUpdate) AddGroups(v ...*Group) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdate) Mutation() *UserMutation {
	return _u.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *UserUpdate) ClearGroups() *UserUpdate {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *UserUpdate) RemoveGroups(v ...*Group) *UserUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *UserUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdate) check() error {
	if _u.mutation.TenantCleared() && len(_u.mutation.TenantIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "User.tenant"`)
	}
	return nil
}

func (_u *UserUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Foods(); ok {
		_spec.SetField(user.FieldFoods, field.TypeJSON, value)
	}
	if value, ok := _u.mutation.AppendedFoods(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFoods, value)
		})
	}
	if _u.mutation.FoodsCleared() {
		_spec.ClearField(user.FieldFoods, field.TypeJSON)
	}
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetName sets the "name" field.
func (_u *UserUpdateOne) SetName(v string) *UserUpdateOne {
	_u.mutation.SetName(v)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableName(v *string) *UserUpdateOne {
	if v != nil {
		_u.SetName(*v)
	}
	return _u
}

// SetFoods sets the "foods" field.
func (_u *UserUpdateOne) SetFoods(v []string) *UserUpdateOne {
	_u.mutation.SetFoods(v)
	return _u
}

// AppendFoods appends value to the "foods" field.
func (_u *UserUpdateOne) AppendFoods(v []string) *UserUpdateOne {
	_u.mutation.AppendFoods(v)
	return _u
}

// ClearFoods clears the value of the "foods" field.
func (_u *UserUpdateOne) ClearFoods() *UserUpdateOne {
	_u.mutation.ClearFoods()
	return _u
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (_u *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddGroupIDs(ids...)
	return _u
}

// AddGroups adds the "groups" edges to the Group entity.
func (_u *UserUpdateOne) AddGroups(v ...*Group) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.AddGroupIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdateOne) Mutation() *UserMutation {
	return _u.mutation
}

// ClearGroups clears all "groups" edges to the Group entity.
func (_u *UserUpdateOne) ClearGroups() *UserUpdateOne {
	_u.mutation.ClearGroups()
	return _u
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (_u *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveGroupIDs(ids...)
	return _u
}

// RemoveGroups removes "groups" edges to Group entities.
func (_u *UserUpdateOne) RemoveGroups(v ...*Group) *UserUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _u.RemoveGroupIDs(ids...)
}

// Where appends a list predicates to the UserUpdate builder.
func (_u *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated User entity.
func (_u *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *UserUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_u *UserUpdateOne) check() error {
	if _u.mutation.TenantCleared() && len(_u.mutation.TenantIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "User.tenant"`)
	}
	return nil
}

func (_u *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := _u.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(user.FieldName, field.TypeString, value)
	}
	if value, ok := _u.mutation.Foods(); ok {
		_spec.SetField(user.FieldFoods, field.TypeJSON, value)
	}
	if value, ok := _u.mutation.AppendedFoods(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFoods, value)
		})
	}
	if _u.mutation.FoodsCleared() {
		_spec.ClearField(user.FieldFoods, field.TypeJSON)
	}
	if _u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !_u.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.GroupsTable,
			Columns: user.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
