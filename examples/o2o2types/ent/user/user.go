// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldAge holds the string denoting the age vertex property in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name vertex property in the database.
	FieldName = "name"

	// Table holds the table name of the user in the database.
	Table = "users"
	// CardTable is the table the holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "owner_id"
)

// Columns holds all SQL columns are user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
}
