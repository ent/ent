package schema

import (
	"fbc/ent"
	"fbc/ent/edge"
	"fbc/ent/field"
)

// User holds the schema for the user entity.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name").
			StructTag(`json:"first_name" graphql:"first_name"`),
		field.String("last").
			Default("unknown").
			StructTag(`graphql:"last_name"`),
		field.String("nickname").
			Optional().
			Unique(),
		field.String("phone").
			Optional().
			Unique(),
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).
			Comment("O2O edge").
			Unique(),
		edge.To("pets", Pet.Type),
		edge.To("files", File.Type),
		edge.To("groups", Group.Type),
		edge.To("friends", User.Type),
		edge.To("following", User.Type).From("followers"),
		edge.To("team", Pet.Type).Unique(),
		edge.To("spouse", User.Type).Unique(),
		edge.To("parent", User.Type).Unique().From("children"),
	}
}
