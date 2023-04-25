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
	"entgo.io/ent/entc/integration/ent/license"
	"entgo.io/ent/schema/field"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (lc *LicenseCreate) SetCreateTime(t time.Time) *LicenseCreate {
	lc.mutation.SetCreateTime(t)
	return lc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableCreateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetCreateTime(*t)
	}
	return lc
}

// SetUpdateTime sets the "update_time" field.
func (lc *LicenseCreate) SetUpdateTime(t time.Time) *LicenseCreate {
	lc.mutation.SetUpdateTime(t)
	return lc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (lc *LicenseCreate) SetNillableUpdateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		lc.SetUpdateTime(*t)
	}
	return lc
}

// SetID sets the "id" field.
func (lc *LicenseCreate) SetID(i int) *LicenseCreate {
	lc.mutation.SetID(i)
	return lc
}

// Mutation returns the LicenseMutation object of the builder.
func (lc *LicenseCreate) Mutation() *LicenseMutation {
	return lc.mutation
}

// Save creates the License in the database.
func (lc *LicenseCreate) Save(ctx context.Context) (*License, error) {
	lc.defaults()
	return withHooks(ctx, lc.sqlSave, lc.mutation, lc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (lc *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := lc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lc *LicenseCreate) Exec(ctx context.Context) error {
	_, err := lc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lc *LicenseCreate) ExecX(ctx context.Context) {
	if err := lc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (lc *LicenseCreate) defaults() {
	if _, ok := lc.mutation.CreateTime(); !ok {
		v := license.DefaultCreateTime()
		lc.mutation.SetCreateTime(v)
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		v := license.DefaultUpdateTime()
		lc.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (lc *LicenseCreate) check() error {
	if _, ok := lc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "License.create_time"`)}
	}
	if _, ok := lc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "License.update_time"`)}
	}
	return nil
}

func (lc *LicenseCreate) sqlSave(ctx context.Context) (*License, error) {
	if err := lc.check(); err != nil {
		return nil, err
	}
	_node, _spec := lc.createSpec()
	if err := sqlgraph.CreateNode(ctx, lc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	lc.mutation.id = &_node.ID
	lc.mutation.done = true
	return _node, nil
}

func (lc *LicenseCreate) createSpec() (*License, *sqlgraph.CreateSpec) {
	var (
		_node = &License{config: lc.config}
		_spec = sqlgraph.NewCreateSpec(license.Table, sqlgraph.NewFieldSpec(license.FieldID, field.TypeInt))
	)
	_spec.OnConflict = lc.conflict
	if id, ok := lc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := lc.mutation.CreateTime(); ok {
		_spec.SetField(license.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := lc.mutation.UpdateTime(); ok {
		_spec.SetField(license.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.License.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LicenseUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (lc *LicenseCreate) OnConflict(opts ...sql.ConflictOption) *LicenseUpsertOne {
	lc.conflict = opts
	return &LicenseUpsertOne{
		create: lc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.License.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lc *LicenseCreate) OnConflictColumns(columns ...string) *LicenseUpsertOne {
	lc.conflict = append(lc.conflict, sql.ConflictColumns(columns...))
	return &LicenseUpsertOne{
		create: lc,
	}
}

type (
	// LicenseUpsertOne is the builder for "upsert"-ing
	//  one License node.
	LicenseUpsertOne struct {
		create *LicenseCreate
	}

	// LicenseUpsert is the "OnConflict" setter.
	LicenseUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *LicenseUpsert) SetUpdateTime(v time.Time) *LicenseUpsert {
	u.Set(license.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LicenseUpsert) UpdateUpdateTime() *LicenseUpsert {
	u.SetExcluded(license.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.License.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(license.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LicenseUpsertOne) UpdateNewValues() *LicenseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(license.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(license.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.License.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *LicenseUpsertOne) Ignore() *LicenseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LicenseUpsertOne) DoNothing() *LicenseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LicenseCreate.OnConflict
// documentation for more info.
func (u *LicenseUpsertOne) Update(set func(*LicenseUpsert)) *LicenseUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LicenseUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *LicenseUpsertOne) SetUpdateTime(v time.Time) *LicenseUpsertOne {
	return u.Update(func(s *LicenseUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LicenseUpsertOne) UpdateUpdateTime() *LicenseUpsertOne {
	return u.Update(func(s *LicenseUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *LicenseUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LicenseCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LicenseUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *LicenseUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *LicenseUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	builders []*LicenseCreate
	conflict []sql.ConflictOption
}

// Save creates the License entities in the database.
func (lcb *LicenseCreateBulk) Save(ctx context.Context) ([]*License, error) {
	specs := make([]*sqlgraph.CreateSpec, len(lcb.builders))
	nodes := make([]*License, len(lcb.builders))
	mutators := make([]Mutator, len(lcb.builders))
	for i := range lcb.builders {
		func(i int, root context.Context) {
			builder := lcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*LicenseMutation)
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
					_, err = mutators[i+1].Mutate(root, lcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = lcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, lcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
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
		if _, err := mutators[0].Mutate(ctx, lcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (lcb *LicenseCreateBulk) SaveX(ctx context.Context) []*License {
	v, err := lcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (lcb *LicenseCreateBulk) Exec(ctx context.Context) error {
	_, err := lcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lcb *LicenseCreateBulk) ExecX(ctx context.Context) {
	if err := lcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.License.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.LicenseUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (lcb *LicenseCreateBulk) OnConflict(opts ...sql.ConflictOption) *LicenseUpsertBulk {
	lcb.conflict = opts
	return &LicenseUpsertBulk{
		create: lcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.License.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (lcb *LicenseCreateBulk) OnConflictColumns(columns ...string) *LicenseUpsertBulk {
	lcb.conflict = append(lcb.conflict, sql.ConflictColumns(columns...))
	return &LicenseUpsertBulk{
		create: lcb,
	}
}

// LicenseUpsertBulk is the builder for "upsert"-ing
// a bulk of License nodes.
type LicenseUpsertBulk struct {
	create *LicenseCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.License.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(license.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *LicenseUpsertBulk) UpdateNewValues() *LicenseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(license.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(license.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.License.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *LicenseUpsertBulk) Ignore() *LicenseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *LicenseUpsertBulk) DoNothing() *LicenseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the LicenseCreateBulk.OnConflict
// documentation for more info.
func (u *LicenseUpsertBulk) Update(set func(*LicenseUpsert)) *LicenseUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&LicenseUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *LicenseUpsertBulk) SetUpdateTime(v time.Time) *LicenseUpsertBulk {
	return u.Update(func(s *LicenseUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *LicenseUpsertBulk) UpdateUpdateTime() *LicenseUpsertBulk {
	return u.Update(func(s *LicenseUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *LicenseUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the LicenseCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for LicenseCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *LicenseUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
