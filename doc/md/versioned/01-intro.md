---
id: intro
title: Introduction
---
## Schema Migration Flows

Ent supports two different workflows for managing schema changes:
* Automatic Migrations - a declarative style of schema migrations which happen entirely at runtime.
 With this flow, Ent calculates the difference between the connected database and the database
 schema needed to satisfy the `ent.Schema` definitions, and then applies the changes to the database.
* Versioned Migrations - a workflow where schema migrations are written as SQL files ahead of time
 and then are applied to the database by a specialized tool such as [Atlas](https://atlasgo.io) or 
 [golang-migrate](https://github.com/golang-migrate/migrate).

Many users start with the automatic migration flow as it is the easiest to get started with, but
as their project grows, they may find that they need more control over the migration process, and
they switch to the versioned migration flow.

This tutorial will walk you through the process of upgrading an existing project from automatic migrations
to versioned migrations. 

## Supporting repository

All of the steps demonstrated in this tutorial can be found in the 
[rotemtam/ent-versioned-migrations-demo](https://github.com/rotemtam/ent-versioned-migrations-demo)
repository on GitHub. In each section we will link to the relevant commit in the repository.

The initial Ent project which we will be upgrading can be found
[here](https://github.com/rotemtam/ent-versioned-migrations-demo/tree/start).

## Automatic Migration

In this tutorial, we assume you have an existing Ent project and that you are using automatic migrations.
Many simpler projects have a block of code similar to this in their `main.go` file:

```go
package main

func main() {
	// Connect to the database (MySQL for example).
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	// highlight-next-line
	if err := client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	// ... Continue with server start.
}
```

This code connects to the database, and then runs the automatic migration tool to create all schema resources.

Next, let's see how to set up our project for versioned migrations.