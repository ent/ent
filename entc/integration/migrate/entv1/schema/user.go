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

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("oid"),
		field.Int32("age"),
		field.String("name").
			MaxLen(10),
		field.Text("description").
			Optional(),
		field.String("nickname").
			Unique(),
		field.String("username").
			Optional(),
		field.String("address").
			Optional(),
		field.String("renamed").
			Optional(),
		field.String("old_token").
			DefaultFunc(uuid.NewString),
		field.Bytes("blob").
			Optional().
			MaxLen(255),
		field.Enum("state").
			Optional().
			Values("logged_in", "logged_out").
			Default("logged_in"),
		field.String("status").
			Optional(),
		field.String("workplace").
			MaxLen(30).
			Optional(),
		field.String("drop_optional").
			Optional(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", User.Type).
			From("parent").
			Unique(),
		edge.To("spouse", User.Type).
			Unique(),
		edge.To("car", Car.Type).
			Unique(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("description").
			Annotations(entsql.Prefix(50)),
		index.Fields("name", "address").
			Unique(),
	}
}

type Car struct {
	ent.Schema
}

// Annotations of the Car.
func (Car) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Car"},
	}
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("car").
			Unique(),
	}
}
