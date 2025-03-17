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
	"entgo.io/ent/examples/privacytenant/ent/group"
	"entgo.io/ent/examples/privacytenant/ent/predicate"
	"entgo.io/ent/examples/privacytenant/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks    []Hook
	mutation *GroupMutation
}

// Where appends a list predicates to the GroupUpdate builder.
func (u *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetName sets the "name" field.
func (m *GroupUpdate) SetName(v string) *GroupUpdate {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *GroupUpdate) SetNillableName(v *string) *GroupUpdate {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (m *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	m.mutation.AddUserIDs(ids...)
	return m
}

// AddUsers adds the "users" edges to the User entity.
func (m *GroupUpdate) AddUsers(v ...*User) *GroupUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (m *GroupUpdate) Mutation() *GroupMutation {
	return m.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (u *GroupUpdate) ClearUsers() *GroupUpdate {
	u.mutation.ClearUsers()
	return u
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (u *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	u.mutation.RemoveUserIDs(ids...)
	return u
}

// RemoveUsers removes "users" edges to User entities.
func (u *GroupUpdate) RemoveUsers(v ...*User) *GroupUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *GroupUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *GroupUpdate) check() error {
	if u.mutation.TenantCleared() && len(u.mutation.TenantIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Group.tenant"`)
	}
	return nil
}

func (u *GroupUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedUsersIDs(); len(nodes) > 0 && !u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupMutation
}

// SetName sets the "name" field.
func (m *GroupUpdateOne) SetName(v string) *GroupUpdateOne {
	m.mutation.SetName(v)
	return m
}

// SetNillableName sets the "name" field if the given value is not nil.
func (m *GroupUpdateOne) SetNillableName(v *string) *GroupUpdateOne {
	if v != nil {
		m.SetName(*v)
	}
	return m
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (m *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	m.mutation.AddUserIDs(ids...)
	return m
}

// AddUsers adds the "users" edges to the User entity.
func (m *GroupUpdateOne) AddUsers(v ...*User) *GroupUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (m *GroupUpdateOne) Mutation() *GroupMutation {
	return m.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (u *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	u.mutation.ClearUsers()
	return u
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (u *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	u.mutation.RemoveUserIDs(ids...)
	return u
}

// RemoveUsers removes "users" edges to User entities.
func (u *GroupUpdateOne) RemoveUsers(v ...*User) *GroupUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (u *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Group entity.
func (u *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *GroupUpdateOne) check() error {
	if u.mutation.TenantCleared() && len(u.mutation.TenantIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Group.tenant"`)
	}
	return nil
}

func (u *GroupUpdateOne) sqlSave(ctx context.Context) (_n *Group, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != group.FieldID {
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
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedUsersIDs(); len(nodes) > 0 && !u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &Group{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
