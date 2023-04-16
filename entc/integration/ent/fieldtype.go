// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/fieldtype"
	"entgo.io/ent/entc/integration/ent/role"
	"entgo.io/ent/entc/integration/ent/schema"
	"github.com/google/uuid"
)

// FieldType is the model entity for the FieldType schema.
type FieldType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
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
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Decimal holds the value of the "decimal" field.
	Decimal float64 `json:"decimal,omitempty"`
	// LinkOther holds the value of the "link_other" field.
	LinkOther *schema.Link `json:"link_other,omitempty"`
	// LinkOtherFunc holds the value of the "link_other_func" field.
	LinkOtherFunc *schema.Link `json:"link_other_func,omitempty"`
	// MAC holds the value of the "mac" field.
	MAC schema.MAC `json:"mac,omitempty"`
	// StringArray holds the value of the "string_array" field.
	StringArray schema.Strings `json:"string_array,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"-"`
	// StringScanner holds the value of the "string_scanner" field.
	StringScanner *schema.StringScanner `json:"string_scanner,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration time.Duration `json:"duration,omitempty"`
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
	Deleted *sql.NullBool `json:"deleted,omitempty"`
	// DeletedAt holds the value of the "deleted_at" field.
	DeletedAt *sql.NullTime `json:"deleted_at,omitempty"`
	// RawData holds the value of the "raw_data" field.
	RawData []byte `json:"raw_data,omitempty"`
	// Sensitive holds the value of the "sensitive" field.
	Sensitive []byte `json:"-"`
	// IP holds the value of the "ip" field.
	IP net.IP `json:"ip,omitempty"`
	// NullInt64 holds the value of the "null_int64" field.
	NullInt64 *sql.NullInt64 `json:"null_int64,omitempty"`
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
	NullFloat *sql.NullFloat64 `json:"null_float,omitempty"`
	// Role holds the value of the "role" field.
	Role role.Role `json:"role,omitempty"`
	// Priority holds the value of the "priority" field.
	Priority role.Priority `json:"priority,omitempty"`
	// OptionalUUID holds the value of the "optional_uuid" field.
	OptionalUUID uuid.UUID `json:"optional_uuid,omitempty"`
	// NillableUUID holds the value of the "nillable_uuid" field.
	NillableUUID *uuid.UUID `json:"nillable_uuid,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
	// Pair holds the value of the "pair" field.
	Pair schema.Pair `json:"pair,omitempty"`
	// NilPair holds the value of the "nil_pair" field.
	NilPair *schema.Pair `json:"nil_pair,omitempty"`
	// Vstring holds the value of the "vstring" field.
	Vstring schema.VString `json:"vstring,omitempty"`
	// Triple holds the value of the "triple" field.
	Triple schema.Triple `json:"triple,omitempty"`
	// BigInt holds the value of the "big_int" field.
	BigInt schema.BigInt `json:"big_int,omitempty"`
	// PasswordOther holds the value of the "password_other" field.
	PasswordOther schema.Password `json:"-"`
	file_field    *int
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FieldType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fieldtype.FieldNullLink:
			values[i] = &sql.NullScanner{S: new(schema.Link)}
		case fieldtype.FieldNilPair:
			values[i] = &sql.NullScanner{S: new(schema.Pair)}
		case fieldtype.FieldStringScanner:
			values[i] = &sql.NullScanner{S: new(schema.StringScanner)}
		case fieldtype.FieldNillableUUID:
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case fieldtype.FieldRawData, fieldtype.FieldSensitive, fieldtype.FieldIP, fieldtype.FieldStrings:
			values[i] = new([]byte)
		case fieldtype.FieldPriority:
			values[i] = new(role.Priority)
		case fieldtype.FieldBigInt:
			values[i] = new(schema.BigInt)
		case fieldtype.FieldLinkOther, fieldtype.FieldLinkOtherFunc, fieldtype.FieldLink:
			values[i] = new(schema.Link)
		case fieldtype.FieldMAC:
			values[i] = new(schema.MAC)
		case fieldtype.FieldPair:
			values[i] = new(schema.Pair)
		case fieldtype.FieldPasswordOther:
			values[i] = new(schema.Password)
		case fieldtype.FieldStringArray:
			values[i] = new(schema.Strings)
		case fieldtype.FieldTriple:
			values[i] = new(schema.Triple)
		case fieldtype.FieldVstring:
			values[i] = new(schema.VString)
		case fieldtype.FieldActive, fieldtype.FieldNullActive, fieldtype.FieldDeleted:
			values[i] = new(sql.NullBool)
		case fieldtype.FieldOptionalFloat, fieldtype.FieldOptionalFloat32, fieldtype.FieldDecimal, fieldtype.FieldSchemaFloat, fieldtype.FieldSchemaFloat32, fieldtype.FieldNullFloat:
			values[i] = new(sql.NullFloat64)
		case fieldtype.FieldID, fieldtype.FieldInt, fieldtype.FieldInt8, fieldtype.FieldInt16, fieldtype.FieldInt32, fieldtype.FieldInt64, fieldtype.FieldOptionalInt, fieldtype.FieldOptionalInt8, fieldtype.FieldOptionalInt16, fieldtype.FieldOptionalInt32, fieldtype.FieldOptionalInt64, fieldtype.FieldNillableInt, fieldtype.FieldNillableInt8, fieldtype.FieldNillableInt16, fieldtype.FieldNillableInt32, fieldtype.FieldNillableInt64, fieldtype.FieldValidateOptionalInt32, fieldtype.FieldOptionalUint, fieldtype.FieldOptionalUint8, fieldtype.FieldOptionalUint16, fieldtype.FieldOptionalUint32, fieldtype.FieldOptionalUint64, fieldtype.FieldDuration, fieldtype.FieldNullInt64, fieldtype.FieldSchemaInt, fieldtype.FieldSchemaInt8, fieldtype.FieldSchemaInt64:
			values[i] = new(sql.NullInt64)
		case fieldtype.FieldState, fieldtype.FieldText, fieldtype.FieldPassword, fieldtype.FieldDir, fieldtype.FieldNdir, fieldtype.FieldStr, fieldtype.FieldNullStr, fieldtype.FieldRole:
			values[i] = new(sql.NullString)
		case fieldtype.FieldDatetime, fieldtype.FieldDeletedAt:
			values[i] = new(sql.NullTime)
		case fieldtype.FieldOptionalUUID:
			values[i] = new(uuid.UUID)
		case fieldtype.ForeignKeys[0]: // file_field
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FieldType fields.
func (ft *FieldType) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fieldtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ft.ID = int(value.Int64)
		case fieldtype.FieldInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int", values[i])
			} else if value.Valid {
				ft.Int = int(value.Int64)
			}
		case fieldtype.FieldInt8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int8", values[i])
			} else if value.Valid {
				ft.Int8 = int8(value.Int64)
			}
		case fieldtype.FieldInt16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int16", values[i])
			} else if value.Valid {
				ft.Int16 = int16(value.Int64)
			}
		case fieldtype.FieldInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int32", values[i])
			} else if value.Valid {
				ft.Int32 = int32(value.Int64)
			}
		case fieldtype.FieldInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int64", values[i])
			} else if value.Valid {
				ft.Int64 = value.Int64
			}
		case fieldtype.FieldOptionalInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int", values[i])
			} else if value.Valid {
				ft.OptionalInt = int(value.Int64)
			}
		case fieldtype.FieldOptionalInt8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int8", values[i])
			} else if value.Valid {
				ft.OptionalInt8 = int8(value.Int64)
			}
		case fieldtype.FieldOptionalInt16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int16", values[i])
			} else if value.Valid {
				ft.OptionalInt16 = int16(value.Int64)
			}
		case fieldtype.FieldOptionalInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int32", values[i])
			} else if value.Valid {
				ft.OptionalInt32 = int32(value.Int64)
			}
		case fieldtype.FieldOptionalInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_int64", values[i])
			} else if value.Valid {
				ft.OptionalInt64 = value.Int64
			}
		case fieldtype.FieldNillableInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int", values[i])
			} else if value.Valid {
				ft.NillableInt = new(int)
				*ft.NillableInt = int(value.Int64)
			}
		case fieldtype.FieldNillableInt8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int8", values[i])
			} else if value.Valid {
				ft.NillableInt8 = new(int8)
				*ft.NillableInt8 = int8(value.Int64)
			}
		case fieldtype.FieldNillableInt16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int16", values[i])
			} else if value.Valid {
				ft.NillableInt16 = new(int16)
				*ft.NillableInt16 = int16(value.Int64)
			}
		case fieldtype.FieldNillableInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int32", values[i])
			} else if value.Valid {
				ft.NillableInt32 = new(int32)
				*ft.NillableInt32 = int32(value.Int64)
			}
		case fieldtype.FieldNillableInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_int64", values[i])
			} else if value.Valid {
				ft.NillableInt64 = new(int64)
				*ft.NillableInt64 = value.Int64
			}
		case fieldtype.FieldValidateOptionalInt32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field validate_optional_int32", values[i])
			} else if value.Valid {
				ft.ValidateOptionalInt32 = int32(value.Int64)
			}
		case fieldtype.FieldOptionalUint:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uint", values[i])
			} else if value.Valid {
				ft.OptionalUint = uint(value.Int64)
			}
		case fieldtype.FieldOptionalUint8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uint8", values[i])
			} else if value.Valid {
				ft.OptionalUint8 = uint8(value.Int64)
			}
		case fieldtype.FieldOptionalUint16:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uint16", values[i])
			} else if value.Valid {
				ft.OptionalUint16 = uint16(value.Int64)
			}
		case fieldtype.FieldOptionalUint32:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uint32", values[i])
			} else if value.Valid {
				ft.OptionalUint32 = uint32(value.Int64)
			}
		case fieldtype.FieldOptionalUint64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uint64", values[i])
			} else if value.Valid {
				ft.OptionalUint64 = uint64(value.Int64)
			}
		case fieldtype.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				ft.State = fieldtype.State(value.String)
			}
		case fieldtype.FieldOptionalFloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_float", values[i])
			} else if value.Valid {
				ft.OptionalFloat = value.Float64
			}
		case fieldtype.FieldOptionalFloat32:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field optional_float32", values[i])
			} else if value.Valid {
				ft.OptionalFloat32 = float32(value.Float64)
			}
		case fieldtype.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				ft.Text = value.String
			}
		case fieldtype.FieldDatetime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field datetime", values[i])
			} else if value.Valid {
				ft.Datetime = value.Time
			}
		case fieldtype.FieldDecimal:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field decimal", values[i])
			} else if value.Valid {
				ft.Decimal = value.Float64
			}
		case fieldtype.FieldLinkOther:
			if value, ok := values[i].(*schema.Link); !ok {
				return fmt.Errorf("unexpected type %T for field link_other", values[i])
			} else if value != nil {
				ft.LinkOther = value
			}
		case fieldtype.FieldLinkOtherFunc:
			if value, ok := values[i].(*schema.Link); !ok {
				return fmt.Errorf("unexpected type %T for field link_other_func", values[i])
			} else if value != nil {
				ft.LinkOtherFunc = value
			}
		case fieldtype.FieldMAC:
			if value, ok := values[i].(*schema.MAC); !ok {
				return fmt.Errorf("unexpected type %T for field mac", values[i])
			} else if value != nil {
				ft.MAC = *value
			}
		case fieldtype.FieldStringArray:
			if value, ok := values[i].(*schema.Strings); !ok {
				return fmt.Errorf("unexpected type %T for field string_array", values[i])
			} else if value != nil {
				ft.StringArray = *value
			}
		case fieldtype.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				ft.Password = value.String
			}
		case fieldtype.FieldStringScanner:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field string_scanner", values[i])
			} else if value.Valid {
				ft.StringScanner = new(schema.StringScanner)
				*ft.StringScanner = *value.S.(*schema.StringScanner)
			}
		case fieldtype.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				ft.Duration = time.Duration(value.Int64)
			}
		case fieldtype.FieldDir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field dir", values[i])
			} else if value.Valid {
				ft.Dir = http.Dir(value.String)
			}
		case fieldtype.FieldNdir:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ndir", values[i])
			} else if value.Valid {
				ft.Ndir = new(http.Dir)
				*ft.Ndir = http.Dir(value.String)
			}
		case fieldtype.FieldStr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field str", values[i])
			} else if value.Valid {
				ft.Str = *value
			}
		case fieldtype.FieldNullStr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field null_str", values[i])
			} else if value.Valid {
				ft.NullStr = value
			}
		case fieldtype.FieldLink:
			if value, ok := values[i].(*schema.Link); !ok {
				return fmt.Errorf("unexpected type %T for field link", values[i])
			} else if value != nil {
				ft.Link = *value
			}
		case fieldtype.FieldNullLink:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field null_link", values[i])
			} else if value.Valid {
				ft.NullLink = value.S.(*schema.Link)
			}
		case fieldtype.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				ft.Active = schema.Status(value.Bool)
			}
		case fieldtype.FieldNullActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field null_active", values[i])
			} else if value.Valid {
				ft.NullActive = new(schema.Status)
				*ft.NullActive = schema.Status(value.Bool)
			}
		case fieldtype.FieldDeleted:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field deleted", values[i])
			} else if value.Valid {
				ft.Deleted = value
			}
		case fieldtype.FieldDeletedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_at", values[i])
			} else if value.Valid {
				ft.DeletedAt = value
			}
		case fieldtype.FieldRawData:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field raw_data", values[i])
			} else if value != nil {
				ft.RawData = *value
			}
		case fieldtype.FieldSensitive:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field sensitive", values[i])
			} else if value != nil {
				ft.Sensitive = *value
			}
		case fieldtype.FieldIP:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field ip", values[i])
			} else if value != nil {
				ft.IP = *value
			}
		case fieldtype.FieldNullInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field null_int64", values[i])
			} else if value.Valid {
				ft.NullInt64 = value
			}
		case fieldtype.FieldSchemaInt:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field schema_int", values[i])
			} else if value.Valid {
				ft.SchemaInt = schema.Int(value.Int64)
			}
		case fieldtype.FieldSchemaInt8:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field schema_int8", values[i])
			} else if value.Valid {
				ft.SchemaInt8 = schema.Int8(value.Int64)
			}
		case fieldtype.FieldSchemaInt64:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field schema_int64", values[i])
			} else if value.Valid {
				ft.SchemaInt64 = schema.Int64(value.Int64)
			}
		case fieldtype.FieldSchemaFloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field schema_float", values[i])
			} else if value.Valid {
				ft.SchemaFloat = schema.Float64(value.Float64)
			}
		case fieldtype.FieldSchemaFloat32:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field schema_float32", values[i])
			} else if value.Valid {
				ft.SchemaFloat32 = schema.Float32(value.Float64)
			}
		case fieldtype.FieldNullFloat:
			if value, ok := values[i].(*sql.NullFloat64); !ok {
				return fmt.Errorf("unexpected type %T for field null_float", values[i])
			} else if value.Valid {
				ft.NullFloat = value
			}
		case fieldtype.FieldRole:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field role", values[i])
			} else if value.Valid {
				ft.Role = role.Role(value.String)
			}
		case fieldtype.FieldPriority:
			if value, ok := values[i].(*role.Priority); !ok {
				return fmt.Errorf("unexpected type %T for field priority", values[i])
			} else if value != nil {
				ft.Priority = *value
			}
		case fieldtype.FieldOptionalUUID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field optional_uuid", values[i])
			} else if value != nil {
				ft.OptionalUUID = *value
			}
		case fieldtype.FieldNillableUUID:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field nillable_uuid", values[i])
			} else if value.Valid {
				ft.NillableUUID = new(uuid.UUID)
				*ft.NillableUUID = *value.S.(*uuid.UUID)
			}
		case fieldtype.FieldStrings:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field strings", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ft.Strings); err != nil {
					return fmt.Errorf("unmarshal field strings: %w", err)
				}
			}
		case fieldtype.FieldPair:
			if value, ok := values[i].(*schema.Pair); !ok {
				return fmt.Errorf("unexpected type %T for field pair", values[i])
			} else if value != nil {
				ft.Pair = *value
			}
		case fieldtype.FieldNilPair:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field nil_pair", values[i])
			} else if value.Valid {
				ft.NilPair = value.S.(*schema.Pair)
			}
		case fieldtype.FieldVstring:
			if value, ok := values[i].(*schema.VString); !ok {
				return fmt.Errorf("unexpected type %T for field vstring", values[i])
			} else if value != nil {
				ft.Vstring = *value
			}
		case fieldtype.FieldTriple:
			if value, ok := values[i].(*schema.Triple); !ok {
				return fmt.Errorf("unexpected type %T for field triple", values[i])
			} else if value != nil {
				ft.Triple = *value
			}
		case fieldtype.FieldBigInt:
			if value, ok := values[i].(*schema.BigInt); !ok {
				return fmt.Errorf("unexpected type %T for field big_int", values[i])
			} else if value != nil {
				ft.BigInt = *value
			}
		case fieldtype.FieldPasswordOther:
			if value, ok := values[i].(*schema.Password); !ok {
				return fmt.Errorf("unexpected type %T for field password_other", values[i])
			} else if value != nil {
				ft.PasswordOther = *value
			}
		case fieldtype.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field file_field", value)
			} else if value.Valid {
				ft.file_field = new(int)
				*ft.file_field = int(value.Int64)
			}
		default:
			ft.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FieldType.
// This includes values selected through modifiers, order, etc.
func (ft *FieldType) Value(name string) (ent.Value, error) {
	return ft.selectValues.Get(name)
}

// Update returns a builder for updating this FieldType.
// Note that you need to call FieldType.Unwrap() before calling this method if this FieldType
// was returned from a transaction, and the transaction was committed or rolled back.
func (ft *FieldType) Update() *FieldTypeUpdateOne {
	return NewFieldTypeClient(ft.config).UpdateOne(ft)
}

// Unwrap unwraps the FieldType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ft *FieldType) Unwrap() *FieldType {
	_tx, ok := ft.config.driver.(*txDriver)
	if !ok {
		panic("ent: FieldType is not a transactional entity")
	}
	ft.config.driver = _tx.drv
	return ft
}

// String implements the fmt.Stringer.
func (ft *FieldType) String() string {
	var builder strings.Builder
	builder.WriteString("FieldType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ft.ID))
	builder.WriteString("int=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int))
	builder.WriteString(", ")
	builder.WriteString("int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int8))
	builder.WriteString(", ")
	builder.WriteString("int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int16))
	builder.WriteString(", ")
	builder.WriteString("int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int32))
	builder.WriteString(", ")
	builder.WriteString("int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.Int64))
	builder.WriteString(", ")
	builder.WriteString("optional_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt))
	builder.WriteString(", ")
	builder.WriteString("optional_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt8))
	builder.WriteString(", ")
	builder.WriteString("optional_int16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt16))
	builder.WriteString(", ")
	builder.WriteString("optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt32))
	builder.WriteString(", ")
	builder.WriteString("optional_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalInt64))
	builder.WriteString(", ")
	if v := ft.NillableInt; v != nil {
		builder.WriteString("nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ft.NillableInt8; v != nil {
		builder.WriteString("nillable_int8=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ft.NillableInt16; v != nil {
		builder.WriteString("nillable_int16=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ft.NillableInt32; v != nil {
		builder.WriteString("nillable_int32=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ft.NillableInt64; v != nil {
		builder.WriteString("nillable_int64=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("validate_optional_int32=")
	builder.WriteString(fmt.Sprintf("%v", ft.ValidateOptionalInt32))
	builder.WriteString(", ")
	builder.WriteString("optional_uint=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint))
	builder.WriteString(", ")
	builder.WriteString("optional_uint8=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint8))
	builder.WriteString(", ")
	builder.WriteString("optional_uint16=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint16))
	builder.WriteString(", ")
	builder.WriteString("optional_uint32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint32))
	builder.WriteString(", ")
	builder.WriteString("optional_uint64=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUint64))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", ft.State))
	builder.WriteString(", ")
	builder.WriteString("optional_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat))
	builder.WriteString(", ")
	builder.WriteString("optional_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalFloat32))
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(ft.Text)
	builder.WriteString(", ")
	builder.WriteString("datetime=")
	builder.WriteString(ft.Datetime.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("decimal=")
	builder.WriteString(fmt.Sprintf("%v", ft.Decimal))
	builder.WriteString(", ")
	builder.WriteString("link_other=")
	builder.WriteString(fmt.Sprintf("%v", ft.LinkOther))
	builder.WriteString(", ")
	builder.WriteString("link_other_func=")
	builder.WriteString(fmt.Sprintf("%v", ft.LinkOtherFunc))
	builder.WriteString(", ")
	builder.WriteString("mac=")
	builder.WriteString(fmt.Sprintf("%v", ft.MAC))
	builder.WriteString(", ")
	builder.WriteString("string_array=")
	builder.WriteString(fmt.Sprintf("%v", ft.StringArray))
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	if v := ft.StringScanner; v != nil {
		builder.WriteString("string_scanner=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", ft.Duration))
	builder.WriteString(", ")
	builder.WriteString("dir=")
	builder.WriteString(fmt.Sprintf("%v", ft.Dir))
	builder.WriteString(", ")
	if v := ft.Ndir; v != nil {
		builder.WriteString("ndir=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("str=")
	builder.WriteString(fmt.Sprintf("%v", ft.Str))
	builder.WriteString(", ")
	if v := ft.NullStr; v != nil {
		builder.WriteString("null_str=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("link=")
	builder.WriteString(fmt.Sprintf("%v", ft.Link))
	builder.WriteString(", ")
	if v := ft.NullLink; v != nil {
		builder.WriteString("null_link=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", ft.Active))
	builder.WriteString(", ")
	if v := ft.NullActive; v != nil {
		builder.WriteString("null_active=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	if v := ft.Deleted; v != nil {
		builder.WriteString("deleted=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("deleted_at=")
	builder.WriteString(fmt.Sprintf("%v", ft.DeletedAt))
	builder.WriteString(", ")
	builder.WriteString("raw_data=")
	builder.WriteString(fmt.Sprintf("%v", ft.RawData))
	builder.WriteString(", ")
	builder.WriteString("sensitive=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("ip=")
	builder.WriteString(fmt.Sprintf("%v", ft.IP))
	builder.WriteString(", ")
	builder.WriteString("null_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullInt64))
	builder.WriteString(", ")
	builder.WriteString("schema_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt))
	builder.WriteString(", ")
	builder.WriteString("schema_int8=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt8))
	builder.WriteString(", ")
	builder.WriteString("schema_int64=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaInt64))
	builder.WriteString(", ")
	builder.WriteString("schema_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat))
	builder.WriteString(", ")
	builder.WriteString("schema_float32=")
	builder.WriteString(fmt.Sprintf("%v", ft.SchemaFloat32))
	builder.WriteString(", ")
	builder.WriteString("null_float=")
	builder.WriteString(fmt.Sprintf("%v", ft.NullFloat))
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", ft.Role))
	builder.WriteString(", ")
	builder.WriteString("priority=")
	builder.WriteString(fmt.Sprintf("%v", ft.Priority))
	builder.WriteString(", ")
	builder.WriteString("optional_uuid=")
	builder.WriteString(fmt.Sprintf("%v", ft.OptionalUUID))
	builder.WriteString(", ")
	if v := ft.NillableUUID; v != nil {
		builder.WriteString("nillable_uuid=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("strings=")
	builder.WriteString(fmt.Sprintf("%v", ft.Strings))
	builder.WriteString(", ")
	builder.WriteString("pair=")
	builder.WriteString(fmt.Sprintf("%v", ft.Pair))
	builder.WriteString(", ")
	if v := ft.NilPair; v != nil {
		builder.WriteString("nil_pair=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteString(", ")
	builder.WriteString("vstring=")
	builder.WriteString(fmt.Sprintf("%v", ft.Vstring))
	builder.WriteString(", ")
	builder.WriteString("triple=")
	builder.WriteString(fmt.Sprintf("%v", ft.Triple))
	builder.WriteString(", ")
	builder.WriteString("big_int=")
	builder.WriteString(fmt.Sprintf("%v", ft.BigInt))
	builder.WriteString(", ")
	builder.WriteString("password_other=<sensitive>")
	builder.WriteByte(')')
	return builder.String()
}

// FieldTypes is a parsable slice of FieldType.
type FieldTypes []*FieldType

// Len returns length of FieldTypes.
func (ft FieldTypes) Len() int { return len(ft) }
