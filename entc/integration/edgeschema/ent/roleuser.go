// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/role"
	"entgo.io/ent/entc/integration/edgeschema/ent/roleuser"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// RoleUser is the model entity for the RoleUser schema.
type RoleUser struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// RoleID holds the value of the "role_id" field.
	RoleID int `json:"role_id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoleUserQuery when eager-loading is set.
	Edges        RoleUserEdges `json:"edges"`
	selectValues sql.SelectValues
}

// RoleUserEdges holds the relations/edges for other nodes in the graph.
type RoleUserEdges struct {
	// Role holds the value of the role edge.
	Role *Role `json:"role,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoleOrErr returns the Role value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleUserEdges) RoleOrErr() (*Role, error) {
	if e.loadedTypes[0] {
		if e.Role == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: role.Label}
		}
		return e.Role, nil
	}
	return nil, &NotLoadedError{edge: "role"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoleUserEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoleUser) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roleuser.FieldRoleID, roleuser.FieldUserID:
			values[i] = new(sql.NullInt64)
		case roleuser.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoleUser fields.
func (ru *RoleUser) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roleuser.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				ru.CreatedAt = value.Time
			}
		case roleuser.FieldRoleID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field role_id", values[i])
			} else if value.Valid {
				ru.RoleID = int(value.Int64)
			}
		case roleuser.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ru.UserID = int(value.Int64)
			}
		default:
			ru.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoleUser.
// This includes values selected through modifiers, order, etc.
func (ru *RoleUser) Value(name string) (ent.Value, error) {
	return ru.selectValues.Get(name)
}

// QueryRole queries the "role" edge of the RoleUser entity.
func (ru *RoleUser) QueryRole() *RoleQuery {
	return NewRoleUserClient(ru.config).QueryRole(ru)
}

// QueryUser queries the "user" edge of the RoleUser entity.
func (ru *RoleUser) QueryUser() *UserQuery {
	return NewRoleUserClient(ru.config).QueryUser(ru)
}

// Update returns a builder for updating this RoleUser.
// Note that you need to call RoleUser.Unwrap() before calling this method if this RoleUser
// was returned from a transaction, and the transaction was committed or rolled back.
func (ru *RoleUser) Update() *RoleUserUpdateOne {
	return NewRoleUserClient(ru.config).UpdateOne(ru)
}

// Unwrap unwraps the RoleUser entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ru *RoleUser) Unwrap() *RoleUser {
	_tx, ok := ru.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoleUser is not a transactional entity")
	}
	ru.config.driver = _tx.drv
	return ru
}

// String implements the fmt.Stringer.
func (ru *RoleUser) String() string {
	var builder strings.Builder
	builder.WriteString("RoleUser(")
	builder.WriteString("created_at=")
	builder.WriteString(ru.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("role_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.RoleID))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ru.UserID))
	builder.WriteByte(')')
	return builder.String()
}

// RoleUsers is a parsable slice of RoleUser.
type RoleUsers []*RoleUser

// Len returns length of RoleUsers.
func (ru RoleUsers) Len() int { return len(ru) }
