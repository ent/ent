// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// FileTypeDelete is the builder for deleting a FileType entity.
type FileTypeDelete struct {
	config
	predicates []predicate.FileType
}

// Where adds a new predicate for the builder.
func (ftd *FileTypeDelete) Where(ps ...predicate.FileType) *FileTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query.
func (ftd *FileTypeDelete) Exec(ctx context.Context) error {
	switch ftd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftd.sqlExec(ctx)
	case dialect.Neptune:
		return ftd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) {
	if err := ftd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftd *FileTypeDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(filetype.Table))
	for _, p := range ftd.predicates {
		p(selector)
	}
	query, args := sql.Delete(filetype.Table).FromSelect(selector).Query()
	return ftd.driver.Exec(ctx, query, args, &res)
}

func (ftd *FileTypeDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := ftd.gremlin().Query()
	return ftd.driver.Exec(ctx, query, bindings, res)
}

func (ftd *FileTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(filetype.Label)
	for _, p := range ftd.predicates {
		p(t)
	}
	return t.Drop()
}

// FileTypeDeleteOne is the builder for deleting a single FileType entity.
type FileTypeDeleteOne struct {
	ftd *FileTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FileTypeDeleteOne) Exec(ctx context.Context) error {
	return ftdo.ftd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
