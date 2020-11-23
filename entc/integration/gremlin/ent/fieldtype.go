// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/facebook/ent/dialect/gremlin"
	"github.com/facebook/ent/entc/integration/ent/role"
	"github.com/facebook/ent/entc/integration/ent/schema"
	"github.com/facebook/ent/entc/integration/gremlin/ent/fieldtype"
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
	// OptionalUint holds the value of the "optional_uint" field.
	OptionalUint uint `json:"optional_uint,omitempty"`
	// OptionalUint8 holds the value of the "optional_uint8" field.
	OptionalUint8 uint8 `json:"optional_uint8,omitempty"`
	// OptionalUint16 holds the value of the "optional_uint16" field.
	OptionalUint16 uint16 `json:"optional_uint16,omitempty"`
	// OptionalUint32 holds the value of the "optional_uint32" field.
	OptionalUint32 uint32 `json:"optional_uint32,omitempty"`
	// OptionalUint64 holds the value of the "optional_uint64" field.
	OptionalUint64 uint64 `json:"optional_uint64,omitempty"`
	// State holds the value of the "state" field.
	State fieldtype.State `json:"state,omitempty"`
	// OptionalFloat holds the value of the "optional_float" field.
	OptionalFloat float64 `json:"optional_float,omitempty"`
	// OptionalFloat32 holds the value of the "optional_float32" field.
	OptionalFloat32 float32 `json:"optional_float32,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Decimal holds the value of the "decimal" field.
	Decimal float64 `json:"decimal,omitempty"`
	// Dir holds the value of the "dir" field.
	Dir http.Dir `json:"dir,omitempty"`
	// Ndir holds the value of the "ndir" field.
	Ndir *http.Dir `json:"ndir,omitempty"`
	// Str holds the value of the "str" field.
	Str sql.NullString `json:"str,omitempty"`
	// NullStr holds the value of the "null_str" field.
	NullStr *sql.NullString `json:"null_str,omitempty"`
	// Link holds the value of the "link" field.
	Link schema.Link `json:"link,omitempty"`
	// NullLink holds the value of the "null_link" field.
	NullLink *schema.Link `json:"null_link,omitempty"`
	// Active holds the value of the "active" field.
	Active schema.Status `json:"active,omitempty"`
	// NullActive holds the value of the "null_active" field.
	NullActive *schema.Status `json:"null_active,omitempty"`
	// Deleted holds the value of the "deleted" field.
	Deleted sql.NullBool `json:"deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt sql.NullTime `json:"deleted_at,omitempty"`
	// IP holds the value of the "ip" field.
	IP net.IP `json:"ip,omitempty"`
	// NullInt64 holds the value of the "null_int64" field.
	NullInt64 sql.NullInt64 `json:"null_int64,omitempty"`
	// SchemaInt holds the value of the "schema_int" field.
	SchemaInt schema.Int `json:"schema_int,omitempty"`
	// SchemaInt8 holds the value of the "schema_int8" field.
	SchemaInt8 schema.Int8 `json:"schema_int8,omitempty"`
	// SchemaInt64 holds the value of the "schema_int64" field.
	SchemaInt64 schema.Int64 `json:"schema_int64,omitempty"`
	// SchemaFloat holds the value of the "schema_float" field.
	SchemaFloat schema.Float64 `json:"schema_float,omitempty"`
	// SchemaFloat32 holds the value of the "schema_float32" field.
	SchemaFloat32 schema.Float32 `json:"schema_float32,omitempty"`
	// NullFloat holds the value of the "null_float" field.
	NullFloat sql.NullFloat64 `json:"null_float,omitempty"`
	// Role holds the value of the "role" field.
	Role role.Role `json:"role,omitempty"`
	// MAC holds the value of the "mac" field.
	MAC schema.MAC `json:"mac,omitempty"`
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
		OptionalUint          uint            `json:"optional_uint,omitempty"`
		OptionalUint8         uint8           `json:"optional_uint8,omitempty"`
		OptionalUint16        uint16          `json:"optional_uint16,omitempty"`
		OptionalUint32        uint32          `json:"optional_uint32,omitempty"`
		OptionalUint64        uint64          `json:"optional_uint64,omitempty"`
		State                 fieldtype.State `json:"state,omitempty"`
		OptionalFloat         float64         `json:"optional_float,omitempty"`
		OptionalFloat32       float32         `json:"optional_float32,omitempty"`
		Datetime              int64           `json:"datetime,omitempty"`
		Decimal               float64         `json:"decimal,omitempty"`
		Dir                   http.Dir        `json:"dir,omitempty"`
		Ndir                  *http.Dir       `json:"ndir,omitempty"`
		Str                   sql.NullString  `json:"str,omitempty"`
		NullStr               *sql.NullString `json:"null_str,omitempty"`
		Link                  schema.Link     `json:"link,omitempty"`
		NullLink              *schema.Link    `json:"null_link,omitempty"`
		Active                schema.Status   `json:"active,omitempty"`
		NullActive            *schema.Status  `json:"null_active,omitempty"`
		Deleted               sql.NullBool    `json:"deleted,omitempty"`
		DeletedAt             sql.NullTime    `json:"deleted_at,omitempty"`
		IP                    net.IP          `json:"ip,omitempty"`
		NullInt64             sql.NullInt64   `json:"null_int64,omitempty"`
		SchemaInt             schema.Int      `json:"schema_int,omitempty"`
		SchemaInt8            schema.Int8     `json:"schema_int8,omitempty"`
		SchemaInt64           schema.Int64    `json:"schema_int64,omitempty"`
		SchemaFloat           schema.Float64  `json:"schema_float,omitempty"`
		SchemaFloat32         schema.Float32  `json:"schema_float32,omitempty"`
		NullFloat             sql.NullFloat64 `json:"null_float,omitempty"`
		Role                  role.Role       `json:"role,omitempty"`
		MAC                   schema.MAC      `json:"mac,omitempty"`
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
	ft.OptionalUint = scanft.OptionalUint
	ft.OptionalUint8 = scanft.OptionalUint8
	ft.OptionalUint16 = scanft.OptionalUint16
	ft.OptionalUint32 = scanft.OptionalUint32
	ft.OptionalUint64 = scanft.OptionalUint64
	ft.State = scanft.State
	ft.OptionalFloat = scanft.OptionalFloat
	ft.OptionalFloat32 = scanft.OptionalFloat32
	ft.Datetime = time.Unix(0, scanft.Datetime)
	ft.Decimal = scanft.Decimal
	ft.Dir = scanft.Dir
	ft.Ndir = scanft.Ndir
	ft.Str = scanft.Str
	ft.NullStr = scanft.NullStr
	ft.Link = scanft.Link
	ft.NullLink = scanft.NullLink
	ft.Active = scanft.Active
	ft.NullActive = scanft.NullActive
	ft.Deleted = scanft.Deleted
	ft.DeletedAt = scanft.DeletedAt
	ft.IP = scanft.IP
	ft.NullInt64 = scanft.NullInt64
	ft.SchemaInt = scanft.SchemaInt
	ft.SchemaInt8 = scanft.SchemaInt8
	ft.SchemaInt64 = scanft.SchemaInt64
	ft.SchemaFloat = scanft.SchemaFloat
	ft.SchemaFloat32 = scanft.SchemaFloat32
	ft.NullFloat = scanft.NullFloat
	ft.Role = scanft.Role
	ft.MAC = scanft.MAC
	return nil
}

// Update returns a builder for updating this FieldType.
// Note that, you need to call FieldType.Unwrap() before calling this method, if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FieldType) Update() *FieldTypeUpdateOne {
	return (&FieldTypeClient{config: ft.config}).UpdateOne(ft)
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
	builder.WriteString(", optional_uint=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint))
	builder.WriteString(", optional_uint8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint8))
	builder.WriteString(", optional_uint16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint16))
	builder.WriteString(", optional_uint32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint32))
	builder.WriteString(", optional_uint64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint64))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteString(", optional_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat))
	builder.WriteString(", optional_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat32))
	builder.WriteString(", datetime=")
	builder.WriteString(ft.Datetime.Format(time.ANSIC))
	builder.WriteString(", decimal=")
	builder.WriteString(fmt.Sprintf("%v", ft.Decimal))
	builder.WriteString(", dir=")
	builder.WriteString(fmt.Sprintf("%v", ft.Dir))
	if v := ft.Ndir; v != nil {
		builder.WriteString(", ndir=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", str=")
	builder.WriteString(fmt.Sprintf("%v", ft.Str))
	if v := ft.NullStr; v != nil {
		builder.WriteString(", null_str=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", link=")
	builder.WriteString(fmt.Sprintf("%v", ft.Link))
	if v := ft.NullLink; v != nil {
		builder.WriteString(", null_link=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", active=")
	builder.WriteString(fmt.Sprintf("%v", ft.Active))
	if v := ft.NullActive; v != nil {
		builder.WriteString(", null_active=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", deleted=")
	builder.WriteString(fmt.Sprintf("%v", ft.Deleted))
	builder.WriteString(", deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ft.DeletedAt))
	builder.WriteString(", ip=")
	builder.WriteString(fmt.Sprintf("%v", ft.IP))
	builder.WriteString(", null_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullInt64))
	builder.WriteString(", schema_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt))
	builder.WriteString(", schema_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt8))
	builder.WriteString(", schema_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt64))
	builder.WriteString(", schema_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat))
	builder.WriteString(", schema_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat32))
	builder.WriteString(", null_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullFloat))
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", ft.Role))
	builder.WriteString(", mac=")
	builder.WriteString(fmt.Sprintf("%v", ft.MAC))
	builder.WriteByte(')')
	return builder.String()
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
		OptionalUint          uint            `json:"optional_uint,omitempty"`
		OptionalUint8         uint8           `json:"optional_uint8,omitempty"`
		OptionalUint16        uint16          `json:"optional_uint16,omitempty"`
		OptionalUint32        uint32          `json:"optional_uint32,omitempty"`
		OptionalUint64        uint64          `json:"optional_uint64,omitempty"`
		State                 fieldtype.State `json:"state,omitempty"`
		OptionalFloat         float64         `json:"optional_float,omitempty"`
		OptionalFloat32       float32         `json:"optional_float32,omitempty"`
		Datetime              int64           `json:"datetime,omitempty"`
		Decimal               float64         `json:"decimal,omitempty"`
		Dir                   http.Dir        `json:"dir,omitempty"`
		Ndir                  *http.Dir       `json:"ndir,omitempty"`
		Str                   sql.NullString  `json:"str,omitempty"`
		NullStr               *sql.NullString `json:"null_str,omitempty"`
		Link                  schema.Link     `json:"link,omitempty"`
		NullLink              *schema.Link    `json:"null_link,omitempty"`
		Active                schema.Status   `json:"active,omitempty"`
		NullActive            *schema.Status  `json:"null_active,omitempty"`
		Deleted               sql.NullBool    `json:"deleted,omitempty"`
		DeletedAt             sql.NullTime    `json:"deleted_at,omitempty"`
		IP                    net.IP          `json:"ip,omitempty"`
		NullInt64             sql.NullInt64   `json:"null_int64,omitempty"`
		SchemaInt             schema.Int      `json:"schema_int,omitempty"`
		SchemaInt8            schema.Int8     `json:"schema_int8,omitempty"`
		SchemaInt64           schema.Int64    `json:"schema_int64,omitempty"`
		SchemaFloat           schema.Float64  `json:"schema_float,omitempty"`
		SchemaFloat32         schema.Float32  `json:"schema_float32,omitempty"`
		NullFloat             sql.NullFloat64 `json:"null_float,omitempty"`
		Role                  role.Role       `json:"role,omitempty"`
		MAC                   schema.MAC      `json:"mac,omitempty"`
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
			OptionalUint:          v.OptionalUint,
			OptionalUint8:         v.OptionalUint8,
			OptionalUint16:        v.OptionalUint16,
			OptionalUint32:        v.OptionalUint32,
			OptionalUint64:        v.OptionalUint64,
			State:                 v.State,
			OptionalFloat:         v.OptionalFloat,
			OptionalFloat32:       v.OptionalFloat32,
			Datetime:              time.Unix(0, v.Datetime),
			Decimal:               v.Decimal,
			Dir:                   v.Dir,
			Ndir:                  v.Ndir,
			Str:                   v.Str,
			NullStr:               v.NullStr,
			Link:                  v.Link,
			NullLink:              v.NullLink,
			Active:                v.Active,
			NullActive:            v.NullActive,
			Deleted:               v.Deleted,
			DeletedAt:             v.DeletedAt,
			IP:                    v.IP,
			NullInt64:             v.NullInt64,
			SchemaInt:             v.SchemaInt,
			SchemaInt8:            v.SchemaInt8,
			SchemaInt64:           v.SchemaInt64,
			SchemaFloat:           v.SchemaFloat,
			SchemaFloat32:         v.SchemaFloat32,
			NullFloat:             v.NullFloat,
			Role:                  v.Role,
			MAC:                   v.MAC,
		})
	}
	return nil
}

func (ft FieldTypes) config(cfg config) {
	for _i := range ft {
		ft[_i].config = cfg
	}
}
