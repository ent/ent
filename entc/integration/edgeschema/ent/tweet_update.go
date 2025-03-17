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
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usertweet"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TweetUpdate is the builder for updating Tweet entities.
type TweetUpdate struct {
	config
	hooks    []Hook
	mutation *TweetMutation
}

// Where appends a list predicates to the TweetUpdate builder.
func (u *TweetUpdate) Where(ps ...predicate.Tweet) *TweetUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetText sets the "text" field.
func (m *TweetUpdate) SetText(v string) *TweetUpdate {
	m.mutation.SetText(v)
	return m
}

// SetNillableText sets the "text" field if the given value is not nil.
func (m *TweetUpdate) SetNillableText(v *string) *TweetUpdate {
	if v != nil {
		m.SetText(*v)
	}
	return m
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (m *TweetUpdate) AddLikedUserIDs(ids ...int) *TweetUpdate {
	m.mutation.AddLikedUserIDs(ids...)
	return m
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (m *TweetUpdate) AddLikedUsers(v ...*User) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddLikedUserIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (m *TweetUpdate) AddUserIDs(ids ...int) *TweetUpdate {
	m.mutation.AddUserIDs(ids...)
	return m
}

// AddUser adds the "user" edges to the User entity.
func (m *TweetUpdate) AddUser(v ...*User) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddUserIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (m *TweetUpdate) AddTagIDs(ids ...int) *TweetUpdate {
	m.mutation.AddTagIDs(ids...)
	return m
}

// AddTags adds the "tags" edges to the Tag entity.
func (m *TweetUpdate) AddTags(v ...*Tag) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTagIDs(ids...)
}

// AddTweetUserIDs adds the "tweet_user" edge to the UserTweet entity by IDs.
func (m *TweetUpdate) AddTweetUserIDs(ids ...int) *TweetUpdate {
	m.mutation.AddTweetUserIDs(ids...)
	return m
}

// AddTweetUser adds the "tweet_user" edges to the UserTweet entity.
func (m *TweetUpdate) AddTweetUser(v ...*UserTweet) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTweetUserIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (m *TweetUpdate) AddTweetTagIDs(ids ...uuid.UUID) *TweetUpdate {
	m.mutation.AddTweetTagIDs(ids...)
	return m
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (m *TweetUpdate) AddTweetTags(v ...*TweetTag) *TweetUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTweetTagIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (m *TweetUpdate) Mutation() *TweetMutation {
	return m.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (u *TweetUpdate) ClearLikedUsers() *TweetUpdate {
	u.mutation.ClearLikedUsers()
	return u
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (u *TweetUpdate) RemoveLikedUserIDs(ids ...int) *TweetUpdate {
	u.mutation.RemoveLikedUserIDs(ids...)
	return u
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (u *TweetUpdate) RemoveLikedUsers(v ...*User) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveLikedUserIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (u *TweetUpdate) ClearUser() *TweetUpdate {
	u.mutation.ClearUser()
	return u
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (u *TweetUpdate) RemoveUserIDs(ids ...int) *TweetUpdate {
	u.mutation.RemoveUserIDs(ids...)
	return u
}

// RemoveUser removes "user" edges to User entities.
func (u *TweetUpdate) RemoveUser(v ...*User) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveUserIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (u *TweetUpdate) ClearTags() *TweetUpdate {
	u.mutation.ClearTags()
	return u
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (u *TweetUpdate) RemoveTagIDs(ids ...int) *TweetUpdate {
	u.mutation.RemoveTagIDs(ids...)
	return u
}

// RemoveTags removes "tags" edges to Tag entities.
func (u *TweetUpdate) RemoveTags(v ...*Tag) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTagIDs(ids...)
}

// ClearTweetUser clears all "tweet_user" edges to the UserTweet entity.
func (u *TweetUpdate) ClearTweetUser() *TweetUpdate {
	u.mutation.ClearTweetUser()
	return u
}

// RemoveTweetUserIDs removes the "tweet_user" edge to UserTweet entities by IDs.
func (u *TweetUpdate) RemoveTweetUserIDs(ids ...int) *TweetUpdate {
	u.mutation.RemoveTweetUserIDs(ids...)
	return u
}

// RemoveTweetUser removes "tweet_user" edges to UserTweet entities.
func (u *TweetUpdate) RemoveTweetUser(v ...*UserTweet) *TweetUpdate {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTweetUserIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (u *TweetUpdate) ClearTweetTags() *TweetUpdate {
	u.mutation.ClearTweetTags()
	return u
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (u *TweetUpdate) RemoveTweetTagIDs(ids ...uuid.UUID) *TweetUpdate {
	u.mutation.RemoveTweetTagIDs(ids...)
	return u
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (u *TweetUpdate) RemoveTweetTags(v ...*TweetTag) *TweetUpdate {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTweetTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *TweetUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *TweetUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *TweetUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *TweetUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if u.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !u.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedUserIDs(); len(nodes) > 0 && !u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTagsIDs(); len(nodes) > 0 && !u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTweetUserIDs(); len(nodes) > 0 && !u.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// TweetUpdateOne is the builder for updating a single Tweet entity.
type TweetUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetMutation
}

// SetText sets the "text" field.
func (m *TweetUpdateOne) SetText(v string) *TweetUpdateOne {
	m.mutation.SetText(v)
	return m
}

// SetNillableText sets the "text" field if the given value is not nil.
func (m *TweetUpdateOne) SetNillableText(v *string) *TweetUpdateOne {
	if v != nil {
		m.SetText(*v)
	}
	return m
}

// AddLikedUserIDs adds the "liked_users" edge to the User entity by IDs.
func (m *TweetUpdateOne) AddLikedUserIDs(ids ...int) *TweetUpdateOne {
	m.mutation.AddLikedUserIDs(ids...)
	return m
}

// AddLikedUsers adds the "liked_users" edges to the User entity.
func (m *TweetUpdateOne) AddLikedUsers(v ...*User) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddLikedUserIDs(ids...)
}

// AddUserIDs adds the "user" edge to the User entity by IDs.
func (m *TweetUpdateOne) AddUserIDs(ids ...int) *TweetUpdateOne {
	m.mutation.AddUserIDs(ids...)
	return m
}

// AddUser adds the "user" edges to the User entity.
func (m *TweetUpdateOne) AddUser(v ...*User) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddUserIDs(ids...)
}

// AddTagIDs adds the "tags" edge to the Tag entity by IDs.
func (m *TweetUpdateOne) AddTagIDs(ids ...int) *TweetUpdateOne {
	m.mutation.AddTagIDs(ids...)
	return m
}

// AddTags adds the "tags" edges to the Tag entity.
func (m *TweetUpdateOne) AddTags(v ...*Tag) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTagIDs(ids...)
}

// AddTweetUserIDs adds the "tweet_user" edge to the UserTweet entity by IDs.
func (m *TweetUpdateOne) AddTweetUserIDs(ids ...int) *TweetUpdateOne {
	m.mutation.AddTweetUserIDs(ids...)
	return m
}

// AddTweetUser adds the "tweet_user" edges to the UserTweet entity.
func (m *TweetUpdateOne) AddTweetUser(v ...*UserTweet) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTweetUserIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (m *TweetUpdateOne) AddTweetTagIDs(ids ...uuid.UUID) *TweetUpdateOne {
	m.mutation.AddTweetTagIDs(ids...)
	return m
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (m *TweetUpdateOne) AddTweetTags(v ...*TweetTag) *TweetUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return m.AddTweetTagIDs(ids...)
}

// Mutation returns the TweetMutation object of the builder.
func (m *TweetUpdateOne) Mutation() *TweetMutation {
	return m.mutation
}

// ClearLikedUsers clears all "liked_users" edges to the User entity.
func (u *TweetUpdateOne) ClearLikedUsers() *TweetUpdateOne {
	u.mutation.ClearLikedUsers()
	return u
}

// RemoveLikedUserIDs removes the "liked_users" edge to User entities by IDs.
func (u *TweetUpdateOne) RemoveLikedUserIDs(ids ...int) *TweetUpdateOne {
	u.mutation.RemoveLikedUserIDs(ids...)
	return u
}

// RemoveLikedUsers removes "liked_users" edges to User entities.
func (u *TweetUpdateOne) RemoveLikedUsers(v ...*User) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveLikedUserIDs(ids...)
}

// ClearUser clears all "user" edges to the User entity.
func (u *TweetUpdateOne) ClearUser() *TweetUpdateOne {
	u.mutation.ClearUser()
	return u
}

// RemoveUserIDs removes the "user" edge to User entities by IDs.
func (u *TweetUpdateOne) RemoveUserIDs(ids ...int) *TweetUpdateOne {
	u.mutation.RemoveUserIDs(ids...)
	return u
}

// RemoveUser removes "user" edges to User entities.
func (u *TweetUpdateOne) RemoveUser(v ...*User) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveUserIDs(ids...)
}

// ClearTags clears all "tags" edges to the Tag entity.
func (u *TweetUpdateOne) ClearTags() *TweetUpdateOne {
	u.mutation.ClearTags()
	return u
}

// RemoveTagIDs removes the "tags" edge to Tag entities by IDs.
func (u *TweetUpdateOne) RemoveTagIDs(ids ...int) *TweetUpdateOne {
	u.mutation.RemoveTagIDs(ids...)
	return u
}

// RemoveTags removes "tags" edges to Tag entities.
func (u *TweetUpdateOne) RemoveTags(v ...*Tag) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTagIDs(ids...)
}

// ClearTweetUser clears all "tweet_user" edges to the UserTweet entity.
func (u *TweetUpdateOne) ClearTweetUser() *TweetUpdateOne {
	u.mutation.ClearTweetUser()
	return u
}

// RemoveTweetUserIDs removes the "tweet_user" edge to UserTweet entities by IDs.
func (u *TweetUpdateOne) RemoveTweetUserIDs(ids ...int) *TweetUpdateOne {
	u.mutation.RemoveTweetUserIDs(ids...)
	return u
}

// RemoveTweetUser removes "tweet_user" edges to UserTweet entities.
func (u *TweetUpdateOne) RemoveTweetUser(v ...*UserTweet) *TweetUpdateOne {
	ids := make([]int, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTweetUserIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (u *TweetUpdateOne) ClearTweetTags() *TweetUpdateOne {
	u.mutation.ClearTweetTags()
	return u
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (u *TweetUpdateOne) RemoveTweetTagIDs(ids ...uuid.UUID) *TweetUpdateOne {
	u.mutation.RemoveTweetTagIDs(ids...)
	return u
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (u *TweetUpdateOne) RemoveTweetTags(v ...*TweetTag) *TweetUpdateOne {
	ids := make([]uuid.UUID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return u.RemoveTweetTagIDs(ids...)
}

// Where appends a list predicates to the TweetUpdate builder.
func (u *TweetUpdateOne) Where(ps ...predicate.Tweet) *TweetUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *TweetUpdateOne) Select(field string, fields ...string) *TweetUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Tweet entity.
func (u *TweetUpdateOne) Save(ctx context.Context) (*Tweet, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *TweetUpdateOne) SaveX(ctx context.Context) *Tweet {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *TweetUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TweetUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (u *TweetUpdateOne) sqlSave(ctx context.Context) (_n *Tweet, err error) {
	_spec := sqlgraph.NewUpdateSpec(tweet.Table, tweet.Columns, sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tweet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweet.FieldID)
		for _, f := range fields {
			if !tweet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweet.FieldID {
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
	if value, ok := u.mutation.Text(); ok {
		_spec.SetField(tweet.FieldText, field.TypeString, value)
	}
	if u.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedLikedUsersIDs(); len(nodes) > 0 && !u.mutation.LikedUsersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.LikedUsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.LikedUsersTable,
			Columns: tweet.LikedUsersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetLikeCreate{config: u.config, mutation: newTweetLikeMutation(u.config, OpCreate)}
		_ = createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedUserIDs(); len(nodes) > 0 && !u.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.UserTable,
			Columns: tweet.UserPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &UserTweetCreate{config: u.config, mutation: newUserTweetMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTagsIDs(); len(nodes) > 0 && !u.mutation.TagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tweet.TagsTable,
			Columns: tweet.TagsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: u.config, mutation: newTweetTagMutation(u.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTweetUserIDs(); len(nodes) > 0 && !u.mutation.TweetUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetUserTable,
			Columns: []string{tweet.TweetUserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(usertweet.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !u.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tweet.TweetTagsTable,
			Columns: []string{tweet.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &Tweet{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
