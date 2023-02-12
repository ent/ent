---
id: code-gen
title: Introduction
---

## Installation

The project comes with a codegen tool called `ent`. In order to install
`ent` run the following command:

```bash
go get -d entgo.io/ent/cmd/ent
``` 

## Initialize A New Schema

In order to generate one or more schema templates, run `ent init` as follows:

```bash
go run -mod=mod entgo.io/ent/cmd/ent new User Pet
```

`init` will create the 2 schemas (`user.go` and `pet.go`) under the `ent/schema` directory.
If the `ent` directory does not exist, it will create it as well. The convention
is to have an `ent` directory under the root directory of the project.

## Generate Assets

After adding a few [fields](schema-fields.md) and [edges](schema-edges), you want to generate
the assets for working with your entities. Run `ent generate` from the root directory of the project,
or use `go generate`:


```bash
go generate ./ent
```

The `generate` command generates the following assets for the schemas:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.mdx) for more info.
- Entity object (Go struct) for each of the schema types.
- Package containing constants and predicates used for interacting with the builders.
- A `migrate` package for SQL dialects. See [Migration](migrate.md) for more info.
- A `hook` package for adding mutation middlewares. See [Hooks](hooks.md) for more info.

## Version Compatibility Between `entc` And `ent`

When working with `ent` CLI in a project, you want to make sure the version being
used by the CLI is **identical** to the `ent` version used by your project.

One of the options for achieving this is asking `go generate` to use the version
mentioned in the `go.mod` file when running `ent`. If your project does not use
[Go modules](https://github.com/golang/go/wiki/Modules#quick-start), setup one as follows:

```console
go mod init <project>
```

And then, re-run the following command in order to add `ent` to your `go.mod` file:

```console
go get -d entgo.io/ent/cmd/ent
```

Add a `generate.go` file to your project under `<project>/ent`:

```go
package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate ./schema
```

Finally, you can run `go generate ./ent` from the root directory of your project
in order to run `ent` code generation on your project schemas.

## Code Generation Options

For more info about codegen options, run `ent generate -h`:

```console
generate go code for the schema directory

Usage:
  ent generate [flags] path

Examples:
  ent generate ./ent/schema
  ent generate github.com/a8m/x

Flags:
      --feature strings       extend codegen with additional features
      --header string         override codegen header
  -h, --help                  help for generate
      --idtype [int string]   type of the id field (default int)
      --storage string        storage driver to support in codegen (default "sql")
      --target string         target directory for codegen
      --template strings      external templates to execute
```

## Storage Options

`ent` can generate assets for both SQL and Gremlin dialect. The default dialect is SQL.

## External Templates

`ent` accepts external Go templates to execute. If the template name already defined by
`ent`, it will override the existing one. Otherwise, it will write the execution output to
a file with the same name as the template. The flag format supports  `file`, `dir` and `glob`
as follows:

```console
go run -mod=mod entgo.io/ent/cmd/ent generate --template <dir-path> --template glob="path/to/*.tmpl" ./ent/schema
```

More information and examples can be found in the [external templates doc](templates.md).

## Use `entc` as a Package

Another option for running `ent` code generation is to create a file named `ent/entc.go` with the following content,
and then the `ent/generate.go` file to execute it:

```go title="ent/entc.go"
// +build ignore

package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
)

func main() {
	if err := entc.Generate("./schema", &gen.Config{}); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

```go title="ent/generate.go"
package ent

//go:generate go run -mod=mod entc.go
```

The full example exists in [GitHub](https://github.com/ent/ent/tree/master/examples/entcpkg).

## Schema Description

In order to get a description of your graph schema, run:

```bash
go run -mod=mod entgo.io/ent/cmd/ent describe ./ent/schema
```

An example for the output is as follows:

```console
Pet:
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| Field |  Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| id    | int     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
	| name  | string  | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          0 |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	+-------+------+---------+---------+----------+--------+----------+
	| Edge  | Type | Inverse | BackRef | Relation | Unique | Optional |
	+-------+------+---------+---------+----------+--------+----------+
	| owner | User | true    | pets    | M2O      | true   | true     |
	+-------+------+---------+---------+----------+--------+----------+
	
User:
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| Field |  Type   | Unique | Optional | Nillable | Default | UpdateDefault | Immutable |       StructTag       | Validators |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	| id    | int     | false  | false    | false    | false   | false         | false     | json:"id,omitempty"   |          0 |
	| age   | int     | false  | false    | false    | false   | false         | false     | json:"age,omitempty"  |          0 |
	| name  | string  | false  | false    | false    | false   | false         | false     | json:"name,omitempty" |          0 |
	+-------+---------+--------+----------+----------+---------+---------------+-----------+-----------------------+------------+
	+------+------+---------+---------+----------+--------+----------+
	| Edge | Type | Inverse | BackRef | Relation | Unique | Optional |
	+------+------+---------+---------+----------+--------+----------+
	| pets | Pet  | false   |         | O2M      | false  | true     |
	+------+------+---------+---------+----------+--------+----------+
```

## Code Generation Hooks

The `entc` package provides an option to add a list of hooks (middlewares) to the code-generation phase.
This option is ideal for adding custom validators for the schema, or for generating additional assets
using the graph schema.

```go
// +build ignore

package main

import (
	"fmt"
	"log"
	"reflect"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Hooks: []gen.Hook{
			EnsureStructTag("json"),
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

// EnsureStructTag ensures all fields in the graph have a specific tag name.
func EnsureStructTag(name string) gen.Hook {
	return func(next gen.Generator) gen.Generator {
		return gen.GenerateFunc(func(g *gen.Graph) error {
			for _, node := range g.Nodes {
				for _, field := range node.Fields {
					tag := reflect.StructTag(field.StructTag)
					if _, ok := tag.Lookup(name); !ok {
						return fmt.Errorf("struct tag %q is missing for field %s.%s", name, node.Name, field.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}
```

## External Dependencies

In order to extend the generated client and builders under the `ent` package, and inject them external
dependencies as struct fields, use the `entc.Dependency` option in your [`ent/entc.go`](#use-entc-as-a-package)
file:

```go title="ent/entc.go" {3-12}
func main() {
	opts := []entc.Option{
		entc.Dependency(
			entc.DependencyType(&http.Client{}),
		),
		entc.Dependency(
			entc.DependencyName("Writer"),
			entc.DependencyTypeInfo(&field.TypeInfo{
				Ident:   "io.Writer",
				PkgPath: "io",
			}),
		),
	}
	if err := entc.Generate("./schema", &gen.Config{}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
```

Then, use it in your application:

```go title="example_test.go" {5-6,15-16}
func Example_Deps() {
	client, err := ent.Open(
		"sqlite3",
		"file:ent?mode=memory&cache=shared&_fk=1",
		ent.Writer(os.Stdout),
		ent.HTTPClient(http.DefaultClient),
	)
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// An example for using the injected dependencies in the generated builders.
	client.User.Use(func(next ent.Mutator) ent.Mutator {
		return hook.UserFunc(func(ctx context.Context, m *ent.UserMutation) (ent.Value, error) {
			_ = m.HTTPClient
			_ = m.Writer
			return next.Mutate(ctx, m)
		})
	})
	// ...
}
```

The full example exists in [GitHub](https://github.com/ent/ent/tree/master/examples/entcpkg).

## Feature Flags

The `entc` package provides a collection of code-generation features that be added or removed using flags.

For more information, please see the [features-flags page](features.md).
