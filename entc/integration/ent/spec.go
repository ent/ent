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
	"entgo.io/ent/entc/integration/ent/spec"
)

// Spec is the model entity for the Spec schema.
type Spec struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the SpecQuery when eager-loading is set.
	Edges        SpecEdges `json:"edges"`
	selectValues sql.SelectValues
}

// SpecEdges holds the relations/edges for other nodes in the graph.
type SpecEdges struct {
	// Card holds the value of the card edge.
	Card []*Card `json:"card,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	namedCard   map[string][]*Card
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading.
func (e SpecEdges) CardOrErr() ([]*Card, error) {
	if e.loadedTypes[0] {
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Spec) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case spec.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Spec fields.
func (_m *Spec) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case spec.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Spec.
// This includes values selected through modifiers, order, etc.
func (_m *Spec) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryCard queries the "card" edge of the Spec entity.
func (_m *Spec) QueryCard() *CardQuery {
	return NewSpecClient(_m.config).QueryCard(_m)
}

// Update returns a builder for updating this Spec.
// Note that you need to call Spec.Unwrap() before calling this method if this Spec
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Spec) Update() *SpecUpdateOne {
	return NewSpecClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Spec entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Spec) Unwrap() *Spec {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Spec is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// TableName returns the table name of the Spec in the database.
func (s *Spec) TableName() string {
	return spec.Table
}

// String implements the fmt.Stringer.
func (_m *Spec) String() string {
	var builder strings.Builder
	builder.WriteString("Spec(")
	builder.WriteString(fmt.Sprintf("id=%v", _m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// NamedCard returns the Card named value or an error if the edge was not
// loaded in eager-loading with this name.
func (_m *Spec) NamedCard(name string) ([]*Card, error) {
	if _m.Edges.namedCard == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := _m.Edges.namedCard[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (_m *Spec) appendNamedCard(name string, edges ...*Card) {
	if _m.Edges.namedCard == nil {
		_m.Edges.namedCard = make(map[string][]*Card)
	}
	if len(edges) == 0 {
		_m.Edges.namedCard[name] = []*Card{}
	} else {
		_m.Edges.namedCard[name] = append(_m.Edges.namedCard[name], edges...)
	}
}

// Specs is a parsable slice of Spec.
type Specs []*Spec
