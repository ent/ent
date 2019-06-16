package schema

import (
	"fbc/ent"
	"fbc/ent/field"
)

// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}

// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.Int("size").
			Positive(),
		field.String("name"),
	}
}
