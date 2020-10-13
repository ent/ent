// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/examples/privacyadmin/ent/privacy"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Mixin of the User schema.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
		TenantMixin{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{},
		Query:    privacy.QueryPolicy{},
	}
}
