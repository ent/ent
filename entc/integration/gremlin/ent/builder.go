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

// Builder is the model entity for the Builder schema.
type Builder struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
}

// FromResponse scans the gremlin response data into Builder.
func (b *Builder) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanb struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scanb); err != nil {
		return err
	}
	b.ID = scanb.ID
	return nil
}

// Update returns a builder for updating this Builder.
// Note that you need to call Builder.Unwrap() before calling this method if this Builder
// was returned from a transaction, and the transaction was committed or rolled back.
func (m *Builder) Update() *BuilderUpdateOne {
	return NewBuilderClient(m.config).UpdateOne(m)
}

// Unwrap unwraps the Builder entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (m *Builder) Unwrap() *Builder {
	_tx, ok := m.config.driver.(*txDriver)
	if !ok {
		panic("ent: Builder is not a transactional entity")
	}
	m.config.driver = _tx.drv
	return m
}

// String implements the fmt.Stringer.
func (m *Builder) String() string {
	var builder strings.Builder
	builder.WriteString("Builder(")
	builder.WriteString(fmt.Sprintf("id=%v", m.ID))
	builder.WriteByte(')')
	return builder.String()
}

// Builders is a parsable slice of Builder.
type Builders []*Builder

// FromResponse scans the gremlin response data into Builders.
func (b *Builders) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanb []struct {
		ID string `json:"id,omitempty"`
	}
	if err := vmap.Decode(&scanb); err != nil {
		return err
	}
	for _, v := range scanb {
		node := &Builder{ID: v.ID}
		*b = append(*b, node)
	}
	return nil
}
