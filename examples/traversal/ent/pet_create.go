// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/traversal/ent/pet"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	name    *string
	friends map[int]struct{}
	owner   map[int]struct{}
}

// SetName sets the name field.
func (pc *PetCreate) SetName(s string) *PetCreate {
	pc.name = &s
	return pc
}

// AddFriendIDs adds the friends edge to Pet by ids.
func (pc *PetCreate) AddFriendIDs(ids ...int) *PetCreate {
	if pc.friends == nil {
		pc.friends = make(map[int]struct{})
	}
	for i := range ids {
		pc.friends[ids[i]] = struct{}{}
	}
	return pc
}

// AddFriends adds the friends edges to Pet.
func (pc *PetCreate) AddFriends(p ...*Pet) *PetCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pc.AddFriendIDs(ids...)
}

// SetOwnerID sets the owner edge to User by id.
func (pc *PetCreate) SetOwnerID(id int) *PetCreate {
	if pc.owner == nil {
		pc.owner = make(map[int]struct{})
	}
	pc.owner[id] = struct{}{}
	return pc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *int) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the owner edge to User.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	if pc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(pc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	var (
		res sql.Result
		pe  = &Pet{config: pc.config}
	)
	tx, err := pc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(pet.Table).Default(pc.driver.Dialect())
	if value := pc.name; value != nil {
		builder.Set(pet.FieldName, *value)
		pe.Name = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	pe.ID = int(id)
	if len(pc.friends) > 0 {
		for eid := range pc.friends {

			query, args := sql.Insert(pet.FriendsTable).
				Columns(pet.FriendsPrimaryKey[0], pet.FriendsPrimaryKey[1]).
				Values(id, eid).
				Values(eid, id).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(pc.owner) > 0 {
		for eid := range pc.owner {
			query, args := sql.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.EQ(pet.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
