// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Mixin of the Pet.
func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional(),
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
