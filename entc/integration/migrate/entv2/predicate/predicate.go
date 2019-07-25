// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fmt"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/sql"
)

// Group is the predicate function for group builders.
type Group func(interface{})

// GroupPerDialect construct a predicate for both gremlin traversal and sql selector.
func GroupPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Group {
	return Group(func(v interface{}) {
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

// Pet is the predicate function for pet builders.
type Pet func(interface{})

// PetPerDialect construct a predicate for both gremlin traversal and sql selector.
func PetPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Pet {
	return Pet(func(v interface{}) {
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
