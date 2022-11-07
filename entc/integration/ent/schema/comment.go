// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	schemadir "entgo.io/ent/entc/integration/ent/schema/dir"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("unique_int").
			Unique(),
		field.Float("unique_float").
			Unique(),
		field.Int("nillable_int").
			Optional().
			Nillable(),
		field.String("table").
			Optional(),
		field.JSON("dir", schemadir.Dir("")).
			Optional(),
		field.String("client").
			Optional(),
	}
}
