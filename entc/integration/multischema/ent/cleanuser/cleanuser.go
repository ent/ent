// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package cleanuser

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the cleanuser type in the database.
	Label = "clean_user"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// Table holds the table name of the cleanuser in the database.
	Table = "clean_users"
)

// Columns holds all SQL columns for cleanuser fields.
var Columns = []string{
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the CleanUser queries.
type OrderOption func(*sql.Selector)

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}
