---
id: tutorial-todo-gql-field-collection
title: GraphQL Field Collection
sidebar_label: Field Collection
---

In this section, we continue our [GraphQL example](tutorial-todo-gql.mdx) by explaining how Ent implements
[GraphQL Field Collection](https://spec.graphql.org/June2018/#sec-Field-Collection) for our GraphQL schema and solves the
"N+1 Problem" in our resolvers.

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example),
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```

## Problem

The *"N+1 problem"* in GraphQL means that a server executes unnecessary database queries to get node associations (i.e. edges)
when it can be avoided. The number of queries that will be potentially executed (N+1) is a factor of the number of the 
nodes returned by the root query, their associations, and so on recursively. Meaning, this can potentially be a very big number (much bigger than N+1).

Let's try to explain this with the following query:

```graphql
query {
    users(first: 50) {
        edges {
            node {
                photos {
                    link
                }
                posts {
                    content
                    comments {
                        content
                    }
                }
            }
        }
    }
}
```

In the query above, we want to fetch the first 50 users with their photos and their posts, including their comments.

**In the naive solution** (the problematic case), a server will fetch the first 50 users in one query, then, for each user
will execute a query for getting their photos (50 queries), and another query for getting their posts (50). Let's say
each user has exactly 10 posts. Therefore, for each post (of each user), the server will execute another query for getting
its comments (500). That means we will have `1+50+50+500=601` queries in total.

![gql-request-tree](https://entgo.io/images/assets/request-tree.png)

## Ent Solution

The Ent extension for field collection adds support for automatic [GraphQL field collection](https://spec.graphql.org/June2018/#sec-Field-Collection)
for associations (i.e. edges) using [eager loading](eager-load.mdx). Meaning, if a query asks for nodes and their edges, 
`entgql` will automatically add [`With<E>`](eager-load.mdx) steps to the root query, and as a result, the client will
execute a constant number of queries to the database - and it works recursively.

In the GraphQL query above, the client will execute 1 query for getting the users, 1 for getting the photos,
and another 2 for getting the posts and their comments **(4 in total!)**. This logic works both for root queries/resolvers
and for the node(s) API.

## Example

For the purpose of the example, we **disable the automatic field collection**, change the `ent.Client` to run in
debug mode in the `Todos` resolver, and restart our GraphQL server:

```diff title="ent.resolvers.go"
func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
-	return r.client.Todo.Query().
+	return r.client.Debug().Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}
```

We execute the GraphQL query from the [pagination tutorial](tutorial-todo-gql-paginate.md), and add the
`parent` edge to the result:

```graphql
query {
    todos(last: 10, orderBy: {direction: DESC, field: TEXT}) {
        edges {
            node {
                id
                text
                parent {
                    id
                }
            }
            cursor
        }
    }
}
```

Check the process output, and you will see that the server executed 11 queries to the database. 1 for getting the last
10 todo items, and another 10 queries for getting the parent of each item:

```sql
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` ORDER BY `id` ASC LIMIT 11
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` JOIN (SELECT `todo_parent` FROM `todos` WHERE `id` = ?) AS `t1` ON `todos`.`id` = `t1`.`todo_parent` LIMIT 2
```

Let's see how Ent can automatically solve our problem: when defining an Ent edge, `entgql` auto binds it to its usage in
GraphQL and generates edge-resolvers for the nodes under the `gql_edge.go` file:

```go title="ent/gql_edge.go"
func (t *Todo) Children(ctx context.Context) ([]*Todo, error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = t.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = t.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}
```

If we check the process' output again without **disabling fields collection**, we will see that this time the server
executed only two queries to the database. One to get the last 10 todo items, and a second for getting
the parent-item of each todo-item that was returned to the first query.

```sql
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority`, `todos`.`todo_parent` FROM `todos` ORDER BY `id` DESC LIMIT 11
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` WHERE `todos`.`id` IN (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
```

If you're having trouble running this example, go to the [first section](#clone-the-code-optional), clone the code
and run the example.

## Field Mappings

The [`entgql.MapsTo`](https://pkg.go.dev/entgo.io/contrib/entgql#MapsTo) allows you to add a custom field/edge mapping
between the Ent schema and the GraphQL schema. This is useful when you want to expose a field or edge with a different
name(s) in the GraphQL schema. For example:

```go
// One to one mapping.
field.Int("priority").
	Annotations(
		entgql.OrderField("PRIORITY_ORDER"),
		entgql.MapsTo("priorityOrder"),
	)

// Multiple GraphQL fields can map to the same Ent field.
field.Int("category_id").
	Annotations(
		entgql.MapsTo("categoryID", "category_id", "categoryX"),
	)
```

---

Well done! By using automatic field collection for our Ent schema definition, we were able to greatly improve the
GraphQL query efficiency in our application. In the next section, we will learn how to make our GraphQL mutations
transactional.
