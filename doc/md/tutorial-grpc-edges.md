---
id: grpc-edges
title: Working with Edges
sidebar_label: Working with Edges
---
Edges enable us to express the relationship between different entities in our ent application. Let's see how they work
together with generated gRPC services.

Let's start by adding a new entity, `Category` and create edges relating our `User` type to it:

```go title="ent/schema/category.go"
package schema

import (
	"entgo.io/contrib/entproto"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Category struct {
	ent.Schema
}

func (Category) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Annotations(entproto.Field(2)),
	}
}

func (Category) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
	}
}

func (Category) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("admin", User.Type).
			Unique().
			Annotations(entproto.Field(3)),
	}
}
```

Creating the inverse relation on the `User`:

```go title="ent/schema/user.go" {4-6}
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("administered", Category.Type).
			Ref("admin").
			Annotations(entproto.Field(5)),
	}
}
```

Notice a few things:

* Our edges also receive an `entproto.Field` annotation. We will see why in a minute.
* We created a one-to-many relationship where a `Category` has a single `admin`, and a `User` can administer multiple
  categories.

Re-generating the project with `go generate ./...`, notice the changes to the `.proto` file:

```protobuf title="ent/proto/entpb/entpb.proto" {1-7,18}
message Category {
  int64 id = 1;

  string name = 2;

  User admin = 3;
}

message User {
  int64 id = 1;

  string name = 2;

  string email_address = 3;

  google.protobuf.StringValue alias = 4;

  repeated Category administered = 5;
}
```

Observe the following changes:

* A new message, `Category` was created. This message has a field named `admin` corresponding to the `admin` edge on
  the `Category` schema. It is a non-repeated field because we set the edge to be `.Unique()`. It's field number is `3`,
  corresponding to the `entproto.Field` annotation on the edge definition.
* A new field `administered` was added to the `User` message definition. It is a `repeated` field, corresponding to the
  fact that we did not mark the edge as `Unique` in this direction. It's field number is `5`, corresponding to the
  `entproto.Field` annotation on the edge.

### Creating Entities with their Edges

Let's demonstrate how to create an entity with its edges by writing a test:

```go
package main

import (
	"context"
	"testing"

	_ "github.com/mattn/go-sqlite3"

	"ent-grpc-example/ent/category"
	"ent-grpc-example/ent/enttest"
	"ent-grpc-example/ent/proto/entpb"
	"ent-grpc-example/ent/user"
)

func TestServiceWithEdges(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and
	// creating a gRPC server and instead we are just calling the library code directly. 
	svc := entpb.NewUserService(client)

	// next, we create a category directly using the ent client.
	// Notice we are initializing it with no relation to a User.
	cat := client.Category.Create().SetName("cat_1").SaveX(ctx)

	// next, we invoke the User service's `Create` method. Notice we are
	// passing a list of entpb.Category instances with only the ID set. 
	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:         "user",
			EmailAddress: "user@service.code",
			Administered: []*entpb.Category{
				{Id: int64(cat.ID)},
			},
		},
	})
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	// to verify everything worked correctly, we query the category table to check
	// we have exactly one category which is administered by the created user.
	count, err := client.Category.
		Query().
		Where(
			category.HasAdminWith(
				user.ID(int(create.Id)),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}
	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}
```


To create the edge from the created `User` to the existing `Category` we do not need to populate the entire `Category`
object. Instead we only populate the `Id` field. This is picked up by the generated service code:

```go title="ent/proto/entpb/entpb_user_service.go" {3-6}
func (svc *UserService) createBuilder(user *User) (*ent.UserCreate, error) {
	  // truncated ...
	for _, item := range user.GetAdministered() {
		administered := int(item.GetId())
		m.AddAdministeredIDs(administered)
	}
	return m, nil
}
```

### Retrieving Edge IDs for Entities

We have seen how to create relations between entities, but how do we retrieve that data from the generated gRPC
service? 

Consider this example test:

```go
func TestGet(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and
	// creating a gRPC server and instead we are just calling the library code directly.
	svc := entpb.NewUserService(client)

	// next, create a user, a category and set that user to be the admin of the category
	user := client.User.Create().
		SetName("rotemtam").
		SetEmailAddress("r@entgo.io").
		SaveX(ctx)

	client.Category.Create().
		SetName("category").
		SetAdmin(user).
		SaveX(ctx)

	// next, retrieve the user without edge information
	get, err := svc.Get(ctx, &entpb.GetUserRequest{
		Id: int64(user.ID),
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 0 {
		t.Fatal("by default edge information is not supposed to be retrieved")
	}

	// next, retrieve the user *WITH* edge information
	get, err = svc.Get(ctx, &entpb.GetUserRequest{
		Id:   int64(user.ID),
		View: entpb.GetUserRequest_WITH_EDGE_IDS,
	})
	if err != nil {
		t.Fatal("failed retrieving the created user", err)
	}
	if len(get.Administered) != 1 {
		t.Fatal("using WITH_EDGE_IDS edges should be returned")
	}
}
```

As you can see in the test, by default, edge information is not returned by the `Get` method of the service. This is
done deliberately because the amount of entities related to an entity is unbound. To allow the caller of to specify
whether or not to return the edge information or not, the generated service adheres to [AIP-157](https://google.aip.dev/157)
(Partial Responses). In short, the `GetUserRequest` message includes an enum named `View`:

```protobuf title="ent/proto/entpb/entpb.proto"
message GetUserRequest {
  int64 id = 1;

  View view = 2;

  enum View {
    VIEW_UNSPECIFIED = 0;

    BASIC = 1;

    WITH_EDGE_IDS = 2;
  }
}
```

Consider the generated code for the `Get` method:

```go title="ent/proto/entpb/entpb_user_service.go"
// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	// .. truncated ..
	switch req.GetView() {
	case GetUserRequest_VIEW_UNSPECIFIED, GetUserRequest_BASIC:
		get, err = svc.client.User.Get(ctx, int(req.GetId()))
	case GetUserRequest_WITH_EDGE_IDS:
		get, err = svc.client.User.Query().
			Where(user.ID(int(req.GetId()))).
			WithAdministered(func(query *ent.CategoryQuery) {
				query.Select(category.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: unknown view")
	}
// .. truncated ..
}
```
By default, `client.User.Get` is invoked, which does not return any edge ID information, but if `WITH_EDGE_IDS` is passed,
the endpoint will retrieve the `ID` field for any `Category` related to the user via the `administered` edge.