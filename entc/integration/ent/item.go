// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Item is the model entity for the Item schema.
type Item struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromRows scans the sql response data into Item.
func (i *Item) FromRows(rows *sql.Rows) error {
	var scani struct {
		ID int
	}
	// the order here should be the same as in the `item.Columns`.
	if err := rows.Scan(
		&scani.ID,
	); err != nil {
		return err
	}
	i.ID = strconv.Itoa(scani.ID)
	return nil
}

// Update returns a builder for updating this Item.
// Note that, you need to call Item.Unwrap() before calling this method, if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{i.config}).UpdateOne(i)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (i *Item) id() int {
	id, _ := strconv.Atoi(i.ID)
	return id
}

// Items is a parsable slice of Item.
type Items []*Item

// FromRows scans the sql response data into Items.
func (i *Items) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scani := &Item{}
		if err := scani.FromRows(rows); err != nil {
			return err
		}
		*i = append(*i, scani)
	}
	return nil
}

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
