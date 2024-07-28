// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// PetUserName represents a user/pet name returned from the pet_user_names view.
type PetUserName struct {
	ent.View
}

// Annotations of the PetUserName.
func (PetUserName) Annotations() []schema.Annotation {
	return []schema.Annotation{
		// The definition below returns all names in the system
		// by unifying the "users" table and the "pets" table.
		entsql.ViewFor(dialect.Postgres, func(s *sql.Selector) {
			s.SelectDistinct("name").
				From(
					s.New().Select("name").From(sql.Table("users")).
						Union(
							s.New().Select("name").From(sql.Table("pets")),
						).
						As("all_names"),
				)
		}),
		// Alternatively, you can use raw definitions to define the view.
		// But note, this definition is skipped if the ViewFor annotation
		// is defined for the dialect we generated migration to (Postgres).
		entsql.View(`SELECT DISTINCT name
FROM (
    SELECT users.name
    FROM users
    UNION
    SELECT pets.name
    FROM pets
) AS all_names;
`),
	}
}

// Fields of the PetUserName.
func (PetUserName) Fields() []ent.Field {
	return []ent.Field{
		// Skip adding the "id" field as
		// it is not needed for this view.
		field.String("name"),
	}
}
