---
id: doc1
title: Quick Introduction
sidebar_label: Quick Introduction
---

`ent` is a simple, yet powerful ORM framework for Go built with the following principles:
- Defining your schema as code.
- Static typing first based on code generation.
- Make the work with graph-like data in Go easier.

## Installation

```console
$ go get github.com/facebookincubator/ent/entc/cmd/entc
```

After installing `entc` (the code generator for `ent`), you should have it in your `PATH`.

## Create Your First Schema
Go to the root directory of your project, and run:

```console
$ entc init User
```
The command above will generate the schema for `User` under `<project>/ent/schema/` directory:

```go
// <project>/ent/schema/user.go

package schema

import "github.com/facebookincubator/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

```

Let's add 2 fields to the `User` schema, and then run `entc generate`:

```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
```

Running `entc generate` from the root directory of the project:

```go
$ entc generate ./ent/schema
```

Will produce the following files:
```
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── example_test.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```


## Create Your First Entity

First thing we need to do, is creating a new `ent.Client`. For the example purpose,
we will use SQLite3.

```go
package main

import (
	"log"

	"<project>/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	drv := dialect.Driver(db)
	if testing.Verbose() {
		drv = dialect.Debug(drv)
	}
	client := ent.NewClient(ent.Driver(db))
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

Now, we're ready to create our user. Let's call this function `Do` for the sake of the example:
```go
func Do(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save()
	if err != nil {
		return nil, fmt.Error("failed creating user: %v", err)
	}
	log.Println("user was created: %v", u)
	return u, nil
}
```

## Query Your Entities

`entc` generates a package for each entity schema that contains its predicates, default values, validators
and information about storage elements (like, column names, primary keys, etc).

```go
package main

import (
	"log"

	"<project>/ent"
	"<project>/ent/user"
)

func Query(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Error("failed querying user: %v", err)
	}
	log.Println("user: %v", u)
	return u, nil
}
```


## Placeholder 1

Placeholder 1

## Placeholder 2

Placeholder 2
