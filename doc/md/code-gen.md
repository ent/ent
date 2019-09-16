---
id: code-gen
title: Introduction
---

## Installation

`ent` comes with a codegen tool called `entc`. In order to install
`entc` run the following command:

```bash
go get github.com/facebookincubator/ent/entc/cmd/entc
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
the assets for working with your entities. Run the following command:

```bash
entc generate ./ent/schema
```

You should note that `goimports` is required for the codegen, and it can be installed using:

```bash
go get -u golang.org/x/tools/cmd/goimports
``` 

The `generate` command generates the following assets for the schemas:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.md) for more info.
- Entity object (Go struct) for each of the schema types.
- Package containing constants and predicates used for interacting with the builders.
- A `migrate` package for SQL dialects. See [Migration](migrate.md) for more info.

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
a file with the same name as the template.

Example of a custom template provides a `Node` API for GraphQL - 
[Github](https://github.com/facebookincubator/ent/blob/master/entc/integration/template/ent/template/node.tmpl).

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