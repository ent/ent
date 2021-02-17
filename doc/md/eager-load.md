---
id: eager-load
title: Eager Loading
---

## Overview

`ent` supports querying entities with their associations (through their edges). The associated entities
are populated to the `Edges` field in the returned object.

Let's give an example hows does the API look like for the following schema:

![er-group-users](https://entgo.io/assets/er_user_pets_groups.png)



**Query all users with their pets:**
```go
users, err := client.User.
	Query().
	WithPets().
	All(ctx)
if err != nil {
	return err
}
// The returned users look as follows:
//
//	[
//		User {
//			ID:   1,
//			Name: "a8m",
//			Edges: {
//				Pets: [Pet(...), ...]
//				...
//			}
//		},
//		...
//	]
//
for _, u := range users {
	for _, p := range u.Edges.Pets {
		fmt.Printf("User(%v) -> Pet(%v)\n", u.ID, p.ID)
		// Output:
		// User(...) -> Pet(...)
	}
} 
```

Eager loading allows to query more than one association (including nested), and also
filter, sort or limit their result. For example:

```go
admins, err := client.User.
	Query().
	Where(user.Admin(true)).
	// Populate the `pets` that associated with the `admins`.
	WithPets().
	// Populate the first 5 `groups` that associated with the `admins`.
	WithGroups(func(q *ent.GroupQuery) {
		q.Limit(5) 				// Limit to 5.
		q.WithUsers().Limit(5)	// Populate the `users` of each `groups`. 
	}).
	All(ctx)
if err != nil {
	return err
}

// The returned users look as follows:
//
//	[
//		User {
//			ID:   1,
//			Name: "admin1",
//			Edges: {
//				Pets:   [Pet(...), ...]
//				Groups: [
//					Group {
//						ID:   7,
//						Name: "GitHub",
//						Edges: {
//							Users: [User(...), ...]
//							...
//						}
//					}
//				]
//			}
//		},
//		...
//	]
//
for _, admin := range admins {
	for _, p := range admin.Edges.Pets {
		fmt.Printf("Admin(%v) -> Pet(%v)\n", u.ID, p.ID)
		// Output:
		// Admin(...) -> Pet(...)
	}
	for _, g := range admin.Edges.Groups {
		for _, u := range g.Edges.Users {
			fmt.Printf("Admin(%v) -> Group(%v) -> User(%v)\n", u.ID, g.ID, u.ID)
			// Output:
			// Admin(...) -> Group(...) -> User(...)
		}
	}
} 
```

## API

Each query-builder has a list of methods in the form of `With<E>(...func(<N>Query))` for each of its edges.
`<E>` stands for the edge name (like, `WithGroups`) and `<N>` for the edge type (like, `GroupQuery`).
 
Note that, only SQL dialects support this feature.

## Implementation

Since a query-builder can load more than one association, it's not possible to load them using one `JOIN` operation.
Therefore, `ent` executes additional queries for loading associations. One query for `M2O/O2M` and `O2O` edges, and
2 queries for loading `M2M` edges.

Note that, we expect to improve this in the next versions of `ent`.
