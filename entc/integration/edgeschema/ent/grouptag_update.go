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
	"entgo.io/ent/schema/field"
)

// GroupTagUpdate is the builder for updating GroupTag entities.
type GroupTagUpdate struct {
	config
	hooks    []Hook
	mutation *GroupTagMutation
}

// Where appends a list predicates to the GroupTagUpdate builder.
func (gtu *GroupTagUpdate) Where(ps ...predicate.GroupTag) *GroupTagUpdate {
	gtu.mutation.Where(ps...)
	return gtu
}

// SetTagID sets the "tag_id" field.
func (gtu *GroupTagUpdate) SetTagID(i int) *GroupTagUpdate {
	gtu.mutation.SetTagID(i)
	return gtu
}

// SetGroupID sets the "group_id" field.
func (gtu *GroupTagUpdate) SetGroupID(i int) *GroupTagUpdate {
	gtu.mutation.SetGroupID(i)
	return gtu
}

// SetTag sets the "tag" edge to the Tag entity.
func (gtu *GroupTagUpdate) SetTag(t *Tag) *GroupTagUpdate {
	return gtu.SetTagID(t.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gtu *GroupTagUpdate) SetGroup(g *Group) *GroupTagUpdate {
	return gtu.SetGroupID(g.ID)
}

// Mutation returns the GroupTagMutation object of the builder.
func (gtu *GroupTagUpdate) Mutation() *GroupTagMutation {
	return gtu.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (gtu *GroupTagUpdate) ClearTag() *GroupTagUpdate {
	gtu.mutation.ClearTag()
	return gtu
}

// ClearGroup clears the "group" edge to the Group entity.
func (gtu *GroupTagUpdate) ClearGroup() *GroupTagUpdate {
	gtu.mutation.ClearGroup()
	return gtu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (gtu *GroupTagUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, GroupTagMutation](ctx, gtu.sqlSave, gtu.mutation, gtu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtu *GroupTagUpdate) SaveX(ctx context.Context) int {
	affected, err := gtu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gtu *GroupTagUpdate) Exec(ctx context.Context) error {
	_, err := gtu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtu *GroupTagUpdate) ExecX(ctx context.Context) {
	if err := gtu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtu *GroupTagUpdate) check() error {
	if _, ok := gtu.mutation.TagID(); gtu.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupTag.tag"`)
	}
	if _, ok := gtu.mutation.GroupID(); gtu.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupTag.group"`)
	}
	return nil
}

func (gtu *GroupTagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := gtu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouptag.Table,
			Columns: grouptag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grouptag.FieldID,
			},
		},
	}
	if ps := gtu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gtu.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.TagTable,
			Columns: []string{grouptag.TagColumn},
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
	if nodes := gtu.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.TagTable,
			Columns: []string{grouptag.TagColumn},
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
	if gtu.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.GroupTable,
			Columns: []string{grouptag.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtu.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.GroupTable,
			Columns: []string{grouptag.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, gtu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouptag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	gtu.mutation.done = true
	return n, nil
}

// GroupTagUpdateOne is the builder for updating a single GroupTag entity.
type GroupTagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *GroupTagMutation
}

// SetTagID sets the "tag_id" field.
func (gtuo *GroupTagUpdateOne) SetTagID(i int) *GroupTagUpdateOne {
	gtuo.mutation.SetTagID(i)
	return gtuo
}

// SetGroupID sets the "group_id" field.
func (gtuo *GroupTagUpdateOne) SetGroupID(i int) *GroupTagUpdateOne {
	gtuo.mutation.SetGroupID(i)
	return gtuo
}

// SetTag sets the "tag" edge to the Tag entity.
func (gtuo *GroupTagUpdateOne) SetTag(t *Tag) *GroupTagUpdateOne {
	return gtuo.SetTagID(t.ID)
}

// SetGroup sets the "group" edge to the Group entity.
func (gtuo *GroupTagUpdateOne) SetGroup(g *Group) *GroupTagUpdateOne {
	return gtuo.SetGroupID(g.ID)
}

// Mutation returns the GroupTagMutation object of the builder.
func (gtuo *GroupTagUpdateOne) Mutation() *GroupTagMutation {
	return gtuo.mutation
}

// ClearTag clears the "tag" edge to the Tag entity.
func (gtuo *GroupTagUpdateOne) ClearTag() *GroupTagUpdateOne {
	gtuo.mutation.ClearTag()
	return gtuo
}

// ClearGroup clears the "group" edge to the Group entity.
func (gtuo *GroupTagUpdateOne) ClearGroup() *GroupTagUpdateOne {
	gtuo.mutation.ClearGroup()
	return gtuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (gtuo *GroupTagUpdateOne) Select(field string, fields ...string) *GroupTagUpdateOne {
	gtuo.fields = append([]string{field}, fields...)
	return gtuo
}

// Save executes the query and returns the updated GroupTag entity.
func (gtuo *GroupTagUpdateOne) Save(ctx context.Context) (*GroupTag, error) {
	return withHooks[*GroupTag, GroupTagMutation](ctx, gtuo.sqlSave, gtuo.mutation, gtuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (gtuo *GroupTagUpdateOne) SaveX(ctx context.Context) *GroupTag {
	node, err := gtuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (gtuo *GroupTagUpdateOne) Exec(ctx context.Context) error {
	_, err := gtuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gtuo *GroupTagUpdateOne) ExecX(ctx context.Context) {
	if err := gtuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gtuo *GroupTagUpdateOne) check() error {
	if _, ok := gtuo.mutation.TagID(); gtuo.mutation.TagCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupTag.tag"`)
	}
	if _, ok := gtuo.mutation.GroupID(); gtuo.mutation.GroupCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "GroupTag.group"`)
	}
	return nil
}

func (gtuo *GroupTagUpdateOne) sqlSave(ctx context.Context) (_node *GroupTag, err error) {
	if err := gtuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   grouptag.Table,
			Columns: grouptag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: grouptag.FieldID,
			},
		},
	}
	id, ok := gtuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "GroupTag.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := gtuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, grouptag.FieldID)
		for _, f := range fields {
			if !grouptag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != grouptag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := gtuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if gtuo.mutation.TagCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.TagTable,
			Columns: []string{grouptag.TagColumn},
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
	if nodes := gtuo.mutation.TagIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.TagTable,
			Columns: []string{grouptag.TagColumn},
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
	if gtuo.mutation.GroupCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.GroupTable,
			Columns: []string{grouptag.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := gtuo.mutation.GroupIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   grouptag.GroupTable,
			Columns: []string{grouptag.GroupColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: group.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &GroupTag{config: gtuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, gtuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{grouptag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	gtuo.mutation.done = true
	return _node, nil
}
