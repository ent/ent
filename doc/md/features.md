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

This option can be added to projects using the `--feature privacy` flag, and its full documentation exists
in the [privacy page](privacy.md).

#### EntQL Filtering

The `entql` option provides a generic and dynamic filtering capability at runtime for the different query builders.

This option can be added to projects using the `--feature entql` flag, and more information about it exists
in the [privacy page](privacy.md#multi-tenancy).

#### Auto-Solve Merge Conflicts

The `schema/snapshot` option tells `entc` (ent codegen) to store a snapshot of the latest schema in an internal package,
and use it to automatically solve merge conflicts when user's schema can't be built.

This option can be added to projects using the `--feature schema/snapshot` flag, but please see
[ent/ent/issues/852](https://github.com/ent/ent/issues/852) to get more context about it.

#### Schema Config

The `sql/schemaconfig` option lets you pass alternate SQL database names to models. This is useful when your models don't all live under one database and are spread out across different schemas.

This option can be added to projects using the `--feature sql/schemaconfig` flag. Once you generate the code, you can now use a new option as such: 

```go
c, err := ent.Open(dialect, conn, ent.AlternateSchema(ent.SchemaConfig{
	User: "usersdb",
	Car: "carsdb",
}))
c.User.Query().All(ctx) // SELECT * FROM `usersdb`.`users`
c.Car.Query().All(ctx) 	// SELECT * FROM `carsdb`.`cars`
```
