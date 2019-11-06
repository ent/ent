// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package group

import (
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
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
			id, _ := strconv.Atoi(id)
			s.Where(sql.EQ(s.C(FieldID), id))
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
			id, _ := strconv.Atoi(id)
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.NEQ(id))
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

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDGTE(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDLT(id string) predicate.Group {
	return predicate.GroupPerDialect(
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
func IDLTE(id string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			id, _ := strconv.Atoi(id)
			s.Where(sql.LTE(s.C(FieldID), id))
		},
		func(t *dsl.Traversal) {
			t.HasID(p.LTE(id))
		},
	)
}

// Active applies equality check predicate on the "active" field. It's identical to ActiveEQ.
func Active(v bool) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldActive), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.EQ(v))
		},
	)
}

// Expire applies equality check predicate on the "expire" field. It's identical to ExpireEQ.
func Expire(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.EQ(v))
		},
	)
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EQ(v))
		},
	)
}

// MaxUsers applies equality check predicate on the "max_users" field. It's identical to MaxUsersEQ.
func MaxUsers(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// ActiveEQ applies the EQ predicate on the "active" field.
func ActiveEQ(v bool) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldActive), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.EQ(v))
		},
	)
}

// ActiveNEQ applies the NEQ predicate on the "active" field.
func ActiveNEQ(v bool) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldActive), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldActive, p.NEQ(v))
		},
	)
}

// ExpireEQ applies the EQ predicate on the "expire" field.
func ExpireEQ(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.EQ(v))
		},
	)
}

// ExpireNEQ applies the NEQ predicate on the "expire" field.
func ExpireNEQ(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.NEQ(v))
		},
	)
}

// ExpireIn applies the In predicate on the "expire" field.
func ExpireIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldExpire), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.Within(v...))
		},
	)
}

// ExpireNotIn applies the NotIn predicate on the "expire" field.
func ExpireNotIn(vs ...time.Time) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldExpire), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.Without(v...))
		},
	)
}

// ExpireGT applies the GT predicate on the "expire" field.
func ExpireGT(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.GT(v))
		},
	)
}

// ExpireGTE applies the GTE predicate on the "expire" field.
func ExpireGTE(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.GTE(v))
		},
	)
}

// ExpireLT applies the LT predicate on the "expire" field.
func ExpireLT(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.LT(v))
		},
	)
}

// ExpireLTE applies the LTE predicate on the "expire" field.
func ExpireLTE(v time.Time) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldExpire), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldExpire, p.LTE(v))
		},
	)
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EQ(v))
		},
	)
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.NEQ(v))
		},
	)
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldType), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Within(v...))
		},
	)
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldType), v...))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Without(v...))
		},
	)
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.GT(v))
		},
	)
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.GTE(v))
		},
	)
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.LT(v))
		},
	)
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.LTE(v))
		},
	)
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.Containing(v))
		},
	)
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.StartingWith(v))
		},
	)
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldType), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldType, p.EndingWith(v))
		},
	)
}

// TypeIsNil applies the IsNil predicate on the "type" field.
func TypeIsNil() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldType)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldType)
		},
	)
}

// TypeNotNil applies the NotNil predicate on the "type" field.
func TypeNotNil() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldType)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldType)
		},
	)
}

// MaxUsersEQ applies the EQ predicate on the "max_users" field.
func MaxUsersEQ(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.EQ(v))
		},
	)
}

// MaxUsersNEQ applies the NEQ predicate on the "max_users" field.
func MaxUsersNEQ(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.NEQ(v))
		},
	)
}

// MaxUsersIn applies the In predicate on the "max_users" field.
func MaxUsersIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
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
func MaxUsersNotIn(vs ...int) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
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

// MaxUsersGT applies the GT predicate on the "max_users" field.
func MaxUsersGT(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GT(v))
		},
	)
}

// MaxUsersGTE applies the GTE predicate on the "max_users" field.
func MaxUsersGTE(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.GTE(v))
		},
	)
}

// MaxUsersLT applies the LT predicate on the "max_users" field.
func MaxUsersLT(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LT(v))
		},
	)
}

// MaxUsersLTE applies the LTE predicate on the "max_users" field.
func MaxUsersLTE(v int) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldMaxUsers), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldMaxUsers, p.LTE(v))
		},
	)
}

// MaxUsersIsNil applies the IsNil predicate on the "max_users" field.
func MaxUsersIsNil() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldMaxUsers)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).HasNot(FieldMaxUsers)
		},
	)
}

// MaxUsersNotNil applies the NotNil predicate on the "max_users" field.
func MaxUsersNotNil() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldMaxUsers)))
		},
		func(t *dsl.Traversal) {
			t.HasLabel(Label).Has(FieldMaxUsers)
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EQ(v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.NEQ(v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
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
func NameNotIn(vs ...string) predicate.Group {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.GroupPerDialect(
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

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GT(v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.GTE(v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LT(v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.LTE(v))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.Containing(v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.StartingWith(v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
		func(t *dsl.Traversal) {
			t.Has(Label, FieldName, p.EndingWith(v))
		},
	)
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			builder := sql.Dialect(s.Dialect())
			s.Where(
				sql.In(
					t1.C(FieldID),
					builder.Select(FilesColumn).
						From(builder.Table(FilesTable)).
						Where(sql.NotNull(FilesColumn)),
				),
			)
		},
		func(t *dsl.Traversal) {
			t.OutE(FilesLabel).OutV()
		},
	)
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FilesColumn).From(builder.Table(FilesTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p(tr)
			}
			t.OutE(FilesLabel).Where(tr).OutV()
		},
	)
}

// HasBlocked applies the HasEdge predicate on the "blocked" edge.
func HasBlocked() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			builder := sql.Dialect(s.Dialect())
			s.Where(
				sql.In(
					t1.C(FieldID),
					builder.Select(BlockedColumn).
						From(builder.Table(BlockedTable)).
						Where(sql.NotNull(BlockedColumn)),
				),
			)
		},
		func(t *dsl.Traversal) {
			t.OutE(BlockedLabel).OutV()
		},
	)
}

// HasBlockedWith applies the HasEdge predicate on the "blocked" edge with a given conditions (other predicates).
func HasBlockedWith(preds ...predicate.User) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(BlockedColumn).From(builder.Table(BlockedTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p(tr)
			}
			t.OutE(BlockedLabel).Where(tr).OutV()
		},
	)
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			builder := sql.Dialect(s.Dialect())
			s.Where(
				sql.In(
					t1.C(FieldID),
					builder.Select(UsersPrimaryKey[1]).
						From(builder.Table(UsersTable)),
				),
			)
		},
		func(t *dsl.Traversal) {
			t.InE(UsersInverseLabel).InV()
		},
	)
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Table(UsersInverseTable)
			t3 := builder.Table(UsersTable)
			t4 := builder.Select(t3.C(UsersPrimaryKey[1])).
				From(t3).
				Join(t2).
				On(t3.C(UsersPrimaryKey[0]), t2.C(FieldID))
			t5 := builder.Select().From(t2)
			for _, p := range preds {
				p(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
		func(t *dsl.Traversal) {
			tr := __.OutV()
			for _, p := range preds {
				p(tr)
			}
			t.InE(UsersInverseLabel).Where(tr).InV()
		},
	)
}

// HasInfo applies the HasEdge predicate on the "info" edge.
func HasInfo() predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(InfoColumn)))
		},
		func(t *dsl.Traversal) {
			t.OutE(InfoLabel).OutV()
		},
	)
}

// HasInfoWith applies the HasEdge predicate on the "info" edge with a given conditions (other predicates).
func HasInfoWith(preds ...predicate.GroupInfo) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			builder := sql.Dialect(s.Dialect())
			t1 := s.Table()
			t2 := builder.Select(FieldID).From(builder.Table(InfoInverseTable))
			for _, p := range preds {
				p(t2)
			}
			s.Where(sql.In(t1.C(InfoColumn), t2))
		},
		func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p(tr)
			}
			t.OutE(InfoLabel).Where(tr).OutV()
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Group) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
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
func Or(predicates ...predicate.Group) predicate.Group {
	return predicate.GroupPerDialect(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
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
