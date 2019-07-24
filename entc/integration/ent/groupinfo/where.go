// Code generated (@generated) by entc, DO NOT EDIT.

package groupinfo

import (
	"strconv"

	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/__"
	"fbc/ent/dialect/gremlin/graph/dsl/p"
	"fbc/ent/dialect/sql"
)

// ID filters vertices based on their identifier.
func ID(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDNEQ(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDGT(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDGTE(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDLT(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDLTE(id string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func IDNotIn(ids ...string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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

// Desc applies equality check predicate on the "desc" field. It's identical to DescEQ.
func Desc(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EQ(v))
		},
	)
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	)
}

// DescEQ applies the EQ predicate on the "desc" field.
func DescEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EQ(v))
		},
	)
}

// DescNEQ applies the NEQ predicate on the "desc" field.
func DescNEQ(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.NEQ(v))
		},
	)
}

// DescGT applies the GT predicate on the "desc" field.
func DescGT(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.GT(v))
		},
	)
}

// DescGTE applies the GTE predicate on the "desc" field.
func DescGTE(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.GTE(v))
		},
	)
}

// DescLT applies the LT predicate on the "desc" field.
func DescLT(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.LT(v))
		},
	)
}

// DescLTE applies the LTE predicate on the "desc" field.
func DescLTE(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.LTE(v))
		},
	)
}

// DescIn applies the In predicate on the "desc" field.
func DescIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldDesc), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Within(v...))
		},
	)
}

// DescNotIn applies the NotIn predicate on the "desc" field.
func DescNotIn(vs ...string) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldDesc), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Without(v...))
		},
	)
}

// DescContains applies the Contains predicate on the "desc" field.
func DescContains(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.Containing(v))
		},
	)
}

// DescHasPrefix applies the HasPrefix predicate on the "desc" field.
func DescHasPrefix(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.StartingWith(v))
		},
	)
}

// DescHasSuffix applies the HasSuffix predicate on the "desc" field.
func DescHasSuffix(v string) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldDesc), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldDesc, p.EndingWith(v))
		},
	)
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	)
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.NEQ(v))
		},
	)
}

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GT(v))
		},
	)
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GTE(v))
		},
	)
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LT(v))
		},
	)
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LTE(v))
		},
	)
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldMaxUsers), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.Within(v...))
		},
	)
}

// MaxUsersNotIn applies the NotIn predicate on the "max_users" field.
func MaxUsersNotIn(vs ...int) predicate.GroupInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldMaxUsers), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.Without(v...))
		},
	)
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
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
		func(t *dsl.Traversal) {
			t.InE(GroupsInverseLabel).InV()
		},
	)
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(GroupsColumn).From(sql.Table(GroupsTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(GroupsInverseLabel).Where(tr).InV()
		},
	)
}

// Or groups list of predicates with the or operator between them.
func Or(predicates ...predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
func Not(p predicate.GroupInfo) predicate.GroupInfo {
	return predicate.GroupInfoPerDialect(
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
