---
id: migrate_versioned
title: Versioned Migrations
---

If you are using the Atlas migration engine you are able to use the versioned migrations feature of it. 

## Configuration

In order to have Ent make the necessary changes to your code, you have to enable this feature with one of the two
options:

1. If you are using the default go generate configuration, simply add the `--feature sql/versioned-migrations` to
   the `ent/generate.go` file as follows:

```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --feature sql/versioned-migrations ./schema
```

2. If you are using the code generation package (e.g. if you are using an Ent extension), add the feature flag as
   follows:

```go
//go:build ignore

package main

import (
	"log"
	
	"entgo.io/ent/entc"
    "entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{}, entc.FeatureNames("sql/versioned-migrations"))
	if err != nil {
       log.Fatalf("running ent codegen: %v", err)
    }
}
```

## Generating Versioned Migration Files

This will add an extra `Diff` method to the Ent client that you can use to inspect the connected database, compare it
with the schema definitions and create sql statements needed to migrate the database to the graph.

```go
package main

import (
    "context"
    "log"

    "<project>/ent"

    "ariga.io/atlas/sql/migrate"
    "entgo.io/ent/dialect/sql/schema"
)

func main() {
    client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/test")
    if err != nil {
        log.Fatalf("failed connecting to mysql: %v", err)
    }
    defer client.Close()
    ctx := context.Background()
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }
    // Run migration.
    err = client.Schema.Diff(ctx, schema.WithDir(dir))
    if err != nil {
        log.Fatalf("failed creating schema resources: %v", err)
    }
}
```

## Apply Migrations

The Atlas migration engine does not support applying the migration files onto a database yet, therefore to manage and
execute the generated migration files, you have to rely on an external tool (or execute them by hand). By default, Atlas
generates one "up" and one "down" migration file for the computed diff. These files are compatible with the popular
[golang-migrate/migrate](https://github.com/golang-migrate/migrate) package, and you can use that tool to manage the
migrations in you deployments.
