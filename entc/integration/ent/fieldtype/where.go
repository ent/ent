// Code generated (@generated) by entc, DO NOT EDIT.

package fieldtype

import (
	"strconv"

	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(id)
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.EQ(id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GT(id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GTE(id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LT(id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Within(v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i], _ = strconv.Atoi(ids[i])
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
		func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Without(v...))
		},
	)
}

// Int applies equality check predicate on the "int" field. It's identical to IntEQ.
func Int(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.EQ(v))
		},
	)
}

// Int8 applies equality check predicate on the "int8" field. It's identical to Int8EQ.
func Int8(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.EQ(v))
		},
	)
}

// Int16 applies equality check predicate on the "int16" field. It's identical to Int16EQ.
func Int16(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.EQ(v))
		},
	)
}

// Int32 applies equality check predicate on the "int32" field. It's identical to Int32EQ.
func Int32(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.EQ(v))
		},
	)
}

// Int64 applies equality check predicate on the "int64" field. It's identical to Int64EQ.
func Int64(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.EQ(v))
		},
	)
}

// OptionalInt applies equality check predicate on the "optional_int" field. It's identical to OptionalIntEQ.
func OptionalInt(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.EQ(v))
		},
	)
}

// OptionalInt8 applies equality check predicate on the "optional_int8" field. It's identical to OptionalInt8EQ.
func OptionalInt8(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.EQ(v))
		},
	)
}

// OptionalInt16 applies equality check predicate on the "optional_int16" field. It's identical to OptionalInt16EQ.
func OptionalInt16(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.EQ(v))
		},
	)
}

// OptionalInt32 applies equality check predicate on the "optional_int32" field. It's identical to OptionalInt32EQ.
func OptionalInt32(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.EQ(v))
		},
	)
}

// OptionalInt64 applies equality check predicate on the "optional_int64" field. It's identical to OptionalInt64EQ.
func OptionalInt64(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.EQ(v))
		},
	)
}

// NullableInt applies equality check predicate on the "nullable_int" field. It's identical to NullableIntEQ.
func NullableInt(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.EQ(v))
		},
	)
}

// NullableInt8 applies equality check predicate on the "nullable_int8" field. It's identical to NullableInt8EQ.
func NullableInt8(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.EQ(v))
		},
	)
}

// NullableInt16 applies equality check predicate on the "nullable_int16" field. It's identical to NullableInt16EQ.
func NullableInt16(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.EQ(v))
		},
	)
}

// NullableInt32 applies equality check predicate on the "nullable_int32" field. It's identical to NullableInt32EQ.
func NullableInt32(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.EQ(v))
		},
	)
}

// NullableInt64 applies equality check predicate on the "nullable_int64" field. It's identical to NullableInt64EQ.
func NullableInt64(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.EQ(v))
		},
	)
}

// IntEQ applies the EQ predicate on the "int" field.
func IntEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.EQ(v))
		},
	)
}

// IntNEQ applies the NEQ predicate on the "int" field.
func IntNEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.NEQ(v))
		},
	)
}

// IntGT applies the GT predicate on the "int" field.
func IntGT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.GT(v))
		},
	)
}

// IntGTE applies the GTE predicate on the "int" field.
func IntGTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.GTE(v))
		},
	)
}

// IntLT applies the LT predicate on the "int" field.
func IntLT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.LT(v))
		},
	)
}

// IntLTE applies the LTE predicate on the "int" field.
func IntLTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.LTE(v))
		},
	)
}

// IntIn applies the In predicate on the "int" field.
func IntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.Within(v...))
		},
	)
}

// IntNotIn applies the NotIn predicate on the "int" field.
func IntNotIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt, p.Without(v...))
		},
	)
}

// Int8EQ applies the EQ predicate on the "int8" field.
func Int8EQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.EQ(v))
		},
	)
}

// Int8NEQ applies the NEQ predicate on the "int8" field.
func Int8NEQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.NEQ(v))
		},
	)
}

// Int8GT applies the GT predicate on the "int8" field.
func Int8GT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.GT(v))
		},
	)
}

// Int8GTE applies the GTE predicate on the "int8" field.
func Int8GTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.GTE(v))
		},
	)
}

// Int8LT applies the LT predicate on the "int8" field.
func Int8LT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.LT(v))
		},
	)
}

// Int8LTE applies the LTE predicate on the "int8" field.
func Int8LTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.LTE(v))
		},
	)
}

// Int8In applies the In predicate on the "int8" field.
func Int8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.Within(v...))
		},
	)
}

// Int8NotIn applies the NotIn predicate on the "int8" field.
func Int8NotIn(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt8, p.Without(v...))
		},
	)
}

// Int16EQ applies the EQ predicate on the "int16" field.
func Int16EQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.EQ(v))
		},
	)
}

// Int16NEQ applies the NEQ predicate on the "int16" field.
func Int16NEQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.NEQ(v))
		},
	)
}

// Int16GT applies the GT predicate on the "int16" field.
func Int16GT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.GT(v))
		},
	)
}

// Int16GTE applies the GTE predicate on the "int16" field.
func Int16GTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.GTE(v))
		},
	)
}

// Int16LT applies the LT predicate on the "int16" field.
func Int16LT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.LT(v))
		},
	)
}

// Int16LTE applies the LTE predicate on the "int16" field.
func Int16LTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.LTE(v))
		},
	)
}

// Int16In applies the In predicate on the "int16" field.
func Int16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.Within(v...))
		},
	)
}

// Int16NotIn applies the NotIn predicate on the "int16" field.
func Int16NotIn(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt16, p.Without(v...))
		},
	)
}

// Int32EQ applies the EQ predicate on the "int32" field.
func Int32EQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.EQ(v))
		},
	)
}

// Int32NEQ applies the NEQ predicate on the "int32" field.
func Int32NEQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.NEQ(v))
		},
	)
}

// Int32GT applies the GT predicate on the "int32" field.
func Int32GT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.GT(v))
		},
	)
}

// Int32GTE applies the GTE predicate on the "int32" field.
func Int32GTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.GTE(v))
		},
	)
}

// Int32LT applies the LT predicate on the "int32" field.
func Int32LT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.LT(v))
		},
	)
}

// Int32LTE applies the LTE predicate on the "int32" field.
func Int32LTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.LTE(v))
		},
	)
}

// Int32In applies the In predicate on the "int32" field.
func Int32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.Within(v...))
		},
	)
}

// Int32NotIn applies the NotIn predicate on the "int32" field.
func Int32NotIn(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt32, p.Without(v...))
		},
	)
}

// Int64EQ applies the EQ predicate on the "int64" field.
func Int64EQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.EQ(v))
		},
	)
}

// Int64NEQ applies the NEQ predicate on the "int64" field.
func Int64NEQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.NEQ(v))
		},
	)
}

// Int64GT applies the GT predicate on the "int64" field.
func Int64GT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.GT(v))
		},
	)
}

// Int64GTE applies the GTE predicate on the "int64" field.
func Int64GTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.GTE(v))
		},
	)
}

// Int64LT applies the LT predicate on the "int64" field.
func Int64LT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.LT(v))
		},
	)
}

// Int64LTE applies the LTE predicate on the "int64" field.
func Int64LTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.LTE(v))
		},
	)
}

// Int64In applies the In predicate on the "int64" field.
func Int64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.Within(v...))
		},
	)
}

// Int64NotIn applies the NotIn predicate on the "int64" field.
func Int64NotIn(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldInt64, p.Without(v...))
		},
	)
}

// OptionalIntEQ applies the EQ predicate on the "optional_int" field.
func OptionalIntEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.EQ(v))
		},
	)
}

// OptionalIntNEQ applies the NEQ predicate on the "optional_int" field.
func OptionalIntNEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.NEQ(v))
		},
	)
}

// OptionalIntGT applies the GT predicate on the "optional_int" field.
func OptionalIntGT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.GT(v))
		},
	)
}

// OptionalIntGTE applies the GTE predicate on the "optional_int" field.
func OptionalIntGTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.GTE(v))
		},
	)
}

// OptionalIntLT applies the LT predicate on the "optional_int" field.
func OptionalIntLT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.LT(v))
		},
	)
}

// OptionalIntLTE applies the LTE predicate on the "optional_int" field.
func OptionalIntLTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldOptionalInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.LTE(v))
		},
	)
}

// OptionalIntIn applies the In predicate on the "optional_int" field.
func OptionalIntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldOptionalInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.Within(v...))
		},
	)
}

// OptionalIntNotIn applies the NotIn predicate on the "optional_int" field.
func OptionalIntNotIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldOptionalInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt, p.Without(v...))
		},
	)
}

// OptionalIntIsNull applies the IsNull predicate on the "optional_int" field.
func OptionalIntIsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOptionalInt)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldOptionalInt)
		},
	)
}

// OptionalIntNotNull applies the NotNull predicate on the "optional_int" field.
func OptionalIntNotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOptionalInt)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldOptionalInt)
		},
	)
}

// OptionalInt8EQ applies the EQ predicate on the "optional_int8" field.
func OptionalInt8EQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.EQ(v))
		},
	)
}

// OptionalInt8NEQ applies the NEQ predicate on the "optional_int8" field.
func OptionalInt8NEQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.NEQ(v))
		},
	)
}

// OptionalInt8GT applies the GT predicate on the "optional_int8" field.
func OptionalInt8GT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.GT(v))
		},
	)
}

// OptionalInt8GTE applies the GTE predicate on the "optional_int8" field.
func OptionalInt8GTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.GTE(v))
		},
	)
}

// OptionalInt8LT applies the LT predicate on the "optional_int8" field.
func OptionalInt8LT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.LT(v))
		},
	)
}

// OptionalInt8LTE applies the LTE predicate on the "optional_int8" field.
func OptionalInt8LTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldOptionalInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.LTE(v))
		},
	)
}

// OptionalInt8In applies the In predicate on the "optional_int8" field.
func OptionalInt8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldOptionalInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.Within(v...))
		},
	)
}

// OptionalInt8NotIn applies the NotIn predicate on the "optional_int8" field.
func OptionalInt8NotIn(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldOptionalInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt8, p.Without(v...))
		},
	)
}

// OptionalInt8IsNull applies the IsNull predicate on the "optional_int8" field.
func OptionalInt8IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOptionalInt8)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldOptionalInt8)
		},
	)
}

// OptionalInt8NotNull applies the NotNull predicate on the "optional_int8" field.
func OptionalInt8NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOptionalInt8)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldOptionalInt8)
		},
	)
}

// OptionalInt16EQ applies the EQ predicate on the "optional_int16" field.
func OptionalInt16EQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.EQ(v))
		},
	)
}

// OptionalInt16NEQ applies the NEQ predicate on the "optional_int16" field.
func OptionalInt16NEQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.NEQ(v))
		},
	)
}

// OptionalInt16GT applies the GT predicate on the "optional_int16" field.
func OptionalInt16GT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.GT(v))
		},
	)
}

// OptionalInt16GTE applies the GTE predicate on the "optional_int16" field.
func OptionalInt16GTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.GTE(v))
		},
	)
}

// OptionalInt16LT applies the LT predicate on the "optional_int16" field.
func OptionalInt16LT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.LT(v))
		},
	)
}

// OptionalInt16LTE applies the LTE predicate on the "optional_int16" field.
func OptionalInt16LTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldOptionalInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.LTE(v))
		},
	)
}

// OptionalInt16In applies the In predicate on the "optional_int16" field.
func OptionalInt16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldOptionalInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.Within(v...))
		},
	)
}

// OptionalInt16NotIn applies the NotIn predicate on the "optional_int16" field.
func OptionalInt16NotIn(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldOptionalInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt16, p.Without(v...))
		},
	)
}

// OptionalInt16IsNull applies the IsNull predicate on the "optional_int16" field.
func OptionalInt16IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOptionalInt16)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldOptionalInt16)
		},
	)
}

// OptionalInt16NotNull applies the NotNull predicate on the "optional_int16" field.
func OptionalInt16NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOptionalInt16)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldOptionalInt16)
		},
	)
}

// OptionalInt32EQ applies the EQ predicate on the "optional_int32" field.
func OptionalInt32EQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.EQ(v))
		},
	)
}

// OptionalInt32NEQ applies the NEQ predicate on the "optional_int32" field.
func OptionalInt32NEQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.NEQ(v))
		},
	)
}

// OptionalInt32GT applies the GT predicate on the "optional_int32" field.
func OptionalInt32GT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.GT(v))
		},
	)
}

// OptionalInt32GTE applies the GTE predicate on the "optional_int32" field.
func OptionalInt32GTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.GTE(v))
		},
	)
}

// OptionalInt32LT applies the LT predicate on the "optional_int32" field.
func OptionalInt32LT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.LT(v))
		},
	)
}

// OptionalInt32LTE applies the LTE predicate on the "optional_int32" field.
func OptionalInt32LTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldOptionalInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.LTE(v))
		},
	)
}

// OptionalInt32In applies the In predicate on the "optional_int32" field.
func OptionalInt32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldOptionalInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.Within(v...))
		},
	)
}

// OptionalInt32NotIn applies the NotIn predicate on the "optional_int32" field.
func OptionalInt32NotIn(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldOptionalInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt32, p.Without(v...))
		},
	)
}

// OptionalInt32IsNull applies the IsNull predicate on the "optional_int32" field.
func OptionalInt32IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOptionalInt32)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldOptionalInt32)
		},
	)
}

// OptionalInt32NotNull applies the NotNull predicate on the "optional_int32" field.
func OptionalInt32NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOptionalInt32)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldOptionalInt32)
		},
	)
}

// OptionalInt64EQ applies the EQ predicate on the "optional_int64" field.
func OptionalInt64EQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.EQ(v))
		},
	)
}

// OptionalInt64NEQ applies the NEQ predicate on the "optional_int64" field.
func OptionalInt64NEQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.NEQ(v))
		},
	)
}

// OptionalInt64GT applies the GT predicate on the "optional_int64" field.
func OptionalInt64GT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.GT(v))
		},
	)
}

// OptionalInt64GTE applies the GTE predicate on the "optional_int64" field.
func OptionalInt64GTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.GTE(v))
		},
	)
}

// OptionalInt64LT applies the LT predicate on the "optional_int64" field.
func OptionalInt64LT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.LT(v))
		},
	)
}

// OptionalInt64LTE applies the LTE predicate on the "optional_int64" field.
func OptionalInt64LTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldOptionalInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.LTE(v))
		},
	)
}

// OptionalInt64In applies the In predicate on the "optional_int64" field.
func OptionalInt64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldOptionalInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.Within(v...))
		},
	)
}

// OptionalInt64NotIn applies the NotIn predicate on the "optional_int64" field.
func OptionalInt64NotIn(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldOptionalInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldOptionalInt64, p.Without(v...))
		},
	)
}

// OptionalInt64IsNull applies the IsNull predicate on the "optional_int64" field.
func OptionalInt64IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldOptionalInt64)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldOptionalInt64)
		},
	)
}

// OptionalInt64NotNull applies the NotNull predicate on the "optional_int64" field.
func OptionalInt64NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldOptionalInt64)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldOptionalInt64)
		},
	)
}

// NullableIntEQ applies the EQ predicate on the "nullable_int" field.
func NullableIntEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.EQ(v))
		},
	)
}

// NullableIntNEQ applies the NEQ predicate on the "nullable_int" field.
func NullableIntNEQ(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.NEQ(v))
		},
	)
}

// NullableIntGT applies the GT predicate on the "nullable_int" field.
func NullableIntGT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.GT(v))
		},
	)
}

// NullableIntGTE applies the GTE predicate on the "nullable_int" field.
func NullableIntGTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.GTE(v))
		},
	)
}

// NullableIntLT applies the LT predicate on the "nullable_int" field.
func NullableIntLT(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.LT(v))
		},
	)
}

// NullableIntLTE applies the LTE predicate on the "nullable_int" field.
func NullableIntLTE(v int) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNullableInt), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.LTE(v))
		},
	)
}

// NullableIntIn applies the In predicate on the "nullable_int" field.
func NullableIntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNullableInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.Within(v...))
		},
	)
}

// NullableIntNotIn applies the NotIn predicate on the "nullable_int" field.
func NullableIntNotIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNullableInt), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt, p.Without(v...))
		},
	)
}

// NullableIntIsNull applies the IsNull predicate on the "nullable_int" field.
func NullableIntIsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNullableInt)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldNullableInt)
		},
	)
}

// NullableIntNotNull applies the NotNull predicate on the "nullable_int" field.
func NullableIntNotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNullableInt)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldNullableInt)
		},
	)
}

// NullableInt8EQ applies the EQ predicate on the "nullable_int8" field.
func NullableInt8EQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.EQ(v))
		},
	)
}

// NullableInt8NEQ applies the NEQ predicate on the "nullable_int8" field.
func NullableInt8NEQ(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.NEQ(v))
		},
	)
}

// NullableInt8GT applies the GT predicate on the "nullable_int8" field.
func NullableInt8GT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.GT(v))
		},
	)
}

// NullableInt8GTE applies the GTE predicate on the "nullable_int8" field.
func NullableInt8GTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.GTE(v))
		},
	)
}

// NullableInt8LT applies the LT predicate on the "nullable_int8" field.
func NullableInt8LT(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.LT(v))
		},
	)
}

// NullableInt8LTE applies the LTE predicate on the "nullable_int8" field.
func NullableInt8LTE(v int8) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNullableInt8), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.LTE(v))
		},
	)
}

// NullableInt8In applies the In predicate on the "nullable_int8" field.
func NullableInt8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNullableInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.Within(v...))
		},
	)
}

// NullableInt8NotIn applies the NotIn predicate on the "nullable_int8" field.
func NullableInt8NotIn(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNullableInt8), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt8, p.Without(v...))
		},
	)
}

// NullableInt8IsNull applies the IsNull predicate on the "nullable_int8" field.
func NullableInt8IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNullableInt8)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldNullableInt8)
		},
	)
}

// NullableInt8NotNull applies the NotNull predicate on the "nullable_int8" field.
func NullableInt8NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNullableInt8)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldNullableInt8)
		},
	)
}

// NullableInt16EQ applies the EQ predicate on the "nullable_int16" field.
func NullableInt16EQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.EQ(v))
		},
	)
}

// NullableInt16NEQ applies the NEQ predicate on the "nullable_int16" field.
func NullableInt16NEQ(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.NEQ(v))
		},
	)
}

// NullableInt16GT applies the GT predicate on the "nullable_int16" field.
func NullableInt16GT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.GT(v))
		},
	)
}

// NullableInt16GTE applies the GTE predicate on the "nullable_int16" field.
func NullableInt16GTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.GTE(v))
		},
	)
}

// NullableInt16LT applies the LT predicate on the "nullable_int16" field.
func NullableInt16LT(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.LT(v))
		},
	)
}

// NullableInt16LTE applies the LTE predicate on the "nullable_int16" field.
func NullableInt16LTE(v int16) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNullableInt16), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.LTE(v))
		},
	)
}

// NullableInt16In applies the In predicate on the "nullable_int16" field.
func NullableInt16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNullableInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.Within(v...))
		},
	)
}

// NullableInt16NotIn applies the NotIn predicate on the "nullable_int16" field.
func NullableInt16NotIn(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNullableInt16), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt16, p.Without(v...))
		},
	)
}

// NullableInt16IsNull applies the IsNull predicate on the "nullable_int16" field.
func NullableInt16IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNullableInt16)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldNullableInt16)
		},
	)
}

// NullableInt16NotNull applies the NotNull predicate on the "nullable_int16" field.
func NullableInt16NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNullableInt16)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldNullableInt16)
		},
	)
}

// NullableInt32EQ applies the EQ predicate on the "nullable_int32" field.
func NullableInt32EQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.EQ(v))
		},
	)
}

// NullableInt32NEQ applies the NEQ predicate on the "nullable_int32" field.
func NullableInt32NEQ(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.NEQ(v))
		},
	)
}

// NullableInt32GT applies the GT predicate on the "nullable_int32" field.
func NullableInt32GT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.GT(v))
		},
	)
}

// NullableInt32GTE applies the GTE predicate on the "nullable_int32" field.
func NullableInt32GTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.GTE(v))
		},
	)
}

// NullableInt32LT applies the LT predicate on the "nullable_int32" field.
func NullableInt32LT(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.LT(v))
		},
	)
}

// NullableInt32LTE applies the LTE predicate on the "nullable_int32" field.
func NullableInt32LTE(v int32) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNullableInt32), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.LTE(v))
		},
	)
}

// NullableInt32In applies the In predicate on the "nullable_int32" field.
func NullableInt32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNullableInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.Within(v...))
		},
	)
}

// NullableInt32NotIn applies the NotIn predicate on the "nullable_int32" field.
func NullableInt32NotIn(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNullableInt32), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt32, p.Without(v...))
		},
	)
}

// NullableInt32IsNull applies the IsNull predicate on the "nullable_int32" field.
func NullableInt32IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNullableInt32)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldNullableInt32)
		},
	)
}

// NullableInt32NotNull applies the NotNull predicate on the "nullable_int32" field.
func NullableInt32NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNullableInt32)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldNullableInt32)
		},
	)
}

// NullableInt64EQ applies the EQ predicate on the "nullable_int64" field.
func NullableInt64EQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.EQ(v))
		},
	)
}

// NullableInt64NEQ applies the NEQ predicate on the "nullable_int64" field.
func NullableInt64NEQ(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.NEQ(v))
		},
	)
}

// NullableInt64GT applies the GT predicate on the "nullable_int64" field.
func NullableInt64GT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.GT(v))
		},
	)
}

// NullableInt64GTE applies the GTE predicate on the "nullable_int64" field.
func NullableInt64GTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.GTE(v))
		},
	)
}

// NullableInt64LT applies the LT predicate on the "nullable_int64" field.
func NullableInt64LT(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.LT(v))
		},
	)
}

// NullableInt64LTE applies the LTE predicate on the "nullable_int64" field.
func NullableInt64LTE(v int64) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNullableInt64), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.LTE(v))
		},
	)
}

// NullableInt64In applies the In predicate on the "nullable_int64" field.
func NullableInt64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNullableInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.Within(v...))
		},
	)
}

// NullableInt64NotIn applies the NotIn predicate on the "nullable_int64" field.
func NullableInt64NotIn(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNullableInt64), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNullableInt64, p.Without(v...))
		},
	)
}

// NullableInt64IsNull applies the IsNull predicate on the "nullable_int64" field.
func NullableInt64IsNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNullableInt64)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldNullableInt64)
		},
	)
}

// NullableInt64NotNull applies the NotNull predicate on the "nullable_int64" field.
func NullableInt64NotNull() predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNullableInt64)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldNullableInt64)
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.FieldType) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.And(trs...))
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.FieldType) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
		func(tr *dsl.Traversal) {
			trs := make([]interface{}, 0, len(predicates))
			for _, p := range predicates {
				t := __.New()
				p(t)
				trs = append(trs, t)
			}
			tr.Where(__.Or(trs...))
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.FieldType) predicate.FieldType {
	return predicate.FieldTypePerDialect(
		func(s *sql.Selector) {
			p(s.Not())
		},
		func(tr *dsl.Traversal) {
			t := __.New()
			p(t)
			tr.Where(__.Not(t))
		},
	)
}
