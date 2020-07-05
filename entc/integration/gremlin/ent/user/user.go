// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
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

	// CardLabel holds the string label denoting the card edge type in the database.
	CardLabel = "user_card"
	// PetsLabel holds the string label denoting the pets edge type in the database.
	PetsLabel = "user_pets"
	// FilesLabel holds the string label denoting the files edge type in the database.
	FilesLabel = "user_files"
	// GroupsLabel holds the string label denoting the groups edge type in the database.
	GroupsLabel = "user_groups"
	// FriendsLabel holds the string label denoting the friends edge type in the database.
	FriendsLabel = "user_friends"
	// FollowersInverseLabel holds the string label denoting the followers inverse edge type in the database.
	FollowersInverseLabel = "user_following"
	// FollowingLabel holds the string label denoting the following edge type in the database.
	FollowingLabel = "user_following"
	// TeamLabel holds the string label denoting the team edge type in the database.
	TeamLabel = "user_team"
	// SpouseLabel holds the string label denoting the spouse edge type in the database.
	SpouseLabel = "user_spouse"
	// ChildrenInverseLabel holds the string label denoting the children inverse edge type in the database.
	ChildrenInverseLabel = "user_parent"
	// ParentLabel holds the string label denoting the parent edge type in the database.
	ParentLabel = "user_parent"
)

var (
	// OptionalIntValidator is a validator for the "optional_int" field. It is called by the builders before save.
	OptionalIntValidator func(int) error
	// DefaultLast holds the default value on creation for the last field.
	DefaultLast string
)

// Role defines the type for the role enum field.
type Role string

// RoleUser is the default Role.
const DefaultRole = RoleUser

// Role values.
const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

// RoleValidator is a validator for the "r" field enum values. It is called by the builders before save.
func RoleValidator(r Role) error {
	switch r {
	case RoleUser, RoleAdmin:
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
