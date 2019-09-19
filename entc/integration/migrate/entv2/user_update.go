// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age           *int
	addage        *int
	name          *string
	phone         *string
	buffer        *[]byte
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
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

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	uu.addage = &i
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.phone = &s
	return uu
}

// SetBuffer sets the buffer field.
func (uu *UserUpdate) SetBuffer(b []byte) *UserUpdate {
	uu.buffer = &b
	return uu
}

// SetTitle sets the title field.
func (uu *UserUpdate) SetTitle(s string) *UserUpdate {
	uu.title = &s
	return uu
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uu *UserUpdate) SetNillableTitle(s *string) *UserUpdate {
	if s != nil {
		uu.SetTitle(*s)
	}
	return uu
}

// SetNewName sets the new_name field.
func (uu *UserUpdate) SetNewName(s string) *UserUpdate {
	uu.new_name = &s
	return uu
}

// SetNillableNewName sets the new_name field if the given value is not nil.
func (uu *UserUpdate) SetNillableNewName(s *string) *UserUpdate {
	if s != nil {
		uu.SetNewName(*s)
	}
	return uu
}

// ClearNewName clears the value of new_name.
func (uu *UserUpdate) ClearNewName() *UserUpdate {
	uu.new_name = nil
	uu.clearnew_name = true
	return uu
}

// SetBlob sets the blob field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.blob = &b
	return uu
}

// ClearBlob clears the value of blob.
func (uu *UserUpdate) ClearBlob() *UserUpdate {
	uu.blob = nil
	uu.clearblob = true
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
			return 0, fmt.Errorf("entv2: failed reading id: %v", err)
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
	if value := uu.age; value != nil {
		builder.Set(user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		builder.Add(user.FieldAge, *value)
	}
	if value := uu.name; value != nil {
		builder.Set(user.FieldName, *value)
	}
	if value := uu.phone; value != nil {
		builder.Set(user.FieldPhone, *value)
	}
	if value := uu.buffer; value != nil {
		builder.Set(user.FieldBuffer, *value)
	}
	if value := uu.title; value != nil {
		builder.Set(user.FieldTitle, *value)
	}
	if value := uu.new_name; value != nil {
		builder.Set(user.FieldNewName, *value)
	}
	if uu.clearnew_name {
		builder.SetNull(user.FieldNewName)
	}
	if value := uu.blob; value != nil {
		builder.Set(user.FieldBlob, *value)
	}
	if uu.clearblob {
		builder.SetNull(user.FieldBlob)
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
	id            int
	age           *int
	addage        *int
	name          *string
	phone         *string
	buffer        *[]byte
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	uuo.addage = &i
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.phone = &s
	return uuo
}

// SetBuffer sets the buffer field.
func (uuo *UserUpdateOne) SetBuffer(b []byte) *UserUpdateOne {
	uuo.buffer = &b
	return uuo
}

// SetTitle sets the title field.
func (uuo *UserUpdateOne) SetTitle(s string) *UserUpdateOne {
	uuo.title = &s
	return uuo
}

// SetNillableTitle sets the title field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTitle(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetTitle(*s)
	}
	return uuo
}

// SetNewName sets the new_name field.
func (uuo *UserUpdateOne) SetNewName(s string) *UserUpdateOne {
	uuo.new_name = &s
	return uuo
}

// SetNillableNewName sets the new_name field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNewName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNewName(*s)
	}
	return uuo
}

// ClearNewName clears the value of new_name.
func (uuo *UserUpdateOne) ClearNewName() *UserUpdateOne {
	uuo.new_name = nil
	uuo.clearnew_name = true
	return uuo
}

// SetBlob sets the blob field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.blob = &b
	return uuo
}

// ClearBlob clears the value of blob.
func (uuo *UserUpdateOne) ClearBlob() *UserUpdateOne {
	uuo.blob = nil
	uuo.clearblob = true
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
			return nil, fmt.Errorf("entv2: failed scanning row into User: %v", err)
		}
		id = u.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("entv2: User not found with id: %v", uuo.id)
	case n > 1:
		return nil, fmt.Errorf("entv2: more than one User with the same id: %v", uuo.id)
	}

	tx, err := uuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.age; value != nil {
		builder.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uuo.addage; value != nil {
		builder.Add(user.FieldAge, *value)
		u.Age += *value
	}
	if value := uuo.name; value != nil {
		builder.Set(user.FieldName, *value)
		u.Name = *value
	}
	if value := uuo.phone; value != nil {
		builder.Set(user.FieldPhone, *value)
		u.Phone = *value
	}
	if value := uuo.buffer; value != nil {
		builder.Set(user.FieldBuffer, *value)
		u.Buffer = *value
	}
	if value := uuo.title; value != nil {
		builder.Set(user.FieldTitle, *value)
		u.Title = *value
	}
	if value := uuo.new_name; value != nil {
		builder.Set(user.FieldNewName, *value)
		u.NewName = *value
	}
	if uuo.clearnew_name {
		var value string
		u.NewName = value
		builder.SetNull(user.FieldNewName)
	}
	if value := uuo.blob; value != nil {
		builder.Set(user.FieldBlob, *value)
		u.Blob = *value
	}
	if uuo.clearblob {
		var value []byte
		u.Blob = value
		builder.SetNull(user.FieldBlob)
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
