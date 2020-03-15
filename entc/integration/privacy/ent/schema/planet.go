package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"
	"github.com/facebookincubator/ent/entc/integration/privacy/rule"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Planet defines the schema of a rule.
type Planet struct {
	ent.Schema
}

func (Planet) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Immutable().
			NotEmpty().
			Unique(),
		field.Uint("age").
			Optional(),
	}
}

func (Planet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("neighbors", Planet.Type),
	}
}

func (Planet) Hooks() []ent.Hook {
	return []ent.Hook{
		rule.LogPlanetMutationHook(),
	}
}

func (Planet) Policy() ent.Policy {
	return privacy.Policy{
		Write: privacy.WritePolicy{
			rule.DenyUpdateRule(),
			rule.DenyPlanetSelfLinkRule(),
			privacy.AlwaysAllowRule(),
		},
		Read: privacy.ReadPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
