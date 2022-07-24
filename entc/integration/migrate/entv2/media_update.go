// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/media"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// MediaUpdate is the builder for updating Media entities.
type MediaUpdate struct {
	config
	hooks    []Hook
	mutation *MediaMutation
}

// Where appends a list predicates to the MediaUpdate builder.
func (mu *MediaUpdate) Where(ps ...predicate.Media) *MediaUpdate {
	mu.mutation.Where(ps...)
	return mu
}

// SetSource sets the "source" field.
func (mu *MediaUpdate) SetSource(s string) *MediaUpdate {
	mu.mutation.SetSource(s)
	return mu
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableSource(s *string) *MediaUpdate {
	if s != nil {
		mu.SetSource(*s)
	}
	return mu
}

// ClearSource clears the value of the "source" field.
func (mu *MediaUpdate) ClearSource() *MediaUpdate {
	mu.mutation.ClearSource()
	return mu
}

// SetSourceURI sets the "source_uri" field.
func (mu *MediaUpdate) SetSourceURI(s string) *MediaUpdate {
	mu.mutation.SetSourceURI(s)
	return mu
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableSourceURI(s *string) *MediaUpdate {
	if s != nil {
		mu.SetSourceURI(*s)
	}
	return mu
}

// ClearSourceURI clears the value of the "source_uri" field.
func (mu *MediaUpdate) ClearSourceURI() *MediaUpdate {
	mu.mutation.ClearSourceURI()
	return mu
}

// SetText sets the "text" field.
func (mu *MediaUpdate) SetText(s string) *MediaUpdate {
	mu.mutation.SetText(s)
	return mu
}

// SetNillableText sets the "text" field if the given value is not nil.
func (mu *MediaUpdate) SetNillableText(s *string) *MediaUpdate {
	if s != nil {
		mu.SetText(*s)
	}
	return mu
}

// ClearText clears the value of the "text" field.
func (mu *MediaUpdate) ClearText() *MediaUpdate {
	mu.mutation.ClearText()
	return mu
}

// Mutation returns the MediaMutation object of the builder.
func (mu *MediaUpdate) Mutation() *MediaMutation {
	return mu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (mu *MediaUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			if mu.hooks[i] == nil {
				return 0, errors.New("entv2: uninitialized hook (forgotten import entv2/runtime?)")
			}
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MediaUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MediaUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MediaUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MediaUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   media.Table,
			Columns: media.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: media.FieldID,
			},
		},
	}
	if ps := mu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSource,
		})
	}
	if mu.mutation.SourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldSource,
		})
	}
	if value, ok := mu.mutation.SourceURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSourceURI,
		})
	}
	if mu.mutation.SourceURICleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldSourceURI,
		})
	}
	if value, ok := mu.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldText,
		})
	}
	if mu.mutation.TextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldText,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// MediaUpdateOne is the builder for updating a single Media entity.
type MediaUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *MediaMutation
}

// SetSource sets the "source" field.
func (muo *MediaUpdateOne) SetSource(s string) *MediaUpdateOne {
	muo.mutation.SetSource(s)
	return muo
}

// SetNillableSource sets the "source" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableSource(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetSource(*s)
	}
	return muo
}

// ClearSource clears the value of the "source" field.
func (muo *MediaUpdateOne) ClearSource() *MediaUpdateOne {
	muo.mutation.ClearSource()
	return muo
}

// SetSourceURI sets the "source_uri" field.
func (muo *MediaUpdateOne) SetSourceURI(s string) *MediaUpdateOne {
	muo.mutation.SetSourceURI(s)
	return muo
}

// SetNillableSourceURI sets the "source_uri" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableSourceURI(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetSourceURI(*s)
	}
	return muo
}

// ClearSourceURI clears the value of the "source_uri" field.
func (muo *MediaUpdateOne) ClearSourceURI() *MediaUpdateOne {
	muo.mutation.ClearSourceURI()
	return muo
}

// SetText sets the "text" field.
func (muo *MediaUpdateOne) SetText(s string) *MediaUpdateOne {
	muo.mutation.SetText(s)
	return muo
}

// SetNillableText sets the "text" field if the given value is not nil.
func (muo *MediaUpdateOne) SetNillableText(s *string) *MediaUpdateOne {
	if s != nil {
		muo.SetText(*s)
	}
	return muo
}

// ClearText clears the value of the "text" field.
func (muo *MediaUpdateOne) ClearText() *MediaUpdateOne {
	muo.mutation.ClearText()
	return muo
}

// Mutation returns the MediaMutation object of the builder.
func (muo *MediaUpdateOne) Mutation() *MediaMutation {
	return muo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (muo *MediaUpdateOne) Select(field string, fields ...string) *MediaUpdateOne {
	muo.fields = append([]string{field}, fields...)
	return muo
}

// Save executes the query and returns the updated Media entity.
func (muo *MediaUpdateOne) Save(ctx context.Context) (*Media, error) {
	var (
		err  error
		node *Media
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MediaMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			if muo.hooks[i] == nil {
				return nil, errors.New("entv2: uninitialized hook (forgotten import entv2/runtime?)")
			}
			mut = muo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, muo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Media)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from MediaMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MediaUpdateOne) SaveX(ctx context.Context) *Media {
	node, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (muo *MediaUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MediaUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MediaUpdateOne) sqlSave(ctx context.Context) (_node *Media, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   media.Table,
			Columns: media.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: media.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "Media.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := muo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, media.FieldID)
		for _, f := range fields {
			if !media.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != media.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := muo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := muo.mutation.Source(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSource,
		})
	}
	if muo.mutation.SourceCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldSource,
		})
	}
	if value, ok := muo.mutation.SourceURI(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldSourceURI,
		})
	}
	if muo.mutation.SourceURICleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldSourceURI,
		})
	}
	if value, ok := muo.mutation.Text(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: media.FieldText,
		})
	}
	if muo.mutation.TextCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: media.FieldText,
		})
	}
	_node = &Media{config: muo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{media.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
