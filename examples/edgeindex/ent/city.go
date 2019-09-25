// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// City is the model entity for the City schema.
type City struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into City.
func (c *City) FromRows(rows *sql.Rows) error {
	var vc struct {
		ID   int
		Name sql.NullString
	}
	// the order here should be the same as in the `city.Columns`.
	if err := rows.Scan(
		&vc.ID,
		&vc.Name,
	); err != nil {
		return err
	}
	c.ID = vc.ID
	c.Name = vc.Name.String
	return nil
}

// QueryStreets queries the streets edge of the City.
func (c *City) QueryStreets() *StreetQuery {
	return (&CityClient{c.config}).QueryStreets(c)
}

// Update returns a builder for updating this City.
// Note that, you need to call City.Unwrap() before calling this method, if this City
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *City) Update() *CityUpdateOne {
	return (&CityClient{c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *City) Unwrap() *City {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: City is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *City) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("City(")
	buf.WriteString(fmt.Sprintf("id=%v", c.ID))
	buf.WriteString(fmt.Sprintf(", name=%v", c.Name))
	buf.WriteString(")")
	return buf.String()
}

// Cities is a parsable slice of City.
type Cities []*City

// FromRows scans the sql response data into Cities.
func (c *Cities) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vc := &City{}
		if err := vc.FromRows(rows); err != nil {
			return err
		}
		*c = append(*c, vc)
	}
	return nil
}

func (c Cities) config(cfg config) {
	for i := range c {
		c[i].config = cfg
	}
}
