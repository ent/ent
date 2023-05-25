package schema

import "entgo.io/ent"

// PC holds the schema definition for the PC entity.
type PC struct {
	ent.Schema
}

// Fields of the PC.
func (PC) Fields() []ent.Field {
	return nil
}

// Edges of the PC.
func (PC) Edges() []ent.Edge {
	return nil
}
