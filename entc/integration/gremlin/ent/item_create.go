// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/item"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	return ic.gremlinSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ic *ItemCreate) gremlinSave(ctx context.Context) (*Item, error) {
	res := &gremlin.Response{}
	query, bindings := ic.gremlin().Query()
	if err := ic.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	i := &Item{config: ic.config}
	if err := i.FromResponse(res); err != nil {
		return nil, err
	}
	return i, nil
}

func (ic *ItemCreate) gremlin() *dsl.Traversal {
	v := g.AddV(item.Label)
	return v.ValueMap(true)
}
