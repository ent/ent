---
id: hooks
title: Hooks
---

Hooks allows adding custom logic before and after operations that mutate the graph.

## Mutation

A mutation operation, is an operation that mutate the database. For example, adding
a new node to the graph, remove an edge between 2 nodes or delete multiple nodes. 

There are 5 types of mutations:
- `Create` - Create node in the graph.
- `UpdateOne` - Update a node in the graph. For example, rename its field.
- `Update` - Update multiple nodes in the graph that match a predicate.
- `DeleteOne` - Delete a node from the graph.
- `Delete` - Delete all nodes that match a predicate.

<br>
Each generated node type has its own type of mutation. For example, the [CRUD builders](crud.md#create-an-entity)
for the `User` entity, share (and use) the generated `UserMutation` object.

However, all builder types implement the generic [`ent.Mutation`](https://pkg.go.dev/github.com/facebookincubator/ent?tab=doc#Mutation) interface.
 
## Hooks

Hooks are functions that get a mutator and return a mutator back. They function as middleware
between mutators. It's similar to the popular middleware pattern that is used in HTTP in Go.

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
**Schema hooks** are mainly used for defining custom logic in the type mutations,
and **runtime hooks** are used for adding things like logging, metrics, tracing, etc.
Let's go over the 2 versions:

## Schema hooks

Schema hooks are defined in the schema and applied only on mutations of the schema type.
The motivation for defining hooks in the schema is to gather all logic regarding the node
type in one place, which is the schema. 

Let's see an example:

```go
package schema

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent"
)

// Card holds the schema definition for the CreditCard entity.
type Card struct {
	ent.Schema
}

func (Card) Hooks() []ent.Hook {
	return []ent.Hook{
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				num, ok := m.Field("number")
				if !ok {
					return nil, fmt.Errorf("missing card number value")
				}
				// Validator in hooks.
				if len(num.(string)) < 4 {
					return nil, fmt.Errorf("card number is too short")
				}
				return next.Mutate(ctx, m)
			})
		},
		func(next ent.Mutator) ent.Mutator {
			return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
				if s, ok := m.(interface{ SetName(string) }); ok {
					s.SetName("boring")
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}
```

**Note that** if you use **schema hooks**, you must add the following import to your
main package:

```go
import _ "<project>/ent/runtime"
```

After running the first code generation for your schema, you'll be able to use the
typed-mutations in your hooks:

```go
package schema

import (
	"context"
	"fmt"

	gen "<project>/ent"
	"<project>/ent/hook"

	"github.com/facebookincubator/ent"
)

func (Card) Hooks() []ent.Hook {
	return []ent.Hook{
        // ...
		func(next ent.Mutator) ent.Mutator {
			return hook.CardFunc(func(ctx context.Context, m *gen.CardMutation) (ent.Value, error) {
				if _, ok := m.Name(); !ok {
					m.SetName("unknown")
				}
				return next.Mutate(ctx, m)
			})
		},
	}
}
```

