// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

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
	// FieldRenamed holds the string denoting the renamed vertex property in the database.
	FieldRenamed = "renamed"
	// FieldBlob holds the string denoting the blob vertex property in the database.
	FieldBlob = "blob"

	// Table holds the table name of the user in the database.
	Table = "users"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldAddress,
	FieldRenamed,
	FieldBlob,
}

var (
	fields = schema.User{}.Fields()

	// descName is the schema descriptor for name field.
	descName = fields[1].Descriptor()
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator = descName.Validators[0].(func(string) error)
)
