// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package grouptag

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// TagID applies equality check predicate on the "tag_id" field. It's identical to TagIDEQ.
func TagID(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// GroupID applies equality check predicate on the "group_id" field. It's identical to GroupIDEQ.
func GroupID(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// TagIDEQ applies the EQ predicate on the "tag_id" field.
func TagIDEQ(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTagID), v))
	})
}

// TagIDNEQ applies the NEQ predicate on the "tag_id" field.
func TagIDNEQ(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTagID), v))
	})
}

// TagIDIn applies the In predicate on the "tag_id" field.
func TagIDIn(vs ...int) predicate.GroupTag {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldTagID), v...))
	})
}

// TagIDNotIn applies the NotIn predicate on the "tag_id" field.
func TagIDNotIn(vs ...int) predicate.GroupTag {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldTagID), v...))
	})
}

// GroupIDEQ applies the EQ predicate on the "group_id" field.
func GroupIDEQ(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGroupID), v))
	})
}

// GroupIDNEQ applies the NEQ predicate on the "group_id" field.
func GroupIDNEQ(v int) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGroupID), v))
	})
}

// GroupIDIn applies the In predicate on the "group_id" field.
func GroupIDIn(vs ...int) predicate.GroupTag {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldGroupID), v...))
	})
}

// GroupIDNotIn applies the NotIn predicate on the "group_id" field.
func GroupIDNotIn(vs ...int) predicate.GroupTag {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupTag(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldGroupID), v...))
	})
}

// HasTag applies the HasEdge predicate on the "tag" edge.
func HasTag() predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTagWith applies the HasEdge predicate on the "tag" edge with a given conditions (other predicates).
func HasTagWith(preds ...predicate.Tag) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(TagInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TagTable, TagColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroup applies the HasEdge predicate on the "group" edge.
func HasGroup() predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupWith applies the HasEdge predicate on the "group" edge with a given conditions (other predicates).
func HasGroupWith(preds ...predicate.Group) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GroupInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, GroupTable, GroupColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.GroupTag) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.GroupTag) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
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
func Not(p predicate.GroupTag) predicate.GroupTag {
	return predicate.GroupTag(func(s *sql.Selector) {
		p(s.Not())
	})
}