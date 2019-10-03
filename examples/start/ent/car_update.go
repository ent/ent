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
	"github.com/facebookincubator/ent/examples/start/ent/car"
	"github.com/facebookincubator/ent/examples/start/ent/predicate"
	"github.com/facebookincubator/ent/examples/start/ent/user"
)

// CarUpdate is the builder for updating Car entities.
type CarUpdate struct {
	config
	model         *string
	registered_at *time.Time
	owner         map[int]struct{}
	clearedOwner  bool
	predicates    []predicate.Car
}

// Where adds a new predicate for the builder.
func (cu *CarUpdate) Where(ps ...predicate.Car) *CarUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetModel sets the model field.
func (cu *CarUpdate) SetModel(s string) *CarUpdate {
	cu.model = &s
	return cu
}

// SetRegisteredAt sets the registered_at field.
func (cu *CarUpdate) SetRegisteredAt(t time.Time) *CarUpdate {
	cu.registered_at = &t
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CarUpdate) SetOwnerID(id int) *CarUpdate {
	if cu.owner == nil {
		cu.owner = make(map[int]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cu *CarUpdate) SetNillableOwnerID(id *int) *CarUpdate {
	if id != nil {
		cu = cu.SetOwnerID(*id)
	}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CarUpdate) SetOwner(u *User) *CarUpdate {
	return cu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cu *CarUpdate) ClearOwner() *CarUpdate {
	cu.clearedOwner = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CarUpdate) Save(ctx context.Context) (int, error) {
	if len(cu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CarUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CarUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CarUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CarUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(car.FieldID).From(sql.Table(car.Table))
	for _, p := range cu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := cu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		builder = sql.Update(car.Table).Where(sql.InInts(car.FieldID, ids...))
	)
	if value := cu.model; value != nil {
		builder.Set(car.FieldModel, *value)
	}
	if value := cu.registered_at; value != nil {
		builder.Set(car.FieldRegisteredAt, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if cu.clearedOwner {
		query, args := sql.Update(car.OwnerTable).
			SetNull(car.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(cu.owner) > 0 {
		for eid := range cu.owner {
			query, args := sql.Update(car.OwnerTable).
				Set(car.OwnerColumn, eid).
				Where(sql.InInts(car.FieldID, ids...)).
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

// CarUpdateOne is the builder for updating a single Car entity.
type CarUpdateOne struct {
	config
	id            int
	model         *string
	registered_at *time.Time
	owner         map[int]struct{}
	clearedOwner  bool
}

// SetModel sets the model field.
func (cuo *CarUpdateOne) SetModel(s string) *CarUpdateOne {
	cuo.model = &s
	return cuo
}

// SetRegisteredAt sets the registered_at field.
func (cuo *CarUpdateOne) SetRegisteredAt(t time.Time) *CarUpdateOne {
	cuo.registered_at = &t
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CarUpdateOne) SetOwnerID(id int) *CarUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[int]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (cuo *CarUpdateOne) SetNillableOwnerID(id *int) *CarUpdateOne {
	if id != nil {
		cuo = cuo.SetOwnerID(*id)
	}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CarUpdateOne) SetOwner(u *User) *CarUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cuo *CarUpdateOne) ClearOwner() *CarUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CarUpdateOne) Save(ctx context.Context) (*Car, error) {
	if len(cuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CarUpdateOne) SaveX(ctx context.Context) *Car {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CarUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CarUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CarUpdateOne) sqlSave(ctx context.Context) (c *Car, err error) {
	selector := sql.Select(car.Columns...).From(sql.Table(car.Table))
	car.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		c = &Car{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Car: %v", err)
		}
		id = c.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Car with id: %v", cuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Car with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(car.Table).Where(sql.InInts(car.FieldID, ids...))
	)
	if value := cuo.model; value != nil {
		builder.Set(car.FieldModel, *value)
		c.Model = *value
	}
	if value := cuo.registered_at; value != nil {
		builder.Set(car.FieldRegisteredAt, *value)
		c.RegisteredAt = *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if cuo.clearedOwner {
		query, args := sql.Update(car.OwnerTable).
			SetNull(car.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(cuo.owner) > 0 {
		for eid := range cuo.owner {
			query, args := sql.Update(car.OwnerTable).
				Set(car.OwnerColumn, eid).
				Where(sql.InInts(car.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
