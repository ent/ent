// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"bytes"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/migrate/entv1/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int32 `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Renamed holds the value of the "renamed" field.
	Renamed string `json:"renamed,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// State holds the value of the "state" field.
	State user.State `json:"state,omitempty"`
}

// FromRows scans the sql response data into User.
func (u *User) FromRows(rows *sql.Rows) error {
	var vu struct {
		ID      int
		Age     sql.NullInt64
		Name    sql.NullString
		Address sql.NullString
		Renamed sql.NullString
		Blob    []byte
		State   sql.NullString
	}
	// the order here should be the same as in the `user.Columns`.
	if err := rows.Scan(
		&vu.ID,
		&vu.Age,
		&vu.Name,
		&vu.Address,
		&vu.Renamed,
		&vu.Blob,
		&vu.State,
	); err != nil {
		return err
	}
	u.ID = vu.ID
	u.Age = int32(vu.Age.Int64)
	u.Name = vu.Name.String
	u.Address = vu.Address.String
	u.Renamed = vu.Renamed.String
	u.Blob = vu.Blob
	u.State = user.State(vu.State.String)
	return nil
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
		panic("entv1: User is not a transactional entity")
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
	buf.WriteString(fmt.Sprintf(", address=%v", u.Address))
	buf.WriteString(fmt.Sprintf(", renamed=%v", u.Renamed))
	buf.WriteString(fmt.Sprintf(", blob=%v", u.Blob))
	buf.WriteString(fmt.Sprintf(", state=%v", u.State))
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
