// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	"entgo.io/ent/entc/integration/gremlin/ent/socialprofile"
	"entgo.io/ent/entc/integration/gremlin/ent/user"
)

// SocialProfileQuery is the builder for querying SocialProfile entities.
type SocialProfileQuery struct {
	config
	ctx        *QueryContext
	order      []socialprofile.OrderOption
	inters     []Interceptor
	predicates []predicate.SocialProfile
	withUser   *UserQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the SocialProfileQuery builder.
func (spq *SocialProfileQuery) Where(ps ...predicate.SocialProfile) *SocialProfileQuery {
	spq.predicates = append(spq.predicates, ps...)
	return spq
}

// Limit the number of records to be returned by this query.
func (spq *SocialProfileQuery) Limit(limit int) *SocialProfileQuery {
	spq.ctx.Limit = &limit
	return spq
}

// Offset to start from.
func (spq *SocialProfileQuery) Offset(offset int) *SocialProfileQuery {
	spq.ctx.Offset = &offset
	return spq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (spq *SocialProfileQuery) Unique(unique bool) *SocialProfileQuery {
	spq.ctx.Unique = &unique
	return spq
}

// Order specifies how the records should be ordered.
func (spq *SocialProfileQuery) Order(o ...socialprofile.OrderOption) *SocialProfileQuery {
	spq.order = append(spq.order, o...)
	return spq
}

// QueryUser chains the current query on the "user" edge.
func (spq *SocialProfileQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: spq.config}).Query()
	query.path = func(ctx context.Context) (fromU *dsl.Traversal, err error) {
		if err := spq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		gremlin := spq.gremlinQuery(ctx)
		fromU = gremlin.InE(user.SocialProfilesLabel).OutV()
		return fromU, nil
	}
	return query
}

// First returns the first SocialProfile entity from the query.
// Returns a *NotFoundError when no SocialProfile was found.
func (spq *SocialProfileQuery) First(ctx context.Context) (*SocialProfile, error) {
	nodes, err := spq.Limit(1).All(setContextOp(ctx, spq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{socialprofile.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (spq *SocialProfileQuery) FirstX(ctx context.Context) *SocialProfile {
	node, err := spq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first SocialProfile ID from the query.
// Returns a *NotFoundError when no SocialProfile ID was found.
func (spq *SocialProfileQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = spq.Limit(1).IDs(setContextOp(ctx, spq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{socialprofile.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (spq *SocialProfileQuery) FirstIDX(ctx context.Context) string {
	id, err := spq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single SocialProfile entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one SocialProfile entity is found.
// Returns a *NotFoundError when no SocialProfile entities are found.
func (spq *SocialProfileQuery) Only(ctx context.Context) (*SocialProfile, error) {
	nodes, err := spq.Limit(2).All(setContextOp(ctx, spq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{socialprofile.Label}
	default:
		return nil, &NotSingularError{socialprofile.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (spq *SocialProfileQuery) OnlyX(ctx context.Context) *SocialProfile {
	node, err := spq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only SocialProfile ID in the query.
// Returns a *NotSingularError when more than one SocialProfile ID is found.
// Returns a *NotFoundError when no entities are found.
func (spq *SocialProfileQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = spq.Limit(2).IDs(setContextOp(ctx, spq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{socialprofile.Label}
	default:
		err = &NotSingularError{socialprofile.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (spq *SocialProfileQuery) OnlyIDX(ctx context.Context) string {
	id, err := spq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of SocialProfiles.
func (spq *SocialProfileQuery) All(ctx context.Context) ([]*SocialProfile, error) {
	ctx = setContextOp(ctx, spq.ctx, "All")
	if err := spq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*SocialProfile, *SocialProfileQuery]()
	return withInterceptors[[]*SocialProfile](ctx, spq, qr, spq.inters)
}

// AllX is like All, but panics if an error occurs.
func (spq *SocialProfileQuery) AllX(ctx context.Context) []*SocialProfile {
	nodes, err := spq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of SocialProfile IDs.
func (spq *SocialProfileQuery) IDs(ctx context.Context) (ids []string, err error) {
	if spq.ctx.Unique == nil && spq.path != nil {
		spq.Unique(true)
	}
	ctx = setContextOp(ctx, spq.ctx, "IDs")
	if err = spq.Select(socialprofile.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (spq *SocialProfileQuery) IDsX(ctx context.Context) []string {
	ids, err := spq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (spq *SocialProfileQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, spq.ctx, "Count")
	if err := spq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, spq, querierCount[*SocialProfileQuery](), spq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (spq *SocialProfileQuery) CountX(ctx context.Context) int {
	count, err := spq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (spq *SocialProfileQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, spq.ctx, "Exist")
	switch _, err := spq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (spq *SocialProfileQuery) ExistX(ctx context.Context) bool {
	exist, err := spq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SocialProfileQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (spq *SocialProfileQuery) Clone() *SocialProfileQuery {
	if spq == nil {
		return nil
	}
	return &SocialProfileQuery{
		config:     spq.config,
		ctx:        spq.ctx.Clone(),
		order:      append([]socialprofile.OrderOption{}, spq.order...),
		inters:     append([]Interceptor{}, spq.inters...),
		predicates: append([]predicate.SocialProfile{}, spq.predicates...),
		withUser:   spq.withUser.Clone(),
		// clone intermediate query.
		gremlin: spq.gremlin.Clone(),
		path:    spq.path,
	}
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (spq *SocialProfileQuery) WithUser(opts ...func(*UserQuery)) *SocialProfileQuery {
	query := (&UserClient{config: spq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	spq.withUser = query
	return spq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.SocialProfile.Query().
//		GroupBy(socialprofile.FieldDesc).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (spq *SocialProfileQuery) GroupBy(field string, fields ...string) *SocialProfileGroupBy {
	spq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SocialProfileGroupBy{build: spq}
	grbuild.flds = &spq.ctx.Fields
	grbuild.label = socialprofile.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Desc string `json:"desc,omitempty"`
//	}
//
//	client.SocialProfile.Query().
//		Select(socialprofile.FieldDesc).
//		Scan(ctx, &v)
func (spq *SocialProfileQuery) Select(fields ...string) *SocialProfileSelect {
	spq.ctx.Fields = append(spq.ctx.Fields, fields...)
	sbuild := &SocialProfileSelect{SocialProfileQuery: spq}
	sbuild.label = socialprofile.Label
	sbuild.flds, sbuild.scan = &spq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SocialProfileSelect configured with the given aggregations.
func (spq *SocialProfileQuery) Aggregate(fns ...AggregateFunc) *SocialProfileSelect {
	return spq.Select().Aggregate(fns...)
}

func (spq *SocialProfileQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range spq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, spq); err != nil {
				return err
			}
		}
	}
	if spq.path != nil {
		prev, err := spq.path(ctx)
		if err != nil {
			return err
		}
		spq.gremlin = prev
	}
	return nil
}

func (spq *SocialProfileQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*SocialProfile, error) {
	res := &gremlin.Response{}
	traversal := spq.gremlinQuery(ctx)
	if len(spq.ctx.Fields) > 0 {
		fields := make([]any, len(spq.ctx.Fields))
		for i, f := range spq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := spq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var sps SocialProfiles
	if err := sps.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range sps {
		sps[i].config = spq.config
	}
	return sps, nil
}

func (spq *SocialProfileQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := spq.gremlinQuery(ctx).Count().Query()
	if err := spq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (spq *SocialProfileQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(socialprofile.Label)
	if spq.gremlin != nil {
		v = spq.gremlin.Clone()
	}
	for _, p := range spq.predicates {
		p(v)
	}
	if len(spq.order) > 0 {
		v.Order()
		for _, p := range spq.order {
			p(v)
		}
	}
	switch limit, offset := spq.ctx.Limit, spq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := spq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// SocialProfileGroupBy is the group-by builder for SocialProfile entities.
type SocialProfileGroupBy struct {
	selector
	build *SocialProfileQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (spgb *SocialProfileGroupBy) Aggregate(fns ...AggregateFunc) *SocialProfileGroupBy {
	spgb.fns = append(spgb.fns, fns...)
	return spgb
}

// Scan applies the selector query and scans the result into the given value.
func (spgb *SocialProfileGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, spgb.build.ctx, "GroupBy")
	if err := spgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SocialProfileQuery, *SocialProfileGroupBy](ctx, spgb.build, spgb, spgb.build.inters, v)
}

func (spgb *SocialProfileGroupBy) gremlinScan(ctx context.Context, root *SocialProfileQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range spgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *spgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*spgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := spgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*spgb.flds)+len(spgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// SocialProfileSelect is the builder for selecting fields of SocialProfile entities.
type SocialProfileSelect struct {
	*SocialProfileQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (sps *SocialProfileSelect) Aggregate(fns ...AggregateFunc) *SocialProfileSelect {
	sps.fns = append(sps.fns, fns...)
	return sps
}

// Scan applies the selector query and scans the result into the given value.
func (sps *SocialProfileSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sps.ctx, "Select")
	if err := sps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SocialProfileQuery, *SocialProfileSelect](ctx, sps.SocialProfileQuery, sps, sps.inters, v)
}

func (sps *SocialProfileSelect) gremlinScan(ctx context.Context, root *SocialProfileQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := sps.ctx.Fields; len(fields) == 1 {
		if fields[0] != socialprofile.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(sps.ctx.Fields))
		for i, f := range sps.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := sps.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(root.ctx.Fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
