// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	predicates []predicate.FieldType
}

// Where adds a new predicate for the builder.
func (ftd *FieldTypeDelete) Where(ps ...predicate.FieldType) *FieldTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query.
func (ftd *FieldTypeDelete) Exec(ctx context.Context) error {
	switch ftd.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return ftd.sqlExec(ctx)
	case dialect.Gremlin:
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
		p(selector)
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
		p(t)
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
