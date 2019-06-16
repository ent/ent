// Code generated (@generated) by entc, DO NOT EDIT.

package boring

import (
	"strconv"

	"fbc/ent"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/p"
)

// ID filters vertices based on their identifier.
func ID(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(id)
		},
	}
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.EQ(id))
		},
	}
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
		},
	}
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.GT(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.GT(id))
		},
	}
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.GTE(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.GTE(id))
		},
	}
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.LT(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.LT(id))
		},
	}
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			v, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	}
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
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
		Gremlin: func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Within(v...))
		},
	}
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
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
		Gremlin: func(t *dsl.Traversal) {
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			t.HasID(p.Without(v...))
		},
	}
}
