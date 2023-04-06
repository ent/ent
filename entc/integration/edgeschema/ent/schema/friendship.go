// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Friendship holds the edge schema definition of the Friendship relationship.
type Friendship struct {
	ent.Schema
}

// Fields of the Friendship.
func (Friendship) Fields() []ent.Field {
	return []ent.Field{
		field.Int("weight").
			Default(1),
		field.Time("created_at").
			Default(time.Now),
		field.Int("user_id").
			Immutable(),
		field.Int("friend_id").
			Immutable(),
	}
}

// Edges of the Friendship.
func (Friendship) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Immutable().
			Field("user_id"),
		edge.To("friend", User.Type).
			Unique().
			Required().
			Immutable().
			Field("friend_id"),
	}
}

// Indexes of the Friendship.
func (Friendship) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
		// By default, Ent generates a unique index named <T>_<FK1>_<FK2>
		// for edge-schemas with an ID field to enforce the uniqueness of
		// the edges reside in the join table. However, in this case it is
		// skipped because we define it explicitly in the definition below.
		index.Fields("user_id", "friend_id").
			Unique().
			StorageKey("friendships_edge"),
	}
}
