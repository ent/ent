// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/license"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// LicenseDelete is the builder for deleting a License entity.
type LicenseDelete struct {
	config
	hooks    []Hook
	mutation *LicenseMutation
}

// Where appends a list predicates to the LicenseDelete builder.
func (_d *LicenseDelete) Where(ps ...predicate.License) *LicenseDelete {
	_d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (_d *LicenseDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, _d.gremlinExec, _d.mutation, _d.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *LicenseDelete) ExecX(ctx context.Context) int {
	n, err := _d.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (_d *LicenseDelete) gremlinExec(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _d.gremlin().Query()
	if err := _d.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	_d.mutation.done = true
	return res.ReadInt()
}

func (_d *LicenseDelete) gremlin() *dsl.Traversal {
	t := g.V().HasLabel(license.Label)
	for _, p := range _d.mutation.predicates {
		p(t)
	}
	return t.SideEffect(__.Drop()).Count()
}

// LicenseDeleteOne is the builder for deleting a single License entity.
type LicenseDeleteOne struct {
	_d *LicenseDelete
}

// Where appends a list predicates to the LicenseDelete builder.
func (_d *LicenseDeleteOne) Where(ps ...predicate.License) *LicenseDeleteOne {
	_d._d.mutation.Where(ps...)
	return _d
}

// Exec executes the deletion query.
func (_d *LicenseDeleteOne) Exec(ctx context.Context) error {
	n, err := _d._d.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{license.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (_d *LicenseDeleteOne) ExecX(ctx context.Context) {
	if err := _d.Exec(ctx); err != nil {
		panic(err)
	}
}
