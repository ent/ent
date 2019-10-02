// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebookincubator/ent"
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
		// changing the type of the field.
		field.Int("age"),
		// extending name field to longtext.
		field.Text("name"),
		// adding new columns.
		field.String("phone"),
		field.Bytes("buffer").
			Default([]byte("{}")),
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

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// deleting old indexes (name, address),
		// and defining a new one.
		index.Fields("phone", "age").
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
