package schema

import (
	"math"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/index"
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
		edge.From("type", FileType.Type).
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
			FromEdges("owner", "type").
			Unique(),
	}
}
