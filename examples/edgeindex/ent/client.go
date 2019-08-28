// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/migrate"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
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
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		City:   NewCityClient(c),
		Street: NewStreetClient(c),
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
	cfg := config{driver: tx, log: c.log, debug: c.debug}
	return &Tx{
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
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config: cfg,
		Schema: migrate.NewSchema(cfg.driver),
		City:   NewCityClient(cfg),
		Street: NewStreetClient(cfg),
	}
}

// CityClient is a client for the City schema.
type CityClient struct {
	config
}

// NewCityClient returns a client for the City from the given config.
func NewCityClient(c config) *CityClient {
	return &CityClient{config: c}
}

// Create returns a create builder for City.
func (c *CityClient) Create() *CityCreate {
	return &CityCreate{config: c.config}
}

// Update returns an update builder for City.
func (c *CityClient) Update() *CityUpdate {
	return &CityUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *CityClient) UpdateOne(ci *City) *CityUpdateOne {
	return c.UpdateOneID(ci.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *CityClient) UpdateOneID(id int) *CityUpdateOne {
	return &CityUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for City.
func (c *CityClient) Delete() *CityDelete {
	return &CityDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CityClient) DeleteOne(ci *City) *CityDeleteOne {
	return c.DeleteOneID(ci.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CityClient) DeleteOneID(id int) *CityDeleteOne {
	return &CityDeleteOne{c.Delete().Where(city.ID(id))}
}

// Create returns a query builder for City.
func (c *CityClient) Query() *CityQuery {
	return &CityQuery{config: c.config}
}

// QueryStreets queries the streets edge of a City.
func (c *CityClient) QueryStreets(ci *City) *StreetQuery {
	query := &StreetQuery{config: c.config}
	id := ci.ID
	query.sql = sql.Select().From(sql.Table(street.Table)).
		Where(sql.EQ(city.StreetsColumn, id))

	return query
}

// StreetClient is a client for the Street schema.
type StreetClient struct {
	config
}

// NewStreetClient returns a client for the Street from the given config.
func NewStreetClient(c config) *StreetClient {
	return &StreetClient{config: c}
}

// Create returns a create builder for Street.
func (c *StreetClient) Create() *StreetCreate {
	return &StreetCreate{config: c.config}
}

// Update returns an update builder for Street.
func (c *StreetClient) Update() *StreetUpdate {
	return &StreetUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *StreetClient) UpdateOne(s *Street) *StreetUpdateOne {
	return c.UpdateOneID(s.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *StreetClient) UpdateOneID(id int) *StreetUpdateOne {
	return &StreetUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Street.
func (c *StreetClient) Delete() *StreetDelete {
	return &StreetDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *StreetClient) DeleteOne(s *Street) *StreetDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *StreetClient) DeleteOneID(id int) *StreetDeleteOne {
	return &StreetDeleteOne{c.Delete().Where(street.ID(id))}
}

// Create returns a query builder for Street.
func (c *StreetClient) Query() *StreetQuery {
	return &StreetQuery{config: c.config}
}

// QueryCity queries the city edge of a Street.
func (c *StreetClient) QueryCity(s *Street) *CityQuery {
	query := &CityQuery{config: c.config}
	id := s.ID
	t1 := sql.Table(city.Table)
	t2 := sql.Select(street.CityColumn).
		From(sql.Table(street.CityTable)).
		Where(sql.EQ(street.FieldID, id))
	query.sql = sql.Select().From(t1).Join(t2).On(t1.C(city.FieldID), t2.C(street.CityColumn))

	return query
}
