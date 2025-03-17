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
	"entgo.io/ent/entc/integration/edgeschema/ent/friendship"
	"entgo.io/ent/entc/integration/edgeschema/ent/user"
)

// Friendship is the model entity for the Friendship schema.
type Friendship struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Weight holds the value of the "weight" field.
	Weight int `json:"weight,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// FriendID holds the value of the "friend_id" field.
	FriendID int `json:"friend_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FriendshipQuery when eager-loading is set.
	Edges        FriendshipEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FriendshipEdges holds the relations/edges for other nodes in the graph.
type FriendshipEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Friend holds the value of the friend edge.
	Friend *User `json:"friend,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) UserOrErr() (*User, error) {
	if e.User != nil {
		return e.User, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "user"}
}

// FriendOrErr returns the Friend value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FriendshipEdges) FriendOrErr() (*User, error) {
	if e.Friend != nil {
		return e.Friend, nil
	} else if e.loadedTypes[1] {
		return nil, &NotFoundError{label: user.Label}
	}
	return nil, &NotLoadedError{edge: "friend"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Friendship) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case friendship.FieldID, friendship.FieldWeight, friendship.FieldUserID, friendship.FieldFriendID:
			values[i] = new(sql.NullInt64)
		case friendship.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Friendship fields.
func (_m *Friendship) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case friendship.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case friendship.FieldWeight:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field weight", values[i])
			} else if value.Valid {
				_m.Weight = int(value.Int64)
			}
		case friendship.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				_m.CreatedAt = value.Time
			}
		case friendship.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				_m.UserID = int(value.Int64)
			}
		case friendship.FieldFriendID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field friend_id", values[i])
			} else if value.Valid {
				_m.FriendID = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Friendship.
// This includes values selected through modifiers, order, etc.
func (_m *Friendship) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryUser queries the "user" edge of the Friendship entity.
func (_m *Friendship) QueryUser() *UserQuery {
	return NewFriendshipClient(_m.config).QueryUser(_m)
}

// QueryFriend queries the "friend" edge of the Friendship entity.
func (_m *Friendship) QueryFriend() *UserQuery {
	return NewFriendshipClient(_m.config).QueryFriend(_m)
}

// Update returns a builder for updating this Friendship.
// Note that you need to call Friendship.Unwrap() before calling this method if this Friendship
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Friendship) Update() *FriendshipUpdateOne {
	return NewFriendshipClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Friendship entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Friendship) Unwrap() *Friendship {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Friendship is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Friendship) String() string {
	var builder strings.Builder
	builder.WriteString("Friendship(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("weight=")
	builder.WriteString(fmt.Sprintf("%v", _m.Weight))
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(_m.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("user_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.UserID))
	builder.WriteString(", ")
	builder.WriteString("friend_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.FriendID))
	builder.WriteByte(')')
	return builder.String()
}

// Friendships is a parsable slice of Friendship.
type Friendships []*Friendship
