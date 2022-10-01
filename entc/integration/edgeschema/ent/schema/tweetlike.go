// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/entc/integration/privacy/ent/privacy"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// TweetLike holds the schema definition for the TweetLike entity.
type TweetLike struct {
	ent.Schema
}

func (TweetLike) Annotations() []schema.Annotation {
	return []schema.Annotation{
		field.ID("user_id", "tweet_id"),
	}
}

// Fields of the TweetLike.
func (TweetLike) Fields() []ent.Field {
	return []ent.Field{
		field.Time("liked_at").
			Default(time.Now),
		field.Int("user_id"),
		field.Int("tweet_id"),
	}
}

// Edges of the TweetLike.
func (TweetLike) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tweet", Tweet.Type).
			Unique().
			Required().
			Field("tweet_id"),
		edge.To("user", User.Type).
			Unique().
			Required().
			Field("user_id"),
	}
}

// Policy defines the privacy policy of the TweetLike.
func (TweetLike) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
