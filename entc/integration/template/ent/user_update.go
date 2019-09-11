// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	name           *string
	pets           map[int]struct{}
	friends        map[int]struct{}
	removedPets    map[int]struct{}
	removedFriends map[int]struct{}
	predicates     []predicate.User
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

// AddPetIDs adds the pets edge to Pet by ids.
func (uu *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	if uu.pets == nil {
		uu.pets = make(map[int]struct{})
	}
	for i := range ids {
		uu.pets[ids[i]] = struct{}{}
	}
	return uu
}

// AddPets adds the pets edges to Pet.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
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

// RemovePetIDs removes the pets edge to Pet by ids.
func (uu *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	if uu.removedPets == nil {
		uu.removedPets = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedPets[ids[i]] = struct{}{}
	}
	return uu
}

// RemovePets removes pets edges to Pet.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
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
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uu.name; value != nil {
		update = true
		builder.Set(user.FieldName, *value)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.removedPets) > 0 {
		eids := make([]int, len(uu.removedPets))
		for eid := range uu.removedPets {
			eids = append(eids, eid)
		}
		query, args := sql.Update(user.PetsTable).
			SetNull(user.PetsColumn).
			Where(sql.InInts(user.PetsColumn, ids...)).
			Where(sql.InInts(pet.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.pets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uu.pets {
				p.Or().EQ(pet.FieldID, eid)
			}
			query, args := sql.Update(user.PetsTable).
				Set(user.PetsColumn, id).
				Where(sql.And(p, sql.IsNull(user.PetsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.pets) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uu.pets))})
			}
		}
	}
	if len(uu.removedFriends) > 0 {
		eids := make([]int, len(uu.removedFriends))
		for eid := range uu.removedFriends {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
		query, args = sql.Delete(user.FriendsTable).
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
		builder := sql.Insert(user.FriendsTable).
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
	name           *string
	pets           map[int]struct{}
	friends        map[int]struct{}
	removedPets    map[int]struct{}
	removedFriends map[int]struct{}
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uuo *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	if uuo.pets == nil {
		uuo.pets = make(map[int]struct{})
	}
	for i := range ids {
		uuo.pets[ids[i]] = struct{}{}
	}
	return uuo
}

// AddPets adds the pets edges to Pet.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
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

// RemovePetIDs removes the pets edge to Pet by ids.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	if uuo.removedPets == nil {
		uuo.removedPets = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedPets[ids[i]] = struct{}{}
	}
	return uuo
}

// RemovePets removes pets edges to Pet.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
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
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.name; value != nil {
		update = true
		builder.Set(user.FieldName, *value)
		u.Name = *value
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.removedPets) > 0 {
		eids := make([]int, len(uuo.removedPets))
		for eid := range uuo.removedPets {
			eids = append(eids, eid)
		}
		query, args := sql.Update(user.PetsTable).
			SetNull(user.PetsColumn).
			Where(sql.InInts(user.PetsColumn, ids...)).
			Where(sql.InInts(pet.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.pets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uuo.pets {
				p.Or().EQ(pet.FieldID, eid)
			}
			query, args := sql.Update(user.PetsTable).
				Set(user.PetsColumn, id).
				Where(sql.And(p, sql.IsNull(user.PetsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(uuo.pets) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uuo.pets))})
			}
		}
	}
	if len(uuo.removedFriends) > 0 {
		eids := make([]int, len(uuo.removedFriends))
		for eid := range uuo.removedFriends {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(user.FriendsTable).
			Where(sql.InInts(user.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(user.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		query, args = sql.Delete(user.FriendsTable).
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
		builder := sql.Insert(user.FriendsTable).
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
