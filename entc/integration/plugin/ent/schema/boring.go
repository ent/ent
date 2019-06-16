package schema

import "fbc/ent"

// Boring holds the schema definition for the Boring entity.
type Boring struct {
	ent.Schema
}

// Fields of the Boring.
func (Boring) Fields() []ent.Field {
	return nil
}

// Edges of the Boring.
func (Boring) Edges() []ent.Edge {
	return nil
}
