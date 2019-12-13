// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package node

import (
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Node {
	return predicate.Node(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i], _ = strconv.Atoi(ids[i])
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		id, _ := strconv.Atoi(id)
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	},
	)
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	},
	)
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	},
	)
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	},
	)
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) predicate.Node {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Node(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	},
	)
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	},
	)
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	},
	)
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	},
	)
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	},
	)
}

// ValueIsNil applies the IsNil predicate on the "value" field.
func ValueIsNil() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldValue)))
	},
	)
}

// ValueNotNil applies the NotNil predicate on the "value" field.
func ValueNotNil() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldValue)))
	},
	)
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(PrevTable, FieldID),
			sql.Edge(sql.O2O, true, PrevTable, PrevColumn),
		)
		sql.HasNeighbors(s, step)
	},
	)
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(Table, FieldID),
			sql.Edge(sql.O2O, true, PrevTable, PrevColumn),
		)
		sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(NextTable, FieldID),
			sql.Edge(sql.O2O, false, NextTable, NextColumn),
		)
		sql.HasNeighbors(s, step)
	},
	)
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...predicate.Node) predicate.Node {
	return predicate.Node(func(s *sql.Selector) {
		step := sql.NewStep(
			sql.From(Table, FieldID),
			sql.To(Table, FieldID),
			sql.Edge(sql.O2O, false, NextTable, NextColumn),
		)
		sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Node) predicate.Node {
	return predicate.Node(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Node) predicate.Node {
	return predicate.Node(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
