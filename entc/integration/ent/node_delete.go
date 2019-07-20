// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/node"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// NodeDelete is the builder for deleting a Node entity.
type NodeDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (nd *NodeDelete) Where(ps ...ent.Predicate) *NodeDelete {
	nd.predicates = append(nd.predicates, ps...)
	return nd
}

// Exec executes the deletion query.
func (nd *NodeDelete) Exec(ctx context.Context) error {
	switch nd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return nd.sqlExec(ctx)
	case dialect.Neptune:
		return nd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
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
		p.SQL(selector)
	}
	query, args := sql.Delete(node.Table).FromSelect(selector).Query()
	return nd.driver.Exec(ctx, query, args, &res)
}

func (nd *NodeDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := nd.gremlin().Query()
	return nd.driver.Exec(ctx, query, bindings, res)
}

func (nd *NodeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(node.Label)
	for _, p := range nd.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
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
