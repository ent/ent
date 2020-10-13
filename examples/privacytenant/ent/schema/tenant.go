// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/examples/privacyadmin/ent/privacy"
	"github.com/facebook/ent/schema/field"
)

// Tenant holds the schema definition for the Tenant entity.
type Tenant struct {
	ent.Schema
}

// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}

// Fields of the Tenant.
func (Tenant) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty(),
	}
}

// Policy defines the privacy policy of the User.
func (Tenant) Policy() ent.Policy {
	return privacy.Policy{
		// For Tenant type, we accepts mutation only from users with admin
		// role, but this constraint is already defined in the BaseMixin.
		Mutation: privacy.MutationPolicy{
			privacy.AlwaysDenyRule(),
		},
		// Same as for mutation, we only allow to admin users to read the
		// global tenant information, and this defined in the BaseMixin.
		Query: privacy.QueryPolicy{
			privacy.AlwaysDenyRule(),
		},
	}
}
