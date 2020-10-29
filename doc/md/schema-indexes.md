---
id: schema-indexes
title: Indexes
---

## Multiple Fields

Indexes can be configured on one or more fields in order to improve 
speed of data retrieval, or defining uniqueness.

```go
package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/index"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Indexes() []ent.Index {
	return []ent.Index{
        // non-unique index.
        index.Fields("field1", "field2"),
        // unique index.
        index.Fields("first_name", "last_name").
            Unique(),
	}
}
```

Note that for setting a single field as unique, use the `Unique`
method on the field builder as follows:

```go
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("phone").
			Unique(),
	}
}
```

## Index On Edges

Indexes can be configured on composition of fields and edges. The main use-case
is setting uniqueness on fields under a specific relation. Let's take an example:

![er-city-streets](https://entgo.io/assets/er_city_streets.png)

In the example above, we have a `City` with many `Street`s, and we want to set the
street name to be unique under each city.

`ent/schema/city.go`
```go
// City holds the schema definition for the City entity.
type City struct {
	ent.Schema
}

// Fields of the City.
func (City) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the City.
func (City) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("streets", Street.Type),
	}
}
```

`ent/schema/street.go`
```go
// Street holds the schema definition for the Street entity.
type Street struct {
	ent.Schema
}

// Fields of the Street.
func (Street) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the Street.
func (Street) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("city", City.Type).
			Ref("streets").
			Unique(),
	}
}

// Indexes of the Street.
func (Street) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("city").
			Unique(),
	}
}
```

`example.go`
```go
func Do(ctx context.Context, client *ent.Client) error {
	// Unlike `Save`, `SaveX` panics if an error occurs.
	tlv := client.City.
		Create().
		SetName("TLV").
		SaveX(ctx)
	nyc := client.City.
		Create().
		SetName("NYC").
		SaveX(ctx)
	// Add a street "ST" to "TLV".
	client.Street.
		Create().
		SetName("ST").
		SetCity(tlv).
		SaveX(ctx)
	// This operation will fail because "ST"
	// is already created under "TLV".
	_, err := client.Street.
		Create().
		SetName("ST").
		SetCity(tlv).
		Save(ctx)
	if err == nil {
		return fmt.Errorf("expecting creation to fail")
	}
	// Add a street "ST" to "NYC".
	client.Street.
		Create().
		SetName("ST").
		SetCity(nyc).
		SaveX(ctx)
	return nil
}
```

The full example exists in [GitHub](https://github.com/facebook/ent/tree/master/examples/edgeindex).

## Dialect Support

Indexes currently support only SQL dialects, and do not support Gremlin.

