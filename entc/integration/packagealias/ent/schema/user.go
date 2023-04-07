// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/entc/integration/packagealias/user"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Other("account", &user.User{}).SchemaType(map[string]string{
			dialect.SQLite:   "text",
			dialect.MySQL:    "text",
			dialect.Postgres: "varchar",
		}).Default(&user.User{Name: "TestName"}),
		field.Enum("state").Values("on", "off").Default("off"),
	}
}
