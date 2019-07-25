// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fmt"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/sql"
)

// Boring is the predicate function for boring builders.
type Boring func(interface{})

// BoringPerDialect construct a predicate for both gremlin traversal and sql selector.
func BoringPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Boring {
	return Boring(func(v interface{}) {
		switch v := v.(type) {
		case *sql.Selector:
			f1(v)
		case *dsl.Traversal:
			f2(v)
		default:
			panic(fmt.Sprintf("unknown type for predicate: %T", v))
		}
	})
}
