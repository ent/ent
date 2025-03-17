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
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
	enttask "entgo.io/ent/entc/integration/gremlin/ent/task"
)

// TaskQuery is the builder for querying Task entities.
type TaskQuery struct {
	config
	ctx        *QueryContext
	order      []enttask.OrderOption
	inters     []Interceptor
	predicates []predicate.Task
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the TaskQuery builder.
func (_q *TaskQuery) Where(ps ...predicate.Task) *TaskQuery {
	_q.predicates = append(_q.predicates, ps...)
	return _q
}

// Limit the number of records to be returned by this query.
func (_q *TaskQuery) Limit(limit int) *TaskQuery {
	_q.ctx.Limit = &limit
	return _q
}

// Offset to start from.
func (_q *TaskQuery) Offset(offset int) *TaskQuery {
	_q.ctx.Offset = &offset
	return _q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (_q *TaskQuery) Unique(unique bool) *TaskQuery {
	_q.ctx.Unique = &unique
	return _q
}

// Order specifies how the records should be ordered.
func (_q *TaskQuery) Order(o ...enttask.OrderOption) *TaskQuery {
	_q.order = append(_q.order, o...)
	return _q
}

// First returns the first Task entity from the query.
// Returns a *NotFoundError when no Task was found.
func (_q *TaskQuery) First(ctx context.Context) (*Task, error) {
	nodes, err := _q.Limit(1).All(setContextOp(ctx, _q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{enttask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (_q *TaskQuery) FirstX(ctx context.Context) *Task {
	node, err := _q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Task ID from the query.
// Returns a *NotFoundError when no Task ID was found.
func (_q *TaskQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(1).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{enttask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (_q *TaskQuery) FirstIDX(ctx context.Context) string {
	id, err := _q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Task entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Task entity is found.
// Returns a *NotFoundError when no Task entities are found.
func (_q *TaskQuery) Only(ctx context.Context) (*Task, error) {
	nodes, err := _q.Limit(2).All(setContextOp(ctx, _q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{enttask.Label}
	default:
		return nil, &NotSingularError{enttask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (_q *TaskQuery) OnlyX(ctx context.Context) *Task {
	node, err := _q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Task ID in the query.
// Returns a *NotSingularError when more than one Task ID is found.
// Returns a *NotFoundError when no entities are found.
func (_q *TaskQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = _q.Limit(2).IDs(setContextOp(ctx, _q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{enttask.Label}
	default:
		err = &NotSingularError{enttask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (_q *TaskQuery) OnlyIDX(ctx context.Context) string {
	id, err := _q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Tasks.
func (_q *TaskQuery) All(ctx context.Context) ([]*Task, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryAll)
	if err := _q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Task, *TaskQuery]()
	return withInterceptors[[]*Task](ctx, _q, qr, _q.inters)
}

// AllX is like All, but panics if an error occurs.
func (_q *TaskQuery) AllX(ctx context.Context) []*Task {
	nodes, err := _q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Task IDs.
func (_q *TaskQuery) IDs(ctx context.Context) (ids []string, err error) {
	if _q.ctx.Unique == nil && _q.path != nil {
		_q.Unique(true)
	}
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryIDs)
	if err = _q.Select(enttask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (_q *TaskQuery) IDsX(ctx context.Context) []string {
	ids, err := _q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (_q *TaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryCount)
	if err := _q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, _q, querierCount[*TaskQuery](), _q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (_q *TaskQuery) CountX(ctx context.Context) int {
	count, err := _q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (_q *TaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, _q.ctx, ent.OpQueryExist)
	switch _, err := _q.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (_q *TaskQuery) ExistX(ctx context.Context) bool {
	exist, err := _q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (_q *TaskQuery) Clone() *TaskQuery {
	if _q == nil {
		return nil
	}
	return &TaskQuery{
		config:     _q.config,
		ctx:        _q.ctx.Clone(),
		order:      append([]enttask.OrderOption{}, _q.order...),
		inters:     append([]Interceptor{}, _q.inters...),
		predicates: append([]predicate.Task{}, _q.predicates...),
		// clone intermediate query.
		gremlin: _q.gremlin.Clone(),
		path:    _q.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Priority task.Priority `json:"priority,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Task.Query().
//		GroupBy(enttask.FieldPriority).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (_q *TaskQuery) GroupBy(field string, fields ...string) *TaskGroupBy {
	_q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TaskGroupBy{build: _q}
	grbuild.flds = &_q.ctx.Fields
	grbuild.label = enttask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Priority task.Priority `json:"priority,omitempty"`
//	}
//
//	client.Task.Query().
//		Select(enttask.FieldPriority).
//		Scan(ctx, &v)
func (_q *TaskQuery) Select(fields ...string) *TaskSelect {
	_q.ctx.Fields = append(_q.ctx.Fields, fields...)
	sbuild := &TaskSelect{TaskQuery: _q}
	sbuild.label = enttask.Label
	sbuild.flds, sbuild.scan = &_q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TaskSelect configured with the given aggregations.
func (_q *TaskQuery) Aggregate(fns ...AggregateFunc) *TaskSelect {
	return _q.Select().Aggregate(fns...)
}

func (_q *TaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range _q.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, _q); err != nil {
				return err
			}
		}
	}
	if _q.path != nil {
		prev, err := _q.path(ctx)
		if err != nil {
			return err
		}
		_q.gremlin = prev
	}
	return nil
}

func (_q *TaskQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Task, error) {
	res := &gremlin.Response{}
	traversal := _q.gremlinQuery(ctx)
	if len(_q.ctx.Fields) > 0 {
		fields := make([]any, len(_q.ctx.Fields))
		for i, f := range _q.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _ms Tasks
	if err := _ms.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range _ms {
		_ms[i].config = _q.config
	}
	return _ms, nil
}

func (_q *TaskQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := _q.gremlinQuery(ctx).Count().Query()
	if err := _q.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (_q *TaskQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(enttask.Label)
	if _q.gremlin != nil {
		v = _q.gremlin.Clone()
	}
	for _, p := range _q.predicates {
		p(v)
	}
	if len(_q.order) > 0 {
		v.Order()
		for _, p := range _q.order {
			p(v)
		}
	}
	switch limit, offset := _q.ctx.Limit, _q.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := _q.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// TaskGroupBy is the group-by builder for Task entities.
type TaskGroupBy struct {
	selector
	build *TaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tgb *TaskGroupBy) Aggregate(fns ...AggregateFunc) *TaskGroupBy {
	tgb.fns = append(tgb.fns, fns...)
	return tgb
}

// Scan applies the selector query and scans the result into the given value.
func (tgb *TaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tgb.build.ctx, ent.OpQueryGroupBy)
	if err := tgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskGroupBy](ctx, tgb.build, tgb, tgb.build.inters, v)
}

func (tgb *TaskGroupBy) gremlinScan(ctx context.Context, root *TaskQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range tgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *tgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*tgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := tgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*tgb.flds)+len(tgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// TaskSelect is the builder for selecting fields of Task entities.
type TaskSelect struct {
	*TaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ts *TaskSelect) Aggregate(fns ...AggregateFunc) *TaskSelect {
	ts.fns = append(ts.fns, fns...)
	return ts
}

// Scan applies the selector query and scans the result into the given value.
func (ts *TaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ts.ctx, ent.OpQuerySelect)
	if err := ts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TaskQuery, *TaskSelect](ctx, ts.TaskQuery, ts, ts.inters, v)
}

func (ts *TaskSelect) gremlinScan(ctx context.Context, root *TaskQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := ts.ctx.Fields; len(fields) == 1 {
		if fields[0] != enttask.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(ts.ctx.Fields))
		for i, f := range ts.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := ts.driver.Exec(ctx, query, bindings, res); err != nil {
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
