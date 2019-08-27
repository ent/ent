// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/o2o2types/ent/card"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/predicate"
	"github.com/facebookincubator/ent/examples/o2o2types/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age         *int
	name        *string
	card        map[int]struct{}
	clearedCard bool
	predicates  []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetCardID sets the card edge to Card by id.
func (uu *UserUpdate) SetCardID(id int) *UserUpdate {
	if uu.card == nil {
		uu.card = make(map[int]struct{})
	}
	uu.card[id] = struct{}{}
	return uu
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uu *UserUpdate) SetNillableCardID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetCardID(*id)
	}
	return uu
}

// SetCard sets the card edge to Card.
func (uu *UserUpdate) SetCard(c *Card) *UserUpdate {
	return uu.SetCardID(c.ID)
}

// ClearCard clears the card edge to Card.
func (uu *UserUpdate) ClearCard() *UserUpdate {
	uu.clearedCard = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if len(uu.card) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"card\"")
	}
	return uu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(user.FieldID).From(sql.Table(user.Table))
	for _, p := range uu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := uu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if uu.age != nil {
		update = true
		builder.Set(user.FieldAge, *uu.age)
	}
	if uu.name != nil {
		update = true
		builder.Set(user.FieldName, *uu.name)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if uu.clearedCard {
		query, args := sql.Update(user.CardTable).
			SetNull(user.CardColumn).
			Where(sql.InInts(card.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.card) > 0 {
		for _, id := range ids {
			eid := keys(uu.card)[0]
			query, args := sql.Update(user.CardTable).
				Set(user.CardColumn, id).
				Where(sql.EQ(card.FieldID, eid).And().IsNull(user.CardColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.card) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"card\" %v already connected to a different \"User\"", keys(uu.card))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id          int
	age         *int
	name        *string
	card        map[int]struct{}
	clearedCard bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetCardID sets the card edge to Card by id.
func (uuo *UserUpdateOne) SetCardID(id int) *UserUpdateOne {
	if uuo.card == nil {
		uuo.card = make(map[int]struct{})
	}
	uuo.card[id] = struct{}{}
	return uuo
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCardID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCardID(*id)
	}
	return uuo
}

// SetCard sets the card edge to Card.
func (uuo *UserUpdateOne) SetCard(c *Card) *UserUpdateOne {
	return uuo.SetCardID(c.ID)
}

// ClearCard clears the card edge to Card.
func (uuo *UserUpdateOne) ClearCard() *UserUpdateOne {
	uuo.clearedCard = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if len(uuo.card) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"card\"")
	}
	return uuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	u, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return u
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (u *User, err error) {
	selector := sql.Select(user.Columns...).From(sql.Table(user.Table))
	user.ID(uuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = uuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		u = &User{config: uuo.config}
		if err := u.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into User: %v", err)
		}
		id = u.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: User not found with id: %v", uuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if uuo.age != nil {
		update = true
		builder.Set(user.FieldAge, *uuo.age)
		u.Age = *uuo.age
	}
	if uuo.name != nil {
		update = true
		builder.Set(user.FieldName, *uuo.name)
		u.Name = *uuo.name
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if uuo.clearedCard {
		query, args := sql.Update(user.CardTable).
			SetNull(user.CardColumn).
			Where(sql.InInts(card.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.card) > 0 {
		for _, id := range ids {
			eid := keys(uuo.card)[0]
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
			if int(affected) < len(uuo.card) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"card\" %v already connected to a different \"User\"", keys(uuo.card))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
