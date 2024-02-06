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

// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}

// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number_hash"),
		field.String("cvv_hash"),
		field.Time("expires_at").
			Optional(),
		field.Int("owner_id").
			Default(0),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("cards").
			Unique().
			Required().
			Field("owner_id"),
		edge.To("payments", Payment.Type),
	}
}

// Annotations of the Card.
func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		entsql.Checks(map[string]string{
			"number_hash_length": "(LENGTH(`number_hash`) = 16)",
		}),
	}
}
