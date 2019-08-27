// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/o2obidi/ent/predicate"
	"github.com/facebookincubator/ent/examples/o2obidi/ent/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age           *int
	name          *string
	spouse        map[int]struct{}
	clearedSpouse bool
	predicates    []predicate.User
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

// SetSpouseID sets the spouse edge to User by id.
func (uu *UserUpdate) SetSpouseID(id int) *UserUpdate {
	if uu.spouse == nil {
		uu.spouse = make(map[int]struct{})
	}
	uu.spouse[id] = struct{}{}
	return uu
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *int) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the spouse edge to User.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// ClearSpouse clears the spouse edge to User.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.clearedSpouse = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if len(uu.spouse) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
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
	if uu.clearedSpouse {
		query, args := sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
		query, args = sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.SpouseColumn, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.spouse) > 0 {
		if n := len(ids); n > 1 {
			return 0, rollback(tx, fmt.Errorf("ent: can't link O2O edge \"spouse\" to %d vertices (> 1)", n))
		}
		for eid := range uu.spouse {
			query, args := sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, ids[0])).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			query, args = sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, ids[0]).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.spouse) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
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
	id            int
	age           *int
	name          *string
	spouse        map[int]struct{}
	clearedSpouse bool
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

// SetSpouseID sets the spouse edge to User by id.
func (uuo *UserUpdateOne) SetSpouseID(id int) *UserUpdateOne {
	if uuo.spouse == nil {
		uuo.spouse = make(map[int]struct{})
	}
	uuo.spouse[id] = struct{}{}
	return uuo
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *int) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the spouse edge to User.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// ClearSpouse clears the spouse edge to User.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.clearedSpouse = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if len(uuo.spouse) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
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
	if uuo.clearedSpouse {
		query, args := sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		query, args = sql.Update(user.SpouseTable).
			SetNull(user.SpouseColumn).
			Where(sql.InInts(user.SpouseColumn, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.spouse) > 0 {
		if n := len(ids); n > 1 {
			return nil, rollback(tx, fmt.Errorf("ent: can't link O2O edge \"spouse\" to %d vertices (> 1)", n))
		}
		for eid := range uuo.spouse {
			query, args := sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, ids[0])).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			query, args = sql.Update(user.SpouseTable).
				Set(user.SpouseColumn, ids[0]).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(uuo.spouse) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
