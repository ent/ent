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

// Task defines the schema of a task.
type Task struct {
	ent.Schema
}

// Mixin list of schemas to the task.
func (Task) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TeamMixin{},
	}
}

// Fields of the task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Enum("status").
			Values("planned", "in_progress", "closed").
			Default("planned"),
	}
}

// Edges of the task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("tasks").
			Unique(),
	}
}

// Hooks for the task.
func (Task) Hooks() []ent.Hook {
	return []ent.Hook{
		rule.LogTaskMutationHook(),
	}
}

// Policy of the task.
func (Task) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowTaskCreateIfOwner(),
			rule.DenyIfStatusChangedByOther(),
			rule.AllowIfViewerInTheSameTeam(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
