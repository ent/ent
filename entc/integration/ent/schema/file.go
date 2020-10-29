// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"math"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/index"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
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
