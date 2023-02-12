---
title: Generate a fully-working Go gRPC server in two minutes with Ent
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---
![ent + gRPC](https://entgo.io/images/assets/ent-grpc.jpg)
## Introduction
Having entity schemas defined in a central, language-neutral format
has [many benefits](https://rotemtam.com/2019/06/28/the-statically-typed-org/) as the scale of software engineering organizations increase. To do this, many organizations use [Protocol Buffers](https://developers.google.com/protocol-buffers) as their [interface definition language](https://en.wikipedia.org/wiki/Interface_description_language) (IDL). In addition, gRPC,
a Protobuf-based RPC framework modeled after Google's internal [Stubby](https://grpc.io/blog/principles/#motivation) is becoming increasingly popular due to its efficiency and code-generation capabilities.

Being an IDL, gRPC does not prescribe any specific guidelines on implementing the data access layer so implementations vary greatly. Ent is a natural candidate for building the data access layer in any Go application and so there is great potential in integrating the two technologies together.

Today we announce an experimental version of `entproto`, a Go package, and a command-line tool to add Protobuf and gRPC support for ent users. With `entproto`, developers can set up a fully working CRUD gRPC server in a few minutes. In this post, we will show exactly how to do just that.

## Setting Up
The final version of this tutorial is available on [GitHub](https://github.com/rotemtam/ent-grpc-example), you can clone it if you prefer following along that way.

Let's start by initializing a new Go module for our project:

```console
mkdir ent-grpc-example
cd ent-grpc-example
go mod init ent-grpc-example
```

Next we use `go run` to invoke the ent code generator to initialize a schema:

```console
go run -mod=mod entgo.io/ent/cmd/ent new User
```

Our directory should now look like:

```console
.
├── ent
│   ├── generate.go
│   └── schema
│       └── user.go
├── go.mod
└── go.sum
```

Next, let's add the `entproto` package to our project:

```console
go get -u entgo.io/contrib/entproto
```

Next, we will define the schema for the `User` entity. Open `ent/schema/user.go` and edit:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
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

In this step, we added two unique fields to our `User` entity: `name` and `email_address`. The `ent.Schema` is just the definition of the schema, to create usable production code from it we need to run Ent's code generation tool on it. Run:

```console
go generate ./...
```

Notice the a bunch of new files were created from our schema definition now:

```console
├── ent
│   ├── client.go
│   ├── config.go
// .... many more
│   ├── user
│   ├── user.go
│   ├── user_create.go
│   ├── user_delete.go
│   ├── user_query.go
│   └── user_update.go
├── go.mod
└── go.sum
```

At this point, we can open a connection to a database, run a migration to create the `users` table, and start reading and writing data to it. This is covered on the [Setup Tutorial](https://entgo.io/docs/tutorial-setup/), so let's cut to the chase and learn about generating Protobuf definitions and gRPC servers from our schema.

## Generating Go Protobufs with  `entproto`

As ent and Protobuf schemas are not identical, we must supply some annotations on our schema to help `entproto` figure out exactly how to generate Protobuf definitions (called "Messages" in protobuf lingo).

The first thing we need to do is to add an `entproto.Message()` annotation. This is our opt-in to Protobuf schema generation, we don't necessarily want to generate proto messages or gRPC service definitions from *all* of our schema entities, and this annotation gives us that control. To add it, append to `ent/schema/user.go`:

```go
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}
```

Next, we need to annotate each field and assign it a field number. Recall that when [defining a protobuf message type](https://developers.google.com/protocol-buffers/docs/proto3#simple), each field must be assigned a unique number.  To do that, we add an `entproto.Field` annotation on each field. Update the `Fields` in `ent/schema/user.go`:

```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Unique().
			Annotations(
				entproto.Field(2),
			),
		field.String("email_address").
			Unique().
			Annotations(
				entproto.Field(3),
			),
	}
}
```

Notice that we did not start our field numbers from 1, this is because `ent` implicitly creates the `ID` field for the entity, and that field is automatically assigned the number 1.  We can now generate our protobuf message type definitions. To do that, we will add to `ent/generate.go` a `go:generate` directive that invokes the `entproto` command-line tool. It should now look like this:

```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
//go:generate go run -mod=mod entgo.io/contrib/entproto/cmd/entproto -path ./schema
```

Let's re-generate our code:

```console
go generate ./...
```

Observe that a new directory was created which will contain all protobuf related generated code: `ent/proto`. It now contains:

```console
ent/proto
└── entpb
    ├── entpb.proto
    └── generate.go
```

Two files were created. Let's look at their contents:

```protobuf
// Code generated by entproto. DO NOT EDIT.
syntax = "proto3";

package entpb;

option go_package = "ent-grpc-example/ent/proto/entpb";

message User {
  int32 id = 1;

  string user_name = 2;

  string email_address = 3;
}
```

Nice! A new `.proto` file containing a message type definition that maps to our `User` schema was created!

```go
package entpb
//go:generate protoc -I=.. --go_out=.. --go-grpc_out=.. --go_opt=paths=source_relative --go-grpc_opt=paths=source_relative --entgrpc_out=.. --entgrpc_opt=paths=source_relative,schema_path=../../schema entpb/entpb.proto
```

A new `generate.go` file was created with an invocation to `protoc`, the protobuf code generator instructing it how to generate Go code from our `.proto` file. For this command to work, we must first install `protoc` as well as 3 protobuf plugins: `protoc-gen-go` (which generates Go Protobuf structs), `protoc-gen-go-grpc` (which generates Go gRPC service interfaces and clients), and `protoc-gen-entgrpc` (which generates an implementation of the service interface). If you do not have these installed, please follow these directions:

- [protoc installation](https://grpc.io/docs/protoc-installation/)
- [protoc-gen-go + protoc-gen-go-grpc installation](https://grpc.io/docs/languages/go/quickstart/)
- Run `go get -u entgo.io/contrib/entproto/cmd/protoc-gen-entgrpc` to install `protoc-gen-entgrpc`

After installing these dependencies, we can re-run code-generation:

```console
go generate ./...
```

Observe that a new file named `ent/proto/entpb/entpb.pb.go` was created which contains the generated Go structs for our entities.

Let's write a test that uses it to make sure everything is wired correctly. Create a new file named `pb_test.go` and write:

```go
package main

import (
	"testing"

	"ent-grpc-example/ent/proto/entpb"
)

func TestUserProto(t *testing.T) {
	user := entpb.User{
		Name:     "rotemtam",
		EmailAddress: "rotemtam@example.com",
	}
	if user.GetName() != "rotemtam" {
		t.Fatal("expected user name to be rotemtam")
	}
	if user.GetEmailAddress() != "rotemtam@example.com" {
		t.Fatal("expected email address to be rotemtam@example.com")
	}
}
```

To run it:

```console
go get -u./... # install deps of the generated package
go test ./...
```

Hooray! The test passes. We have successfully generated working Go Protobuf structs from our Ent schema. Next, let's see how to automatically generate a working CRUD gRPC *server* from our schema.

## Generating a Fully Working gRPC Server from our Schema

Having Protobuf structs generated from our `ent.Schema` can be useful, but what we're really interested in is getting an actual server that can create, read, update, and delete entities from an actual database. To do that, we need to update just one line of code! When we annotate a schema with `entproto.Service`, we tell the `entproto` code-gen that we are interested in generating a gRPC service definition, from the `protoc-gen-entgrpc` will read our definition and generate a service implementation. Edit `ent/schema/user.go` and modify the schema's `Annotations`:

```diff
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
+		entproto.Service(), // <-- add this
	}
}
```

Now re-run code-generation:

```console
go generate ./...
```

Observe some interesting changes in `ent/proto/entpb`:

```console
ent/proto/entpb
├── entpb.pb.go
├── entpb.proto
├── entpb_grpc.pb.go
├── entpb_user_service.go
└── generate.go
```

First, `entproto` added a service definition to `entpb.proto`:

```protobuf
service UserService {
  rpc Create ( CreateUserRequest ) returns ( User );

  rpc Get ( GetUserRequest ) returns ( User );

  rpc Update ( UpdateUserRequest ) returns ( User );

  rpc Delete ( DeleteUserRequest ) returns ( google.protobuf.Empty );
}
```

In addition, two new files were created. The first, `ent_grpc.pb.go`, contains the gRPC client stub and the interface definition. If you open the file, you will find in it (among many other things):

```go
// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
}
```

The second file, `entpub_user_service.go` contains a generated implementation for this interface. For example, an implementation for the `Get` method:

```go
// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	get, err := svc.client.User.Get(ctx, int(req.GetId()))
	switch {
	case err == nil:
		return toProtoUser(get), nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}
```

Not bad! Next, let's create a gRPC server that can serve requests to our service.

### Creating the Server

Create a new file `cmd/server/main.go` and write:

```go
package main

import (
	"context"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"ent-grpc-example/ent"
	"ent-grpc-example/ent/proto/entpb"
	"google.golang.org/grpc"
)

func main() {
	// Initialize an ent client.
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	// Initialize the generated User service.
	svc := entpb.NewUserService(client)

	// Create a new gRPC server (you can wire multiple services to a single server).
	server := grpc.NewServer()

	// Register the User service with the server.
	entpb.RegisterUserServiceServer(server, svc)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
```

Notice that we added an import of `github.com/mattn/go-sqlite3`, so we need to add it to our module:

```console
go get -u github.com/mattn/go-sqlite3
```

Next, let's run the server, while we write a client that will communicate with it:

```console
go run -mod=mod ./cmd/server
```

### Creating the Client

Let's create a simple client that will make some calls to our server. Create a new file named `cmd/client/main.go` and write:

```go
package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"ent-grpc-example/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Open a connection to the server.
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	// Create a User service Client on the connection.
	client := entpb.NewUserServiceClient(conn)

	// Ask the server to create a random User.
	ctx := context.Background()
	user := randomUser()
	created, err := client.Create(ctx, &entpb.CreateUserRequest{
		User: user,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("user created with id: %d", created.Id)

	// On a separate RPC invocation, retrieve the user we saved previously.
	get, err := client.Get(ctx, &entpb.GetUserRequest{
		Id: created.Id,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed retrieving user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("retrieved user with id=%d: %v", get.Id, get)
}

func randomUser() *entpb.User {
	return &entpb.User{
		Name:         fmt.Sprintf("user_%d", rand.Int()),
		EmailAddress: fmt.Sprintf("user_%d@example.com", rand.Int()),
	}
}
```

Our client creates a connection to port 5000, where our server is listening, then issues a `Create` request to create a new user, and then issues a second `Get` request to retrieve it from the database. Let's run our client code:

```console
go run ./cmd/client
```

Observe the output:

```console
2021/03/18 10:42:58 user created with id: 1
2021/03/18 10:42:58 retrieved user with id=1: id:1 name:"user_730811260095307266" email_address:"user_7338662242574055998@example.com"
```

Amazing! With a few annotations on our schema, we used the super-powers of code generation to create a working gRPC server in no time!

## Caveats and Limitations

`entproto` is still experimental stage and lacks some basic functionality. For example, many applications will probably want a `List` or `Find` method on their service, but these are not yet supported. In addition, some other issues we plan to tackle in the near future:

- Currently only "unique" edges are supported (O2O, O2M).
- The generated "mutating" methods (Create/Update) currently set all fields, disregarding zero/null values and field nullability.
- All fields are copied from the gRPC request to the ent client, support for configuring some fields to be unsettable via the service by adding a field/edge annotation is also planned.

## Next Steps

We believe that `ent` + gRPC can be a great way to build server applications in Go. For example, to set granular access control to the entities managed by our application, developers can already use [Privacy Policies](https://entgo.io/docs/privacy/) that work out-of-the-box with the gRPC integration. To run any arbitrary Go code on the different lifecycle events of entities, developers can utilize custom [Hooks](https://entgo.io/docs/hooks/).

Do you want to build gRPC servers with `ent`? If you want some help setting up or want the integration to support your use case, please reach out to us via our [Discussions Page on GitHub](https://github.com/ent/ent/discussions) or in the #ent channel on the [Gophers Slack](https://app.slack.com/client/T029RQSE6/C01FMSQDT53) or our [Discord server](https://discord.gg/qZmPgTE6RX).


:::note For more Ent news and updates:
- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://app.slack.com/client/T029RQSE6/C01FMSQDT53)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
