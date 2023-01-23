// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/edgeschema/ent/group"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
	"entgo.io/ent/entc/integration/edgeschema/ent/usergroup"
)

// UserGroup is the model entity for the UserGroup schema.
type UserGroup struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// JoinedAt holds the value of the "joined_at" field.
	JoinedAt time.Time `json:"joined_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// GroupID holds the value of the "group_id" field.
	GroupID int `json:"group_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserGroupQuery when eager-loading is set.
	Edges UserGroupEdges `json:"edges"`
}

// UserGroupEdges holds the relations/edges for other nodes in the graph.
type UserGroupEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Group holds the value of the group edge.
	Group *Group `json:"group,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserGroupEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// GroupOrErr returns the Group value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserGroupEdges) GroupOrErr() (*Group, error) {
	if e.loadedTypes[1] {
		if e.Group == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: group.Label}
		}
		return e.Group, nil
	}
	return nil, &NotLoadedError{edge: "group"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*UserGroup) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case usergroup.FieldID, usergroup.FieldUserID, usergroup.FieldGroupID:
			values[i] = new(sql.NullInt64)
		case usergroup.FieldJoinedAt:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type UserGroup", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the UserGroup fields.
func (ug *UserGroup) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case usergroup.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ug.ID = int(value.Int64)
		case usergroup.FieldJoinedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field joined_at", values[i])
			} else if value.Valid {
				ug.JoinedAt = value.Time
			}
		case usergroup.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				ug.UserID = int(value.Int64)
			}
		case usergroup.FieldGroupID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field group_id", values[i])
			} else if value.Valid {
				ug.GroupID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the UserGroup entity.
func (ug *UserGroup) QueryUser() *UserQuery {
	return NewUserGroupClient(ug.config).QueryUser(ug)
}

// QueryGroup queries the "group" edge of the UserGroup entity.
func (ug *UserGroup) QueryGroup() *GroupQuery {
	return NewUserGroupClient(ug.config).QueryGroup(ug)
}

// Update returns a builder for updating this UserGroup.
// Note that you need to call UserGroup.Unwrap() before calling this method if this UserGroup
// was returned from a transaction, and the transaction was committed or rolled back.
func (ug *UserGroup) Update() *UserGroupUpdateOne {
	return NewUserGroupClient(ug.config).UpdateOne(ug)
}

// Unwrap unwraps the UserGroup entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ug *UserGroup) Unwrap() *UserGroup {
	_tx, ok := ug.config.driver.(*txDriver)
	if !ok {
		panic("ent: UserGroup is not a transactional entity")
	}
	ug.config.driver = _tx.drv
	return ug
}

// String implements the fmt.Stringer.
func (ug *UserGroup) String() string {
	var builder strings.Builder
	builder.WriteString("UserGroup(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ug.ID))
	builder.WriteString("joined_at=")
	builder.WriteString(ug.JoinedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", ug.UserID))
	builder.WriteString(", ")
	builder.WriteString("group_id=")
	builder.WriteString(fmt.Sprintf("%v", ug.GroupID))
	builder.WriteByte(')')
	return builder.String()
}

// UserGroups is a parsable slice of UserGroup.
type UserGroups []*UserGroup

func (ug UserGroups) config(cfg config) {
	for _i := range ug {
		ug[_i].config = cfg
	}
}

func (ug UserGroups) IDs() []int {
	ids := make([]int, len(ug))
	for _i := range ug {
		ids[_i] = ug[_i].ID
	}
	return ids
}
