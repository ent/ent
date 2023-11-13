// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// An example for an ent/schema that uses struct embedding to define default table schema (in the database).

// Group holds the schema definition for the Group entity.
type Group struct {
	base
}

// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("unknown"),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		// The edge schema (join table) is defaults to the edge owner schema (Group).
		edge.To("users", User.Type),
	}
}
