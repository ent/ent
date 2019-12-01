// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(id)
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.EQ(id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
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
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Within(v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
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
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Without(v...))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GT(id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GTE(id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LT(id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	)
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.EQ(v))
		},
	)
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.EQ(v))
		},
	)
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.EQ(v))
		},
	)
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.NEQ(v))
		},
	)
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldCreateTime), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.Within(v...))
		},
	)
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.Without(v...))
		},
	)
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.GT(v))
		},
	)
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.GTE(v))
		},
	)
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.LT(v))
		},
	)
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldCreateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldCreateTime, p.LTE(v))
		},
	)
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.EQ(v))
		},
	)
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.NEQ(v))
		},
	)
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldUpdateTime), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.Within(v...))
		},
	)
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.Without(v...))
		},
	)
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.GT(v))
		},
	)
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.GTE(v))
		},
	)
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.LT(v))
		},
	)
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldUpdateTime), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldUpdateTime, p.LTE(v))
		},
	)
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	)
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.NEQ(v))
		},
	)
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNumber), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Within(v...))
		},
	)
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNumber), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Without(v...))
		},
	)
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GT(v))
		},
	)
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GTE(v))
		},
	)
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LT(v))
		},
	)
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LTE(v))
		},
	)
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Containing(v))
		},
	)
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.StartingWith(v))
		},
	)
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EndingWith(v))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Within(v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Without(v...))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	)
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldName)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldName)
		},
	)
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldName)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldName)
		},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(OwnerTable, FieldID),
				sql.Edge(sql.O2O, true, OwnerTable, OwnerColumn),
			)
			sql.HasNeighbors(s, step)
		},
		func(t *dsl.Traversal) {
			t.InE(OwnerInverseLabel).InV()
		},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			step := sql.NewStep(
				sql.From(Table, FieldID),
				sql.To(OwnerInverseTable, FieldID),
				sql.Edge(sql.O2O, true, OwnerTable, OwnerColumn),
			)
			sql.HasNeighborsWith(s, step, func(s *sql.Selector) {
				for _, p := range preds {
					p(s)
				}
			})
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(OwnerInverseLabel).Where(tr).InV()
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.And(trs...))
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
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
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.Or(trs...))
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			p(s.Not())
		},
		func(tr *dsl.Traversal) {
			t := __.New()
			p(t)
			tr.Where(__.Not(t))
		},
	)
}
