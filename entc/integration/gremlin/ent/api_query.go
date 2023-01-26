// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	context "context"
	fmt "fmt"
	math "math"

	gremlin "entgo.io/ent/dialect/gremlin"
	dsl "entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	g "entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/api"
	predicate "entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// APIQuery is the builder for querying Api entities.
type APIQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Api
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the APIQuery builder.
func (aq *APIQuery) Where(ps ...predicate.Api) *APIQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

// Limit the number of records to be returned by this query.
func (aq *APIQuery) Limit(limit int) *APIQuery {
	aq.ctx.Limit = &limit
	return aq
}

// Offset to start from.
func (aq *APIQuery) Offset(offset int) *APIQuery {
	aq.ctx.Offset = &offset
	return aq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (aq *APIQuery) Unique(unique bool) *APIQuery {
	aq.ctx.Unique = &unique
	return aq
}

// Order specifies how the records should be ordered.
func (aq *APIQuery) Order(o ...OrderFunc) *APIQuery {
	aq.order = append(aq.order, o...)
	return aq
}

// First returns the first Api entity from the query.
// Returns a *NotFoundError when no Api was found.
func (aq *APIQuery) First(ctx context.Context) (*Api, error) {
	nodes, err := aq.Limit(1).All(setContextOp(ctx, aq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{api.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (aq *APIQuery) FirstX(ctx context.Context) *Api {
	node, err := aq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Api ID from the query.
// Returns a *NotFoundError when no Api ID was found.
func (aq *APIQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(1).IDs(setContextOp(ctx, aq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{api.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (aq *APIQuery) FirstIDX(ctx context.Context) string {
	id, err := aq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Api entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Api entity is found.
// Returns a *NotFoundError when no Api entities are found.
func (aq *APIQuery) Only(ctx context.Context) (*Api, error) {
	nodes, err := aq.Limit(2).All(setContextOp(ctx, aq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{api.Label}
	default:
		return nil, &NotSingularError{api.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (aq *APIQuery) OnlyX(ctx context.Context) *Api {
	node, err := aq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Api ID in the query.
// Returns a *NotSingularError when more than one Api ID is found.
// Returns a *NotFoundError when no entities are found.
func (aq *APIQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = aq.Limit(2).IDs(setContextOp(ctx, aq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{api.Label}
	default:
		err = &NotSingularError{api.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (aq *APIQuery) OnlyIDX(ctx context.Context) string {
	id, err := aq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Apis.
func (aq *APIQuery) All(ctx context.Context) ([]*Api, error) {
	ctx = setContextOp(ctx, aq.ctx, "All")
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Api, *APIQuery]()
	return withInterceptors[[]*Api](ctx, aq, qr, aq.inters)
}

// AllX is like All, but panics if an error occurs.
func (aq *APIQuery) AllX(ctx context.Context) []*Api {
	nodes, err := aq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Api IDs.
func (aq *APIQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = setContextOp(ctx, aq.ctx, "IDs")
	if err := aq.Select(api.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (aq *APIQuery) IDsX(ctx context.Context) []string {
	ids, err := aq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (aq *APIQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, aq.ctx, "Count")
	if err := aq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, aq, querierCount[*APIQuery](), aq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (aq *APIQuery) CountX(ctx context.Context) int {
	count, err := aq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (aq *APIQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, aq.ctx, "Exist")
	switch _, err := aq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (aq *APIQuery) ExistX(ctx context.Context) bool {
	exist, err := aq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the APIQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (aq *APIQuery) Clone() *APIQuery {
	if aq == nil {
		return nil
	}
	return &APIQuery{
		config:     aq.config,
		ctx:        aq.ctx.Clone(),
		order:      append([]OrderFunc{}, aq.order...),
		inters:     append([]Interceptor{}, aq.inters...),
		predicates: append([]predicate.Api{}, aq.predicates...),
		// clone intermediate query.
		gremlin: aq.gremlin.Clone(),
		path:    aq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (aq *APIQuery) GroupBy(field string, fields ...string) *APIGroupBy {
	aq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &APIGroupBy{build: aq}
	grbuild.flds = &aq.ctx.Fields
	grbuild.label = api.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (aq *APIQuery) Select(fields ...string) *APISelect {
	aq.ctx.Fields = append(aq.ctx.Fields, fields...)
	sbuild := &APISelect{APIQuery: aq}
	sbuild.label = api.Label
	sbuild.flds, sbuild.scan = &aq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a APISelect configured with the given aggregations.
func (aq *APIQuery) Aggregate(fns ...AggregateFunc) *APISelect {
	return aq.Select().Aggregate(fns...)
}

func (aq *APIQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range aq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, aq); err != nil {
				return err
			}
		}
	}
	if aq.path != nil {
		prev, err := aq.path(ctx)
		if err != nil {
			return err
		}
		aq.gremlin = prev
	}
	return nil
}

func (aq *APIQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Api, error) {
	res := &gremlin.Response{}
	traversal := aq.gremlinQuery(ctx)
	if len(aq.ctx.Fields) > 0 {
		fields := make([]any, len(aq.ctx.Fields))
		for i, f := range aq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := aq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var as Apis
	if err := as.FromResponse(res); err != nil {
		return nil, err
	}
	as.config(aq.config)
	return as, nil
}

func (aq *APIQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := aq.gremlinQuery(ctx).Count().Query()
	if err := aq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (aq *APIQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(api.Label)
	if aq.gremlin != nil {
		v = aq.gremlin.Clone()
	}
	for _, p := range aq.predicates {
		p(v)
	}
	if len(aq.order) > 0 {
		v.Order()
		for _, p := range aq.order {
			p(v)
		}
	}
	switch limit, offset := aq.ctx.Limit, aq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := aq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// APIGroupBy is the group-by builder for Api entities.
type APIGroupBy struct {
	selector
	build *APIQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (agb *APIGroupBy) Aggregate(fns ...AggregateFunc) *APIGroupBy {
	agb.fns = append(agb.fns, fns...)
	return agb
}

// Scan applies the selector query and scans the result into the given value.
func (agb *APIGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, agb.build.ctx, "GroupBy")
	if err := agb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*APIQuery, *APIGroupBy](ctx, agb.build, agb, agb.build.inters, v)
}

func (agb *APIGroupBy) gremlinScan(ctx context.Context, root *APIQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range agb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *agb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*agb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := agb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*agb.flds)+len(agb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// APISelect is the builder for selecting fields of API entities.
type APISelect struct {
	*APIQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (as *APISelect) Aggregate(fns ...AggregateFunc) *APISelect {
	as.fns = append(as.fns, fns...)
	return as
}

// Scan applies the selector query and scans the result into the given value.
func (as *APISelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, as.ctx, "Select")
	if err := as.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*APIQuery, *APISelect](ctx, as.APIQuery, as, as.inters, v)
}

func (as *APISelect) gremlinScan(ctx context.Context, root *APIQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := as.ctx.Fields; len(fields) == 1 {
		if fields[0] != api.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(as.ctx.Fields))
		for i, f := range as.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := as.driver.Exec(ctx, query, bindings, res); err != nil {
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
