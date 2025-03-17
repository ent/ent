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
	"entgo.io/ent/examples/privacytenant/ent/group"
	"entgo.io/ent/examples/privacytenant/ent/tenant"
)

// Group is the model entity for the Group schema.
type Group struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// TenantID holds the value of the "tenant_id" field.
	TenantID int `json:"tenant_id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupQuery when eager-loading is set.
	Edges        GroupEdges `json:"edges"`
	selectValues sql.SelectValues
}

// GroupEdges holds the relations/edges for other nodes in the graph.
type GroupEdges struct {
	// Tenant holds the value of the tenant edge.
	Tenant *Tenant `json:"tenant,omitempty"`
	// Users holds the value of the users edge.
	Users []*User `json:"users,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// TenantOrErr returns the Tenant value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e GroupEdges) TenantOrErr() (*Tenant, error) {
	if e.Tenant != nil {
		return e.Tenant, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: tenant.Label}
	}
	return nil, &NotLoadedError{edge: "tenant"}
}

// UsersOrErr returns the Users value or an error if the edge
// was not loaded in eager-loading.
func (e GroupEdges) UsersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Users, nil
	}
	return nil, &NotLoadedError{edge: "users"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Group) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case group.FieldID, group.FieldTenantID:
			values[i] = new(sql.NullInt64)
		case group.FieldName:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Group fields.
func (_m *Group) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case group.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case group.FieldTenantID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tenant_id", values[i])
			} else if value.Valid {
				_m.TenantID = int(value.Int64)
			}
		case group.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				_m.Name = value.String
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Group.
// This includes values selected through modifiers, order, etc.
func (_m *Group) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryTenant queries the "tenant" edge of the Group entity.
func (_m *Group) QueryTenant() *TenantQuery {
	return NewGroupClient(_m.config).QueryTenant(_m)
}

// QueryUsers queries the "users" edge of the Group entity.
func (_m *Group) QueryUsers() *UserQuery {
	return NewGroupClient(_m.config).QueryUsers(_m)
}

// Update returns a builder for updating this Group.
// Note that you need to call Group.Unwrap() before calling this method if this Group
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Group) Update() *GroupUpdateOne {
	return NewGroupClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Group entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Group) Unwrap() *Group {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Group is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Group) String() string {
	var builder strings.Builder
	builder.WriteString("Group(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("tenant_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.TenantID))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(_m.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Groups is a parsable slice of Group.
type Groups []*Group
