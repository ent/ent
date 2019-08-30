// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package predicate

import (
	"fmt"

	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Card is the predicate function for card builders.
type Card func(interface{})

// CardPerDialect construct a predicate for graph traversals based on dialect type.
func CardPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Card {
	return Card(func(v interface{}) {
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

// Comment is the predicate function for comment builders.
type Comment func(interface{})

// CommentPerDialect construct a predicate for graph traversals based on dialect type.
func CommentPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Comment {
	return Comment(func(v interface{}) {
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

// FieldType is the predicate function for fieldtype builders.
type FieldType func(interface{})

// FieldTypePerDialect construct a predicate for graph traversals based on dialect type.
func FieldTypePerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) FieldType {
	return FieldType(func(v interface{}) {
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

// File is the predicate function for file builders.
type File func(interface{})

// FilePerDialect construct a predicate for graph traversals based on dialect type.
func FilePerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) File {
	return File(func(v interface{}) {
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

// FileType is the predicate function for filetype builders.
type FileType func(interface{})

// FileTypePerDialect construct a predicate for graph traversals based on dialect type.
func FileTypePerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) FileType {
	return FileType(func(v interface{}) {
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

// Group is the predicate function for group builders.
type Group func(interface{})

// GroupPerDialect construct a predicate for graph traversals based on dialect type.
func GroupPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Group {
	return Group(func(v interface{}) {
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

// GroupInfo is the predicate function for groupinfo builders.
type GroupInfo func(interface{})

// GroupInfoPerDialect construct a predicate for graph traversals based on dialect type.
func GroupInfoPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) GroupInfo {
	return GroupInfo(func(v interface{}) {
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

// Node is the predicate function for node builders.
type Node func(interface{})

// NodePerDialect construct a predicate for graph traversals based on dialect type.
func NodePerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Node {
	return Node(func(v interface{}) {
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

// Pet is the predicate function for pet builders.
type Pet func(interface{})

// PetPerDialect construct a predicate for graph traversals based on dialect type.
func PetPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) Pet {
	return Pet(func(v interface{}) {
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

// User is the predicate function for user builders.
type User func(interface{})

// UserPerDialect construct a predicate for graph traversals based on dialect type.
func UserPerDialect(f0 func(*sql.Selector), f1 func(*dsl.Traversal)) User {
	return User(func(v interface{}) {
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
