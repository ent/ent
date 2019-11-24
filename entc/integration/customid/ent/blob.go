// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/google/uuid"
)

// Blob is the model entity for the Blob schema.
type Blob struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// UUID holds the value of the "uuid" field.
	UUID uuid.UUID `json:"uuid,omitempty"`
}

// FromRows scans the sql response data into Blob.
func (b *Blob) FromRows(rows *sql.Rows) error {
	var scanb struct {
		ID   uuid.UUID
		UUID uuid.UUID
	}
	// the order here should be the same as in the `blob.Columns`.
	if err := rows.Scan(
		&scanb.ID,
		&scanb.UUID,
	); err != nil {
		return err
	}
	b.ID = scanb.ID
	b.UUID = scanb.UUID
	return nil
}

// Update returns a builder for updating this Blob.
// Note that, you need to call Blob.Unwrap() before calling this method, if this Blob
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Blob) Update() *BlobUpdateOne {
	return (&BlobClient{b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Blob) Unwrap() *Blob {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Blob is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Blob) String() string {
	var builder strings.Builder
	builder.WriteString("Blob(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", uuid=")
	builder.WriteString(fmt.Sprintf("%v", b.UUID))
	builder.WriteByte(')')
	return builder.String()
}

// Blobs is a parsable slice of Blob.
type Blobs []*Blob

// FromRows scans the sql response data into Blobs.
func (b *Blobs) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanb := &Blob{}
		if err := scanb.FromRows(rows); err != nil {
			return err
		}
		*b = append(*b, scanb)
	}
	return nil
}

func (b Blobs) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
