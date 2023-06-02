// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"

	"ariga.io/atlas/sql/postgres"
	"github.com/google/uuid"
)

type Mixin struct {
	mixin.Schema
}

func (m Mixin) Fields() []ent.Field {
	return []ent.Field{
		field.String("mixed_string").
			Default("default"),
		field.Enum("mixed_enum").
			Values("on", "off").
			Default("on"),
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("oid"),
		// add a new column.
		field.Bool("active").
			Default(true),
		// changing the type of the field.
		field.Int("age"),
		// extending name field to longtext.
		field.Text("name"),
		// extending the index prefix below (on MySQL).
		field.Text("description").
			Optional(),
		// changing nickname from unique no non-unique.
		field.String("nickname").
			MaxLen(255),
		// adding new columns (must be either optional, or with a default value).
		field.String("phone").
			Default("unknown"),
		field.Bytes("buffer").
			Optional().
			DefaultFunc(func() []byte { return []byte("null") }),
		// adding new column with supported default value
		// in the database side, will append this value to
		// all existing rows.
		field.String("title").
			Default("SWE"),
		// change column name and reference it to the
		// previous one using storage-key ("renamed").
		field.String("new_name").
			Optional().
			StorageKey("renamed"),
		// change column name from "old_token" to "new_token"
		// and use Atlas diff hook in the migration.
		field.String("new_token").
			DefaultFunc(uuid.NewString),
		// extending the blob size.
		field.Bytes("blob").
			Optional().
			MaxLen(1000),
		// adding enum to the `state` column.
		field.Enum("state").
			Optional().
			Values("logged_in", "logged_out", "online").
			Default("logged_in"),
		// convert string to enum.
		field.Enum("status").
			Optional().
			Values("done", "pending"),
		// remove the max-length constraint from varchar.
		field.String("workplace").
			Optional(),
		// JSON field with database-default value.
		field.Strings("roles").
			Optional().
			Annotations(entsql.Default(`[]`)),
		field.String("default_expr").
			Optional().
			Annotations(entsql.DefaultExpr("lower('hello')")),
		field.String("default_exprs").
			Optional().
			Annotations(entsql.DefaultExprs(map[string]string{
				dialect.MySQL:    "TO_BASE64('ent')",
				dialect.SQLite:   "hex('ent')",
				dialect.Postgres: "md5('ent')",
			})),
		// add a new column with generated values by the database.
		field.Time("created_at").
			Default(time.Now).
			Annotations(&entsql.Annotation{
				Default: "CURRENT_TIMESTAMP",
			}),
		// nullable field was changed to non-nullable without a static
		// default value, and it requires apply hook to fix this.
		field.String("drop_optional").
			DefaultFunc(uuid.NewString),
		// deleting the `address` column.
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// Edge(children<-M2O->parent) to be dropped.
		// Edge(spouse<-O2O->spouse) to be dropped.
		edge.To("car", Car.Type),
		// New edges to added.
		edge.To("pets", Pet.Type).
			StorageKey(edge.Column("owner_id"), edge.Symbol("user_pet_id")).
			Unique(),
		edge.To("friends", User.Type).
			StorageKey(
				edge.Table("friends"),
				edge.Columns("user", "friend"),
				edge.Symbols("user_friend_id1", "user_friend_id2"),
			),
	}
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
		// Extend the column prefix by drop and create
		// this index on MySQL.
		index.Fields("description").
			Annotations(entsql.Prefix(100)),
		// Deleting old indexes (name, address),
		// and defining a new one.
		index.Fields("phone", "age").
			Unique(),
		index.Fields("age").
			Annotations(entsql.Desc()),
		// Enable FULLTEXT search on "nickname"
		// field only in MySQL.
		index.Fields("nickname").
			Annotations(
				entsql.IndexTypes(map[string]string{
					dialect.MySQL: "FULLTEXT",
				}),
			),
		// For PostgreSQL, we can include in the index non-key columns.
		index.Fields("workplace").
			Annotations(
				entsql.IncludeColumns("nickname"),
			),
		// For PostgreSQL and SQLite, users can define partial indexes.
		index.Fields("phone").
			Annotations(
				entsql.IndexWhere(`active AND "phone" <> ''`),
			),
		// For PostgreSQL, operator classes can be configured for each field.
		index.Fields("age", "phone").
			Annotations(
				entsql.OpClassColumn("phone", "bpchar_pattern_ops"),
			),
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

func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Unique(),
	}
}

func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("car").
			Unique().
			// Make the M20 edge from nullable to required edge.
			// Requires column and foreign-key migration.
			Required(),
	}
}

// Group schema.
type Group struct{ ent.Schema }

// Blog schema.
type Blog struct{ ent.Schema }

func (Blog) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeSerial,
			}),
		field.Int("oid").
			SchemaType(map[string]string{
				dialect.Postgres: postgres.TypeSerial,
			}),
	}
}

func (Blog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admins", User.Type),
	}
}

// Pet schema.
type Pet struct {
	ent.Schema
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Unique(),
	}
}

func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
