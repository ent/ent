// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/sql"
)

// Comment is the model entity for the Comment schema.
type Comment struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Comment.
func (c *Comment) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	c.ID = vc.ID
	return nil
}

// FromRows scans the sql response data into Comment.
func (c *Comment) FromRows(rows *sql.Rows) error {
	var vc struct {
		ID int
	}
	// the order here should be the same as in the `comment.Columns`.
	if err := rows.Scan(
		&vc.ID,
	); err != nil {
		return err
	}
	c.ID = strconv.Itoa(vc.ID)
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
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Comment(")
	buf.WriteString(fmt.Sprintf("id=%v", c.ID))
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (c *Comment) id() int {
	id, _ := strconv.Atoi(c.ID)
	return id
}

// Comments is a parsable slice of Comment.
type Comments []*Comment

// FromResponse scans the gremlin response data into Comments.
func (c *Comments) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	for _, v := range vc {
		*c = append(*c, &Comment{
			ID: v.ID,
		})
	}
	return nil
}

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

func (c Comments) config(cfg config) {
	for i := range c {
		c[i].config = cfg
	}
}
