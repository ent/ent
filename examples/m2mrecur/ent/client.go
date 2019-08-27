// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/examples/m2mrecur/ent/migrate"

	"github.com/facebookincubator/ent/examples/m2mrecur/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config: c,
		Schema: migrate.NewSchema(c.driver),
		User:   NewUserClient(c),
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
		User:   NewUserClient(cfg),
	}, nil
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Create returns a create builder for User.
func (c *UserClient) Create() *UserCreate {
	return &UserCreate{config: c.config}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	return &UserUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	return c.UpdateOneID(u.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	return &UserUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	return &UserDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	return &UserDeleteOne{c.Delete().Where(user.ID(id))}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.ID
	t1 := sql.Table(user.Table)
	t2 := sql.Table(user.Table)
	t3 := sql.Table(user.FollowersTable)
	t4 := sql.Select(t3.C(user.FollowersPrimaryKey[0])).
		From(t3).
		Join(t2).
		On(t3.C(user.FollowersPrimaryKey[1]), t2.C(user.FieldID)).
		Where(sql.EQ(t2.C(user.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(user.FieldID), t4.C(user.FollowersPrimaryKey[0]))

	return query
}

// QueryFollowing queries the following edge of a User.
func (c *UserClient) QueryFollowing(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.ID
	t1 := sql.Table(user.Table)
	t2 := sql.Table(user.Table)
	t3 := sql.Table(user.FollowingTable)
	t4 := sql.Select(t3.C(user.FollowingPrimaryKey[1])).
		From(t3).
		Join(t2).
		On(t3.C(user.FollowingPrimaryKey[0]), t2.C(user.FieldID)).
		Where(sql.EQ(t2.C(user.FieldID), id))
	query.sql = sql.Select().
		From(t1).
		Join(t4).
		On(t1.C(user.FieldID), t4.C(user.FollowingPrimaryKey[1]))

	return query
}
