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
	"entgo.io/ent/entc/integration/edgeschema/ent/relationship"
	"entgo.io/ent/entc/integration/edgeschema/ent/relationshipinfo"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// RelationshipUpdate is the builder for updating Relationship entities.
type RelationshipUpdate struct {
	config
	hooks    []Hook
	mutation *RelationshipMutation
}

// Where appends a list predicates to the RelationshipUpdate builder.
func (ru *RelationshipUpdate) Where(ps ...predicate.Relationship) *RelationshipUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetWeight sets the "weight" field.
func (ru *RelationshipUpdate) SetWeight(i int) *RelationshipUpdate {
	ru.mutation.ResetWeight()
	ru.mutation.SetWeight(i)
	return ru
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (ru *RelationshipUpdate) SetNillableWeight(i *int) *RelationshipUpdate {
	if i != nil {
		ru.SetWeight(*i)
	}
	return ru
}

// AddWeight adds i to the "weight" field.
func (ru *RelationshipUpdate) AddWeight(i int) *RelationshipUpdate {
	ru.mutation.AddWeight(i)
	return ru
}

// SetUserID sets the "user_id" field.
func (ru *RelationshipUpdate) SetUserID(i int) *RelationshipUpdate {
	ru.mutation.SetUserID(i)
	return ru
}

// SetRelativeID sets the "relative_id" field.
func (ru *RelationshipUpdate) SetRelativeID(i int) *RelationshipUpdate {
	ru.mutation.SetRelativeID(i)
	return ru
}

// SetInfoID sets the "info_id" field.
func (ru *RelationshipUpdate) SetInfoID(i int) *RelationshipUpdate {
	ru.mutation.SetInfoID(i)
	return ru
}

// SetNillableInfoID sets the "info_id" field if the given value is not nil.
func (ru *RelationshipUpdate) SetNillableInfoID(i *int) *RelationshipUpdate {
	if i != nil {
		ru.SetInfoID(*i)
	}
	return ru
}

// ClearInfoID clears the value of the "info_id" field.
func (ru *RelationshipUpdate) ClearInfoID() *RelationshipUpdate {
	ru.mutation.ClearInfoID()
	return ru
}

// SetUser sets the "user" edge to the User entity.
func (ru *RelationshipUpdate) SetUser(u *User) *RelationshipUpdate {
	return ru.SetUserID(u.ID)
}

// SetRelative sets the "relative" edge to the User entity.
func (ru *RelationshipUpdate) SetRelative(u *User) *RelationshipUpdate {
	return ru.SetRelativeID(u.ID)
}

// SetInfo sets the "info" edge to the RelationshipInfo entity.
func (ru *RelationshipUpdate) SetInfo(r *RelationshipInfo) *RelationshipUpdate {
	return ru.SetInfoID(r.ID)
}

// Mutation returns the RelationshipMutation object of the builder.
func (ru *RelationshipUpdate) Mutation() *RelationshipMutation {
	return ru.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ru *RelationshipUpdate) ClearUser() *RelationshipUpdate {
	ru.mutation.ClearUser()
	return ru
}

// ClearRelative clears the "relative" edge to the User entity.
func (ru *RelationshipUpdate) ClearRelative() *RelationshipUpdate {
	ru.mutation.ClearRelative()
	return ru
}

// ClearInfo clears the "info" edge to the RelationshipInfo entity.
func (ru *RelationshipUpdate) ClearInfo() *RelationshipUpdate {
	ru.mutation.ClearInfo()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RelationshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, RelationshipMutation](ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RelationshipUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RelationshipUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RelationshipUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RelationshipUpdate) check() error {
	if _, ok := ru.mutation.UserID(); ru.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Relationship.user"`)
	}
	if _, ok := ru.mutation.RelativeID(); ru.mutation.RelativeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Relationship.relative"`)
	}
	return nil
}

func (ru *RelationshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relationship.Table,
			Columns: relationship.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: relationship.FieldUserID,
				},
				{
					Type:   field.TypeInt,
					Column: relationship.FieldRelativeID,
				},
			},
		},
	}
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Weight(); ok {
		_spec.SetField(relationship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedWeight(); ok {
		_spec.AddField(relationship.FieldWeight, field.TypeInt, value)
	}
	if ru.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.RelativeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RelativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.InfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.InfoTable,
			Columns: []string{relationship.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relationshipinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.InfoTable,
			Columns: []string{relationship.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relationshipinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RelationshipUpdateOne is the builder for updating a single Relationship entity.
type RelationshipUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RelationshipMutation
}

// SetWeight sets the "weight" field.
func (ruo *RelationshipUpdateOne) SetWeight(i int) *RelationshipUpdateOne {
	ruo.mutation.ResetWeight()
	ruo.mutation.SetWeight(i)
	return ruo
}

// SetNillableWeight sets the "weight" field if the given value is not nil.
func (ruo *RelationshipUpdateOne) SetNillableWeight(i *int) *RelationshipUpdateOne {
	if i != nil {
		ruo.SetWeight(*i)
	}
	return ruo
}

// AddWeight adds i to the "weight" field.
func (ruo *RelationshipUpdateOne) AddWeight(i int) *RelationshipUpdateOne {
	ruo.mutation.AddWeight(i)
	return ruo
}

// SetUserID sets the "user_id" field.
func (ruo *RelationshipUpdateOne) SetUserID(i int) *RelationshipUpdateOne {
	ruo.mutation.SetUserID(i)
	return ruo
}

// SetRelativeID sets the "relative_id" field.
func (ruo *RelationshipUpdateOne) SetRelativeID(i int) *RelationshipUpdateOne {
	ruo.mutation.SetRelativeID(i)
	return ruo
}

// SetInfoID sets the "info_id" field.
func (ruo *RelationshipUpdateOne) SetInfoID(i int) *RelationshipUpdateOne {
	ruo.mutation.SetInfoID(i)
	return ruo
}

// SetNillableInfoID sets the "info_id" field if the given value is not nil.
func (ruo *RelationshipUpdateOne) SetNillableInfoID(i *int) *RelationshipUpdateOne {
	if i != nil {
		ruo.SetInfoID(*i)
	}
	return ruo
}

// ClearInfoID clears the value of the "info_id" field.
func (ruo *RelationshipUpdateOne) ClearInfoID() *RelationshipUpdateOne {
	ruo.mutation.ClearInfoID()
	return ruo
}

// SetUser sets the "user" edge to the User entity.
func (ruo *RelationshipUpdateOne) SetUser(u *User) *RelationshipUpdateOne {
	return ruo.SetUserID(u.ID)
}

// SetRelative sets the "relative" edge to the User entity.
func (ruo *RelationshipUpdateOne) SetRelative(u *User) *RelationshipUpdateOne {
	return ruo.SetRelativeID(u.ID)
}

// SetInfo sets the "info" edge to the RelationshipInfo entity.
func (ruo *RelationshipUpdateOne) SetInfo(r *RelationshipInfo) *RelationshipUpdateOne {
	return ruo.SetInfoID(r.ID)
}

// Mutation returns the RelationshipMutation object of the builder.
func (ruo *RelationshipUpdateOne) Mutation() *RelationshipMutation {
	return ruo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ruo *RelationshipUpdateOne) ClearUser() *RelationshipUpdateOne {
	ruo.mutation.ClearUser()
	return ruo
}

// ClearRelative clears the "relative" edge to the User entity.
func (ruo *RelationshipUpdateOne) ClearRelative() *RelationshipUpdateOne {
	ruo.mutation.ClearRelative()
	return ruo
}

// ClearInfo clears the "info" edge to the RelationshipInfo entity.
func (ruo *RelationshipUpdateOne) ClearInfo() *RelationshipUpdateOne {
	ruo.mutation.ClearInfo()
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RelationshipUpdateOne) Select(field string, fields ...string) *RelationshipUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Relationship entity.
func (ruo *RelationshipUpdateOne) Save(ctx context.Context) (*Relationship, error) {
	return withHooks[*Relationship, RelationshipMutation](ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RelationshipUpdateOne) SaveX(ctx context.Context) *Relationship {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RelationshipUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RelationshipUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RelationshipUpdateOne) check() error {
	if _, ok := ruo.mutation.UserID(); ruo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Relationship.user"`)
	}
	if _, ok := ruo.mutation.RelativeID(); ruo.mutation.RelativeCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Relationship.relative"`)
	}
	return nil
}

func (ruo *RelationshipUpdateOne) sqlSave(ctx context.Context) (_node *Relationship, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   relationship.Table,
			Columns: relationship.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: relationship.FieldUserID,
				},
				{
					Type:   field.TypeInt,
					Column: relationship.FieldRelativeID,
				},
			},
		},
	}
	if id, ok := ruo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "Relationship.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ruo.mutation.RelativeID(); !ok {
		return nil, &ValidationError{Name: "relative_id", err: errors.New(`ent: missing "Relationship.relative_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !relationship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Weight(); ok {
		_spec.SetField(relationship.FieldWeight, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedWeight(); ok {
		_spec.AddField(relationship.FieldWeight, field.TypeInt, value)
	}
	if ruo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.UserTable,
			Columns: []string{relationship.UserColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.RelativeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RelativeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.RelativeTable,
			Columns: []string{relationship.RelativeColumn},
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.InfoCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.InfoTable,
			Columns: []string{relationship.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relationshipinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.InfoIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   relationship.InfoTable,
			Columns: []string{relationship.InfoColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: relationshipinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Relationship{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{relationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
