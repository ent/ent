package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Dog.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("team", User.Type).
			Unique().Ref("team"),
		edge.From("owner", User.Type).
			Unique().
			Ref("pets"),
	}
}
