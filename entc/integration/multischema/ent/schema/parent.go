// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Parent holds the schema definition for the Parent entity.
type Parent struct {
	base
}

// Fields of the Parent.
func (Parent) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("by_adoption").
			Default(false),
		field.Int("user_id").
			Immutable(),
		field.Int("parent_id").
			Immutable(),
	}
}

// Edges of the Parent.
func (Parent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("child", User.Type).
			Unique().
			Required().
			Immutable().
			Field("user_id"),
		edge.To("parent", User.Type).
			Unique().
			Required().
			Immutable().
			Field("parent_id"),
	}
}
