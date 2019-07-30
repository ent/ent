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

	// Table holds the table name of the file in the database.
	Table = "files"
)

// Columns holds all SQL columns are file fields.
var Columns = []string{
	FieldID,
	FieldSize,
	FieldName,
}

var (
	fields = schema.File{}.Fields()
	// SizeValidator is a validator for the "size" field. It is called by the builders before save.
	SizeValidator = fields[0].Validators()[0].(func(int) error)
)
