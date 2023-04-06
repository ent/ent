// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/entc/integration/privacy/ent/privacy"
	"entgo.io/ent/schema/edge"
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
			Default("Unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type).
			Through("joined_groups", UserGroup.Type),
		edge.To("friends", User.Type).
			Through("friendships", Friendship.Type),
		edge.To("relatives", User.Type).
			Through("relationship", Relationship.Type),
		edge.To("liked_tweets", Tweet.Type).
			Through("likes", TweetLike.Type),
		edge.To("tweets", Tweet.Type).
			Through("user_tweets", UserTweet.Type),
		edge.To("roles", Role.Type).
			Through("roles_users", RoleUser.Type),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysAllowRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
