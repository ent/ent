---
id: crud
title: CRUD API
---

As mentioned in the [introduction](code-gen.md) section, running `ent` on the schemas,
will generate the following assets:

- `Client` and `Tx` objects used for interacting with the graph.
- CRUD builders for each schema type. See [CRUD](crud.md) for more info.
- Entity object (Go struct) for each of the schema type.
- Package containing constants and predicates used for interacting with the builders.
- A `migrate` package for SQL dialects. See [Migration](migrate.md) for more info.

## Create A New Client

**MySQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	client, err := ent.Open("mysql", "<user>:<pass>@tcp(<host>:<port>)/<database>?parseTime=True")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```

**PostgreSQL**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/lib/pq"
)

func main() {
	client, err := ent.Open("postgres","host=<host> port=<port> user=<user> dbname=<database> password=<pass>")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```

**SQLite**

```go
package main

import (
	"log"

	"<project>/ent"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
```


**Gremlin (AWS Neptune)**

```go
package main

import (
	"log"

	"<project>/ent"
)

func main() {
	client, err := ent.Open("gremlin", "http://localhost:8182")
	if err != nil {
		log.Fatal(err)
	}
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

**SaveX** a pet; Unlike **Save**, **SaveX** panics if an error occurs.

```go
pedro := client.Pet.	// PetClient.
	Create().			// Pet create builder.
	SetName("pedro").	// Set field value.
	SetOwner(a8m).		// Set owner (unique edge).
	SaveX(ctx)			// Create and return.
```

## Create Many

**Save** a bulk of pets.

```go
names := []string{"pedro", "xabi", "layla"}
bulk := make([]*ent.PetCreate, len(names))
for i, name := range names {
    bulk[i] = client.Pet.Create().SetName(name).SetOwner(a8m)
}
pets, err := client.Pet.CreateBulk(bulk...).Save(ctx)
```

## Update One

Update an entity that was returned from the database.

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
			user.AgeGT(30), 	//
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

## Upsert One

Ent supports [upsert](https://en.wikipedia.org/wiki/Merge_(SQL)) records using the [`sql/upsert`](features.md#upsert)
feature-flag.

```go
err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflict().
	// Use the new values that were set on create.
	UpdateNewValues().
	Exec(ctx)

id, err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflict().
	// Use the "age" that was set on create.
	UpdateAge().
	// Set a different "name" in case of conflict.
	SetName("Mashraki").
	ID(ctx)

// Customize the UPDATE clause.
err := client.User.
	Create().
	SetAge(30).
	SetName("Ariel").
	OnConflict().
	UpdateNewValues().
	// Override some of the fields with a custom update.
	Update(func(u *ent.UserUpsert) {
		u.SetAddress("localhost")
		u.AddCount(1)
		u.ClearPhone()
	}).
	Exec(ctx)
```

In PostgreSQL, the [conflict target](https://www.postgresql.org/docs/current/sql-insert.html#SQL-ON-CONFLICT) is required:

```go
// Setting the column names using the fluent API.
err := client.User.
	Create().
	SetName("Ariel").
	OnConflictColumns(user.FieldName).
	UpdateNewValues().
	Exec(ctx)

// Setting the column names using the SQL API.
err := client.User.
	Create().
	SetName("Ariel").
	OnConflict(
	    sql.ConflictColumns(user.FieldName),	
	).
	UpdateNewValues().
	Exec(ctx)

// Setting the constraint name using the SQL API.
err := client.User.
	Create().
	SetName("Ariel").
	OnConflict(
	    sql.ConflictConstraint(constraint),	
	).
	UpdateNewValues().
	Exec(ctx)
```

In order to customize the executed statement, use the SQL API:

```go
id, err := client.User.
	Create().
	OnConflict(
		sql.ConflictColumns(...),
		sql.ConflictWhere(...),
		sql.UpdateWhere(...),
	).
	Update(func(u *ent.UserUpsert) {
		u.SetAge(30)
		u.UpdateName()
	}).
	ID(ctx)

// INSERT INTO "users" (...) VALUES (...) ON CONFLICT WHERE ... DO UPDATE SET ... WHERE ...
```

:::info
Since the upsert API is implemented using the `ON CONFLICT` clause (and `ON DUPLICATE KEY` in MySQL),
Ent executes only one statement to the database, and therefore, only create [hooks](hooks.md) are applied
for such operations.
:::

## Upsert Many

```go
err := client.User.             // UserClient
	CreateBulk(builders...).    // User bulk create.
	OnConflict().               // User bulk upsert.
	UpdateNewValues().          // Use the values that were set on create in case of conflict.
	Exec(ctx)                   // Execute the statement.
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

Count the number of posts without comments.
```go
n, err := client.Post.
	Query().
	Where(
		post.Not(
		    post.HasComments(),	
		)
	).
	Count(ctx)
```

More advance traversals can be found in the [next section](traversals.md). 

## Field Selection

Get all pet names.

```go
names, err := client.Pet.
	Query().
	Select(pet.FieldName).
	Strings(ctx)
```

Get all unique pet names.

```go
names, err := client.Pet.
	Query().
	Unique(true).
	Select(pet.FieldName).
	Strings(ctx)
```

Count the number of unique pet names.

```go
n, err := client.Pet.
	Query().
	Unique(true).
	Select(pet.FieldName).
	Count(ctx)
```

Select partial objects and partial associations.gs
Get all pets and their owners, but select and fill only the `ID` and `Name` fields.

```go
pets, err := client.Pet.
    Query().
    Select(pet.FieldName).
    WithOwner(func (q *ent.UserQuery) {
        q.Select(user.FieldName)
    }).
    All(ctx)
```

Scan all pet names and ages to custom struct.

```go
var v []struct {
	Age  int    `json:"age"`
	Name string `json:"name"`
}
err := client.Pet.
	Query().
	Select(pet.FieldAge, pet.FieldName).
	Scan(ctx, &v)
if err != nil {
	log.Fatal(err)
}
```

Update an entity and return a partial of it.

```go
pedro, err := client.Pet.
	UpdateOneID(id).
	SetAge(9).
	SetName("pedro").
	// Select allows selecting one or more fields (columns) of the returned entity.
	// The default is selecting all fields defined in the entity schema.
	Select(pet.FieldName).
	Save(ctx)
```

## Delete One 

Delete an entity.

```go
err := client.User.
	DeleteOne(a8m).
	Exec(ctx)
```

Delete by ID.

```go
err := client.User.
	DeleteOneID(id).
	Exec(ctx)
```

## Delete Many

Delete using predicates.

```go
_, err := client.File.
	Delete().
	Where(file.UpdatedAtLT(date)).
	Exec(ctx)
```

## Mutation

Each generated node type has its own type of mutation. For example, all [`User` builders](crud.md#create-an-entity), share
the same generated `UserMutation` object.
However, all builder types implement the generic <a target="_blank" href="https://pkg.go.dev/entgo.io/ent?tab=doc#Mutation">`ent.Mutation`</a> interface.

For example, in order to write a generic code that apply a set of methods on both `ent.UserCreate`
and `ent.UserUpdate`, use the `UserMutation` object:

```go
func Do() {
    creator := client.User.Create()
    SetAgeName(creator.Mutation())
	updater := client.User.UpdateOneID(id)
	SetAgeName(updater.Mutation())
}

// SetAgeName sets the age and the name for any mutation.
func SetAgeName(m *ent.UserMutation) {
    m.SetAge(32)
    m.SetName("Ariel")
}
```

In some cases, you want to apply a set of methods on multiple types.
For cases like this, either use the generic `ent.Mutation` interface,
or create your own interface.

```go
func Do() {
	creator1 := client.User.Create()
	SetName(creator1.Mutation(), "a8m")

	creator2 := client.Pet.Create()
	SetName(creator2.Mutation(), "pedro")
}

// SetNamer wraps the 2 methods for getting
// and setting the "name" field in mutations.
type SetNamer interface {
	SetName(string)
	Name() (string, bool)
}

func SetName(m SetNamer, name string) {
    if _, exist := m.Name(); !exist {
    	m.SetName(name)
    }
}
```
