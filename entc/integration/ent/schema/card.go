// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

type TimeMixin struct{}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Immutable(),
	}
}

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
	}
}

// Fields of the Comment.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			NotEmpty(),
	}
}

// Edges of the Comment.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Comment("O2O inverse edge").
			Ref("card").
			Unique(),
	}
}
