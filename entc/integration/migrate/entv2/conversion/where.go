// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package conversion

import (
	"entgo.io/ent/dialect/sql"
	predicate "entgo.io/ent/entc/integration/migrate/entv2/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldName, v))
}

// Int8ToString applies equality check predicate on the "int8_to_string" field. It's identical to Int8ToStringEQ.
func Int8ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt8ToString, v))
}

// Uint8ToString applies equality check predicate on the "uint8_to_string" field. It's identical to Uint8ToStringEQ.
func Uint8ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint8ToString, v))
}

// Int16ToString applies equality check predicate on the "int16_to_string" field. It's identical to Int16ToStringEQ.
func Int16ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt16ToString, v))
}

// Uint16ToString applies equality check predicate on the "uint16_to_string" field. It's identical to Uint16ToStringEQ.
func Uint16ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint16ToString, v))
}

// Int32ToString applies equality check predicate on the "int32_to_string" field. It's identical to Int32ToStringEQ.
func Int32ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt32ToString, v))
}

// Uint32ToString applies equality check predicate on the "uint32_to_string" field. It's identical to Uint32ToStringEQ.
func Uint32ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint32ToString, v))
}

// Int64ToString applies equality check predicate on the "int64_to_string" field. It's identical to Int64ToStringEQ.
func Int64ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt64ToString, v))
}

// Uint64ToString applies equality check predicate on the "uint64_to_string" field. It's identical to Uint64ToStringEQ.
func Uint64ToString(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint64ToString, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldName, v))
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldName))
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldName))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldName, v))
}

// Int8ToStringEQ applies the EQ predicate on the "int8_to_string" field.
func Int8ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt8ToString, v))
}

// Int8ToStringNEQ applies the NEQ predicate on the "int8_to_string" field.
func Int8ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldInt8ToString, v))
}

// Int8ToStringIn applies the In predicate on the "int8_to_string" field.
func Int8ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldInt8ToString, vs...))
}

// Int8ToStringNotIn applies the NotIn predicate on the "int8_to_string" field.
func Int8ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldInt8ToString, vs...))
}

// Int8ToStringGT applies the GT predicate on the "int8_to_string" field.
func Int8ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldInt8ToString, v))
}

// Int8ToStringGTE applies the GTE predicate on the "int8_to_string" field.
func Int8ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldInt8ToString, v))
}

// Int8ToStringLT applies the LT predicate on the "int8_to_string" field.
func Int8ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldInt8ToString, v))
}

// Int8ToStringLTE applies the LTE predicate on the "int8_to_string" field.
func Int8ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldInt8ToString, v))
}

// Int8ToStringContains applies the Contains predicate on the "int8_to_string" field.
func Int8ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldInt8ToString, v))
}

// Int8ToStringHasPrefix applies the HasPrefix predicate on the "int8_to_string" field.
func Int8ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldInt8ToString, v))
}

// Int8ToStringHasSuffix applies the HasSuffix predicate on the "int8_to_string" field.
func Int8ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldInt8ToString, v))
}

// Int8ToStringIsNil applies the IsNil predicate on the "int8_to_string" field.
func Int8ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldInt8ToString))
}

// Int8ToStringNotNil applies the NotNil predicate on the "int8_to_string" field.
func Int8ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldInt8ToString))
}

// Int8ToStringEqualFold applies the EqualFold predicate on the "int8_to_string" field.
func Int8ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldInt8ToString, v))
}

// Int8ToStringContainsFold applies the ContainsFold predicate on the "int8_to_string" field.
func Int8ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldInt8ToString, v))
}

// Uint8ToStringEQ applies the EQ predicate on the "uint8_to_string" field.
func Uint8ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint8ToString, v))
}

// Uint8ToStringNEQ applies the NEQ predicate on the "uint8_to_string" field.
func Uint8ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldUint8ToString, v))
}

// Uint8ToStringIn applies the In predicate on the "uint8_to_string" field.
func Uint8ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldUint8ToString, vs...))
}

// Uint8ToStringNotIn applies the NotIn predicate on the "uint8_to_string" field.
func Uint8ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldUint8ToString, vs...))
}

// Uint8ToStringGT applies the GT predicate on the "uint8_to_string" field.
func Uint8ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldUint8ToString, v))
}

// Uint8ToStringGTE applies the GTE predicate on the "uint8_to_string" field.
func Uint8ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldUint8ToString, v))
}

// Uint8ToStringLT applies the LT predicate on the "uint8_to_string" field.
func Uint8ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldUint8ToString, v))
}

// Uint8ToStringLTE applies the LTE predicate on the "uint8_to_string" field.
func Uint8ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldUint8ToString, v))
}

// Uint8ToStringContains applies the Contains predicate on the "uint8_to_string" field.
func Uint8ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldUint8ToString, v))
}

// Uint8ToStringHasPrefix applies the HasPrefix predicate on the "uint8_to_string" field.
func Uint8ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldUint8ToString, v))
}

// Uint8ToStringHasSuffix applies the HasSuffix predicate on the "uint8_to_string" field.
func Uint8ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldUint8ToString, v))
}

// Uint8ToStringIsNil applies the IsNil predicate on the "uint8_to_string" field.
func Uint8ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldUint8ToString))
}

// Uint8ToStringNotNil applies the NotNil predicate on the "uint8_to_string" field.
func Uint8ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldUint8ToString))
}

// Uint8ToStringEqualFold applies the EqualFold predicate on the "uint8_to_string" field.
func Uint8ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldUint8ToString, v))
}

// Uint8ToStringContainsFold applies the ContainsFold predicate on the "uint8_to_string" field.
func Uint8ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldUint8ToString, v))
}

// Int16ToStringEQ applies the EQ predicate on the "int16_to_string" field.
func Int16ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt16ToString, v))
}

// Int16ToStringNEQ applies the NEQ predicate on the "int16_to_string" field.
func Int16ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldInt16ToString, v))
}

// Int16ToStringIn applies the In predicate on the "int16_to_string" field.
func Int16ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldInt16ToString, vs...))
}

// Int16ToStringNotIn applies the NotIn predicate on the "int16_to_string" field.
func Int16ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldInt16ToString, vs...))
}

// Int16ToStringGT applies the GT predicate on the "int16_to_string" field.
func Int16ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldInt16ToString, v))
}

// Int16ToStringGTE applies the GTE predicate on the "int16_to_string" field.
func Int16ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldInt16ToString, v))
}

// Int16ToStringLT applies the LT predicate on the "int16_to_string" field.
func Int16ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldInt16ToString, v))
}

// Int16ToStringLTE applies the LTE predicate on the "int16_to_string" field.
func Int16ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldInt16ToString, v))
}

// Int16ToStringContains applies the Contains predicate on the "int16_to_string" field.
func Int16ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldInt16ToString, v))
}

// Int16ToStringHasPrefix applies the HasPrefix predicate on the "int16_to_string" field.
func Int16ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldInt16ToString, v))
}

// Int16ToStringHasSuffix applies the HasSuffix predicate on the "int16_to_string" field.
func Int16ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldInt16ToString, v))
}

// Int16ToStringIsNil applies the IsNil predicate on the "int16_to_string" field.
func Int16ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldInt16ToString))
}

// Int16ToStringNotNil applies the NotNil predicate on the "int16_to_string" field.
func Int16ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldInt16ToString))
}

// Int16ToStringEqualFold applies the EqualFold predicate on the "int16_to_string" field.
func Int16ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldInt16ToString, v))
}

// Int16ToStringContainsFold applies the ContainsFold predicate on the "int16_to_string" field.
func Int16ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldInt16ToString, v))
}

// Uint16ToStringEQ applies the EQ predicate on the "uint16_to_string" field.
func Uint16ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint16ToString, v))
}

// Uint16ToStringNEQ applies the NEQ predicate on the "uint16_to_string" field.
func Uint16ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldUint16ToString, v))
}

// Uint16ToStringIn applies the In predicate on the "uint16_to_string" field.
func Uint16ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldUint16ToString, vs...))
}

// Uint16ToStringNotIn applies the NotIn predicate on the "uint16_to_string" field.
func Uint16ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldUint16ToString, vs...))
}

// Uint16ToStringGT applies the GT predicate on the "uint16_to_string" field.
func Uint16ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldUint16ToString, v))
}

// Uint16ToStringGTE applies the GTE predicate on the "uint16_to_string" field.
func Uint16ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldUint16ToString, v))
}

// Uint16ToStringLT applies the LT predicate on the "uint16_to_string" field.
func Uint16ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldUint16ToString, v))
}

// Uint16ToStringLTE applies the LTE predicate on the "uint16_to_string" field.
func Uint16ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldUint16ToString, v))
}

// Uint16ToStringContains applies the Contains predicate on the "uint16_to_string" field.
func Uint16ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldUint16ToString, v))
}

// Uint16ToStringHasPrefix applies the HasPrefix predicate on the "uint16_to_string" field.
func Uint16ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldUint16ToString, v))
}

// Uint16ToStringHasSuffix applies the HasSuffix predicate on the "uint16_to_string" field.
func Uint16ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldUint16ToString, v))
}

// Uint16ToStringIsNil applies the IsNil predicate on the "uint16_to_string" field.
func Uint16ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldUint16ToString))
}

// Uint16ToStringNotNil applies the NotNil predicate on the "uint16_to_string" field.
func Uint16ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldUint16ToString))
}

// Uint16ToStringEqualFold applies the EqualFold predicate on the "uint16_to_string" field.
func Uint16ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldUint16ToString, v))
}

// Uint16ToStringContainsFold applies the ContainsFold predicate on the "uint16_to_string" field.
func Uint16ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldUint16ToString, v))
}

// Int32ToStringEQ applies the EQ predicate on the "int32_to_string" field.
func Int32ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt32ToString, v))
}

// Int32ToStringNEQ applies the NEQ predicate on the "int32_to_string" field.
func Int32ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldInt32ToString, v))
}

// Int32ToStringIn applies the In predicate on the "int32_to_string" field.
func Int32ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldInt32ToString, vs...))
}

// Int32ToStringNotIn applies the NotIn predicate on the "int32_to_string" field.
func Int32ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldInt32ToString, vs...))
}

// Int32ToStringGT applies the GT predicate on the "int32_to_string" field.
func Int32ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldInt32ToString, v))
}

// Int32ToStringGTE applies the GTE predicate on the "int32_to_string" field.
func Int32ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldInt32ToString, v))
}

// Int32ToStringLT applies the LT predicate on the "int32_to_string" field.
func Int32ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldInt32ToString, v))
}

// Int32ToStringLTE applies the LTE predicate on the "int32_to_string" field.
func Int32ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldInt32ToString, v))
}

// Int32ToStringContains applies the Contains predicate on the "int32_to_string" field.
func Int32ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldInt32ToString, v))
}

// Int32ToStringHasPrefix applies the HasPrefix predicate on the "int32_to_string" field.
func Int32ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldInt32ToString, v))
}

// Int32ToStringHasSuffix applies the HasSuffix predicate on the "int32_to_string" field.
func Int32ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldInt32ToString, v))
}

// Int32ToStringIsNil applies the IsNil predicate on the "int32_to_string" field.
func Int32ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldInt32ToString))
}

// Int32ToStringNotNil applies the NotNil predicate on the "int32_to_string" field.
func Int32ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldInt32ToString))
}

// Int32ToStringEqualFold applies the EqualFold predicate on the "int32_to_string" field.
func Int32ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldInt32ToString, v))
}

// Int32ToStringContainsFold applies the ContainsFold predicate on the "int32_to_string" field.
func Int32ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldInt32ToString, v))
}

// Uint32ToStringEQ applies the EQ predicate on the "uint32_to_string" field.
func Uint32ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint32ToString, v))
}

// Uint32ToStringNEQ applies the NEQ predicate on the "uint32_to_string" field.
func Uint32ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldUint32ToString, v))
}

// Uint32ToStringIn applies the In predicate on the "uint32_to_string" field.
func Uint32ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldUint32ToString, vs...))
}

// Uint32ToStringNotIn applies the NotIn predicate on the "uint32_to_string" field.
func Uint32ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldUint32ToString, vs...))
}

// Uint32ToStringGT applies the GT predicate on the "uint32_to_string" field.
func Uint32ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldUint32ToString, v))
}

// Uint32ToStringGTE applies the GTE predicate on the "uint32_to_string" field.
func Uint32ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldUint32ToString, v))
}

// Uint32ToStringLT applies the LT predicate on the "uint32_to_string" field.
func Uint32ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldUint32ToString, v))
}

// Uint32ToStringLTE applies the LTE predicate on the "uint32_to_string" field.
func Uint32ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldUint32ToString, v))
}

// Uint32ToStringContains applies the Contains predicate on the "uint32_to_string" field.
func Uint32ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldUint32ToString, v))
}

// Uint32ToStringHasPrefix applies the HasPrefix predicate on the "uint32_to_string" field.
func Uint32ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldUint32ToString, v))
}

// Uint32ToStringHasSuffix applies the HasSuffix predicate on the "uint32_to_string" field.
func Uint32ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldUint32ToString, v))
}

// Uint32ToStringIsNil applies the IsNil predicate on the "uint32_to_string" field.
func Uint32ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldUint32ToString))
}

// Uint32ToStringNotNil applies the NotNil predicate on the "uint32_to_string" field.
func Uint32ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldUint32ToString))
}

// Uint32ToStringEqualFold applies the EqualFold predicate on the "uint32_to_string" field.
func Uint32ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldUint32ToString, v))
}

// Uint32ToStringContainsFold applies the ContainsFold predicate on the "uint32_to_string" field.
func Uint32ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldUint32ToString, v))
}

// Int64ToStringEQ applies the EQ predicate on the "int64_to_string" field.
func Int64ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldInt64ToString, v))
}

// Int64ToStringNEQ applies the NEQ predicate on the "int64_to_string" field.
func Int64ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldInt64ToString, v))
}

// Int64ToStringIn applies the In predicate on the "int64_to_string" field.
func Int64ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldInt64ToString, vs...))
}

// Int64ToStringNotIn applies the NotIn predicate on the "int64_to_string" field.
func Int64ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldInt64ToString, vs...))
}

// Int64ToStringGT applies the GT predicate on the "int64_to_string" field.
func Int64ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldInt64ToString, v))
}

// Int64ToStringGTE applies the GTE predicate on the "int64_to_string" field.
func Int64ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldInt64ToString, v))
}

// Int64ToStringLT applies the LT predicate on the "int64_to_string" field.
func Int64ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldInt64ToString, v))
}

// Int64ToStringLTE applies the LTE predicate on the "int64_to_string" field.
func Int64ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldInt64ToString, v))
}

// Int64ToStringContains applies the Contains predicate on the "int64_to_string" field.
func Int64ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldInt64ToString, v))
}

// Int64ToStringHasPrefix applies the HasPrefix predicate on the "int64_to_string" field.
func Int64ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldInt64ToString, v))
}

// Int64ToStringHasSuffix applies the HasSuffix predicate on the "int64_to_string" field.
func Int64ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldInt64ToString, v))
}

// Int64ToStringIsNil applies the IsNil predicate on the "int64_to_string" field.
func Int64ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldInt64ToString))
}

// Int64ToStringNotNil applies the NotNil predicate on the "int64_to_string" field.
func Int64ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldInt64ToString))
}

// Int64ToStringEqualFold applies the EqualFold predicate on the "int64_to_string" field.
func Int64ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldInt64ToString, v))
}

// Int64ToStringContainsFold applies the ContainsFold predicate on the "int64_to_string" field.
func Int64ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldInt64ToString, v))
}

// Uint64ToStringEQ applies the EQ predicate on the "uint64_to_string" field.
func Uint64ToStringEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEQ(FieldUint64ToString, v))
}

// Uint64ToStringNEQ applies the NEQ predicate on the "uint64_to_string" field.
func Uint64ToStringNEQ(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNEQ(FieldUint64ToString, v))
}

// Uint64ToStringIn applies the In predicate on the "uint64_to_string" field.
func Uint64ToStringIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldIn(FieldUint64ToString, vs...))
}

// Uint64ToStringNotIn applies the NotIn predicate on the "uint64_to_string" field.
func Uint64ToStringNotIn(vs ...string) predicate.Conversion {
	return predicate.Conversion(sql.FieldNotIn(FieldUint64ToString, vs...))
}

// Uint64ToStringGT applies the GT predicate on the "uint64_to_string" field.
func Uint64ToStringGT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGT(FieldUint64ToString, v))
}

// Uint64ToStringGTE applies the GTE predicate on the "uint64_to_string" field.
func Uint64ToStringGTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldGTE(FieldUint64ToString, v))
}

// Uint64ToStringLT applies the LT predicate on the "uint64_to_string" field.
func Uint64ToStringLT(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLT(FieldUint64ToString, v))
}

// Uint64ToStringLTE applies the LTE predicate on the "uint64_to_string" field.
func Uint64ToStringLTE(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldLTE(FieldUint64ToString, v))
}

// Uint64ToStringContains applies the Contains predicate on the "uint64_to_string" field.
func Uint64ToStringContains(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContains(FieldUint64ToString, v))
}

// Uint64ToStringHasPrefix applies the HasPrefix predicate on the "uint64_to_string" field.
func Uint64ToStringHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasPrefix(FieldUint64ToString, v))
}

// Uint64ToStringHasSuffix applies the HasSuffix predicate on the "uint64_to_string" field.
func Uint64ToStringHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldHasSuffix(FieldUint64ToString, v))
}

// Uint64ToStringIsNil applies the IsNil predicate on the "uint64_to_string" field.
func Uint64ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldIsNull(FieldUint64ToString))
}

// Uint64ToStringNotNil applies the NotNil predicate on the "uint64_to_string" field.
func Uint64ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(sql.FieldNotNull(FieldUint64ToString))
}

// Uint64ToStringEqualFold applies the EqualFold predicate on the "uint64_to_string" field.
func Uint64ToStringEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldEqualFold(FieldUint64ToString, v))
}

// Uint64ToStringContainsFold applies the ContainsFold predicate on the "uint64_to_string" field.
func Uint64ToStringContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(sql.FieldContainsFold(FieldUint64ToString, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Conversion) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Conversion) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
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
func Not(p predicate.Conversion) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		p(s.Not())
	})
}
