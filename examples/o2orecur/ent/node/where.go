// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package node

import (
	"github.com/facebook/ent/dialect/sql"
	"github.com/facebook/ent/dialect/sql/sqlgraph"
	"github.com/facebook/ent/examples/o2orecur/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if no arguments were provided, append the FALSE constants to make
		// the predicate "falsy", because we can't apply `IN` on empty list.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if no arguments were provided, append the FALSE constants to make
		// the predicate "falsy", because we can't apply `IN` on empty list.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if no arguments were provided, append the FALSE constants to make
		// the predicate "falsy", because we can't apply `IN` on empty list.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if no arguments were provided, append the FALSE constants to make
		// the predicate "falsy", because we can't apply `IN` on empty list.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrevTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrevTable, PrevColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(NextTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, NextTable, NextColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		p(s.Not())
	})
}
