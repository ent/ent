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
	"entgo.io/ent/entc/integration/edgeschema/ent/attachedfile"
	"entgo.io/ent/entc/integration/edgeschema/ent/file"
	"entgo.io/ent/entc/integration/edgeschema/ent/process"
	"entgo.io/ent/schema/field"
)

// ProcessCreate is the builder for creating a Process entity.
type ProcessCreate struct {
	config
	mutation *ProcessMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (m *ProcessCreate) AddFileIDs(ids ...int) *ProcessCreate {
	m.mutation.AddFileIDs(ids...)
	return m
}

// AddFiles adds the "files" edges to the File entity.
func (m *ProcessCreate) AddFiles(v ...*File) *ProcessCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddFileIDs(ids...)
}

// AddAttachedFileIDs adds the "attached_files" edge to the AttachedFile entity by IDs.
func (m *ProcessCreate) AddAttachedFileIDs(ids ...int) *ProcessCreate {
	m.mutation.AddAttachedFileIDs(ids...)
	return m
}

// AddAttachedFiles adds the "attached_files" edges to the AttachedFile entity.
func (m *ProcessCreate) AddAttachedFiles(v ...*AttachedFile) *ProcessCreate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddAttachedFileIDs(ids...)
}

// Mutation returns the ProcessMutation object of the builder.
func (m *ProcessCreate) Mutation() *ProcessMutation {
	return m.mutation
}

// Save creates the Process in the database.
func (c *ProcessCreate) Save(ctx context.Context) (*Process, error) {
	return withHooks(ctx, c.sqlSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *ProcessCreate) SaveX(ctx context.Context) *Process {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *ProcessCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *ProcessCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *ProcessCreate) check() error {
	return nil
}

func (c *ProcessCreate) sqlSave(ctx context.Context) (*Process, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	_node, _spec := c.createSpec()
	if err := sqlgraph.CreateNode(ctx, c.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	c.mutation.id = &_node.ID
	c.mutation.done = true
	return _node, nil
}

func (c *ProcessCreate) createSpec() (*Process, *sqlgraph.CreateSpec) {
	var (
		_node = &Process{config: c.config}
		_spec = sqlgraph.NewCreateSpec(process.Table, sqlgraph.NewFieldSpec(process.FieldID, field.TypeInt))
	)
	_spec.OnConflict = c.conflict
	if nodes := c.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   process.FilesTable,
			Columns: process.FilesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &AttachedFileCreate{config: c.config, mutation: newAttachedFileMutation(c.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := c.mutation.AttachedFilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   process.AttachedFilesTable,
			Columns: []string{process.AttachedFilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(attachedfile.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Process.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (c *ProcessCreate) OnConflict(opts ...sql.ConflictOption) *ProcessUpsertOne {
	c.conflict = opts
	return &ProcessUpsertOne{create: c}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Process.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (c *ProcessCreate) OnConflictColumns(columns ...string) *ProcessUpsertOne {
	c.conflict = append(c.conflict, sql.ConflictColumns(columns...))
	return &ProcessUpsertOne{create: c}
}

type (
	// ProcessUpsertOne is the builder for "upsert"-ing
	//  one Process node.
	ProcessUpsertOne struct {
		create *ProcessCreate
	}

	// ProcessUpsert is the "OnConflict" setter.
	ProcessUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Process.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProcessUpsertOne) UpdateNewValues() *ProcessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Process.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProcessUpsertOne) Ignore() *ProcessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProcessUpsertOne) DoNothing() *ProcessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProcessCreate.OnConflict
// documentation for more info.
func (u *ProcessUpsertOne) Update(set func(*ProcessUpsert)) *ProcessUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProcessUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ProcessUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProcessCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProcessUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProcessUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProcessUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProcessCreateBulk is the builder for creating many Process entities in bulk.
type ProcessCreateBulk struct {
	config
	err      error
	builders []*ProcessCreate
	conflict []sql.ConflictOption
}

// Save creates the Process entities in the database.
func (c *ProcessCreateBulk) Save(ctx context.Context) ([]*Process, error) {
	if c.err != nil {
		return nil, c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(c.builders))
	nodes := make([]*Process, len(c.builders))
	mutators := make([]Mutator, len(c.builders))
	for i := range c.builders {
		func(i int, root context.Context) {
			builder := c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProcessMutation)
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
					_, err = mutators[i+1].Mutate(root, c.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = c.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, c.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, c.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (c *ProcessCreateBulk) SaveX(ctx context.Context) []*Process {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *ProcessCreateBulk) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *ProcessCreateBulk) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Process.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (c *ProcessCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProcessUpsertBulk {
	c.conflict = opts
	return &ProcessUpsertBulk{create: c}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Process.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (c *ProcessCreateBulk) OnConflictColumns(columns ...string) *ProcessUpsertBulk {
	c.conflict = append(c.conflict, sql.ConflictColumns(columns...))
	return &ProcessUpsertBulk{create: c}
}

// ProcessUpsertBulk is the builder for "upsert"-ing
// a bulk of Process nodes.
type ProcessUpsertBulk struct {
	create *ProcessCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Process.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *ProcessUpsertBulk) UpdateNewValues() *ProcessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Process.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProcessUpsertBulk) Ignore() *ProcessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProcessUpsertBulk) DoNothing() *ProcessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProcessCreateBulk.OnConflict
// documentation for more info.
func (u *ProcessUpsertBulk) Update(set func(*ProcessUpsert)) *ProcessUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProcessUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ProcessUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ProcessCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ProcessCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProcessUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
