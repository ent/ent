// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"time"

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
func (_c *NodeCreate) SetValue(i int) *NodeCreate {
	_c.mutation.SetValue(i)
	return _c
}

// SetNillableValue sets the "value" field if the given value is not nil.
func (_c *NodeCreate) SetNillableValue(i *int) *NodeCreate {
	if i != nil {
		_c.SetValue(*i)
	}
	return _c
}

// SetUpdatedAt sets the "updated_at" field.
func (_c *NodeCreate) SetUpdatedAt(t time.Time) *NodeCreate {
	_c.mutation.SetUpdatedAt(t)
	return _c
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (_c *NodeCreate) SetNillableUpdatedAt(t *time.Time) *NodeCreate {
	if t != nil {
		_c.SetUpdatedAt(*t)
	}
	return _c
}

// SetPrevID sets the "prev" edge to the Node entity by ID.
func (_c *NodeCreate) SetPrevID(id string) *NodeCreate {
	_c.mutation.SetPrevID(id)
	return _c
}

// SetNillablePrevID sets the "prev" edge to the Node entity by ID if the given value is not nil.
func (_c *NodeCreate) SetNillablePrevID(id *string) *NodeCreate {
	if id != nil {
		_c = _c.SetPrevID(*id)
	}
	return _c
}

// SetPrev sets the "prev" edge to the Node entity.
func (_c *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return _c.SetPrevID(n.ID)
}

// SetNextID sets the "next" edge to the Node entity by ID.
func (_c *NodeCreate) SetNextID(id string) *NodeCreate {
	_c.mutation.SetNextID(id)
	return _c
}

// SetNillableNextID sets the "next" edge to the Node entity by ID if the given value is not nil.
func (_c *NodeCreate) SetNillableNextID(id *string) *NodeCreate {
	if id != nil {
		_c = _c.SetNextID(*id)
	}
	return _c
}

// SetNext sets the "next" edge to the Node entity.
func (_c *NodeCreate) SetNext(n *Node) *NodeCreate {
	return _c.SetNextID(n.ID)
}

// Mutation returns the NodeMutation object of the builder.
func (_c *NodeCreate) Mutation() *NodeMutation {
	return _c.mutation
}

// Save creates the Node in the database.
func (_c *NodeCreate) Save(ctx context.Context) (*Node, error) {
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *NodeCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *NodeCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *NodeCreate) check() error {
	return nil
}

func (_c *NodeCreate) gremlinSave(ctx context.Context) (*Node, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Node{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *NodeCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 2)
	v := g.AddV(node.Label)
	if value, ok := _c.mutation.Value(); ok {
		v.Property(dsl.Single, node.FieldValue, value)
	}
	if value, ok := _c.mutation.UpdatedAt(); ok {
		v.Property(dsl.Single, node.FieldUpdatedAt, value)
	}
	for _, id := range _c.mutation.PrevIDs() {
		v.AddE(node.NextLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(node.NextLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(node.Label, node.NextLabel, id)),
		})
	}
	for _, id := range _c.mutation.NextIDs() {
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
	err      error
	builders []*NodeCreate
}
