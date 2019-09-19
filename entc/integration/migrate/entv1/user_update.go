// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/predicate"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age          *int32
	addage       *int32
	name         *string
	address      *string
	clearaddress bool
	renamed      *string
	clearrenamed bool
	blob         *[]byte
	clearblob    bool
	predicates   []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int32) *UserUpdate {
	uu.age = &i
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int32) *UserUpdate {
	uu.addage = &i
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetAddress sets the address field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.address = &s
	return uu
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of address.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.address = nil
	uu.clearaddress = true
	return uu
}

// SetRenamed sets the renamed field.
func (uu *UserUpdate) SetRenamed(s string) *UserUpdate {
	uu.renamed = &s
	return uu
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uu *UserUpdate) SetNillableRenamed(s *string) *UserUpdate {
	if s != nil {
		uu.SetRenamed(*s)
	}
	return uu
}

// ClearRenamed clears the value of renamed.
func (uu *UserUpdate) ClearRenamed() *UserUpdate {
	uu.renamed = nil
	uu.clearrenamed = true
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
	if uu.name != nil {
		if err := user.NameValidator(*uu.name); err != nil {
			return 0, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
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
			return 0, fmt.Errorf("entv1: failed reading id: %v", err)
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
	if value := uu.address; value != nil {
		builder.Set(user.FieldAddress, *value)
	}
	if uu.clearaddress {
		builder.SetNull(user.FieldAddress)
	}
	if value := uu.renamed; value != nil {
		builder.Set(user.FieldRenamed, *value)
	}
	if uu.clearrenamed {
		builder.SetNull(user.FieldRenamed)
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
	id           int
	age          *int32
	addage       *int32
	name         *string
	address      *string
	clearaddress bool
	renamed      *string
	clearrenamed bool
	blob         *[]byte
	clearblob    bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int32) *UserUpdateOne {
	uuo.age = &i
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int32) *UserUpdateOne {
	uuo.addage = &i
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetAddress sets the address field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.address = &s
	return uuo
}

// SetNillableAddress sets the address field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of address.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.address = nil
	uuo.clearaddress = true
	return uuo
}

// SetRenamed sets the renamed field.
func (uuo *UserUpdateOne) SetRenamed(s string) *UserUpdateOne {
	uuo.renamed = &s
	return uuo
}

// SetNillableRenamed sets the renamed field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRenamed(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetRenamed(*s)
	}
	return uuo
}

// ClearRenamed clears the value of renamed.
func (uuo *UserUpdateOne) ClearRenamed() *UserUpdateOne {
	uuo.renamed = nil
	uuo.clearrenamed = true
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
	if uuo.name != nil {
		if err := user.NameValidator(*uuo.name); err != nil {
			return nil, fmt.Errorf("entv1: validator failed for field \"name\": %v", err)
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
			return nil, fmt.Errorf("entv1: failed scanning row into User: %v", err)
		}
		id = u.ID
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, fmt.Errorf("entv1: User not found with id: %v", uuo.id)
	case n > 1:
		return nil, fmt.Errorf("entv1: more than one User with the same id: %v", uuo.id)
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
	if value := uuo.address; value != nil {
		builder.Set(user.FieldAddress, *value)
		u.Address = *value
	}
	if uuo.clearaddress {
		var value string
		u.Address = value
		builder.SetNull(user.FieldAddress)
	}
	if value := uuo.renamed; value != nil {
		builder.Set(user.FieldRenamed, *value)
		u.Renamed = *value
	}
	if uuo.clearrenamed {
		var value string
		u.Renamed = value
		builder.SetNull(user.FieldRenamed)
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
