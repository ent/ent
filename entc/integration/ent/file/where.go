// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// SetID applies equality check predicate on the "set_id" field. It's identical to SetIDEQ.
func SetID(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSetID, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// User applies equality check predicate on the "user" field. It's identical to UserEQ.
func User(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUser, v))
}

// Group applies equality check predicate on the "group" field. It's identical to GroupEQ.
func Group(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldGroup, v))
}

// Op applies equality check predicate on the "op" field. It's identical to OpEQ.
func Op(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldOp, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreateTime, v))
}

// SetIDEQ applies the EQ predicate on the "set_id" field.
func SetIDEQ(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSetID, v))
}

// SetIDNEQ applies the NEQ predicate on the "set_id" field.
func SetIDNEQ(v int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSetID, v))
}

// SetIDIn applies the In predicate on the "set_id" field.
func SetIDIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldSetID, vs...))
}

// SetIDNotIn applies the NotIn predicate on the "set_id" field.
func SetIDNotIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSetID, vs...))
}

// SetIDGT applies the GT predicate on the "set_id" field.
func SetIDGT(v int) predicate.File {
	return predicate.File(sql.FieldGT(FieldSetID, v))
}

// SetIDGTE applies the GTE predicate on the "set_id" field.
func SetIDGTE(v int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSetID, v))
}

// SetIDLT applies the LT predicate on the "set_id" field.
func SetIDLT(v int) predicate.File {
	return predicate.File(sql.FieldLT(FieldSetID, v))
}

// SetIDLTE applies the LTE predicate on the "set_id" field.
func SetIDLTE(v int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSetID, v))
}

// SetIDIsNil applies the IsNil predicate on the "set_id" field.
func SetIDIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldSetID))
}

// SetIDNotNil applies the NotNil predicate on the "set_id" field.
func SetIDNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldSetID))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.File {
	return predicate.File(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.File {
	return predicate.File(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSize, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldName, v))
}

// NameRegex applies the Regex predicate on the "name" field.
func NameRegex(v string) predicate.File {
	return predicate.File(sql.FieldRegex(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldName, v))
}

// UserEQ applies the EQ predicate on the "user" field.
func UserEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUser, v))
}

// UserNEQ applies the NEQ predicate on the "user" field.
func UserNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUser, v))
}

// UserIn applies the In predicate on the "user" field.
func UserIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldUser, vs...))
}

// UserNotIn applies the NotIn predicate on the "user" field.
func UserNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUser, vs...))
}

// UserGT applies the GT predicate on the "user" field.
func UserGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldUser, v))
}

// UserGTE applies the GTE predicate on the "user" field.
func UserGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUser, v))
}

// UserLT applies the LT predicate on the "user" field.
func UserLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldUser, v))
}

// UserLTE applies the LTE predicate on the "user" field.
func UserLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUser, v))
}

// UserContains applies the Contains predicate on the "user" field.
func UserContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldUser, v))
}

// UserHasPrefix applies the HasPrefix predicate on the "user" field.
func UserHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldUser, v))
}

// UserHasSuffix applies the HasSuffix predicate on the "user" field.
func UserHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldUser, v))
}

// UserRegex applies the Regex predicate on the "user" field.
func UserRegex(v string) predicate.File {
	return predicate.File(sql.FieldRegex(FieldUser, v))
}

// UserIsNil applies the IsNil predicate on the "user" field.
func UserIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldUser))
}

// UserNotNil applies the NotNil predicate on the "user" field.
func UserNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldUser))
}

// UserEqualFold applies the EqualFold predicate on the "user" field.
func UserEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldUser, v))
}

// UserContainsFold applies the ContainsFold predicate on the "user" field.
func UserContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldUser, v))
}

// GroupEQ applies the EQ predicate on the "group" field.
func GroupEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldGroup, v))
}

// GroupNEQ applies the NEQ predicate on the "group" field.
func GroupNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldGroup, v))
}

// GroupIn applies the In predicate on the "group" field.
func GroupIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldGroup, vs...))
}

// GroupNotIn applies the NotIn predicate on the "group" field.
func GroupNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldGroup, vs...))
}

// GroupGT applies the GT predicate on the "group" field.
func GroupGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldGroup, v))
}

// GroupGTE applies the GTE predicate on the "group" field.
func GroupGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldGroup, v))
}

// GroupLT applies the LT predicate on the "group" field.
func GroupLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldGroup, v))
}

// GroupLTE applies the LTE predicate on the "group" field.
func GroupLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldGroup, v))
}

// GroupContains applies the Contains predicate on the "group" field.
func GroupContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldGroup, v))
}

// GroupHasPrefix applies the HasPrefix predicate on the "group" field.
func GroupHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldGroup, v))
}

// GroupHasSuffix applies the HasSuffix predicate on the "group" field.
func GroupHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldGroup, v))
}

// GroupRegex applies the Regex predicate on the "group" field.
func GroupRegex(v string) predicate.File {
	return predicate.File(sql.FieldRegex(FieldGroup, v))
}

// GroupIsNil applies the IsNil predicate on the "group" field.
func GroupIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldGroup))
}

// GroupNotNil applies the NotNil predicate on the "group" field.
func GroupNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldGroup))
}

// GroupEqualFold applies the EqualFold predicate on the "group" field.
func GroupEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldGroup, v))
}

// GroupContainsFold applies the ContainsFold predicate on the "group" field.
func GroupContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldGroup, v))
}

// OpEQ applies the EQ predicate on the "op" field.
func OpEQ(v bool) predicate.File {
	return predicate.File(sql.FieldEQ(FieldOp, v))
}

// OpNEQ applies the NEQ predicate on the "op" field.
func OpNEQ(v bool) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldOp, v))
}

// OpIsNil applies the IsNil predicate on the "op" field.
func OpIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldOp))
}

// OpNotNil applies the NotNil predicate on the "op" field.
func OpNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldOp))
}

// FieldIDEQ applies the EQ predicate on the "field_id" field.
func FieldIDEQ(v int) predicate.File {
	return predicate.File(sql.FieldEQ(FieldFieldID, v))
}

// FieldIDNEQ applies the NEQ predicate on the "field_id" field.
func FieldIDNEQ(v int) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldFieldID, v))
}

// FieldIDIn applies the In predicate on the "field_id" field.
func FieldIDIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldIn(FieldFieldID, vs...))
}

// FieldIDNotIn applies the NotIn predicate on the "field_id" field.
func FieldIDNotIn(vs ...int) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldFieldID, vs...))
}

// FieldIDGT applies the GT predicate on the "field_id" field.
func FieldIDGT(v int) predicate.File {
	return predicate.File(sql.FieldGT(FieldFieldID, v))
}

// FieldIDGTE applies the GTE predicate on the "field_id" field.
func FieldIDGTE(v int) predicate.File {
	return predicate.File(sql.FieldGTE(FieldFieldID, v))
}

// FieldIDLT applies the LT predicate on the "field_id" field.
func FieldIDLT(v int) predicate.File {
	return predicate.File(sql.FieldLT(FieldFieldID, v))
}

// FieldIDLTE applies the LTE predicate on the "field_id" field.
func FieldIDLTE(v int) predicate.File {
	return predicate.File(sql.FieldLTE(FieldFieldID, v))
}

// FieldIDIsNil applies the IsNil predicate on the "field_id" field.
func FieldIDIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldFieldID))
}

// FieldIDNotNil applies the NotNil predicate on the "field_id" field.
func FieldIDNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldFieldID))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreateTime, v))
}

// CreateTimeIsNil applies the IsNil predicate on the "create_time" field.
func CreateTimeIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldCreateTime))
}

// CreateTimeNotNil applies the NotNil predicate on the "create_time" field.
func CreateTimeNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldCreateTime))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasType applies the HasEdge predicate on the "type" edge.
func HasType() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, TypeTable, TypeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTypeWith applies the HasEdge predicate on the "type" edge with a given conditions (other predicates).
func HasTypeWith(preds ...predicate.FileType) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newTypeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasField applies the HasEdge predicate on the "field" edge.
func HasField() predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FieldTable, FieldColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFieldWith applies the HasEdge predicate on the "field" edge with a given conditions (other predicates).
func HasFieldWith(preds ...predicate.FieldType) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		step := newFieldStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(sql.NotPredicates(p))
}
