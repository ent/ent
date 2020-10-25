---
id: schema-edges
title: Edges
---

## Quick Summary

Edges are the relations (or associations) of entities. For example, user's pets, or group's users.


![er-group-users](https://entgo.io/assets/er_user_pets_groups.png)

In the example above, you can see 2 relations declared using edges. Let's go over them.

1\. `pets` / `owner` edges; user's pets and pet's owner - 

`ent/schema/user.go`
```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
```


`ent/schema/pet.go`
```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Pet holds the schema definition for the Pet entity.
type Pet struct {
	ent.Schema
}

// Fields of the Pet.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
```

As you can see, a `User` entity can **have many** pets, but a `Pet` entity can **have only one** owner.  
In relationship definition, the `pets` edge is a *O2M* (one-to-many) relationship, and the `owner` edge
is a *M2O* (many-to-one) relationship.

The `User` schema **owns** the `pets/owner` relationship because it uses `edge.To`, and the `Pet` schema
just has a back-reference to it, declared using `edge.From` with the `Ref` method.

The `Ref` method describes which edge of the `User` schema we're referencing because there can be multiple
references from one schema to other. 

The cardinality of the edge/relationship can be controlled using the `Unique` method, and it's explained
more widely below. 

2\. `users` / `groups` edges; group's users and user's groups - 

`ent/schema/group.go`
```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// Group schema.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Edges of the group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
```

`ent/schema/user.go`
```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users"),
		// "pets" declared in the example above.
		edge.To("pets", Pet.Type),
	}
}
```

As you can see, a Group entity can **have many** users, and a User entity can have **have many** groups.  
In relationship definition, the `users` edge is a *M2M* (many-to-many) relationship, and the `groups`
edge is also a *M2M* (many-to-many) relationship.

## To and From

`edge.To` and `edge.From` are the 2 builders for creating edges/relations.

A schema that defines an edge using the `edge.To` builder owns the relation,
unlike using the `edge.From` builder that gives only a back-reference for the relation (with a different name).

Let's go over a few examples that show how to define different relation types using edges.

## Relationship

- [O2O Two Types](#o2o-two-types)
- [O2O Same Type](#o2o-same-type)
- [O2O Bidirectional](#o2o-bidirectional)
- [O2M Two Types](#o2m-two-types)
- [O2M Same Type](#o2m-same-type)
- [M2M Two Types](#m2m-two-types)
- [M2M Same Type](#m2m-same-type)
- [M2M Bidirectional](#m2m-bidirectional)

## O2O Two Types

![er-user-card](https://entgo.io/assets/er_user_card.png)

In this example, a user **has only one** credit-card, and a card **has only one** owner.

The `User` schema defines an `edge.To` card named `card`, and the `Card` schema
defines a back-reference to this edge using `edge.From` named `owner`. 


`ent/schema/user.go`
```go
// Edges of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("card", Card.Type).
			Unique(),
	}
}
```

`ent/schema/card.go`
```go
// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Unique().
			// We add the "Required" method to the builder
			// to make this edge required on entity creation.
			// i.e. Card cannot be created without its owner.
			Required(),
	}
}
```

The API for interacting with these edges is as follows:
```go
func Do(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("Mashraki").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}
	log.Println("user:", a8m)
	card1, err := client.Card.
		Create().
		SetOwner(a8m).
		SetNumber("1020").
		SetExpired(time.Now().Add(time.Minute)).
		Save(ctx)
	if err != nil {
    	return fmt.Errorf("creating card: %v", err)
    }
	log.Println("card:", card1)
	// Only returns the card of the user,
	// and expects that there's only one.
	card2, err := a8m.QueryCard().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying card: %v", err)
    }
	log.Println("card:", card2)
	// The Card entity is able to query its owner using
	// its back-reference.
	owner, err := card2.QueryOwner().Only(ctx)
	if err != nil {
		return fmt.Errorf("querying owner: %v", err)
    }
	log.Println("owner:", owner)
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/o2o2types).

## O2O Same Type

![er-linked-list](https://entgo.io/assets/er_linked_list.png)

In this linked-list example, we have a **recursive relation** named `next`/`prev`. Each node in the list can
**have only one** `next` node. If a node A points (using `next`) to node B, B can get its pointer using `prev` (the back-reference edge).   

`ent/schema/node.go`
```go
// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("next", Node.Type).
			Unique().
			From("prev").
			Unique(),
	}
}
```

As you can see, in cases of relations of the same type, you can declare the edge and its
reference in the same builder.

```diff
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
+		edge.To("next", Node.Type).
+			Unique().
+			From("prev").
+			Unique(),

-		edge.To("next", Node.Type).
-			Unique(),
-		edge.From("prev", Node.Type).
-			Ref("next).
-			Unique(),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	head, err := client.Node.
		Create().
		SetValue(1).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the head: %v", err)
	}
	curr := head
	// Generate the following linked-list: 1<->2<->3<->4<->5.
	for i := 0; i < 4; i++ {
		curr, err = client.Node.
			Create().
			SetValue(curr.Value + 1).
			SetPrev(curr).
			Save(ctx)
		if err != nil {
			return err
		}
	}

	// Loop over the list and print it. `FirstX` panics if an error occur.
	for curr = head; curr != nil; curr = curr.QueryNext().FirstX(ctx) {
		fmt.Printf("%d ", curr.Value)
	}
	// Output: 1 2 3 4 5

	// Make the linked-list circular:
	// The tail of the list, has no "next".
	tail, err := client.Node.
		Query().
		Where(node.Not(node.HasNext())).
		Only(ctx)
	if err != nil {
		return fmt.Errorf("getting the tail of the list: %v", tail)
	}
	tail, err = tail.Update().SetNext(head).Save(ctx)
	if err != nil {
		return err
	}
	// Check that the change actually applied:
	prev, err := head.QueryPrev().Only(ctx)
	if err != nil {
		return fmt.Errorf("getting head's prev: %v", err)
	}
	fmt.Printf("\n%v", prev.Value == tail.Value)
	// Output: true
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/o2orecur).

## O2O Bidirectional

![er-user-spouse](https://entgo.io/assets/er_user_spouse.png)

In this user-spouse example, we have a **symmetric O2O relation** named `spouse`. Each user can **have only one** spouse.
If user A sets its spouse (using `spouse`) to B, B can get its spouse using the `spouse` edge.

Note that there are no owner/inverse terms in cases of bidirectional edges.

`ent/schema/user.go`
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("spouse", User.Type).
			Unique(),
	}
}
```

The API for interacting with this edge is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}
	nati, err := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		SetSpouse(a8m).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}

	// Query the spouse edge.
	// Unlike `Only`, `OnlyX` panics if an error occurs.
	spouse := nati.QuerySpouse().OnlyX(ctx)
	fmt.Println(spouse.Name)
	// Output: a8m

	spouse = a8m.QuerySpouse().OnlyX(ctx)
	fmt.Println(spouse.Name)
	// Output: nati

	// Query how many users have a spouse.
	// Unlike `Count`, `CountX` panics if an error occurs.
	count := client.User.
		Query().
		Where(user.HasSpouse()).
		CountX(ctx)
	fmt.Println(count)
	// Output: 2

	// Get the user, that has a spouse with name="a8m".
	spouse = client.User.
		Query().
		Where(user.HasSpouseWith(user.Name("a8m"))).
		OnlyX(ctx)
	fmt.Println(spouse.Name)
	// Output: nati
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/o2obidi).

## O2M Two Types

![er-user-pets](https://entgo.io/assets/er_user_pets.png)

In this user-pets example, we have a O2M relation between user and its pets.
Each user **has many** pets, and a pet **has one** owner.
If user A adds a pet B using the `pets` edge, B can get its owner using the `owner` edge (the back-reference edge).

Note that this relation is also a M2O (many-to-one) from the point of view of the `Pet` schema. 

`ent/schema/user.go`
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type),
	}
}
```

`ent/schema/pet.go`
```go
// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Create the 2 pets.
	pedro, err := client.Pet.
		Create().
		SetName("pedro").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %v", err)
	}
	lola, err := client.Pet.
		Create().
		SetName("lola").
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating pet: %v", err)
	}
	// Create the user, and add its pets on the creation.
	a8m, err := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddPets(pedro, lola).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating user: %v", err)
	}
	fmt.Println("User created:", a8m)
	// Output: User(id=1, age=30, name=a8m)

	// Query the owner. Unlike `Only`, `OnlyX` panics if an error occurs.
	owner := pedro.QueryOwner().OnlyX(ctx)
	fmt.Println(owner.Name)
	// Output: a8m

	// Traverse the sub-graph. Unlike `Count`, `CountX` panics if an error occurs.
	count := pedro.
		QueryOwner(). // a8m
		QueryPets().  // pedro, lola
		CountX(ctx)   // count
	fmt.Println(count)
	// Output: 2
	return nil
}
```
The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/o2m2types).

## O2M Same Type

![er-tree](https://entgo.io/assets/er_tree.png)

In this example, we have a recursive O2M relation between tree's nodes and their children (or their parent).  
Each node in the tree **has many** children, and **has one** parent. If node A adds B to its children,
B can get its owner using the `owner` edge.


`ent/schema/node.go`
```go
// Edges of the Node.
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("children", Node.Type).
			From("parent").
			Unique(),
	}
}
```

As you can see, in cases of relations of the same type, you can declare the edge and its
reference in the same builder.

```diff
func (Node) Edges() []ent.Edge {
	return []ent.Edge{
+		edge.To("children", Node.Type).
+			From("parent").
+			Unique(),

-		edge.To("children", Node.Type),
-		edge.From("parent", Node.Type).
-			Ref("children").
-			Unique(),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	root, err := client.Node.
		Create().
		SetValue(2).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("creating the root: %v", err)
	}
	// Add additional nodes to the tree:
	//
	//       2
	//     /   \
	//    1     4
	//        /   \
	//       3     5
	//
	// Unlike `Save`, `SaveX` panics if an error occurs.
	n1 := client.Node.
		Create().
		SetValue(1).
		SetParent(root).
		SaveX(ctx)
	n4 := client.Node.
		Create().
		SetValue(4).
		SetParent(root).
		SaveX(ctx)
	n3 := client.Node.
		Create().
		SetValue(3).
		SetParent(n4).
		SaveX(ctx)
	n5 := client.Node.
		Create().
		SetValue(5).
		SetParent(n4).
		SaveX(ctx)

	fmt.Println("Tree leafs", []int{n1.Value, n3.Value, n5.Value})
	// Output: Tree leafs [1 3 5]

	// Get all leafs (nodes without children).
	// Unlike `Int`, `IntX` panics if an error occurs.
	ints := client.Node.
		Query().                             // All nodes.
		Where(node.Not(node.HasChildren())). // Only leafs.
		Order(ent.Asc(node.FieldValue)).     // Order by their `value` field.
		GroupBy(node.FieldValue).            // Extract only the `value` field.
		IntsX(ctx)
	fmt.Println(ints)
	// Output: [1 3 5]

	// Get orphan nodes (nodes without parent).
	// Unlike `Only`, `OnlyX` panics if an error occurs.
	orphan := client.Node.
		Query().
		Where(node.Not(node.HasParent())).
		OnlyX(ctx)
	fmt.Println(orphan)
	// Output: Node(id=1, value=2)

	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/o2mrecur).

## M2M Two Types

![er-user-groups](https://entgo.io/assets/er_user_groups.png)

In this groups-users example, we have a M2M relation between groups and their users.
Each group **has many** users, and each user can be joined to **many** groups.

`ent/schema/group.go`
```go
// Edges of the Group.
func (Group) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("users", User.Type),
	}
}
```

`ent/schema/user.go`
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users"),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	hub := client.Group.
		Create().
		SetName("GitHub").
		SaveX(ctx)
	lab := client.Group.
		Create().
		SetName("GitLab").
		SaveX(ctx)
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		AddGroups(hub, lab).
		SaveX(ctx)
	nati := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		AddGroups(hub).
		SaveX(ctx)

	// Query the edges.
	groups, err := a8m.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying a8m groups: %v", err)
	}
	fmt.Println(groups)
	// Output: [Group(id=1, name=GitHub) Group(id=2, name=GitLab)]

	groups, err = nati.
		QueryGroups().
		All(ctx)
	if err != nil {
		return fmt.Errorf("querying nati groups: %v", err)
	}
	fmt.Println(groups)
	// Output: [Group(id=1, name=GitHub)]

	// Traverse the graph.
	users, err := a8m.
		QueryGroups().                                           // [hub, lab]
		Where(group.Not(group.HasUsersWith(user.Name("nati")))). // [lab]
		QueryUsers().                                            // [a8m]
		QueryGroups().                                           // [hub, lab]
		QueryUsers().                                            // [a8m, nati]
		All(ctx)
	if err != nil {
		return fmt.Errorf("traversing the graph: %v", err)
	}
	fmt.Println(users)
	// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/m2m2types).

## M2M Same Type

![er-following-followers](https://entgo.io/assets/er_following_followers.png)

In this following-followers example, we have a M2M relation between users to their followers. Each user 
can follow **many** users, and can have **many** followers.

`ent/schema/user.go`
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("following", User.Type).
			From("followers"),
	}
}
```


As you can see, in cases of relations of the same type, you can declare the edge and its
reference in the same builder.

```diff
func (User) Edges() []ent.Edge {
	return []ent.Edge{
+		edge.To("following", User.Type).
+			From("followers"),

-		edge.To("following", User.Type),
-		edge.From("followers", User.Type).
-			Ref("following"),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SaveX(ctx)
	nati := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		AddFollowers(a8m).
		SaveX(ctx)

	// Query following/followers:

	flw := a8m.QueryFollowing().AllX(ctx)
	fmt.Println(flw)
	// Output: [User(id=2, age=28, name=nati)]

	flr := a8m.QueryFollowers().AllX(ctx)
	fmt.Println(flr)
	// Output: []

	flw = nati.QueryFollowing().AllX(ctx)
	fmt.Println(flw)
	// Output: []

	flr = nati.QueryFollowers().AllX(ctx)
	fmt.Println(flr)
	// Output: [User(id=1, age=30, name=a8m)]

	// Traverse the graph:

	ages := nati.
		QueryFollowers().       // [a8m]
		QueryFollowing().       // [nati]
		GroupBy(user.FieldAge). // [28]
		IntsX(ctx)
	fmt.Println(ages)
	// Output: [28]

	names := client.User.
		Query().
		Where(user.Not(user.HasFollowers())).
		GroupBy(user.FieldName).
		StringsX(ctx)
	fmt.Println(names)
	// Output: [a8m]
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/m2mrecur).


## M2M Bidirectional

![er-user-friends](https://entgo.io/assets/er_user_friends.png)

In this user-friends example, we have a **symmetric M2M relation** named `friends`.
Each user can **have many** friends. If user A becomes a friend of B, B is also a friend of A.

Note that there are no owner/inverse terms in cases of bidirectional edges.

`ent/schema/user.go`
```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("friends", User.Type),
	}
}
```

The API for interacting with these edges is as follows:

```go
func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	a8m := client.User.
		Create().
		SetAge(30).
		SetName("a8m").
		SaveX(ctx)
	nati := client.User.
		Create().
		SetAge(28).
		SetName("nati").
		AddFriends(a8m).
		SaveX(ctx)

	// Query friends. Unlike `All`, `AllX` panics if an error occurs.
	friends := nati.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m)]

	friends = a8m.
		QueryFriends().
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=2, age=28, name=nati)]

	// Query the graph:
	friends = client.User.
		Query().
		Where(user.HasFriends()).
		AllX(ctx)
	fmt.Println(friends)
	// Output: [User(id=1, age=30, name=a8m) User(id=2, age=28, name=nati)]
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/m2mbidi).

## Required

Edges can be defined as required in the entity creation using the `Required` method on the builder.

```go
// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Unique().
			Required(),
	}
}
```

If the example above, a card entity cannot be created without its owner. 

## StorageKey

Custom storage configuration can be provided for edges using the `StorageKey` method.

```go
// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pets", Pet.Type).
			// Set the column name in the "pets" table for O2M relationship.
			StorageKey(edge.Column("owner_id")),
		edge.To("friends", User.Type).
			// Set the join-table and the column names for M2M relationship.
			StorageKey(edge.Table("friends"), edge.Columns("user_id", "friend_id")),
	}
}
```

## Indexes

Indexes can be defined on multi fields and some types of edges as well.
However, you should note, that this is currently an SQL-only feature.

Read more about this in the [Indexes](schema-indexes.md) section.

## Annotations

`Annotations` is used to attach arbitrary metadata to the edge object in code generation.
Template extensions can retrieve this metadata and use it inside their templates.

Note that the metadata object must be serializable to a JSON raw value (e.g. struct, map or slice).

```go
// Pet schema.
type Pet struct {
	ent.Schema
}

// Edges of the Pet.
func (Pet) Edges() []ent.Edge {
	return []ent.Field{
		edge.To("owner", User.Type).
			Ref("pets").
			Unique().
			Annotations(entgql.Annotation{
				OrderField: "OWNER",
			}),
	}
}
```

Read more about annotations and their usage in templates in the [template doc](templates.md#annotations).
