// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package mixinid

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.MixinID {
	return predicate.MixinID(sql.FieldLTE(FieldID, id))
}

// SomeField applies equality check predicate on the "some_field" field. It's identical to SomeFieldEQ.
func SomeField(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldSomeField, v))
}

// MixinField applies equality check predicate on the "mixin_field" field. It's identical to MixinFieldEQ.
func MixinField(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldMixinField, v))
}

// SomeFieldEQ applies the EQ predicate on the "some_field" field.
func SomeFieldEQ(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldSomeField, v))
}

// SomeFieldNEQ applies the NEQ predicate on the "some_field" field.
func SomeFieldNEQ(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldNEQ(FieldSomeField, v))
}

// SomeFieldIn applies the In predicate on the "some_field" field.
func SomeFieldIn(vs ...string) predicate.MixinID {
	return predicate.MixinID(sql.FieldIn(FieldSomeField, vs...))
}

// SomeFieldNotIn applies the NotIn predicate on the "some_field" field.
func SomeFieldNotIn(vs ...string) predicate.MixinID {
	return predicate.MixinID(sql.FieldNotIn(FieldSomeField, vs...))
}

// SomeFieldGT applies the GT predicate on the "some_field" field.
func SomeFieldGT(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldGT(FieldSomeField, v))
}

// SomeFieldGTE applies the GTE predicate on the "some_field" field.
func SomeFieldGTE(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldGTE(FieldSomeField, v))
}

// SomeFieldLT applies the LT predicate on the "some_field" field.
func SomeFieldLT(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldLT(FieldSomeField, v))
}

// SomeFieldLTE applies the LTE predicate on the "some_field" field.
func SomeFieldLTE(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldLTE(FieldSomeField, v))
}

// SomeFieldEqualFold applies the EqualFold predicate on the "some_field" field.
func SomeFieldEqualFold(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEqualFold(FieldSomeField, v))
}

// SomeFieldContains applies the Contains predicate on the "some_field" field.
func SomeFieldContains(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldContains(FieldSomeField, v))
}

// SomeFieldContainsFold applies the ContainsFold predicate on the "some_field" field.
func SomeFieldContainsFold(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldContainsFold(FieldSomeField, v))
}

// SomeFieldHasPrefix applies the HasPrefix predicate on the "some_field" field.
func SomeFieldHasPrefix(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldHasPrefix(FieldSomeField, v))
}

// SomeFieldHasSuffix applies the HasSuffix predicate on the "some_field" field.
func SomeFieldHasSuffix(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldHasSuffix(FieldSomeField, v))
}

// MixinFieldEQ applies the EQ predicate on the "mixin_field" field.
func MixinFieldEQ(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEQ(FieldMixinField, v))
}

// MixinFieldNEQ applies the NEQ predicate on the "mixin_field" field.
func MixinFieldNEQ(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldNEQ(FieldMixinField, v))
}

// MixinFieldIn applies the In predicate on the "mixin_field" field.
func MixinFieldIn(vs ...string) predicate.MixinID {
	return predicate.MixinID(sql.FieldIn(FieldMixinField, vs...))
}

// MixinFieldNotIn applies the NotIn predicate on the "mixin_field" field.
func MixinFieldNotIn(vs ...string) predicate.MixinID {
	return predicate.MixinID(sql.FieldNotIn(FieldMixinField, vs...))
}

// MixinFieldGT applies the GT predicate on the "mixin_field" field.
func MixinFieldGT(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldGT(FieldMixinField, v))
}

// MixinFieldGTE applies the GTE predicate on the "mixin_field" field.
func MixinFieldGTE(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldGTE(FieldMixinField, v))
}

// MixinFieldLT applies the LT predicate on the "mixin_field" field.
func MixinFieldLT(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldLT(FieldMixinField, v))
}

// MixinFieldLTE applies the LTE predicate on the "mixin_field" field.
func MixinFieldLTE(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldLTE(FieldMixinField, v))
}

// MixinFieldEqualFold applies the EqualFold predicate on the "mixin_field" field.
func MixinFieldEqualFold(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldEqualFold(FieldMixinField, v))
}

// MixinFieldContains applies the Contains predicate on the "mixin_field" field.
func MixinFieldContains(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldContains(FieldMixinField, v))
}

// MixinFieldContainsFold applies the ContainsFold predicate on the "mixin_field" field.
func MixinFieldContainsFold(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldContainsFold(FieldMixinField, v))
}

// MixinFieldHasPrefix applies the HasPrefix predicate on the "mixin_field" field.
func MixinFieldHasPrefix(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldHasPrefix(FieldMixinField, v))
}

// MixinFieldHasSuffix applies the HasSuffix predicate on the "mixin_field" field.
func MixinFieldHasSuffix(v string) predicate.MixinID {
	return predicate.MixinID(sql.FieldHasSuffix(FieldMixinField, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.MixinID) predicate.MixinID {
	return predicate.MixinID(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.MixinID) predicate.MixinID {
	return predicate.MixinID(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.MixinID) predicate.MixinID {
	return predicate.MixinID(sql.NotPredicates(p))
}
