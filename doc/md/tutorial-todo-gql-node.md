---
id: tutorial-todo-gql-node
title: Relay Node Interface
sidebar_label: Relay Node Interface
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.mdx) by explaining how to implement the
[Relay Node Interface](https://relay.dev/graphql/objectidentification.htm). If you're not familiar with the
Node interface, read the following paragraphs that were taken from [relay.dev](https://relay.dev/graphql/objectidentification.htm#sel-DABDDBAADLA0Cl0c):

> To provide options for GraphQL clients to elegantly handle for caching and data refetching GraphQL servers need to expose object identifiers in a standardized way. In the query, the schema should provide a standard mechanism for asking for an object by ID. In the response, the schema provides a standard way of providing these IDs.
>
> We refer to objects with identifiers as “nodes”. An example of both of those is the following query:
>
>  ```graphql
>   {
>       node(id: "4") {
>           id
>          ... on User {
>               name
>           }
>       }
>   }
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

## Implementation

Ent supports the Node interface through its GraphQL integration. By following a few simple steps you can add support
for it in your application. We start by telling `gqlgen` that Ent provides the `Node` interface by editing the
`gqlgen.yaml` file as follows:

```diff title="gqlgen.yml" {7-9}
# This section declares type mapping between the GraphQL and Go type systems.
models:
  # Defines the ID field as Go 'int'.
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.IntID
  Node:
    model:
      - todo/ent.Noder
```

To apply these changes, we rerun the code generation:

```console
go generate .
```

Like before, we need to implement the GraphQL resolvers in `ent.resolvers.go`. With a one-liner change, we can
implement those by replacing the generated `gqlgen` code with the following:

```diff title="ent.resolvers.go"
func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
-	panic(fmt.Errorf("not implemented: Node - node"))
+	return r.client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
-	panic(fmt.Errorf("not implemented: Nodes - nodes"))
+	return r.client.Noders(ctx, ids)
}
```

## Query Nodes

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

# Query Variables: { "input": { "text":"Create GraphQL Example", "status": "IN_PROGRESS", "priority": 1 } }
# Output: { "data": { "createTodo": { "id": "2", "text": "Create GraphQL Example", "createdAt": "2021-03-10T15:02:18+02:00", "priority": 1, "parent": null } } }
```

Running the **Node** API on one of the todo items will return:

````graphql
query {
  node(id: 1) {
    id
    ... on Todo {
      text
    }
  }
}

# Output: { "data": { "node": { "id": "1", "text": "Create GraphQL Example" } } }
````

Running the **Nodes** API on one of the todo items will return:

```graphql
query {
  nodes(ids: [1, 2]) {
    id
    ... on Todo {
      text
    }
  }
}

# Output: { "data": { "nodes": [ { "id": "1", "text": "Create GraphQL Example" }, { "id": "2", "text": "Create Tracing Example" } ] } }
```

---

Well done! As you can see, by changing a few lines of code our application now implements the Relay Node Interface. 
In the next section, we will show how to implement the Relay Cursor Connections spec using Ent, which is very useful 
if we want our application to support slicing and pagination of query results.
