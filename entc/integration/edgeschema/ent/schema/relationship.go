// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/edgeschema/ent/privacy"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Relationship holds the schema definition for the Relationship entity.
type Relationship struct {
	ent.Schema
}

func (Relationship) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "relative_id"),
	}
}

// Fields of the Relationship.
func (Relationship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weight").
			Default(1),

		// Edge fields for the fields that compose this edge.
		// They also function as the primary key of the table.
		field.Int("user_id"),
		field.Int("relative_id"),

		// An edge field to external node that holds
		// additional information about this edge.
		field.Int("info_id").
			Optional(),
	}
}

// Edges of the Relationship.
func (Relationship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Required().
			Unique().
			Field("user_id"),
		edge.To("relative", User.Type).
			Required().
			Unique().
			Field("relative_id"),

		// An optional edge to an entity that holds
		// information about this relationship.
		edge.To("info", RelationshipInfo.Type).
			Field("info_id").
			Unique(),
	}
}

// Indexes of the Relationship.
func (Relationship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("weight"),

		// A relationship-info can be connected to no more
		// than one relationship object (and edge schema).
		index.Edges("info").
			Unique(),
	}
}

// Policy defines the privacy policy of the Relationship.
func (Relationship) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
