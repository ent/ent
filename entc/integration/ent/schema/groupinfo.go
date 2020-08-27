// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// GroupInfo holds the schema for the group-info entity.
type GroupInfo struct {
	ent.Schema
}

// Fields of the group.
func (GroupInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
		field.Int("max_users").
			Default(1e4),
	}
}

// Edges of the group.
func (GroupInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("info"),
	}
}
