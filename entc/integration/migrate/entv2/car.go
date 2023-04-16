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
	"entgo.io/ent/entc/integration/migrate/entv2/car"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
)

// Car is the model entity for the Car schema.
type Car struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CarQuery when eager-loading is set.
	Edges        CarEdges `json:"edges"`
	user_car     *int
	selectValues sql.SelectValues
}

// CarEdges holds the relations/edges for other nodes in the graph.
type CarEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CarEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Car) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			values[i] = new(sql.NullInt64)
		case car.FieldName:
			values[i] = new(sql.NullString)
		case car.ForeignKeys[0]: // user_car
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Car fields.
func (c *Car) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case car.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case car.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				c.Name = value.String
			}
		case car.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field user_car", value)
			} else if value.Valid {
				c.user_car = new(int)
				*c.user_car = int(value.Int64)
			}
		default:
			c.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Car.
// This includes values selected through modifiers, order, etc.
func (c *Car) Value(name string) (ent.Value, error) {
	return c.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Car entity.
func (c *Car) QueryOwner() *UserQuery {
	return NewCarClient(c.config).QueryOwner(c)
}

// Update returns a builder for updating this Car.
// Note that you need to call Car.Unwrap() before calling this method if this Car
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Car) Update() *CarUpdateOne {
	return NewCarClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Car entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Car) Unwrap() *Car {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("entv2: Car is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Car) String() string {
	var builder strings.Builder
	builder.WriteString("Car(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Cars is a parsable slice of Car.
type Cars []*Car

// Len returns length of Cars.
func (c Cars) Len() int { return len(c) }
