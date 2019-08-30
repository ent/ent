// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/traversal/ent/group"
	"github.com/facebookincubator/ent/examples/traversal/ent/pet"
	"github.com/facebookincubator/ent/examples/traversal/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age     *int
	name    *string
	pets    map[int]struct{}
	friends map[int]struct{}
	groups  map[int]struct{}
	manage  map[int]struct{}
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.age = &i
	return uc
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

// AddGroupIDs adds the groups edge to Group by ids.
func (uc *UserCreate) AddGroupIDs(ids ...int) *UserCreate {
	if uc.groups == nil {
		uc.groups = make(map[int]struct{})
	}
	for i := range ids {
		uc.groups[ids[i]] = struct{}{}
	}
	return uc
}

// AddGroups adds the groups edges to Group.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddManageIDs adds the manage edge to Group by ids.
func (uc *UserCreate) AddManageIDs(ids ...int) *UserCreate {
	if uc.manage == nil {
		uc.manage = make(map[int]struct{})
	}
	for i := range ids {
		uc.manage[ids[i]] = struct{}{}
	}
	return uc
}

// AddManage adds the manage edges to Group.
func (uc *UserCreate) AddManage(g ...*Group) *UserCreate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddManageIDs(ids...)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
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
	builder := sql.Insert(user.Table).Default(uc.driver.Dialect())
	if uc.age != nil {
		builder.Set(user.FieldAge, *uc.age)
		u.Age = *uc.age
	}
	if uc.name != nil {
		builder.Set(user.FieldName, *uc.name)
		u.Name = *uc.name
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
	if len(uc.groups) > 0 {
		for eid := range uc.groups {

			query, args := sql.Insert(user.GroupsTable).
				Columns(user.GroupsPrimaryKey[1], user.GroupsPrimaryKey[0]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.manage) > 0 {
		p := sql.P()
		for eid := range uc.manage {
			p.Or().EQ(group.FieldID, eid)
		}
		query, args := sql.Update(user.ManageTable).
			Set(user.ManageColumn, id).
			Where(sql.And(p, sql.IsNull(user.ManageColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.manage) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"manage\" %v already connected to a different \"User\"", keys(uc.manage))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
