---
title: Quickly visualize your Ent schemas with entviz
author: Rotem Tamir
authorURL: "https://github.com/rotemtam"
authorImageURL: "https://s.gravatar.com/avatar/36b3739951a27d2e37251867b7d44b1a?s=80"
authorTwitter: _rtam
image: "https://entgo.io/images/assets/entviz-v2.png"
---

### TL;DR

To get a public link to a visualization of your Ent schema, run:

```
go run -mod=mod ariga.io/entviz ./path/to/ent/schema 
```

![](https://entgo.io/images/assets/erd/edges-quick-summary.png)

### Visualizing Ent schemas

Ent enables developers to build complex application data models
using [graph semantics](https://en.wikipedia.org/wiki/Graph_theory): instead of defining tables, columns, association
tables and foreign keys, Ent models are simply defined in terms of [Nodes](https://entgo.io/docs/schema-fields)
and [Edges](https://entgo.io/docs/schema-edges):

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
```

Modeling data this way has many benefits such as being able to
easily [traverse](https://entgo.io/docs/traversals) an application's data graph in an intuitive API, automatically
generating [GraphQL](https://entgo.io/docs/tutorial-todo-gql) servers and more.

While Ent can use a Graph database as its storage layer, most Ent users use common relational databases such as MySQL,
PostgreSQL or MariaDB for their applications. In these use-cases, developers often ponder, *what actual database schema
will Ent create from my application's schema?*

Whether you're a new Ent user learning the basics of creating Ent schemas or an expert dealing with optimizing the
resulting database schema for performance reasons, being able to easily visualize your Ent schema's backing database
schema can be very useful.

#### Introducing the new `entviz`

A year and a half ago
we [shared an Ent extension named entviz](https://entgo.io/blog/2021/08/26/visualizing-your-data-graph-using-entviz),
that extension enabled users to generate simple, local HTML documents containing entity-relationship diagrams describing
an application's Ent schema.

Today, we're happy to share a [super cool tool](https://github.com/ariga/entviz) by the same name created
by [Pedro Henrique (crossworth)](https://github.com/crossworth) which is a completely fresh take on the same problem.
With (the new) entviz you run a simple Go command:

```
go run -mod=mod ariga.io/entviz ./path/to/ent/schema 
```

The tool will analyze your Ent schema and create a visualization on the [Atlas Playground](https://gh.atlasgo.cloud) and
create a shareable, public [link](https://gh.atlasgo.cloud/explore/saved/60129542154) for you:

```
Here is a public link to your schema visualization:
	    https://gh.atlasgo.cloud/explore/saved/60129542154
```

In this link you will be able to see your schema visually as an ERD or textually as either a SQL
or [Atlas HCL](https://atlasgo.io/atlas-schema/sql-resources) document.

### Wrapping up

In this blog post we discussed some scenarios where you might find it useful to quickly get a visualization of your Ent
application's schema, we then showed how creating such visualizations can be achieved
using [entviz](https://github.com/ariga/entviz). If you like the idea, we'd be super happy if you tried it today and
gave us feedback!

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
  :::
