// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

type Mixin struct {
	mixin.Schema
}

// Annotations of the Mixin schema.
func (Mixin) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Charset: "utf8mb4"},
	}
}

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User schema.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Optional().
			Annotations(entsql.Annotation{
				Size: 128,
			}),
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
	return []schema.Annotation{
		entsql.Annotation{
			Table: "Users",
		},
	}
}
