// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"errors"
	"strconv"

	"fbc/ent/entc/integration/migrate/entv2/pet"

	"fbc/ent/dialect"
	"fbc/ent/dialect/sql"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
	"fbc/lib/go/gremlin/graph/dsl/g"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	switch pc.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return pc.sqlSave(ctx)
	case dialect.Neptune:
		return pc.gremlinSave(ctx)
	default:
		return nil, errors.New("entv2: unsupported dialect")
	}
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	var (
		res sql.Result
		pe  = &Pet{config: pc.config}
	)
	tx, err := pc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(pet.Table).Default(pc.driver.Dialect())
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	pe.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pc *PetCreate) gremlinSave(ctx context.Context) (*Pet, error) {
	res := &gremlin.Response{}
	query, bindings := pc.gremlin().Query()
	if err := pc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	pe := &Pet{config: pc.config}
	if err := pe.FromResponse(res); err != nil {
		return nil, err
	}
	return pe, nil
}

func (pc *PetCreate) gremlin() *dsl.Traversal {
	v := g.AddV(pet.Label)
	return v.ValueMap(true)
}
