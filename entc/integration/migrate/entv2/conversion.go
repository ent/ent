// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"database/sql"
	fmt "fmt"
	strings "strings"

	conversion "entgo.io/ent/entc/integration/migrate/entv2/conversion"
)

// Conversion is the model entity for the Conversion schema.
type Conversion struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Int8ToString holds the value of the "int8_to_string" field.
	Int8ToString string `json:"int8_to_string,omitempty"`
	// Uint8ToString holds the value of the "uint8_to_string" field.
	Uint8ToString string `json:"uint8_to_string,omitempty"`
	// Int16ToString holds the value of the "int16_to_string" field.
	Int16ToString string `json:"int16_to_string,omitempty"`
	// Uint16ToString holds the value of the "uint16_to_string" field.
	Uint16ToString string `json:"uint16_to_string,omitempty"`
	// Int32ToString holds the value of the "int32_to_string" field.
	Int32ToString string `json:"int32_to_string,omitempty"`
	// Uint32ToString holds the value of the "uint32_to_string" field.
	Uint32ToString string `json:"uint32_to_string,omitempty"`
	// Int64ToString holds the value of the "int64_to_string" field.
	Int64ToString string `json:"int64_to_string,omitempty"`
	// Uint64ToString holds the value of the "uint64_to_string" field.
	Uint64ToString string `json:"uint64_to_string,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Conversion) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case conversion.FieldID:
			values[i] = new(sql.NullInt64)
		case conversion.FieldName, conversion.FieldInt8ToString, conversion.FieldUint8ToString, conversion.FieldInt16ToString, conversion.FieldUint16ToString, conversion.FieldInt32ToString, conversion.FieldUint32ToString, conversion.FieldInt64ToString, conversion.FieldUint64ToString:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Conversion", columns[i])
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
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field int8_to_string", values[i])
			} else if value.Valid {
				c.Int8ToString = value.String
			}
		case conversion.FieldUint8ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uint8_to_string", values[i])
			} else if value.Valid {
				c.Uint8ToString = value.String
			}
		case conversion.FieldInt16ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field int16_to_string", values[i])
			} else if value.Valid {
				c.Int16ToString = value.String
			}
		case conversion.FieldUint16ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uint16_to_string", values[i])
			} else if value.Valid {
				c.Uint16ToString = value.String
			}
		case conversion.FieldInt32ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field int32_to_string", values[i])
			} else if value.Valid {
				c.Int32ToString = value.String
			}
		case conversion.FieldUint32ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uint32_to_string", values[i])
			} else if value.Valid {
				c.Uint32ToString = value.String
			}
		case conversion.FieldInt64ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field int64_to_string", values[i])
			} else if value.Valid {
				c.Int64ToString = value.String
			}
		case conversion.FieldUint64ToString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field uint64_to_string", values[i])
			} else if value.Valid {
				c.Uint64ToString = value.String
			}
		}
	}
	return nil
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
		panic("entv2: Conversion is not a transactional entity")
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
	builder.WriteString(c.Int8ToString)
	builder.WriteString(", ")
	builder.WriteString("uint8_to_string=")
	builder.WriteString(c.Uint8ToString)
	builder.WriteString(", ")
	builder.WriteString("int16_to_string=")
	builder.WriteString(c.Int16ToString)
	builder.WriteString(", ")
	builder.WriteString("uint16_to_string=")
	builder.WriteString(c.Uint16ToString)
	builder.WriteString(", ")
	builder.WriteString("int32_to_string=")
	builder.WriteString(c.Int32ToString)
	builder.WriteString(", ")
	builder.WriteString("uint32_to_string=")
	builder.WriteString(c.Uint32ToString)
	builder.WriteString(", ")
	builder.WriteString("int64_to_string=")
	builder.WriteString(c.Int64ToString)
	builder.WriteString(", ")
	builder.WriteString("uint64_to_string=")
	builder.WriteString(c.Uint64ToString)
	builder.WriteByte(')')
	return builder.String()
}

// Conversions is a parsable slice of Conversion.
type Conversions []*Conversion

func (c Conversions) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
