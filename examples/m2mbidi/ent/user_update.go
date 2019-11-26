// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/m2mbidi/ent/predicate"
	"github.com/facebookincubator/ent/examples/m2mbidi/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age            *int
	addage         *int
	name           *string
	friends        map[int]struct{}
	removedFriends map[int]struct{}
	predicates     []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	uu.addage = nil
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	if uu.addage == nil {
		uu.addage = &i
	} else {
		*uu.addage += i
	}
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// AddFriendIDs adds the friends edge to User by ids.
func (uu *UserUpdate) AddFriendIDs(ids ...int) *UserUpdate {
	if uu.friends == nil {
		uu.friends = make(map[int]struct{})
	}
	for i := range ids {
		uu.friends[ids[i]] = struct{}{}
	}
	return uu
}

// AddFriends adds the friends edges to User.
func (uu *UserUpdate) AddFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFriendIDs(ids...)
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (uu *UserUpdate) RemoveFriendIDs(ids ...int) *UserUpdate {
	if uu.removedFriends == nil {
		uu.removedFriends = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedFriends[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFriends removes friends edges to User.
func (uu *UserUpdate) RemoveFriends(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFriendIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(uu.driver.Dialect())
		selector = builder.Select(user.FieldID).From(builder.Table(user.Table))
	)
	for _, p := range uu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := uu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(user.Table)
	)
	updater = updater.Where(sql.InInts(user.FieldID, ids...))
	if value := uu.age; value != nil {
		updater.Set(user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		updater.Add(user.FieldAge, *value)
	}
	if value := uu.name; value != nil {
		updater.Set(user.FieldName, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.removedFriends) > 0 {
		eids := make([]int, len(uu.removedFriends))
		for eid := range uu.removedFriends {
			eids = append(eids, eid)
		}
		query, args := builder.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
		query, args = builder.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[1], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.friends) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uu.friends {
				values = append(values, []int{id, eid}, []int{eid, id})
			}
		}
		builder := builder.Insert(user.FriendsTable).
			Columns(user.FriendsPrimaryKey[0], user.FriendsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id             int
	age            *int
	addage         *int
	name           *string
	friends        map[int]struct{}
	removedFriends map[int]struct{}
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	uuo.addage = nil
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	if uuo.addage == nil {
		uuo.addage = &i
	} else {
		*uuo.addage += i
	}
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// AddFriendIDs adds the friends edge to User by ids.
func (uuo *UserUpdateOne) AddFriendIDs(ids ...int) *UserUpdateOne {
	if uuo.friends == nil {
		uuo.friends = make(map[int]struct{})
	}
	for i := range ids {
		uuo.friends[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFriends adds the friends edges to User.
func (uuo *UserUpdateOne) AddFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFriendIDs(ids...)
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (uuo *UserUpdateOne) RemoveFriendIDs(ids ...int) *UserUpdateOne {
	if uuo.removedFriends == nil {
		uuo.removedFriends = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedFriends[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFriends removes friends edges to User.
func (uuo *UserUpdateOne) RemoveFriends(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFriendIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	var (
		builder  = sql.Dialect(uuo.driver.Dialect())
		selector = builder.Select(user.Columns...).From(builder.Table(user.Table))
	)
	user.ID(uuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		u = &User{config: uuo.config}
		if err := u.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into User: %v", err)
		}
		id = u.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("User with id: %v", uuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(user.Table)
	)
	updater = updater.Where(sql.InInts(user.FieldID, ids...))
	if value := uuo.age; value != nil {
		updater.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uuo.addage; value != nil {
		updater.Add(user.FieldAge, *value)
		u.Age += *value
	}
	if value := uuo.name; value != nil {
		updater.Set(user.FieldName, *value)
		u.Name = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.removedFriends) > 0 {
		eids := make([]int, len(uuo.removedFriends))
		for eid := range uuo.removedFriends {
			eids = append(eids, eid)
		}
		query, args := builder.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		query, args = builder.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[1], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.friends) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uuo.friends {
				values = append(values, []int{id, eid}, []int{eid, id})
			}
		}
		builder := builder.Insert(user.FriendsTable).
			Columns(user.FriendsPrimaryKey[0], user.FriendsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
