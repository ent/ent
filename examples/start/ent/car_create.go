// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/start/ent/car"
)

// CarCreate is the builder for creating a Car entity.
type CarCreate struct {
	config
	model         *string
	registered_at *time.Time
	owner         map[int]struct{}
}

// SetModel sets the model field.
func (cc *CarCreate) SetModel(s string) *CarCreate {
	cc.model = &s
	return cc
}

// SetRegisteredAt sets the registered_at field.
func (cc *CarCreate) SetRegisteredAt(t time.Time) *CarCreate {
	cc.registered_at = &t
	return cc
}

// SetOwnerID sets the owner edge to User by id.
func (cc *CarCreate) SetOwnerID(id int) *CarCreate {
	if cc.owner == nil {
		cc.owner = make(map[int]struct{})
	}
	cc.owner[id] = struct{}{}
	return cc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cc *CarCreate) SetNillableOwnerID(id *int) *CarCreate {
	if id != nil {
		cc = cc.SetOwnerID(*id)
	}
	return cc
}

// SetOwner sets the owner edge to User.
func (cc *CarCreate) SetOwner(u *User) *CarCreate {
	return cc.SetOwnerID(u.ID)
}

// Save creates the Car in the database.
func (cc *CarCreate) Save(ctx context.Context) (*Car, error) {
	if cc.model == nil {
		return nil, errors.New("ent: missing required field \"model\"")
	}
	if cc.registered_at == nil {
		return nil, errors.New("ent: missing required field \"registered_at\"")
	}
	if len(cc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CarCreate) SaveX(ctx context.Context) *Car {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (cc *CarCreate) sqlSave(ctx context.Context) (*Car, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(cc.driver.Dialect())
		c       = &Car{config: cc.config}
	)
	tx, err := cc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(car.Table).Default()
	if value := cc.model; value != nil {
		insert.Set(car.FieldModel, *value)
		c.Model = *value
	}
	if value := cc.registered_at; value != nil {
		insert.Set(car.FieldRegisteredAt, *value)
		c.RegisteredAt = *value
	}
	id, err := insertLastID(ctx, tx, insert.Returning(car.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	c.ID = int(id)
	if len(cc.owner) > 0 {
		for eid := range cc.owner {
			query, args := builder.Update(car.OwnerTable).
				Set(car.OwnerColumn, eid).
				Where(sql.EQ(car.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
