// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Card is the model entity for the Card schema.
type Card struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Number holds the value of the "number" field.
	Number string `json:"number,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// additional struct fields defined in the schema.
	RequestID string      // RequestID.
	Logger    *log.Logger // Logger.
}

// FromRows scans the sql response data into Card.
func (c *Card) FromRows(rows *sql.Rows) error {
	var vc struct {
		ID        int
		Number    sql.NullString
		CreatedAt sql.NullTime
		UpdatedAt sql.NullTime
	}
	// the order here should be the same as in the `card.Columns`.
	if err := rows.Scan(
		&vc.ID,
		&vc.Number,
		&vc.CreatedAt,
		&vc.UpdatedAt,
	); err != nil {
		return err
	}
	c.ID = strconv.Itoa(vc.ID)
	c.Number = vc.Number.String
	c.CreatedAt = vc.CreatedAt.Time
	c.UpdatedAt = vc.UpdatedAt.Time
	return nil
}

// FromResponse scans the gremlin response data into Card.
func (c *Card) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc struct {
		ID        string `json:"id,omitempty"`
		Number    string `json:"number,omitempty"`
		CreatedAt int64  `json:"created_at,omitempty"`
		UpdatedAt int64  `json:"updated_at,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	c.ID = vc.ID
	c.Number = vc.Number
	c.CreatedAt = time.Unix(0, vc.CreatedAt)
	c.UpdatedAt = time.Unix(0, vc.UpdatedAt)
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
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Card(")
	buf.WriteString(fmt.Sprintf("id=%v", c.ID))
	buf.WriteString(fmt.Sprintf(", number=%v", c.Number))
	buf.WriteString(fmt.Sprintf(", created_at=%v", c.CreatedAt))
	buf.WriteString(fmt.Sprintf(", updated_at=%v", c.UpdatedAt))
	buf.WriteString(")")
	return buf.String()
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
		vc := &Card{}
		if err := vc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, vc)
	}
	return nil
}

// FromResponse scans the gremlin response data into Cards.
func (c *Cards) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vc []struct {
		ID        string `json:"id,omitempty"`
		Number    string `json:"number,omitempty"`
		CreatedAt int64  `json:"created_at,omitempty"`
		UpdatedAt int64  `json:"updated_at,omitempty"`
	}
	if err := vmap.Decode(&vc); err != nil {
		return err
	}
	for _, v := range vc {
		*c = append(*c, &Card{
			ID:        v.ID,
			Number:    v.Number,
			CreatedAt: time.Unix(0, v.CreatedAt),
			UpdatedAt: time.Unix(0, v.UpdatedAt),
		})
	}
	return nil
}

func (c Cards) config(cfg config) {
	for i := range c {
		c[i].config = cfg
	}
}
