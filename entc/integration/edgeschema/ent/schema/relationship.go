// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
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
		field.Int("user_id"),
		field.Int("relative_id"),
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
	}
}

// Indexes of the Relationship.
func (Relationship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("weight"),
	}
}
