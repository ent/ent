## Ent

> Note: if you edit this file, don't forget to update the [Wiki][wiki] as well.

### First Installation

If it is the first time you work with `entc`, you need to compile it manually, 
since we don't have any official binary distribution.

```
cd fbsource/fbcode/fbc/ent/entc/cmd/entc
go build
sudo mv entc /usr/local/bin
```

### Creating Schema
If you came here to see how to create your first schema, it's preferred to give `entc` 
to do it for you in order to keep the same standard for all projects that use it.
Run the following (replace `User/Group` with your entities):

```
entc init User Group
```

The schema that was created has 2 methods: `User.Fields` and `User.Edges`. The first defines the fields/properties
of the entity in the graph, and the second defines the edges for other entities (or to itself) in the graph.

Here are a few examples that will help you to understand what field/edge to declare in your schema
```
type User struct{
	ent.Schema
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		// "age" defines a field of type int, with a positive validator.
		// The validator is called on create or update on this field.
		field.Int("age").
			Positive(),

		// "name" defines a field of type string, and overrides the standard
		// tag for the generated entity.
		field.String("name").
			StructTag(`json:"first_name" graphql:"first_name"`),

		// "last" defines a field of type string, with default (on creation) to "unknown",
		// and 2 validators.
		field.String("last").
			Default("unknown").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("last name must begin with uppercase")
				}
				return nil
			}),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		// "groups" defines an edge to the Group entity (also a schema in this package).
		// The relation type for this edge is many-2-many. 
		edge.To("groups", Group.Type),
		
		// "workplace" defines an edge from the Company entity (also a schema in this package).
		// The relation type for this edge is many-2-one, and the owner of this edge, is the
		// Company entity. 
		edge.From("workplace", Company.Type).Unique().Ref("employees"),
		
		// "parent" defines an edge from a User to itself.
		edge.To("parent", User.Type).Unique().From("children"),
	}
}
```

### Code Generation

After running init, run the codegen on the directory the was created (`ent/schema`).

```
entc generate ./ent/schema
```

In addition to the "production" code, `entc` generates for you also an `example_test.go` file
with example for each ent in the graph.

### Working with the generated code

First, you need to create the `ent.Client` in order to interact with the different builders,
then, use this client to create, update, delete, or query entities.

```
package main

import (
	"log"
	
	"<project>/ent"
	"<project>/ent/user"
	"fbc/ent/dialect/sql"
)

func main() {
	ctx := context.Backgorund()
	drv, err := sql.Open("mysql", "root:pass@tcp(localhost:3306)/test?charset=utf8&parseTime=True")
	if err != nil {
    	log.Fatal(err)
    }
	defer drv.Close()
	client := ent.NewClient(drv)
	
	// Create:
	
	// `client.User` holds the `UserClient`, and `client.User.Create()` returns a new User creator.
	a8m, err := client.User.
		Create().
		SetAge(30)
		SetName("a8m").
		Save(ctx)	
	if err != nil {
		log.Fatal(err)
	}	
	// If you want to ignore the error checks in the code, replace `Save` with `SaveX`.
	
	// Delete:
	
	// delete one.
	client.User.DeleteOne(a8m).ExecX(ctx)
	// delete all.
	client.User.Delete().ExecX(ctx)
	// delete with condition.
	client.User.Delete().Where(user.Name("a8m")).ExecX(ctx)
	
	// Update:
	
	// add a user to a group.
    a8m = client.User.UpdateOne(a8m).AddGroups(grp).SaveX(ctx)
    // delete a user from a group.
    a8m = client.User.UpdateOne(a8m).RemoveGroups(grp).SaveX(ctx)
    // add user to all groups.
    client.Group.Update().AddUsers(a8m).ExecX(ctx)
	
	// Query:
	
	// get all groups.
	groups := client.Group.Query().AllX(ctx) 
	// get all groups of a specific user.
	groups = a8m.QueryGroups().AllX(ctx)
	// query by path.
	users := client.Group.
		Query().
		Where(group.HasUsers(), group.NameHasPrefix("fb")).
		QueryUsers().
		AllX(ctx)
		
	// Aggregation:
	
	var v []struct{
		Name  string `json:"name"`
		Count int    `json:"count"`
	}
	client.User.
		Query().
		GroupBy(user.FieldName).
		Aggregate(ent.Count()).
		ScanX(&v)
}
```


[wiki]: https://our.internmc.facebook.com/intern/wiki/Facebook_Connectivity_(FBC)/Entity_Framework/
