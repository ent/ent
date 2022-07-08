// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/groupinfo"
)

// GroupInfo is the model entity for the GroupInfo schema.
type GroupInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Desc holds the value of the "desc" field.
	Desc string `json:"desc,omitempty"`
	// MaxUsers holds the value of the "max_users" field.
	MaxUsers int `json:"max_users,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the GroupInfoQuery when eager-loading is set.
	Edges GroupInfoEdges `json:"edges"`
}

// GroupInfoEdges holds the relations/edges for other nodes in the graph.
type GroupInfoEdges struct {
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e GroupInfoEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[0] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*GroupInfo) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case groupinfo.FieldID, groupinfo.FieldMaxUsers:
			values[i] = new(sql.NullInt64)
		case groupinfo.FieldDesc:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type GroupInfo", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the GroupInfo fields.
func (gi *GroupInfo) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case groupinfo.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			gi.ID = int(value.Int64)
		case groupinfo.FieldDesc:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field desc", values[i])
			} else if value.Valid {
				gi.Desc = value.String
			}
		case groupinfo.FieldMaxUsers:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field max_users", values[i])
			} else if value.Valid {
				gi.MaxUsers = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryGroups queries the "groups" edge of the GroupInfo entity.
func (gi *GroupInfo) QueryGroups() *GroupQuery {
	return (&GroupInfoClient{config: gi.config}).QueryGroups(gi)
}

// Update returns a builder for updating this GroupInfo.
// Note that you need to call GroupInfo.Unwrap() before calling this method if this GroupInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (gi *GroupInfo) Update() *GroupInfoUpdateOne {
	return (&GroupInfoClient{config: gi.config}).UpdateOne(gi)
}

// Unwrap unwraps the GroupInfo entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (gi *GroupInfo) Unwrap() *GroupInfo {
	_tx, ok := gi.config.driver.(*txDriver)
	if !ok {
		panic("ent: GroupInfo is not a transactional entity")
	}
	gi.config.driver = _tx.drv
	return gi
}

// String implements the fmt.Stringer.
func (gi *GroupInfo) String() string {
	var builder strings.Builder
	builder.WriteString("GroupInfo(")
	builder.WriteString(fmt.Sprintf("id=%v, ", gi.ID))
	builder.WriteString("desc=")
	builder.WriteString(gi.Desc)
	builder.WriteString(", ")
	builder.WriteString("max_users=")
	builder.WriteString(fmt.Sprintf("%v", gi.MaxUsers))
	builder.WriteByte(')')
	return builder.String()
}

// GroupInfos is a parsable slice of GroupInfo.
type GroupInfos []*GroupInfo

func (gi GroupInfos) config(cfg config) {
	for _i := range gi {
		gi[_i].config = cfg
	}
}
