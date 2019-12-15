// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/encoding/graphson"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
)

// Order applies an ordering on either graph traversal or sql selector.
type Order func(*dsl.Traversal)

// Asc applies the given fields in ASC order.
func Asc(fields ...string) Order {
	return func(tr *dsl.Traversal) {
		for _, f := range fields {
			tr.By(f, dsl.Incr)
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) Order {
	return func(tr *dsl.Traversal) {
		for _, f := range fields {
			tr.By(f, dsl.Decr)
		}
	}
}

// Aggregate applies an aggregation step on the group-by traversal/selector.
// It gets two labels as parameters. The first used in the `As` step for the predicate,
// and the second is an optional name for the next predicates (or for later usage).
type Aggregate func(string, string) (string, *dsl.Traversal)

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(ent.As(ent.Sum(field1), "sum_field1"), (ent.As(ent.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn Aggregate, end string) Aggregate {
	return func(start, _ string) (string, *dsl.Traversal) {
		return fn(start, end)
	}
}

// DefaultCountLabel is the default label name for the Count aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "count" 2 or more fields and avoid conflicting, use the `ent.As(ent.Count(field), "custom_name")`
// function with custom name in order to override it.
const DefaultCountLabel = "count"

// Count applies the "count" aggregation function on each group.
func Count() Aggregate {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultCountLabel
		}
		return end, __.As(start).Count(dsl.Local).As(end)
	}
}

// DefaultMaxLabel is the default label name for the Max aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "max" 2 or more fields and avoid conflicting, use the `ent.As(ent.Max(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMaxLabel = "max"

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) Aggregate {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMaxLabel
		}
		return end, __.As(start).Unfold().Values(field).Max().As(end)
	}
}

// DefaultMeanLabel is the default label name for the Mean aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "mean" 2 or more fields and avoid conflicting, use the `ent.As(ent.Mean(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMeanLabel = "mean"

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) Aggregate {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMeanLabel
		}
		return end, __.As(start).Unfold().Values(field).Mean().As(end)
	}
}

// DefaultMinLabel is the default label name for the Min aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "min" 2 or more fields and avoid conflicting, use the `ent.As(ent.Min(field), "custom_name")`
// function with custom name in order to override it.
const DefaultMinLabel = "min"

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) Aggregate {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultMinLabel
		}
		return end, __.As(start).Unfold().Values(field).Min().As(end)
	}
}

// DefaultSumLabel is the default label name for the Sum aggregation function.
// It should be used as the struct-tag for decoding, or a map key for interaction with the returned response.
// In order to "sum" 2 or more fields and avoid conflicting, use the `ent.As(ent.Sum(field), "custom_name")`
// function with custom name in order to override it.
const DefaultSumLabel = "sum"

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) Aggregate {
	return func(start, end string) (string, *dsl.Traversal) {
		if end == "" {
			end = DefaultSumLabel
		}
		return end, __.As(start).Unfold().Values(field).Sum().As(end)
	}
}

// ErrNotFound returns when trying to fetch a specific entity and it was not found in the database.
type ErrNotFound struct {
	label string
}

// Error implements the error interface.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("ent: %s not found", e.label)
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	_, ok := err.(*ErrNotFound)
	return ok
}

// MaskNotFound masks nor found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// ErrNotSingular returns when trying to fetch a singular entity and more then one was found in the database.
type ErrNotSingular struct {
	label string
}

// Error implements the error interface.
func (e *ErrNotSingular) Error() string {
	return fmt.Sprintf("ent: %s not singular", e.label)
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	_, ok := err.(*ErrNotSingular)
	return ok
}

// ErrConstraintFailed returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or field uniqueness.
type ErrConstraintFailed struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ErrConstraintFailed) Error() string {
	return fmt.Sprintf("ent: unique constraint failed: %s", e.msg)
}

// Unwrap implements the errors.Wrapper interface.
func (e *ErrConstraintFailed) Unwrap() error {
	return e.wrap
}

// IsConstraintFailure returns a boolean indicating whether the error is a constraint failure.
func IsConstraintFailure(err error) bool {
	_, ok := err.(*ErrConstraintFailed)
	return ok
}

// Code implements the dsl.Node interface.
func (e ErrConstraintFailed) Code() (string, []interface{}) {
	return strconv.Quote(e.prefix() + e.msg), nil
}

func (e *ErrConstraintFailed) UnmarshalGraphson(b []byte) error {
	var v [1]*string
	if err := graphson.Unmarshal(b, &v); err != nil {
		return err
	}
	if v[0] == nil {
		return fmt.Errorf("ent: missing string value")
	}
	if !strings.HasPrefix(*v[0], e.prefix()) {
		return fmt.Errorf("ent: invalid string for error: %s", *v[0])
	}
	e.msg = strings.TrimPrefix(*v[0], e.prefix())
	return nil
}

// prefix returns the prefix used for gremlin constants.
func (ErrConstraintFailed) prefix() string { return "Error: " }

// NewErrUniqueField creates a constraint error for unique fields.
func NewErrUniqueField(label, field string, v interface{}) *ErrConstraintFailed {
	return &ErrConstraintFailed{msg: fmt.Sprintf("field %s.%s with value: %#v", label, field, v)}
}

// NewErrUniqueEdge creates a constraint error for unique edges.
func NewErrUniqueEdge(label, edge, id string) *ErrConstraintFailed {
	return &ErrConstraintFailed{msg: fmt.Sprintf("edge %s.%s with id: %#v", label, edge, id)}
}

// isConstantError indicates if the given response holds a gremlin constant containing an error.
func isConstantError(r *gremlin.Response) (*ErrConstraintFailed, bool) {
	e := &ErrConstraintFailed{}
	if err := graphson.Unmarshal(r.Result.Data, e); err != nil {
		return nil, false
	}
	return e, true
}

// keys returns the keys/ids from the edge map.
func keys(m map[string]struct{}) []string {
	s := make([]string, 0, len(m))
	for id := range m {
		s = append(s, id)
	}
	return s
}
