// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/dialect/gremlin/graph/dsl/p"
	"entgo.io/ent/entc/integration/gremlin/ent/node"
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

// SetNillableValue sets the "value" field if the given value is not nil.
func (nc *NodeCreate) SetNillableValue(i *int) *NodeCreate {
	if i != nil {
		nc.SetValue(*i)
	}
	return nc
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (nc *NodeCreate) SetPrevID(id string) *NodeCreate {
	nc.mutation.SetPrevID(id)
	return nc
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetPrevID(*id)
	}
	return nc
}

// SetPrev sets the "prev" edge to the Node entity.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (nc *NodeCreate) SetNextID(id string) *NodeCreate {
	nc.mutation.SetNextID(id)
	return nc
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *string) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the "next" edge to the Node entity.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
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
		node, err = nc.gremlinSave(ctx)
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
			node, err = nc.gremlinSave(ctx)
			mutation.id = &node.ID
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
	return nil
}

func (nc *NodeCreate) gremlinSave(ctx context.Context) (*Node, error) {
	res := &gremlin.Response{}
	query, bindings := nc.gremlin().Query()
	if err := nc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	n := &Node{config: nc.config}
	if err := n.FromResponse(res); err != nil {
		return nil, err
	}
	return n, nil
}

func (nc *NodeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(node.Label)
	if value, ok := nc.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	for _, id := range nc.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	for _, id := range nc.mutation.NextIDs() {
		v.AddE(node.NextLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}

// NodeCreateBulk is the builder for creating many Node entities in bulk.
type NodeCreateBulk struct {
	config
	builders []*NodeCreate
}
