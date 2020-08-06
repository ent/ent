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

`Offset` sets the first vertex to return from the query. 

```go
users, err := client.User.
	Query().
	Offset(10).
	All(ctx)
```

## Ordering

`Order` returns the entities sorted by the values of one or more fields.

```go
users, err := client.User.Query().
	Order(ent.Asc(user.FieldName)).
	All(ctx)
```
