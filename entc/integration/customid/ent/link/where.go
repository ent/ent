// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package link

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	uuidc "entgo.io/ent/entc/integration/customid/uuidcompatible"
)

// ID filters vertices based on their ID field.
func ID(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuidc.UUIDC) predicate.Link {
	return predicate.Link(sql.FieldLTE(FieldID, id))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Link) predicate.Link {
	return predicate.Link(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Link) predicate.Link {
	return predicate.Link(sql.NotPredicates(p))
}
