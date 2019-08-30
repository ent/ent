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

// Where adds a new predicate for the builder.
func (sd *StreetDelete) Where(ps ...predicate.Street) *StreetDelete {
	sd.predicates = append(sd.predicates, ps...)
	return sd
}

// Exec executes the deletion query.
func (sd *StreetDelete) Exec(ctx context.Context) error {
	return sd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sd *StreetDelete) ExecX(ctx context.Context) {
	if err := sd.Exec(ctx); err != nil {
		panic(err)
	}
}

func (sd *StreetDelete) sqlExec(ctx context.Context) error {
	var res sql.Result
	selector := sql.Select().From(sql.Table(street.Table))
	for _, p := range sd.predicates {
		p(selector)
	}
	query, args := sql.Delete(street.Table).FromSelect(selector).Query()
	return sd.driver.Exec(ctx, query, args, &res)
}

// StreetDeleteOne is the builder for deleting a single Street entity.
type StreetDeleteOne struct {
	sd *StreetDelete
}

// Exec executes the deletion query.
func (sdo *StreetDeleteOne) Exec(ctx context.Context) error {
	return sdo.sd.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdo *StreetDeleteOne) ExecX(ctx context.Context) {
	sdo.sd.ExecX(ctx)
}
