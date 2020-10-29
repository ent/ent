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
)

// User defines the schema of a user.
type User struct {
	ent.Schema
}

// Mixin list of schemas to the user.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TeamMixin{},
	}
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Immutable().
			NotEmpty().
			Unique(),
		field.Uint("age").
			Optional(),
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
	}
}

// Policy of the user.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowUserCreateIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
