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
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	name    *string
	pets    map[int]struct{}
	friends map[int]struct{}
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uc *UserCreate) AddPetIDs(ids ...int) *UserCreate {
	if uc.pets == nil {
		uc.pets = make(map[int]struct{})
	}
	for i := range ids {
		uc.pets[ids[i]] = struct{}{}
	}
	return uc
}

// AddPets adds the pets edges to Pet.
func (uc *UserCreate) AddPets(p ...*Pet) *UserCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPetIDs(ids...)
}

// AddFriendIDs adds the friends edge to User by ids.
func (uc *UserCreate) AddFriendIDs(ids ...int) *UserCreate {
	if uc.friends == nil {
		uc.friends = make(map[int]struct{})
	}
	for i := range ids {
		uc.friends[ids[i]] = struct{}{}
	}
	return uc
}

// AddFriends adds the friends edges to User.
func (uc *UserCreate) AddFriends(u ...*User) *UserCreate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFriendIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		res sql.Result
		u   = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(uc.driver.Dialect()).
		Insert(user.Table).
		Default()
	if value := uc.name; value != nil {
		builder.Set(user.FieldName, *value)
		u.Name = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if len(uc.pets) > 0 {
		p := sql.P()
		for eid := range uc.pets {
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
		if int(affected) < len(uc.pets) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uc.pets))})
		}
	}
	if len(uc.friends) > 0 {
		for eid := range uc.friends {

			query, args := sql.Insert(user.FriendsTable).
				Columns(user.FriendsPrimaryKey[0], user.FriendsPrimaryKey[1]).
				Values(id, eid).
				Values(eid, id).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
