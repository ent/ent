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
	"entgo.io/ent"
	"entgo.io/ent/schema/index"
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

![er-city-streets](https://entgo.io/images/assets/er_city_streets.png)

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
	// This operation fails because "ST"
	// was already created under "TLV".
	if err := client.Street.
		Create().
		SetName("ST").
		SetCity(tlv).
		Exec(ctx); err == nil {
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

The full example exists in [GitHub](https://github.com/ent/ent/tree/master/examples/edgeindex).

## Index On Edge Fields

Currently `Edges` columns are always added after `Fields` columns. However, some indexes require these columns to come first in order to achieve specific optimizations. You can work around this problem by making use of [Edge Fields](schema-edges.mdx#edge-field). 

```go
// Card holds the schema definition for the Card entity.
type Card struct {
	ent.Schema
}
// Fields of the Card.
func (Card) Fields() []ent.Field {
	return []ent.Field{
		field.String("number").
			Optional(),
		field.Int("owner_id").
			Optional(),
	}
}
// Edges of the Card.
func (Card) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("card").
			Field("owner_id").
 			Unique(),
 	}
}
// Indexes of the Card.
func (Card) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("owner_id", "number"),
	}
}
```

## Dialect Support

Dialect specific features are allowed using [annotations](schema-annotations.md). For example, in order to use [index prefixes](https://dev.mysql.com/doc/refman/8.0/en/column-indexes.html#column-indexes-prefix)
in MySQL, use the following configuration:

```go
// Indexes of the User.
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("description").
			Annotations(entsql.Prefix(128)),
		index.Fields("c1", "c2", "c3").
			Annotations(
				entsql.PrefixColumn("c1", 100),
				entsql.PrefixColumn("c2", 200),
			)
	}
}
```

The code above generates the following SQL statements:

```sql
CREATE INDEX `users_description` ON `users`(`description`(128))

CREATE INDEX `users_c1_c2_c3` ON `users`(`c1`(100), `c2`(200), `c3`)
```

## Atlas Support

Starting with v0.10, Ent running migration with [Atlas](https://github.com/ariga/atlas). This option provides
more control on indexes such as, configuring their types or define indexes in a reverse order.

```go
func (User) Indexes() []ent.Index {
    return []ent.Index{
        index.Fields("c1").
            Annotations(entsql.Desc()),
        index.Fields("c1", "c2", "c3").
            Annotations(entsql.DescColumns("c1", "c2")),
        index.Fields("c4").
            Annotations(entsql.IndexType("HASH")),
        // Enable FULLTEXT search on MySQL,
        // and GIN on PostgreSQL.
        index.Fields("c5").
            Annotations(
                entsql.IndexTypes(map[string]string{
                    dialect.MySQL:    "FULLTEXT",
                    dialect.Postgres: "GIN",
                }),
            ),
		// For PostgreSQL, we can include in the index
		// non-key columns.
		index.Fields("workplace").
			Annotations(
				entsql.IncludeColumns("address"),
			),
		// Define a partial index on SQLite and PostgreSQL.
		index.Fields("nickname").
			Annotations(
				entsql.IndexWhere("active"),
			),	
		// Define a custom operator class.
		index.Fields("phone").
			Annotations(
				entsql.OpClass("bpchar_pattern_ops"),
			),
    }
}
```

The code above generates the following SQL statements:

```sql
CREATE INDEX `users_c1` ON `users` (`c1` DESC)

CREATE INDEX `users_c1_c2_c3` ON `users` (`c1` DESC, `c2` DESC, `c3`)

CREATE INDEX `users_c4` ON `users` USING HASH (`c4`)

-- MySQL only.
CREATE FULLTEXT INDEX `users_c5` ON `users` (`c5`)

-- PostgreSQL only.
CREATE INDEX "users_c5" ON "users" USING GIN ("c5")

-- Include index-only scan on PostgreSQL.
CREATE INDEX "users_workplace" ON "users" ("workplace") INCLUDE ("address")

-- Define partial index on SQLite and PostgreSQL.
CREATE INDEX "users_nickname" ON "users" ("nickname") WHERE "active"

-- PostgreSQL only.
CREATE INDEX "users_phone" ON "users" ("phone" bpchar_pattern_ops)
```

## Functional Indexes

The Ent schema supports defining indexes on fields and edges (foreign-keys), but there is no API for defining index
parts as expressions, such as function calls. If you are using [Atlas](https://atlasgo.io/docs) for managing schema
migrations, you can define functional indexes as described in [this guide](/docs/migration/functional-indexes).

## Storage Key

Like Fields, custom index name can be configured using the `StorageKey` method.
It's mapped to an index name in SQL dialects.

```go
func (User) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("field1", "field2").
			StorageKey("custom_index"),
	}
}
```
