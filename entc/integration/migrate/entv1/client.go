// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"
	"log"

	"entgo.io/ent/entc/integration/migrate/entv1/migrate"

	"entgo.io/ent/entc/integration/migrate/entv1/car"
	"entgo.io/ent/entc/integration/migrate/entv1/conversion"
	"entgo.io/ent/entc/integration/migrate/entv1/customtype"
	"entgo.io/ent/entc/integration/migrate/entv1/user"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Car is the client for interacting with the Car builders.
	Car *CarClient
	// Conversion is the client for interacting with the Conversion builders.
	Conversion *ConversionClient
	// CustomType is the client for interacting with the CustomType builders.
	CustomType *CustomTypeClient
	// User is the client for interacting with the User builders.
	User *UserClient
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
	c.Conversion = NewConversionClient(c.config)
	c.CustomType = NewCustomTypeClient(c.config)
	c.User = NewUserClient(c.config)
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
		return nil, fmt.Errorf("entv1: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("entv1: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:        ctx,
		config:     cfg,
		Car:        NewCarClient(cfg),
		Conversion: NewConversionClient(cfg),
		CustomType: NewCustomTypeClient(cfg),
		User:       NewUserClient(cfg),
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
		config:     cfg,
		Car:        NewCarClient(cfg),
		Conversion: NewConversionClient(cfg),
		CustomType: NewCustomTypeClient(cfg),
		User:       NewUserClient(cfg),
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
	c.Car.Use(hooks...)
	c.Conversion.Use(hooks...)
	c.CustomType.Use(hooks...)
	c.User.Use(hooks...)
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

// CreateBulk returns a builder for creating a bulk of Car entities.
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
	return &CarQuery{
		config: c.config,
	}
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

// QueryOwner queries the owner edge of a Car.
func (c *CarClient) QueryOwner(ca *Car) *UserQuery {
	return c.QueryOwnerId(ca.ID)
}

// QueryOwnerId queries the owner edge of a Car by its id.
func (c *CarClient) QueryOwnerId(id int) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		step := sqlgraph.NewStep(
			sqlgraph.From(car.Table, car.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, car.OwnerTable, car.OwnerColumn),
		)
		fromV = sqlgraph.Neighbors(c.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *CarClient) Hooks() []Hook {
	return c.hooks.Car
}

// ConversionClient is a client for the Conversion schema.
type ConversionClient struct {
	config
}

// NewConversionClient returns a client for the Conversion from the given config.
func NewConversionClient(c config) *ConversionClient {
	return &ConversionClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `conversion.Hooks(f(g(h())))`.
func (c *ConversionClient) Use(hooks ...Hook) {
	c.hooks.Conversion = append(c.hooks.Conversion, hooks...)
}

// Create returns a create builder for Conversion.
func (c *ConversionClient) Create() *ConversionCreate {
	mutation := newConversionMutation(c.config, OpCreate)
	return &ConversionCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Conversion entities.
func (c *ConversionClient) CreateBulk(builders ...*ConversionCreate) *ConversionCreateBulk {
	return &ConversionCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Conversion.
func (c *ConversionClient) Update() *ConversionUpdate {
	mutation := newConversionMutation(c.config, OpUpdate)
	return &ConversionUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *ConversionClient) UpdateOne(co *Conversion) *ConversionUpdateOne {
	mutation := newConversionMutation(c.config, OpUpdateOne, withConversion(co))
	return &ConversionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *ConversionClient) UpdateOneID(id int) *ConversionUpdateOne {
	mutation := newConversionMutation(c.config, OpUpdateOne, withConversionID(id))
	return &ConversionUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Conversion.
func (c *ConversionClient) Delete() *ConversionDelete {
	mutation := newConversionMutation(c.config, OpDelete)
	return &ConversionDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ConversionClient) DeleteOne(co *Conversion) *ConversionDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ConversionClient) DeleteOneID(id int) *ConversionDeleteOne {
	builder := c.Delete().Where(conversion.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &ConversionDeleteOne{builder}
}

// Query returns a query builder for Conversion.
func (c *ConversionClient) Query() *ConversionQuery {
	return &ConversionQuery{
		config: c.config,
	}
}

// Get returns a Conversion entity by its id.
func (c *ConversionClient) Get(ctx context.Context, id int) (*Conversion, error) {
	return c.Query().Where(conversion.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ConversionClient) GetX(ctx context.Context, id int) *Conversion {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *ConversionClient) Hooks() []Hook {
	return c.hooks.Conversion
}

// CustomTypeClient is a client for the CustomType schema.
type CustomTypeClient struct {
	config
}

// NewCustomTypeClient returns a client for the CustomType from the given config.
func NewCustomTypeClient(c config) *CustomTypeClient {
	return &CustomTypeClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `customtype.Hooks(f(g(h())))`.
func (c *CustomTypeClient) Use(hooks ...Hook) {
	c.hooks.CustomType = append(c.hooks.CustomType, hooks...)
}

// Create returns a create builder for CustomType.
func (c *CustomTypeClient) Create() *CustomTypeCreate {
	mutation := newCustomTypeMutation(c.config, OpCreate)
	return &CustomTypeCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of CustomType entities.
func (c *CustomTypeClient) CreateBulk(builders ...*CustomTypeCreate) *CustomTypeCreateBulk {
	return &CustomTypeCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for CustomType.
func (c *CustomTypeClient) Update() *CustomTypeUpdate {
	mutation := newCustomTypeMutation(c.config, OpUpdate)
	return &CustomTypeUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *CustomTypeClient) UpdateOne(ct *CustomType) *CustomTypeUpdateOne {
	mutation := newCustomTypeMutation(c.config, OpUpdateOne, withCustomType(ct))
	return &CustomTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *CustomTypeClient) UpdateOneID(id int) *CustomTypeUpdateOne {
	mutation := newCustomTypeMutation(c.config, OpUpdateOne, withCustomTypeID(id))
	return &CustomTypeUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for CustomType.
func (c *CustomTypeClient) Delete() *CustomTypeDelete {
	mutation := newCustomTypeMutation(c.config, OpDelete)
	return &CustomTypeDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CustomTypeClient) DeleteOne(ct *CustomType) *CustomTypeDeleteOne {
	return c.DeleteOneID(ct.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CustomTypeClient) DeleteOneID(id int) *CustomTypeDeleteOne {
	builder := c.Delete().Where(customtype.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &CustomTypeDeleteOne{builder}
}

// Query returns a query builder for CustomType.
func (c *CustomTypeClient) Query() *CustomTypeQuery {
	return &CustomTypeQuery{
		config: c.config,
	}
}

// Get returns a CustomType entity by its id.
func (c *CustomTypeClient) Get(ctx context.Context, id int) (*CustomType, error) {
	return c.Query().Where(customtype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CustomTypeClient) GetX(ctx context.Context, id int) *CustomType {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *CustomTypeClient) Hooks() []Hook {
	return c.hooks.CustomType
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	return c.QueryParentId(u.ID)
}

// QueryParentId queries the parent edge of a User by its id.
func (c *UserClient) QueryParentId(id int) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, user.ParentTable, user.ParentColumn),
		)
		fromV = sqlgraph.Neighbors(c.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	return c.QueryChildrenId(u.ID)
}

// QueryChildrenId queries the children edge of a User by its id.
func (c *UserClient) QueryChildrenId(id int) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, user.ChildrenTable, user.ChildrenColumn),
		)
		fromV = sqlgraph.Neighbors(c.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QuerySpouse queries the spouse edge of a User.
func (c *UserClient) QuerySpouse(u *User) *UserQuery {
	return c.QuerySpouseId(u.ID)
}

// QuerySpouseId queries the spouse edge of a User by its id.
func (c *UserClient) QuerySpouseId(id int) *UserQuery {
	query := &UserQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.SpouseTable, user.SpouseColumn),
		)
		fromV = sqlgraph.Neighbors(c.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryCar queries the car edge of a User.
func (c *UserClient) QueryCar(u *User) *CarQuery {
	return c.QueryCarId(u.ID)
}

// QueryCarId queries the car edge of a User by its id.
func (c *UserClient) QueryCarId(id int) *CarQuery {
	query := &CarQuery{config: c.config}
	query.path = func(ctx context.Context) (fromV *sql.Selector, _ error) {
		step := sqlgraph.NewStep(
			sqlgraph.From(user.Table, user.FieldID, id),
			sqlgraph.To(car.Table, car.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, user.CarTable, user.CarColumn),
		)
		fromV = sqlgraph.Neighbors(c.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}
