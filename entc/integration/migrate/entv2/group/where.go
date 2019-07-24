// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"strconv"

	"fbc/ent/entc/integration/migrate/entv2/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDEQ(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDNEQ(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDGT(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDGTE(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDLT(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDLTE(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDIn(ids ...string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDNotIn(ids ...string) predicate.Group {
	return predicate.GroupPerDialect(
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

// Or groups list of predicates with the or operator between them.
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.GroupPerDialect(
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
func Not(p predicate.Group) predicate.Group {
	return predicate.GroupPerDialect(
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
