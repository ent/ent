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
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
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

// Fields of the user.
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
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
)

// User schema.
type Pet struct {
	ent.Schema
}

// Fields of the user.
func (Pet) Fields() []ent.Field {
	return []ent.Field{
		// ...
	}
}

// Fields of the user.
func (Pet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("pets").
			Unique(),
	}
}
```

As you can see, a `User` entity can have many pets, but a `Pet` entity can have only one owner.  
In relationship definition, the `pets` edge is a *O2M* (one-to-many) relationship, and the `owner` edge
is a *M2O* (many-to-one) relationship.

The `User` schema **owns** the `pets/owner` relationship because it uses `edge.To`, and the `Pet` schema
just have a back-reference to it, declared using `edge.From` with the `Ref` method.

The `Ref` method describes which edge of the `User` schema we're referencing to, because, there can be multiple
references from one schema to other. 

The cardinality of the edge/relationship can be controlled using the `Unique` method, and it's explained
more widely below. 

2\. `users` / `groups` edges; group's users and user's groups - 

`ent/schema/group.go`
```go
package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
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

// Fields of the group.
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
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
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

// Fields of the user.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("groups", Group.Type).
			Ref("users"),
		// "pets" declared in the example above.
		edge.To("pets", Pet.Type),
	}
}
```

As you can see, a Group entity can have many users, and a User entity can have have many groups.  
In relationship definition, the `users` edge is a *M2M* (many-to-many) relationship, and the `groups`
edge is also a *M2M* (many-to-many) relationship.

## Relationship

## To

## From

## Required

## Indexes

## Examples