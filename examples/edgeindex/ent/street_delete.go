// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"

	"github.com/facebookincubator/ent/dialect/sql"
)

// StreetDelete is the builder for deleting a Street entity.
type StreetDelete struct {
	config
	predicates []predicate.Street
}

// Where adds a new predicate to the delete builder.
func (sd *StreetDelete) Where(ps ...predicate.Street) *StreetDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sd *StreetDelete) Exec(ctx context.Context) (int, error) {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StreetDelete) ExecX(ctx context.Context) int {
	n, err := sd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sd *StreetDelete) sqlExec(ctx context.Context) (int, error) {
	var res sql.Result
	selector := sql.Select().From(sql.Table(street.Table))
	for _, p := range sd.predicates {
		p(selector)
	}
	query, args := sql.Delete(street.Table).FromSelect(selector).Query()
	if err := sd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// StreetDeleteOne is the builder for deleting a single Street entity.
type StreetDeleteOne struct {
	sd *StreetDelete
}

// Exec executes the deletion query.
func (sdo *StreetDeleteOne) Exec(ctx context.Context) error {
	n, err := sdo.sd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{street.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StreetDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
