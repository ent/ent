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
  updateTodo(id: 2, input: {parentID: 1}) {
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

## Create edges with mutations

To create the edges of a node in the same mutation, you can extend the GQL mutation input with the edge fields:

```graphql title="extended.graphql"
extend input CreateTodoInput {
  createChildren: [CreateTodoInput!]
}
```

Next, run code generation again:
```go
go generate .
```

GQLGen will generate the resolver for the `createChildren` field, allowing you to use it in your resolver:

```go title="extended.resolvers.go"
// CreateChildren is the resolver for the createChildren field.
func (r *createTodoInputResolver) CreateChildren(ctx context.Context, obj *ent.CreateTodoInput, data []*ent.CreateTodoInput) error {
	panic(fmt.Errorf("not implemented: CreateChildren - createChildren"))
}
```

Now, we need to implement the logic to create the children:

```go title="extended.resolvers.go"
// CreateChildren is the resolver for the createChildren field.
func (r *createTodoInputResolver) CreateChildren(ctx context.Context, obj *ent.CreateTodoInput, data []*ent.CreateTodoInput) error {
	// highlight-start
	// NOTE: We need to use the Ent client from the context.
	// To ensure we create all of the children in the same transaction.
	// See: Transactional Mutations for more information.
	c := ent.FromContext(ctx)
	// highlight-end
	builders := make([]*ent.TodoCreate, len(data))
	for i := range data {
		builders[i] = c.Todo.Create().SetInput(*data[i])
	}
	todos, err := c.Todo.CreateBulk(builders...).Save(ctx)
	if err != nil {
		return err
	}
	ids := make([]int, len(todos))
	for i := range todos {
		ids[i] = todos[i].ID
	}
	obj.ChildIDs = append(obj.ChildIDs, ids...)
	return nil
}
```

Change the following lines to use the transactional client:

```go title="todo.resolvers.go"
// CreateTodo is the resolver for the createTodo field.
func (r *mutationResolver) CreateTodo(ctx context.Context, input ent.CreateTodoInput) (*ent.Todo, error) {
	// highlight-next-line
	return ent.FromContext(ctx).Todo.Create().SetInput(input).Save(ctx)
}

// UpdateTodo is the resolver for the updateTodo field.
func (r *mutationResolver) UpdateTodo(ctx context.Context, id int, input ent.UpdateTodoInput) (*ent.Todo, error) {
	// highlight-next-line
	return ent.FromContext(ctx).Todo.UpdateOneID(id).SetInput(input).Save(ctx)
}
```

Test the mutation with the children:

**Mutation**
```graphql
mutation {
  createTodo(input: {
    text: "parent", status:IN_PROGRESS,
    createChildren: [
      { text: "children1", status: IN_PROGRESS },
      { text: "children2", status: COMPLETED }
    ]
  }) {
    id
    text
    children {
      id
      text
      status
    }
  }
}
```

**Output**
```json
{
  "data": {
    "createTodo": {
      "id": "3",
      "text": "parent",
      "children": [
        {
          "id": "1",
          "text": "children1",
          "status": "IN_PROGRESS"
        },
        {
          "id": "2",
          "text": "children2",
          "status": "COMPLETED"
        }
      ]
    }
  }
}
```

If you enable the debug Client, you'll see that the children are created in the same transaction:
```log
2022/12/14 00:27:41 driver.Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312): started
2022/12/14 00:27:41 Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312).Query: query=INSERT INTO `todos` (`created_at`, `priority`, `status`, `text`) VALUES (?, ?, ?, ?), (?, ?, ?, ?) RETURNING `id` args=[2022-12-14 00:27:41.046344 +0700 +07 m=+5.283557793 0 IN_PROGRESS children1 2022-12-14 00:27:41.046345 +0700 +07 m=+5.283558626 0 COMPLETED children2]
2022/12/14 00:27:41 Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312).Query: query=INSERT INTO `todos` (`text`, `created_at`, `status`, `priority`) VALUES (?, ?, ?, ?) RETURNING `id` args=[parent 2022-12-14 00:27:41.047455 +0700 +07 m=+5.284669251 IN_PROGRESS 0]
2022/12/14 00:27:41 Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312).Exec: query=UPDATE `todos` SET `todo_parent` = ? WHERE `id` IN (?, ?) AND `todo_parent` IS NULL args=[3 1 2]
2022/12/14 00:27:41 Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312).Query: query=SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` WHERE `todo_parent` = ? args=[3]
2022/12/14 00:27:41 Tx(7e04b00b-7941-41c5-9aee-41c8c2d85312): committed
```
