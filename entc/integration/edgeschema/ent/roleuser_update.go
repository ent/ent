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
	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"entgo.io/ent/entc/integration/edgeschema/ent/roleuser"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/schema/field"
)

// RoleUserUpdate is the builder for updating RoleUser entities.
type RoleUserUpdate struct {
	config
	hooks    []Hook
	mutation *RoleUserMutation
}

// Where appends a list predicates to the RoleUserUpdate builder.
func (ruu *RoleUserUpdate) Where(ps ...predicate.RoleUser) *RoleUserUpdate {
	ruu.mutation.Where(ps...)
	return ruu
}

// SetCreatedAt sets the "created_at" field.
func (ruu *RoleUserUpdate) SetCreatedAt(t time.Time) *RoleUserUpdate {
	ruu.mutation.SetCreatedAt(t)
	return ruu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruu *RoleUserUpdate) SetNillableCreatedAt(t *time.Time) *RoleUserUpdate {
	if t != nil {
		ruu.SetCreatedAt(*t)
	}
	return ruu
}

// SetRoleID sets the "role_id" field.
func (ruu *RoleUserUpdate) SetRoleID(i int) *RoleUserUpdate {
	ruu.mutation.SetRoleID(i)
	return ruu
}

// SetUserID sets the "user_id" field.
func (ruu *RoleUserUpdate) SetUserID(i int) *RoleUserUpdate {
	ruu.mutation.SetUserID(i)
	return ruu
}

// SetRole sets the "role" edge to the Role entity.
func (ruu *RoleUserUpdate) SetRole(r *Role) *RoleUserUpdate {
	return ruu.SetRoleID(r.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ruu *RoleUserUpdate) SetUser(u *User) *RoleUserUpdate {
	return ruu.SetUserID(u.ID)
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruu *RoleUserUpdate) Mutation() *RoleUserMutation {
	return ruu.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (ruu *RoleUserUpdate) ClearRole() *RoleUserUpdate {
	ruu.mutation.ClearRole()
	return ruu
}

// ClearUser clears the "user" edge to the User entity.
func (ruu *RoleUserUpdate) ClearUser() *RoleUserUpdate {
	ruu.mutation.ClearUser()
	return ruu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ruu *RoleUserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ruu.hooks) == 0 {
		if err = ruu.check(); err != nil {
			return 0, err
		}
		affected, err = ruu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruu.check(); err != nil {
				return 0, err
			}
			ruu.mutation = mutation
			affected, err = ruu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ruu.hooks) - 1; i >= 0; i-- {
			if ruu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ruu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruu *RoleUserUpdate) SaveX(ctx context.Context) int {
	affected, err := ruu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ruu *RoleUserUpdate) Exec(ctx context.Context) error {
	_, err := ruu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruu *RoleUserUpdate) ExecX(ctx context.Context) {
	if err := ruu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruu *RoleUserUpdate) check() error {
	if _, ok := ruu.mutation.RoleID(); ruu.mutation.RoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoleUser.role"`)
	}
	if _, ok := ruu.mutation.UserID(); ruu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoleUser.user"`)
	}
	return nil
}

func (ruu *RoleUserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleuser.Table,
			Columns: roleuser.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: roleuser.FieldUserID,
				},
				{
					Type:   field.TypeInt,
					Column: roleuser.FieldRoleID,
				},
			},
		},
	}
	if ps := ruu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: roleuser.FieldCreatedAt,
		})
	}
	if ruu.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.RoleTable,
			Columns: []string{roleuser.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruu.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.RoleTable,
			Columns: []string{roleuser.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.UserTable,
			Columns: []string{roleuser.UserColumn},
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
	if nodes := ruu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.UserTable,
			Columns: []string{roleuser.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ruu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RoleUserUpdateOne is the builder for updating a single RoleUser entity.
type RoleUserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoleUserMutation
}

// SetCreatedAt sets the "created_at" field.
func (ruuo *RoleUserUpdateOne) SetCreatedAt(t time.Time) *RoleUserUpdateOne {
	ruuo.mutation.SetCreatedAt(t)
	return ruuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ruuo *RoleUserUpdateOne) SetNillableCreatedAt(t *time.Time) *RoleUserUpdateOne {
	if t != nil {
		ruuo.SetCreatedAt(*t)
	}
	return ruuo
}

// SetRoleID sets the "role_id" field.
func (ruuo *RoleUserUpdateOne) SetRoleID(i int) *RoleUserUpdateOne {
	ruuo.mutation.SetRoleID(i)
	return ruuo
}

// SetUserID sets the "user_id" field.
func (ruuo *RoleUserUpdateOne) SetUserID(i int) *RoleUserUpdateOne {
	ruuo.mutation.SetUserID(i)
	return ruuo
}

// SetRole sets the "role" edge to the Role entity.
func (ruuo *RoleUserUpdateOne) SetRole(r *Role) *RoleUserUpdateOne {
	return ruuo.SetRoleID(r.ID)
}

// SetUser sets the "user" edge to the User entity.
func (ruuo *RoleUserUpdateOne) SetUser(u *User) *RoleUserUpdateOne {
	return ruuo.SetUserID(u.ID)
}

// Mutation returns the RoleUserMutation object of the builder.
func (ruuo *RoleUserUpdateOne) Mutation() *RoleUserMutation {
	return ruuo.mutation
}

// ClearRole clears the "role" edge to the Role entity.
func (ruuo *RoleUserUpdateOne) ClearRole() *RoleUserUpdateOne {
	ruuo.mutation.ClearRole()
	return ruuo
}

// ClearUser clears the "user" edge to the User entity.
func (ruuo *RoleUserUpdateOne) ClearUser() *RoleUserUpdateOne {
	ruuo.mutation.ClearUser()
	return ruuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruuo *RoleUserUpdateOne) Select(field string, fields ...string) *RoleUserUpdateOne {
	ruuo.fields = append([]string{field}, fields...)
	return ruuo
}

// Save executes the query and returns the updated RoleUser entity.
func (ruuo *RoleUserUpdateOne) Save(ctx context.Context) (*RoleUser, error) {
	var (
		err  error
		node *RoleUser
	)
	if len(ruuo.hooks) == 0 {
		if err = ruuo.check(); err != nil {
			return nil, err
		}
		node, err = ruuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RoleUserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ruuo.check(); err != nil {
				return nil, err
			}
			ruuo.mutation = mutation
			node, err = ruuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ruuo.hooks) - 1; i >= 0; i-- {
			if ruuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ruuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ruuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RoleUser)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RoleUserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ruuo *RoleUserUpdateOne) SaveX(ctx context.Context) *RoleUser {
	node, err := ruuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruuo *RoleUserUpdateOne) Exec(ctx context.Context) error {
	_, err := ruuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruuo *RoleUserUpdateOne) ExecX(ctx context.Context) {
	if err := ruuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruuo *RoleUserUpdateOne) check() error {
	if _, ok := ruuo.mutation.RoleID(); ruuo.mutation.RoleCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoleUser.role"`)
	}
	if _, ok := ruuo.mutation.UserID(); ruuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoleUser.user"`)
	}
	return nil
}

func (ruuo *RoleUserUpdateOne) sqlSave(ctx context.Context) (_node *RoleUser, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   roleuser.Table,
			Columns: roleuser.Columns,
			CompositeID: []*sqlgraph.FieldSpec{
				{
					Type:   field.TypeInt,
					Column: roleuser.FieldUserID,
				},
				{
					Type:   field.TypeInt,
					Column: roleuser.FieldRoleID,
				},
			},
		},
	}
	if id, ok := ruuo.mutation.UserID(); !ok {
		return nil, &ValidationError{Name: "user_id", err: errors.New(`ent: missing "RoleUser.user_id" for update`)}
	} else {
		_spec.Node.CompositeID[0].Value = id
	}
	if id, ok := ruuo.mutation.RoleID(); !ok {
		return nil, &ValidationError{Name: "role_id", err: errors.New(`ent: missing "RoleUser.role_id" for update`)}
	} else {
		_spec.Node.CompositeID[1].Value = id
	}
	if fields := ruuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, len(fields))
		for i, f := range fields {
			if !roleuser.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			_spec.Node.Columns[i] = f
		}
	}
	if ps := ruuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: roleuser.FieldCreatedAt,
		})
	}
	if ruuo.mutation.RoleCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.RoleTable,
			Columns: []string{roleuser.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruuo.mutation.RoleIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.RoleTable,
			Columns: []string{roleuser.RoleColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: role.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.UserTable,
			Columns: []string{roleuser.UserColumn},
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
	if nodes := ruuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   roleuser.UserTable,
			Columns: []string{roleuser.UserColumn},
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
	_node = &RoleUser{config: ruuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roleuser.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
