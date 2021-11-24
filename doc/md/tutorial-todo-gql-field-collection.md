---
id: tutorial-todo-gql-field-collection
title: GraphQL Field Collection
sidebar_label: Field Collection
---

In this section, we continue our [GraphQL example](tutorial-todo-gql.md) by explaining how to implement 
[GraphQL Field Collection](https://spec.graphql.org/June2018/#sec-Field-Collection) for our Ent schema and solve the
"N+1 Problem" in our GraphQL resolvers.

#### Clone the code (optional)

The code for this tutorial is available under [github.com/a8m/ent-graphql-example](https://github.com/a8m/ent-graphql-example),
and tagged (using Git) in each step. If you want to skip the basic setup and start with the initial version of the GraphQL
server, you can clone the repository and checkout `v0.1.0` as follows:

```console
git clone git@github.com:a8m/ent-graphql-example.git
cd ent-graphql-example 
go run ./cmd/todo/
```

## Problem

The *"N+1 problem"* in GraphQL means that a server executes unnecessary database queries to get node associations (i.e. edges)
when it can be avoided. The number of queries that potentially executed (N+1) is a factor of the number of the nodes returned
by the root query, their associations, and so on recursively. That means, this can be a very big number (much bigger than N+1).

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

In the query above, we want to fetch the first 50 users with their photos and their posts including their comments.

**In the naive solution** (the problematic case), a server will fetch the first 50 users in 1 query, then, for each user
will execute a query for getting their photos (50 queries), and another query for getting their posts (50). Let's say,
each user has exactly 10 posts. Therefore, For each post (of each user), the server will execute another query for getting
its comments (500). That means, we have `1+50+50+500=601` queries in total.

![gql-request-tree](https://entgo.io/images/assets/request-tree.png)

## Ent Solution

The Ent extension for field collection adds support for automatic [GraphQL fields collection](https://spec.graphql.org/June2018/#sec-Field-Collection)
for associations (i.e. edges) using [eager loading](eager-load.md). That means, if a query asks for nodes and their edges, 
`entgql` will automatically add [`With<E>`](eager-load.md#api) steps to the root query, and as a result, the client will
execute constant number of queries to the database - and it works recursively.

That means, in the GraphQL query above, the client will execute 1 query for getting the users, 1 for getting the photos,
and another 2 for getting the posts, and their comments **(4 in total!)**. This logic works both for root queries/resolvers
and for the node(s) API.

## Example

Before we go over the example, we change the `ent.Client` to run in debug mode in the `Todos` resolver and restart
our GraphQL server:

```diff
func (r *queryResolver) Todos(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TodoOrder) (*ent.TodoConnection, error) {
-	return r.client.Todo.Query().
+	return r.client.Debug().Todo.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithTodoOrder(orderBy),
		)
}
```

Then, we execute the GraphQL query from the [pagination tutorial](tutorial-todo-gql-paginate.md), but we add the
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

We check the process output, and we'll see that the server executed 11 queries to the database. 1 for getting the last
10 todo items, and another 10 for getting the parent of each item:

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

Let's see how Ent can automatically solve our problem. All we need to do is to add the following
`entql` annotations to our edges:

```diff
func (Todo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", Todo.Type).
+			Annotations(entgql.Bind()).
			Unique().
			From("children").
+			Annotations(entgql.Bind()),
	}
}
```

After adding these annotations, `entgql` will do the binding mentioned in the [section](#ent-solution) above. Additionally, it
will also generate edge-resolvers for the nodes under the `edge.go` file:

```go
func (t *Todo) Children(ctx context.Context) ([]*Todo, error) {
	result, err := t.Edges.ChildrenOrErr()
	if IsNotLoaded(err) {
		result, err = t.QueryChildren().All(ctx)
	}
	return result, err
}
```

Let's run the code generation again and re-run our GraphQL server:

```console
go generate ./...
go run ./cmd/todo
```

If we check the process's output again, we will see that this time the server executed only two queries to the database. One, in order to get the last 10 todo items, and a second one for getting the parent-item of each todo-item that was returned in the
first query.

```sql
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority`, `todos`.`todo_parent` FROM `todos` ORDER BY `id` DESC LIMIT 11
SELECT DISTINCT `todos`.`id`, `todos`.`text`, `todos`.`created_at`, `todos`.`status`, `todos`.`priority` FROM `todos` WHERE `todos`.`id` IN (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
```

If you're having troubles running this example, go to the [first section](#clone-the-code-optional), clone the code
and run the example.

---

Well done! By using `entgql.Bind()` in the Ent schema definition, we were able to greatly improve the efficiency of
queries to our application. In the next section, we will learn how to make our GraphQL mutations transactional.
