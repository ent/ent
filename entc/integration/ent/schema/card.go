package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
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
			Default(time.Now).
			Immutable(),
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
