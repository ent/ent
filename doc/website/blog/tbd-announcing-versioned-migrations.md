---
title: Announcing Versioned Migrations
author: MasseElch
authorURL: "https://github.com/masseelch"
authorImageURL: "https://avatars.githubusercontent.com/u/12862103?v=4"
image: tbd
---

When [Ariel](https://github.com/a8m) released Ent v0.10.0 end of January,
he [introduced](2022-01-20-announcing-new-migration-engine.md) a new migration engine for Ent based on another
OpenSource tool called [Atlas](https://github.com/ariga/atlas). A month ago we added support for Atlas to generate
versioned migration files to be used with popular migration management solutions
like [golang-migrate/migrate](https://github.com/golang-migrate/migrate), [Flyway](https://flywaydb.org/)
or [Liquibase](https://liquibase.org/).

In this post I want to show you how to configure Ent to generate versioned migration files for both existing and new
projects. Furthermore, I will demonstrate the workflow with `golang-migrate/migrate`. 

### Getting Started

The very first thing to do, is to make sure you have an up-to-date Ent version:

```shell
go get -u entgo.io/ent@master
```

There are two ways to have Ent generate migration files for schema changes. The first one is to use an instantiated Ent
client and the second one to generate the changes from a parsed schema graph. This post will take the second approach,
if you want to learn how to use the first one you can have a look at
the [documentation](./docs/versioned-migrations#from-client).

### Generating Versioned Migration Files

Since we have enabled the versioned migrations feature now, let's create a small schema and generate the initial set of
migration files. Consider the following schema for a fresh Ent project:

```go title="ent/schema/user.go"
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email").Unique(),
	}
}
```

As I told you earlier, wa want to use the parsed schema graph to compute the difference between our schema and the
connected database. Here is an example of a (semi-)persistent MySQL docker container to use if you want to follow along:

```shell
docker run --rm --name ent-versioned-migrations --detach --env MYSQL_ROOT_PASSWORD=pass --env MYSQL_DATABASE=ent -p 3306:3306 mysql
```

Once you are done you can shut down the container and remove all resources with `docker stop ent-versioned-migrations`.

Now, let's create a small function, that loads the schema graph and generates the migration files. Create a new Go file
named `main.go` and copy the following contents:

```go title="main.go"
package main

import (
	"context"
	"log"
	"os"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// We need a name for the new migration file.
	if len(os.Args) < 2 {
		log.Fatalln("no name given")
	}
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalln(err)
	}
	// Load the graph.
	graph, err := entc.LoadGraph("./ent/schema", &gen.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	tbls, err := graph.Tables()
	if err != nil {
		log.Fatalln(err)
	}
	// Open connection to the database.
	drv, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/ent")
	if err != nil {
		log.Fatalln(err)
	}
	// Inspect the current database state and compare it with the graph.
	m, err := schema.NewMigrate(drv, schema.WithDir(dir))
	if err != nil {
		log.Fatalln(err)
	}
	if err := m.NamedDiff(context.Background(), os.Args[1], tbls...); err != nil {
		log.Fatalln(err)
	}
}
```

Everything we have to do now is creating the migration directory and executing the above Go file:

```shell
mkdir migrations
go run -mod=mod main.go initial
```

You will now see two new files in the `migrations` directory: `<timestamp>_initial.down.sql`
and `<timestamp>_initial.up.sql`.


### Wrapping Up

 -- TODO

Have questions? Need help with getting started? Feel free to [join our Slack channel](https://entgo.io/docs/slack/).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://www.getrevue.co/profile/ent)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
