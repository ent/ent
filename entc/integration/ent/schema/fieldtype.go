// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"net/http"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/dialect"
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
	}
}
