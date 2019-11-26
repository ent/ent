// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

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
		join := builder.Table(s.Edge.Table)
		matches := builder.Select(join.C(pk2)).
			From(join).
			Join(to).
			On(join.C(pk1), to.C(s.To.Column))
		pred(matches)
		q.Where(In(from.C(s.From.Column), matches))
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
