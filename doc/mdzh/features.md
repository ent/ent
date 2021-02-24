---
id: feature-flags
title: Feature Flags
sidebar_label: Feature Flags
---

此框架提供一系列代码生产特性，这些特性可以使用标志位添加或者移除。

## 使用

特性标志位能通过CLI标识或者作为`gen`包的参数被提供。

#### 命令行

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
		Features: []*gen.Feature{
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

## 特性列表

#### 私有层

私有层允许配置针对数据库实体的查询及变化的私有策略

这个选项能够使用`--feature privacy`标识被添加到项目。它的完整文档在[privacy page](privacy.md)。

#### EntQL过滤器

`entql`选项在运行时为不同的查询构建器提供了通用的动态过滤功能。

这个选项能够使用`--feature entql`标识被添加到项目，更多的信息在[privacy page](privacy.md#multi-tenancy)

#### 自动解决合并冲突

`schema/snapshot`选项告诉`entc`（ent生成）在内部的包去存储最新的schema的快照，当用户的schema不能建造时，使用它会自动解决合并冲突。

这个选项能够使用`--feature schema/snapshot`标识被添加到项目，但是想要获取更多的信息请看[facebook/ent/issues/852](https://entgo.io/ent/issues/852)

#### Schema配置

`sql/schemaconfig`选项允许将备用sql数据库名称传递给模型。当模型不是都在一个数据库中，而是分布在不同的模式中时，这是非常有用的。

这个选项能够使用`--feature sql/schemaconfig`标识被添加到项目，一旦你生成了代码，可以使用这样的新选项:

```golang
c, err := ent.Open(dialect, conn, ent.AlternateSchema(ent.SchemaConfig{
	User: "usersdb",
	Car: "carsdb",
}))
c.User.Query().All(ctx) // SELECT * FROM `usersdb`.`users`
c.Car.Query().All(ctx) 	// SELECT * FROM `carsdb`.`cars`
```
