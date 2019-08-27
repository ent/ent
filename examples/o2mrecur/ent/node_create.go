// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"

	"github.com/facebookincubator/ent/dialect/sql"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	value    *int
	parent   map[int]struct{}
	children map[int]struct{}
}

// SetValue sets the value field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.value = &i
	return nc
}

// SetParentID sets the parent edge to Node by id.
func (nc *NodeCreate) SetParentID(id int) *NodeCreate {
	if nc.parent == nil {
		nc.parent = make(map[int]struct{})
	}
	nc.parent[id] = struct{}{}
	return nc
}

// SetNillableParentID sets the parent edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillableParentID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetParentID(*id)
	}
	return nc
}

// SetParent sets the parent edge to Node.
func (nc *NodeCreate) SetParent(n *Node) *NodeCreate {
	return nc.SetParentID(n.ID)
}

// AddChildIDs adds the children edge to Node by ids.
func (nc *NodeCreate) AddChildIDs(ids ...int) *NodeCreate {
	if nc.children == nil {
		nc.children = make(map[int]struct{})
	}
	for i := range ids {
		nc.children[ids[i]] = struct{}{}
	}
	return nc
}

// AddChildren adds the children edges to Node.
func (nc *NodeCreate) AddChildren(n ...*Node) *NodeCreate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nc.AddChildIDs(ids...)
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	if nc.value == nil {
		return nil, errors.New("ent: missing required field \"value\"")
	}
	if len(nc.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return nc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (nc *NodeCreate) SaveX(ctx context.Context) *Node {
	v, err := nc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (nc *NodeCreate) sqlSave(ctx context.Context) (*Node, error) {
	var (
		res sql.Result
		n   = &Node{config: nc.config}
	)
	tx, err := nc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(node.Table).Default(nc.driver.Dialect())
	if nc.value != nil {
		builder.Set(node.FieldValue, *nc.value)
		n.Value = *nc.value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	n.ID = int(id)
	if len(nc.parent) > 0 {
		for eid := range nc.parent {
			query, args := sql.Update(node.ParentTable).
				Set(node.ParentColumn, eid).
				Where(sql.EQ(node.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(nc.children) > 0 {
		p := sql.P()
		for eid := range nc.children {
			p.Or().EQ(node.FieldID, eid)
		}
		query, args := sql.Update(node.ChildrenTable).
			Set(node.ChildrenColumn, id).
			Where(sql.And(p, sql.IsNull(node.ChildrenColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(nc.children) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"children\" %v already connected to a different \"Node\"", keys(nc.children))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return n, nil
}
