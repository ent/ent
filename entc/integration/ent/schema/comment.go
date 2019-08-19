package schema

import (
	"fbc/ent"
	"fbc/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.Int("unique_int").
			Unique(),
		field.Float("unique_float").
			Unique(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
