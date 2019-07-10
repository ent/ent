package schema

import (
	"fbc/ent"
	"fbc/ent/field"
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
		field.Int("nullable_int").Optional().Nullable(),
		field.Int8("nullable_int8").Optional().Nullable(),
		field.Int16("nullable_int16").Optional().Nullable(),
		field.Int32("nullable_int32").Optional().Nullable(),
		field.Int64("nullable_int64").Optional().Nullable(),
	}
}
