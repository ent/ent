---
id: schema-annotations
title: Annotations
---

Schema annotations allow attaching metadata to schema objects like fields and edges and inject them to external templates.
An annotation must be a Go type that is serializable to JSON raw value (e.g. struct, map or slice)
and implement the [Annotation](https://pkg.go.dev/entgo.io/ent/schema?tab=doc#Annotation) interface.

The builtin annotations allow configuring the different storage drivers (like SQL) and control the code generation output. 

## Custom Table Name

A custom table name can be provided for types using the `entsql` annotation as follows:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
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

## Foreign Keys Configuration

Ent allows to customize the foreign key creation and provide a [referential action](https://dev.mysql.com/doc/refman/8.0/en/create-table-foreign-keys.html#foreign-key-referential-actions)
for the `ON DELETE` clause:

```go
package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Default("Unknown"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("posts", Post.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Cascade,
			}),
	}
}
```

The example above configures the foreign key to cascade the deletion of rows in the parent table to the matching
rows in the child table.
