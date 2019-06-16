// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"strconv"

	"fbc/ent"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/__"
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

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	}
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	}
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.NEQ(v))
		},
	}
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GT(v))
		},
	}
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GTE(v))
		},
	}
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LT(v))
		},
	}
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LTE(v))
		},
	}
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldNumber), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Within(v...))
		},
	}
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldNumber), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Without(v...))
		},
	}
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Containing(v))
		},
	}
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.StartingWith(v))
		},
	}
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNumber), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EndingWith(v))
		},
	}
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(OwnerColumn)))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(OwnerInverseLabel).InV()
		},
	}
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(OwnerInverseTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(OwnerColumn), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.InE(OwnerInverseLabel).Where(tr).InV()
		},
	}
}
