// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/entc/integration/privacy/ent/privacy"
	"github.com/facebook/ent/entc/integration/privacy/rule"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// Team defines the schema of a team.
type Team struct {
	ent.Schema
}

// Mixin list of schemas to the team.
func (Team) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the team.
func (Team) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Edges of the team.
func (Team) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("tasks", Task.Type).
			Ref("teams"),
		edge.From("users", User.Type).
			Ref("teams"),
	}
}

// Policy of the team.
func (Team) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNotAdmin(),
			rule.DenyUpdateRule(),
		},
	}
}

// TeamMixin shared between task and user.
type TeamMixin struct {
	mixin.Schema
}

// Edges of the team-mixin.
func (TeamMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("teams", Team.Type),
	}
}

// Policy of the team-mixin.
func (TeamMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.FilterTeamRule(),
		},
	}
}
