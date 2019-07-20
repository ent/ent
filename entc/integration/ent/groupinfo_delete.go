// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/groupinfo"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (gid *GroupInfoDelete) Where(ps ...ent.Predicate) *GroupInfoDelete {
	gid.predicates = append(gid.predicates, ps...)
	return gid
}

// Exec executes the deletion query.
func (gid *GroupInfoDelete) Exec(ctx context.Context) error {
	switch gid.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gid.sqlExec(ctx)
	case dialect.Neptune:
		return gid.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) {
	if err := gid.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gid *GroupInfoDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(groupinfo.Table))
	for _, p := range gid.predicates {
		p.SQL(selector)
	}
	query, args := sql.Delete(groupinfo.Table).FromSelect(selector).Query()
	return gid.driver.Exec(ctx, query, args, &res)
}

func (gid *GroupInfoDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := gid.gremlin().Query()
	return gid.driver.Exec(ctx, query, bindings, res)
}

func (gid *GroupInfoDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(groupinfo.Label)
	for _, p := range gid.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	return gido.gid.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	gido.gid.ExecX(ctx)
}
