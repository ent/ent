---
id: tutorial-todo-gql-paginate
title: Relay Cursor Connections (Pagination)
sidebar_label: Relay Cursor Connections
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.mdx) by explaining how to implement the
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
server, you can clone the repository as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```


## Add Annotations To Schema

Ordering can be defined on any comparable field of Ent by annotating it with `entgql.Annotation`.
Note that the given `OrderField` name must be uppercase and match its enum value in the GraphQL schema.

```go title="ent/schema/todo.go"
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

## Order By Multiple Fields

By default, the `orderBy` argument only accepts a single `<T>Order` value. To enable sorting by multiple fields, simply
add the `entgql.MultiOrder()` annotation to desired schema.

```go title="ent/schema/todo.go"
func (Todo) Annotations() []schema.Annotation {
    return []schema.Annotation{
        //highlight-next-line
        entgql.MultiOrder(),
    }
}
```

By adding this annotation to the `Todo` schema, the `orderBy` argument will be changed from `TodoOrder` to `[TodoOrder!]`.

## Order By Edge Count

Non-unique edges can be annotated with the `OrderField` annotation to enable sorting nodes based on the count of specific
edge types.

```go title="ent/schema/todo/go"
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Todo.Type).
			Annotations(
				entgql.RelayConnection(),
				// highlight-next-line
				entgql.OrderField("CHILDREN_COUNT"),
			).
			From("parent").
			Unique(),
	}
}
```

:::info
The naming convention for this ordering term is: `UPPER(<edge-name>)_COUNT`. For example, `CHILDREN_COUNT`
or `POSTS_COUNT`.
:::

## Order By Edge Field

Unique edges can be annotated with the `OrderField` annotation to allow sorting nodes by their associated edge fields.
For example, _sorting posts by their author's name_, or _ordering todos based on their parent's priority_. Note that
in order to sort by an edge field, the field must be annotated with `OrderField` within the referenced type.

The naming convention for this ordering term is: `UPPER(<edge-name>)_<edge-field>`. For example, `PARENT_PRIORITY`.

```go title="ent/schema/todo.go"
// Fields returns todo fields.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		// ...
		field.Int("priority").
			Default(0).
			Annotations(
				// highlight-next-line
				entgql.OrderField("PRIORITY"),
			),
    }
}

// Edges returns todo edges.
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Todo.Type).
			From("parent").
			Annotations(
				// highlight-next-line
				entgql.OrderField("PARENT_PRIORITY"),
			).
			Unique(),
    }
}
```

:::info
The naming convention for this ordering term is: `UPPER(<edge-name>)_<edge-field>`. For example, `PARENT_PRIORITY` or
`AUTHOR_NAME`.
:::

## Add Pagination Support For Query

1\. The next step for enabling pagination is to tell Ent that the `Todo` type is a Relay Connection.

```go title="ent/schema/todo.go"
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
        //highlight-next-line
		entgql.RelayConnection(),
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate()),
	}
}
```

2\. Then, run `go generate .` and you'll notice that `ent.resolvers.go` was changed. Head over to the `Todos` resolver
and update it to pass pagination arguments to `.Paginate()`:

```go title="ent.resolvers.go" {2-5}
func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
	return r.client.Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}
```

:::info Relay Connection Configuration

The `entgql.RelayConnection()` function indicates that the node or edge should support pagination.
Hence,the returned result is a Relay connection rather than a list of nodes (`[T!]!` => `<T>Connection!`).

Setting this annotation on schema `T` (reside in ent/schema), enables pagination for this node and therefore, Ent will
generate all Relay types for this schema, such as: `<T>Edge`, `<T>Connection`, and `PageInfo`. For example:

```go
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		entgql.QueryField(),
	}
}
```

Setting this annotation on an edge indicates that the GraphQL field for this edge should support nested pagination
and the returned type is a Relay connection. For example:

```go
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
			edge.To("parent", Todo.Type).
				Unique().
				From("children").
				Annotations(entgql.RelayConnection()),
	}
}
```

The generated GraphQL schema will be:

```diff
-children: [Todo!]!
+children(first: Int, last: Int, after: Cursor, before: Cursor): TodoConnection!
```

:::

## Pagination Usage

Now, we're ready to test our new GraphQL resolvers. Let's start with creating a few todo items by running this
query multiple times (changing variables is optional):

```graphql
mutation CreateTodo($input: CreateTodoInput!) {
    createTodo(input: $input) {
        id
        text
        createdAt
        priority
        parent {
            id
        }
    }
}

# Query Variables: { "input": { "text": "Create GraphQL Example", "status": "IN_PROGRESS", "priority": 1 } }
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

We can also use the cursor we got in the query above to get all items that come after it.

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

Great! With a few simple changes, our application now supports pagination. Please continue to the next section where we
explain how to implement GraphQL field collections and learn how Ent solves the *"N+1 problem"* in GraphQL resolvers.
