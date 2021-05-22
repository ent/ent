## ent - An Entity Framework For Go

[![Twitter](https://img.shields.io/twitter/url/https/twitter.com/entgo_io.svg?style=social&label=Follow%20%40entgo_io)](https://twitter.com/entgo_io)

[English](README.md) | [中文](README_zh.md) | [日本語](README_jp.md)

<img width="50%"
align="right"
style="display: block; margin:40px auto;"
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

Simple, yet powerful entity framework for Go, that makes it easy to build and maintain applications
with large data-models.

- **Schema As Code** - model any database schema as Go objects.
- **Easily Traverse Any Graph** - run queries, aggregations and traverse any graph structure easily.
- **Statically Typed And Explicit API** - 100% statically typed and explicit API using code generation.
- **Multi Storage Driver** - supports MySQL, PostgreSQL, SQLite and Gremlin.
- **Extendable** - simple to extend and customize using Go templates.

## Quick Installation
```console
go get entgo.io/ent/cmd/ent
```

For proper installation using [Go modules], visit [entgo.io website][entgo instal].

## Docs and Support
The documentation for developing and using ent is available at: https://entgo.io

For discussion and support, [open an issue](https://github.com/ent/ent/issues/new/choose) or join our [channel](https://gophers.slack.com/archives/C01FMSQDT53) in the gophers Slack.

## Join the ent Community
Building `ent` would not have been possible without the collective work of our entire community. We maintain a [contributors page](doc/md/contributors.md)
which lists the contributors to this `ent`. 

In order to contribute to `ent`, see the [CONTRIBUTING](CONTRIBUTING.md) file for how to go get started.
If your company or your product is using `ent`, please let us know by adding yourself to the [ent users page](https://github.com/ent/ent/wiki/ent-users).

For updates, follow us on Twitter at https://twitter.com/entgo_io



## About the Project
The `ent` project was inspired by Ent, an entity framework we use internally. It is developed and maintained
by [a8m](https://github.com/a8m) and [alexsn](https://github.com/alexsn)
from the [Facebook Connectivity][fbc] team. It is used by multiple teams and projects in production,
and the roadmap for its v1 release is described [here](https://github.com/ent/ent/issues/46).
Read more about the motivation of the project [here](https://entgo.io/blog/2019/10/03/introducing-ent).

## License
ent is licensed under Apache 2.0 as found in the [LICENSE file](LICENSE).


[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
