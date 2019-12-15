// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	return ic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	var (
		builder = sql.Dialect(ic.driver.Dialect())
		i       = &Item{config: ic.config}
	)
	tx, err := ic.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(item.Table).Default()

	id, err := insertLastID(ctx, tx, insert.Returning(item.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	i.ID = strconv.FormatInt(id, 10)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return i, nil
}
