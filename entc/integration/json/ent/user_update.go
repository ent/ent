// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/entc/integration/json/ent/predicate"
	"entgo.io/ent/entc/integration/json/ent/schema"
	"entgo.io/ent/entc/integration/json/ent/user"
	"entgo.io/ent/schema/field"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetT sets the "t" field.
func (uu *UserUpdate) SetT(s *schema.T) *UserUpdate {
	uu.mutation.SetT(s)
	return uu
}

// ClearT clears the value of the "t" field.
func (uu *UserUpdate) ClearT() *UserUpdate {
	uu.mutation.ClearT()
	return uu
}

// SetURL sets the "url" field.
func (uu *UserUpdate) SetURL(u *url.URL) *UserUpdate {
	uu.mutation.SetURL(u)
	return uu
}

// ClearURL clears the value of the "url" field.
func (uu *UserUpdate) ClearURL() *UserUpdate {
	uu.mutation.ClearURL()
	return uu
}

// SetURLs sets the "URLs" field.
func (uu *UserUpdate) SetURLs(u []*url.URL) *UserUpdate {
	uu.mutation.SetURLs(u)
	return uu
}

// AppendURLs appends u to the "URLs" field.
func (uu *UserUpdate) AppendURLs(u []*url.URL) *UserUpdate {
	uu.mutation.AppendURLs(u)
	return uu
}

// ClearURLs clears the value of the "URLs" field.
func (uu *UserUpdate) ClearURLs() *UserUpdate {
	uu.mutation.ClearURLs()
	return uu
}

// SetRaw sets the "raw" field.
func (uu *UserUpdate) SetRaw(jm json.RawMessage) *UserUpdate {
	uu.mutation.SetRaw(jm)
	return uu
}

// AppendRaw appends jm to the "raw" field.
func (uu *UserUpdate) AppendRaw(jm json.RawMessage) *UserUpdate {
	uu.mutation.AppendRaw(jm)
	return uu
}

// ClearRaw clears the value of the "raw" field.
func (uu *UserUpdate) ClearRaw() *UserUpdate {
	uu.mutation.ClearRaw()
	return uu
}

// SetDirs sets the "dirs" field.
func (uu *UserUpdate) SetDirs(h []http.Dir) *UserUpdate {
	uu.mutation.SetDirs(h)
	return uu
}

// AppendDirs appends h to the "dirs" field.
func (uu *UserUpdate) AppendDirs(h []http.Dir) *UserUpdate {
	uu.mutation.AppendDirs(h)
	return uu
}

// SetInts sets the "ints" field.
func (uu *UserUpdate) SetInts(i []int) *UserUpdate {
	uu.mutation.SetInts(i)
	return uu
}

// AppendInts appends i to the "ints" field.
func (uu *UserUpdate) AppendInts(i []int) *UserUpdate {
	uu.mutation.AppendInts(i)
	return uu
}

// ClearInts clears the value of the "ints" field.
func (uu *UserUpdate) ClearInts() *UserUpdate {
	uu.mutation.ClearInts()
	return uu
}

// SetFloats sets the "floats" field.
func (uu *UserUpdate) SetFloats(f []float64) *UserUpdate {
	uu.mutation.SetFloats(f)
	return uu
}

// AppendFloats appends f to the "floats" field.
func (uu *UserUpdate) AppendFloats(f []float64) *UserUpdate {
	uu.mutation.AppendFloats(f)
	return uu
}

// ClearFloats clears the value of the "floats" field.
func (uu *UserUpdate) ClearFloats() *UserUpdate {
	uu.mutation.ClearFloats()
	return uu
}

// SetStrings sets the "strings" field.
func (uu *UserUpdate) SetStrings(s []string) *UserUpdate {
	uu.mutation.SetStrings(s)
	return uu
}

// AppendStrings appends s to the "strings" field.
func (uu *UserUpdate) AppendStrings(s []string) *UserUpdate {
	uu.mutation.AppendStrings(s)
	return uu
}

// ClearStrings clears the value of the "strings" field.
func (uu *UserUpdate) ClearStrings() *UserUpdate {
	uu.mutation.ClearStrings()
	return uu
}

// SetAddr sets the "addr" field.
func (uu *UserUpdate) SetAddr(s schema.Addr) *UserUpdate {
	uu.mutation.SetAddr(s)
	return uu
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddr(s *schema.Addr) *UserUpdate {
	if s != nil {
		uu.SetAddr(*s)
	}
	return uu
}

// ClearAddr clears the value of the "addr" field.
func (uu *UserUpdate) ClearAddr() *UserUpdate {
	uu.mutation.ClearAddr()
	return uu
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(uu.hooks) == 0 {
		affected, err = uu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uu.mutation = mutation
			affected, err = uu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uu.hooks) - 1; i >= 0; i-- {
			if uu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uu *UserUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdate {
	uu.modifiers = append(uu.modifiers, modifiers...)
	return uu
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
	}
	if uu.mutation.TCleared() {
		_spec.ClearField(user.FieldT, field.TypeJSON)
	}
	if value, ok := uu.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uu.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uu.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedURLs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldURLs, value)
		})
	}
	if uu.mutation.URLsCleared() {
		_spec.ClearField(user.FieldURLs, field.TypeJSON)
	}
	if value, ok := uu.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedRaw(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldRaw, value)
		})
	}
	if uu.mutation.RawCleared() {
		_spec.ClearField(user.FieldRaw, field.TypeJSON)
	}
	if value, ok := uu.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedDirs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldDirs, value)
		})
	}
	if value, ok := uu.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldInts, value)
		})
	}
	if uu.mutation.IntsCleared() {
		_spec.ClearField(user.FieldInts, field.TypeJSON)
	}
	if value, ok := uu.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloats, value)
		})
	}
	if uu.mutation.FloatsCleared() {
		_spec.ClearField(user.FieldFloats, field.TypeJSON)
	}
	if value, ok := uu.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uu.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uu.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uu.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
	}
	if uu.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeJSON)
	}
	_spec.AddModifiers(uu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *UserMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetT sets the "t" field.
func (uuo *UserUpdateOne) SetT(s *schema.T) *UserUpdateOne {
	uuo.mutation.SetT(s)
	return uuo
}

// ClearT clears the value of the "t" field.
func (uuo *UserUpdateOne) ClearT() *UserUpdateOne {
	uuo.mutation.ClearT()
	return uuo
}

// SetURL sets the "url" field.
func (uuo *UserUpdateOne) SetURL(u *url.URL) *UserUpdateOne {
	uuo.mutation.SetURL(u)
	return uuo
}

// ClearURL clears the value of the "url" field.
func (uuo *UserUpdateOne) ClearURL() *UserUpdateOne {
	uuo.mutation.ClearURL()
	return uuo
}

// SetURLs sets the "URLs" field.
func (uuo *UserUpdateOne) SetURLs(u []*url.URL) *UserUpdateOne {
	uuo.mutation.SetURLs(u)
	return uuo
}

// AppendURLs appends u to the "URLs" field.
func (uuo *UserUpdateOne) AppendURLs(u []*url.URL) *UserUpdateOne {
	uuo.mutation.AppendURLs(u)
	return uuo
}

// ClearURLs clears the value of the "URLs" field.
func (uuo *UserUpdateOne) ClearURLs() *UserUpdateOne {
	uuo.mutation.ClearURLs()
	return uuo
}

// SetRaw sets the "raw" field.
func (uuo *UserUpdateOne) SetRaw(jm json.RawMessage) *UserUpdateOne {
	uuo.mutation.SetRaw(jm)
	return uuo
}

// AppendRaw appends jm to the "raw" field.
func (uuo *UserUpdateOne) AppendRaw(jm json.RawMessage) *UserUpdateOne {
	uuo.mutation.AppendRaw(jm)
	return uuo
}

// ClearRaw clears the value of the "raw" field.
func (uuo *UserUpdateOne) ClearRaw() *UserUpdateOne {
	uuo.mutation.ClearRaw()
	return uuo
}

// SetDirs sets the "dirs" field.
func (uuo *UserUpdateOne) SetDirs(h []http.Dir) *UserUpdateOne {
	uuo.mutation.SetDirs(h)
	return uuo
}

// AppendDirs appends h to the "dirs" field.
func (uuo *UserUpdateOne) AppendDirs(h []http.Dir) *UserUpdateOne {
	uuo.mutation.AppendDirs(h)
	return uuo
}

// SetInts sets the "ints" field.
func (uuo *UserUpdateOne) SetInts(i []int) *UserUpdateOne {
	uuo.mutation.SetInts(i)
	return uuo
}

// AppendInts appends i to the "ints" field.
func (uuo *UserUpdateOne) AppendInts(i []int) *UserUpdateOne {
	uuo.mutation.AppendInts(i)
	return uuo
}

// ClearInts clears the value of the "ints" field.
func (uuo *UserUpdateOne) ClearInts() *UserUpdateOne {
	uuo.mutation.ClearInts()
	return uuo
}

// SetFloats sets the "floats" field.
func (uuo *UserUpdateOne) SetFloats(f []float64) *UserUpdateOne {
	uuo.mutation.SetFloats(f)
	return uuo
}

// AppendFloats appends f to the "floats" field.
func (uuo *UserUpdateOne) AppendFloats(f []float64) *UserUpdateOne {
	uuo.mutation.AppendFloats(f)
	return uuo
}

// ClearFloats clears the value of the "floats" field.
func (uuo *UserUpdateOne) ClearFloats() *UserUpdateOne {
	uuo.mutation.ClearFloats()
	return uuo
}

// SetStrings sets the "strings" field.
func (uuo *UserUpdateOne) SetStrings(s []string) *UserUpdateOne {
	uuo.mutation.SetStrings(s)
	return uuo
}

// AppendStrings appends s to the "strings" field.
func (uuo *UserUpdateOne) AppendStrings(s []string) *UserUpdateOne {
	uuo.mutation.AppendStrings(s)
	return uuo
}

// ClearStrings clears the value of the "strings" field.
func (uuo *UserUpdateOne) ClearStrings() *UserUpdateOne {
	uuo.mutation.ClearStrings()
	return uuo
}

// SetAddr sets the "addr" field.
func (uuo *UserUpdateOne) SetAddr(s schema.Addr) *UserUpdateOne {
	uuo.mutation.SetAddr(s)
	return uuo
}

// SetNillableAddr sets the "addr" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddr(s *schema.Addr) *UserUpdateOne {
	if s != nil {
		uuo.SetAddr(*s)
	}
	return uuo
}

// ClearAddr clears the value of the "addr" field.
func (uuo *UserUpdateOne) ClearAddr() *UserUpdateOne {
	uuo.mutation.ClearAddr()
	return uuo
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	var (
		err  error
		node *User
	)
	if len(uuo.hooks) == 0 {
		node, err = uuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uuo.mutation = mutation
			node, err = uuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uuo.hooks) - 1; i >= 0; i-- {
			if uuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = uuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, uuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*User)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from UserMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (uuo *UserUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *UserUpdateOne {
	uuo.modifiers = append(uuo.modifiers, modifiers...)
	return uuo
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
	}
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.T(); ok {
		_spec.SetField(user.FieldT, field.TypeJSON, value)
	}
	if uuo.mutation.TCleared() {
		_spec.ClearField(user.FieldT, field.TypeJSON)
	}
	if value, ok := uuo.mutation.URL(); ok {
		_spec.SetField(user.FieldURL, field.TypeJSON, value)
	}
	if uuo.mutation.URLCleared() {
		_spec.ClearField(user.FieldURL, field.TypeJSON)
	}
	if value, ok := uuo.mutation.URLs(); ok {
		_spec.SetField(user.FieldURLs, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedURLs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldURLs, value)
		})
	}
	if uuo.mutation.URLsCleared() {
		_spec.ClearField(user.FieldURLs, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Raw(); ok {
		_spec.SetField(user.FieldRaw, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedRaw(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldRaw, value)
		})
	}
	if uuo.mutation.RawCleared() {
		_spec.ClearField(user.FieldRaw, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Dirs(); ok {
		_spec.SetField(user.FieldDirs, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedDirs(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldDirs, value)
		})
	}
	if value, ok := uuo.mutation.Ints(); ok {
		_spec.SetField(user.FieldInts, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedInts(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldInts, value)
		})
	}
	if uuo.mutation.IntsCleared() {
		_spec.ClearField(user.FieldInts, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Floats(); ok {
		_spec.SetField(user.FieldFloats, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedFloats(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldFloats, value)
		})
	}
	if uuo.mutation.FloatsCleared() {
		_spec.ClearField(user.FieldFloats, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Strings(); ok {
		_spec.SetField(user.FieldStrings, field.TypeJSON, value)
	}
	if value, ok := uuo.mutation.AppendedStrings(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, user.FieldStrings, value)
		})
	}
	if uuo.mutation.StringsCleared() {
		_spec.ClearField(user.FieldStrings, field.TypeJSON)
	}
	if value, ok := uuo.mutation.Addr(); ok {
		_spec.SetField(user.FieldAddr, field.TypeJSON, value)
	}
	if uuo.mutation.AddrCleared() {
		_spec.ClearField(user.FieldAddr, field.TypeJSON)
	}
	_spec.AddModifiers(uuo.modifiers...)
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
