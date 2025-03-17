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
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/grouptag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usergroup"
	"entgo.io/ent/schema/field"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (_c *GroupCreate) SetName(s string) *GroupCreate {
	_c.mutation.SetName(s)
	return _c
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_c *GroupCreate) SetNillableName(s *string) *GroupCreate {
	if s != nil {
		_c.SetName(*s)
	}
	return _c
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (_c *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	_c.mutation.AddUserIDs(ids...)
	return _c
}

// AddUsers adds the "users" edges to the User entity.
func (_c *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _c.AddUserIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (_c *GroupCreate) AddTagIDs(ids ...int) *GroupCreate {
	_c.mutation.AddTagIDs(ids...)
	return _c
}

// AddTags adds the "tags" edges to the Tag entity.
func (_c *GroupCreate) AddTags(t ...*Tag) *GroupCreate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return _c.AddTagIDs(ids...)
}

// AddJoinedUserIDs adds the "joined_users" edge to the UserGroup entity by IDs.
func (_c *GroupCreate) AddJoinedUserIDs(ids ...int) *GroupCreate {
	_c.mutation.AddJoinedUserIDs(ids...)
	return _c
}

// AddJoinedUsers adds the "joined_users" edges to the UserGroup entity.
func (_c *GroupCreate) AddJoinedUsers(u ...*UserGroup) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _c.AddJoinedUserIDs(ids...)
}

// AddGroupTagIDs adds the "group_tags" edge to the GroupTag entity by IDs.
func (_c *GroupCreate) AddGroupTagIDs(ids ...int) *GroupCreate {
	_c.mutation.AddGroupTagIDs(ids...)
	return _c
}

// AddGroupTags adds the "group_tags" edges to the GroupTag entity.
func (_c *GroupCreate) AddGroupTags(g ...*GroupTag) *GroupCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return _c.AddGroupTagIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (_c *GroupCreate) Mutation() *GroupMutation {
	return _c.mutation
}

// Save creates the Group in the database.
func (_c *GroupCreate) Save(ctx context.Context) (*Group, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *GroupCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *GroupCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *GroupCreate) defaults() {
	if _, ok := _c.mutation.Name(); !ok {
		v := group.DefaultName
		_c.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *GroupCreate) check() error {
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	return nil
}

func (_c *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
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

func (_c *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := _c.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserGroupCreate{config: _c.config, mutation: newUserGroupMutation(_c.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.TagsTable,
			Columns: group.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.JoinedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.JoinedUsersTable,
			Columns: []string{group.JoinedUsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usergroup.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.GroupTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   group.GroupTagsTable,
			Columns: []string{group.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
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
//	client.Group.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (_c *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	_c.conflict = opts
	return &GroupUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: _c,
	}
}

type (
	// GroupUpsertOne is the builder for "upsert"-ing
	//  one Group node.
	GroupUpsertOne struct {
		create *GroupCreate
	}

	// GroupUpsert is the "OnConflict" setter.
	GroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *GroupUpsert) SetName(v string) *GroupUpsert {
	u.Set(group.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsert) UpdateName() *GroupUpsert {
	u.SetExcluded(group.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupUpsertOne) Ignore() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertOne) DoNothing() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreate.OnConflict
// documentation for more info.
func (u *GroupUpsertOne) Update(set func(*GroupUpsert)) *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertOne) SetName(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateName() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *GroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (_c *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*Group, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
func (_c *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (_c *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	_c.conflict = opts
	return &GroupUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: _c,
	}
}

// GroupUpsertBulk is the builder for "upsert"-ing
// a bulk of Group nodes.
type GroupUpsertBulk struct {
	create *GroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *GroupUpsertBulk) UpdateNewValues() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupUpsertBulk) Ignore() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertBulk) DoNothing() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreateBulk.OnConflict
// documentation for more info.
func (u *GroupUpsertBulk) Update(set func(*GroupUpsert)) *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertBulk) SetName(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateName() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *GroupUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
