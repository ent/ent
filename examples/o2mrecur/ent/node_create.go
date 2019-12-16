// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"
	"github.com/facebookincubator/ent/schema/field"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	value    *int
	parent   map[int]struct{}
	children map[int]struct{}
}

// SetValue sets the value field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.value = &i
	return nc
}

// SetParentID sets the parent edge to Node by id.
func (nc *NodeCreate) SetParentID(id int) *NodeCreate {
	if nc.parent == nil {
		nc.parent = make(map[int]struct{})
	}
	nc.parent[id] = struct{}{}
	return nc
}

// SetNillableParentID sets the parent edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillableParentID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetParentID(*id)
	}
	return nc
}

// SetParent sets the parent edge to Node.
func (nc *NodeCreate) SetParent(n *Node) *NodeCreate {
	return nc.SetParentID(n.ID)
}

// AddChildIDs adds the children edge to Node by ids.
func (nc *NodeCreate) AddChildIDs(ids ...int) *NodeCreate {
	if nc.children == nil {
		nc.children = make(map[int]struct{})
	}
	for i := range ids {
		nc.children[ids[i]] = struct{}{}
	}
	return nc
}

// AddChildren adds the children edges to Node.
func (nc *NodeCreate) AddChildren(n ...*Node) *NodeCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddChildIDs(ids...)
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	if nc.value == nil {
		return nil, errors.New("ent: missing required field \"value\"")
	}
	if len(nc.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return nc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	var (
		n    = &Node{config: nc.config}
		spec = &sqlgraph.CreateSpec{
			Table: node.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		}
	)
	if value := nc.value; value != nil {
		spec.Fields = append(spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  *value,
			Column: node.FieldValue,
		})
		n.Value = *value
	}
	if nodes := nc.parent; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   node.ParentTable,
			Columns: []string{node.ParentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := nc.children; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   node.ChildrenTable,
			Columns: []string{node.ChildrenColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: node.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, nc.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}

	id := spec.ID.Value.(int64)
	n.ID = int(id)

	return n, nil
}
