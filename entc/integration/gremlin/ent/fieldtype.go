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
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/fieldtype"
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

// FromResponse scans the gremlin response data into FieldType.
func (ft *FieldType) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanft struct {
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
	if err := vmap.Decode(&scanft); err != nil {
		return err
	}
	ft.ID = scanft.ID
	ft.Int = scanft.Int
	ft.Int8 = scanft.Int8
	ft.Int16 = scanft.Int16
	ft.Int32 = scanft.Int32
	ft.Int64 = scanft.Int64
	ft.OptionalInt = scanft.OptionalInt
	ft.OptionalInt8 = scanft.OptionalInt8
	ft.OptionalInt16 = scanft.OptionalInt16
	ft.OptionalInt32 = scanft.OptionalInt32
	ft.OptionalInt64 = scanft.OptionalInt64
	ft.NillableInt = scanft.NillableInt
	ft.NillableInt8 = scanft.NillableInt8
	ft.NillableInt16 = scanft.NillableInt16
	ft.NillableInt32 = scanft.NillableInt32
	ft.NillableInt64 = scanft.NillableInt64
	ft.ValidateOptionalInt32 = scanft.ValidateOptionalInt32
	ft.State = scanft.State
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

// FromResponse scans the gremlin response data into FieldTypes.
func (ft *FieldTypes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanft []struct {
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
	if err := vmap.Decode(&scanft); err != nil {
		return err
	}
	for _, v := range scanft {
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
