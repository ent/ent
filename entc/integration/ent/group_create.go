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
func (_c *GroupCreate) SetActive(b bool) *GroupCreate {
	_c.mutation.SetActive(b)
	return _c
}

// SetNillableActive sets the "active" field if the given value is not nil.
func (_c *GroupCreate) SetNillableActive(b *bool) *GroupCreate {
	if b != nil {
		_c.SetActive(*b)
	}
	return _c
}

// SetExpire sets the "expire" field.
func (_c *GroupCreate) SetExpire(t time.Time) *GroupCreate {
	_c.mutation.SetExpire(t)
	return _c
}

// SetType sets the "type" field.
func (_c *GroupCreate) SetType(s string) *GroupCreate {
	_c.mutation.SetType(s)
	return _c
}

// SetNillableType sets the "type" field if the given value is not nil.
func (_c *GroupCreate) SetNillableType(s *string) *GroupCreate {
	if s != nil {
		_c.SetType(*s)
	}
	return _c
}

// SetMaxUsers sets the "max_users" field.
func (_c *GroupCreate) SetMaxUsers(i int) *GroupCreate {
	_c.mutation.SetMaxUsers(i)
	return _c
}

// SetNillableMaxUsers sets the "max_users" field if the given value is not nil.
func (_c *GroupCreate) SetNillableMaxUsers(i *int) *GroupCreate {
	if i != nil {
		_c.SetMaxUsers(*i)
	}
	return _c
}

// SetName sets the "name" field.
func (_c *GroupCreate) SetName(s string) *GroupCreate {
	_c.mutation.SetName(s)
	return _c
}

// AddFileIDs adds the "files" edge to the File entity by IDs.
func (_c *GroupCreate) AddFileIDs(ids ...int) *GroupCreate {
	_c.mutation.AddFileIDs(ids...)
	return _c
}

// AddFiles adds the "files" edges to the File entity.
func (_c *GroupCreate) AddFiles(f ...*File) *GroupCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return _c.AddFileIDs(ids...)
}

// AddBlockedIDs adds the "blocked" edge to the User entity by IDs.
func (_c *GroupCreate) AddBlockedIDs(ids ...int) *GroupCreate {
	_c.mutation.AddBlockedIDs(ids...)
	return _c
}

// AddBlocked adds the "blocked" edges to the User entity.
func (_c *GroupCreate) AddBlocked(u ...*User) *GroupCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return _c.AddBlockedIDs(ids...)
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

// SetInfoID sets the "info" edge to the GroupInfo entity by ID.
func (_c *GroupCreate) SetInfoID(id int) *GroupCreate {
	_c.mutation.SetInfoID(id)
	return _c
}

// SetInfo sets the "info" edge to the GroupInfo entity.
func (_c *GroupCreate) SetInfo(g *GroupInfo) *GroupCreate {
	return _c.SetInfoID(g.ID)
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
	if _, ok := _c.mutation.Active(); !ok {
		v := group.DefaultActive
		_c.mutation.SetActive(v)
	}
	if _, ok := _c.mutation.MaxUsers(); !ok {
		v := group.DefaultMaxUsers
		_c.mutation.SetMaxUsers(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *GroupCreate) check() error {
	if _, ok := _c.mutation.Active(); !ok {
		return &ValidationError{Name: "active", err: errors.New(`ent: missing required field "Group.active"`)}
	}
	if _, ok := _c.mutation.Expire(); !ok {
		return &ValidationError{Name: "expire", err: errors.New(`ent: missing required field "Group.expire"`)}
	}
	if v, ok := _c.mutation.GetType(); ok {
		if err := group.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`ent: validator failed for field "Group.type": %w`, err)}
		}
	}
	if v, ok := _c.mutation.MaxUsers(); ok {
		if err := group.MaxUsersValidator(v); err != nil {
			return &ValidationError{Name: "max_users", err: fmt.Errorf(`ent: validator failed for field "Group.max_users": %w`, err)}
		}
	}
	if _, ok := _c.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if v, ok := _c.mutation.Name(); ok {
		if err := group.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Group.name": %w`, err)}
		}
	}
	if len(_c.mutation.InfoIDs()) == 0 {
		return &ValidationError{Name: "info", err: errors.New(`ent: missing required edge "Group.info"`)}
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
	if value, ok := _c.mutation.Active(); ok {
		_spec.SetField(group.FieldActive, field.TypeBool, value)
		_node.Active = value
	}
	if value, ok := _c.mutation.Expire(); ok {
		_spec.SetField(group.FieldExpire, field.TypeTime, value)
		_node.Expire = value
	}
	if value, ok := _c.mutation.GetType(); ok {
		_spec.SetField(group.FieldType, field.TypeString, value)
		_node.Type = &value
	}
	if value, ok := _c.mutation.MaxUsers(); ok {
		_spec.SetField(group.FieldMaxUsers, field.TypeInt, value)
		_node.MaxUsers = value
	}
	if value, ok := _c.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if nodes := _c.mutation.FilesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.FilesTable,
			Columns: []string{group.FilesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(file.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.BlockedIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.BlockedTable,
			Columns: []string{group.BlockedColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   group.InfoTable,
			Columns: []string{group.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(groupinfo.FieldID, field.TypeInt),
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
//			SetActive(v+v).
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
