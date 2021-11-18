---
id: grpc-service-generation-options
title: Configuring Service Method Generation
sidebar_label: Service Generation Options
---
By default, entproto will generate CRUD service methods for an `ent.Schema` annotated with `ent.Service()`. Method generation can be customized by including the argument `entproto.Methods()` in the `entproto.Service()` annotation. `entproto.Methods()` accepts bit flags to determine what service methods should be generated. The flags include:
```go
// Generates a Create gRPC service method for the entproto.Service.
entproto.MethodCreate

// Generates a Get gRPC service method for the entproto.Service.
entproto.MethodGet

// Generates an Update gRPC service method for the entproto.Service.
entproto.MethodUpdate

// Generates a Delete gRPC service method for the entproto.Service.
entproto.MethodDelete

// Generates all service methods for the entproto.Service.
// This is the same behavior as not including entproto.Methods.
entproto.MethodAll
```
To generate a service with multiple methods, bitwise OR the flags.


To see this in action, we can modify our ent schema. Let's say we wanted to prevent our gRPC client from mutating entries. We can accomplish this by modifying `ent/schema/user.go`:
```go {5}
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entproto.Message(),
		entproto.Service(
			entproto.Methods(entproto.MethodCreate | entproto.MethodGet),
        ),
	}
}
```

Re-running `go generate ./...` will give us the following service definition in `entpb.proto`:
```protobuf
service UserService {
  rpc Create ( CreateUserRequest ) returns ( User );

  rpc Get ( GetUserRequest ) returns ( User );
}
```

Notice that the service no longer includes `Update` and `Delete` methods. Perfect! 