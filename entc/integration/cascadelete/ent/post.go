// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"

	"entgo.io/ent/entc/integration/cascadelete/ent/post"
	"entgo.io/ent/entc/integration/cascadelete/ent/user"
)

// Post is the model entity for the Post schema.
type Post struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// AuthorID holds the value of the "author_id" field.
	AuthorID int `json:"author_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PostQuery when eager-loading is set.
	Edges PostEdges `json:"edges"`
}

// PostEdges holds the relations/edges for other nodes in the graph.
type PostEdges struct {
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PostEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e PostEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Post) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case post.FieldID, post.FieldAuthorID:
			values[i] = new(sql.NullInt64)
		case post.FieldText:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Post", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Post fields.
func (po *Post) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case post.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			po.ID = int(value.Int64)
		case post.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				po.Text = value.String
			}
		case post.FieldAuthorID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field author_id", values[i])
			} else if value.Valid {
				po.AuthorID = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAuthor queries the "author" edge of the Post entity.
func (po *Post) QueryAuthor() *UserQuery {
	return (&PostClient{config: po.config}).QueryAuthor(po)
}

// QueryComments queries the "comments" edge of the Post entity.
func (po *Post) QueryComments() *CommentQuery {
	return (&PostClient{config: po.config}).QueryComments(po)
}

// Update returns a builder for updating this Post.
// Note that you need to call Post.Unwrap() before calling this method if this Post
// was returned from a transaction, and the transaction was committed or rolled back.
func (po *Post) Update() *PostUpdateOne {
	return (&PostClient{config: po.config}).UpdateOne(po)
}

// Unwrap unwraps the Post entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (po *Post) Unwrap() *Post {
	tx, ok := po.config.driver.(*txDriver)
	if !ok {
		panic("ent: Post is not a transactional entity")
	}
	po.config.driver = tx.drv
	return po
}

// String implements the fmt.Stringer.
func (po *Post) String() string {
	var builder strings.Builder
	builder.WriteString("Post(")
	builder.WriteString(fmt.Sprintf("id=%v", po.ID))
	builder.WriteString(", text=")
	builder.WriteString(po.Text)
	builder.WriteString(", author_id=")
	builder.WriteString(fmt.Sprintf("%v", po.AuthorID))
	builder.WriteByte(')')
	return builder.String()
}

// Posts is a parsable slice of Post.
type Posts []*Post

func (po Posts) config(cfg config) {
	for _i := range po {
		po[_i].config = cfg
	}
}
