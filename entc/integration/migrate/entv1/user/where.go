// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"strconv"

	"fbc/ent/entc/integration/migrate/entv1/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.User {
	return predicate.UserPerDialect(
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
func IDEQ(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.EQ(id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GT(id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.GTE(id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LT(id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), v))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.UserPerDialect(
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
func IDNotIn(ids ...string) predicate.User {
	return predicate.UserPerDialect(
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

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.EQ(v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.EQ(v))
		},
	)
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.EQ(v))
		},
	)
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.NEQ(v))
		},
	)
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.GT(v))
		},
	)
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.GTE(v))
		},
	)
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.LT(v))
		},
	)
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int32) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAge), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.LTE(v))
		},
	)
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int32) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldAge), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.Within(v...))
		},
	)
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int32) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldAge), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.Without(v...))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Within(v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Without(v...))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	)
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.EQ(v))
		},
	)
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.NEQ(v))
		},
	)
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.GT(v))
		},
	)
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.GTE(v))
		},
	)
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.LT(v))
		},
	)
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.LTE(v))
		},
	)
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldAddress), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.Within(v...))
		},
	)
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldAddress), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.Without(v...))
		},
	)
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.Containing(v))
		},
	)
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.StartingWith(v))
		},
	)
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.UserPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldAddress), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldAddress, p.EndingWith(v))
		},
	)
}

// Or groups list of predicates with the or operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.UserPerDialect(
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
func Not(p predicate.User) predicate.User {
	return predicate.UserPerDialect(
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
