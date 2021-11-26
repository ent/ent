// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"

	"entgo.io/ent/entc/integration/customid/ent/mixinid"
)

// MixinID is the model entity for the MixinID schema.
type MixinID struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// SomeField holds the value of the "some_field" field.
	SomeField string `json:"some_field,omitempty"`
	// MixinField holds the value of the "mixin_field" field.
	MixinField string `json:"mixin_field,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MixinID) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case mixinid.FieldSomeField, mixinid.FieldMixinField:
			values[i] = new(sql.NullString)
		case mixinid.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type MixinID", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MixinID fields.
func (mi *MixinID) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case mixinid.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				mi.ID = *value
			}
		case mixinid.FieldSomeField:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field some_field", values[i])
			} else if value.Valid {
				mi.SomeField = value.String
			}
		case mixinid.FieldMixinField:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mixin_field", values[i])
			} else if value.Valid {
				mi.MixinField = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this MixinID.
// Note that you need to call MixinID.Unwrap() before calling this method if this MixinID
// was returned from a transaction, and the transaction was committed or rolled back.
func (mi *MixinID) Update() *MixinIDUpdateOne {
	return (&MixinIDClient{config: mi.config}).UpdateOne(mi)
}

// Unwrap unwraps the MixinID entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (mi *MixinID) Unwrap() *MixinID {
	tx, ok := mi.config.driver.(*txDriver)
	if !ok {
		panic("ent: MixinID is not a transactional entity")
	}
	mi.config.driver = tx.drv
	return mi
}

// String implements the fmt.Stringer.
func (mi *MixinID) String() string {
	var builder strings.Builder
	builder.WriteString("MixinID(")
	builder.WriteString(fmt.Sprintf("id=%v", mi.ID))
	builder.WriteString(", some_field=")
	builder.WriteString(mi.SomeField)
	builder.WriteString(", mixin_field=")
	builder.WriteString(mi.MixinField)
	builder.WriteByte(')')
	return builder.String()
}

// MixinIDs is a parsable slice of MixinID.
type MixinIDs []*MixinID

func (mi MixinIDs) config(cfg config) {
	for _i := range mi {
		mi[_i].config = cfg
	}
}
