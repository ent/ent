---
title: Serverless GraphQL using with AWS and ent
author: Bodo Kaiser
authorURL: "https://github.com/bodokaiser"
authorImageURL: "https://avatars.githubusercontent.com/u/1780466?v=4"
---

[GraphQL][1] is a query language for HTTP APIs, providing a statically-typed interface to conveniently represent today's complex data hierarchies.
One way to use GraphQL is to import a library implementing a GraphQL server to which one registers custom resolvers implementing the database interface.
An alternative way is to use a GraphQL cloud service to implement the GraphQL server and register serverless cloud functions as resolvers.
Among the many benefits of cloud services, one of the biggest practical advantages is the resolvers' independence and composability.
For example, we can write one resolver to a relational database and another to a search database.

We consider such a setup using [Amazon Web Services (AWS)][2] in the following. In particular, we use [AWS AppSync][3] as the GraphQL service and [AWS Lambda][4] to run a relational database resolver, which we implement using [Go][5] with [Ent][6] as the entity framework.
Compared to Nodejs, the most popular runtime for AWS Lambda, Go offers faster start times, higher performance, and, in my opinion, an improved developer experience.
On the other hand, Ent presents an innovative approach towards type-safe access to relational databases, which in my opinion, is unmatched in the Go ecosystem.
In conclusion, running Ent with AWS Lambda as AWS AppSync resolvers is an extremely powerful setup to face today's demanding API requirements.

In the next sections, we set up GraphQL in AWS AppSync and the AWS Lambda function running Ent.
Subsequently, we propose a Go implementation integrating Ent and the AWS Lambda event handler, followed by performing a quick test of the Ent function.
Finally, we register it as a data source to our AWS AppSync API and configure the resolvers, which define the mapping from GraphQL requests to AWS Lambda events.
Be aware that this tutorial requires an AWS account, a public accessible Postgres database, which may incur costs.

### Setting up AWS AppSync schema

To set up the GraphQL schema in AWS AppSync, sign in to your AWS account and select the AppSync service through the navbar.
The landing page of the AppSync service should render you a "Create API" button, which you may click to arrive at the "Getting Started" page as depicted in the screenshot below.
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of getting started with AWS AppSync from scratch" src="https://entgo.io/images/assets/appsync/from-scratch.png" />
  <p style={{fontSize: 12}}>Getting started from sratch with AWS AppSync</p>
</div>
In the top panel reading "Customize your API or import from Amazon DynamoDB" select the option "Build from scratch" and click the "Start" button belonging to the panel.
You should now see a form where you may insert the API name.
For the present tutorial, we type "Todo", see the screenshot below, and click the "Create" button.
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of creating a new AWS AppSync API resource" src="https://entgo.io/images/assets/appsync/create-resources.png" />
  <p style={{fontSize: 12}}>Creating a new API resource in AWS AppSync</p>
</div>
After creating the AppSync API, you should see a landing page showing a panel to define the schema, a panel to query the API, and a panel on integrating AppSync into your app as captured in the screenshot below.
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of the landing page of the AWS AppSync API" src="https://entgo.io/images/assets/appsync/getting-started.png" />
  <p style={{fontSize: 12}}>Landing page of the AWS AppSync API</p>
</div>

Click the "Edit Schema" button in the first panel and replace the previous schema with the following GraphQL schema:
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
After replacing the schema, a short validation runs and you should be able to click a "Save Schema" button on the top right corner and you should find yourself with the following view:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot AWS AppSync: Final GraphQL schema for AWS AppSync API" src="https://entgo.io/images/assets/appsync/final-schema.png" />
  <p style={{fontSize: 12}}>Final GraphQL schema of AWS AppSync API</p>
</div>
If we sent GraphQL requests to our AppSync API, the API would return errors as no resolvers have been attached to the schema.
We will configure the resolvers after deploying the Ent function via AWS Lambda.

Explaining the present GraphQL schema in detail is beyond the scope of this tutorial.
In short, the GraphQL schema implements a list todos operation via `Query.todos`, a single read todo operation via `Query.todo`, a create todo operation via `Mutation.createTodo`, and a delete operation via `Mutation.deleteTodo`.
The GraphQL API is similar to a simple REST API design of an `/todos` resource, where we would use `GET /todos`, `GET /todos/:id`, `POST /todos`, and `DELETE /todos/:id`.
For details on the GraphQL schema design, I obtain inspiration from the [GitHub GraphQL API](https://docs.github.com/en/graphql/reference/queries).

### Setting up AWS Lambda

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/function-list.png" />
  <p style={{fontSize: 12}}>AWS Lambda landing page showing functions.</p>
</div>

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/function-overview.png" />
  <p style={{fontSize: 12}}>AWS Lambda function overview of Ent function.</p>
</div>

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/runtime-settings.png" />
  <p style={{fontSize: 12}}>AWS Lambda runtime settings of Ent function.</p>
</div>

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/envars.png" />
  <p style={{fontSize: 12}}>AWS Lambda environemnt variables settings of Ent function.</p>
</div>

### Setting up Ent and deploying AWS Lambda

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
package handler

import (
	"context"
	"encoding/json"
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/resolver"
	"fmt"
	"log"
)

type Action string

const (
	ActionMigrate Action = "migrate"

	ActionTodos      = "todos"
	ActionTodoByID   = "todoById"
	ActionAddTodo    = "addTodo"
	ActionRemoveTodo = "removeTodo"
)

type Event struct {
	Action Action          `json:"action"`
	Input  json.RawMessage `json:"input"`
}

type Handler struct {
	client *ent.Client
}

func New(c *ent.Client) *Handler {
	return &Handler{
		client: c,
	}
}

func (h *Handler) Handle(ctx context.Context, e Event) (interface{}, error) {
	log.Printf("action: %s", e.Action)
	log.Printf("payload: %s", e.Input)

	switch e.Action {
	case ActionMigrate:
		return nil, h.client.Schema.Create(ctx)
	case ActionTodos:
		input := resolver.TodosInput{}
		return resolver.Todos(ctx, h.client, input)
	case ActionTodoByID:
		input := resolver.TodoByIDInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionTodoByID, err)
		}
		return resolver.TodoByID(ctx, h.client, input)
	case ActionAddTodo:
		input := resolver.AddTodoInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionAddTodo, err)
		}
		return resolver.AddTodo(ctx, h.client, input)
	case ActionRemoveTodo:
		input := resolver.RemoveTodoInput{}
		if err := json.Unmarshal(e.Input, &input); err != nil {
			return nil, fmt.Errorf("failed parsing %s params: %w", ActionRemoveTodo, err)
		}
		return resolver.RemoveTodo(ctx, h.client, input)
	}

	return nil, fmt.Errorf("invalid action %q", e.Action)
}
```

Write the event handler:
```go title="internal/handler/resolver.go"
package resolver

import (
	"context"
	"fmt"
	"strconv"

	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/ent/todo"
)

type TodosInput struct{}

func Todos(ctx context.Context, client *ent.Client, input TodosInput) ([]*ent.Todo, error) {
	return client.Todo.
		Query().
		All(ctx)
}

type TodoByIDInput struct {
	ID string `json:"id"`
}

func TodoByID(ctx context.Context, client *ent.Client, input TodoByIDInput) (*ent.Todo, error) {
	tid, err := strconv.Atoi(input.ID)
	if err != nil {
		return nil, fmt.Errorf("failed parsing todo id: %w", err)
	}
	return client.Todo.
		Query().
		Where(todo.ID(tid)).
		Only(ctx)
}

type AddTodoInput struct {
	Title string `json:"title"`
}

type AddTodoOutput struct {
	Todo *ent.Todo `json:"todo"`
}

func AddTodo(ctx context.Context, client *ent.Client, input AddTodoInput) (*AddTodoOutput, error) {
	t, err := client.Todo.
		Create().
		SetTitle(input.Title).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating todo: %w", err)
	}
	return &AddTodoOutput{Todo: t}, nil
}

type RemoveTodoInput struct {
	TodoID string `json:"todoId"`
}

type RemoveTodoOutput struct {
	Todo *ent.Todo `json:"todo"`
}

func RemoveTodo(ctx context.Context, client *ent.Client, input RemoveTodoInput) (*RemoveTodoOutput, error) {
	t, err := TodoByID(ctx, client, TodoByIDInput{ID: input.TodoID})
	if err != nil {
		return nil, fmt.Errorf("failed querying todo with id %q: %w", input.TodoID, err)
	}
	err = client.Todo.
		DeleteOne(t).
		Exec(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed deleting todo with id %q: %w", input.TodoID, err)
	}
	return &RemoveTodoOutput{Todo: t}, nil
}
```

Bootstrap the Ent client and event handler for AWS Lambda:
```go title="lambda/main.go"
package main

import (
	"database/sql"
	"entgo-aws-appsync/ent"
	"entgo-aws-appsync/internal/handler"
	"log"
	"os"

	entsql "entgo.io/ent/dialect/sql"

	"entgo.io/ent/dialect"
	"github.com/aws/aws-lambda-go/lambda"
	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {
	db, err := sql.Open("pgx", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalf("failed opening database connection: %v", err)
	}

	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, db)))
	defer client.Close()

	lambda.Start(handler.New(client).Handle)
}
```

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of invoking the Ent Lambda with a migrate action" src="https://entgo.io/images/assets/appsync/execution-result.png" />
  <p style={{fontSize: 12}}>Invoking Lambda with a "migrate" action</p>
</div>

<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of invoking the Ent Lambda with a todos action" src="https://entgo.io/images/assets/appsync/execution-result2.png" />
  <p style={{fontSize: 12}}>Invoking Lambda with a "todos" action</p>
</div>


### Configuring AWS AppSync resolvers

```vtl title="Query.todos"
{
  "version" : "2017-02-28",
  "operation": "Invoke",
  "payload": {
  	"action": "todos"
  }
}
```

```vtl title="Query.todo"
{
  "version" : "2017-02-28",
  "operation": "Invoke",
  "payload": {
  	"action": "todo",
    "input": $util.toJson($context.args.input)
  }
}
```

```vtl title="Mutation.addTodo"
{
  "version" : "2017-02-28",
  "operation": "Invoke",
  "payload": {
  	"action": "addTodo",
    "input": $util.toJson($context.args.input)
  }
}
```

```vtl title="Mutation.removeTodo"
{
  "version" : "2017-02-28",
  "operation": "Invoke",
  "payload": {
  	"action": "removeTodo",
    "input": $util.toJson($context.args.input)
  }
}
```

### Testing AppSync using the Query explorer

```graphql
mutation MyMutation {
  addTodo(input: {title: "foo"}) {
    todo {
      id
      title
    }
  }
}

```

```graphql
query MyQuery {
  todos {
    title
    id
  }
}

```

```graphql
query MyQuery {
  todo(id: "1") {
    title
    id
  }
}
```

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
