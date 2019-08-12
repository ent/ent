package schema

import (
	"fbc/ent"
	"fbc/ent/schema/edge"
	"fbc/ent/schema/field"
)

// Node holds the schema definition for the linked-list Node entity.
type Node struct {
	ent.Schema
}

// Fields of the Node.
func (Node) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value").
			Optional(),
	}
}

// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Node.Type).
			Unique().
			From("prev").
			Unique(),
	}
}
