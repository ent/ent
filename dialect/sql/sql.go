// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package sql

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect"
)

// The following helpers exist to simplify the way raw predicates
// are defined and used in both ent/schema and generated code. For
// full predicates API, check out the sql.P in builder.go.

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

// FieldsGT returns a raw predicate to check if field1 is greater than field2.
func FieldsGT(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsGT(s.C(field1), s.C(field2)))
	}
}

// FieldGTE returns a raw predicate to check if the given field is greater than or equal the given value.
func FieldGTE(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(GTE(s.C(name), v))
	}
}

// FieldsGTE returns a raw predicate to check if field1 is greater than or equal field2.
func FieldsGTE(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsGTE(s.C(field1), s.C(field2)))
	}
}

// FieldLT returns a raw predicate to check if the value of the field is less than the given value.
func FieldLT(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(LT(s.C(name), v))
	}
}

// FieldsLT returns a raw predicate to check if field1 is lower than field2.
func FieldsLT(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsLT(s.C(field1), s.C(field2)))
	}
}

// FieldLTE returns a raw predicate to check if the value of the field is less than the given value.
func FieldLTE(name string, v any) func(*Selector) {
	return func(s *Selector) {
		s.Where(LTE(s.C(name), v))
	}
}

// FieldsLTE returns a raw predicate to check if field1 is lower than or equal field2.
func FieldsLTE(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsLTE(s.C(field1), s.C(field2)))
	}
}

// FieldsHasPrefix returns a raw predicate to checks if field1 begins with the value of field2.
func FieldsHasPrefix(field1, field2 string) func(*Selector) {
	return func(s *Selector) {
		s.Where(ColumnsHasPrefix(s.C(field1), s.C(field2)))
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

// FieldEqualFold returns a raw predicate to check if the field is equal to the given string under case-folding.
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

// FieldHasPrefixFold returns a raw predicate to check if the field has the given prefix with case-folding
func FieldHasPrefixFold(name string, prefix string) func(*Selector) {
	return func(s *Selector) {
		s.Where(HasPrefixFold(s.C(name), prefix))
	}
}

// FieldHasSuffix returns a raw predicate to check if the field has the given suffix.
func FieldHasSuffix(name string, suffix string) func(*Selector) {
	return func(s *Selector) {
		s.Where(HasSuffix(s.C(name), suffix))
	}
}

// FieldHasSuffixFold returns a raw predicate to check if the field has the given suffix with case-folding
func FieldHasSuffixFold(name string, suffix string) func(*Selector) {
	return func(s *Selector) {
		s.Where(HasSuffixFold(s.C(name), suffix))
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

// AndPredicates returns a new predicate for joining multiple generated predicates with AND between them.
func AndPredicates[P ~func(*Selector)](predicates ...P) func(*Selector) {
	return func(s *Selector) {
		s.CollectPredicates()
		for _, p := range predicates {
			p(s)
		}
		collected := s.CollectedPredicates()
		s.UncollectedPredicates()
		switch len(collected) {
		case 0:
		case 1:
			s.Where(collected[0])
		default:
			s.Where(And(collected...))
		}
	}
}

// OrPredicates returns a new predicate for joining multiple generated predicates with OR between them.
func OrPredicates[P ~func(*Selector)](predicates ...P) func(*Selector) {
	return func(s *Selector) {
		s.CollectPredicates()
		for _, p := range predicates {
			p(s)
		}
		collected := s.CollectedPredicates()
		s.UncollectedPredicates()
		switch len(collected) {
		case 0:
		case 1:
			s.Where(collected[0])
		default:
			s.Where(Or(collected...))
		}
	}
}

// NotPredicates wraps the generated predicates with NOT. For example, NOT(P), NOT((P1 AND P2)).
func NotPredicates[P ~func(*Selector)](predicates ...P) func(*Selector) {
	return func(s *Selector) {
		s.CollectPredicates()
		for _, p := range predicates {
			p(s)
		}
		collected := s.CollectedPredicates()
		s.UncollectedPredicates()
		switch len(collected) {
		case 0:
		case 1:
			s.Where(Not(collected[0]))
		default:
			s.Where(Not(And(collected...)))
		}
	}
}

// ColumnCheck is a function that verifies whether the
// specified column exists within the given table.
type ColumnCheck func(table, column string) error

// NewColumnCheck returns a function that verifies whether the specified column exists
// within the given table. This function is utilized by the generated code to validate
// column names in ordering functions.
func NewColumnCheck(checks map[string]func(string) bool) ColumnCheck {
	return func(table, column string) error {
		check, ok := checks[table]
		if !ok {
			return fmt.Errorf("unknown table %q", table)
		}
		if !check(column) {
			return fmt.Errorf("unknown column %q for table %q", column, table)
		}
		return nil
	}
}

type (
	// OrderFieldTerm represents an ordering by a field.
	OrderFieldTerm struct {
		OrderTermOptions
		Field string // Field name.
	}
	// OrderExprTerm represents an ordering by an expression.
	OrderExprTerm struct {
		OrderTermOptions
		Expr func(*Selector) Querier // Expression.
	}
	// OrderTerm represents an ordering by a term.
	OrderTerm interface {
		term()
	}
	// OrderTermOptions represents options for ordering by a term.
	OrderTermOptions struct {
		Desc       bool   // Whether to sort in descending order.
		As         string // Optional alias.
		Selected   bool   // Whether the term should be selected.
		NullsFirst bool   // Whether to sort nulls first.
		NullsLast  bool   // Whether to sort nulls last.
	}
	// OrderTermOption is an option for ordering by a term.
	OrderTermOption func(*OrderTermOptions)
)

// OrderDesc returns an option to sort in descending order.
func OrderDesc() OrderTermOption {
	return func(o *OrderTermOptions) {
		o.Desc = true
	}
}

// OrderAsc returns an option to sort in ascending order.
func OrderAsc() OrderTermOption {
	return func(o *OrderTermOptions) {
		o.Desc = false
	}
}

// OrderAs returns an option to set the alias for the ordering.
func OrderAs(as string) OrderTermOption {
	return func(o *OrderTermOptions) {
		o.As = as
	}
}

// OrderSelected returns an option to select the ordering term.
func OrderSelected() OrderTermOption {
	return func(o *OrderTermOptions) {
		o.Selected = true
	}
}

// OrderSelectAs returns an option to set and select the alias for the ordering.
func OrderSelectAs(as string) OrderTermOption {
	return func(o *OrderTermOptions) {
		o.As = as
		o.Selected = true
	}
}

// OrderNullsFirst returns an option to sort nulls first.
func OrderNullsFirst() OrderTermOption {
	return func(o *OrderTermOptions) {
		o.NullsFirst = true
	}
}

// OrderNullsLast returns an option to sort nulls last.
func OrderNullsLast() OrderTermOption {
	return func(o *OrderTermOptions) {
		o.NullsLast = true
	}
}

// NewOrderTermOptions returns a new OrderTermOptions from the given options.
func NewOrderTermOptions(opts ...OrderTermOption) *OrderTermOptions {
	o := &OrderTermOptions{}
	for _, opt := range opts {
		opt(o)
	}
	return o
}

// OrderByField returns an ordering by the given field.
func OrderByField(field string, opts ...OrderTermOption) *OrderFieldTerm {
	return &OrderFieldTerm{Field: field, OrderTermOptions: *NewOrderTermOptions(opts...)}
}

// OrderBySum returns an ordering by the sum of the given field.
func OrderBySum(field string, opts ...OrderTermOption) *OrderExprTerm {
	return orderByAgg("SUM", field, opts...)
}

// OrderByCount returns an ordering by the count of the given field.
func OrderByCount(field string, opts ...OrderTermOption) *OrderExprTerm {
	return orderByAgg("COUNT", field, opts...)
}

// orderByAgg returns an ordering by the aggregation of the given field.
func orderByAgg(fn, field string, opts ...OrderTermOption) *OrderExprTerm {
	return &OrderExprTerm{
		OrderTermOptions: *NewOrderTermOptions(
			append(
				// Default alias is "<func>_<field>".
				[]OrderTermOption{OrderAs(fmt.Sprintf("%s_%s", strings.ToLower(fn), field))},
				opts...,
			)...,
		),
		Expr: func(s *Selector) Querier {
			var c string
			switch {
			case field == "*", isFunc(field):
				c = field
			default:
				c = s.C(field)
			}
			return Raw(fmt.Sprintf("%s(%s)", fn, c))
		},
	}
}

// OrderByRand returns a term to natively order by a random value.
func OrderByRand() func(*Selector) {
	return func(s *Selector) {
		s.OrderExprFunc(func(b *Builder) {
			switch s.Dialect() {
			case dialect.MySQL:
				b.WriteString("RAND()")
			default:
				b.WriteString("RANDOM()")
			}
		})
	}
}

// ToFunc returns a function that sets the ordering on the given selector.
// This is used by the generated code.
func (f *OrderFieldTerm) ToFunc() func(*Selector) {
	return func(s *Selector) {
		s.OrderExprFunc(func(b *Builder) {
			b.WriteString(s.C(f.Field))
			if f.Desc {
				b.WriteString(" DESC")
			}
			if f.NullsFirst {
				b.WriteString(" NULLS FIRST")
			} else if f.NullsLast {
				b.WriteString(" NULLS LAST")
			}
		})
	}
}

func (OrderFieldTerm) term() {}
func (OrderExprTerm) term()  {}
