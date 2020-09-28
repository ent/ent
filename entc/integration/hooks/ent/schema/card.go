// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	"github.com/facebook/ent"
	gen "github.com/facebook/ent/entc/integration/hooks/ent"
	"github.com/facebook/ent/entc/integration/hooks/ent/card"
	"github.com/facebook/ent/entc/integration/hooks/ent/hook"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// RejectMany rejects all update operations
// that are not on a specific entity.
type RejectUpdate struct {
	mixin.Schema
}

func (RejectUpdate) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.Reject(ent.OpUpdate),
	}
}

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Mixin() []ent.Mixin {
	return []ent.Mixin{
		RejectUpdate{},
	}
}

func (Card) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
					num, ok := m.Field(card.FieldNumber)
					if !ok {
						return nil, fmt.Errorf("missing card number value")
					}
					// Validator in hooks.
					if len(num.(string)) < 4 {
						return nil, fmt.Errorf("card number is too short")
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
		func(next ent.Mutator) ent.Mutator {
			return hook.CardFunc(func(ctx context.Context, m *gen.CardMutation) (ent.Value, error) {
				m.SetInHook("value was set in hook")
				if _, ok := m.Name(); !ok {
					m.SetName("unknown")
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}

func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			Immutable().
			Default("unknown").
			NotEmpty(),
		field.String("name").
			Optional().
			Comment("Exact name written on card"),
		field.Time("created_at").
			Default(time.Now),
		field.String("in_hook").
			Comment("A mandatory field that is set by the hook"),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("cards").
			Unique(),
	}
}
