// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/ent/schema/task"

	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
}

// SetPriority sets the "priority" field.
func (tc *TaskCreate) SetPriority(t task.Priority) *TaskCreate {
	tc.mutation.SetPriority(t)
	return tc
}

// SetNillablePriority sets the "priority" field if the given value is not nil.
func (tc *TaskCreate) SetNillablePriority(t *task.Priority) *TaskCreate {
	if t != nil {
		tc.SetPriority(*t)
	}
	return tc
}

// SetPriorities sets the "priorities" field.
func (tc *TaskCreate) SetPriorities(m map[string]task.Priority) *TaskCreate {
	tc.mutation.SetPriorities(m)
	return tc
}

// SetCreatedAt sets the "created_at" field.
func (tc *TaskCreate) SetCreatedAt(t time.Time) *TaskCreate {
	tc.mutation.SetCreatedAt(t)
	return tc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (tc *TaskCreate) SetNillableCreatedAt(t *time.Time) *TaskCreate {
	if t != nil {
		tc.SetCreatedAt(*t)
	}
	return tc
}

// Mutation returns the TaskMutation object of the builder.
func (tc *TaskCreate) Mutation() *TaskMutation {
	return tc.mutation
}

// Save creates the Task in the database.
func (tc *TaskCreate) Save(ctx context.Context) (*Task, error) {
	tc.defaults()
	return withHooks[*Task, TaskMutation](ctx, tc.gremlinSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TaskCreate) SaveX(ctx context.Context) *Task {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TaskCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TaskCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TaskCreate) defaults() {
	if _, ok := tc.mutation.Priority(); !ok {
		v := enttask.DefaultPriority
		tc.mutation.SetPriority(v)
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		v := enttask.DefaultCreatedAt()
		tc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (tc *TaskCreate) check() error {
	if _, ok := tc.mutation.Priority(); !ok {
		return &ValidationError{Name: "priority", err: errors.New(`ent: missing required field "Task.priority"`)}
	}
	if v, ok := tc.mutation.Priority(); ok {
		if err := enttask.PriorityValidator(int(v)); err != nil {
			return &ValidationError{Name: "priority", err: fmt.Errorf(`ent: validator failed for field "Task.priority": %w`, err)}
		}
	}
	if _, ok := tc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Task.created_at"`)}
	}
	return nil
}

func (tc *TaskCreate) gremlinSave(ctx context.Context) (*Task, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := tc.gremlin().Query()
	if err := tc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	t := &Task{config: tc.config}
	if err := t.FromResponse(res); err != nil {
		return nil, err
	}
	tc.mutation.id = &t.ID
	tc.mutation.done = true
	return t, nil
}

func (tc *TaskCreate) gremlin() *dsl.Traversal {
	v := g.AddV(enttask.Label)
	if value, ok := tc.mutation.Priority(); ok {
		v.Property(dsl.Single, enttask.FieldPriority, value)
	}
	if value, ok := tc.mutation.Priorities(); ok {
		v.Property(dsl.Single, enttask.FieldPriorities, value)
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		v.Property(dsl.Single, enttask.FieldCreatedAt, value)
	}
	return v.ValueMap(true)
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
}
