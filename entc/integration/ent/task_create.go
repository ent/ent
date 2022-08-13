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

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/schema/task"
	"entgo.io/ent/schema/field"

	enttask "entgo.io/ent/entc/integration/ent/task"
)

// TaskCreate is the builder for creating a Task entity.
type TaskCreate struct {
	config
	mutation *TaskMutation
	hooks    []Hook
	conflict []sql.ConflictOption
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
	var (
		err  error
		node *Task
	)
	tc.defaults()
	if len(tc.hooks) == 0 {
		if err = tc.check(); err != nil {
			return nil, err
		}
		node, err = tc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TaskMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = tc.check(); err != nil {
				return nil, err
			}
			tc.mutation = mutation
			if node, err = tc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(tc.hooks) - 1; i >= 0; i-- {
			if tc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, tc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Task)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from TaskMutation", v)
		}
		node = nv
	}
	return node, err
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

func (tc *TaskCreate) sqlSave(ctx context.Context) (*Task, error) {
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (tc *TaskCreate) createSpec() (*Task, *sqlgraph.CreateSpec) {
	var (
		_node = &Task{config: tc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: enttask.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: enttask.FieldID,
			},
		}
	)
	_spec.OnConflict = tc.conflict
	if value, ok := tc.mutation.Priority(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: enttask.FieldPriority,
		})
		_node.Priority = value
	}
	if value, ok := tc.mutation.Priorities(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: enttask.FieldPriorities,
		})
		_node.Priorities = value
	}
	if value, ok := tc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: enttask.FieldCreatedAt,
		})
		_node.CreatedAt = &value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.Create().
//		SetPriority(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetPriority(v+v).
//		}).
//		Exec(ctx)
func (tc *TaskCreate) OnConflict(opts ...sql.ConflictOption) *TaskUpsertOne {
	tc.conflict = opts
	return &TaskUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TaskCreate) OnConflictColumns(columns ...string) *TaskUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertOne{
		create: tc,
	}
}

type (
	// TaskUpsertOne is the builder for "upsert"-ing
	//  one Task node.
	TaskUpsertOne struct {
		create *TaskCreate
	}

	// TaskUpsert is the "OnConflict" setter.
	TaskUpsert struct {
		*sql.UpdateSet
	}
)

// SetPriority sets the "priority" field.
func (u *TaskUpsert) SetPriority(v task.Priority) *TaskUpsert {
	u.Set(enttask.FieldPriority, v)
	return u
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsert) UpdatePriority() *TaskUpsert {
	u.SetExcluded(enttask.FieldPriority)
	return u
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsert) AddPriority(v task.Priority) *TaskUpsert {
	u.Add(enttask.FieldPriority, v)
	return u
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsert) SetPriorities(v map[string]task.Priority) *TaskUpsert {
	u.Set(enttask.FieldPriorities, v)
	return u
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsert) UpdatePriorities() *TaskUpsert {
	u.SetExcluded(enttask.FieldPriorities)
	return u
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsert) ClearPriorities() *TaskUpsert {
	u.SetNull(enttask.FieldPriorities)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsert) SetCreatedAt(v time.Time) *TaskUpsert {
	u.Set(enttask.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsert) UpdateCreatedAt() *TaskUpsert {
	u.SetExcluded(enttask.FieldCreatedAt)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertOne) UpdateNewValues() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreatedAt(); exists {
			s.SetIgnore(enttask.FieldCreatedAt)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TaskUpsertOne) Ignore() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertOne) DoNothing() *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreate.OnConflict
// documentation for more info.
func (u *TaskUpsertOne) Update(set func(*TaskUpsert)) *TaskUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriority sets the "priority" field.
func (u *TaskUpsertOne) SetPriority(v task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsertOne) AddPriority(v task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdatePriority() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriority()
	})
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsertOne) SetPriorities(v map[string]task.Priority) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriorities(v)
	})
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdatePriorities() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriorities()
	})
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsertOne) ClearPriorities() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.ClearPriorities()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertOne) SetCreatedAt(v time.Time) *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertOne) UpdateCreatedAt() *TaskUpsertOne {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TaskUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TaskUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TaskCreateBulk is the builder for creating many Task entities in bulk.
type TaskCreateBulk struct {
	config
	builders []*TaskCreate
	conflict []sql.ConflictOption
}

// Save creates the Task entities in the database.
func (tcb *TaskCreateBulk) Save(ctx context.Context) ([]*Task, error) {
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Task, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TaskCreateBulk) SaveX(ctx context.Context) []*Task {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TaskCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TaskCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Task.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TaskUpsert) {
//			SetPriority(v+v).
//		}).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflict(opts ...sql.ConflictOption) *TaskUpsertBulk {
	tcb.conflict = opts
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TaskCreateBulk) OnConflictColumns(columns ...string) *TaskUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TaskUpsertBulk{
		create: tcb,
	}
}

// TaskUpsertBulk is the builder for "upsert"-ing
// a bulk of Task nodes.
type TaskUpsertBulk struct {
	create *TaskCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *TaskUpsertBulk) UpdateNewValues() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreatedAt(); exists {
				s.SetIgnore(enttask.FieldCreatedAt)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Task.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TaskUpsertBulk) Ignore() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TaskUpsertBulk) DoNothing() *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TaskCreateBulk.OnConflict
// documentation for more info.
func (u *TaskUpsertBulk) Update(set func(*TaskUpsert)) *TaskUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TaskUpsert{UpdateSet: update})
	}))
	return u
}

// SetPriority sets the "priority" field.
func (u *TaskUpsertBulk) SetPriority(v task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriority(v)
	})
}

// AddPriority adds v to the "priority" field.
func (u *TaskUpsertBulk) AddPriority(v task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.AddPriority(v)
	})
}

// UpdatePriority sets the "priority" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdatePriority() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriority()
	})
}

// SetPriorities sets the "priorities" field.
func (u *TaskUpsertBulk) SetPriorities(v map[string]task.Priority) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetPriorities(v)
	})
}

// UpdatePriorities sets the "priorities" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdatePriorities() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdatePriorities()
	})
}

// ClearPriorities clears the value of the "priorities" field.
func (u *TaskUpsertBulk) ClearPriorities() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.ClearPriorities()
	})
}

// SetCreatedAt sets the "created_at" field.
func (u *TaskUpsertBulk) SetCreatedAt(v time.Time) *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *TaskUpsertBulk) UpdateCreatedAt() *TaskUpsertBulk {
	return u.Update(func(s *TaskUpsert) {
		s.UpdateCreatedAt()
	})
}

// Exec executes the query.
func (u *TaskUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the TaskCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for TaskCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TaskUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
