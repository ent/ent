// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(id)
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.EQ(id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.NEQ(id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Within(v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		t.HasID(p.Without(v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.GT(id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.GTE(id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.LT(id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasID(p.LTE(id))
	})
}

// OptionalInt applies equality check predicate on the "optional_int" field. It's identical to OptionalIntEQ.
func OptionalInt(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.EQ(v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// Last applies equality check predicate on the "last" field. It's identical to LastEQ.
func Last(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.EQ(v))
	})
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EQ(v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.EQ(v))
	})
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.EQ(v))
	})
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.EQ(v))
	})
}

// SSOCert applies equality check predicate on the "SSOCert" field. It's identical to SSOCertEQ.
func SSOCert(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.EQ(v))
	})
}

// OptionalIntEQ applies the EQ predicate on the "optional_int" field.
func OptionalIntEQ(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.EQ(v))
	})
}

// OptionalIntNEQ applies the NEQ predicate on the "optional_int" field.
func OptionalIntNEQ(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.NEQ(v))
	})
}

// OptionalIntIn applies the In predicate on the "optional_int" field.
func OptionalIntIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.Within(v...))
	})
}

// OptionalIntNotIn applies the NotIn predicate on the "optional_int" field.
func OptionalIntNotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.Without(v...))
	})
}

// OptionalIntGT applies the GT predicate on the "optional_int" field.
func OptionalIntGT(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.GT(v))
	})
}

// OptionalIntGTE applies the GTE predicate on the "optional_int" field.
func OptionalIntGTE(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.GTE(v))
	})
}

// OptionalIntLT applies the LT predicate on the "optional_int" field.
func OptionalIntLT(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.LT(v))
	})
}

// OptionalIntLTE applies the LTE predicate on the "optional_int" field.
func OptionalIntLTE(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldOptionalInt, p.LTE(v))
	})
}

// OptionalIntIsNil applies the IsNil predicate on the "optional_int" field.
func OptionalIntIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldOptionalInt)
	})
}

// OptionalIntNotNil applies the NotNil predicate on the "optional_int" field.
func OptionalIntNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldOptionalInt)
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.EQ(v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.NEQ(v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Within(v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.Without(v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GT(v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.GTE(v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LT(v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAge, p.LTE(v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EQ(v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.NEQ(v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Within(v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Without(v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GT(v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.GTE(v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LT(v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.LTE(v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.Containing(v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.StartingWith(v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldName, p.EndingWith(v))
	})
}

// LastEQ applies the EQ predicate on the "last" field.
func LastEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.EQ(v))
	})
}

// LastNEQ applies the NEQ predicate on the "last" field.
func LastNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.NEQ(v))
	})
}

// LastIn applies the In predicate on the "last" field.
func LastIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.Within(v...))
	})
}

// LastNotIn applies the NotIn predicate on the "last" field.
func LastNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.Without(v...))
	})
}

// LastGT applies the GT predicate on the "last" field.
func LastGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.GT(v))
	})
}

// LastGTE applies the GTE predicate on the "last" field.
func LastGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.GTE(v))
	})
}

// LastLT applies the LT predicate on the "last" field.
func LastLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.LT(v))
	})
}

// LastLTE applies the LTE predicate on the "last" field.
func LastLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.LTE(v))
	})
}

// LastContains applies the Contains predicate on the "last" field.
func LastContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.Containing(v))
	})
}

// LastHasPrefix applies the HasPrefix predicate on the "last" field.
func LastHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.StartingWith(v))
	})
}

// LastHasSuffix applies the HasSuffix predicate on the "last" field.
func LastHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldLast, p.EndingWith(v))
	})
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EQ(v))
	})
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.NEQ(v))
	})
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Within(v...))
	})
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Without(v...))
	})
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.GT(v))
	})
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.GTE(v))
	})
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.LT(v))
	})
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.LTE(v))
	})
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.Containing(v))
	})
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.StartingWith(v))
	})
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldNickname, p.EndingWith(v))
	})
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldNickname)
	})
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldNickname)
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.EQ(v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.NEQ(v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.Within(v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.Without(v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.GT(v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.GTE(v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.LT(v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.LTE(v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.Containing(v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.StartingWith(v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldAddress, p.EndingWith(v))
	})
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldAddress)
	})
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldAddress)
	})
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.EQ(v))
	})
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.NEQ(v))
	})
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.Within(v...))
	})
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.Without(v...))
	})
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.GT(v))
	})
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.GTE(v))
	})
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.LT(v))
	})
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.LTE(v))
	})
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.Containing(v))
	})
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.StartingWith(v))
	})
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPhone, p.EndingWith(v))
	})
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldPhone)
	})
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldPhone)
	})
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.EQ(v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.NEQ(v))
	})
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.Within(v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.Without(v...))
	})
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.GT(v))
	})
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.GTE(v))
	})
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.LT(v))
	})
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.LTE(v))
	})
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.Containing(v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.StartingWith(v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldPassword, p.EndingWith(v))
	})
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldPassword)
	})
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldPassword)
	})
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldRole, p.EQ(v))
	})
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldRole, p.NEQ(v))
	})
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldRole, p.Within(v...))
	})
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldRole, p.Without(v...))
	})
}

// EmploymentEQ applies the EQ predicate on the "employment" field.
func EmploymentEQ(v Employment) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmployment, p.EQ(v))
	})
}

// EmploymentNEQ applies the NEQ predicate on the "employment" field.
func EmploymentNEQ(v Employment) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmployment, p.NEQ(v))
	})
}

// EmploymentIn applies the In predicate on the "employment" field.
func EmploymentIn(vs ...Employment) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmployment, p.Within(v...))
	})
}

// EmploymentNotIn applies the NotIn predicate on the "employment" field.
func EmploymentNotIn(vs ...Employment) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldEmployment, p.Without(v...))
	})
}

// SSOCertEQ applies the EQ predicate on the "SSOCert" field.
func SSOCertEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.EQ(v))
	})
}

// SSOCertNEQ applies the NEQ predicate on the "SSOCert" field.
func SSOCertNEQ(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.NEQ(v))
	})
}

// SSOCertIn applies the In predicate on the "SSOCert" field.
func SSOCertIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.Within(v...))
	})
}

// SSOCertNotIn applies the NotIn predicate on the "SSOCert" field.
func SSOCertNotIn(vs ...string) predicate.User {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.Without(v...))
	})
}

// SSOCertGT applies the GT predicate on the "SSOCert" field.
func SSOCertGT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.GT(v))
	})
}

// SSOCertGTE applies the GTE predicate on the "SSOCert" field.
func SSOCertGTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.GTE(v))
	})
}

// SSOCertLT applies the LT predicate on the "SSOCert" field.
func SSOCertLT(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.LT(v))
	})
}

// SSOCertLTE applies the LTE predicate on the "SSOCert" field.
func SSOCertLTE(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.LTE(v))
	})
}

// SSOCertContains applies the Contains predicate on the "SSOCert" field.
func SSOCertContains(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.Containing(v))
	})
}

// SSOCertHasPrefix applies the HasPrefix predicate on the "SSOCert" field.
func SSOCertHasPrefix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.StartingWith(v))
	})
}

// SSOCertHasSuffix applies the HasSuffix predicate on the "SSOCert" field.
func SSOCertHasSuffix(v string) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Has(Label, FieldSSOCert, p.EndingWith(v))
	})
}

// SSOCertIsNil applies the IsNil predicate on the "SSOCert" field.
func SSOCertIsNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).HasNot(FieldSSOCert)
	})
}

// SSOCertNotNil applies the NotNil predicate on the "SSOCert" field.
func SSOCertNotNil() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.HasLabel(Label).Has(FieldSSOCert)
	})
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(CardLabel).OutV()
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Card) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(CardLabel).Where(tr).OutV()
	})
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(PetsLabel).OutV()
	})
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(PetsLabel).Where(tr).OutV()
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(FilesLabel).OutV()
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(FilesLabel).Where(tr).OutV()
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(GroupsLabel).OutV()
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(GroupsLabel).Where(tr).OutV()
	})
}

// HasFriends applies the HasEdge predicate on the "friends" edge.
func HasFriends() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Both(FriendsLabel)
	})
}

// HasFriendsWith applies the HasEdge predicate on the "friends" edge with a given conditions (other predicates).
func HasFriendsWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		in, out := __.InV(), __.OutV()
		for _, p := range preds {
			p(in)
			p(out)
		}
		t.Where(
			__.Or(
				__.OutE(FriendsLabel).Where(in),
				__.InE(FriendsLabel).Where(out),
			),
		)
	})
}

// HasFollowers applies the HasEdge predicate on the "followers" edge.
func HasFollowers() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.InE(FollowersInverseLabel).InV()
	})
}

// HasFollowersWith applies the HasEdge predicate on the "followers" edge with a given conditions (other predicates).
func HasFollowersWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(FollowersInverseLabel).Where(tr).InV()
	})
}

// HasFollowing applies the HasEdge predicate on the "following" edge.
func HasFollowing() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(FollowingLabel).OutV()
	})
}

// HasFollowingWith applies the HasEdge predicate on the "following" edge with a given conditions (other predicates).
func HasFollowingWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(FollowingLabel).Where(tr).OutV()
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(TeamLabel).OutV()
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(TeamLabel).Where(tr).OutV()
	})
}

// HasSpouse applies the HasEdge predicate on the "spouse" edge.
func HasSpouse() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.Both(SpouseLabel)
	})
}

// HasSpouseWith applies the HasEdge predicate on the "spouse" edge with a given conditions (other predicates).
func HasSpouseWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		in, out := __.InV(), __.OutV()
		for _, p := range preds {
			p(in)
			p(out)
		}
		t.Where(
			__.Or(
				__.OutE(SpouseLabel).Where(in),
				__.InE(SpouseLabel).Where(out),
			),
		)
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.InE(ChildrenInverseLabel).InV()
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.OutV()
		for _, p := range preds {
			p(tr)
		}
		t.InE(ChildrenInverseLabel).Where(tr).InV()
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		t.OutE(ParentLabel).OutV()
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(t *dsl.Traversal) {
		tr := __.InV()
		for _, p := range preds {
			p(tr)
		}
		t.OutE(ParentLabel).Where(tr).OutV()
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.And(trs...))
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(func(tr *dsl.Traversal) {
		trs := make([]interface{}, 0, len(predicates))
		for _, p := range predicates {
			t := __.New()
			p(t)
			trs = append(trs, t)
		}
		tr.Where(__.Or(trs...))
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(func(tr *dsl.Traversal) {
		t := __.New()
		p(t)
		tr.Where(__.Not(t))
	})
}
