// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/entc/integration/template/ent/pet"

	"github.com/facebookincubator/ent/dialect/sql"
)

// PetCreate is the builder for creating a Pet entity.
type PetCreate struct {
	config
	age   *int
	owner map[int]struct{}
}

// SetAge sets the age field.
func (pc *PetCreate) SetAge(i int) *PetCreate {
	pc.age = &i
	return pc
}

// SetOwnerID sets the owner edge to User by id.
func (pc *PetCreate) SetOwnerID(id int) *PetCreate {
	if pc.owner == nil {
		pc.owner = make(map[int]struct{})
	}
	pc.owner[id] = struct{}{}
	return pc
}

// SetNillableOwnerID sets the owner edge to User by id if the given value is not nil.
func (pc *PetCreate) SetNillableOwnerID(id *int) *PetCreate {
	if id != nil {
		pc = pc.SetOwnerID(*id)
	}
	return pc
}

// SetOwner sets the owner edge to User.
func (pc *PetCreate) SetOwner(u *User) *PetCreate {
	return pc.SetOwnerID(u.ID)
}

// Save creates the Pet in the database.
func (pc *PetCreate) Save(ctx context.Context) (*Pet, error) {
	if pc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if len(pc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return pc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PetCreate) SaveX(ctx context.Context) *Pet {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PetCreate) sqlSave(ctx context.Context) (*Pet, error) {
	var (
		res sql.Result
		pe  = &Pet{config: pc.config}
	)
	tx, err := pc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(pet.Table).Default(pc.driver.Dialect())
	if pc.age != nil {
		builder.Set(pet.FieldAge, *pc.age)
		pe.Age = *pc.age
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	pe.ID = int(id)
	if len(pc.owner) > 0 {
		for eid := range pc.owner {
			query, args := sql.Update(pet.OwnerTable).
				Set(pet.OwnerColumn, eid).
				Where(sql.EQ(pet.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return pe, nil
}
