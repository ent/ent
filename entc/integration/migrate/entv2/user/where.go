// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldID), id))
		},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldID), id))
		},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldID), id))
		},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldID), id))
		},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldID), id))
		},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.In(s.C(FieldID), v...))
		},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(ids) == 0 {
				s.Where(sql.False())
				return
			}
			v := make([]interface{}, len(ids))
			for i := range v {
				v[i] = ids[i]
			}
			s.Where(sql.NotIn(s.C(FieldID), v...))
		},
	)
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
	)
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPhone), v))
		},
	)
}

// Buffer applies equality check predicate on the "buffer" field. It's identical to BufferEQ.
func Buffer(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldBuffer), v))
		},
	)
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldTitle), v))
		},
	)
}

// NewName applies equality check predicate on the "new_name" field. It's identical to NewNameEQ.
func NewName(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNewName), v))
		},
	)
}

// Blob applies equality check predicate on the "blob" field. It's identical to BlobEQ.
func Blob(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldBlob), v))
		},
	)
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
	)
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAge), v))
		},
	)
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAge), v))
		},
	)
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAge), v))
		},
	)
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAge), v))
		},
	)
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAge), v))
		},
	)
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldAge), v...))
		},
	)
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldAge), v...))
		},
	)
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldName), v))
		},
	)
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldName), v))
		},
	)
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldName), v))
		},
	)
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldName), v))
		},
	)
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldName), v))
		},
	)
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldName), v))
		},
	)
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldName), v...))
		},
	)
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldName), v...))
		},
	)
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldName), v))
		},
	)
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldName), v))
		},
	)
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldName), v))
		},
	)
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldName), v))
		},
	)
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldName), v))
		},
	)
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldPhone), v))
		},
	)
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldPhone), v))
		},
	)
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldPhone), v))
		},
	)
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldPhone), v))
		},
	)
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldPhone), v))
		},
	)
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldPhone), v))
		},
	)
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldPhone), v...))
		},
	)
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldPhone), v...))
		},
	)
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldPhone), v))
		},
	)
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldPhone), v))
		},
	)
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldPhone), v))
		},
	)
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldPhone), v))
		},
	)
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldPhone), v))
		},
	)
}

// BufferEQ applies the EQ predicate on the "buffer" field.
func BufferEQ(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldBuffer), v))
		},
	)
}

// BufferNEQ applies the NEQ predicate on the "buffer" field.
func BufferNEQ(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldBuffer), v))
		},
	)
}

// BufferGT applies the GT predicate on the "buffer" field.
func BufferGT(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldBuffer), v))
		},
	)
}

// BufferGTE applies the GTE predicate on the "buffer" field.
func BufferGTE(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldBuffer), v))
		},
	)
}

// BufferLT applies the LT predicate on the "buffer" field.
func BufferLT(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldBuffer), v))
		},
	)
}

// BufferLTE applies the LTE predicate on the "buffer" field.
func BufferLTE(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldBuffer), v))
		},
	)
}

// BufferIn applies the In predicate on the "buffer" field.
func BufferIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldBuffer), v...))
		},
	)
}

// BufferNotIn applies the NotIn predicate on the "buffer" field.
func BufferNotIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldBuffer), v...))
		},
	)
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldTitle), v))
		},
	)
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldTitle), v))
		},
	)
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldTitle), v))
		},
	)
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldTitle), v))
		},
	)
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldTitle), v))
		},
	)
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldTitle), v))
		},
	)
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldTitle), v...))
		},
	)
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldTitle), v...))
		},
	)
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldTitle), v))
		},
	)
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldTitle), v))
		},
	)
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldTitle), v))
		},
	)
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldTitle), v))
		},
	)
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldTitle), v))
		},
	)
}

// NewNameEQ applies the EQ predicate on the "new_name" field.
func NewNameEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldNewName), v))
		},
	)
}

// NewNameNEQ applies the NEQ predicate on the "new_name" field.
func NewNameNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldNewName), v))
		},
	)
}

// NewNameGT applies the GT predicate on the "new_name" field.
func NewNameGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldNewName), v))
		},
	)
}

// NewNameGTE applies the GTE predicate on the "new_name" field.
func NewNameGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldNewName), v))
		},
	)
}

// NewNameLT applies the LT predicate on the "new_name" field.
func NewNameLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldNewName), v))
		},
	)
}

// NewNameLTE applies the LTE predicate on the "new_name" field.
func NewNameLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldNewName), v))
		},
	)
}

// NewNameIn applies the In predicate on the "new_name" field.
func NewNameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldNewName), v...))
		},
	)
}

// NewNameNotIn applies the NotIn predicate on the "new_name" field.
func NewNameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldNewName), v...))
		},
	)
}

// NewNameContains applies the Contains predicate on the "new_name" field.
func NewNameContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldNewName), v))
		},
	)
}

// NewNameHasPrefix applies the HasPrefix predicate on the "new_name" field.
func NewNameHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldNewName), v))
		},
	)
}

// NewNameHasSuffix applies the HasSuffix predicate on the "new_name" field.
func NewNameHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldNewName), v))
		},
	)
}

// NewNameIsNil applies the IsNil predicate on the "new_name" field.
func NewNameIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldNewName)))
		},
	)
}

// NewNameNotNil applies the NotNil predicate on the "new_name" field.
func NewNameNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldNewName)))
		},
	)
}

// NewNameEqualFold applies the EqualFold predicate on the "new_name" field.
func NewNameEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldNewName), v))
		},
	)
}

// NewNameContainsFold applies the ContainsFold predicate on the "new_name" field.
func NewNameContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldNewName), v))
		},
	)
}

// BlobEQ applies the EQ predicate on the "blob" field.
func BlobEQ(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldBlob), v))
		},
	)
}

// BlobNEQ applies the NEQ predicate on the "blob" field.
func BlobNEQ(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldBlob), v))
		},
	)
}

// BlobGT applies the GT predicate on the "blob" field.
func BlobGT(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldBlob), v))
		},
	)
}

// BlobGTE applies the GTE predicate on the "blob" field.
func BlobGTE(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldBlob), v))
		},
	)
}

// BlobLT applies the LT predicate on the "blob" field.
func BlobLT(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldBlob), v))
		},
	)
}

// BlobLTE applies the LTE predicate on the "blob" field.
func BlobLTE(v []byte) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldBlob), v))
		},
	)
}

// BlobIn applies the In predicate on the "blob" field.
func BlobIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.In(s.C(FieldBlob), v...))
		},
	)
}

// BlobNotIn applies the NotIn predicate on the "blob" field.
func BlobNotIn(vs ...[]byte) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(
		func(s *sql.Selector) {
			// if not arguments were provided, append the FALSE constants,
			// since we can't apply "IN ()". This will make this predicate falsy.
			if len(vs) == 0 {
				s.Where(sql.False())
				return
			}
			s.Where(sql.NotIn(s.C(FieldBlob), v...))
		},
	)
}

// BlobIsNil applies the IsNil predicate on the "blob" field.
func BlobIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldBlob)))
		},
	)
}

// BlobNotNil applies the NotNil predicate on the "blob" field.
func BlobNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldBlob)))
		},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			for _, p := range predicates {
				p(s)
			}
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			for i, p := range predicates {
				if i > 0 {
					s.Or()
				}
				p(s)
			}
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}
