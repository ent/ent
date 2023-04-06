// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"encoding/hex"
	"errors"

	"gocloud.dev/secrets"

	"entgo.io/ent"
	gen "entgo.io/ent/examples/encryptfield/ent"
	"entgo.io/ent/examples/encryptfield/ent/hook"
	"entgo.io/ent/examples/encryptfield/ent/intercept"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Comment("This PII field is encrypted before store in the database").
			Optional(),
		field.String("nickname").
			Comment("This field is stored as-is in the database"),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hook.If(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					v, ok := m.Name()
					if !ok || v == "" {
						return nil, errors.New("unexpected 'name' value")
					}
					c, err := m.SecretsKeeper.Encrypt(ctx, []byte(v))
					if err != nil {
						return nil, err
					}
					m.SetName(hex.EncodeToString(c))
					u, err := next.Mutate(ctx, m)
					if err != nil {
						return nil, err
					}
					if u, ok := u.(*gen.User); ok {
						// Another option, is to assign `u.Name = v` here.
						err = decrypt(ctx, m.SecretsKeeper, u)
					}
					return u, err
				})
			},
			hook.HasFields("name"),
		),
	}
}

// Interceptors of the User.
func (User) Interceptors() []ent.Interceptor {
	return []ent.Interceptor{
		ent.InterceptFunc(func(next ent.Querier) ent.Querier {
			return intercept.UserFunc(func(ctx context.Context, query *gen.UserQuery) (gen.Value, error) {
				v, err := next.Query(ctx, query)
				if err != nil {
					return nil, err
				}
				users, ok := v.([]*gen.User)
				// Skip all query types besides node queries (e.g., Count, Scan, GroupBy).
				if !ok {
					return v, nil
				}
				for _, u := range users {
					if err := decrypt(ctx, query.SecretsKeeper, u); err != nil {
						return nil, err
					}
				}
				return users, nil
			})
		}),
	}
}

func decrypt(ctx context.Context, k *secrets.Keeper, u *gen.User) error {
	b, err := hex.DecodeString(u.Name)
	if err != nil {
		return err
	}
	plain, err := k.Decrypt(ctx, b)
	if err != nil {
		return err
	}
	u.Name = string(plain)
	return nil
}
