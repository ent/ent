---
id: doc1
title: Quick Introduction
sidebar_label: Quick Introduction
---

`ent` is a simple, yet powerful entity framework for Go built with the following principles:
- Defining your schema as code.
- Static typing first based on code generation.
- Make the work with graph-like data in Go easier.

## Installation

```console
$ go get github.com/facebookincubator/ent/entc/cmd/entc
```

After installing `entc` (the code generator for `ent`), you should have it in your `PATH`.

## Create Your First Schema
Go to the root directory of your project, and run:

```console
$ entc init User
```
The command above will generate the schema for `User` under `<project>/ent/schema/` directory:

```go
// <project>/ent/schema/user.go

package schema

import "github.com/facebookincubator/ent"

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return nil
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return nil
}

```

Let's add 2 fields to the `User` schema, and then run `entc generate`:

```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			Default("unknown"),
	}
}
```

Running `entc generate` from the root directory of the project:

```go
$ entc generate ./ent/schema
```

Will produce the following files:
```
ent
├── client.go
├── config.go
├── context.go
├── ent.go
├── example_test.go
├── migrate
│   ├── migrate.go
│   └── schema.go
├── predicate
│   └── predicate.go
├── schema
│   └── user.go
├── tx.go
├── user
│   ├── user.go
│   └── where.go
├── user.go
├── user_create.go
├── user_delete.go
├── user_query.go
└── user_update.go
```


## Create Your First Entity

First thing we need to do, is creating a new `ent.Client`. For the example purpose,
we will use SQLite3.

```go
package main

import (
	"log"

	"<project>/ent"

	"github.com/facebookincubator/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer db.Close()
	drv := dialect.Driver(db)
	if testing.Verbose() {
		drv = dialect.Debug(drv)
	}
	client := ent.NewClient(ent.Driver(db))
	// run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

Now, we're ready to create our user. Let's call this function `Do` for the sake of the example:
```go
func Do(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save()
	if err != nil {
		return nil, fmt.Error("failed creating user: %v", err)
	}
	log.Println("user was created: %v", u)
	return u, nil
}
```

## Query Your Entities

`entc` generates a package for each entity schema that contains its predicates, default values, validators
and information about storage elements (like, column names, primary keys, etc).

```go
package main

import (
	"log"

	"<project>/ent"
	"<project>/ent/user"
)

func Query(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Error("failed querying user: %v", err)
	}
	log.Println("user: %v", u)
	return u, nil
}
```


## Add Your First Edge (Relation)
In this part of the tutorial, we want to declare an edge to another entity in the schema.  
Let's create 2 additional entities named `Car` and `Group` with a few fields. We use `entc`
to generate the initial schema:

```console
$ entc init Car Group
```

And then, we add the rest of the fields manually:
```go
import (
	"log"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/field"
)

// Fields of the Car.
func (Car) Fields() []ent.Field {
	return []ent.Field{
		field.String("model"),
		field.Time("registered_at"),
	}
}


// Fields of the Group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			// regexp validation for group name.
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
	}
}
```

Let's define our first relation. An edge from `User` to `Car` defining that a user
can have 1 or more cars, but a car has only one owner (one-to-many relation).

![er-user-cars](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_user_cars.png)

Let's add the `"cars"` edge to the `User` schema, and run `entc generate ./ent/schema`:

 ```go
 import (
 	"log"

 	"github.com/facebookincubator/ent"
 	"github.com/facebookincubator/ent/edge"
 )

 // Edges of the User.
 func (User) Edges() []ent.Edge {
 	return []ent.Edge{
		edge.To("cars", Car.Type),
 	}
 }
 ```

We continue our example, by creating 2 cars, and add them to a user.
```go
func Do(ctx context.Context, client *ent.Client) error {
	// creating new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating car: %v", err)
	}

	// creating new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: %v", ford)

	// create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: %v", a8m)
}
```
But, what about querying the "cars" edge? Here's how we do it:
```go
import (
	"log"

	"<project>/ent"
	"<project>/ent/car"
)

func Do(ctx context.Context, client *ent.Client) error {
	// <continuation of the code block above>
	// ...

	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(cars...)

	// what about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.NameEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
}
```

## Add Your First Inverse Edge (BackRef)
Assume we have a `Car` object and we want to get its owner; The user that this car belongs to.
For this, we have another type of edge called "inverse edge" that is defined using the `edge.From`
function.

![er-cars-owner](https://s3.eu-central-1.amazonaws.com/entgo.io/assets/re_cars_owner.png)

The new edge created in the diagram above is transparent, to emphasis that we don't create another
edge in the database, and it is just a back-reference to the real edge.

Let's add an inverse edge named `"owner"` to the `Car` schema, reference it to the `"cars"` edge
in the `User` schema, and run `entc generate ./ent/schema`.

```go
import (
 "log"

 "github.com/facebookincubator/ent"
 "github.com/facebookincubator/ent/edge"
)

// Edges of the Car.
func (Car) Edges() []ent.Edge {
 return []ent.Edge{
	 // create an inverse-edge called "owner" of type `User`
	 // and reference it to the "cars" edge (in User schema)
	 // explicitly using the `Ref` method.
	 edge.From("owner", User.Type).
	 	Ref("cars").
		// setting the edge to unique, ensure
		// that a car can have only one owner.
		Unique(),
	}
}
```
We'll continue the user/cars example above by querying the inverse edge.

```go
import (
	"log"

	"<project>/ent"
)

func Do(ctx context.Context, client *ent.Client) error {
	// <continuation of the code block above>
	// ...

	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}

	// query the inverse edge.
	for _, car := range cars {
		owner, err := car.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %v", car.Model, err)
		}
		log.Printf("car %q owner: %q", car.Model, owner.Name)
	}
}
```

## Create Your Second Edge

TODO: User/Group example.
