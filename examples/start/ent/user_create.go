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
	"github.com/facebookincubator/ent/examples/start/ent/car"
	"github.com/facebookincubator/ent/examples/start/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age    *int
	name   *string
	cars   map[int]struct{}
	groups map[int]struct{}
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

// SetNillableName sets the name field if the given value is not nil.
func (uc *UserCreate) SetNillableName(s *string) *UserCreate {
	if s != nil {
		uc.SetName(*s)
	}
	return uc
}

// AddCarIDs adds the cars edge to Car by ids.
func (uc *UserCreate) AddCarIDs(ids ...int) *UserCreate {
	if uc.cars == nil {
		uc.cars = make(map[int]struct{})
	}
	for i := range ids {
		uc.cars[ids[i]] = struct{}{}
	}
	return uc
}

// AddCars adds the cars edges to Car.
func (uc *UserCreate) AddCars(c ...*Car) *UserCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uc.AddCarIDs(ids...)
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

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if err := user.AgeValidator(*uc.age); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
	}
	if uc.name == nil {
		v := user.DefaultName
		uc.name = &v
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
		res     sql.Result
		builder = sql.Dialect(uc.driver.Dialect())
		u       = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(user.Table).Default()
	if value := uc.age; value != nil {
		insert.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uc.name; value != nil {
		insert.Set(user.FieldName, *value)
		u.Name = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(user.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = int(id)
	if len(uc.cars) > 0 {
		p := sql.P()
		for eid := range uc.cars {
			p.Or().EQ(car.FieldID, eid)
		}
		query, args := builder.Update(user.CarsTable).
			Set(user.CarsColumn, id).
			Where(sql.And(p, sql.IsNull(user.CarsColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.cars) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"cars\" %v already connected to a different \"User\"", keys(uc.cars))})
		}
	}
	if len(uc.groups) > 0 {
		for eid := range uc.groups {

			query, args := builder.Insert(user.GroupsTable).
				Columns(user.GroupsPrimaryKey[1], user.GroupsPrimaryKey[0]).
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
