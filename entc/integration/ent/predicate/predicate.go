// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/sql"
	"fmt"
)

// Card is the predicate function for card builders.
type Card func(interface{})

// CardPerDialect construct a predicate for both gremlin traversal and sql selector.
func CardPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Card {
	return Card(func(v interface{}) {
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

// Comment is the predicate function for comment builders.
type Comment func(interface{})

// CommentPerDialect construct a predicate for both gremlin traversal and sql selector.
func CommentPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Comment {
	return Comment(func(v interface{}) {
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

// FieldType is the predicate function for fieldtype builders.
type FieldType func(interface{})

// FieldTypePerDialect construct a predicate for both gremlin traversal and sql selector.
func FieldTypePerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) FieldType {
	return FieldType(func(v interface{}) {
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

// File is the predicate function for file builders.
type File func(interface{})

// FilePerDialect construct a predicate for both gremlin traversal and sql selector.
func FilePerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) File {
	return File(func(v interface{}) {
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

// GroupInfo is the predicate function for groupinfo builders.
type GroupInfo func(interface{})

// GroupInfoPerDialect construct a predicate for both gremlin traversal and sql selector.
func GroupInfoPerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) GroupInfo {
	return GroupInfo(func(v interface{}) {
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

// Node is the predicate function for node builders.
type Node func(interface{})

// NodePerDialect construct a predicate for both gremlin traversal and sql selector.
func NodePerDialect(f1 func(*sql.Selector), f2 func(*dsl.Traversal)) Node {
	return Node(func(v interface{}) {
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
