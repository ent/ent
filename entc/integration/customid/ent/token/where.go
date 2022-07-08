// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package token

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/sid"
)

// ID filters vertices based on their ID field.
func ID(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
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
func IDNotIn(ids ...sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
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
func IDGT(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id sid.ID) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Token {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Token(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBody), v))
	})
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBody), v))
	})
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBody), v))
	})
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBody), v))
	})
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBody), v))
	})
}

// HasAccount applies the HasEdge predicate on the "account" edge.
func HasAccount() predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAccountWith applies the HasEdge predicate on the "account" edge with a given conditions (other predicates).
func HasAccountWith(preds ...predicate.Account) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AccountInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AccountTable, AccountColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
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
func Not(p predicate.Token) predicate.Token {
	return predicate.Token(func(s *sql.Selector) {
		p(s.Not())
	})
}
