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

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/schema/field"
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
	}
}

type Status bool

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
		err = fmt.Errorf("unexpcted type %T", v)
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
