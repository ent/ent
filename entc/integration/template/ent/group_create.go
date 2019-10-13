// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/template/ent/group"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	max_users *int
}

// SetMaxUsers sets the max_users field.
func (gc *GroupCreate) SetMaxUsers(i int) *GroupCreate {
	gc.max_users = &i
	return gc
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if gc.max_users == nil {
		return nil, errors.New("ent: missing required field \"max_users\"")
	}
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
		res sql.Result
		gr  = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(gc.driver.Dialect()).
		Insert(group.Table).
		Default()
	if value := gc.max_users; value != nil {
		builder.Set(group.FieldMaxUsers, *value)
		gr.MaxUsers = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = int(id)
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
