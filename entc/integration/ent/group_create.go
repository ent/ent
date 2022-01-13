// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/file"
	"entgo.io/ent/entc/integration/ent/group"
	"entgo.io/ent/entc/integration/ent/groupinfo"
	"entgo.io/ent/entc/integration/ent/user"
	"entgo.io/ent/schema/field"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetActive sets the "active" field.
func (gc *GroupCreate) SetActive(b bool) *GroupCreate {
	gc.mutation.SetActive(b)
	return gc
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (gc *GroupCreate) SetNillableActive(b *bool) *GroupCreate {
	if b != nil {
		gc.SetActive(*b)
	}
	return gc
}

// SetExpire sets the "expire" field.
func (gc *GroupCreate) SetExpire(t time.Time) *GroupCreate {
	gc.mutation.SetExpire(t)
	return gc
}

// SetType sets the "type" field.
func (gc *GroupCreate) SetType(s string) *GroupCreate {
	gc.mutation.SetType(s)
	return gc
}

// SetMaxUsers sets the "max_users" field.
func (gc *GroupCreate) SetMaxUsers(i int) *GroupCreate {
	gc.mutation.SetMaxUsers(i)
	return gc
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (gc *GroupCreate) SetNillableMaxUsers(i *int) *GroupCreate {
	if i != nil {
		gc.SetMaxUsers(*i)
	}
	return gc
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (gc *GroupCreate) AddFileIDs(ids ...int) *GroupCreate {
	gc.mutation.AddFileIDs(ids...)
	return gc
}

// AddFiles adds the "files" edges to the File entity.
func (gc *GroupCreate) AddFiles(f ...*File) *GroupCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gc.AddFileIDs(ids...)
}

// AddBlockedIDs adds the "blocked" edge to the User entity by IDs.
func (gc *GroupCreate) AddBlockedIDs(ids ...int) *GroupCreate {
	gc.mutation.AddBlockedIDs(ids...)
	return gc
}

// AddBlocked adds the "blocked" edges to the User entity.
func (gc *GroupCreate) AddBlocked(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddBlockedIDs(ids...)
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gc *GroupCreate) AddUserIDs(ids ...int) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the "users" edges to the User entity.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// SetInfoID sets the "info" edge to the GroupInfo entity by ID.
func (gc *GroupCreate) SetInfoID(id int) *GroupCreate {
	gc.mutation.SetInfoID(id)
	return gc
}

// SetInfo sets the "info" edge to the GroupInfo entity.
func (gc *GroupCreate) SetInfo(g *GroupInfo) *GroupCreate {
	return gc.SetInfoID(g.ID)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	var (
		err  error
		node *Group
	)
	gc.defaults()
	if len(gc.hooks) == 0 {
		if err = gc.check(); err != nil {
			return nil, err
		}
		node, err = gc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*GroupMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = gc.check(); err != nil {
				return nil, err
			}
			gc.mutation = mutation
			if node, err = gc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(gc.hooks) - 1; i >= 0; i-- {
			if gc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = gc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, gc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() {
	if _, ok := gc.mutation.Active(); !ok {
		v := group.DefaultActive
		gc.mutation.SetActive(v)
	}
	if _, ok := gc.mutation.MaxUsers(); !ok {
		v := group.DefaultMaxUsers
		gc.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Group.active"`)}
	}
	if _, ok := gc.mutation.Expire(); !ok {
		return &ValidationError{Name: "expire", err: errors.New(`ent: missing required field "Group.expire"`)}
	}
	if v, ok := gc.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Group.type": %w`, err)}
		}
	}
	if v, ok := gc.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return &ValidationError{Name: "max_users", err: fmt.Errorf(`ent: validator failed for field "Group.max_users": %w`, err)}
		}
	}
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if v, ok := gc.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
		}
	}
	if _, ok := gc.mutation.InfoID(); !ok {
		return &ValidationError{Name: "info", err: errors.New(`ent: missing required edge "Group.info"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: group.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: group.FieldID,
			},
		}
	)
	_spec.OnConflict = gc.conflict
	if value, ok := gc.mutation.Active(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: group.FieldActive,
		})
		_node.Active = value
	}
	if value, ok := gc.mutation.Expire(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: group.FieldExpire,
		})
		_node.Expire = value
	}
	if value, ok := gc.mutation.GetType(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldType,
		})
		_node.Type = &value
	}
	if value, ok := gc.mutation.MaxUsers(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: group.FieldMaxUsers,
		})
		_node.MaxUsers = value
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: group.FieldName,
		})
		_node.Name = value
	}
	if nodes := gc.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.FilesTable,
			Columns: []string{group.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: file.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.BlockedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.BlockedTable,
			Columns: []string{group.BlockedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   group.UsersTable,
			Columns: group.UsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := gc.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.InfoTable,
			Columns: []string{group.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: groupinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.group_info = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.Create().
//		SetActive(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetActive(v+v).
//		}).
//		Exec(ctx)
//
func (gc *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	gc.conflict = opts
	return &GroupUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gc *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: gc,
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

// SetActive sets the "active" field.
func (u *GroupUpsert) SetActive(v bool) *GroupUpsert {
	u.Set(group.FieldActive, v)
	return u
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *GroupUpsert) UpdateActive() *GroupUpsert {
	u.SetExcluded(group.FieldActive)
	return u
}

// SetExpire sets the "expire" field.
func (u *GroupUpsert) SetExpire(v time.Time) *GroupUpsert {
	u.Set(group.FieldExpire, v)
	return u
}

// UpdateExpire sets the "expire" field to the value that was provided on create.
func (u *GroupUpsert) UpdateExpire() *GroupUpsert {
	u.SetExcluded(group.FieldExpire)
	return u
}

// SetType sets the "type" field.
func (u *GroupUpsert) SetType(v string) *GroupUpsert {
	u.Set(group.FieldType, v)
	return u
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GroupUpsert) UpdateType() *GroupUpsert {
	u.SetExcluded(group.FieldType)
	return u
}

// ClearType clears the value of the "type" field.
func (u *GroupUpsert) ClearType() *GroupUpsert {
	u.SetNull(group.FieldType)
	return u
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupUpsert) SetMaxUsers(v int) *GroupUpsert {
	u.Set(group.FieldMaxUsers, v)
	return u
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupUpsert) UpdateMaxUsers() *GroupUpsert {
	u.SetExcluded(group.FieldMaxUsers)
	return u
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupUpsert) AddMaxUsers(v int) *GroupUpsert {
	u.Add(group.FieldMaxUsers, v)
	return u
}

// ClearMaxUsers clears the value of the "max_users" field.
func (u *GroupUpsert) ClearMaxUsers() *GroupUpsert {
	u.SetNull(group.FieldMaxUsers)
	return u
}

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
//
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Group.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
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

// SetActive sets the "active" field.
func (u *GroupUpsertOne) SetActive(v bool) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateActive() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateActive()
	})
}

// SetExpire sets the "expire" field.
func (u *GroupUpsertOne) SetExpire(v time.Time) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetExpire(v)
	})
}

// UpdateExpire sets the "expire" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateExpire() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateExpire()
	})
}

// SetType sets the "type" field.
func (u *GroupUpsertOne) SetType(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateType() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *GroupUpsertOne) ClearType() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearType()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupUpsertOne) SetMaxUsers(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupUpsertOne) AddMaxUsers(v int) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateMaxUsers() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMaxUsers()
	})
}

// ClearMaxUsers clears the value of the "max_users" field.
func (u *GroupUpsertOne) ClearMaxUsers() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.ClearMaxUsers()
	})
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
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
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
//			SetActive(v+v).
//		}).
//		Exec(ctx)
//
func (gcb *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	gcb.conflict = opts
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (gcb *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: gcb,
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
//
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
//
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

// SetActive sets the "active" field.
func (u *GroupUpsertBulk) SetActive(v bool) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetActive(v)
	})
}

// UpdateActive sets the "active" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateActive() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateActive()
	})
}

// SetExpire sets the "expire" field.
func (u *GroupUpsertBulk) SetExpire(v time.Time) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetExpire(v)
	})
}

// UpdateExpire sets the "expire" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateExpire() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateExpire()
	})
}

// SetType sets the "type" field.
func (u *GroupUpsertBulk) SetType(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetType(v)
	})
}

// UpdateType sets the "type" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateType() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateType()
	})
}

// ClearType clears the value of the "type" field.
func (u *GroupUpsertBulk) ClearType() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearType()
	})
}

// SetMaxUsers sets the "max_users" field.
func (u *GroupUpsertBulk) SetMaxUsers(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetMaxUsers(v)
	})
}

// AddMaxUsers adds v to the "max_users" field.
func (u *GroupUpsertBulk) AddMaxUsers(v int) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.AddMaxUsers(v)
	})
}

// UpdateMaxUsers sets the "max_users" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateMaxUsers() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateMaxUsers()
	})
}

// ClearMaxUsers clears the value of the "max_users" field.
func (u *GroupUpsertBulk) ClearMaxUsers() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.ClearMaxUsers()
	})
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
