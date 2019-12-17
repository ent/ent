// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"strconv"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"
)

// GroupInfoUpdate is the builder for updating GroupInfo entities.
type GroupInfoUpdate struct {
	config
	desc          *string
	max_users     *int
	addmax_users  *int
	groups        map[string]struct{}
	removedGroups map[string]struct{}
	predicates    []predicate.GroupInfo
}

// Where adds a new predicate for the builder.
func (giu *GroupInfoUpdate) Where(ps ...predicate.GroupInfo) *GroupInfoUpdate {
	giu.predicates = append(giu.predicates, ps...)
	return giu
}

// SetDesc sets the desc field.
func (giu *GroupInfoUpdate) SetDesc(s string) *GroupInfoUpdate {
	giu.desc = &s
	return giu
}

// SetMaxUsers sets the max_users field.
func (giu *GroupInfoUpdate) SetMaxUsers(i int) *GroupInfoUpdate {
	giu.max_users = &i
	giu.addmax_users = nil
	return giu
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (giu *GroupInfoUpdate) SetNillableMaxUsers(i *int) *GroupInfoUpdate {
	if i != nil {
		giu.SetMaxUsers(*i)
	}
	return giu
}

// AddMaxUsers adds i to max_users.
func (giu *GroupInfoUpdate) AddMaxUsers(i int) *GroupInfoUpdate {
	if giu.addmax_users == nil {
		giu.addmax_users = &i
	} else {
		*giu.addmax_users += i
	}
	return giu
}

// AddGroupIDs adds the groups edge to Group by ids.
func (giu *GroupInfoUpdate) AddGroupIDs(ids ...string) *GroupInfoUpdate {
	if giu.groups == nil {
		giu.groups = make(map[string]struct{})
	}
	for i := range ids {
		giu.groups[ids[i]] = struct{}{}
	}
	return giu
}

// AddGroups adds the groups edges to Group.
func (giu *GroupInfoUpdate) AddGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.AddGroupIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (giu *GroupInfoUpdate) RemoveGroupIDs(ids ...string) *GroupInfoUpdate {
	if giu.removedGroups == nil {
		giu.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		giu.removedGroups[ids[i]] = struct{}{}
	}
	return giu
}

// RemoveGroups removes groups edges to Group.
func (giu *GroupInfoUpdate) RemoveGroups(g ...*Group) *GroupInfoUpdate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giu.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (giu *GroupInfoUpdate) Save(ctx context.Context) (int, error) {
	return giu.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (giu *GroupInfoUpdate) SaveX(ctx context.Context) int {
	affected, err := giu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (giu *GroupInfoUpdate) Exec(ctx context.Context) error {
	_, err := giu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giu *GroupInfoUpdate) ExecX(ctx context.Context) {
	if err := giu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giu *GroupInfoUpdate) sqlSave(ctx context.Context) (n int, err error) {
	var (
		builder  = sql.Dialect(giu.driver.Dialect())
		selector = builder.Select(groupinfo.FieldID).From(builder.Table(groupinfo.Table))
	)
	for _, p := range giu.predicates {
		p(selector)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = giu.driver.Query(ctx, query, args, rows); err != nil {
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

	tx, err := giu.driver.Tx(ctx)
	if err != nil {
		return 0, err
	}
	var (
		res     sql.Result
		updater = builder.Update(groupinfo.Table)
	)
	updater = updater.Where(sql.InInts(groupinfo.FieldID, ids...))
	if value := giu.desc; value != nil {
		updater.Set(groupinfo.FieldDesc, *value)
	}
	if value := giu.max_users; value != nil {
		updater.Set(groupinfo.FieldMaxUsers, *value)
	}
	if value := giu.addmax_users; value != nil {
		updater.Add(groupinfo.FieldMaxUsers, *value)
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(giu.removedGroups) > 0 {
		eids := make([]int, len(giu.removedGroups))
		for eid := range giu.removedGroups {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(groupinfo.GroupsTable).
			SetNull(groupinfo.GroupsColumn).
			Where(sql.InInts(groupinfo.GroupsColumn, ids...)).
			Where(sql.InInts(group.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return 0, rollback(tx, err)
		}
	}
	if len(giu.groups) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range giu.groups {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(group.FieldID, eid)
			}
			query, args := builder.Update(groupinfo.GroupsTable).
				Set(groupinfo.GroupsColumn, id).
				Where(sql.And(p, sql.IsNull(groupinfo.GroupsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return 0, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return 0, rollback(tx, err)
			}
			if int(affected) < len(giu.groups) {
				return 0, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"groups\" %v already connected to a different \"GroupInfo\"", keys(giu.groups))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

// GroupInfoUpdateOne is the builder for updating a single GroupInfo entity.
type GroupInfoUpdateOne struct {
	config
	id            string
	desc          *string
	max_users     *int
	addmax_users  *int
	groups        map[string]struct{}
	removedGroups map[string]struct{}
}

// SetDesc sets the desc field.
func (giuo *GroupInfoUpdateOne) SetDesc(s string) *GroupInfoUpdateOne {
	giuo.desc = &s
	return giuo
}

// SetMaxUsers sets the max_users field.
func (giuo *GroupInfoUpdateOne) SetMaxUsers(i int) *GroupInfoUpdateOne {
	giuo.max_users = &i
	giuo.addmax_users = nil
	return giuo
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (giuo *GroupInfoUpdateOne) SetNillableMaxUsers(i *int) *GroupInfoUpdateOne {
	if i != nil {
		giuo.SetMaxUsers(*i)
	}
	return giuo
}

// AddMaxUsers adds i to max_users.
func (giuo *GroupInfoUpdateOne) AddMaxUsers(i int) *GroupInfoUpdateOne {
	if giuo.addmax_users == nil {
		giuo.addmax_users = &i
	} else {
		*giuo.addmax_users += i
	}
	return giuo
}

// AddGroupIDs adds the groups edge to Group by ids.
func (giuo *GroupInfoUpdateOne) AddGroupIDs(ids ...string) *GroupInfoUpdateOne {
	if giuo.groups == nil {
		giuo.groups = make(map[string]struct{})
	}
	for i := range ids {
		giuo.groups[ids[i]] = struct{}{}
	}
	return giuo
}

// AddGroups adds the groups edges to Group.
func (giuo *GroupInfoUpdateOne) AddGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.AddGroupIDs(ids...)
}

// RemoveGroupIDs removes the groups edge to Group by ids.
func (giuo *GroupInfoUpdateOne) RemoveGroupIDs(ids ...string) *GroupInfoUpdateOne {
	if giuo.removedGroups == nil {
		giuo.removedGroups = make(map[string]struct{})
	}
	for i := range ids {
		giuo.removedGroups[ids[i]] = struct{}{}
	}
	return giuo
}

// RemoveGroups removes groups edges to Group.
func (giuo *GroupInfoUpdateOne) RemoveGroups(g ...*Group) *GroupInfoUpdateOne {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return giuo.RemoveGroupIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (giuo *GroupInfoUpdateOne) Save(ctx context.Context) (*GroupInfo, error) {
	return giuo.sqlSave(ctx)
}

// SaveX is like Save, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) SaveX(ctx context.Context) *GroupInfo {
	gi, err := giuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return gi
}

// Exec executes the query on the entity.
func (giuo *GroupInfoUpdateOne) Exec(ctx context.Context) error {
	_, err := giuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (giuo *GroupInfoUpdateOne) ExecX(ctx context.Context) {
	if err := giuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (giuo *GroupInfoUpdateOne) sqlSave(ctx context.Context) (gi *GroupInfo, err error) {
	var (
		builder  = sql.Dialect(giuo.driver.Dialect())
		selector = builder.Select(groupinfo.Columns...).From(builder.Table(groupinfo.Table))
	)
	groupinfo.ID(giuo.id)(selector)
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err = giuo.driver.Query(ctx, query, args, rows); err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		gi = &GroupInfo{config: giuo.config}
		if err := gi.FromRows(rows); err != nil {
			return nil, fmt.Errorf("ent: failed scanning row into GroupInfo: %v", err)
		}
		id = gi.id()
		ids = append(ids, id)
	}
	switch n := len(ids); {
	case n == 0:
		return nil, &ErrNotFound{fmt.Sprintf("GroupInfo with id: %v", giuo.id)}
	case n > 1:
		return nil, fmt.Errorf("ent: more than one GroupInfo with the same id: %v", giuo.id)
	}

	tx, err := giuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		res     sql.Result
		updater = builder.Update(groupinfo.Table)
	)
	updater = updater.Where(sql.InInts(groupinfo.FieldID, ids...))
	if value := giuo.desc; value != nil {
		updater.Set(groupinfo.FieldDesc, *value)
		gi.Desc = *value
	}
	if value := giuo.max_users; value != nil {
		updater.Set(groupinfo.FieldMaxUsers, *value)
		gi.MaxUsers = *value
	}
	if value := giuo.addmax_users; value != nil {
		updater.Add(groupinfo.FieldMaxUsers, *value)
		gi.MaxUsers += *value
	}
	if !updater.Empty() {
		query, args := updater.Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(giuo.removedGroups) > 0 {
		eids := make([]int, len(giuo.removedGroups))
		for eid := range giuo.removedGroups {
			eid, serr := strconv.Atoi(eid)
			if serr != nil {
				err = rollback(tx, serr)
				return
			}
			eids = append(eids, eid)
		}
		query, args := builder.Update(groupinfo.GroupsTable).
			SetNull(groupinfo.GroupsColumn).
			Where(sql.InInts(groupinfo.GroupsColumn, ids...)).
			Where(sql.InInts(group.FieldID, eids...)).
			Query()
		if err := tx.Exec(ctx, query, args, &res); err != nil {
			return nil, rollback(tx, err)
		}
	}
	if len(giuo.groups) > 0 {
		for _, id := range ids {
			p := sql.P()
			for eid := range giuo.groups {
				eid, serr := strconv.Atoi(eid)
				if serr != nil {
					err = rollback(tx, serr)
					return
				}
				p.Or().EQ(group.FieldID, eid)
			}
			query, args := builder.Update(groupinfo.GroupsTable).
				Set(groupinfo.GroupsColumn, id).
				Where(sql.And(p, sql.IsNull(groupinfo.GroupsColumn))).
				Query()
			if err := tx.Exec(ctx, query, args, &res); err != nil {
				return nil, rollback(tx, err)
			}
			affected, err := res.RowsAffected()
			if err != nil {
				return nil, rollback(tx, err)
			}
			if int(affected) < len(giuo.groups) {
				return nil, rollback(tx, &ConstraintError{msg: fmt.Sprintf("one of \"groups\" %v already connected to a different \"GroupInfo\"", keys(giuo.groups))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gi, nil
}
