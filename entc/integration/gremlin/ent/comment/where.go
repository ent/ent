// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package comment

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// UniqueInt applies equality check predicate on the "unique_int" field. It's identical to UniqueIntEQ.
func UniqueInt(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.EQ(v))
	})
}

// UniqueFloat applies equality check predicate on the "unique_float" field. It's identical to UniqueFloatEQ.
func UniqueFloat(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.EQ(v))
	})
}

// NillableInt applies equality check predicate on the "nillable_int" field. It's identical to NillableIntEQ.
func NillableInt(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.EQ(v))
	})
}

// Client applies equality check predicate on the "client" field. It's identical to ClientEQ.
func Client(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.EQ(v))
	})
}

// UniqueIntEQ applies the EQ predicate on the "unique_int" field.
func UniqueIntEQ(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.EQ(v))
	})
}

// UniqueIntNEQ applies the NEQ predicate on the "unique_int" field.
func UniqueIntNEQ(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.NEQ(v))
	})
}

// UniqueIntIn applies the In predicate on the "unique_int" field.
func UniqueIntIn(vs ...int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.Within(vs...))
	})
}

// UniqueIntNotIn applies the NotIn predicate on the "unique_int" field.
func UniqueIntNotIn(vs ...int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.Without(vs...))
	})
}

// UniqueIntGT applies the GT predicate on the "unique_int" field.
func UniqueIntGT(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.GT(v))
	})
}

// UniqueIntGTE applies the GTE predicate on the "unique_int" field.
func UniqueIntGTE(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.GTE(v))
	})
}

// UniqueIntLT applies the LT predicate on the "unique_int" field.
func UniqueIntLT(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.LT(v))
	})
}

// UniqueIntLTE applies the LTE predicate on the "unique_int" field.
func UniqueIntLTE(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueInt, p.LTE(v))
	})
}

// UniqueFloatEQ applies the EQ predicate on the "unique_float" field.
func UniqueFloatEQ(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.EQ(v))
	})
}

// UniqueFloatNEQ applies the NEQ predicate on the "unique_float" field.
func UniqueFloatNEQ(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.NEQ(v))
	})
}

// UniqueFloatIn applies the In predicate on the "unique_float" field.
func UniqueFloatIn(vs ...float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.Within(vs...))
	})
}

// UniqueFloatNotIn applies the NotIn predicate on the "unique_float" field.
func UniqueFloatNotIn(vs ...float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.Without(vs...))
	})
}

// UniqueFloatGT applies the GT predicate on the "unique_float" field.
func UniqueFloatGT(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.GT(v))
	})
}

// UniqueFloatGTE applies the GTE predicate on the "unique_float" field.
func UniqueFloatGTE(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.GTE(v))
	})
}

// UniqueFloatLT applies the LT predicate on the "unique_float" field.
func UniqueFloatLT(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.LT(v))
	})
}

// UniqueFloatLTE applies the LTE predicate on the "unique_float" field.
func UniqueFloatLTE(v float64) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldUniqueFloat, p.LTE(v))
	})
}

// NillableIntEQ applies the EQ predicate on the "nillable_int" field.
func NillableIntEQ(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.EQ(v))
	})
}

// NillableIntNEQ applies the NEQ predicate on the "nillable_int" field.
func NillableIntNEQ(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.NEQ(v))
	})
}

// NillableIntIn applies the In predicate on the "nillable_int" field.
func NillableIntIn(vs ...int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.Within(vs...))
	})
}

// NillableIntNotIn applies the NotIn predicate on the "nillable_int" field.
func NillableIntNotIn(vs ...int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.Without(vs...))
	})
}

// NillableIntGT applies the GT predicate on the "nillable_int" field.
func NillableIntGT(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.GT(v))
	})
}

// NillableIntGTE applies the GTE predicate on the "nillable_int" field.
func NillableIntGTE(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.GTE(v))
	})
}

// NillableIntLT applies the LT predicate on the "nillable_int" field.
func NillableIntLT(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.LT(v))
	})
}

// NillableIntLTE applies the LTE predicate on the "nillable_int" field.
func NillableIntLTE(v int) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.LTE(v))
	})
}

// NillableIntIsNil applies the IsNil predicate on the "nillable_int" field.
func NillableIntIsNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt)
	})
}

// NillableIntNotNil applies the NotNil predicate on the "nillable_int" field.
func NillableIntNotNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt)
	})
}

// TableEQ applies the EQ predicate on the "table" field.
func TableEQ(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.EQ(v))
	})
}

// TableNEQ applies the NEQ predicate on the "table" field.
func TableNEQ(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.NEQ(v))
	})
}

// TableIn applies the In predicate on the "table" field.
func TableIn(vs ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.Within(vs...))
	})
}

// TableNotIn applies the NotIn predicate on the "table" field.
func TableNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.Without(vs...))
	})
}

// TableGT applies the GT predicate on the "table" field.
func TableGT(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.GT(v))
	})
}

// TableGTE applies the GTE predicate on the "table" field.
func TableGTE(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.GTE(v))
	})
}

// TableLT applies the LT predicate on the "table" field.
func TableLT(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.LT(v))
	})
}

// TableLTE applies the LTE predicate on the "table" field.
func TableLTE(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.LTE(v))
	})
}

// TableEqualFold applies the EqualFold predicate on the "table" field.
func TableEqualFold(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.EqualFold(v))
	})
}

// TableContains applies the Contains predicate on the "table" field.
func TableContains(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.Containing(v))
	})
}

// TableContainsFold applies the ContainsFold predicate on the "table" field.
func TableContainsFold(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.ContainsFold(v))
	})
}

// TableHasPrefix applies the HasPrefix predicate on the "table" field.
func TableHasPrefix(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.StartingWith(v))
	})
}

// TableHasSuffix applies the HasSuffix predicate on the "table" field.
func TableHasSuffix(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldTable, p.EndingWith(v))
	})
}

// TableIsNil applies the IsNil predicate on the "table" field.
func TableIsNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldTable)
	})
}

// TableNotNil applies the NotNil predicate on the "table" field.
func TableNotNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldTable)
	})
}

// DirIsNil applies the IsNil predicate on the "dir" field.
func DirIsNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldDir)
	})
}

// DirNotNil applies the NotNil predicate on the "dir" field.
func DirNotNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldDir)
	})
}

// ClientEQ applies the EQ predicate on the "client" field.
func ClientEQ(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.EQ(v))
	})
}

// ClientNEQ applies the NEQ predicate on the "client" field.
func ClientNEQ(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.NEQ(v))
	})
}

// ClientIn applies the In predicate on the "client" field.
func ClientIn(vs ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.Within(vs...))
	})
}

// ClientNotIn applies the NotIn predicate on the "client" field.
func ClientNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.Without(vs...))
	})
}

// ClientGT applies the GT predicate on the "client" field.
func ClientGT(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.GT(v))
	})
}

// ClientGTE applies the GTE predicate on the "client" field.
func ClientGTE(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.GTE(v))
	})
}

// ClientLT applies the LT predicate on the "client" field.
func ClientLT(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.LT(v))
	})
}

// ClientLTE applies the LTE predicate on the "client" field.
func ClientLTE(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.LTE(v))
	})
}

// ClientEqualFold applies the EqualFold predicate on the "client" field.
func ClientEqualFold(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.EqualFold(v))
	})
}

// ClientContains applies the Contains predicate on the "client" field.
func ClientContains(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.Containing(v))
	})
}

// ClientContainsFold applies the ContainsFold predicate on the "client" field.
func ClientContainsFold(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.ContainsFold(v))
	})
}

// ClientHasPrefix applies the HasPrefix predicate on the "client" field.
func ClientHasPrefix(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.StartingWith(v))
	})
}

// ClientHasSuffix applies the HasSuffix predicate on the "client" field.
func ClientHasSuffix(v string) predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.Has(Label, FieldClient, p.EndingWith(v))
	})
}

// ClientIsNil applies the IsNil predicate on the "client" field.
func ClientIsNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldClient)
	})
}

// ClientNotNil applies the NotNil predicate on the "client" field.
func ClientNotNil() predicate.Comment {
	return predicate.Comment(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldClient)
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(tr *dsl.Traversal) {
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
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(func(tr *dsl.Traversal) {
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
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
