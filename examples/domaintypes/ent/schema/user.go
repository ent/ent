// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("postal_code").
			SchemaType(map[string]string{
				// Set the database column type to "us_postal_code" only in PostgreSQL.
				// In case this schema is used with other databases, it falls back to the
				// default type (e.g., "varchar").
				dialect.Postgres: "us_postal_code",
			}),
	}
}
