// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"database/sql/driver"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/entc/integration/ent/role"
	"github.com/facebook/ent/schema/field"

	"github.com/google/uuid"
)

// FieldType holds the schema definition for the FieldType entity.
// used for testing field types.
type FieldType struct {
	ent.Schema
}

// Fields of the File.
func (FieldType) Fields() []ent.Field {
	return []ent.Field{
		field.Int("int"),
		field.Int8("int8"),
		field.Int16("int16"),
		field.Int32("int32"),
		field.Int64("int64"),
		field.Int("optional_int").Optional(),
		field.Int8("optional_int8").Optional(),
		field.Int16("optional_int16").Optional(),
		field.Int32("optional_int32").Optional(),
		field.Int64("optional_int64").Optional(),
		field.Int("nillable_int").Optional().Nillable(),
		field.Int8("nillable_int8").Optional().Nillable(),
		field.Int16("nillable_int16").Optional().Nillable(),
		field.Int32("nillable_int32").Optional().Nillable(),
		field.Int64("nillable_int64").Optional().Nillable(),
		field.Int32("validate_optional_int32").
			Optional().
			Max(100),
		field.Uint("optional_uint").Optional(),
		field.Uint8("optional_uint8").Optional(),
		field.Uint16("optional_uint16").Optional(),
		field.Uint32("optional_uint32").Optional(),
		field.Uint64("optional_uint64").Optional(),
		field.Enum("state").
			Values("on", "off").
			Optional(),
		field.Float("optional_float").Optional(),
		field.Float32("optional_float32").Optional(),
		field.Time("datetime").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "datetime",
				dialect.Postgres: "date",
			}),
		field.Float("decimal").
			Optional().
			SchemaType(map[string]string{
				dialect.MySQL:    "decimal(6,2)",
				dialect.Postgres: "numeric",
			}),
		field.String("dir").
			Optional().
			GoType(http.Dir("dir")),
		field.String("ndir").
			Optional().
			Nillable().
			NotEmpty().
			GoType(http.Dir("ndir")),
		field.String("str").
			Optional().
			GoType(&sql.NullString{}),
		field.String("null_str").
			Optional().
			Nillable().
			GoType(&sql.NullString{}),
		field.String("link").
			Optional().
			NotEmpty().
			GoType(&Link{}),
		field.String("null_link").
			Optional().
			Nillable().
			GoType(&Link{}),
		field.Bool("active").
			Optional().
			GoType(Status(false)),
		field.Bool("null_active").
			Optional().
			Nillable().
			GoType(Status(false)),
		field.Bool("deleted").
			Optional().
			GoType(&sql.NullBool{}),
		field.Time("deleted_at").
			Optional().
			GoType(&sql.NullTime{}),
		field.Bytes("ip").
			Optional().
			GoType(net.IP("127.0.0.1")),
		field.Int("null_int64").
			Optional().
			GoType(&sql.NullInt64{}),
		field.Int("schema_int").
			Optional().
			GoType(Int(0)),
		field.Int8("schema_int8").
			Optional().
			GoType(Int8(0)),
		field.Int64("schema_int64").
			Optional().
			GoType(Int64(0)),
		field.Float("schema_float").
			Optional().
			GoType(Float64(0)),
		field.Float32("schema_float32").
			Optional().
			GoType(Float32(0)),
		field.Float("null_float").
			Optional().
			GoType(&sql.NullFloat64{}),
		field.Enum("role").
			Default(string(role.Read)).
			GoType(role.Role("role")),
		field.String("mac").
			Optional().
			GoType(&MAC{}).
			SchemaType(map[string]string{
				dialect.Postgres: "macaddr",
			}).
			Validate(func(s string) error {
				_, err := net.ParseMAC(s)
				return err
			}),
		field.UUID("uuid", uuid.UUID{}).
			Optional(),
	}
}

type (
	Int     int
	Int8    int8
	Int64   int64
	Status  bool
	Float64 float64
	Float32 float32
)

type Link struct {
	*url.URL
}

// Scan implements the Scanner interface.
func (l *Link) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		l.URL, err = url.Parse(string(v))
	case string:
		l.URL, err = url.Parse(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (l Link) Value() (driver.Value, error) {
	if l.URL == nil {
		return nil, nil
	}
	return l.String(), nil
}

type MAC struct {
	net.HardwareAddr
}

// Scan implements the Scanner interface.
func (m *MAC) Scan(value interface{}) (err error) {
	switch v := value.(type) {
	case nil:
	case []byte:
		m.HardwareAddr, err = net.ParseMAC(string(v))
	case string:
		m.HardwareAddr, err = net.ParseMAC(v)
	default:
		err = fmt.Errorf("unexpected type %T", v)
	}
	return
}

// Value implements the driver Valuer interface.
func (m MAC) Value() (driver.Value, error) {
	return m.HardwareAddr.String(), nil
}
