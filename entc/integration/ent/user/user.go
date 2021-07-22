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
	FieldID = "id"
	// FieldOptionalInt holds the string denoting the optional_int field in the database.
	FieldOptionalInt = "optional_int"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldLast holds the string denoting the last field in the database.
	FieldLast = "last"
	// FieldNickname holds the string denoting the nickname field in the database.
	FieldNickname = "nickname"
	// FieldAddress holds the string denoting the address field in the database.
	FieldAddress = "address"
	// FieldPhone holds the string denoting the phone field in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password field in the database.
	FieldPassword = "password"
	// FieldRole holds the string denoting the role field in the database.
	FieldRole = "role"
	// FieldSSOCert holds the string denoting the ssocert field in the database.
	FieldSSOCert = "sso_cert"
	// EdgeCard holds the string denoting the card edge name in mutations.
	EdgeCard = "card"
	// EdgePets holds the string denoting the pets edge name in mutations.
	EdgePets = "pets"
	// EdgeFiles holds the string denoting the files edge name in mutations.
	EdgeFiles = "files"
	// EdgeGroups holds the string denoting the groups edge name in mutations.
	EdgeGroups = "groups"
	// EdgeFriends holds the string denoting the friends edge name in mutations.
	EdgeFriends = "friends"
	// EdgeFollowers holds the string denoting the followers edge name in mutations.
	EdgeFollowers = "followers"
	// EdgeFollowing holds the string denoting the following edge name in mutations.
	EdgeFollowing = "following"
	// EdgeTeam holds the string denoting the team edge name in mutations.
	EdgeTeam = "team"
	// EdgeSpouse holds the string denoting the spouse edge name in mutations.
	EdgeSpouse = "spouse"
	// EdgeChildren holds the string denoting the children edge name in mutations.
	EdgeChildren = "children"
	// EdgeParent holds the string denoting the parent edge name in mutations.
	EdgeParent = "parent"
	// Table holds the table name of the user in the database.
	Table = "users"
	// CardTable is the table that holds the card relation/edge.
	CardTable = "cards"
	// CardInverseTable is the table name for the Card entity.
	// It exists in this package in order to avoid circular dependency with the "card" package.
	CardInverseTable = "cards"
	// CardColumn is the table column denoting the card relation/edge.
	CardColumn = "user_card"
	// PetsTable is the table that holds the pets relation/edge.
	PetsTable = "pet"
	// PetsInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	PetsInverseTable = "pet"
	// PetsColumn is the table column denoting the pets relation/edge.
	PetsColumn = "user_pets"
	// FilesTable is the table that holds the files relation/edge.
	FilesTable = "files"
	// FilesInverseTable is the table name for the File entity.
	// It exists in this package in order to avoid circular dependency with the "file" package.
	FilesInverseTable = "files"
	// FilesColumn is the table column denoting the files relation/edge.
	FilesColumn = "user_files"
	// GroupsTable is the table that holds the groups relation/edge. The primary key declared below.
	GroupsTable = "user_groups"
	// GroupsInverseTable is the table name for the Group entity.
	// It exists in this package in order to avoid circular dependency with the "group" package.
	GroupsInverseTable = "groups"
	// FriendsTable is the table that holds the friends relation/edge. The primary key declared below.
	FriendsTable = "user_friends"
	// FollowersTable is the table that holds the followers relation/edge. The primary key declared below.
	FollowersTable = "user_following"
	// FollowingTable is the table that holds the following relation/edge. The primary key declared below.
	FollowingTable = "user_following"
	// TeamTable is the table that holds the team relation/edge.
	TeamTable = "pet"
	// TeamInverseTable is the table name for the Pet entity.
	// It exists in this package in order to avoid circular dependency with the "pet" package.
	TeamInverseTable = "pet"
	// TeamColumn is the table column denoting the team relation/edge.
	TeamColumn = "user_team"
	// SpouseTable is the table that holds the spouse relation/edge.
	SpouseTable = "users"
	// SpouseColumn is the table column denoting the spouse relation/edge.
	SpouseColumn = "user_spouse"
	// ChildrenTable is the table that holds the children relation/edge.
	ChildrenTable = "users"
	// ChildrenColumn is the table column denoting the children relation/edge.
	ChildrenColumn = "user_parent"
	// ParentTable is the table that holds the parent relation/edge.
	ParentTable = "users"
	// ParentColumn is the table column denoting the parent relation/edge.
	ParentColumn = "user_parent"
)

// Columns holds all SQL columns for user fields.
var Columns = []string{
	FieldID,
	FieldOptionalInt,
	FieldAge,
	FieldName,
	FieldLast,
	FieldNickname,
	FieldAddress,
	FieldPhone,
	FieldPassword,
	FieldRole,
	FieldSSOCert,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "users"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"group_blocked",
	"user_spouse",
	"user_parent",
}

var (
	// GroupsPrimaryKey and GroupsColumn2 are the table columns denoting the
	// primary key for the groups relation (M2M).
	GroupsPrimaryKey = []string{"user_id", "group_id"}
	// FriendsPrimaryKey and FriendsColumn2 are the table columns denoting the
	// primary key for the friends relation (M2M).
	FriendsPrimaryKey = []string{"user_id", "friend_id"}
	// FollowersPrimaryKey and FollowersColumn2 are the table columns denoting the
	// primary key for the followers relation (M2M).
	FollowersPrimaryKey = []string{"user_id", "follower_id"}
	// FollowingPrimaryKey and FollowingColumn2 are the table columns denoting the
	// primary key for the following relation (M2M).
	FollowingPrimaryKey = []string{"user_id", "follower_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	OptionalIntValidator func(int) error
	// DefaultLast holds the default value on creation for the "last" field.
	DefaultLast string
	// DefaultAddress holds the default value on creation for the "address" field.
	DefaultAddress func() string
)

// Role defines the type for the "role" enum field.
type Role string

// RoleUser is the default value of the Role enum.
const DefaultRole = RoleUser

// Role values.
const (
	RoleUser     Role = "user"
	RoleAdmin    Role = "admin"
	RoleFreeUser Role = "free-user"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "role" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUser, RoleAdmin, RoleFreeUser:
		return nil
	default:
		return fmt.Errorf("user: invalid enum value for role field: %q", r)
	}
}

// Ptr returns a new pointer to the enum value.
func (r Role) Ptr() *Role {
	return &r
}

// comment from another template.
