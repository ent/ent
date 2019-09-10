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

	"github.com/facebookincubator/ent/entc/integration/ent/group"
	"github.com/facebookincubator/ent/entc/integration/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/ent/predicate"

	"github.com/facebookincubator/ent/dialect"
	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/dialect/sql"
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
	giu.addmax_users = &i
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
	switch giu.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giu.sqlSave(ctx)
	case dialect.Gremlin:
		return giu.gremlinSave(ctx)
	default:
		return 0, errors.New("ent: unsupported dialect")
	}
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
	selector := sql.Select(groupinfo.FieldID).From(sql.Table(groupinfo.Table))
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
		update  bool
		res     sql.Result
		builder = sql.Update(groupinfo.Table).Where(sql.InInts(groupinfo.FieldID, ids...))
	)
	if value := giu.desc; value != nil {
		update = true
		builder.Set(groupinfo.FieldDesc, *value)
	}
	if value := giu.max_users; value != nil {
		update = true
		builder.Set(groupinfo.FieldMaxUsers, *value)
	}
	if value := giu.addmax_users; value != nil {
		update = true
		builder.Add(groupinfo.FieldMaxUsers, *value)
	}
	if update {
		query, args := builder.Query()
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
		query, args := sql.Update(groupinfo.GroupsTable).
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
			query, args := sql.Update(groupinfo.GroupsTable).
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
				return 0, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"groups\" %v already connected to a different \"GroupInfo\"", keys(giu.groups))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return len(ids), nil
}

func (giu *GroupInfoUpdate) gremlinSave(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := giu.gremlin().Query()
	if err := giu.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	if err, ok := isConstantError(res); ok {
		return 0, err
	}
	return res.ReadInt()
}

func (giu *GroupInfoUpdate) gremlin() *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V().HasLabel(groupinfo.Label)
	for _, p := range giu.predicates {
		p(v)
	}
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := giu.desc; value != nil {
		v.Property(dsl.Single, groupinfo.FieldDesc, *value)
	}
	if value := giu.max_users; value != nil {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, *value)
	}
	if value := giu.addmax_users; value != nil {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, __.Union(__.Values(groupinfo.FieldMaxUsers), __.Constant(*value)).Sum())
	}
	for id := range giu.removedGroups {
		tr := rv.Clone().InE(group.InfoLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range giu.groups {
		v.AddE(group.InfoLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.InfoLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(groupinfo.Label, group.InfoLabel, id)),
		})
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
	giuo.addmax_users = &i
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
	switch giuo.driver.Dialect() {
	case dialect.MySQL, dialect.SQLite:
		return giuo.sqlSave(ctx)
	case dialect.Gremlin:
		return giuo.gremlinSave(ctx)
	default:
		return nil, errors.New("ent: unsupported dialect")
	}
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
	selector := sql.Select(groupinfo.Columns...).From(sql.Table(groupinfo.Table))
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
		return nil, fmt.Errorf("ent: GroupInfo not found with id: %v", giuo.id)
	case n > 1:
		return nil, fmt.Errorf("ent: more than one GroupInfo with the same id: %v", giuo.id)
	}

	tx, err := giuo.driver.Tx(ctx)
	if err != nil {
		return nil, err
	}
	var (
		update  bool
		res     sql.Result
		builder = sql.Update(groupinfo.Table).Where(sql.InInts(groupinfo.FieldID, ids...))
	)
	if value := giuo.desc; value != nil {
		update = true
		builder.Set(groupinfo.FieldDesc, *value)
		gi.Desc = *value
	}
	if value := giuo.max_users; value != nil {
		update = true
		builder.Set(groupinfo.FieldMaxUsers, *value)
		gi.MaxUsers = *value
	}
	if value := giuo.addmax_users; value != nil {
		update = true
		builder.Add(groupinfo.FieldMaxUsers, *value)
		gi.MaxUsers += *value
	}
	if update {
		query, args := builder.Query()
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
		query, args := sql.Update(groupinfo.GroupsTable).
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
			query, args := sql.Update(groupinfo.GroupsTable).
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
				return nil, rollback(tx, &ErrConstraintFailed{msg: fmt.Sprintf("one of \"groups\" %v already connected to a different \"GroupInfo\"", keys(giuo.groups))})
			}
		}
	}
	if err = tx.Commit(); err != nil {
		return nil, err
	}
	return gi, nil
}

func (giuo *GroupInfoUpdateOne) gremlinSave(ctx context.Context) (*GroupInfo, error) {
	res := &gremlin.Response{}
	query, bindings := giuo.gremlin(giuo.id).Query()
	if err := giuo.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	gi := &GroupInfo{config: giuo.config}
	if err := gi.FromResponse(res); err != nil {
		return nil, err
	}
	return gi, nil
}

func (giuo *GroupInfoUpdateOne) gremlin(id string) *dsl.Traversal {
	type constraint struct {
		pred *dsl.Traversal // constraint predicate.
		test *dsl.Traversal // test matches and its constant.
	}
	constraints := make([]*constraint, 0, 1)
	v := g.V(id)
	var (
		rv = v.Clone()
		_  = rv

		trs []*dsl.Traversal
	)
	if value := giuo.desc; value != nil {
		v.Property(dsl.Single, groupinfo.FieldDesc, *value)
	}
	if value := giuo.max_users; value != nil {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, *value)
	}
	if value := giuo.addmax_users; value != nil {
		v.Property(dsl.Single, groupinfo.FieldMaxUsers, __.Union(__.Values(groupinfo.FieldMaxUsers), __.Constant(*value)).Sum())
	}
	for id := range giuo.removedGroups {
		tr := rv.Clone().InE(group.InfoLabel).Where(__.OtherV().HasID(id)).Drop().Iterate()
		trs = append(trs, tr)
	}
	for id := range giuo.groups {
		v.AddE(group.InfoLabel).From(g.V(id)).InV()
		constraints = append(constraints, &constraint{
			pred: g.E().HasLabel(group.InfoLabel).OutV().HasID(id).Count(),
			test: __.Is(p.NEQ(0)).Constant(NewErrUniqueEdge(groupinfo.Label, group.InfoLabel, id)),
		})
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
