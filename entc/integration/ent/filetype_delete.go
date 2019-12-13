// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// FileTypeDelete is the builder for deleting a FileType entity.
type FileTypeDelete struct {
	config
	predicates []predicate.FileType
}

// Where adds a new predicate to the delete builder.
func (ftd *FileTypeDelete) Where(ps ...predicate.FileType) *FileTypeDelete {
	ftd.predicates = append(ftd.predicates, ps...)
	return ftd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ftd *FileTypeDelete) Exec(ctx context.Context) (int, error) {
	return ftd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (ftd *FileTypeDelete) ExecX(ctx context.Context) int {
	n, err := ftd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ftd *FileTypeDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(ftd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(filetype.Table))
	for _, p := range ftd.predicates {
		p(selector)
	}
	query, args := builder.Delete(filetype.Table).FromSelect(selector).Query()
	if err := ftd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// FileTypeDeleteOne is the builder for deleting a single FileType entity.
type FileTypeDeleteOne struct {
	ftd *FileTypeDelete
}

// Exec executes the deletion query.
func (ftdo *FileTypeDeleteOne) Exec(ctx context.Context) error {
	n, err := ftdo.ftd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{filetype.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ftdo *FileTypeDeleteOne) ExecX(ctx context.Context) {
	ftdo.ftd.ExecX(ctx)
}
