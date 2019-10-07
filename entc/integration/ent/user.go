// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"bytes"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `graphql:"-" json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"first_name" graphql:"first_name"`
	// Last holds the value of the "last" field.
	Last string `json:"last,omitempty" graphql:"last_name"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Password holds the value of the "password" field.
	Password string `graphql:"-" json:"-"`
}

// FromRows scans the sql response data into User.
func (u *User) FromRows(rows *sql.Rows) error {
	var vu struct {
		ID       int
		Age      sql.NullInt64
		Name     sql.NullString
		Last     sql.NullString
		Nickname sql.NullString
		Phone    sql.NullString
		Password sql.NullString
	}
	// the order here should be the same as in the `user.Columns`.
	if err := rows.Scan(
		&vu.ID,
		&vu.Age,
		&vu.Name,
		&vu.Last,
		&vu.Nickname,
		&vu.Phone,
		&vu.Password,
	); err != nil {
		return err
	}
	u.ID = strconv.Itoa(vu.ID)
	u.Age = int(vu.Age.Int64)
	u.Name = vu.Name.String
	u.Last = vu.Last.String
	u.Nickname = vu.Nickname.String
	u.Phone = vu.Phone.String
	u.Password = vu.Password.String
	return nil
}

// FromResponse scans the gremlin response data into User.
func (u *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vu struct {
		ID       string `json:"id,omitempty"`
		Age      int    `json:"age,omitempty"`
		Name     string `json:"name,omitempty"`
		Last     string `json:"last,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Phone    string `json:"phone,omitempty"`
		Password string `json:"password,omitempty"`
	}
	if err := vmap.Decode(&vu); err != nil {
		return err
	}
	u.ID = vu.ID
	u.Age = vu.Age
	u.Name = vu.Name
	u.Last = vu.Last
	u.Nickname = vu.Nickname
	u.Phone = vu.Phone
	u.Password = vu.Password
	return nil
}

// QueryCard queries the card edge of the User.
func (u *User) QueryCard() *CardQuery {
	return (&UserClient{u.config}).QueryCard(u)
}

// QueryPets queries the pets edge of the User.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{u.config}).QueryPets(u)
}

// QueryFiles queries the files edge of the User.
func (u *User) QueryFiles() *FileQuery {
	return (&UserClient{u.config}).QueryFiles(u)
}

// QueryGroups queries the groups edge of the User.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{u.config}).QueryGroups(u)
}

// QueryFriends queries the friends edge of the User.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{u.config}).QueryFriends(u)
}

// QueryFollowers queries the followers edge of the User.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{u.config}).QueryFollowers(u)
}

// QueryFollowing queries the following edge of the User.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{u.config}).QueryFollowing(u)
}

// QueryTeam queries the team edge of the User.
func (u *User) QueryTeam() *PetQuery {
	return (&UserClient{u.config}).QueryTeam(u)
}

// QuerySpouse queries the spouse edge of the User.
func (u *User) QuerySpouse() *UserQuery {
	return (&UserClient{u.config}).QuerySpouse(u)
}

// QueryChildren queries the children edge of the User.
func (u *User) QueryChildren() *UserQuery {
	return (&UserClient{u.config}).QueryChildren(u)
}

// QueryParent queries the parent edge of the User.
func (u *User) QueryParent() *UserQuery {
	return (&UserClient{u.config}).QueryParent(u)
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
	buf.WriteString(fmt.Sprintf(", last=%v", u.Last))
	buf.WriteString(fmt.Sprintf(", nickname=%v", u.Nickname))
	buf.WriteString(fmt.Sprintf(", phone=%v", u.Phone))
	buf.WriteString(", password=<sensitive>")
	buf.WriteString(")")
	return buf.String()
}

// id returns the int representation of the ID field.
func (u *User) id() int {
	id, _ := strconv.Atoi(u.ID)
	return id
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

// FromResponse scans the gremlin response data into Users.
func (u *Users) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var vu []struct {
		ID       string `json:"id,omitempty"`
		Age      int    `json:"age,omitempty"`
		Name     string `json:"name,omitempty"`
		Last     string `json:"last,omitempty"`
		Nickname string `json:"nickname,omitempty"`
		Phone    string `json:"phone,omitempty"`
		Password string `json:"password,omitempty"`
	}
	if err := vmap.Decode(&vu); err != nil {
		return err
	}
	for _, v := range vu {
		*u = append(*u, &User{
			ID:       v.ID,
			Age:      v.Age,
			Name:     v.Name,
			Last:     v.Last,
			Nickname: v.Nickname,
			Phone:    v.Phone,
			Password: v.Password,
		})
	}
	return nil
}

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
