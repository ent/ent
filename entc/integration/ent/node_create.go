// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/node"
	"github.com/facebookincubator/ent/schema/field"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	value *int
	prev  map[string]struct{}
	next  map[string]struct{}
}

// SetValue sets the value field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.value = &i
	return nc
}

// SetNillableValue sets the value field if the given value is not nil.
func (nc *NodeCreate) SetNillableValue(i *int) *NodeCreate {
	if i != nil {
		nc.SetValue(*i)
	}
	return nc
}

// SetPrevID sets the prev edge to Node by id.
func (nc *NodeCreate) SetPrevID(id string) *NodeCreate {
	if nc.prev == nil {
		nc.prev = make(map[string]struct{})
	}
	nc.prev[id] = struct{}{}
	return nc
}

// SetNillablePrevID sets the prev edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetPrevID(*id)
	}
	return nc
}

// SetPrev sets the prev edge to Node.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the next edge to Node by id.
func (nc *NodeCreate) SetNextID(id string) *NodeCreate {
	if nc.next == nil {
		nc.next = make(map[string]struct{})
	}
	nc.next[id] = struct{}{}
	return nc
}

// SetNillableNextID sets the next edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the next edge to Node.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	if len(nc.prev) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(nc.next) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"next\"")
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
				Type:   field.TypeString,
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
	if nodes := nc.prev; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   node.PrevTable,
			Columns: []string{node.PrevColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: node.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		spec.Edges = append(spec.Edges, edge)
	}
	if nodes := nc.next; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   node.NextTable,
			Columns: []string{node.NextColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeString,
					Column: node.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			k, err := strconv.Atoi(k)
			if err != nil {
				return nil, err
			}
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
	n.ID = strconv.FormatInt(id, 10)
	return n, nil
}
