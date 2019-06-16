// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"fbc/ent/entc/integration/plugin/ent/migrate"

	"fbc/ent/entc/integration/plugin/ent/boring"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Boring is the client for interacting with the Boring builders.
	Boring *BoringClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		Boring: NewBoringClient(c),
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
		Boring: NewBoringClient(cfg),
	}, nil
}

// BoringClient is a client for the Boring schema.
type BoringClient struct {
	config
}

// NewBoringClient returns a client for the Boring from the given config.
func NewBoringClient(c config) *BoringClient {
	return &BoringClient{config: c}
}

// Create returns a create builder for Boring.
func (c *BoringClient) Create() *BoringCreate {
	return &BoringCreate{config: c.config}
}

// Update returns an update builder for Boring.
func (c *BoringClient) Update() *BoringUpdate {
	return &BoringUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *BoringClient) UpdateOne(b *Boring) *BoringUpdateOne {
	return c.UpdateOneID(b.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *BoringClient) UpdateOneID(id string) *BoringUpdateOne {
	return &BoringUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Boring.
func (c *BoringClient) Delete() *BoringDelete {
	return &BoringDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *BoringClient) DeleteOne(b *Boring) *BoringDeleteOne {
	return c.DeleteOneID(b.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *BoringClient) DeleteOneID(id string) *BoringDeleteOne {
	return &BoringDeleteOne{c.Delete().Where(boring.ID(id))}
}

// Create returns a query builder for Boring.
func (c *BoringClient) Query() *BoringQuery {
	return &BoringQuery{config: c.config}
}
