// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	mutation *GroupInfoMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetDesc sets the "desc" field.
func (gic *GroupInfoCreate) SetDesc(s string) *GroupInfoCreate {
	gic.mutation.SetDesc(s)
	return gic
}

// SetMaxUsers sets the "max_users" field.
func (gic *GroupInfoCreate) SetMaxUsers(i int) *GroupInfoCreate {
	gic.mutation.SetMaxUsers(i)
	return gic
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (gic *GroupInfoCreate) SetNillableMaxUsers(i *int) *GroupInfoCreate {
	if i != nil {
		gic.SetMaxUsers(*i)
	}
	return gic
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (gic *GroupInfoCreate) AddGroupIDs(ids ...int) *GroupInfoCreate {
	gic.mutation.AddGroupIDs(ids...)
	return gic
}

// AddGroups adds the "groups" edges to the Group entity.
func (gic *GroupInfoCreate) AddGroups(g ...*Group) *GroupInfoCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gic.AddGroupIDs(ids...)
}

// Mutation returns the GroupInfoMutation object of the builder.
func (gic *GroupInfoCreate) Mutation() *GroupInfoMutation {
	return gic.mutation
}

// Save creates the GroupInfo in the database.
func (gic *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	var (
		err  error
		node *GroupInfo
	)
	gic.defaults()
	if len(gic.hooks) == 0 {
		if err = gic.check(); err != nil {
			return nil, err
		}
		node, err = gic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupInfoMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gic.check(); err != nil {
				return nil, err
			}
			gic.mutation = mutation
			if node, err = gic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gic.hooks) - 1; i >= 0; i-- {
			if gic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gic *GroupInfoCreate) Exec(ctx context.Context) error {
	_, err := gic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gic *GroupInfoCreate) ExecX(ctx context.Context) {
	if err := gic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gic *GroupInfoCreate) defaults() {
	if _, ok := gic.mutation.MaxUsers(); !ok {
		v := groupinfo.DefaultMaxUsers
		gic.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gic *GroupInfoCreate) check() error {
	if _, ok := gic.mutation.Desc(); !ok {
		return &ValidationError{Name: "desc", err: errors.New(`ent: missing required field "GroupInfo.desc"`)}
	}
	if _, ok := gic.mutation.MaxUsers(); !ok {
		return &ValidationError{Name: "max_users", err: errors.New(`ent: missing required field "GroupInfo.max_users"`)}
	}
	return nil
}

func (gic *GroupInfoCreate) sqlSave(ctx context.Context) (*GroupInfo, error) {
	_node, _spec := gic.createSpec()
	if err := sqlgraph.CreateNode(ctx, gic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gic *GroupInfoCreate) createSpec() (*GroupInfo, *sqlgraph.CreateSpec) {
	var (
		_node = &GroupInfo{config: gic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: groupinfo.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: groupinfo.FieldID,
			},
		}
	)
	_spec.OnConflict = gic.conflict
	if value, ok := gic.mutation.Desc(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: groupinfo.FieldDesc,
		})
		_node.Desc = value
	}
	if value, ok := gic.mutation.MaxUsers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: groupinfo.FieldMaxUsers,
		})
		_node.MaxUsers = value
	}
	if nodes := gic.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   groupinfo.GroupsTable,
			Columns: []string{groupinfo.GroupsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
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
//	client.GroupInfo.Create().
//		SetDesc(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupInfoUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
//
func (gic *GroupInfoCreate) OnConflict(opts ...sql.ConflictOption) *GroupInfoUpsertOne {
	gic.conflict = opts
	return &GroupInfoUpsertOne{
		create: gic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gic *GroupInfoCreate) OnConflictColumns(columns ...string) *GroupInfoUpsertOne {
	gic.conflict = append(gic.conflict, sql.ConflictColumns(columns...))
	return &GroupInfoUpsertOne{
		create: gic,
	}
}

type (
	// GroupInfoUpsertOne is the builder for "upsert"-ing
	//  one GroupInfo node.
	GroupInfoUpsertOne struct {
		create *GroupInfoCreate
	}

	// GroupInfoUpsert is the "OnConflict" setter.
	GroupInfoUpsert struct {
		*sql.UpdateSet
	}
)

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsert) SetDesc(v string) *GroupInfoUpsert {
	u.Set(groupinfo.FieldDesc, v)
	return u
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsert) UpdateDesc() *GroupInfoUpsert {
	u.SetExcluded(groupinfo.FieldDesc)
	return u
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsert) SetMaxUsers(v int) *GroupInfoUpsert {
	u.Set(groupinfo.FieldMaxUsers, v)
	return u
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsert) UpdateMaxUsers() *GroupInfoUpsert {
	u.SetExcluded(groupinfo.FieldMaxUsers)
	return u
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsert) AddMaxUsers(v int) *GroupInfoUpsert {
	u.Add(groupinfo.FieldMaxUsers, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *GroupInfoUpsertOne) UpdateNewValues() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.GroupInfo.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *GroupInfoUpsertOne) Ignore() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupInfoUpsertOne) DoNothing() *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupInfoCreate.OnConflict
// documentation for more info.
func (u *GroupInfoUpsertOne) Update(set func(*GroupInfoUpsert)) *GroupInfoUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsertOne) SetDesc(v string) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsertOne) UpdateDesc() *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateDesc()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsertOne) SetMaxUsers(v int) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsertOne) AddMaxUsers(v int) *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsertOne) UpdateMaxUsers() *GroupInfoUpsertOne {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateMaxUsers()
	})
}

// Exec executes the query.
func (u *GroupInfoUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupInfoCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupInfoUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupInfoUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupInfoUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupInfoCreateBulk is the builder for creating many GroupInfo entities in bulk.
type GroupInfoCreateBulk struct {
	config
	builders []*GroupInfoCreate
	conflict []sql.ConflictOption
}

// Save creates the GroupInfo entities in the database.
func (gicb *GroupInfoCreateBulk) Save(ctx context.Context) ([]*GroupInfo, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gicb.builders))
	nodes := make([]*GroupInfo, len(gicb.builders))
	mutators := make([]Mutator, len(gicb.builders))
	for i := range gicb.builders {
		func(i int, root context.Context) {
			builder := gicb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupInfoMutation)
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
					_, err = mutators[i+1].Mutate(root, gicb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gicb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gicb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, gicb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gicb *GroupInfoCreateBulk) SaveX(ctx context.Context) []*GroupInfo {
	v, err := gicb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gicb *GroupInfoCreateBulk) Exec(ctx context.Context) error {
	_, err := gicb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gicb *GroupInfoCreateBulk) ExecX(ctx context.Context) {
	if err := gicb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.GroupInfo.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupInfoUpsert) {
//			SetDesc(v+v).
//		}).
//		Exec(ctx)
//
func (gicb *GroupInfoCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupInfoUpsertBulk {
	gicb.conflict = opts
	return &GroupInfoUpsertBulk{
		create: gicb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gicb *GroupInfoCreateBulk) OnConflictColumns(columns ...string) *GroupInfoUpsertBulk {
	gicb.conflict = append(gicb.conflict, sql.ConflictColumns(columns...))
	return &GroupInfoUpsertBulk{
		create: gicb,
	}
}

// GroupInfoUpsertBulk is the builder for "upsert"-ing
// a bulk of GroupInfo nodes.
type GroupInfoUpsertBulk struct {
	create *GroupInfoCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *GroupInfoUpsertBulk) UpdateNewValues() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.GroupInfo.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *GroupInfoUpsertBulk) Ignore() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupInfoUpsertBulk) DoNothing() *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupInfoCreateBulk.OnConflict
// documentation for more info.
func (u *GroupInfoUpsertBulk) Update(set func(*GroupInfoUpsert)) *GroupInfoUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupInfoUpsert{UpdateSet: update})
	}))
	return u
}

// SetDesc sets the "desc" field.
func (u *GroupInfoUpsertBulk) SetDesc(v string) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetDesc(v)
	})
}

// UpdateDesc sets the "desc" field to the value that was provided on create.
func (u *GroupInfoUpsertBulk) UpdateDesc() *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateDesc()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupInfoUpsertBulk) SetMaxUsers(v int) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupInfoUpsertBulk) AddMaxUsers(v int) *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupInfoUpsertBulk) UpdateMaxUsers() *GroupInfoUpsertBulk {
	return u.Update(func(s *GroupInfoUpsert) {
		s.UpdateMaxUsers()
	})
}

// Exec executes the query.
func (u *GroupInfoUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupInfoCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupInfoCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupInfoUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
