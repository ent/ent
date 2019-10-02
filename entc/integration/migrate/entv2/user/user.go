// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/schema"
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
	// FieldNewName holds the string denoting the new_name vertex property in the database.
	FieldNewName = "renamed"
	// FieldBlob holds the string denoting the blob vertex property in the database.
	FieldBlob = "blob"
	// FieldState holds the string denoting the state vertex property in the database.
	FieldState = "state"

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
	FieldNewName,
	FieldBlob,
	FieldState,
}

var (
	fields = schema.User{}.Fields()

	// descBuffer is the schema descriptor for buffer field.
	descBuffer = fields[3].Descriptor()
	// DefaultBuffer holds the default value on creation for the buffer field.
	DefaultBuffer = descBuffer.Default.([]byte)

	// descTitle is the schema descriptor for title field.
	descTitle = fields[4].Descriptor()
	// DefaultTitle holds the default value on creation for the title field.
	DefaultTitle = descTitle.Default.(string)
)

// State defines the type for the state enum field.
type State string

const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
	StateOnline    State = "online"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "state" field enum values. It is called by the builders before save.
func StateValidator(state State) error {
	switch state {
	case StateLoggedIn, StateLoggedOut, StateOnline:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", state)
	}
}
