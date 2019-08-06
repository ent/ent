// Code generated (@generated) by entc, DO NOT EDIT.

package file

import (
	"strconv"

	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDEQ(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDNEQ(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDGT(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDGTE(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDLT(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDLTE(id string) predicate.File {
	return predicate.FilePerDialect(
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
func IDIn(ids ...string) predicate.File {
	return predicate.FilePerDialect(
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
func IDNotIn(ids ...string) predicate.File {
	return predicate.FilePerDialect(
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

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.EQ(v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.EQ(v))
		},
	)
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.NEQ(v))
		},
	)
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.GT(v))
		},
	)
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.GTE(v))
		},
	)
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.LT(v))
		},
	)
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v int) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldSize), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.LTE(v))
		},
	)
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldSize), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.Within(v...))
		},
	)
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...int) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldSize), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldSize, p.Without(v...))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
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
func NameNotIn(vs ...string) predicate.File {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.FilePerDialect(
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
func NameContains(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.FilePerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.FilePerDialect(
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
func Or(predicates ...predicate.File) predicate.File {
	return predicate.FilePerDialect(
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
func Not(p predicate.File) predicate.File {
	return predicate.FilePerDialect(
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
