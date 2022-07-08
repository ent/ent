// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text"),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("liked_users", User.Type).
			Ref("liked_tweets").
			Through("likes", TweetLike.Type),
		edge.From("user", User.Type).
			Ref("tweets").
			Through("tweet_user", UserTweet.Type).
			Comment("The uniqueness is enforced on the edge schema"),
		edge.From("tags", Tag.Type).
			Ref("tweets").
			Through("tweet_tags", TweetTag.Type),
	}
}
