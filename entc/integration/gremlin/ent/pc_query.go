// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/pc"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// PCQuery is the builder for querying PC entities.
type PCQuery struct {
	config
	ctx        *QueryContext
	order      []pc.OrderOption
	inters     []Interceptor
	predicates []predicate.PC
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the PCQuery builder.
func (q *PCQuery) Where(ps ...predicate.PC) *PCQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *PCQuery) Limit(limit int) *PCQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *PCQuery) Offset(offset int) *PCQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *PCQuery) Unique(unique bool) *PCQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *PCQuery) Order(o ...pc.OrderOption) *PCQuery {
	q.order = append(q.order, o...)
	return q
}

// First returns the first PC entity from the query.
// Returns a *NotFoundError when no PC was found.
func (q *PCQuery) First(ctx context.Context) (*PC, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pc.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *PCQuery) FirstX(ctx context.Context) *PC {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PC ID from the query.
// Returns a *NotFoundError when no PC ID was found.
func (q *PCQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pc.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *PCQuery) FirstIDX(ctx context.Context) string {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PC entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PC entity is found.
// Returns a *NotFoundError when no PC entities are found.
func (q *PCQuery) Only(ctx context.Context) (*PC, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pc.Label}
	default:
		return nil, &NotSingularError{pc.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *PCQuery) OnlyX(ctx context.Context) *PC {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PC ID in the query.
// Returns a *NotSingularError when more than one PC ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *PCQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pc.Label}
	default:
		err = &NotSingularError{pc.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *PCQuery) OnlyIDX(ctx context.Context) string {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PCs.
func (q *PCQuery) All(ctx context.Context) ([]*PC, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PC, *PCQuery]()
	return withInterceptors[[]*PC](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *PCQuery) AllX(ctx context.Context) []*PC {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PC IDs.
func (q *PCQuery) IDs(ctx context.Context) (ids []string, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(pc.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *PCQuery) IDsX(ctx context.Context) []string {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *PCQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*PCQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *PCQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *PCQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryExist)
	switch _, err := q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (q *PCQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PCQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *PCQuery) Clone() *PCQuery {
	if q == nil {
		return nil
	}
	return &PCQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]pc.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.PC{}, q.predicates...),
		// clone intermediate query.
		gremlin: q.gremlin.Clone(),
		path:    q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (q *PCQuery) GroupBy(field string, fields ...string) *PCGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PCGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = pc.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (q *PCQuery) Select(fields ...string) *PCSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &PCSelect{PCQuery: q}
	sbuild.label = pc.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PCSelect configured with the given aggregations.
func (q *PCQuery) Aggregate(fns ...AggregateFunc) *PCSelect {
	return q.Select().Aggregate(fns...)
}

func (q *PCQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, q); err != nil {
				return err
			}
		}
	}
	if q.path != nil {
		prev, err := q.path(ctx)
		if err != nil {
			return err
		}
		q.gremlin = prev
	}
	return nil
}

func (q *PCQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*PC, error) {
	res := &gremlin.Response{}
	traversal := q.gremlinQuery(ctx)
	if len(q.ctx.Fields) > 0 {
		fields := make([]any, len(q.ctx.Fields))
		for i, f := range q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var results PCs
	if err := results.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range results {
		results[i].config = q.config
	}
	return results, nil
}

func (q *PCQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := q.gremlinQuery(ctx).Count().Query()
	if err := q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (q *PCQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(pc.Label)
	if q.gremlin != nil {
		v = q.gremlin.Clone()
	}
	for _, p := range q.predicates {
		p(v)
	}
	if len(q.order) > 0 {
		v.Order()
		for _, p := range q.order {
			p(v)
		}
	}
	switch limit, offset := q.ctx.Limit, q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// PCGroupBy is the group-by builder for PC entities.
type PCGroupBy struct {
	selector
	build *PCQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PCGroupBy) Aggregate(fns ...AggregateFunc) *PCGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PCGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, ent.OpQueryGroupBy)
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PCQuery, *PCGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PCGroupBy) gremlinScan(ctx context.Context, root *PCQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range pgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *pgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*pgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := pgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*pgb.flds)+len(pgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// PCSelect is the builder for selecting fields of PC entities.
type PCSelect struct {
	*PCQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PCSelect) Aggregate(fns ...AggregateFunc) *PCSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PCSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, ent.OpQuerySelect)
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PCQuery, *PCSelect](ctx, ps.PCQuery, ps, ps.inters, v)
}

func (ps *PCSelect) gremlinScan(ctx context.Context, root *PCQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ps.ctx.Fields; len(fields) == 1 {
		if fields[0] != pc.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ps.ctx.Fields))
		for i, f := range ps.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ps.driver.Exec(ctx, query, bindings, res); err != nil {
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
