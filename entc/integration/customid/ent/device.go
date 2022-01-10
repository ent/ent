// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
)

// Device is the model entity for the Device schema.
type Device struct {
	config
	// ID of the ent.
	ID schema.ID `json:"id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DeviceQuery when eager-loading is set.
	Edges                 DeviceEdges `json:"edges"`
	device_active_session *schema.ID
}

// DeviceEdges holds the relations/edges for other nodes in the graph.
type DeviceEdges struct {
	// ActiveSession holds the value of the active_session edge.
	ActiveSession *Session `json:"active_session,omitempty"`
	// Sessions holds the value of the sessions edge.
	Sessions []*Session `json:"sessions,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// ActiveSessionOrErr returns the ActiveSession value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DeviceEdges) ActiveSessionOrErr() (*Session, error) {
	if e.loadedTypes[0] {
		if e.ActiveSession == nil {
			// The edge active_session was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: session.Label}
		}
		return e.ActiveSession, nil
	}
	return nil, &NotLoadedError{edge: "active_session"}
}

// SessionsOrErr returns the Sessions value or an error if the edge
// was not loaded in eager-loading.
func (e DeviceEdges) SessionsOrErr() ([]*Session, error) {
	if e.loadedTypes[1] {
		return e.Sessions, nil
	}
	return nil, &NotLoadedError{edge: "sessions"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Device) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			values[i] = new(schema.ID)
		case device.ForeignKeys[0]: // device_active_session
			values[i] = &sql.NullScanner{S: new(schema.ID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Device", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Device fields.
func (d *Device) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case device.FieldID:
			if value, ok := values[i].(*schema.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				d.ID = *value
			}
		case device.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field device_active_session", values[i])
			} else if value.Valid {
				d.device_active_session = new(schema.ID)
				*d.device_active_session = *value.S.(*schema.ID)
			}
		}
	}
	return nil
}

// QueryActiveSession queries the "active_session" edge of the Device entity.
func (d *Device) QueryActiveSession() *SessionQuery {
	return (&DeviceClient{config: d.config}).QueryActiveSession(d)
}

// QuerySessions queries the "sessions" edge of the Device entity.
func (d *Device) QuerySessions() *SessionQuery {
	return (&DeviceClient{config: d.config}).QuerySessions(d)
}

// Update returns a builder for updating this Device.
// Note that you need to call Device.Unwrap() before calling this method if this Device
// was returned from a transaction, and the transaction was committed or rolled back.
func (d *Device) Update() *DeviceUpdateOne {
	return (&DeviceClient{config: d.config}).UpdateOne(d)
}

// Unwrap unwraps the Device entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (d *Device) Unwrap() *Device {
	tx, ok := d.config.driver.(*txDriver)
	if !ok {
		panic("ent: Device is not a transactional entity")
	}
	d.config.driver = tx.drv
	return d
}

// TableName returns the table name of the Device in the database.
func (d *Device) TableName() string {
	return device.Table
}

// String implements the fmt.Stringer.
func (d *Device) String() string {
	var builder strings.Builder
	builder.WriteString("Device(")
	builder.WriteString(fmt.Sprintf("id=%v", d.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Devices is a parsable slice of Device.
type Devices []*Device

func (d Devices) config(cfg config) {
	for _i := range d {
		d[_i].config = cfg
	}
}
