// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package pet

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.EQ(v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EQ(v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.NEQ(v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...float64) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Within(v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...float64) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Without(v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GT(v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GTE(v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LT(v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v float64) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LTE(v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.NEQ(v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Within(v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Without(v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GT(v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GTE(v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LT(v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LTE(v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Containing(v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.StartingWith(v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EndingWith(v))
	})
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.EQ(v))
	})
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.NEQ(v))
	})
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...uuid.UUID) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.Within(v...))
	})
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...uuid.UUID) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.Without(v...))
	})
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.GT(v))
	})
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.GTE(v))
	})
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.LT(v))
	})
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v uuid.UUID) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldUUID, p.LTE(v))
	})
}

// UUIDIsNil applies the IsNil predicate on the "uuid" field.
func UUIDIsNil() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldUUID)
	})
}

// UUIDNotNil applies the NotNil predicate on the "uuid" field.
func UUIDNotNil() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldUUID)
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EQ(v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.NEQ(v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Within(v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.Pet {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Without(v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.GT(v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.GTE(v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.LT(v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.LTE(v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Containing(v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.StartingWith(v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EndingWith(v))
	})
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNickname)
	})
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNickname)
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.InE(TeamInverseLabel).InV()
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(TeamInverseLabel).Where(tr).InV()
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		t.InE(OwnerInverseLabel).InV()
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Pet {
	return predicate.Pet(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(OwnerInverseLabel).Where(tr).InV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(tr *dsl.Traversal) {
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
func Or(predicates ...predicate.Pet) predicate.Pet {
	return predicate.Pet(func(tr *dsl.Traversal) {
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
func Not(p predicate.Pet) predicate.Pet {
	return predicate.Pet(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
