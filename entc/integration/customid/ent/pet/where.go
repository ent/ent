// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package pet

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
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
func IDNotIn(ids ...string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
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
func IDGT(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCars applies the HasEdge predicate on the "cars" edge.
func HasCars() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CarsTable, CarsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCarsWith applies the HasEdge predicate on the "cars" edge with a given conditions (other predicates).
func HasCarsWith(preds ...predicate.Car) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CarsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, CarsTable, CarsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriends applies the HasEdge predicate on the "friends" edge.
func HasFriends() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(FriendsTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendsWith applies the HasEdge predicate on the "friends" edge with a given conditions (other predicates).
func HasFriendsWith(preds ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBestFriend applies the HasEdge predicate on the "best_friend" edge.
func HasBestFriend() predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BestFriendTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BestFriendTable, BestFriendColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBestFriendWith applies the HasEdge predicate on the "best_friend" edge with a given conditions (other predicates).
func HasBestFriendWith(preds ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BestFriendTable, BestFriendColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
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
func Not(p predicate.Pet) predicate.Pet {
	return predicate.Pet(func(s *sql.Selector) {
		p(s.Not())
	})
}
