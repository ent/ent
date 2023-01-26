// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	fmt "fmt"
	strings "strings"

	"entgo.io/ent/dialect/sql"
	api "entgo.io/ent/entc/integration/ent/api"
)

// Api is the model entity for the Api schema.
type Api struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Api) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case api.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Api", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Api fields.
func (a *Api) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case api.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this Api.
// Note that you need to call Api.Unwrap() before calling this method if this Api
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Api) Update() *APIUpdateOne {
	return NewAPIClient(a.config).UpdateOne(a)
}

// Unwrap unwraps the Api entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Api) Unwrap() *Api {
	_tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Api is not a transactional entity")
	}
	a.config.driver = _tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Api) String() string {
	var builder strings.Builder
	builder.WriteString("Api(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Apis is a parsable slice of Api.
type Apis []*Api

func (a Apis) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
