// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package card

import (
	"time"

	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.EQ(v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.EQ(v))
	})
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.EQ(v))
	})
}

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.EQ(v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.EQ(v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.NEQ(v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.Within(vs...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.Without(vs...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.GT(v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.GTE(v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.LT(v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.LTE(v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.EQ(v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.NEQ(v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.Within(vs...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.Without(vs...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.GT(v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.GTE(v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.LT(v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.LTE(v))
	})
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.EQ(v))
	})
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.NEQ(v))
	})
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.Within(vs...))
	})
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.Without(vs...))
	})
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.GT(v))
	})
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.GTE(v))
	})
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.LT(v))
	})
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldBalance, p.LTE(v))
	})
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.EQ(v))
	})
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.NEQ(v))
	})
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.Within(vs...))
	})
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.Without(vs...))
	})
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.GT(v))
	})
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.GTE(v))
	})
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.LT(v))
	})
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.LTE(v))
	})
}

// NumberEqualFold applies the EqualFold predicate on the "number" field.
func NumberEqualFold(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.EqualFold(v))
	})
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.Containing(v))
	})
}

// NumberContainsFold applies the ContainsFold predicate on the "number" field.
func NumberContainsFold(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.ContainsFold(v))
	})
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.StartingWith(v))
	})
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldNumber, p.EndingWith(v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.NEQ(v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Within(vs...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Without(vs...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GT(v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GTE(v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LT(v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LTE(v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EqualFold(v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Containing(v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.ContainsFold(v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.StartingWith(v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EndingWith(v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldName)
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldName)
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.InE(OwnerInverseLabel).InV()
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(OwnerInverseLabel).Where(tr).InV()
	})
}

// HasSpec applies the HasEdge predicate on the "spec" edge.
func HasSpec() predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		t.InE(SpecInverseLabel).InV()
	})
}

// HasSpecWith applies the HasEdge predicate on the "spec" edge with a given conditions (other predicates).
func HasSpecWith(preds ...predicate.Spec) predicate.Card {
	return predicate.Card(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(SpecInverseLabel).Where(tr).InV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.Card(func(tr *dsl.Traversal) {
		trs := make([]any, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Card) predicate.Card {
	return predicate.Card(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
