// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package note

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
)

// ID filters vertices based on their ID field.
func ID(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id schema.NoteID) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldID, id))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldText, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Note {
	return predicate.Note(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Note {
	return predicate.Note(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Note {
	return predicate.Note(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Note {
	return predicate.Note(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Note {
	return predicate.Note(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Note {
	return predicate.Note(sql.FieldLTE(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Note {
	return predicate.Note(sql.FieldEqualFold(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Note {
	return predicate.Note(sql.FieldContains(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Note {
	return predicate.Note(sql.FieldContainsFold(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Note {
	return predicate.Note(sql.FieldHasSuffix(FieldText, v))
}

// TextIsNil applies the IsNil predicate on the "text" field.
func TextIsNil() predicate.Note {
	return predicate.Note(sql.FieldIsNull(FieldText))
}

// TextNotNil applies the NotNil predicate on the "text" field.
func TextNotNil() predicate.Note {
	return predicate.Note(sql.FieldNotNull(FieldText))
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.Note) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.Note) predicate.Note {
	return predicate.Note(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Note) predicate.Note {
	return predicate.Note(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Note) predicate.Note {
	return predicate.Note(sql.NotPredicates(p))
}
