// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/customid/sid"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// IntSid holds the schema definition for the IntSid entity.
type IntSid struct {
	ent.Schema
}

// Fields of the IntSid.
func (IntSid) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("id").
			GoType(sid.New()).
			Unique().
			Immutable(),
	}
}

// Edges of the IntSid.
func (IntSid) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", IntSid.Type).
			Unique(),
		edge.From("children", IntSid.Type).Ref("parent"),
	}
}
