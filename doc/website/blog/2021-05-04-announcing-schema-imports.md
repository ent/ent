---
title: Announcing the "Schema Import Initiative" and protoc-gen-ent
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
---

Migrating to a new ORM is not an easy process, and the transition cost can be prohibitive to many organizations. As much
as we developers are enamoured by "Shiny New Things", the truth is that we rarely get a chance to work on a
truly "green-field" project. Most of our careers, we operate in contexts where many technical and business constraints 
(a.k.a legacy systems) dictate and limit our options for moving forward. Developers of new technologies that want to 
succeed must offer interoperability capability and integration paths to help organizations seamlessly transition to a 
new way of solving an existing problem.

To help lower the cost of transitioning to Ent (or simply experimenting with it), we have started the
"**Schema Import Initiative**" to help support many use cases for generating Ent schemas from external resources. 
The centrepiece of this effort is the `schemast` package ([source code](https://github.com/ent/contrib/tree/master/schemast), 
[docs](https://entgo.io/docs/generating-ent-schemas)) which enables developers to easily write programs that generate
and manipulate Ent schemas. Using this package, developers can program in a high-level API, relieving them from worrying
about code parsing and AST manipulations.

### Protobuf Import Support

The first project to use this new API, is `protoc-gen-ent`, a `protoc` plugin to generate Ent schemas from `.proto` 
files ([docs](https://github.com/ent/contrib/tree/master/entproto/cmd/protoc-gen-ent)).  Organizations that have existing 
schemas defined in Protobuf can use this tool to generate Ent code automatically. For example, taking a simple
message definition:

```protobuf
syntax = "proto3";

package entpb;

option go_package = "github.com/yourorg/project/ent/proto/entpb";

message User {
  string name = 1;
  string email_address = 2;
}
```

And setting the `ent.schema.gen` option to true:

```diff
syntax = "proto3";

package entpb;

+import "options/opts.proto";
 
option go_package = "github.com/yourorg/project/ent/proto/entpb";  

message User {
+  option (ent.schema).gen = true; // <-- tell protoc-gen-ent you want to generate a schema from this message
  string name = 1;
  string email_address = 2;
}
```

Developers can invoke the standard `protoc` (protobuf compiler) command to use this plugin:

```shell
protoc -I=proto/ --ent_out=. --ent_opt=schemadir=./schema proto/entpb/user.proto
```

To generate Ent schemas from these definitions:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.String("name"), field.String("email_address")}
}
func (User) Edges() []ent.Edge {
	return nil
}
```

To start using `protoc-gen-ent` today, and read about all of the different configuration options, head over to 
the [documentation](https://github.com/ent/contrib/tree/master/entproto/cmd/protoc-gen-ent)!

### Join the Schema Import Initiative

Do you have schemas defined elsewhere that you would like to automatically import in to Ent?  With the `schemast`
package, it is easier than ever to write the tool that you need to do that. Not sure how to start? Want to collaborate
with the community in planning and building out your idea? Reach out to our great community via our 
[Discord server](https://discord.gg/qZmPgTE6RX), [Slack channel](https://app.slack.com/client/T029RQSE6/C01FMSQDT53) or start a [discussion on GitHub](https://github.com/ent/ent/discussions)!

:::note For more Ent news and updates:
- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://app.slack.com/client/T029RQSE6/C01FMSQDT53)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
