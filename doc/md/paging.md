---
id: paging
title: Paging And Ordering
---

## Limit

`Limit` limits the query result to `n` entities.

```go
users, err := client.User.
	Query().
	Limit(n).
	All(ctx)
```


## Offset

`Offset` sets the first node to return from the query. 

```go
users, err := client.User.
	Query().
	Offset(10).
	All(ctx)
```

## Ordering

`Order` returns the entities sorted by the values of one or more fields. Note that, an error
is returned if the given fields are not valid columns or foreign-keys.

```go
users, err := client.User.Query().
	Order(ent.Asc(user.FieldName)).
	All(ctx)
```

## Edge Ordering

In order to sort by fields of an edge (relation), start the traversal from the edge (you want to order by),
apply the ordering, and then jump to the neighbours (target type).

The following shows how to order the users by the `"name"` of their `"pets"` in ascending order.
```go
users, err := client.Pet.Query().
	Order(ent.Asc(pet.FieldName)).
	QueryOwner().
	All(ctx)
```