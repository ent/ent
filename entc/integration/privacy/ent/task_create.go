// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/privacy/ent/task"
	"entgo.io/ent/entc/integration/privacy/ent/team"
	"entgo.io/ent/entc/integration/privacy/ent/user"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetTitle sets the "title" field.
func (_c *TaskCreate) SetTitle(v string) *TaskCreate {
	_c.mutation.SetTitle(v)
	return _c
}

// SetDescription sets the "description" field.
func (_c *TaskCreate) SetDescription(v string) *TaskCreate {
	_c.mutation.SetDescription(v)
	return _c
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (_c *TaskCreate) SetNillableDescription(v *string) *TaskCreate {
	if v != nil {
		_c.SetDescription(*v)
	}
	return _c
}

// SetStatus sets the "status" field.
func (_c *TaskCreate) SetStatus(v task.Status) *TaskCreate {
	_c.mutation.SetStatus(v)
	return _c
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (_c *TaskCreate) SetNillableStatus(v *task.Status) *TaskCreate {
	if v != nil {
		_c.SetStatus(*v)
	}
	return _c
}

// SetUUID sets the "uuid" field.
func (_c *TaskCreate) SetUUID(v uuid.UUID) *TaskCreate {
	_c.mutation.SetUUID(v)
	return _c
}

// SetNillableUUID sets the "uuid" field if the given value is not nil.
func (_c *TaskCreate) SetNillableUUID(v *uuid.UUID) *TaskCreate {
	if v != nil {
		_c.SetUUID(*v)
	}
	return _c
}

// AddTeamIDs adds the "teams" edge to the Team entity by IDs.
func (_c *TaskCreate) AddTeamIDs(ids ...int) *TaskCreate {
	_c.mutation.AddTeamIDs(ids...)
	return _c
}

// AddTeams adds the "teams" edges to the Team entity.
func (_c *TaskCreate) AddTeams(v ...*Team) *TaskCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return _c.AddTeamIDs(ids...)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (_c *TaskCreate) SetOwnerID(id int) *TaskCreate {
	_c.mutation.SetOwnerID(id)
	return _c
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (_c *TaskCreate) SetNillableOwnerID(id *int) *TaskCreate {
	if id != nil {
		_c = _c.SetOwnerID(*id)
	}
	return _c
}

// SetOwner sets the "owner" edge to the User entity.
func (_c *TaskCreate) SetOwner(v *User) *TaskCreate {
	return _c.SetOwnerID(v.ID)
}

// Mutation returns the TaskMutation object of the builder.
func (_c *TaskCreate) Mutation() *TaskMutation {
	return _c.mutation
}

// Save creates the Task in the database.
func (_c *TaskCreate) Save(ctx context.Context) (*Task, error) {
	if err := _c.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *TaskCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *TaskCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *TaskCreate) defaults() error {
	if _, ok := _c.mutation.Status(); !ok {
		v := task.DefaultStatus
		_c.mutation.SetStatus(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (_c *TaskCreate) check() error {
	if _, ok := _c.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Task.title"`)}
	}
	if v, ok := _c.mutation.Title(); ok {
		if err := task.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Task.title": %w`, err)}
		}
	}
	if _, ok := _c.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`ent: missing required field "Task.status"`)}
	}
	if v, ok := _c.mutation.Status(); ok {
		if err := task.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`ent: validator failed for field "Task.status": %w`, err)}
		}
	}
	return nil
}

func (_c *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	_node, _spec := _c.createSpec()
	if err := sqlgraph.CreateNode(ctx, _c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	_c.mutation.id = &_node.ID
	_c.mutation.done = true
	return _node, nil
}

func (_c *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(task.Table, sqlgraph.NewFieldSpec(task.FieldID, field.TypeInt))
	)
	if value, ok := _c.mutation.Title(); ok {
		_spec.SetField(task.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := _c.mutation.Description(); ok {
		_spec.SetField(task.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := _c.mutation.Status(); ok {
		_spec.SetField(task.FieldStatus, field.TypeEnum, value)
		_node.Status = value
	}
	if value, ok := _c.mutation.UUID(); ok {
		_spec.SetField(task.FieldUUID, field.TypeUUID, value)
		_node.UUID = value
	}
	if nodes := _c.mutation.TeamsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   task.TeamsTable,
			Columns: task.TeamsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(team.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   task.OwnerTable,
			Columns: []string{task.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_tasks = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	err      error
	builders []*TaskCreate
}

// Save creates the Task entities in the database.
func (_c *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Task, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TaskMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, _c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, _c.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, _c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (_c *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}
