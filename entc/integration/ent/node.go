// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/ent/node"
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
	Prev *Node `json:"prev,omitempty" gqlgen:"prev"`
	// Next holds the value of the next edge.
	Next *Node `json:"next,omitempty" gqlgen:"next"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// PrevOrErr returns the Prev value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e NodeEdges) PrevOrErr() (*Node, error) {
	if e.loadedTypes[0] {
		if e.Prev == nil {
			// Edge was loaded but was not found.
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
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: node.Label}
		}
		return e.Next, nil
	}
	return nil, &NotLoadedError{edge: "next"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Node) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case node.FieldID, node.FieldValue:
			values[i] = new(sql.NullInt64)
		case node.ForeignKeys[0]: // node_next
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Node", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Node fields.
func (n *Node) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case node.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int(value.Int64)
		case node.FieldValue:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field value", values[i])
			} else if value.Valid {
				n.Value = int(value.Int64)
			}
		case node.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field node_next", value)
			} else if value.Valid {
				n.node_next = new(int)
				*n.node_next = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryPrev queries the "prev" edge of the Node entity.
func (n *Node) QueryPrev() *NodeQuery {
	return NewNodeClient(n.config).QueryPrev(n)
}

// QueryNext queries the "next" edge of the Node entity.
func (n *Node) QueryNext() *NodeQuery {
	return NewNodeClient(n.config).QueryNext(n)
}

// Update returns a builder for updating this Node.
// Note that you need to call Node.Unwrap() before calling this method if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return NewNodeClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the Node entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *Node) Unwrap() *Node {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: Node is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *Node) String() string {
	var builder strings.Builder
	builder.WriteString("Node(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("value=")
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

func (n Nodes) IDs() []int {
	ids := make([]int, len(n))
	for _i := range n {
		ids[_i] = n[_i].ID
	}
	return ids
}
