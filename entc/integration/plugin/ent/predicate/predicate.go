// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fmt"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/sql"
)

// Boring is the predicate function for boring builders.
type Boring func(interface{})

// BoringPerDialect construct a predicate for graph traversals based on dialect type.
func BoringPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Boring {
	return Boring(func(v interface{}) {
		switch v := v.(type) {
		case *sql.Selector:
			f0(v)
		case *dsl.Traversal:
			f1(v)
		default:
			panic(fmt.Sprintf("unknown type for predicate: %T", v))
		}
	})
}
