---
title: Serverless GraphQL using with AWS and ent
author: Bodo Kaiser
authorURL: "https://github.com/bodokaiser"
authorImageURL: "https://avatars.githubusercontent.com/u/1780466?v=4"
---

[Graphql][1] is a query language for HTTP APIs, providing a statically-typed interface to conveniently represent today's complex data hierarchies.
One way to use GraphQL is to import a library implementing a GraphQL server to which one registers custom resolvers implementing the database interface.
An alternative way is to use a GraphQL cloud service to implement the GraphQL server and register serverless cloud functions as resolvers.
Among the many benefits of cloud services, one of the biggest practical advantages is the resolvers' independence and composability.
For example, we can write one resolver to a relational database and another to a search database.

We consider such a setup using [Amazon Web Services (AWS)][2] in the following. In particular, we use [AWS AppSync][3] as the GraphQL service and [AWS Lambda][4] to run a relational database resolver, which we implement using [Go][5] with [Ent][6] as the entity framework.
Compared to Nodejs, the most popular runtime for AWS Lambda, Go offers faster start times, higher performance, and, in my opinion, an improved developer experience.
On the other hand, Ent presents an innovative approach towards type-safe access to relational databases, which in my opinion, is unmatched in the Go ecosystem.
In conclusion, running Ent with AWS Lambda as AWS AppSync resolvers is an extremely powerful setup to face today's demanding API requirements.

### Setting up AWS AppSync

```graphql
input AddTodoInput {
	title: String!
}

type AddTodoOutput {
	todo: Todo!
}

type Mutation {
	addTodo(input: AddTodoInput!): AddTodoOutput!
	removeTodo(input: RemoveTodoInput!): RemoveTodoOutput!
}

type Query {
	todos: [Todo!]!
	todo(id: ID!): Todo
}

input RemoveTodoInput {
	todoId: ID!
}

type RemoveTodoOutput {
	todo: Todo!
}

type Todo {
	id: ID!
	title: String!
}

schema {
	query: Query
	mutation: Mutation
}
```

### Developing AWS Lambda with Ent

Create an empty directory and change directory:
```console
mkdir entgo-aws-appsync && cd $_
```
Setup go modules and Ent:
```console
go mod init entgo-aws-appsync
go mod tidy
go get -d entgo.io/ent/cmd/ent
```

Create the `Todo` schema
```console
go run entgo.io/ent/cmd/ent init Todo
```
and add the `title` field:
```go title="ent/schema/todo.go"
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Todo holds the schema definition for the Todo entity.
type Todo struct {
	ent.Schema
}

// Fields of the Todo.
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.String("title"),
	}
}

// Edges of the Todo.
func (Todo) Edges() []ent.Edge {
	return nil
}
```
Finally, generate the schema:
```console
go generate ./ent
```

Write the resolvers:
```go title="internal/resolver/resolver.go"

```

Write the event handler:
```go title="internal/handler/resolver.go"

```

Bootstrap the Ent client and event handler for AWS Lambda:
```go title="lambda/main.go"
package main

import (
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/handler"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	lambda.Start(handler.New(client))
}
```

### Deploying AWS Lambda

### Wrapping Up

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::

[1]: https://graphql.org
[2]: https://aws.amazon.com
[3]: https://aws.amazon.com/appsync/
[4]: https://aws.amazon.com/lambda/
[5]: https://go.dev
[6]: https://entgo.io
