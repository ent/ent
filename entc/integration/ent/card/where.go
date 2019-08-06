// Code generated (@generated) by entc, DO NOT EDIT.

package card

import (
	"strconv"

	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDEQ(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDNEQ(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDGT(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDGTE(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDLT(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDLTE(id string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDIn(ids ...string) predicate.Card {
	return predicate.CardPerDialect(
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
func IDNotIn(ids ...string) predicate.Card {
	return predicate.CardPerDialect(
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

// Number applies equality check predicate on the "number" field. It's identical to NumberEQ.
func Number(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	)
}

// NumberEQ applies the EQ predicate on the "number" field.
func NumberEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EQ(v))
		},
	)
}

// NumberNEQ applies the NEQ predicate on the "number" field.
func NumberNEQ(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.NEQ(v))
		},
	)
}

// NumberGT applies the GT predicate on the "number" field.
func NumberGT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GT(v))
		},
	)
}

// NumberGTE applies the GTE predicate on the "number" field.
func NumberGTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.GTE(v))
		},
	)
}

// NumberLT applies the LT predicate on the "number" field.
func NumberLT(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LT(v))
		},
	)
}

// NumberLTE applies the LTE predicate on the "number" field.
func NumberLTE(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.LTE(v))
		},
	)
}

// NumberIn applies the In predicate on the "number" field.
func NumberIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNumber), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Within(v...))
		},
	)
}

// NumberNotIn applies the NotIn predicate on the "number" field.
func NumberNotIn(vs ...string) predicate.Card {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNumber), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Without(v...))
		},
	)
}

// NumberContains applies the Contains predicate on the "number" field.
func NumberContains(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.Containing(v))
		},
	)
}

// NumberHasPrefix applies the HasPrefix predicate on the "number" field.
func NumberHasPrefix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.StartingWith(v))
		},
	)
}

// NumberHasSuffix applies the HasSuffix predicate on the "number" field.
func NumberHasSuffix(v string) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNumber), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldNumber, p.EndingWith(v))
		},
	)
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(OwnerColumn)))
		},
		func(t *dsl.Traversal) {
			t.InE(OwnerInverseLabel).InV()
		},
	)
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Card {
	return predicate.CardPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(OwnerInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(OwnerColumn), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(OwnerInverseLabel).Where(tr).InV()
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
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
func Or(predicates ...predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
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
func Not(p predicate.Card) predicate.Card {
	return predicate.CardPerDialect(
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
