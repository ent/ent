// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
	}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("metadata", Metadata.Type).
			Unique().
			StorageKey(edge.Column("id")),
		edge.From("info", Info.Type).
			Ref("user"),
	}
}
