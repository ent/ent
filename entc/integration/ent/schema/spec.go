package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type Spec struct {
	ent.Schema
}

// Edges of the Spec.
func (Spec) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type),
	}
}
