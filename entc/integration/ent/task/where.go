// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package enttask

import (
	time "time"

	"entgo.io/ent/dialect/sql"
	predicate "entgo.io/ent/entc/integration/ent/predicate"
	task "entgo.io/ent/entc/integration/ent/schema/task"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldID, id))
}

// Priority applies equality check predicate on the "priority" field. It's identical to PriorityEQ.
func Priority(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldEQ(FieldPriority, vc))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// PriorityEQ applies the EQ predicate on the "priority" field.
func PriorityEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldEQ(FieldPriority, vc))
}

// PriorityNEQ applies the NEQ predicate on the "priority" field.
func PriorityNEQ(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldNEQ(FieldPriority, vc))
}

// PriorityIn applies the In predicate on the "priority" field.
func PriorityIn(vs ...task.Priority) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(sql.FieldIn(FieldPriority, v...))
}

// PriorityNotIn applies the NotIn predicate on the "priority" field.
func PriorityNotIn(vs ...task.Priority) predicate.Task {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = int(vs[i])
	}
	return predicate.Task(sql.FieldNotIn(FieldPriority, v...))
}

// PriorityGT applies the GT predicate on the "priority" field.
func PriorityGT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldGT(FieldPriority, vc))
}

// PriorityGTE applies the GTE predicate on the "priority" field.
func PriorityGTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldGTE(FieldPriority, vc))
}

// PriorityLT applies the LT predicate on the "priority" field.
func PriorityLT(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldLT(FieldPriority, vc))
}

// PriorityLTE applies the LTE predicate on the "priority" field.
func PriorityLTE(v task.Priority) predicate.Task {
	vc := int(v)
	return predicate.Task(sql.FieldLTE(FieldPriority, vc))
}

// PrioritiesIsNil applies the IsNil predicate on the "priorities" field.
func PrioritiesIsNil() predicate.Task {
	return predicate.Task(sql.FieldIsNull(FieldPriorities))
}

// PrioritiesNotNil applies the NotNil predicate on the "priorities" field.
func PrioritiesNotNil() predicate.Task {
	return predicate.Task(sql.FieldNotNull(FieldPriorities))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Task {
	return predicate.Task(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Task {
	return predicate.Task(sql.FieldLTE(FieldCreatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
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
func Not(p predicate.Task) predicate.Task {
	return predicate.Task(func(s *sql.Selector) {
		p(s.Not())
	})
}
