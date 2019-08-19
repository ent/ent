// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/schema"
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
	// FieldAddress holds the string denoting the address vertex property in the database.
	FieldAddress = "address"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldAddress,
}

var (
	fields = schema.User{}.Fields()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = fields[1].Validators()[0].(func(string) error)
)
