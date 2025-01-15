// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"

	"github.com/google/uuid"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Annotations of the Pet.
func (Pet) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "pet"},
	}
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Float("age").
			Default(0),
		field.String("name"),
		field.UUID("uuid", uuid.UUID{}).
			Optional(),
		field.String("nickname").
			Optional(),
		field.Bool("trained").
			Default(false),
		field.Time("optional_time").
			Optional(),
	}
}

// Edges of the Dog.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", User.Type).
			Ref("team").
			Unique(),
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}

func (Pet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("owner"),
		index.Fields("nickname").
			Unique(),
	}
}
