// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Node holds the schema definition for the linked-list Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").
			Optional(),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Node.Type).
			StructTag("gqlgen:next").
			Unique().
			From("prev").
			StructTag("gqlgen:prev").
			Unique(),
	}
}
