---
title: Automatic GraphQL Filters Generation
author: Ariel Mashraki
authorURL: "https://github.com/a8m"
authorImageURL: "https://avatars0.githubusercontent.com/u/7413593"
authorTwitter: arielmashraki
---

#### TL;DR

We added a new integration to the Ent GraphQL extension, that generates type-safe GraphQL filters (i.e. `Where` predicates)
from an `ent/schema`, and allows users to seamlessly map GraphQL queries to Ent queries.

For example, in order to get all `COMPLETED` todo items, we execute the following:

````graphql
query QueryAllCompletedTodos {
  todos(
    where: {
      status: COMPLETED,
    },
  ) {
    edges {
      node {
        id
      }
    }
  }
}
````

The generated GraphQL filters follow the Ent syntax. This means, the following query is also valid:

```graphql
query FilterTodos {
  todos(
    where: {
      or: [
        {
          has_parent: false,
          status: COMPLETED,
        },
        {
          status: IN_PROGRESS,
          has_parent_with: {
            priority_lt: 1,
            status_neq: COMPLETED,
          },
        }
      ]
    },
  ) {
    edges {
      node {
        id
      }
    }
  }
}
```

### Background
Many libraries that deal with data in Go choose the path of passing around empty interace instances
(`interface{}`) and use reflection at runtime to figure out how to map data to struct fields. Aside from the
performance penalty of using reflection everywhere, the big negative impact on teams is the
loss of type-safety. 

When APIs are explicit, known at compile-time (or even as we type), the feedback a developer receives around a 
large class of errors is almost immediate. Many defects are found early, and development is also much more fun!

Ent was designed to provide an excellent developer experience for teams working on applications with
large data-models. To facilitate this, we decided early on that one of the core design principles
of Ent is "statically typed and explicit API using code generation". This means, that for every
entity a developer defines in their `ent/schema`, explicit, type-safe code is generated for the
developer to efficiently interact with their data. For example, In the
[Filesystem Example in the ent repository](https://github.com/ent/ent/blob/master/examples/fs/), you will
find a schema named `File`:

```go
// File holds the schema definition for the File entity.
type File struct {
	ent.Schema
}
// Fields of the File.
func (File) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Bool("deleted").
			Default(false),
		field.Int("parent_id").
			Optional(),
	}
}
```
When the Ent code-gen runs, it will generate many predicate functions. For example, the following function which
can be used to filter `File`s by their `name` field:

```go
package file
// .. truncated ..

// Name applies the EQ predicate on the "name" field.
func Name(v string) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}
```

[GraphQL](https://graphql.org) is a query language for APIs originally created at Facebook. Similar to Ent,
GraphQL models data in graph concepts and facilitates type-safe queries. Around a year ago, we
released an integration between Ent and GraphQL. Similar to the [gRPC Integration](2021-06-28-gprc-ready-for-use.md),
the goal for this integration is to allow developers to easily create API servers that map to Ent, to mutate
and query data in their databases.

### Automatic GraphQL Filters Generation

In a recent community survey, the Ent + GraphQL integration was mentioned as one of the most
loved features of the Ent project. Until today, the integration allowed users to perform useful, albeit
basic queries against their data. Today, we announce the release of a feature that we think will
open up many interesting new use cases for Ent users: "Automatic GraphQL Filters Generation".

As we have seen above, the Ent code-gen maintains for us a suite of predicate functions in our Go codebase
that allow us to easily and explicitly filter data from our database tables. This power was,
until recently, not available (at least not automatically) to users of the Ent + GraphQL integration.
With automatic GraphQL filter generation, by making a single-line configuration change, developers
can now add to their GraphQL schema a complete set of "Filter Input Types" that can be used as predicates in their
GraphQL queries. In addition, the implementation provides runtime code that parses these predicates and maps them into
Ent queries. Let's see this in action:

### Generate Filter Input Types

In order to generate input filters (e.g. `TodoWhereInput`) for each type in your `ent/schema` package,
edit the `ent/entc.go` configuration file as follows:

```go
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithWhereFilters(true),
		entgql.WithSchemaPath("<PATH-TO-GRAPHQL-SCHEMA>"),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	err = entc.Generate("./schema", &gen.Config{}, entc.Extensions(ex))
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

If you're new to Ent and GraphQL, please follow the [Getting Started Tutorial](https://entgo.io/docs/tutorial-todo-gql).

Then, after running `go generate ./ent/...`, Ent will generate `<T>WhereInput` for each type in your schema, and
update the GraphQL schema as well, so you won't need to `autobind` them to `gqlgen` manually. For example:

```go
// TodoWhereInput represents a where input for filtering Todo queries.
type TodoWhereInput struct {
	Not *TodoWhereInput   `json:"not,omitempty"`
	Or  []*TodoWhereInput `json:"or,omitempty"`
	And []*TodoWhereInput `json:"and,omitempty"`

	// "created_at" field predicates.
	CreatedAt      *time.Time  `json:"created_at,omitempty"`
	CreatedAtNEQ   *time.Time  `json:"created_at_neq,omitempty"`
	CreatedAtIn    []time.Time `json:"created_at_in,omitempty"`
	CreatedAtNotIn []time.Time `json:"created_at_not_in,omitempty"`
	CreatedAtGT    *time.Time  `json:"created_at_gt,omitempty"`
	CreatedAtGTE   *time.Time  `json:"created_at_gte,omitempty"`
	CreatedAtLT    *time.Time  `json:"created_at_lt,omitempty"`
	CreatedAtLTE   *time.Time  `json:"created_at_lte,omitempty"`

	// "status" field predicates.
	Status      *todo.Status  `json:"status,omitempty"`
	StatusNEQ   *todo.Status  `json:"status_neq,omitempty"`
	StatusIn    []todo.Status `json:"status_in,omitempty"`
	StatusNotIn []todo.Status `json:"status_not_in,omitempty"`

    // .. truncated ..
}
```

```graphql
"""
TodoWhereInput is used for filtering Todo objects.
Input was generated by ent.
"""
input TodoWhereInput {
  not: TodoWhereInput
  and: [TodoWhereInput!]
  or: [TodoWhereInput!]
  
  """created_at field predicates"""
  created_at: Time
  created_at_neq: Time
  created_at_in: [Time!]
  created_at_not_in: [Time!]
  created_at_gt: Time
  created_at_gte: Time
  created_at_lt: Time
  created_at_lte: Time
  
  """status field predicates"""
  status: Status
  status_neq: Status
  status_in: [Status!]
  status_not_in: [Status!]
    
  # .. truncated ..
}
```

Next, to complete the integration we need to change 2 more lines:

1\. Edit the GraphQL schema to accept the new filter types:
```graphql {8}
type Query {
  todos(
    after: Cursor,
    first: Int,
    before: Cursor,
    last: Int,
    orderBy: TodoOrder,
    where: TodoWhereInput,
  ): TodoConnection
}
```

2\. Use the new filter types in GraphQL resolvers:
```go {5}
func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder, where *ent.TodoWhereInput) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
			ent.WithTodoFilter(where.Filter),
		)
}
```

### Filter Specification

As mentioned above, with the new GraphQL filter types, you can express the same Ent filters you use in Go.

#### Conjunction, disjunction and negation

The `Not`, `And` and `Or` operators can be added using the `not`, `and` and `or` fields. For example:

```graphql
{
  or: [
    {
      status: COMPLETED,
    },
    {
      not: {
        has_parent: true,
        status: IN_PROGRESS,
      }
    }
  ]
}
```

When multiple filter fields are provided, the `And` operator is added as conjunction implicitly.

```graphql
{
  status: COMPLETED,
  text_has_prefix: "GraphQL",
}
```
The above code will produce the following Ent query:

```go
client.Todo.
	Query().
	Where(
		todo.And(
			todo.StatusEQ(todo.StatusCompleted),
			todo.TextHasPrefix("GraphQL"),
		)
	).
	All(ctx)
```

#### Edge/Relation filters

[Edge (relation) predicates](https://entgo.io/docs/predicates#edge-predicates) can be expressed in the same Ent syntax:

```graphql
{
  has_parent: false,
  has_children_with: {
    status: IN_PROGRESS,
  }
}
```

The above code will produce the following Ent query:

```go
client.Todo.
	Query().
	Where(
		todo.HasParent(),
		todo.HasChildrenWith(
			todo.StatusEQ(todo.StatusInProgress),
		),
	).
	All(ctx)
```

### Implementation Example

A working example exists in [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example). 