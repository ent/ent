// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/node"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// NodeDelete is the builder for deleting a Node entity.
type NodeDelete struct {
	config
	hooks    []Hook
	mutation *NodeMutation
}

// Where appends a list predicates to the NodeDelete builder.
func (nd *NodeDelete) Where(ps ...predicate.Node) *NodeDelete {
	nd.mutation.Where(ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NodeDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, NodeMutation](ctx, nd.gremlinExec, nd.mutation, nd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NodeDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NodeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nd.gremlin().Query()
	if err := nd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	nd.mutation.done = true
	return res.ReadInt()
}

func (nd *NodeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(node.Label)
	for _, p := range nd.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// NodeDeleteOne is the builder for deleting a single Node entity.
type NodeDeleteOne struct {
	nd *NodeDelete
}

// Where appends a list predicates to the NodeDelete builder.
func (ndo *NodeDeleteOne) Where(ps ...predicate.Node) *NodeDeleteOne {
	ndo.nd.mutation.Where(ps...)
	return ndo
}

// Exec executes the deletion query.
func (ndo *NodeDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{node.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NodeDeleteOne) ExecX(ctx context.Context) {
	if err := ndo.Exec(ctx); err != nil {
		panic(err)
	}
}
