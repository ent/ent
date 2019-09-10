// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv2

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv2/user"

	"github.com/facebookincubator/ent/dialect/sql"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age        *int
	addage     *int
	name       *string
	phone      *string
	buffer     *[]byte
	title      *string
	blob       *[]byte
	predicates []predicate.User
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

// SetBlob sets the blob field.
func (uu *UserUpdate) SetBlob(b []byte) *UserUpdate {
	uu.blob = &b
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
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uu.age; value != nil {
		update = true
		builder.Set(user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		update = true
		builder.Add(user.FieldAge, *value)
	}
	if value := uu.name; value != nil {
		update = true
		builder.Set(user.FieldName, *value)
	}
	if value := uu.phone; value != nil {
		update = true
		builder.Set(user.FieldPhone, *value)
	}
	if value := uu.buffer; value != nil {
		update = true
		builder.Set(user.FieldBuffer, *value)
	}
	if value := uu.title; value != nil {
		update = true
		builder.Set(user.FieldTitle, *value)
	}
	if value := uu.blob; value != nil {
		update = true
		builder.Set(user.FieldBlob, *value)
	}
	if update {
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
	id     int
	age    *int
	addage *int
	name   *string
	phone  *string
	buffer *[]byte
	title  *string
	blob   *[]byte
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

// SetBlob sets the blob field.
func (uuo *UserUpdateOne) SetBlob(b []byte) *UserUpdateOne {
	uuo.blob = &b
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
		update  bool
		res     sql.Result
		builder = sql.Update(user.Table).Where(sql.InInts(user.FieldID, ids...))
	)
	if value := uuo.age; value != nil {
		update = true
		builder.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uuo.addage; value != nil {
		update = true
		builder.Add(user.FieldAge, *value)
		u.Age += *value
	}
	if value := uuo.name; value != nil {
		update = true
		builder.Set(user.FieldName, *value)
		u.Name = *value
	}
	if value := uuo.phone; value != nil {
		update = true
		builder.Set(user.FieldPhone, *value)
		u.Phone = *value
	}
	if value := uuo.buffer; value != nil {
		update = true
		builder.Set(user.FieldBuffer, *value)
		u.Buffer = *value
	}
	if value := uuo.title; value != nil {
		update = true
		builder.Set(user.FieldTitle, *value)
		u.Title = *value
	}
	if value := uuo.blob; value != nil {
		update = true
		builder.Set(user.FieldBlob, *value)
		u.Blob = *value
	}
	if update {
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
