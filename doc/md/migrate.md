---
id: migrate
title: Automatic Migration
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
can share the same ID. Unlike AWS Neptune, where node IDs are UUIDs. [Read this](features.md#globally-unique-id) to 
learn how to enable universally unique ids when using Ent with a SQL database.

## Offline Mode

**With Atlas becoming the default migration engine soon, offline migration will be replaced
by [versioned migrations](versioned-migrations.mdx).**

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

## Migration Hooks

The framework provides an option to add hooks (middlewares) to the migration phase.
This option is ideal for modifying or filtering the tables that the migration is working on,
or for creating custom resources in the database.

```go
package main

import (
    "context"
    "log"

    "<project>/ent"
    "<project>/ent/migrate"

    "entgo.io/ent/dialect/sql/schema"
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
        schema.WithHooks(func(next schema.Creator) schema.Creator {
            return schema.CreateFunc(func(ctx context.Context, tables ...*schema.Table) error {
                // Run custom code here.
                return next.Create(ctx, tables...)
            })
        }),
    )
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

## Atlas Integration

Starting with v0.10, Ent supports running migration with [Atlas](https://atlasgo.io), which is a more robust
migration framework that covers many features that are not supported by current Ent migrate package. In order
to execute a migration with the Atlas engine, use the `WithAtlas(true)` option.

```go {21}
package main

import (
    "context"
    "log"

    "<project>/ent"
    "<project>/ent/migrate"

    "entgo.io/ent/dialect/sql/schema"
)

func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err = client.Schema.Create(ctx, schema.WithAtlas(true))
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

In addition to the standard options (e.g. `WithDropColumn`, `WithGlobalUniqueID`), the Atlas integration provides additional
options for hooking into schema migration steps.

![atlas-migration-process](https://entgo.io/images/assets/migrate-atlas-process.png)


#### Atlas `Diff` and `Apply` Hooks

Here are two examples that show how to hook into the Atlas `Diff` and `Apply` steps.

```go
package main

import (
    "context"
    "log"

    "<project>/ent"
    "<project>/ent/migrate"

	"ariga.io/atlas/sql/migrate"
	atlas "ariga.io/atlas/sql/schema"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
)

func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
    // Run migration.
    err := 	client.Schema.Create(
		ctx,
		// Hook into Atlas Diff process.
		schema.WithDiffHook(func(next schema.Differ) schema.Differ {
			return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
				// Before calculating changes.
				changes, err := next.Diff(current, desired)
				if err != nil {
					return nil, err
				}
				// After diff, you can filter
				// changes or return new ones.
				return changes, nil
			})
		}),
		// Hook into Atlas Apply process.
		schema.WithApplyHook(func(next schema.Applier) schema.Applier {
			return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
				// Example to hook into the apply process, or implement
				// a custom applier. For example, write to a file.
				//
				//	for _, c := range plan.Changes {
				//		fmt.Printf("%s: %s", c.Comment, c.Cmd)
				//		if err := conn.Exec(ctx, c.Cmd, c.Args, nil); err != nil {
				//			return err
				//		}
				//	}
				//
				return next.Apply(ctx, conn, plan)
			})
		}),
	)
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

#### `Diff` Hook Example

In case a field was renamed in the `ent/schema`, Ent won't detect this change as renaming and will propose `DropColumn`
and `AddColumn` changes in the diff stage. One way to get over this is to use the
[StorageKey](schema-fields.mdx#storage-key) option on the field and keep the old column name in the database table.
However, using Atlas `Diff` hooks allow replacing the `DropColumn` and `AddColumn` changes with a `RenameColumn` change.

```go
func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    // ...
    if err := client.Schema.Create(ctx, schema.WithDiffHook(renameColumnHook)); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}

func renameColumnHook(next schema.Differ) schema.Differ {
    return schema.DiffFunc(func(current, desired *atlas.Schema) ([]atlas.Change, error) {
        changes, err := next.Diff(current, desired)
        if err != nil {
            return nil, err
        }
        for _, c := range changes {
            m, ok := c.(*atlas.ModifyTable)
            // Skip if the change is not a ModifyTable,
            // or if the table is not the "users" table.
            if !ok || m.T.Name != user.Table {
                continue
            }
            changes := atlas.Changes(m.Changes)
            switch i, j := changes.IndexDropColumn("old_name"), changes.IndexAddColumn("new_name"); {
            case i != -1 && j != -1:
                // Append a new renaming change.
                changes = append(changes, &atlas.RenameColumn{
                    From: changes[i].(*atlas.DropColumn).C,
                    To: changes[j].(*atlas.AddColumn).C,
                })
                // Remove the drop and add changes.
                changes.RemoveIndex(i, j)
                m.Changes = changes
            case i != -1 || j != -1:
                return nil, errors.New("old_name and new_name must be present or absent")
            }
        }
        return changes, nil
    })
}
```

#### `Apply` Hook Example

The `Apply` hook allows accessing and mutating the migration plan and its raw changes (SQL statements), but in addition
to that it is also useful for executing custom SQL statements before or after the plan is applied. For example, changing
a nullable column to non-nullable without a default value is not allowed by default. However, we can work around this
using an `Apply` hook that `UPDATE`s all rows that contain `NULL` value in this column:

```go
func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    // ...
    if err := client.Schema.Create(ctx, schema.WithApplyHook(fillNulls)); err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}

func fillNulls(next schema.Applier) schema.Applier {
	return schema.ApplyFunc(func(ctx context.Context, conn dialect.ExecQuerier, plan *migrate.Plan) error {
		// There are three ways to UPDATE the NULL values to "Unknown" in this stage.
		// Append a custom migrate.Change to the plan, execute an SQL statement directly
		// on the dialect.ExecQuerier, or use the ent.Client used by the project.

		// Execute a custom SQL statement.
		query, args := sql.Dialect(dialect.MySQL).
			Update(user.Table).
			Set(user.FieldDropOptional, "Unknown").
			Where(sql.IsNull(user.FieldDropOptional)).
			Query()
		if err := conn.Exec(ctx, query, args, nil); err != nil {
			return err
		}

		// Append a custom statement to migrate.Plan.
		//
		//  plan.Changes = append([]*migrate.Change{
		//	    {
		//		    Cmd: fmt.Sprintf("UPDATE users SET %[1]s = '%[2]s' WHERE %[1]s IS NULL", user.FieldDropOptional, "Unknown"),
		//	    },
		//  }, plan.Changes...)

		// Use the ent.Client used by the project.
		//
		//  drv := sql.NewDriver(dialect.MySQL, sql.Conn{ExecQuerier: conn.(*sql.Tx)})
		//  if err := ent.NewClient(ent.Driver(drv)).
		//  	User.
		//  	Update().
		//  	SetDropOptional("Unknown").
		//  	Where(/* Add predicate to filter only rows with NULL values */).
		//  	Exec(ctx); err != nil {
		//  	return fmt.Errorf("fix default values to uppercase: %w", err)
		//  }

		return next.Apply(ctx, conn, plan)
	})
}
```
