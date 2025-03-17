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
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
)

// UserTweetUpdate is the builder for updating UserTweet entities.
type UserTweetUpdate struct {
	config
	hooks    []Hook
	mutation *UserTweetMutation
}

// Where appends a list predicates to the UserTweetUpdate builder.
func (u *UserTweetUpdate) Where(ps ...predicate.UserTweet) *UserTweetUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetCreatedAt sets the "created_at" field.
func (m *UserTweetUpdate) SetCreatedAt(v time.Time) *UserTweetUpdate {
	m.mutation.SetCreatedAt(v)
	return m
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (m *UserTweetUpdate) SetNillableCreatedAt(v *time.Time) *UserTweetUpdate {
	if v != nil {
		m.SetCreatedAt(*v)
	}
	return m
}

// SetUserID sets the "user_id" field.
func (m *UserTweetUpdate) SetUserID(v int) *UserTweetUpdate {
	m.mutation.SetUserID(v)
	return m
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (m *UserTweetUpdate) SetNillableUserID(v *int) *UserTweetUpdate {
	if v != nil {
		m.SetUserID(*v)
	}
	return m
}

// SetTweetID sets the "tweet_id" field.
func (m *UserTweetUpdate) SetTweetID(v int) *UserTweetUpdate {
	m.mutation.SetTweetID(v)
	return m
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (m *UserTweetUpdate) SetNillableTweetID(v *int) *UserTweetUpdate {
	if v != nil {
		m.SetTweetID(*v)
	}
	return m
}

// SetUser sets the "user" edge to the User entity.
func (m *UserTweetUpdate) SetUser(v *User) *UserTweetUpdate {
	return m.SetUserID(v.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (m *UserTweetUpdate) SetTweet(v *Tweet) *UserTweetUpdate {
	return m.SetTweetID(v.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (m *UserTweetUpdate) Mutation() *UserTweetMutation {
	return m.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (u *UserTweetUpdate) ClearUser() *UserTweetUpdate {
	u.mutation.ClearUser()
	return u
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (u *UserTweetUpdate) ClearTweet() *UserTweetUpdate {
	u.mutation.ClearTweet()
	return u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *UserTweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *UserTweetUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *UserTweetUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserTweetUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *UserTweetUpdate) check() error {
	if u.mutation.UserCleared() && len(u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserTweet.user"`)
	}
	if u.mutation.TweetCleared() && len(u.mutation.TweetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserTweet.tweet"`)
	}
	return nil
}

func (u *UserTweetUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usertweet.Table, usertweet.Columns, sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
	}
	if u.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// UserTweetUpdateOne is the builder for updating a single UserTweet entity.
type UserTweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserTweetMutation
}

// SetCreatedAt sets the "created_at" field.
func (m *UserTweetUpdateOne) SetCreatedAt(v time.Time) *UserTweetUpdateOne {
	m.mutation.SetCreatedAt(v)
	return m
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (m *UserTweetUpdateOne) SetNillableCreatedAt(v *time.Time) *UserTweetUpdateOne {
	if v != nil {
		m.SetCreatedAt(*v)
	}
	return m
}

// SetUserID sets the "user_id" field.
func (m *UserTweetUpdateOne) SetUserID(v int) *UserTweetUpdateOne {
	m.mutation.SetUserID(v)
	return m
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (m *UserTweetUpdateOne) SetNillableUserID(v *int) *UserTweetUpdateOne {
	if v != nil {
		m.SetUserID(*v)
	}
	return m
}

// SetTweetID sets the "tweet_id" field.
func (m *UserTweetUpdateOne) SetTweetID(v int) *UserTweetUpdateOne {
	m.mutation.SetTweetID(v)
	return m
}

// SetNillableTweetID sets the "tweet_id" field if the given value is not nil.
func (m *UserTweetUpdateOne) SetNillableTweetID(v *int) *UserTweetUpdateOne {
	if v != nil {
		m.SetTweetID(*v)
	}
	return m
}

// SetUser sets the "user" edge to the User entity.
func (m *UserTweetUpdateOne) SetUser(v *User) *UserTweetUpdateOne {
	return m.SetUserID(v.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (m *UserTweetUpdateOne) SetTweet(v *Tweet) *UserTweetUpdateOne {
	return m.SetTweetID(v.ID)
}

// Mutation returns the UserTweetMutation object of the builder.
func (m *UserTweetUpdateOne) Mutation() *UserTweetMutation {
	return m.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (u *UserTweetUpdateOne) ClearUser() *UserTweetUpdateOne {
	u.mutation.ClearUser()
	return u
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (u *UserTweetUpdateOne) ClearTweet() *UserTweetUpdateOne {
	u.mutation.ClearTweet()
	return u
}

// Where appends a list predicates to the UserTweetUpdate builder.
func (u *UserTweetUpdateOne) Where(ps ...predicate.UserTweet) *UserTweetUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *UserTweetUpdateOne) Select(field string, fields ...string) *UserTweetUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated UserTweet entity.
func (u *UserTweetUpdateOne) Save(ctx context.Context) (*UserTweet, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *UserTweetUpdateOne) SaveX(ctx context.Context) *UserTweet {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *UserTweetUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UserTweetUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *UserTweetUpdateOne) check() error {
	if u.mutation.UserCleared() && len(u.mutation.UserIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserTweet.user"`)
	}
	if u.mutation.TweetCleared() && len(u.mutation.TweetIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "UserTweet.tweet"`)
	}
	return nil
}

func (u *UserTweetUpdateOne) sqlSave(ctx context.Context) (_n *UserTweet, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(usertweet.Table, usertweet.Columns, sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UserTweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, usertweet.FieldID)
		for _, f := range fields {
			if !usertweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != usertweet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.CreatedAt(); ok {
		_spec.SetField(usertweet.FieldCreatedAt, field.TypeTime, value)
	}
	if u.mutation.UserCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UserIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &UserTweet{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{usertweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
