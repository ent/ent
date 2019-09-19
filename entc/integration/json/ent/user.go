// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// URL holds the value of the "url" field.
	URL *url.URL `json:"url,omitempty"`
	// Raw holds the value of the "raw" field.
	Raw json.RawMessage `json:"raw,omitempty"`
	// Dirs holds the value of the "dirs" field.
	Dirs []http.Dir `json:"dirs,omitempty"`
	// Ints holds the value of the "ints" field.
	Ints []int `json:"ints,omitempty"`
	// Floats holds the value of the "floats" field.
	Floats []float64 `json:"floats,omitempty"`
	// Strings holds the value of the "strings" field.
	Strings []string `json:"strings,omitempty"`
}

// FromRows scans the sql response data into User.
func (u *User) FromRows(rows *sql.Rows) error {
	var vu struct {
		ID      int
		URL     []byte
		Raw     []byte
		Dirs    []byte
		Ints    []byte
		Floats  []byte
		Strings []byte
	}
	// the order here should be the same as in the `user.Columns`.
	if err := rows.Scan(
		&vu.ID,
		&vu.URL,
		&vu.Raw,
		&vu.Dirs,
		&vu.Ints,
		&vu.Floats,
		&vu.Strings,
	); err != nil {
		return err
	}
	u.ID = vu.ID
	if value := vu.URL; len(value) > 0 {
		if err := json.Unmarshal(value, &u.URL); err != nil {
			return fmt.Errorf("unmarshal field url: %v", err)
		}
	}
	if value := vu.Raw; len(value) > 0 {
		if err := json.Unmarshal(value, &u.Raw); err != nil {
			return fmt.Errorf("unmarshal field raw: %v", err)
		}
	}
	if value := vu.Dirs; len(value) > 0 {
		if err := json.Unmarshal(value, &u.Dirs); err != nil {
			return fmt.Errorf("unmarshal field dirs: %v", err)
		}
	}
	if value := vu.Ints; len(value) > 0 {
		if err := json.Unmarshal(value, &u.Ints); err != nil {
			return fmt.Errorf("unmarshal field ints: %v", err)
		}
	}
	if value := vu.Floats; len(value) > 0 {
		if err := json.Unmarshal(value, &u.Floats); err != nil {
			return fmt.Errorf("unmarshal field floats: %v", err)
		}
	}
	if value := vu.Strings; len(value) > 0 {
		if err := json.Unmarshal(value, &u.Strings); err != nil {
			return fmt.Errorf("unmarshal field strings: %v", err)
		}
	}
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
	buf.WriteString(fmt.Sprintf(", url=%v", u.URL))
	buf.WriteString(fmt.Sprintf(", raw=%v", u.Raw))
	buf.WriteString(fmt.Sprintf(", dirs=%v", u.Dirs))
	buf.WriteString(fmt.Sprintf(", ints=%v", u.Ints))
	buf.WriteString(fmt.Sprintf(", floats=%v", u.Floats))
	buf.WriteString(fmt.Sprintf(", strings=%v", u.Strings))
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
