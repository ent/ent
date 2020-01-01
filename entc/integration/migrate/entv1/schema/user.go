// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int32("age"),
		field.String("name").
			MaxLen(10),
		field.String("nickname").
			Unique(),
		field.String("address").
			Optional(),
		field.String("renamed").
			Optional(),
		field.Bytes("blob").
			Optional().
			MaxLen(255),
		field.Enum("state").
			Optional().
			Values("logged_in", "logged_out"),
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
		index.Fields("name", "address").
			Unique(),
	}
}

type Car struct {
	ent.Schema
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("car").
			Unique(),
	}
}
