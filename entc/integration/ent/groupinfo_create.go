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
	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
)

// GroupInfoCreate is the builder for creating a GroupInfo entity.
type GroupInfoCreate struct {
	config
	desc      *string
	max_users *int
	groups    map[string]struct{}
}

// SetDesc sets the desc field.
func (gic *GroupInfoCreate) SetDesc(s string) *GroupInfoCreate {
	gic.desc = &s
	return gic
}

// SetMaxUsers sets the max_users field.
func (gic *GroupInfoCreate) SetMaxUsers(i int) *GroupInfoCreate {
	gic.max_users = &i
	return gic
}

// SetNillableMaxUsers sets the max_users field if the given value is not nil.
func (gic *GroupInfoCreate) SetNillableMaxUsers(i *int) *GroupInfoCreate {
	if i != nil {
		gic.SetMaxUsers(*i)
	}
	return gic
}

// AddGroupIDs adds the groups edge to Group by ids.
func (gic *GroupInfoCreate) AddGroupIDs(ids ...string) *GroupInfoCreate {
	if gic.groups == nil {
		gic.groups = make(map[string]struct{})
	}
	for i := range ids {
		gic.groups[ids[i]] = struct{}{}
	}
	return gic
}

// AddGroups adds the groups edges to Group.
func (gic *GroupInfoCreate) AddGroups(g ...*Group) *GroupInfoCreate {
	ids := make([]string, len(g))
	for i := range g {
		ids[i] = g[i].ID
	}
	return gic.AddGroupIDs(ids...)
}

// Save creates the GroupInfo in the database.
func (gic *GroupInfoCreate) Save(ctx context.Context) (*GroupInfo, error) {
	if gic.desc == nil {
		return nil, errors.New("ent: missing required field \"desc\"")
	}
	if gic.max_users == nil {
		v := groupinfo.DefaultMaxUsers
		gic.max_users = &v
	}
	return gic.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (gic *GroupInfoCreate) SaveX(ctx context.Context) *GroupInfo {
	v, err := gic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gic *GroupInfoCreate) sqlSave(ctx context.Context) (*GroupInfo, error) {
	var (
		res     sql.Result
		builder = sql.Dialect(gic.driver.Dialect())
		gi      = &GroupInfo{config: gic.config}
	)
	tx, err := gic.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	insert := builder.Insert(groupinfo.Table).Default()
	if value := gic.desc; value != nil {
		insert.Set(groupinfo.FieldDesc, *value)
		gi.Desc = *value
	}
	if value := gic.max_users; value != nil {
		insert.Set(groupinfo.FieldMaxUsers, *value)
		gi.MaxUsers = *value
	}

	id, err := insertLastID(ctx, tx, insert.Returning(groupinfo.FieldID))
	if err != nil {
		return nil, rollback(tx, err)
	}
	gi.ID = strconv.FormatInt(id, 10)
	if len(gic.groups) > 0 {
		p := sql.P()
		for eid := range gic.groups {
			eid, err := strconv.Atoi(eid)
			if err != nil {
				return nil, rollback(tx, err)
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
		if int(affected) < len(gic.groups) {
			return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"groups\" %v already connected to a different \"GroupInfo\"", keys(gic.groups))})
		}
	}
	if err := tx.Commit(); err != nil {
		return nil, err
	}
	return gi, nil
}
