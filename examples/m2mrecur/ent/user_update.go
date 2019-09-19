// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/m2mrecur/ent/predicate"
	"github.com/facebookincubator/ent/examples/m2mrecur/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age              *int
	addage           *int
	name             *string
	followers        map[int]struct{}
	following        map[int]struct{}
	removedFollowers map[int]struct{}
	removedFollowing map[int]struct{}
	predicates       []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.addage = &i
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uu *UserUpdate) AddFollowerIDs(ids ...int) *UserUpdate {
	if uu.followers == nil {
		uu.followers = make(map[int]struct{})
	}
	for i := range ids {
		uu.followers[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowers adds the followers edges to User.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uu *UserUpdate) AddFollowingIDs(ids ...int) *UserUpdate {
	if uu.following == nil {
		uu.following = make(map[int]struct{})
	}
	for i := range ids {
		uu.following[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowing adds the following edges to User.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...int) *UserUpdate {
	if uu.removedFollowers == nil {
		uu.removedFollowers = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedFollowers[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowers removes followers edges to User.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...int) *UserUpdate {
	if uu.removedFollowing == nil {
		uu.removedFollowing = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedFollowing[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowing removes following edges to User.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
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
	selector := sql.Select(user.FieldID).From(sql.Table(user.Table))
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
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uu.age; value != nil {
		builder.Set(user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		builder.Add(user.FieldAge, *value)
	}
	if value := uu.name; value != nil {
		builder.Set(user.FieldName, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.removedFollowers) > 0 {
		eids := make([]int, len(uu.removedFollowers))
		for eid := range uu.removedFollowers {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FollowersTable).
			Where(sql.InInts(user.FollowersPrimaryKey[1], ids...)).
			Where(sql.InInts(user.FollowersPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.followers) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uu.followers {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(user.FollowersTable).
			Columns(user.FollowersPrimaryKey[1], user.FollowersPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.removedFollowing) > 0 {
		eids := make([]int, len(uu.removedFollowing))
		for eid := range uu.removedFollowing {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FollowingTable).
			Where(sql.InInts(user.FollowingPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FollowingPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.following) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uu.following {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(user.FollowingTable).
			Columns(user.FollowingPrimaryKey[0], user.FollowingPrimaryKey[1])
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
	id               int
	age              *int
	addage           *int
	name             *string
	followers        map[int]struct{}
	following        map[int]struct{}
	removedFollowers map[int]struct{}
	removedFollowing map[int]struct{}
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.addage = &i
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...int) *UserUpdateOne {
	if uuo.followers == nil {
		uuo.followers = make(map[int]struct{})
	}
	for i := range ids {
		uuo.followers[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowers adds the followers edges to User.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...int) *UserUpdateOne {
	if uuo.following == nil {
		uuo.following = make(map[int]struct{})
	}
	for i := range ids {
		uuo.following[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowing adds the following edges to User.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...int) *UserUpdateOne {
	if uuo.removedFollowers == nil {
		uuo.removedFollowers = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedFollowers[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowers removes followers edges to User.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...int) *UserUpdateOne {
	if uuo.removedFollowing == nil {
		uuo.removedFollowing = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedFollowing[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowing removes following edges to User.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
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
	selector := sql.Select(user.Columns...).From(sql.Table(user.Table))
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
		return nil, fmt.Errorf("ent: User not found with id: %v", uuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.age; value != nil {
		builder.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uuo.addage; value != nil {
		builder.Add(user.FieldAge, *value)
		u.Age += *value
	}
	if value := uuo.name; value != nil {
		builder.Set(user.FieldName, *value)
		u.Name = *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.removedFollowers) > 0 {
		eids := make([]int, len(uuo.removedFollowers))
		for eid := range uuo.removedFollowers {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FollowersTable).
			Where(sql.InInts(user.FollowersPrimaryKey[1], ids...)).
			Where(sql.InInts(user.FollowersPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.followers) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uuo.followers {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(user.FollowersTable).
			Columns(user.FollowersPrimaryKey[1], user.FollowersPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.removedFollowing) > 0 {
		eids := make([]int, len(uuo.removedFollowing))
		for eid := range uuo.removedFollowing {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FollowingTable).
			Where(sql.InInts(user.FollowingPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FollowingPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.following) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uuo.following {
				values = append(values, []int{id, eid})
			}
		}
		builder := sql.Insert(user.FollowingTable).
			Columns(user.FollowingPrimaryKey[0], user.FollowingPrimaryKey[1])
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
