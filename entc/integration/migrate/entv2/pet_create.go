// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"

	"fbc/ent/entc/integration/migrate/entv2/pet"

	"fbc/ent/dialect/sql"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	return pc.sqlSave(ctx)
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
	pe.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
