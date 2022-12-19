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
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetAge sets the "age" field.
func (uu *UserUpdate) SetAge(u uint) *UserUpdate {
	uu.mutation.ResetAge()
	uu.mutation.SetAge(u)
	return uu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAge(u *uint) *UserUpdate {
	if u != nil {
		uu.SetAge(*u)
	}
	return uu
}

// AddAge adds u to the "age" field.
func (uu *UserUpdate) AddAge(u int) *UserUpdate {
	uu.mutation.AddAge(u)
	return uu
}

// ClearAge clears the value of the "age" field.
func (uu *UserUpdate) ClearAge() *UserUpdate {
	uu.mutation.ClearAge()
	return uu
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (uu *UserUpdate) AddTeamIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTeamIDs(ids...)
	return uu
}

// AddTeams adds the "teams" edges to the Team entity.
func (uu *UserUpdate) AddTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTeamIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (uu *UserUpdate) AddTaskIDs(ids ...int) *UserUpdate {
	uu.mutation.AddTaskIDs(ids...)
	return uu
}

// AddTasks adds the "tasks" edges to the Task entity.
func (uu *UserUpdate) AddTasks(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.AddTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (uu *UserUpdate) ClearTeams() *UserUpdate {
	uu.mutation.ClearTeams()
	return uu
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (uu *UserUpdate) RemoveTeamIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTeamIDs(ids...)
	return uu
}

// RemoveTeams removes "teams" edges to Team entities.
func (uu *UserUpdate) RemoveTeams(t ...*Team) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTeamIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (uu *UserUpdate) ClearTasks() *UserUpdate {
	uu.mutation.ClearTasks()
	return uu
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (uu *UserUpdate) RemoveTaskIDs(ids ...int) *UserUpdate {
	uu.mutation.RemoveTaskIDs(ids...)
	return uu
}

// RemoveTasks removes "tasks" edges to Task entities.
func (uu *UserUpdate) RemoveTasks(t ...*Task) *UserUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uu.RemoveTaskIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, UserMutation](ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uu.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if uu.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeUint)
	}
	if uu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !uu.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.RemovedTasksIDs(); len(nodes) > 0 && !uu.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
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
func (uuo *UserUpdateOne) SetAge(u uint) *UserUpdateOne {
	uuo.mutation.ResetAge()
	uuo.mutation.SetAge(u)
	return uuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAge(u *uint) *UserUpdateOne {
	if u != nil {
		uuo.SetAge(*u)
	}
	return uuo
}

// AddAge adds u to the "age" field.
func (uuo *UserUpdateOne) AddAge(u int) *UserUpdateOne {
	uuo.mutation.AddAge(u)
	return uuo
}

// ClearAge clears the value of the "age" field.
func (uuo *UserUpdateOne) ClearAge() *UserUpdateOne {
	uuo.mutation.ClearAge()
	return uuo
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (uuo *UserUpdateOne) AddTeamIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTeamIDs(ids...)
	return uuo
}

// AddTeams adds the "teams" edges to the Team entity.
func (uuo *UserUpdateOne) AddTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTeamIDs(ids...)
}

// AddTaskIDs adds the "tasks" edge to the Task entity by IDs.
func (uuo *UserUpdateOne) AddTaskIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.AddTaskIDs(ids...)
	return uuo
}

// AddTasks adds the "tasks" edges to the Task entity.
func (uuo *UserUpdateOne) AddTasks(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.AddTaskIDs(ids...)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearTeams clears all "teams" edges to the Team entity.
func (uuo *UserUpdateOne) ClearTeams() *UserUpdateOne {
	uuo.mutation.ClearTeams()
	return uuo
}

// RemoveTeamIDs removes the "teams" edge to Team entities by IDs.
func (uuo *UserUpdateOne) RemoveTeamIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTeamIDs(ids...)
	return uuo
}

// RemoveTeams removes "teams" edges to Team entities.
func (uuo *UserUpdateOne) RemoveTeams(t ...*Team) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTeamIDs(ids...)
}

// ClearTasks clears all "tasks" edges to the Task entity.
func (uuo *UserUpdateOne) ClearTasks() *UserUpdateOne {
	uuo.mutation.ClearTasks()
	return uuo
}

// RemoveTaskIDs removes the "tasks" edge to Task entities by IDs.
func (uuo *UserUpdateOne) RemoveTaskIDs(ids ...int) *UserUpdateOne {
	uuo.mutation.RemoveTaskIDs(ids...)
	return uuo
}

// RemoveTasks removes "tasks" edges to Task entities.
func (uuo *UserUpdateOne) RemoveTasks(t ...*Task) *UserUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return uuo.RemoveTaskIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return withHooks[*User, UserMutation](ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
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
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.Age(); ok {
		_spec.SetField(user.FieldAge, field.TypeUint, value)
	}
	if value, ok := uuo.mutation.AddedAge(); ok {
		_spec.AddField(user.FieldAge, field.TypeUint, value)
	}
	if uuo.mutation.AgeCleared() {
		_spec.ClearField(user.FieldAge, field.TypeUint)
	}
	if uuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTeamsIDs(); len(nodes) > 0 && !uuo.mutation.TeamsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.TeamsTable,
			Columns: user.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: team.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if uuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.RemovedTasksIDs(); len(nodes) > 0 && !uuo.mutation.TasksCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.TasksIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TasksTable,
			Columns: []string{user.TasksColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: task.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}
