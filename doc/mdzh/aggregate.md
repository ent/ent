---
id: aggregate
title: Aggregation
---

## Group By

通过`name`和`age`字段给所有用户信息分组并求年龄的总和

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

根据一个字段分组

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
