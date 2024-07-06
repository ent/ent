// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package file

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the file type in the database.
	Label = "file"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeProcesses holds the string denoting the processes edge name in mutations.
	EdgeProcesses = "processes"
	// Table holds the table name of the file in the database.
	Table = "files"
	// ProcessesTable is the table that holds the processes relation/edge. The primary key declared below.
	ProcessesTable = "attached_files"
	// ProcessesInverseTable is the table name for the Process entity.
	// It exists in this package in order to avoid circular dependency with the "process" package.
	ProcessesInverseTable = "processes"
)

// Columns holds all SQL columns for file fields.
var Columns = []string{
	FieldID,
	FieldName,
}

var (
	// ProcessesPrimaryKey and ProcessesColumn2 are the table columns denoting the
	// primary key for the processes relation (M2M).
	ProcessesPrimaryKey = []string{"proc_id", "f_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// OrderOption defines the ordering options for the File queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByProcessesCount orders the results by processes count.
func ByProcessesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProcessesStep(), opts...)
	}
}

// ByProcesses orders the results by processes terms.
func ByProcesses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProcessesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

func newProcessesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProcessesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ProcessesTable, ProcessesPrimaryKey...),
	)
}
