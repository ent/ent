// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// CustomType holds the schema definition for the CustomType entity.
type CustomType struct {
	ent.Schema
}

// Fields of the CustomType.
func (CustomType) Fields() []ent.Field {
	return []ent.Field{
		field.String("custom").
			Optional().
			SchemaType(map[string]string{
				dialect.Postgres: "customtype",
			}),
	}
}
