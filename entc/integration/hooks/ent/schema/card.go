// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent"
	gen "entgo.io/ent/entc/integration/hooks/ent"
	"entgo.io/ent/entc/integration/hooks/ent/card"
	"entgo.io/ent/entc/integration/hooks/ent/hook"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// RejectUpdate rejects ~all update operations
// that are not on a specific entity.
type RejectUpdate struct {
	mixin.Schema
}

func (RejectUpdate) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			hook.Reject(ent.OpUpdate),
			// Accept only updates that contains the "expired_at" field.
			hook.Not(hook.HasFields(card.FieldExpiredAt)),
		),
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
			Comment("InHook is a mandatory field that is set by the hook."),
		field.Time("expired_at").
			Optional(),
	}
}

func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("cards").
			Unique(),
	}
}
