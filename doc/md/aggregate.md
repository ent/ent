---
id: aggregate
title: Aggregation
---

## Group By

Group by `name` and `age` fields of all users, and sum their total age.

```go
package main

import (
	"context"
	
	"<project>/ent"
	"<project>/ent/user"
)

func Do(ctx context.Context, client *ent.Client) {
	var v []struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Sum   int    `json:"sum"`
		Count int    `json:"count"`
	}
	err := client.User.Query().
		GroupBy(user.FieldName, user.FieldAge).
		Aggregate(ent.Count(), ent.Sum(user.FieldAge)).
		Scan(ctx, &v)
}
```

Group by one field.

```go
package main

import (
	"context"
	
	"<project>/ent"
	"<project>/ent/user"
)

func Do(ctx context.Context, client *ent.Client) {
	names, err := client.User.
		Query().
		GroupBy(user.FieldName).
		Strings(ctx)
}
```

## Group By Edge

Custom aggregation functions can be useful if you want to write your own storage-specific logic.

The following shows how to group by the `id` and the `name` of all users and calculate the average `age` of their pets.

```go
package main

import (
	"context"
	"log"

	"<project>/ent"
	"<project>/ent/pet"
	"<project>/ent/user"
)

func Do(ctx context.Context, client *ent.Client) {
	var users []struct {
		ID      int
		Name    string
		Average float64
	}
	err := client.User.Query().
		GroupBy(user.FieldID, user.FieldName).
		Aggregate(func(s *sql.Selector) string {
			t := sql.Table(pet.Table)
			s.Join(t).On(s.C(user.FieldID), t.C(pet.OwnerColumn))
			return sql.As(sql.Avg(t.C(pet.FieldAge)), "average")
		}).
		Scan(ctx, &users)
}
```