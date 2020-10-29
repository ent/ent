---
id: migrate
title: Database Migration
---

The migration support for `ent` provides the option for keeping the database schema
aligned with the schema objects defined in `ent/migrate/schema.go` under the root of your project.

## Auto Migration

Run the auto-migration logic in the initialization of the application:

```go
if err := client.Schema.Create(ctx); err != nil {
	log.Fatalf("failed creating schema resources: %v", err)
}
```

`Create` creates all database resources needed for your `ent` project. By default, `Create` works
in an *"append-only"* mode; which means, it only creates new tables and indexes, appends columns to tables or 
extends column types. For example, changing `int` to `bigint`.

What about dropping columns or indexes?

## Drop Resources

`WithDropIndex` and `WithDropColumn` are 2 options for dropping table columns and indexes.

```go
package main

import (
	"context"
	"log"
	
	"<project>/ent"
	"<project>/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(
		ctx, 
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true), 
	)
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

In order to run the migration in debug mode (printing all SQL queries), run:

```go
err := client.Debug().Schema.Create(
	ctx, 
	migrate.WithDropIndex(true),
	migrate.WithDropColumn(true),
)
if err != nil {
	log.Fatalf("failed creating schema resources: %v", err)
}
```

## Universal IDs

By default, SQL primary-keys start from 1 for each table; which means that multiple entities of different types
can share the same ID. Unlike AWS Neptune, where node IDs are UUIDs.

This does not work well if you work with [GraphQL](https://graphql.org/learn/schema/#scalar-types), which requires
the object ID to be unique.

To enable the Universal-IDs support for your project, pass the `WithGlobalUniqueID` option to the migration.

```go
package main

import (
	"context"
	"log"
	
	"<project>/ent"
	"<project>/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	if err := client.Schema.Create(ctx, migrate.WithGlobalUniqueID(true)); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

**How does it work?** `ent` migration allocates a 1<<32 range for the IDs of each entity (table),
and store this information in a table named `ent_types`. For example, type `A` will have the range
of `[1,4294967296)` for its IDs, and type `B` will have the range of `[4294967296,8589934592)`, etc.

Note that if this option is enabled, the maximum number of possible tables is **65535**. 

## Offline Mode

Offline mode allows you to write the schema changes to an `io.Writer` before executing them on the database.
It's useful for verifying the SQL commands before they're executed on the database, or to get an SQL script
to run manually. 

**Print changes**
```go
package main

import (
	"context"
	"log"
	"os"
	
	"<project>/ent"
	"<project>/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Dump migration changes to stdout.
	if err := client.Schema.WriteTo(ctx, os.Stdout); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
```

**Write changes to a file**
```go
package main

import (
	"context"
	"log"
	"os"
	
	"<project>/ent"
	"<project>/ent/migrate"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Dump migration changes to an SQL script.
	f, err := os.Create("migrate.sql")
	if err != nil {
		log.Fatalf("create migrate file: %v", err)
	}
	defer f.Close()
	if err := client.Schema.WriteTo(ctx, f); err != nil {
		log.Fatalf("failed printing schema changes: %v", err)
	}
}
```

## Foreign Keys

By default, `ent` uses foreign-keys when defining relationships (edges) to enforce correctness and consistency on the
database side.

However, `ent` also provide an option to disable this functionality using the `WithForeignKeys` option.
You should note that setting this option to `false`, will tell the migration to not create foreign-keys in the
schema DDL and the edges validation and clearing must be handled manually by the developer.

We expect to provide a set of hooks for implementing the foreign-key constraints in the application level in the near future.

```go
package main

import (
    "context"
    "log"

    "<project>/ent"
    "<project>/ent/migrate"
)

func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err = client.Schema.Create(
        ctx,
        migrate.WithForeignKeys(false), // Disable foreign keys.
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```