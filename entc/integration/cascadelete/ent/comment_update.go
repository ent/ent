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
	"entgo.io/ent/entc/integration/cascadelete/ent/comment"
	"entgo.io/ent/entc/integration/cascadelete/ent/post"
	"entgo.io/ent/entc/integration/cascadelete/ent/predicate"
	"entgo.io/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where appends a list predicates to the CommentUpdate builder.
func (u *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	u.mutation.Where(ps...)
	return u
}

// SetText sets the "text" field.
func (m *CommentUpdate) SetText(v string) *CommentUpdate {
	m.mutation.SetText(v)
	return m
}

// SetNillableText sets the "text" field if the given value is not nil.
func (m *CommentUpdate) SetNillableText(v *string) *CommentUpdate {
	if v != nil {
		m.SetText(*v)
	}
	return m
}

// SetPostID sets the "post_id" field.
func (m *CommentUpdate) SetPostID(v int) *CommentUpdate {
	m.mutation.SetPostID(v)
	return m
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (m *CommentUpdate) SetNillablePostID(v *int) *CommentUpdate {
	if v != nil {
		m.SetPostID(*v)
	}
	return m
}

// SetPost sets the "post" edge to the Post entity.
func (m *CommentUpdate) SetPost(v *Post) *CommentUpdate {
	return m.SetPostID(v.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (m *CommentUpdate) Mutation() *CommentMutation {
	return m.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (u *CommentUpdate) ClearPost() *CommentUpdate {
	u.mutation.ClearPost()
	return u
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (u *CommentUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (u *CommentUpdate) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpdate) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *CommentUpdate) check() error {
	if u.mutation.PostCleared() && len(u.mutation.PostIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Comment.post"`)
	}
	return nil
}

func (u *CommentUpdate) sqlSave(ctx context.Context) (_n int, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	if ps := u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := u.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if u.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if _n, err = sqlgraph.UpdateNodes(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	u.mutation.done = true
	return _n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CommentMutation
}

// SetText sets the "text" field.
func (m *CommentUpdateOne) SetText(v string) *CommentUpdateOne {
	m.mutation.SetText(v)
	return m
}

// SetNillableText sets the "text" field if the given value is not nil.
func (m *CommentUpdateOne) SetNillableText(v *string) *CommentUpdateOne {
	if v != nil {
		m.SetText(*v)
	}
	return m
}

// SetPostID sets the "post_id" field.
func (m *CommentUpdateOne) SetPostID(v int) *CommentUpdateOne {
	m.mutation.SetPostID(v)
	return m
}

// SetNillablePostID sets the "post_id" field if the given value is not nil.
func (m *CommentUpdateOne) SetNillablePostID(v *int) *CommentUpdateOne {
	if v != nil {
		m.SetPostID(*v)
	}
	return m
}

// SetPost sets the "post" edge to the Post entity.
func (m *CommentUpdateOne) SetPost(v *Post) *CommentUpdateOne {
	return m.SetPostID(v.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (m *CommentUpdateOne) Mutation() *CommentMutation {
	return m.mutation
}

// ClearPost clears the "post" edge to the Post entity.
func (u *CommentUpdateOne) ClearPost() *CommentUpdateOne {
	u.mutation.ClearPost()
	return u
}

// Where appends a list predicates to the CommentUpdate builder.
func (u *CommentUpdateOne) Where(ps ...predicate.Comment) *CommentUpdateOne {
	u.mutation.Where(ps...)
	return u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (u *CommentUpdateOne) Select(field string, fields ...string) *CommentUpdateOne {
	u.fields = append([]string{field}, fields...)
	return u
}

// Save executes the query and returns the updated Comment entity.
func (u *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, u.sqlSave, u.mutation, u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (u *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (u *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (u *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := u.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (u *CommentUpdateOne) check() error {
	if u.mutation.PostCleared() && len(u.mutation.PostIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "Comment.post"`)
	}
	return nil
}

func (u *CommentUpdateOne) sqlSave(ctx context.Context) (_n *Comment, err error) {
	if err := u.check(); err != nil {
		return _n, err
	}
	_spec := sqlgraph.NewUpdateSpec(comment.Table, comment.Columns, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeInt))
	id, ok := u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Comment.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, comment.FieldID)
		for _, f := range fields {
			if !comment.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != comment.FieldID {
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
		_spec.SetField(comment.FieldText, field.TypeString, value)
	}
	if u.mutation.PostCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := u.mutation.PostIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.PostTable,
			Columns: []string{comment.PostColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_n = &Comment{config: u.config}
	_spec.Assign = _n.assignValues
	_spec.ScanValues = _n.scanValues
	if err = sqlgraph.UpdateNode(ctx, u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	u.mutation.done = true
	return _n, nil
}
