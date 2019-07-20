// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/plugin/ent/boring"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// BoringDelete is the builder for deleting a Boring entity.
type BoringDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (bd *BoringDelete) Where(ps ...ent.Predicate) *BoringDelete {
	bd.predicates = append(bd.predicates, ps...)
	return bd
}

// Exec executes the deletion query.
func (bd *BoringDelete) Exec(ctx context.Context) error {
	switch bd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bd.sqlExec(ctx)
	case dialect.Neptune:
		return bd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (bd *BoringDelete) ExecX(ctx context.Context) {
	if err := bd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bd *BoringDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(boring.Table))
	for _, p := range bd.predicates {
		p.SQL(selector)
	}
	query, args := sql.Delete(boring.Table).FromSelect(selector).Query()
	return bd.driver.Exec(ctx, query, args, &res)
}

func (bd *BoringDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := bd.gremlin().Query()
	return bd.driver.Exec(ctx, query, bindings, res)
}

func (bd *BoringDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(boring.Label)
	for _, p := range bd.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
}

// BoringDeleteOne is the builder for deleting a single Boring entity.
type BoringDeleteOne struct {
	bd *BoringDelete
}

// Exec executes the deletion query.
func (bdo *BoringDeleteOne) Exec(ctx context.Context) error {
	return bdo.bd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (bdo *BoringDeleteOne) ExecX(ctx context.Context) {
	bdo.bd.ExecX(ctx)
}
