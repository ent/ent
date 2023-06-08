// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package builder

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the builder type in the database.
	Label = "builder"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the builder in the database.
	Table = "builders"
)

// Columns holds all SQL columns for builder fields.
var Columns = []string{
	FieldID,
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

// OrderOption defines the ordering options for the Builder queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// comment from another template.
