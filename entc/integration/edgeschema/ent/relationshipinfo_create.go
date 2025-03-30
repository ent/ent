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
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/schema/field"
)

// RelationshipInfoCreate is the builder for creating a RelationshipInfo entity.
type RelationshipInfoCreate struct {
	config
	mutation *RelationshipInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetText sets the "text" field.
func (_c *RelationshipInfoCreate) SetText(v string) *RelationshipInfoCreate {
	_c.mutation.SetText(v)
	return _c
}

// Mutation returns the RelationshipInfoMutation object of the builder.
func (_c *RelationshipInfoCreate) Mutation() *RelationshipInfoMutation {
	return _c.mutation
}

// Save creates the RelationshipInfo in the database.
func (_c *RelationshipInfoCreate) Save(ctx context.Context) (*RelationshipInfo, error) {
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *RelationshipInfoCreate) SaveX(ctx context.Context) *RelationshipInfo {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *RelationshipInfoCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *RelationshipInfoCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *RelationshipInfoCreate) check() error {
	if _, ok := _c.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "RelationshipInfo.text"`)}
	}
	return nil
}

func (_c *RelationshipInfoCreate) sqlSave(ctx context.Context) (*RelationshipInfo, error) {
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

func (_c *RelationshipInfoCreate) createSpec() (*RelationshipInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &RelationshipInfo{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(relationshipinfo.Table, sqlgraph.NewFieldSpec(relationshipinfo.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.Text(); ok {
		_spec.SetField(relationshipinfo.FieldText, field.TypeString, value)
		_node.Text = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RelationshipInfo.Create().
//		SetText(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RelationshipInfoUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (_c *RelationshipInfoCreate) OnConflict(opts ...sql.ConflictOption) *RelationshipInfoUpsertOne {
	_c.conflict = opts
	return &RelationshipInfoUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *RelationshipInfoCreate) OnConflictColumns(columns ...string) *RelationshipInfoUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &RelationshipInfoUpsertOne{
		create: _c,
	}
}

type (
	// RelationshipInfoUpsertOne is the builder for "upsert"-ing
	//  one RelationshipInfo node.
	RelationshipInfoUpsertOne struct {
		create *RelationshipInfoCreate
	}

	// RelationshipInfoUpsert is the "OnConflict" setter.
	RelationshipInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetText sets the "text" field.
func (u *RelationshipInfoUpsert) SetText(v string) *RelationshipInfoUpsert {
	u.Set(relationshipinfo.FieldText, v)
	return u
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *RelationshipInfoUpsert) UpdateText() *RelationshipInfoUpsert {
	u.SetExcluded(relationshipinfo.FieldText)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RelationshipInfoUpsertOne) UpdateNewValues() *RelationshipInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *RelationshipInfoUpsertOne) Ignore() *RelationshipInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RelationshipInfoUpsertOne) DoNothing() *RelationshipInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RelationshipInfoCreate.OnConflict
// documentation for more info.
func (u *RelationshipInfoUpsertOne) Update(set func(*RelationshipInfoUpsert)) *RelationshipInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RelationshipInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *RelationshipInfoUpsertOne) SetText(v string) *RelationshipInfoUpsertOne {
	return u.Update(func(s *RelationshipInfoUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *RelationshipInfoUpsertOne) UpdateText() *RelationshipInfoUpsertOne {
	return u.Update(func(s *RelationshipInfoUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *RelationshipInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RelationshipInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RelationshipInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *RelationshipInfoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *RelationshipInfoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// RelationshipInfoCreateBulk is the builder for creating many RelationshipInfo entities in bulk.
type RelationshipInfoCreateBulk struct {
	config
	err      error
	builders []*RelationshipInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the RelationshipInfo entities in the database.
func (_c *RelationshipInfoCreateBulk) Save(ctx context.Context) ([]*RelationshipInfo, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*RelationshipInfo, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RelationshipInfoMutation)
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
					spec.OnConflict = _c.conflict
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
func (_c *RelationshipInfoCreateBulk) SaveX(ctx context.Context) []*RelationshipInfo {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *RelationshipInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *RelationshipInfoCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.RelationshipInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.RelationshipInfoUpsert) {
//			SetText(v+v).
//		}).
//		Exec(ctx)
func (_c *RelationshipInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *RelationshipInfoUpsertBulk {
	_c.conflict = opts
	return &RelationshipInfoUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *RelationshipInfoCreateBulk) OnConflictColumns(columns ...string) *RelationshipInfoUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &RelationshipInfoUpsertBulk{
		create: _c,
	}
}

// RelationshipInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of RelationshipInfo nodes.
type RelationshipInfoUpsertBulk struct {
	create *RelationshipInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *RelationshipInfoUpsertBulk) UpdateNewValues() *RelationshipInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.RelationshipInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *RelationshipInfoUpsertBulk) Ignore() *RelationshipInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *RelationshipInfoUpsertBulk) DoNothing() *RelationshipInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the RelationshipInfoCreateBulk.OnConflict
// documentation for more info.
func (u *RelationshipInfoUpsertBulk) Update(set func(*RelationshipInfoUpsert)) *RelationshipInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&RelationshipInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetText sets the "text" field.
func (u *RelationshipInfoUpsertBulk) SetText(v string) *RelationshipInfoUpsertBulk {
	return u.Update(func(s *RelationshipInfoUpsert) {
		s.SetText(v)
	})
}

// UpdateText sets the "text" field to the value that was provided on create.
func (u *RelationshipInfoUpsertBulk) UpdateText() *RelationshipInfoUpsertBulk {
	return u.Update(func(s *RelationshipInfoUpsert) {
		s.UpdateText()
	})
}

// Exec executes the query.
func (u *RelationshipInfoUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if b == nil {
			return fmt.Errorf("ent: missing builder at index %d, unexpected nil builder passed to CreateBulk", i)
		}
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the RelationshipInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for RelationshipInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *RelationshipInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
