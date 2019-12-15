// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated (@generated) by entc, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/p"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/group"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
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
	return giu.gremlinSave(ctx)
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
	return giuo.gremlinSave(ctx)
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
