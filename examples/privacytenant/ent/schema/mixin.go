// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/examples/privacytenant/ent/privacy"
	"entgo.io/ent/examples/privacytenant/rule"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
			// Allow admins to query any information.
			rule.AllowIfAdmin(),
		},
		Mutation: privacy.MutationPolicy{
			// Deny any operation in case there is no "viewer context".
			rule.DenyIfNoViewer(),
		},
	}
}

// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Fields for all schemas that embed TenantMixin.
func (TenantMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("tenant_id").
			Immutable(),
	}
}

// Edges for all schemas that embed TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Field("tenant_id").
			Unique().
			Required().
			Immutable(),
	}
}

// Policy for all schemas that embed TenantMixin.
func (TenantMixin) Policy() ent.Policy {
	return rule.FilterTenantRule()
}
