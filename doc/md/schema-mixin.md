---
id: schema-mixin
title: Mixin
---
 
A `Mixin` allows you to create reusable pieces of `ent.Schema` code.

The `ent.Mixin` interface is as follows:

```go
type Mixin interface {
	Fields() []ent.Field
}
```

## Example

A common use case for `Mixin` is to mix-in a list of common fields to your schema.

```go
// -------------------------------------------------
// Mixin definition

// TimeMixin implements the ent.Mixin for sharing
// time fields with package schemas.
type TimeMixin struct{}

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

// TimeMixin implements the ent.Mixin for sharing
// entity details fields with package schemas.
type DetailsMixin struct{}

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
