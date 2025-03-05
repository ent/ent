package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
)

// base holds the default configuration for most schemas in this package.
type base struct {
	ent.Schema
}

// Annotations of the base schema.
func (base) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Schema("db1"),
	}
}
