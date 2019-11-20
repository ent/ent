// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
}

// FromRows scans the sql response data into Node.
func (n *Node) FromRows(rows *sql.Rows) error {
	var scann struct {
		ID    int
		Value sql.NullInt64
	}
	// the order here should be the same as in the `node.Columns`.
	if err := rows.Scan(
		&scann.ID,
		&scann.Value,
	); err != nil {
		return err
	}
	n.ID = scann.ID
	n.Value = int(scann.Value.Int64)
	return nil
}

// QueryParent queries the parent edge of the Node.
func (n *Node) QueryParent() *NodeQuery {
	return (&NodeClient{n.config}).QueryParent(n)
}

// QueryChildren queries the children edge of the Node.
func (n *Node) QueryChildren() *NodeQuery {
	return (&NodeClient{n.config}).QueryChildren(n)
}

// Update returns a builder for updating this Node.
// Note that, you need to call Node.Unwrap() before calling this method, if this Node
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *Node) Update() *NodeUpdateOne {
	return (&NodeClient{n.config}).UpdateOne(n)
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

// FromRows scans the sql response data into Nodes.
func (n *Nodes) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scann := &Node{}
		if err := scann.FromRows(rows); err != nil {
			return err
		}
		*n = append(*n, scann)
	}
	return nil
}

func (n Nodes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
