// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/privacy/ent/privacy"
	"entgo.io/ent/entc/integration/privacy/rule"
	"entgo.io/ent/schema/field"
)

type Note struct {
	ent.Schema
}

// Fields of the note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			NotEmpty(),
		field.String("description").
			Optional(),
		field.Bool("readonly").
			Default(false),
	}
}

// Hooks for the note.
func (Note) Hooks() []ent.Hook {
	return []ent.Hook{
		rule.NoteMockHook(),
	}
}

// Policy of the note.
func (Note) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.FilterReadonlyNoteRule(),
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
