// Code generated (@generated) by entc, DO NOT EDIT.

package node

import (
	"strconv"

	"fbc/ent"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
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

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.EQ(v))
		},
	}
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.EQ(v))
		},
	}
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.NEQ(v))
		},
	}
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.GT(v))
		},
	}
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.GTE(v))
		},
	}
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.LT(v))
		},
	}
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldValue), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.LTE(v))
		},
	}
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...int) ent.Predicate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldValue), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.Within(v...))
		},
	}
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...int) ent.Predicate {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldValue), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldValue, p.Without(v...))
		},
	}
}

// HasPrev applies the HasEdge predicate on the "prev" edge.
func HasPrev() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(PrevColumn)))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(PrevInverseLabel).InV()
		},
	}
}

// HasPrevWith applies the HasEdge predicate on the "prev" edge with a given conditions (other predicates).
func HasPrevWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(PrevTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(PrevColumn), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.InE(PrevInverseLabel).Where(tr).InV()
		},
	}
}

// HasNext applies the HasEdge predicate on the "next" edge.
func HasNext() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(NextColumn).
						From(sql.Table(NextTable)).
						Where(sql.NotNull(NextColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(NextLabel).OutV()
		},
	}
}

// HasNextWith applies the HasEdge predicate on the "next" edge with a given conditions (other predicates).
func HasNextWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(NextColumn).From(sql.Table(NextTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.OutE(NextLabel).Where(tr).OutV()
		},
	}
}
