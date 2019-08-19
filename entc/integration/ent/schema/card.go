package schema

import (
	"time"

	"fbc/ent"
	"fbc/ent/schema/edge"
	"fbc/ent/schema/field"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

// Fields of the Comment.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			MinLen(1),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Comment.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Comment("O2O inverse edge").
			Ref("card").
			Unique(),
	}
}
