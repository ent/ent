// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// SocialProfile holds the schema for the social-profile entity.
type SocialProfile struct {
	ent.Schema
}

// Fields of the social-profile.
func (SocialProfile) Fields() []ent.Field {
	return []ent.Field{
		field.String("desc"),
	}
}

// Edges of the social-profile.
func (SocialProfile) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("social_profiles").
			Unique().
			Required(), // Mark as required as profile should not exist without a user.
	}
}
