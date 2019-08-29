---
id: crud
title: CRUD API
---

As mentioned in the [introduction](code-gen.md) section, running `entc` on the schemas,
will generate the following assets:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.md) for more info.
- Entity object (Go struct) for each of the schema type.
- Package contains constants and predicates used for interacting with the builders.
- A `migrate` package, for SQL dialects. See [Migration](migrate.md) for more info.

## Create A New Client

**MySQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/go-sql-driver/mysql"
	"github.com/facebookincubator/ent/dialect/sql"
)

func main() {
	drv, err := sql.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer drv.Close()
	client := ent.NewClient(ent.Driver(drv))
}
```

**SQLite**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
	"github.com/facebookincubator/ent/dialect/sql"
)

func main() {
	drv, err := sql.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer drv.Close()
	client := ent.NewClient(ent.Driver(drv))
}
```


**Gremlin (AWS Neptune)**

```go
package main

import (
	"log"
	"net/url"

	"<project>/ent"

	"github.com/facebookincubator/ent/dialect/gremlin"
)

func main() {
	c, err := gremlin.NewClient(gremlin.Config{
		Endpoint: gremlin.Endpoint{
			URL: &url.URL{
				Scheme: "http",
				Host:   "localhost:8182",
			},
		},
	})
	if err != nil {
		log.Fatal(err)
	}
	client := ent.NewClient(ent.Driver(gremlin.NewDriver(c)))
}
```

## Create An Entity

**Save** a user.

```go
a8m, err := client.User.	// UserClient.
	Create().				// User create builder.
	SetName("a8m").			// Set field value.
	SetNillableAge(age).	// Avoid nil checks.
	AddGroups(g1, g2).		// Add many edges.
	SetSpouse(nati).		// Set unique edge.
	Save(ctx)				// Create and return.
```

**SaveX** a user; Unlike **Save**, **SaveX** panics if an error occurs.

```go
pedro := client.Pet.	// PetClient.
	Create().			// User create builder.
	SetName("pedro").	// Set field value.
	SetOwner(a8m).		// Set owner (unique edge).
	SaveX(ctx)			// Create and return.
```

## Update One

If an entity that already returned from the database.

```go
a8m, err = a8m.Update().	// User update builder.
	RemoveGroup(g2).		// Remove specific edge.
	ClearCard().			// Clear unique edge.
	SetAge(30).				// Set field value
	Save(ctx)				// Save and return.
```


## Update By ID

```go
pedro, err := client.Pet.	// PetClient.
	UpdateOneID(id).		// Pet update builder.
	SetName("pedro").		// Set field name.
	SetOwnerID(owner).		// Set unique edge, using id.
	Save(ctx)				// Save and return.
```

## Update Many

Filter using predicates.

```go
n, err := client.User.			// UserClient.
	Update().					// Pet update builder.
	Where(						//
		user.Or(				// (age >= 30 OR name = "bar") 
			user.AgeEQ(30), 	//
			user.Name("bar"),	// AND
		),						//  
		user.HasFollowers(),	// UserHasFollowers()  
	).							//
	SetName("foo").				// Set field name.
	Save(ctx)					// exec and return.
```

Query edge-predicates.

```go
n, err := client.User.			// UserClient.
	Update().					// Pet update builder.
	Where(						// 
		user.HasFriendsWith(	// UserHasFriendsWith (
			user.Or(			//   age = 20
				user.Age(20),	//      OR
				user.Age(30),	//   age = 30
			)					// )
		), 						//
	).							//
	SetName("a8m").				// Set field name.
	Save(ctx)					// exec and return.
```

## Query The Graph

Get all users with followers.
```go
users, err := client.User.		// UserClient.
	Query().					// User query builder.
	Where(user.HasFollowers()).	// filter only users with followers.
	All(ctx)					// query and return.
```

Get all followers of a specific user; Start the traversal from a node in the graph.
```go
users, err := a8m.
	QueryFollowers().
	All(ctx)
```

Get all pets of the followers of a user.
```go
users, err := a8m.
	QueryFollowers().
	QueryPets().
	All(ctx)
```
More advance traversals can be found in the [next section](traversals.md). 

## Delete One 

Delete an entity.

```go
err := client.User.
	DeleteOne(a8m).
	Exec(ctx)
```

Delete by id.

```go
err := client.User.
	DeleteOneID(id).
	Exec(ctx)
```

## Delete Many

Delete using predicates.

```go
err := client.File.
	Delete().
	Where(file.UpdatedAtLT(date))
	Exec(ctx)
```