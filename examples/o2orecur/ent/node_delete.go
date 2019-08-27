// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/examples/o2orecur/ent/node"
	"github.com/facebookincubator/ent/examples/o2orecur/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// NodeDelete is the builder for deleting a Node entity.
type NodeDelete struct {
	config
	predicates []predicate.Node
}

// Where adds a new predicate for the builder.
func (nd *NodeDelete) Where(ps ...predicate.Node) *NodeDelete {
	nd.predicates = append(nd.predicates, ps...)
	return nd
}

// Exec executes the deletion query.
func (nd *NodeDelete) Exec(ctx context.Context) error {
	return nd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (nd *NodeDelete) ExecX(ctx context.Context) {
	if err := nd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nd *NodeDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(node.Table))
	for _, p := range nd.predicates {
		p(selector)
	}
	query, args := sql.Delete(node.Table).FromSelect(selector).Query()
	return nd.driver.Exec(ctx, query, args, &res)
}

// NodeDeleteOne is the builder for deleting a single Node entity.
type NodeDeleteOne struct {
	nd *NodeDelete
}

// Exec executes the deletion query.
func (ndo *NodeDeleteOne) Exec(ctx context.Context) error {
	return ndo.nd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ndo *NodeDeleteOne) ExecX(ctx context.Context) {
	ndo.nd.ExecX(ctx)
}
