// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Group holds the schema for the group entity.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("active").
			Default(true),
		field.String("name"),
	}
}

// Edges of the group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("blocked", User.Type),
		edge.From("users", User.Type).Ref("groups"),
	}
}
