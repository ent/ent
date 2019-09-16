// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/node"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// NodeDelete is the builder for deleting a Node entity.
type NodeDelete struct {
	config
	predicates []predicate.Node
}

// Where adds a new predicate to the delete builder.
func (nd *NodeDelete) Where(ps ...predicate.Node) *NodeDelete {
	nd.predicates = append(nd.predicates, ps...)
	return nd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (nd *NodeDelete) Exec(ctx context.Context) (int, error) {
	switch nd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return nd.sqlExec(ctx)
	case dialect.Gremlin:
		return nd.gremlinExec(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NodeDelete) ExecX(ctx context.Context) int {
	n, err := nd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (nd *NodeDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(node.Table))
	for _, p := range nd.predicates {
		p(selector)
	}
	query, args := sql.Delete(node.Table).FromSelect(selector).Query()
	if err := nd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

func (nd *NodeDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := nd.gremlin().Query()
	if err := nd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (nd *NodeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(node.Label)
	for _, p := range nd.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// NodeDeleteOne is the builder for deleting a single Node entity.
type NodeDeleteOne struct {
	nd *NodeDelete
}

// Exec executes the deletion query.
func (ndo *NodeDeleteOne) Exec(ctx context.Context) error {
	n, err := ndo.nd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{node.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NodeDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
