// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ent aliases to avoid import conflicts in user's code.
type (
	Op         = ent.Op
	Hook       = ent.Hook
	Value      = ent.Value
	Query      = ent.Query
	Policy     = ent.Policy
	Mutator    = ent.Mutator
	Mutation   = ent.Mutation
	MutateFunc = ent.MutateFunc
	LockType   = ent.LockType
)

const (
	LockForUpdate = ent.LockForUpdate
	LockForShare  = ent.LockForShare
)

// OrderFunc applies an ordering on the sql selector.
type OrderFunc func(*sql.Selector, func(string) bool)

// Asc applies the given fields in ASC order.
func Asc(fields ...string) OrderFunc {
	return func(s *sql.Selector, check func(string) bool) {
		for _, f := range fields {
			if check(f) {
				s.OrderBy(sql.Asc(f))
			} else {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("invalid field %q for ordering", f)})
			}
		}
	}
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) OrderFunc {
	return func(s *sql.Selector, check func(string) bool) {
		for _, f := range fields {
			if check(f) {
				s.OrderBy(sql.Desc(f))
			} else {
				s.AddError(&ValidationError{Name: f, err: fmt.Errorf("invalid field %q for ordering", f)})
			}
		}
	}
}

// AggregateFunc applies an aggregation step on the group-by traversal/selector.
type AggregateFunc func(*sql.Selector, func(string) bool) string

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(entv1.As(entv1.Sum(field1), "sum_field1"), (entv1.As(entv1.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn AggregateFunc, end string) AggregateFunc {
	return func(s *sql.Selector, check func(string) bool) string {
		return sql.As(fn(s, check), end)
	}
}

// Count applies the "count" aggregation function on each group.
func Count() AggregateFunc {
	return func(s *sql.Selector, _ func(string) bool) string {
		return sql.Count("*")
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) AggregateFunc {
	return func(s *sql.Selector, check func(string) bool) string {
		if !check(field) {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("invalid field %q for grouping", field)})
			return ""
		}
		return sql.Max(s.C(field))
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) AggregateFunc {
	return func(s *sql.Selector, check func(string) bool) string {
		if !check(field) {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("invalid field %q for grouping", field)})
			return ""
		}
		return sql.Avg(s.C(field))
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) AggregateFunc {
	return func(s *sql.Selector, check func(string) bool) string {
		if !check(field) {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("invalid field %q for grouping", field)})
			return ""
		}
		return sql.Min(s.C(field))
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) AggregateFunc {
	return func(s *sql.Selector, check func(string) bool) string {
		if !check(field) {
			s.AddError(&ValidationError{Name: field, err: fmt.Errorf("invalid field %q for grouping", field)})
			return ""
		}
		return sql.Sum(s.C(field))
	}
}

// ValidationError returns when validating a field fails.
type ValidationError struct {
	Name string // Field or edge name.
	err  error
}

// Error implements the error interface.
func (e *ValidationError) Error() string {
	return e.err.Error()
}

// Unwrap implements the errors.Wrapper interface.
func (e *ValidationError) Unwrap() error {
	return e.err
}

// IsValidationError returns a boolean indicating whether the error is a validaton error.
func IsValidationError(err error) bool {
	if err == nil {
		return false
	}
	var e *ValidationError
	return errors.As(err, &e)
}

// NotFoundError returns when trying to fetch a specific entity and it was not found in the database.
type NotFoundError struct {
	label string
}

// Error implements the error interface.
func (e *NotFoundError) Error() string {
	return "entv1: " + e.label + " not found"
}

// IsNotFound returns a boolean indicating whether the error is a not found error.
func IsNotFound(err error) bool {
	if err == nil {
		return false
	}
	var e *NotFoundError
	return errors.As(err, &e)
}

// MaskNotFound masks not found error.
func MaskNotFound(err error) error {
	if IsNotFound(err) {
		return nil
	}
	return err
}

// NotSingularError returns when trying to fetch a singular entity and more then one was found in the database.
type NotSingularError struct {
	label string
}

// Error implements the error interface.
func (e *NotSingularError) Error() string {
	return "entv1: " + e.label + " not singular"
}

// IsNotSingular returns a boolean indicating whether the error is a not singular error.
func IsNotSingular(err error) bool {
	if err == nil {
		return false
	}
	var e *NotSingularError
	return errors.As(err, &e)
}

// NotLoadedError returns when trying to get a node that was not loaded by the query.
type NotLoadedError struct {
	edge string
}

// Error implements the error interface.
func (e *NotLoadedError) Error() string {
	return "entv1: " + e.edge + " edge was not loaded"
}

// IsNotLoaded returns a boolean indicating whether the error is a not loaded error.
func IsNotLoaded(err error) bool {
	if err == nil {
		return false
	}
	var e *NotLoadedError
	return errors.As(err, &e)
}

// ConstraintError returns when trying to create/update one or more entities and
// one or more of their constraints failed. For example, violation of edge or
// field uniqueness.
type ConstraintError struct {
	msg  string
	wrap error
}

// Error implements the error interface.
func (e ConstraintError) Error() string {
	return "entv1: constraint failed: " + e.msg
}

// Unwrap implements the errors.Wrapper interface.
func (e *ConstraintError) Unwrap() error {
	return e.wrap
}

// IsConstraintError returns a boolean indicating whether the error is a constraint failure.
func IsConstraintError(err error) bool {
	if err == nil {
		return false
	}
	var e *ConstraintError
	return errors.As(err, &e)
}

func isSQLConstraintError(err error) (*ConstraintError, bool) {
	if sqlgraph.IsConstraintError(err) {
		return &ConstraintError{err.Error(), err}, true
	}
	return nil, false
}

// rollback calls tx.Rollback and wraps the given error with the rollback error if present.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%w: %v", err, rerr)
	}
	if err, ok := isSQLConstraintError(err); ok {
		return err
	}
	return err
}
