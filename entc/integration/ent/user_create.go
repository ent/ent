// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/card"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/pet"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
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
	return uc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (uc *UserCreate) SaveX(ctx context.Context) *User {
	v, err := uc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uc *UserCreate) sqlSave(ctx context.Context) (*User, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(uc.driver.Dialect())
		u       = &User{config: uc.config}
	)
	tx, err := uc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(user.Table).Default()
	if value := uc.age; value != nil {
		insert.Set(user.FieldAge, *value)
		u.Age = *value
	}
	if value := uc.name; value != nil {
		insert.Set(user.FieldName, *value)
		u.Name = *value
	}
	if value := uc.last; value != nil {
		insert.Set(user.FieldLast, *value)
		u.Last = *value
	}
	if value := uc.nickname; value != nil {
		insert.Set(user.FieldNickname, *value)
		u.Nickname = *value
	}
	if value := uc.phone; value != nil {
		insert.Set(user.FieldPhone, *value)
		u.Phone = *value
	}
	if value := uc.password; value != nil {
		insert.Set(user.FieldPassword, *value)
		u.Password = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(user.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	u.ID = strconv.FormatInt(id, 10)
	if len(uc.card) > 0 {
		eid, err := strconv.Atoi(keys(uc.card)[0])
		if err != nil {
			return nil, err
		}
		query, args := builder.Update(user.CardTable).
			Set(user.CardColumn, id).
			Where(sql.EQ(card.FieldID, eid).And().IsNull(user.CardColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.card) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"card\" %v already connected to a different \"User\"", keys(uc.card))})
		}
	}
	if len(uc.pets) > 0 {
		p := sql.P()
		for eid := range uc.pets {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(pet.FieldID, eid)
		}
		query, args := builder.Update(user.PetsTable).
			Set(user.PetsColumn, id).
			Where(sql.And(p, sql.IsNull(user.PetsColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.pets) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"pets\" %v already connected to a different \"User\"", keys(uc.pets))})
		}
	}
	if len(uc.files) > 0 {
		p := sql.P()
		for eid := range uc.files {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(file.FieldID, eid)
		}
		query, args := builder.Update(user.FilesTable).
			Set(user.FilesColumn, id).
			Where(sql.And(p, sql.IsNull(user.FilesColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.files) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"User\"", keys(uc.files))})
		}
	}
	if len(uc.groups) > 0 {
		for eid := range uc.groups {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}

			query, args := builder.Insert(user.GroupsTable).
				Columns(user.GroupsPrimaryKey[0], user.GroupsPrimaryKey[1]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.friends) > 0 {
		for eid := range uc.friends {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}

			query, args := builder.Insert(user.FriendsTable).
				Columns(user.FriendsPrimaryKey[0], user.FriendsPrimaryKey[1]).
				Values(id, eid).
				Values(eid, id).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.followers) > 0 {
		for eid := range uc.followers {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}

			query, args := builder.Insert(user.FollowersTable).
				Columns(user.FollowersPrimaryKey[1], user.FollowersPrimaryKey[0]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.following) > 0 {
		for eid := range uc.following {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}

			query, args := builder.Insert(user.FollowingTable).
				Columns(user.FollowingPrimaryKey[0], user.FollowingPrimaryKey[1]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(uc.team) > 0 {
		eid, err := strconv.Atoi(keys(uc.team)[0])
		if err != nil {
			return nil, err
		}
		query, args := builder.Update(user.TeamTable).
			Set(user.TeamColumn, id).
			Where(sql.EQ(pet.FieldID, eid).And().IsNull(user.TeamColumn)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.team) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"team\" %v already connected to a different \"User\"", keys(uc.team))})
		}
	}
	if len(uc.spouse) > 0 {
		for eid := range uc.spouse {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(user.SpouseTable).
				Set(user.SpouseColumn, eid).
				Where(sql.EQ(user.FieldID, id)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			query, args = builder.Update(user.SpouseTable).
				Set(user.SpouseColumn, id).
				Where(sql.EQ(user.FieldID, eid).And().IsNull(user.SpouseColumn)).Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(uc.spouse) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("\"spouse\" (%v) already connected to a different \"User\"", eid)})
			}
		}
	}
	if len(uc.children) > 0 {
		p := sql.P()
		for eid := range uc.children {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(user.FieldID, eid)
		}
		query, args := builder.Update(user.ChildrenTable).
			Set(user.ChildrenColumn, id).
			Where(sql.And(p, sql.IsNull(user.ChildrenColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(uc.children) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"children\" %v already connected to a different \"User\"", keys(uc.children))})
		}
	}
	if len(uc.parent) > 0 {
		for eid := range uc.parent {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(user.ParentTable).
				Set(user.ParentColumn, eid).
				Where(sql.EQ(user.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return u, nil
}
