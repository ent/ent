---
id: schema-config
title: Config
---

## Custom Table Name

A custom table name can be provided for types using the `Table` option as follows:

```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
)

type User struct {
	ent.Schema
}

func (User) Config() ent.Config {
	return ent.Config{
		Table: "Users",
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}
```  
