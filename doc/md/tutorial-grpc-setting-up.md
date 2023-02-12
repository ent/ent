---
id: grpc-setting-up
title: Setting Up
sidebar_label: Setting Up 
---

Let's start by initializing a new Go module for our project:

```console
mkdir ent-grpc-example
cd ent-grpc-example
go mod init ent-grpc-example
```

Next, we use `go run` to invoke the ent code generator to initialize a schema:

```console
go run -mod=mod entgo.io/ent/cmd/ent new User
```

Our directory should now look like:

```console
.
├── ent
│   ├── generate.go
│   └── schema
│       └── user.go
├── go.mod
└── go.sum
```

Next, let's add the `entproto` package to our project:

```console
go get -u entgo.io/contrib/entproto
```

Next, we will define the schema for the `User` entity. Open `ent/schema/user.go` and edit:

```go title="ent/schema/user.go"
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique(),
		field.String("email_address").
			Unique(),
	}
}
```

In this step, we added two unique fields to our `User` entity: `name` and `email_address`. The `ent.Schema` is just the definition of the schema. To create usable production code from it we need to run Ent's code generation tool on it. Run:

```console
go generate ./...
```

Notice that new files were created from our schema definition:

```console
├── ent
│   ├── client.go
│   ├── config.go
// .... many more
│   ├── user
│   ├── user.go
│   ├── user_create.go
│   ├── user_delete.go
│   ├── user_query.go
│   └── user_update.go
├── go.mod
└── go.sum
```

At this point, we can open a connection to a database, run a migration to create the `users` table, and start reading and writing data to it. This is covered on the [Setup Tutorial](tutorial-setup.md), so let's cut to the chase and learn about generating Protobuf definitions and gRPC servers from our schema.
