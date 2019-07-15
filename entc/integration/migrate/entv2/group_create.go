// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/migrate/entv2/group"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	switch gc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return gc.sqlSave(ctx)
	case dialect.Neptune:
		return gc.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	var (
		res sql.Result
		gr  = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(group.Table).Default(gc.driver.Dialect())
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gc *GroupCreate) gremlinSave(ctx context.Context) (*Group, error) {
	res := &gremlin.Response{}
	query, bindings := gc.gremlin().Query()
	if err := gc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gr := &Group{config: gc.config}
	if err := gr.FromResponse(res); err != nil {
		return nil, err
	}
	return gr, nil
}

func (gc *GroupCreate) gremlin() *dsl.Traversal {
	v := g.AddV(group.Label)
	return v.ValueMap(true)
}
