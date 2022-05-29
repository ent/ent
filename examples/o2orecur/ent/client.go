// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/o2orecur/ent/migrate"

	"entgo.io/ent/examples/o2orecur/ent/node"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Node = NewNodeClient(c.config)
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Node:   NewNodeClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Node:   NewNodeClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Node.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Node.Use(hooks...)
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `node.Hooks(f(g(h())))`.
func (c *NodeClient) Use(hooks ...Hook) {
	c.hooks.Node = append(c.hooks.Node, hooks...)
}

// Create returns a builder for creating a Node entity.
func (c *NodeClient) Create() *NodeCreate {
	mutation := newNodeMutation(c.config, OpCreate)
	return &NodeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Node entities.
func (c *NodeClient) CreateBulk(builders ...*NodeCreate) *NodeCreateBulk {
	return &NodeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	mutation := newNodeMutation(c.config, OpUpdate)
	return &NodeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNode(n))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id int) *NodeUpdateOne {
	mutation := newNodeMutation(c.config, OpUpdateOne, withNodeID(id))
	return &NodeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	mutation := newNodeMutation(c.config, OpDelete)
	return &NodeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *NodeClient) DeleteOneID(id int) *NodeDeleteOne {
	builder := c.Delete().Where(node.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &NodeDeleteOne{builder}
}

// Query returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{
		config: c.config,
	}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id int) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id int) *Node {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryPrev queries the prev edge of a Node.
func (c *NodeClient) QueryPrev(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, node.PrevTable, node.PrevColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryNext queries the next edge of a Node.
func (c *NodeClient) QueryNext(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := n.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(node.Table, node.FieldID, id),
			sqlgraph.To(node.Table, node.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, node.NextTable, node.NextColumn),
		)
		fromV = sqlgraph.Neighbors(n.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *NodeClient) Hooks() []Hook {
	return c.hooks.Node
}
