// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
)

// FieldTypeCreate is the builder for creating a FieldType entity.
type FieldTypeCreate struct {
	config
	int                     *int
	int8                    *int8
	int16                   *int16
	int32                   *int32
	int64                   *int64
	optional_int            *int
	optional_int8           *int8
	optional_int16          *int16
	optional_int32          *int32
	optional_int64          *int64
	nillable_int            *int
	nillable_int8           *int8
	nillable_int16          *int16
	nillable_int32          *int32
	nillable_int64          *int64
	validate_optional_int32 *int32
	state                   *fieldtype.State
}

// SetInt sets the int field.
func (ftc *FieldTypeCreate) SetInt(i int) *FieldTypeCreate {
	ftc.int = &i
	return ftc
}

// SetInt8 sets the int8 field.
func (ftc *FieldTypeCreate) SetInt8(i int8) *FieldTypeCreate {
	ftc.int8 = &i
	return ftc
}

// SetInt16 sets the int16 field.
func (ftc *FieldTypeCreate) SetInt16(i int16) *FieldTypeCreate {
	ftc.int16 = &i
	return ftc
}

// SetInt32 sets the int32 field.
func (ftc *FieldTypeCreate) SetInt32(i int32) *FieldTypeCreate {
	ftc.int32 = &i
	return ftc
}

// SetInt64 sets the int64 field.
func (ftc *FieldTypeCreate) SetInt64(i int64) *FieldTypeCreate {
	ftc.int64 = &i
	return ftc
}

// SetOptionalInt sets the optional_int field.
func (ftc *FieldTypeCreate) SetOptionalInt(i int) *FieldTypeCreate {
	ftc.optional_int = &i
	return ftc
}

// SetNillableOptionalInt sets the optional_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt(*i)
	}
	return ftc
}

// SetOptionalInt8 sets the optional_int8 field.
func (ftc *FieldTypeCreate) SetOptionalInt8(i int8) *FieldTypeCreate {
	ftc.optional_int8 = &i
	return ftc
}

// SetNillableOptionalInt8 sets the optional_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt8(*i)
	}
	return ftc
}

// SetOptionalInt16 sets the optional_int16 field.
func (ftc *FieldTypeCreate) SetOptionalInt16(i int16) *FieldTypeCreate {
	ftc.optional_int16 = &i
	return ftc
}

// SetNillableOptionalInt16 sets the optional_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt16(*i)
	}
	return ftc
}

// SetOptionalInt32 sets the optional_int32 field.
func (ftc *FieldTypeCreate) SetOptionalInt32(i int32) *FieldTypeCreate {
	ftc.optional_int32 = &i
	return ftc
}

// SetNillableOptionalInt32 sets the optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt32(*i)
	}
	return ftc
}

// SetOptionalInt64 sets the optional_int64 field.
func (ftc *FieldTypeCreate) SetOptionalInt64(i int64) *FieldTypeCreate {
	ftc.optional_int64 = &i
	return ftc
}

// SetNillableOptionalInt64 sets the optional_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableOptionalInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetOptionalInt64(*i)
	}
	return ftc
}

// SetNillableInt sets the nillable_int field.
func (ftc *FieldTypeCreate) SetNillableInt(i int) *FieldTypeCreate {
	ftc.nillable_int = &i
	return ftc
}

// SetNillableNillableInt sets the nillable_int field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt(i *int) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt(*i)
	}
	return ftc
}

// SetNillableInt8 sets the nillable_int8 field.
func (ftc *FieldTypeCreate) SetNillableInt8(i int8) *FieldTypeCreate {
	ftc.nillable_int8 = &i
	return ftc
}

// SetNillableNillableInt8 sets the nillable_int8 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt8(i *int8) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt8(*i)
	}
	return ftc
}

// SetNillableInt16 sets the nillable_int16 field.
func (ftc *FieldTypeCreate) SetNillableInt16(i int16) *FieldTypeCreate {
	ftc.nillable_int16 = &i
	return ftc
}

// SetNillableNillableInt16 sets the nillable_int16 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt16(i *int16) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt16(*i)
	}
	return ftc
}

// SetNillableInt32 sets the nillable_int32 field.
func (ftc *FieldTypeCreate) SetNillableInt32(i int32) *FieldTypeCreate {
	ftc.nillable_int32 = &i
	return ftc
}

// SetNillableNillableInt32 sets the nillable_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt32(*i)
	}
	return ftc
}

// SetNillableInt64 sets the nillable_int64 field.
func (ftc *FieldTypeCreate) SetNillableInt64(i int64) *FieldTypeCreate {
	ftc.nillable_int64 = &i
	return ftc
}

// SetNillableNillableInt64 sets the nillable_int64 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableNillableInt64(i *int64) *FieldTypeCreate {
	if i != nil {
		ftc.SetNillableInt64(*i)
	}
	return ftc
}

// SetValidateOptionalInt32 sets the validate_optional_int32 field.
func (ftc *FieldTypeCreate) SetValidateOptionalInt32(i int32) *FieldTypeCreate {
	ftc.validate_optional_int32 = &i
	return ftc
}

// SetNillableValidateOptionalInt32 sets the validate_optional_int32 field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableValidateOptionalInt32(i *int32) *FieldTypeCreate {
	if i != nil {
		ftc.SetValidateOptionalInt32(*i)
	}
	return ftc
}

// SetState sets the state field.
func (ftc *FieldTypeCreate) SetState(f fieldtype.State) *FieldTypeCreate {
	ftc.state = &f
	return ftc
}

// SetNillableState sets the state field if the given value is not nil.
func (ftc *FieldTypeCreate) SetNillableState(f *fieldtype.State) *FieldTypeCreate {
	if f != nil {
		ftc.SetState(*f)
	}
	return ftc
}

// Save creates the FieldType in the database.
func (ftc *FieldTypeCreate) Save(ctx context.Context) (*FieldType, error) {
	if ftc.int == nil {
		return nil, errors.New("ent: missing required field \"int\"")
	}
	if ftc.int8 == nil {
		return nil, errors.New("ent: missing required field \"int8\"")
	}
	if ftc.int16 == nil {
		return nil, errors.New("ent: missing required field \"int16\"")
	}
	if ftc.int32 == nil {
		return nil, errors.New("ent: missing required field \"int32\"")
	}
	if ftc.int64 == nil {
		return nil, errors.New("ent: missing required field \"int64\"")
	}
	if ftc.validate_optional_int32 != nil {
		if err := fieldtype.ValidateOptionalInt32Validator(*ftc.validate_optional_int32); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"validate_optional_int32\": %v", err)
		}
	}
	if ftc.state != nil {
		if err := fieldtype.StateValidator(*ftc.state); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"state\": %v", err)
		}
	}
	switch ftc.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return ftc.sqlSave(ctx)
	case dialect.Gremlin:
		return ftc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (ftc *FieldTypeCreate) SaveX(ctx context.Context) *FieldType {
	v, err := ftc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ftc *FieldTypeCreate) sqlSave(ctx context.Context) (*FieldType, error) {
	var (
		builder = sql.Dialect(ftc.driver.Dialect())
		ft      = &FieldType{config: ftc.config}
	)
	tx, err := ftc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(fieldtype.Table).Default()
	if value := ftc.int; value != nil {
		insert.Set(fieldtype.FieldInt, *value)
		ft.Int = *value
	}
	if value := ftc.int8; value != nil {
		insert.Set(fieldtype.FieldInt8, *value)
		ft.Int8 = *value
	}
	if value := ftc.int16; value != nil {
		insert.Set(fieldtype.FieldInt16, *value)
		ft.Int16 = *value
	}
	if value := ftc.int32; value != nil {
		insert.Set(fieldtype.FieldInt32, *value)
		ft.Int32 = *value
	}
	if value := ftc.int64; value != nil {
		insert.Set(fieldtype.FieldInt64, *value)
		ft.Int64 = *value
	}
	if value := ftc.optional_int; value != nil {
		insert.Set(fieldtype.FieldOptionalInt, *value)
		ft.OptionalInt = *value
	}
	if value := ftc.optional_int8; value != nil {
		insert.Set(fieldtype.FieldOptionalInt8, *value)
		ft.OptionalInt8 = *value
	}
	if value := ftc.optional_int16; value != nil {
		insert.Set(fieldtype.FieldOptionalInt16, *value)
		ft.OptionalInt16 = *value
	}
	if value := ftc.optional_int32; value != nil {
		insert.Set(fieldtype.FieldOptionalInt32, *value)
		ft.OptionalInt32 = *value
	}
	if value := ftc.optional_int64; value != nil {
		insert.Set(fieldtype.FieldOptionalInt64, *value)
		ft.OptionalInt64 = *value
	}
	if value := ftc.nillable_int; value != nil {
		insert.Set(fieldtype.FieldNillableInt, *value)
		ft.NillableInt = value
	}
	if value := ftc.nillable_int8; value != nil {
		insert.Set(fieldtype.FieldNillableInt8, *value)
		ft.NillableInt8 = value
	}
	if value := ftc.nillable_int16; value != nil {
		insert.Set(fieldtype.FieldNillableInt16, *value)
		ft.NillableInt16 = value
	}
	if value := ftc.nillable_int32; value != nil {
		insert.Set(fieldtype.FieldNillableInt32, *value)
		ft.NillableInt32 = value
	}
	if value := ftc.nillable_int64; value != nil {
		insert.Set(fieldtype.FieldNillableInt64, *value)
		ft.NillableInt64 = value
	}
	if value := ftc.validate_optional_int32; value != nil {
		insert.Set(fieldtype.FieldValidateOptionalInt32, *value)
		ft.ValidateOptionalInt32 = *value
	}
	if value := ftc.state; value != nil {
		insert.Set(fieldtype.FieldState, *value)
		ft.State = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(fieldtype.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	ft.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FieldTypeCreate) gremlinSave(ctx context.Context) (*FieldType, error) {
	res := &gremlin.Response{}
	query, bindings := ftc.gremlin().Query()
	if err := ftc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	ft := &FieldType{config: ftc.config}
	if err := ft.FromResponse(res); err != nil {
		return nil, err
	}
	return ft, nil
}

func (ftc *FieldTypeCreate) gremlin() *dsl.Traversal {
	v := g.AddV(fieldtype.Label)
	if ftc.int != nil {
		v.Property(dsl.Single, fieldtype.FieldInt, *ftc.int)
	}
	if ftc.int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt8, *ftc.int8)
	}
	if ftc.int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt16, *ftc.int16)
	}
	if ftc.int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt32, *ftc.int32)
	}
	if ftc.int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldInt64, *ftc.int64)
	}
	if ftc.optional_int != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt, *ftc.optional_int)
	}
	if ftc.optional_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt8, *ftc.optional_int8)
	}
	if ftc.optional_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt16, *ftc.optional_int16)
	}
	if ftc.optional_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt32, *ftc.optional_int32)
	}
	if ftc.optional_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldOptionalInt64, *ftc.optional_int64)
	}
	if ftc.nillable_int != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt, *ftc.nillable_int)
	}
	if ftc.nillable_int8 != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt8, *ftc.nillable_int8)
	}
	if ftc.nillable_int16 != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt16, *ftc.nillable_int16)
	}
	if ftc.nillable_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt32, *ftc.nillable_int32)
	}
	if ftc.nillable_int64 != nil {
		v.Property(dsl.Single, fieldtype.FieldNillableInt64, *ftc.nillable_int64)
	}
	if ftc.validate_optional_int32 != nil {
		v.Property(dsl.Single, fieldtype.FieldValidateOptionalInt32, *ftc.validate_optional_int32)
	}
	if ftc.state != nil {
		v.Property(dsl.Single, fieldtype.FieldState, *ftc.state)
	}
	return v.ValueMap(true)
}
