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
	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age  *int
	name *string
	card map[int]struct{}
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

// SetCardID sets the card edge to Card by id.
func (uc *UserCreate) SetCardID(id int) *UserCreate {
	if uc.card == nil {
		uc.card = make(map[int]struct{})
	}
	uc.card[id] = struct{}{}
	return uc
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uc *UserCreate) SetNillableCardID(id *int) *UserCreate {
	if id != nil {
		uc = uc.SetCardID(*id)
	}
	return uc
}

// SetCard sets the card edge to Card.
func (uc *UserCreate) SetCard(c *Card) *UserCreate {
	return uc.SetCardID(c.ID)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(uc.card) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"card\"")
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
	if value := uc.age; value != nil {
		builder.Set(user.FieldAge, *value)
		u.Age = *value
	}
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
	if len(uc.card) > 0 {
		eid := keys(uc.card)[0]
		query, args := sql.Update(user.CardTable).
			Set(user.CardColumn, id).
			Where(sql.EQ(card.FieldID, eid).And().IsNull(user.CardColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.card) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"card\" %v already connected to a different \"User\"", keys(uc.card))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
