// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgefield/ent/info"
	"entgo.io/ent/entc/integration/edgefield/ent/user"
)

// Info is the model entity for the Info schema.
type Info struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Content holds the value of the "content" field.
	Content json.RawMessage `json:"content,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the InfoQuery when eager-loading is set.
	Edges        InfoEdges `json:"edges"`
	selectValues sql.SelectValues
}

// InfoEdges holds the relations/edges for other nodes in the graph.
type InfoEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e InfoEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Info) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case info.FieldContent:
			values[i] = new([]byte)
		case info.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Info fields.
func (_m *Info) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case info.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case info.FieldContent:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &_m.Content); err != nil {
					return fmt.Errorf("unmarshal field content: %w", err)
				}
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Info.
// This includes values selected through modifiers, order, etc.
func (_m *Info) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Info entity.
func (_m *Info) QueryUser() *UserQuery {
	return NewInfoClient(_m.config).QueryUser(_m)
}

// Update returns a builder for updating this Info.
// Note that you need to call Info.Unwrap() before calling this method if this Info
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Info) Update() *InfoUpdateOne {
	return NewInfoClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Info entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Info) Unwrap() *Info {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Info is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// TableName returns the table name of the Info in the database.
func (i *Info) TableName() string {
	return info.Table
}

// String implements the fmt.Stringer.
func (_m *Info) String() string {
	var builder strings.Builder
	builder.WriteString("Info(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", _m.Content))
	builder.WriteByte(')')
	return builder.String()
}

// Infos is a parsable slice of Info.
type Infos []*Info
