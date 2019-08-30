// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/traversal/ent/pet"
	"github.com/facebookincubator/ent/examples/traversal/ent/predicate"
	"github.com/facebookincubator/ent/examples/traversal/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	name           *string
	friends        map[int]struct{}
	owner          map[int]struct{}
	removedFriends map[int]struct{}
	clearedOwner   bool
	predicates     []predicate.Pet
}

// Where adds a new predicate for the builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.name = &s
	return pu
}

// AddFriendIDs adds the friends edge to Pet by ids.
func (pu *PetUpdate) AddFriendIDs(ids ...int) *PetUpdate {
	if pu.friends == nil {
		pu.friends = make(map[int]struct{})
	}
	for i := range ids {
		pu.friends[ids[i]] = struct{}{}
	}
	return pu
}

// AddFriends adds the friends edges to Pet.
func (pu *PetUpdate) AddFriends(p ...*Pet) *PetUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.AddFriendIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (pu *PetUpdate) SetOwnerID(id int) *PetUpdate {
	if pu.owner == nil {
		pu.owner = make(map[int]struct{})
	}
	pu.owner[id] = struct{}{}
	return pu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to User.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// RemoveFriendIDs removes the friends edge to Pet by ids.
func (pu *PetUpdate) RemoveFriendIDs(ids ...int) *PetUpdate {
	if pu.removedFriends == nil {
		pu.removedFriends = make(map[int]struct{})
	}
	for i := range ids {
		pu.removedFriends[ids[i]] = struct{}{}
	}
	return pu
}

// RemoveFriends removes friends edges to Pet.
func (pu *PetUpdate) RemoveFriends(p ...*Pet) *PetUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pu.RemoveFriendIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.clearedOwner = true
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	if len(pu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(pet.FieldID).From(sql.Table(pet.Table))
	for _, p := range pu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = pu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := pu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(pet.Table).Where(sql.InInts(pet.FieldID, ids...))
	)
	if pu.name != nil {
		update = true
		builder.Set(pet.FieldName, *pu.name)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.removedFriends) > 0 {
		eids := make([]int, len(pu.removedFriends))
		for eid := range pu.removedFriends {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(pet.FriendsTable).
			Where(sql.InInts(pet.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(pet.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
		query, args = sql.Delete(pet.FriendsTable).
			Where(sql.InInts(pet.FriendsPrimaryKey[1], ids...)).
			Where(sql.InInts(pet.FriendsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.friends) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range pu.friends {
				values = append(values, []int{id, eid}, []int{eid, id})
			}
		}
		builder := sql.Insert(pet.FriendsTable).
			Columns(pet.FriendsPrimaryKey[0], pet.FriendsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if pu.clearedOwner {
		query, args := sql.Update(pet.OwnerTable).
			SetNull(pet.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.owner) > 0 {
		for eid := range pu.owner {
			query, args := sql.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.InInts(pet.FieldID, ids...)).
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

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	id             int
	name           *string
	friends        map[int]struct{}
	owner          map[int]struct{}
	removedFriends map[int]struct{}
	clearedOwner   bool
}

// SetName sets the name field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.name = &s
	return puo
}

// AddFriendIDs adds the friends edge to Pet by ids.
func (puo *PetUpdateOne) AddFriendIDs(ids ...int) *PetUpdateOne {
	if puo.friends == nil {
		puo.friends = make(map[int]struct{})
	}
	for i := range ids {
		puo.friends[ids[i]] = struct{}{}
	}
	return puo
}

// AddFriends adds the friends edges to Pet.
func (puo *PetUpdateOne) AddFriends(p ...*Pet) *PetUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.AddFriendIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (puo *PetUpdateOne) SetOwnerID(id int) *PetUpdateOne {
	if puo.owner == nil {
		puo.owner = make(map[int]struct{})
	}
	puo.owner[id] = struct{}{}
	return puo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to User.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// RemoveFriendIDs removes the friends edge to Pet by ids.
func (puo *PetUpdateOne) RemoveFriendIDs(ids ...int) *PetUpdateOne {
	if puo.removedFriends == nil {
		puo.removedFriends = make(map[int]struct{})
	}
	for i := range ids {
		puo.removedFriends[ids[i]] = struct{}{}
	}
	return puo
}

// RemoveFriends removes friends edges to Pet.
func (puo *PetUpdateOne) RemoveFriends(p ...*Pet) *PetUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return puo.RemoveFriendIDs(ids...)
}

// ClearOwner clears the owner edge to User.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.clearedOwner = true
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	if len(puo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (pe *Pet, err error) {
	selector := sql.Select(pet.Columns...).From(sql.Table(pet.Table))
	pet.ID(puo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = puo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		pe = &Pet{config: puo.config}
		if err := pe.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Pet: %v", err)
		}
		id = pe.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Pet not found with id: %v", puo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Pet with the same id: %v", puo.id)
	}

	tx, err := puo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(pet.Table).Where(sql.InInts(pet.FieldID, ids...))
	)
	if puo.name != nil {
		update = true
		builder.Set(pet.FieldName, *puo.name)
		pe.Name = *puo.name
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.removedFriends) > 0 {
		eids := make([]int, len(puo.removedFriends))
		for eid := range puo.removedFriends {
			eids = append(eids, eid)
		}
		query, args := sql.Delete(pet.FriendsTable).
			Where(sql.InInts(pet.FriendsPrimaryKey[0], ids...)).
			Where(sql.InInts(pet.FriendsPrimaryKey[1], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		query, args = sql.Delete(pet.FriendsTable).
			Where(sql.InInts(pet.FriendsPrimaryKey[1], ids...)).
			Where(sql.InInts(pet.FriendsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.friends) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range puo.friends {
				values = append(values, []int{id, eid}, []int{eid, id})
			}
		}
		builder := sql.Insert(pet.FriendsTable).
			Columns(pet.FriendsPrimaryKey[0], pet.FriendsPrimaryKey[1])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if puo.clearedOwner {
		query, args := sql.Update(pet.OwnerTable).
			SetNull(pet.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.owner) > 0 {
		for eid := range puo.owner {
			query, args := sql.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.InInts(pet.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
