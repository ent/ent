// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// PetUpdate is the builder for updating Pet entities.
type PetUpdate struct {
	config
	name         *string
	team         map[string]struct{}
	owner        map[string]struct{}
	clearedTeam  bool
	clearedOwner bool
	predicates   []predicate.Pet
}

// Where adds a new predicate for the builder.
func (pu *PetUpdate) Where(ps ...predicate.Pet) *PetUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetName sets the name field.
func (pu *PetUpdate) SetName(s string) *PetUpdate {
	pu.name = &s
	return pu
}

// SetTeamID sets the team edge to User by id.
func (pu *PetUpdate) SetTeamID(id string) *PetUpdate {
	if pu.team == nil {
		pu.team = make(map[string]struct{})
	}
	pu.team[id] = struct{}{}
	return pu
}

// SetNillableTeamID sets the team edge to User by id if the given value is not nil.
func (pu *PetUpdate) SetNillableTeamID(id *string) *PetUpdate {
	if id != nil {
		pu = pu.SetTeamID(*id)
	}
	return pu
}

// SetTeam sets the team edge to User.
func (pu *PetUpdate) SetTeam(u *User) *PetUpdate {
	return pu.SetTeamID(u.ID)
}

// SetOwnerID sets the owner edge to User by id.
func (pu *PetUpdate) SetOwnerID(id string) *PetUpdate {
	if pu.owner == nil {
		pu.owner = make(map[string]struct{})
	}
	pu.owner[id] = struct{}{}
	return pu
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pu *PetUpdate) SetNillableOwnerID(id *string) *PetUpdate {
	if id != nil {
		pu = pu.SetOwnerID(*id)
	}
	return pu
}

// SetOwner sets the owner edge to User.
func (pu *PetUpdate) SetOwner(u *User) *PetUpdate {
	return pu.SetOwnerID(u.ID)
}

// ClearTeam clears the team edge to User.
func (pu *PetUpdate) ClearTeam() *PetUpdate {
	pu.clearedTeam = true
	return pu
}

// ClearOwner clears the owner edge to User.
func (pu *PetUpdate) ClearOwner() *PetUpdate {
	pu.clearedOwner = true
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PetUpdate) Save(ctx context.Context) (int, error) {
	if len(pu.team) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"team\"")
	}
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
	if value := pu.name; value != nil {
		updater.Set(pet.FieldName, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if pu.clearedTeam {
		query, args := builder.Update(pet.TeamTable).
			SetNull(pet.TeamColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(pu.team) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(pu.team)[0])
			if serr != nil {
				return 0, rollback(tx, err)
			}
			query, args := builder.Update(pet.TeamTable).
				Set(pet.TeamColumn, eid).
				Where(sql.EQ(pet.FieldID, id).And().IsNull(pet.TeamColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(pu.team) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"team\" %v already connected to a different \"Pet\"", keys(pu.team))})
			}
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
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
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
	id           string
	name         *string
	team         map[string]struct{}
	owner        map[string]struct{}
	clearedTeam  bool
	clearedOwner bool
}

// SetName sets the name field.
func (puo *PetUpdateOne) SetName(s string) *PetUpdateOne {
	puo.name = &s
	return puo
}

// SetTeamID sets the team edge to User by id.
func (puo *PetUpdateOne) SetTeamID(id string) *PetUpdateOne {
	if puo.team == nil {
		puo.team = make(map[string]struct{})
	}
	puo.team[id] = struct{}{}
	return puo
}

// SetNillableTeamID sets the team edge to User by id if the given value is not nil.
func (puo *PetUpdateOne) SetNillableTeamID(id *string) *PetUpdateOne {
	if id != nil {
		puo = puo.SetTeamID(*id)
	}
	return puo
}

// SetTeam sets the team edge to User.
func (puo *PetUpdateOne) SetTeam(u *User) *PetUpdateOne {
	return puo.SetTeamID(u.ID)
}

// SetOwnerID sets the owner edge to User by id.
func (puo *PetUpdateOne) SetOwnerID(id string) *PetUpdateOne {
	if puo.owner == nil {
		puo.owner = make(map[string]struct{})
	}
	puo.owner[id] = struct{}{}
	return puo
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (puo *PetUpdateOne) SetNillableOwnerID(id *string) *PetUpdateOne {
	if id != nil {
		puo = puo.SetOwnerID(*id)
	}
	return puo
}

// SetOwner sets the owner edge to User.
func (puo *PetUpdateOne) SetOwner(u *User) *PetUpdateOne {
	return puo.SetOwnerID(u.ID)
}

// ClearTeam clears the team edge to User.
func (puo *PetUpdateOne) ClearTeam() *PetUpdateOne {
	puo.clearedTeam = true
	return puo
}

// ClearOwner clears the owner edge to User.
func (puo *PetUpdateOne) ClearOwner() *PetUpdateOne {
	puo.clearedOwner = true
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PetUpdateOne) Save(ctx context.Context) (*Pet, error) {
	if len(puo.team) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"team\"")
	}
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
		id = pe.id()
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
	if value := puo.name; value != nil {
		updater.Set(pet.FieldName, *value)
		pe.Name = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if puo.clearedTeam {
		query, args := builder.Update(pet.TeamTable).
			SetNull(pet.TeamColumn).
			Where(sql.InInts(user.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(puo.team) > 0 {
		for _, id := range ids {
			eid, serr := strconv.Atoi(keys(puo.team)[0])
			if serr != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(pet.TeamTable).
				Set(pet.TeamColumn, eid).
				Where(sql.EQ(pet.FieldID, id).And().IsNull(pet.TeamColumn)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(puo.team) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"team\" %v already connected to a different \"Pet\"", keys(puo.team))})
			}
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
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
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
