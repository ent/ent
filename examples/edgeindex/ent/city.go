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
	"entgo.io/ent/examples/edgeindex/ent/city"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CityQuery when eager-loading is set.
	Edges        CityEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CityEdges holds the relations/edges for other nodes in the graph.
type CityEdges struct {
	// Streets holds the value of the streets edge.
	Streets []*Street `json:"streets,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// StreetsOrErr returns the Streets value or an error if the edge
// was not loaded in eager-loading.
func (e CityEdges) StreetsOrErr() ([]*Street, error) {
	if e.loadedTypes[0] {
		return e.Streets, nil
	}
	return nil, &NotLoadedError{edge: "streets"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*City) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case city.FieldID:
			values[i] = new(sql.NullInt64)
		case city.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the City fields.
func (m *City) assignValues(columns []string, values []any) error {
	if v, c := len(values), len(columns); v < c {
		return fmt.Errorf("mismatch number of scan values: %d != %d", v, c)
	}
	for i := range columns {
		switch columns[i] {
		case city.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			m.ID = int(value.Int64)
		case city.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				m.Name = value.String
			}
		default:
			m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the City.
// This includes values selected through modifiers, order, etc.
func (m *City) Value(name string) (ent.Value, error) {
	return m.selectValues.Get(name)
}

// QueryStreets queries the "streets" edge of the City entity.
func (m *City) QueryStreets() *StreetQuery {
	return NewCityClient(m.config).QueryStreets(m)
}

// Update returns a builder for updating this City.
// Note that you need to call City.Unwrap() before calling this method if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *City) Update() *CityUpdateOne {
	return NewCityClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the City entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *City) Unwrap() *City {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *City) String() string {
	var builder strings.Builder
	builder.WriteString("City(")
	builder.WriteString(fmt.Sprintf("id=%v, ", m.ID))
	builder.WriteString("name=")
	builder.WriteString(m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cities is a parsable slice of City.
type Cities []*City
