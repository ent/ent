// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/pet"
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
		builder = sql.Dialect(pc.driver.Dialect())
		pe      = &Pet{config: pc.config}
	)
	tx, err := pc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(pet.Table).Default()
	id, err := insertLastID(ctx, tx, insert.Returning(pet.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	pe.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
