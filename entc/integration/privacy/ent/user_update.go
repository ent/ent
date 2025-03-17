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
	"entgo.io/ent/entc/integration/privacy/ent/predicate"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/ent/team"
	"entgo.io/ent/entc/integration/privacy/ent/user"
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

// SetAge sets the "age" field.
func (_u *UserUpdate) SetAge(u uint) *UserUpdate {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(u)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdate) SetNillableAge(u *uint) *UserUpdate {
	if u != nil {
		_u.SetAge(*u)
	}
	return _u
}

// AddAge adds u to the "age" field.
func (_u *UserUpdate) AddAge(u int) *UserUpdate {
	_u.mutation.AddAge(u)
	return _u
}

// ClearAge clears the value of the "age" field.
func (_u *UserUpdate) ClearAge() *UserUpdate {
	_u.mutation.ClearAge()
	return _u
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (_u *UserUpdate) AddTeamIDs(ids ...int) *UserUpdate {
	_u.mutation.AddTeamIDs(ids...)
	return _u
}

// AddTeams adds the "teams" edges to the Team entity.
func (_u *UserUpdate) AddTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.AddTeamIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (_u *UserUpdate) AddTaskIDs(ids ...int) *UserUpdate {
	_u.mutation.AddTaskIDs(ids...)
	return _u
}

// AddTasks adds the "tasks" edges to the Task entity.
func (_u *UserUpdate) AddTasks(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.AddTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdate) Mutation() *UserMutation {
	return _u.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (_u *UserUpdate) ClearTeams() *UserUpdate {
	_u.mutation.ClearTeams()
	return _u
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (_u *UserUpdate) RemoveTeamIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveTeamIDs(ids...)
	return _u
}

// RemoveTeams removes "teams" edges to Team entities.
func (_u *UserUpdate) RemoveTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.RemoveTeamIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (_u *UserUpdate) ClearTasks() *UserUpdate {
	_u.mutation.ClearTasks()
	return _u
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (_u *UserUpdate) RemoveTaskIDs(ids ...int) *UserUpdate {
	_u.mutation.RemoveTaskIDs(ids...)
	return _u
}

// RemoveTasks removes "tasks" edges to Task entities.
func (_u *UserUpdate) RemoveTasks(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.RemoveTaskIDs(ids...)
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

func (_u *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if _u.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeUint)
	}
	if _u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !_u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTasksIDs(); len(nodes) > 0 && !_u.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetAge sets the "age" field.
func (_u *UserUpdateOne) SetAge(u uint) *UserUpdateOne {
	_u.mutation.ResetAge()
	_u.mutation.SetAge(u)
	return _u
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (_u *UserUpdateOne) SetNillableAge(u *uint) *UserUpdateOne {
	if u != nil {
		_u.SetAge(*u)
	}
	return _u
}

// AddAge adds u to the "age" field.
func (_u *UserUpdateOne) AddAge(u int) *UserUpdateOne {
	_u.mutation.AddAge(u)
	return _u
}

// ClearAge clears the value of the "age" field.
func (_u *UserUpdateOne) ClearAge() *UserUpdateOne {
	_u.mutation.ClearAge()
	return _u
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (_u *UserUpdateOne) AddTeamIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddTeamIDs(ids...)
	return _u
}

// AddTeams adds the "teams" edges to the Team entity.
func (_u *UserUpdateOne) AddTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.AddTeamIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (_u *UserUpdateOne) AddTaskIDs(ids ...int) *UserUpdateOne {
	_u.mutation.AddTaskIDs(ids...)
	return _u
}

// AddTasks adds the "tasks" edges to the Task entity.
func (_u *UserUpdateOne) AddTasks(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.AddTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (_u *UserUpdateOne) Mutation() *UserMutation {
	return _u.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (_u *UserUpdateOne) ClearTeams() *UserUpdateOne {
	_u.mutation.ClearTeams()
	return _u
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (_u *UserUpdateOne) RemoveTeamIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveTeamIDs(ids...)
	return _u
}

// RemoveTeams removes "teams" edges to Team entities.
func (_u *UserUpdateOne) RemoveTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.RemoveTeamIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (_u *UserUpdateOne) ClearTasks() *UserUpdateOne {
	_u.mutation.ClearTasks()
	return _u
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (_u *UserUpdateOne) RemoveTaskIDs(ids ...int) *UserUpdateOne {
	_u.mutation.RemoveTaskIDs(ids...)
	return _u
}

// RemoveTasks removes "tasks" edges to Task entities.
func (_u *UserUpdateOne) RemoveTasks(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _u.RemoveTaskIDs(ids...)
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

func (_u *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
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
	if value, ok := _u.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := _u.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if _u.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeUint)
	}
	if _u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !_u.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _u.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.RemovedTasksIDs(); len(nodes) > 0 && !_u.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := _u.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt),
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
