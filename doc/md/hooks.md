---
id: hooks
title: Hooks
---

The `Hooks` option allows adding custom logic before and after operations that mutate the graph.

## Mutation

A mutation operation is an operation that mutate the database. For example, adding
a new node to the graph, remove an edge between 2 nodes or delete multiple nodes. 

There are 5 types of mutations:
- `Create` - Create node in the graph.
- `UpdateOne` - Update a node in the graph. For example, increment its field.
- `Update` - Update multiple nodes in the graph that match a predicate.
- `DeleteOne` - Delete a node from the graph.
- `Delete` - Delete all nodes that match a predicate.

<br>
Each generated node type has its own type of mutation. For example, all [`User` builders](crud.md#create-an-entity), share
the same generated `UserMutation` object.

However, all builder types implement the generic <a target="_blank" href="https://pkg.go.dev/github.com/facebook/ent?tab=doc#Mutation">`ent.Mutation`<a> interface.
 
## Hooks

Hooks are functions that get an <a target="_blank" href="https://pkg.go.dev/github.com/facebook/ent?tab=doc#Mutator">`ent.Mutator`<a> and return a mutator back.
They function as middleware between mutators. It's similar to the popular HTTP middleware pattern.

```go
type (
	// Mutator is the interface that wraps the Mutate method.
	Mutator interface {
		// Mutate apply the given mutation on the graph.
		Mutate(context.Context, Mutation) (Value, error)
	}

	// Hook defines the "mutation middleware". A function that gets a Mutator
	// and returns a Mutator. For example:
	//
	//	hook := func(next ent.Mutator) ent.Mutator {
	//		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	//			fmt.Printf("Type: %s, Operation: %s, ConcreteType: %T\n", m.Type(), m.Op(), m)
	//			return next.Mutate(ctx, m)
	//		})
	//	}
	//
	Hook func(Mutator) Mutator
)
```

There are 2 types of mutation hooks - **schema hooks** and **runtime hooks**.
**Schema hooks** are mainly used for defining custom mutation logic in the schema,
and **runtime hooks** are used for adding things like logging, metrics, tracing, etc.
Let's go over the 2 versions:

## Runtime hooks

Let's start with a short example that logs all mutation operations of all types:

```go
func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the auto migration tool.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
    // Add a global hook that runs on all types and all operations.
	client.Use(func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			start := time.Now()
			defer func() {
				log.Printf("Op=%s\tType=%s\tTime=%s\tConcreteType=%T\n", m.Op(), m.Type(), time.Since(start), m)
			}()
			return next.Mutate(ctx, m)
		})
	})
    client.User.Create().SetName("a8m").SaveX(ctx)
    // Output:
    // 2020/03/21 10:59:10 Op=Create	Type=Card	Time=46.23µs	ConcreteType=*ent.UserMutation
}
```

Global hooks are useful for adding traces, metrics, logs and more. But sometimes, users want more granularity:  

```go
func main() {
    // <client was defined in the previous block>

    // Add a hook only on user mutations.
	client.User.Use(func(next ent.Mutator) ent.Mutator {
        // Use the "<project>/ent/hook" to get the concrete type of the mutation.
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			return next.Mutate(ctx, m)
		})
	})
    
    // Add a hook only on update operations.
    client.Use(hook.On(Logger(), ent.OpUpdate|ent.OpUpdateOne))
    
    // Reject delete operations.
    client.Use(hook.Reject(ent.OpDelete|ent.OpDeleteOne))
}
```

Assume you want to share a hook that mutate a field between multiple types (e.g. `Group` and `User`).
There are ~2 ways to do this:

```go
// Option 1: use type assertion.
client.Use(func(next ent.Mutator) ent.Mutator {
    type NameSetter interface {
        SetName(value string)
    }
    return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
        // A schema with a "name" field must implement the NameSetter interface. 
        if ns, ok := m.(NameSetter); ok {
            ns.SetName("Ariel Mashraki")
        }
        return next.Mutate(ctx, m)
    })
})

// Option 2: use the generic ent.Mutation interface.
client.Use(func(next ent.Mutator) ent.Mutator {
	return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
        if err := m.SetField("name", "Ariel Mashraki"); err != nil {
            // An error is returned, if the field is not defined in
			// the schema, or if the type mismatch the field type.
        }
        return next.Mutate(ctx, m)
    })
})
```

## Schema hooks

Schema hooks are defined in the type schema and applied only on mutations that match the
schema type. The motivation for defining hooks in the schema is to gather all logic
regarding the node type in one place, which is the schema. 

```go
package schema

import (
	"context"
	"fmt"

    gen "<project>/ent"
    "<project>/ent/hook"

	"github.com/facebook/ent"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

// Hooks of the Card.
func (Card) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.CardFunc(func(ctx context.Context, m *gen.CardMutation) (ent.Value, error) {
					if num, ok := m.Number(); ok && len(num) < 10 {
						return nil, fmt.Errorf("card number is too short")
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		// Second hook.
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if s, ok := m.(interface{ SetName(string) }); ok {
					s.SetName("Boring")
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}
```
> **Note that** if you use **schema hooks**, you **MUST** add the following import in the
> main package, because a circular import is possible.
>
> ```go
> import _ "<project>/ent/runtime"
> ```

## Evaluation order

Hooks are called in the order they were registered to the client. Thus, `client.Use(f, g, h)` 
executes `f(g(h(...)))` on mutations.

Also note, that **runtime hooks** are called before **schema hooks**. That is, if `g`,
and `h` were defined in the schema, and `f` was registered using `client.Use(...)`,
they will be executed as follows: `f(g(h(...)))`. 

## Hook helpers

The generated hooks package provides several helpers that can help you control when a hook will
be executed.

```go
package schema

import (
	"context"
	"fmt"

	"<project>/ent/hook"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/mixin"
)


type SomeMixin struct {
	mixin.Schema
}

func (SomeMixin) Hooks() []ent.Hook {
    return []ent.Hook{
        // Execute "HookA" only for the UpdateOne and DeleteOne operations.
        hook.On(HookA(), ent.OpUpdateOne|ent.OpDeleteOne),
        // Don't execute "HookB" on Create operation.
        hook.Unless(HookB(), ent.OpCreate),
        // Execute "HookC" only if the ent.Mutation is changing the "status" field,
        // and clearing the "dirty" field.
        hook.If(HookC(), hook.And(hook.HasFields("status"), hook.HasClearedFields("dirty"))),
    }
}
```

## Transaction Hooks

Hooks can also be registered on active transactions, and will be executed on `Tx.Commit` or `Tx.Rollback`.
For more information, read about it in the [transactions page](transactions.md#hooks). 

## Codegen Hooks

The `entc` package provides an option to add a list of hooks (middlewares) to the code-generation phase.
For more information, read about it in the [codegen page](code-gen.md#code-generation-hooks). 
