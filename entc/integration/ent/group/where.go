// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"strconv"
	"time"

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

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldActive), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.EQ(v))
		},
	}
}

// Expire applies equality check predicate on the "expire" field. It's identical to ExpireEQ.
func Expire(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.EQ(v))
		},
	}
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EQ(v))
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

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	}
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldActive), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.EQ(v))
		},
	}
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldActive), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.NEQ(v))
		},
	}
}

// ExpireEQ applies the EQ predicate on the "expire" field.
func ExpireEQ(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.EQ(v))
		},
	}
}

// ExpireNEQ applies the NEQ predicate on the "expire" field.
func ExpireNEQ(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.NEQ(v))
		},
	}
}

// ExpireGT applies the GT predicate on the "expire" field.
func ExpireGT(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.GT(v))
		},
	}
}

// ExpireGTE applies the GTE predicate on the "expire" field.
func ExpireGTE(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.GTE(v))
		},
	}
}

// ExpireLT applies the LT predicate on the "expire" field.
func ExpireLT(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.LT(v))
		},
	}
}

// ExpireLTE applies the LTE predicate on the "expire" field.
func ExpireLTE(v time.Time) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExpire), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.LTE(v))
		},
	}
}

// ExpireIn applies the In predicate on the "expire" field.
func ExpireIn(vs ...time.Time) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldExpire), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.Within(v...))
		},
	}
}

// ExpireNotIn applies the NotIn predicate on the "expire" field.
func ExpireNotIn(vs ...time.Time) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldExpire), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.Without(v...))
		},
	}
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EQ(v))
		},
	}
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.NEQ(v))
		},
	}
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.GT(v))
		},
	}
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.GTE(v))
		},
	}
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.LT(v))
		},
	}
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.LTE(v))
		},
	}
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldType), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Within(v...))
		},
	}
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldType), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Without(v...))
		},
	}
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Containing(v))
		},
	}
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.StartingWith(v))
		},
	}
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldType), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EndingWith(v))
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

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	}
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	}
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	}
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	}
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	}
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	}
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldName), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Within(v...))
		},
	}
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Without(v...))
		},
	}
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	}
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	}
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	}
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(FilesColumn).
						From(sql.Table(FilesTable)).
						Where(sql.NotNull(FilesColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(FilesLabel).OutV()
		},
	}
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FilesColumn).From(sql.Table(FilesTable))
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
			t.OutE(FilesLabel).Where(tr).OutV()
		},
	}
}

// HasBlocked applies the HasEdge predicate on the "blocked" edge.
func HasBlocked() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(BlockedColumn).
						From(sql.Table(BlockedTable)).
						Where(sql.NotNull(BlockedColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(BlockedLabel).OutV()
		},
	}
}

// HasBlockedWith applies the HasEdge predicate on the "blocked" edge with a given conditions (other predicates).
func HasBlockedWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(BlockedColumn).From(sql.Table(BlockedTable))
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
			t.OutE(BlockedLabel).Where(tr).OutV()
		},
	}
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(UsersPrimaryKey[1]).From(sql.Table(UsersTable)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(UsersInverseLabel).InV()
		},
	}
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Table(UsersInverseTable)
			t3 := sql.Table(UsersTable)
			t4 := sql.Select(t3.C(UsersPrimaryKey[1])).
				From(t3).
				Join(t2).
				On(t3.C(UsersPrimaryKey[0]), t2.C(FieldID))
			t5 := sql.Select().From(t2)
			for _, p := range preds {
				p.SQL(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.InE(UsersInverseLabel).Where(tr).InV()
		},
	}
}

// HasInfo applies the HasEdge predicate on the "info" edge.
func HasInfo() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(InfoColumn)))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(InfoLabel).OutV()
		},
	}
}

// HasInfoWith applies the HasEdge predicate on the "info" edge with a given conditions (other predicates).
func HasInfoWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(InfoInverseTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(InfoColumn), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.OutE(InfoLabel).Where(tr).OutV()
		},
	}
}
