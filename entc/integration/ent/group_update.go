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
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
	"github.com/facebookincubator/ent/entc/integration/ent/user"
)

// GroupUpdate is the builder for updating Group entities.
type GroupUpdate struct {
	config
	active         *bool
	expire         *time.Time
	_type          *string
	clear_type     bool
	max_users      *int
	addmax_users   *int
	clearmax_users bool
	name           *string
	files          map[string]struct{}
	blocked        map[string]struct{}
	users          map[string]struct{}
	info           map[string]struct{}
	removedFiles   map[string]struct{}
	removedBlocked map[string]struct{}
	removedUsers   map[string]struct{}
	clearedInfo    bool
	predicates     []predicate.Group
}

// Where adds a new predicate for the builder.
func (gu *GroupUpdate) Where(ps ...predicate.Group) *GroupUpdate {
	gu.predicates = append(gu.predicates, ps...)
	return gu
}

// SetActive sets the active field.
func (gu *GroupUpdate) SetActive(b bool) *GroupUpdate {
	gu.active = &b
	return gu
}

// SetNillableActive sets the active field if the given value is not nil.
func (gu *GroupUpdate) SetNillableActive(b *bool) *GroupUpdate {
	if b != nil {
		gu.SetActive(*b)
	}
	return gu
}

// SetExpire sets the expire field.
func (gu *GroupUpdate) SetExpire(t time.Time) *GroupUpdate {
	gu.expire = &t
	return gu
}

// SetType sets the type field.
func (gu *GroupUpdate) SetType(s string) *GroupUpdate {
	gu._type = &s
	return gu
}

// SetNillableType sets the type field if the given value is not nil.
func (gu *GroupUpdate) SetNillableType(s *string) *GroupUpdate {
	if s != nil {
		gu.SetType(*s)
	}
	return gu
}

// ClearType clears the value of type.
func (gu *GroupUpdate) ClearType() *GroupUpdate {
	gu._type = nil
	gu.clear_type = true
	return gu
}

// SetMaxUsers sets the max_users field.
func (gu *GroupUpdate) SetMaxUsers(i int) *GroupUpdate {
	gu.max_users = &i
	gu.addmax_users = nil
	return gu
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gu *GroupUpdate) SetNillableMaxUsers(i *int) *GroupUpdate {
	if i != nil {
		gu.SetMaxUsers(*i)
	}
	return gu
}

// AddMaxUsers adds i to max_users.
func (gu *GroupUpdate) AddMaxUsers(i int) *GroupUpdate {
	if gu.addmax_users == nil {
		gu.addmax_users = &i
	} else {
		*gu.addmax_users += i
	}
	return gu
}

// ClearMaxUsers clears the value of max_users.
func (gu *GroupUpdate) ClearMaxUsers() *GroupUpdate {
	gu.max_users = nil
	gu.clearmax_users = true
	return gu
}

// SetName sets the name field.
func (gu *GroupUpdate) SetName(s string) *GroupUpdate {
	gu.name = &s
	return gu
}

// AddFileIDs adds the files edge to File by ids.
func (gu *GroupUpdate) AddFileIDs(ids ...string) *GroupUpdate {
	if gu.files == nil {
		gu.files = make(map[string]struct{})
	}
	for i := range ids {
		gu.files[ids[i]] = struct{}{}
	}
	return gu
}

// AddFiles adds the files edges to File.
func (gu *GroupUpdate) AddFiles(f ...*File) *GroupUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gu.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (gu *GroupUpdate) AddBlockedIDs(ids ...string) *GroupUpdate {
	if gu.blocked == nil {
		gu.blocked = make(map[string]struct{})
	}
	for i := range ids {
		gu.blocked[ids[i]] = struct{}{}
	}
	return gu
}

// AddBlocked adds the blocked edges to User.
func (gu *GroupUpdate) AddBlocked(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (gu *GroupUpdate) AddUserIDs(ids ...string) *GroupUpdate {
	if gu.users == nil {
		gu.users = make(map[string]struct{})
	}
	for i := range ids {
		gu.users[ids[i]] = struct{}{}
	}
	return gu
}

// AddUsers adds the users edges to User.
func (gu *GroupUpdate) AddUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (gu *GroupUpdate) SetInfoID(id string) *GroupUpdate {
	if gu.info == nil {
		gu.info = make(map[string]struct{})
	}
	gu.info[id] = struct{}{}
	return gu
}

// SetInfo sets the info edge to GroupInfo.
func (gu *GroupUpdate) SetInfo(g *GroupInfo) *GroupUpdate {
	return gu.SetInfoID(g.ID)
}

// RemoveFileIDs removes the files edge to File by ids.
func (gu *GroupUpdate) RemoveFileIDs(ids ...string) *GroupUpdate {
	if gu.removedFiles == nil {
		gu.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		gu.removedFiles[ids[i]] = struct{}{}
	}
	return gu
}

// RemoveFiles removes files edges to File.
func (gu *GroupUpdate) RemoveFiles(f ...*File) *GroupUpdate {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return gu.RemoveFileIDs(ids...)
}

// RemoveBlockedIDs removes the blocked edge to User by ids.
func (gu *GroupUpdate) RemoveBlockedIDs(ids ...string) *GroupUpdate {
	if gu.removedBlocked == nil {
		gu.removedBlocked = make(map[string]struct{})
	}
	for i := range ids {
		gu.removedBlocked[ids[i]] = struct{}{}
	}
	return gu
}

// RemoveBlocked removes blocked edges to User.
func (gu *GroupUpdate) RemoveBlocked(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveBlockedIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (gu *GroupUpdate) RemoveUserIDs(ids ...string) *GroupUpdate {
	if gu.removedUsers == nil {
		gu.removedUsers = make(map[string]struct{})
	}
	for i := range ids {
		gu.removedUsers[ids[i]] = struct{}{}
	}
	return gu
}

// RemoveUsers removes users edges to User.
func (gu *GroupUpdate) RemoveUsers(u ...*User) *GroupUpdate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gu.RemoveUserIDs(ids...)
}

// ClearInfo clears the info edge to GroupInfo.
func (gu *GroupUpdate) ClearInfo() *GroupUpdate {
	gu.clearedInfo = true
	return gu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (gu *GroupUpdate) Save(ctx context.Context) (int, error) {
	if gu._type != nil {
		if err := group.TypeValidator(*gu._type); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	if gu.max_users != nil {
		if err := group.MaxUsersValidator(*gu.max_users); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"max_users\": %v", err)
		}
	}
	if gu.name != nil {
		if err := group.NameValidator(*gu.name); err != nil {
			return 0, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(gu.info) > 1 {
		return 0, errors.New("ent: multiple assignments on a unique edge \"info\"")
	}
	if gu.clearedInfo && gu.info == nil {
		return 0, errors.New("ent: clearing a unique edge \"info\"")
	}
	return gu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (gu *GroupUpdate) SaveX(ctx context.Context) int {
	affected, err := gu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (gu *GroupUpdate) Exec(ctx context.Context) error {
	_, err := gu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gu *GroupUpdate) ExecX(ctx context.Context) {
	if err := gu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (gu *GroupUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(gu.driver.Dialect())
		selector = builder.Select(group.FieldID).From(builder.Table(group.Table))
	)
	for _, p := range gu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = gu.driver.Query(ctx, query, args, rows); err != nil {
		return 0, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("ent: failed reading id: %v", err)
		}
		ids = append(ids, id)
	}
	if len(ids) == 0 {
		return 0, nil
	}

	tx, err := gu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(group.Table)
	)
	updater = updater.Where(sql.InInts(group.FieldID, ids...))
	if value := gu.active; value != nil {
		updater.Set(group.FieldActive, *value)
	}
	if value := gu.expire; value != nil {
		updater.Set(group.FieldExpire, *value)
	}
	if value := gu._type; value != nil {
		updater.Set(group.FieldType, *value)
	}
	if gu.clear_type {
		updater.SetNull(group.FieldType)
	}
	if value := gu.max_users; value != nil {
		updater.Set(group.FieldMaxUsers, *value)
	}
	if value := gu.addmax_users; value != nil {
		updater.Add(group.FieldMaxUsers, *value)
	}
	if gu.clearmax_users {
		updater.SetNull(group.FieldMaxUsers)
	}
	if value := gu.name; value != nil {
		updater.Set(group.FieldName, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.removedFiles) > 0 {
		eids := make([]int, len(gu.removedFiles))
		for eid := range gu.removedFiles {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(group.FilesTable).
			SetNull(group.FilesColumn).
			Where(sql.InInts(group.FilesColumn, ids...)).
			Where(sql.InInts(file.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.files) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range gu.files {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(file.FieldID, eid)
			}
			query, args := builder.Update(group.FilesTable).
				Set(group.FilesColumn, id).
				Where(sql.And(p, sql.IsNull(group.FilesColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(gu.files) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"Group\"", keys(gu.files))})
			}
		}
	}
	if len(gu.removedBlocked) > 0 {
		eids := make([]int, len(gu.removedBlocked))
		for eid := range gu.removedBlocked {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(group.BlockedTable).
			SetNull(group.BlockedColumn).
			Where(sql.InInts(group.BlockedColumn, ids...)).
			Where(sql.InInts(user.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.blocked) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range gu.blocked {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(user.FieldID, eid)
			}
			query, args := builder.Update(group.BlockedTable).
				Set(group.BlockedColumn, id).
				Where(sql.And(p, sql.IsNull(group.BlockedColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(gu.blocked) {
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"blocked\" %v already connected to a different \"Group\"", keys(gu.blocked))})
			}
		}
	}
	if len(gu.removedUsers) > 0 {
		eids := make([]int, len(gu.removedUsers))
		for eid := range gu.removedUsers {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Delete(group.UsersTable).
			Where(sql.InInts(group.UsersPrimaryKey[1], ids...)).
			Where(sql.InInts(group.UsersPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.users) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range gu.users {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				values = append(values, []int{id, eid})
			}
		}
		builder := builder.Insert(group.UsersTable).
			Columns(group.UsersPrimaryKey[1], group.UsersPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if gu.clearedInfo {
		query, args := builder.Update(group.InfoTable).
			SetNull(group.InfoColumn).
			Where(sql.InInts(groupinfo.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(gu.info) > 0 {
		for eid := range gu.info {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(group.InfoTable).
				Set(group.InfoColumn, eid).
				Where(sql.InInts(group.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// GroupUpdateOne is the builder for updating a single Group entity.
type GroupUpdateOne struct {
	config
	id             string
	active         *bool
	expire         *time.Time
	_type          *string
	clear_type     bool
	max_users      *int
	addmax_users   *int
	clearmax_users bool
	name           *string
	files          map[string]struct{}
	blocked        map[string]struct{}
	users          map[string]struct{}
	info           map[string]struct{}
	removedFiles   map[string]struct{}
	removedBlocked map[string]struct{}
	removedUsers   map[string]struct{}
	clearedInfo    bool
}

// SetActive sets the active field.
func (guo *GroupUpdateOne) SetActive(b bool) *GroupUpdateOne {
	guo.active = &b
	return guo
}

// SetNillableActive sets the active field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableActive(b *bool) *GroupUpdateOne {
	if b != nil {
		guo.SetActive(*b)
	}
	return guo
}

// SetExpire sets the expire field.
func (guo *GroupUpdateOne) SetExpire(t time.Time) *GroupUpdateOne {
	guo.expire = &t
	return guo
}

// SetType sets the type field.
func (guo *GroupUpdateOne) SetType(s string) *GroupUpdateOne {
	guo._type = &s
	return guo
}

// SetNillableType sets the type field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableType(s *string) *GroupUpdateOne {
	if s != nil {
		guo.SetType(*s)
	}
	return guo
}

// ClearType clears the value of type.
func (guo *GroupUpdateOne) ClearType() *GroupUpdateOne {
	guo._type = nil
	guo.clear_type = true
	return guo
}

// SetMaxUsers sets the max_users field.
func (guo *GroupUpdateOne) SetMaxUsers(i int) *GroupUpdateOne {
	guo.max_users = &i
	guo.addmax_users = nil
	return guo
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (guo *GroupUpdateOne) SetNillableMaxUsers(i *int) *GroupUpdateOne {
	if i != nil {
		guo.SetMaxUsers(*i)
	}
	return guo
}

// AddMaxUsers adds i to max_users.
func (guo *GroupUpdateOne) AddMaxUsers(i int) *GroupUpdateOne {
	if guo.addmax_users == nil {
		guo.addmax_users = &i
	} else {
		*guo.addmax_users += i
	}
	return guo
}

// ClearMaxUsers clears the value of max_users.
func (guo *GroupUpdateOne) ClearMaxUsers() *GroupUpdateOne {
	guo.max_users = nil
	guo.clearmax_users = true
	return guo
}

// SetName sets the name field.
func (guo *GroupUpdateOne) SetName(s string) *GroupUpdateOne {
	guo.name = &s
	return guo
}

// AddFileIDs adds the files edge to File by ids.
func (guo *GroupUpdateOne) AddFileIDs(ids ...string) *GroupUpdateOne {
	if guo.files == nil {
		guo.files = make(map[string]struct{})
	}
	for i := range ids {
		guo.files[ids[i]] = struct{}{}
	}
	return guo
}

// AddFiles adds the files edges to File.
func (guo *GroupUpdateOne) AddFiles(f ...*File) *GroupUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return guo.AddFileIDs(ids...)
}

// AddBlockedIDs adds the blocked edge to User by ids.
func (guo *GroupUpdateOne) AddBlockedIDs(ids ...string) *GroupUpdateOne {
	if guo.blocked == nil {
		guo.blocked = make(map[string]struct{})
	}
	for i := range ids {
		guo.blocked[ids[i]] = struct{}{}
	}
	return guo
}

// AddBlocked adds the blocked edges to User.
func (guo *GroupUpdateOne) AddBlocked(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddBlockedIDs(ids...)
}

// AddUserIDs adds the users edge to User by ids.
func (guo *GroupUpdateOne) AddUserIDs(ids ...string) *GroupUpdateOne {
	if guo.users == nil {
		guo.users = make(map[string]struct{})
	}
	for i := range ids {
		guo.users[ids[i]] = struct{}{}
	}
	return guo
}

// AddUsers adds the users edges to User.
func (guo *GroupUpdateOne) AddUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.AddUserIDs(ids...)
}

// SetInfoID sets the info edge to GroupInfo by id.
func (guo *GroupUpdateOne) SetInfoID(id string) *GroupUpdateOne {
	if guo.info == nil {
		guo.info = make(map[string]struct{})
	}
	guo.info[id] = struct{}{}
	return guo
}

// SetInfo sets the info edge to GroupInfo.
func (guo *GroupUpdateOne) SetInfo(g *GroupInfo) *GroupUpdateOne {
	return guo.SetInfoID(g.ID)
}

// RemoveFileIDs removes the files edge to File by ids.
func (guo *GroupUpdateOne) RemoveFileIDs(ids ...string) *GroupUpdateOne {
	if guo.removedFiles == nil {
		guo.removedFiles = make(map[string]struct{})
	}
	for i := range ids {
		guo.removedFiles[ids[i]] = struct{}{}
	}
	return guo
}

// RemoveFiles removes files edges to File.
func (guo *GroupUpdateOne) RemoveFiles(f ...*File) *GroupUpdateOne {
	ids := make([]string, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return guo.RemoveFileIDs(ids...)
}

// RemoveBlockedIDs removes the blocked edge to User by ids.
func (guo *GroupUpdateOne) RemoveBlockedIDs(ids ...string) *GroupUpdateOne {
	if guo.removedBlocked == nil {
		guo.removedBlocked = make(map[string]struct{})
	}
	for i := range ids {
		guo.removedBlocked[ids[i]] = struct{}{}
	}
	return guo
}

// RemoveBlocked removes blocked edges to User.
func (guo *GroupUpdateOne) RemoveBlocked(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveBlockedIDs(ids...)
}

// RemoveUserIDs removes the users edge to User by ids.
func (guo *GroupUpdateOne) RemoveUserIDs(ids ...string) *GroupUpdateOne {
	if guo.removedUsers == nil {
		guo.removedUsers = make(map[string]struct{})
	}
	for i := range ids {
		guo.removedUsers[ids[i]] = struct{}{}
	}
	return guo
}

// RemoveUsers removes users edges to User.
func (guo *GroupUpdateOne) RemoveUsers(u ...*User) *GroupUpdateOne {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return guo.RemoveUserIDs(ids...)
}

// ClearInfo clears the info edge to GroupInfo.
func (guo *GroupUpdateOne) ClearInfo() *GroupUpdateOne {
	guo.clearedInfo = true
	return guo
}

// Save executes the query and returns the updated entity.
func (guo *GroupUpdateOne) Save(ctx context.Context) (*Group, error) {
	if guo._type != nil {
		if err := group.TypeValidator(*guo._type); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"type\": %v", err)
		}
	}
	if guo.max_users != nil {
		if err := group.MaxUsersValidator(*guo.max_users); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"max_users\": %v", err)
		}
	}
	if guo.name != nil {
		if err := group.NameValidator(*guo.name); err != nil {
			return nil, fmt.Errorf("ent: validator failed for field \"name\": %v", err)
		}
	}
	if len(guo.info) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"info\"")
	}
	if guo.clearedInfo && guo.info == nil {
		return nil, errors.New("ent: clearing a unique edge \"info\"")
	}
	return guo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (guo *GroupUpdateOne) SaveX(ctx context.Context) *Group {
	gr, err := guo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gr
}

// Exec executes the query on the entity.
func (guo *GroupUpdateOne) Exec(ctx context.Context) error {
	_, err := guo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (guo *GroupUpdateOne) ExecX(ctx context.Context) {
	if err := guo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (guo *GroupUpdateOne) sqlSave(ctx context.Context) (gr *Group, err error) {
	var (
		builder  = sql.Dialect(guo.driver.Dialect())
		selector = builder.Select(group.Columns...).From(builder.Table(group.Table))
	)
	group.ID(guo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = guo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		gr = &Group{config: guo.config}
		if err := gr.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into Group: %v", err)
		}
		id = gr.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("Group with id: %v", guo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one Group with the same id: %v", guo.id)
	}

	tx, err := guo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(group.Table)
	)
	updater = updater.Where(sql.InInts(group.FieldID, ids...))
	if value := guo.active; value != nil {
		updater.Set(group.FieldActive, *value)
		gr.Active = *value
	}
	if value := guo.expire; value != nil {
		updater.Set(group.FieldExpire, *value)
		gr.Expire = *value
	}
	if value := guo._type; value != nil {
		updater.Set(group.FieldType, *value)
		gr.Type = value
	}
	if guo.clear_type {
		gr.Type = nil
		updater.SetNull(group.FieldType)
	}
	if value := guo.max_users; value != nil {
		updater.Set(group.FieldMaxUsers, *value)
		gr.MaxUsers = *value
	}
	if value := guo.addmax_users; value != nil {
		updater.Add(group.FieldMaxUsers, *value)
		gr.MaxUsers += *value
	}
	if guo.clearmax_users {
		var value int
		gr.MaxUsers = value
		updater.SetNull(group.FieldMaxUsers)
	}
	if value := guo.name; value != nil {
		updater.Set(group.FieldName, *value)
		gr.Name = *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.removedFiles) > 0 {
		eids := make([]int, len(guo.removedFiles))
		for eid := range guo.removedFiles {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(group.FilesTable).
			SetNull(group.FilesColumn).
			Where(sql.InInts(group.FilesColumn, ids...)).
			Where(sql.InInts(file.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.files) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range guo.files {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
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
			if int(affected) < len(guo.files) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"files\" %v already connected to a different \"Group\"", keys(guo.files))})
			}
		}
	}
	if len(guo.removedBlocked) > 0 {
		eids := make([]int, len(guo.removedBlocked))
		for eid := range guo.removedBlocked {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(group.BlockedTable).
			SetNull(group.BlockedColumn).
			Where(sql.InInts(group.BlockedColumn, ids...)).
			Where(sql.InInts(user.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.blocked) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range guo.blocked {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
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
			if int(affected) < len(guo.blocked) {
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"blocked\" %v already connected to a different \"Group\"", keys(guo.blocked))})
			}
		}
	}
	if len(guo.removedUsers) > 0 {
		eids := make([]int, len(guo.removedUsers))
		for eid := range guo.removedUsers {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Delete(group.UsersTable).
			Where(sql.InInts(group.UsersPrimaryKey[1], ids...)).
			Where(sql.InInts(group.UsersPrimaryKey[0], eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.users) > 0 {
		values := make([][]int, 0, len(ids))
		for _, id := range ids {
			for eid := range guo.users {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				values = append(values, []int{id, eid})
			}
		}
		builder := builder.Insert(group.UsersTable).
			Columns(group.UsersPrimaryKey[1], group.UsersPrimaryKey[0])
		for _, v := range values {
			builder.Values(v[0], v[1])
		}
		query, args := builder.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if guo.clearedInfo {
		query, args := builder.Update(group.InfoTable).
			SetNull(group.InfoColumn).
			Where(sql.InInts(groupinfo.FieldID, ids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(guo.info) > 0 {
		for eid := range guo.info {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			query, args := builder.Update(group.InfoTable).
				Set(group.InfoColumn, eid).
				Where(sql.InInts(group.FieldID, ids...)).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gr, nil
}
