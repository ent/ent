// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// GroupInfoDelete is the builder for deleting a GroupInfo entity.
type GroupInfoDelete struct {
	config
	predicates []predicate.GroupInfo
}

// Where adds a new predicate to the delete builder.
func (gid *GroupInfoDelete) Where(ps ...predicate.GroupInfo) *GroupInfoDelete {
	gid.predicates = append(gid.predicates, ps...)
	return gid
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (gid *GroupInfoDelete) Exec(ctx context.Context) (int, error) {
	return gid.sqlExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (gid *GroupInfoDelete) ExecX(ctx context.Context) int {
	n, err := gid.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (gid *GroupInfoDelete) sqlExec(ctx context.Context) (int, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gid.driver.Dialect())
	)
	selector := builder.Select().From(sql.Table(groupinfo.Table))
	for _, p := range gid.predicates {
		p(selector)
	}
	query, args := builder.Delete(groupinfo.Table).FromSelect(selector).Query()
	if err := gid.driver.Exec(ctx, query, args, &res); err != nil {
		return 0, err
	}
	affected, err := res.RowsAffected()
	if err != nil {
		return 0, err
	}
	return int(affected), nil
}

// GroupInfoDeleteOne is the builder for deleting a single GroupInfo entity.
type GroupInfoDeleteOne struct {
	gid *GroupInfoDelete
}

// Exec executes the deletion query.
func (gido *GroupInfoDeleteOne) Exec(ctx context.Context) error {
	n, err := gido.gid.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{groupinfo.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (gido *GroupInfoDeleteOne) ExecX(ctx context.Context) {
	gido.gid.ExecX(ctx)
}
