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
