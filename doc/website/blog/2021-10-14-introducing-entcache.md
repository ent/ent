---
title: Announcing entcache - a Cache Driver for Ent
author: Ariel Mashraki
authorURL: "https://github.com/a8m"
authorImageURL: "https://avatars0.githubusercontent.com/u/7413593"
authorTwitter: arielmashraki
---

While working on [Ariga's](https://ariga.io) operational data graph query engine, we saw the opportunity to greatly
improve the performance of many use cases by building a robust caching library. As heavy users of Ent, it was only
natural for us to implement this layer as an extension to Ent. In this post, I will briefly explain what caches are,
how they fit into software architectures, and present `entcache` - a cache driver for Ent.

Caching is a popular strategy for improving application performance. It is based on the observation that the speed for
retrieving data using different types of media can vary within many orders of magnitude.
[Jeff Dean](https://twitter.com/jeffdean?lang=en) famously presented the following numbers in a
[lecture](http://static.googleusercontent.com/media/research.google.com/en/us/people/jeff/stanford-295-talk.pdf) about
"Software Engineering Advice from Building Large-Scale Distributed Systems":

![cache numbers](https://entgo.io/images/assets/entcache/cache-numbers.png)

These numbers show things that experienced software engineers know intuitively: reading from memory is faster than
reading from disk, retrieving data from the same data center is faster than going out to the internet to fetch it.
We add to that, that some calculations are expensive and slow, and that fetching a precomputed result can be much faster
(and less expensive) than recomputing it every time.

The collective intelligence of [Wikipedia](https://en.wikipedia.org/wiki/Cache_(computing)) tells us that a Cache is
"a hardware or software component that stores data so that future requests for that data can be served faster".
In other words, if we can store a query result in RAM, we can fulfill a request that depends on it much faster than
if we need to go over the network to our database, have it read data from disk, run some computation on it, and only
then send it back to us (over a network).

However, as software engineers, we should remember that caching is a notoriously complicated topic. As the phrase
coined by early-day Netscape engineer [Phil Karlton](https://martinfowler.com/bliki/TwoHardThings.html) says: _"There
are only two hard things in Computer Science: cache invalidation and naming things"_. For instance, in systems that rely
on strong consistency, a cache entry may be stale, therefore causing the system to behave incorrectly. For this reason,
take great care and pay attention to detail when you are designing caches into your system architectures.

### Presenting `entcache`

The `entcache` package provides its users with a new Ent driver that can wrap one of the existing SQL drivers available
for Ent. On a high level, it decorates the Query method of the given driver, and for each call:

1. Generates a cache key (i.e. hash) from its arguments (i.e. statement and parameters).

2. Checks the cache to see if the results for this query are already available. If they are (this is called a
   cache-hit), the database is skipped and results are returned to the caller from memory.

3. If the cache does not contain an entry for the query, the query is passed to the database.

4. After the query is executed, the driver records the raw values of the returned rows (`sql.Rows`), and stores them in
   the cache with the generated cache key.

The package provides a variety of options to configure the TTL of the cache entries, control the hash function, provide
custom and multi-level cache stores, evict and skip cache entries. See the full documentation in
[https://pkg.go.dev/ariga.io/entcache](https://pkg.go.dev/ariga.io/entcache).

As we mentioned above, correctly configuring caching for an application is a delicate task, and so `entcache` provides
developers with different caching levels that can be used with it:

1. A `context.Context`-based cache. Usually, attached to a request and does not work with other cache levels.
   It is used to eliminate duplicate queries that are executed by the same request.

2. A driver-level cache used by the `ent.Client`. An application usually creates a driver per database,
   and therefore, we treat it as a process-level cache.

3. A remote cache. For example, a Redis database that provides a persistence layer for storing and sharing cache
   entries between multiple processes. A remote cache layer is resistant to application deployment changes or failures,
   and allows reducing the number of identical queries executed on the database by different process.

4. A cache hierarchy, or multi-level cache allows structuring the cache in hierarchical way. The hierarchy of cache
   stores is mostly based on access speeds and cache sizes. For example, a 2-level cache that composed of an LRU-cache
   in the application memory, and a remote-level cache backed by a Redis database.

Let's demonstrate this by explaining the `context.Context` based cache.

### Context-Level Cache

The `ContextLevel` option configures the driver to work with a `context.Context` level cache. The context is usually
attached to a request (e.g. `*http.Request`) and is not available in multi-level mode. When this option is used as
a cache store, the attached `context.Context` carries an LRU cache (can be configured differently), and the driver
stores and searches entries in the LRU cache when queries are executed.

This option is ideal for applications that require strong consistency, but still want to avoid executing duplicate
database queries on the same request. For example, given the following GraphQL query:

```graphql
query($ids: [ID!]!) {
    nodes(ids: $ids) {
        ... on User {
            id
            name
            todos {
                id
                owner {
                    id
                    name
                }
            }
        }
    }
}
```

A naive solution for resolving the above query will execute, 1 for getting N users, another N queries for getting
the todos of each user, and a query for each todo item for getting its owner (read more about the
[_N+1 Problem_](https://entgo.io/docs/tutorial-todo-gql-field-collection/#problem)).

However, Ent provides a unique approach for resolving such queries(read more in
[Ent website](https://entgo.io/docs/tutorial-todo-gql-field-collection)) and therefore, only 3 queries will be executed
in this case. 1 for getting N users, 1 for getting the todo items of **all** users, and 1 query for getting the owners
of **all** todo items.

With `entcache`, the number of queries may be reduced to 2, as the first and last queries are identical (see
[code example](https://github.com/ariga/entcache/blob/master/internal/examples/ctxlevel/main_test.go)).

![context-level-cache](https://entgo.io/images/assets/entcache/ctxlevel.png)

The different levels are explained in depth in the repository
[README](https://github.com/ariga/entcache/blob/master/README.md).

### Getting Started

> If you are not familiar with how to set up a new Ent project, complete Ent
> [Setting Up tutorial](https://entgo.io/docs/tutorial-setup) first.

First, `go get` the package using the following command.

```shell
go get ariga.io/entcache
```

After installing `entcache`, you can easily add it to your project with the snippet below:

```go
// Open the database connection.
db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
if err != nil {
	log.Fatal("opening database", err)
}
// Decorates the sql.Driver with entcache.Driver.
drv := entcache.NewDriver(db)
// Create an ent.Client.
client := ent.NewClient(ent.Driver(drv))

// Tell the entcache.Driver to skip the caching layer
// when running the schema migration.
if client.Schema.Create(entcache.Skip(ctx)); err != nil {
	log.Fatal("running schema migration", err)
}

// Run queries.
if u, err := client.User.Get(ctx, id); err != nil {
	log.Fatal("querying user", err)
}
// The query below is cached.
if u, err := client.User.Get(ctx, id); err != nil {
	log.Fatal("querying user", err)
}
```

To see more advanced examples, head over to the repo's
[examples directory](https://github.com/ariga/entcache/tree/master/internal/examples).

### Wrapping Up

In this post, I presented “entcache” a new cache driver for Ent that I developed while working on [Ariga's Operational
Data Graph](https://ariga.io) query engine. We started the discussion by briefly mentioning the motivation for including
caches in software systems. Following that, we described the features and capabilities of `entcache` and concluded with
a short example of how you can set it up in your application.

There are a few features we are working on, and wish to work on, but need help from the community to design them
properly (solving cache invalidation, anyone? ;)). If you are interested to contribute, reach out to me on the Ent
Slack channel.

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
