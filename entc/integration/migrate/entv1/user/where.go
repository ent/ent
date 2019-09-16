// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package user

import (
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
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
func Age(v int32) predicate.User {
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

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAddress), v))
		},
	)
}

// Renamed applies equality check predicate on the "renamed" field. It's identical to RenamedEQ.
func Renamed(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldRenamed), v))
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
func AgeEQ(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAge), v))
		},
	)
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAge), v))
		},
	)
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAge), v))
		},
	)
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAge), v))
		},
	)
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAge), v))
		},
	)
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int32) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAge), v))
		},
	)
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int32) predicate.User {
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
func AgeNotIn(vs ...int32) predicate.User {
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

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldAddress), v))
		},
	)
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldAddress), v))
		},
	)
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldAddress), v))
		},
	)
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldAddress), v))
		},
	)
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldAddress), v))
		},
	)
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldAddress), v))
		},
	)
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
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
			s.Where(sql.In(s.C(FieldAddress), v...))
		},
	)
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
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
			s.Where(sql.NotIn(s.C(FieldAddress), v...))
		},
	)
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldAddress), v))
		},
	)
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldAddress), v))
		},
	)
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldAddress), v))
		},
	)
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldAddress)))
		},
	)
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldAddress)))
		},
	)
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldAddress), v))
		},
	)
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldAddress), v))
		},
	)
}

// RenamedEQ applies the EQ predicate on the "renamed" field.
func RenamedEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldRenamed), v))
		},
	)
}

// RenamedNEQ applies the NEQ predicate on the "renamed" field.
func RenamedNEQ(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NEQ(s.C(FieldRenamed), v))
		},
	)
}

// RenamedGT applies the GT predicate on the "renamed" field.
func RenamedGT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GT(s.C(FieldRenamed), v))
		},
	)
}

// RenamedGTE applies the GTE predicate on the "renamed" field.
func RenamedGTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.GTE(s.C(FieldRenamed), v))
		},
	)
}

// RenamedLT applies the LT predicate on the "renamed" field.
func RenamedLT(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LT(s.C(FieldRenamed), v))
		},
	)
}

// RenamedLTE applies the LTE predicate on the "renamed" field.
func RenamedLTE(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.LTE(s.C(FieldRenamed), v))
		},
	)
}

// RenamedIn applies the In predicate on the "renamed" field.
func RenamedIn(vs ...string) predicate.User {
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
			s.Where(sql.In(s.C(FieldRenamed), v...))
		},
	)
}

// RenamedNotIn applies the NotIn predicate on the "renamed" field.
func RenamedNotIn(vs ...string) predicate.User {
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
			s.Where(sql.NotIn(s.C(FieldRenamed), v...))
		},
	)
}

// RenamedContains applies the Contains predicate on the "renamed" field.
func RenamedContains(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.Contains(s.C(FieldRenamed), v))
		},
	)
}

// RenamedHasPrefix applies the HasPrefix predicate on the "renamed" field.
func RenamedHasPrefix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasPrefix(s.C(FieldRenamed), v))
		},
	)
}

// RenamedHasSuffix applies the HasSuffix predicate on the "renamed" field.
func RenamedHasSuffix(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.HasSuffix(s.C(FieldRenamed), v))
		},
	)
}

// RenamedIsNil applies the IsNil predicate on the "renamed" field.
func RenamedIsNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.IsNull(s.C(FieldRenamed)))
		},
	)
}

// RenamedNotNil applies the NotNil predicate on the "renamed" field.
func RenamedNotNil() predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.NotNull(s.C(FieldRenamed)))
		},
	)
}

// RenamedEqualFold applies the EqualFold predicate on the "renamed" field.
func RenamedEqualFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.EqualFold(s.C(FieldRenamed), v))
		},
	)
}

// RenamedContainsFold applies the ContainsFold predicate on the "renamed" field.
func RenamedContainsFold(v string) predicate.User {
	return predicate.User(
		func(s *sql.Selector) {
			s.Where(sql.ContainsFold(s.C(FieldRenamed), v))
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
