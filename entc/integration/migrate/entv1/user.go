// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package entv1

import (
	"fmt"
	"strings"

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
	var scanu struct {
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
		&scanu.ID,
		&scanu.Age,
		&scanu.Name,
		&scanu.Address,
		&scanu.Renamed,
		&scanu.Blob,
		&scanu.State,
	); err != nil {
		return err
	}
	u.ID = scanu.ID
	u.Age = int32(scanu.Age.Int64)
	u.Name = scanu.Name.String
	u.Address = scanu.Address.String
	u.Renamed = scanu.Renamed.String
	u.Blob = scanu.Blob
	u.State = user.State(scanu.State.String)
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
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", address=")
	builder.WriteString(u.Address)
	builder.WriteString(", renamed=")
	builder.WriteString(u.Renamed)
	builder.WriteString(", blob=")
	builder.WriteString(fmt.Sprintf("%v", u.Blob))
	builder.WriteString(", state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

// FromRows scans the sql response data into Users.
func (u *Users) FromRows(rows *sql.Rows) error {
	for rows.Next() {
		scanu := &User{}
		if err := scanu.FromRows(rows); err != nil {
			return err
		}
		*u = append(*u, scanu)
	}
	return nil
}

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
