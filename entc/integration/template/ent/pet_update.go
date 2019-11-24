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
	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/template/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/template/ent/user"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	age              *int
	addage           *int
	licensed_at      *time.Time
	clearlicensed_at bool
	owner            map[int]struct{}
	clearedOwner     bool
	predicates       []predicate.Pet
}

// Where adds a new predicate for the builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetAge sets the age field.
func (pu *PetUpdate) SetAge(i int) *PetUpdate {
	pu.age = &i
	pu.addage = nil
	return pu
}

// AddAge adds i to age.
func (pu *PetUpdate) AddAge(i int) *PetUpdate {
	if pu.addage == nil {
		pu.addage = &i
	} else {
		*pu.addage += i
	}
	return pu
}

// SetLicensedAt sets the licensed_at field.
func (pu *PetUpdate) SetLicensedAt(t time.Time) *PetUpdate {
	pu.licensed_at = &t
	return pu
}

// SetNillableLicensedAt sets the licensed_at field if the given value is not nil.
func (pu *PetUpdate) SetNillableLicensedAt(t *time.Time) *PetUpdate {
	if t != nil {
		pu.SetLicensedAt(*t)
	}
	return pu
}

// ClearLicensedAt clears the value of licensed_at.
func (pu *PetUpdate) ClearLicensedAt() *PetUpdate {
	pu.licensed_at = nil
	pu.clearlicensed_at = true
	return pu
}

// SetOwnerID sets the owner edge to User by id.
func (pu *PetUpdate) SetOwnerID(id int) *PetUpdate {
	if pu.owner == nil {
		pu.owner = make(map[int]struct{})
	}
	pu.owner[id] = struct{}{}
	return pu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *int) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to User.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.clearedOwner = true
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	if len(pu.owner) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PetUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PetUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PetUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(pu.driver.Dialect())
		selector = builder.Select(pet.FieldID).From(builder.Table(pet.Table))
	)
	for _, p := range pu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = pu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := pu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(pet.Table)
	)
	updater = updater.Where(sql.InInts(pet.FieldID, ids...))
	if value := pu.age; value != nil {
		updater.Set(pet.FieldAge, *value)
	}
	if value := pu.addage; value != nil {
		updater.Add(pet.FieldAge, *value)
	}
	if value := pu.licensed_at; value != nil {
		updater.Set(pet.FieldLicensedAt, *value)
	}
	if pu.clearlicensed_at {
		updater.SetNull(pet.FieldLicensedAt)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if pu.clearedOwner {
		query, args := builder.Update(pet.OwnerTable).
			SetNull(pet.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.owner) > 0 {
		for eid := range pu.owner {
			query, args := builder.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.InInts(pet.FieldID, ids...)).
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

// PetUpdateOne is the builder for updating a single Pet entity.
type PetUpdateOne struct {
	config
	id               int
	age              *int
	addage           *int
	licensed_at      *time.Time
	clearlicensed_at bool
	owner            map[int]struct{}
	clearedOwner     bool
}

// SetAge sets the age field.
func (puo *PetUpdateOne) SetAge(i int) *PetUpdateOne {
	puo.age = &i
	puo.addage = nil
	return puo
}

// AddAge adds i to age.
func (puo *PetUpdateOne) AddAge(i int) *PetUpdateOne {
	if puo.addage == nil {
		puo.addage = &i
	} else {
		*puo.addage += i
	}
	return puo
}

// SetLicensedAt sets the licensed_at field.
func (puo *PetUpdateOne) SetLicensedAt(t time.Time) *PetUpdateOne {
	puo.licensed_at = &t
	return puo
}

// SetNillableLicensedAt sets the licensed_at field if the given value is not nil.
func (puo *PetUpdateOne) SetNillableLicensedAt(t *time.Time) *PetUpdateOne {
	if t != nil {
		puo.SetLicensedAt(*t)
	}
	return puo
}

// ClearLicensedAt clears the value of licensed_at.
func (puo *PetUpdateOne) ClearLicensedAt() *PetUpdateOne {
	puo.licensed_at = nil
	puo.clearlicensed_at = true
	return puo
}

// SetOwnerID sets the owner edge to User by id.
func (puo *PetUpdateOne) SetOwnerID(id int) *PetUpdateOne {
	if puo.owner == nil {
		puo.owner = make(map[int]struct{})
	}
	puo.owner[id] = struct{}{}
	return puo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *int) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to User.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// ClearOwner clears the owner edge to User.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.clearedOwner = true
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	if len(puo.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return puo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PetUpdateOne) SaveX(ctx context.Context) *Pet {
	pe, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pe
}

// Exec executes the query on the entity.
func (puo *PetUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PetUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PetUpdateOne) sqlSave(ctx context.Context) (pe *Pet, err error) {
	var (
		builder  = sql.Dialect(puo.driver.Dialect())
		selector = builder.Select(pet.Columns...).From(builder.Table(pet.Table))
	)
	pet.ID(puo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = puo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		pe = &Pet{config: puo.config}
		if err := pe.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Pet: %v", err)
		}
		id = pe.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Pet with id: %v", puo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Pet with the same id: %v", puo.id)
	}

	tx, err := puo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(pet.Table)
	)
	updater = updater.Where(sql.InInts(pet.FieldID, ids...))
	if value := puo.age; value != nil {
		updater.Set(pet.FieldAge, *value)
		pe.Age = *value
	}
	if value := puo.addage; value != nil {
		updater.Add(pet.FieldAge, *value)
		pe.Age += *value
	}
	if value := puo.licensed_at; value != nil {
		updater.Set(pet.FieldLicensedAt, *value)
		pe.LicensedAt = value
	}
	if puo.clearlicensed_at {
		pe.LicensedAt = nil
		updater.SetNull(pet.FieldLicensedAt)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if puo.clearedOwner {
		query, args := builder.Update(pet.OwnerTable).
			SetNull(pet.OwnerColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.owner) > 0 {
		for eid := range puo.owner {
			query, args := builder.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.InInts(pet.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
