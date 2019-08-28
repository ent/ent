// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// PetDelete is the builder for deleting a Pet entity.
type PetDelete struct {
	config
	predicates []predicate.Pet
}

// Where adds a new predicate for the builder.
func (pd *PetDelete) Where(ps ...predicate.Pet) *PetDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query.
func (pd *PetDelete) Exec(ctx context.Context) error {
	switch pd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pd.sqlExec(ctx)
	case dialect.Gremlin:
		return pd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pd *PetDelete) ExecX(ctx context.Context) {
	if err := pd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pd *PetDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(pet.Table))
	for _, p := range pd.predicates {
		p(selector)
	}
	query, args := sql.Delete(pet.Table).FromSelect(selector).Query()
	return pd.driver.Exec(ctx, query, args, &res)
}

func (pd *PetDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := pd.gremlin().Query()
	return pd.driver.Exec(ctx, query, bindings, res)
}

func (pd *PetDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(pet.Label)
	for _, p := range pd.predicates {
		p(t)
	}
	return t.Drop()
}

// PetDeleteOne is the builder for deleting a single Pet entity.
type PetDeleteOne struct {
	pd *PetDelete
}

// Exec executes the deletion query.
func (pdo *PetDeleteOne) Exec(ctx context.Context) error {
	return pdo.pd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (pdo *PetDeleteOne) ExecX(ctx context.Context) {
	pdo.pd.ExecX(ctx)
}
