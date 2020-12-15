// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package pet

const (
	// Label holds the string label denoting the pet type in the database.
	Label = "pet"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// EdgeFavouriteFood holds the string denoting the favourite_food edge name in mutations.
	EdgeFavouriteFood = "favourite_food"

	// Table holds the table name of the pet in the database.
	Table = "pets"
	// OwnerTable is the table the holds the owner relation/edge.
	OwnerTable = "pets"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_pets"
	// FavouriteFoodTable is the table the holds the favourite_food relation/edge.
	FavouriteFoodTable = "pets"
	// FavouriteFoodInverseTable is the table name for the Food entity.
	// It exists in this package in order to avoid circular dependency with the "food" package.
	FavouriteFoodInverseTable = "foods"
	// FavouriteFoodColumn is the table column denoting the favourite_food relation/edge.
	FavouriteFoodColumn = "pet_favourite_food"
)

// Columns holds all SQL columns for pet fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Pet type.
var ForeignKeys = []string{
	"user_pets",
	"pet_favourite_food",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
