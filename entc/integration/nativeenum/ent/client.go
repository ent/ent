// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebook/ent/entc/integration/nativeenum/ent/migrate"

	"github.com/facebook/ent/entc/integration/nativeenum/ent/car"
	"github.com/facebook/ent/entc/integration/nativeenum/ent/person"

	"github.com/facebook/ent/dialect"
	"github.com/facebook/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Person is the client for interacting with the Person builders.
	Person *PersonClient
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
	c.Car = NewCarClient(c.config)
	c.Person = NewPersonClient(c.config)
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
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: tx, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		ctx:    ctx,
		config: cfg,
		Car:    NewCarClient(cfg),
		Person: NewPersonClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, fmt.Errorf("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(*sql.Driver).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %v", err)
	}
	cfg := config{driver: &txDriver{tx: tx, drv: c.driver}, log: c.log, debug: c.debug, hooks: c.hooks}
	return &Tx{
		config: cfg,
		Car:    NewCarClient(cfg),
		Person: NewPersonClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Car.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true, hooks: c.hooks}
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
	c.Car.Use(hooks...)
	c.Person.Use(hooks...)
}

// CarClient is a client for the Car schema.
type CarClient struct {
	config
}

// NewCarClient returns a client for the Car from the given config.
func NewCarClient(c config) *CarClient {
	return &CarClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `car.Hooks(f(g(h())))`.
func (c *CarClient) Use(hooks ...Hook) {
	c.hooks.Car = append(c.hooks.Car, hooks...)
}

// Create returns a create builder for Car.
func (c *CarClient) Create() *CarCreate {
	mutation := newCarMutation(c.config, OpCreate)
	return &CarCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Car entities.
func (c *CarClient) CreateBulk(builders ...*CarCreate) *CarCreateBulk {
	return &CarCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Car.
func (c *CarClient) Update() *CarUpdate {
	mutation := newCarMutation(c.config, OpUpdate)
	return &CarUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CarClient) UpdateOne(ca *Car) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCar(ca))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CarClient) UpdateOneID(id int) *CarUpdateOne {
	mutation := newCarMutation(c.config, OpUpdateOne, withCarID(id))
	return &CarUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Car.
func (c *CarClient) Delete() *CarDelete {
	mutation := newCarMutation(c.config, OpDelete)
	return &CarDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CarClient) DeleteOne(ca *Car) *CarDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CarClient) DeleteOneID(id int) *CarDeleteOne {
	builder := c.Delete().Where(car.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CarDeleteOne{builder}
}

// Query returns a query builder for Car.
func (c *CarClient) Query() *CarQuery {
	return &CarQuery{config: c.config}
}

// Get returns a Car entity by its id.
func (c *CarClient) Get(ctx context.Context, id int) (*Car, error) {
	return c.Query().Where(car.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CarClient) GetX(ctx context.Context, id int) *Car {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// PersonClient is a client for the Person schema.
type PersonClient struct {
	config
}

// NewPersonClient returns a client for the Person from the given config.
func NewPersonClient(c config) *PersonClient {
	return &PersonClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `person.Hooks(f(g(h())))`.
func (c *PersonClient) Use(hooks ...Hook) {
	c.hooks.Person = append(c.hooks.Person, hooks...)
}

// Create returns a create builder for Person.
func (c *PersonClient) Create() *PersonCreate {
	mutation := newPersonMutation(c.config, OpCreate)
	return &PersonCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// BulkCreate returns a builder for creating a bulk of Person entities.
func (c *PersonClient) CreateBulk(builders ...*PersonCreate) *PersonCreateBulk {
	return &PersonCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Person.
func (c *PersonClient) Update() *PersonUpdate {
	mutation := newPersonMutation(c.config, OpUpdate)
	return &PersonUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *PersonClient) UpdateOne(pe *Person) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPerson(pe))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *PersonClient) UpdateOneID(id int) *PersonUpdateOne {
	mutation := newPersonMutation(c.config, OpUpdateOne, withPersonID(id))
	return &PersonUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Person.
func (c *PersonClient) Delete() *PersonDelete {
	mutation := newPersonMutation(c.config, OpDelete)
	return &PersonDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PersonClient) DeleteOne(pe *Person) *PersonDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PersonClient) DeleteOneID(id int) *PersonDeleteOne {
	builder := c.Delete().Where(person.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &PersonDeleteOne{builder}
}

// Query returns a query builder for Person.
func (c *PersonClient) Query() *PersonQuery {
	return &PersonQuery{config: c.config}
}

// Get returns a Person entity by its id.
func (c *PersonClient) Get(ctx context.Context, id int) (*Person, error) {
	return c.Query().Where(person.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PersonClient) GetX(ctx context.Context, id int) *Person {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *PersonClient) Hooks() []Hook {
	return c.hooks.Person
}
