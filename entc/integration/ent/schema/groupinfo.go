package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// GroupInfo holds the schema for the group-info entity.
type GroupInfo struct {
	ent.Schema
}

// Fields of the group.
func (GroupInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
		field.Int("max_users").
			Default(1e4),
	}
}

// Edges of the group.
func (GroupInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("info"),
	}
}
