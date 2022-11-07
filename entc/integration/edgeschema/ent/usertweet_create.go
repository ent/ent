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
func (utc *UserTweetCreate) SetCreatedAt(t time.Time) *UserTweetCreate {
	utc.mutation.SetCreatedAt(t)
	return utc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (utc *UserTweetCreate) SetNillableCreatedAt(t *time.Time) *UserTweetCreate {
	if t != nil {
		utc.SetCreatedAt(*t)
	}
	return utc
}

// SetUserID sets the "user_id" field.
func (utc *UserTweetCreate) SetUserID(i int) *UserTweetCreate {
	utc.mutation.SetUserID(i)
	return utc
}

// SetTweetID sets the "tweet_id" field.
func (utc *UserTweetCreate) SetTweetID(i int) *UserTweetCreate {
	utc.mutation.SetTweetID(i)
	return utc
}

// SetUser sets the "user" edge to the User entity.
func (utc *UserTweetCreate) SetUser(u *User) *UserTweetCreate {
	return utc.SetUserID(u.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (utc *UserTweetCreate) SetTweet(t *Tweet) *UserTweetCreate {
	return utc.SetTweetID(t.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (utc *UserTweetCreate) Mutation() *UserTweetMutation {
	return utc.mutation
}

// Save creates the UserTweet in the database.
func (utc *UserTweetCreate) Save(ctx context.Context) (*UserTweet, error) {
	var (
		err  error
		node *UserTweet
	)
	utc.defaults()
	if len(utc.hooks) == 0 {
		if err = utc.check(); err != nil {
			return nil, err
		}
		node, err = utc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserTweetMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = utc.check(); err != nil {
				return nil, err
			}
			utc.mutation = mutation
			if node, err = utc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(utc.hooks) - 1; i >= 0; i-- {
			if utc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = utc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, utc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*UserTweet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserTweetMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (utc *UserTweetCreate) SaveX(ctx context.Context) *UserTweet {
	v, err := utc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utc *UserTweetCreate) Exec(ctx context.Context) error {
	_, err := utc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utc *UserTweetCreate) ExecX(ctx context.Context) {
	if err := utc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (utc *UserTweetCreate) defaults() {
	if _, ok := utc.mutation.CreatedAt(); !ok {
		v := usertweet.DefaultCreatedAt()
		utc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (utc *UserTweetCreate) check() error {
	if _, ok := utc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "UserTweet.created_at"`)}
	}
	if _, ok := utc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "UserTweet.user_id"`)}
	}
	if _, ok := utc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet_id", err: errors.New(`ent: missing required field "UserTweet.tweet_id"`)}
	}
	if _, ok := utc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "UserTweet.user"`)}
	}
	if _, ok := utc.mutation.TweetID(); !ok {
		return &ValidationError{Name: "tweet", err: errors.New(`ent: missing required edge "UserTweet.tweet"`)}
	}
	return nil
}

func (utc *UserTweetCreate) sqlSave(ctx context.Context) (*UserTweet, error) {
	_node, _spec := utc.createSpec()
	if err := sqlgraph.CreateNode(ctx, utc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (utc *UserTweetCreate) createSpec() (*UserTweet, *sqlgraph.CreateSpec) {
	var (
		_node = &UserTweet{config: utc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: usertweet.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: usertweet.FieldID,
			},
		}
	)
	_spec.OnConflict = utc.conflict
	if value, ok := utc.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := utc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.UserTable,
			Columns: []string{usertweet.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := utc.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   usertweet.TweetTable,
			Columns: []string{usertweet.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
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
func (utc *UserTweetCreate) OnConflict(opts ...sql.ConflictOption) *UserTweetUpsertOne {
	utc.conflict = opts
	return &UserTweetUpsertOne{
		create: utc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (utc *UserTweetCreate) OnConflictColumns(columns ...string) *UserTweetUpsertOne {
	utc.conflict = append(utc.conflict, sql.ConflictColumns(columns...))
	return &UserTweetUpsertOne{
		create: utc,
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
	builders []*UserTweetCreate
	conflict []sql.ConflictOption
}

// Save creates the UserTweet entities in the database.
func (utcb *UserTweetCreateBulk) Save(ctx context.Context) ([]*UserTweet, error) {
	specs := make([]*sqlgraph.CreateSpec, len(utcb.builders))
	nodes := make([]*UserTweet, len(utcb.builders))
	mutators := make([]Mutator, len(utcb.builders))
	for i := range utcb.builders {
		func(i int, root context.Context) {
			builder := utcb.builders[i]
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
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, utcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = utcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, utcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, utcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (utcb *UserTweetCreateBulk) SaveX(ctx context.Context) []*UserTweet {
	v, err := utcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (utcb *UserTweetCreateBulk) Exec(ctx context.Context) error {
	_, err := utcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (utcb *UserTweetCreateBulk) ExecX(ctx context.Context) {
	if err := utcb.Exec(ctx); err != nil {
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
func (utcb *UserTweetCreateBulk) OnConflict(opts ...sql.ConflictOption) *UserTweetUpsertBulk {
	utcb.conflict = opts
	return &UserTweetUpsertBulk{
		create: utcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UserTweet.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (utcb *UserTweetCreateBulk) OnConflictColumns(columns ...string) *UserTweetUpsertBulk {
	utcb.conflict = append(utcb.conflict, sql.ConflictColumns(columns...))
	return &UserTweetUpsertBulk{
		create: utcb,
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
	for i, b := range u.create.builders {
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
