// Code generated (@generated) by entc, DO NOT EDIT.

package user

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
