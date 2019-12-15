// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"log"

	"github.com/facebookincubator/ent/entc/integration/ent/migrate"

	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/comment"
	"github.com/facebookincubator/ent/entc/integration/ent/fieldtype"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/filetype"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/item"
	"github.com/facebookincubator/ent/entc/integration/ent/node"
	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/sql"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Card is the client for interacting with the Card builders.
	Card *CardClient
	// Comment is the client for interacting with the Comment builders.
	Comment *CommentClient
	// FieldType is the client for interacting with the FieldType builders.
	FieldType *FieldTypeClient
	// File is the client for interacting with the File builders.
	File *FileClient
	// FileType is the client for interacting with the FileType builders.
	FileType *FileTypeClient
	// Group is the client for interacting with the Group builders.
	Group *GroupClient
	// GroupInfo is the client for interacting with the GroupInfo builders.
	GroupInfo *GroupInfoClient
	// Item is the client for interacting with the Item builders.
	Item *ItemClient
	// Node is the client for interacting with the Node builders.
	Node *NodeClient
	// Pet is the client for interacting with the Pet builders.
	Pet *PetClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	c := config{log: log.Println}
	c.options(opts...)
	return &Client{
		config:    c,
		Schema:    migrate.NewSchema(c.driver),
		Card:      NewCardClient(c),
		Comment:   NewCommentClient(c),
		FieldType: NewFieldTypeClient(c),
		File:      NewFileClient(c),
		FileType:  NewFileTypeClient(c),
		Group:     NewGroupClient(c),
		GroupInfo: NewGroupInfoClient(c),
		Item:      NewItemClient(c),
		Node:      NewNodeClient(c),
		Pet:       NewPetClient(c),
		User:      NewUserClient(c),
	}
}

// Open opens a connection to the database specified by the driver name and a
// driver-specific data source name, and returns a new client attached to it.
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
		config:    cfg,
		Card:      NewCardClient(cfg),
		Comment:   NewCommentClient(cfg),
		FieldType: NewFieldTypeClient(cfg),
		File:      NewFileClient(cfg),
		FileType:  NewFileTypeClient(cfg),
		Group:     NewGroupClient(cfg),
		GroupInfo: NewGroupInfoClient(cfg),
		Item:      NewItemClient(cfg),
		Node:      NewNodeClient(cfg),
		Pet:       NewPetClient(cfg),
		User:      NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Card.
//		Query().
//		Count(ctx)
//
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := config{driver: dialect.Debug(c.driver, c.log), log: c.log, debug: true}
	return &Client{
		config:    cfg,
		Schema:    migrate.NewSchema(cfg.driver),
		Card:      NewCardClient(cfg),
		Comment:   NewCommentClient(cfg),
		FieldType: NewFieldTypeClient(cfg),
		File:      NewFileClient(cfg),
		FileType:  NewFileTypeClient(cfg),
		Group:     NewGroupClient(cfg),
		GroupInfo: NewGroupInfoClient(cfg),
		Item:      NewItemClient(cfg),
		Node:      NewNodeClient(cfg),
		Pet:       NewPetClient(cfg),
		User:      NewUserClient(cfg),
	}
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// CardClient is a client for the Card schema.
type CardClient struct {
	config
}

// NewCardClient returns a client for the Card from the given config.
func NewCardClient(c config) *CardClient {
	return &CardClient{config: c}
}

// Create returns a create builder for Card.
func (c *CardClient) Create() *CardCreate {
	return &CardCreate{config: c.config}
}

// Update returns an update builder for Card.
func (c *CardClient) Update() *CardUpdate {
	return &CardUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *CardClient) UpdateOne(ca *Card) *CardUpdateOne {
	return c.UpdateOneID(ca.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *CardClient) UpdateOneID(id string) *CardUpdateOne {
	return &CardUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Card.
func (c *CardClient) Delete() *CardDelete {
	return &CardDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CardClient) DeleteOne(ca *Card) *CardDeleteOne {
	return c.DeleteOneID(ca.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CardClient) DeleteOneID(id string) *CardDeleteOne {
	return &CardDeleteOne{c.Delete().Where(card.ID(id))}
}

// Create returns a query builder for Card.
func (c *CardClient) Query() *CardQuery {
	return &CardQuery{config: c.config}
}

// Get returns a Card entity by its id.
func (c *CardClient) Get(ctx context.Context, id string) (*Card, error) {
	return c.Query().Where(card.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CardClient) GetX(ctx context.Context, id string) *Card {
	ca, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ca
}

// QueryOwner queries the owner edge of a Card.
func (c *CardClient) QueryOwner(ca *Card) *UserQuery {
	query := &UserQuery{config: c.config}
	id := ca.id()
	step := sql.NewStep(
		sql.From(card.Table, card.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.O2O, true, card.OwnerTable, card.OwnerColumn),
	)
	query.sql = sql.Neighbors(ca.driver.Dialect(), step)

	return query
}

// CommentClient is a client for the Comment schema.
type CommentClient struct {
	config
}

// NewCommentClient returns a client for the Comment from the given config.
func NewCommentClient(c config) *CommentClient {
	return &CommentClient{config: c}
}

// Create returns a create builder for Comment.
func (c *CommentClient) Create() *CommentCreate {
	return &CommentCreate{config: c.config}
}

// Update returns an update builder for Comment.
func (c *CommentClient) Update() *CommentUpdate {
	return &CommentUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *CommentClient) UpdateOne(co *Comment) *CommentUpdateOne {
	return c.UpdateOneID(co.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *CommentClient) UpdateOneID(id string) *CommentUpdateOne {
	return &CommentUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Comment.
func (c *CommentClient) Delete() *CommentDelete {
	return &CommentDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *CommentClient) DeleteOne(co *Comment) *CommentDeleteOne {
	return c.DeleteOneID(co.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *CommentClient) DeleteOneID(id string) *CommentDeleteOne {
	return &CommentDeleteOne{c.Delete().Where(comment.ID(id))}
}

// Create returns a query builder for Comment.
func (c *CommentClient) Query() *CommentQuery {
	return &CommentQuery{config: c.config}
}

// Get returns a Comment entity by its id.
func (c *CommentClient) Get(ctx context.Context, id string) (*Comment, error) {
	return c.Query().Where(comment.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *CommentClient) GetX(ctx context.Context, id string) *Comment {
	co, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return co
}

// FieldTypeClient is a client for the FieldType schema.
type FieldTypeClient struct {
	config
}

// NewFieldTypeClient returns a client for the FieldType from the given config.
func NewFieldTypeClient(c config) *FieldTypeClient {
	return &FieldTypeClient{config: c}
}

// Create returns a create builder for FieldType.
func (c *FieldTypeClient) Create() *FieldTypeCreate {
	return &FieldTypeCreate{config: c.config}
}

// Update returns an update builder for FieldType.
func (c *FieldTypeClient) Update() *FieldTypeUpdate {
	return &FieldTypeUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *FieldTypeClient) UpdateOne(ft *FieldType) *FieldTypeUpdateOne {
	return c.UpdateOneID(ft.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *FieldTypeClient) UpdateOneID(id string) *FieldTypeUpdateOne {
	return &FieldTypeUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for FieldType.
func (c *FieldTypeClient) Delete() *FieldTypeDelete {
	return &FieldTypeDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FieldTypeClient) DeleteOne(ft *FieldType) *FieldTypeDeleteOne {
	return c.DeleteOneID(ft.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FieldTypeClient) DeleteOneID(id string) *FieldTypeDeleteOne {
	return &FieldTypeDeleteOne{c.Delete().Where(fieldtype.ID(id))}
}

// Create returns a query builder for FieldType.
func (c *FieldTypeClient) Query() *FieldTypeQuery {
	return &FieldTypeQuery{config: c.config}
}

// Get returns a FieldType entity by its id.
func (c *FieldTypeClient) Get(ctx context.Context, id string) (*FieldType, error) {
	return c.Query().Where(fieldtype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FieldTypeClient) GetX(ctx context.Context, id string) *FieldType {
	ft, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ft
}

// FileClient is a client for the File schema.
type FileClient struct {
	config
}

// NewFileClient returns a client for the File from the given config.
func NewFileClient(c config) *FileClient {
	return &FileClient{config: c}
}

// Create returns a create builder for File.
func (c *FileClient) Create() *FileCreate {
	return &FileCreate{config: c.config}
}

// Update returns an update builder for File.
func (c *FileClient) Update() *FileUpdate {
	return &FileUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileClient) UpdateOne(f *File) *FileUpdateOne {
	return c.UpdateOneID(f.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *FileClient) UpdateOneID(id string) *FileUpdateOne {
	return &FileUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for File.
func (c *FileClient) Delete() *FileDelete {
	return &FileDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FileClient) DeleteOne(f *File) *FileDeleteOne {
	return c.DeleteOneID(f.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FileClient) DeleteOneID(id string) *FileDeleteOne {
	return &FileDeleteOne{c.Delete().Where(file.ID(id))}
}

// Create returns a query builder for File.
func (c *FileClient) Query() *FileQuery {
	return &FileQuery{config: c.config}
}

// Get returns a File entity by its id.
func (c *FileClient) Get(ctx context.Context, id string) (*File, error) {
	return c.Query().Where(file.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileClient) GetX(ctx context.Context, id string) *File {
	f, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return f
}

// QueryOwner queries the owner edge of a File.
func (c *FileClient) QueryOwner(f *File) *UserQuery {
	query := &UserQuery{config: c.config}
	id := f.id()
	step := sql.NewStep(
		sql.From(file.Table, file.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2O, true, file.OwnerTable, file.OwnerColumn),
	)
	query.sql = sql.Neighbors(f.driver.Dialect(), step)

	return query
}

// QueryType queries the type edge of a File.
func (c *FileClient) QueryType(f *File) *FileTypeQuery {
	query := &FileTypeQuery{config: c.config}
	id := f.id()
	step := sql.NewStep(
		sql.From(file.Table, file.FieldID, id),
		sql.To(filetype.Table, filetype.FieldID),
		sql.Edge(sql.M2O, true, file.TypeTable, file.TypeColumn),
	)
	query.sql = sql.Neighbors(f.driver.Dialect(), step)

	return query
}

// FileTypeClient is a client for the FileType schema.
type FileTypeClient struct {
	config
}

// NewFileTypeClient returns a client for the FileType from the given config.
func NewFileTypeClient(c config) *FileTypeClient {
	return &FileTypeClient{config: c}
}

// Create returns a create builder for FileType.
func (c *FileTypeClient) Create() *FileTypeCreate {
	return &FileTypeCreate{config: c.config}
}

// Update returns an update builder for FileType.
func (c *FileTypeClient) Update() *FileTypeUpdate {
	return &FileTypeUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *FileTypeClient) UpdateOne(ft *FileType) *FileTypeUpdateOne {
	return c.UpdateOneID(ft.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *FileTypeClient) UpdateOneID(id string) *FileTypeUpdateOne {
	return &FileTypeUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for FileType.
func (c *FileTypeClient) Delete() *FileTypeDelete {
	return &FileTypeDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *FileTypeClient) DeleteOne(ft *FileType) *FileTypeDeleteOne {
	return c.DeleteOneID(ft.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *FileTypeClient) DeleteOneID(id string) *FileTypeDeleteOne {
	return &FileTypeDeleteOne{c.Delete().Where(filetype.ID(id))}
}

// Create returns a query builder for FileType.
func (c *FileTypeClient) Query() *FileTypeQuery {
	return &FileTypeQuery{config: c.config}
}

// Get returns a FileType entity by its id.
func (c *FileTypeClient) Get(ctx context.Context, id string) (*FileType, error) {
	return c.Query().Where(filetype.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *FileTypeClient) GetX(ctx context.Context, id string) *FileType {
	ft, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return ft
}

// QueryFiles queries the files edge of a FileType.
func (c *FileTypeClient) QueryFiles(ft *FileType) *FileQuery {
	query := &FileQuery{config: c.config}
	id := ft.id()
	step := sql.NewStep(
		sql.From(filetype.Table, filetype.FieldID, id),
		sql.To(file.Table, file.FieldID),
		sql.Edge(sql.O2M, false, filetype.FilesTable, filetype.FilesColumn),
	)
	query.sql = sql.Neighbors(ft.driver.Dialect(), step)

	return query
}

// GroupClient is a client for the Group schema.
type GroupClient struct {
	config
}

// NewGroupClient returns a client for the Group from the given config.
func NewGroupClient(c config) *GroupClient {
	return &GroupClient{config: c}
}

// Create returns a create builder for Group.
func (c *GroupClient) Create() *GroupCreate {
	return &GroupCreate{config: c.config}
}

// Update returns an update builder for Group.
func (c *GroupClient) Update() *GroupUpdate {
	return &GroupUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupClient) UpdateOne(gr *Group) *GroupUpdateOne {
	return c.UpdateOneID(gr.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupClient) UpdateOneID(id string) *GroupUpdateOne {
	return &GroupUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Group.
func (c *GroupClient) Delete() *GroupDelete {
	return &GroupDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupClient) DeleteOne(gr *Group) *GroupDeleteOne {
	return c.DeleteOneID(gr.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupClient) DeleteOneID(id string) *GroupDeleteOne {
	return &GroupDeleteOne{c.Delete().Where(group.ID(id))}
}

// Create returns a query builder for Group.
func (c *GroupClient) Query() *GroupQuery {
	return &GroupQuery{config: c.config}
}

// Get returns a Group entity by its id.
func (c *GroupClient) Get(ctx context.Context, id string) (*Group, error) {
	return c.Query().Where(group.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupClient) GetX(ctx context.Context, id string) *Group {
	gr, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gr
}

// QueryFiles queries the files edge of a Group.
func (c *GroupClient) QueryFiles(gr *Group) *FileQuery {
	query := &FileQuery{config: c.config}
	id := gr.id()
	step := sql.NewStep(
		sql.From(group.Table, group.FieldID, id),
		sql.To(file.Table, file.FieldID),
		sql.Edge(sql.O2M, false, group.FilesTable, group.FilesColumn),
	)
	query.sql = sql.Neighbors(gr.driver.Dialect(), step)

	return query
}

// QueryBlocked queries the blocked edge of a Group.
func (c *GroupClient) QueryBlocked(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	id := gr.id()
	step := sql.NewStep(
		sql.From(group.Table, group.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.O2M, false, group.BlockedTable, group.BlockedColumn),
	)
	query.sql = sql.Neighbors(gr.driver.Dialect(), step)

	return query
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	id := gr.id()
	step := sql.NewStep(
		sql.From(group.Table, group.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2M, true, group.UsersTable, group.UsersPrimaryKey...),
	)
	query.sql = sql.Neighbors(gr.driver.Dialect(), step)

	return query
}

// QueryInfo queries the info edge of a Group.
func (c *GroupClient) QueryInfo(gr *Group) *GroupInfoQuery {
	query := &GroupInfoQuery{config: c.config}
	id := gr.id()
	step := sql.NewStep(
		sql.From(group.Table, group.FieldID, id),
		sql.To(groupinfo.Table, groupinfo.FieldID),
		sql.Edge(sql.M2O, false, group.InfoTable, group.InfoColumn),
	)
	query.sql = sql.Neighbors(gr.driver.Dialect(), step)

	return query
}

// GroupInfoClient is a client for the GroupInfo schema.
type GroupInfoClient struct {
	config
}

// NewGroupInfoClient returns a client for the GroupInfo from the given config.
func NewGroupInfoClient(c config) *GroupInfoClient {
	return &GroupInfoClient{config: c}
}

// Create returns a create builder for GroupInfo.
func (c *GroupInfoClient) Create() *GroupInfoCreate {
	return &GroupInfoCreate{config: c.config}
}

// Update returns an update builder for GroupInfo.
func (c *GroupInfoClient) Update() *GroupInfoUpdate {
	return &GroupInfoUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *GroupInfoClient) UpdateOne(gi *GroupInfo) *GroupInfoUpdateOne {
	return c.UpdateOneID(gi.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *GroupInfoClient) UpdateOneID(id string) *GroupInfoUpdateOne {
	return &GroupInfoUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for GroupInfo.
func (c *GroupInfoClient) Delete() *GroupInfoDelete {
	return &GroupInfoDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *GroupInfoClient) DeleteOne(gi *GroupInfo) *GroupInfoDeleteOne {
	return c.DeleteOneID(gi.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *GroupInfoClient) DeleteOneID(id string) *GroupInfoDeleteOne {
	return &GroupInfoDeleteOne{c.Delete().Where(groupinfo.ID(id))}
}

// Create returns a query builder for GroupInfo.
func (c *GroupInfoClient) Query() *GroupInfoQuery {
	return &GroupInfoQuery{config: c.config}
}

// Get returns a GroupInfo entity by its id.
func (c *GroupInfoClient) Get(ctx context.Context, id string) (*GroupInfo, error) {
	return c.Query().Where(groupinfo.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *GroupInfoClient) GetX(ctx context.Context, id string) *GroupInfo {
	gi, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return gi
}

// QueryGroups queries the groups edge of a GroupInfo.
func (c *GroupInfoClient) QueryGroups(gi *GroupInfo) *GroupQuery {
	query := &GroupQuery{config: c.config}
	id := gi.id()
	step := sql.NewStep(
		sql.From(groupinfo.Table, groupinfo.FieldID, id),
		sql.To(group.Table, group.FieldID),
		sql.Edge(sql.O2M, true, groupinfo.GroupsTable, groupinfo.GroupsColumn),
	)
	query.sql = sql.Neighbors(gi.driver.Dialect(), step)

	return query
}

// ItemClient is a client for the Item schema.
type ItemClient struct {
	config
}

// NewItemClient returns a client for the Item from the given config.
func NewItemClient(c config) *ItemClient {
	return &ItemClient{config: c}
}

// Create returns a create builder for Item.
func (c *ItemClient) Create() *ItemCreate {
	return &ItemCreate{config: c.config}
}

// Update returns an update builder for Item.
func (c *ItemClient) Update() *ItemUpdate {
	return &ItemUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *ItemClient) UpdateOne(i *Item) *ItemUpdateOne {
	return c.UpdateOneID(i.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *ItemClient) UpdateOneID(id string) *ItemUpdateOne {
	return &ItemUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Item.
func (c *ItemClient) Delete() *ItemDelete {
	return &ItemDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *ItemClient) DeleteOne(i *Item) *ItemDeleteOne {
	return c.DeleteOneID(i.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *ItemClient) DeleteOneID(id string) *ItemDeleteOne {
	return &ItemDeleteOne{c.Delete().Where(item.ID(id))}
}

// Create returns a query builder for Item.
func (c *ItemClient) Query() *ItemQuery {
	return &ItemQuery{config: c.config}
}

// Get returns a Item entity by its id.
func (c *ItemClient) Get(ctx context.Context, id string) (*Item, error) {
	return c.Query().Where(item.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *ItemClient) GetX(ctx context.Context, id string) *Item {
	i, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return i
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
func (c *NodeClient) UpdateOneID(id string) *NodeUpdateOne {
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
func (c *NodeClient) DeleteOneID(id string) *NodeDeleteOne {
	return &NodeDeleteOne{c.Delete().Where(node.ID(id))}
}

// Create returns a query builder for Node.
func (c *NodeClient) Query() *NodeQuery {
	return &NodeQuery{config: c.config}
}

// Get returns a Node entity by its id.
func (c *NodeClient) Get(ctx context.Context, id string) (*Node, error) {
	return c.Query().Where(node.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *NodeClient) GetX(ctx context.Context, id string) *Node {
	n, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return n
}

// QueryPrev queries the prev edge of a Node.
func (c *NodeClient) QueryPrev(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	id := n.id()
	step := sql.NewStep(
		sql.From(node.Table, node.FieldID, id),
		sql.To(node.Table, node.FieldID),
		sql.Edge(sql.O2O, true, node.PrevTable, node.PrevColumn),
	)
	query.sql = sql.Neighbors(n.driver.Dialect(), step)

	return query
}

// QueryNext queries the next edge of a Node.
func (c *NodeClient) QueryNext(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	id := n.id()
	step := sql.NewStep(
		sql.From(node.Table, node.FieldID, id),
		sql.To(node.Table, node.FieldID),
		sql.Edge(sql.O2O, false, node.NextTable, node.NextColumn),
	)
	query.sql = sql.Neighbors(n.driver.Dialect(), step)

	return query
}

// PetClient is a client for the Pet schema.
type PetClient struct {
	config
}

// NewPetClient returns a client for the Pet from the given config.
func NewPetClient(c config) *PetClient {
	return &PetClient{config: c}
}

// Create returns a create builder for Pet.
func (c *PetClient) Create() *PetCreate {
	return &PetCreate{config: c.config}
}

// Update returns an update builder for Pet.
func (c *PetClient) Update() *PetUpdate {
	return &PetUpdate{config: c.config}
}

// UpdateOne returns an update builder for the given entity.
func (c *PetClient) UpdateOne(pe *Pet) *PetUpdateOne {
	return c.UpdateOneID(pe.ID)
}

// UpdateOneID returns an update builder for the given id.
func (c *PetClient) UpdateOneID(id string) *PetUpdateOne {
	return &PetUpdateOne{config: c.config, id: id}
}

// Delete returns a delete builder for Pet.
func (c *PetClient) Delete() *PetDelete {
	return &PetDelete{config: c.config}
}

// DeleteOne returns a delete builder for the given entity.
func (c *PetClient) DeleteOne(pe *Pet) *PetDeleteOne {
	return c.DeleteOneID(pe.ID)
}

// DeleteOneID returns a delete builder for the given id.
func (c *PetClient) DeleteOneID(id string) *PetDeleteOne {
	return &PetDeleteOne{c.Delete().Where(pet.ID(id))}
}

// Create returns a query builder for Pet.
func (c *PetClient) Query() *PetQuery {
	return &PetQuery{config: c.config}
}

// Get returns a Pet entity by its id.
func (c *PetClient) Get(ctx context.Context, id string) (*Pet, error) {
	return c.Query().Where(pet.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *PetClient) GetX(ctx context.Context, id string) *Pet {
	pe, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return pe
}

// QueryTeam queries the team edge of a Pet.
func (c *PetClient) QueryTeam(pe *Pet) *UserQuery {
	query := &UserQuery{config: c.config}
	id := pe.id()
	step := sql.NewStep(
		sql.From(pet.Table, pet.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.O2O, true, pet.TeamTable, pet.TeamColumn),
	)
	query.sql = sql.Neighbors(pe.driver.Dialect(), step)

	return query
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *UserQuery {
	query := &UserQuery{config: c.config}
	id := pe.id()
	step := sql.NewStep(
		sql.From(pet.Table, pet.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2O, true, pet.OwnerTable, pet.OwnerColumn),
	)
	query.sql = sql.Neighbors(pe.driver.Dialect(), step)

	return query
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
func (c *UserClient) UpdateOneID(id string) *UserUpdateOne {
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
func (c *UserClient) DeleteOneID(id string) *UserDeleteOne {
	return &UserDeleteOne{c.Delete().Where(user.ID(id))}
}

// Create returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{config: c.config}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id string) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id string) *User {
	u, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return u
}

// QueryCard queries the card edge of a User.
func (c *UserClient) QueryCard(u *User) *CardQuery {
	query := &CardQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(card.Table, card.FieldID),
		sql.Edge(sql.O2O, false, user.CardTable, user.CardColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryPets queries the pets edge of a User.
func (c *UserClient) QueryPets(u *User) *PetQuery {
	query := &PetQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(pet.Table, pet.FieldID),
		sql.Edge(sql.O2M, false, user.PetsTable, user.PetsColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryFiles queries the files edge of a User.
func (c *UserClient) QueryFiles(u *User) *FileQuery {
	query := &FileQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(file.Table, file.FieldID),
		sql.Edge(sql.O2M, false, user.FilesTable, user.FilesColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(group.Table, group.FieldID),
		sql.Edge(sql.M2M, false, user.GroupsTable, user.GroupsPrimaryKey...),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryFriends queries the friends edge of a User.
func (c *UserClient) QueryFriends(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2M, false, user.FriendsTable, user.FriendsPrimaryKey...),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2M, true, user.FollowersTable, user.FollowersPrimaryKey...),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryFollowing queries the following edge of a User.
func (c *UserClient) QueryFollowing(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2M, false, user.FollowingTable, user.FollowingPrimaryKey...),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryTeam queries the team edge of a User.
func (c *UserClient) QueryTeam(u *User) *PetQuery {
	query := &PetQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(pet.Table, pet.FieldID),
		sql.Edge(sql.O2O, false, user.TeamTable, user.TeamColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QuerySpouse queries the spouse edge of a User.
func (c *UserClient) QuerySpouse(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.O2O, false, user.SpouseTable, user.SpouseColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.O2M, true, user.ChildrenTable, user.ChildrenColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	id := u.id()
	step := sql.NewStep(
		sql.From(user.Table, user.FieldID, id),
		sql.To(user.Table, user.FieldID),
		sql.Edge(sql.M2O, false, user.ParentTable, user.ParentColumn),
	)
	query.sql = sql.Neighbors(u.driver.Dialect(), step)

	return query
}
