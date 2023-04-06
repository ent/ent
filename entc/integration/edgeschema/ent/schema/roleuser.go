// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// RoleUser holds the schema definition for the RoleUser entity.
type RoleUser struct {
	ent.Schema
}

func (RoleUser) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "role_id"),
	}
}

// Fields of the RoleUser.
func (RoleUser) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		// Edge fields.
		field.Int("role_id"),
		field.Int("user_id"),
	}
}

// Edges of the RoleUser.
func (RoleUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("role", Role.Type).
			Unique().
			Required().
			Field("role_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}
