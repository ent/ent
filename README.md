## ent - An Entity Framework For Go

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
go get github.com/facebook/ent/cmd/entc
```

For proper installation using [Go modules], visit [entgo.io website][entgo instal].

## Docs
The documentation for developing and using ent is available at: https://entgo.io

## Join the ent Community
In order to contribute to `ent`, see the [CONTRIBUTING](CONTRIBUTING.md) file for how to go get started.  
If your company or your product is using `ent`, please let us know by adding yourself to the [ent users page](https://github.com/facebook/ent/wiki/ent-users).

## About the Project
The `ent` project was inspired by Ent, an entity framework we use internally. It is developed and maintained
by [a8m](https://github.com/a8m) and [alexsn](https://github.com/alexsn)
from the [Facebook Connectivity][fbc] team. It is used by multiple teams and projects in production,
and the roadmap for its v1 release is described [here](https://github.com/facebook/ent/issues/46). 
Read more about the motivation of the project [here](https://entgo.io/blog/2019/10/03/introducing-ent).

## License
ent is licensed under Apache 2.0 as found in the [LICENSE file](LICENSE).


[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
[fbc]: https://connectivity.fb.com
