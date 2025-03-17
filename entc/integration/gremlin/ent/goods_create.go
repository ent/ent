// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/goods"
)

// GoodsCreate is the builder for creating a Goods entity.
type GoodsCreate struct {
	config
	mutation *GoodsMutation
	hooks    []Hook
}

// Mutation returns the GoodsMutation object of the builder.
func (m *GoodsCreate) Mutation() *GoodsMutation {
	return m.mutation
}

// Save creates the Goods in the database.
func (c *GoodsCreate) Save(ctx context.Context) (*Goods, error) {
	return withHooks(ctx, c.gremlinSave, c.mutation, c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (c *GoodsCreate) SaveX(ctx context.Context) *Goods {
	v, err := c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (c *GoodsCreate) Exec(ctx context.Context) error {
	_, err := c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (c *GoodsCreate) ExecX(ctx context.Context) {
	if err := c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (c *GoodsCreate) check() error {
	return nil
}

func (c *GoodsCreate) gremlinSave(ctx context.Context) (*Goods, error) {
	if err := c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := c.gremlin().Query()
	if err := c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &Goods{config: c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	c.mutation.id = &rnode.ID
	c.mutation.done = true
	return rnode, nil
}

func (c *GoodsCreate) gremlin() *dsl.Traversal {
	v := g.AddV(goods.Label)
	return v.ValueMap(true)
}

// GoodsCreateBulk is the builder for creating many Goods entities in bulk.
type GoodsCreateBulk struct {
	config
	err      error
	builders []*GoodsCreate
}
