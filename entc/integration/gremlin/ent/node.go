// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
)

// Node is the model entity for the Node schema.
type Node struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
}

// FromResponse scans the gremlin response data into Node.
func (n *Node) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scann struct {
		ID    string `json:"id,omitempty"`
		Value int    `json:"value,omitempty"`
	}
	if err := vmap.Decode(&scann); err != nil {
		return err
	}
	n.ID = scann.ID
	n.Value = scann.Value
	return nil
}

// QueryPrev queries the prev edge of the Node.
func (n *Node) QueryPrev() *NodeQuery {
	return (&NodeClient{n.config}).QueryPrev(n)
}

// QueryNext queries the next edge of the Node.
func (n *Node) QueryNext() *NodeQuery {
	return (&NodeClient{n.config}).QueryNext(n)
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

// id returns the int representation of the ID field.
func (n *Node) id() int {
	id, _ := strconv.Atoi(n.ID)
	return id
}

// Nodes is a parsable slice of Node.
type Nodes []*Node

// FromResponse scans the gremlin response data into Nodes.
func (n *Nodes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scann []struct {
		ID    string `json:"id,omitempty"`
		Value int    `json:"value,omitempty"`
	}
	if err := vmap.Decode(&scann); err != nil {
		return err
	}
	for _, v := range scann {
		*n = append(*n, &Node{
			ID:    v.ID,
			Value: v.Value,
		})
	}
	return nil
}

func (n Nodes) config(cfg config) {
	for _i := range n {
		n[_i].config = cfg
	}
}
