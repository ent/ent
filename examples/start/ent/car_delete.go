// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/start/ent/car"
	"github.com/facebookincubator/ent/examples/start/ent/predicate"
)

// CarDelete is the builder for deleting a Car entity.
type CarDelete struct {
	config
	predicates []predicate.Car
}

// Where adds a new predicate to the delete builder.
func (cd *CarDelete) Where(ps ...predicate.Car) *CarDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CarDelete) Exec(ctx context.Context) (int, error) {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CarDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CarDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(car.Table))
	for _, p := range cd.predicates {
		p(selector)
	}
	query, args := sql.Delete(car.Table).FromSelect(selector).Query()
	if err := cd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// CarDeleteOne is the builder for deleting a single Car entity.
type CarDeleteOne struct {
	cd *CarDelete
}

// Exec executes the deletion query.
func (cdo *CarDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{car.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CarDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
