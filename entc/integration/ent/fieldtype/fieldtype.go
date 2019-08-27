// Code generated (@generated) by entc, DO NOT EDIT.

package fieldtype

import (
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the fieldtype type in the database.
	Label = "field_type"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldInt holds the string denoting the int vertex property in the database.
	FieldInt = "int"
	// FieldInt8 holds the string denoting the int8 vertex property in the database.
	FieldInt8 = "int8"
	// FieldInt16 holds the string denoting the int16 vertex property in the database.
	FieldInt16 = "int16"
	// FieldInt32 holds the string denoting the int32 vertex property in the database.
	FieldInt32 = "int32"
	// FieldInt64 holds the string denoting the int64 vertex property in the database.
	FieldInt64 = "int64"
	// FieldOptionalInt holds the string denoting the optional_int vertex property in the database.
	FieldOptionalInt = "optional_int"
	// FieldOptionalInt8 holds the string denoting the optional_int8 vertex property in the database.
	FieldOptionalInt8 = "optional_int8"
	// FieldOptionalInt16 holds the string denoting the optional_int16 vertex property in the database.
	FieldOptionalInt16 = "optional_int16"
	// FieldOptionalInt32 holds the string denoting the optional_int32 vertex property in the database.
	FieldOptionalInt32 = "optional_int32"
	// FieldOptionalInt64 holds the string denoting the optional_int64 vertex property in the database.
	FieldOptionalInt64 = "optional_int64"
	// FieldNillableInt holds the string denoting the nillable_int vertex property in the database.
	FieldNillableInt = "nillable_int"
	// FieldNillableInt8 holds the string denoting the nillable_int8 vertex property in the database.
	FieldNillableInt8 = "nillable_int8"
	// FieldNillableInt16 holds the string denoting the nillable_int16 vertex property in the database.
	FieldNillableInt16 = "nillable_int16"
	// FieldNillableInt32 holds the string denoting the nillable_int32 vertex property in the database.
	FieldNillableInt32 = "nillable_int32"
	// FieldNillableInt64 holds the string denoting the nillable_int64 vertex property in the database.
	FieldNillableInt64 = "nillable_int64"
	// FieldValidateOptionalInt32 holds the string denoting the validate_optional_int32 vertex property in the database.
	FieldValidateOptionalInt32 = "validate_optional_int32"

	// Table holds the table name of the fieldtype in the database.
	Table = "field_types"
)

// Columns holds all SQL columns are fieldtype fields.
var Columns = []string{
	FieldID,
	FieldInt,
	FieldInt8,
	FieldInt16,
	FieldInt32,
	FieldInt64,
	FieldOptionalInt,
	FieldOptionalInt8,
	FieldOptionalInt16,
	FieldOptionalInt32,
	FieldOptionalInt64,
	FieldNillableInt,
	FieldNillableInt8,
	FieldNillableInt16,
	FieldNillableInt32,
	FieldNillableInt64,
	FieldValidateOptionalInt32,
}

var (
	fields = schema.FieldType{}.Fields()
	// descValidateOptionalInt32 is the schema descriptor for validate_optional_int32 field.
	descValidateOptionalInt32 = fields[15].Descriptor()
	// ValidateOptionalInt32Validator is a validator for the "validate_optional_int32" field. It is called by the builders before save.
	ValidateOptionalInt32Validator = descValidateOptionalInt32.Validators[0].(func(int32) error)
)
