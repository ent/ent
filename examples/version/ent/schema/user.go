// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"context"
	"fmt"
	"time"

	gen "entgo.io/ent/examples/version/ent"
	"entgo.io/ent/examples/version/ent/hook"
	"entgo.io/ent/examples/version/ent/predicate"
	"entgo.io/ent/examples/version/ent/user"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
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
		field.Enum("status").
			Values("online", "offline"),
	}
}

// VersionMixin provides an optimistic concurrency
// control mechanism using a "version" field.
type VersionMixin struct {
	mixin.Schema
}

// Fields of the VersionMixin.
func (VersionMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int64("version").
			DefaultFunc(func() int64 {
				return time.Now().UnixNano()
			}).
			Comment("Unix time of when the latest update occurred"),
	}
}

// Hooks of the VersionMixin.
func (VersionMixin) Hooks() []ent.Hook {
	return []ent.Hook{
		// Apply the `OptimisticLock` hook only on `UpdateOne` operations,
		// and block all `Update` (update-many) operations as we don't have
		// access to the nodes that are affected by these mutation.
		hook.On(OptimisticLock(), ent.OpUpdateOne),
		hook.Reject(ent.OpUpdate),
	}
}

func OptimisticLock() ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (ent.Value, error) {
			oldV, err := m.OldVersion(ctx)
			if err != nil {
				return nil, err
			}
			curV, exists := m.Version()
			if !exists {
				curV = time.Now().UnixNano()
				m.SetVersion(curV)
			} else if curV <= oldV {
				return nil, fmt.Errorf("version field must be > previous value: %v <= %v", curV, oldV)
			}
			// Add an SQL predicate that validates the "version" column is equal
			// to "oldV" (it wasn't changed during the mutation by another process).
			m.Where(MatchVersion(oldV, curV))
			v, err := next.Mutate(ctx, m)
			if gen.IsNotFound(err) {
				id, _ := m.ID()
				return nil, fmt.Errorf("user %d was changed by another process", id)
			}
			return v, err
		})
	}
}

// MatchVersion returns a "dynamic User predicate". First, it checks that the "version"
// field was not changed by another process when executing the `UPDATE` operation. Next,
// it checks that the new value matches when `SELECT`-ing the node from the database.
func MatchVersion(oldV, curV int64) predicate.User {
	p := user.Version(oldV)
	return func(s *sql.Selector) {
		p(s)
		p = user.Version(curV)
	}
}
