// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/sql"
	"fmt"
)

// User is the predicate function for user builders.
type User func(interface{})

// UserPerDialect construct a predicate for both gremlin traversal and sql selector.
func UserPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) User {
	return User(func(v interface{}) {
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
