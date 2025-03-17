// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
)

// PC is the model entity for the PC schema.
type PC struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into PC.
func (_pc *PC) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_pc struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_pc); err != nil {
		return err
	}
	_pc.ID = scan_pc.ID
	return nil
}

// Update returns a builder for updating this PC.
// Note that you need to call PC.Unwrap() before calling this method if this PC
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *PC) Update() *PCUpdateOne {
	return NewPCClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the PC entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *PC) Unwrap() *PC {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: PC is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *PC) String() string {
	var builder strings.Builder
	builder.WriteString("PC(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// PCs is a parsable slice of PC.
type PCs []*PC

// FromResponse scans the gremlin response data into PCs.
func (_pc *PCs) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scan_pc []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scan_pc); err != nil {
		return err
	}
	for _, v := range scan_pc {
		node := &PC{ID: v.ID}
		*_pc = append(*_pc, node)
	}
	return nil
}
