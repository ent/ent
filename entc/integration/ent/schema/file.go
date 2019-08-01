package schema

import (
	"math"

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
			Default(math.MaxInt32).
			Positive(),
		field.String("name"),
	}
}
