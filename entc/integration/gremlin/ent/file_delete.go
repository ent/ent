// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/file"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// FileDelete is the builder for deleting a File entity.
type FileDelete struct {
	config
	predicates []predicate.File
}

// Where adds a new predicate to the delete builder.
func (fd *FileDelete) Where(ps ...predicate.File) *FileDelete {
	fd.predicates = append(fd.predicates, ps...)
	return fd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fd *FileDelete) Exec(ctx context.Context) (int, error) {
	return fd.gremlinExec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (fd *FileDelete) ExecX(ctx context.Context) int {
	n, err := fd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fd *FileDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := fd.gremlin().Query()
	if err := fd.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (fd *FileDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(file.Label)
	for _, p := range fd.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// FileDeleteOne is the builder for deleting a single File entity.
type FileDeleteOne struct {
	fd *FileDelete
}

// Exec executes the deletion query.
func (fdo *FileDeleteOne) Exec(ctx context.Context) error {
	n, err := fdo.fd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &ErrNotFound{file.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fdo *FileDeleteOne) ExecX(ctx context.Context) {
	fdo.fd.ExecX(ctx)
}
