// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/examples/edgeindex/ent/migrate"

	"entgo.io/ent/examples/edgeindex/ent/city"
	"entgo.io/ent/examples/edgeindex/ent/street"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// City is the client for interacting with the City builders.
	City *CityClient
	// Street is the client for interacting with the Street builders.
	Street *StreetClient
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
	c.City = NewCityClient(c.config)
	c.Street = NewStreetClient(c.config)
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
		City:   NewCityClient(cfg),
		Street: NewStreetClient(cfg),
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
		City:   NewCityClient(cfg),
		Street: NewStreetClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		City.
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
	c.City.Use(hooks...)
	c.Street.Use(hooks...)
}

// CityClient is a client for the City schema.
type CityClient struct {
	config
}

// NewCityClient returns a client for the City from the given config.
func NewCityClient(c config) *CityClient {
	return &CityClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `city.Hooks(f(g(h())))`.
func (c *CityClient) Use(hooks ...Hook) {
	c.hooks.City = append(c.hooks.City, hooks...)
}

// Create returns a builder for creating a City entity.
func (c *CityClient) Create() *CityCreate {
	mutation := newCityMutation(c.config, OpCreate)
	return &CityCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of City entities.
func (c *CityClient) CreateBulk(builders ...*CityCreate) *CityCreateBulk {
	return &CityCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for City.
func (c *CityClient) Update() *CityUpdate {
	mutation := newCityMutation(c.config, OpUpdate)
	return &CityUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CityClient) UpdateOne(ci *City) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCity(ci))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CityClient) UpdateOneID(id int) *CityUpdateOne {
	mutation := newCityMutation(c.config, OpUpdateOne, withCityID(id))
	return &CityUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for City.
func (c *CityClient) Delete() *CityDelete {
	mutation := newCityMutation(c.config, OpDelete)
	return &CityDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *CityClient) DeleteOne(ci *City) *CityDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *CityClient) DeleteOneID(id int) *CityDeleteOne {
	builder := c.Delete().Where(city.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CityDeleteOne{builder}
}

// Query returns a query builder for City.
func (c *CityClient) Query() *CityQuery {
	return &CityQuery{
		config: c.config,
	}
}

// Get returns a City entity by its id.
func (c *CityClient) Get(ctx context.Context, id int) (*City, error) {
	return c.Query().Where(city.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CityClient) GetX(ctx context.Context, id int) *City {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryStreets queries the streets edge of a City.
func (c *CityClient) QueryStreets(ci *City) *StreetQuery {
	query := &StreetQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := ci.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(city.Table, city.FieldID, id),
			sqlgraph.To(street.Table, street.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, city.StreetsTable, city.StreetsColumn),
		)
		fromV = sqlgraph.Neighbors(ci.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CityClient) Hooks() []Hook {
	return c.hooks.City
}

// StreetClient is a client for the Street schema.
type StreetClient struct {
	config
}

// NewStreetClient returns a client for the Street from the given config.
func NewStreetClient(c config) *StreetClient {
	return &StreetClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `street.Hooks(f(g(h())))`.
func (c *StreetClient) Use(hooks ...Hook) {
	c.hooks.Street = append(c.hooks.Street, hooks...)
}

// Create returns a builder for creating a Street entity.
func (c *StreetClient) Create() *StreetCreate {
	mutation := newStreetMutation(c.config, OpCreate)
	return &StreetCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Street entities.
func (c *StreetClient) CreateBulk(builders ...*StreetCreate) *StreetCreateBulk {
	return &StreetCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Street.
func (c *StreetClient) Update() *StreetUpdate {
	mutation := newStreetMutation(c.config, OpUpdate)
	return &StreetUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *StreetClient) UpdateOne(s *Street) *StreetUpdateOne {
	mutation := newStreetMutation(c.config, OpUpdateOne, withStreet(s))
	return &StreetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *StreetClient) UpdateOneID(id int) *StreetUpdateOne {
	mutation := newStreetMutation(c.config, OpUpdateOne, withStreetID(id))
	return &StreetUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Street.
func (c *StreetClient) Delete() *StreetDelete {
	mutation := newStreetMutation(c.config, OpDelete)
	return &StreetDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *StreetClient) DeleteOne(s *Street) *StreetDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOne returns a builder for deleting the given entity by its id.
func (c *StreetClient) DeleteOneID(id int) *StreetDeleteOne {
	builder := c.Delete().Where(street.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &StreetDeleteOne{builder}
}

// Query returns a query builder for Street.
func (c *StreetClient) Query() *StreetQuery {
	return &StreetQuery{
		config: c.config,
	}
}

// Get returns a Street entity by its id.
func (c *StreetClient) Get(ctx context.Context, id int) (*Street, error) {
	return c.Query().Where(street.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *StreetClient) GetX(ctx context.Context, id int) *Street {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryCity queries the city edge of a Street.
func (c *StreetClient) QueryCity(s *Street) *CityQuery {
	query := &CityQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(street.Table, street.FieldID, id),
			sqlgraph.To(city.Table, city.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, street.CityTable, street.CityColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *StreetClient) Hooks() []Hook {
	return c.hooks.Street
}
