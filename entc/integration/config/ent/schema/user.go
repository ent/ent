// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

var (
	incrementalDisabled = false
)

type Mixin struct {
	mixin.Schema
}

// Annotations of the Mixin schema.
func (Mixin) Annotations() []schema.Annotation {
	false := false
	return []schema.Annotation{
		entsql.Annotation{Charset: "utf8mb4"},
		entsql.Annotation{Unique: &false},
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User schema.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id").
			StorageKey("user_id").
			Annotations(entsql.Annotation{
				Incremental: &incrementalDisabled,
			}),
		field.String("name").
			Optional().
			Annotations(entsql.Annotation{
				Size: 128,
			}).Comment(`Name of the user.
Comment line1
Comment line2`),
		field.String("label").
			Optional(),
	}
}

// Mixin of the User schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		Mixin{},
	}
}

// Annotations of the User schema.
func (User) Annotations() []schema.Annotation {
	incremental := false
	return []schema.Annotation{
		entsql.Annotation{
			Table:       "Users",
			Incremental: &incremental,
		},
	}
}
