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
	"github.com/facebookincubator/ent/entc/integration/idtype/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	name      *string
	spouse    map[uint64]struct{}
	followers map[uint64]struct{}
	following map[uint64]struct{}
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetSpouseID sets the spouse edge to User by id.
func (uc *UserCreate) SetSpouseID(id uint64) *UserCreate {
	if uc.spouse == nil {
		uc.spouse = make(map[uint64]struct{})
	}
	uc.spouse[id] = struct{}{}
	return uc
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableSpouseID(id *uint64) *UserCreate {
	if id != nil {
		uc = uc.SetSpouseID(*id)
	}
	return uc
}

// SetSpouse sets the spouse edge to User.
func (uc *UserCreate) SetSpouse(u *User) *UserCreate {
	return uc.SetSpouseID(u.ID)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uc *UserCreate) AddFollowerIDs(ids ...uint64) *UserCreate {
	if uc.followers == nil {
		uc.followers = make(map[uint64]struct{})
	}
	for i := range ids {
		uc.followers[ids[i]] = struct{}{}
	}
	return uc
}

// AddFollowers adds the followers edges to User.
func (uc *UserCreate) AddFollowers(u ...*User) *UserCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uc *UserCreate) AddFollowingIDs(ids ...uint64) *UserCreate {
	if uc.following == nil {
		uc.following = make(map[uint64]struct{})
	}
	for i := range ids {
		uc.following[ids[i]] = struct{}{}
	}
	return uc
}

// AddFollowing adds the following edges to User.
func (uc *UserCreate) AddFollowing(u ...*User) *UserCreate {
	ids := make([]uint64, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowingIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(uc.spouse) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
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
	builder := sql.Insert(user.Table).Default(uc.driver.Dialect())
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
	u.ID = uint64(id)
	if len(uc.spouse) > 0 {
		for eid := range uc.spouse {
			query, args := sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, id)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			query, args = sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, id).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(uc.spouse) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
			}
		}
	}
	if len(uc.followers) > 0 {
		for eid := range uc.followers {

			query, args := sql.Insert(user.FollowersTable).
				Columns(user.FollowersPrimaryKey[1], user.FollowersPrimaryKey[0]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.following) > 0 {
		for eid := range uc.following {

			query, args := sql.Insert(user.FollowingTable).
				Columns(user.FollowingPrimaryKey[0], user.FollowingPrimaryKey[1]).
				Values(id, eid).
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
