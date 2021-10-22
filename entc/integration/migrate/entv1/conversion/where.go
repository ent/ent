// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package conversion

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv1/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Int8ToString applies equality check predicate on the "int8_to_string" field. It's identical to Int8ToStringEQ.
func Int8ToString(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt8ToString), v))
	})
}

// Uint8ToString applies equality check predicate on the "uint8_to_string" field. It's identical to Uint8ToStringEQ.
func Uint8ToString(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint8ToString), v))
	})
}

// Int16ToString applies equality check predicate on the "int16_to_string" field. It's identical to Int16ToStringEQ.
func Int16ToString(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt16ToString), v))
	})
}

// Uint16ToString applies equality check predicate on the "uint16_to_string" field. It's identical to Uint16ToStringEQ.
func Uint16ToString(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint16ToString), v))
	})
}

// Int32ToString applies equality check predicate on the "int32_to_string" field. It's identical to Int32ToStringEQ.
func Int32ToString(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt32ToString), v))
	})
}

// Uint32ToString applies equality check predicate on the "uint32_to_string" field. It's identical to Uint32ToStringEQ.
func Uint32ToString(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint32ToString), v))
	})
}

// Int64ToString applies equality check predicate on the "int64_to_string" field. It's identical to Int64ToStringEQ.
func Int64ToString(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt64ToString), v))
	})
}

// Uint64ToString applies equality check predicate on the "uint64_to_string" field. It's identical to Uint64ToStringEQ.
func Uint64ToString(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint64ToString), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameIsNil applies the IsNil predicate on the "name" field.
func NameIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldName)))
	})
}

// NameNotNil applies the NotNil predicate on the "name" field.
func NameNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldName)))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// NameHasPrefixFold applies the HasPrefixFold predicate on the "name" field.
func NameHasPrefixFold(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.HasPrefixFold(s.C(FieldName), v))
	})
}

// NameHasSuffixFold applies the HasSuffixFold predicate on the "name" field.
func NameHasSuffixFold(v string) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.HasSuffixFold(s.C(FieldName), v))
	})
}

// Int8ToStringEQ applies the EQ predicate on the "int8_to_string" field.
func Int8ToStringEQ(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringNEQ applies the NEQ predicate on the "int8_to_string" field.
func Int8ToStringNEQ(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringIn applies the In predicate on the "int8_to_string" field.
func Int8ToStringIn(vs ...int8) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInt8ToString), v...))
	})
}

// Int8ToStringNotIn applies the NotIn predicate on the "int8_to_string" field.
func Int8ToStringNotIn(vs ...int8) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInt8ToString), v...))
	})
}

// Int8ToStringGT applies the GT predicate on the "int8_to_string" field.
func Int8ToStringGT(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringGTE applies the GTE predicate on the "int8_to_string" field.
func Int8ToStringGTE(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringLT applies the LT predicate on the "int8_to_string" field.
func Int8ToStringLT(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringLTE applies the LTE predicate on the "int8_to_string" field.
func Int8ToStringLTE(v int8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInt8ToString), v))
	})
}

// Int8ToStringIsNil applies the IsNil predicate on the "int8_to_string" field.
func Int8ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInt8ToString)))
	})
}

// Int8ToStringNotNil applies the NotNil predicate on the "int8_to_string" field.
func Int8ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInt8ToString)))
	})
}

// Uint8ToStringEQ applies the EQ predicate on the "uint8_to_string" field.
func Uint8ToStringEQ(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringNEQ applies the NEQ predicate on the "uint8_to_string" field.
func Uint8ToStringNEQ(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringIn applies the In predicate on the "uint8_to_string" field.
func Uint8ToStringIn(vs ...uint8) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUint8ToString), v...))
	})
}

// Uint8ToStringNotIn applies the NotIn predicate on the "uint8_to_string" field.
func Uint8ToStringNotIn(vs ...uint8) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUint8ToString), v...))
	})
}

// Uint8ToStringGT applies the GT predicate on the "uint8_to_string" field.
func Uint8ToStringGT(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringGTE applies the GTE predicate on the "uint8_to_string" field.
func Uint8ToStringGTE(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringLT applies the LT predicate on the "uint8_to_string" field.
func Uint8ToStringLT(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringLTE applies the LTE predicate on the "uint8_to_string" field.
func Uint8ToStringLTE(v uint8) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUint8ToString), v))
	})
}

// Uint8ToStringIsNil applies the IsNil predicate on the "uint8_to_string" field.
func Uint8ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUint8ToString)))
	})
}

// Uint8ToStringNotNil applies the NotNil predicate on the "uint8_to_string" field.
func Uint8ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUint8ToString)))
	})
}

// Int16ToStringEQ applies the EQ predicate on the "int16_to_string" field.
func Int16ToStringEQ(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringNEQ applies the NEQ predicate on the "int16_to_string" field.
func Int16ToStringNEQ(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringIn applies the In predicate on the "int16_to_string" field.
func Int16ToStringIn(vs ...int16) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInt16ToString), v...))
	})
}

// Int16ToStringNotIn applies the NotIn predicate on the "int16_to_string" field.
func Int16ToStringNotIn(vs ...int16) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInt16ToString), v...))
	})
}

// Int16ToStringGT applies the GT predicate on the "int16_to_string" field.
func Int16ToStringGT(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringGTE applies the GTE predicate on the "int16_to_string" field.
func Int16ToStringGTE(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringLT applies the LT predicate on the "int16_to_string" field.
func Int16ToStringLT(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringLTE applies the LTE predicate on the "int16_to_string" field.
func Int16ToStringLTE(v int16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInt16ToString), v))
	})
}

// Int16ToStringIsNil applies the IsNil predicate on the "int16_to_string" field.
func Int16ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInt16ToString)))
	})
}

// Int16ToStringNotNil applies the NotNil predicate on the "int16_to_string" field.
func Int16ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInt16ToString)))
	})
}

// Uint16ToStringEQ applies the EQ predicate on the "uint16_to_string" field.
func Uint16ToStringEQ(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringNEQ applies the NEQ predicate on the "uint16_to_string" field.
func Uint16ToStringNEQ(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringIn applies the In predicate on the "uint16_to_string" field.
func Uint16ToStringIn(vs ...uint16) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUint16ToString), v...))
	})
}

// Uint16ToStringNotIn applies the NotIn predicate on the "uint16_to_string" field.
func Uint16ToStringNotIn(vs ...uint16) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUint16ToString), v...))
	})
}

// Uint16ToStringGT applies the GT predicate on the "uint16_to_string" field.
func Uint16ToStringGT(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringGTE applies the GTE predicate on the "uint16_to_string" field.
func Uint16ToStringGTE(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringLT applies the LT predicate on the "uint16_to_string" field.
func Uint16ToStringLT(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringLTE applies the LTE predicate on the "uint16_to_string" field.
func Uint16ToStringLTE(v uint16) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUint16ToString), v))
	})
}

// Uint16ToStringIsNil applies the IsNil predicate on the "uint16_to_string" field.
func Uint16ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUint16ToString)))
	})
}

// Uint16ToStringNotNil applies the NotNil predicate on the "uint16_to_string" field.
func Uint16ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUint16ToString)))
	})
}

// Int32ToStringEQ applies the EQ predicate on the "int32_to_string" field.
func Int32ToStringEQ(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringNEQ applies the NEQ predicate on the "int32_to_string" field.
func Int32ToStringNEQ(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringIn applies the In predicate on the "int32_to_string" field.
func Int32ToStringIn(vs ...int32) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInt32ToString), v...))
	})
}

// Int32ToStringNotIn applies the NotIn predicate on the "int32_to_string" field.
func Int32ToStringNotIn(vs ...int32) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInt32ToString), v...))
	})
}

// Int32ToStringGT applies the GT predicate on the "int32_to_string" field.
func Int32ToStringGT(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringGTE applies the GTE predicate on the "int32_to_string" field.
func Int32ToStringGTE(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringLT applies the LT predicate on the "int32_to_string" field.
func Int32ToStringLT(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringLTE applies the LTE predicate on the "int32_to_string" field.
func Int32ToStringLTE(v int32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInt32ToString), v))
	})
}

// Int32ToStringIsNil applies the IsNil predicate on the "int32_to_string" field.
func Int32ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInt32ToString)))
	})
}

// Int32ToStringNotNil applies the NotNil predicate on the "int32_to_string" field.
func Int32ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInt32ToString)))
	})
}

// Uint32ToStringEQ applies the EQ predicate on the "uint32_to_string" field.
func Uint32ToStringEQ(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringNEQ applies the NEQ predicate on the "uint32_to_string" field.
func Uint32ToStringNEQ(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringIn applies the In predicate on the "uint32_to_string" field.
func Uint32ToStringIn(vs ...uint32) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUint32ToString), v...))
	})
}

// Uint32ToStringNotIn applies the NotIn predicate on the "uint32_to_string" field.
func Uint32ToStringNotIn(vs ...uint32) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUint32ToString), v...))
	})
}

// Uint32ToStringGT applies the GT predicate on the "uint32_to_string" field.
func Uint32ToStringGT(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringGTE applies the GTE predicate on the "uint32_to_string" field.
func Uint32ToStringGTE(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringLT applies the LT predicate on the "uint32_to_string" field.
func Uint32ToStringLT(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringLTE applies the LTE predicate on the "uint32_to_string" field.
func Uint32ToStringLTE(v uint32) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUint32ToString), v))
	})
}

// Uint32ToStringIsNil applies the IsNil predicate on the "uint32_to_string" field.
func Uint32ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUint32ToString)))
	})
}

// Uint32ToStringNotNil applies the NotNil predicate on the "uint32_to_string" field.
func Uint32ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUint32ToString)))
	})
}

// Int64ToStringEQ applies the EQ predicate on the "int64_to_string" field.
func Int64ToStringEQ(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringNEQ applies the NEQ predicate on the "int64_to_string" field.
func Int64ToStringNEQ(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringIn applies the In predicate on the "int64_to_string" field.
func Int64ToStringIn(vs ...int64) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldInt64ToString), v...))
	})
}

// Int64ToStringNotIn applies the NotIn predicate on the "int64_to_string" field.
func Int64ToStringNotIn(vs ...int64) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldInt64ToString), v...))
	})
}

// Int64ToStringGT applies the GT predicate on the "int64_to_string" field.
func Int64ToStringGT(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringGTE applies the GTE predicate on the "int64_to_string" field.
func Int64ToStringGTE(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringLT applies the LT predicate on the "int64_to_string" field.
func Int64ToStringLT(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringLTE applies the LTE predicate on the "int64_to_string" field.
func Int64ToStringLTE(v int64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldInt64ToString), v))
	})
}

// Int64ToStringIsNil applies the IsNil predicate on the "int64_to_string" field.
func Int64ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInt64ToString)))
	})
}

// Int64ToStringNotNil applies the NotNil predicate on the "int64_to_string" field.
func Int64ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInt64ToString)))
	})
}

// Uint64ToStringEQ applies the EQ predicate on the "uint64_to_string" field.
func Uint64ToStringEQ(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringNEQ applies the NEQ predicate on the "uint64_to_string" field.
func Uint64ToStringNEQ(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringIn applies the In predicate on the "uint64_to_string" field.
func Uint64ToStringIn(vs ...uint64) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUint64ToString), v...))
	})
}

// Uint64ToStringNotIn applies the NotIn predicate on the "uint64_to_string" field.
func Uint64ToStringNotIn(vs ...uint64) predicate.Conversion {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Conversion(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUint64ToString), v...))
	})
}

// Uint64ToStringGT applies the GT predicate on the "uint64_to_string" field.
func Uint64ToStringGT(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringGTE applies the GTE predicate on the "uint64_to_string" field.
func Uint64ToStringGTE(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringLT applies the LT predicate on the "uint64_to_string" field.
func Uint64ToStringLT(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringLTE applies the LTE predicate on the "uint64_to_string" field.
func Uint64ToStringLTE(v uint64) predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUint64ToString), v))
	})
}

// Uint64ToStringIsNil applies the IsNil predicate on the "uint64_to_string" field.
func Uint64ToStringIsNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUint64ToString)))
	})
}

// Uint64ToStringNotNil applies the NotNil predicate on the "uint64_to_string" field.
func Uint64ToStringNotNil() predicate.Conversion {
	return predicate.Conversion(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUint64ToString)))
	})
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
