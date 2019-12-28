// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `graphql:"-" json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// OptionalInt holds the value of the "optional_int" field.
	OptionalInt int `json:"optional_int,omitempty"`
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
<<<<<<< HEAD:entc/integration/gremlin/ent/user.go
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
=======
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges struct {
		// Card holds the value of the card edge.
		Card *Card
		// Pets holds the value of the pets edge.
		Pets []*Pet
		// Files holds the value of the files edge.
		Files []*File
		// Groups holds the value of the groups edge.
		Groups []*Group
		// Friends holds the value of the friends edge.
		Friends []*User
		// Followers holds the value of the followers edge.
		Followers []*User
		// Following holds the value of the following edge.
		Following []*User
		// Team holds the value of the team edge.
		Team *Pet
		// Spouse holds the value of the spouse edge.
		Spouse    *User
		spouse_id int
		// Children holds the value of the children edge.
		Children []*User
		// Parent holds the value of the parent edge.
		Parent    *User
		parent_id int
	}
>>>>>>> entc/gen: add With<T> method to query-builder template:entc/integration/ent/user.go
}

// FromResponse scans the gremlin response data into User.
func (u *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanu struct {
		ID          string    `json:"id,omitempty"`
		OptionalInt int       `json:"optional_int,omitempty"`
		Age         int       `json:"age,omitempty"`
		Name        string    `json:"name,omitempty"`
		Last        string    `json:"last,omitempty"`
		Nickname    string    `json:"nickname,omitempty"`
		Phone       string    `json:"phone,omitempty"`
		Password    string    `json:"password,omitempty"`
		Role        user.Role `json:"role,omitempty"`
	}
	if err := vmap.Decode(&scanu); err != nil {
		return err
	}
	u.ID = scanu.ID
	u.OptionalInt = scanu.OptionalInt
	u.Age = scanu.Age
	u.Name = scanu.Name
	u.Last = scanu.Last
	u.Nickname = scanu.Nickname
	u.Phone = scanu.Phone
	u.Password = scanu.Password
	u.Role = scanu.Role
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
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", optional_int=")
	builder.WriteString(fmt.Sprintf("%v", u.OptionalInt))
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteString(", last=")
	builder.WriteString(u.Last)
	builder.WriteString(", nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", password=<sensitive>")
	builder.WriteString(", role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteByte(')')
	return builder.String()
}

// id returns the int representation of the ID field.
func (u *User) id() int {
	id, _ := strconv.Atoi(u.ID)
	return id
}

// Users is a parsable slice of User.
type Users []*User

// FromResponse scans the gremlin response data into Users.
func (u *Users) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanu []struct {
		ID          string    `json:"id,omitempty"`
		OptionalInt int       `json:"optional_int,omitempty"`
		Age         int       `json:"age,omitempty"`
		Name        string    `json:"name,omitempty"`
		Last        string    `json:"last,omitempty"`
		Nickname    string    `json:"nickname,omitempty"`
		Phone       string    `json:"phone,omitempty"`
		Password    string    `json:"password,omitempty"`
		Role        user.Role `json:"role,omitempty"`
	}
	if err := vmap.Decode(&scanu); err != nil {
		return err
	}
	for _, v := range scanu {
		*u = append(*u, &User{
			ID:          v.ID,
			OptionalInt: v.OptionalInt,
			Age:         v.Age,
			Name:        v.Name,
			Last:        v.Last,
			Nickname:    v.Nickname,
			Phone:       v.Phone,
			Password:    v.Password,
			Role:        v.Role,
		})
	}
	return nil
}

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
