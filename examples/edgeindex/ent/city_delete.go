// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CityDelete is the builder for deleting a City entity.
type CityDelete struct {
	config
	predicates []predicate.City
}

// Where adds a new predicate for the builder.
func (cd *CityDelete) Where(ps ...predicate.City) *CityDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query.
func (cd *CityDelete) Exec(ctx context.Context) error {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CityDelete) ExecX(ctx context.Context) {
	if err := cd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cd *CityDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(city.Table))
	for _, p := range cd.predicates {
		p(selector)
	}
	query, args := sql.Delete(city.Table).FromSelect(selector).Query()
	return cd.driver.Exec(ctx, query, args, &res)
}

// CityDeleteOne is the builder for deleting a single City entity.
type CityDeleteOne struct {
	cd *CityDelete
}

// Exec executes the deletion query.
func (cdo *CityDeleteOne) Exec(ctx context.Context) error {
	return cdo.cd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CityDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
