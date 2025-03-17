// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math/big"
	"net/url"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/exvaluescan"
)

// ExValueScanCreate is the builder for creating a ExValueScan entity.
type ExValueScanCreate struct {
	config
	mutation *ExValueScanMutation
	hooks    []Hook
}

// SetBinary sets the "binary" field.
func (_c *ExValueScanCreate) SetBinary(u *url.URL) *ExValueScanCreate {
	_c.mutation.SetBinary(u)
	return _c
}

// SetBinaryBytes sets the "binary_bytes" field.
func (_c *ExValueScanCreate) SetBinaryBytes(u *url.URL) *ExValueScanCreate {
	_c.mutation.SetBinaryBytes(u)
	return _c
}

// SetBinaryOptional sets the "binary_optional" field.
func (_c *ExValueScanCreate) SetBinaryOptional(u *url.URL) *ExValueScanCreate {
	_c.mutation.SetBinaryOptional(u)
	return _c
}

// SetText sets the "text" field.
func (_c *ExValueScanCreate) SetText(b *big.Int) *ExValueScanCreate {
	_c.mutation.SetText(b)
	return _c
}

// SetTextOptional sets the "text_optional" field.
func (_c *ExValueScanCreate) SetTextOptional(b *big.Int) *ExValueScanCreate {
	_c.mutation.SetTextOptional(b)
	return _c
}

// SetBase64 sets the "base64" field.
func (_c *ExValueScanCreate) SetBase64(s string) *ExValueScanCreate {
	_c.mutation.SetBase64(s)
	return _c
}

// SetCustom sets the "custom" field.
func (_c *ExValueScanCreate) SetCustom(s string) *ExValueScanCreate {
	_c.mutation.SetCustom(s)
	return _c
}

// SetCustomOptional sets the "custom_optional" field.
func (_c *ExValueScanCreate) SetCustomOptional(s string) *ExValueScanCreate {
	_c.mutation.SetCustomOptional(s)
	return _c
}

// SetNillableCustomOptional sets the "custom_optional" field if the given value is not nil.
func (_c *ExValueScanCreate) SetNillableCustomOptional(s *string) *ExValueScanCreate {
	if s != nil {
		_c.SetCustomOptional(*s)
	}
	return _c
}

// Mutation returns the ExValueScanMutation object of the builder.
func (_c *ExValueScanCreate) Mutation() *ExValueScanMutation {
	return _c.mutation
}

// Save creates the ExValueScan in the database.
func (_c *ExValueScanCreate) Save(ctx context.Context) (*ExValueScan, error) {
	return withHooks(ctx, _c.gremlinSave, _c.mutation, _c.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (_c *ExValueScanCreate) SaveX(ctx context.Context) *ExValueScan {
	v, err := _c.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (_c *ExValueScanCreate) Exec(ctx context.Context) error {
	_, err := _c.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (_c *ExValueScanCreate) ExecX(ctx context.Context) {
	if err := _c.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (_c *ExValueScanCreate) check() error {
	if _, ok := _c.mutation.Binary(); !ok {
		return &ValidationError{Name: "binary", err: errors.New(`ent: missing required field "ExValueScan.binary"`)}
	}
	if _, ok := _c.mutation.BinaryBytes(); !ok {
		return &ValidationError{Name: "binary_bytes", err: errors.New(`ent: missing required field "ExValueScan.binary_bytes"`)}
	}
	if _, ok := _c.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "ExValueScan.text"`)}
	}
	if _, ok := _c.mutation.Base64(); !ok {
		return &ValidationError{Name: "base64", err: errors.New(`ent: missing required field "ExValueScan.base64"`)}
	}
	if _, ok := _c.mutation.Custom(); !ok {
		return &ValidationError{Name: "custom", err: errors.New(`ent: missing required field "ExValueScan.custom"`)}
	}
	return nil
}

func (_c *ExValueScanCreate) gremlinSave(ctx context.Context) (*ExValueScan, error) {
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
	rnode := &ExValueScan{config: _c.config}
	if err := rnode.FromResponse(res); err != nil {
		return nil, err
	}
	_c.mutation.id = &rnode.ID
	_c.mutation.done = true
	return rnode, nil
}

func (_c *ExValueScanCreate) gremlin() *dsl.Traversal {
	v := g.AddV(exvaluescan.Label)
	if value, ok := _c.mutation.Binary(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinary, value)
	}
	if value, ok := _c.mutation.BinaryBytes(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryBytes, value)
	}
	if value, ok := _c.mutation.BinaryOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryOptional, value)
	}
	if value, ok := _c.mutation.Text(); ok {
		v.Property(dsl.Single, exvaluescan.FieldText, value)
	}
	if value, ok := _c.mutation.TextOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldTextOptional, value)
	}
	if value, ok := _c.mutation.Base64(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBase64, value)
	}
	if value, ok := _c.mutation.Custom(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustom, value)
	}
	if value, ok := _c.mutation.CustomOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustomOptional, value)
	}
	return v.ValueMap(true)
}

// ExValueScanCreateBulk is the builder for creating many ExValueScan entities in bulk.
type ExValueScanCreateBulk struct {
	config
	err      error
	builders []*ExValueScanCreate
}
