// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/migrate/entv1/user"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// UserDelete is the builder for deleting a User entity.
type UserDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (ud *UserDelete) Where(ps ...ent.Predicate) *UserDelete {
	ud.predicates = append(ud.predicates, ps...)
	return ud
}

// Exec executes the deletion query.
func (ud *UserDelete) Exec(ctx context.Context) error {
	switch ud.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ud.sqlExec(ctx)
	case dialect.Neptune:
		return ud.gremlinExec(ctx)
	default:
		return errors.New("entv1: unsupported dialect")
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
		p.SQL(selector)
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
		p.Gremlin(t)
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
