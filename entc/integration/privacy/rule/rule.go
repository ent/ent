package rule

import (
	"context"

	"github.com/facebookincubator/ent/entc/integration/privacy/ent"
	"github.com/facebookincubator/ent/entc/integration/privacy/ent/privacy"
)

// DenyUpdateOperationRule is a write rule that denies update many operations.
func DenyUpdateOperationRule() privacy.WriteRule {
	return privacy.WriteRuleFunc(func(_ context.Context, m privacy.Mutation) error {
		if m.Op() == ent.OpUpdate {
			return privacy.Denyf("ent/privacy: update operation not allowed")
		}
		return privacy.Skip
	})
}

// DenyPlanetSelfLinkRule is a write rule that prevents planet self link via neighbor edge.
func DenyPlanetSelfLinkRule() privacy.WriteRule {
	return privacy.WriteRuleFunc(func(ctx context.Context, m privacy.Mutation) error {
		if m.Op() != ent.OpUpdateOne {
			return privacy.Skip
		}
		mutation, ok := m.(*ent.PlanetMutation)
		if !ok {
			return privacy.Denyf("ent/privacy: not a planet mutation")
		}
		id, exists := mutation.ID()
		if !exists {
			return privacy.Denyf("ent/privacy: planet id not provided")
		}
		for _, neighbor := range mutation.NeighborsIDs() {
			if id == neighbor {
				return privacy.Denyf("ent/privacy: planet %d cannot have itself as a neighbor", id)
			}
		}
		return privacy.Skip
	})
}
