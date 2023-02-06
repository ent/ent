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
	"entgo.io/ent/entc/integration/gremlin/ent/comment"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// CommentQuery is the builder for querying Comment entities.
type CommentQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Comment
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the CommentQuery builder.
func (cq *CommentQuery) Where(ps ...predicate.Comment) *CommentQuery {
	cq.predicates = append(cq.predicates, ps...)
	return cq
}

// Limit the number of records to be returned by this query.
func (cq *CommentQuery) Limit(limit int) *CommentQuery {
	cq.ctx.Limit = &limit
	return cq
}

// Offset to start from.
func (cq *CommentQuery) Offset(offset int) *CommentQuery {
	cq.ctx.Offset = &offset
	return cq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cq *CommentQuery) Unique(unique bool) *CommentQuery {
	cq.ctx.Unique = &unique
	return cq
}

// Order specifies how the records should be ordered.
func (cq *CommentQuery) Order(o ...OrderFunc) *CommentQuery {
	cq.order = append(cq.order, o...)
	return cq
}

// First returns the first Comment entity from the query.
// Returns a *NotFoundError when no Comment was found.
func (cq *CommentQuery) First(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(1).All(setContextOp(ctx, cq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{comment.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cq *CommentQuery) FirstX(ctx context.Context) *Comment {
	node, err := cq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Comment ID from the query.
// Returns a *NotFoundError when no Comment ID was found.
func (cq *CommentQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(1).IDs(setContextOp(ctx, cq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{comment.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cq *CommentQuery) FirstIDX(ctx context.Context) string {
	id, err := cq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Comment entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Comment entity is found.
// Returns a *NotFoundError when no Comment entities are found.
func (cq *CommentQuery) Only(ctx context.Context) (*Comment, error) {
	nodes, err := cq.Limit(2).All(setContextOp(ctx, cq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{comment.Label}
	default:
		return nil, &NotSingularError{comment.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cq *CommentQuery) OnlyX(ctx context.Context) *Comment {
	node, err := cq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Comment ID in the query.
// Returns a *NotSingularError when more than one Comment ID is found.
// Returns a *NotFoundError when no entities are found.
func (cq *CommentQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = cq.Limit(2).IDs(setContextOp(ctx, cq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{comment.Label}
	default:
		err = &NotSingularError{comment.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cq *CommentQuery) OnlyIDX(ctx context.Context) string {
	id, err := cq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Comments.
func (cq *CommentQuery) All(ctx context.Context) ([]*Comment, error) {
	ctx = setContextOp(ctx, cq.ctx, "All")
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Comment, *CommentQuery]()
	return withInterceptors[[]*Comment](ctx, cq, qr, cq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cq *CommentQuery) AllX(ctx context.Context) []*Comment {
	nodes, err := cq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Comment IDs.
func (cq *CommentQuery) IDs(ctx context.Context) (ids []string, err error) {
	if cq.ctx.Unique == nil && cq.path != nil {
		cq.Unique(true)
	}
	ctx = setContextOp(ctx, cq.ctx, "IDs")
	if err = cq.Select(comment.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cq *CommentQuery) IDsX(ctx context.Context) []string {
	ids, err := cq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cq *CommentQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cq.ctx, "Count")
	if err := cq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cq, querierCount[*CommentQuery](), cq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cq *CommentQuery) CountX(ctx context.Context) int {
	count, err := cq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cq *CommentQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cq.ctx, "Exist")
	switch _, err := cq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cq *CommentQuery) ExistX(ctx context.Context) bool {
	exist, err := cq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the CommentQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cq *CommentQuery) Clone() *CommentQuery {
	if cq == nil {
		return nil
	}
	return &CommentQuery{
		config:     cq.config,
		ctx:        cq.ctx.Clone(),
		order:      append([]OrderFunc{}, cq.order...),
		inters:     append([]Interceptor{}, cq.inters...),
		predicates: append([]predicate.Comment{}, cq.predicates...),
		// clone intermediate query.
		gremlin: cq.gremlin.Clone(),
		path:    cq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UniqueInt int `json:"unique_int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Comment.Query().
//		GroupBy(comment.FieldUniqueInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cq *CommentQuery) GroupBy(field string, fields ...string) *CommentGroupBy {
	cq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &CommentGroupBy{build: cq}
	grbuild.flds = &cq.ctx.Fields
	grbuild.label = comment.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UniqueInt int `json:"unique_int,omitempty"`
//	}
//
//	client.Comment.Query().
//		Select(comment.FieldUniqueInt).
//		Scan(ctx, &v)
func (cq *CommentQuery) Select(fields ...string) *CommentSelect {
	cq.ctx.Fields = append(cq.ctx.Fields, fields...)
	sbuild := &CommentSelect{CommentQuery: cq}
	sbuild.label = comment.Label
	sbuild.flds, sbuild.scan = &cq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a CommentSelect configured with the given aggregations.
func (cq *CommentQuery) Aggregate(fns ...AggregateFunc) *CommentSelect {
	return cq.Select().Aggregate(fns...)
}

func (cq *CommentQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cq); err != nil {
				return err
			}
		}
	}
	if cq.path != nil {
		prev, err := cq.path(ctx)
		if err != nil {
			return err
		}
		cq.gremlin = prev
	}
	return nil
}

func (cq *CommentQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*Comment, error) {
	res := &gremlin.Response{}
	traversal := cq.gremlinQuery(ctx)
	if len(cq.ctx.Fields) > 0 {
		fields := make([]any, len(cq.ctx.Fields))
		for i, f := range cq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := cq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var cs Comments
	if err := cs.FromResponse(res); err != nil {
		return nil, err
	}
	for i := range cs {
		cs[i].config = cq.config
	}
	return cs, nil
}

func (cq *CommentQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := cq.gremlinQuery(ctx).Count().Query()
	if err := cq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (cq *CommentQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(comment.Label)
	if cq.gremlin != nil {
		v = cq.gremlin.Clone()
	}
	for _, p := range cq.predicates {
		p(v)
	}
	if len(cq.order) > 0 {
		v.Order()
		for _, p := range cq.order {
			p(v)
		}
	}
	switch limit, offset := cq.ctx.Limit, cq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := cq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// CommentGroupBy is the group-by builder for Comment entities.
type CommentGroupBy struct {
	selector
	build *CommentQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cgb *CommentGroupBy) Aggregate(fns ...AggregateFunc) *CommentGroupBy {
	cgb.fns = append(cgb.fns, fns...)
	return cgb
}

// Scan applies the selector query and scans the result into the given value.
func (cgb *CommentGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cgb.build.ctx, "GroupBy")
	if err := cgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentGroupBy](ctx, cgb.build, cgb, cgb.build.inters, v)
}

func (cgb *CommentGroupBy) gremlinScan(ctx context.Context, root *CommentQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range cgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *cgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*cgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := cgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*cgb.flds)+len(cgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// CommentSelect is the builder for selecting fields of Comment entities.
type CommentSelect struct {
	*CommentQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cs *CommentSelect) Aggregate(fns ...AggregateFunc) *CommentSelect {
	cs.fns = append(cs.fns, fns...)
	return cs
}

// Scan applies the selector query and scans the result into the given value.
func (cs *CommentSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cs.ctx, "Select")
	if err := cs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*CommentQuery, *CommentSelect](ctx, cs.CommentQuery, cs, cs.inters, v)
}

func (cs *CommentSelect) gremlinScan(ctx context.Context, root *CommentQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := cs.ctx.Fields; len(fields) == 1 {
		if fields[0] != comment.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(cs.ctx.Fields))
		for i, f := range cs.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := cs.driver.Exec(ctx, query, bindings, res); err != nil {
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
