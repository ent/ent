---
id: tutorial-todo-gql-tx-mutation
title: Transactional Mutations
sidebar_label: Transactional Mutations
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.mdx) by explaining how to set our GraphQL mutations
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

```diff title="cmd/todo/main.go"
srv := handler.NewDefaultServer(todo.NewSchema(client))
+srv.Use(entgql.Transactioner{TxOpener: client})
```

2\. Then, in the GraphQL mutations, use the client from context as follows:
```diff title="todo.resolvers.go"
}
+func (mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
+	client := ent.FromContext(ctx)
+	return client.Todo.Create().SetInput(input).Save(ctx)
-func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
-	return r.client.Todo.Create().SetInput(input).Save(ctx)
}
```

## Isolation Levels

If you'd like to tweak the transaction's isolation level, you can do so by implementing your own `TxOpener`. For example:

```go title="cmd/todo/main.go"
srv.Use(entgql.Transactioner{
	TxOpener: entgql.TxOpenerFunc(func(ctx context.Context) (context.Context, driver.Tx, error) {
		tx, err := client.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelRepeatableRead})
		if err != nil {
			return nil, nil, err
		}
		ctx = ent.NewTxContext(ctx, tx)
		ctx = ent.NewContext(ctx, tx.Client())
		return ctx, tx, nil
	}),
})
```

## Skip Operations

By default, `entgql.Transactioner` wraps all mutations within a transaction. However, there are mutations or operations
that don't require database access or need special handling. In these cases, you can instruct `entgql.Transactioner` to
skip the transaction by setting a custom `SkipTxFunc` function or using one of the built-in ones.

```go title="cmd/todo/main.go" {4,10,16-18}
srv.Use(entgql.Transactioner{
	TxOpener: client,
	// Skip the given operation names from running under a transaction.
	SkipTxFunc: entgql.SkipOperations("operation1", "operation2"),
})

srv.Use(entgql.Transactioner{
	TxOpener: client,
	// Skip if the operation has a mutation field with the given names.
	SkipTxFunc: entgql.SkipIfHasFields("field1", "field2"),
})

srv.Use(entgql.Transactioner{
	TxOpener: client,
	// Custom skip function.
	SkipTxFunc: func(*ast.OperationDefinition) bool {
	    // ...
    },
})
```

---

Great! With a few lines of code, our application now supports automatic transactional mutations. Please continue to the
next section where we explain how to extend the Ent code generator and generate [GraphQL input types](https://graphql.org/graphql-js/mutations-and-input-types/)
for our GraphQL mutations.