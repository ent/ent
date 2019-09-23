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
	"github.com/facebookincubator/ent/entc/integration/idtype/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/idtype/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	name             *string
	spouse           map[uint64]struct{}
	followers        map[uint64]struct{}
	following        map[uint64]struct{}
	clearedSpouse    bool
	removedFollowers map[uint64]struct{}
	removedFollowing map[uint64]struct{}
	predicates       []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetSpouseID sets the spouse edge to User by id.
func (uu *UserUpdate) SetSpouseID(id uint64) *UserUpdate {
	if uu.spouse == nil {
		uu.spouse = make(map[uint64]struct{})
	}
	uu.spouse[id] = struct{}{}
	return uu
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *uint64) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the spouse edge to User.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uu *UserUpdate) AddFollowerIDs(ids ...uint64) *UserUpdate {
	if uu.followers == nil {
		uu.followers = make(map[uint64]struct{})
	}
	for i := range ids {
		uu.followers[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowers adds the followers edges to User.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uu *UserUpdate) AddFollowingIDs(ids ...uint64) *UserUpdate {
	if uu.following == nil {
		uu.following = make(map[uint64]struct{})
	}
	for i := range ids {
		uu.following[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowing adds the following edges to User.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// ClearSpouse clears the spouse edge to User.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.clearedSpouse = true
	return uu
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...uint64) *UserUpdate {
	if uu.removedFollowers == nil {
		uu.removedFollowers = make(map[uint64]struct{})
	}
	for i := range ids {
		uu.removedFollowers[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowers removes followers edges to User.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...uint64) *UserUpdate {
	if uu.removedFollowing == nil {
		uu.removedFollowing = make(map[uint64]struct{})
	}
	for i := range ids {
		uu.removedFollowing[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowing removes following edges to User.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if len(uu.spouse) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
	}
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
	if value := uu.name; value != nil {
		builder.Set(user.FieldName, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if uu.clearedSpouse {
		query, args := sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
		query, args = sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.SpouseColumn, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.spouse) > 0 {
		if n := len(ids); n > 1 {
			return 0, rollback(tx, fmt.Errorf("ent: can't link O2O edge \"spouse\" to %d vertices (> 1)", n))
		}
		for eid := range uu.spouse {
			eid := int(eid)
			query, args := sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, ids[0])).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			query, args = sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, ids[0]).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.spouse) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
			}
		}
	}
	if len(uu.removedFollowers) > 0 {
		eids := make([]int, len(uu.removedFollowers))
		for eid := range uu.removedFollowers {
			eid := int(eid)
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
				eid := int(eid)
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
			eid := int(eid)
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
				eid := int(eid)
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
	id               uint64
	name             *string
	spouse           map[uint64]struct{}
	followers        map[uint64]struct{}
	following        map[uint64]struct{}
	clearedSpouse    bool
	removedFollowers map[uint64]struct{}
	removedFollowing map[uint64]struct{}
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetSpouseID sets the spouse edge to User by id.
func (uuo *UserUpdateOne) SetSpouseID(id uint64) *UserUpdateOne {
	if uuo.spouse == nil {
		uuo.spouse = make(map[uint64]struct{})
	}
	uuo.spouse[id] = struct{}{}
	return uuo
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *uint64) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the spouse edge to User.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...uint64) *UserUpdateOne {
	if uuo.followers == nil {
		uuo.followers = make(map[uint64]struct{})
	}
	for i := range ids {
		uuo.followers[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowers adds the followers edges to User.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...uint64) *UserUpdateOne {
	if uuo.following == nil {
		uuo.following = make(map[uint64]struct{})
	}
	for i := range ids {
		uuo.following[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowing adds the following edges to User.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// ClearSpouse clears the spouse edge to User.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.clearedSpouse = true
	return uuo
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...uint64) *UserUpdateOne {
	if uuo.removedFollowers == nil {
		uuo.removedFollowers = make(map[uint64]struct{})
	}
	for i := range ids {
		uuo.removedFollowers[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowers removes followers edges to User.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...uint64) *UserUpdateOne {
	if uuo.removedFollowing == nil {
		uuo.removedFollowing = make(map[uint64]struct{})
	}
	for i := range ids {
		uuo.removedFollowing[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowing removes following edges to User.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if len(uuo.spouse) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
	}
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
		id = int(u.ID)
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
	if uuo.clearedSpouse {
		query, args := sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		query, args = sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.SpouseColumn, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.spouse) > 0 {
		if n := len(ids); n > 1 {
			return nil, rollback(tx, fmt.Errorf("ent: can't link O2O edge \"spouse\" to %d vertices (> 1)", n))
		}
		for eid := range uuo.spouse {
			eid := int(eid)
			query, args := sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, ids[0])).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			query, args = sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, ids[0]).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(uuo.spouse) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
			}
		}
	}
	if len(uuo.removedFollowers) > 0 {
		eids := make([]int, len(uuo.removedFollowers))
		for eid := range uuo.removedFollowers {
			eid := int(eid)
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
				eid := int(eid)
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
			eid := int(eid)
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
				eid := int(eid)
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
