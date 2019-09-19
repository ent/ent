// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/json/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/json/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	url          **url.URL
	clearurl     bool
	raw          *json.RawMessage
	clearraw     bool
	dirs         *[]http.Dir
	cleardirs    bool
	ints         *[]int
	clearints    bool
	floats       *[]float64
	clearfloats  bool
	strings      *[]string
	clearstrings bool
	predicates   []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetURL sets the url field.
func (uu *UserUpdate) SetURL(u *url.URL) *UserUpdate {
	uu.url = &u
	return uu
}

// ClearURL clears the value of url.
func (uu *UserUpdate) ClearURL() *UserUpdate {
	uu.url = nil
	uu.clearurl = true
	return uu
}

// SetRaw sets the raw field.
func (uu *UserUpdate) SetRaw(jm json.RawMessage) *UserUpdate {
	uu.raw = &jm
	return uu
}

// ClearRaw clears the value of raw.
func (uu *UserUpdate) ClearRaw() *UserUpdate {
	uu.raw = nil
	uu.clearraw = true
	return uu
}

// SetDirs sets the dirs field.
func (uu *UserUpdate) SetDirs(h []http.Dir) *UserUpdate {
	uu.dirs = &h
	return uu
}

// ClearDirs clears the value of dirs.
func (uu *UserUpdate) ClearDirs() *UserUpdate {
	uu.dirs = nil
	uu.cleardirs = true
	return uu
}

// SetInts sets the ints field.
func (uu *UserUpdate) SetInts(i []int) *UserUpdate {
	uu.ints = &i
	return uu
}

// ClearInts clears the value of ints.
func (uu *UserUpdate) ClearInts() *UserUpdate {
	uu.ints = nil
	uu.clearints = true
	return uu
}

// SetFloats sets the floats field.
func (uu *UserUpdate) SetFloats(f []float64) *UserUpdate {
	uu.floats = &f
	return uu
}

// ClearFloats clears the value of floats.
func (uu *UserUpdate) ClearFloats() *UserUpdate {
	uu.floats = nil
	uu.clearfloats = true
	return uu
}

// SetStrings sets the strings field.
func (uu *UserUpdate) SetStrings(s []string) *UserUpdate {
	uu.strings = &s
	return uu
}

// ClearStrings clears the value of strings.
func (uu *UserUpdate) ClearStrings() *UserUpdate {
	uu.strings = nil
	uu.clearstrings = true
	return uu
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
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uu.url; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldURL, buf)
	}
	if uu.clearurl {
		builder.SetNull(user.FieldURL)
	}
	if value := uu.raw; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldRaw, buf)
	}
	if uu.clearraw {
		builder.SetNull(user.FieldRaw)
	}
	if value := uu.dirs; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldDirs, buf)
	}
	if uu.cleardirs {
		builder.SetNull(user.FieldDirs)
	}
	if value := uu.ints; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldInts, buf)
	}
	if uu.clearints {
		builder.SetNull(user.FieldInts)
	}
	if value := uu.floats; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldFloats, buf)
	}
	if uu.clearfloats {
		builder.SetNull(user.FieldFloats)
	}
	if value := uu.strings; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return 0, err
		}
		builder.Set(user.FieldStrings, buf)
	}
	if uu.clearstrings {
		builder.SetNull(user.FieldStrings)
	}
	if !builder.Empty() {
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
	id           int
	url          **url.URL
	clearurl     bool
	raw          *json.RawMessage
	clearraw     bool
	dirs         *[]http.Dir
	cleardirs    bool
	ints         *[]int
	clearints    bool
	floats       *[]float64
	clearfloats  bool
	strings      *[]string
	clearstrings bool
}

// SetURL sets the url field.
func (uuo *UserUpdateOne) SetURL(u *url.URL) *UserUpdateOne {
	uuo.url = &u
	return uuo
}

// ClearURL clears the value of url.
func (uuo *UserUpdateOne) ClearURL() *UserUpdateOne {
	uuo.url = nil
	uuo.clearurl = true
	return uuo
}

// SetRaw sets the raw field.
func (uuo *UserUpdateOne) SetRaw(jm json.RawMessage) *UserUpdateOne {
	uuo.raw = &jm
	return uuo
}

// ClearRaw clears the value of raw.
func (uuo *UserUpdateOne) ClearRaw() *UserUpdateOne {
	uuo.raw = nil
	uuo.clearraw = true
	return uuo
}

// SetDirs sets the dirs field.
func (uuo *UserUpdateOne) SetDirs(h []http.Dir) *UserUpdateOne {
	uuo.dirs = &h
	return uuo
}

// ClearDirs clears the value of dirs.
func (uuo *UserUpdateOne) ClearDirs() *UserUpdateOne {
	uuo.dirs = nil
	uuo.cleardirs = true
	return uuo
}

// SetInts sets the ints field.
func (uuo *UserUpdateOne) SetInts(i []int) *UserUpdateOne {
	uuo.ints = &i
	return uuo
}

// ClearInts clears the value of ints.
func (uuo *UserUpdateOne) ClearInts() *UserUpdateOne {
	uuo.ints = nil
	uuo.clearints = true
	return uuo
}

// SetFloats sets the floats field.
func (uuo *UserUpdateOne) SetFloats(f []float64) *UserUpdateOne {
	uuo.floats = &f
	return uuo
}

// ClearFloats clears the value of floats.
func (uuo *UserUpdateOne) ClearFloats() *UserUpdateOne {
	uuo.floats = nil
	uuo.clearfloats = true
	return uuo
}

// SetStrings sets the strings field.
func (uuo *UserUpdateOne) SetStrings(s []string) *UserUpdateOne {
	uuo.strings = &s
	return uuo
}

// ClearStrings clears the value of strings.
func (uuo *UserUpdateOne) ClearStrings() *UserUpdateOne {
	uuo.strings = nil
	uuo.clearstrings = true
	return uuo
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
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.url; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldURL, buf)
		u.URL = *value
	}
	if uuo.clearurl {
		var value *url.URL
		u.URL = value
		builder.SetNull(user.FieldURL)
	}
	if value := uuo.raw; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldRaw, buf)
		u.Raw = *value
	}
	if uuo.clearraw {
		var value json.RawMessage
		u.Raw = value
		builder.SetNull(user.FieldRaw)
	}
	if value := uuo.dirs; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldDirs, buf)
		u.Dirs = *value
	}
	if uuo.cleardirs {
		var value []http.Dir
		u.Dirs = value
		builder.SetNull(user.FieldDirs)
	}
	if value := uuo.ints; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldInts, buf)
		u.Ints = *value
	}
	if uuo.clearints {
		var value []int
		u.Ints = value
		builder.SetNull(user.FieldInts)
	}
	if value := uuo.floats; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldFloats, buf)
		u.Floats = *value
	}
	if uuo.clearfloats {
		var value []float64
		u.Floats = value
		builder.SetNull(user.FieldFloats)
	}
	if value := uuo.strings; value != nil {
		buf, err := json.Marshal(*value)
		if err != nil {
			return nil, err
		}
		builder.Set(user.FieldStrings, buf)
		u.Strings = *value
	}
	if uuo.clearstrings {
		var value []string
		u.Strings = value
		builder.SetNull(user.FieldStrings)
	}
	if !builder.Empty() {
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
