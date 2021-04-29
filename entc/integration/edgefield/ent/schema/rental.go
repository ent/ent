// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Rental holds the schema definition for the Rental entity.
type Rental struct {
	ent.Schema
}

// Fields of the Rental.
func (Rental) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date").
			Default(time.Now),
		field.Int("car_id"),
		field.Int("user_id"),
	}
}

// Edges of the Rental.
func (Rental) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("rentals").
			Field("user_id").
			Unique().
			Required(),
		edge.From("car", Car.Type).
			Ref("rentals").
			Field("car_id").
			Unique().
			Required(),
	}
}

// Indexes of the Rental.
func (Rental) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("car_id", "user_id").
			Unique(),
	}
}
