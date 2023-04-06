---
id: tutorial-todo-crud
title: Query and Mutation
sidebar_label: Query and Mutation
---

After setting up our project, we're ready to create our Todo list and query it.

## Create a Todo

Let's create a Todo in our testable example. We do it by adding the following code to `example_test.go`:

```go
func Example_Todo() {
	// ...
	task1, err := client.Todo.Create().Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Println(task1)
	// Output:
	// Todo(id=1)
}
```

Running `go test` should pass successfully. 

## Add Fields To The Schema

As you can see, our Todos are too boring as they contain only the `ID` field. Let's improve this example by adding
multiple fields to the schema in `todo/ent/schema/todo.go`:

```go
func (Todo) Fields() []ent.Field {
	return []ent.Field{
		field.Text("text").
			NotEmpty(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Enum("status").
			NamedValues(
				"InProgress", "IN_PROGRESS",
				"Completed", "COMPLETED",
			).
			Default("IN_PROGRESS"),
		field.Int("priority").
			Default(0),
	}
}
```

After adding these fields, we need to run the code-generation as before:

```console
go generate ./ent
```

As you may notice, all fields have a default value on creation except the `text` field, which must be provided by
the user. Let's change our `example_test.go` to follow these changes:

```go
func Example_Todo() {
	// ...
	task1, err := client.Todo.Create().SetText("Add GraphQL Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Printf("%d: %q\n", task1.ID, task1.Text)
	task2, err := client.Todo.Create().SetText("Add Tracing Example").Save(ctx)
	if err != nil {
		log.Fatalf("failed creating a todo: %v", err)
	}
	fmt.Printf("%d: %q\n", task2.ID, task2.Text)
    // Output:
    // 1: "Add GraphQL Example"
    // 2: "Add Tracing Example"
}
```

Wonderful! We created a schema in the database with 5 columns (`id`, `text`, `created_at`, `status`, `priority`)
and created 2 items in our todo list, by inserting 2 rows to the table.

![tutorial-todo-create](https://entgo.io/images/assets/tutorial-todo-create-items.png)

## Add Edges To The Schema

Let’s say we want to design our todo list so that an item can depend on another item. Therefore, we'll add a `parent`
edge to each Todo item, to get the item it depends on, and a back-reference edge named `children` in order to get all
items that depend on it.

Let's change our schema again in `todo/ent/schema/todo.go`:

```go
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Todo.Type).
			Unique().
			From("children"),
	}
}
```

After adding these edges, we need to run the code-generation as before:

```console
go generate ./ent
```

## Connect 2 Todos

We continue our edges example, by updating the 2 todo items we just created. We define that item-2 (*"Add Tracing Example"*)
depends on item-1 (*"Add GraphQL Example"*). 

![tutorial-todo-create](https://entgo.io/images/assets/tutorial-todo-create-edges.png)

```go
func Example_Todo() {
	// ...
	if err := task2.Update().SetParent(task1).Exec(ctx); err != nil {
		log.Fatalf("failed connecting todo2 to its parent: %v", err)
	}
    // Output:
    // 1: "Add GraphQL Example"
    // 2: "Add Tracing Example"
}
```

## Query Todos

After connecting item-2 to item-1, we're ready to start querying our todo list. 

#### Query all todo items:

```go
func Example_Todo() {
	// ...

	// Query all todo items.
	items, err := client.Todo.Query().All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// Output:
	// 1: "Add GraphQL Example"
	// 2: "Add Tracing Example"
}
```

#### Query all todo items that depend on other items:

```go
func Example_Todo() {
	// ...

	// Query all todo items that depend on other items.
	items, err := client.Todo.Query().Where(todo.HasParent()).All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// Output:
	// 2: "Add Tracing Example"
}
```

#### Query all todo items that don't depend on other items and have items that depend on them:

```go
func Example_Todo() {
	// ...

	// Query all todo items that don't depend on other items and have items that depend them.
	items, err := client.Todo.Query().
		Where(
			todo.Not(
				todo.HasParent(),
			),
			todo.HasChildren(),
		).
		All(ctx)
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	for _, t := range items {
		fmt.Printf("%d: %q\n", t.ID, t.Text)
	}
	// Output:
	// 1: "Add GraphQL Example"
}
```

#### Query parent through its children:

```go
func Example_Todo() {
	// ...
	
	// Get a parent item through its children and expect the
	// query to return exactly one item.
	parent, err := client.Todo.Query(). // Query all todos.
		Where(todo.HasParent()).        // Filter only those with parents.
		QueryParent().                  // Continue traversals to the parents.
		Only(ctx)                       // Expect exactly one item.
	if err != nil {
		log.Fatalf("failed querying todos: %v", err)
	}
	fmt.Printf("%d: %q\n", parent.ID, parent.Text)
	// Output:
	// 1: "Add GraphQL Example"
}
```
