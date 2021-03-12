// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/examples/o2mrecur/ent/node"
	"entgo.io/ent/schema/field"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	mutation *NodeMutation
	hooks    []Hook
}

// SetValue sets the "value" field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.mutation.SetValue(i)
	return nc
}

// SetParentID sets the "parent" edge to the Node entity by ID.
func (nc *NodeCreate) SetParentID(id int) *NodeCreate {
	nc.mutation.SetParentID(id)
	return nc
}

// SetNillableParentID sets the "parent" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillableParentID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetParentID(*id)
	}
	return nc
}

// SetParent sets the "parent" edge to the Node entity.
func (nc *NodeCreate) SetParent(n *Node) *NodeCreate {
	return nc.SetParentID(n.ID)
}

// AddChildIDs adds the "children" edge to the Node entity by IDs.
func (nc *NodeCreate) AddChildIDs(ids ...int) *NodeCreate {
	nc.mutation.AddChildIDs(ids...)
	return nc
}

// AddChildren adds the "children" edges to the Node entity.
func (nc *NodeCreate) AddChildren(n ...*Node) *NodeCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddChildIDs(ids...)
}

// Mutation returns the NodeMutation object of the builder.
func (nc *NodeCreate) Mutation() *NodeMutation {
	return nc.mutation
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	var (
		err  error
		node *Node
	)
	if len(nc.hooks) == 0 {
		if err = nc.check(); err != nil {
			return nil, err
		}
		node, err = nc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*NodeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = nc.check(); err != nil {
				return nil, err
			}
			nc.mutation = mutation
			node, err = nc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(nc.hooks) - 1; i >= 0; i-- {
			mut = nc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, nc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// check runs all checks and user-defined validators on the builder.
func (nc *NodeCreate) check() error {
	if _, ok := nc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New("ent: missing required field \"value\"")}
	}
	return nil
}

// OnConflict specifies how to handle inserts that conflict with a unique constraint on Node entities in the database.
func (nc *NodeCreate) OnConflict(fields ...string) *NodeCreate {
	return nc
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	_node, _spec := nc.createSpec()
	if err := sqlgraph.CreateNode(ctx, nc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (nc *NodeCreate) createSpec() (*Node, *sqlgraph.CreateSpec) {
	var (
		_node = &Node{config: nc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: node.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: node.FieldID,
			},
		}
	)
	if value, ok := nc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: node.FieldValue,
		})
		_node.Value = value
	}
	if nodes := nc.mutation.ParentIDs(); len(nodes) > 0 {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.node_children = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := nc.mutation.ChildrenIDs(); len(nodes) > 0 {
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
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
}

// OnConflict specifies how to handle inserts that conflict with a unique constraint on Node entities in the database.
func (ncb *NodeCreateBulk) OnConflict() *NodeCreateBulk {
	// for i := range ncb.builders {

	// }

	return ncb
}

// Save creates the Node entities in the database.
func (ncb *NodeCreateBulk) Save(ctx context.Context) ([]*Node, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ncb.builders))
	nodes := make([]*Node, len(ncb.builders))
	mutators := make([]Mutator, len(ncb.builders))
	for i := range ncb.builders {
		func(i int, root context.Context) {
			builder := ncb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*NodeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ncb.builders[i+1].mutation)
				} else {
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ncb.driver, &sqlgraph.BatchCreateSpec{Nodes: specs}); err != nil {
						if cerr, ok := isSQLConstraintError(err); ok {
							err = cerr
						}
					}
				}
				mutation.done = true
				if err != nil {
					return nil, err
				}
				id := specs[i].ID.Value.(int64)
				nodes[i].ID = int(id)
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ncb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ncb *NodeCreateBulk) SaveX(ctx context.Context) []*Node {
	v, err := ncb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}
