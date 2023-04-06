---
id: grpc-server-and-client
title: Creating the Server and Client
sidebar_label: Server and Client
---

Getting an automatically generated gRPC service definition is super cool, but we still need to register it to a
concrete gRPC server, that listens on some TCP port for traffic and is able to respond to RPC calls. 

We decided not to generate this part automatically because it typically involves some team/org specific
behavior such as wiring in different middlewares. This may change in the future. In the meantime, this section
describes how to create a simple gRPC server that will serve our service code.

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

Let's create a simple client that makes some calls to our server. Create a new file named `cmd/client/main.go` and write:

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
	conn, err := grpc.Dial(":5000", grpc.WithTransportCredentials(insecure.NewCredentials()))
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

Our client creates a connection to port 5000, where our server is listening, then issues a `Create`
request to create a new user, and then issues a second `Get` request to retrieve it from the database.
Let's run our client code:

```console
go run ./cmd/client
```

Observe the output:

```console
2021/03/18 10:42:58 user created with id: 1
2021/03/18 10:42:58 retrieved user with id=1: id:1 name:"user_730811260095307266" email_address:"user_7338662242574055998@example.com"
```

Hooray! We have successfully created a real gRPC client to talk to our real gRPC server! In the next sections, we will
see how the ent/gRPC integration deals with more advanced ent schema definitions.
