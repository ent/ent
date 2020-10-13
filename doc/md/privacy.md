---
id: privacy
title: Privacy
---

The `Policy` option in the schema allows configuring privacy policy for queries and mutations of entities in the database.

![gopher-privacy](https://entgo.io/assets/gopher-privacy-opacity.png)

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
ets access to the target nodes.

![privacy-rules](https://entgo.io/assets/permission_1.png)

However, if one of the evaluated rules returns an error or a `privacy.Deny` decision (see below), the executed operation
returns an error, and it is cancelled. 

![privacy-deny](https://entgo.io/assets/permission_2.png)

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

![privacy-allow](https://entgo.io/assets/permission_3.png)

Now, that we’ve covered the basic terms, let’s start writing some code.

## Configuration

In order to enable the privacy option in your code generation, enable the `privacy` feature with one of two options:

1\. If you are using the default go generate config, add `--feature privacy` option to the `ent/generate.go` file as follows:

```go
package ent
  
//go:generate go run github.com/facebook/ent/cmd/entc generate --feature privacy ./schema
```

2\. If you are using the configuration from the GraphQL documentation, add the feature flag as follows:

```go
// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// +build ignore

package main


import (
    "log"

    "github.com/facebook/ent/entc"
    "github.com/facebook/ent/entc/gen"
    "github.com/facebookincubator/ent-contrib/entgql"
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

> You should notice that, similar to schema hooks, if you use the **`Policy`** option in your schema,
> you **MUST** add the following import in the main package, because a circular import is possible:
>
> ```go
> import _ "<project>/ent/runtime"
> ```

## Examples

### Admin Only

We start with a simple example of an application that lets any user read any data, and accepts mutations only from users
with admin role. We will create 2 additional packages for the purpose of the examples:

- `rule` - for holding the different privacy rules in our schema.
- `viewer` - for getting and setting the user/viewer who's executing the operation. In this simple example, it can be
   either a normal user or an admin.
    
<br/>
After running the code-generation (with the feature-flag for privacy), we add the `Policy` method with 2 generated policy rules.

```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/examples/privacyadmin/ent/privacy"
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

```go
package rule

import (
	"context"

	"github.com/facebook/ent/examples/privacyadmin/ent/privacy"
	"github.com/facebook/ent/examples/privacyadmin/viewer"
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

```go
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

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Expect operation to fail, because viewer-context
	// is missing (first mutation rule check).
	if _, err := client.User.Create().Save(ctx); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
	}
	// Apply the same operation with "Admin" role.
	admin := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.Admin})
	if _, err := client.User.Create().Save(admin); err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
	// Apply the same operation with "ViewOnly" role.
	viewOnly := viewer.NewContext(ctx, viewer.UserViewer{Role: viewer.View})
	if _, err := client.User.Create().Save(viewOnly); !errors.Is(err, privacy.Deny) {
		return fmt.Errorf("expect operation to fail, but got %v", err)
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

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Bind a privacy decision to the context (bypass all other rules).
	allow := privacy.DecisionContext(ctx, privacy.Allow)
	if _, err := client.User.Create().Save(allow); err != nil {
		return fmt.Errorf("expect operation to pass, but got %v", err)
	}
    return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/privacyadmin).

Please note that this documentation is under active development.