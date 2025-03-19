// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/migrate/entv2/customtype"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// CustomTypeUpdate is the builder for updating CustomType entities.
type CustomTypeUpdate struct {
	config
	hooks    []Hook
	mutation *CustomTypeMutation
}

// Where appends a list predicates to the CustomTypeUpdate builder.
func (_u *CustomTypeUpdate) Where(ps ...predicate.CustomType) *CustomTypeUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetCustom sets the "custom" field.
func (_u *CustomTypeUpdate) SetCustom(v string) *CustomTypeUpdate {
	_u.mutation.SetCustom(v)
	return _u
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (_u *CustomTypeUpdate) SetNillableCustom(v *string) *CustomTypeUpdate {
	if v != nil {
		_u.SetCustom(*v)
	}
	return _u
}

// ClearCustom clears the value of the "custom" field.
func (_u *CustomTypeUpdate) ClearCustom() *CustomTypeUpdate {
	_u.mutation.ClearCustom()
	return _u
}

// SetTz0 sets the "tz0" field.
func (_u *CustomTypeUpdate) SetTz0(v time.Time) *CustomTypeUpdate {
	_u.mutation.SetTz0(v)
	return _u
}

// SetNillableTz0 sets the "tz0" field if the given value is not nil.
func (_u *CustomTypeUpdate) SetNillableTz0(v *time.Time) *CustomTypeUpdate {
	if v != nil {
		_u.SetTz0(*v)
	}
	return _u
}

// ClearTz0 clears the value of the "tz0" field.
func (_u *CustomTypeUpdate) ClearTz0() *CustomTypeUpdate {
	_u.mutation.ClearTz0()
	return _u
}

// SetTz3 sets the "tz3" field.
func (_u *CustomTypeUpdate) SetTz3(v time.Time) *CustomTypeUpdate {
	_u.mutation.SetTz3(v)
	return _u
}

// SetNillableTz3 sets the "tz3" field if the given value is not nil.
func (_u *CustomTypeUpdate) SetNillableTz3(v *time.Time) *CustomTypeUpdate {
	if v != nil {
		_u.SetTz3(*v)
	}
	return _u
}

// ClearTz3 clears the value of the "tz3" field.
func (_u *CustomTypeUpdate) ClearTz3() *CustomTypeUpdate {
	_u.mutation.ClearTz3()
	return _u
}

// Mutation returns the CustomTypeMutation object of the builder.
func (_u *CustomTypeUpdate) Mutation() *CustomTypeMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *CustomTypeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CustomTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *CustomTypeUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CustomTypeUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *CustomTypeUpdate) sqlSave(ctx context.Context) (_node int, err error) {
	_spec := sqlgraph.NewUpdateSpec(customtype.Table, customtype.Columns, sqlgraph.NewFieldSpec(customtype.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
	}
	if _u.mutation.CustomCleared() {
		_spec.ClearField(customtype.FieldCustom, field.TypeString)
	}
	if value, ok := _u.mutation.Tz0(); ok {
		_spec.SetField(customtype.FieldTz0, field.TypeTime, value)
	}
	if _u.mutation.Tz0Cleared() {
		_spec.ClearField(customtype.FieldTz0, field.TypeTime)
	}
	if value, ok := _u.mutation.Tz3(); ok {
		_spec.SetField(customtype.FieldTz3, field.TypeTime, value)
	}
	if _u.mutation.Tz3Cleared() {
		_spec.ClearField(customtype.FieldTz3, field.TypeTime)
	}
	if _node, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return _node, nil
}

// CustomTypeUpdateOne is the builder for updating a single CustomType entity.
type CustomTypeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *CustomTypeMutation
}

// SetCustom sets the "custom" field.
func (_u *CustomTypeUpdateOne) SetCustom(v string) *CustomTypeUpdateOne {
	_u.mutation.SetCustom(v)
	return _u
}

// SetNillableCustom sets the "custom" field if the given value is not nil.
func (_u *CustomTypeUpdateOne) SetNillableCustom(v *string) *CustomTypeUpdateOne {
	if v != nil {
		_u.SetCustom(*v)
	}
	return _u
}

// ClearCustom clears the value of the "custom" field.
func (_u *CustomTypeUpdateOne) ClearCustom() *CustomTypeUpdateOne {
	_u.mutation.ClearCustom()
	return _u
}

// SetTz0 sets the "tz0" field.
func (_u *CustomTypeUpdateOne) SetTz0(v time.Time) *CustomTypeUpdateOne {
	_u.mutation.SetTz0(v)
	return _u
}

// SetNillableTz0 sets the "tz0" field if the given value is not nil.
func (_u *CustomTypeUpdateOne) SetNillableTz0(v *time.Time) *CustomTypeUpdateOne {
	if v != nil {
		_u.SetTz0(*v)
	}
	return _u
}

// ClearTz0 clears the value of the "tz0" field.
func (_u *CustomTypeUpdateOne) ClearTz0() *CustomTypeUpdateOne {
	_u.mutation.ClearTz0()
	return _u
}

// SetTz3 sets the "tz3" field.
func (_u *CustomTypeUpdateOne) SetTz3(v time.Time) *CustomTypeUpdateOne {
	_u.mutation.SetTz3(v)
	return _u
}

// SetNillableTz3 sets the "tz3" field if the given value is not nil.
func (_u *CustomTypeUpdateOne) SetNillableTz3(v *time.Time) *CustomTypeUpdateOne {
	if v != nil {
		_u.SetTz3(*v)
	}
	return _u
}

// ClearTz3 clears the value of the "tz3" field.
func (_u *CustomTypeUpdateOne) ClearTz3() *CustomTypeUpdateOne {
	_u.mutation.ClearTz3()
	return _u
}

// Mutation returns the CustomTypeMutation object of the builder.
func (_u *CustomTypeUpdateOne) Mutation() *CustomTypeMutation {
	return _u.mutation
}

// Where appends a list predicates to the CustomTypeUpdate builder.
func (_u *CustomTypeUpdateOne) Where(ps ...predicate.CustomType) *CustomTypeUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *CustomTypeUpdateOne) Select(field string, fields ...string) *CustomTypeUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated CustomType entity.
func (_u *CustomTypeUpdateOne) Save(ctx context.Context) (*CustomType, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *CustomTypeUpdateOne) SaveX(ctx context.Context) *CustomType {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *CustomTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *CustomTypeUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *CustomTypeUpdateOne) sqlSave(ctx context.Context) (_node *CustomType, err error) {
	_spec := sqlgraph.NewUpdateSpec(customtype.Table, customtype.Columns, sqlgraph.NewFieldSpec(customtype.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "CustomType.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, customtype.FieldID)
		for _, f := range fields {
			if !customtype.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != customtype.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Custom(); ok {
		_spec.SetField(customtype.FieldCustom, field.TypeString, value)
	}
	if _u.mutation.CustomCleared() {
		_spec.ClearField(customtype.FieldCustom, field.TypeString)
	}
	if value, ok := _u.mutation.Tz0(); ok {
		_spec.SetField(customtype.FieldTz0, field.TypeTime, value)
	}
	if _u.mutation.Tz0Cleared() {
		_spec.ClearField(customtype.FieldTz0, field.TypeTime)
	}
	if value, ok := _u.mutation.Tz3(); ok {
		_spec.SetField(customtype.FieldTz3, field.TypeTime, value)
	}
	if _u.mutation.Tz3Cleared() {
		_spec.ClearField(customtype.FieldTz3, field.TypeTime)
	}
	_node = &CustomType{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{customtype.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
