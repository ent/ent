// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package role

import (
	time "time"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeRolesUsers holds the string denoting the roles_users edge name in mutations.
	EdgeRolesUsers = "roles_users"
	// Table holds the table name of the role in the database.
	Table = "roles"
	// UserTable is the table that holds the user relation/edge. The primary key declared below.
	UserTable = "role_users"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// RolesUsersTable is the table that holds the roles_users relation/edge.
	RolesUsersTable = "role_users"
	// RolesUsersInverseTable is the table name for the RoleUser entity.
	// It exists in this package in order to avoid circular dependency with the "roleuser" package.
	RolesUsersInverseTable = "role_users"
	// RolesUsersColumn is the table column denoting the roles_users relation/edge.
	RolesUsersColumn = "role_id"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldCreatedAt,
}

var (
	// UserPrimaryKey and UserColumn2 are the table columns denoting the
	// primary key for the user relation (M2M).
	UserPrimaryKey = []string{"user_id", "role_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
)
