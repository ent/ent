---
title: Automatic migration planning
id: auto-plan
---
## Supporting repository

The change described in this section can be found in PR [#4](https://github.com/rotemtam/ent-versioned-migrations-demo/pull/4/files)
in the supporting repository.

## Automatic migration planning

One of the convenient features of Automatic Migrations is that developers do not
need to write the SQL statements to create or modify the database schema. To 
achieve similar benefits, we will now add a script to our project that will 
automatically plan migration files for us based on the changes to our schema. 

To do this, Ent uses [Atlas](https://atlasgo.io), an open-source tool for managing database
schemas that was created by the same people behind Ent. 

If you have been following our example repo, we have been using SQLite as our database
until this point. To demonstrate a more realistic use case, we will now switch to MySQL.
See this change in [PR #3](https://github.com/rotemtam/ent-versioned-migrations-demo/pull/3/files).

## Dev database

To be able to plan accurate and consistent migration files, Atlas introduces the
concept of a [Dev database](https://atlasgo.io/concepts/dev-database), a temporary
database which is used to simulate the state of the database after different changes.
Therefore, to use Atlas to automatically plan migrations, we need to supply a connection
string to such a database to our migration planning script. Such a database is most commonly 
spun up using a local Docker container. Let's do this now by running the following command:

```shell
docker run --rm --name atlas-db-dev -d -p 3306:3306 -e MYSQL_DATABASE=dev -e MYSQL_ROOT_PASSWORD=pass mysql:8
```

## Migration planning script

Now that we have a Dev database, we can write a script that will use Atlas to plan
migration files for us. Let's create a new file called `main.go` in the `ent/migrate` directory
of our project:

```go title=ent/migrate/main.go
//go:build ignore

package main

import (
    "context"
    "log"
    "os"
    
    // highlight-next-line
    "<project>/ent/migrate"

    atlas "ariga.io/atlas/sql/migrate"
    "entgo.io/ent/dialect"
    "entgo.io/ent/dialect/sql/schema"
    _ "github.com/go-sql-driver/mysql"
)

func main() {
    ctx := context.Background()
    // Create a local migration directory able to understand Atlas migration file format for replay.
    dir, err := atlas.NewLocalDir("ent/migrate/migrations")
    if err != nil {
        log.Fatalf("failed creating atlas migration directory: %v", err)
    }
    // Migrate diff options.
    opts := []schema.MigrateOption{
        schema.WithDir(dir),                         // provide migration directory
        schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
        schema.WithDialect(dialect.MySQL),           // Ent dialect to use
        schema.WithFormatter(atlas.DefaultFormatter),
    }
    if len(os.Args) != 2 {
        log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
    }
    // Generate migrations using Atlas support for MySQL (note the Ent dialect option passed above).
    //highlight-next-line
    err = migrate.NamedDiff(ctx, "mysql://root:pass@localhost:3306/test", os.Args[1], opts...)
    if err != nil {
        log.Fatalf("failed generating migration file: %v", err)
    }
}
```

:::info

Notice that you need to make some modifications to the script above in the highlighted lines.
Edit the import path of the `migrate` package to match your project and to supply the connection 
string to your Dev database.

:::

To run the script, first create a `migrations` directory in the `ent/migrate` directory of your
project:

```text
mkdir ent/migrate/migrations
```

Then, run the script to create the initial migration file for your project:

```shell
go run -mod=mod ent/migrate/main.go initial
```
Notice that `initial` here is just a label for the migration file. You can use any name you want.

Observe that after running the script, two new files were created in the `ent/migrate/migrations`
directory. The first file is named `atlas.sum`, which is a checksum file used by Atlas to enforce
a linear history of migrations:

```text title=ent/migrate/migrations/atlas.sum
h1:Dt6N5dIebSto365ZEyIqiBKDqp4INvd7xijLIokqWqA=
20221114165732_initialize.sql h1:/33+7ubMlxuTkW6Ry55HeGEZQ58JqrzaAl2x1TmUTdE=
```

The second file is the actual migration file, which is named after the label we passed to the
script:

```sql title=ent/migrate/migrations/20221114165732_initial.sql
-- create "users" table
CREATE TABLE `users` (`id` bigint NOT NULL AUTO_INCREMENT, `name` varchar(255) NOT NULL, `email` varchar(255) NOT NULL, PRIMARY KEY (`id`), UNIQUE INDEX `email` (`email`)) CHARSET utf8mb4 COLLATE utf8mb4_bin;
-- create "blogs" table
CREATE TABLE `blogs` (`id` bigint NOT NULL AUTO_INCREMENT, `title` varchar(255) NOT NULL, `body` longtext NOT NULL, `created_at` timestamp NOT NULL, `user_blog_posts` bigint NULL, PRIMARY KEY (`id`), CONSTRAINT `blogs_users_blog_posts` FOREIGN KEY (`user_blog_posts`) REFERENCES `users` (`id`) ON DELETE SET NULL) CHARSET utf8mb4 COLLATE utf8mb4_bin;
```

## Other migration tools

Atlas integrates very well with Ent, but it is not the only migration tool that can be used
to manage database schemas in Ent projects. The following is a list of other migration tools
that are supported:
* [Goose](https://github.com/pressly/goose)
* [Golang Migrate](https://github.com/golang-migrate/migrate)
* [Flyway](https://flywaydb.org)
* [Liquibase](https://www.liquibase.org)
* [dbmate](https://github.com/amacneil/dbmate)
To learn more about how to use these tools with Ent, see the
* [docs](https://entgo.io/docs/versioned-migrations#create-a-migration-files-generator) on this subject.

Next, let's see how to upgrade an existing production database to be managed with versioned migrations. 