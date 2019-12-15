// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`

	// StaticField defined by templates.
	StaticField string `json:"boring,omitempty"`
}

// FromRows scans the sql response data into Card.
func (c *Card) FromRows(rows *sql.Rows) error {
	var scanc struct {
		ID         int
		CreateTime sql.NullTime
		UpdateTime sql.NullTime
		Number     sql.NullString
		Name       sql.NullString
	}
	// the order here should be the same as in the `card.Columns`.
	if err := rows.Scan(
		&scanc.ID,
		&scanc.CreateTime,
		&scanc.UpdateTime,
		&scanc.Number,
		&scanc.Name,
	); err != nil {
		return err
	}
	c.ID = strconv.Itoa(scanc.ID)
	c.CreateTime = scanc.CreateTime.Time
	c.UpdateTime = scanc.UpdateTime.Time
	c.Number = scanc.Number.String
	c.Name = scanc.Name.String
	return nil
}

// QueryOwner queries the owner edge of the Card.
func (c *Card) QueryOwner() *UserQuery {
	return (&CardClient{c.config}).QueryOwner(c)
}

// Update returns a builder for updating this Card.
// Note that, you need to call Card.Unwrap() before calling this method, if this Card
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Card) Update() *CardUpdateOne {
	return (&CardClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Card) Unwrap() *Card {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Card is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Card) String() string {
	var builder strings.Builder
	builder.WriteString("Card(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(c.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(c.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", number=")
	builder.WriteString(c.Number)
	builder.WriteString(", name=")
	builder.WriteString(c.Name)
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (c *Card) id() int {
	id, _ := strconv.Atoi(c.ID)
	return id
}

// Cards is a parsable slice of Card.
type Cards []*Card

// FromRows scans the sql response data into Cards.
func (c *Cards) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanc := &Card{}
		if err := scanc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, scanc)
	}
	return nil
}

func (c Cards) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
