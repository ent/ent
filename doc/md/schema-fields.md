---
id: schema-fields
title: Fields
---

## Quick Summary

Fields (or properties) in the schema are the attributes of the vertex. For example, a `User`
with 4 fields: `age`, `name`, `username` and `created_at`:

![re-fields-properties](https://entgo.io/assets/er_fields_properties.png)

Fields are returned from the schema using the `Fields` method. For example:

```go
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.String("username").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}
``` 

All fields are required by default, and can be set to optional using the `Optional` method.

## Types

The following types are currently supported by the framework:

- All Go numeric types. Like, `int`, `uint8`, `float64`, etc.
- `bool`
- `string`
- `time.Time`
- `[]byte` (only support by SQL dialects).

```go
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.Float("rank").
			Optional(),
		field.Bool("active").
			Default(false),
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}
```

To read more about how each type is mapped to its database-type, go to the [Migration](migrate.md) section.

## Default Values

**Non-unique** fields support default values using the `.Default` method.

## Validators

A field validator is a function from type `func(T) error` that is defined in the schema
using the `Validate` method, and applied on the given field value before creating or updating
the entity.

The supported types for field validators are `string` and all numeric types.

```go
package schema

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)


// Group schema.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("group name must begin with uppercase")
				}
				return nil
			}),
	}
}
```

## Builtin Validators

The framework provides a few builtin validators for each type:

- Numeric types:
  - `Positive()` 
  - `Negative()`
  - `Min(i)` - Validate that the given value is > i.
  - `Max(i)` - Validate that the given value is < i.
  - `Range(i, j)` - Validate that the given value is within the range [i, j]. 

- `string`
  - `MinLen(i)` 
  - `MaxLen(i)`
  - `Match(regexp.Regexp)`

## Optional

## Nillable

## Immutable

## Uniqueness

## Indexes

## Struct Tags