// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// OptionalInt applies equality check predicate on the "optional_int" field. It's identical to OptionalIntEQ.
func OptionalInt(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptionalInt, v))
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// Last applies equality check predicate on the "last" field. It's identical to LastEQ.
func Last(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLast, v))
}

// Nickname applies equality check predicate on the "nickname" field. It's identical to NicknameEQ.
func Nickname(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// Password applies equality check predicate on the "password" field. It's identical to PasswordEQ.
func Password(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// SSOCert applies equality check predicate on the "SSOCert" field. It's identical to SSOCertEQ.
func SSOCert(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSSOCert, v))
}

// FilesCount applies equality check predicate on the "files_count" field. It's identical to FilesCountEQ.
func FilesCount(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFilesCount, v))
}

// OptionalIntEQ applies the EQ predicate on the "optional_int" field.
func OptionalIntEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldOptionalInt, v))
}

// OptionalIntNEQ applies the NEQ predicate on the "optional_int" field.
func OptionalIntNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldOptionalInt, v))
}

// OptionalIntIn applies the In predicate on the "optional_int" field.
func OptionalIntIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldOptionalInt, vs...))
}

// OptionalIntNotIn applies the NotIn predicate on the "optional_int" field.
func OptionalIntNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldOptionalInt, vs...))
}

// OptionalIntGT applies the GT predicate on the "optional_int" field.
func OptionalIntGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldOptionalInt, v))
}

// OptionalIntGTE applies the GTE predicate on the "optional_int" field.
func OptionalIntGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldOptionalInt, v))
}

// OptionalIntLT applies the LT predicate on the "optional_int" field.
func OptionalIntLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldOptionalInt, v))
}

// OptionalIntLTE applies the LTE predicate on the "optional_int" field.
func OptionalIntLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldOptionalInt, v))
}

// OptionalIntIsNil applies the IsNil predicate on the "optional_int" field.
func OptionalIntIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldOptionalInt))
}

// OptionalIntNotNil applies the NotNil predicate on the "optional_int" field.
func OptionalIntNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldOptionalInt))
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAge, v))
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAge, v))
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldAge, vs...))
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAge, vs...))
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldAge, v))
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAge, v))
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldAge, v))
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAge, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldName, v))
}

// LastEQ applies the EQ predicate on the "last" field.
func LastEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldLast, v))
}

// LastNEQ applies the NEQ predicate on the "last" field.
func LastNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldLast, v))
}

// LastIn applies the In predicate on the "last" field.
func LastIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldLast, vs...))
}

// LastNotIn applies the NotIn predicate on the "last" field.
func LastNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldLast, vs...))
}

// LastGT applies the GT predicate on the "last" field.
func LastGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldLast, v))
}

// LastGTE applies the GTE predicate on the "last" field.
func LastGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldLast, v))
}

// LastLT applies the LT predicate on the "last" field.
func LastLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldLast, v))
}

// LastLTE applies the LTE predicate on the "last" field.
func LastLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldLast, v))
}

// LastEqualFold applies the EqualFold predicate on the "last" field.
func LastEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldLast, v))
}

// LastContains applies the Contains predicate on the "last" field.
func LastContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldLast, v))
}

// LastContainsFold applies the ContainsFold predicate on the "last" field.
func LastContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldLast, v))
}

// LastHasPrefix applies the HasPrefix predicate on the "last" field.
func LastHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldLast, v))
}

// LastHasSuffix applies the HasSuffix predicate on the "last" field.
func LastHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldLast, v))
}

// NicknameEQ applies the EQ predicate on the "nickname" field.
func NicknameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldNickname, v))
}

// NicknameNEQ applies the NEQ predicate on the "nickname" field.
func NicknameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldNickname, v))
}

// NicknameIn applies the In predicate on the "nickname" field.
func NicknameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldNickname, vs...))
}

// NicknameNotIn applies the NotIn predicate on the "nickname" field.
func NicknameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldNickname, vs...))
}

// NicknameGT applies the GT predicate on the "nickname" field.
func NicknameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldNickname, v))
}

// NicknameGTE applies the GTE predicate on the "nickname" field.
func NicknameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldNickname, v))
}

// NicknameLT applies the LT predicate on the "nickname" field.
func NicknameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldNickname, v))
}

// NicknameLTE applies the LTE predicate on the "nickname" field.
func NicknameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldNickname, v))
}

// NicknameEqualFold applies the EqualFold predicate on the "nickname" field.
func NicknameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldNickname, v))
}

// NicknameContains applies the Contains predicate on the "nickname" field.
func NicknameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldNickname, v))
}

// NicknameContainsFold applies the ContainsFold predicate on the "nickname" field.
func NicknameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldNickname, v))
}

// NicknameHasPrefix applies the HasPrefix predicate on the "nickname" field.
func NicknameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldNickname, v))
}

// NicknameHasSuffix applies the HasSuffix predicate on the "nickname" field.
func NicknameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldNickname, v))
}

// NicknameIsNil applies the IsNil predicate on the "nickname" field.
func NicknameIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldNickname))
}

// NicknameNotNil applies the NotNil predicate on the "nickname" field.
func NicknameNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldNickname))
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldAddress, v))
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldAddress, v))
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldAddress, vs...))
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldAddress, vs...))
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldAddress, v))
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldAddress, v))
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldAddress, v))
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldAddress, v))
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldAddress, v))
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldAddress, v))
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldAddress, v))
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldAddress, v))
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldAddress, v))
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldAddress))
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldAddress))
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPhone, v))
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPhone, v))
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPhone, vs...))
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPhone, vs...))
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPhone, v))
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPhone, v))
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPhone, v))
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPhone, v))
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPhone, v))
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPhone, v))
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPhone, v))
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPhone, v))
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPhone, v))
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPhone))
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPhone))
}

// PasswordEQ applies the EQ predicate on the "password" field.
func PasswordEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPassword, v))
}

// PasswordNEQ applies the NEQ predicate on the "password" field.
func PasswordNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPassword, v))
}

// PasswordIn applies the In predicate on the "password" field.
func PasswordIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPassword, vs...))
}

// PasswordNotIn applies the NotIn predicate on the "password" field.
func PasswordNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPassword, vs...))
}

// PasswordGT applies the GT predicate on the "password" field.
func PasswordGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPassword, v))
}

// PasswordGTE applies the GTE predicate on the "password" field.
func PasswordGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPassword, v))
}

// PasswordLT applies the LT predicate on the "password" field.
func PasswordLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPassword, v))
}

// PasswordLTE applies the LTE predicate on the "password" field.
func PasswordLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPassword, v))
}

// PasswordEqualFold applies the EqualFold predicate on the "password" field.
func PasswordEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPassword, v))
}

// PasswordContains applies the Contains predicate on the "password" field.
func PasswordContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPassword, v))
}

// PasswordContainsFold applies the ContainsFold predicate on the "password" field.
func PasswordContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPassword, v))
}

// PasswordHasPrefix applies the HasPrefix predicate on the "password" field.
func PasswordHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPassword, v))
}

// PasswordHasSuffix applies the HasSuffix predicate on the "password" field.
func PasswordHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPassword, v))
}

// PasswordIsNil applies the IsNil predicate on the "password" field.
func PasswordIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldPassword))
}

// PasswordNotNil applies the NotNil predicate on the "password" field.
func PasswordNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldPassword))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v Role) predicate.User {
	return predicate.User(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v Role) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...Role) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldRole, vs...))
}

// EmploymentEQ applies the EQ predicate on the "employment" field.
func EmploymentEQ(v Employment) predicate.User {
	return predicate.User(sql.FieldEQ(FieldEmployment, v))
}

// EmploymentNEQ applies the NEQ predicate on the "employment" field.
func EmploymentNEQ(v Employment) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldEmployment, v))
}

// EmploymentIn applies the In predicate on the "employment" field.
func EmploymentIn(vs ...Employment) predicate.User {
	return predicate.User(sql.FieldIn(FieldEmployment, vs...))
}

// EmploymentNotIn applies the NotIn predicate on the "employment" field.
func EmploymentNotIn(vs ...Employment) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldEmployment, vs...))
}

// SSOCertEQ applies the EQ predicate on the "SSOCert" field.
func SSOCertEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldSSOCert, v))
}

// SSOCertNEQ applies the NEQ predicate on the "SSOCert" field.
func SSOCertNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldSSOCert, v))
}

// SSOCertIn applies the In predicate on the "SSOCert" field.
func SSOCertIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldSSOCert, vs...))
}

// SSOCertNotIn applies the NotIn predicate on the "SSOCert" field.
func SSOCertNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldSSOCert, vs...))
}

// SSOCertGT applies the GT predicate on the "SSOCert" field.
func SSOCertGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldSSOCert, v))
}

// SSOCertGTE applies the GTE predicate on the "SSOCert" field.
func SSOCertGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldSSOCert, v))
}

// SSOCertLT applies the LT predicate on the "SSOCert" field.
func SSOCertLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldSSOCert, v))
}

// SSOCertLTE applies the LTE predicate on the "SSOCert" field.
func SSOCertLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldSSOCert, v))
}

// SSOCertEqualFold applies the EqualFold predicate on the "SSOCert" field.
func SSOCertEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldSSOCert, v))
}

// SSOCertContains applies the Contains predicate on the "SSOCert" field.
func SSOCertContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldSSOCert, v))
}

// SSOCertContainsFold applies the ContainsFold predicate on the "SSOCert" field.
func SSOCertContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldSSOCert, v))
}

// SSOCertHasPrefix applies the HasPrefix predicate on the "SSOCert" field.
func SSOCertHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldSSOCert, v))
}

// SSOCertHasSuffix applies the HasSuffix predicate on the "SSOCert" field.
func SSOCertHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldSSOCert, v))
}

// SSOCertIsNil applies the IsNil predicate on the "SSOCert" field.
func SSOCertIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldSSOCert))
}

// SSOCertNotNil applies the NotNil predicate on the "SSOCert" field.
func SSOCertNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldSSOCert))
}

// FilesCountEQ applies the EQ predicate on the "files_count" field.
func FilesCountEQ(v int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldFilesCount, v))
}

// FilesCountNEQ applies the NEQ predicate on the "files_count" field.
func FilesCountNEQ(v int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldFilesCount, v))
}

// FilesCountIn applies the In predicate on the "files_count" field.
func FilesCountIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldFilesCount, vs...))
}

// FilesCountNotIn applies the NotIn predicate on the "files_count" field.
func FilesCountNotIn(vs ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldFilesCount, vs...))
}

// FilesCountGT applies the GT predicate on the "files_count" field.
func FilesCountGT(v int) predicate.User {
	return predicate.User(sql.FieldGT(FieldFilesCount, v))
}

// FilesCountGTE applies the GTE predicate on the "files_count" field.
func FilesCountGTE(v int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldFilesCount, v))
}

// FilesCountLT applies the LT predicate on the "files_count" field.
func FilesCountLT(v int) predicate.User {
	return predicate.User(sql.FieldLT(FieldFilesCount, v))
}

// FilesCountLTE applies the LTE predicate on the "files_count" field.
func FilesCountLTE(v int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldFilesCount, v))
}

// FilesCountIsNil applies the IsNil predicate on the "files_count" field.
func FilesCountIsNil() predicate.User {
	return predicate.User(sql.FieldIsNull(FieldFilesCount))
}

// FilesCountNotNil applies the NotNil predicate on the "files_count" field.
func FilesCountNotNil() predicate.User {
	return predicate.User(sql.FieldNotNull(FieldFilesCount))
}

// HasCard applies the HasEdge predicate on the "card" edge.
func HasCard() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, CardTable, CardColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCardWith applies the HasEdge predicate on the "card" edge with a given conditions (other predicates).
func HasCardWith(preds ...predicate.Card) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newCardStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPets applies the HasEdge predicate on the "pets" edge.
func HasPets() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PetsTable, PetsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPetsWith applies the HasEdge predicate on the "pets" edge with a given conditions (other predicates).
func HasPetsWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newPetsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFiles applies the HasEdge predicate on the "files" edge.
func HasFiles() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, FilesTable, FilesColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFilesWith applies the HasEdge predicate on the "files" edge with a given conditions (other predicates).
func HasFilesWith(preds ...predicate.File) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newFilesStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasGroups applies the HasEdge predicate on the "groups" edge.
func HasGroups() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GroupsTable, GroupsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGroupsWith applies the HasEdge predicate on the "groups" edge with a given conditions (other predicates).
func HasGroupsWith(preds ...predicate.Group) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newGroupsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFriends applies the HasEdge predicate on the "friends" edge.
func HasFriends() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FriendsTable, FriendsPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFriendsWith applies the HasEdge predicate on the "friends" edge with a given conditions (other predicates).
func HasFriendsWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newFriendsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFollowers applies the HasEdge predicate on the "followers" edge.
func HasFollowers() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, FollowersTable, FollowersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFollowersWith applies the HasEdge predicate on the "followers" edge with a given conditions (other predicates).
func HasFollowersWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newFollowersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasFollowing applies the HasEdge predicate on the "following" edge.
func HasFollowing() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, FollowingTable, FollowingPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasFollowingWith applies the HasEdge predicate on the "following" edge with a given conditions (other predicates).
func HasFollowingWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newFollowingStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTeam applies the HasEdge predicate on the "team" edge.
func HasTeam() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, TeamTable, TeamColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTeamWith applies the HasEdge predicate on the "team" edge with a given conditions (other predicates).
func HasTeamWith(preds ...predicate.Pet) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newTeamStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSpouse applies the HasEdge predicate on the "spouse" edge.
func HasSpouse() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpouseWith applies the HasEdge predicate on the "spouse" edge with a given conditions (other predicates).
func HasSpouseWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newSpouseStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasChildren applies the HasEdge predicate on the "children" edge.
func HasChildren() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, ChildrenTable, ChildrenColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasChildrenWith applies the HasEdge predicate on the "children" edge with a given conditions (other predicates).
func HasChildrenWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newChildrenStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasParent applies the HasEdge predicate on the "parent" edge.
func HasParent() predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ParentTable, ParentColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasParentWith applies the HasEdge predicate on the "parent" edge with a given conditions (other predicates).
func HasParentWith(preds ...predicate.User) predicate.User {
	return predicate.User(func(s *sql.Selector) {
		step := newParentStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
