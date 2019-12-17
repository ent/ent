// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/o2mrecur/ent/node"
	"github.com/facebookincubator/ent/examples/o2mrecur/ent/predicate"
)

// NodeUpdate is the builder for updating Node entities.
type NodeUpdate struct {
	config
	value           *int
	addvalue        *int
	parent          map[int]struct{}
	children        map[int]struct{}
	clearedParent   bool
	removedChildren map[int]struct{}
	predicates      []predicate.Node
}

// Where adds a new predicate for the builder.
func (nu *NodeUpdate) Where(ps ...predicate.Node) *NodeUpdate {
	nu.predicates = append(nu.predicates, ps...)
	return nu
}

// SetValue sets the value field.
func (nu *NodeUpdate) SetValue(i int) *NodeUpdate {
	nu.value = &i
	nu.addvalue = nil
	return nu
}

// AddValue adds i to value.
func (nu *NodeUpdate) AddValue(i int) *NodeUpdate {
	if nu.addvalue == nil {
		nu.addvalue = &i
	} else {
		*nu.addvalue += i
	}
	return nu
}

// SetParentID sets the parent edge to Node by id.
func (nu *NodeUpdate) SetParentID(id int) *NodeUpdate {
	if nu.parent == nil {
		nu.parent = make(map[int]struct{})
	}
	nu.parent[id] = struct{}{}
	return nu
}

// SetNillableParentID sets the parent edge to Node by id if the given value is not nil.
func (nu *NodeUpdate) SetNillableParentID(id *int) *NodeUpdate {
	if id != nil {
		nu = nu.SetParentID(*id)
	}
	return nu
}

// SetParent sets the parent edge to Node.
func (nu *NodeUpdate) SetParent(n *Node) *NodeUpdate {
	return nu.SetParentID(n.ID)
}

// AddChildIDs adds the children edge to Node by ids.
func (nu *NodeUpdate) AddChildIDs(ids ...int) *NodeUpdate {
	if nu.children == nil {
		nu.children = make(map[int]struct{})
	}
	for i := range ids {
		nu.children[ids[i]] = struct{}{}
	}
	return nu
}

// AddChildren adds the children edges to Node.
func (nu *NodeUpdate) AddChildren(n ...*Node) *NodeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nu.AddChildIDs(ids...)
}

// ClearParent clears the parent edge to Node.
func (nu *NodeUpdate) ClearParent() *NodeUpdate {
	nu.clearedParent = true
	return nu
}

// RemoveChildIDs removes the children edge to Node by ids.
func (nu *NodeUpdate) RemoveChildIDs(ids ...int) *NodeUpdate {
	if nu.removedChildren == nil {
		nu.removedChildren = make(map[int]struct{})
	}
	for i := range ids {
		nu.removedChildren[ids[i]] = struct{}{}
	}
	return nu
}

// RemoveChildren removes children edges to Node.
func (nu *NodeUpdate) RemoveChildren(n ...*Node) *NodeUpdate {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nu.RemoveChildIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (nu *NodeUpdate) Save(ctx context.Context) (int, error) {
	if len(nu.parent) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return nu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NodeUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NodeUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NodeUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NodeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(nu.driver.Dialect())
		selector = builder.Select(node.FieldID).From(builder.Table(node.Table))
	)
	for _, p := range nu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = nu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := nu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(node.Table)
	)
	updater = updater.Where(sql.InInts(node.FieldID, ids...))
	if value := nu.value; value != nil {
		updater.Set(node.FieldValue, *value)
	}
	if value := nu.addvalue; value != nil {
		updater.Add(node.FieldValue, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if nu.clearedParent {
		query, args := builder.Update(node.ParentTable).
			SetNull(node.ParentColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(nu.parent) > 0 {
		for eid := range nu.parent {
			query, args := builder.Update(node.ParentTable).
				Set(node.ParentColumn, eid).
				Where(sql.InInts(node.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if len(nu.removedChildren) > 0 {
		eids := make([]int, len(nu.removedChildren))
		for eid := range nu.removedChildren {
			eids = append(eids, eid)
		}
		query, args := builder.Update(node.ChildrenTable).
			SetNull(node.ChildrenColumn).
			Where(sql.InInts(node.ChildrenColumn, ids...)).
			Where(sql.InInts(node.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(nu.children) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range nu.children {
				p.Or().EQ(node.FieldID, eid)
			}
			query, args := builder.Update(node.ChildrenTable).
				Set(node.ChildrenColumn, id).
				Where(sql.And(p, sql.IsNull(node.ChildrenColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(nu.children) {
				return 0, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"children\" %v already connected to a different \"Node\"", keys(nu.children))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// NodeUpdateOne is the builder for updating a single Node entity.
type NodeUpdateOne struct {
	config
	id              int
	value           *int
	addvalue        *int
	parent          map[int]struct{}
	children        map[int]struct{}
	clearedParent   bool
	removedChildren map[int]struct{}
}

// SetValue sets the value field.
func (nuo *NodeUpdateOne) SetValue(i int) *NodeUpdateOne {
	nuo.value = &i
	nuo.addvalue = nil
	return nuo
}

// AddValue adds i to value.
func (nuo *NodeUpdateOne) AddValue(i int) *NodeUpdateOne {
	if nuo.addvalue == nil {
		nuo.addvalue = &i
	} else {
		*nuo.addvalue += i
	}
	return nuo
}

// SetParentID sets the parent edge to Node by id.
func (nuo *NodeUpdateOne) SetParentID(id int) *NodeUpdateOne {
	if nuo.parent == nil {
		nuo.parent = make(map[int]struct{})
	}
	nuo.parent[id] = struct{}{}
	return nuo
}

// SetNillableParentID sets the parent edge to Node by id if the given value is not nil.
func (nuo *NodeUpdateOne) SetNillableParentID(id *int) *NodeUpdateOne {
	if id != nil {
		nuo = nuo.SetParentID(*id)
	}
	return nuo
}

// SetParent sets the parent edge to Node.
func (nuo *NodeUpdateOne) SetParent(n *Node) *NodeUpdateOne {
	return nuo.SetParentID(n.ID)
}

// AddChildIDs adds the children edge to Node by ids.
func (nuo *NodeUpdateOne) AddChildIDs(ids ...int) *NodeUpdateOne {
	if nuo.children == nil {
		nuo.children = make(map[int]struct{})
	}
	for i := range ids {
		nuo.children[ids[i]] = struct{}{}
	}
	return nuo
}

// AddChildren adds the children edges to Node.
func (nuo *NodeUpdateOne) AddChildren(n ...*Node) *NodeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nuo.AddChildIDs(ids...)
}

// ClearParent clears the parent edge to Node.
func (nuo *NodeUpdateOne) ClearParent() *NodeUpdateOne {
	nuo.clearedParent = true
	return nuo
}

// RemoveChildIDs removes the children edge to Node by ids.
func (nuo *NodeUpdateOne) RemoveChildIDs(ids ...int) *NodeUpdateOne {
	if nuo.removedChildren == nil {
		nuo.removedChildren = make(map[int]struct{})
	}
	for i := range ids {
		nuo.removedChildren[ids[i]] = struct{}{}
	}
	return nuo
}

// RemoveChildren removes children edges to Node.
func (nuo *NodeUpdateOne) RemoveChildren(n ...*Node) *NodeUpdateOne {
	ids := make([]int, len(n))
	for i := range n {
		ids[i] = n[i].ID
	}
	return nuo.RemoveChildIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (nuo *NodeUpdateOne) Save(ctx context.Context) (*Node, error) {
	if len(nuo.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return nuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NodeUpdateOne) SaveX(ctx context.Context) *Node {
	n, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

// Exec executes the query on the entity.
func (nuo *NodeUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NodeUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NodeUpdateOne) sqlSave(ctx context.Context) (n *Node, err error) {
	var (
		builder  = sql.Dialect(nuo.driver.Dialect())
		selector = builder.Select(node.Columns...).From(builder.Table(node.Table))
	)
	node.ID(nuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = nuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		n = &Node{config: nuo.config}
		if err := n.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Node: %v", err)
		}
		id = n.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Node with id: %v", nuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Node with the same id: %v", nuo.id)
	}

	tx, err := nuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(node.Table)
	)
	updater = updater.Where(sql.InInts(node.FieldID, ids...))
	if value := nuo.value; value != nil {
		updater.Set(node.FieldValue, *value)
		n.Value = *value
	}
	if value := nuo.addvalue; value != nil {
		updater.Add(node.FieldValue, *value)
		n.Value += *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if nuo.clearedParent {
		query, args := builder.Update(node.ParentTable).
			SetNull(node.ParentColumn).
			Where(sql.InInts(node.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(nuo.parent) > 0 {
		for eid := range nuo.parent {
			query, args := builder.Update(node.ParentTable).
				Set(node.ParentColumn, eid).
				Where(sql.InInts(node.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(nuo.removedChildren) > 0 {
		eids := make([]int, len(nuo.removedChildren))
		for eid := range nuo.removedChildren {
			eids = append(eids, eid)
		}
		query, args := builder.Update(node.ChildrenTable).
			SetNull(node.ChildrenColumn).
			Where(sql.InInts(node.ChildrenColumn, ids...)).
			Where(sql.InInts(node.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(nuo.children) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range nuo.children {
				p.Or().EQ(node.FieldID, eid)
			}
			query, args := builder.Update(node.ChildrenTable).
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
			if int(affected) < len(nuo.children) {
				return nil, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"children\" %v already connected to a different \"Node\"", keys(nuo.children))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return n, nil
}
