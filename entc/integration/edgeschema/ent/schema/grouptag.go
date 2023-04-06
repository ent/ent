// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// GroupTag holds the schema definition for the GroupTag entity.
type GroupTag struct {
	ent.Schema
}

// Fields of the GroupTag.
func (GroupTag) Fields() []ent.Field {
	return []ent.Field{
		// An edge schema with the builtin ID
		// field, but without any other field.
		field.Int("tag_id"),
		field.Int("group_id"),
	}
}

// Edges of the GroupTag.
func (GroupTag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tag", Tag.Type).
			Unique().
			Required().
			Field("tag_id"),
		edge.To("group", Group.Type).
			Unique().
			Required().
			Field("group_id"),
	}
}
