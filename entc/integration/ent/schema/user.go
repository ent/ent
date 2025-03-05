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
	"entgo.io/ent/schema/mixin"
)

// User holds the schema for the user entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		UserMixin{},
	}
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name").
			StructTag(`json:"first_name" graphql:"first_name"`),
		field.String("last").
			Default("unknown").
			StructTag(`graphql:"last_name"`),
		field.String("nickname").
			Optional().
			Unique(),
		field.String("address").
			Optional().
			DefaultFunc(func() string { return "static" }),
		field.String("phone").
			Optional().
			Unique(),
		field.String("password").
			Optional().
			Sensitive(),
		field.Enum("role").
			Values("user", "admin", "free-user", "test user").
			Default("user"),
		field.Enum("employment").
			Values("Full-Time", "Part-Time", "Contract").
			Default("Full-Time"),
		field.String("SSOCert").
			Optional(),
		// Some users store the associations
		// count as a separate field.
		field.Int("files_count").
			Optional(),
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).Comment("Cards associated with this user. O2O edge").Unique(),
		edge.To("pets", Pet.Type),
		edge.To("files", File.Type),
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("team", Pet.Type).Unique(),
		edge.To("spouse", User.Type).Unique(),
		edge.To("parent", User.Type).Unique().From("children"),
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.IncrementStart(2 << 32),
	}
}

// UserMixin composes create/update time mixin.
type UserMixin struct {
	mixin.Schema
}

// Fields of the time mixin.
func (UserMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("optional_int").
			Optional().
			Positive(),
	}
}
