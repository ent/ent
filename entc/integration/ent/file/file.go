// Code generated (@generated) by entc, DO NOT EDIT.

package file

import (
	"fbc/ent/entc/integration/ent/schema"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSize holds the string denoting the size vertex property in the database.
	FieldSize = "size"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldUser holds the string denoting the user vertex property in the database.
	FieldUser = "user"
	// FieldGroup holds the string denoting the group vertex property in the database.
	FieldGroup = "group"

	// Table holds the table name of the file in the database.
	Table = "files"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "files"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "owner_id"

	// OwnerInverseLabel holds the string label denoting the owner inverse edge type in the database.
	OwnerInverseLabel = "user_files"
)

// Columns holds all SQL columns are file fields.
var Columns = []string{
	FieldID,
	FieldSize,
	FieldName,
	FieldUser,
	FieldGroup,
}

var (
	fields = schema.File{}.Fields()
	// DefaultSize holds the default value for the size field.
	DefaultSize = fields[0].Value().(int)
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator = fields[0].Validators()[0].(func(int) error)
)
