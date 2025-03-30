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
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetCreate is the builder for creating a UserTweet entity.
type UserTweetCreate struct {
	config
	mutation *UserTweetMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreatedAt sets the "created_at" field.
func (_c *UserTweetCreate) SetCreatedAt(v time.Time) *UserTweetCreate {
	_c.mutation.SetCreatedAt(v)
	return _c
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (_c *UserTweetCreate) SetNillableCreatedAt(v *time.Time) *UserTweetCreate {
	if v != nil {
		_c.SetCreatedAt(*v)
	}
	return _c
}

// SetUserID sets the "user_id" field.
func (_c *UserTweetCreate) SetUserID(v int) *UserTweetCreate {
	_c.mutation.SetUserID(v)
	return _c
}

// SetTweetID sets the "tweet_id" field.
func (_c *UserTweetCreate) SetTweetID(v int) *UserTweetCreate {
	_c.mutation.SetTweetID(v)
	return _c
}

// SetUser sets the "user" edge to the User entity.
func (_c *UserTweetCreate) SetUser(v *User) *UserTweetCreate {
	return _c.SetUserID(v.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (_c *UserTweetCreate) SetTweet(v *Tweet) *UserTweetCreate {
	return _c.SetTweetID(v.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (_c *UserTweetCreate) Mutation() *UserTweetMutation {
	return _c.mutation
}

// Save creates the UserTweet in the database.
func (_c *UserTweetCreate) Save(ctx context.Context) (*UserTweet, error) {
	_c.defaults()
	return withHooks(ctx, _c.sqlSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *UserTweetCreate) SaveX(ctx context.Context) *UserTweet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserTweetCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserTweetCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *UserTweetCreate) defaults() {
	if _, ok := _c.mutation.CreatedAt(); !ok {
		v := usertweet.DefaultCreatedAt()
		_c.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *UserTweetCreate) check() error {
	if _, ok := _c.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserTweet.created_at"`)}
	}
	if _, ok := _c.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserTweet.user_id"`)}
	}
	if _, ok := _c.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "UserTweet.tweet_id"`)}
	}
	if len(_c.mutation.UserIDs()) == 0 {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserTweet.user"`)}
	}
	if len(_c.mutation.TweetIDs()) == 0 {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "UserTweet.tweet"`)}
	}
	return nil
}

func (_c *UserTweetCreate) sqlSave(ctx context.Context) (*UserTweet, error) {
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

func (_c *UserTweetCreate) createSpec() (*UserTweet, *sqlgraph.CreateSpec) {
	var (
		_node = &UserTweet{config: _c.config}
		_spec = sqlgraph.NewCreateSpec(usertweet.Table, sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt))
	)
	_spec.OnConflict = _c.conflict
	if value, ok := _c.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := _c.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := _c.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.TweetID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserTweet.Create().
//		SetCreatedAt(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserTweetUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (_c *UserTweetCreate) OnConflict(opts ...sql.ConflictOption) *UserTweetUpsertOne {
	_c.conflict = opts
	return &UserTweetUpsertOne{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *UserTweetCreate) OnConflictColumns(columns ...string) *UserTweetUpsertOne {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &UserTweetUpsertOne{
		create: _c,
	}
}

type (
	// UserTweetUpsertOne is the builder for "upsert"-ing
	//  one UserTweet node.
	UserTweetUpsertOne struct {
		create *UserTweetCreate
	}

	// UserTweetUpsert is the "OnConflict" setter.
	UserTweetUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreatedAt sets the "created_at" field.
func (u *UserTweetUpsert) SetCreatedAt(v time.Time) *UserTweetUpsert {
	u.Set(usertweet.FieldCreatedAt, v)
	return u
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserTweetUpsert) UpdateCreatedAt() *UserTweetUpsert {
	u.SetExcluded(usertweet.FieldCreatedAt)
	return u
}

// SetUserID sets the "user_id" field.
func (u *UserTweetUpsert) SetUserID(v int) *UserTweetUpsert {
	u.Set(usertweet.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserTweetUpsert) UpdateUserID() *UserTweetUpsert {
	u.SetExcluded(usertweet.FieldUserID)
	return u
}

// SetTweetID sets the "tweet_id" field.
func (u *UserTweetUpsert) SetTweetID(v int) *UserTweetUpsert {
	u.Set(usertweet.FieldTweetID, v)
	return u
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *UserTweetUpsert) UpdateTweetID() *UserTweetUpsert {
	u.SetExcluded(usertweet.FieldTweetID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserTweetUpsertOne) UpdateNewValues() *UserTweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UserTweetUpsertOne) Ignore() *UserTweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserTweetUpsertOne) DoNothing() *UserTweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserTweetCreate.OnConflict
// documentation for more info.
func (u *UserTweetUpsertOne) Update(set func(*UserTweetUpsert)) *UserTweetUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserTweetUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserTweetUpsertOne) SetCreatedAt(v time.Time) *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserTweetUpsertOne) UpdateCreatedAt() *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserTweetUpsertOne) SetUserID(v int) *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserTweetUpsertOne) UpdateUserID() *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateUserID()
	})
}

// SetTweetID sets the "tweet_id" field.
func (u *UserTweetUpsertOne) SetTweetID(v int) *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetTweetID(v)
	})
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *UserTweetUpsertOne) UpdateTweetID() *UserTweetUpsertOne {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateTweetID()
	})
}

// Exec executes the query.
func (u *UserTweetUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserTweetCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserTweetUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UserTweetUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UserTweetUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UserTweetCreateBulk is the builder for creating many UserTweet entities in bulk.
type UserTweetCreateBulk struct {
	config
	err      error
	builders []*UserTweetCreate
	conflict []sql.ConflictOption
}

// Save creates the UserTweet entities in the database.
func (_c *UserTweetCreateBulk) Save(ctx context.Context) ([]*UserTweet, error) {
	if _c.err != nil {
		return nil, _c.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(_c.builders))
	nodes := make([]*UserTweet, len(_c.builders))
	mutators := make([]Mutator, len(_c.builders))
	for i := range _c.builders {
		func(i int, root context.Context) {
			builder := _c.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserTweetMutation)
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
func (_c *UserTweetCreateBulk) SaveX(ctx context.Context) []*UserTweet {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *UserTweetCreateBulk) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *UserTweetCreateBulk) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UserTweet.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.UserTweetUpsert) {
//			SetCreatedAt(v+v).
//		}).
//		Exec(ctx)
func (_c *UserTweetCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserTweetUpsertBulk {
	_c.conflict = opts
	return &UserTweetUpsertBulk{
		create: _c,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (_c *UserTweetCreateBulk) OnConflictColumns(columns ...string) *UserTweetUpsertBulk {
	_c.conflict = append(_c.conflict, sql.ConflictColumns(columns...))
	return &UserTweetUpsertBulk{
		create: _c,
	}
}

// UserTweetUpsertBulk is the builder for "upsert"-ing
// a bulk of UserTweet nodes.
type UserTweetUpsertBulk struct {
	create *UserTweetCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UserTweetUpsertBulk) UpdateNewValues() *UserTweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UserTweetUpsertBulk) Ignore() *UserTweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UserTweetUpsertBulk) DoNothing() *UserTweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UserTweetCreateBulk.OnConflict
// documentation for more info.
func (u *UserTweetUpsertBulk) Update(set func(*UserTweetUpsert)) *UserTweetUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UserTweetUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreatedAt sets the "created_at" field.
func (u *UserTweetUpsertBulk) SetCreatedAt(v time.Time) *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetCreatedAt(v)
	})
}

// UpdateCreatedAt sets the "created_at" field to the value that was provided on create.
func (u *UserTweetUpsertBulk) UpdateCreatedAt() *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateCreatedAt()
	})
}

// SetUserID sets the "user_id" field.
func (u *UserTweetUpsertBulk) SetUserID(v int) *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *UserTweetUpsertBulk) UpdateUserID() *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateUserID()
	})
}

// SetTweetID sets the "tweet_id" field.
func (u *UserTweetUpsertBulk) SetTweetID(v int) *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.SetTweetID(v)
	})
}

// UpdateTweetID sets the "tweet_id" field to the value that was provided on create.
func (u *UserTweetUpsertBulk) UpdateTweetID() *UserTweetUpsertBulk {
	return u.Update(func(s *UserTweetUpsert) {
		s.UpdateTweetID()
	})
}

// Exec executes the query.
func (u *UserTweetUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if b == nil {
			return fmt.Errorf("ent: missing builder at index %d, unexpected nil builder passed to CreateBulk", i)
		}
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UserTweetCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UserTweetCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserTweetUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
