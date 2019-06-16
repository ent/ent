// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/plugin/ent/boring"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// BoringCreate is the builder for creating a Boring entity.
type BoringCreate struct {
	config
}

// Save creates the Boring in the database.
func (bc *BoringCreate) Save(ctx context.Context) (*Boring, error) {
	switch bc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return bc.sqlSave(ctx)
	case dialect.Neptune:
		return bc.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BoringCreate) SaveX(ctx context.Context) *Boring {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BoringCreate) sqlSave(ctx context.Context) (*Boring, error) {
	var (
		res sql.Result
		b   = &Boring{config: bc.config}
	)
	tx, err := bc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(boring.Table).Default(bc.driver.Dialect())
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	b.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return b, nil
}

func (bc *BoringCreate) gremlinSave(ctx context.Context) (*Boring, error) {
	res := &gremlin.Response{}
	query, bindings := bc.gremlin().Query()
	if err := bc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	b := &Boring{config: bc.config}
	if err := b.FromResponse(res); err != nil {
		return nil, err
	}
	return b, nil
}

func (bc *BoringCreate) gremlin() *dsl.Traversal {
	v := g.AddV(boring.Label)
	return v.ValueMap(true)
}
