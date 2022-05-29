// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/entc/integration/gremlin/ent/card"
	"entgo.io/ent/entc/integration/gremlin/ent/pet"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
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
	// Address holds the value of the "address" field.
	Address string `json:"address,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Password holds the value of the "password" field.
	Password string `graphql:"-" json:"-"`
	// Role holds the value of the "role" field.
	Role user.Role `json:"role,omitempty"`
	// Employment holds the value of the "employment" field.
	Employment user.Employment `json:"employment,omitempty"`
	// SSOCert holds the value of the "SSOCert" field.
	SSOCert string `json:"SSOCert,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Card holds the value of the card edge.
	Card *Card `json:"card,omitempty"`
	// Pets holds the value of the pets edge.
	Pets []*Pet `json:"pets,omitempty"`
	// Files holds the value of the files edge.
	Files []*File `json:"files,omitempty"`
	// Groups holds the value of the groups edge.
	Groups []*Group `json:"groups,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Team holds the value of the team edge.
	Team *Pet `json:"team,omitempty"`
	// Spouse holds the value of the spouse edge.
	Spouse *User `json:"spouse,omitempty"`
	// Children holds the value of the children edge.
	Children []*User `json:"children,omitempty"`
	// Parent holds the value of the parent edge.
	Parent *User `json:"parent,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [11]bool
}

// CardOrErr returns the Card value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) CardOrErr() (*Card, error) {
	if e.loadedTypes[0] {
		if e.Card == nil {
			// The edge card was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: card.Label}
		}
		return e.Card, nil
	}
	return nil, &NotLoadedError{edge: "card"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PetsOrErr() ([]*Pet, error) {
	if e.loadedTypes[1] {
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// FilesOrErr returns the Files value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FilesOrErr() ([]*File, error) {
	if e.loadedTypes[2] {
		return e.Files, nil
	}
	return nil, &NotLoadedError{edge: "files"}
}

// GroupsOrErr returns the Groups value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) GroupsOrErr() ([]*Group, error) {
	if e.loadedTypes[3] {
		return e.Groups, nil
	}
	return nil, &NotLoadedError{edge: "groups"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[4] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[5] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[6] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// TeamOrErr returns the Team value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) TeamOrErr() (*Pet, error) {
	if e.loadedTypes[7] {
		if e.Team == nil {
			// The edge team was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Team, nil
	}
	return nil, &NotLoadedError{edge: "team"}
}

// SpouseOrErr returns the Spouse value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) SpouseOrErr() (*User, error) {
	if e.loadedTypes[8] {
		if e.Spouse == nil {
			// The edge spouse was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Spouse, nil
	}
	return nil, &NotLoadedError{edge: "spouse"}
}

// ChildrenOrErr returns the Children value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) ChildrenOrErr() ([]*User, error) {
	if e.loadedTypes[9] {
		return e.Children, nil
	}
	return nil, &NotLoadedError{edge: "children"}
}

// ParentOrErr returns the Parent value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) ParentOrErr() (*User, error) {
	if e.loadedTypes[10] {
		if e.Parent == nil {
			// The edge parent was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Parent, nil
	}
	return nil, &NotLoadedError{edge: "parent"}
}

// FromResponse scans the gremlin response data into User.
func (u *User) FromResponse(res *gremlin.Response) error {
	vmap, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	var scanu struct {
		ID          string          `json:"id,omitempty"`
		OptionalInt int             `json:"optional_int,omitempty"`
		Age         int             `json:"age,omitempty"`
		Name        string          `json:"name,omitempty"`
		Last        string          `json:"last,omitempty"`
		Nickname    string          `json:"nickname,omitempty"`
		Address     string          `json:"address,omitempty"`
		Phone       string          `json:"phone,omitempty"`
		Password    string          `json:"password,omitempty"`
		Role        user.Role       `json:"role,omitempty"`
		Employment  user.Employment `json:"employment,omitempty"`
		SSOCert     string          `json:"sso_cert,omitempty"`
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
	u.Address = scanu.Address
	u.Phone = scanu.Phone
	u.Password = scanu.Password
	u.Role = scanu.Role
	u.Employment = scanu.Employment
	u.SSOCert = scanu.SSOCert
	return nil
}

// QueryCard queries the "card" edge of the User entity.
func (u *User) QueryCard() *CardQuery {
	return (&UserClient{config: u.config}).QueryCard(u)
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{config: u.config}).QueryPets(u)
}

// QueryFiles queries the "files" edge of the User entity.
func (u *User) QueryFiles() *FileQuery {
	return (&UserClient{config: u.config}).QueryFiles(u)
}

// QueryGroups queries the "groups" edge of the User entity.
func (u *User) QueryGroups() *GroupQuery {
	return (&UserClient{config: u.config}).QueryGroups(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowing(u)
}

// QueryTeam queries the "team" edge of the User entity.
func (u *User) QueryTeam() *PetQuery {
	return (&UserClient{config: u.config}).QueryTeam(u)
}

// QuerySpouse queries the "spouse" edge of the User entity.
func (u *User) QuerySpouse() *UserQuery {
	return (&UserClient{config: u.config}).QuerySpouse(u)
}

// QueryChildren queries the "children" edge of the User entity.
func (u *User) QueryChildren() *UserQuery {
	return (&UserClient{config: u.config}).QueryChildren(u)
}

// QueryParent queries the "parent" edge of the User entity.
func (u *User) QueryParent() *UserQuery {
	return (&UserClient{config: u.config}).QueryParent(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
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
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("optional_int=")
	builder.WriteString(fmt.Sprintf("%v", u.OptionalInt))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("last=")
	builder.WriteString(u.Last)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("address=")
	builder.WriteString(u.Address)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("password=<sensitive>")
	builder.WriteString(", ")
	builder.WriteString("role=")
	builder.WriteString(fmt.Sprintf("%v", u.Role))
	builder.WriteString(", ")
	builder.WriteString("employment=")
	builder.WriteString(fmt.Sprintf("%v", u.Employment))
	builder.WriteString(", ")
	builder.WriteString("SSOCert=")
	builder.WriteString(u.SSOCert)
	builder.WriteByte(')')
	return builder.String()
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
		ID          string          `json:"id,omitempty"`
		OptionalInt int             `json:"optional_int,omitempty"`
		Age         int             `json:"age,omitempty"`
		Name        string          `json:"name,omitempty"`
		Last        string          `json:"last,omitempty"`
		Nickname    string          `json:"nickname,omitempty"`
		Address     string          `json:"address,omitempty"`
		Phone       string          `json:"phone,omitempty"`
		Password    string          `json:"password,omitempty"`
		Role        user.Role       `json:"role,omitempty"`
		Employment  user.Employment `json:"employment,omitempty"`
		SSOCert     string          `json:"sso_cert,omitempty"`
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
			Address:     v.Address,
			Phone:       v.Phone,
			Password:    v.Password,
			Role:        v.Role,
			Employment:  v.Employment,
			SSOCert:     v.SSOCert,
		})
	}
	return nil
}

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
