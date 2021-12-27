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
Be aware that this tutorial requires an AWS account and **the URL to a public-accessible Postgres database**, which may incur costs.

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
For details on the GraphQL schema design, e.g., the arguments and returns from the `Query` and `Mutation` objects, I obtain inspiration from the [GitHub GraphQL API](https://docs.github.com/en/graphql/reference/queries).

### Setting up AWS Lambda

With the AppSync API in place, our next stop is the AWS Lambda function to run Ent.
For this, we navigate to the AWS Lambda service through the navbar, which leads us to the landing page of the AWS Lambda service showing our list of functions:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/function-list.png" />
  <p style={{fontSize: 12}}>AWS Lambda landing page showing functions.</p>
</div>
We click the "Create function" button on the top right and select "Author from scratch" in the upper panel.
Furthermore, we name the function "ent", set the runtime to "Go 1.x", and click the "Create function" button at the bottom.
We should then find ourselves viewing the landing page of our "ent" function:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/function-overview.png" />
  <p style={{fontSize: 12}}>AWS Lambda function overview of Ent function.</p>
</div>
Before reviewing the Go code and uploading the compiled Go code, we need to adjust some default settings of the function.
First, we change the default handler name from `hello` to `main`, which equals the filename of the compiled Go binary:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/runtime-settings.png" />
  <p style={{fontSize: 12}}>AWS Lambda runtime settings of Ent function.</p>
</div>
Second, we add an environment the variable `DATABASE_URL` encoding the database network parameters and credentials:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of AWS Lambda landing page listing functions" src="https://entgo.io/images/assets/appsync/envars.png" />
  <p style={{fontSize: 12}}>AWS Lambda environemnt variables settings of Ent function.</p>
</div>
The database URL follows the schema `postgres://username:password@hostname/dbname`.
By default, AWS Lambda encrypts the environment variables, making them a fast and safe mechanism to supply database connection parameters.
Alternatively, one can supply an ARN ID to an AWS Secretsmanager secret from which the function requests new credentials.
Using AWS Secretmanager supports rotating credentials but requires an additional request at the start of each Lambda.
A third option is to use AWS IAM to handle the database authorization.
The former two options make only sense when using AWS RDS.

If you created your Postgres database in AWS RDS, the default username and database name are `postgres`.
The password can be reset by modifying the AWS RDS instance.
Be aware that AWS RDS is a rather expensive service and that the network configuration is highly non-trivial!

### Setting up Ent and deploying AWS Lambda

We now review, compile and deploy the missing Go program, which you can also find at [bodokaiser/entgo-aws-appsync](https://github.com/bodokaiser/entgo-aws-appsync), for the "ent" function to run.

First, we create an empty directory to which we change:
```console
mkdir entgo-aws-appsync && cd $_
```
Second, we enable go modules and install the Ent toolset:
```console
go mod init entgo-aws-appsync
go mod tidy
go get -d entgo.io/ent/cmd/ent
```
Third, we create the `Todo` schema
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
Finally, we perform the Ent code generation:
```console
go generate ./ent
```

Using the Ent, we write a set of resolver functions, which implement the create, read, and delete operations on the todos:
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
Using input structs for the resolver functions allows mapping the GraphQL request arguments.
Using output structs allows us to return multiple objects for more complex operations.

To map the Lambda event to a resolver function, we implement a Handler, which performs the mapping according to an `action` field in the event:
```go title="internal/handler/handler.go"
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
In addition to the resolver actions, we also added a migration action, which is a convenient way to expose database migrations.

Finally, we need to register our `Handler` as AWS Lambda event handler:
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
The function body of `main` is executed whenever an AWS Lambda performs a cold starts.
After the cold start, a Lambda function is considered "warm," and only the event handler runs, making Lambda executions extremely efficient.

To compile and deploy the Go code, we run:
```console
GOOS=linux go build -o main ./lambda
zip function.zip main
aws lambda update-function-code --function-name ent --zip-file fileb://function.zip
```
The first command creates a compiled binary named `main`.
The second command compresses the binary to a ZIP archive, required by AWS Lambda.
The third command replaces the function code of the AWS Lambda named `ent` with the new ZIP archive.
If you work with multiple AWS accounts you want to use the `--profile <your aws profile>` swiitch.

After you successfully deployed the AWS Lambda, open the "Test" tab of the `ent` function in the web console and invoke it with a "migrate" action:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of invoking the Ent Lambda with a migrate action" src="https://entgo.io/images/assets/appsync/execution-result.png" />
  <p style={{fontSize: 12}}>Invoking Lambda with a "migrate" action</p>
</div>
On success, you should get a green feedback box and test the result of a "todos" action:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of invoking the Ent Lambda with a todos action" src="https://entgo.io/images/assets/appsync/execution-result2.png" />
  <p style={{fontSize: 12}}>Invoking Lambda with a "todos" action</p>
</div>
In case the test executions fail, you most probably have an issue with your database connection.

### Configuring AWS AppSync resolvers

With the ent Lambda successfully deployed, we are left to register the ent Lambda as a data source to our AppSync API and configure the schema resolvers to map the AppSync requests to Lambda events.
First, open our AWS AppSync API in the web console and move to the "data sources" configuration, which you find in the navigation pane on the left.
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of the list of data sources registered to the AWS AppSync API" src="https://entgo.io/images/assets/appsync/data-sources.png" />
  <p style={{fontSize: 12}}>List of data sources registered to the AWS AppSync API</p>
</div>
Click the "Create data source" button in the top right to start registering the Ent function as data source:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot registering the ent Lambda as data source to the AWS AppSync API" src="https://entgo.io/images/assets/appsync/new-data-source.png" />
  <p style={{fontSize: 12}}>Registering the ent Lambda as data source to the AWS AppSync API</p>
</div>
Now, open the GraphQL schema of the AppSync API and search for the `Query` type in the sidebar to the right.
Click the "Attach" button next to the `Query.Todos` type:
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot attaching a resolver to Query type in the AWS AppSync API" src="https://entgo.io/images/assets/appsync/todo-schema.png" />
  <p style={{fontSize: 12}}>Attaching a resolver for the todos Query in the AWS AppSync API</p>
</div>
In the resolver view for `Query.todos`, select the Lambda function as data source, enable the request mapping template option,
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot configuring the resolver mapping for the todos Query in the AWS AppSync API" src="https://entgo.io/images/assets/appsync/edit-resolver.png" />
  <p style={{fontSize: 12}}>Configuring the resolver mapping for the todos Query in the AWS AppSync API</p>
</div>
and copy the following template:

```vtl title="Query.todos"
{
  "version" : "2017-02-28",
  "operation": "Invoke",
  "payload": {
  	"action": "todos"
  }
}
```

Repeat the same procedure for the remaining `Query` and `Mutation` types:


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

The request mapping templates let us construct the event objects with which we invoke the Lambda functions.
Through the `$context` object, we have access to the GraphQL request and the authentication session.
In addition, it is possible to arange multiple resolvers sequentielly and reference the respectie outputs via the `$context` objct.
In principle, it is also possible to define response mapping templates.
However, in most cases it is suffucient to return the response object "as is".

### Testing AppSync using the Query explorer

The easiest way to test the API is to use the Query Explorer in AWS AppSync.
Alternatively, one can register an API key in the settings of our AppSync API and use any standard GraphQL client.

Let us first create a todo with the title `foo`:
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
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of an executed addTodo Mutation using the AppSync Query Explorer" src="https://entgo.io/images/assets/appsync/todo-queries.png" />
  <p style={{fontSize: 12}}>"addTodo" Mutation using the AppSync Query Explorer</p>
</div>

Requesting a list of the todos should return a single todo with title `foo`:

```graphql
query MyQuery {
  todos {
    title
    id
  }
}

```
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of an executed addTodo Mutation using the AppSync Query Explorer" src="https://entgo.io/images/assets/appsync/todo-queries-3.png" />
  <p style={{fontSize: 12}}>"addTodo" Mutation using the AppSync Query Explorer</p>
</div>

Requesting the `foo` todo by id should work too:

```graphql
query MyQuery {
  todo(id: "1") {
    title
    id
  }
}
```
<div style={{textAlign: 'center'}}>
  <img alt="Screenshot of an executed addTodo Mutation using the AppSync Query Explorer" src="https://entgo.io/images/assets/appsync/todo-queries-4.png" />
  <p style={{fontSize: 12}}>"addTodo" Mutation using the AppSync Query Explorer</p>
</div>

### Wrapping Up

We successfully deployed a serverless GraphQL API for managing simple todos using AWS AppSync, AWS Lambda, and Ent.
In particular, we provided step-by-step instructions on configuring AWS AppSync and AWS Lambda through the web console.
Furthermore, we presented a proposal for how to structure our Go code.

I successfully employed a similar but more complex setup for a B2B SaaS project I am working on part-time in a small team of two.
The developer experience is consistently pleasurable:
Writing SQL queries with Ent is almost a no-brainer because of the type-safe code generation and well-designed API.
Complex features like pagination are provided through the GraphQL extension.
Using a serverless infrastructure keeps the fixed cost down to a minimum and embraces separation of concerns.

Some potential challenges arise with a serverless architecture, which needs to be addressed.
First, having many Lambda functions spawn at the same time quickly exhausts the connection pool of the database, and it is advised to use AWS RDS Proxy as a connection pool.
Second, testing becomes more elaborate as we can not run the GraphQL server or Lambda outside AWS.
In my experience, it is often sufficient to have thorough end-to-end tests for the GraphQL API and local tests for small independent parts of your code, e.g., validations.

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
