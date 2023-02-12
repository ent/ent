---
id: tutorial-setup
title: Setting Up
sidebar_label: Setting Up
---

This guide is intended for first-time users who want instructions on how to set up an Ent project from scratch.
Before we get started, make sure you have the following prerequisites installed on your machine.

## Prerequisites

- [Go](https://go.dev/doc/install)
- [Docker](https://docs.docker.com/get-docker) (optional)

After installing these dependencies, create a directory for the project and initialize a Go module:

```console
mkdir todo
cd $_
go mod init todo
```

## Installation

Run the following Go commands to install Ent, and tell it to initialize the project structure along with a `Todo` schema.

```console
go get -d entgo.io/ent/cmd/ent
```

```console
go run -mod=mod entgo.io/ent/cmd/ent new Todo
```

After installing Ent and running `ent init`, your project directory should look like this:

```console
.
├── ent
│   ├── generate.go
│   └── schema
│       └── todo.go
├── go.mod
└── go.sum
```

The `ent` directory holds the generated assets (see the next section), and the `ent/schema` directory contains your
entity schemas.

## Code Generation

When we ran `ent new Todo` above, a schema named `Todo`  was created in the `todo.go` file under the`todo/ent/schema/` directory:

```go
package schema

import "entgo.io/ent"

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return nil
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
```

As you can see, initially, the schema has no fields or edges defined. Let's run the command for generating assets to interact with
the `Todo` entity:

```console
go generate ./ent
```

## Create a Test Case

Running `go generate ./ent` invoked Ent's automatic code generation tool, which uses the schemas we define in our `schema` package to generate the actual Go code which we will now use to interact with a database. At this stage, you can find under `./ent/client.go`, client code that is capable of  querying and mutating the `Todo` entities. Let's create a
[testable example](https://go.dev/blog/examples) to use this. We'll use [SQLite](https://github.com/mattn/go-sqlite3)
in this test-case for testing Ent.

```console
go get github.com/mattn/go-sqlite3
touch example_test.go
```

Paste the following code in `example_test.go` that instantiates an `ent.Client` and automatically creates all schema resources
in the database (tables, columns, etc).

```go
package todo

import (
	"context"
	"log"
	"todo/ent"

	"entgo.io/ent/dialect"
	_ "github.com/mattn/go-sqlite3"
)

func Example_Todo() {
	// Create an ent.Client with in-memory SQLite database.
	client, err := ent.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run the automatic migration tool to create all schema resources.
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// Output:
}
```

Then, run `go test` to verify that everything works as expected.

```console
go test
```

After setting up our project, we're ready to create our Todo list.
