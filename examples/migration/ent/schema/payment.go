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
	"entgo.io/ent/schema/index"
)

// Payment holds the schema definition for the Payment entity.
type Payment struct {
	ent.Schema
}

// Fields of the Payment.
func (Payment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("card_id"),
		field.Float("amount").
			Positive(),
		field.Enum("currency").
			Values("USD", "EUR", "VND", "ILS"),
		field.Time("time"),
		field.String("description"),
		field.Enum("status").
			Values("pending", "completed", "failed"),
	}
}

// Edges of the Payment.
func (Payment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("card", Card.Type).
			Ref("payments").
			Unique().
			Required().
			Field("card_id"),
	}
}

// Indexes of the Payment.
func (Payment) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("status", "time"),
	}
}

// Annotations of the Payment.
func (Payment) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// Named check constraints are compared by their name.
		// Thus, the definition does not need to be normalized.
		entsql.Checks(map[string]string{
			"amount_positive": "(`amount` > 0)",
		}),
	}
}
