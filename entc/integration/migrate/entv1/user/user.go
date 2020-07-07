// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"fmt"
)

const (
	// Label holds the string label denoting the user type in the database.
	Label = "user"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "oid"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldRenamed holds the string denoting the renamed field in the database.
	FieldRenamed = "renamed"
	// FieldBlob holds the string denoting the blob field in the database.
	FieldBlob = "blob"
	// FieldState holds the string denoting the state field in the database.
	FieldState = "state"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"

	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeSpouse holds the string denoting the spouse edge name in mutations.
	EdgeSpouse = "spouse"
	// EdgeCar holds the string denoting the car edge name in mutations.
	EdgeCar = "car"

	// Table holds the table name of the user in the database.
	Table = "users"
	// ParentTable is the table the holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "user_children"
	// ChildrenTable is the table the holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "user_children"
	// SpouseTable is the table the holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "user_spouse"
	// CarTable is the table the holds the car relation/edge.
	CarTable = "cars"
	// CarInverseTable is the table name for the Car entity.
	// It exists in this package in order to avoid circular dependency with the "car" package.
	CarInverseTable = "cars"
	// CarColumn is the table column denoting the car relation/edge.
	CarColumn = "user_car"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldAge,
	FieldName,
	FieldNickname,
	FieldAddress,
	FieldRenamed,
	FieldBlob,
	FieldState,
	FieldStatus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the User type.
var ForeignKeys = []string{
	"user_children",
	"user_spouse",
}

var (
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// State defines the type for the state enum field.
type State string

// State values.
const (
	StateLoggedIn  State = "logged_in"
	StateLoggedOut State = "logged_out"
)

func (s State) String() string {
	return string(s)
}

// StateValidator is a validator for the "s" field enum values. It is called by the builders before save.
func StateValidator(s State) error {
	switch s {
	case StateLoggedIn, StateLoggedOut:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for state field: %q", s)
	}
}
