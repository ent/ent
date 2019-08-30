// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/examples/edgeindex/ent/city"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/predicate"
	"github.com/facebookincubator/ent/examples/edgeindex/ent/street"

	"github.com/facebookincubator/ent/dialect/sql"
)

// CityUpdate is the builder for updating City entities.
type CityUpdate struct {
	config
	name           *string
	streets        map[int]struct{}
	removedStreets map[int]struct{}
	predicates     []predicate.City
}

// Where adds a new predicate for the builder.
func (cu *CityUpdate) Where(ps ...predicate.City) *CityUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the name field.
func (cu *CityUpdate) SetName(s string) *CityUpdate {
	cu.name = &s
	return cu
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cu *CityUpdate) AddStreetIDs(ids ...int) *CityUpdate {
	if cu.streets == nil {
		cu.streets = make(map[int]struct{})
	}
	for i := range ids {
		cu.streets[ids[i]] = struct{}{}
	}
	return cu
}

// AddStreets adds the streets edges to Street.
func (cu *CityUpdate) AddStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.AddStreetIDs(ids...)
}

// RemoveStreetIDs removes the streets edge to Street by ids.
func (cu *CityUpdate) RemoveStreetIDs(ids ...int) *CityUpdate {
	if cu.removedStreets == nil {
		cu.removedStreets = make(map[int]struct{})
	}
	for i := range ids {
		cu.removedStreets[ids[i]] = struct{}{}
	}
	return cu
}

// RemoveStreets removes streets edges to Street.
func (cu *CityUpdate) RemoveStreets(s ...*Street) *CityUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cu.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CityUpdate) Save(ctx context.Context) (int, error) {
	return cu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CityUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CityUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CityUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	selector := sql.Select(city.FieldID).From(sql.Table(city.Table))
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
		update  bool
		res     sql.Result
		builder = sql.Update(city.Table).Where(sql.InInts(city.FieldID, ids...))
	)
	if cu.name != nil {
		update = true
		builder.Set(city.FieldName, *cu.name)
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(cu.removedStreets) > 0 {
		eids := make([]int, len(cu.removedStreets))
		for eid := range cu.removedStreets {
			eids = append(eids, eid)
		}
		query, args := sql.Update(city.StreetsTable).
			SetNull(city.StreetsColumn).
			Where(sql.InInts(city.StreetsColumn, ids...)).
			Where(sql.InInts(street.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(cu.streets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range cu.streets {
				p.Or().EQ(street.FieldID, eid)
			}
			query, args := sql.Update(city.StreetsTable).
				Set(city.StreetsColumn, id).
				Where(sql.And(p, sql.IsNull(city.StreetsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(cu.streets) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"streets\" %v already connected to a different \"City\"", keys(cu.streets))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// CityUpdateOne is the builder for updating a single City entity.
type CityUpdateOne struct {
	config
	id             int
	name           *string
	streets        map[int]struct{}
	removedStreets map[int]struct{}
}

// SetName sets the name field.
func (cuo *CityUpdateOne) SetName(s string) *CityUpdateOne {
	cuo.name = &s
	return cuo
}

// AddStreetIDs adds the streets edge to Street by ids.
func (cuo *CityUpdateOne) AddStreetIDs(ids ...int) *CityUpdateOne {
	if cuo.streets == nil {
		cuo.streets = make(map[int]struct{})
	}
	for i := range ids {
		cuo.streets[ids[i]] = struct{}{}
	}
	return cuo
}

// AddStreets adds the streets edges to Street.
func (cuo *CityUpdateOne) AddStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.AddStreetIDs(ids...)
}

// RemoveStreetIDs removes the streets edge to Street by ids.
func (cuo *CityUpdateOne) RemoveStreetIDs(ids ...int) *CityUpdateOne {
	if cuo.removedStreets == nil {
		cuo.removedStreets = make(map[int]struct{})
	}
	for i := range ids {
		cuo.removedStreets[ids[i]] = struct{}{}
	}
	return cuo
}

// RemoveStreets removes streets edges to Street.
func (cuo *CityUpdateOne) RemoveStreets(s ...*Street) *CityUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return cuo.RemoveStreetIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CityUpdateOne) Save(ctx context.Context) (*City, error) {
	return cuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CityUpdateOne) SaveX(ctx context.Context) *City {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CityUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CityUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CityUpdateOne) sqlSave(ctx context.Context) (c *City, err error) {
	selector := sql.Select(city.Columns...).From(sql.Table(city.Table))
	city.ID(cuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = cuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		c = &City{config: cuo.config}
		if err := c.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into City: %v", err)
		}
		id = c.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("ent: City not found with id: %v", cuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one City with the same id: %v", cuo.id)
	}

	tx, err := cuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(city.Table).Where(sql.InInts(city.FieldID, ids...))
	)
	if cuo.name != nil {
		update = true
		builder.Set(city.FieldName, *cuo.name)
		c.Name = *cuo.name
	}
	if update {
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(cuo.removedStreets) > 0 {
		eids := make([]int, len(cuo.removedStreets))
		for eid := range cuo.removedStreets {
			eids = append(eids, eid)
		}
		query, args := sql.Update(city.StreetsTable).
			SetNull(city.StreetsColumn).
			Where(sql.InInts(city.StreetsColumn, ids...)).
			Where(sql.InInts(street.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(cuo.streets) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range cuo.streets {
				p.Or().EQ(street.FieldID, eid)
			}
			query, args := sql.Update(city.StreetsTable).
				Set(city.StreetsColumn, id).
				Where(sql.And(p, sql.IsNull(city.StreetsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(cuo.streets) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"streets\" %v already connected to a different \"City\"", keys(cuo.streets))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return c, nil
}
