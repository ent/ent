// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/entc/integration/ent/template"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"card_edges" mashraki:"edges"`,
		},
		field.Annotation{
			StructTag: map[string]string{
				"id":     `json:"-"`,
				"number": `json:"-"`,
			},
		},
	}
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

// Fields of the Comment.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			Immutable().
			NotEmpty().
			Annotations(&template.Extension{
				Type: "string",
			}),
		field.String("name").
			Optional().
			Comment("Exact name written on card").
			NotEmpty().
			Annotations(&template.Extension{
				Type: "string",
			}),
	}
}

// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Comment("O2O inverse edge").
			Ref("card").
			Unique(),
		edge.From("spec", Spec.Type).
			Ref("card").
			Annotations(&template.Extension{
				Type: "int",
			}),
	}
}
