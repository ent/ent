// Code generated (@generated) by entc, DO NOT EDIT.

package groupinfo

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

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EQ(v))
		},
	}
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	}
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EQ(v))
		},
	}
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.NEQ(v))
		},
	}
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.GT(v))
		},
	}
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.GTE(v))
		},
	}
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.LT(v))
		},
	}
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.LTE(v))
		},
	}
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldDesc), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Within(v...))
		},
	}
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldDesc), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Without(v...))
		},
	}
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Containing(v))
		},
	}
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.StartingWith(v))
		},
	}
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldDesc), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EndingWith(v))
		},
	}
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	}
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.NEQ(v))
		},
	}
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GT(v))
		},
	}
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GTE(v))
		},
	}
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LT(v))
		},
	}
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldMaxUsers), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LTE(v))
		},
	}
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldMaxUsers), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.Within(v...))
		},
	}
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldMaxUsers), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.Without(v...))
		},
	}
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(GroupsColumn).
						From(sql.Table(GroupsTable)).
						Where(sql.NotNull(GroupsColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(GroupsInverseLabel).InV()
		},
	}
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(GroupsColumn).From(sql.Table(GroupsTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.InE(GroupsInverseLabel).Where(tr).InV()
		},
	}
}
