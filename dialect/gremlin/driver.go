package gremlin

import (
	"context"
	"fmt"

	"fbc/ent/dialect"
	"fbc/ent/dialect/gremlin/graph/dsl"
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
func (Driver) Dialect() string { return dialect.Neptune }

// Exec implements the dialect.Exec method.
func (c *Driver) Exec(ctx context.Context, query string, args, v interface{}) error {
	vr, ok := v.(*Response)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect *gremlin.Response", v)
	}
	bindings, ok := args.(dsl.Bindings)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect map[string]interface{} for bindings", args)
	}
	res, err := c.Do(ctx, NewEvalRequest(query, WithBindings(bindings)))
	if err != nil {
		return err
	}
	*vr = *res
	return nil
}

// Query implements the dialect.Query method.
func (c *Driver) Query(ctx context.Context, query string, args, v interface{}) error {
	return c.Exec(ctx, query, args, v)
}

// Close is a nop close call. It should close the connection in case of WS client.
func (c *Driver) Close() error { return nil }

// Tx returns a nop transaction.
func (c *Driver) Tx(context.Context) (dialect.Tx, error) { return c, nil }

// Commit is a nop commit.
func (c *Driver) Commit() error { return nil }

// Rollback is a nop rollback.
func (c *Driver) Rollback() error { return nil }

var _ dialect.Driver = (*Driver)(nil)
