// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/builder"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// BuilderUpdate is the builder for updating Builder entities.
type BuilderUpdate struct {
	config
	hooks    []Hook
	mutation *BuilderMutation
}

// Where appends a list predicates to the BuilderUpdate builder.
func (_u *BuilderUpdate) Where(ps ...predicate.Builder) *BuilderUpdate {
	_u.mutation.Where(ps...)
	return _u
}

// Mutation returns the BuilderMutation object of the builder.
func (_u *BuilderUpdate) Mutation() *BuilderMutation {
	return _u.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (_u *BuilderUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *BuilderUpdate) SaveX(ctx context.Context) int {
	affected, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (_u *BuilderUpdate) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *BuilderUpdate) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *BuilderUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _u.gremlin().Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	_u.mutation.done = true
	return res.ReadInt()
}

func (_u *BuilderUpdate) gremlin() *dsl.Traversal {
	v := g.V().HasLabel(builder.Label)
	for _, p := range _u.mutation.predicates {
		p(v)
	}
	var (
		trs []*dsl.Traversal
	)
	v.Count()
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// BuilderUpdateOne is the builder for updating a single Builder entity.
type BuilderUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *BuilderMutation
}

// Mutation returns the BuilderMutation object of the builder.
func (_u *BuilderUpdateOne) Mutation() *BuilderMutation {
	return _u.mutation
}

// Where appends a list predicates to the BuilderUpdate builder.
func (_u *BuilderUpdateOne) Where(ps ...predicate.Builder) *BuilderUpdateOne {
	_u.mutation.Where(ps...)
	return _u
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (_u *BuilderUpdateOne) Select(field string, fields ...string) *BuilderUpdateOne {
	_u.fields = append([]string{field}, fields...)
	return _u
}

// Save executes the query and returns the updated Builder entity.
func (_u *BuilderUpdateOne) Save(ctx context.Context) (*Builder, error) {
	return withHooks(ctx, _u.gremlinSave, _u.mutation, _u.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (_u *BuilderUpdateOne) SaveX(ctx context.Context) *Builder {
	node, err := _u.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (_u *BuilderUpdateOne) Exec(ctx context.Context) error {
	_, err := _u.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_u *BuilderUpdateOne) ExecX(ctx context.Context) {
	if err := _u.Exec(ctx); err != nil {
		panic(err)
	}
}

func (_u *BuilderUpdateOne) gremlinSave(ctx context.Context) (*Builder, error) {
	res := &gremlin.Response{}
	id, ok := _u.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Builder.id" for update`)}
	}
	query, bindings := _u.gremlin(id).Query()
	if err := _u.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	_u.mutation.done = true
	_m := &Builder{config: _u.config}
	if err := _m.FromResponse(res); err != nil {
		return nil, err
	}
	return _m, nil
}

func (_u *BuilderUpdateOne) gremlin(id string) *dsl.Traversal {
	v := g.V(id)
	var (
		trs []*dsl.Traversal
	)
	if len(_u.fields) > 0 {
		fields := make([]any, 0, len(_u.fields)+1)
		fields = append(fields, true)
		for _, f := range _u.fields {
			fields = append(fields, f)
		}
		v.ValueMap(fields...)
	} else {
		v.ValueMap(true)
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
