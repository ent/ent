// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/o2orecur/ent/node"

	"github.com/facebookincubator/ent/dialect/sql"
)

// NodeCreate is the builder for creating a Node entity.
type NodeCreate struct {
	config
	value *int
	prev  map[int]struct{}
	next  map[int]struct{}
}

// SetValue sets the value field.
func (nc *NodeCreate) SetValue(i int) *NodeCreate {
	nc.value = &i
	return nc
}

// SetPrevID sets the prev edge to Node by id.
func (nc *NodeCreate) SetPrevID(id int) *NodeCreate {
	if nc.prev == nil {
		nc.prev = make(map[int]struct{})
	}
	nc.prev[id] = struct{}{}
	return nc
}

// SetNillablePrevID sets the prev edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillablePrevID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetPrevID(*id)
	}
	return nc
}

// SetPrev sets the prev edge to Node.
func (nc *NodeCreate) SetPrev(n *Node) *NodeCreate {
	return nc.SetPrevID(n.ID)
}

// SetNextID sets the next edge to Node by id.
func (nc *NodeCreate) SetNextID(id int) *NodeCreate {
	if nc.next == nil {
		nc.next = make(map[int]struct{})
	}
	nc.next[id] = struct{}{}
	return nc
}

// SetNillableNextID sets the next edge to Node by id if the given value is not nil.
func (nc *NodeCreate) SetNillableNextID(id *int) *NodeCreate {
	if id != nil {
		nc = nc.SetNextID(*id)
	}
	return nc
}

// SetNext sets the next edge to Node.
func (nc *NodeCreate) SetNext(n *Node) *NodeCreate {
	return nc.SetNextID(n.ID)
}

// Save creates the Node in the database.
func (nc *NodeCreate) Save(ctx context.Context) (*Node, error) {
	if nc.value == nil {
		return nil, errors.New("ent: missing required field \"value\"")
	}
	if len(nc.prev) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"prev\"")
	}
	if len(nc.next) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"next\"")
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
	if len(nc.prev) > 0 {
		eid := keys(nc.prev)[0]
		query, args := sql.Update(node.PrevTable).
			Set(node.PrevColumn, eid).
			Where(sql.EQ(node.FieldID, id).And().IsNull(node.PrevColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(nc.prev) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"prev\" %v already connected to a different \"Node\"", keys(nc.prev))})
		}
	}
	if len(nc.next) > 0 {
		eid := keys(nc.next)[0]
		query, args := sql.Update(node.NextTable).
			Set(node.NextColumn, id).
			Where(sql.EQ(node.FieldID, eid).And().IsNull(node.NextColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(nc.next) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"next\" %v already connected to a different \"Node\"", keys(nc.next))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return n, nil
}
