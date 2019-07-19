// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/migrate/entv2/pet"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// PetDelete is the builder for deleting a Pet entity.
type PetDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (pd *PetDelete) Where(ps ...ent.Predicate) *PetDelete {
	pd.predicates = append(pd.predicates, ps...)
	return pd
}

// Exec executes the deletion query.
func (pd *PetDelete) Exec(ctx context.Context) error {
	switch pd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pd.sqlExec(ctx)
	case dialect.Neptune:
		return pd.gremlinExec(ctx)
	default:
		return errors.New("entv2: unsupported dialect")
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
		p.SQL(selector)
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
		p.Gremlin(t)
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
