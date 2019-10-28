// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	predicates []predicate.Item
}

// Where adds a new predicate for the builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.predicates = append(iu.predicates, ps...)
	return iu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	switch iu.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return iu.sqlSave(ctx)
	case dialect.Gremlin:
		return iu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(iu.driver.Dialect())
		selector = builder.Select(item.FieldID).From(builder.Table(item.Table))
	)
	for _, p := range iu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = iu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := iu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (iu *ItemUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := iu.gremlin().Query()
	if err := iu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (iu *ItemUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(item.Label)
	for _, p := range iu.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	id string
}

// Save executes the query and returns the updated entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	switch iuo.driver.Dialect() {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		return iuo.sqlSave(ctx)
	case dialect.Gremlin:
		return iuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	i, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return i
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (i *Item, err error) {
	var (
		builder  = sql.Dialect(iuo.driver.Dialect())
		selector = builder.Select(item.Columns...).From(builder.Table(item.Table))
	)
	item.ID(iuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = iuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		i = &Item{config: iuo.config}
		if err := i.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Item: %v", err)
		}
		id = i.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Item with id: %v", iuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Item with the same id: %v", iuo.id)
	}

	tx, err := iuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return i, nil
}

func (iuo *ItemUpdateOne) gremlinSave(ctx context.Context) (*Item, error) {
	res := &gremlin.Response{}
	query, bindings := iuo.gremlin(iuo.id).Query()
	if err := iuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	i := &Item{config: iuo.config}
	if err := i.FromResponse(res); err != nil {
		return nil, err
	}
	return i, nil
}

func (iuo *ItemUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	v.ValueMap(true)
	trs = append(trs, v)
	return dsl.Join(trs...)
}
