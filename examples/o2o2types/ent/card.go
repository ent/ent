// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Card is the model entity for the Card schema.
type Card struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Expired holds the value of the "expired" field.
	Expired time.Time `json:"expired,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
}

// FromRows scans the sql response data into Card.
func (c *Card) FromRows(rows *sql.Rows) error {
	var scanc struct {
		ID      int
		Expired sql.NullTime
		Number  sql.NullString
	}
	// the order here should be the same as in the `card.Columns`.
	if err := rows.Scan(
		&scanc.ID,
		&scanc.Expired,
		&scanc.Number,
	); err != nil {
		return err
	}
	c.ID = scanc.ID
	c.Expired = scanc.Expired.Time
	c.Number = scanc.Number.String
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
	builder.WriteString(", expired=")
	builder.WriteString(c.Expired.Format(time.ANSIC))
	builder.WriteString(", number=")
	builder.WriteString(c.Number)
	builder.WriteByte(')')
	return builder.String()
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
