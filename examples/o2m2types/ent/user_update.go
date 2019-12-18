// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/o2m2types/ent/pet"
	"github.com/facebookincubator/ent/examples/o2m2types/ent/predicate"
	"github.com/facebookincubator/ent/examples/o2m2types/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age         *int
	addage      *int
	name        *string
	pets        map[int]struct{}
	removedPets map[int]struct{}
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

// AddPetIDs adds the pets edge to Pet by ids.
func (uu *UserUpdate) AddPetIDs(ids ...int) *UserUpdate {
	if uu.pets == nil {
		uu.pets = make(map[int]struct{})
	}
	for i := range ids {
		uu.pets[ids[i]] = struct{}{}
	}
	return uu
}

// AddPets adds the pets edges to Pet.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (uu *UserUpdate) RemovePetIDs(ids ...int) *UserUpdate {
	if uu.removedPets == nil {
		uu.removedPets = make(map[int]struct{})
	}
	for i := range ids {
		uu.removedPets[ids[i]] = struct{}{}
	}
	return uu
}

// RemovePets removes pets edges to Pet.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
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
		updater = builder.Update(user.Table)
	)
	updater = updater.Where(sql.InInts(user.FieldID, ids...))
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
	if len(uu.removedPets) > 0 {
		eids := make([]int, len(uu.removedPets))
		for eid := range uu.removedPets {
			eids = append(eids, eid)
		}
		query, args := builder.Update(user.PetsTable).
			SetNull(user.PetsColumn).
			Where(sql.InInts(user.PetsColumn, ids...)).
			Where(sql.InInts(pet.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(uu.pets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uu.pets {
				p.Or().EQ(pet.FieldID, eid)
			}
			query, args := builder.Update(user.PetsTable).
				Set(user.PetsColumn, id).
				Where(sql.And(p, sql.IsNull(user.PetsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(uu.pets) {
				return 0, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uu.pets))})
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
	addage      *int
	name        *string
	pets        map[int]struct{}
	removedPets map[int]struct{}
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

// AddPetIDs adds the pets edge to Pet by ids.
func (uuo *UserUpdateOne) AddPetIDs(ids ...int) *UserUpdateOne {
	if uuo.pets == nil {
		uuo.pets = make(map[int]struct{})
	}
	for i := range ids {
		uuo.pets[ids[i]] = struct{}{}
	}
	return uuo
}

// AddPets adds the pets edges to Pet.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...int) *UserUpdateOne {
	if uuo.removedPets == nil {
		uuo.removedPets = make(map[int]struct{})
	}
	for i := range ids {
		uuo.removedPets[ids[i]] = struct{}{}
	}
	return uuo
}

// RemovePets removes pets edges to Pet.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
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
		updater = builder.Update(user.Table)
	)
	updater = updater.Where(sql.InInts(user.FieldID, ids...))
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
	if len(uuo.removedPets) > 0 {
		eids := make([]int, len(uuo.removedPets))
		for eid := range uuo.removedPets {
			eids = append(eids, eid)
		}
		query, args := builder.Update(user.PetsTable).
			SetNull(user.PetsColumn).
			Where(sql.InInts(user.PetsColumn, ids...)).
			Where(sql.InInts(pet.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(uuo.pets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range uuo.pets {
				p.Or().EQ(pet.FieldID, eid)
			}
			query, args := builder.Update(user.PetsTable).
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
			if int(affected) < len(uuo.pets) {
				return nil, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uuo.pets))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
