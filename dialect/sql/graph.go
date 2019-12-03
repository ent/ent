// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/schema/field"
)

// Rel is a relation type of an edge.
type Rel int

// Relation types.
const (
	Unk Rel = iota // Unknown.
	O2O            // One to one / has one.
	O2M            // One to many / has many.
	M2O            // Many to one (inverse perspective for O2M).
	M2M            // Many to many.
)

// String returns the relation name.
func (r Rel) String() (s string) {
	switch r {
	case O2O:
		s = "O2O"
	case O2M:
		s = "O2M"
	case M2O:
		s = "M2O"
	case M2M:
		s = "M2M"
	default:
		s = "Unknown"
	}
	return s
}

// A Step provides a path-step information to the traversal functions.
type Step struct {
	// From is the source of the step.
	From struct {
		// V can be either one vertex or set of vertices.
		// It can be a pre-processed step (sql.Query) or a simple Go type (integer or string).
		V interface{}
		// Table holds the table name of V (from).
		Table string
		// Column to join with. Usually the "id" column.
		Column string
	}
	// Edge holds the edge information for getting the neighbors.
	Edge struct {
		// Rel of the edge.
		Rel Rel
		// Table name of where this edge columns reside.
		Table string
		// Columns of the edge.
		// In O2O and M2O, it holds the foreign-key column. Hence, len == 1.
		// In M2M, it holds the primary-key columns of the join table. Hence, len == 2.
		Columns []string
		// Inverse indicates if the edge is an inverse edge.
		Inverse bool
	}
	// To is the dest of the path (the neighbors).
	To struct {
		// Table holds the table name of the neighbors (to).
		Table string
		// Column to join with. Usually the "id" column.
		Column string
	}
}

// StepOption allows configuring Steps using functional options.
type StepOption func(*Step)

// From sets the source of the step.
func From(table, column string, v ...interface{}) StepOption {
	return func(s *Step) {
		s.From.Table = table
		s.From.Column = column
		if len(v) > 0 {
			s.From.V = v[0]
		}
	}
}

// To sets the destination of the step.
func To(table, column string) StepOption {
	return func(s *Step) {
		s.To.Table = table
		s.To.Column = column
	}
}

// Edge sets the edge info for getting the neighbors.
func Edge(rel Rel, inverse bool, table string, columns ...string) StepOption {
	return func(s *Step) {
		s.Edge.Rel = rel
		s.Edge.Table = table
		s.Edge.Columns = columns
		s.Edge.Inverse = inverse
	}
}

// NewStep gets list of options and returns a configured step.
//
//	NewStep(
//		From("table", "pk", V),
//		To("table", "pk"),
//		Edge("name", O2M, "fk"),
//	)
//
func NewStep(opts ...StepOption) *Step {
	s := &Step{}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// Neighbors returns a Selector for evaluating the path-step
// and getting the neighbors of one vertex.
func Neighbors(dialect string, s *Step) (q *Selector) {
	builder := Dialect(dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		to := builder.Table(s.To.Table)
		join := builder.Table(s.Edge.Table)
		match := builder.Select(join.C(pk1)).
			From(join).
			Where(EQ(join.C(pk2), s.From.V))
		q = builder.Select().
			From(to).
			Join(match).
			On(to.C(s.To.Column), match.C(pk1))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		t2 := builder.Select(s.Edge.Columns[0]).
			From(builder.Table(s.Edge.Table)).
			Where(EQ(s.From.Column, s.From.V))
		q = builder.Select().
			From(t1).
			Join(t2).
			On(t1.C(s.From.Column), t2.C(s.Edge.Columns[0]))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		q = builder.Select().
			From(builder.Table(s.To.Table)).
			Where(EQ(s.Edge.Columns[0], s.From.V))
	}
	return q
}

// SetNeighbors returns a Selector for evaluating the path-step
// and getting the neighbors of set of vertices.
func SetNeighbors(dialect string, s *Step) (q *Selector) {
	set := s.From.V.(*Selector)
	builder := Dialect(dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		to := builder.Table(s.To.Table)
		set.Select(set.C(s.From.Column))
		join := builder.Table(s.Edge.Table)
		match := builder.Select(join.C(pk1)).
			From(join).
			Join(set).
			On(join.C(pk2), set.C(s.From.Column))
		q = builder.Select().
			From(to).
			Join(match).
			On(to.C(s.To.Column), match.C(pk1))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		set.Select(set.C(s.Edge.Columns[0]))
		q = builder.Select().
			From(t1).
			Join(set).
			On(t1.C(s.To.Column), set.C(s.Edge.Columns[0]))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		t1 := builder.Table(s.To.Table)
		set.Select(set.C(s.From.Column))
		q = builder.Select().
			From(t1).
			Join(set).
			On(t1.C(s.Edge.Columns[0]), set.C(s.From.Column))
	}
	return q
}

// HasNeighbors applies on the given Selector a neighbors check.
func HasNeighbors(q *Selector, s *Step) {
	builder := Dialect(q.dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1 := s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1 = s.Edge.Columns[1]
		}
		from := q.Table()
		join := builder.Table(s.Edge.Table)
		q.Where(
			In(
				from.C(s.From.Column),
				builder.Select(join.C(pk1)).From(join),
			),
		)
	case r == M2O || (r == O2O && s.Edge.Inverse):
		from := q.Table()
		q.Where(NotNull(from.C(s.Edge.Columns[0])))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.Edge.Table)
		q.Where(
			In(
				from.C(s.From.Column),
				builder.Select(to.C(s.Edge.Columns[0])).
					From(to).
					Where(NotNull(to.C(s.Edge.Columns[0]))),
			),
		)
	}
}

// HasNeighborsWith applies on the given Selector a neighbors check.
// The given predicate applies its filtering on the selector.
func HasNeighborsWith(q *Selector, s *Step, pred func(*Selector)) {
	builder := Dialect(q.dialect)
	switch r := s.Edge.Rel; {
	case r == M2M:
		pk1, pk2 := s.Edge.Columns[1], s.Edge.Columns[0]
		if s.Edge.Inverse {
			pk1, pk2 = pk2, pk1
		}
		from := q.Table()
		to := builder.Table(s.To.Table)
		edge := builder.Table(s.Edge.Table)
		join := builder.Select(edge.C(pk2)).
			From(edge).
			Join(to).
			On(edge.C(pk1), to.C(s.To.Column))
		matches := builder.Select().From(to)
		pred(matches)
		join.FromSelect(matches)
		q.Where(In(from.C(s.From.Column), join))
	case r == M2O || (r == O2O && s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.To.Table)
		matches := builder.Select(to.C(s.To.Column)).
			From(to)
		pred(matches)
		q.Where(In(from.C(s.Edge.Columns[0]), matches))
	case r == O2M || (r == O2O && !s.Edge.Inverse):
		from := q.Table()
		to := builder.Table(s.Edge.Table)
		matches := builder.Select(to.C(s.Edge.Columns[0])).
			From(to)
		pred(matches)
		q.Where(In(from.C(s.From.Column), matches))
	}
}

type (
	// FieldSpec holds the information for updating a field
	// column in the database.
	FieldSpec struct {
		Column string
		Type   field.Type
		Value  driver.Value // value to be stored.
	}

	// EdgeSpec holds the information for updating a field
	// column in the database.
	EdgeSpec struct {
		Rel     Rel
		Table   string
		Columns []string
		Inverse bool
		Value   driver.Value
	}

	// CreateSpec holds the information for creating a node
	// in the graph.
	CreateSpec struct {
		// Type or table name.
		Table string
		// ID field.
		ID *FieldSpec
		// Fields.
		Fields []*FieldSpec
		// Edges.
		Edges []*EdgeSpec
	}
)

// CreateNode applies the spec on the graph.
func CreateNode(ctx context.Context, drv dialect.Driver, spec *CreateSpec) error {
	tx, err := drv.Tx(ctx)
	if err != nil {
		return err
	}
	insert := Dialect(drv.Dialect()).Insert(spec.Table).Default()
	for _, fi := range spec.Fields {
		value := fi.Value
		if fi.Type == field.TypeJSON {
			if value, err = json.Marshal(value); err != nil {
				return err
			}
		}
		insert.Set(fi.Column, value)
	}
	// ID was provided by the user.
	if spec.ID.Value != nil {
		insert.Set(spec.ID.Column, spec.ID.Value)
		query, args := insert.Query()
		if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
			return rollback(tx, err)
		}
	} else {
		id, err := insertLastID(ctx, tx, insert.Returning(spec.ID.Column))
		if err != nil {
			return rollback(tx, err)
		}
		spec.ID.Value = id
	}
	return tx.Commit()
}

// insertLastID invokes the insert query on the transaction and returns the LastInsertID.
func insertLastID(ctx context.Context, tx dialect.ExecQuerier, insert *InsertBuilder) (int64, error) {
	query, args := insert.Query()
	// PostgreSQL does not support the LastInsertId() method of sql.Result
	// on Exec, and should be extracted manually using the `RETURNING` clause.
	if insert.Dialect() == dialect.Postgres {
		rows := &sql.Rows{}
		if err := tx.Query(ctx, query, args, rows); err != nil {
			return 0, err
		}
		defer rows.Close()
		return ScanInt64(rows)
	}
	// MySQL, SQLite, etc.
	var res sql.Result
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	return res.LastInsertId()
}

// rollback calls to tx.Rollback and wraps the given error with the rollback error if occurred.
func rollback(tx dialect.Tx, err error) error {
	if rerr := tx.Rollback(); rerr != nil {
		err = fmt.Errorf("%s: %v", err.Error(), rerr)
	}
	return err
}
