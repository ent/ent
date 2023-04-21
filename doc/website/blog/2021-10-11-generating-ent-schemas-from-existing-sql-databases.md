---
title: Generating Ent Schemas from Existing SQL Databases 
author: Zeev Manilovich
authorURL: "https://github.com/zeevmoney"
authorImageURL: "https://avatars.githubusercontent.com/u/7361100?v=4"
---

A few months ago the Ent project announced
the [Schema Import Initiative](https://entgo.io/blog/2021/05/04/announcing-schema-imports), its goal is to help support
many use cases for generating Ent schemas from external resources. Today, I'm happy to share a project I’ve been working
on: **entimport** - an _importent_ (pun intended) command line tool designed to create Ent schemas from existing SQL
databases. This is a feature that has been requested by the community for some time, so I hope many people find it
useful. It can help ease the transition of an existing setup from another language or ORM to Ent. It can also help with
use cases where you would like to access the same data from different platforms (such as to automatically sync between
them).  
The first version supports both MySQL and PostgreSQL databases, with some limitations described below. Support for other
relational databases such as SQLite is in the works.

## Getting Started

To give you an idea of how `entimport` works, I want to share a quick example of end to end usage with a MySQL database.
On a high-level, this is what we’re going to do:

1. Create a Database and Schema - we want to show how `entimport` can generate an Ent schema for an existing database.
   We will first create a database, then define some tables in it that we can import into Ent.
2. Initialize an Ent Project - we will use the Ent CLI to create the needed directory structure and an Ent schema
   generation script.
3. Install `entimport`
4. Run `entimport` against our demo database - next, we will import the database schema that we’ve created into our Ent
   project.
5. Explain how to use Ent with our generated schemas.  

Let's get started.

### Create a Database

We’re going to start by creating a database. The way I prefer to do it is to use
a [Docker](https://docs.docker.com/get-docker/) container. We will use a `docker-compose` which will automatically pass
all needed parameters to the MySQL container.

Start the project in a new directory called `entimport-example`. Create a file named `docker-compose.yaml` and paste the
following content inside:

```yaml
version: "3.7"

services:

  mysql8:
    platform: linux/amd64
    image: mysql
    environment:
      MYSQL_DATABASE: entimport
      MYSQL_ROOT_PASSWORD: pass
    healthcheck:
      test: mysqladmin ping -ppass
    ports:
      - "3306:3306"
```

This file contains the service configuration for a MySQL docker container. Run it with the following command:

```shell
docker-compose up -d
```

Next, we will create a simple schema. For this example we will use a relation between two entities:

- User
- Car

Connect to the database using MySQL shell, you can do it with the following command:
> Make sure you run it from the root project directory

```shell
docker-compose exec mysql8 mysql --database=entimport -ppass
```

```sql
create table users
(
    id        bigint auto_increment primary key,
    age       bigint       not null,
    name      varchar(255) not null,
    last_name varchar(255) null comment 'surname'
);

create table cars
(
    id          bigint auto_increment primary key,
    model       varchar(255) not null,
    color       varchar(255) not null,
    engine_size mediumint    not null,
    user_id     bigint       null,
    constraint cars_owners foreign key (user_id) references users (id) on delete set null
);
```

Let's validate that we've created the tables mentioned above, in your MySQL shell, run:

```sql
show tables;
+---------------------+
| Tables_in_entimport |
+---------------------+
| cars                |
| users               |
+---------------------+
```

We should see two tables: `users` & `cars`

### Initialize Ent Project

Now that we've created our database, and a baseline schema to demonstrate our example, we need to create
a [Go](https://golang.org/doc/install) project with Ent. In this phase I will explain how to do it. Since eventually we
would like to use our imported schema, we need to create the Ent directory structure.

Initialize a new Go project inside a directory called `entimport-example`

```shell
go mod init entimport-example
```

Run Ent Init:

```shell
go run -mod=mod entgo.io/ent/cmd/ent new 
```

The project should look like this:

```
├── docker-compose.yaml
├── ent
│   ├── generate.go
│   └── schema
└── go.mod
```

### Install entimport

OK, now the fun begins! We are finally ready to install `entimport` and see it in action.  
Let’s start by running `entimport`:

```shell
go run -mod=mod ariga.io/entimport/cmd/entimport -h
```

`entimport` will be downloaded and the command will print:

```
Usage of entimport:
  -dialect string
        database dialect (default "mysql")
  -dsn string
        data source name (connection information)
  -schema-path string
        output path for ent schema (default "./ent/schema")
  -tables value
        comma-separated list of tables to inspect (all if empty)
```

### Run entimport

We are now ready to import our MySQL schema to Ent!

We will do it with the following command:
> This command will import all tables in our schema, you can also limit to specific tables using `-tables` flag.

```shell
go run ariga.io/entimport/cmd/entimport -dialect mysql -dsn "root:pass@tcp(localhost:3306)/entimport"
```

Like many unix tools, `entimport` doesn't print anything on a successful run. To verify that it ran properly, we will
check the file system, and more specifically `ent/schema` directory.

```console {5-6}
├── docker-compose.yaml
├── ent
│   ├── generate.go
│   └── schema
│       ├── car.go
│       └── user.go
├── go.mod
└── go.sum
```

Let’s see what this gives us - remember that we had two schemas: the `users` schema and the `cars` schema with a one to
many relationship. Let’s see how `entimport` performed.

```go title="entimport-example/ent/schema/user.go"
type User struct {
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.Int("age"), field.String("name"), field.String("last_name").Optional().Comment("surname")}
}
func (User) Edges() []ent.Edge {
	return []ent.Edge{edge.To("cars", Car.Type)}
}
func (User) Annotations() []schema.Annotation {
	return nil
}
```

```go title="entimport-example/ent/schema/car.go"
type Car struct {
	ent.Schema
}

func (Car) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), field.String("model"), field.String("color"), field.Int32("engine_size"), field.Int("user_id").Optional()}
}
func (Car) Edges() []ent.Edge {
	return []ent.Edge{edge.From("user", User.Type).Ref("cars").Unique().Field("user_id")}
}
func (Car) Annotations() []schema.Annotation {
	return nil
}
```

> **`entimport` successfully created entities and their relation!**

So far looks good, now let’s actually try them out. First we must generate the Ent schema. We do it because Ent is a
**schema first** ORM that [generates](https://entgo.io/docs/code-gen) Go code for interacting with different databases.

To run the Ent code generation:

```shell
go generate ./ent
```

Let's see our `ent` directory:

```
...
├── ent
│   ├── car
│   │   ├── car.go
│   │   └── where.go
...
│   ├── schema
│   │   ├── car.go
│   │   └── user.go
...
│   ├── user
│   │   ├── user.go
│   │   └── where.go
...
```

### Ent Example

Let’s run a quick example to verify that our schema works:

Create a file named `example.go` in the root of the project, with the following content:

> This part of the example can be found [here](https://github.com/zeevmoney/entimport-example/blob/master/part1/example.go)

```go title="entimport-example/example.go"
package main

import (
	"context"
	"fmt"
	"log"

	"entimport-example/ent"

	"entgo.io/ent/dialect"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open(dialect.MySQL, "root:pass@tcp(localhost:3306)/entimport?parseTime=True")
	if err != nil {
		log.Fatalf("failed opening connection to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	example(ctx, client)
}
```

Let's try to add a user, write the following code at the end of the file:

```go title="entimport-example/example.go"
func example(ctx context.Context, client *ent.Client) {
	// Create a User.
	zeev := client.User.
		Create().
		SetAge(33).
		SetName("Zeev").
		SetLastName("Manilovich").
		SaveX(ctx)
	fmt.Println("User created:", zeev)
}
```

Then run:

```shell
go run example.go
```

This should output:

`# User created: User(id=1, age=33, name=Zeev, last_name=Manilovich)`

Let's check with the database if the user was really added

```sql
SELECT *
FROM users
WHERE name = 'Zeev';

+--+---+----+----------+
|id|age|name|last_name |
+--+---+----+----------+
|1 |33 |Zeev|Manilovich|
+--+---+----+----------+
```

Great! now let's play a little more with Ent and add some relations, add the following code at the end of
the `example()` func:
> make sure you add `"entimport-example/ent/user"` to the import() declaration

```go title="entimport-example/example.go"
// Create Car.
vw := client.Car.
    Create().
    SetModel("volkswagen").
    SetColor("blue").
    SetEngineSize(1400).
    SaveX(ctx)
fmt.Println("First car created:", vw)

// Update the user - add the car relation.
client.User.Update().Where(user.ID(zeev.ID)).AddCars(vw).SaveX(ctx)

// Query all cars that belong to the user.
cars := zeev.QueryCars().AllX(ctx)
fmt.Println("User cars:", cars)

// Create a second Car.
delorean := client.Car.
    Create().
    SetModel("delorean").
    SetColor("silver").
    SetEngineSize(9999).
    SaveX(ctx)
fmt.Println("Second car created:", delorean)

// Update the user - add another car relation.
client.User.Update().Where(user.ID(zeev.ID)).AddCars(delorean).SaveX(ctx)

// Traverse the sub-graph.
cars = delorean.
    QueryUser().
    QueryCars().
    AllX(ctx)
fmt.Println("User cars:", cars)
```

> This part of the example can be found [here](https://github.com/zeevmoney/entimport-example/blob/master/part2/example.go)

Now do: `go run example.go`.  
After Running the code above, the database should hold a user with 2 cars in a O2M relation.

```sql
SELECT *
FROM users;

+--+---+----+----------+
|id|age|name|last_name |
+--+---+----+----------+
|1 |33 |Zeev|Manilovich|
+--+---+----+----------+

SELECT *
FROM cars;

+--+----------+------+-----------+-------+
|id|model     |color |engine_size|user_id|
+--+----------+------+-----------+-------+
|1 |volkswagen|blue  |1400       |1      |
|2 |delorean  |silver|9999       |1      |
+--+----------+------+-----------+-------+
```

### Syncing DB changes

Since we want to keep the database in sync, we want `entimport` to be able to change the schema after the database was
changed. Let's see how it works.

Run the following SQL code to add a `phone` column with a `unique` index to the `users` table:

```sql
alter table users
    add phone varchar(255) null;

create unique index users_phone_uindex
    on users (phone);
```

The table should look like this:

```sql
describe users;
+-----------+--------------+------+-----+---------+----------------+
| Field     | Type         | Null | Key | Default | Extra          |
+-----------+--------------+------+-----+---------+----------------+
| id        | bigint       | NO   | PRI | NULL    | auto_increment |
| age       | bigint       | NO   |     | NULL    |                |
| name      | varchar(255) | NO   |     | NULL    |                |
| last_name | varchar(255) | YES  |     | NULL    |                |
| phone     | varchar(255) | YES  | UNI | NULL    |                |
+-----------+--------------+------+-----+---------+----------------+
```

Now let's run `entimport` again to get the latest schema from our database:

```shell
go run -mod=mod ariga.io/entimport/cmd/entimport -dialect mysql -dsn "root:pass@tcp(localhost:3306)/entimport"
```

We can see that the `user.go` file was changed:

```go title="entimport-example/ent/schema/user.go"
func (User) Fields() []ent.Field {
	return []ent.Field{field.Int("id"), ..., field.String("phone").Optional().Unique()}
}
```

Now we can run `go generate ./ent` again and use the new schema to add a `phone` to the User entity.

## Future Plans

As mentioned above this initial version supports MySQL and PostgreSQL databases.  
It also supports all types of SQL relations. I have plans to further upgrade the tool and add features such as missing
PostgreSQL fields, default values, and more.

## Wrapping Up

In this post, I presented `entimport`, a tool that was anticipated and requested many times by the Ent community. I
showed an example of how to use it with Ent. This tool is another addition to Ent schema import tools, which are
designed to make the integration of ent even easier. For discussion and
support, [open an issue](https://github.com/ariga/entimport/issues/new). The full example can be
found [in here](https://github.com/zeevmoney/entimport-example). I hope you found this blog post useful!

:::note For more Ent news and updates:

- Subscribe to our [Newsletter](https://entgo.substack.com/)
- Follow us on [Twitter](https://twitter.com/entgo_io)
- Join us on #ent on the [Gophers Slack](https://entgo.io/docs/slack)
- Join us on the [Ent Discord Server](https://discord.gg/qZmPgTE6RX)

:::
