// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

type IDMixin struct {
	mixin.Schema
}

func (IDMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Float("before_id").
			Optional().
			Positive(),
		field.Int("id").
			Positive().
			Immutable(),
		field.Float("after_id").
			Optional().
			Negative(),
	}
}

// Car holds the schema definition for the Car entity.
type Car struct {
	ent.Schema
}

// Mixin of the Car.
func (Car) Mixin() []ent.Mixin {
	return []ent.Mixin{
		IDMixin{},
	}
}

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
	}
}

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", Pet.Type).
			Ref("cars").
			Unique(),
	}
}
