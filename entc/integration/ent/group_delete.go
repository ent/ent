// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GroupDelete is the builder for deleting a Group entity.
type GroupDelete struct {
	config
	predicates []predicate.Group
}

// Where adds a new predicate for the builder.
func (gd *GroupDelete) Where(ps ...predicate.Group) *GroupDelete {
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
		p(selector)
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
		p(t)
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
