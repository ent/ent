// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/entc/integration/template/ent/group"
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"
)

// Noder wraps the basic Node method.
type Noder interface {
	Node(context.Context) (*Node, error)
}

// Node in the graph.
type Node struct {
	ID     int      `json:"id,omitemty"`      // node id.
	Type   string   `json:"type,omitempty"`   // node type.
	Fields []*Field `json:"fields,omitempty"` // node fields.
	Edges  []*Edge  `json:"edges,omitempty"`  // node edges.
}

// Field of a node.
type Field struct {
	Type  string `json:"type,omitempty"`  // field type.
	Name  string `json:"name,omitempty"`  // field name (as in struct).
	Value string `json:"value,omitempty"` // stringified value.
}

// Edges between two nodes.
type Edge struct {
	Type string `json:"type,omitempty"` // edge type.
	Name string `json:"name,omitempty"` // edge name.
	IDs  []int  `json:"ids,omitempty"`  // node ids (where this edge point to).
}

func (gr *Group) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     gr.ID,
		Type:   "Group",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 0),
	}
	var buf []byte
	if buf, err = json.Marshal(gr.MaxUsers); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "MaxUsers",
		Value: string(buf),
	}
	return node, nil
}

func (pe *Pet) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     pe.ID,
		Type:   "Pet",
		Fields: make([]*Field, 2),
		Edges:  make([]*Edge, 1),
	}
	var buf []byte
	if buf, err = json.Marshal(pe.Age); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "int",
		Name:  "Age",
		Value: string(buf),
	}
	if buf, err = json.Marshal(pe.LicensedAt); err != nil {
		return nil, err
	}
	node.Fields[1] = &Field{
		Type:  "time.Time",
		Name:  "LicensedAt",
		Value: string(buf),
	}
	var ids []int
	ids, err = pe.QueryOwner().
		Select(user.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[0] = &Edge{
		IDs:  ids,
		Type: "User",
		Name: "Owner",
	}
	return node, nil
}

func (u *User) Node(ctx context.Context) (node *Node, err error) {
	node = &Node{
		ID:     u.ID,
		Type:   "User",
		Fields: make([]*Field, 1),
		Edges:  make([]*Edge, 2),
	}
	var buf []byte
	if buf, err = json.Marshal(u.Name); err != nil {
		return nil, err
	}
	node.Fields[0] = &Field{
		Type:  "string",
		Name:  "Name",
		Value: string(buf),
	}
	var ids []int
	ids, err = u.QueryPets().
		Select(pet.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[0] = &Edge{
		IDs:  ids,
		Type: "Pet",
		Name: "Pets",
	}
	ids, err = u.QueryFriends().
		Select(user.FieldID).
		Ints(ctx)
	if err != nil {
		return nil, err
	}
	node.Edges[1] = &Edge{
		IDs:  ids,
		Type: "User",
		Name: "Friends",
	}
	return node, nil
}

var (
	once      sync.Once
	types     []string
	typeNodes = make(map[string]func(context.Context, int) (*Node, error))
)

func (c *Client) Node(ctx context.Context, id int) (*Node, error) {
	var err error
	once.Do(func() {
		err = c.loadTypes(ctx)
	})
	if err != nil {
		return nil, err
	}
	idx := id / (1<<32 - 1)
	return typeNodes[types[idx]](ctx, id)
}

func (c *Client) loadTypes(ctx context.Context) error {
	rows := &sql.Rows{}
	query, args := sql.Select("type").
		From(sql.Table(schema.TypeTable)).
		OrderBy(sql.Asc("id")).
		Query()
	if err := c.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	if err := sql.ScanSlice(rows, &types); err != nil {
		return err
	}
	typeNodes[group.Table] = func(ctx context.Context, id int) (*Node, error) {
		nv, err := c.Group.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return nv.Node(ctx)
	}
	typeNodes[pet.Table] = func(ctx context.Context, id int) (*Node, error) {
		nv, err := c.Pet.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return nv.Node(ctx)
	}
	typeNodes[user.Table] = func(ctx context.Context, id int) (*Node, error) {
		nv, err := c.User.Get(ctx, id)
		if err != nil {
			return nil, err
		}
		return nv.Node(ctx)
	}
	return nil
}
