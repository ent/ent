// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
)

// FieldType is the model entity for the FieldType schema.
type FieldType struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Int holds the value of the "int" field.
	Int int `json:"int,omitempty"`
	// Int8 holds the value of the "int8" field.
	Int8 int8 `json:"int8,omitempty"`
	// Int16 holds the value of the "int16" field.
	Int16 int16 `json:"int16,omitempty"`
	// Int32 holds the value of the "int32" field.
	Int32 int32 `json:"int32,omitempty"`
	// Int64 holds the value of the "int64" field.
	Int64 int64 `json:"int64,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
	// OptionalInt8 holds the value of the "optional_int8" field.
	OptionalInt8 int8 `json:"optional_int8,omitempty"`
	// OptionalInt16 holds the value of the "optional_int16" field.
	OptionalInt16 int16 `json:"optional_int16,omitempty"`
	// OptionalInt32 holds the value of the "optional_int32" field.
	OptionalInt32 int32 `json:"optional_int32,omitempty"`
	// OptionalInt64 holds the value of the "optional_int64" field.
	OptionalInt64 int64 `json:"optional_int64,omitempty"`
	// NillableInt holds the value of the "nillable_int" field.
	NillableInt *int `json:"nillable_int,omitempty"`
	// NillableInt8 holds the value of the "nillable_int8" field.
	NillableInt8 *int8 `json:"nillable_int8,omitempty"`
	// NillableInt16 holds the value of the "nillable_int16" field.
	NillableInt16 *int16 `json:"nillable_int16,omitempty"`
	// NillableInt32 holds the value of the "nillable_int32" field.
	NillableInt32 *int32 `json:"nillable_int32,omitempty"`
	// NillableInt64 holds the value of the "nillable_int64" field.
	NillableInt64 *int64 `json:"nillable_int64,omitempty"`
	// ValidateOptionalInt32 holds the value of the "validate_optional_int32" field.
	ValidateOptionalInt32 int32 `json:"validate_optional_int32,omitempty"`
	// State holds the value of the "state" field.
	State fieldtype.State `json:"state,omitempty"`
}

// FromRows scans the sql response data into FieldType.
func (ft *FieldType) FromRows(rows *sql.Rows) error {
	var vft struct {
		ID                    int
		Int                   sql.NullInt64
		Int8                  sql.NullInt64
		Int16                 sql.NullInt64
		Int32                 sql.NullInt64
		Int64                 sql.NullInt64
		OptionalInt           sql.NullInt64
		OptionalInt8          sql.NullInt64
		OptionalInt16         sql.NullInt64
		OptionalInt32         sql.NullInt64
		OptionalInt64         sql.NullInt64
		NillableInt           sql.NullInt64
		NillableInt8          sql.NullInt64
		NillableInt16         sql.NullInt64
		NillableInt32         sql.NullInt64
		NillableInt64         sql.NullInt64
		ValidateOptionalInt32 sql.NullInt64
		State                 sql.NullString
	}
	// the order here should be the same as in the `fieldtype.Columns`.
	if err := rows.Scan(
		&vft.ID,
		&vft.Int,
		&vft.Int8,
		&vft.Int16,
		&vft.Int32,
		&vft.Int64,
		&vft.OptionalInt,
		&vft.OptionalInt8,
		&vft.OptionalInt16,
		&vft.OptionalInt32,
		&vft.OptionalInt64,
		&vft.NillableInt,
		&vft.NillableInt8,
		&vft.NillableInt16,
		&vft.NillableInt32,
		&vft.NillableInt64,
		&vft.ValidateOptionalInt32,
		&vft.State,
	); err != nil {
		return err
	}
	ft.ID = strconv.Itoa(vft.ID)
	ft.Int = int(vft.Int.Int64)
	ft.Int8 = int8(vft.Int8.Int64)
	ft.Int16 = int16(vft.Int16.Int64)
	ft.Int32 = int32(vft.Int32.Int64)
	ft.Int64 = vft.Int64.Int64
	ft.OptionalInt = int(vft.OptionalInt.Int64)
	ft.OptionalInt8 = int8(vft.OptionalInt8.Int64)
	ft.OptionalInt16 = int16(vft.OptionalInt16.Int64)
	ft.OptionalInt32 = int32(vft.OptionalInt32.Int64)
	ft.OptionalInt64 = vft.OptionalInt64.Int64
	if vft.NillableInt.Valid {
		ft.NillableInt = new(int)
		*ft.NillableInt = int(vft.NillableInt.Int64)
	}
	if vft.NillableInt8.Valid {
		ft.NillableInt8 = new(int8)
		*ft.NillableInt8 = int8(vft.NillableInt8.Int64)
	}
	if vft.NillableInt16.Valid {
		ft.NillableInt16 = new(int16)
		*ft.NillableInt16 = int16(vft.NillableInt16.Int64)
	}
	if vft.NillableInt32.Valid {
		ft.NillableInt32 = new(int32)
		*ft.NillableInt32 = int32(vft.NillableInt32.Int64)
	}
	if vft.NillableInt64.Valid {
		ft.NillableInt64 = new(int64)
		*ft.NillableInt64 = vft.NillableInt64.Int64
	}
	ft.ValidateOptionalInt32 = int32(vft.ValidateOptionalInt32.Int64)
	ft.State = fieldtype.State(vft.State.String)
	return nil
}

// FromResponse scans the gremlin response data into FieldType.
func (ft *FieldType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft struct {
		ID                    string          `json:"id,omitempty"`
		Int                   int             `json:"int,omitempty"`
		Int8                  int8            `json:"int8,omitempty"`
		Int16                 int16           `json:"int16,omitempty"`
		Int32                 int32           `json:"int32,omitempty"`
		Int64                 int64           `json:"int64,omitempty"`
		OptionalInt           int             `json:"optional_int,omitempty"`
		OptionalInt8          int8            `json:"optional_int8,omitempty"`
		OptionalInt16         int16           `json:"optional_int16,omitempty"`
		OptionalInt32         int32           `json:"optional_int32,omitempty"`
		OptionalInt64         int64           `json:"optional_int64,omitempty"`
		NillableInt           *int            `json:"nillable_int,omitempty"`
		NillableInt8          *int8           `json:"nillable_int8,omitempty"`
		NillableInt16         *int16          `json:"nillable_int16,omitempty"`
		NillableInt32         *int32          `json:"nillable_int32,omitempty"`
		NillableInt64         *int64          `json:"nillable_int64,omitempty"`
		ValidateOptionalInt32 int32           `json:"validate_optional_int32,omitempty"`
		State                 fieldtype.State `json:"state,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	ft.ID = vft.ID
	ft.Int = vft.Int
	ft.Int8 = vft.Int8
	ft.Int16 = vft.Int16
	ft.Int32 = vft.Int32
	ft.Int64 = vft.Int64
	ft.OptionalInt = vft.OptionalInt
	ft.OptionalInt8 = vft.OptionalInt8
	ft.OptionalInt16 = vft.OptionalInt16
	ft.OptionalInt32 = vft.OptionalInt32
	ft.OptionalInt64 = vft.OptionalInt64
	ft.NillableInt = vft.NillableInt
	ft.NillableInt8 = vft.NillableInt8
	ft.NillableInt16 = vft.NillableInt16
	ft.NillableInt32 = vft.NillableInt32
	ft.NillableInt64 = vft.NillableInt64
	ft.ValidateOptionalInt32 = vft.ValidateOptionalInt32
	ft.State = vft.State
	return nil
}

// Update returns a builder for updating this FieldType.
// Note that, you need to call FieldType.Unwrap() before calling this method, if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FieldType) Update() *FieldTypeUpdateOne {
	return (&FieldTypeClient{ft.config}).UpdateOne(ft)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (ft *FieldType) Unwrap() *FieldType {
	tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FieldType is not a transactional entity")
	}
	ft.config.driver = tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FieldType) String() string {
	var builder strings.Builder
	builder.WriteString("FieldType(")
	builder.WriteString(fmt.Sprintf("id=%v", ft.ID))
	builder.WriteString(", int=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int))
	builder.WriteString(", int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int8))
	builder.WriteString(", int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int16))
	builder.WriteString(", int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int32))
	builder.WriteString(", int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int64))
	builder.WriteString(", optional_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt))
	builder.WriteString(", optional_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt8))
	builder.WriteString(", optional_int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt16))
	builder.WriteString(", optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt32))
	builder.WriteString(", optional_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt64))
	if v := ft.NillableInt; v != nil {
		builder.WriteString(", nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt8; v != nil {
		builder.WriteString(", nillable_int8=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt16; v != nil {
		builder.WriteString(", nillable_int16=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt32; v != nil {
		builder.WriteString(", nillable_int32=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	if v := ft.NillableInt64; v != nil {
		builder.WriteString(", nillable_int64=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", validate_optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.ValidateOptionalInt32))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (ft *FieldType) id() int {
	id, _ := strconv.Atoi(ft.ID)
	return id
}

// FieldTypes is a parsable slice of FieldType.
type FieldTypes []*FieldType

// FromRows scans the sql response data into FieldTypes.
func (ft *FieldTypes) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vft := &FieldType{}
		if err := vft.FromRows(rows); err != nil {
			return err
		}
		*ft = append(*ft, vft)
	}
	return nil
}

// FromResponse scans the gremlin response data into FieldTypes.
func (ft *FieldTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vft []struct {
		ID                    string          `json:"id,omitempty"`
		Int                   int             `json:"int,omitempty"`
		Int8                  int8            `json:"int8,omitempty"`
		Int16                 int16           `json:"int16,omitempty"`
		Int32                 int32           `json:"int32,omitempty"`
		Int64                 int64           `json:"int64,omitempty"`
		OptionalInt           int             `json:"optional_int,omitempty"`
		OptionalInt8          int8            `json:"optional_int8,omitempty"`
		OptionalInt16         int16           `json:"optional_int16,omitempty"`
		OptionalInt32         int32           `json:"optional_int32,omitempty"`
		OptionalInt64         int64           `json:"optional_int64,omitempty"`
		NillableInt           *int            `json:"nillable_int,omitempty"`
		NillableInt8          *int8           `json:"nillable_int8,omitempty"`
		NillableInt16         *int16          `json:"nillable_int16,omitempty"`
		NillableInt32         *int32          `json:"nillable_int32,omitempty"`
		NillableInt64         *int64          `json:"nillable_int64,omitempty"`
		ValidateOptionalInt32 int32           `json:"validate_optional_int32,omitempty"`
		State                 fieldtype.State `json:"state,omitempty"`
	}
	if err := vmap.Decode(&vft); err != nil {
		return err
	}
	for _, v := range vft {
		*ft = append(*ft, &FieldType{
			ID:                    v.ID,
			Int:                   v.Int,
			Int8:                  v.Int8,
			Int16:                 v.Int16,
			Int32:                 v.Int32,
			Int64:                 v.Int64,
			OptionalInt:           v.OptionalInt,
			OptionalInt8:          v.OptionalInt8,
			OptionalInt16:         v.OptionalInt16,
			OptionalInt32:         v.OptionalInt32,
			OptionalInt64:         v.OptionalInt64,
			NillableInt:           v.NillableInt,
			NillableInt8:          v.NillableInt8,
			NillableInt16:         v.NillableInt16,
			NillableInt32:         v.NillableInt32,
			NillableInt64:         v.NillableInt64,
			ValidateOptionalInt32: v.ValidateOptionalInt32,
			State:                 v.State,
		})
	}
	return nil
}

func (ft FieldTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
