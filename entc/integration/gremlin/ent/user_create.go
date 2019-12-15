// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// UserCreate is the builder for creating a User entity.
type UserCreate struct {
	config
	age       *int
	name      *string
	last      *string
	nickname  *string
	phone     *string
	password  *string
	card      map[string]struct{}
	pets      map[string]struct{}
	files     map[string]struct{}
	groups    map[string]struct{}
	friends   map[string]struct{}
	followers map[string]struct{}
	following map[string]struct{}
	team      map[string]struct{}
	spouse    map[string]struct{}
	children  map[string]struct{}
	parent    map[string]struct{}
}

// SetAge sets the age field.
func (uc *UserCreate) SetAge(i int) *UserCreate {
	uc.age = &i
	return uc
}

// SetName sets the name field.
func (uc *UserCreate) SetName(s string) *UserCreate {
	uc.name = &s
	return uc
}

// SetLast sets the last field.
func (uc *UserCreate) SetLast(s string) *UserCreate {
	uc.last = &s
	return uc
}

// SetNillableLast sets the last field if the given value is not nil.
func (uc *UserCreate) SetNillableLast(s *string) *UserCreate {
	if s != nil {
		uc.SetLast(*s)
	}
	return uc
}

// SetNickname sets the nickname field.
func (uc *UserCreate) SetNickname(s string) *UserCreate {
	uc.nickname = &s
	return uc
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (uc *UserCreate) SetNillableNickname(s *string) *UserCreate {
	if s != nil {
		uc.SetNickname(*s)
	}
	return uc
}

// SetPhone sets the phone field.
func (uc *UserCreate) SetPhone(s string) *UserCreate {
	uc.phone = &s
	return uc
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uc *UserCreate) SetNillablePhone(s *string) *UserCreate {
	if s != nil {
		uc.SetPhone(*s)
	}
	return uc
}

// SetPassword sets the password field.
func (uc *UserCreate) SetPassword(s string) *UserCreate {
	uc.password = &s
	return uc
}

// SetNillablePassword sets the password field if the given value is not nil.
func (uc *UserCreate) SetNillablePassword(s *string) *UserCreate {
	if s != nil {
		uc.SetPassword(*s)
	}
	return uc
}

// SetCardID sets the card edge to Card by id.
func (uc *UserCreate) SetCardID(id string) *UserCreate {
	if uc.card == nil {
		uc.card = make(map[string]struct{})
	}
	uc.card[id] = struct{}{}
	return uc
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uc *UserCreate) SetNillableCardID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetCardID(*id)
	}
	return uc
}

// SetCard sets the card edge to Card.
func (uc *UserCreate) SetCard(c *Card) *UserCreate {
	return uc.SetCardID(c.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uc *UserCreate) AddPetIDs(ids ...string) *UserCreate {
	if uc.pets == nil {
		uc.pets = make(map[string]struct{})
	}
	for i := range ids {
		uc.pets[ids[i]] = struct{}{}
	}
	return uc
}

// AddPets adds the pets edges to Pet.
func (uc *UserCreate) AddPets(p ...*Pet) *UserCreate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uc.AddPetIDs(ids...)
}

// AddFileIDs adds the files edge to File by ids.
func (uc *UserCreate) AddFileIDs(ids ...string) *UserCreate {
	if uc.files == nil {
		uc.files = make(map[string]struct{})
	}
	for i := range ids {
		uc.files[ids[i]] = struct{}{}
	}
	return uc
}

// AddFiles adds the files edges to File.
func (uc *UserCreate) AddFiles(f ...*File) *UserCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uc.AddFileIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uc *UserCreate) AddGroupIDs(ids ...string) *UserCreate {
	if uc.groups == nil {
		uc.groups = make(map[string]struct{})
	}
	for i := range ids {
		uc.groups[ids[i]] = struct{}{}
	}
	return uc
}

// AddGroups adds the groups edges to Group.
func (uc *UserCreate) AddGroups(g ...*Group) *UserCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uc.AddGroupIDs(ids...)
}

// AddFriendIDs adds the friends edge to User by ids.
func (uc *UserCreate) AddFriendIDs(ids ...string) *UserCreate {
	if uc.friends == nil {
		uc.friends = make(map[string]struct{})
	}
	for i := range ids {
		uc.friends[ids[i]] = struct{}{}
	}
	return uc
}

// AddFriends adds the friends edges to User.
func (uc *UserCreate) AddFriends(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFriendIDs(ids...)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uc *UserCreate) AddFollowerIDs(ids ...string) *UserCreate {
	if uc.followers == nil {
		uc.followers = make(map[string]struct{})
	}
	for i := range ids {
		uc.followers[ids[i]] = struct{}{}
	}
	return uc
}

// AddFollowers adds the followers edges to User.
func (uc *UserCreate) AddFollowers(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uc *UserCreate) AddFollowingIDs(ids ...string) *UserCreate {
	if uc.following == nil {
		uc.following = make(map[string]struct{})
	}
	for i := range ids {
		uc.following[ids[i]] = struct{}{}
	}
	return uc
}

// AddFollowing adds the following edges to User.
func (uc *UserCreate) AddFollowing(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddFollowingIDs(ids...)
}

// SetTeamID sets the team edge to Pet by id.
func (uc *UserCreate) SetTeamID(id string) *UserCreate {
	if uc.team == nil {
		uc.team = make(map[string]struct{})
	}
	uc.team[id] = struct{}{}
	return uc
}

// SetNillableTeamID sets the team edge to Pet by id if the given value is not nil.
func (uc *UserCreate) SetNillableTeamID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetTeamID(*id)
	}
	return uc
}

// SetTeam sets the team edge to Pet.
func (uc *UserCreate) SetTeam(p *Pet) *UserCreate {
	return uc.SetTeamID(p.ID)
}

// SetSpouseID sets the spouse edge to User by id.
func (uc *UserCreate) SetSpouseID(id string) *UserCreate {
	if uc.spouse == nil {
		uc.spouse = make(map[string]struct{})
	}
	uc.spouse[id] = struct{}{}
	return uc
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableSpouseID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetSpouseID(*id)
	}
	return uc
}

// SetSpouse sets the spouse edge to User.
func (uc *UserCreate) SetSpouse(u *User) *UserCreate {
	return uc.SetSpouseID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uc *UserCreate) AddChildIDs(ids ...string) *UserCreate {
	if uc.children == nil {
		uc.children = make(map[string]struct{})
	}
	for i := range ids {
		uc.children[ids[i]] = struct{}{}
	}
	return uc
}

// AddChildren adds the children edges to User.
func (uc *UserCreate) AddChildren(u ...*User) *UserCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uc.AddChildIDs(ids...)
}

// SetParentID sets the parent edge to User by id.
func (uc *UserCreate) SetParentID(id string) *UserCreate {
	if uc.parent == nil {
		uc.parent = make(map[string]struct{})
	}
	uc.parent[id] = struct{}{}
	return uc
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uc *UserCreate) SetNillableParentID(id *string) *UserCreate {
	if id != nil {
		uc = uc.SetParentID(*id)
	}
	return uc
}

// SetParent sets the parent edge to User.
func (uc *UserCreate) SetParent(u *User) *UserCreate {
	return uc.SetParentID(u.ID)
}

// Save creates the User in the database.
func (uc *UserCreate) Save(ctx context.Context) (*User, error) {
	if uc.age == nil {
		return nil, errors.New("ent: missing required field \"age\"")
	}
	if uc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if uc.last == nil {
		v := user.DefaultLast
		uc.last = &v
	}
	if len(uc.card) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"card\"")
	}
	if len(uc.team) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"team\"")
	}
	if len(uc.spouse) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
	}
	if len(uc.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return uc.gremlinSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) gremlinSave(ctx context.Context) (*User, error) {
	res := &gremlin.Response{}
	query, bindings := uc.gremlin().Query()
	if err := uc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	u := &User{config: uc.config}
	if err := u.FromResponse(res); err != nil {
		return nil, err
	}
	return u, nil
}

func (uc *UserCreate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 8)
	v := g.AddV(user.Label)
	if uc.age != nil {
		v.Property(dsl.Single, user.FieldAge, *uc.age)
	}
	if uc.name != nil {
		v.Property(dsl.Single, user.FieldName, *uc.name)
	}
	if uc.last != nil {
		v.Property(dsl.Single, user.FieldLast, *uc.last)
	}
	if uc.nickname != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldNickname, *uc.nickname).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldNickname, *uc.nickname)),
		})
		v.Property(dsl.Single, user.FieldNickname, *uc.nickname)
	}
	if uc.phone != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldPhone, *uc.phone).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldPhone, *uc.phone)),
		})
		v.Property(dsl.Single, user.FieldPhone, *uc.phone)
	}
	if uc.password != nil {
		v.Property(dsl.Single, user.FieldPassword, *uc.password)
	}
	for id := range uc.card {
		v.AddE(user.CardLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.CardLabel, id)),
		})
	}
	for id := range uc.pets {
		v.AddE(user.PetsLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.PetsLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.PetsLabel, id)),
		})
	}
	for id := range uc.files {
		v.AddE(user.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.FilesLabel, id)),
		})
	}
	for id := range uc.groups {
		v.AddE(user.GroupsLabel).To(g.V(id)).OutV()
	}
	for id := range uc.friends {
		v.AddE(user.FriendsLabel).To(g.V(id)).OutV()
	}
	for id := range uc.followers {
		v.AddE(user.FollowingLabel).From(g.V(id)).InV()
	}
	for id := range uc.following {
		v.AddE(user.FollowingLabel).To(g.V(id)).OutV()
	}
	for id := range uc.team {
		v.AddE(user.TeamLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.TeamLabel, id)),
		})
	}
	for id := range uc.spouse {
		v.AddE(user.SpouseLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.SpouseLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
	}
	for id := range uc.children {
		v.AddE(user.ParentLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.ParentLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.ParentLabel, id)),
		})
	}
	for id := range uc.parent {
		v.AddE(user.ParentLabel).To(g.V(id)).OutV()
	}
	if len(constraints) == 0 {
		return v.ValueMap(true)
	}
	tr := constraints[0].pred.Coalesce(constraints[0].test, v.ValueMap(true))
	for _, cr := range constraints[1:] {
		tr = cr.pred.Coalesce(cr.test, tr)
	}
	return tr
}
