// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/multischema/ent/cleanuser"
)

// CleanUser is the model entity for the CleanUser schema.
type CleanUser struct {
	config `json:"-"`
	// Name holds the value of the "name" field.
	Name         string `json:"name,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CleanUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case cleanuser.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CleanUser fields.
func (cu *CleanUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case cleanuser.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				cu.Name = value.String
			}
		default:
			cu.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CleanUser.
// This includes values selected through modifiers, order, etc.
func (cu *CleanUser) Value(name string) (ent.Value, error) {
	return cu.selectValues.Get(name)
}

// Unwrap unwraps the CleanUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (cu *CleanUser) Unwrap() *CleanUser {
	_tx, ok := cu.config.driver.(*txDriver)
	if !ok {
		panic("ent: CleanUser is not a transactional entity")
	}
	cu.config.driver = _tx.drv
	return cu
}

// String implements the fmt.Stringer.
func (cu *CleanUser) String() string {
	var builder strings.Builder
	builder.WriteString("CleanUser(")
	builder.WriteString("name=")
	builder.WriteString(cu.Name)
	builder.WriteByte(')')
	return builder.String()
}

// CleanUsers is a parsable slice of CleanUser.
type CleanUsers []*CleanUser
