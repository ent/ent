// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"
	errors "errors"
	fmt "fmt"
	time "time"

	"entgo.io/ent/dialect/sql"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	predicate "entgo.io/ent/entc/integration/edgeschema/ent/predicate"
	tag "entgo.io/ent/entc/integration/edgeschema/ent/tag"
	tweet "entgo.io/ent/entc/integration/edgeschema/ent/tweet"
	tweettag "entgo.io/ent/entc/integration/edgeschema/ent/tweettag"
	field "entgo.io/ent/schema/field"
)

// TweetTagUpdate is the builder for updating TweetTag entities.
type TweetTagUpdate struct {
	config
	hooks    []Hook
	mutation *TweetTagMutation
}

// Where appends a list predicates to the TweetTagUpdate builder.
func (ttu *TweetTagUpdate) Where(ps ...predicate.TweetTag) *TweetTagUpdate {
	ttu.mutation.Where(ps...)
	return ttu
}

// SetAddedAt sets the "added_at" field.
func (ttu *TweetTagUpdate) SetAddedAt(t time.Time) *TweetTagUpdate {
	ttu.mutation.SetAddedAt(t)
	return ttu
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (ttu *TweetTagUpdate) SetNillableAddedAt(t *time.Time) *TweetTagUpdate {
	if t != nil {
		ttu.SetAddedAt(*t)
	}
	return ttu
}

// SetTagID sets the "tag_id" field.
func (ttu *TweetTagUpdate) SetTagID(i int) *TweetTagUpdate {
	ttu.mutation.SetTagID(i)
	return ttu
}

// SetTweetID sets the "tweet_id" field.
func (ttu *TweetTagUpdate) SetTweetID(i int) *TweetTagUpdate {
	ttu.mutation.SetTweetID(i)
	return ttu
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttu *TweetTagUpdate) SetTag(t *Tag) *TweetTagUpdate {
	return ttu.SetTagID(t.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (ttu *TweetTagUpdate) SetTweet(t *Tweet) *TweetTagUpdate {
	return ttu.SetTweetID(t.ID)
}

// Mutation returns the TweetTagMutation object of the builder.
func (ttu *TweetTagUpdate) Mutation() *TweetTagMutation {
	return ttu.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttu *TweetTagUpdate) ClearTag() *TweetTagUpdate {
	ttu.mutation.ClearTag()
	return ttu
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (ttu *TweetTagUpdate) ClearTweet() *TweetTagUpdate {
	ttu.mutation.ClearTweet()
	return ttu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ttu *TweetTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, TweetTagMutation](ctx, ttu.sqlSave, ttu.mutation, ttu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ttu *TweetTagUpdate) SaveX(ctx context.Context) int {
	affected, err := ttu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ttu *TweetTagUpdate) Exec(ctx context.Context) error {
	_, err := ttu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttu *TweetTagUpdate) ExecX(ctx context.Context) {
	if err := ttu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttu *TweetTagUpdate) check() error {
	if _, ok := ttu.mutation.TagID(); ttu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TweetTag.tag"`)
	}
	if _, ok := ttu.mutation.TweetID(); ttu.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TweetTag.tweet"`)
	}
	return nil
}

func (ttu *TweetTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ttu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweettag.Table,
			Columns: tweettag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		},
	}
	if ps := ttu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttu.mutation.AddedAt(); ok {
		_spec.SetField(tweettag.FieldAddedAt, field.TypeTime, value)
	}
	if ttu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttu.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttu.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ttu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweettag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ttu.mutation.done = true
	return n, nil
}

// TweetTagUpdateOne is the builder for updating a single TweetTag entity.
type TweetTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TweetTagMutation
}

// SetAddedAt sets the "added_at" field.
func (ttuo *TweetTagUpdateOne) SetAddedAt(t time.Time) *TweetTagUpdateOne {
	ttuo.mutation.SetAddedAt(t)
	return ttuo
}

// SetNillableAddedAt sets the "added_at" field if the given value is not nil.
func (ttuo *TweetTagUpdateOne) SetNillableAddedAt(t *time.Time) *TweetTagUpdateOne {
	if t != nil {
		ttuo.SetAddedAt(*t)
	}
	return ttuo
}

// SetTagID sets the "tag_id" field.
func (ttuo *TweetTagUpdateOne) SetTagID(i int) *TweetTagUpdateOne {
	ttuo.mutation.SetTagID(i)
	return ttuo
}

// SetTweetID sets the "tweet_id" field.
func (ttuo *TweetTagUpdateOne) SetTweetID(i int) *TweetTagUpdateOne {
	ttuo.mutation.SetTweetID(i)
	return ttuo
}

// SetTag sets the "tag" edge to the Tag entity.
func (ttuo *TweetTagUpdateOne) SetTag(t *Tag) *TweetTagUpdateOne {
	return ttuo.SetTagID(t.ID)
}

// SetTweet sets the "tweet" edge to the Tweet entity.
func (ttuo *TweetTagUpdateOne) SetTweet(t *Tweet) *TweetTagUpdateOne {
	return ttuo.SetTweetID(t.ID)
}

// Mutation returns the TweetTagMutation object of the builder.
func (ttuo *TweetTagUpdateOne) Mutation() *TweetTagMutation {
	return ttuo.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (ttuo *TweetTagUpdateOne) ClearTag() *TweetTagUpdateOne {
	ttuo.mutation.ClearTag()
	return ttuo
}

// ClearTweet clears the "tweet" edge to the Tweet entity.
func (ttuo *TweetTagUpdateOne) ClearTweet() *TweetTagUpdateOne {
	ttuo.mutation.ClearTweet()
	return ttuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ttuo *TweetTagUpdateOne) Select(field string, fields ...string) *TweetTagUpdateOne {
	ttuo.fields = append([]string{field}, fields...)
	return ttuo
}

// Save executes the query and returns the updated TweetTag entity.
func (ttuo *TweetTagUpdateOne) Save(ctx context.Context) (*TweetTag, error) {
	return withHooks[*TweetTag, TweetTagMutation](ctx, ttuo.sqlSave, ttuo.mutation, ttuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ttuo *TweetTagUpdateOne) SaveX(ctx context.Context) *TweetTag {
	node, err := ttuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ttuo *TweetTagUpdateOne) Exec(ctx context.Context) error {
	_, err := ttuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ttuo *TweetTagUpdateOne) ExecX(ctx context.Context) {
	if err := ttuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ttuo *TweetTagUpdateOne) check() error {
	if _, ok := ttuo.mutation.TagID(); ttuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TweetTag.tag"`)
	}
	if _, ok := ttuo.mutation.TweetID(); ttuo.mutation.TweetCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "TweetTag.tweet"`)
	}
	return nil
}

func (ttuo *TweetTagUpdateOne) sqlSave(ctx context.Context) (_node *TweetTag, err error) {
	if err := ttuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tweettag.Table,
			Columns: tweettag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: tweettag.FieldID,
			},
		},
	}
	id, ok := ttuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TweetTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ttuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tweettag.FieldID)
		for _, f := range fields {
			if !tweettag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tweettag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ttuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ttuo.mutation.AddedAt(); ok {
		_spec.SetField(tweettag.FieldAddedAt, field.TypeTime, value)
	}
	if ttuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TagTable,
			Columns: []string{tweettag.TagColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tag.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ttuo.mutation.TweetCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: tweet.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ttuo.mutation.TweetIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   tweettag.TweetTable,
			Columns: []string{tweettag.TweetColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &TweetTag{config: ttuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ttuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tweettag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ttuo.mutation.done = true
	return _node, nil
}
