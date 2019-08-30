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
	"github.com/facebookincubator/ent/entc/integration/ent/node"
	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/user"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
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
		Node:      NewNodeClient(c),
		Pet:       NewPetClient(c),
		User:      NewUserClient(c),
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
		Node:      NewNodeClient(cfg),
		Pet:       NewPetClient(cfg),
		User:      NewUserClient(cfg),
	}
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

// QueryOwner queries the owner edge of a Card.
func (c *CardClient) QueryOwner(ca *Card) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := ca.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Select(card.OwnerColumn).
			From(sql.Table(card.OwnerTable)).
			Where(sql.EQ(card.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(user.FieldID), t2.C(card.OwnerColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(ca.ID).InE(user.CardLabel).OutV()

	}
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

// QueryOwner queries the owner edge of a File.
func (c *FileClient) QueryOwner(f *File) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := f.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Select(file.OwnerColumn).
			From(sql.Table(file.OwnerTable)).
			Where(sql.EQ(file.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(user.FieldID), t2.C(file.OwnerColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(f.ID).InE(user.FilesLabel).OutV()

	}
	return query
}

// QueryType queries the type edge of a File.
func (c *FileClient) QueryType(f *File) *FileTypeQuery {
	query := &FileTypeQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := f.id()
		t1 := sql.Table(filetype.Table)
		t2 := sql.Select(file.TypeColumn).
			From(sql.Table(file.TypeTable)).
			Where(sql.EQ(file.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(filetype.FieldID), t2.C(file.TypeColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(f.ID).InE(filetype.FilesLabel).OutV()

	}
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

// QueryFiles queries the files edge of a FileType.
func (c *FileTypeClient) QueryFiles(ft *FileType) *FileQuery {
	query := &FileQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := ft.id()
		query.sql = sql.Select().From(sql.Table(file.Table)).
			Where(sql.EQ(filetype.FilesColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(ft.ID).OutE(filetype.FilesLabel).InV()

	}
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

// QueryFiles queries the files edge of a Group.
func (c *GroupClient) QueryFiles(gr *Group) *FileQuery {
	query := &FileQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := gr.id()
		query.sql = sql.Select().From(sql.Table(file.Table)).
			Where(sql.EQ(group.FilesColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(gr.ID).OutE(group.FilesLabel).InV()

	}
	return query
}

// QueryBlocked queries the blocked edge of a Group.
func (c *GroupClient) QueryBlocked(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := gr.id()
		query.sql = sql.Select().From(sql.Table(user.Table)).
			Where(sql.EQ(group.BlockedColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(gr.ID).OutE(group.BlockedLabel).InV()

	}
	return query
}

// QueryUsers queries the users edge of a Group.
func (c *GroupClient) QueryUsers(gr *Group) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := gr.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Table(group.Table)
		t3 := sql.Table(group.UsersTable)
		t4 := sql.Select(t3.C(group.UsersPrimaryKey[0])).
			From(t3).
			Join(t2).
			On(t3.C(group.UsersPrimaryKey[1]), t2.C(group.FieldID)).
			Where(sql.EQ(t2.C(group.FieldID), id))
		query.sql = sql.Select().
			From(t1).
			Join(t4).
			On(t1.C(user.FieldID), t4.C(group.UsersPrimaryKey[0]))

	case dialect.Gremlin:
		query.gremlin = g.V(gr.ID).InE(user.GroupsLabel).OutV()

	}
	return query
}

// QueryInfo queries the info edge of a Group.
func (c *GroupClient) QueryInfo(gr *Group) *GroupInfoQuery {
	query := &GroupInfoQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := gr.id()
		t1 := sql.Table(groupinfo.Table)
		t2 := sql.Select(group.InfoColumn).
			From(sql.Table(group.InfoTable)).
			Where(sql.EQ(group.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(groupinfo.FieldID), t2.C(group.InfoColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(gr.ID).OutE(group.InfoLabel).InV()

	}
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

// QueryGroups queries the groups edge of a GroupInfo.
func (c *GroupInfoClient) QueryGroups(gi *GroupInfo) *GroupQuery {
	query := &GroupQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := gi.id()
		query.sql = sql.Select().From(sql.Table(group.Table)).
			Where(sql.EQ(groupinfo.GroupsColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(gi.ID).InE(group.InfoLabel).OutV()

	}
	return query
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

// QueryPrev queries the prev edge of a Node.
func (c *NodeClient) QueryPrev(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := n.id()
		t1 := sql.Table(node.Table)
		t2 := sql.Select(node.PrevColumn).
			From(sql.Table(node.PrevTable)).
			Where(sql.EQ(node.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(node.FieldID), t2.C(node.PrevColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(n.ID).InE(node.NextLabel).OutV()

	}
	return query
}

// QueryNext queries the next edge of a Node.
func (c *NodeClient) QueryNext(n *Node) *NodeQuery {
	query := &NodeQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := n.id()
		query.sql = sql.Select().From(sql.Table(node.Table)).
			Where(sql.EQ(node.NextColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(n.ID).OutE(node.NextLabel).InV()

	}
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

// QueryTeam queries the team edge of a Pet.
func (c *PetClient) QueryTeam(pe *Pet) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := pe.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Select(pet.TeamColumn).
			From(sql.Table(pet.TeamTable)).
			Where(sql.EQ(pet.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(user.FieldID), t2.C(pet.TeamColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(pe.ID).InE(user.TeamLabel).OutV()

	}
	return query
}

// QueryOwner queries the owner edge of a Pet.
func (c *PetClient) QueryOwner(pe *Pet) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := pe.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Select(pet.OwnerColumn).
			From(sql.Table(pet.OwnerTable)).
			Where(sql.EQ(pet.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(user.FieldID), t2.C(pet.OwnerColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(pe.ID).InE(user.PetsLabel).OutV()

	}
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

// QueryCard queries the card edge of a User.
func (c *UserClient) QueryCard(u *User) *CardQuery {
	query := &CardQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(card.Table)).
			Where(sql.EQ(user.CardColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.CardLabel).InV()

	}
	return query
}

// QueryPets queries the pets edge of a User.
func (c *UserClient) QueryPets(u *User) *PetQuery {
	query := &PetQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(pet.Table)).
			Where(sql.EQ(user.PetsColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.PetsLabel).InV()

	}
	return query
}

// QueryFiles queries the files edge of a User.
func (c *UserClient) QueryFiles(u *User) *FileQuery {
	query := &FileQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(file.Table)).
			Where(sql.EQ(user.FilesColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.FilesLabel).InV()

	}
	return query
}

// QueryGroups queries the groups edge of a User.
func (c *UserClient) QueryGroups(u *User) *GroupQuery {
	query := &GroupQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		t1 := sql.Table(group.Table)
		t2 := sql.Table(user.Table)
		t3 := sql.Table(user.GroupsTable)
		t4 := sql.Select(t3.C(user.GroupsPrimaryKey[1])).
			From(t3).
			Join(t2).
			On(t3.C(user.GroupsPrimaryKey[0]), t2.C(user.FieldID)).
			Where(sql.EQ(t2.C(user.FieldID), id))
		query.sql = sql.Select().
			From(t1).
			Join(t4).
			On(t1.C(group.FieldID), t4.C(user.GroupsPrimaryKey[1]))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.GroupsLabel).InV()

	}
	return query
}

// QueryFriends queries the friends edge of a User.
func (c *UserClient) QueryFriends(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Table(user.Table)
		t3 := sql.Table(user.FriendsTable)
		t4 := sql.Select(t3.C(user.FriendsPrimaryKey[1])).
			From(t3).
			Join(t2).
			On(t3.C(user.FriendsPrimaryKey[0]), t2.C(user.FieldID)).
			Where(sql.EQ(t2.C(user.FieldID), id))
		query.sql = sql.Select().
			From(t1).
			Join(t4).
			On(t1.C(user.FieldID), t4.C(user.FriendsPrimaryKey[1]))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).Both(user.FriendsLabel)

	}
	return query
}

// QueryFollowers queries the followers edge of a User.
func (c *UserClient) QueryFollowers(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
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

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).InE(user.FollowingLabel).OutV()

	}
	return query
}

// QueryFollowing queries the following edge of a User.
func (c *UserClient) QueryFollowing(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
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

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.FollowingLabel).InV()

	}
	return query
}

// QueryTeam queries the team edge of a User.
func (c *UserClient) QueryTeam(u *User) *PetQuery {
	query := &PetQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(pet.Table)).
			Where(sql.EQ(user.TeamColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.TeamLabel).InV()

	}
	return query
}

// QuerySpouse queries the spouse edge of a User.
func (c *UserClient) QuerySpouse(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(user.Table)).
			Where(sql.EQ(user.SpouseColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).Both(user.SpouseLabel)

	}
	return query
}

// QueryChildren queries the children edge of a User.
func (c *UserClient) QueryChildren(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		query.sql = sql.Select().From(sql.Table(user.Table)).
			Where(sql.EQ(user.ChildrenColumn, id))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).InE(user.ParentLabel).OutV()

	}
	return query
}

// QueryParent queries the parent edge of a User.
func (c *UserClient) QueryParent(u *User) *UserQuery {
	query := &UserQuery{config: c.config}
	switch c.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		id := u.id()
		t1 := sql.Table(user.Table)
		t2 := sql.Select(user.ParentColumn).
			From(sql.Table(user.ParentTable)).
			Where(sql.EQ(user.FieldID, id))
		query.sql = sql.Select().From(t1).Join(t2).On(t1.C(user.FieldID), t2.C(user.ParentColumn))

	case dialect.Gremlin:
		query.gremlin = g.V(u.ID).OutE(user.ParentLabel).InV()

	}
	return query
}
