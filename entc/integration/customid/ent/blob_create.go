// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/customid/ent/blob"
	"github.com/google/uuid"
)

// BlobCreate is the builder for creating a Blob entity.
type BlobCreate struct {
	config
	id   *uuid.UUID
	uuid *uuid.UUID
}

// SetUUID sets the uuid field.
func (bc *BlobCreate) SetUUID(u uuid.UUID) *BlobCreate {
	bc.uuid = &u
	return bc
}

// SetID sets the id field.
func (bc *BlobCreate) SetID(u uuid.UUID) *BlobCreate {
	bc.id = &u
	return bc
}

// Save creates the Blob in the database.
func (bc *BlobCreate) Save(ctx context.Context) (*Blob, error) {
	if bc.uuid == nil {
		v := blob.DefaultUUID()
		bc.uuid = &v
	}
	return bc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BlobCreate) SaveX(ctx context.Context) *Blob {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BlobCreate) sqlSave(ctx context.Context) (*Blob, error) {
	var (
		builder = sql.Dialect(bc.driver.Dialect())
		b       = &Blob{config: bc.config}
	)
	tx, err := bc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(blob.Table).Default()
	if value := bc.uuid; value != nil {
		insert.Set(blob.FieldUUID, *value)
		b.UUID = *value
	}
	if value := bc.id; value != nil {
		insert.Set(blob.FieldID, *value)
		b.ID = *value
	}

	query, args := insert.Query()
	if err := tx.Exec(ctx, query, args, new(sql.Result)); err != nil {
		return nil, rollback(tx, err)
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return b, nil
}
