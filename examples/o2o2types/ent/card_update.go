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
	"github.com/facebookincubator/ent/examples/o2o2types/ent/predicate"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/user"
)

// CardUpdate is the builder for updating Card entities.
type CardUpdate struct {
	config
	expired      *time.Time
	number       *string
	owner        map[int]struct{}
	clearedOwner bool
	predicates   []predicate.Card
}

// Where adds a new predicate for the builder.
func (cu *CardUpdate) Where(ps ...predicate.Card) *CardUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetExpired sets the expired field.
func (cu *CardUpdate) SetExpired(t time.Time) *CardUpdate {
	cu.expired = &t
	return cu
}

// SetNumber sets the number field.
func (cu *CardUpdate) SetNumber(s string) *CardUpdate {
	cu.number = &s
	return cu
}

// SetOwnerID sets the owner edge to User by id.
func (cu *CardUpdate) SetOwnerID(id int) *CardUpdate {
	if cu.owner == nil {
		cu.owner = make(map[int]struct{})
	}
	cu.owner[id] = struct{}{}
	return cu
}

// SetOwner sets the owner edge to User.
func (cu *CardUpdate) SetOwner(u *User) *CardUpdate {
	return cu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cu *CardUpdate) ClearOwner() *CardUpdate {
	cu.clearedOwner = true
	return cu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CardUpdate) Save(ctx context.Context) (int, error) {
	if len(cu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if cu.clearedOwner && cu.owner == nil {
		return 0, errors.New("ent: clearing a unique edge \"owner\"")
	}
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CardUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CardUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CardUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CardUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(card.FieldID).From(sql.Table(card.Table))
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
		builder = sql.Update(card.Table).Where(sql.InInts(card.FieldID, ids...))
	)
	if value := cu.expired; value != nil {
		builder.Set(card.FieldExpired, *value)
	}
	if value := cu.number; value != nil {
		builder.Set(card.FieldNumber, *value)
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if cu.clearedOwner {
		query, args := sql.Update(card.OwnerTable).
			SetNull(card.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(cu.owner) > 0 {
		for _, id := range ids {
			eid := keys(cu.owner)[0]
			query, args := sql.Update(card.OwnerTable).
				Set(card.OwnerColumn, eid).
				Where(sql.EQ(card.FieldID, id).And().IsNull(card.OwnerColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(cu.owner) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"owner\" %v already connected to a different \"Card\"", keys(cu.owner))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// CardUpdateOne is the builder for updating a single Card entity.
type CardUpdateOne struct {
	config
	id           int
	expired      *time.Time
	number       *string
	owner        map[int]struct{}
	clearedOwner bool
}

// SetExpired sets the expired field.
func (cuo *CardUpdateOne) SetExpired(t time.Time) *CardUpdateOne {
	cuo.expired = &t
	return cuo
}

// SetNumber sets the number field.
func (cuo *CardUpdateOne) SetNumber(s string) *CardUpdateOne {
	cuo.number = &s
	return cuo
}

// SetOwnerID sets the owner edge to User by id.
func (cuo *CardUpdateOne) SetOwnerID(id int) *CardUpdateOne {
	if cuo.owner == nil {
		cuo.owner = make(map[int]struct{})
	}
	cuo.owner[id] = struct{}{}
	return cuo
}

// SetOwner sets the owner edge to User.
func (cuo *CardUpdateOne) SetOwner(u *User) *CardUpdateOne {
	return cuo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (cuo *CardUpdateOne) ClearOwner() *CardUpdateOne {
	cuo.clearedOwner = true
	return cuo
}

// Save executes the query and returns the updated entity.
func (cuo *CardUpdateOne) Save(ctx context.Context) (*Card, error) {
	if len(cuo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	if cuo.clearedOwner && cuo.owner == nil {
		return nil, errors.New("ent: clearing a unique edge \"owner\"")
	}
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CardUpdateOne) SaveX(ctx context.Context) *Card {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CardUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CardUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CardUpdateOne) sqlSave(ctx context.Context) (c *Card, err error) {
	selector := sql.Select(card.Columns...).From(sql.Table(card.Table))
	card.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		c = &Card{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Card: %v", err)
		}
		id = c.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Card not found with id: %v", cuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Card with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(card.Table).Where(sql.InInts(card.FieldID, ids...))
	)
	if value := cuo.expired; value != nil {
		builder.Set(card.FieldExpired, *value)
		c.Expired = *value
	}
	if value := cuo.number; value != nil {
		builder.Set(card.FieldNumber, *value)
		c.Number = *value
	}
	if !builder.Empty() {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if cuo.clearedOwner {
		query, args := sql.Update(card.OwnerTable).
			SetNull(card.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(cuo.owner) > 0 {
		for _, id := range ids {
			eid := keys(cuo.owner)[0]
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
			if int(affected) < len(cuo.owner) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"owner\" %v already connected to a different \"Card\"", keys(cuo.owner))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
