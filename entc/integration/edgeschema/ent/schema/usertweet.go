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

// UserTweet holds the schema definition for the UserTweet entity.
type UserTweet struct {
	ent.Schema
}

// Fields of the UserTweet.
func (UserTweet) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the UserTweet.
func (UserTweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
	}
}

// Indexes of the UserTweet.
func (UserTweet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tweet_id").
			Unique(),
	}
}
