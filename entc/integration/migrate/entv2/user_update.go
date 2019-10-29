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
	clearbuffer   bool
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
	state         *user.State
	clearstate    bool
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

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.phone = &s
	return uu
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetBuffer sets the buffer field.
func (uu *UserUpdate) SetBuffer(b []byte) *UserUpdate {
	uu.buffer = &b
	return uu
}

// ClearBuffer clears the value of buffer.
func (uu *UserUpdate) ClearBuffer() *UserUpdate {
	uu.buffer = nil
	uu.clearbuffer = true
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

// SetState sets the state field.
func (uu *UserUpdate) SetState(u user.State) *UserUpdate {
	uu.state = &u
	return uu
}

// SetNillableState sets the state field if the given value is not nil.
func (uu *UserUpdate) SetNillableState(u *user.State) *UserUpdate {
	if u != nil {
		uu.SetState(*u)
	}
	return uu
}

// ClearState clears the value of state.
func (uu *UserUpdate) ClearState() *UserUpdate {
	uu.state = nil
	uu.clearstate = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if uu.state != nil {
		if err := user.StateValidator(*uu.state); err != nil {
			return 0, fmt.Errorf("entv2: validator failed for field \"state\": %v", err)
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
	if value := uu.phone; value != nil {
		updater.Set(user.FieldPhone, *value)
	}
	if value := uu.buffer; value != nil {
		updater.Set(user.FieldBuffer, *value)
	}
	if uu.clearbuffer {
		updater.SetNull(user.FieldBuffer)
	}
	if value := uu.title; value != nil {
		updater.Set(user.FieldTitle, *value)
	}
	if value := uu.new_name; value != nil {
		updater.Set(user.FieldNewName, *value)
	}
	if uu.clearnew_name {
		updater.SetNull(user.FieldNewName)
	}
	if value := uu.blob; value != nil {
		updater.Set(user.FieldBlob, *value)
	}
	if uu.clearblob {
		updater.SetNull(user.FieldBlob)
	}
	if value := uu.state; value != nil {
		updater.Set(user.FieldState, *value)
	}
	if uu.clearstate {
		updater.SetNull(user.FieldState)
	}
	if !updater.Empty() {
		query, args := updater.Query()
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
	clearbuffer   bool
	title         *string
	new_name      *string
	clearnew_name bool
	blob          *[]byte
	clearblob     bool
	state         *user.State
	clearstate    bool
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

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.phone = &s
	return uuo
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetBuffer sets the buffer field.
func (uuo *UserUpdateOne) SetBuffer(b []byte) *UserUpdateOne {
	uuo.buffer = &b
	return uuo
}

// ClearBuffer clears the value of buffer.
func (uuo *UserUpdateOne) ClearBuffer() *UserUpdateOne {
	uuo.buffer = nil
	uuo.clearbuffer = true
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

// SetState sets the state field.
func (uuo *UserUpdateOne) SetState(u user.State) *UserUpdateOne {
	uuo.state = &u
	return uuo
}

// SetNillableState sets the state field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableState(u *user.State) *UserUpdateOne {
	if u != nil {
		uuo.SetState(*u)
	}
	return uuo
}

// ClearState clears the value of state.
func (uuo *UserUpdateOne) ClearState() *UserUpdateOne {
	uuo.state = nil
	uuo.clearstate = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if uuo.state != nil {
		if err := user.StateValidator(*uuo.state); err != nil {
			return nil, fmt.Errorf("entv2: validator failed for field \"state\": %v", err)
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
			return nil, fmt.Errorf("entv2: failed scanning row into User: %v", err)
		}
		id = u.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("User with id: %v", uuo.id)}
	case n > 1:
		return nil, fmt.Errorf("entv2: more than one User with the same id: %v", uuo.id)
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
	if value := uuo.phone; value != nil {
		updater.Set(user.FieldPhone, *value)
		u.Phone = *value
	}
	if value := uuo.buffer; value != nil {
		updater.Set(user.FieldBuffer, *value)
		u.Buffer = *value
	}
	if uuo.clearbuffer {
		var value []byte
		u.Buffer = value
		updater.SetNull(user.FieldBuffer)
	}
	if value := uuo.title; value != nil {
		updater.Set(user.FieldTitle, *value)
		u.Title = *value
	}
	if value := uuo.new_name; value != nil {
		updater.Set(user.FieldNewName, *value)
		u.NewName = *value
	}
	if uuo.clearnew_name {
		var value string
		u.NewName = value
		updater.SetNull(user.FieldNewName)
	}
	if value := uuo.blob; value != nil {
		updater.Set(user.FieldBlob, *value)
		u.Blob = *value
	}
	if uuo.clearblob {
		var value []byte
		u.Blob = value
		updater.SetNull(user.FieldBlob)
	}
	if value := uuo.state; value != nil {
		updater.Set(user.FieldState, *value)
		u.State = *value
	}
	if uuo.clearstate {
		var value user.State
		u.State = value
		updater.SetNull(user.FieldState)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
