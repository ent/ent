---
title: Introducing ent
author: Ariel Mashraki
authorURL: "https://github.com/a8m"
authorImageURL: "https://avatars0.githubusercontent.com/u/7413593"
authorTwitter: arielmashraki
---
## The state of Go in Facebook Connectivity Tel Aviv
20 months ago, I joined Facebook Connectivity (FBC) team in Tel Aviv after ~5 years
of programming in Go and embedding it in a few companies.  
I joined a team that was working on a new project and we needed to choose a language
for this mission. We compared a few languages and decided to go with Go.

Since then, Go continued to spread across other FBC projects and became a big success
with around 15 Go engineers in Tel Aviv alone. **New services are now written in Go**.

## The motivation for writing a new ORM in Go

Most of my work in my 5 years before Facebook was on infra tooling and micro-services without
too much data-model work. A service that was needed to do a little amount of work with an SQL
database used one of the existing open-source solutions, but one that had worked with a
complicated data model was written in a different language with a robust ORM. For example,
Python with SQLAlchemy. 

At Facebook we like to think about our data-model in graph concepts. We've had a good experience
with this model internally.  
The lack of a proper Graph-based ORM for Go, led us to write one here with the following principles:

- **Schema As Code** - defining types, relations and constraints should be in Go code (not struct
  tags), and should be validated using a CLI tool. We have good experience with a similar tool
  internally at Facebook.
- **Statically typed and explicit API** using codegen - API with `interface{}`s everywhere affects
  developers efficiency; especially project newbies.
- **Queries, aggregations and graph traversals** should be simple - developers don’t want to deal
  with raw SQL queries nor SQL terms.
- **Predicates should be statically typed**. No strings everywhere.
- Full support for `context.Context` - This helps us to get full visibility in our traces and logs
  systems, and it’s important for other features like cancellation.
- **Storage agnostic** - we tried to keep the storage layer dynamic using codegen templates,
  since the development initially started on Gremlin (AWS Neptune) and switched later to MySQL.
  
## Open-sourcing ent

**ent** is an entity framework (ORM) for Go, built with the principles described above.
**ent** makes it possible to define any data model or graph-structure in Go code easily; The
schema configuration is verified by **entc** (the ent codegen) that generates an idiomatic and
statically-typed API that keeps Go developers productive and happy.
It supports MySQL, MariaDB, PostgreSQL, SQLite, and Gremlin-based graph databases.

We’re open-sourcing **ent** today, and invite you to get started → [entgo.io/docs/getting-started](/docs/getting-started).
