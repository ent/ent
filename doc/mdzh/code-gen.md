---
id: code-gen
title: Introduction
---

## 下载

这个项目有一个叫做`ent`的代码生成工具。下载`ent`需要执行如下命令：

```bash
go get entgo.io/ent/cmd/ent
``` 

## 初始化新的Schema

生成一个或多个schema模板，运行`ent init`命令，如下：

```bash
go run entgo.io/ent/cmd/ent init User Pet
```

`init`将在`ent/schema`目录下创建2个schema(`user.go` 和 `pet.go`)，如果`ent`目录不存在，它也会创建。惯例是在项目的根目录下有一个`ent`目录。

## 生成Assets

在给实体添加了[字段](schema-fields.md)以及[各实体之间的关系](schema-edges.md)之后。你要生成一些用来操作实体的资源文件。
项目根目录下运行`ent generate`命令或者`go generate` 

```bash
go generate ./ent
```

`generate` 命令为schemas生成了如下资源：

- `Client` 和 `Tx` 用于图形之间的交互
- 每个schema类型的CRUD构建器。更多内容阅读[CRUD](crud.md) 
- 每个schema类型的实体对象(Go struct)
- 包含用于与构建器交互的常量和断言的包
- `migrate` 一个进行数据迁移的包，更多内容阅读[Migration](migrate.md)

## `entc` 和 `ent`的版本兼容问题

当在项目中使用`ent`命令时，你要确保它的版本号和项目中使用的`ent`的版本要**完全一致**。

当运行`ent`时，执行`go generate`的选项之一是使用 `go.mod`文件提及到的版本。
如果你的项目没有使用[Go modules](https://github.com/golang/go/wiki/Modules#quick-start)，执行如下命令：

```console
go mod init <project>
```

然后，重新运行如下命令，为了将`ent` 添加到`go.mod`中

```console
go get entgo.io/ent/cmd/ent
```

在`<project>/ent`下添加一个`generate.go`文件。

```go
package ent

//go:generate go run entgo.io/ent/cmd/ent generate ./schema
```

最后，在项目根目录下执行`go generate ./ent`，可以在你项目的schemas下生成可以运行的`ent`代码。

## 代码生成选项

关于更多的代码生成的选项，执行`ent generate -h`:

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

## 存储选项

`ent` 能够生成SQL和Gremlin(Gremlin 是在某些领域专用的语言，用来遍历属性图)的资源。默认是SQL。

## 外部模板

`ent` 接受外部Go模板去执行。如果模板名已经通过`ent`定义，将会覆盖掉之前的模板。否则，它将把执行输出写入与模板同名的文件中。
标志格式支持 `file`, `dir` 和 `glob`

```console
go run entgo.io/ent/cmd/ent generate --template <dir-path> --template glob="path/to/*.tmpl" ./ent/schema
```

更多的信息和例子请参考[external templates doc](templates.md)

## Use `entc` As A Package

运行`ent` 命令行的另一个选项是将它作为一个包使用，如下所示:

```go
package main

import (
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"entgo.io/ent/schema/field"
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

所有的案例在[GitHub](https://github.com/facebook/ent/tree/master/examples/entcpkg)

## Schema描述

得到图模式的描述，执行：

```bash
go run entgo.io/ent/cmd/ent describe ./ent/schema
```

输出如下：

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

## 代码生成Hooks

`entc` 包提供了一个选项，可以将钩子(中间件)列表添加到代码生成阶段。这个选项非常适合为schema添加自定义验证器，或者生成额外的assets
使用图形schema。

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
						return fmt.Errorf("struct tag %q is missing for field %s.%s", name, node.Name, f.Name)
					}
				}
			}
			return next.Generate(g)
		})
	}
}
```

## 特征标识

`entc`包提供了代码成的特性集合，能够使用标识添加或者删除。

更多的信息，请阅读[features-flags page](features.md)。
