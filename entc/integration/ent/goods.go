// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/goods"
)

// Goods is the model entity for the Goods schema.
type Goods struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Goods) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case goods.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Goods", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Goods fields.
func (_go *Goods) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case goods.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			_go.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Goods.
// Note that you need to call Goods.Unwrap() before calling this method if this Goods
// was returned from a transaction, and the transaction was committed or rolled back.
func (_go *Goods) Update() *GoodsUpdateOne {
	return (&GoodsClient{config: _go.config}).UpdateOne(_go)
}

// Unwrap unwraps the Goods entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (_go *Goods) Unwrap() *Goods {
	_tx, ok := _go.config.driver.(*txDriver)
	if !ok {
		panic("ent: Goods is not a transactional entity")
	}
	_go.config.driver = _tx.drv
	return _go
}

// String implements the fmt.Stringer.
func (_go *Goods) String() string {
	var builder strings.Builder
	builder.WriteString("Goods(")
	builder.WriteString(fmt.Sprintf("id=%v", _go.ID))
	builder.WriteByte(')')
	return builder.String()
}

// GoodsSlice is a parsable slice of Goods.
type GoodsSlice []*Goods

func (_go GoodsSlice) config(cfg config) {
	for _i := range _go {
		_go[_i].config = cfg
	}
}
