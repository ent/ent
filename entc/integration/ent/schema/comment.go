package schema

import "fbc/ent"

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return nil
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return nil
}
