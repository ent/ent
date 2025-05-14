// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	base
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
		edge.From("groups", Group.Type).
			Ref("users"),
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type),
		edge.To("children", User.Type).
			Through("parent_hood", Parent.Type).
			From("parents"),
	}
}

// CleanUser represents a user without its PII field.
type CleanUser struct {
	ent.View
}

// Annotations of the CleanUser.
func (CleanUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.View("SELECT id, name FROM users"),
		entsql.Schema("db1"),
	}
}

// Fields of the CleanUser.
func (CleanUser) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}
