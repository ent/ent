// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package license

import (
	time "time"

	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	p "entgo.io/ent/dialect/gremlin/graph/dsl/p"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.EQ(v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.EQ(v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.EQ(v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.NEQ(v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.Within(vs...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.Without(vs...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.GT(v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.GTE(v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.LT(v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldCreateTime, p.LTE(v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.EQ(v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.NEQ(v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.Within(vs...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.Without(vs...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.GT(v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.GTE(v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.LT(v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.License {
	return predicate.License(func(t *dsl.Traversal) {
		t.Has(Label, FieldUpdateTime, p.LTE(v))
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.License) predicate.License {
	return predicate.License(func(tr *dsl.Traversal) {
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
func Or(predicates ...predicate.License) predicate.License {
	return predicate.License(func(tr *dsl.Traversal) {
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
func Not(p predicate.License) predicate.License {
	return predicate.License(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
