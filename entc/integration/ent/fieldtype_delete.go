// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// FieldTypeDelete is the builder for deleting a FieldType entity.
type FieldTypeDelete struct {
	config
	predicates []predicate.FieldType
}

// Where adds a new predicate to the delete builder.
func (ftd *FieldTypeDelete) Where(ps ...predicate.FieldType) *FieldTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FieldTypeDelete) Exec(ctx context.Context) (int, error) {
	return ftd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FieldTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FieldTypeDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(ftd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(fieldtype.Table))
	for _, p := range ftd.predicates {
		p(selector)
	}
	query, args := builder.Delete(fieldtype.Table).FromSelect(selector).Query()
	if err := ftd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// FieldTypeDeleteOne is the builder for deleting a single FieldType entity.
type FieldTypeDeleteOne struct {
	ftd *FieldTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FieldTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{fieldtype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FieldTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
