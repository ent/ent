// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/schema/field"
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
		i    = &Item{config: ic.config}
		spec = &sqlgraph.CreateSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: item.FieldID,
			},
		}
	)
	if err := sqlgraph.CreateNode(ctx, ic.driver, spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := spec.ID.Value.(int64)
	i.ID = strconv.FormatInt(id, 10)
	return i, nil
}
