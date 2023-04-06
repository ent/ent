// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

// This file provides extra helpers to simplify the way raw predicates
// are defined and used in both ent/schema and generated code. For full
// predicates, check out the sql.P in builder.go.

// FieldIsNull returns a raw predicate to check if the given field is NULL.
func FieldIsNull(name string) func(*Selector) {
	return func(s *Selector) {
		s.Where(IsNull(s.C(name)))
	}
}

// FieldNotNull returns a raw predicate to check if the given field is not NULL.
func FieldNotNull(name string) func(*Selector) {
	return func(s *Selector) {
		s.Where(NotNull(s.C(name)))
	}
}

// FieldEQ returns a raw predicate to check if the given field equals to the given value.
func FieldEQ(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(EQ(s.C(name), v))
	}
}

// FieldsEQ returns a raw predicate to check if the given fields (columns) are equal.
func FieldsEQ(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsEQ(s.C(field1), s.C(field2)))
	}
}

// FieldNEQ returns a raw predicate to check if the given field does not equal to the given value.
func FieldNEQ(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(NEQ(s.C(name), v))
	}
}

// FieldsNEQ returns a raw predicate to check if the given fields (columns) are not equal.
func FieldsNEQ(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsNEQ(s.C(field1), s.C(field2)))
	}
}

// FieldGT returns a raw predicate to check if the given field is greater than the given value.
func FieldGT(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(GT(s.C(name), v))
	}
}

// FieldGTE returns a raw predicate to check if the given field is greater than or equal the given value.
func FieldGTE(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(GTE(s.C(name), v))
	}
}

// FieldLT returns a raw predicate to check if the value of the field is less than the given value.
func FieldLT(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(LT(s.C(name), v))
	}
}

// FieldLTE returns a raw predicate to check if the value of the field is less than the given value.
func FieldLTE(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(LTE(s.C(name), v))
	}
}

// FieldIn returns a raw predicate to check if the value of the field is IN the given values.
func FieldIn[T any](name string, vs ...T) func(*Selector) {
	return func(s *Selector) {
		v := make([]any, len(vs))
		for i := range v {
			v[i] = vs[i]
		}
		s.Where(In(s.C(name), v...))
	}
}

// FieldNotIn returns a raw predicate to check if the value of the field is NOT IN the given values.
func FieldNotIn[T any](name string, vs ...T) func(*Selector) {
	return func(s *Selector) {
		v := make([]any, len(vs))
		for i := range v {
			v[i] = vs[i]
		}
		s.Where(NotIn(s.C(name), v...))
	}
}

// FieldEqualFold returns a raw predicate to check if the field has the given prefix with case-folding.
func FieldEqualFold(name string, substr string) func(*Selector) {
	return func(s *Selector) {
		s.Where(EqualFold(s.C(name), substr))
	}
}

// FieldHasPrefix returns a raw predicate to check if the field has the given prefix.
func FieldHasPrefix(name string, prefix string) func(*Selector) {
	return func(s *Selector) {
		s.Where(HasPrefix(s.C(name), prefix))
	}
}

// FieldHasSuffix returns a raw predicate to check if the field has the given suffix.
func FieldHasSuffix(name string, suffix string) func(*Selector) {
	return func(s *Selector) {
		s.Where(HasSuffix(s.C(name), suffix))
	}
}

// FieldContains returns a raw predicate to check if the field contains the given substring.
func FieldContains(name string, substr string) func(*Selector) {
	return func(s *Selector) {
		s.Where(Contains(s.C(name), substr))
	}
}

// FieldContainsFold returns a raw predicate to check if the field contains the given substring with case-folding.
func FieldContainsFold(name string, substr string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ContainsFold(s.C(name), substr))
	}
}
