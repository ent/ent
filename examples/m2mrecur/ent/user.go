// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
}

// FromRows scans the sql response data into User.
func (u *User) FromRows(rows *sql.Rows) error {
	var vu struct {
		ID   int
		Age  sql.NullInt64
		Name sql.NullString
	}
	// the order here should be the same as in the `user.Columns`.
	if err := rows.Scan(
		&vu.ID,
		&vu.Age,
		&vu.Name,
	); err != nil {
		return err
	}
	u.ID = vu.ID
	u.Age = int(vu.Age.Int64)
	u.Name = vu.Name.String
	return nil
}

// QueryFollowers queries the followers edge of the User.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{u.config}).QueryFollowers(u)
}

// QueryFollowing queries the following edge of the User.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{u.config}).QueryFollowing(u)
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	buf := bytes.NewBuffer(nil)
	buf.WriteString("User(")
	buf.WriteString(fmt.Sprintf("id=%v", u.ID))
	buf.WriteString(fmt.Sprintf(", age=%v", u.Age))
	buf.WriteString(fmt.Sprintf(", name=%v", u.Name))
	buf.WriteString(")")
	return buf.String()
}

// Users is a parsable slice of User.
type Users []*User

// FromRows scans the sql response data into Users.
func (u *Users) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		vu := &User{}
		if err := vu.FromRows(rows); err != nil {
			return err
		}
		*u = append(*u, vu)
	}
	return nil
}

func (u Users) config(cfg config) {
	for i := range u {
		u[i].config = cfg
	}
}
