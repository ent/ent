package schema

import (
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Group holds the schema definition for the Group entity.
type Group struct {
	ent.Schema
}

// Fields of the Group.
// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			// regexp validation for group name.
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
	}
}

// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
