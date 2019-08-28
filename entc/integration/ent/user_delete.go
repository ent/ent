// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// UserDelete is the builder for deleting a User entity.
type UserDelete struct {
	config
	predicates []predicate.User
}

// Where adds a new predicate for the builder.
func (ud *UserDelete) Where(ps ...predicate.User) *UserDelete {
	ud.predicates = append(ud.predicates, ps...)
	return ud
}

// Exec executes the deletion query.
func (ud *UserDelete) Exec(ctx context.Context) error {
	switch ud.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ud.sqlExec(ctx)
	case dialect.Gremlin:
		return ud.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ud *UserDelete) ExecX(ctx context.Context) {
	if err := ud.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ud *UserDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(user.Table))
	for _, p := range ud.predicates {
		p(selector)
	}
	query, args := sql.Delete(user.Table).FromSelect(selector).Query()
	return ud.driver.Exec(ctx, query, args, &res)
}

func (ud *UserDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := ud.gremlin().Query()
	return ud.driver.Exec(ctx, query, bindings, res)
}

func (ud *UserDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(user.Label)
	for _, p := range ud.predicates {
		p(t)
	}
	return t.Drop()
}

// UserDeleteOne is the builder for deleting a single User entity.
type UserDeleteOne struct {
	ud *UserDelete
}

// Exec executes the deletion query.
func (udo *UserDeleteOne) Exec(ctx context.Context) error {
	return udo.ud.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (udo *UserDeleteOne) ExecX(ctx context.Context) {
	udo.ud.ExecX(ctx)
}
