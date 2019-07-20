// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/group"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// GroupDelete is the builder for deleting a Group entity.
type GroupDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (gd *GroupDelete) Where(ps ...ent.Predicate) *GroupDelete {
	gd.predicates = append(gd.predicates, ps...)
	return gd
}

// Exec executes the deletion query.
func (gd *GroupDelete) Exec(ctx context.Context) error {
	switch gd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gd.sqlExec(ctx)
	case dialect.Neptune:
		return gd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gd *GroupDelete) ExecX(ctx context.Context) {
	if err := gd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gd *GroupDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(group.Table))
	for _, p := range gd.predicates {
		p.SQL(selector)
	}
	query, args := sql.Delete(group.Table).FromSelect(selector).Query()
	return gd.driver.Exec(ctx, query, args, &res)
}

func (gd *GroupDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := gd.gremlin().Query()
	return gd.driver.Exec(ctx, query, bindings, res)
}

func (gd *GroupDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(group.Label)
	for _, p := range gd.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
}

// GroupDeleteOne is the builder for deleting a single Group entity.
type GroupDeleteOne struct {
	gd *GroupDelete
}

// Exec executes the deletion query.
func (gdo *GroupDeleteOne) Exec(ctx context.Context) error {
	return gdo.gd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gdo *GroupDeleteOne) ExecX(ctx context.Context) {
	gdo.gd.ExecX(ctx)
}
