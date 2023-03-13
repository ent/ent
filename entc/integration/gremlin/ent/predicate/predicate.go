// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

// Api is the predicate function for api builders.
type Api func(*dsl.Traversal)

// Card is the predicate function for card builders.
type Card func(*dsl.Traversal)

// Comment is the predicate function for comment builders.
type Comment func(*dsl.Traversal)

// ExValueScan is the predicate function for exvaluescan builders.
type ExValueScan func(*dsl.Traversal)

// ExValueScanOrErr calls the predicate only if the error is not nit.
func ExValueScanOrErr(p ExValueScan, err error) ExValueScan {
	return func(s *dsl.Traversal) {
		if err != nil {
			s.AddError(err)
			return
		}
		p(s)
	}
}

// FieldType is the predicate function for fieldtype builders.
type FieldType func(*dsl.Traversal)

// File is the predicate function for file builders.
type File func(*dsl.Traversal)

// FileType is the predicate function for filetype builders.
type FileType func(*dsl.Traversal)

// Goods is the predicate function for goods builders.
type Goods func(*dsl.Traversal)

// Group is the predicate function for group builders.
type Group func(*dsl.Traversal)

// GroupInfo is the predicate function for groupinfo builders.
type GroupInfo func(*dsl.Traversal)

// Item is the predicate function for item builders.
type Item func(*dsl.Traversal)

// License is the predicate function for license builders.
type License func(*dsl.Traversal)

// Node is the predicate function for node builders.
type Node func(*dsl.Traversal)

// Pet is the predicate function for pet builders.
type Pet func(*dsl.Traversal)

// Spec is the predicate function for spec builders.
type Spec func(*dsl.Traversal)

// Task is the predicate function for enttask builders.
type Task func(*dsl.Traversal)

// User is the predicate function for user builders.
type User func(*dsl.Traversal)
