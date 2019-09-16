// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"

	"github.com/facebookincubator/ent/dialect/sql"
)

// StreetUpdate is the builder for updating Street entities.
type StreetUpdate struct {
	config
	name        *string
	city        map[int]struct{}
	clearedCity bool
	predicates  []predicate.Street
}

// Where adds a new predicate for the builder.
func (su *StreetUpdate) Where(ps ...predicate.Street) *StreetUpdate {
	su.predicates = append(su.predicates, ps...)
	return su
}

// SetName sets the name field.
func (su *StreetUpdate) SetName(s string) *StreetUpdate {
	su.name = &s
	return su
}

// SetCityID sets the city edge to City by id.
func (su *StreetUpdate) SetCityID(id int) *StreetUpdate {
	if su.city == nil {
		su.city = make(map[int]struct{})
	}
	su.city[id] = struct{}{}
	return su
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (su *StreetUpdate) SetNillableCityID(id *int) *StreetUpdate {
	if id != nil {
		su = su.SetCityID(*id)
	}
	return su
}

// SetCity sets the city edge to City.
func (su *StreetUpdate) SetCity(c *City) *StreetUpdate {
	return su.SetCityID(c.ID)
}

// ClearCity clears the city edge to City.
func (su *StreetUpdate) ClearCity() *StreetUpdate {
	su.clearedCity = true
	return su
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (su *StreetUpdate) Save(ctx context.Context) (int, error) {
	if len(su.city) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"city\"")
	}
	return su.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (su *StreetUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *StreetUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *StreetUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *StreetUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(street.FieldID).From(sql.Table(street.Table))
	for _, p := range su.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = su.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := su.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(street.Table).Where(sql.InInts(street.FieldID, ids...))
	)
	if value := su.name; value != nil {
		update = true
		builder.Set(street.FieldName, *value)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if su.clearedCity {
		query, args := sql.Update(street.CityTable).
			SetNull(street.CityColumn).
			Where(sql.InInts(city.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(su.city) > 0 {
		for eid := range su.city {
			query, args := sql.Update(street.CityTable).
				Set(street.CityColumn, eid).
				Where(sql.InInts(street.FieldID, ids...)).
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

// StreetUpdateOne is the builder for updating a single Street entity.
type StreetUpdateOne struct {
	config
	id          int
	name        *string
	city        map[int]struct{}
	clearedCity bool
}

// SetName sets the name field.
func (suo *StreetUpdateOne) SetName(s string) *StreetUpdateOne {
	suo.name = &s
	return suo
}

// SetCityID sets the city edge to City by id.
func (suo *StreetUpdateOne) SetCityID(id int) *StreetUpdateOne {
	if suo.city == nil {
		suo.city = make(map[int]struct{})
	}
	suo.city[id] = struct{}{}
	return suo
}

// SetNillableCityID sets the city edge to City by id if the given value is not nil.
func (suo *StreetUpdateOne) SetNillableCityID(id *int) *StreetUpdateOne {
	if id != nil {
		suo = suo.SetCityID(*id)
	}
	return suo
}

// SetCity sets the city edge to City.
func (suo *StreetUpdateOne) SetCity(c *City) *StreetUpdateOne {
	return suo.SetCityID(c.ID)
}

// ClearCity clears the city edge to City.
func (suo *StreetUpdateOne) ClearCity() *StreetUpdateOne {
	suo.clearedCity = true
	return suo
}

// Save executes the query and returns the updated entity.
func (suo *StreetUpdateOne) Save(ctx context.Context) (*Street, error) {
	if len(suo.city) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"city\"")
	}
	return suo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (suo *StreetUpdateOne) SaveX(ctx context.Context) *Street {
	s, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return s
}

// Exec executes the query on the entity.
func (suo *StreetUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *StreetUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *StreetUpdateOne) sqlSave(ctx context.Context) (s *Street, err error) {
	selector := sql.Select(street.Columns...).From(sql.Table(street.Table))
	street.ID(suo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = suo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		s = &Street{config: suo.config}
		if err := s.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Street: %v", err)
		}
		id = s.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: Street not found with id: %v", suo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Street with the same id: %v", suo.id)
	}

	tx, err := suo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(street.Table).Where(sql.InInts(street.FieldID, ids...))
	)
	if value := suo.name; value != nil {
		update = true
		builder.Set(street.FieldName, *value)
		s.Name = *value
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if suo.clearedCity {
		query, args := sql.Update(street.CityTable).
			SetNull(street.CityColumn).
			Where(sql.InInts(city.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(suo.city) > 0 {
		for eid := range suo.city {
			query, args := sql.Update(street.CityTable).
				Set(street.CityColumn, eid).
				Where(sql.InInts(street.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return s, nil
}
