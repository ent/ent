// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"fbc/ent/entc/integration/ent/fieldtype"

	"fbc/ent"
	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin"
	"fbc/ent/dialect/gremlin/graph/dsl"
	"fbc/ent/dialect/gremlin/graph/dsl/g"
	"fbc/ent/dialect/sql"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	predicates []ent.Predicate
}

// Where adds a new predicate for the builder.
func (ftd *FieldTypeDelete) Where(ps ...ent.Predicate) *FieldTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query.
func (ftd *FieldTypeDelete) Exec(ctx context.Context) error {
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
func (ftd *FieldTypeDelete) ExecX(ctx context.Context) {
	if err := ftd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (ftd *FieldTypeDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(fieldtype.Table))
	for _, p := range ftd.predicates {
		p.SQL(selector)
	}
	query, args := sql.Delete(fieldtype.Table).FromSelect(selector).Query()
	return ftd.driver.Exec(ctx, query, args, &res)
}

func (ftd *FieldTypeDelete) gremlinExec(ctx context.Context) error {
	res := &gremlin.Response{}
	query, bindings := ftd.gremlin().Query()
	return ftd.driver.Exec(ctx, query, bindings, res)
}

func (ftd *FieldTypeDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(fieldtype.Label)
	for _, p := range ftd.predicates {
		p.Gremlin(t)
	}
	return t.Drop()
}

// FieldTypeDeleteOne is the builder for deleting a single FieldType entity.
type FieldTypeDeleteOne struct {
	ftd *FieldTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FieldTypeDeleteOne) Exec(ctx context.Context) error {
	return ftdo.ftd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FieldTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
