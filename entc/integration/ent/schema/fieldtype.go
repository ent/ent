package schema

import (
	"github.com/facebookincubator/ent"
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
	}
}
