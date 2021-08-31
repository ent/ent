// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Metadata holds the schema definition for the Metadata entity.
type Metadata struct {
	ent.Schema
}

func (Metadata) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.Int("age").
			Default(0),
		field.Int("parent_id").
			Optional(),
	}
}
func (Metadata) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("metadata").
			Unique(),
		edge.To("parent", Metadata.Type).
			Field("parent_id").
			Unique().
			From("children"),
	}
}
