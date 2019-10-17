// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UniqueInt holds the value of the "unique_int" field.
	UniqueInt int `json:"unique_int,omitempty"`
	// UniqueFloat holds the value of the "unique_float" field.
	UniqueFloat float64 `json:"unique_float,omitempty"`
	// NillableInt holds the value of the "nillable_int" field.
	NillableInt *int `json:"nillable_int,omitempty"`
}

// FromRows scans the sql response data into Comment.
func (c *Comment) FromRows(rows *sql.Rows) error {
	var vc struct {
		ID          int
		UniqueInt   sql.NullInt64
		UniqueFloat sql.NullFloat64
		NillableInt sql.NullInt64
	}
	// the order here should be the same as in the `comment.Columns`.
	if err := rows.Scan(
		&vc.ID,
		&vc.UniqueInt,
		&vc.UniqueFloat,
		&vc.NillableInt,
	); err != nil {
		return err
	}
	c.ID = strconv.Itoa(vc.ID)
	c.UniqueInt = int(vc.UniqueInt.Int64)
	c.UniqueFloat = vc.UniqueFloat.Float64
	if vc.NillableInt.Valid {
		c.NillableInt = new(int)
		*c.NillableInt = int(vc.NillableInt.Int64)
	}
	return nil
}

// FromResponse scans the gremlin response data into Comment.
func (c *Comment) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc struct {
		ID          string  `json:"id,omitempty"`
		UniqueInt   int     `json:"unique_int,omitempty"`
		UniqueFloat float64 `json:"unique_float,omitempty"`
		NillableInt *int    `json:"nillable_int,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	c.ID = vc.ID
	c.UniqueInt = vc.UniqueInt
	c.UniqueFloat = vc.UniqueFloat
	c.NillableInt = vc.NillableInt
	return nil
}

// Update returns a builder for updating this Comment.
// Note that, you need to call Comment.Unwrap() before calling this method, if this Comment
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Comment) Update() *CommentUpdateOne {
	return (&CommentClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Comment) Unwrap() *Comment {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Comment is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Comment) String() string {
	var builder strings.Builder
	builder.WriteString("Comment(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", unique_int=")
	builder.WriteString(fmt.Sprintf("%v", c.UniqueInt))
	builder.WriteString(", unique_float=")
	builder.WriteString(fmt.Sprintf("%v", c.UniqueFloat))
	if v := c.NillableInt; v != nil {
		builder.WriteString(", nillable_int=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (c *Comment) id() int {
	id, _ := strconv.Atoi(c.ID)
	return id
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

// FromRows scans the sql response data into Comments.
func (c *Comments) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vc := &Comment{}
		if err := vc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, vc)
	}
	return nil
}

// FromResponse scans the gremlin response data into Comments.
func (c *Comments) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc []struct {
		ID          string  `json:"id,omitempty"`
		UniqueInt   int     `json:"unique_int,omitempty"`
		UniqueFloat float64 `json:"unique_float,omitempty"`
		NillableInt *int    `json:"nillable_int,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	for _, v := range vc {
		*c = append(*c, &Comment{
			ID:          v.ID,
			UniqueInt:   v.UniqueInt,
			UniqueFloat: v.UniqueFloat,
			NillableInt: v.NillableInt,
		})
	}
	return nil
}

func (c Comments) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
