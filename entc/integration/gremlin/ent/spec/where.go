// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package spec

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"

	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		t.OutE(CardLabel).OutV()
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Card) predicate.Spec {
	return predicate.Spec(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(CardLabel).Where(tr).OutV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Spec) predicate.Spec {
	return predicate.Spec(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Spec) predicate.Spec {
	return predicate.Spec(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Spec) predicate.Spec {
	return predicate.Spec(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
