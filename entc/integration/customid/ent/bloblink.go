// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/customid/ent/blob"
	"entgo.io/ent/entc/integration/customid/ent/bloblink"
	"github.com/google/uuid"
)

// BlobLink is the model entity for the BlobLink schema.
type BlobLink struct {
	config `json:"-"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// BlobID holds the value of the "blob_id" field.
	BlobID uuid.UUID `json:"blob_id,omitempty"`
	// LinkID holds the value of the "link_id" field.
	LinkID uuid.UUID `json:"link_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BlobLinkQuery when eager-loading is set.
	Edges BlobLinkEdges `json:"edges"`
}

// BlobLinkEdges holds the relations/edges for other nodes in the graph.
type BlobLinkEdges struct {
	// Blob holds the value of the blob edge.
	Blob *Blob `json:"blob,omitempty"`
	// Link holds the value of the link edge.
	Link *Blob `json:"link,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// BlobOrErr returns the Blob value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlobLinkEdges) BlobOrErr() (*Blob, error) {
	if e.loadedTypes[0] {
		if e.Blob == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blob.Label}
		}
		return e.Blob, nil
	}
	return nil, &NotLoadedError{edge: "blob"}
}

// LinkOrErr returns the Link value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BlobLinkEdges) LinkOrErr() (*Blob, error) {
	if e.loadedTypes[1] {
		if e.Link == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: blob.Label}
		}
		return e.Link, nil
	}
	return nil, &NotLoadedError{edge: "link"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BlobLink) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case bloblink.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case bloblink.FieldBlobID, bloblink.FieldLinkID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type BlobLink", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BlobLink fields.
func (bl *BlobLink) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bloblink.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				bl.CreatedAt = value.Time
			}
		case bloblink.FieldBlobID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field blob_id", values[i])
			} else if value != nil {
				bl.BlobID = *value
			}
		case bloblink.FieldLinkID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field link_id", values[i])
			} else if value != nil {
				bl.LinkID = *value
			}
		}
	}
	return nil
}

// QueryBlob queries the "blob" edge of the BlobLink entity.
func (bl *BlobLink) QueryBlob() *BlobQuery {
	return (&BlobLinkClient{config: bl.config}).QueryBlob(bl)
}

// QueryLink queries the "link" edge of the BlobLink entity.
func (bl *BlobLink) QueryLink() *BlobQuery {
	return (&BlobLinkClient{config: bl.config}).QueryLink(bl)
}

// Update returns a builder for updating this BlobLink.
// Note that you need to call BlobLink.Unwrap() before calling this method if this BlobLink
// was returned from a transaction, and the transaction was committed or rolled back.
func (bl *BlobLink) Update() *BlobLinkUpdateOne {
	return (&BlobLinkClient{config: bl.config}).UpdateOne(bl)
}

// Unwrap unwraps the BlobLink entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bl *BlobLink) Unwrap() *BlobLink {
	_tx, ok := bl.config.driver.(*txDriver)
	if !ok {
		panic("ent: BlobLink is not a transactional entity")
	}
	bl.config.driver = _tx.drv
	return bl
}

// String implements the fmt.Stringer.
func (bl *BlobLink) String() string {
	var builder strings.Builder
	builder.WriteString("BlobLink(")
	builder.WriteString("created_at=")
	builder.WriteString(bl.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("blob_id=")
	builder.WriteString(fmt.Sprintf("%v", bl.BlobID))
	builder.WriteString(", ")
	builder.WriteString("link_id=")
	builder.WriteString(fmt.Sprintf("%v", bl.LinkID))
	builder.WriteByte(')')
	return builder.String()
}

// BlobLinks is a parsable slice of BlobLink.
type BlobLinks []*BlobLink

func (bl BlobLinks) config(cfg config) {
	for _i := range bl {
		bl[_i].config = cfg
	}
}
