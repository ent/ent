// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/entc/integration/ent/comment"
	"github.com/facebook/ent/entc/integration/ent/predicate"
	"github.com/facebook/ent/schema/field"
)

// CommentUpdate is the builder for updating Comment entities.
type CommentUpdate struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// Where adds a new predicate for the builder.
func (cu *CommentUpdate) Where(ps ...predicate.Comment) *CommentUpdate {
	cu.mutation.predicates = append(cu.mutation.predicates, ps...)
	return cu
}

// SetUniqueInt sets the unique_int field.
func (cu *CommentUpdate) SetUniqueInt(i int) *CommentUpdate {
	cu.mutation.ResetUniqueInt()
	cu.mutation.SetUniqueInt(i)
	return cu
}

// AddUniqueInt adds i to unique_int.
func (cu *CommentUpdate) AddUniqueInt(i int) *CommentUpdate {
	cu.mutation.AddUniqueInt(i)
	return cu
}

// SetUniqueFloat sets the unique_float field.
func (cu *CommentUpdate) SetUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.ResetUniqueFloat()
	cu.mutation.SetUniqueFloat(f)
	return cu
}

// AddUniqueFloat adds f to unique_float.
func (cu *CommentUpdate) AddUniqueFloat(f float64) *CommentUpdate {
	cu.mutation.AddUniqueFloat(f)
	return cu
}

// SetNillableInt sets the nillable_int field.
func (cu *CommentUpdate) SetNillableInt(i int) *CommentUpdate {
	cu.mutation.ResetNillableInt()
	cu.mutation.SetNillableInt(i)
	return cu
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (cu *CommentUpdate) SetNillableNillableInt(i *int) *CommentUpdate {
	if i != nil {
		cu.SetNillableInt(*i)
	}
	return cu
}

// AddNillableInt adds i to nillable_int.
func (cu *CommentUpdate) AddNillableInt(i int) *CommentUpdate {
	cu.mutation.AddNillableInt(i)
	return cu
}

// ClearNillableInt clears the value of nillable_int.
func (cu *CommentUpdate) ClearNillableInt() *CommentUpdate {
	cu.mutation.ClearNillableInt()
	return cu
}

// Mutation returns the CommentMutation object of the builder.
func (cu *CommentUpdate) Mutation() *CommentMutation {
	return cu.mutation
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CommentUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CommentUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CommentUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CommentUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CommentUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UniqueInt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUniqueInt,
		})
	}
	if value, ok := cu.mutation.AddedUniqueInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUniqueInt,
		})
	}
	if value, ok := cu.mutation.UniqueFloat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: comment.FieldUniqueFloat,
		})
	}
	if value, ok := cu.mutation.AddedUniqueFloat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: comment.FieldUniqueFloat,
		})
	}
	if value, ok := cu.mutation.NillableInt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldNillableInt,
		})
	}
	if value, ok := cu.mutation.AddedNillableInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldNillableInt,
		})
	}
	if cu.mutation.NillableIntCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: comment.FieldNillableInt,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CommentUpdateOne is the builder for updating a single Comment entity.
type CommentUpdateOne struct {
	config
	hooks    []Hook
	mutation *CommentMutation
}

// SetUniqueInt sets the unique_int field.
func (cuo *CommentUpdateOne) SetUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetUniqueInt()
	cuo.mutation.SetUniqueInt(i)
	return cuo
}

// AddUniqueInt adds i to unique_int.
func (cuo *CommentUpdateOne) AddUniqueInt(i int) *CommentUpdateOne {
	cuo.mutation.AddUniqueInt(i)
	return cuo
}

// SetUniqueFloat sets the unique_float field.
func (cuo *CommentUpdateOne) SetUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.ResetUniqueFloat()
	cuo.mutation.SetUniqueFloat(f)
	return cuo
}

// AddUniqueFloat adds f to unique_float.
func (cuo *CommentUpdateOne) AddUniqueFloat(f float64) *CommentUpdateOne {
	cuo.mutation.AddUniqueFloat(f)
	return cuo
}

// SetNillableInt sets the nillable_int field.
func (cuo *CommentUpdateOne) SetNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.ResetNillableInt()
	cuo.mutation.SetNillableInt(i)
	return cuo
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (cuo *CommentUpdateOne) SetNillableNillableInt(i *int) *CommentUpdateOne {
	if i != nil {
		cuo.SetNillableInt(*i)
	}
	return cuo
}

// AddNillableInt adds i to nillable_int.
func (cuo *CommentUpdateOne) AddNillableInt(i int) *CommentUpdateOne {
	cuo.mutation.AddNillableInt(i)
	return cuo
}

// ClearNillableInt clears the value of nillable_int.
func (cuo *CommentUpdateOne) ClearNillableInt() *CommentUpdateOne {
	cuo.mutation.ClearNillableInt()
	return cuo
}

// Mutation returns the CommentMutation object of the builder.
func (cuo *CommentUpdateOne) Mutation() *CommentMutation {
	return cuo.mutation
}

// Save executes the query and returns the updated entity.
func (cuo *CommentUpdateOne) Save(ctx context.Context) (*Comment, error) {
	var (
		err  error
		node *Comment
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CommentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CommentUpdateOne) SaveX(ctx context.Context) *Comment {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *CommentUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CommentUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CommentUpdateOne) sqlSave(ctx context.Context) (_node *Comment, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   comment.Table,
			Columns: comment.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: comment.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Comment.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.UniqueInt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUniqueInt,
		})
	}
	if value, ok := cuo.mutation.AddedUniqueInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldUniqueInt,
		})
	}
	if value, ok := cuo.mutation.UniqueFloat(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: comment.FieldUniqueFloat,
		})
	}
	if value, ok := cuo.mutation.AddedUniqueFloat(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: comment.FieldUniqueFloat,
		})
	}
	if value, ok := cuo.mutation.NillableInt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldNillableInt,
		})
	}
	if value, ok := cuo.mutation.AddedNillableInt(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: comment.FieldNillableInt,
		})
	}
	if cuo.mutation.NillableIntCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: comment.FieldNillableInt,
		})
	}
	_node = &Comment{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{comment.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return _node, nil
}
