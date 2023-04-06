---
id: grpc-optional-fields
title: Optional Fields
sidebar_label: Optional Fields
---
A common issue with Protobufs is that the way that nil values are represented: a zero-valued primitive field isn't
encoded into the binary representation, this means that applications cannot distinguish between zero and not-set for
primitive fields. 

To support this, the Protobuf project supports some [Well-Known types](https://developers.google.com/protocol-buffers/docs/reference/google.protobuf) called "wrapper types".
For example, the wrapper type for a `bool`, is called `google.protobuf.BoolValue` and is [defined as](https://github.com/protocolbuffers/protobuf/blob/991bcada050d7e9919503adef5b52547ec249d35/src/google/protobuf/wrappers.proto#L103-L107):
```protobuf title="ent/proto/entpb/entpb.proto"
// Wrapper message for `bool`.
//
// The JSON representation for `BoolValue` is JSON `true` and `false`.
message BoolValue {
  // The bool value.
  bool value = 1;
}
```
When `entproto` generates a Protobuf message definition, it uses these wrapper types to represent "Optional" ent fields.

Let's see this in action, modifying our ent schema to include an optional field:

```go title="ent/schema/user.go" {14-16}
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
		field.String("alias").
			Optional().
			Annotations(entproto.Field(4)),
	}
}
```

Re-running `go generate ./...`, observe that our Protobuf definition for `User` now looks like:

```protobuf title="ent/proto/entpb/entpb.proto" {8}
message User {
  int32 id = 1;

  string name = 2;

  string email_address = 3;

  google.protobuf.StringValue alias = 4; // <-- this is new 

  repeated Category administered = 5;
}
```

The generated service implementation also utilize this field. Observe in `entpb_user_service.go`:

```go title="ent/proto/entpb/entpb_user_service.go" {3-6}
func (svc *UserService) createBuilder(user *User) (*ent.UserCreate, error) {
	m := svc.client.User.Create()
	if user.GetAlias() != nil {
		userAlias := user.GetAlias().GetValue()
		m.SetAlias(userAlias)
	}
	userEmailAddress := user.GetEmailAddress()
	m.SetEmailAddress(userEmailAddress)
	userName := user.GetName()
	m.SetName(userName)
	for _, item := range user.GetAdministered() {
		administered := int(item.GetId())
		m.AddAdministeredIDs(administered)
	}
	return m, nil
}
```

To use the wrapper types in our client code, we can use helper methods supplied by the [wrapperspb](https://github.com/protocolbuffers/protobuf-go/blob/3f51f05e40d61e930a5416f1ed7092cef14cc058/types/known/wrapperspb/wrappers.pb.go#L458-L460)
package to easily build instances of these types. For example in `cmd/client/main.go`:
```go {5}
func randomUser() *entpb.User {
	return &entpb.User{
		Name:         fmt.Sprintf("user_%d", rand.Int()),
		EmailAddress: fmt.Sprintf("user_%d@example.com", rand.Int()),
		Alias:        wrapperspb.String("John Doe"),
	}
}
```