// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv1

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv1/customtype"
)

// CustomType is the model entity for the CustomType schema.
type CustomType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Custom holds the value of the "custom" field.
	Custom       string `json:"custom,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*CustomType) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case customtype.FieldID:
			values[i] = new(sql.NullInt64)
		case customtype.FieldCustom:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the CustomType fields.
func (m *CustomType) assignValues(columns []string, values []any) error {
	if v, c := len(values), len(columns); v < c {
		return fmt.Errorf("mismatch number of scan values: %d != %d", v, c)
	}
	for i := range columns {
		switch columns[i] {
		case customtype.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case customtype.FieldCustom:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field custom", values[i])
			} else if value.Valid {
				m.Custom = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the CustomType.
// This includes values selected through modifiers, order, etc.
func (m *CustomType) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// Update returns a builder for updating this CustomType.
// Note that you need to call CustomType.Unwrap() before calling this method if this CustomType
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *CustomType) Update() *CustomTypeUpdateOne {
	return NewCustomTypeClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the CustomType entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *CustomType) Unwrap() *CustomType {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("entv1: CustomType is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *CustomType) String() string {
	var builder strings.Builder
	builder.WriteString("CustomType(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("custom=")
	builder.WriteString(m.Custom)
	builder.WriteByte(')')
	return builder.String()
}

// CustomTypes is a parsable slice of CustomType.
type CustomTypes []*CustomType
