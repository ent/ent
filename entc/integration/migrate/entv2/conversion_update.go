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
	"entgo.io/ent/entc/integration/migrate/entv2/conversion"
	"entgo.io/ent/entc/integration/migrate/entv2/predicate"
	"entgo.io/ent/schema/field"
)

// ConversionUpdate is the builder for updating Conversion entities.
type ConversionUpdate struct {
	config
	hooks    []Hook
	mutation *ConversionMutation
}

// Where appends a list predicates to the ConversionUpdate builder.
func (_u *ConversionUpdate) Where(ps ...predicate.Conversion) *ConversionUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// SetName sets the "name" field.
func (_u *ConversionUpdate) SetName(s string) *ConversionUpdate {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableName(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// ClearName clears the value of the "name" field.
func (_u *ConversionUpdate) ClearName() *ConversionUpdate {
	_u.mutation.ClearName()
	return _u
}

// SetInt8ToString sets the "int8_to_string" field.
func (_u *ConversionUpdate) SetInt8ToString(s string) *ConversionUpdate {
	_u.mutation.SetInt8ToString(s)
	return _u
}

// SetNillableInt8ToString sets the "int8_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableInt8ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetInt8ToString(*s)
	}
	return _u
}

// ClearInt8ToString clears the value of the "int8_to_string" field.
func (_u *ConversionUpdate) ClearInt8ToString() *ConversionUpdate {
	_u.mutation.ClearInt8ToString()
	return _u
}

// SetUint8ToString sets the "uint8_to_string" field.
func (_u *ConversionUpdate) SetUint8ToString(s string) *ConversionUpdate {
	_u.mutation.SetUint8ToString(s)
	return _u
}

// SetNillableUint8ToString sets the "uint8_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableUint8ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetUint8ToString(*s)
	}
	return _u
}

// ClearUint8ToString clears the value of the "uint8_to_string" field.
func (_u *ConversionUpdate) ClearUint8ToString() *ConversionUpdate {
	_u.mutation.ClearUint8ToString()
	return _u
}

// SetInt16ToString sets the "int16_to_string" field.
func (_u *ConversionUpdate) SetInt16ToString(s string) *ConversionUpdate {
	_u.mutation.SetInt16ToString(s)
	return _u
}

// SetNillableInt16ToString sets the "int16_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableInt16ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetInt16ToString(*s)
	}
	return _u
}

// ClearInt16ToString clears the value of the "int16_to_string" field.
func (_u *ConversionUpdate) ClearInt16ToString() *ConversionUpdate {
	_u.mutation.ClearInt16ToString()
	return _u
}

// SetUint16ToString sets the "uint16_to_string" field.
func (_u *ConversionUpdate) SetUint16ToString(s string) *ConversionUpdate {
	_u.mutation.SetUint16ToString(s)
	return _u
}

// SetNillableUint16ToString sets the "uint16_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableUint16ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetUint16ToString(*s)
	}
	return _u
}

// ClearUint16ToString clears the value of the "uint16_to_string" field.
func (_u *ConversionUpdate) ClearUint16ToString() *ConversionUpdate {
	_u.mutation.ClearUint16ToString()
	return _u
}

// SetInt32ToString sets the "int32_to_string" field.
func (_u *ConversionUpdate) SetInt32ToString(s string) *ConversionUpdate {
	_u.mutation.SetInt32ToString(s)
	return _u
}

// SetNillableInt32ToString sets the "int32_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableInt32ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetInt32ToString(*s)
	}
	return _u
}

// ClearInt32ToString clears the value of the "int32_to_string" field.
func (_u *ConversionUpdate) ClearInt32ToString() *ConversionUpdate {
	_u.mutation.ClearInt32ToString()
	return _u
}

// SetUint32ToString sets the "uint32_to_string" field.
func (_u *ConversionUpdate) SetUint32ToString(s string) *ConversionUpdate {
	_u.mutation.SetUint32ToString(s)
	return _u
}

// SetNillableUint32ToString sets the "uint32_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableUint32ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetUint32ToString(*s)
	}
	return _u
}

// ClearUint32ToString clears the value of the "uint32_to_string" field.
func (_u *ConversionUpdate) ClearUint32ToString() *ConversionUpdate {
	_u.mutation.ClearUint32ToString()
	return _u
}

// SetInt64ToString sets the "int64_to_string" field.
func (_u *ConversionUpdate) SetInt64ToString(s string) *ConversionUpdate {
	_u.mutation.SetInt64ToString(s)
	return _u
}

// SetNillableInt64ToString sets the "int64_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableInt64ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetInt64ToString(*s)
	}
	return _u
}

// ClearInt64ToString clears the value of the "int64_to_string" field.
func (_u *ConversionUpdate) ClearInt64ToString() *ConversionUpdate {
	_u.mutation.ClearInt64ToString()
	return _u
}

// SetUint64ToString sets the "uint64_to_string" field.
func (_u *ConversionUpdate) SetUint64ToString(s string) *ConversionUpdate {
	_u.mutation.SetUint64ToString(s)
	return _u
}

// SetNillableUint64ToString sets the "uint64_to_string" field if the given value is not nil.
func (_u *ConversionUpdate) SetNillableUint64ToString(s *string) *ConversionUpdate {
	if s != nil {
		_u.SetUint64ToString(*s)
	}
	return _u
}

// ClearUint64ToString clears the value of the "uint64_to_string" field.
func (_u *ConversionUpdate) ClearUint64ToString() *ConversionUpdate {
	_u.mutation.ClearUint64ToString()
	return _u
}

// Mutation returns the ConversionMutation object of the builder.
func (_u *ConversionUpdate) Mutation() *ConversionMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *ConversionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *ConversionUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *ConversionUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *ConversionUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *ConversionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(conversion.Table, conversion.Columns, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	if ps := _u.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(conversion.FieldName, field.TypeString, value)
	}
	if _u.mutation.NameCleared() {
		_spec.ClearField(conversion.FieldName, field.TypeString)
	}
	if value, ok := _u.mutation.Int8ToString(); ok {
		_spec.SetField(conversion.FieldInt8ToString, field.TypeString, value)
	}
	if _u.mutation.Int8ToStringCleared() {
		_spec.ClearField(conversion.FieldInt8ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint8ToString(); ok {
		_spec.SetField(conversion.FieldUint8ToString, field.TypeString, value)
	}
	if _u.mutation.Uint8ToStringCleared() {
		_spec.ClearField(conversion.FieldUint8ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int16ToString(); ok {
		_spec.SetField(conversion.FieldInt16ToString, field.TypeString, value)
	}
	if _u.mutation.Int16ToStringCleared() {
		_spec.ClearField(conversion.FieldInt16ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint16ToString(); ok {
		_spec.SetField(conversion.FieldUint16ToString, field.TypeString, value)
	}
	if _u.mutation.Uint16ToStringCleared() {
		_spec.ClearField(conversion.FieldUint16ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int32ToString(); ok {
		_spec.SetField(conversion.FieldInt32ToString, field.TypeString, value)
	}
	if _u.mutation.Int32ToStringCleared() {
		_spec.ClearField(conversion.FieldInt32ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint32ToString(); ok {
		_spec.SetField(conversion.FieldUint32ToString, field.TypeString, value)
	}
	if _u.mutation.Uint32ToStringCleared() {
		_spec.ClearField(conversion.FieldUint32ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int64ToString(); ok {
		_spec.SetField(conversion.FieldInt64ToString, field.TypeString, value)
	}
	if _u.mutation.Int64ToStringCleared() {
		_spec.ClearField(conversion.FieldInt64ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint64ToString(); ok {
		_spec.SetField(conversion.FieldUint64ToString, field.TypeString, value)
	}
	if _u.mutation.Uint64ToStringCleared() {
		_spec.ClearField(conversion.FieldUint64ToString, field.TypeString)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	_u.mutation.done = true
	return n, nil
}

// ConversionUpdateOne is the builder for updating a single Conversion entity.
type ConversionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ConversionMutation
}

// SetName sets the "name" field.
func (_u *ConversionUpdateOne) SetName(s string) *ConversionUpdateOne {
	_u.mutation.SetName(s)
	return _u
}

// SetNillableName sets the "name" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableName(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetName(*s)
	}
	return _u
}

// ClearName clears the value of the "name" field.
func (_u *ConversionUpdateOne) ClearName() *ConversionUpdateOne {
	_u.mutation.ClearName()
	return _u
}

// SetInt8ToString sets the "int8_to_string" field.
func (_u *ConversionUpdateOne) SetInt8ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetInt8ToString(s)
	return _u
}

// SetNillableInt8ToString sets the "int8_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableInt8ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetInt8ToString(*s)
	}
	return _u
}

// ClearInt8ToString clears the value of the "int8_to_string" field.
func (_u *ConversionUpdateOne) ClearInt8ToString() *ConversionUpdateOne {
	_u.mutation.ClearInt8ToString()
	return _u
}

// SetUint8ToString sets the "uint8_to_string" field.
func (_u *ConversionUpdateOne) SetUint8ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetUint8ToString(s)
	return _u
}

// SetNillableUint8ToString sets the "uint8_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableUint8ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetUint8ToString(*s)
	}
	return _u
}

// ClearUint8ToString clears the value of the "uint8_to_string" field.
func (_u *ConversionUpdateOne) ClearUint8ToString() *ConversionUpdateOne {
	_u.mutation.ClearUint8ToString()
	return _u
}

// SetInt16ToString sets the "int16_to_string" field.
func (_u *ConversionUpdateOne) SetInt16ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetInt16ToString(s)
	return _u
}

// SetNillableInt16ToString sets the "int16_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableInt16ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetInt16ToString(*s)
	}
	return _u
}

// ClearInt16ToString clears the value of the "int16_to_string" field.
func (_u *ConversionUpdateOne) ClearInt16ToString() *ConversionUpdateOne {
	_u.mutation.ClearInt16ToString()
	return _u
}

// SetUint16ToString sets the "uint16_to_string" field.
func (_u *ConversionUpdateOne) SetUint16ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetUint16ToString(s)
	return _u
}

// SetNillableUint16ToString sets the "uint16_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableUint16ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetUint16ToString(*s)
	}
	return _u
}

// ClearUint16ToString clears the value of the "uint16_to_string" field.
func (_u *ConversionUpdateOne) ClearUint16ToString() *ConversionUpdateOne {
	_u.mutation.ClearUint16ToString()
	return _u
}

// SetInt32ToString sets the "int32_to_string" field.
func (_u *ConversionUpdateOne) SetInt32ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetInt32ToString(s)
	return _u
}

// SetNillableInt32ToString sets the "int32_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableInt32ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetInt32ToString(*s)
	}
	return _u
}

// ClearInt32ToString clears the value of the "int32_to_string" field.
func (_u *ConversionUpdateOne) ClearInt32ToString() *ConversionUpdateOne {
	_u.mutation.ClearInt32ToString()
	return _u
}

// SetUint32ToString sets the "uint32_to_string" field.
func (_u *ConversionUpdateOne) SetUint32ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetUint32ToString(s)
	return _u
}

// SetNillableUint32ToString sets the "uint32_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableUint32ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetUint32ToString(*s)
	}
	return _u
}

// ClearUint32ToString clears the value of the "uint32_to_string" field.
func (_u *ConversionUpdateOne) ClearUint32ToString() *ConversionUpdateOne {
	_u.mutation.ClearUint32ToString()
	return _u
}

// SetInt64ToString sets the "int64_to_string" field.
func (_u *ConversionUpdateOne) SetInt64ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetInt64ToString(s)
	return _u
}

// SetNillableInt64ToString sets the "int64_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableInt64ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetInt64ToString(*s)
	}
	return _u
}

// ClearInt64ToString clears the value of the "int64_to_string" field.
func (_u *ConversionUpdateOne) ClearInt64ToString() *ConversionUpdateOne {
	_u.mutation.ClearInt64ToString()
	return _u
}

// SetUint64ToString sets the "uint64_to_string" field.
func (_u *ConversionUpdateOne) SetUint64ToString(s string) *ConversionUpdateOne {
	_u.mutation.SetUint64ToString(s)
	return _u
}

// SetNillableUint64ToString sets the "uint64_to_string" field if the given value is not nil.
func (_u *ConversionUpdateOne) SetNillableUint64ToString(s *string) *ConversionUpdateOne {
	if s != nil {
		_u.SetUint64ToString(*s)
	}
	return _u
}

// ClearUint64ToString clears the value of the "uint64_to_string" field.
func (_u *ConversionUpdateOne) ClearUint64ToString() *ConversionUpdateOne {
	_u.mutation.ClearUint64ToString()
	return _u
}

// Mutation returns the ConversionMutation object of the builder.
func (_u *ConversionUpdateOne) Mutation() *ConversionMutation {
	return _u.mutation
}

// Where appends a list predicates to the ConversionUpdate builder.
func (_u *ConversionUpdateOne) Where(ps ...predicate.Conversion) *ConversionUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *ConversionUpdateOne) Select(field string, fields ...string) *ConversionUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Conversion entity.
func (_u *ConversionUpdateOne) Save(ctx context.Context) (*Conversion, error) {
	return withHooks(ctx, _u.sqlSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *ConversionUpdateOne) SaveX(ctx context.Context) *Conversion {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *ConversionUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *ConversionUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *ConversionUpdateOne) sqlSave(ctx context.Context) (_node *Conversion, err error) {
	_spec := sqlgraph.NewUpdateSpec(conversion.Table, conversion.Columns, sqlgraph.NewFieldSpec(conversion.FieldID, field.TypeInt))
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`entv2: missing "Conversion.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := _u.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, conversion.FieldID)
		for _, f := range fields {
			if !conversion.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("entv2: invalid field %q for query", f)}
			}
			if f != conversion.FieldID {
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
	if value, ok := _u.mutation.Name(); ok {
		_spec.SetField(conversion.FieldName, field.TypeString, value)
	}
	if _u.mutation.NameCleared() {
		_spec.ClearField(conversion.FieldName, field.TypeString)
	}
	if value, ok := _u.mutation.Int8ToString(); ok {
		_spec.SetField(conversion.FieldInt8ToString, field.TypeString, value)
	}
	if _u.mutation.Int8ToStringCleared() {
		_spec.ClearField(conversion.FieldInt8ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint8ToString(); ok {
		_spec.SetField(conversion.FieldUint8ToString, field.TypeString, value)
	}
	if _u.mutation.Uint8ToStringCleared() {
		_spec.ClearField(conversion.FieldUint8ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int16ToString(); ok {
		_spec.SetField(conversion.FieldInt16ToString, field.TypeString, value)
	}
	if _u.mutation.Int16ToStringCleared() {
		_spec.ClearField(conversion.FieldInt16ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint16ToString(); ok {
		_spec.SetField(conversion.FieldUint16ToString, field.TypeString, value)
	}
	if _u.mutation.Uint16ToStringCleared() {
		_spec.ClearField(conversion.FieldUint16ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int32ToString(); ok {
		_spec.SetField(conversion.FieldInt32ToString, field.TypeString, value)
	}
	if _u.mutation.Int32ToStringCleared() {
		_spec.ClearField(conversion.FieldInt32ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint32ToString(); ok {
		_spec.SetField(conversion.FieldUint32ToString, field.TypeString, value)
	}
	if _u.mutation.Uint32ToStringCleared() {
		_spec.ClearField(conversion.FieldUint32ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Int64ToString(); ok {
		_spec.SetField(conversion.FieldInt64ToString, field.TypeString, value)
	}
	if _u.mutation.Int64ToStringCleared() {
		_spec.ClearField(conversion.FieldInt64ToString, field.TypeString)
	}
	if value, ok := _u.mutation.Uint64ToString(); ok {
		_spec.SetField(conversion.FieldUint64ToString, field.TypeString, value)
	}
	if _u.mutation.Uint64ToStringCleared() {
		_spec.ClearField(conversion.FieldUint64ToString, field.TypeString)
	}
	_node = &Conversion{config: _u.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, _u.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{conversion.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	_u.mutation.done = true
	return _node, nil
}
