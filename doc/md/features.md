---
id: feature-flags
title: Feature Flags
sidebar_label: Feature Flags
---

The framework provides a collection of code-generation features that be added or removed using flags.

## Usage

Feature flags can be provided either by CLI flags or as arguments to the `gen` package. 

#### CLI

```console
go run entgo.io/ent/cmd/ent generate --feature privacy,entql ./ent/schema
```

#### Go

```go
// +build ignore

package main

import (
	"log"
	"text/template"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeaturePrivacy,
			gen.FeatureEntQL,
		},
		Templates: []*gen.Template{
			gen.MustParse(gen.NewTemplate("static").
				Funcs(template.FuncMap{"title": strings.ToTitle}).
				ParseFiles("template/static.tmpl")),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

## List of Features

#### Privacy Layer

The privacy layer allows configuring privacy policy for queries and mutations of entities in the database.

This option can added to a project using the `--feature privacy` flag, and its full documentation exists
in the [privacy page](privacy.md).

#### EntQL Filtering

The `entql` option provides a generic and dynamic filtering capability at runtime for the different query builders.

This option can be added to a project using the `--feature entql` flag, and more information about it exists
in the [privacy page](privacy.md#multi-tenancy).

#### Auto-Solve Merge Conflicts

The `schema/snapshot` option tells `entc` (ent codegen) to store a snapshot of the latest schema in an internal package,
and use it to automatically solve merge conflicts when user's schema can't be built.

This option can be added to a project using the `--feature schema/snapshot` flag, but please see
[ent/ent/issues/852](https://github.com/ent/ent/issues/852) to get more context about it.

#### Schema Config

The `sql/schemaconfig` option lets you pass alternate SQL database names to models. This is useful when your models don't all live under one database and are spread out across different schemas.

This option can be added to a project using the `--feature sql/schemaconfig` flag. Once you generate the code, you can now use a new option as such: 

```go
c, err := ent.Open(dialect, conn, ent.AlternateSchema(ent.SchemaConfig{
	User: "usersdb",
	Car: "carsdb",
}))
c.User.Query().All(ctx) // SELECT * FROM `usersdb`.`users`
c.Car.Query().All(ctx) 	// SELECT * FROM `carsdb`.`cars`
```

#### Row-level Locks

The `sql/lock` option lets configure row-level locking using the SQL `SELECT ... FOR {UPDATE | SHARE}` syntax.

This option can be added to a project using the `--feature sql/lock` flag.

```go
tx, err := client.Tx(ctx)
if err != nil {
	log.Fatal(err)
}

tx.Pet.Query().
	Where(pet.Name(name)).
	ForUpdate().
	Only(ctx)

tx.Pet.Query().
	Where(pet.ID(id)).
	ForShare(
		sql.WithLockTables(pet.Table),
		sql.WithLockAction(sql.NoWait),
	).
	Only(ctx)
```

#### Custom SQL Modifiers

The `sql/modifier` option lets add custom SQL modifiers to the builders and mutate the statements before they are executed.

This option can be added to a project using the `--feature sql/modifier` flag.

```go
client.Pet.
	Query().
	Modify(func(s *sql.Selector) {
		s.Select("SUM(LENGTH(name))")
	}).
	IntX(ctx)

// SELECT SUM(LENGTH(name)) FROM `pet`

var v []struct {
	Count     int       `json:"count"`
	Price     int       `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}
client.User.
	Query().
	Where(
        user.CreatedAtGT(x),
        user.CreatedAtLT(y),
	).
	Modify(func(s *sql.Selector) {
		s.Select(
			sql.As(sql.Count("*"), "count"),
			sql.As(sql.Sum("price"), "price"),
			sql.As("DATE(created_at)", "created_at"),
		).
		GroupBy("DATE(created_at)").
		OrderBy(sql.Desc("DATE(created_at)"))
	}).
	ScanX(ctx, &v)

// SELECT COUNT(*) AS `count`, SUM(`price`) AS `price`, DATE(created_at) AS `created_at`
// FROM `users` WHERE created_at > x AND created_at < y
// GROUP BY DATE(created_at)
// ORDER BY DATE(created_at) DESC
```

#### Upsert

The `sql/upsert` option lets configure upsert and bulk-upsert logic using the SQL `ON CONFLICT` / `ON DUPLICATE KEY`
syntax. For full documentation, go to the [Upsert API](crud.md#upsert-one).

This option can be added to a project using the `--feature sql/upsert` flag.

```go
// Use the new values that were set on create.
id, err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflict().
	UpdateNewValues().
	ID(ctx)

// In PostgreSQL, the conflict target is required.
err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflictColumns(user.FieldName).
	UpdateNewValues().
	Exec(ctx)

// Bulk upsert is also supported.
client.User.
	CreateBulk(builders...).
	OnConflict(
		sql.ConflictWhere(...),
		sql.UpdateWhere(...),
	).
	UpdateNewValues().
	Exec(ctx)

// INSERT INTO "users" (...) VALUES ... ON CONFLICT WHERE ... DO UPDATE SET ... WHERE ...
```
