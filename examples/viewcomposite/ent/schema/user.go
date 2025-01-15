// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("public_info"),
		field.String("private_info"),
	}
}

// CleanUser represents a user without its PII field.
type CleanUser struct {
	ent.View
}

// Fields of the CleanUser.
func (CleanUser) Fields() []ent.Field {
	return []ent.Field{
		// Note, unlike real schemas (tables, defined with ent.Schema),
		// the "id" field should be defined manually if needed.
		field.Int("id"),
		field.String("name"),
		field.String("public_info"),
	}
}
