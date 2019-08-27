// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent/migrate"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"

	"github.com/facebookincubator/ent/dialect/sql"
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
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Node:   NewNodeClient(c),
	}
}

// Tx returns a new transactional client.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, verbose: c.verbose}
	return &Tx{
		config: cfg,
		Node:   NewNodeClient(cfg),
	}, nil
}

// NodeClient is a client for the Node schema.
type NodeClient struct {
	config
}

// NewNodeClient returns a client for the Node from the given config.
func NewNodeClient(c config) *NodeClient {
	return &NodeClient{config: c}
}

// Create returns a create builder for Node.
func (c *NodeClient) Create() *NodeCreate {
	return &NodeCreate{config: c.config}
}

// Update returns an update builder for Node.
func (c *NodeClient) Update() *NodeUpdate {
	return &NodeUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *NodeClient) UpdateOne(n *Node) *NodeUpdateOne {
	return c.UpdateOneID(n.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *NodeClient) UpdateOneID(id int) *NodeUpdateOne {
	return &NodeUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Node.
func (c *NodeClient) Delete() *NodeDelete {
	return &NodeDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *NodeClient) DeleteOne(n *Node) *NodeDeleteOne {
	return c.DeleteOneID(n.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *NodeClient) DeleteOneID(id int) *NodeDeleteOne {
	return &NodeDeleteOne{c.Delete().Where(node.ID(id))}
}

// Create returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{config: c.config}
}

// QueryParent queries the parent edge of a Node.
func (c *NodeClient) QueryParent(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	id := n.ID
	t1 := sql.Table(node.Table)
	t2 := sql.Select(node.ParentColumn).
		From(sql.Table(node.ParentTable)).
		Where(sql.EQ(node.FieldID, id))
	query.sql = sql.Select().From(t1).Join(t2).On(t1.C(node.FieldID), t2.C(node.ParentColumn))

	return query
}

// QueryChildren queries the children edge of a Node.
func (c *NodeClient) QueryChildren(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	id := n.ID
	query.sql = sql.Select().From(sql.Table(node.Table)).
		Where(sql.EQ(node.ChildrenColumn, id))

	return query
}
