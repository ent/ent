// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package p

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

// EQ is the equal predicate.
func EQ(v any) *dsl.Traversal {
	return op("eq", v)
}

// NEQ is the not-equal predicate.
func NEQ(v any) *dsl.Traversal {
	return op("neq", v)
}

// GT is the greater than predicate.
func GT(v any) *dsl.Traversal {
	return op("gt", v)
}

// GTE is the greater than or equal predicate.
func GTE(v any) *dsl.Traversal {
	return op("gte", v)
}

// LT is the less than predicate.
func LT(v any) *dsl.Traversal {
	return op("lt", v)
}

// LTE is the less than or equal predicate.
func LTE(v any) *dsl.Traversal {
	return op("lte", v)
}

// Between is the between/contains predicate.
func Between(v, u any) *dsl.Traversal {
	return op("between", v, u)
}

// StartingWith is the prefix test predicate.
func StartingWith(prefix string) *dsl.Traversal {
	return op("startingWith", prefix)
}

// EndingWith is the suffix test predicate.
func EndingWith(suffix string) *dsl.Traversal {
	return op("endingWith", suffix)
}

// Containing is the sub string test predicate.
func Containing(substr string) *dsl.Traversal {
	return op("containing", substr)
}

// NotStartingWith is the negation of StartingWith.
func NotStartingWith(prefix string) *dsl.Traversal {
	return op("notStartingWith", prefix)
}

// NotEndingWith is the negation of EndingWith.
func NotEndingWith(suffix string) *dsl.Traversal {
	return op("notEndingWith", suffix)
}

// NotContaining is the negation of Containing.
func NotContaining(substr string) *dsl.Traversal {
	return op("notContaining", substr)
}

// Within Determines if a value is within the specified list of values.
func Within[T any](args ...T) *dsl.Traversal {
	return op("within", args...)
}

// Without determines if a value is not within the specified list of values.
func Without[T any](args ...T) *dsl.Traversal {
	return op("without", args...)
}

func op[T any](name string, args ...T) *dsl.Traversal {
	t := &dsl.Traversal{}
	vs := make([]any, len(args))
	for i, arg := range args {
		vs[i] = arg
	}
	return t.Add(dsl.NewFunc(name, vs...))
}
