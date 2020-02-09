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
		field.Int("id").
			StorageKey("oid"),
		// changing the type of the field.
		field.Int("age"),
		// extending name field to longtext.
		field.Text("name"),
		// changing nickname from unique no non-unique.
		field.String("nickname"),
		// adding new columns (must be either optional, or with a default value).
		field.String("phone").
			Default("unknown"),
		field.Bytes("buffer").
			Optional(),
		// adding new column with supported default value
		// in the database side, will append this value to
		// all existing rows.
		field.String("title").
			Default("SWE"),
		// change column name and reference it to the
		// previous one ("renamed").
		field.String("new_name").
			Optional().
			StorageKey("renamed"),
		// extending the blob size.
		field.Bytes("blob").
			Optional().
			MaxLen(1000),
		// adding enum to the `state` column.
		field.Enum("state").
			Optional().
			Values("logged_in", "logged_out", "online"),
		// deleting the `address` column.
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge(children<-M2O->parent) to be dropped.
		// Edge(spouse<-O2O->spouse) to be dropped.
		edge.To("car", Car.Type),
		// A new edge was added.
		edge.To("pets", Pet.Type).
			Unique(),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// deleting old indexes (name, address),
		// and defining a new one.
		index.Fields("phone", "age").
			Unique(),
	}
}

type Car struct {
	ent.Schema
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Car now can have more than 1 owner (not unique anymore).
		edge.From("owner", User.Type).
			Ref("car").
			Unique(),
	}
}

// Additional types to be added to the schema.
type (
	// Pet schema.
	Pet struct{ ent.Schema }
	// Group schema.
	Group struct{ ent.Schema }
)
