// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// An example for an edge-schema (AttachedFile) that uses a custom name for its edges
// (not foreign-keys) and Ent matches by the schema types and not foreign-key names.

// Process holds the edge schema definition of the Process relationship.
type Process struct {
	ent.Schema
}

// Edges of the Process.
func (Process) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type).
			Through("attached_files", AttachedFile.Type).
			Comment("Files that were attached by this process"),
	}
}

// File holds the edge schema definition of the File relationship.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("processes", Process.Type).
			Ref("files"),
	}
}

// AttachedFile holds the edge schema definition of the File relationship.
type AttachedFile struct {
	ent.Schema
}

// Fields of the AttachedFile.
func (AttachedFile) Fields() []ent.Field {
	return []ent.Field{
		field.Time("attach_time").
			Default(time.Now),
		field.Int("f_id"),
		field.Int("proc_id"),
	}
}

// Edges of the AttachedFile.
func (AttachedFile) Edges() []ent.Edge {
	return []ent.Edge{
		// Note: the two following edges use different name conventions (e.g., f_id <> file_id), but
		// Ent knows how to resolve this as there is only one usage of each type in this declaration.
		edge.To("fi", File.Type).
			Required().
			Unique().
			Field("f_id"),
		edge.To("proc", Process.Type).
			Required().
			Unique().
			Field("proc_id"),
	}
}
