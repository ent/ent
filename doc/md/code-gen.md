---
id: code-gen
title: Introduction
---

## Installation

`ent` comes with a codegen tool called `entc`. In order to install
`entc` run the following command:

```bash
go get github.com/facebook/ent/cmd/entc
``` 

## Initialize A New Schema

In order to generate one or more schema templates, run `entc init` as follows:

```bash
entc init User Pet
```

`init` will create the 2 schemas (`user.go` and `pet.go`) under the `ent/schema` directory.
If the `ent` directory does not exist, it will create it as well. The convention
is to have an `ent` directory under the root directory of the project.

## Generate Assets

After adding a few [fields](schema-fields.md) and [edges](schema-edges.md), you want to generate
the assets for working with your entities. Run `entc generate` from the root directory of the project,
or use `go generate`:


```bash
go generate ./ent
```

The `generate` command generates the following assets for the schemas:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.md) for more info.
- Entity object (Go struct) for each of the schema types.
- Package containing constants and predicates used for interacting with the builders.
- A `migrate` package for SQL dialects. See [Migration](migrate.md) for more info.

## Version Compatibility Between `entc` And `ent`

When working with `entc` in a project, you want to make sure that the version being
used by `entc` is **identical** to the `ent` version used by your project.

One of the options for achieving this is asking `go generate` to use the version
mentioned in the `go.mod` file when running `entc`. If your project does not use
[Go modules](https://github.com/golang/go/wiki/Modules#quick-start), setup one as follows:

```console
go mod init <project>
```

And then, re-run the following command in order to add `ent` to your `go.mod` file:

```console
go get github.com/facebook/ent/cmd/entc
```

Add a `generate.go` file to your project under `<project>/ent`:

```go
package ent

//go:generate go run github.com/facebook/ent/cmd/entc generate ./schema
```

Finally, you can run `go generate ./ent` from the root directory of your project
in order to run `entc` code generation on your project schemas.

## Code Generation Options

For more info about codegen options, run `entc generate -h`:

```console
generate go code for the schema directory

Usage:
  entc generate [flags] path

Examples:
  entc generate ./ent/schema
  entc generate github.com/a8m/x

Flags:
      --header string         override codegen header
  -h, --help                  help for generate
      --idtype [int string]   type of the id field (default int)
      --storage strings       list of storage drivers to support (default [sql])
      --target string         target directory for codegen
      --template strings      external templates to execute
```

## Storage Options

`entc` can generate assets for both SQL and Gremlin dialect. The default dialect is SQL.

## External Templates

`entc` accepts external Go templates to execute. If the template name is already defined by
`entc`, it will override the existing one. Otherwise, it will write the execution output to
a file with the same name as the template. The flag format supports  `file`, `dir` and `glob`
as follows:

```console
entc generate --template <dir-path> --template glob="path/to/*.tmpl" ./ent/schema
```

More information and examples can be found in the [external templates doc](templates.md).

## Use `entc` As A Package

Another option for running `entc` is to use it as a package as follows:

```go
package main

import (
	"log"

	"github.com/facebook/ent/entc"
	"github.com/facebook/ent/entc/gen"
	"github.com/facebook/ent/schema/field"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Header: "// Your Custom Header",
		IDType: &field.TypeInfo{Type: field.TypeInt},
	})
	if err != nil {
		log.Fatal("running ent codegen:", err)
	}
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/entcpkg).


## Schema Description

In order to get a description of your graph schema, run:

```bash
entc describe ./ent/schema
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
