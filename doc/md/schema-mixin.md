---
id: schema-mixin
title: Mixin
---
 
A `Mixin` allows you to create reusable pieces of `ent.Schema` code.

The `ent.Mixin` interface is as follows:

```go
type Mixin interface {
	// Fields returns a slice of fields to add to the schema.
	Fields() []Field
	// Edges returns a slice of edges to add to the schema.
	Edges() []Edge
	// Indexes returns a slice of indexes to add to the schema.
	Indexes() []Index
	// Hooks returns a slice of hooks to add to the schema.
	// Note that mixin hooks are executed before schema hooks.
	Hooks() []Hook
	// Policy returns a privacy policy to add to the schema.
	// Note that mixin policy are executed before schema policy.
	Policy() Policy
	// Annotations returns a list of schema annotations to add
	// to the schema annotations.
	Annotations() []schema.Annotation
}
```

## Example

A common use case for `Mixin` is to mix-in a list of common fields to your schema.

```go
package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/field"
	"github.com/facebook/ent/schema/mixin"
)

// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct{
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (TimeMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable().
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// DetailsMixin implements the ent.Mixin for sharing
// entity details fields with package schemas.
type DetailsMixin struct{
	// We embed the `mixin.Schema` to avoid
	// implementing the rest of the methods.
	mixin.Schema
}

func (DetailsMixin) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.String("name").
			NotEmpty(),
	}
}

// -------------------------------------------------
// Schema definition

// User schema mixed-in the TimeMixin and DetailsMixin fields and therefore
// has 5 fields: `created_at`, `updated_at`, `age`, `name` and `nickname`.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		TimeMixin{},
		DetailsMixin{},
	}
}

func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").
			Unique(),
	}
}

// Pet schema mixed-in the DetailsMixin fields and therefore
// has 3 fields: `age`, `name` and `weight`.
type Pet struct {
	ent.Schema
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		DetailsMixin{},
	}
}

func (Pet) Fields() []ent.Field {
	return []ent.Field{
		field.Float("weight"),
	}
}
```

## Builtin Mixin

Package `mixin` provides a few builtin mixins that can be used
for adding the `create_time` and `update_time` fields to the schema.

In order to use them, add the `mixin.Time` mixin to your schema as follows:
```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/mixin"
)

type Pet struct {
	ent.Schema
}

func (Pet) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
		// Or, mixin.CreateTime only for create_time
		// and mixin.UpdateTime only for update_time.
	}
}
```
