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
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	age              *int
	addage           *int
	name             *string
	last             *string
	nickname         *string
	clearnickname    bool
	phone            *string
	clearphone       bool
	password         *string
	clearpassword    bool
	card             map[string]struct{}
	pets             map[string]struct{}
	files            map[string]struct{}
	groups           map[string]struct{}
	friends          map[string]struct{}
	followers        map[string]struct{}
	following        map[string]struct{}
	team             map[string]struct{}
	spouse           map[string]struct{}
	children         map[string]struct{}
	parent           map[string]struct{}
	clearedCard      bool
	removedPets      map[string]struct{}
	removedFiles     map[string]struct{}
	removedGroups    map[string]struct{}
	removedFriends   map[string]struct{}
	removedFollowers map[string]struct{}
	removedFollowing map[string]struct{}
	clearedTeam      bool
	clearedSpouse    bool
	removedChildren  map[string]struct{}
	clearedParent    bool
	predicates       []predicate.User
}

// Where adds a new predicate for the builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.predicates = append(uu.predicates, ps...)
	return uu
}

// SetAge sets the age field.
func (uu *UserUpdate) SetAge(i int) *UserUpdate {
	uu.age = &i
	uu.addage = nil
	return uu
}

// AddAge adds i to age.
func (uu *UserUpdate) AddAge(i int) *UserUpdate {
	if uu.addage == nil {
		uu.addage = &i
	} else {
		*uu.addage += i
	}
	return uu
}

// SetName sets the name field.
func (uu *UserUpdate) SetName(s string) *UserUpdate {
	uu.name = &s
	return uu
}

// SetLast sets the last field.
func (uu *UserUpdate) SetLast(s string) *UserUpdate {
	uu.last = &s
	return uu
}

// SetNillableLast sets the last field if the given value is not nil.
func (uu *UserUpdate) SetNillableLast(s *string) *UserUpdate {
	if s != nil {
		uu.SetLast(*s)
	}
	return uu
}

// SetNickname sets the nickname field.
func (uu *UserUpdate) SetNickname(s string) *UserUpdate {
	uu.nickname = &s
	return uu
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (uu *UserUpdate) SetNillableNickname(s *string) *UserUpdate {
	if s != nil {
		uu.SetNickname(*s)
	}
	return uu
}

// ClearNickname clears the value of nickname.
func (uu *UserUpdate) ClearNickname() *UserUpdate {
	uu.nickname = nil
	uu.clearnickname = true
	return uu
}

// SetPhone sets the phone field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.phone = &s
	return uu
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// ClearPhone clears the value of phone.
func (uu *UserUpdate) ClearPhone() *UserUpdate {
	uu.phone = nil
	uu.clearphone = true
	return uu
}

// SetPassword sets the password field.
func (uu *UserUpdate) SetPassword(s string) *UserUpdate {
	uu.password = &s
	return uu
}

// SetNillablePassword sets the password field if the given value is not nil.
func (uu *UserUpdate) SetNillablePassword(s *string) *UserUpdate {
	if s != nil {
		uu.SetPassword(*s)
	}
	return uu
}

// ClearPassword clears the value of password.
func (uu *UserUpdate) ClearPassword() *UserUpdate {
	uu.password = nil
	uu.clearpassword = true
	return uu
}

// SetCardID sets the card edge to Card by id.
func (uu *UserUpdate) SetCardID(id string) *UserUpdate {
	if uu.card == nil {
		uu.card = make(map[string]struct{})
	}
	uu.card[id] = struct{}{}
	return uu
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uu *UserUpdate) SetNillableCardID(id *string) *UserUpdate {
	if id != nil {
		uu = uu.SetCardID(*id)
	}
	return uu
}

// SetCard sets the card edge to Card.
func (uu *UserUpdate) SetCard(c *Card) *UserUpdate {
	return uu.SetCardID(c.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uu *UserUpdate) AddPetIDs(ids ...string) *UserUpdate {
	if uu.pets == nil {
		uu.pets = make(map[string]struct{})
	}
	for i := range ids {
		uu.pets[ids[i]] = struct{}{}
	}
	return uu
}

// AddPets adds the pets edges to Pet.
func (uu *UserUpdate) AddPets(p ...*Pet) *UserUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.AddPetIDs(ids...)
}

// AddFileIDs adds the files edge to File by ids.
func (uu *UserUpdate) AddFileIDs(ids ...string) *UserUpdate {
	if uu.files == nil {
		uu.files = make(map[string]struct{})
	}
	for i := range ids {
		uu.files[ids[i]] = struct{}{}
	}
	return uu
}

// AddFiles adds the files edges to File.
func (uu *UserUpdate) AddFiles(f ...*File) *UserUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.AddFileIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uu *UserUpdate) AddGroupIDs(ids ...string) *UserUpdate {
	if uu.groups == nil {
		uu.groups = make(map[string]struct{})
	}
	for i := range ids {
		uu.groups[ids[i]] = struct{}{}
	}
	return uu
}

// AddGroups adds the groups edges to Group.
func (uu *UserUpdate) AddGroups(g ...*Group) *UserUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.AddGroupIDs(ids...)
}

// AddFriendIDs adds the friends edge to User by ids.
func (uu *UserUpdate) AddFriendIDs(ids ...string) *UserUpdate {
	if uu.friends == nil {
		uu.friends = make(map[string]struct{})
	}
	for i := range ids {
		uu.friends[ids[i]] = struct{}{}
	}
	return uu
}

// AddFriends adds the friends edges to User.
func (uu *UserUpdate) AddFriends(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFriendIDs(ids...)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uu *UserUpdate) AddFollowerIDs(ids ...string) *UserUpdate {
	if uu.followers == nil {
		uu.followers = make(map[string]struct{})
	}
	for i := range ids {
		uu.followers[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowers adds the followers edges to User.
func (uu *UserUpdate) AddFollowers(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uu *UserUpdate) AddFollowingIDs(ids ...string) *UserUpdate {
	if uu.following == nil {
		uu.following = make(map[string]struct{})
	}
	for i := range ids {
		uu.following[ids[i]] = struct{}{}
	}
	return uu
}

// AddFollowing adds the following edges to User.
func (uu *UserUpdate) AddFollowing(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddFollowingIDs(ids...)
}

// SetTeamID sets the team edge to Pet by id.
func (uu *UserUpdate) SetTeamID(id string) *UserUpdate {
	if uu.team == nil {
		uu.team = make(map[string]struct{})
	}
	uu.team[id] = struct{}{}
	return uu
}

// SetNillableTeamID sets the team edge to Pet by id if the given value is not nil.
func (uu *UserUpdate) SetNillableTeamID(id *string) *UserUpdate {
	if id != nil {
		uu = uu.SetTeamID(*id)
	}
	return uu
}

// SetTeam sets the team edge to Pet.
func (uu *UserUpdate) SetTeam(p *Pet) *UserUpdate {
	return uu.SetTeamID(p.ID)
}

// SetSpouseID sets the spouse edge to User by id.
func (uu *UserUpdate) SetSpouseID(id string) *UserUpdate {
	if uu.spouse == nil {
		uu.spouse = make(map[string]struct{})
	}
	uu.spouse[id] = struct{}{}
	return uu
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableSpouseID(id *string) *UserUpdate {
	if id != nil {
		uu = uu.SetSpouseID(*id)
	}
	return uu
}

// SetSpouse sets the spouse edge to User.
func (uu *UserUpdate) SetSpouse(u *User) *UserUpdate {
	return uu.SetSpouseID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uu *UserUpdate) AddChildIDs(ids ...string) *UserUpdate {
	if uu.children == nil {
		uu.children = make(map[string]struct{})
	}
	for i := range ids {
		uu.children[ids[i]] = struct{}{}
	}
	return uu
}

// AddChildren adds the children edges to User.
func (uu *UserUpdate) AddChildren(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.AddChildIDs(ids...)
}

// SetParentID sets the parent edge to User by id.
func (uu *UserUpdate) SetParentID(id string) *UserUpdate {
	if uu.parent == nil {
		uu.parent = make(map[string]struct{})
	}
	uu.parent[id] = struct{}{}
	return uu
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uu *UserUpdate) SetNillableParentID(id *string) *UserUpdate {
	if id != nil {
		uu = uu.SetParentID(*id)
	}
	return uu
}

// SetParent sets the parent edge to User.
func (uu *UserUpdate) SetParent(u *User) *UserUpdate {
	return uu.SetParentID(u.ID)
}

// ClearCard clears the card edge to Card.
func (uu *UserUpdate) ClearCard() *UserUpdate {
	uu.clearedCard = true
	return uu
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (uu *UserUpdate) RemovePetIDs(ids ...string) *UserUpdate {
	if uu.removedPets == nil {
		uu.removedPets = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedPets[ids[i]] = struct{}{}
	}
	return uu
}

// RemovePets removes pets edges to Pet.
func (uu *UserUpdate) RemovePets(p ...*Pet) *UserUpdate {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uu.RemovePetIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (uu *UserUpdate) RemoveFileIDs(ids ...string) *UserUpdate {
	if uu.removedFiles == nil {
		uu.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedFiles[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFiles removes files edges to File.
func (uu *UserUpdate) RemoveFiles(f ...*File) *UserUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uu.RemoveFileIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uu *UserUpdate) RemoveGroupIDs(ids ...string) *UserUpdate {
	if uu.removedGroups == nil {
		uu.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedGroups[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveGroups removes groups edges to Group.
func (uu *UserUpdate) RemoveGroups(g ...*Group) *UserUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uu.RemoveGroupIDs(ids...)
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (uu *UserUpdate) RemoveFriendIDs(ids ...string) *UserUpdate {
	if uu.removedFriends == nil {
		uu.removedFriends = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedFriends[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFriends removes friends edges to User.
func (uu *UserUpdate) RemoveFriends(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFriendIDs(ids...)
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uu *UserUpdate) RemoveFollowerIDs(ids ...string) *UserUpdate {
	if uu.removedFollowers == nil {
		uu.removedFollowers = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedFollowers[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowers removes followers edges to User.
func (uu *UserUpdate) RemoveFollowers(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uu *UserUpdate) RemoveFollowingIDs(ids ...string) *UserUpdate {
	if uu.removedFollowing == nil {
		uu.removedFollowing = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedFollowing[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveFollowing removes following edges to User.
func (uu *UserUpdate) RemoveFollowing(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveFollowingIDs(ids...)
}

// ClearTeam clears the team edge to Pet.
func (uu *UserUpdate) ClearTeam() *UserUpdate {
	uu.clearedTeam = true
	return uu
}

// ClearSpouse clears the spouse edge to User.
func (uu *UserUpdate) ClearSpouse() *UserUpdate {
	uu.clearedSpouse = true
	return uu
}

// RemoveChildIDs removes the children edge to User by ids.
func (uu *UserUpdate) RemoveChildIDs(ids ...string) *UserUpdate {
	if uu.removedChildren == nil {
		uu.removedChildren = make(map[string]struct{})
	}
	for i := range ids {
		uu.removedChildren[ids[i]] = struct{}{}
	}
	return uu
}

// RemoveChildren removes children edges to User.
func (uu *UserUpdate) RemoveChildren(u ...*User) *UserUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uu.RemoveChildIDs(ids...)
}

// ClearParent clears the parent edge to User.
func (uu *UserUpdate) ClearParent() *UserUpdate {
	uu.clearedParent = true
	return uu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	if len(uu.card) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"card\"")
	}
	if len(uu.team) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"team\"")
	}
	if len(uu.spouse) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
	}
	if len(uu.parent) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return uu.gremlinSave(ctx)
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

func (uu *UserUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := uu.gremlin().Query()
	if err := uu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (uu *UserUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 8)
	v := g.V().HasLabel(user.Label)
	for _, p := range uu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := uu.age; value != nil {
		v.Property(dsl.Single, user.FieldAge, *value)
	}
	if value := uu.addage; value != nil {
		v.Property(dsl.Single, user.FieldAge, __.Union(__.Values(user.FieldAge), __.Constant(*value)).Sum())
	}
	if value := uu.name; value != nil {
		v.Property(dsl.Single, user.FieldName, *value)
	}
	if value := uu.last; value != nil {
		v.Property(dsl.Single, user.FieldLast, *value)
	}
	if value := uu.nickname; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldNickname, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldNickname, *value)),
		})
		v.Property(dsl.Single, user.FieldNickname, *value)
	}
	if value := uu.phone; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldPhone, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldPhone, *value)),
		})
		v.Property(dsl.Single, user.FieldPhone, *value)
	}
	if value := uu.password; value != nil {
		v.Property(dsl.Single, user.FieldPassword, *value)
	}
	var properties []interface{}
	if uu.clearnickname {
		properties = append(properties, user.FieldNickname)
	}
	if uu.clearphone {
		properties = append(properties, user.FieldPhone)
	}
	if uu.clearpassword {
		properties = append(properties, user.FieldPassword)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if uu.clearedCard {
		tr := rv.Clone().OutE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.card {
		v.AddE(user.CardLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.CardLabel, id)),
		})
	}
	for id := range uu.removedPets {
		tr := rv.Clone().OutE(user.PetsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.pets {
		v.AddE(user.PetsLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.PetsLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.PetsLabel, id)),
		})
	}
	for id := range uu.removedFiles {
		tr := rv.Clone().OutE(user.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.files {
		v.AddE(user.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.FilesLabel, id)),
		})
	}
	for id := range uu.removedGroups {
		tr := rv.Clone().OutE(user.GroupsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.groups {
		v.AddE(user.GroupsLabel).To(g.V(id)).OutV()
	}
	for id := range uu.removedFriends {
		tr := rv.Clone().BothE(user.FriendsLabel).Where(__.Or(__.InV().HasID(id), __.OutV().HasID(id))).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.friends {
		v.AddE(user.FriendsLabel).To(g.V(id)).OutV()
	}
	for id := range uu.removedFollowers {
		tr := rv.Clone().InE(user.FollowingLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.followers {
		v.AddE(user.FollowingLabel).From(g.V(id)).InV()
	}
	for id := range uu.removedFollowing {
		tr := rv.Clone().OutE(user.FollowingLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.following {
		v.AddE(user.FollowingLabel).To(g.V(id)).OutV()
	}
	if uu.clearedTeam {
		tr := rv.Clone().OutE(user.TeamLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.team {
		v.AddE(user.TeamLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.TeamLabel, id)),
		})
	}
	if uu.clearedSpouse {
		tr := rv.Clone().BothE(user.SpouseLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.spouse {
		v.AddE(user.SpouseLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: rv.Clone().Both(user.SpouseLabel).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.SpouseLabel).Where(__.Or(__.InV().HasID(id), __.OutV().HasID(id))).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
	}
	for id := range uu.removedChildren {
		tr := rv.Clone().InE(user.ParentLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.children {
		v.AddE(user.ParentLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.ParentLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.ParentLabel, id)),
		})
	}
	if uu.clearedParent {
		tr := rv.Clone().OutE(user.ParentLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uu.parent {
		v.AddE(user.ParentLabel).To(g.V(id)).OutV()
	}
	v.Count()
	if len(constraints) > 0 {
		constraints = append(constraints, &constraint{
			pred: rv.Count(),
			test: __.Is(p.GT(1)).Constant(&ErrConstraintFailed{msg: "update traversal contains more than one vertex"}),
		})
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	id               string
	age              *int
	addage           *int
	name             *string
	last             *string
	nickname         *string
	clearnickname    bool
	phone            *string
	clearphone       bool
	password         *string
	clearpassword    bool
	card             map[string]struct{}
	pets             map[string]struct{}
	files            map[string]struct{}
	groups           map[string]struct{}
	friends          map[string]struct{}
	followers        map[string]struct{}
	following        map[string]struct{}
	team             map[string]struct{}
	spouse           map[string]struct{}
	children         map[string]struct{}
	parent           map[string]struct{}
	clearedCard      bool
	removedPets      map[string]struct{}
	removedFiles     map[string]struct{}
	removedGroups    map[string]struct{}
	removedFriends   map[string]struct{}
	removedFollowers map[string]struct{}
	removedFollowing map[string]struct{}
	clearedTeam      bool
	clearedSpouse    bool
	removedChildren  map[string]struct{}
	clearedParent    bool
}

// SetAge sets the age field.
func (uuo *UserUpdateOne) SetAge(i int) *UserUpdateOne {
	uuo.age = &i
	uuo.addage = nil
	return uuo
}

// AddAge adds i to age.
func (uuo *UserUpdateOne) AddAge(i int) *UserUpdateOne {
	if uuo.addage == nil {
		uuo.addage = &i
	} else {
		*uuo.addage += i
	}
	return uuo
}

// SetName sets the name field.
func (uuo *UserUpdateOne) SetName(s string) *UserUpdateOne {
	uuo.name = &s
	return uuo
}

// SetLast sets the last field.
func (uuo *UserUpdateOne) SetLast(s string) *UserUpdateOne {
	uuo.last = &s
	return uuo
}

// SetNillableLast sets the last field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLast(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLast(*s)
	}
	return uuo
}

// SetNickname sets the nickname field.
func (uuo *UserUpdateOne) SetNickname(s string) *UserUpdateOne {
	uuo.nickname = &s
	return uuo
}

// SetNillableNickname sets the nickname field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNickname(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNickname(*s)
	}
	return uuo
}

// ClearNickname clears the value of nickname.
func (uuo *UserUpdateOne) ClearNickname() *UserUpdateOne {
	uuo.nickname = nil
	uuo.clearnickname = true
	return uuo
}

// SetPhone sets the phone field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.phone = &s
	return uuo
}

// SetNillablePhone sets the phone field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// ClearPhone clears the value of phone.
func (uuo *UserUpdateOne) ClearPhone() *UserUpdateOne {
	uuo.phone = nil
	uuo.clearphone = true
	return uuo
}

// SetPassword sets the password field.
func (uuo *UserUpdateOne) SetPassword(s string) *UserUpdateOne {
	uuo.password = &s
	return uuo
}

// SetNillablePassword sets the password field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePassword(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPassword(*s)
	}
	return uuo
}

// ClearPassword clears the value of password.
func (uuo *UserUpdateOne) ClearPassword() *UserUpdateOne {
	uuo.password = nil
	uuo.clearpassword = true
	return uuo
}

// SetCardID sets the card edge to Card by id.
func (uuo *UserUpdateOne) SetCardID(id string) *UserUpdateOne {
	if uuo.card == nil {
		uuo.card = make(map[string]struct{})
	}
	uuo.card[id] = struct{}{}
	return uuo
}

// SetNillableCardID sets the card edge to Card by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCardID(id *string) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetCardID(*id)
	}
	return uuo
}

// SetCard sets the card edge to Card.
func (uuo *UserUpdateOne) SetCard(c *Card) *UserUpdateOne {
	return uuo.SetCardID(c.ID)
}

// AddPetIDs adds the pets edge to Pet by ids.
func (uuo *UserUpdateOne) AddPetIDs(ids ...string) *UserUpdateOne {
	if uuo.pets == nil {
		uuo.pets = make(map[string]struct{})
	}
	for i := range ids {
		uuo.pets[ids[i]] = struct{}{}
	}
	return uuo
}

// AddPets adds the pets edges to Pet.
func (uuo *UserUpdateOne) AddPets(p ...*Pet) *UserUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.AddPetIDs(ids...)
}

// AddFileIDs adds the files edge to File by ids.
func (uuo *UserUpdateOne) AddFileIDs(ids ...string) *UserUpdateOne {
	if uuo.files == nil {
		uuo.files = make(map[string]struct{})
	}
	for i := range ids {
		uuo.files[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFiles adds the files edges to File.
func (uuo *UserUpdateOne) AddFiles(f ...*File) *UserUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.AddFileIDs(ids...)
}

// AddGroupIDs adds the groups edge to Group by ids.
func (uuo *UserUpdateOne) AddGroupIDs(ids ...string) *UserUpdateOne {
	if uuo.groups == nil {
		uuo.groups = make(map[string]struct{})
	}
	for i := range ids {
		uuo.groups[ids[i]] = struct{}{}
	}
	return uuo
}

// AddGroups adds the groups edges to Group.
func (uuo *UserUpdateOne) AddGroups(g ...*Group) *UserUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.AddGroupIDs(ids...)
}

// AddFriendIDs adds the friends edge to User by ids.
func (uuo *UserUpdateOne) AddFriendIDs(ids ...string) *UserUpdateOne {
	if uuo.friends == nil {
		uuo.friends = make(map[string]struct{})
	}
	for i := range ids {
		uuo.friends[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFriends adds the friends edges to User.
func (uuo *UserUpdateOne) AddFriends(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFriendIDs(ids...)
}

// AddFollowerIDs adds the followers edge to User by ids.
func (uuo *UserUpdateOne) AddFollowerIDs(ids ...string) *UserUpdateOne {
	if uuo.followers == nil {
		uuo.followers = make(map[string]struct{})
	}
	for i := range ids {
		uuo.followers[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowers adds the followers edges to User.
func (uuo *UserUpdateOne) AddFollowers(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowerIDs(ids...)
}

// AddFollowingIDs adds the following edge to User by ids.
func (uuo *UserUpdateOne) AddFollowingIDs(ids ...string) *UserUpdateOne {
	if uuo.following == nil {
		uuo.following = make(map[string]struct{})
	}
	for i := range ids {
		uuo.following[ids[i]] = struct{}{}
	}
	return uuo
}

// AddFollowing adds the following edges to User.
func (uuo *UserUpdateOne) AddFollowing(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddFollowingIDs(ids...)
}

// SetTeamID sets the team edge to Pet by id.
func (uuo *UserUpdateOne) SetTeamID(id string) *UserUpdateOne {
	if uuo.team == nil {
		uuo.team = make(map[string]struct{})
	}
	uuo.team[id] = struct{}{}
	return uuo
}

// SetNillableTeamID sets the team edge to Pet by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableTeamID(id *string) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetTeamID(*id)
	}
	return uuo
}

// SetTeam sets the team edge to Pet.
func (uuo *UserUpdateOne) SetTeam(p *Pet) *UserUpdateOne {
	return uuo.SetTeamID(p.ID)
}

// SetSpouseID sets the spouse edge to User by id.
func (uuo *UserUpdateOne) SetSpouseID(id string) *UserUpdateOne {
	if uuo.spouse == nil {
		uuo.spouse = make(map[string]struct{})
	}
	uuo.spouse[id] = struct{}{}
	return uuo
}

// SetNillableSpouseID sets the spouse edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableSpouseID(id *string) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetSpouseID(*id)
	}
	return uuo
}

// SetSpouse sets the spouse edge to User.
func (uuo *UserUpdateOne) SetSpouse(u *User) *UserUpdateOne {
	return uuo.SetSpouseID(u.ID)
}

// AddChildIDs adds the children edge to User by ids.
func (uuo *UserUpdateOne) AddChildIDs(ids ...string) *UserUpdateOne {
	if uuo.children == nil {
		uuo.children = make(map[string]struct{})
	}
	for i := range ids {
		uuo.children[ids[i]] = struct{}{}
	}
	return uuo
}

// AddChildren adds the children edges to User.
func (uuo *UserUpdateOne) AddChildren(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.AddChildIDs(ids...)
}

// SetParentID sets the parent edge to User by id.
func (uuo *UserUpdateOne) SetParentID(id string) *UserUpdateOne {
	if uuo.parent == nil {
		uuo.parent = make(map[string]struct{})
	}
	uuo.parent[id] = struct{}{}
	return uuo
}

// SetNillableParentID sets the parent edge to User by id if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableParentID(id *string) *UserUpdateOne {
	if id != nil {
		uuo = uuo.SetParentID(*id)
	}
	return uuo
}

// SetParent sets the parent edge to User.
func (uuo *UserUpdateOne) SetParent(u *User) *UserUpdateOne {
	return uuo.SetParentID(u.ID)
}

// ClearCard clears the card edge to Card.
func (uuo *UserUpdateOne) ClearCard() *UserUpdateOne {
	uuo.clearedCard = true
	return uuo
}

// RemovePetIDs removes the pets edge to Pet by ids.
func (uuo *UserUpdateOne) RemovePetIDs(ids ...string) *UserUpdateOne {
	if uuo.removedPets == nil {
		uuo.removedPets = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedPets[ids[i]] = struct{}{}
	}
	return uuo
}

// RemovePets removes pets edges to Pet.
func (uuo *UserUpdateOne) RemovePets(p ...*Pet) *UserUpdateOne {
	ids := make([]string, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return uuo.RemovePetIDs(ids...)
}

// RemoveFileIDs removes the files edge to File by ids.
func (uuo *UserUpdateOne) RemoveFileIDs(ids ...string) *UserUpdateOne {
	if uuo.removedFiles == nil {
		uuo.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedFiles[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFiles removes files edges to File.
func (uuo *UserUpdateOne) RemoveFiles(f ...*File) *UserUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return uuo.RemoveFileIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (uuo *UserUpdateOne) RemoveGroupIDs(ids ...string) *UserUpdateOne {
	if uuo.removedGroups == nil {
		uuo.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedGroups[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveGroups removes groups edges to Group.
func (uuo *UserUpdateOne) RemoveGroups(g ...*Group) *UserUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return uuo.RemoveGroupIDs(ids...)
}

// RemoveFriendIDs removes the friends edge to User by ids.
func (uuo *UserUpdateOne) RemoveFriendIDs(ids ...string) *UserUpdateOne {
	if uuo.removedFriends == nil {
		uuo.removedFriends = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedFriends[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFriends removes friends edges to User.
func (uuo *UserUpdateOne) RemoveFriends(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFriendIDs(ids...)
}

// RemoveFollowerIDs removes the followers edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowerIDs(ids ...string) *UserUpdateOne {
	if uuo.removedFollowers == nil {
		uuo.removedFollowers = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedFollowers[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowers removes followers edges to User.
func (uuo *UserUpdateOne) RemoveFollowers(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowerIDs(ids...)
}

// RemoveFollowingIDs removes the following edge to User by ids.
func (uuo *UserUpdateOne) RemoveFollowingIDs(ids ...string) *UserUpdateOne {
	if uuo.removedFollowing == nil {
		uuo.removedFollowing = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedFollowing[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveFollowing removes following edges to User.
func (uuo *UserUpdateOne) RemoveFollowing(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveFollowingIDs(ids...)
}

// ClearTeam clears the team edge to Pet.
func (uuo *UserUpdateOne) ClearTeam() *UserUpdateOne {
	uuo.clearedTeam = true
	return uuo
}

// ClearSpouse clears the spouse edge to User.
func (uuo *UserUpdateOne) ClearSpouse() *UserUpdateOne {
	uuo.clearedSpouse = true
	return uuo
}

// RemoveChildIDs removes the children edge to User by ids.
func (uuo *UserUpdateOne) RemoveChildIDs(ids ...string) *UserUpdateOne {
	if uuo.removedChildren == nil {
		uuo.removedChildren = make(map[string]struct{})
	}
	for i := range ids {
		uuo.removedChildren[ids[i]] = struct{}{}
	}
	return uuo
}

// RemoveChildren removes children edges to User.
func (uuo *UserUpdateOne) RemoveChildren(u ...*User) *UserUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return uuo.RemoveChildIDs(ids...)
}

// ClearParent clears the parent edge to User.
func (uuo *UserUpdateOne) ClearParent() *UserUpdateOne {
	uuo.clearedParent = true
	return uuo
}

// Save executes the query and returns the updated entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	if len(uuo.card) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"card\"")
	}
	if len(uuo.team) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"team\"")
	}
	if len(uuo.spouse) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"spouse\"")
	}
	if len(uuo.parent) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"parent\"")
	}
	return uuo.gremlinSave(ctx)
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

func (uuo *UserUpdateOne) gremlinSave(ctx context.Context) (*User, error) {
	res := &gremlin.Response{}
	query, bindings := uuo.gremlin(uuo.id).Query()
	if err := uuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	u := &User{config: uuo.config}
	if err := u.FromResponse(res); err != nil {
		return nil, err
	}
	return u, nil
}

func (uuo *UserUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 8)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := uuo.age; value != nil {
		v.Property(dsl.Single, user.FieldAge, *value)
	}
	if value := uuo.addage; value != nil {
		v.Property(dsl.Single, user.FieldAge, __.Union(__.Values(user.FieldAge), __.Constant(*value)).Sum())
	}
	if value := uuo.name; value != nil {
		v.Property(dsl.Single, user.FieldName, *value)
	}
	if value := uuo.last; value != nil {
		v.Property(dsl.Single, user.FieldLast, *value)
	}
	if value := uuo.nickname; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldNickname, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldNickname, *value)),
		})
		v.Property(dsl.Single, user.FieldNickname, *value)
	}
	if value := uuo.phone; value != nil {
		constraints = append(constraints, &constraint{
			pred: g.V().Has(user.Label, user.FieldPhone, *value).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueField(user.Label, user.FieldPhone, *value)),
		})
		v.Property(dsl.Single, user.FieldPhone, *value)
	}
	if value := uuo.password; value != nil {
		v.Property(dsl.Single, user.FieldPassword, *value)
	}
	var properties []interface{}
	if uuo.clearnickname {
		properties = append(properties, user.FieldNickname)
	}
	if uuo.clearphone {
		properties = append(properties, user.FieldPhone)
	}
	if uuo.clearpassword {
		properties = append(properties, user.FieldPassword)
	}
	if len(properties) > 0 {
		v.SideEffect(__.Properties(properties...).Drop())
	}
	if uuo.clearedCard {
		tr := rv.Clone().OutE(user.CardLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.card {
		v.AddE(user.CardLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.CardLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.CardLabel, id)),
		})
	}
	for id := range uuo.removedPets {
		tr := rv.Clone().OutE(user.PetsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.pets {
		v.AddE(user.PetsLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.PetsLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.PetsLabel, id)),
		})
	}
	for id := range uuo.removedFiles {
		tr := rv.Clone().OutE(user.FilesLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.files {
		v.AddE(user.FilesLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.FilesLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.FilesLabel, id)),
		})
	}
	for id := range uuo.removedGroups {
		tr := rv.Clone().OutE(user.GroupsLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.groups {
		v.AddE(user.GroupsLabel).To(g.V(id)).OutV()
	}
	for id := range uuo.removedFriends {
		tr := rv.Clone().BothE(user.FriendsLabel).Where(__.Or(__.InV().HasID(id), __.OutV().HasID(id))).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.friends {
		v.AddE(user.FriendsLabel).To(g.V(id)).OutV()
	}
	for id := range uuo.removedFollowers {
		tr := rv.Clone().InE(user.FollowingLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.followers {
		v.AddE(user.FollowingLabel).From(g.V(id)).InV()
	}
	for id := range uuo.removedFollowing {
		tr := rv.Clone().OutE(user.FollowingLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.following {
		v.AddE(user.FollowingLabel).To(g.V(id)).OutV()
	}
	if uuo.clearedTeam {
		tr := rv.Clone().OutE(user.TeamLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.team {
		v.AddE(user.TeamLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.TeamLabel).InV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.TeamLabel, id)),
		})
	}
	if uuo.clearedSpouse {
		tr := rv.Clone().BothE(user.SpouseLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.spouse {
		v.AddE(user.SpouseLabel).To(g.V(id)).OutV()
		constraints = append(constraints, &constraint{
			pred: rv.Clone().Both(user.SpouseLabel).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.SpouseLabel).Where(__.Or(__.InV().HasID(id), __.OutV().HasID(id))).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.SpouseLabel, id)),
		})
	}
	for id := range uuo.removedChildren {
		tr := rv.Clone().InE(user.ParentLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.children {
		v.AddE(user.ParentLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(user.ParentLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(user.Label, user.ParentLabel, id)),
		})
	}
	if uuo.clearedParent {
		tr := rv.Clone().OutE(user.ParentLabel).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range uuo.parent {
		v.AddE(user.ParentLabel).To(g.V(id)).OutV()
	}
	v.ValueMap(true)
	if len(constraints) > 0 {
		v = constraints[0].pred.Coalesce(constraints[0].test, v)
		for _, cr := range constraints[1:] {
			v = cr.pred.Coalesce(cr.test, v)
		}
	}
	trs = append(trs, v)
	return dsl.Join(trs...)
}
