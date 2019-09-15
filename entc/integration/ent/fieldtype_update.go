// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// FieldTypeUpdate is the builder for updating FieldType entities.
type FieldTypeUpdate struct {
	config
	int                          *int
	addint                       *int
	int8                         *int8
	int16                        *int16
	int32                        *int32
	int64                        *int64
	addint64                     *int64
	optional_int                 *int
	addoptional_int              *int
	clearoptional_int            bool
	optional_int8                *int8
	clearoptional_int8           bool
	optional_int16               *int16
	clearoptional_int16          bool
	optional_int32               *int32
	clearoptional_int32          bool
	optional_int64               *int64
	addoptional_int64            *int64
	clearoptional_int64          bool
	nillable_int                 *int
	addnillable_int              *int
	clearnillable_int            bool
	nillable_int8                *int8
	clearnillable_int8           bool
	nillable_int16               *int16
	clearnillable_int16          bool
	nillable_int32               *int32
	clearnillable_int32          bool
	nillable_int64               *int64
	addnillable_int64            *int64
	clearnillable_int64          bool
	validate_optional_int32      *int32
	clearvalidate_optional_int32 bool
	predicates                   []predicate.FieldType
}

// Where adds a new predicate for the builder.
func (ftu *FieldTypeUpdate) Where(ps ...predicate.FieldType) *FieldTypeUpdate {
	ftu.predicates = append(ftu.predicates, ps...)
	return ftu
}

// SetInt sets the int field.
func (ftu *FieldTypeUpdate) SetInt(i int) *FieldTypeUpdate {
	ftu.int = &i
	return ftu
}

// AddInt adds i to int.
func (ftu *FieldTypeUpdate) AddInt(i int) *FieldTypeUpdate {
	ftu.addint = &i
	return ftu
}

// SetInt8 sets the int8 field.
func (ftu *FieldTypeUpdate) SetInt8(i int8) *FieldTypeUpdate {
	ftu.int8 = &i
	return ftu
}

// SetInt16 sets the int16 field.
func (ftu *FieldTypeUpdate) SetInt16(i int16) *FieldTypeUpdate {
	ftu.int16 = &i
	return ftu
}

// SetInt32 sets the int32 field.
func (ftu *FieldTypeUpdate) SetInt32(i int32) *FieldTypeUpdate {
	ftu.int32 = &i
	return ftu
}

// SetInt64 sets the int64 field.
func (ftu *FieldTypeUpdate) SetInt64(i int64) *FieldTypeUpdate {
	ftu.int64 = &i
	return ftu
}

// AddInt64 adds i to int64.
func (ftu *FieldTypeUpdate) AddInt64(i int64) *FieldTypeUpdate {
	ftu.addint64 = &i
	return ftu
}

// SetOptionalInt sets the optional_int field.
func (ftu *FieldTypeUpdate) SetOptionalInt(i int) *FieldTypeUpdate {
	ftu.optional_int = &i
	return ftu
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt(i *int) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt(*i)
	}
	return ftu
}

// AddOptionalInt adds i to optional_int.
func (ftu *FieldTypeUpdate) AddOptionalInt(i int) *FieldTypeUpdate {
	ftu.addoptional_int = &i
	return ftu
}

// ClearOptionalInt clears the value of optional_int.
func (ftu *FieldTypeUpdate) ClearOptionalInt() *FieldTypeUpdate {
	ftu.optional_int = nil
	ftu.clearoptional_int = true
	return ftu
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftu *FieldTypeUpdate) SetOptionalInt8(i int8) *FieldTypeUpdate {
	ftu.optional_int8 = &i
	return ftu
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt8(i *int8) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt8(*i)
	}
	return ftu
}

// ClearOptionalInt8 clears the value of optional_int8.
func (ftu *FieldTypeUpdate) ClearOptionalInt8() *FieldTypeUpdate {
	ftu.optional_int8 = nil
	ftu.clearoptional_int8 = true
	return ftu
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftu *FieldTypeUpdate) SetOptionalInt16(i int16) *FieldTypeUpdate {
	ftu.optional_int16 = &i
	return ftu
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt16(i *int16) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt16(*i)
	}
	return ftu
}

// ClearOptionalInt16 clears the value of optional_int16.
func (ftu *FieldTypeUpdate) ClearOptionalInt16() *FieldTypeUpdate {
	ftu.optional_int16 = nil
	ftu.clearoptional_int16 = true
	return ftu
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftu *FieldTypeUpdate) SetOptionalInt32(i int32) *FieldTypeUpdate {
	ftu.optional_int32 = &i
	return ftu
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt32(i *int32) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt32(*i)
	}
	return ftu
}

// ClearOptionalInt32 clears the value of optional_int32.
func (ftu *FieldTypeUpdate) ClearOptionalInt32() *FieldTypeUpdate {
	ftu.optional_int32 = nil
	ftu.clearoptional_int32 = true
	return ftu
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftu *FieldTypeUpdate) SetOptionalInt64(i int64) *FieldTypeUpdate {
	ftu.optional_int64 = &i
	return ftu
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableOptionalInt64(i *int64) *FieldTypeUpdate {
	if i != nil {
		ftu.SetOptionalInt64(*i)
	}
	return ftu
}

// AddOptionalInt64 adds i to optional_int64.
func (ftu *FieldTypeUpdate) AddOptionalInt64(i int64) *FieldTypeUpdate {
	ftu.addoptional_int64 = &i
	return ftu
}

// ClearOptionalInt64 clears the value of optional_int64.
func (ftu *FieldTypeUpdate) ClearOptionalInt64() *FieldTypeUpdate {
	ftu.optional_int64 = nil
	ftu.clearoptional_int64 = true
	return ftu
}

// SetNillableInt sets the nillable_int field.
func (ftu *FieldTypeUpdate) SetNillableInt(i int) *FieldTypeUpdate {
	ftu.nillable_int = &i
	return ftu
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNillableInt(i *int) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNillableInt(*i)
	}
	return ftu
}

// AddNillableInt adds i to nillable_int.
func (ftu *FieldTypeUpdate) AddNillableInt(i int) *FieldTypeUpdate {
	ftu.addnillable_int = &i
	return ftu
}

// ClearNillableInt clears the value of nillable_int.
func (ftu *FieldTypeUpdate) ClearNillableInt() *FieldTypeUpdate {
	ftu.nillable_int = nil
	ftu.clearnillable_int = true
	return ftu
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftu *FieldTypeUpdate) SetNillableInt8(i int8) *FieldTypeUpdate {
	ftu.nillable_int8 = &i
	return ftu
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNillableInt8(i *int8) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNillableInt8(*i)
	}
	return ftu
}

// ClearNillableInt8 clears the value of nillable_int8.
func (ftu *FieldTypeUpdate) ClearNillableInt8() *FieldTypeUpdate {
	ftu.nillable_int8 = nil
	ftu.clearnillable_int8 = true
	return ftu
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftu *FieldTypeUpdate) SetNillableInt16(i int16) *FieldTypeUpdate {
	ftu.nillable_int16 = &i
	return ftu
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNillableInt16(i *int16) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNillableInt16(*i)
	}
	return ftu
}

// ClearNillableInt16 clears the value of nillable_int16.
func (ftu *FieldTypeUpdate) ClearNillableInt16() *FieldTypeUpdate {
	ftu.nillable_int16 = nil
	ftu.clearnillable_int16 = true
	return ftu
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftu *FieldTypeUpdate) SetNillableInt32(i int32) *FieldTypeUpdate {
	ftu.nillable_int32 = &i
	return ftu
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNillableInt32(i *int32) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNillableInt32(*i)
	}
	return ftu
}

// ClearNillableInt32 clears the value of nillable_int32.
func (ftu *FieldTypeUpdate) ClearNillableInt32() *FieldTypeUpdate {
	ftu.nillable_int32 = nil
	ftu.clearnillable_int32 = true
	return ftu
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftu *FieldTypeUpdate) SetNillableInt64(i int64) *FieldTypeUpdate {
	ftu.nillable_int64 = &i
	return ftu
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableNillableInt64(i *int64) *FieldTypeUpdate {
	if i != nil {
		ftu.SetNillableInt64(*i)
	}
	return ftu
}

// AddNillableInt64 adds i to nillable_int64.
func (ftu *FieldTypeUpdate) AddNillableInt64(i int64) *FieldTypeUpdate {
	ftu.addnillable_int64 = &i
	return ftu
}

// ClearNillableInt64 clears the value of nillable_int64.
func (ftu *FieldTypeUpdate) ClearNillableInt64() *FieldTypeUpdate {
	ftu.nillable_int64 = nil
	ftu.clearnillable_int64 = true
	return ftu
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftu *FieldTypeUpdate) SetValidateOptionalInt32(i int32) *FieldTypeUpdate {
	ftu.validate_optional_int32 = &i
	return ftu
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftu *FieldTypeUpdate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeUpdate {
	if i != nil {
		ftu.SetValidateOptionalInt32(*i)
	}
	return ftu
}

// ClearValidateOptionalInt32 clears the value of validate_optional_int32.
func (ftu *FieldTypeUpdate) ClearValidateOptionalInt32() *FieldTypeUpdate {
	ftu.validate_optional_int32 = nil
	ftu.clearvalidate_optional_int32 = true
	return ftu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (ftu *FieldTypeUpdate) Save(ctx context.Context) (int, error) {
	if ftu.validate_optional_int32 != nil {
		if err := fieldtype.ValidateOptionalInt32Validator(*ftu.validate_optional_int32); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %v", err)
		}
	}
	switch ftu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftu.sqlSave(ctx)
	case dialect.Gremlin:
		return ftu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftu *FieldTypeUpdate) SaveX(ctx context.Context) int {
	affected, err := ftu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ftu *FieldTypeUpdate) Exec(ctx context.Context) error {
	_, err := ftu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftu *FieldTypeUpdate) ExecX(ctx context.Context) {
	if err := ftu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftu *FieldTypeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(fieldtype.FieldID).From(sql.Table(fieldtype.Table))
	for _, p := range ftu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := ftu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(fieldtype.Table).Where(sql.InInts(fieldtype.FieldID, ids...))
	)
	if value := ftu.int; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt, *value)
	}
	if value := ftu.addint; value != nil {
		update = true
		builder.Add(fieldtype.FieldInt, *value)
	}
	if value := ftu.int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt8, *value)
	}
	if value := ftu.int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt16, *value)
	}
	if value := ftu.int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt32, *value)
	}
	if value := ftu.int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt64, *value)
	}
	if value := ftu.addint64; value != nil {
		update = true
		builder.Add(fieldtype.FieldInt64, *value)
	}
	if value := ftu.optional_int; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt, *value)
	}
	if value := ftu.addoptional_int; value != nil {
		update = true
		builder.Add(fieldtype.FieldOptionalInt, *value)
	}
	if ftu.clearoptional_int {
		update = true
		builder.SetNull(fieldtype.FieldOptionalInt)
	}
	if value := ftu.optional_int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt8, *value)
	}
	if ftu.clearoptional_int8 {
		update = true
		builder.SetNull(fieldtype.FieldOptionalInt8)
	}
	if value := ftu.optional_int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt16, *value)
	}
	if ftu.clearoptional_int16 {
		update = true
		builder.SetNull(fieldtype.FieldOptionalInt16)
	}
	if value := ftu.optional_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt32, *value)
	}
	if ftu.clearoptional_int32 {
		update = true
		builder.SetNull(fieldtype.FieldOptionalInt32)
	}
	if value := ftu.optional_int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt64, *value)
	}
	if value := ftu.addoptional_int64; value != nil {
		update = true
		builder.Add(fieldtype.FieldOptionalInt64, *value)
	}
	if ftu.clearoptional_int64 {
		update = true
		builder.SetNull(fieldtype.FieldOptionalInt64)
	}
	if value := ftu.nillable_int; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt, *value)
	}
	if value := ftu.addnillable_int; value != nil {
		update = true
		builder.Add(fieldtype.FieldNillableInt, *value)
	}
	if ftu.clearnillable_int {
		update = true
		builder.SetNull(fieldtype.FieldNillableInt)
	}
	if value := ftu.nillable_int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt8, *value)
	}
	if ftu.clearnillable_int8 {
		update = true
		builder.SetNull(fieldtype.FieldNillableInt8)
	}
	if value := ftu.nillable_int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt16, *value)
	}
	if ftu.clearnillable_int16 {
		update = true
		builder.SetNull(fieldtype.FieldNillableInt16)
	}
	if value := ftu.nillable_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt32, *value)
	}
	if ftu.clearnillable_int32 {
		update = true
		builder.SetNull(fieldtype.FieldNillableInt32)
	}
	if value := ftu.nillable_int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt64, *value)
	}
	if value := ftu.addnillable_int64; value != nil {
		update = true
		builder.Add(fieldtype.FieldNillableInt64, *value)
	}
	if ftu.clearnillable_int64 {
		update = true
		builder.SetNull(fieldtype.FieldNillableInt64)
	}
	if value := ftu.validate_optional_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldValidateOptionalInt32, *value)
	}
	if ftu.clearvalidate_optional_int32 {
		update = true
		builder.SetNull(fieldtype.FieldValidateOptionalInt32)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (ftu *FieldTypeUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftu.gremlin().Query()
	if err := ftu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (ftu *FieldTypeUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(fieldtype.Label)
	for _, p := range ftu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	if value := ftu.int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *value)
	}
	if value := ftu.addint; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, __.Union(__.Values(fieldtype.FieldInt), __.Constant(*value)).Sum())
	}
	if value := ftu.int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *value)
	}
	if value := ftu.int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *value)
	}
	if value := ftu.int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *value)
	}
	if value := ftu.int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *value)
	}
	if value := ftu.addint64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, __.Union(__.Values(fieldtype.FieldInt64), __.Constant(*value)).Sum())
	}
	if value := ftu.optional_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *value)
	}
	if value := ftu.addoptional_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, __.Union(__.Values(fieldtype.FieldOptionalInt), __.Constant(*value)).Sum())
	}
	if value := ftu.optional_int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *value)
	}
	if value := ftu.optional_int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *value)
	}
	if value := ftu.optional_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *value)
	}
	if value := ftu.optional_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *value)
	}
	if value := ftu.addoptional_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, __.Union(__.Values(fieldtype.FieldOptionalInt64), __.Constant(*value)).Sum())
	}
	if value := ftu.nillable_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, *value)
	}
	if value := ftu.addnillable_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, __.Union(__.Values(fieldtype.FieldNillableInt), __.Constant(*value)).Sum())
	}
	if value := ftu.nillable_int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt8, *value)
	}
	if value := ftu.nillable_int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt16, *value)
	}
	if value := ftu.nillable_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt32, *value)
	}
	if value := ftu.nillable_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, *value)
	}
	if value := ftu.addnillable_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, __.Union(__.Values(fieldtype.FieldNillableInt64), __.Constant(*value)).Sum())
	}
	if value := ftu.validate_optional_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldValidateOptionalInt32, *value)
	}
	var properties []interface{}
	if ftu.clearoptional_int {
		properties = append(properties, fieldtype.FieldOptionalInt)
	}
	if ftu.clearoptional_int8 {
		properties = append(properties, fieldtype.FieldOptionalInt8)
	}
	if ftu.clearoptional_int16 {
		properties = append(properties, fieldtype.FieldOptionalInt16)
	}
	if ftu.clearoptional_int32 {
		properties = append(properties, fieldtype.FieldOptionalInt32)
	}
	if ftu.clearoptional_int64 {
		properties = append(properties, fieldtype.FieldOptionalInt64)
	}
	if ftu.clearnillable_int {
		properties = append(properties, fieldtype.FieldNillableInt)
	}
	if ftu.clearnillable_int8 {
		properties = append(properties, fieldtype.FieldNillableInt8)
	}
	if ftu.clearnillable_int16 {
		properties = append(properties, fieldtype.FieldNillableInt16)
	}
	if ftu.clearnillable_int32 {
		properties = append(properties, fieldtype.FieldNillableInt32)
	}
	if ftu.clearnillable_int64 {
		properties = append(properties, fieldtype.FieldNillableInt64)
	}
	if ftu.clearvalidate_optional_int32 {
		properties = append(properties, fieldtype.FieldValidateOptionalInt32)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// FieldTypeUpdateOne is the builder for updating a single FieldType entity.
type FieldTypeUpdateOne struct {
	config
	id                           string
	int                          *int
	addint                       *int
	int8                         *int8
	int16                        *int16
	int32                        *int32
	int64                        *int64
	addint64                     *int64
	optional_int                 *int
	addoptional_int              *int
	clearoptional_int            bool
	optional_int8                *int8
	clearoptional_int8           bool
	optional_int16               *int16
	clearoptional_int16          bool
	optional_int32               *int32
	clearoptional_int32          bool
	optional_int64               *int64
	addoptional_int64            *int64
	clearoptional_int64          bool
	nillable_int                 *int
	addnillable_int              *int
	clearnillable_int            bool
	nillable_int8                *int8
	clearnillable_int8           bool
	nillable_int16               *int16
	clearnillable_int16          bool
	nillable_int32               *int32
	clearnillable_int32          bool
	nillable_int64               *int64
	addnillable_int64            *int64
	clearnillable_int64          bool
	validate_optional_int32      *int32
	clearvalidate_optional_int32 bool
}

// SetInt sets the int field.
func (ftuo *FieldTypeUpdateOne) SetInt(i int) *FieldTypeUpdateOne {
	ftuo.int = &i
	return ftuo
}

// AddInt adds i to int.
func (ftuo *FieldTypeUpdateOne) AddInt(i int) *FieldTypeUpdateOne {
	ftuo.addint = &i
	return ftuo
}

// SetInt8 sets the int8 field.
func (ftuo *FieldTypeUpdateOne) SetInt8(i int8) *FieldTypeUpdateOne {
	ftuo.int8 = &i
	return ftuo
}

// SetInt16 sets the int16 field.
func (ftuo *FieldTypeUpdateOne) SetInt16(i int16) *FieldTypeUpdateOne {
	ftuo.int16 = &i
	return ftuo
}

// SetInt32 sets the int32 field.
func (ftuo *FieldTypeUpdateOne) SetInt32(i int32) *FieldTypeUpdateOne {
	ftuo.int32 = &i
	return ftuo
}

// SetInt64 sets the int64 field.
func (ftuo *FieldTypeUpdateOne) SetInt64(i int64) *FieldTypeUpdateOne {
	ftuo.int64 = &i
	return ftuo
}

// AddInt64 adds i to int64.
func (ftuo *FieldTypeUpdateOne) AddInt64(i int64) *FieldTypeUpdateOne {
	ftuo.addint64 = &i
	return ftuo
}

// SetOptionalInt sets the optional_int field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt(i int) *FieldTypeUpdateOne {
	ftuo.optional_int = &i
	return ftuo
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt(i *int) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt(*i)
	}
	return ftuo
}

// AddOptionalInt adds i to optional_int.
func (ftuo *FieldTypeUpdateOne) AddOptionalInt(i int) *FieldTypeUpdateOne {
	ftuo.addoptional_int = &i
	return ftuo
}

// ClearOptionalInt clears the value of optional_int.
func (ftuo *FieldTypeUpdateOne) ClearOptionalInt() *FieldTypeUpdateOne {
	ftuo.optional_int = nil
	ftuo.clearoptional_int = true
	return ftuo
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt8(i int8) *FieldTypeUpdateOne {
	ftuo.optional_int8 = &i
	return ftuo
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt8(i *int8) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt8(*i)
	}
	return ftuo
}

// ClearOptionalInt8 clears the value of optional_int8.
func (ftuo *FieldTypeUpdateOne) ClearOptionalInt8() *FieldTypeUpdateOne {
	ftuo.optional_int8 = nil
	ftuo.clearoptional_int8 = true
	return ftuo
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt16(i int16) *FieldTypeUpdateOne {
	ftuo.optional_int16 = &i
	return ftuo
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt16(i *int16) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt16(*i)
	}
	return ftuo
}

// ClearOptionalInt16 clears the value of optional_int16.
func (ftuo *FieldTypeUpdateOne) ClearOptionalInt16() *FieldTypeUpdateOne {
	ftuo.optional_int16 = nil
	ftuo.clearoptional_int16 = true
	return ftuo
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt32(i int32) *FieldTypeUpdateOne {
	ftuo.optional_int32 = &i
	return ftuo
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt32(i *int32) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt32(*i)
	}
	return ftuo
}

// ClearOptionalInt32 clears the value of optional_int32.
func (ftuo *FieldTypeUpdateOne) ClearOptionalInt32() *FieldTypeUpdateOne {
	ftuo.optional_int32 = nil
	ftuo.clearoptional_int32 = true
	return ftuo
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftuo *FieldTypeUpdateOne) SetOptionalInt64(i int64) *FieldTypeUpdateOne {
	ftuo.optional_int64 = &i
	return ftuo
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableOptionalInt64(i *int64) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetOptionalInt64(*i)
	}
	return ftuo
}

// AddOptionalInt64 adds i to optional_int64.
func (ftuo *FieldTypeUpdateOne) AddOptionalInt64(i int64) *FieldTypeUpdateOne {
	ftuo.addoptional_int64 = &i
	return ftuo
}

// ClearOptionalInt64 clears the value of optional_int64.
func (ftuo *FieldTypeUpdateOne) ClearOptionalInt64() *FieldTypeUpdateOne {
	ftuo.optional_int64 = nil
	ftuo.clearoptional_int64 = true
	return ftuo
}

// SetNillableInt sets the nillable_int field.
func (ftuo *FieldTypeUpdateOne) SetNillableInt(i int) *FieldTypeUpdateOne {
	ftuo.nillable_int = &i
	return ftuo
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNillableInt(i *int) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNillableInt(*i)
	}
	return ftuo
}

// AddNillableInt adds i to nillable_int.
func (ftuo *FieldTypeUpdateOne) AddNillableInt(i int) *FieldTypeUpdateOne {
	ftuo.addnillable_int = &i
	return ftuo
}

// ClearNillableInt clears the value of nillable_int.
func (ftuo *FieldTypeUpdateOne) ClearNillableInt() *FieldTypeUpdateOne {
	ftuo.nillable_int = nil
	ftuo.clearnillable_int = true
	return ftuo
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftuo *FieldTypeUpdateOne) SetNillableInt8(i int8) *FieldTypeUpdateOne {
	ftuo.nillable_int8 = &i
	return ftuo
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNillableInt8(i *int8) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNillableInt8(*i)
	}
	return ftuo
}

// ClearNillableInt8 clears the value of nillable_int8.
func (ftuo *FieldTypeUpdateOne) ClearNillableInt8() *FieldTypeUpdateOne {
	ftuo.nillable_int8 = nil
	ftuo.clearnillable_int8 = true
	return ftuo
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftuo *FieldTypeUpdateOne) SetNillableInt16(i int16) *FieldTypeUpdateOne {
	ftuo.nillable_int16 = &i
	return ftuo
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNillableInt16(i *int16) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNillableInt16(*i)
	}
	return ftuo
}

// ClearNillableInt16 clears the value of nillable_int16.
func (ftuo *FieldTypeUpdateOne) ClearNillableInt16() *FieldTypeUpdateOne {
	ftuo.nillable_int16 = nil
	ftuo.clearnillable_int16 = true
	return ftuo
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftuo *FieldTypeUpdateOne) SetNillableInt32(i int32) *FieldTypeUpdateOne {
	ftuo.nillable_int32 = &i
	return ftuo
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNillableInt32(i *int32) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNillableInt32(*i)
	}
	return ftuo
}

// ClearNillableInt32 clears the value of nillable_int32.
func (ftuo *FieldTypeUpdateOne) ClearNillableInt32() *FieldTypeUpdateOne {
	ftuo.nillable_int32 = nil
	ftuo.clearnillable_int32 = true
	return ftuo
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftuo *FieldTypeUpdateOne) SetNillableInt64(i int64) *FieldTypeUpdateOne {
	ftuo.nillable_int64 = &i
	return ftuo
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableNillableInt64(i *int64) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetNillableInt64(*i)
	}
	return ftuo
}

// AddNillableInt64 adds i to nillable_int64.
func (ftuo *FieldTypeUpdateOne) AddNillableInt64(i int64) *FieldTypeUpdateOne {
	ftuo.addnillable_int64 = &i
	return ftuo
}

// ClearNillableInt64 clears the value of nillable_int64.
func (ftuo *FieldTypeUpdateOne) ClearNillableInt64() *FieldTypeUpdateOne {
	ftuo.nillable_int64 = nil
	ftuo.clearnillable_int64 = true
	return ftuo
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftuo *FieldTypeUpdateOne) SetValidateOptionalInt32(i int32) *FieldTypeUpdateOne {
	ftuo.validate_optional_int32 = &i
	return ftuo
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftuo *FieldTypeUpdateOne) SetNillableValidateOptionalInt32(i *int32) *FieldTypeUpdateOne {
	if i != nil {
		ftuo.SetValidateOptionalInt32(*i)
	}
	return ftuo
}

// ClearValidateOptionalInt32 clears the value of validate_optional_int32.
func (ftuo *FieldTypeUpdateOne) ClearValidateOptionalInt32() *FieldTypeUpdateOne {
	ftuo.validate_optional_int32 = nil
	ftuo.clearvalidate_optional_int32 = true
	return ftuo
}

// Save executes the query and returns the updated entity.
func (ftuo *FieldTypeUpdateOne) Save(ctx context.Context) (*FieldType, error) {
	if ftuo.validate_optional_int32 != nil {
		if err := fieldtype.ValidateOptionalInt32Validator(*ftuo.validate_optional_int32); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %v", err)
		}
	}
	switch ftuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftuo.sqlSave(ctx)
	case dialect.Gremlin:
		return ftuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (ftuo *FieldTypeUpdateOne) SaveX(ctx context.Context) *FieldType {
	ft, err := ftuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return ft
}

// Exec executes the query on the entity.
func (ftuo *FieldTypeUpdateOne) Exec(ctx context.Context) error {
	_, err := ftuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ftuo *FieldTypeUpdateOne) ExecX(ctx context.Context) {
	if err := ftuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftuo *FieldTypeUpdateOne) sqlSave(ctx context.Context) (ft *FieldType, err error) {
	selector := sql.Select(fieldtype.Columns...).From(sql.Table(fieldtype.Table))
	fieldtype.ID(ftuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = ftuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		ft = &FieldType{config: ftuo.config}
		if err := ft.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into FieldType: %v", err)
		}
		id = ft.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: FieldType not found with id: %v", ftuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one FieldType with the same id: %v", ftuo.id)
	}

	tx, err := ftuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(fieldtype.Table).Where(sql.InInts(fieldtype.FieldID, ids...))
	)
	if value := ftuo.int; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt, *value)
		ft.Int = *value
	}
	if value := ftuo.addint; value != nil {
		update = true
		builder.Add(fieldtype.FieldInt, *value)
		ft.Int += *value
	}
	if value := ftuo.int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt8, *value)
		ft.Int8 = *value
	}
	if value := ftuo.int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt16, *value)
		ft.Int16 = *value
	}
	if value := ftuo.int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt32, *value)
		ft.Int32 = *value
	}
	if value := ftuo.int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldInt64, *value)
		ft.Int64 = *value
	}
	if value := ftuo.addint64; value != nil {
		update = true
		builder.Add(fieldtype.FieldInt64, *value)
		ft.Int64 += *value
	}
	if value := ftuo.optional_int; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt, *value)
		ft.OptionalInt = *value
	}
	if value := ftuo.addoptional_int; value != nil {
		update = true
		builder.Add(fieldtype.FieldOptionalInt, *value)
		ft.OptionalInt += *value
	}
	if ftuo.clearoptional_int {
		update = true
		var value int
		ft.OptionalInt = value
		builder.SetNull(fieldtype.FieldOptionalInt)
	}
	if value := ftuo.optional_int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt8, *value)
		ft.OptionalInt8 = *value
	}
	if ftuo.clearoptional_int8 {
		update = true
		var value int8
		ft.OptionalInt8 = value
		builder.SetNull(fieldtype.FieldOptionalInt8)
	}
	if value := ftuo.optional_int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt16, *value)
		ft.OptionalInt16 = *value
	}
	if ftuo.clearoptional_int16 {
		update = true
		var value int16
		ft.OptionalInt16 = value
		builder.SetNull(fieldtype.FieldOptionalInt16)
	}
	if value := ftuo.optional_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt32, *value)
		ft.OptionalInt32 = *value
	}
	if ftuo.clearoptional_int32 {
		update = true
		var value int32
		ft.OptionalInt32 = value
		builder.SetNull(fieldtype.FieldOptionalInt32)
	}
	if value := ftuo.optional_int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldOptionalInt64, *value)
		ft.OptionalInt64 = *value
	}
	if value := ftuo.addoptional_int64; value != nil {
		update = true
		builder.Add(fieldtype.FieldOptionalInt64, *value)
		ft.OptionalInt64 += *value
	}
	if ftuo.clearoptional_int64 {
		update = true
		var value int64
		ft.OptionalInt64 = value
		builder.SetNull(fieldtype.FieldOptionalInt64)
	}
	if value := ftuo.nillable_int; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt, *value)
		ft.NillableInt = value
	}
	if value := ftuo.addnillable_int; value != nil {
		update = true
		builder.Add(fieldtype.FieldNillableInt, *value)
		if ft.NillableInt != nil {
			*ft.NillableInt += *value
		} else {
			ft.NillableInt = value
		}
	}
	if ftuo.clearnillable_int {
		update = true
		ft.NillableInt = nil
		builder.SetNull(fieldtype.FieldNillableInt)
	}
	if value := ftuo.nillable_int8; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt8, *value)
		ft.NillableInt8 = value
	}
	if ftuo.clearnillable_int8 {
		update = true
		ft.NillableInt8 = nil
		builder.SetNull(fieldtype.FieldNillableInt8)
	}
	if value := ftuo.nillable_int16; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt16, *value)
		ft.NillableInt16 = value
	}
	if ftuo.clearnillable_int16 {
		update = true
		ft.NillableInt16 = nil
		builder.SetNull(fieldtype.FieldNillableInt16)
	}
	if value := ftuo.nillable_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt32, *value)
		ft.NillableInt32 = value
	}
	if ftuo.clearnillable_int32 {
		update = true
		ft.NillableInt32 = nil
		builder.SetNull(fieldtype.FieldNillableInt32)
	}
	if value := ftuo.nillable_int64; value != nil {
		update = true
		builder.Set(fieldtype.FieldNillableInt64, *value)
		ft.NillableInt64 = value
	}
	if value := ftuo.addnillable_int64; value != nil {
		update = true
		builder.Add(fieldtype.FieldNillableInt64, *value)
		if ft.NillableInt64 != nil {
			*ft.NillableInt64 += *value
		} else {
			ft.NillableInt64 = value
		}
	}
	if ftuo.clearnillable_int64 {
		update = true
		ft.NillableInt64 = nil
		builder.SetNull(fieldtype.FieldNillableInt64)
	}
	if value := ftuo.validate_optional_int32; value != nil {
		update = true
		builder.Set(fieldtype.FieldValidateOptionalInt32, *value)
		ft.ValidateOptionalInt32 = *value
	}
	if ftuo.clearvalidate_optional_int32 {
		update = true
		var value int32
		ft.ValidateOptionalInt32 = value
		builder.SetNull(fieldtype.FieldValidateOptionalInt32)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FieldTypeUpdateOne) gremlinSave(ctx context.Context) (*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftuo.gremlin(ftuo.id).Query()
	if err := ftuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FieldType{config: ftuo.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftuo *FieldTypeUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if value := ftuo.int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *value)
	}
	if value := ftuo.addint; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, __.Union(__.Values(fieldtype.FieldInt), __.Constant(*value)).Sum())
	}
	if value := ftuo.int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *value)
	}
	if value := ftuo.int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *value)
	}
	if value := ftuo.int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *value)
	}
	if value := ftuo.int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *value)
	}
	if value := ftuo.addint64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, __.Union(__.Values(fieldtype.FieldInt64), __.Constant(*value)).Sum())
	}
	if value := ftuo.optional_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *value)
	}
	if value := ftuo.addoptional_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, __.Union(__.Values(fieldtype.FieldOptionalInt), __.Constant(*value)).Sum())
	}
	if value := ftuo.optional_int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *value)
	}
	if value := ftuo.optional_int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *value)
	}
	if value := ftuo.optional_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *value)
	}
	if value := ftuo.optional_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *value)
	}
	if value := ftuo.addoptional_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, __.Union(__.Values(fieldtype.FieldOptionalInt64), __.Constant(*value)).Sum())
	}
	if value := ftuo.nillable_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, *value)
	}
	if value := ftuo.addnillable_int; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, __.Union(__.Values(fieldtype.FieldNillableInt), __.Constant(*value)).Sum())
	}
	if value := ftuo.nillable_int8; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt8, *value)
	}
	if value := ftuo.nillable_int16; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt16, *value)
	}
	if value := ftuo.nillable_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt32, *value)
	}
	if value := ftuo.nillable_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, *value)
	}
	if value := ftuo.addnillable_int64; value != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, __.Union(__.Values(fieldtype.FieldNillableInt64), __.Constant(*value)).Sum())
	}
	if value := ftuo.validate_optional_int32; value != nil {
		v.Property(dsl.Single, fieldtype.FieldValidateOptionalInt32, *value)
	}
	var properties []interface{}
	if ftuo.clearoptional_int {
		properties = append(properties, fieldtype.FieldOptionalInt)
	}
	if ftuo.clearoptional_int8 {
		properties = append(properties, fieldtype.FieldOptionalInt8)
	}
	if ftuo.clearoptional_int16 {
		properties = append(properties, fieldtype.FieldOptionalInt16)
	}
	if ftuo.clearoptional_int32 {
		properties = append(properties, fieldtype.FieldOptionalInt32)
	}
	if ftuo.clearoptional_int64 {
		properties = append(properties, fieldtype.FieldOptionalInt64)
	}
	if ftuo.clearnillable_int {
		properties = append(properties, fieldtype.FieldNillableInt)
	}
	if ftuo.clearnillable_int8 {
		properties = append(properties, fieldtype.FieldNillableInt8)
	}
	if ftuo.clearnillable_int16 {
		properties = append(properties, fieldtype.FieldNillableInt16)
	}
	if ftuo.clearnillable_int32 {
		properties = append(properties, fieldtype.FieldNillableInt32)
	}
	if ftuo.clearnillable_int64 {
		properties = append(properties, fieldtype.FieldNillableInt64)
	}
	if ftuo.clearvalidate_optional_int32 {
		properties = append(properties, fieldtype.FieldValidateOptionalInt32)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
