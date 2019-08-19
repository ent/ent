// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/comment"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// CommentDelete is the builder for deleting a Comment entity.
type CommentDelete struct {
	config
	predicates []predicate.Comment
}

// Where adds a new predicate for the builder.
func (cd *CommentDelete) Where(ps ...predicate.Comment) *CommentDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query.
func (cd *CommentDelete) Exec(ctx context.Context) error {
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
func (cd *CommentDelete) ExecX(ctx context.Context) {
	if err := cd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cd *CommentDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(comment.Table))
	for _, p := range cd.predicates {
		p(selector)
	}
	query, args := sql.Delete(comment.Table).FromSelect(selector).Query()
	return cd.driver.Exec(ctx, query, args, &res)
}

func (cd *CommentDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := cd.gremlin().Query()
	return cd.driver.Exec(ctx, query, bindings, res)
}

func (cd *CommentDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(comment.Label)
	for _, p := range cd.predicates {
		p(t)
	}
	return t.Drop()
}

// CommentDeleteOne is the builder for deleting a single Comment entity.
type CommentDeleteOne struct {
	cd *CommentDelete
}

// Exec executes the deletion query.
func (cdo *CommentDeleteOne) Exec(ctx context.Context) error {
	return cdo.cd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CommentDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
