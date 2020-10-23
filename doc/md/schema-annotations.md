---
id: schema-annotations
title: Annotations
---

Schema annotations allow attaching metadata to schema objects like fields and edges and inject them to external templates.
An annotation must be a Go type that is serializable to JSON raw value (e.g. struct, map or slice)
and implement the [Annotation](https://pkg.go.dev/github.com/facebook/ent/schema/field?tab=doc#Annotation) interface.

The builtin annotations allow configuring the different storage drivers (like SQL) and control the code generation output. 

## Custom Table Name

A custom table name can be provided for types using the `entsql` annotation as follows:

```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/dialect/entsql"
	"github.com/facebook/ent/schema"
	"github.com/facebook/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "Users"},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
	}
}
```  
