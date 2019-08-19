// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Order applies an ordering on either graph traversal or sql selector.
type Order func(*sql.Selector)

// Asc applies the given fields in ASC order.
func Asc(fields ...string) Order {
	return Order(
		func(s *sql.Selector) {
			for _, f := range fields {
				s.OrderBy(sql.Asc(f))
			}
		},
	)
}

// Desc applies the given fields in DESC order.
func Desc(fields ...string) Order {
	return Order(
		func(s *sql.Selector) {
			for _, f := range fields {
				s.OrderBy(sql.Desc(f))
			}
		},
	)
}

// Aggregate applies an aggregation step on the group-by traversal/selector.
type Aggregate struct {
	// SQL the column wrapped with the aggregation function.
	SQL func(*sql.Selector) string
}

// As is a pseudo aggregation function for renaming another other functions with custom names. For example:
//
//	GroupBy(field1, field2).
//	Aggregate(entv2.As(entv2.Sum(field1), "sum_field1"), (entv2.As(entv2.Sum(field2), "sum_field2")).
//	Scan(ctx, &v)
//
func As(fn Aggregate, end string) Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.As(fn.SQL(s), end)
		},
	}
}

// Count applies the "count" aggregation function on each group.
func Count() Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.Count("*")
		},
	}
}

// Max applies the "max" aggregation function on the given field of each group.
func Max(field string) Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.Max(s.C(field))
		},
	}
}

// Mean applies the "mean" aggregation function on the given field of each group.
func Mean(field string) Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.Avg(s.C(field))
		},
	}
}

// Min applies the "min" aggregation function on the given field of each group.
func Min(field string) Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.Min(s.C(field))
		},
	}
}

// Sum applies the "sum" aggregation function on the given field of each group.
func Sum(field string) Aggregate {
	return Aggregate{
		SQL: func(s *sql.Selector) string {
			return sql.Sum(s.C(field))
		},
	}
}

// ErrNotFound returns when trying to fetch a specific entity and it was not found in the database.
type ErrNotFound struct {
	label string
}

// Error implements the error interface.
func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("entv2: %s not found", e.label)
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
	return fmt.Sprintf("entv2: %s not singular", e.label)
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
	return fmt.Sprintf("entv2: unique constraint failed: %s", e.msg)
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

func isSQLConstraintError(err error) (*ErrConstraintFailed, bool) {
	// Error number 1062 is ER_DUP_ENTRY in mysql, and "UNIQUE constraint failed" is SQLite prefix.
	if msg := err.Error(); strings.HasPrefix(msg, "Error 1062") || strings.HasPrefix(msg, "UNIQUE constraint failed") {
		return &ErrConstraintFailed{msg, err}, true
	}
	return nil, false
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	if err, ok := isSQLConstraintError(err); ok {
		return err
	}
	return err
}

// keys returns the keys/ids from the edge map.
func keys(m map[int]struct{}) []int {
	s := make([]int, 0, len(m))
	for id, _ := range m {
		s = append(s, id)
	}
	return s
}
