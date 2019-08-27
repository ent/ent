// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/examples/m2m2types/ent/group"
	"github.com/facebookincubator/ent/examples/m2m2types/ent/predicate"

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
	return gd.sqlExec(ctx)
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
