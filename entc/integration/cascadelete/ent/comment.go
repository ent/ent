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
	"entgo.io/ent/entc/integration/cascadelete/ent/comment"
	"entgo.io/ent/entc/integration/cascadelete/ent/post"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// PostID holds the value of the "post_id" field.
	PostID int `json:"post_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CommentQuery when eager-loading is set.
	Edges        CommentEdges `json:"edges"`
	selectValues sql.SelectValues
}

// CommentEdges holds the relations/edges for other nodes in the graph.
type CommentEdges struct {
	// Post holds the value of the post edge.
	Post *Post `json:"post,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// PostOrErr returns the Post value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CommentEdges) PostOrErr() (*Post, error) {
	if e.Post != nil {
		return e.Post, nil
	} else if e.loadedTypes[0] {
		return nil, &NotFoundError{label: post.Label}
	}
	return nil, &NotLoadedError{edge: "post"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Comment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case comment.FieldID, comment.FieldPostID:
			values[i] = new(sql.NullInt64)
		case comment.FieldText:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Comment fields.
func (_m *Comment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case comment.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_m.ID = int(value.Int64)
		case comment.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				_m.Text = value.String
			}
		case comment.FieldPostID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field post_id", values[i])
			} else if value.Valid {
				_m.PostID = int(value.Int64)
			}
		default:
			_m.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Comment.
// This includes values selected through modifiers, order, etc.
func (_m *Comment) Value(name string) (ent.Value, error) {
	return _m.selectValues.Get(name)
}

// QueryPost queries the "post" edge of the Comment entity.
func (_m *Comment) QueryPost() *PostQuery {
	return NewCommentClient(_m.config).QueryPost(_m)
}

// Update returns a builder for updating this Comment.
// Note that you need to call Comment.Unwrap() before calling this method if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (_m *Comment) Update() *CommentUpdateOne {
	return NewCommentClient(_m.config).UpdateOne(_m)
}

// Unwrap unwraps the Comment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_m *Comment) Unwrap() *Comment {
	_tx, ok := _m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	_m.config.driver = _tx.drv
	return _m
}

// String implements the fmt.Stringer.
func (_m *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", _m.ID))
	builder.WriteString("text=")
	builder.WriteString(_m.Text)
	builder.WriteString(", ")
	builder.WriteString("post_id=")
	builder.WriteString(fmt.Sprintf("%v", _m.PostID))
	builder.WriteByte(')')
	return builder.String()
}

// Comments is a parsable slice of Comment.
type Comments []*Comment
