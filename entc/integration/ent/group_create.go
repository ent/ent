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
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/file"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	active    *bool
	expire    *time.Time
	_type     *string
	max_users *int
	name      *string
	files     map[string]struct{}
	blocked   map[string]struct{}
	users     map[string]struct{}
	info      map[string]struct{}
}

// SetActive sets the active field.
func (gc *GroupCreate) SetActive(b bool) *GroupCreate {
	gc.active = &b
	return gc
}

// SetNillableActive sets the active field if the given value is not nil.
func (gc *GroupCreate) SetNillableActive(b *bool) *GroupCreate {
	if b != nil {
		gc.SetActive(*b)
	}
	return gc
}

// SetExpire sets the expire field.
func (gc *GroupCreate) SetExpire(t time.Time) *GroupCreate {
	gc.expire = &t
	return gc
}

// SetType sets the type field.
func (gc *GroupCreate) SetType(s string) *GroupCreate {
	gc._type = &s
	return gc
}

// SetNillableType sets the type field if the given value is not nil.
func (gc *GroupCreate) SetNillableType(s *string) *GroupCreate {
	if s != nil {
		gc.SetType(*s)
	}
	return gc
}

// SetMaxUsers sets the max_users field.
func (gc *GroupCreate) SetMaxUsers(i int) *GroupCreate {
	gc.max_users = &i
	return gc
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gc *GroupCreate) SetNillableMaxUsers(i *int) *GroupCreate {
	if i != nil {
		gc.SetMaxUsers(*i)
	}
	return gc
}

// SetName sets the name field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.name = &s
	return gc
}

// AddFileIDs adds the files edge to File by ids.
func (gc *GroupCreate) AddFileIDs(ids ...string) *GroupCreate {
	if gc.files == nil {
		gc.files = make(map[string]struct{})
	}
	for i := range ids {
		gc.files[ids[i]] = struct{}{}
	}
	return gc
}

// AddFiles adds the files edges to File.
func (gc *GroupCreate) AddFiles(f ...*File) *GroupCreate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gc.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (gc *GroupCreate) AddBlockedIDs(ids ...string) *GroupCreate {
	if gc.blocked == nil {
		gc.blocked = make(map[string]struct{})
	}
	for i := range ids {
		gc.blocked[ids[i]] = struct{}{}
	}
	return gc
}

// AddBlocked adds the blocked edges to User.
func (gc *GroupCreate) AddBlocked(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (gc *GroupCreate) AddUserIDs(ids ...string) *GroupCreate {
	if gc.users == nil {
		gc.users = make(map[string]struct{})
	}
	for i := range ids {
		gc.users[ids[i]] = struct{}{}
	}
	return gc
}

// AddUsers adds the users edges to User.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (gc *GroupCreate) SetInfoID(id string) *GroupCreate {
	if gc.info == nil {
		gc.info = make(map[string]struct{})
	}
	gc.info[id] = struct{}{}
	return gc
}

// SetInfo sets the info edge to GroupInfo.
func (gc *GroupCreate) SetInfo(g *GroupInfo) *GroupCreate {
	return gc.SetInfoID(g.ID)
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	if gc.active == nil {
		v := group.DefaultActive
		gc.active = &v
	}
	if gc.expire == nil {
		return nil, errors.New("ent: missing required field \"expire\"")
	}
	if gc._type != nil {
		if err := group.TypeValidator(*gc._type); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	if gc.max_users == nil {
		v := group.DefaultMaxUsers
		gc.max_users = &v
	}
	if err := group.MaxUsersValidator(*gc.max_users); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"max_users\": %v", err)
	}
	if gc.name == nil {
		return nil, errors.New("ent: missing required field \"name\"")
	}
	if err := group.NameValidator(*gc.name); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
	}
	if len(gc.info) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"info\"")
	}
	if gc.info == nil {
		return nil, errors.New("ent: missing required edge \"info\"")
	}
	return gc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gc.driver.Dialect())
		gr      = &Group{config: gc.config}
	)
	tx, err := gc.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(group.Table).Default()
	if value := gc.active; value != nil {
		insert.Set(group.FieldActive, *value)
		gr.Active = *value
	}
	if value := gc.expire; value != nil {
		insert.Set(group.FieldExpire, *value)
		gr.Expire = *value
	}
	if value := gc._type; value != nil {
		insert.Set(group.FieldType, *value)
		gr.Type = value
	}
	if value := gc.max_users; value != nil {
		insert.Set(group.FieldMaxUsers, *value)
		gr.MaxUsers = *value
	}
	if value := gc.name; value != nil {
		insert.Set(group.FieldName, *value)
		gr.Name = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(group.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	gr.ID = strconv.FormatInt(id, 10)
	if len(gc.files) > 0 {
		p := sql.P()
		for eid := range gc.files {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(file.FieldID, eid)
		}
		query, args := builder.Update(group.FilesTable).
			Set(group.FilesColumn, id).
			Where(sql.And(p, sql.IsNull(group.FilesColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(gc.files) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"Group\"", keys(gc.files))})
		}
	}
	if len(gc.blocked) > 0 {
		p := sql.P()
		for eid := range gc.blocked {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			p.Or().EQ(user.FieldID, eid)
		}
		query, args := builder.Update(group.BlockedTable).
			Set(group.BlockedColumn, id).
			Where(sql.And(p, sql.IsNull(group.BlockedColumn))).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
		affected, err := res.RowsAffected()
		if err != nil {
			return nil, rollback(tx, err)
		}
		if int(affected) < len(gc.blocked) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"blocked\" %v already connected to a different \"Group\"", keys(gc.blocked))})
		}
	}
	if len(gc.users) > 0 {
		for eid := range gc.users {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}

			query, args := builder.Insert(group.UsersTable).
				Columns(group.UsersPrimaryKey[1], group.UsersPrimaryKey[0]).
				Values(id, eid).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if len(gc.info) > 0 {
		for eid := range gc.info {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
			}
			query, args := builder.Update(group.InfoTable).
				Set(group.InfoColumn, eid).
				Where(sql.EQ(group.FieldID, id)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
