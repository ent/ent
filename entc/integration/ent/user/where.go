// Code generated (@generated) by entc, DO NOT EDIT.

package user

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

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.EQ(v))
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

// Last applies equality check predicate on the "last" field. It's identical to LastEQ.
func Last(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.EQ(v))
		},
	}
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.EQ(v))
		},
	}
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.EQ(v))
		},
	}
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.EQ(v))
		},
	}
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.NEQ(v))
		},
	}
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.GT(v))
		},
	}
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.GTE(v))
		},
	}
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.LT(v))
		},
	}
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAge), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.LTE(v))
		},
	}
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldAge), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.Within(v...))
		},
	}
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldAge), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldAge, p.Without(v...))
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

// LastEQ applies the EQ predicate on the "last" field.
func LastEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.EQ(v))
		},
	}
}

// LastNEQ applies the NEQ predicate on the "last" field.
func LastNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.NEQ(v))
		},
	}
}

// LastGT applies the GT predicate on the "last" field.
func LastGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.GT(v))
		},
	}
}

// LastGTE applies the GTE predicate on the "last" field.
func LastGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.GTE(v))
		},
	}
}

// LastLT applies the LT predicate on the "last" field.
func LastLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.LT(v))
		},
	}
}

// LastLTE applies the LTE predicate on the "last" field.
func LastLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.LTE(v))
		},
	}
}

// LastIn applies the In predicate on the "last" field.
func LastIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldLast), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.Within(v...))
		},
	}
}

// LastNotIn applies the NotIn predicate on the "last" field.
func LastNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldLast), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.Without(v...))
		},
	}
}

// LastContains applies the Contains predicate on the "last" field.
func LastContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.Containing(v))
		},
	}
}

// LastHasPrefix applies the HasPrefix predicate on the "last" field.
func LastHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.StartingWith(v))
		},
	}
}

// LastHasSuffix applies the HasSuffix predicate on the "last" field.
func LastHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldLast), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldLast, p.EndingWith(v))
		},
	}
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.EQ(v))
		},
	}
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.NEQ(v))
		},
	}
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.GT(v))
		},
	}
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.GTE(v))
		},
	}
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.LT(v))
		},
	}
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.LTE(v))
		},
	}
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldNickname), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.Within(v...))
		},
	}
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldNickname), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.Without(v...))
		},
	}
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.Containing(v))
		},
	}
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.StartingWith(v))
		},
	}
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNickname), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldNickname, p.EndingWith(v))
		},
	}
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.EQ(v))
		},
	}
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.NEQ(v))
		},
	}
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.GT(v))
		},
	}
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.GTE(v))
		},
	}
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.LT(v))
		},
	}
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.LTE(v))
		},
	}
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) ent.Predicate {
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
			s.Where(sql.In(s.C(FieldPhone), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.Within(v...))
		},
	}
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) ent.Predicate {
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
			s.Where(sql.NotIn(s.C(FieldPhone), v...))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.Without(v...))
		},
	}
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.Containing(v))
		},
	}
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.StartingWith(v))
		},
	}
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldPhone), v))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Has(Label, FieldPhone, p.EndingWith(v))
		},
	}
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(CardColumn).
						From(sql.Table(CardTable)).
						Where(sql.NotNull(CardColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(CardLabel).OutV()
		},
	}
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(CardColumn).From(sql.Table(CardTable))
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
			t.OutE(CardLabel).Where(tr).OutV()
		},
	}
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(PetsColumn).
						From(sql.Table(PetsTable)).
						Where(sql.NotNull(PetsColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(PetsLabel).OutV()
		},
	}
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(PetsColumn).From(sql.Table(PetsTable))
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
			t.OutE(PetsLabel).Where(tr).OutV()
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

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(GroupsPrimaryKey[0]).From(sql.Table(GroupsTable)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(GroupsLabel).OutV()
		},
	}
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Table(GroupsInverseTable)
			t3 := sql.Table(GroupsTable)
			t4 := sql.Select(t3.C(GroupsPrimaryKey[0])).
				From(t3).
				Join(t2).
				On(t3.C(GroupsPrimaryKey[1]), t2.C(FieldID))
			t5 := sql.Select().From(t2)
			for _, p := range preds {
				p.SQL(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.OutE(GroupsLabel).Where(tr).OutV()
		},
	}
}

// HasFriends applies the HasEdge predicate on the "friends" edge.
func HasFriends() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(FriendsPrimaryKey[0]).From(sql.Table(FriendsTable)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Both(FriendsLabel)
		},
	}
}

// HasFriendsWith applies the HasEdge predicate on the "friends" edge with a given conditions (other predicates).
func HasFriendsWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Table(Table)
			t3 := sql.Table(FriendsTable)
			t4 := sql.Select(t3.C(FriendsPrimaryKey[0])).
				From(t3).
				Join(t2).
				On(t3.C(FriendsPrimaryKey[1]), t2.C(FieldID))
			t5 := sql.Select().From(t2)
			for _, p := range preds {
				p.SQL(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
		Gremlin: func(t *dsl.Traversal) {
			in, out := __.InV(), __.OutV()
			for _, p := range preds {
				p.Gremlin(in)
				p.Gremlin(out)
			}
			t.Where(
				__.Or(
					__.OutE(FriendsLabel).Where(in),
					__.InE(FriendsLabel).Where(out),
				),
			)
		},
	}
}

// HasFollowers applies the HasEdge predicate on the "followers" edge.
func HasFollowers() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(FollowersPrimaryKey[1]).From(sql.Table(FollowersTable)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(FollowersInverseLabel).InV()
		},
	}
}

// HasFollowersWith applies the HasEdge predicate on the "followers" edge with a given conditions (other predicates).
func HasFollowersWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Table(Table)
			t3 := sql.Table(FollowersTable)
			t4 := sql.Select(t3.C(FollowersPrimaryKey[1])).
				From(t3).
				Join(t2).
				On(t3.C(FollowersPrimaryKey[0]), t2.C(FieldID))
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
			t.InE(FollowersInverseLabel).Where(tr).InV()
		},
	}
}

// HasFollowing applies the HasEdge predicate on the "following" edge.
func HasFollowing() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(FollowingPrimaryKey[0]).From(sql.Table(FollowingTable)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(FollowingLabel).OutV()
		},
	}
}

// HasFollowingWith applies the HasEdge predicate on the "following" edge with a given conditions (other predicates).
func HasFollowingWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Table(Table)
			t3 := sql.Table(FollowingTable)
			t4 := sql.Select(t3.C(FollowingPrimaryKey[0])).
				From(t3).
				Join(t2).
				On(t3.C(FollowingPrimaryKey[1]), t2.C(FieldID))
			t5 := sql.Select().From(t2)
			for _, p := range preds {
				p.SQL(t5)
			}
			t4.FromSelect(t5)
			s.Where(sql.In(t1.C(FieldID), t4))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.OutE(FollowingLabel).Where(tr).OutV()
		},
	}
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(TeamColumn).
						From(sql.Table(TeamTable)).
						Where(sql.NotNull(TeamColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(TeamLabel).OutV()
		},
	}
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(TeamColumn).From(sql.Table(TeamTable))
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
			t.OutE(TeamLabel).Where(tr).OutV()
		},
	}
}

// HasSpouse applies the HasEdge predicate on the "spouse" edge.
func HasSpouse() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(SpouseColumn).
						From(sql.Table(SpouseTable)).
						Where(sql.NotNull(SpouseColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.Both(SpouseLabel)
		},
	}
}

// HasSpouseWith applies the HasEdge predicate on the "spouse" edge with a given conditions (other predicates).
func HasSpouseWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(SpouseColumn).From(sql.Table(SpouseTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(FieldID), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			in, out := __.InV(), __.OutV()
			for _, p := range preds {
				p.Gremlin(in)
				p.Gremlin(out)
			}
			t.Where(
				__.Or(
					__.OutE(SpouseLabel).Where(in),
					__.InE(SpouseLabel).Where(out),
				),
			)
		},
	}
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(
				sql.In(
					t1.C(FieldID),
					sql.Select(ChildrenColumn).
						From(sql.Table(ChildrenTable)).
						Where(sql.NotNull(ChildrenColumn)),
				),
			)
		},
		Gremlin: func(t *dsl.Traversal) {
			t.InE(ChildrenInverseLabel).InV()
		},
	}
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(ChildrenColumn).From(sql.Table(ChildrenTable))
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
			t.InE(ChildrenInverseLabel).Where(tr).InV()
		},
	}
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			s.Where(sql.NotNull(t1.C(ParentColumn)))
		},
		Gremlin: func(t *dsl.Traversal) {
			t.OutE(ParentLabel).OutV()
		},
	}
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...ent.Predicate) ent.Predicate {
	return ent.Predicate{
		SQL: func(s *sql.Selector) {
			t1 := s.Table()
			t2 := sql.Select(FieldID).From(sql.Table(ParentTable))
			for _, p := range preds {
				p.SQL(t2)
			}
			s.Where(sql.In(t1.C(ParentColumn), t2))
		},
		Gremlin: func(t *dsl.Traversal) {
			tr := __.InV()
			for _, p := range preds {
				p.Gremlin(tr)
			}
			t.OutE(ParentLabel).Where(tr).OutV()
		},
	}
}
