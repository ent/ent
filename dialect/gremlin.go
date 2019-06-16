package dialect

import (
	"context"
	"fmt"

	"fbc/lib/go/gremlin"
	"fbc/lib/go/gremlin/graph/dsl"
)

// Gremlin is a dialect.Client implementation for TinkerPop gremlin.
type Gremlin struct {
	*gremlin.Client
}

// NewGremlin returns a new dialect.Gremlin implementation for gremlin.
func NewGremlin(c *gremlin.Client) *Gremlin {
	c.Transport = gremlin.ExpandBindings(c.Transport)
	return &Gremlin{c}
}

// Dialect implements the dialect.Dialect method.
func (Gremlin) Dialect() string { return Neptune }

// Exec implements the dialect.Exec method.
func (c *Gremlin) Exec(ctx context.Context, query string, args interface{}, v interface{}) error {
	vr, ok := v.(*gremlin.Response)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect *gremlin.Response", v)
	}
	bindings, ok := args.(dsl.Bindings)
	if !ok {
		return fmt.Errorf("dialect/gremlin: invalid type %T. expect map[string]interface{} for bindings", args)
	}
	res, err := c.Do(ctx, gremlin.NewEvalRequest(query, gremlin.WithBindings(bindings)))
	if err != nil {
		return err
	}
	*vr = *res
	return nil
}

// Query implements the dialect.Query method.
func (c *Gremlin) Query(ctx context.Context, query string, args interface{}, v interface{}) error {
	return c.Exec(ctx, query, args, v)
}

// Close is a nop close call. It should close the connection in case of WS client.
func (c *Gremlin) Close() error { return nil }

// Tx returns a nop transaction.
func (c *Gremlin) Tx(context.Context) (Tx, error) { return c, nil }

// Commit is a nop commit.
func (c *Gremlin) Commit() error { return nil }

// Rollback is a nop rollback.
func (c *Gremlin) Rollback() error { return nil }

var _ Driver = (*Gremlin)(nil)
