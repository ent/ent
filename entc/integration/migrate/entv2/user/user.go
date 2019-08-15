// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"fbc/ent/entc/integration/migrate/entv2/schema"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"
	// FieldPhone holds the string denoting the phone vertex property in the database.
	FieldPhone = "phone"
	// FieldBuffer holds the string denoting the buffer vertex property in the database.
	FieldBuffer = "buffer"
	// FieldTitle holds the string denoting the title vertex property in the database.
	FieldTitle = "title"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldPhone,
	FieldBuffer,
	FieldTitle,
}

var (
	fields = schema.User{}.Fields()
	// DefaultBuffer holds the default value for the buffer field.
	DefaultBuffer = fields[3].Value().([]byte)
	// DefaultTitle holds the default value for the title field.
	DefaultTitle = fields[4].Value().(string)
)
