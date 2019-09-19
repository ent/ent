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
	"github.com/facebookincubator/ent/examples/traversal/ent/group"
	"github.com/facebookincubator/ent/examples/traversal/ent/predicate"
	"github.com/facebookincubator/ent/examples/traversal/ent/user"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	name         *string
	users        map[int]struct{}
	admin        map[int]struct{}
	removedUsers map[int]struct{}
	clearedAdmin bool
	predicates   []predicate.Group
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetName sets the name field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.name = &s
	return gu
}

// AddUserIDs adds the users edge to User by ids.
func (gu *GroupUpdate) AddUserIDs(ids ...int) *GroupUpdate {
	if gu.users == nil {
		gu.users = make(map[int]struct{})
	}
	for i := range ids {
		gu.users[ids[i]] = struct{}{}
	}
	return gu
}

// AddUsers adds the users edges to User.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetAdminID sets the admin edge to User by id.
func (gu *GroupUpdate) SetAdminID(id int) *GroupUpdate {
	if gu.admin == nil {
		gu.admin = make(map[int]struct{})
	}
	gu.admin[id] = struct{}{}
	return gu
}

// SetNillableAdminID sets the admin edge to User by id if the given value is not nil.
func (gu *GroupUpdate) SetNillableAdminID(id *int) *GroupUpdate {
	if id != nil {
		gu = gu.SetAdminID(*id)
	}
	return gu
}

// SetAdmin sets the admin edge to User.
func (gu *GroupUpdate) SetAdmin(u *User) *GroupUpdate {
	return gu.SetAdminID(u.ID)
}

// RemoveUserIDs removes the users edge to User by ids.
func (gu *GroupUpdate) RemoveUserIDs(ids ...int) *GroupUpdate {
	if gu.removedUsers == nil {
		gu.removedUsers = make(map[int]struct{})
	}
	for i := range ids {
		gu.removedUsers[ids[i]] = struct{}{}
	}
	return gu
}

// RemoveUsers removes users edges to User.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearAdmin clears the admin edge to User.
func (gu *GroupUpdate) ClearAdmin() *GroupUpdate {
	gu.clearedAdmin = true
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	if len(gu.admin) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"admin\"")
	}
	return gu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(group.FieldID).From(sql.Table(group.Table))
	for _, p := range gu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = gu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := gu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(group.Table).Where(sql.InInts(group.FieldID, ids...))
	)
	if value := gu.name; value != nil {
		builder.Set(group.FieldName, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.removedUsers) > 0 {
		eids := make([]int, len(gu.removedUsers))
		for eid := range gu.removedUsers {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(group.UsersTable).
			Where(sql.InInts(group.UsersPrimaryKey[0], ids...)).
			Where(sql.InInts(group.UsersPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.users) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range gu.users {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(group.UsersTable).
			Columns(group.UsersPrimaryKey[0], group.UsersPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if gu.clearedAdmin {
		query, args := sql.Update(group.AdminTable).
			SetNull(group.AdminColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.admin) > 0 {
		for eid := range gu.admin {
			query, args := sql.Update(group.AdminTable).
				Set(group.AdminColumn, eid).
				Where(sql.InInts(group.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	id           int
	name         *string
	users        map[int]struct{}
	admin        map[int]struct{}
	removedUsers map[int]struct{}
	clearedAdmin bool
}

// SetName sets the name field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.name = &s
	return guo
}

// AddUserIDs adds the users edge to User by ids.
func (guo *GroupUpdateOne) AddUserIDs(ids ...int) *GroupUpdateOne {
	if guo.users == nil {
		guo.users = make(map[int]struct{})
	}
	for i := range ids {
		guo.users[ids[i]] = struct{}{}
	}
	return guo
}

// AddUsers adds the users edges to User.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetAdminID sets the admin edge to User by id.
func (guo *GroupUpdateOne) SetAdminID(id int) *GroupUpdateOne {
	if guo.admin == nil {
		guo.admin = make(map[int]struct{})
	}
	guo.admin[id] = struct{}{}
	return guo
}

// SetNillableAdminID sets the admin edge to User by id if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableAdminID(id *int) *GroupUpdateOne {
	if id != nil {
		guo = guo.SetAdminID(*id)
	}
	return guo
}

// SetAdmin sets the admin edge to User.
func (guo *GroupUpdateOne) SetAdmin(u *User) *GroupUpdateOne {
	return guo.SetAdminID(u.ID)
}

// RemoveUserIDs removes the users edge to User by ids.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...int) *GroupUpdateOne {
	if guo.removedUsers == nil {
		guo.removedUsers = make(map[int]struct{})
	}
	for i := range ids {
		guo.removedUsers[ids[i]] = struct{}{}
	}
	return guo
}

// RemoveUsers removes users edges to User.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearAdmin clears the admin edge to User.
func (guo *GroupUpdateOne) ClearAdmin() *GroupUpdateOne {
	guo.clearedAdmin = true
	return guo
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	if len(guo.admin) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"admin\"")
	}
	return guo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (gr *Group, err error) {
	selector := sql.Select(group.Columns...).From(sql.Table(group.Table))
	group.ID(guo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = guo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		gr = &Group{config: guo.config}
		if err := gr.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Group: %v", err)
		}
		id = gr.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Group not found with id: %v", guo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Group with the same id: %v", guo.id)
	}

	tx, err := guo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(group.Table).Where(sql.InInts(group.FieldID, ids...))
	)
	if value := guo.name; value != nil {
		builder.Set(group.FieldName, *value)
		gr.Name = *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.removedUsers) > 0 {
		eids := make([]int, len(guo.removedUsers))
		for eid := range guo.removedUsers {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(group.UsersTable).
			Where(sql.InInts(group.UsersPrimaryKey[0], ids...)).
			Where(sql.InInts(group.UsersPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.users) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range guo.users {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(group.UsersTable).
			Columns(group.UsersPrimaryKey[0], group.UsersPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if guo.clearedAdmin {
		query, args := sql.Update(group.AdminTable).
			SetNull(group.AdminColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.admin) > 0 {
		for eid := range guo.admin {
			query, args := sql.Update(group.AdminTable).
				Set(group.AdminColumn, eid).
				Where(sql.InInts(group.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
