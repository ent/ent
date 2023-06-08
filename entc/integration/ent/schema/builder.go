package schema

import "entgo.io/ent"

// Builder holds the schema definition for the Builder entity.
type Builder struct {
	ent.Schema
}

// Fields of the Builder.
func (Builder) Fields() []ent.Field {
	return nil
}

// Edges of the Builder.
func (Builder) Edges() []ent.Edge {
	return nil
}
