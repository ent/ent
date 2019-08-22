package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// FileType holds the schema definition for the FileType entity.
type FileType struct {
	ent.Schema
}

// Fields of the FileType.
func (FileType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
	}
}

// Edges of the FileType.
func (FileType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("files", File.Type),
	}
}
