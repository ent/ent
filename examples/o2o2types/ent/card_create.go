// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
)

// CardCreate is the builder for creating a Card entity.
type CardCreate struct {
	config
	expired *time.Time
	number  *string
	owner   map[int]struct{}
}

// SetExpired sets the expired field.
func (cc *CardCreate) SetExpired(t time.Time) *CardCreate {
	cc.expired = &t
	return cc
}

// SetNumber sets the number field.
func (cc *CardCreate) SetNumber(s string) *CardCreate {
	cc.number = &s
	return cc
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CardCreate) SetOwnerID(id int) *CardCreate {
	if cc.owner == nil {
		cc.owner = make(map[int]struct{})
	}
	cc.owner[id] = struct{}{}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CardCreate) SetOwner(u *User) *CardCreate {
	return cc.SetOwnerID(u.ID)
}

// Save creates the Card in the database.
func (cc *CardCreate) Save(ctx context.Context) (*Card, error) {
	if cc.expired == nil {
		return nil, errors.New("ent: missing required field \"expired\"")
	}
	if cc.number == nil {
		return nil, errors.New("ent: missing required field \"number\"")
	}
	if len(cc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if cc.owner == nil {
		return nil, errors.New("ent: missing required edge \"owner\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CardCreate) SaveX(ctx context.Context) *Card {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CardCreate) sqlSave(ctx context.Context) (*Card, error) {
	var (
		res sql.Result
		c   = &Card{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Dialect(cc.driver.Dialect()).
		Insert(card.Table).
		Default()
	if value := cc.expired; value != nil {
		builder.Set(card.FieldExpired, *value)
		c.Expired = *value
	}
	if value := cc.number; value != nil {
		builder.Set(card.FieldNumber, *value)
		c.Number = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if len(cc.owner) > 0 {
		eid := keys(cc.owner)[0]
		query, args := sql.Update(card.OwnerTable).
			Set(card.OwnerColumn, eid).
			Where(sql.EQ(card.FieldID, id).And().IsNull(card.OwnerColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(cc.owner) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"owner\" %v already connected to a different \"Card\"", keys(cc.owner))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
