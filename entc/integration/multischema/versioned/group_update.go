// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package versioned

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/multischema/versioned/group"
	"entgo.io/ent/entc/integration/multischema/versioned/internal"
	"entgo.io/ent/entc/integration/multischema/versioned/predicate"
	"entgo.io/ent/entc/integration/multischema/versioned/user"
	"entgo.io/ent/schema/field"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	hooks     []Hook
	mutation  *GroupMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the GroupUpdate builder.
func (_u *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetName sets the "name" field.
func (_u *GroupUpdate) SetName(s string) *GroupUpdate {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *GroupUpdate) SetNillableName(s *string) *GroupUpdate {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (_u *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	_u.mutation.AddUserIDs(ids...)
	return _u
}

// AddUsers adds the "users" edges to the User entity.
func (_u *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (_u *GroupUpdate) Mutation() *GroupMutation {
	return _u.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (_u *GroupUpdate) ClearUsers() *GroupUpdate {
	_u.mutation.ClearUsers()
	return _u
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (_u *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	_u.mutation.RemoveUserIDs(ids...)
	return _u
}

// RemoveUsers removes "users" edges to User entities.
func (_u *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.RemoveUserIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *GroupUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *GroupUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *GroupUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *GroupUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupUpdate {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if _u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedUsersIDs(); len(nodes) > 0 && !_u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.Group
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *GroupMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetName sets the "name" field.
func (_u *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *GroupUpdateOne) SetNillableName(s *string) *GroupUpdateOne {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (_u *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	_u.mutation.AddUserIDs(ids...)
	return _u
}

// AddUsers adds the "users" edges to the User entity.
func (_u *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (_u *GroupUpdateOne) Mutation() *GroupMutation {
	return _u.mutation
}

// ClearUsers clears all "users" edges to the User entity.
func (_u *GroupUpdateOne) ClearUsers() *GroupUpdateOne {
	_u.mutation.ClearUsers()
	return _u
}

// RemoveUserIDs removes the "users" edge to User entities by IDs.
func (_u *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	_u.mutation.RemoveUserIDs(ids...)
	return _u
}

// RemoveUsers removes "users" edges to User entities.
func (_u *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _u.RemoveUserIDs(ids...)
}

// Where appends a list predicates to the GroupUpdate builder.
func (_u *GroupUpdateOne) Where(ps ...predicate.Group) *GroupUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *GroupUpdateOne) Select(field string, fields ...string) *GroupUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Group entity.
func (_u *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (_u *GroupUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *GroupUpdateOne {
	_u.modifiers = append(_u.modifiers, modifiers...)
	return _u
}

func (_u *GroupUpdateOne) sqlSave(ctx context.Context) (_node *Group, err error) {
	_spec := sqlgraph.NewUpdateSpec(group.Table, group.Columns, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`versioned: missing "Group.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, group.FieldID)
		for _, f := range fields {
			if !group.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("versioned: invalid field %q for query", f)}
			}
			if f != group.FieldID {
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
		_spec.SetField(group.FieldName, field.TypeString, value)
	}
	if _u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedUsersIDs(); len(nodes) > 0 && !_u.mutation.UsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		edge.Schema = _u.schemaConfig.GroupUsers
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Node.Schema = _u.schemaConfig.Group
	ctx = internal.NewSchemaConfigContext(ctx, _u.schemaConfig)
	_spec.AddModifiers(_u.modifiers...)
	_node = &Group{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{group.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
