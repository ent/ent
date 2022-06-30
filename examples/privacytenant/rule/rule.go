// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package rule

import (
	"context"

	"entgo.io/ent/entql"
	"entgo.io/ent/examples/privacytenant/ent"
	"entgo.io/ent/examples/privacytenant/ent/privacy"
	"entgo.io/ent/examples/privacytenant/ent/user"
	"entgo.io/ent/examples/privacytenant/viewer"
)

// DenyIfNoViewer is a rule that returns deny decision if the viewer is missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereTenantID()
	// predicate that is used by both `Group` and `User` schemas.
	type TenantsFilter interface {
		WhereTenantID(entql.IntP)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		tid, ok := view.Tenant()
		if !ok {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that have an edge to it.
		tf.WhereTenantID(entql.IntEQ(tid))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}

// DenyMismatchedTenants is a rule that runs only on create operations and returns a deny
// decision if the operation tries to add users to groups that are not in the same tenant.
func DenyMismatchedTenants() privacy.MutationRule {
	return privacy.GroupMutationRuleFunc(func(ctx context.Context, m *ent.GroupMutation) error {
		tid, exists := m.TenantID()
		if !exists {
			return privacy.Denyf("missing tenant information in mutation")
		}
		users := m.UsersIDs()
		// If there are no users in the mutation, skip this rule-check.
		if len(users) == 0 {
			return privacy.Skip
		}
		// Query the tenant-ids of all attached users. Expect all users to be connected to the same tenant
		// as the group. Note, we use privacy.DecisionContext to skip the FilterTenantRule defined above.
		ids, err := m.Client().User.Query().Where(user.IDIn(users...)).Select(user.FieldTenantID).Ints(privacy.DecisionContext(ctx, privacy.Allow))
		if err != nil {
			return privacy.Denyf("querying the tenant-ids %v", err)
		}
		if len(ids) != len(users) {
			return privacy.Denyf("one the attached users is not connected to a tenant %v", err)
		}
		for _, id := range ids {
			if id != tid {
				return privacy.Denyf("mismatch tenant-ids for group/users %d != %d", tid, id)
			}
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
