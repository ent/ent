// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Media holds the schema definition for the Media entity.
type Media struct {
	ent.Schema
}

// Fields of the Media.
func (Media) Fields() []ent.Field {
	return []ent.Field{
		field.String("source").
			Optional(),
		field.String("source_uri").
			Optional(),
		field.Text("text").
			Optional(),
	}
}

// Indexes of the Media.
func (Media) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("source", "source_uri").
			Annotations(entsql.PrefixColumn("source", 100)).
			Unique(),
		// MySQL allow indexing text column prefix.
		index.Fields("text").
			Annotations(entsql.Prefix(100)),
	}
}

func (Media) Annotations() []schema.Annotation {
	return []schema.Annotation{
		&entsql.Annotation{
			Check: "text <> 'boring'",
			Checks: map[string]string{
				"boring_check": "source_uri <> 'entgo.io'",
			},
		},
	}
}
