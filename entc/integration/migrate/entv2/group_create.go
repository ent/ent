// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/group"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	return gc.sqlSave(ctx)
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
		builder = sql.Dialect(gc.driver.Dialect())
		gr      = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(group.Table).Default()

	id, err := insertLastID(ctx, tx, insert.Returning(group.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
