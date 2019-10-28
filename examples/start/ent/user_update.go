// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/start/ent/car"
	"github.com/facebookincubator/ent/examples/start/ent/predicate"
	"github.com/facebookincubator/ent/examples/start/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age           *int
	addage        *int
	name          *string
	cars          map[int]struct{}
	groups        map[int]struct{}
	removedCars   map[int]struct{}
	removedGroups map[int]struct{}
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
	uu.addage = nil
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	if uu.addage == nil {
		uu.addage = &i
	} else {
		*uu.addage += i
	}
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetNillableName sets the name field if the given value is not nil.
func (uu *UserUpdate) SetNillableName(s *string) *UserUpdate {
	if s != nil {
		uu.SetName(*s)
	}
	return uu
}

// AddCarIDs adds the cars edge to Car by ids.
func (uu *UserUpdate) AddCarIDs(ids ...int) *UserUpdate {
	if uu.cars == nil {
		uu.cars = make(map[int]struct{})
	}
	for i := range ids {
		uu.cars[ids[i]] = struct{}{}
	}
	return uu
}

// AddCars adds the cars edges to Car.
func (uu *UserUpdate) AddCars(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.AddCarIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uu *UserUpdate) AddGroupIDs(ids ...int) *UserUpdate {
	if uu.groups == nil {
		uu.groups = make(map[int]struct{})
	}
	for i := range ids {
		uu.groups[ids[i]] = struct{}{}
	}
	return uu
}

// AddGroups adds the groups edges to Group.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// RemoveCarIDs removes the cars edge to Car by ids.
func (uu *UserUpdate) RemoveCarIDs(ids ...int) *UserUpdate {
	if uu.removedCars == nil {
		uu.removedCars = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedCars[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveCars removes cars edges to Car.
func (uu *UserUpdate) RemoveCars(c ...*Car) *UserUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uu.RemoveCarIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uu *UserUpdate) RemoveGroupIDs(ids ...int) *UserUpdate {
	if uu.removedGroups == nil {
		uu.removedGroups = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedGroups[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveGroups removes groups edges to Group.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.age != nil {
		if err := user.AgeValidator(*uu.age); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
		}
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
	var (
		builder  = sql.Dialect(uu.driver.Dialect())
		selector = builder.Select(user.FieldID).From(builder.Table(user.Table))
	)
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
		res     sql.Result
		updater = builder.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uu.age; value != nil {
		updater.Set(user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		updater.Add(user.FieldAge, *value)
	}
	if value := uu.name; value != nil {
		updater.Set(user.FieldName, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.removedCars) > 0 {
		eids := make([]int, len(uu.removedCars))
		for eid := range uu.removedCars {
			eids = append(eids, eid)
		}
		query, args := builder.Update(user.CarsTable).
			SetNull(user.CarsColumn).
			Where(sql.InInts(user.CarsColumn, ids...)).
			Where(sql.InInts(car.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.cars) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uu.cars {
				p.Or().EQ(car.FieldID, eid)
			}
			query, args := builder.Update(user.CarsTable).
				Set(user.CarsColumn, id).
				Where(sql.And(p, sql.IsNull(user.CarsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.cars) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"cars\" %v already connected to a different \"User\"", keys(uu.cars))})
			}
		}
	}
	if len(uu.removedGroups) > 0 {
		eids := make([]int, len(uu.removedGroups))
		for eid := range uu.removedGroups {
			eids = append(eids, eid)
		}
		query, args := builder.Delete(user.GroupsTable).
			Where(sql.InInts(user.GroupsPrimaryKey[1], ids...)).
			Where(sql.InInts(user.GroupsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.groups) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uu.groups {
				values = append(values, []int{id, eid})
			}
		}
		builder := builder.Insert(user.GroupsTable).
			Columns(user.GroupsPrimaryKey[1], user.GroupsPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
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
	addage        *int
	name          *string
	cars          map[int]struct{}
	groups        map[int]struct{}
	removedCars   map[int]struct{}
	removedGroups map[int]struct{}
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	uuo.addage = nil
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	if uuo.addage == nil {
		uuo.addage = &i
	} else {
		*uuo.addage += i
	}
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetNillableName sets the name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetName(*s)
	}
	return uuo
}

// AddCarIDs adds the cars edge to Car by ids.
func (uuo *UserUpdateOne) AddCarIDs(ids ...int) *UserUpdateOne {
	if uuo.cars == nil {
		uuo.cars = make(map[int]struct{})
	}
	for i := range ids {
		uuo.cars[ids[i]] = struct{}{}
	}
	return uuo
}

// AddCars adds the cars edges to Car.
func (uuo *UserUpdateOne) AddCars(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.AddCarIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...int) *UserUpdateOne {
	if uuo.groups == nil {
		uuo.groups = make(map[int]struct{})
	}
	for i := range ids {
		uuo.groups[ids[i]] = struct{}{}
	}
	return uuo
}

// AddGroups adds the groups edges to Group.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// RemoveCarIDs removes the cars edge to Car by ids.
func (uuo *UserUpdateOne) RemoveCarIDs(ids ...int) *UserUpdateOne {
	if uuo.removedCars == nil {
		uuo.removedCars = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedCars[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveCars removes cars edges to Car.
func (uuo *UserUpdateOne) RemoveCars(c ...*Car) *UserUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return uuo.RemoveCarIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...int) *UserUpdateOne {
	if uuo.removedGroups == nil {
		uuo.removedGroups = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedGroups[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveGroups removes groups edges to Group.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]int, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.age != nil {
		if err := user.AgeValidator(*uuo.age); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"age\": %v", err)
		}
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
	var (
		builder  = sql.Dialect(uuo.driver.Dialect())
		selector = builder.Select(user.Columns...).From(builder.Table(user.Table))
	)
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
		return nil, &ErrNotFound{fmt.Sprintf("User with id: %v", uuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.age; value != nil {
		updater.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uuo.addage; value != nil {
		updater.Add(user.FieldAge, *value)
		u.Age += *value
	}
	if value := uuo.name; value != nil {
		updater.Set(user.FieldName, *value)
		u.Name = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.removedCars) > 0 {
		eids := make([]int, len(uuo.removedCars))
		for eid := range uuo.removedCars {
			eids = append(eids, eid)
		}
		query, args := builder.Update(user.CarsTable).
			SetNull(user.CarsColumn).
			Where(sql.InInts(user.CarsColumn, ids...)).
			Where(sql.InInts(car.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.cars) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uuo.cars {
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
			if int(affected) < len(uuo.cars) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"cars\" %v already connected to a different \"User\"", keys(uuo.cars))})
			}
		}
	}
	if len(uuo.removedGroups) > 0 {
		eids := make([]int, len(uuo.removedGroups))
		for eid := range uuo.removedGroups {
			eids = append(eids, eid)
		}
		query, args := builder.Delete(user.GroupsTable).
			Where(sql.InInts(user.GroupsPrimaryKey[1], ids...)).
			Where(sql.InInts(user.GroupsPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.groups) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range uuo.groups {
				values = append(values, []int{id, eid})
			}
		}
		builder := builder.Insert(user.GroupsTable).
			Columns(user.GroupsPrimaryKey[1], user.GroupsPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
