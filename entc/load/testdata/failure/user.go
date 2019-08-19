package failure

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

type User struct {
	ent.Schema
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("panic", User{}.Type),
	}
}
