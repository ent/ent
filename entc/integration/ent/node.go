// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/node"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the NodeQuery when eager-loading is set.
	Edges     NodeEdges `json:"edges"`
	node_next *int
}

// NodeEdges holds the relations/edges for other nodes in the graph.
type NodeEdges struct {
	// Prev holds the value of the prev edge.
	Prev *Node `gqlgen:prev`
	// Next holds the value of the next edge.
	Next *Node `gqlgen:next`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) PrevOrErr() (*Node, error) {
	if e.loadedTypes[0] {
		if e.Prev == nil {
			// The edge prev was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.Prev, nil
	}
	return nil, &NotLoadedError{edge: "prev"}
}

// NextOrErr returns the Next value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) NextOrErr() (*Node, error) {
	if e.loadedTypes[1] {
		if e.Next == nil {
			// The edge next was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // id
		&sql.NullInt64{}, // value
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Node) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // node_next
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(values ...interface{}) error {
	if m, n := len(values), len(node.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	n.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field value", values[0])
	} else if value.Valid {
		n.Value = int(value.Int64)
	}
	values = values[1:]
	if len(values) == len(node.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field node_next", value)
		} else if value.Valid {
			n.node_next = new(int)
			*n.node_next = int(value.Int64)
		}
	}
	return nil
}

// QueryPrev queries the prev edge of the Node.
func (n *Node) QueryPrev() *NodeQuery {
	return (&NodeClient{config: n.config}).QueryPrev(n)
}

// QueryNext queries the next edge of the Node.
func (n *Node) QueryNext() *NodeQuery {
	return (&NodeClient{config: n.config}).QueryNext(n)
}

// Update returns a builder for updating this Node.
// Note that, you need to call Node.Unwrap() before calling this method, if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return (&NodeClient{config: n.config}).UpdateOne(n)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v", n.ID))
	builder.WriteString(", value=")
	builder.WriteString(fmt.Sprintf("%v", n.Value))
	builder.WriteByte(')')
	return builder.String()
}

// Nodes is a parsable slice of Node.
type Nodes []*Node

func (n Nodes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
