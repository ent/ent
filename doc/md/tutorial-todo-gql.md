---
id: tutorial-todo-gql
title: Introduction
sidebar_label: Introduction
---

In this section, we will learn how to connect Ent to [GraphQL](https://graphql.org). If you're not familiar with GraphQL,
it's recommended to go over its [introduction guide](https://graphql.org/learn/) before going over this tutorial.

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example), 
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository and checkout `v0.1.0` as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example
git checkout v0.1.0
go run ./cmd/todo/
```

## Basic Skeleton

[gqlgen](https://gqlgen.com/) is a framework for easily generating GraphQL servers in Go. In this tutorial, we will review Ent's official integration with it.

This tutorial begins where the previous one ended (with a working Todo-list schema). We start by creating a simple GraphQL schema for our todo list, then install the [99designs/gqlgen](https://github.com/99designs/gqlgen)
package and configure it. Let's create a file named `todo.graphql` and paste the following:

```graphql
# Maps a Time GraphQL scalar to a Go time.Time struct.
scalar Time

# Define an enumeration type and map it later to Ent enum (Go type).
# https://graphql.org/learn/schema/#enumeration-types
enum Status {
    IN_PROGRESS
    COMPLETED
}

# Define an object type and map it later to the generated Ent model.
# https://graphql.org/learn/schema/#object-types-and-fields
type Todo {
    id: ID!
    createdAt: Time
    status: Status!
    priority: Int!
    text: String!
    parent: Todo
    children: [Todo!]
}

# Define an input type for the mutation below.
# https://graphql.org/learn/schema/#input-types
input TodoInput {
    status: Status! = IN_PROGRESS
    priority: Int
    text: String!
    parent: ID
}

# Define a mutation for creating todos.
# https://graphql.org/learn/queries/#mutations
type Mutation {
    createTodo(todo: TodoInput!): Todo!
}

# Define a query for getting all todos.
type Query {
    todos: [Todo!]
}
```

Install [99designs/gqlgen](https://github.com/99designs/gqlgen):

```console
go get github.com/99designs/gqlgen
```

The gqlgen package can be configured using a `gqlgen.yml` file that it automatically loads from the current directory.
Let's add this file. Follow the comments in this file to understand what each config directive means:

```yaml
# schema tells gqlgen where the GraphQL schema is located.
schema:
  - todo.graphql

# resolver reports where the resolver implementations go.
resolver:
  layout: follow-schema
  dir: .

# gqlgen will search for any type names in the schema in these go packages
# if they match it will use them, otherwise it will generate them.

# autobind tells gqlgen to search for any type names in the GraphQL schema in the
# provided Go package. If they match it will use them, otherwise it will generate new ones.
autobind:
  - todo/ent

# This section declares type mapping between the GraphQL and Go type systems.
models:
  # Defines the ID field as Go 'int'.
  ID:
    model:
      - github.com/99designs/gqlgen/graphql.IntID
  # Map the Status type that was defined in the schema
  Status:
    model:
      - todo/ent/todo.Status
```

Now, we're ready to run gqlgen code generation. Execute this command from the root of the project:

```console
go run github.com/99designs/gqlgen
```

The command above will execute the gqlgen code-generator, and if that finished successfully, your project directory
should look like this:

```console
➜ tree -L 1   
.
├── ent
├── example_test.go
├── generated.go
├── go.mod
├── go.sum
├── gqlgen.yml
├── models_gen.go
├── resolver.go
├── todo.graphql
└── todo.resolvers.go

1 directories, 9 files
```

## Connect Ent to GQL

After the gqlgen assets were generated, we're ready to connect Ent to gqlgen and start running our server.
This section contains 5 steps, follow them carefully :). 

**1\.** Install the GraphQL extension for Ent

```console
go get entgo.io/contrib/entgql
```

**2\.** Create a new Go file named `ent/entc.go`, and paste the following content:

```go
// +build ignore

package main

import (
    "log"

    "entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
    "entgo.io/contrib/entgql"
)

func main() {
	ex, err := entgql.NewExtension()
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	opts := []entc.Option{
		entc.Extensions(ex),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

**3\.** Edit the `ent/generate.go` file to execute the `ent/entc.go` file:

```go
package ent

//go:generate go run entc.go
```

Note that `ent/entc.go` is ignored using a build tag, and it's executed by the go generate command through the
`generate.go` file.

**4\.** In order to execute `gqlgen` through `go generate`, we create a new `generate.go` file (in the root
of the project) with the following:

```go
package todo

//go:generate go run github.com/99designs/gqlgen
```

Now, running `go generate ./...` from the root of the project, triggers both Ent and gqlgen code generation.

```console
go generate ./...
```

**5\.** `gqlgen` allows changing the generated `Resolver` and add additional dependencies to it. Let's add
the `ent.Client` as a dependency by pasting the following in `resolver.go`:

```go
package todo

import (
	"todo/ent"
	
	"github.com/99designs/gqlgen/graphql"
)

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	return NewExecutableSchema(Config{
		Resolvers: &Resolver{client},
	})
}
```

## Run the server

We create a new directory `cmd/todo` and a `main.go` file with the following code to create the GraphQL server:

```go
package main

import (
	"context"
	"log"
	"net/http"

	"todo/ent"
	"todo/ent/migrate"

	"entgo.io/ent/dialect"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Create ent.Client and run the schema migration.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal("opening ent client", err)
	}
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		log.Fatal("opening ent client", err)
	}

	// Configure the server and start listening on :8081.
	srv := handler.NewDefaultServer(NewSchema(client))
	http.Handle("/",
		playground.Handler("Todo", "/query"),
	)
	http.Handle("/query", srv)
	log.Println("listening on :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatal("http server terminated", err)
	}
}

```

Run the server using the command below, and open [localhost:8081](http://localhost:8081):

```console
go run ./cmd/todo
```

You should see the interactive playground:

![tutorial-todo-playground](https://entgo.io/images/assets/tutorial-gql-playground.png)

If you're having troubles with getting the playground to run, go to [first section](#clone-the-code-optional) and clone the
example repository.

## Query Todos

If we try to query our todo list, we'll get an error as the resolver method is not yet implemented. 
Let's implement the resolver by replacing the `Todos` implementation in the query resolver:

```diff
func (r *queryResolver) Todos(ctx context.Context) ([]*ent.Todo, error) {
-	panic(fmt.Errorf("not implemented"))
+	return r.client.Todo.Query().All(ctx)
}
```

Then, running this GraphQL query should return an empty todo list:

```graphql
query AllTodos {
    todos {
        id
    }
}

# Output: { "data": { "todos": [] } }
```

## Create a Todo

Same as before, if we try to create a todo item in GraphQL, we'll get an error as the resolver is not yet implemented.
Let's implement the resolver by changing the `CreateTodo` implementation in the mutation resolver:

```go
func (r *mutationResolver) CreateTodo(ctx context.Context, todo TodoInput) (*ent.Todo, error) {
	return r.client.Todo.Create().
		SetText(todo.Text).
		SetStatus(todo.Status).
		SetNillablePriority(todo.Priority). // Set the "priority" field if provided.
		SetNillableParentID(todo.Parent).   // Set the "parent_id" field if provided.
		Save(ctx)
}
```

Now, creating a todo item should work:

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

If you're having troubles with getting this example to work, go to [first section](#clone-the-code-optional) and clone the
example repository.

---

Please continue to the next section where we explain how to implement the
[Relay Node Interface](https://relay.dev/graphql/objectidentification.htm) and learn how Ent automatically supports this.