// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/sql"
)

// Node is the model entity for the Node schema.
type Node struct {
	config
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Value holds the value of the "value" field.
	Value int `json:"value,omitempty"`
}

// FromRows scans the sql response data into Node.
func (n *Node) FromRows(rows *sql.Rows) error {
	var vn struct {
		ID    int
		Value sql.NullInt64
	}
	// the order here should be the same as in the `node.Columns`.
	if err := rows.Scan(
		&vn.ID,
		&vn.Value,
	); err != nil {
		return err
	}
	n.ID = strconv.Itoa(vn.ID)
	n.Value = int(vn.Value.Int64)
	return nil
}

// FromResponse scans the gremlin response data into Node.
func (n *Node) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vn struct {
		ID    string `json:"id,omitempty"`
		Value int    `json:"value,omitempty"`
	}
	if err := vmap.Decode(&vn); err != nil {
		return err
	}
	n.ID = vn.ID
	n.Value = vn.Value
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
	buf := bytes.NewBuffer(nil)
	buf.WriteString("Node(")
	buf.WriteString(fmt.Sprintf("id=%v", n.ID))
	buf.WriteString(fmt.Sprintf(", value=%v", n.Value))
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (n *Node) id() int {
	id, _ := strconv.Atoi(n.ID)
	return id
}

// Nodes is a parsable slice of Node.
type Nodes []*Node

// FromRows scans the sql response data into Nodes.
func (n *Nodes) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vn := &Node{}
		if err := vn.FromRows(rows); err != nil {
			return err
		}
		*n = append(*n, vn)
	}
	return nil
}

// FromResponse scans the gremlin response data into Nodes.
func (n *Nodes) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vn []struct {
		ID    string `json:"id,omitempty"`
		Value int    `json:"value,omitempty"`
	}
	if err := vmap.Decode(&vn); err != nil {
		return err
	}
	for _, v := range vn {
		*n = append(*n, &Node{
			ID:    v.ID,
			Value: v.Value,
		})
	}
	return nil
}

func (n Nodes) config(cfg config) {
	for i := range n {
		n[i].config = cfg
	}
}
