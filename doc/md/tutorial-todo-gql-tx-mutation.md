---
id: tutorial-todo-gql-tx-mutation
title: Transactional Mutations
sidebar_label: Transactional Mutations
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.md) by explaining how to set our GraphQL mutations
to be transactional. That means, to automatically wrap our GraphQL mutations with a database transaction and either
commit at the end, or rollback the transaction in case of a GraphQL error.

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example),
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository and checkout `v0.1.0` as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```

## Usage

The GraphQL extensions provides a handler named `entgql.Transactioner` that executes each GraphQL mutation in a
transaction. The injected client for the resolver is a [transactional `ent.Client`](transactions.md#transactional-client).
Hence, GraphQL resolvers that uses `ent.Client` won't need to be changed. In order to add it to our todo list application
we follow these steps:

1\. Edit the `cmd/todo/main.go` and add to the GraphQL server initialization the `entgql.Transactioner` handler as
follows:

```diff
srv := handler.NewDefaultServer(todo.NewSchema(client))
+srv.Use(entgql.Transactioner{TxOpener: client})
```

2\. Then, in the GraphQL mutations, use the client from context as follows:
```diff
func (mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
+	client := ent.FromContext(ctx)
+	return client.Todo.
-	return r.client.Todo.
		Create().
		SetText(todo.Text).
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority). // Set the "priority" field if provided.
		SetNillableParentID(todo.Parent).   // Set the "parent_id" field if provided.
		Save(ctx)
}
```
