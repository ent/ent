// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

package gremlin

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
)

// Driver is a dialect.Driver implementation for TinkerPop gremlin.
type Driver struct {
	*Client
}

// NewDriver returns a new dialect.Driver implementation for gremlin.
func NewDriver(c *Client) *Driver {
	c.Transport = ExpandBindings(c.Transport)
	return &Driver{c}
}

// Dialect implements the dialect.Dialect method.
func (Driver) Dialect() string { return dialect.Gremlin }

// Exec implements the dialect.Exec method.
func (c *Driver) Exec(ctx context.Context, query string, args, v any) error {
	vr, ok := v.(*Response)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect *gremlin.Response", v)
	}
	bindings, ok := args.(dsl.Bindings)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect map[string]any for bindings", args)
	}
	res, err := c.Do(ctx, NewEvalRequest(query, WithBindings(bindings)))
	if err != nil {
		return err
	}
	*vr = *res
	return nil
}

// Query implements the dialect.Query method.
func (c *Driver) Query(ctx context.Context, query string, args, v any) error {
	return c.Exec(ctx, query, args, v)
}

// Close is a nop close call. It should close the connection in case of WS client.
func (Driver) Close() error { return nil }

// Tx returns a nop transaction.
func (c *Driver) Tx(context.Context) (dialect.Tx, error) { return dialect.NopTx(c), nil }

var _ dialect.Driver = (*Driver)(nil)
