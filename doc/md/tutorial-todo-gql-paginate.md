---
id: tutorial-todo-gql-paginate
title: Relay Cursor Connections (Pagination)
sidebar_label: Relay Cursor Connections
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.md) by explaining how to implement the 
[Relay Cursor Connections Spec](https://relay.dev/graphql/connections.htm). If you're not familiar with the
Cursor Connections interface, read the following paragraphs that were taken from [relay.dev](https://relay.dev/graphql/connections.htm#sel-DABDDDAADFA0E3kM):

> In the query, the connection model provides a standard mechanism for slicing and paginating the result set.
> 
> In the response, the connection model provides a standard way of providing cursors, and a way of telling the client
> when more results are available.
> 
> An example of all four of those is the following query:
> ```graphql
> {
>   user {
>     id
>     name
>     friends(first: 10, after: "opaqueCursor") {
>       edges {
>         cursor
>         node {
>           id
>           name
>         }
>       }
>       pageInfo {
>         hasNextPage
>       }
>     }
>   }
> }
> ```

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example),
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository and checkout `v0.1.0` as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```


## Add Annotations To Schema

Ordering can be defined on any comparable field of ent by annotating it with `entgql.Annotation`.
Note that the given `OrderField` name must match its enum value in GraphQL schema (see
[next section](#define-ordering-types-in-graphql-schema) below).

```go
func (Todo) Fields() []ent.Field {
    return []ent.Field{
		field.Text("text").
			NotEmpty().
			Annotations(
				entgql.OrderField("TEXT"),
			),
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS").
			Annotations(
				entgql.OrderField("STATUS"),
			),
		field.Int("priority").
			Default(0).
			Annotations(
				entgql.OrderField("PRIORITY"),
			),
    }
}
```

## Define Types In GraphQL Schema

Next, we need to define the ordering types along with the [Relay Connection Types](https://relay.dev/graphql/connections.htm#sec-Connection-Types)
in the GraphQL schema:

```graphql
# Define a Relay Cursor type:
# https://relay.dev/graphql/connections.htm#sec-Cursor
scalar Cursor

type PageInfo {
    hasNextPage: Boolean!
    hasPreviousPage: Boolean!
    startCursor: Cursor
    endCursor: Cursor
}

type TodoConnection {
    totalCount: Int!
    pageInfo: PageInfo!
    edges: [TodoEdge]
}

type TodoEdge {
    node: Todo
    cursor: Cursor!
}

# These enums are matched the entgql annotations in the ent/schema.
enum TodoOrderField {
    CREATED_AT
    PRIORITY
    STATUS
    TEXT
}

enum OrderDirection {
    ASC
    DESC
}

input TodoOrder {
    direction: OrderDirection!
    field: TodoOrderField
}
```

Note that the naming must take the form of `<T>OrderField` / `<T>Order` for `autobind`ing to the generated ent types.
Alternatively [@goModel](https://gqlgen.com/config/#inline-config-with-directives) directive can be used for manual type binding.

## Add Pagination Support For Query

```graphql
type Query {
    todos(
        after: Cursor
        first: Int
        before: Cursor
        last: Int
        orderBy: TodoOrder
    ): TodoConnection
}
```
That's all for the GraphQL schema changes, let's run `gqlgen` code generation.

## Update The GraphQL Resolver

After changing our Ent and GraphQL schemas, we're ready to run the codegen and use the `Paginate` API:

```console
go generate ./...
```

Head over to the `Todos` resolver and update it to pass `orderBy` argument to `.Paginate()` call:

```go
func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}
```

## Pagination Usage

Now, we're ready to test our new GraphQL resolvers. Let's start with creating a few todo items by running this
query multiple times (changing variables is optional):

```graphql
mutation CreateTodo($todo: TodoInput!) {
    createTodo(todo: $todo) {
        id
        text
        createdAt
        priority
        parent {
            id
        }
    }
}

# Query Variables: { "todo": { "text": "Create GraphQL Example", "status": "IN_PROGRESS", "priority": 1 } }
# Output: { "data": { "createTodo": { "id": "2", "text": "Create GraphQL Example", "createdAt": "2021-03-10T15:02:18+02:00", "priority": 1, "parent": null } } }
```

Then, we can query our todo list using the pagination API:

```graphql
query {
    todos(first: 3, orderBy: {direction: DESC, field: TEXT}) {
        edges {
            node {
                id
                text
            }
            cursor
        }
    }
}

# Output: { "data": { "todos": { "edges": [ { "node": { "id": "16", "text": "Create GraphQL Example" }, "cursor": "gqFpEKF2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" }, { "node": { "id": "15", "text": "Create GraphQL Example" }, "cursor": "gqFpD6F2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" }, { "node": { "id": "14", "text": "Create GraphQL Example" }, "cursor": "gqFpDqF2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" } ] } } }
```

We can also use the cursor we got in the query above to get all items after that cursor:

```graphql
query {
    todos(first: 3, after:"gqFpEKF2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU", orderBy: {direction: DESC, field: TEXT}) {
        edges {
            node {
                id
                text
            }
            cursor
        }
    }
}

# Output: { "data": { "todos": { "edges": [ { "node": { "id": "15", "text": "Create GraphQL Example" }, "cursor": "gqFpD6F2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" }, { "node": { "id": "14", "text": "Create GraphQL Example" }, "cursor": "gqFpDqF2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" }, { "node": { "id": "13", "text": "Create GraphQL Example" }, "cursor": "gqFpDaF2tkNyZWF0ZSBHcmFwaFFMIEV4YW1wbGU" } ] } } }
```

---

Great! With a few simple changes, our application now supports pagination! Please continue to the next section where we explain how to implement GraphQL field collections and learn how Ent solves
the *"N+1 problem"* in GraphQL resolvers.
