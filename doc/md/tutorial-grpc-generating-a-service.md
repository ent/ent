---
id: grpc-generating-a-service
title: Generating a gRPC Service
sidebar_label: Generating a Service
---
Generating Protobuf structs generated from our `ent.Schema` can be useful, but what we're really interested in is getting an actual server that can create, read, update, and delete entities from an actual database. To do that, we need to update just one line of code! When we annotate a schema with `entproto.Service`, we tell the `entproto` code-gen that we are interested in generating a gRPC service definition, from the `protoc-gen-entgrpc` will read our definition and generate a service implementation. Edit `ent/schema/user.go` and modify the schema's `Annotations`:

```go title="ent/schema/user.go" {4}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(), // <-- add this
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

```protobuf title="ent/proto/entpb/entpb.proto"
service UserService {
  rpc Create ( CreateUserRequest ) returns ( User );

  rpc Get ( GetUserRequest ) returns ( User );

  rpc Update ( UpdateUserRequest ) returns ( User );

  rpc Delete ( DeleteUserRequest ) returns ( google.protobuf.Empty );

  rpc List ( ListUserRequest ) returns ( ListUserResponse );

  rpc BatchCreate ( BatchCreateUsersRequest ) returns ( BatchCreateUsersResponse );
}
```

In addition, two new files were created. The first, `entpb_grpc.pb.go`, contains the gRPC client stub and the interface definition. If you open the file, you will find in it (among many other things):

```go title="ent/proto/entpb/entpb_grpc.pb.go"
// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please
// refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Create(ctx context.Context, in *CreateUserRequest, opts ...grpc.CallOption) (*User, error)
	Get(ctx context.Context, in *GetUserRequest, opts ...grpc.CallOption) (*User, error)
	Update(ctx context.Context, in *UpdateUserRequest, opts ...grpc.CallOption) (*User, error)
	Delete(ctx context.Context, in *DeleteUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	List(ctx context.Context, in *ListUserRequest, opts ...grpc.CallOption) (*ListUserResponse, error)
	BatchCreate(ctx context.Context, in *BatchCreateUsersRequest, opts ...grpc.CallOption) (*BatchCreateUsersResponse, error)
}
```

The second file, `entpub_user_service.go` contains a generated implementation for this interface. For example, an implementation for the `Get` method:

```go title="ent/proto/entpb/entpb_user_service.go"
// Get implements UserServiceServer.Get
func (svc *UserService) Get(ctx context.Context, req *GetUserRequest) (*User, error) {
	var (
		err error
		get *ent.User
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetUserRequest_VIEW_UNSPECIFIED, GetUserRequest_BASIC:
		get, err = svc.client.User.Get(ctx, id)
	case GetUserRequest_WITH_EDGE_IDS:
		get, err = svc.client.User.Query().
			Where(user.ID(id)).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoUser(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}
}

```

Not bad! Next, let's create a gRPC server that can serve requests to our service.
