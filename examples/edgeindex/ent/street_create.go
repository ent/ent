// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"
)

// StreetCreate is the builder for creating a Street entity.
type StreetCreate struct {
	config
	name *string
	city map[int]struct{}
}

// SetName sets the name field.
func (sc *StreetCreate) SetName(s string) *StreetCreate {
	sc.name = &s
	return sc
}

// SetCityID sets the city edge to City by id.
func (sc *StreetCreate) SetCityID(id int) *StreetCreate {
	if sc.city == nil {
		sc.city = make(map[int]struct{})
	}
	sc.city[id] = struct{}{}
	return sc
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (sc *StreetCreate) SetNillableCityID(id *int) *StreetCreate {
	if id != nil {
		sc = sc.SetCityID(*id)
	}
	return sc
}

// SetCity sets the city edge to City.
func (sc *StreetCreate) SetCity(c *City) *StreetCreate {
	return sc.SetCityID(c.ID)
}

// Save creates the Street in the database.
func (sc *StreetCreate) Save(ctx context.Context) (*Street, error) {
	if sc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if len(sc.city) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"city\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *StreetCreate) SaveX(ctx context.Context) *Street {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *StreetCreate) sqlSave(ctx context.Context) (*Street, error) {
	var (
		res sql.Result
		s   = &Street{config: sc.config}
	)
	tx, err := sc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	builder := sql.Insert(street.Table).Default(sc.driver.Dialect())
	if value := sc.name; value != nil {
		builder.Set(street.FieldName, *value)
		s.Name = *value
	}
	query, args := builder.Query()
	if err := tx.Exec(ctx, query, args, &res); err != nil {
		return nil, rollback(tx, err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		return nil, rollback(tx, err)
	}
	s.ID = int(id)
	if len(sc.city) > 0 {
		for eid := range sc.city {
			query, args := sql.Update(street.CityTable).
				Set(street.CityColumn, eid).
				Where(sql.EQ(street.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return s, nil
}
