---
title: Versioned Migrations Management and Migration Directory Integrity
author: Jannik Clausen (MasseElch)
authorURL: "https://github.com/masseelch"
authorImageURL: "https://avatars.githubusercontent.com/u/12862103?v=4"
image: "https://entgo.io/images/assets/migrate/atlas-validate.png"
---

Five weeks ago we released a long awaited feature for managing database changes in Ent: **Versioned Migrations**. In
the [announcement blog post](2022-03-14-announcing-versioned-migrations.md) we gave a brief introduction into both the
declarative and change-based approach to keep database schemas in sync with the consuming applications, as well as their
drawbacks and why [Atlas'](https://atlasgo.io) (Ents underlying migration engine) attempt of bringing the best of both
worlds into one workflow is worth a try. We call it **Versioned Migration Authoring** and if you haven't read it, now is
a good time!

With versioned migration authoring, the resulting migration files are still "change-based", but have been safely planned
by the Atlas engine. This means that you can still use your favorite migration management tool,
like [Flyway](https://flywaydb.org/), [Liquibase](https://liquibase.org/), 
[golang-migrate/migrate](https://github.com/golang-migrate/migrate), or 
[pressly/goose](https://github.com/pressly/goose) when developing services with Ent.

In this blog post I want to show you another new feature of the Atlas project we call the **Migration Directory 
Integrity File**, which is now supported in Ent, and how you can use it with any of the migration management tools you 
are already used to and like. 

### The Problem

When using versioned migrations, developers need to be careful of doing the following in order to not break the database:

1. Retroactively changing migrations that have already run.
2. Accidentally changing the order in which migrations are organized.
3. Checking in semantically incorrect SQL scripts.
Theoretically, code review should guard teams from merging migrations with these issues. In my experience, however, there are many kinds of errors that can slip the human eye, making this approach error-prone.
Therefore, an automated way of preventing these errors is much safer.

The first issue (changing history) is addressed by most management tools by saving a hash of the applied migration file to the managed
database and comparing it with the files. If they don't match, the migration can be aborted. However, this happens in a
very late stage in the development cycle (during deployment), and it could save both time and resources if this can be detected
earlier.

For the second (and third) issue, consider the following scenario:

![atlas-versioned-migrations-no-conflict](https://entgo.io/images/assets/migrate/no-conflict-2.svg)

This diagram shows two possible errors that go undetected. The first one being the order of the migration files. 

Team A and Team B both branch a feature roughly at the same time. Team B generates a migration file with a version
timestamp **x** and continues to work on the feature. Team A generates a migration file at a later point in time and
therefore has the migration version timestamp **x+1**. Team A finishes the feature and merges it into master,
possibly automatically deploying it in production with the migration version **x+1** applied. No problem so far.

Now, Team B merges its feature with the migration version **x**, which predates the already applied version **x+1**. If the code
review process does not detect this, the migration file lands in production, and it now depends on the specific migration
management tool to decide what happens.

Most tools have their own solution to that problem, `pressly/goose` for example takes an approach they
call [hybrid versioning](https://github.com/pressly/goose/issues/63#issuecomment-428681694). Before I introduce you to
Atlas' (Ent's) unique way of handling this problem, let's have a quick look at the third issue:

If both Team A and Team B develop a feature where they need new tables or columns, and they give them the same name, (e.g.
`users`) they could both generate a statement to create that table. While the team that merges first will have a
successful migration, the second team's migration will fail since the table or column already exists.

### The Solution

Atlas has a unique way of handling the above problems. The goal is to raise awareness about the issues as soon as
possible. In our opinion, the best place to do so is in version control and continuous integration (CI) parts of a
product. Atlas' solution to this is the introduction of a new file we call the **Migration Directory Integrity File**.
It is simply another file named `atlas.sum` that is stored together with the migration files and contains some
metadata about the migration directory. Its format is inspired by the `go.sum` file of a Go module, and it would look
similar to this: 

```text
h1:KRFsSi68ZOarsQAJZ1mfSiMSkIOZlMq4RzyF//Pwf8A=
20220318104614_team_A.sql h1:EGknG5Y6GQYrc4W8e/r3S61Aqx2p+NmQyVz/2m8ZNwA=
```

The `atlas.sum` file contains a sum of the whole directory as its first entry, and a checksum for each of the migration
files (implemented by a reverse, one branch merkle hash tree). Let's see how we can use this file to detect the cases
above in version control and CI. Our goal is to raise awareness that both teams added migrations and that they most
likely have to be checked before proceeding the merge.

:::note
To follow along, run the following commands to quickly have an example to work with. They will:

1. Create a Go module and download all needed dependencies
2. Create a very basic User schema
3. Enable the versioned migrations feature
4. Run the codegen
5. Start a MySQL docker container to use (remove with `docker stop atlas-sum`)

```shell
mkdir ent-sum-file
cd ent-sum-file
go mod init ent-sum-file
go install entgo.io/ent/cmd/ent@master
go run entgo.io/ent/cmd/ent new User
sed -i -E 's|^//go(.*)$|//go\1 --feature sql/versioned-migration|' ent/generate.go
go generate ./...
docker run --rm --name atlas-sum --detach --env MYSQL_ROOT_PASSWORD=pass --env MYSQL_DATABASE=ent -p 3306:3306 mysql
```
:::

The first step is to tell the migration engine to create and manage the `atlas.sum` by using the `schema.WithSumFile()`
option. The below example uses an [instantiated Ent client](/docs/versioned-migrations#from-client) to generate new
migration files:

```go
package main

import (
	"context"
	"log"
	"os"

	"ent-sum-file/ent"

	"ariga.io/atlas/sql/migrate"
	"entgo.io/ent/dialect/sql/schema"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "root:pass@tcp(localhost:3306)/ent")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Create a local migration directory.
	dir, err := migrate.NewLocalDir("migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Write migration diff.
	// highlight-start
	err = client.Schema.NamedDiff(ctx, os.Args[1], schema.WithDir(dir), schema.WithSumFile())
	// highlight-end
	if err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

After creating a migrations directory and running the above commands you should see `golang-migrate/migrate` compatible
migration files and in addition, the `atlas.sum` file with the following contents:

```shell
mkdir migrations
go run -mod=mod main.go initial
```

```sql title="20220504114411_initial.up.sql"
-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, PRIMARY KEY (`id`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;

```

```sql title="20220504114411_initial.down.sql"
-- reverse: create "users" table
DROP TABLE `users`;

```

```text title="atlas.sum"
h1:SxbWjP6gufiBpBjOVtFXgXy7q3pq1X11XYUxvT4ErxM=
20220504114411_initial.down.sql h1:OllnelRaqecTrPbd2YpDbBEymCpY/l6ihbyd/tVDgeY=
20220504114411_initial.up.sql h1:o/6yOczGSNYQLlvALEU9lK2/L6/ws65FrHJkEk/tjBk=
```

As you can see the `atlas.sum` file contains one entry for each migration file generated. With the `atlas.sum`
generation file enabled, both Team A and Team B will have such a file once they generate migrations for a schema change.
Now the version control will raise a merge conflict once the second Team attempts to merge their feature. 

![atlas-versioned-migrations-no-conflict](https://entgo.io/images/assets/migrate/conflict-2.svg)

:::note
In the following steps we invoke the Atlas CLI by calling `go run -mod=mod ariga.io/atlas/cmd/atlas`, but you can also
install the CLI globally (and then simply invoke it by calling `atlas`) to your system by following the installation 
instructions [here](https://atlasgo.io/cli/getting-started/setting-up#install-the-cli).
:::

You can check at any time, if your `atlas.sum` file is in sync with the migration directory with the following command (
which should not output any errors now):

```shell
go run -mod=mod ariga.io/atlas/cmd/atlas migrate validate
```

However, if you happen to make a manual change to your migration files, like adding a new SQL statement, editing an
existing one or even creating a completely new file, the `atlas.sum` file is no longer in sync with the migration
directory's contents. Attempting to generate new migration files for a schema change will now be blocked by the Atlas
migration engine. Try it out by creating a new empty migration file and run the `main.go` once again:

```shell
go run -mod=mod ariga.io/atlas/cmd/atlas migrate new migrations/manual_version.sql --format golang-migrate
go run -mod=mod main.go initial
# 2022/05/04 15:08:09 failed creating schema resources: validating migration directory: checksum mismatch
# exit status 1

```

The `atlas migrate validate` command will tell you the same:

```shell
go run -mod=mod ariga.io/atlas/cmd/atlas migrate validate
# Error: checksum mismatch
# 
# You have a checksum error in your migration directory.
# This happens if you manually create or edit a migration file.
# Please check your migration files and run
# 
# 'atlas migrate hash --force'
# 
# to re-hash the contents and resolve the error.
# 
# exit status 1
```

In order to get the `atlas.sum` file back in sync with the migration directory, we can once again use the Atlas CLI:

```shell
go run -mod=mod ariga.io/atlas/cmd/atlas migrate hash --force
```

As a safety measure, the Atlas CLI does not operate on a migration directory that is not in sync with its `atlas.sum`
file. Therefore, you need to add the `--force` flag to the command. 

For cases, where a developer forgets to update the `atlas.sum` file after making a manual change, you can add
an `atlas migrate validate` call to your CI. We are actively working on a GitHub action and CI solution, that does this 
(among and other things) for you _out-of-the-box_.

### Wrapping Up

In this post, we gave a brief introduction to common sources of schema migration when working with change based SQL
files and introduced a solution based on the Atlas project to make migrations more safe.

Have questions? Need help with getting started? Feel free to join
our [Ent Discord Server](https://discord.gg/qZmPgTE6RX).

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)
:::
