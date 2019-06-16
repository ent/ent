// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/card"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// CardDelete is the builder for deleting a Card entity.
type CardDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (cd *CardDelete) Where(ps ...ent.Predicate) *CardDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query.
func (cd *CardDelete) Exec(ctx context.Context) error {
	switch cd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return cd.sqlExec(ctx)
	case dialect.Neptune:
		return cd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CardDelete) ExecX(ctx context.Context) {
	if err := cd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cd *CardDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(card.Table))
	for _, p := range cd.predicates {
		p.SQL(selector)
	}
	query, args := sql.Delete(card.Table).FromSelect(selector).Query()
	return cd.driver.Exec(ctx, query, args, &res)
}

func (cd *CardDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := cd.gremlin().Query()
	return cd.driver.Exec(ctx, query, bindings, res)
}

func (cd *CardDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(card.Label)
	for _, p := range cd.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
}

// CardDeleteOne is the builder for deleting a single Card entity.
type CardDeleteOne struct {
	cd *CardDelete
}

// Exec executes the deletion query.
func (cdo *CardDeleteOne) Exec(ctx context.Context) error {
	return cdo.cd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CardDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
