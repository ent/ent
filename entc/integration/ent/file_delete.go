// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/file"
	"fbc/ent/entc/integration/ent/predicate"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FileDelete is the builder for deleting a File entity.
type FileDelete struct {
	config
	predicates []predicate.File
}

// Where adds a new predicate for the builder.
func (fd *FileDelete) Where(ps ...predicate.File) *FileDelete {
	fd.predicates = append(fd.predicates, ps...)
	return fd
}

// Exec executes the deletion query.
func (fd *FileDelete) Exec(ctx context.Context) error {
	switch fd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return fd.sqlExec(ctx)
	case dialect.Neptune:
		return fd.gremlinExec(ctx)
	default:
		return errors.New("ent: unsupported dialect")
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FileDelete) ExecX(ctx context.Context) {
	if err := fd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (fd *FileDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(file.Table))
	for _, p := range fd.predicates {
		p(selector)
	}
	query, args := sql.Delete(file.Table).FromSelect(selector).Query()
	return fd.driver.Exec(ctx, query, args, &res)
}

func (fd *FileDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := fd.gremlin().Query()
	return fd.driver.Exec(ctx, query, bindings, res)
}

func (fd *FileDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(file.Label)
	for _, p := range fd.predicates {
		p(t)
	}
	return t.Drop()
}

// FileDeleteOne is the builder for deleting a single File entity.
type FileDeleteOne struct {
	fd *FileDelete
}

// Exec executes the deletion query.
func (fdo *FileDeleteOne) Exec(ctx context.Context) error {
	return fdo.fd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FileDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
