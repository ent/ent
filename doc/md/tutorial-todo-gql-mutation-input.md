---
id: tutorial-todo-gql-mutation-input
title: Mutation Inputs
sidebar_label: Mutation Inputs
---

In this section, we continue the [GraphQL example](tutorial-todo-gql.mdx) by explaining how to extend the Ent code
generator using Go templates and generate [input type](https://graphql.org/graphql-js/mutations-and-input-types/)
objects for our GraphQL mutations that can be applied directly on Ent mutations.

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example),
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository and run the program as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```

## Mutation Types

Ent supports generating mutation types. A mutation type can be accepted as an input for GraphQL mutations, and it is
handled and verified by Ent. Let's tell Ent that our GraphQL `Todo` type supports create and update operations:

```go title="ent/schema/todo.go"
func (Todo) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.QueryField(),
		//highlight-next-line
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}
```

Then, run code generation:

```go
go generate .
```

You'll notice that Ent generated for you 2 types: `ent.CreateTodoInput` and `ent.UpdateTodoInput`.

## Mutations

After generating our mutation inputs, we can connect them to the GraphQL mutations:

```graphql title="todo.graphql"
type Mutation {
  createTodo(input: CreateTodoInput!): Todo!
  updateTodo(id: ID!, input: UpdateTodoInput!): Todo!
}
```

Running code generation we'll generate the actual mutations and the only thing left after that is to bind the resolvers
to Ent.
```go
go generate .
```

```go title="todo.resolvers.go"
// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	return r.client.Todo.Create().SetInput(input).Save(ctx)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input ent.UpdateTodoInput) (*ent.Todo, error) {
	return r.client.Todo.UpdateOneID(id).SetInput(input).Save(ctx)
}
```


## Test the `CreateTodo` Resolver

Let's start with creating 2 todo items by executing the `createTodo` mutations twice.

#### Mutation

```graphql
mutation CreateTodo {
   createTodo(input: {text: "Create GraphQL Example", status: IN_PROGRESS, priority: 2}) {
     id
     text
     createdAt
     priority
     parent {
       id
     }
   }
 }
```

#### Output

```json
{
  "data": {
    "createTodo": {
      "id": "1",
      "text": "Create GraphQL Example",
      "createdAt": "2021-04-19T10:49:52+03:00",
      "priority": 2,
      "parent": null
    }
  }
}
```

#### Mutation

```graphql
mutation CreateTodo {
   createTodo(input: {text: "Create Tracing Example", status: IN_PROGRESS, priority: 2}) {
     id
     text
     createdAt
     priority
     parent {
       id
     }
   }
 }
```

#### Output

```json
{
  "data": {
    "createTodo": {
      "id": "2",
      "text": "Create Tracing Example",
      "createdAt": "2021-04-19T10:50:01+03:00",
      "priority": 2,
      "parent": null
    }
  }
}
```

## Test the `UpdateTodo` Resolver

The only thing left is to test the `UpdateTodo` resolver. Let's use it to update the `parent` of the 2nd todo item to `1`.

```graphql
mutation UpdateTodo {
  updateTodo(id: 2, input: {parent: 1}) {
    id
    text
    createdAt
    priority
    parent {
      id
      text
    }
  }
}
```

#### Output

```json
{
  "data": {
    "updateTodo": {
      "id": "2",
      "text": "Create Tracing Example",
      "createdAt": "2021-04-19T10:50:01+03:00",
      "priority": 1,
      "parent": {
        "id": "1",
        "text": "Create GraphQL Example"
      }
    }
  }
}
```