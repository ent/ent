package schema

import (
	"context"
	"log"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"
	"github.com/facebookincubator/ent/entc/integration/privacy/rule"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Planet defines the schema of a planet.
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
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				log.Println("Planet mutation of type", m.Type())
				return next.Mutate(ctx, m)
			})
		},
	}
}

func (Planet) Policy() ent.Policy {
	return privacy.Policy{
		WritePolicy: privacy.WritePolicy{
			rule.DenyUpdateOperationRule(),
			rule.DenyPlanetSelfLinkRule(),
			privacy.AlwaysAllowRule(),
		},
		ReadPolicy: privacy.ReadPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
