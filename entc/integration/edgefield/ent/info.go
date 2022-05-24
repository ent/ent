// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

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
	Edges InfoEdges `json:"edges"`
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
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Info) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case info.FieldContent:
			values[i] = new([]byte)
		case info.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Info", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Info fields.
func (i *Info) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case info.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case info.FieldContent:
			if value, ok := values[j].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field content", values[j])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &i.Content); err != nil {
					return fmt.Errorf("unmarshal field content: %w", err)
				}
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Info entity.
func (i *Info) QueryUser() *UserQuery {
	return (&InfoClient{config: i.config}).QueryUser(i)
}

// Update returns a builder for updating this Info.
// Note that you need to call Info.Unwrap() before calling this method if this Info
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Info) Update() *InfoUpdateOne {
	return (&InfoClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Info entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Info) Unwrap() *Info {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Info is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Info) String() string {
	var builder strings.Builder
	builder.WriteString("Info(")
	builder.WriteString(fmt.Sprintf("id=%v, ", i.ID))
	builder.WriteString("content=")
	builder.WriteString(fmt.Sprintf("%v", i.Content))
	builder.WriteByte(')')
	return builder.String()
}

// Infos is a parsable slice of Info.
type Infos []*Info

func (i Infos) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
