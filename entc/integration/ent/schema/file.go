// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"math"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

func (File) Annotations() []schema.Annotation {
	return []schema.Annotation{
		edge.Annotation{
			StructTag: `json:"file_edges"`,
		},
	}
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int("set_id").
			Max(10).
			Optional(),
		field.Int("size").
			StorageKey("fsize").
			Default(math.MaxInt32).
			Positive(),
		field.String("name"),
		field.String("user").
			Optional().
			Nillable(),
		field.String("group").
			Optional(),
		field.Bool("op").
			Optional(),
		// Skip generating the "FieldID" predicate
		// as it conflicts with the "FieldID" constant.
		field.Int("field_id").
			Optional(),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("files").
			Unique(),
		edge.From("type", FileType.Type).
			Ref("files").
			Unique(),
		edge.To("field", FieldType.Type),
	}
}

// Indexes of a file.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index should not prevent duplicates.
		index.Fields("name", "size").
			StorageKey("file_name_size"),
		// unique index prevents duplicates records.
		index.Fields("name", "user").
			Unique(),
		// index on edges only.
		index.Edges("owner", "type"),
		// unique index under the "owner" sub-tree.
		// user/owner can't have files with duplicate names.
		index.Fields("name").
			Edges("owner", "type").
			Unique(),
		index.Fields("name").
			Edges("owner"),
	}
}
