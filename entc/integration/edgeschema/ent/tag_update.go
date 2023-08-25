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
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	"entgo.io/ent/entc/integration/edgeschema/ent/tag"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	"entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetValue sets the "value" field.
func (tu *TagUpdate) SetValue(s string) *TagUpdate {
	tu.mutation.SetValue(s)
	return tu
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by IDs.
func (tu *TagUpdate) AddTweetIDs(ids ...int) *TagUpdate {
	tu.mutation.AddTweetIDs(ids...)
	return tu
}

// AddTweets adds the "tweets" edges to the Tweet entity.
func (tu *TagUpdate) AddTweets(t ...*Tweet) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTweetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (tu *TagUpdate) AddGroupIDs(ids ...int) *TagUpdate {
	tu.mutation.AddGroupIDs(ids...)
	return tu
}

// AddGroups adds the "groups" edges to the Group entity.
func (tu *TagUpdate) AddGroups(g ...*Group) *TagUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddGroupIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (tu *TagUpdate) AddTweetTagIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.AddTweetTagIDs(ids...)
	return tu
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (tu *TagUpdate) AddTweetTags(t ...*TweetTag) *TagUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.AddTweetTagIDs(ids...)
}

// AddGroupTagIDs adds the "group_tags" edge to the GroupTag entity by IDs.
func (tu *TagUpdate) AddGroupTagIDs(ids ...int) *TagUpdate {
	tu.mutation.AddGroupTagIDs(ids...)
	return tu
}

// AddGroupTags adds the "group_tags" edges to the GroupTag entity.
func (tu *TagUpdate) AddGroupTags(g ...*GroupTag) *TagUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.AddGroupTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearTweets clears all "tweets" edges to the Tweet entity.
func (tu *TagUpdate) ClearTweets() *TagUpdate {
	tu.mutation.ClearTweets()
	return tu
}

// RemoveTweetIDs removes the "tweets" edge to Tweet entities by IDs.
func (tu *TagUpdate) RemoveTweetIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveTweetIDs(ids...)
	return tu
}

// RemoveTweets removes "tweets" edges to Tweet entities.
func (tu *TagUpdate) RemoveTweets(t ...*Tweet) *TagUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTweetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (tu *TagUpdate) ClearGroups() *TagUpdate {
	tu.mutation.ClearGroups()
	return tu
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (tu *TagUpdate) RemoveGroupIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveGroupIDs(ids...)
	return tu
}

// RemoveGroups removes "groups" edges to Group entities.
func (tu *TagUpdate) RemoveGroups(g ...*Group) *TagUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveGroupIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (tu *TagUpdate) ClearTweetTags() *TagUpdate {
	tu.mutation.ClearTweetTags()
	return tu
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (tu *TagUpdate) RemoveTweetTagIDs(ids ...uuid.UUID) *TagUpdate {
	tu.mutation.RemoveTweetTagIDs(ids...)
	return tu
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (tu *TagUpdate) RemoveTweetTags(t ...*TweetTag) *TagUpdate {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tu.RemoveTweetTagIDs(ids...)
}

// ClearGroupTags clears all "group_tags" edges to the GroupTag entity.
func (tu *TagUpdate) ClearGroupTags() *TagUpdate {
	tu.mutation.ClearGroupTags()
	return tu
}

// RemoveGroupTagIDs removes the "group_tags" edge to GroupTag entities by IDs.
func (tu *TagUpdate) RemoveGroupTagIDs(ids ...int) *TagUpdate {
	tu.mutation.RemoveGroupTagIDs(ids...)
	return tu
}

// RemoveGroupTags removes "group_tags" edges to GroupTag entities.
func (tu *TagUpdate) RemoveGroupTags(g ...*GroupTag) *TagUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tu.RemoveGroupTagIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.Value(); ok {
		_spec.SetField(tag.FieldValue, field.TypeString, value)
	}
	if tu.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTweetsIDs(); len(nodes) > 0 && !tu.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TweetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tu.config, mutation: newTweetTagMutation(tu.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !tu.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !tu.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tu.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedGroupTagsIDs(); len(nodes) > 0 && !tu.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.GroupTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetValue sets the "value" field.
func (tuo *TagUpdateOne) SetValue(s string) *TagUpdateOne {
	tuo.mutation.SetValue(s)
	return tuo
}

// AddTweetIDs adds the "tweets" edge to the Tweet entity by IDs.
func (tuo *TagUpdateOne) AddTweetIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddTweetIDs(ids...)
	return tuo
}

// AddTweets adds the "tweets" edges to the Tweet entity.
func (tuo *TagUpdateOne) AddTweets(t ...*Tweet) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTweetIDs(ids...)
}

// AddGroupIDs adds the "groups" edge to the Group entity by IDs.
func (tuo *TagUpdateOne) AddGroupIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddGroupIDs(ids...)
	return tuo
}

// AddGroups adds the "groups" edges to the Group entity.
func (tuo *TagUpdateOne) AddGroups(g ...*Group) *TagUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddGroupIDs(ids...)
}

// AddTweetTagIDs adds the "tweet_tags" edge to the TweetTag entity by IDs.
func (tuo *TagUpdateOne) AddTweetTagIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.AddTweetTagIDs(ids...)
	return tuo
}

// AddTweetTags adds the "tweet_tags" edges to the TweetTag entity.
func (tuo *TagUpdateOne) AddTweetTags(t ...*TweetTag) *TagUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.AddTweetTagIDs(ids...)
}

// AddGroupTagIDs adds the "group_tags" edge to the GroupTag entity by IDs.
func (tuo *TagUpdateOne) AddGroupTagIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.AddGroupTagIDs(ids...)
	return tuo
}

// AddGroupTags adds the "group_tags" edges to the GroupTag entity.
func (tuo *TagUpdateOne) AddGroupTags(g ...*GroupTag) *TagUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.AddGroupTagIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearTweets clears all "tweets" edges to the Tweet entity.
func (tuo *TagUpdateOne) ClearTweets() *TagUpdateOne {
	tuo.mutation.ClearTweets()
	return tuo
}

// RemoveTweetIDs removes the "tweets" edge to Tweet entities by IDs.
func (tuo *TagUpdateOne) RemoveTweetIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveTweetIDs(ids...)
	return tuo
}

// RemoveTweets removes "tweets" edges to Tweet entities.
func (tuo *TagUpdateOne) RemoveTweets(t ...*Tweet) *TagUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTweetIDs(ids...)
}

// ClearGroups clears all "groups" edges to the Group entity.
func (tuo *TagUpdateOne) ClearGroups() *TagUpdateOne {
	tuo.mutation.ClearGroups()
	return tuo
}

// RemoveGroupIDs removes the "groups" edge to Group entities by IDs.
func (tuo *TagUpdateOne) RemoveGroupIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveGroupIDs(ids...)
	return tuo
}

// RemoveGroups removes "groups" edges to Group entities.
func (tuo *TagUpdateOne) RemoveGroups(g ...*Group) *TagUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveGroupIDs(ids...)
}

// ClearTweetTags clears all "tweet_tags" edges to the TweetTag entity.
func (tuo *TagUpdateOne) ClearTweetTags() *TagUpdateOne {
	tuo.mutation.ClearTweetTags()
	return tuo
}

// RemoveTweetTagIDs removes the "tweet_tags" edge to TweetTag entities by IDs.
func (tuo *TagUpdateOne) RemoveTweetTagIDs(ids ...uuid.UUID) *TagUpdateOne {
	tuo.mutation.RemoveTweetTagIDs(ids...)
	return tuo
}

// RemoveTweetTags removes "tweet_tags" edges to TweetTag entities.
func (tuo *TagUpdateOne) RemoveTweetTags(t ...*TweetTag) *TagUpdateOne {
	ids := make([]uuid.UUID, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return tuo.RemoveTweetTagIDs(ids...)
}

// ClearGroupTags clears all "group_tags" edges to the GroupTag entity.
func (tuo *TagUpdateOne) ClearGroupTags() *TagUpdateOne {
	tuo.mutation.ClearGroupTags()
	return tuo
}

// RemoveGroupTagIDs removes the "group_tags" edge to GroupTag entities by IDs.
func (tuo *TagUpdateOne) RemoveGroupTagIDs(ids ...int) *TagUpdateOne {
	tuo.mutation.RemoveGroupTagIDs(ids...)
	return tuo
}

// RemoveGroupTags removes "group_tags" edges to GroupTag entities.
func (tuo *TagUpdateOne) RemoveGroupTags(g ...*GroupTag) *TagUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return tuo.RemoveGroupTagIDs(ids...)
}

// Where appends a list predicates to the TagUpdate builder.
func (tuo *TagUpdateOne) Where(ps ...predicate.Tag) *TagUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := sqlgraph.NewUpdateSpec(tag.Table, tag.Columns, sqlgraph.NewFieldSpec(tag.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Tag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.Value(); ok {
		_spec.SetField(tag.FieldValue, field.TypeString, value)
	}
	if tuo.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTweetsIDs(); len(nodes) > 0 && !tuo.mutation.TweetsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TweetsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.TweetsTable,
			Columns: tag.TweetsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweet.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		createE := &TweetTagCreate{config: tuo.config, mutation: newTweetTagMutation(tuo.config, OpCreate)}
		createE.defaults()
		_, specE := createE.createSpec()
		edge.Target.Fields = specE.Fields
		if specE.ID.Value != nil {
			edge.Target.Fields = append(edge.Target.Fields, specE.ID)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedGroupsIDs(); len(nodes) > 0 && !tuo.mutation.GroupsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GroupsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tag.GroupsTable,
			Columns: tag.GroupsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(group.FieldID, field.TypeInt),
			},
			RefRequired: false,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedTweetTagsIDs(); len(nodes) > 0 && !tuo.mutation.TweetTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.TweetTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.TweetTagsTable,
			Columns: []string{tag.TweetTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(tweettag.FieldID, field.TypeUUID),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if tuo.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedGroupTagsIDs(); len(nodes) > 0 && !tuo.mutation.GroupTagsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.GroupTagsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   tag.GroupTagsTable,
			Columns: []string{tag.GroupTagsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(grouptag.FieldID, field.TypeInt),
			},
			RefRequired: true,
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
