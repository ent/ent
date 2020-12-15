package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			Immutable().
			Default("unknown").
			NotEmpty(),
		field.String("name").
			Optional().
			Comment("Exact name written on card"),
		field.Time("created_at").
			Default(time.Now),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("card").Unique(),
	}
}
