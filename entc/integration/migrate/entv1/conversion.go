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
	"entgo.io/ent/entc/integration/migrate/entv1/conversion"
)

// Conversion is the model entity for the Conversion schema.
type Conversion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Int8ToString holds the value of the "int8_to_string" field.
	Int8ToString int8 `json:"int8_to_string,omitempty"`
	// Uint8ToString holds the value of the "uint8_to_string" field.
	Uint8ToString uint8 `json:"uint8_to_string,omitempty"`
	// Int16ToString holds the value of the "int16_to_string" field.
	Int16ToString int16 `json:"int16_to_string,omitempty"`
	// Uint16ToString holds the value of the "uint16_to_string" field.
	Uint16ToString uint16 `json:"uint16_to_string,omitempty"`
	// Int32ToString holds the value of the "int32_to_string" field.
	Int32ToString int32 `json:"int32_to_string,omitempty"`
	// Uint32ToString holds the value of the "uint32_to_string" field.
	Uint32ToString uint32 `json:"uint32_to_string,omitempty"`
	// Int64ToString holds the value of the "int64_to_string" field.
	Int64ToString int64 `json:"int64_to_string,omitempty"`
	// Uint64ToString holds the value of the "uint64_to_string" field.
	Uint64ToString uint64 `json:"uint64_to_string,omitempty"`
	selectValues   sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Conversion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case conversion.FieldID, conversion.FieldInt8ToString, conversion.FieldUint8ToString, conversion.FieldInt16ToString, conversion.FieldUint16ToString, conversion.FieldInt32ToString, conversion.FieldUint32ToString, conversion.FieldInt64ToString, conversion.FieldUint64ToString:
			values[i] = new(sql.NullInt64)
		case conversion.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Conversion fields.
func (c *Conversion) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case conversion.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case conversion.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case conversion.FieldInt8ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int8_to_string", values[i])
			} else if value.Valid {
				c.Int8ToString = int8(value.Int64)
			}
		case conversion.FieldUint8ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint8_to_string", values[i])
			} else if value.Valid {
				c.Uint8ToString = uint8(value.Int64)
			}
		case conversion.FieldInt16ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int16_to_string", values[i])
			} else if value.Valid {
				c.Int16ToString = int16(value.Int64)
			}
		case conversion.FieldUint16ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint16_to_string", values[i])
			} else if value.Valid {
				c.Uint16ToString = uint16(value.Int64)
			}
		case conversion.FieldInt32ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int32_to_string", values[i])
			} else if value.Valid {
				c.Int32ToString = int32(value.Int64)
			}
		case conversion.FieldUint32ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint32_to_string", values[i])
			} else if value.Valid {
				c.Uint32ToString = uint32(value.Int64)
			}
		case conversion.FieldInt64ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field int64_to_string", values[i])
			} else if value.Valid {
				c.Int64ToString = value.Int64
			}
		case conversion.FieldUint64ToString:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field uint64_to_string", values[i])
			} else if value.Valid {
				c.Uint64ToString = uint64(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Conversion.
// This includes values selected through modifiers, order, etc.
func (c *Conversion) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// Update returns a builder for updating this Conversion.
// Note that you need to call Conversion.Unwrap() before calling this method if this Conversion
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Conversion) Update() *ConversionUpdateOne {
	return NewConversionClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Conversion entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Conversion) Unwrap() *Conversion {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entv1: Conversion is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Conversion) String() string {
	var builder strings.Builder
	builder.WriteString("Conversion(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteString(", ")
	builder.WriteString("int8_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Int8ToString))
	builder.WriteString(", ")
	builder.WriteString("uint8_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Uint8ToString))
	builder.WriteString(", ")
	builder.WriteString("int16_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Int16ToString))
	builder.WriteString(", ")
	builder.WriteString("uint16_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Uint16ToString))
	builder.WriteString(", ")
	builder.WriteString("int32_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Int32ToString))
	builder.WriteString(", ")
	builder.WriteString("uint32_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Uint32ToString))
	builder.WriteString(", ")
	builder.WriteString("int64_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Int64ToString))
	builder.WriteString(", ")
	builder.WriteString("uint64_to_string=")
	builder.WriteString(fmt.Sprintf("%v", c.Uint64ToString))
	builder.WriteByte(')')
	return builder.String()
}

// Conversions is a parsable slice of Conversion.
type Conversions []*Conversion
