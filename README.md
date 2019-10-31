## ent - An Entity Framework For Go

<img width="50%" 
align="right"
style="display: block; margin:40px auto;" 
src="https://s3.eu-central-1.amazonaws.com/entgo.io/assets/gopher_graph.png"/>

Simple, yet powerful ORM for modeling and querying data.

- **Schema As Code** - model any graph schema as Go objects.
- **Easily Traverse Any Graph** - run queries, aggregations and traverse any graph structure easily.
- **Statically Typed And Explicit API** - 100% statically typed and explicit API using code generation.
- **Multi Storage Driver** - supports MySQL, PostgreSQL, SQLite and Gremlin.

## Quick Installation
```console
go get github.com/facebookincubator/ent/cmd/entc
```

For proper installation using [Go modules], visit [entgo.io website][entgo instal].

## Docs
The documentation for developing and using ent is available at: https://entgo.io

## Join the ent Community
See the [CONTRIBUTING](CONTRIBUTING.md) file for how to help out.

## Project Status
`ent` was developed and maintained by [a8m](https://github.com/a8m) and [alexsn](https://github.com/alexsn)
from the Facebook Connectivity team. It's currently considered experimental (although we're using it in production),
and the roadmap for v1 release is described [here](https://github.com/facebookincubator/ent/issues/46).  
Read more about the motivation of the project [here](https://entgo.io/blog/2019/10/03/introducing-ent).

## License
ent is licensed under Apache 2.0 as found in the [LICENSE file](LICENSE).


[entgo instal]: https://entgo.io/docs/code-gen/#version-compatibility-between-entc-and-ent
[Go modules]: https://github.com/golang/go/wiki/Modules#quick-start
