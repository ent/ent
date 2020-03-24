// Copyright (c) Facebook, Inc. and its affiliates. All Rights Reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math"

	"github.com/facebookincubator/ent/dialect/gremlin"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/__"
	"github.com/facebookincubator/ent/dialect/gremlin/graph/dsl/g"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/group"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/groupinfo"
	"github.com/facebookincubator/ent/entc/integration/gremlin/ent/predicate"
)

// GroupInfoQuery is the builder for querying GroupInfo entities.
type GroupInfoQuery struct {
	config
	err        error
	limit      *int
	offset     *int
	order      []Order
	unique     []string
	predicates []predicate.GroupInfo
	// eager-loading edges.
	withGroups *GroupQuery
	// intermediate query.
	gremlin *dsl.Traversal
}

// Where adds a new predicate for the builder.
func (giq *GroupInfoQuery) Where(ps ...predicate.GroupInfo) *GroupInfoQuery {
	giq.predicates = append(giq.predicates, ps...)
	return giq
}

// Limit adds a limit step to the query.
func (giq *GroupInfoQuery) Limit(limit int) *GroupInfoQuery {
	giq.limit = &limit
	return giq
}

// Offset adds an offset step to the query.
func (giq *GroupInfoQuery) Offset(offset int) *GroupInfoQuery {
	giq.offset = &offset
	return giq
}

// Order adds an order step to the query.
func (giq *GroupInfoQuery) Order(o ...Order) *GroupInfoQuery {
	giq.order = append(giq.order, o...)
	return giq
}

// QueryGroups chains the current query on the groups edge.
func (giq *GroupInfoQuery) QueryGroups() *GroupQuery {
	query := &GroupQuery{
		config: giq.config,
		err:    giq.err,
	}
	gremlin := giq.gremlinQuery()
	query.gremlin = gremlin.InE(group.InfoLabel).OutV()
	return query
}

// First returns the first GroupInfo entity in the query. Returns *NotFoundError when no groupinfo was found.
func (giq *GroupInfoQuery) First(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(gis) == 0 {
		return nil, &NotFoundError{groupinfo.Label}
	}
	return gis[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstX(ctx context.Context) *GroupInfo {
	gi, err := giq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return gi
}

// FirstID returns the first GroupInfo id in the query. Returns *NotFoundError when no id was found.
func (giq *GroupInfoQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{groupinfo.Label}
		return
	}
	return ids[0], nil
}

// FirstXID is like FirstID, but panics if an error occurs.
func (giq *GroupInfoQuery) FirstXID(ctx context.Context) string {
	id, err := giq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns the only GroupInfo entity in the query, returns an error if not exactly one entity was returned.
func (giq *GroupInfoQuery) Only(ctx context.Context) (*GroupInfo, error) {
	gis, err := giq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(gis) {
	case 1:
		return gis[0], nil
	case 0:
		return nil, &NotFoundError{groupinfo.Label}
	default:
		return nil, &NotSingularError{groupinfo.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyX(ctx context.Context) *GroupInfo {
	gi, err := giq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return gi
}

// OnlyID returns the only GroupInfo id in the query, returns an error if not exactly one id was returned.
func (giq *GroupInfoQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = giq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{groupinfo.Label}
	default:
		err = &NotSingularError{groupinfo.Label}
	}
	return
}

// OnlyXID is like OnlyID, but panics if an error occurs.
func (giq *GroupInfoQuery) OnlyXID(ctx context.Context) string {
	id, err := giq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GroupInfos.
func (giq *GroupInfoQuery) All(ctx context.Context) ([]*GroupInfo, error) {
	if giq.err != nil {
		return nil, giq.err
	}
	return giq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (giq *GroupInfoQuery) AllX(ctx context.Context) []*GroupInfo {
	gis, err := giq.All(ctx)
	if err != nil {
		panic(err)
	}
	return gis
}

// IDs executes the query and returns a list of GroupInfo ids.
func (giq *GroupInfoQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := giq.Select(groupinfo.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (giq *GroupInfoQuery) IDsX(ctx context.Context) []string {
	ids, err := giq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (giq *GroupInfoQuery) Count(ctx context.Context) (int, error) {
	if giq.err != nil {
		return 0, giq.err
	}
	return giq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (giq *GroupInfoQuery) CountX(ctx context.Context) int {
	count, err := giq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (giq *GroupInfoQuery) Exist(ctx context.Context) (bool, error) {
	if giq.err != nil {
		return false, giq.err
	}
	return giq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (giq *GroupInfoQuery) ExistX(ctx context.Context) bool {
	exist, err := giq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the query builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (giq *GroupInfoQuery) Clone() *GroupInfoQuery {
	return &GroupInfoQuery{
		config:     giq.config,
		err:        giq.err,
		limit:      giq.limit,
		offset:     giq.offset,
		order:      append([]Order(nil), giq.order...),
		unique:     append([]string(nil), giq.unique...),
		predicates: append([]predicate.GroupInfo(nil), giq.predicates...),
		// clone intermediate query.
		gremlin: giq.gremlin.Clone(),
	}
}

//  WithGroups tells the query-builder to eager-loads the nodes that are connected to
// the "groups" edge. The optional arguments used to configure the query builder of the edge.
func (giq *GroupInfoQuery) WithGroups(opts ...func(*GroupQuery)) *GroupInfoQuery {
	query := &GroupQuery{
		config: giq.config,
		err:    giq.err,
	}
	for _, opt := range opts {
		opt(query)
	}
	giq.withGroups = query
	return giq
}

// GroupBy used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		GroupBy(groupinfo.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) GroupBy(field string, fields ...string) *GroupInfoGroupBy {
	group := &GroupInfoGroupBy{
		config: giq.config,
		err:    giq.err,
		fields: append([]string{field}, fields...),
	}
	group.gremlin = giq.gremlinQuery()
	return group
}

// Select one or more fields from the given query.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//	}
//
//	client.GroupInfo.Query().
//		Select(groupinfo.FieldDesc).
//		Scan(ctx, &v)
//
func (giq *GroupInfoQuery) Select(field string, fields ...string) *GroupInfoSelect {
	selector := &GroupInfoSelect{
		config: giq.config,
		err:    giq.err,
		fields: append([]string{field}, fields...),
	}
	selector.gremlin = giq.gremlinQuery()
	return selector
}

func (giq *GroupInfoQuery) gremlinAll(ctx context.Context) ([]*GroupInfo, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().ValueMap(true).Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var gis GroupInfos
	if err := gis.FromResponse(res); err != nil {
		return nil, err
	}
	gis.config(giq.config)
	return gis, nil
}

func (giq *GroupInfoQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().Count().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (giq *GroupInfoQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := giq.gremlinQuery().HasNext().Query()
	if err := giq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (giq *GroupInfoQuery) gremlinQuery() *dsl.Traversal {
	v := g.V().HasLabel(groupinfo.Label)
	if giq.gremlin != nil {
		v = giq.gremlin.Clone()
	}
	for _, p := range giq.predicates {
		p(v)
	}
	if len(giq.order) > 0 {
		v.Order()
		for _, p := range giq.order {
			p(v)
		}
	}
	switch limit, offset := giq.limit, giq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := giq.unique; len(unique) == 0 {
		v.Dedup()
	}
	return v
}

// GroupInfoGroupBy is the builder for group-by GroupInfo entities.
type GroupInfoGroupBy struct {
	config
	err    error
	fields []string
	fns    []Aggregate
	// intermediate query.
	gremlin *dsl.Traversal
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gigb *GroupInfoGroupBy) Aggregate(fns ...Aggregate) *GroupInfoGroupBy {
	gigb.fns = append(gigb.fns, fns...)
	return gigb
}

// Scan applies the group-by query and scan the result into the given value.
func (gigb *GroupInfoGroupBy) Scan(ctx context.Context, v interface{}) error {
	if gigb.err != nil {
		return gigb.err
	}
	return gigb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gigb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) StringsX(ctx context.Context) []string {
	v, err := gigb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) IntsX(ctx context.Context) []int {
	v, err := gigb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gigb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by. It is only allowed when querying group-by with one field.
func (gigb *GroupInfoGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gigb.fields) > 1 {
		return nil, errors.New("ent: GroupInfoGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gigb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gigb *GroupInfoGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gigb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gigb *GroupInfoGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := gigb.gremlinQuery().Query()
	if err := gigb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(gigb.fields)+len(gigb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (gigb *GroupInfoGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range gigb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range gigb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return gigb.gremlin.Group().
		By(__.Values(gigb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// GroupInfoSelect is the builder for select fields of GroupInfo entities.
type GroupInfoSelect struct {
	config
	err    error
	fields []string
	// intermediate queries.
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scan the result into the given value.
func (gis *GroupInfoSelect) Scan(ctx context.Context, v interface{}) error {
	if gis.err != nil {
		return gis.err
	}
	return gis.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gis *GroupInfoSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gis.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gis *GroupInfoSelect) StringsX(ctx context.Context) []string {
	v, err := gis.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gis *GroupInfoSelect) IntsX(ctx context.Context) []int {
	v, err := gis.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gis *GroupInfoSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gis.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from selector. It is only allowed when selecting one field.
func (gis *GroupInfoSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gis.fields) > 1 {
		return nil, errors.New("ent: GroupInfoSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gis.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gis *GroupInfoSelect) BoolsX(ctx context.Context) []bool {
	v, err := gis.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gis *GroupInfoSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(gis.fields) == 1 {
		if gis.fields[0] != groupinfo.FieldID {
			traversal = gis.gremlin.Values(gis.fields...)
		} else {
			traversal = gis.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(gis.fields))
		for i, f := range gis.fields {
			fields[i] = f
		}
		traversal = gis.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := gis.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(gis.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
