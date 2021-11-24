---
id: privacy
title: Privacy
---

The `Policy` option in the schema allows configuring privacy policy for queries and mutations of entities in the database.

![gopher-privacy](https://entgo.io/images/assets/gopher-privacy-opacity.png)

The main advantage of the privacy layer is that, you write the privacy policy **once** (in the schema), and it is **always**
evaluated. No matter where queries and mutations are performed in your codebase, it will always go through the privacy layer.

In this tutorial, we will start by going over the basic terms we use in the framework, continue with a section for configuring
the policy feature to your project, and finish with a few examples.

## Basic Terms

### Policy 

The `ent.Policy` interface contains two methods: `EvalQuery` and `EvalMutation`. The first defines the read-policy, and
the second defines the write-policy. A policy contains zero or more privacy rules (see below). These rules are evaluated
in the same order they are declared in the schema.

If all rules are evaluated without returning an error, the evaluation finishes successfully, and the executed operation 
gets access to the target nodes.

![privacy-rules](https://entgo.io/images/assets/permission_1.png)

However, if one of the evaluated rules returns an error or a `privacy.Deny` decision (see below), the executed operation
returns an error, and it is cancelled.

![privacy-deny](https://entgo.io/images/assets/permission_2.png)

### Privacy Rules

Each policy (mutation or query) includes one or more privacy rules. The function signature for these rules is as follows:

```go
// EvalQuery defines the a read-policy rule.
func(Policy) EvalQuery(context.Context, Query) error

// EvalMutation defines the a write-policy rule.
func(Policy) EvalMutation(context.Context, Mutation) error
```

### Privacy Decisions

There are three types of decision that can help you control the privacy rules evaluation.

- `privacy.Allow` - If returned from a privacy rule, the evaluation stops (next rules will be skipped), and the executed
   operation (query or mutation) gets access to the target nodes.
   
- `privacy.Deny` - If returned from a privacy rule, the evaluation stops (next rules will be skipped), and the executed
  operation is cancelled. This equivalent to returning any error. 
  
- `privacy.Skip` - Skip the current rule, and jump to the next privacy rule. This equivalent to returning a `nil` error.

![privacy-allow](https://entgo.io/images/assets/permission_3.png)

Now that we’ve covered the basic terms, let’s start writing some code.

## Configuration

In order to enable the privacy option in your code generation, enable the `privacy` feature with one of two options:

1\. If you are using the default go generate config, add `--feature privacy` option to the `ent/generate.go` file as follows:

```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature privacy ./schema
```

It is recommended to add the [`schema/snapshot`](features.md#auto-solve-merge-conflicts) feature-flag along with the
`privacy` to enhance the development experience (e.g. `--feature privacy,schema/snapshot`)

2\. If you are using the configuration from the GraphQL documentation, add the feature flag as follows:

```go {20}
// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// +build ignore

package main


import (
    "log"

    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "entgo.io/contrib/entgql"
)

func main() {
	opts := []entc.Option{
        entc.FeatureNames("privacy"),
	}
    err := entc.Generate("./schema", &gen.Config{
        Templates: entgql.AllTemplates,
    }, opts...)
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

:::important
You should notice that, similar to [schema hooks](hooks.md#hooks-registration), if you use the **`Policy`** option in your schema,
you **MUST** add the following import in the main package, because a circular import is possible between the schema package,
and the generated ent package:

```go
import _ "<project>/ent/runtime"
```
:::

## Examples

### Admin Only

We start with a simple example of an application that lets any user read any data, and accepts mutations only from users
with admin role. We will create 2 additional packages for the purpose of the examples:

- `rule` - for holding the different privacy rules in our schema.
- `viewer` - for getting and setting the user/viewer who's executing the operation. In this simple example, it can be
   either a normal user or an admin.

After running the code-generation (with the feature-flag for privacy), we add the `Policy` method with 2 generated policy rules.

```go title="examples/privacyadmin/ent/schema/user.go"
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/examples/privacyadmin/ent/privacy"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
            // Deny if not set otherwise. 
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
            // Allow any viewer to read anything.
			privacy.AlwaysAllowRule(),
		},
	}
}
```

We defined a policy that rejects any mutation and accepts any query. However, as mentioned above, in this example,
we accept mutations only from viewers with admin role. Let's create 2 privacy rules to enforce this:

```go title="examples/privacyadmin/rule/rule.go"
package rule

import (
	"context"

	"entgo.io/ent/examples/privacyadmin/ent/privacy"
	"entgo.io/ent/examples/privacyadmin/viewer"
)

// DenyIfNoViewer is a rule that returns Deny decision if the viewer is
// missing in the context.
func DenyIfNoViewer() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view == nil {
			return privacy.Denyf("viewer-context is missing")
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}

// AllowIfAdmin is a rule that returns Allow decision if the viewer is admin.
func AllowIfAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		view := viewer.FromContext(ctx)
		if view.Admin() {
			return privacy.Allow
		}
		// Skip to the next privacy rule (equivalent to returning nil).
		return privacy.Skip
	})
}
```

As you can see, the first rule `DenyIfNoViewer`, makes sure every operation has a viewer in its context,
otherwise, the operation rejected. The second rule `AllowIfAdmin`, accepts any operation from viewer with
admin role. Let's add them to the schema, and run the code-generation:

```go title="examples/privacyadmin/ent/schema/user.go"
// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
			rule.AllowIfAdmin(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
```

Since we define the `DenyIfNoViewer` first, it will be executed before all other rules, and accessing the 
`viewer.Viewer` object is safe in the `AllowIfAdmin` rule.

After adding the rules above and running the code-generation, we expect the privacy-layer logic to be applied on
`ent.Client` operations.

```go title="examples/privacyadmin/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if err := client.User.Create().Exec(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}
	// Apply the same operation with "Admin" role.
	admin := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	if err := client.User.Create().Exec(admin); err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
	// Apply the same operation with "ViewOnly" role.
	viewOnly := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if err := client.User.Create().Exec(viewOnly); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}
	// Allow all viewers to query users.
	for _, ctx := range []context.Context{ctx, viewOnly, admin} {
		// Operation should pass for all viewers.
		count := client.User.Query().CountX(ctx)
		fmt.Println(count)
	}
	return nil
}
```

### Decision Context

Sometimes, we want to bind a specific privacy decision to the `context.Context`. In cases like this, we
can use the `privacy.DecisionContext` function to create a new context with a privacy decision attached to it.

```go title="examples/privacyadmin/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
	// Bind a privacy decision to the context (bypass all other rules).
	allow := privacy.DecisionContext(ctx, privacy.Allow)
	if err := client.User.Create().Exec(allow); err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
    return nil
}
```

The full example exists in [GitHub](https://github.com/ent/ent/tree/master/examples/privacyadmin).

### Multi Tenancy

In this example, we're going to create a schema with 3 entity types - `Tenant`, `User` and `Group`.
The helper packages `viewer` and `rule` (as mentioned above) also exist in this example to help us structure the application.

![tenant-example](https://entgo.io/images/assets/tenant_medium.png)

Let's start building this application piece by piece. We begin by creating 3 different schemas (see the full code [here](https://github.com/ent/ent/tree/master/examples/privacytenant/ent/schema)),
and since we want to share some logic between them, we create another [mixed-in schema](schema-mixin.md) and add it to all other schemas as follows:

```go title="examples/privacytenant/ent/schema/mixin.go"
// BaseMixin for all schemas in the graph.
type BaseMixin struct {
	mixin.Schema
}

// Policy defines the privacy policy of the BaseMixin.
func (BaseMixin) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.DenyIfNoViewer(),
		},
		Query: privacy.QueryPolicy{
			rule.DenyIfNoViewer(),
		},
	}
}
```

```go title="examples/privacytenant/ent/schema/tenant.go"
// Mixin of the Tenant schema.
func (Tenant) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseMixin{},
	}
}
```

As explained in the first example, the `DenyIfNoViewer` privacy rule, denies the operation if the `context.Context` does not
contain the `viewer.Viewer` information.

Similar to the previous example, we want add a constraint that only admin users can create tenants (and deny otherwise).
We do it by copying the `AllowIfAdmin` rule from above, and adding it to the `Policy` of the `Tenant` schema:

```go title="examples/privacytenant/ent/schema/tenant.go"
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
```

Then, we expect the following code to run successfully:

```go title="examples/privacytenant/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if err := client.Tenant.Create().Exec(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}
	// Deny tenant creation if the viewer is not admin.
	viewCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if err := client.Tenant.Create().Exec(viewCtx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %w", err)
	}
	// Apply the same operation with "Admin" role, expect it to pass.
	adminCtx := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	hub, err := client.Tenant.Create().SetName("GitHub").Save(adminCtx)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
	fmt.Println(hub)
	lab, err := client.Tenant.Create().SetName("GitLab").Save(adminCtx)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
	fmt.Println(lab)
    return nil
}
```

We continue by adding the rest of the edges in our data-model (see image above), and since both `User` and `Group` have
an edge to the `Tenant` schema, we create a shared [mixed-in schema](schema-mixin.md) named `TenantMixin` for this:

```go title="examples/privacytenant/ent/schema/mixin.go"
// TenantMixin for embedding the tenant info in different schemas.
type TenantMixin struct {
	mixin.Schema
}

// Edges for all schemas that embed TenantMixin.
func (TenantMixin) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tenant", Tenant.Type).
			Unique().
			Required(),
	}
}
```

Next, we may want to enforce a rule that will limit viewers to only query groups and users that are connected to the tenant they belong to.
For use cases like this, Ent has an additional type of privacy rule named `Filter`.
We can use `Filter` rules to filter out entities based on the identity of the viewer.
Unlike the rules we previously discussed, `Filter` rules can limit the scope of the queries a viewer can make, in addition to returning privacy decisions. 

> Note, the privacy filtering option needs to be enabled using the [`entql`](features.md#entql-filtering) feature-flag (see instructions [above](#configuration)).

```go title="examples/privacytenant/rule/rule.go"
// FilterTenantRule is a query/mutation rule that filters out entities that are not in the tenant.
func FilterTenantRule() privacy.QueryMutationRule {
	// TenantsFilter is an interface to wrap WhereHasTenantWith()
	// predicate that is used by both `Group` and `User` schemas.
	type TenantsFilter interface {
		WhereHasTenantWith(...predicate.Tenant)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		view := viewer.FromContext(ctx)
		if view.Tenant() == "" {
			return privacy.Denyf("missing tenant information in viewer")
		}
		tf, ok := f.(TenantsFilter)
		if !ok {
			return privacy.Denyf("unexpected filter type %T", f)
		}
		// Make sure that a tenant reads only entities that has an edge to it.
		tf.WhereHasTenantWith(tenant.Name(view.Tenant()))
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
```

After creating the `FilterTenantRule` privacy rule, we add it to the `TenantMixin` to make sure **all schemas**
that use this mixin, will also have this privacy rule.

```go title="examples/privacytenant/ent/schema/mixin.go"
// Policy for all schemas that embed TenantMixin.
func (TenantMixin) Policy() ent.Policy {
	return privacy.Policy{
		Query: privacy.QueryPolicy{
			rule.AllowIfAdmin(),
			// Filter out entities that are not connected to the tenant.
			// If the viewer is admin, this policy rule is skipped above.
			rule.FilterTenantRule(),
		},
	}
}
```

Then, after running the code-generation, we expect the privacy-rules to take effect on the client operations.

```go title="examples/privacytenant/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
    // A continuation of the code-block above.

	// Create 2 users connected to the 2 tenants we created above
	hubUser := client.User.Create().SetName("a8m").SetTenant(hub).SaveX(adminCtx)
	labUser := client.User.Create().SetName("nati").SetTenant(lab).SaveX(adminCtx)

	hubView := viewer.NewContext(ctx, viewer.UserViewer{T: hub})
	out := client.User.Query().OnlyX(hubView)
	// Expect that "GitHub" tenant to read only its users (i.e. a8m).
	if out.ID != hubUser.ID {
		return fmt.Errorf("expect result for user query, got %v", out)
	}
	fmt.Println(out)

	labView := viewer.NewContext(ctx, viewer.UserViewer{T: lab})
	out = client.User.Query().OnlyX(labView)
	// Expect that "GitLab" tenant to read only its users (i.e. nati).
	if out.ID != labUser.ID {
		return fmt.Errorf("expect result for user query, got %v", out)
	}
	fmt.Println(out)
	return nil
}
```

We finish our example with another privacy-rule named `DenyMismatchedTenants` on the `Group` schema.
The `DenyMismatchedTenants` rule rejects group creation if the associated users don't belong to
the same tenant as the group.

```go title="examples/privacytenant/rule/rule.go"
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
		// Query the tenant-id of all users. Expect to have exact 1 result,
		// and it matches the tenant-id of the group above.
		id, err := m.Client().User.Query().Where(user.IDIn(users...)).QueryTenant().OnlyID(ctx)
		if err != nil {
			return privacy.Denyf("querying the tenant-id %v", err)
		}
		if id != tid {
			return privacy.Denyf("mismatch tenant-ids for group/users %d != %d", tid, id)
		}
		// Skip to the next privacy rule (equivalent to return nil).
		return privacy.Skip
	})
}
```

We add this rule to the `Group` schema and run code-generation.

```go title="examples/privacytenant/ent/schema/group.go"
// Policy defines the privacy policy of the Group.
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operation
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				ent.OpCreate,
			),
		},
	}
}
```

Again, we expect the privacy-rules to take effect on the client operations.

```go title="examples/privacytenant/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
    // A continuation of the code-block above.

	// Expect operation to fail because the DenyMismatchedTenants rule
	// makes sure the group and the users are connected to the same tenant.
	err = client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(labUser).Exec(adminCtx)
	if !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, since user (nati) is not connected to the same tenant")
	}
	err = client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(labUser, hubUser).Exec(adminCtx)
	if !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, since some users (nati) are not connected to the same tenant")
	}
	entgo, err := client.Group.Create().SetName("entgo.io").SetTenant(hub).AddUsers(hubUser).Save(adminCtx)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
	fmt.Println(entgo)
	return nil
}
```

In some cases, we want to reject user operations on entities that don't belong to their tenant **without loading
these entities from the database** (unlike the `DenyMismatchedTenants` example above). 
To achieve this, we can use the `FilterTenantRule` rule for mutations as well, but limit it to specific operations as follows:

```go title="examples/privacytenant/ent/schema/group.go"
// Policy defines the privacy policy of the Group.
func (Group) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			// Limit DenyMismatchedTenants only for
			// Create operations
			privacy.OnMutationOperation(
				rule.DenyMismatchedTenants(),
				ent.OpCreate,
			),
			// Limit the FilterTenantRule only for
			// UpdateOne and DeleteOne operations.
			privacy.OnMutationOperation(
				rule.FilterTenantRule(),
				ent.OpUpdateOne|ent.OpDeleteOne,
			),
		},
	}
}
```

Then, we expect the privacy-rules to take effect on the client operations.

```go title="examples/privacytenant/example_test.go"
func Do(ctx context.Context, client *ent.Client) error {
    // A continuation of the code-block above.

	// Expect operation to fail, because the FilterTenantRule rule makes sure
	// that tenants can update and delete only their groups.
	err = entgo.Update().SetName("fail.go").Exec(labView)
	if !ent.IsNotFound(err) {
		return fmt.Errorf("expect operation to fail, since the group (entgo) is managed by a different tenant (hub), but got %w", err)
	}
	entgo, err = entgo.Update().SetName("entgo").Save(hubView)
	if err != nil {
		return fmt.Errorf("expect operation to pass, but got %w", err)
	}
	fmt.Println(entgo)
	return nil
}
```

The full example exists in [GitHub](https://github.com/ent/ent/tree/master/examples/privacytenant).

Please note that this documentation is under active development.
