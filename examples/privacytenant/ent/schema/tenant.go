// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/examples/privacytenant/ent/privacy"
	"github.com/facebook/ent/examples/privacytenant/rule"
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
		Mutation: privacy.MutationPolicy{
			// For Tenant type, we only allow admin users to mutate
			// the tenant information and deny otherwise.
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
	}
}
