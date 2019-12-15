// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package fieldtype

import (
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.FieldType {
	return predicate.FieldType(
		func(t *dsl.Traversal) {
			t.HasID(id)
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	},
	)
}

// Int applies equality check predicate on the "int" field. It's identical to IntEQ.
func Int(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.EQ(v))
	},
	)
}

// Int8 applies equality check predicate on the "int8" field. It's identical to Int8EQ.
func Int8(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.EQ(v))
	},
	)
}

// Int16 applies equality check predicate on the "int16" field. It's identical to Int16EQ.
func Int16(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.EQ(v))
	},
	)
}

// Int32 applies equality check predicate on the "int32" field. It's identical to Int32EQ.
func Int32(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.EQ(v))
	},
	)
}

// Int64 applies equality check predicate on the "int64" field. It's identical to Int64EQ.
func Int64(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.EQ(v))
	},
	)
}

// OptionalInt applies equality check predicate on the "optional_int" field. It's identical to OptionalIntEQ.
func OptionalInt(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.EQ(v))
	},
	)
}

// OptionalInt8 applies equality check predicate on the "optional_int8" field. It's identical to OptionalInt8EQ.
func OptionalInt8(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.EQ(v))
	},
	)
}

// OptionalInt16 applies equality check predicate on the "optional_int16" field. It's identical to OptionalInt16EQ.
func OptionalInt16(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.EQ(v))
	},
	)
}

// OptionalInt32 applies equality check predicate on the "optional_int32" field. It's identical to OptionalInt32EQ.
func OptionalInt32(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.EQ(v))
	},
	)
}

// OptionalInt64 applies equality check predicate on the "optional_int64" field. It's identical to OptionalInt64EQ.
func OptionalInt64(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.EQ(v))
	},
	)
}

// NillableInt applies equality check predicate on the "nillable_int" field. It's identical to NillableIntEQ.
func NillableInt(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.EQ(v))
	},
	)
}

// NillableInt8 applies equality check predicate on the "nillable_int8" field. It's identical to NillableInt8EQ.
func NillableInt8(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.EQ(v))
	},
	)
}

// NillableInt16 applies equality check predicate on the "nillable_int16" field. It's identical to NillableInt16EQ.
func NillableInt16(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.EQ(v))
	},
	)
}

// NillableInt32 applies equality check predicate on the "nillable_int32" field. It's identical to NillableInt32EQ.
func NillableInt32(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.EQ(v))
	},
	)
}

// NillableInt64 applies equality check predicate on the "nillable_int64" field. It's identical to NillableInt64EQ.
func NillableInt64(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.EQ(v))
	},
	)
}

// ValidateOptionalInt32 applies equality check predicate on the "validate_optional_int32" field. It's identical to ValidateOptionalInt32EQ.
func ValidateOptionalInt32(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.EQ(v))
	},
	)
}

// IntEQ applies the EQ predicate on the "int" field.
func IntEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.EQ(v))
	},
	)
}

// IntNEQ applies the NEQ predicate on the "int" field.
func IntNEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.NEQ(v))
	},
	)
}

// IntIn applies the In predicate on the "int" field.
func IntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.Without(v...))
	},
	)
}

// IntGT applies the GT predicate on the "int" field.
func IntGT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.GT(v))
	},
	)
}

// IntGTE applies the GTE predicate on the "int" field.
func IntGTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.GTE(v))
	},
	)
}

// IntLT applies the LT predicate on the "int" field.
func IntLT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.LT(v))
	},
	)
}

// IntLTE applies the LTE predicate on the "int" field.
func IntLTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt, p.LTE(v))
	},
	)
}

// Int8EQ applies the EQ predicate on the "int8" field.
func Int8EQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.EQ(v))
	},
	)
}

// Int8NEQ applies the NEQ predicate on the "int8" field.
func Int8NEQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.NEQ(v))
	},
	)
}

// Int8In applies the In predicate on the "int8" field.
func Int8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.Without(v...))
	},
	)
}

// Int8GT applies the GT predicate on the "int8" field.
func Int8GT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.GT(v))
	},
	)
}

// Int8GTE applies the GTE predicate on the "int8" field.
func Int8GTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.GTE(v))
	},
	)
}

// Int8LT applies the LT predicate on the "int8" field.
func Int8LT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.LT(v))
	},
	)
}

// Int8LTE applies the LTE predicate on the "int8" field.
func Int8LTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt8, p.LTE(v))
	},
	)
}

// Int16EQ applies the EQ predicate on the "int16" field.
func Int16EQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.EQ(v))
	},
	)
}

// Int16NEQ applies the NEQ predicate on the "int16" field.
func Int16NEQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.NEQ(v))
	},
	)
}

// Int16In applies the In predicate on the "int16" field.
func Int16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.Without(v...))
	},
	)
}

// Int16GT applies the GT predicate on the "int16" field.
func Int16GT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.GT(v))
	},
	)
}

// Int16GTE applies the GTE predicate on the "int16" field.
func Int16GTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.GTE(v))
	},
	)
}

// Int16LT applies the LT predicate on the "int16" field.
func Int16LT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.LT(v))
	},
	)
}

// Int16LTE applies the LTE predicate on the "int16" field.
func Int16LTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt16, p.LTE(v))
	},
	)
}

// Int32EQ applies the EQ predicate on the "int32" field.
func Int32EQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.EQ(v))
	},
	)
}

// Int32NEQ applies the NEQ predicate on the "int32" field.
func Int32NEQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.NEQ(v))
	},
	)
}

// Int32In applies the In predicate on the "int32" field.
func Int32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.Without(v...))
	},
	)
}

// Int32GT applies the GT predicate on the "int32" field.
func Int32GT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.GT(v))
	},
	)
}

// Int32GTE applies the GTE predicate on the "int32" field.
func Int32GTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.GTE(v))
	},
	)
}

// Int32LT applies the LT predicate on the "int32" field.
func Int32LT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.LT(v))
	},
	)
}

// Int32LTE applies the LTE predicate on the "int32" field.
func Int32LTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt32, p.LTE(v))
	},
	)
}

// Int64EQ applies the EQ predicate on the "int64" field.
func Int64EQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.EQ(v))
	},
	)
}

// Int64NEQ applies the NEQ predicate on the "int64" field.
func Int64NEQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.NEQ(v))
	},
	)
}

// Int64In applies the In predicate on the "int64" field.
func Int64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.Without(v...))
	},
	)
}

// Int64GT applies the GT predicate on the "int64" field.
func Int64GT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.GT(v))
	},
	)
}

// Int64GTE applies the GTE predicate on the "int64" field.
func Int64GTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.GTE(v))
	},
	)
}

// Int64LT applies the LT predicate on the "int64" field.
func Int64LT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.LT(v))
	},
	)
}

// Int64LTE applies the LTE predicate on the "int64" field.
func Int64LTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldInt64, p.LTE(v))
	},
	)
}

// OptionalIntEQ applies the EQ predicate on the "optional_int" field.
func OptionalIntEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.EQ(v))
	},
	)
}

// OptionalIntNEQ applies the NEQ predicate on the "optional_int" field.
func OptionalIntNEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.NEQ(v))
	},
	)
}

// OptionalIntIn applies the In predicate on the "optional_int" field.
func OptionalIntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.Without(v...))
	},
	)
}

// OptionalIntGT applies the GT predicate on the "optional_int" field.
func OptionalIntGT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.GT(v))
	},
	)
}

// OptionalIntGTE applies the GTE predicate on the "optional_int" field.
func OptionalIntGTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.GTE(v))
	},
	)
}

// OptionalIntLT applies the LT predicate on the "optional_int" field.
func OptionalIntLT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.LT(v))
	},
	)
}

// OptionalIntLTE applies the LTE predicate on the "optional_int" field.
func OptionalIntLTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.LTE(v))
	},
	)
}

// OptionalIntIsNil applies the IsNil predicate on the "optional_int" field.
func OptionalIntIsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt)
	},
	)
}

// OptionalIntNotNil applies the NotNil predicate on the "optional_int" field.
func OptionalIntNotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt)
	},
	)
}

// OptionalInt8EQ applies the EQ predicate on the "optional_int8" field.
func OptionalInt8EQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.EQ(v))
	},
	)
}

// OptionalInt8NEQ applies the NEQ predicate on the "optional_int8" field.
func OptionalInt8NEQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.NEQ(v))
	},
	)
}

// OptionalInt8In applies the In predicate on the "optional_int8" field.
func OptionalInt8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.Without(v...))
	},
	)
}

// OptionalInt8GT applies the GT predicate on the "optional_int8" field.
func OptionalInt8GT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.GT(v))
	},
	)
}

// OptionalInt8GTE applies the GTE predicate on the "optional_int8" field.
func OptionalInt8GTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.GTE(v))
	},
	)
}

// OptionalInt8LT applies the LT predicate on the "optional_int8" field.
func OptionalInt8LT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.LT(v))
	},
	)
}

// OptionalInt8LTE applies the LTE predicate on the "optional_int8" field.
func OptionalInt8LTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt8, p.LTE(v))
	},
	)
}

// OptionalInt8IsNil applies the IsNil predicate on the "optional_int8" field.
func OptionalInt8IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt8)
	},
	)
}

// OptionalInt8NotNil applies the NotNil predicate on the "optional_int8" field.
func OptionalInt8NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt8)
	},
	)
}

// OptionalInt16EQ applies the EQ predicate on the "optional_int16" field.
func OptionalInt16EQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.EQ(v))
	},
	)
}

// OptionalInt16NEQ applies the NEQ predicate on the "optional_int16" field.
func OptionalInt16NEQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.NEQ(v))
	},
	)
}

// OptionalInt16In applies the In predicate on the "optional_int16" field.
func OptionalInt16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.Without(v...))
	},
	)
}

// OptionalInt16GT applies the GT predicate on the "optional_int16" field.
func OptionalInt16GT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.GT(v))
	},
	)
}

// OptionalInt16GTE applies the GTE predicate on the "optional_int16" field.
func OptionalInt16GTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.GTE(v))
	},
	)
}

// OptionalInt16LT applies the LT predicate on the "optional_int16" field.
func OptionalInt16LT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.LT(v))
	},
	)
}

// OptionalInt16LTE applies the LTE predicate on the "optional_int16" field.
func OptionalInt16LTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt16, p.LTE(v))
	},
	)
}

// OptionalInt16IsNil applies the IsNil predicate on the "optional_int16" field.
func OptionalInt16IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt16)
	},
	)
}

// OptionalInt16NotNil applies the NotNil predicate on the "optional_int16" field.
func OptionalInt16NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt16)
	},
	)
}

// OptionalInt32EQ applies the EQ predicate on the "optional_int32" field.
func OptionalInt32EQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.EQ(v))
	},
	)
}

// OptionalInt32NEQ applies the NEQ predicate on the "optional_int32" field.
func OptionalInt32NEQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.NEQ(v))
	},
	)
}

// OptionalInt32In applies the In predicate on the "optional_int32" field.
func OptionalInt32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.Without(v...))
	},
	)
}

// OptionalInt32GT applies the GT predicate on the "optional_int32" field.
func OptionalInt32GT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.GT(v))
	},
	)
}

// OptionalInt32GTE applies the GTE predicate on the "optional_int32" field.
func OptionalInt32GTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.GTE(v))
	},
	)
}

// OptionalInt32LT applies the LT predicate on the "optional_int32" field.
func OptionalInt32LT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.LT(v))
	},
	)
}

// OptionalInt32LTE applies the LTE predicate on the "optional_int32" field.
func OptionalInt32LTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt32, p.LTE(v))
	},
	)
}

// OptionalInt32IsNil applies the IsNil predicate on the "optional_int32" field.
func OptionalInt32IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt32)
	},
	)
}

// OptionalInt32NotNil applies the NotNil predicate on the "optional_int32" field.
func OptionalInt32NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt32)
	},
	)
}

// OptionalInt64EQ applies the EQ predicate on the "optional_int64" field.
func OptionalInt64EQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.EQ(v))
	},
	)
}

// OptionalInt64NEQ applies the NEQ predicate on the "optional_int64" field.
func OptionalInt64NEQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.NEQ(v))
	},
	)
}

// OptionalInt64In applies the In predicate on the "optional_int64" field.
func OptionalInt64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
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
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.Without(v...))
	},
	)
}

// OptionalInt64GT applies the GT predicate on the "optional_int64" field.
func OptionalInt64GT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.GT(v))
	},
	)
}

// OptionalInt64GTE applies the GTE predicate on the "optional_int64" field.
func OptionalInt64GTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.GTE(v))
	},
	)
}

// OptionalInt64LT applies the LT predicate on the "optional_int64" field.
func OptionalInt64LT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.LT(v))
	},
	)
}

// OptionalInt64LTE applies the LTE predicate on the "optional_int64" field.
func OptionalInt64LTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt64, p.LTE(v))
	},
	)
}

// OptionalInt64IsNil applies the IsNil predicate on the "optional_int64" field.
func OptionalInt64IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt64)
	},
	)
}

// OptionalInt64NotNil applies the NotNil predicate on the "optional_int64" field.
func OptionalInt64NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt64)
	},
	)
}

// NillableIntEQ applies the EQ predicate on the "nillable_int" field.
func NillableIntEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.EQ(v))
	},
	)
}

// NillableIntNEQ applies the NEQ predicate on the "nillable_int" field.
func NillableIntNEQ(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.NEQ(v))
	},
	)
}

// NillableIntIn applies the In predicate on the "nillable_int" field.
func NillableIntIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.Within(v...))
	},
	)
}

// NillableIntNotIn applies the NotIn predicate on the "nillable_int" field.
func NillableIntNotIn(vs ...int) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.Without(v...))
	},
	)
}

// NillableIntGT applies the GT predicate on the "nillable_int" field.
func NillableIntGT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.GT(v))
	},
	)
}

// NillableIntGTE applies the GTE predicate on the "nillable_int" field.
func NillableIntGTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.GTE(v))
	},
	)
}

// NillableIntLT applies the LT predicate on the "nillable_int" field.
func NillableIntLT(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.LT(v))
	},
	)
}

// NillableIntLTE applies the LTE predicate on the "nillable_int" field.
func NillableIntLTE(v int) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt, p.LTE(v))
	},
	)
}

// NillableIntIsNil applies the IsNil predicate on the "nillable_int" field.
func NillableIntIsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt)
	},
	)
}

// NillableIntNotNil applies the NotNil predicate on the "nillable_int" field.
func NillableIntNotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt)
	},
	)
}

// NillableInt8EQ applies the EQ predicate on the "nillable_int8" field.
func NillableInt8EQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.EQ(v))
	},
	)
}

// NillableInt8NEQ applies the NEQ predicate on the "nillable_int8" field.
func NillableInt8NEQ(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.NEQ(v))
	},
	)
}

// NillableInt8In applies the In predicate on the "nillable_int8" field.
func NillableInt8In(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.Within(v...))
	},
	)
}

// NillableInt8NotIn applies the NotIn predicate on the "nillable_int8" field.
func NillableInt8NotIn(vs ...int8) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.Without(v...))
	},
	)
}

// NillableInt8GT applies the GT predicate on the "nillable_int8" field.
func NillableInt8GT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.GT(v))
	},
	)
}

// NillableInt8GTE applies the GTE predicate on the "nillable_int8" field.
func NillableInt8GTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.GTE(v))
	},
	)
}

// NillableInt8LT applies the LT predicate on the "nillable_int8" field.
func NillableInt8LT(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.LT(v))
	},
	)
}

// NillableInt8LTE applies the LTE predicate on the "nillable_int8" field.
func NillableInt8LTE(v int8) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt8, p.LTE(v))
	},
	)
}

// NillableInt8IsNil applies the IsNil predicate on the "nillable_int8" field.
func NillableInt8IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt8)
	},
	)
}

// NillableInt8NotNil applies the NotNil predicate on the "nillable_int8" field.
func NillableInt8NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt8)
	},
	)
}

// NillableInt16EQ applies the EQ predicate on the "nillable_int16" field.
func NillableInt16EQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.EQ(v))
	},
	)
}

// NillableInt16NEQ applies the NEQ predicate on the "nillable_int16" field.
func NillableInt16NEQ(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.NEQ(v))
	},
	)
}

// NillableInt16In applies the In predicate on the "nillable_int16" field.
func NillableInt16In(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.Within(v...))
	},
	)
}

// NillableInt16NotIn applies the NotIn predicate on the "nillable_int16" field.
func NillableInt16NotIn(vs ...int16) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.Without(v...))
	},
	)
}

// NillableInt16GT applies the GT predicate on the "nillable_int16" field.
func NillableInt16GT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.GT(v))
	},
	)
}

// NillableInt16GTE applies the GTE predicate on the "nillable_int16" field.
func NillableInt16GTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.GTE(v))
	},
	)
}

// NillableInt16LT applies the LT predicate on the "nillable_int16" field.
func NillableInt16LT(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.LT(v))
	},
	)
}

// NillableInt16LTE applies the LTE predicate on the "nillable_int16" field.
func NillableInt16LTE(v int16) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt16, p.LTE(v))
	},
	)
}

// NillableInt16IsNil applies the IsNil predicate on the "nillable_int16" field.
func NillableInt16IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt16)
	},
	)
}

// NillableInt16NotNil applies the NotNil predicate on the "nillable_int16" field.
func NillableInt16NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt16)
	},
	)
}

// NillableInt32EQ applies the EQ predicate on the "nillable_int32" field.
func NillableInt32EQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.EQ(v))
	},
	)
}

// NillableInt32NEQ applies the NEQ predicate on the "nillable_int32" field.
func NillableInt32NEQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.NEQ(v))
	},
	)
}

// NillableInt32In applies the In predicate on the "nillable_int32" field.
func NillableInt32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.Within(v...))
	},
	)
}

// NillableInt32NotIn applies the NotIn predicate on the "nillable_int32" field.
func NillableInt32NotIn(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.Without(v...))
	},
	)
}

// NillableInt32GT applies the GT predicate on the "nillable_int32" field.
func NillableInt32GT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.GT(v))
	},
	)
}

// NillableInt32GTE applies the GTE predicate on the "nillable_int32" field.
func NillableInt32GTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.GTE(v))
	},
	)
}

// NillableInt32LT applies the LT predicate on the "nillable_int32" field.
func NillableInt32LT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.LT(v))
	},
	)
}

// NillableInt32LTE applies the LTE predicate on the "nillable_int32" field.
func NillableInt32LTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt32, p.LTE(v))
	},
	)
}

// NillableInt32IsNil applies the IsNil predicate on the "nillable_int32" field.
func NillableInt32IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt32)
	},
	)
}

// NillableInt32NotNil applies the NotNil predicate on the "nillable_int32" field.
func NillableInt32NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt32)
	},
	)
}

// NillableInt64EQ applies the EQ predicate on the "nillable_int64" field.
func NillableInt64EQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.EQ(v))
	},
	)
}

// NillableInt64NEQ applies the NEQ predicate on the "nillable_int64" field.
func NillableInt64NEQ(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.NEQ(v))
	},
	)
}

// NillableInt64In applies the In predicate on the "nillable_int64" field.
func NillableInt64In(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.Within(v...))
	},
	)
}

// NillableInt64NotIn applies the NotIn predicate on the "nillable_int64" field.
func NillableInt64NotIn(vs ...int64) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.Without(v...))
	},
	)
}

// NillableInt64GT applies the GT predicate on the "nillable_int64" field.
func NillableInt64GT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.GT(v))
	},
	)
}

// NillableInt64GTE applies the GTE predicate on the "nillable_int64" field.
func NillableInt64GTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.GTE(v))
	},
	)
}

// NillableInt64LT applies the LT predicate on the "nillable_int64" field.
func NillableInt64LT(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.LT(v))
	},
	)
}

// NillableInt64LTE applies the LTE predicate on the "nillable_int64" field.
func NillableInt64LTE(v int64) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldNillableInt64, p.LTE(v))
	},
	)
}

// NillableInt64IsNil applies the IsNil predicate on the "nillable_int64" field.
func NillableInt64IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNillableInt64)
	},
	)
}

// NillableInt64NotNil applies the NotNil predicate on the "nillable_int64" field.
func NillableInt64NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNillableInt64)
	},
	)
}

// ValidateOptionalInt32EQ applies the EQ predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32EQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.EQ(v))
	},
	)
}

// ValidateOptionalInt32NEQ applies the NEQ predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32NEQ(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.NEQ(v))
	},
	)
}

// ValidateOptionalInt32In applies the In predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32In(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.Within(v...))
	},
	)
}

// ValidateOptionalInt32NotIn applies the NotIn predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32NotIn(vs ...int32) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.Without(v...))
	},
	)
}

// ValidateOptionalInt32GT applies the GT predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32GT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.GT(v))
	},
	)
}

// ValidateOptionalInt32GTE applies the GTE predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32GTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.GTE(v))
	},
	)
}

// ValidateOptionalInt32LT applies the LT predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32LT(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.LT(v))
	},
	)
}

// ValidateOptionalInt32LTE applies the LTE predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32LTE(v int32) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldValidateOptionalInt32, p.LTE(v))
	},
	)
}

// ValidateOptionalInt32IsNil applies the IsNil predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32IsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldValidateOptionalInt32)
	},
	)
}

// ValidateOptionalInt32NotNil applies the NotNil predicate on the "validate_optional_int32" field.
func ValidateOptionalInt32NotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldValidateOptionalInt32)
	},
	)
}

// StateEQ applies the EQ predicate on the "state" field.
func StateEQ(v State) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldState, p.EQ(v))
	},
	)
}

// StateNEQ applies the NEQ predicate on the "state" field.
func StateNEQ(v State) predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldState, p.NEQ(v))
	},
	)
}

// StateIn applies the In predicate on the "state" field.
func StateIn(vs ...State) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldState, p.Within(v...))
	},
	)
}

// StateNotIn applies the NotIn predicate on the "state" field.
func StateNotIn(vs ...State) predicate.FieldType {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.Has(Label, FieldState, p.Without(v...))
	},
	)
}

// StateIsNil applies the IsNil predicate on the "state" field.
func StateIsNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldState)
	},
	)
}

// StateNotNil applies the NotNil predicate on the "state" field.
func StateNotNil() predicate.FieldType {
	return predicate.FieldType(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldState)
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.FieldType) predicate.FieldType {
	return predicate.FieldType(
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
	return predicate.FieldType(
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
	return predicate.FieldType(
		func(tr *dsl.Traversal) {
			t := __.New()
			p(t)
			tr.Where(__.Not(t))
		},
	)
}
