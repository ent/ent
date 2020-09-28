---
id: getting-started
title: Quick Introduction
sidebar_label: Quick Introduction
---

**ent** is a simple, yet powerful entity framework for Go, that makes it easy to build
and maintain applications with large data-models and sticks with the following principles:

- Easily model database schema as a graph structure.
- Define schema as a programmatic Go code.
- Static typing based on code generation.
- Database queries and graph traversals are easy to write.
- Simple to extend and customize using Go templates.

<br/>

![gopher-schema-as-code](https://entgo.io/assets/gopher-schema-as-code.png)

## Installation

```console
go get github.com/facebook/ent/cmd/entc
```

After installing `entc` (the code generator for `ent`), you should have it in your `PATH`.

## Setup A Go Environment

If your project directory is outside [GOPATH](https://github.com/golang/go/wiki/GOPATH) or you are not familiar with
GOPATH, setup a [Go module](https://github.com/golang/go/wiki/Modules#quick-start) project as follows:

```console
go mod init <project>
```

## Create Your First Schema

Go to the root directory of your project, and run:

```console
entc init User
```
The command above will generate the schema for `User` under `<project>/ent/schema/` directory:

```go
// <project>/ent/schema/user.go

package schema

import "github.com/facebook/ent"

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

Add 2 fields to the `User` schema:

```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)


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

Run `go generate` from the root directory of the project as follows:

```go
go generate ./ent
```

This produces the following files:
```
ent
├── client.go
├── config.go
├── context.go
├── ent.go
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

To get started, create a new `ent.Client`. For this example, we will use SQLite3.

```go
package main

import (
	"context"
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}
```

Now, we're ready to create our user. Let's call this function `CreateUser` for the sake of example:
```go
func CreateUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", u)
	return u, nil
}
```

## Query Your Entities

`entc` generates a package for each entity schema that contains its predicates, default values, validators
and additional information about storage elements (column names, primary keys, etc).

```go
package main

import (
	"log"

	"<project>/ent"
	"<project>/ent/user"
)

func QueryUser(ctx context.Context, client *ent.Client) (*ent.User, error) {
	u, err := client.User.
		Query().
		Where(user.NameEQ("a8m")).
		// `Only` fails if no user found,
		// or more than 1 user returned.
		Only(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed querying user: %v", err)
	}
	log.Println("user returned: ", u)
	return u, nil
}

```


## Add Your First Edge (Relation)
In this part of the tutorial, we want to declare an edge (relation) to another entity in the schema.  
Let's create 2 additional entities named `Car` and `Group` with a few fields. We use `entc`
to generate the initial schemas:

```console
entc init Car Group
```

And then we add the rest of the fields manually:
```go
import (
	"regexp"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
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
			// Regexp validation for group name.
			Match(regexp.MustCompile("[a-zA-Z_]+$")),
	}
}
```

Let's define our first relation. An edge from `User` to `Car` defining that a user
can **have 1 or more** cars, but a car **has only one** owner (one-to-many relation).

![er-user-cars](https://entgo.io/assets/re_user_cars.png)

Let's add the `"cars"` edge to the `User` schema, and run `go generate ./ent`:

 ```go
 import (
 	"log"

 	"github.com/facebook/ent"
 	"github.com/facebook/ent/schema/edge"
 )

 // Edges of the User.
 func (User) Edges() []ent.Edge {
 	return []ent.Edge{
		edge.To("cars", Car.Type),
 	}
 }
 ```

We continue our example by creating 2 cars and adding them to a user.
```go
func CreateCars(ctx context.Context, client *ent.Client) (*ent.User, error) {
	// Create a new car with model "Tesla".
	tesla, err := client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}

	// Create a new car with model "Ford".
	ford, err := client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating car: %v", err)
	}
	log.Println("car was created: ", ford)

	// Create a new user, and add it the 2 cars.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddCars(tesla, ford).
		Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed creating user: %v", err)
	}
	log.Println("user was created: ", a8m)
	return a8m, nil
}
```
But what about querying the `cars` edge (relation)? Here's how we do it:
```go
import (
	"log"

	"<project>/ent"
	"<project>/ent/car"
)

func QueryCars(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println("returned cars:", cars)

	// What about filtering specific cars.
	ford, err := a8m.QueryCars().
		Where(car.ModelEQ("Ford")).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	log.Println(ford)
	return nil
}
```

## Add Your First Inverse Edge (BackRef)
Assume we have a `Car` object and we want to get its owner; the user that this car belongs to.
For this, we have another type of edge called "inverse edge" that is defined using the `edge.From`
function.

![er-cars-owner](https://entgo.io/assets/re_cars_owner.png)

The new edge created in the diagram above is translucent, to emphasize that we don't create another
edge in the database. It's just a back-reference to the real edge (relation).

Let's add an inverse edge named `owner` to the `Car` schema, reference it to the `cars` edge
in the `User` schema, and run `go generate ./ent`.

```go
import (
	"log"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Edges of the Car.
func (Car) Edges() []ent.Edge {
	return []ent.Edge{
		// Create an inverse-edge called "owner" of type `User`
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

func QueryCarUsers(ctx context.Context, a8m *ent.User) error {
	cars, err := a8m.QueryCars().All(ctx)
	if err != nil {
		return fmt.Errorf("failed querying user cars: %v", err)
	}
	// Query the inverse edge.
	for _, ca := range cars {
		owner, err := ca.QueryOwner().Only(ctx)
		if err != nil {
			return fmt.Errorf("failed querying car %q owner: %v", ca.Model, err)
		}
		log.Printf("car %q owner: %q\n", ca.Model, owner.Name)
	}
	return nil
}
```

## Create Your Second Edge

We'll continue our example by creating a M2M (many-to-many) relationship between users and groups.

![er-group-users](https://entgo.io/assets/re_group_users.png)

As you can see, each group entity can **have many** users, and a user can **be connected to many** groups;
a simple "many-to-many" relationship. In the above illustration, the `Group` schema is the owner
of the `users` edge (relation), and the `User` entity has a back-reference/inverse edge to this
relationship named `groups`. Let's define this relationship in our schemas:

- `<project>/ent/schema/group.go`:

	```go
	 import (
		"log"
	
		"github.com/facebook/ent"
		"github.com/facebook/ent/schema/edge"
	 )
	
	 // Edges of the Group.
	 func (Group) Edges() []ent.Edge {
		return []ent.Edge{
			edge.To("users", User.Type),
		}
	 }
	```

- `<project>/ent/schema/user.go`:   
	```go
	 import (
	 	"log"
	
	 	"github.com/facebook/ent"
	 	"github.com/facebook/ent/schema/edge"
	 )
	
	 // Edges of the User.
	 func (User) Edges() []ent.Edge {
	 	return []ent.Edge{
			edge.To("cars", Car.Type),
		 	// Create an inverse-edge called "groups" of type `Group`
		 	// and reference it to the "users" edge (in Group schema)
		 	// explicitly using the `Ref` method.
			edge.From("groups", Group.Type).
				Ref("users"),
	 	}
	 }
	```

We run `entc` on the schema directory to re-generate the assets.
```console
go generate ./ent
```

## Run Your First Graph Traversal

In order to run our first graph traversal, we need to generate some data (nodes and edges, or in other words, 
entities and relations). Let's create the following graph using the framework:

![re-graph](https://entgo.io/assets/re_graph_getting_started.png)


```go

func CreateGraph(ctx context.Context, client *ent.Client) error {
	// First, create the users.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Ariel").
		Save(ctx)
	if err != nil {
		return err
	}
	neta, err := client.User.
		Create().
		SetAge(28).
		SetName("Neta").
		Save(ctx)
	if err != nil {
		return err
	}
	// Then, create the cars, and attach them to the users in the creation.
	_, err = client.Car.
		Create().
		SetModel("Tesla").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Mazda").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(a8m).               // attach this graph to Ariel.
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Car.
		Create().
		SetModel("Ford").
		SetRegisteredAt(time.Now()). // ignore the time in the graph.
		SetOwner(neta).              // attach this graph to Neta.
		Save(ctx)
	if err != nil {
		return err
	}
	// Create the groups, and add their users in the creation.
	_, err = client.Group.
		Create().
		SetName("GitLab").
		AddUsers(neta, a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	_, err = client.Group.
		Create().
		SetName("GitHub").
		AddUsers(a8m).
		Save(ctx)
	if err != nil {
		return err
	}
	log.Println("The graph was created successfully")
	return nil
}
```

Now when we have a graph with data, we can run a few queries on it:

1. Get all user's cars within the group named "GitHub":

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/group"
	)

	func QueryGithub(ctx context.Context, client *ent.Client) error {
		cars, err := client.Group.
			Query().
			Where(group.Name("GitHub")). // (Group(Name=GitHub),)
			QueryUsers().                // (User(Name=Ariel, Age=30),)
			QueryCars().                 // (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
			All(ctx)
		if err != nil {
			return fmt.Errorf("failed getting cars: %v", err)
		}
		log.Println("cars returned:", cars)
		// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Mazda, RegisteredAt=<Time>),)
		return nil
	}
	```

2. Change the query above, so that the source of the traversal is the user *Ariel*:

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/car"
	)

	func QueryArielCars(ctx context.Context, client *ent.Client) error {
		// Get "Ariel" from previous steps.
		a8m := client.User.
			Query().
			Where(
				user.HasCars(),
				user.Name("Ariel"),
			).
			OnlyX(ctx)
		cars, err := a8m. 						// Get the groups, that a8m is connected to:
				QueryGroups(). 					// (Group(Name=GitHub), Group(Name=GitLab),)
				QueryUsers().  					// (User(Name=Ariel, Age=30), User(Name=Neta, Age=28),)
				QueryCars().   					//
				Where(         					//
					car.Not( 					//	Get Neta and Ariel cars, but filter out
						car.ModelEQ("Mazda"),	//	those who named "Mazda"
					), 							//
				). 								//
				All(ctx)
		if err != nil {
			return fmt.Errorf("failed getting cars: %v", err)
		}
		log.Println("cars returned:", cars)
		// Output: (Car(Model=Tesla, RegisteredAt=<Time>), Car(Model=Ford, RegisteredAt=<Time>),)
		return nil
	}
	```

3. Get all groups that have users (query with a look-aside predicate):

	```go
	import (
		"log"
		
		"<project>/ent"
		"<project>/ent/group"
	)

	func QueryGroupWithUsers(ctx context.Context, client *ent.Client) error {
    	groups, err := client.Group.
    		Query().
    		Where(group.HasUsers()).
    		All(ctx)
    	if err != nil {
    		return fmt.Errorf("failed getting groups: %v", err)
    	}
    	log.Println("groups returned:", groups)
    	// Output: (Group(Name=GitHub), Group(Name=GitLab),)
    	return nil
    }
    ```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/start).
