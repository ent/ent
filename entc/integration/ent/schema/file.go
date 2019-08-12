package schema

import (
	"math"

	"fbc/ent"
	"fbc/ent/schema/edge"
	"fbc/ent/schema/field"
	"fbc/ent/schema/index"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int("size").
			Default(math.MaxInt32).
			Positive(),
		field.String("name"),
		field.String("user").
			Optional().
			Nillable(),
		field.String("group").
			Optional(),
	}
}

// Edges of the File.
func (File) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("files").
			Unique(),
	}
}

// Indexes of a file.
func (File) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index should not prevent duplicates.
		index.Fields("name", "size"),
		// unique index prevents duplicates records.
		index.Fields("name", "user").
			Unique(),
		// unique index under the "owner" sub-tree.
		// user/owner can't have files with duplicate names.
		index.Fields("name").
			FromEdge("owner").
			Unique(),
	}
}
