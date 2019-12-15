// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/entc/integration/ent/schema"
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
	// FieldLast holds the string denoting the last vertex property in the database.
	FieldLast = "last"
	// FieldNickname holds the string denoting the nickname vertex property in the database.
	FieldNickname = "nickname"
	// FieldPhone holds the string denoting the phone vertex property in the database.
	FieldPhone = "phone"
	// FieldPassword holds the string denoting the password vertex property in the database.
	FieldPassword = "password"

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
	fields = schema.User{}.Fields()

	// descLast is the schema descriptor for last field.
	descLast = fields[2].Descriptor()
	// DefaultLast holds the default value on creation for the last field.
	DefaultLast = descLast.Default.(string)
)
