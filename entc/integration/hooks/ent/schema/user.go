// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"

	"github.com/facebook/ent/entc/integration/hooks/ent/hook"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		VersionMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Uint("worth").
			Optional(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("cards", Card.Type),
		edge.To("friends", User.Type),
		edge.To("best_friend", User.Type).
			Unique(),
	}
}

type VersionMixin struct {
	mixin.Schema
}

func (VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("version").
			Default(0),
	}
}

func (VersionMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(VersionHook(), ent.OpUpdateOne),
	}
}

func VersionHook() ent.Hook {
	type OldSetVersion interface {
		SetVersion(int)
		Version() (int, bool)
		OldVersion(context.Context) (int, error)
	}
	return func(next ent.Mutator) ent.Mutator {
		// A hook that validates the "version" field is incremented by 1 on each update.
		// Note that this is just a dummy example, and it doesn't promise consistency in
		// the database.
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			ver, ok := m.(OldSetVersion)
			if !ok {
				return next.Mutate(ctx, m)
			}
			oldV, err := ver.OldVersion(ctx)
			if err != nil {
				return nil, err
			}
			curV, exists := ver.Version()
			if !exists {
				return nil, fmt.Errorf("version field is required in update mutation")
			}
			if curV != oldV+1 {
				return nil, fmt.Errorf("version field must be incremented by 1")
			}
			// Add an SQL predicate that validates the "version" column is equal
			// to "oldV" (it wasn't changed during the mutation by other process).
			return next.Mutate(ctx, m)
		})
	}
}
