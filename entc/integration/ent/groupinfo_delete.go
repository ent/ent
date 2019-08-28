// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	predicates []predicate.GroupInfo
}

// Where adds a new predicate for the builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.predicates = append(gid.predicates, ps...)
	return gid
}

// Exec executes the deletion query.
func (gid *GroupInfoDelete) Exec(ctx context.Context) error {
	switch gid.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gid.sqlExec(ctx)
	case dialect.Gremlin:
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
		p(selector)
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
		p(t)
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
