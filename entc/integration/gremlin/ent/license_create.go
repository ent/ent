// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/license"
)

// LicenseCreate is the builder for creating a License entity.
type LicenseCreate struct {
	config
	mutation *LicenseMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (_c *LicenseCreate) SetCreateTime(t time.Time) *LicenseCreate {
	_c.mutation.SetCreateTime(t)
	return _c
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (_c *LicenseCreate) SetNillableCreateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		_c.SetCreateTime(*t)
	}
	return _c
}

// SetUpdateTime sets the "update_time" field.
func (_c *LicenseCreate) SetUpdateTime(t time.Time) *LicenseCreate {
	_c.mutation.SetUpdateTime(t)
	return _c
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (_c *LicenseCreate) SetNillableUpdateTime(t *time.Time) *LicenseCreate {
	if t != nil {
		_c.SetUpdateTime(*t)
	}
	return _c
}

// SetID sets the "id" field.
func (_c *LicenseCreate) SetID(i int) *LicenseCreate {
	_c.mutation.SetID(i)
	return _c
}

// Mutation returns the LicenseMutation object of the builder.
func (_c *LicenseCreate) Mutation() *LicenseMutation {
	return _c.mutation
}

// Save creates the License in the database.
func (_c *LicenseCreate) Save(ctx context.Context) (*License, error) {
	_c.defaults()
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *LicenseCreate) SaveX(ctx context.Context) *License {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *LicenseCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *LicenseCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (_c *LicenseCreate) defaults() {
	if _, ok := _c.mutation.CreateTime(); !ok {
		v := license.DefaultCreateTime()
		_c.mutation.SetCreateTime(v)
	}
	if _, ok := _c.mutation.UpdateTime(); !ok {
		v := license.DefaultUpdateTime()
		_c.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *LicenseCreate) check() error {
	if _, ok := _c.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "License.create_time"`)}
	}
	if _, ok := _c.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "License.update_time"`)}
	}
	return nil
}

func (_c *LicenseCreate) gremlinSave(ctx context.Context) (*License, error) {
	if err := _c.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := _c.gremlin().Query()
	if err := _c.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	rnode := &License{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *LicenseCreate) gremlin() *dsl.Traversal {
	v := g.AddV(license.Label)
	if id, ok := _c.mutation.ID(); ok {
		v.Property(dsl.ID, id)
	}
	if value, ok := _c.mutation.CreateTime(); ok {
		v.Property(dsl.Single, license.FieldCreateTime, value)
	}
	if value, ok := _c.mutation.UpdateTime(); ok {
		v.Property(dsl.Single, license.FieldUpdateTime, value)
	}
	return v.ValueMap(true)
}

// LicenseCreateBulk is the builder for creating many License entities in bulk.
type LicenseCreateBulk struct {
	config
	err      error
	builders []*LicenseCreate
}
