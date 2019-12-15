// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/comment"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// CommentDelete is the builder for deleting a Comment entity.
type CommentDelete struct {
	config
	predicates []predicate.Comment
}

// Where adds a new predicate to the delete builder.
func (cd *CommentDelete) Where(ps ...predicate.Comment) *CommentDelete {
	cd.predicates = append(cd.predicates, ps...)
	return cd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cd *CommentDelete) Exec(ctx context.Context) (int, error) {
	return cd.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (cd *CommentDelete) ExecX(ctx context.Context) int {
	n, err := cd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cd *CommentDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(cd.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(comment.Table))
	for _, p := range cd.predicates {
		p(selector)
	}
	query, args := builder.Delete(comment.Table).FromSelect(selector).Query()
	if err := cd.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// CommentDeleteOne is the builder for deleting a single Comment entity.
type CommentDeleteOne struct {
	cd *CommentDelete
}

// Exec executes the deletion query.
func (cdo *CommentDeleteOne) Exec(ctx context.Context) error {
	n, err := cdo.cd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{comment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cdo *CommentDeleteOne) ExecX(ctx context.Context) {
	cdo.cd.ExecX(ctx)
}
