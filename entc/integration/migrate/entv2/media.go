// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/media"
)

// Comment that appears in both the schema and the generated code
type Media struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Source holds the value of the "source" field.
	Source string `json:"source,omitempty"`
	// source_ui text
	SourceURI string `json:"source_uri,omitempty"`
	// media text
	Text         string `json:"text,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Media) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case media.FieldID:
			values[i] = new(sql.NullInt64)
		case media.FieldSource, media.FieldSourceURI, media.FieldText:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Media fields.
func (_m *Media) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case media.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case media.FieldSource:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source", values[i])
			} else if value.Valid {
				_m.Source = value.String
			}
		case media.FieldSourceURI:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field source_uri", values[i])
			} else if value.Valid {
				_m.SourceURI = value.String
			}
		case media.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				_m.Text = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Media.
// This includes values selected through modifiers, order, etc.
func (_m *Media) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// Update returns a builder for updating this Media.
// Note that you need to call Media.Unwrap() before calling this method if this Media
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Media) Update() *MediaUpdateOne {
	return NewMediaClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Media entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Media) Unwrap() *Media {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Media is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// TableName returns the table name of the Media in the database.
func (m *Media) TableName() string {
	return media.Table
}

// String implements the fmt.Stringer.
func (_m *Media) String() string {
	var builder strings.Builder
	builder.WriteString("Media(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("source=")
	builder.WriteString(_m.Source)
	builder.WriteString(", ")
	builder.WriteString("source_uri=")
	builder.WriteString(_m.SourceURI)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(_m.Text)
	builder.WriteByte(')')
	return builder.String()
}

// MediaSlice is a parsable slice of Media.
type MediaSlice []*Media
