// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.Nil).
			Default(uuid.New),
		field.UUID("best_friend_id", uuid.Nil).
			Annotations(
				entsql.Default(uuid.Nil.String()),
			),
		field.Int("owner_id").
			Default(0),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("best_friend", Pet.Type).
			Unique().
			Required().
			Field("best_friend_id"),
		edge.To("owner", User.Type).
			Unique().
			Required().
			Field("owner_id"),
	}
}
