// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"encoding/json"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Info holds the schema definition for the Info entity.
type Info struct {
	ent.Schema
}

func (Info) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.JSON("content", json.RawMessage{}),
	}
}

func (Info) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			StorageKey(edge.Column("id")),
	}
}
