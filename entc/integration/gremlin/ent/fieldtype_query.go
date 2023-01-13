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
	"entgo.io/ent/entc/integration/gremlin/ent/fieldtype"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// FieldTypeQuery is the builder for querying FieldType entities.
type FieldTypeQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.FieldType
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the FieldTypeQuery builder.
func (ftq *FieldTypeQuery) Where(ps ...predicate.FieldType) *FieldTypeQuery {
	ftq.predicates = append(ftq.predicates, ps...)
	return ftq
}

// Limit the number of records to be returned by this query.
func (ftq *FieldTypeQuery) Limit(limit int) *FieldTypeQuery {
	ftq.ctx.Limit = &limit
	return ftq
}

// Offset to start from.
func (ftq *FieldTypeQuery) Offset(offset int) *FieldTypeQuery {
	ftq.ctx.Offset = &offset
	return ftq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ftq *FieldTypeQuery) Unique(unique bool) *FieldTypeQuery {
	ftq.ctx.Unique = &unique
	return ftq
}

// Order specifies how the records should be ordered.
func (ftq *FieldTypeQuery) Order(o ...OrderFunc) *FieldTypeQuery {
	ftq.order = append(ftq.order, o...)
	return ftq
}

// First returns the first FieldType entity from the query.
// Returns a *NotFoundError when no FieldType was found.
func (ftq *FieldTypeQuery) First(ctx context.Context) (*FieldType, error) {
	nodes, err := ftq.Limit(1).All(setContextOp(ctx, ftq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{fieldtype.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstX(ctx context.Context) *FieldType {
	node, err := ftq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first FieldType ID from the query.
// Returns a *NotFoundError when no FieldType ID was found.
func (ftq *FieldTypeQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(1).IDs(setContextOp(ctx, ftq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{fieldtype.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ftq *FieldTypeQuery) FirstIDX(ctx context.Context) string {
	id, err := ftq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single FieldType entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one FieldType entity is found.
// Returns a *NotFoundError when no FieldType entities are found.
func (ftq *FieldTypeQuery) Only(ctx context.Context) (*FieldType, error) {
	nodes, err := ftq.Limit(2).All(setContextOp(ctx, ftq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{fieldtype.Label}
	default:
		return nil, &NotSingularError{fieldtype.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyX(ctx context.Context) *FieldType {
	node, err := ftq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only FieldType ID in the query.
// Returns a *NotSingularError when more than one FieldType ID is found.
// Returns a *NotFoundError when no entities are found.
func (ftq *FieldTypeQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = ftq.Limit(2).IDs(setContextOp(ctx, ftq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{fieldtype.Label}
	default:
		err = &NotSingularError{fieldtype.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ftq *FieldTypeQuery) OnlyIDX(ctx context.Context) string {
	id, err := ftq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of FieldTypes.
func (ftq *FieldTypeQuery) All(ctx context.Context) ([]*FieldType, error) {
	ctx = setContextOp(ctx, ftq.ctx, "All")
	if err := ftq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*FieldType, *FieldTypeQuery]()
	return withInterceptors[[]*FieldType](ctx, ftq, qr, ftq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ftq *FieldTypeQuery) AllX(ctx context.Context) []*FieldType {
	nodes, err := ftq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of FieldType IDs.
func (ftq *FieldTypeQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	ctx = setContextOp(ctx, ftq.ctx, "IDs")
	if err := ftq.Select(fieldtype.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ftq *FieldTypeQuery) IDsX(ctx context.Context) []string {
	ids, err := ftq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ftq *FieldTypeQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Count")
	if err := ftq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ftq, querierCount[*FieldTypeQuery](), ftq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ftq *FieldTypeQuery) CountX(ctx context.Context) int {
	count, err := ftq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ftq *FieldTypeQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ftq.ctx, "Exist")
	switch _, err := ftq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ftq *FieldTypeQuery) ExistX(ctx context.Context) bool {
	exist, err := ftq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FieldTypeQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ftq *FieldTypeQuery) Clone() *FieldTypeQuery {
	if ftq == nil {
		return nil
	}
	return &FieldTypeQuery{
		config:     ftq.config,
		ctx:        ftq.ctx.Clone(),
		order:      append([]OrderFunc{}, ftq.order...),
		inters:     append([]Interceptor{}, ftq.inters...),
		predicates: append([]predicate.FieldType{}, ftq.predicates...),
		// clone intermediate query.
		gremlin: ftq.gremlin.Clone(),
		path:    ftq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.FieldType.Query().
//		GroupBy(fieldtype.FieldInt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ftq *FieldTypeQuery) GroupBy(field string, fields ...string) *FieldTypeGroupBy {
	ftq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FieldTypeGroupBy{build: ftq}
	grbuild.flds = &ftq.ctx.Fields
	grbuild.label = fieldtype.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Int int `json:"int,omitempty"`
//	}
//
//	client.FieldType.Query().
//		Select(fieldtype.FieldInt).
//		Scan(ctx, &v)
func (ftq *FieldTypeQuery) Select(fields ...string) *FieldTypeSelect {
	ftq.ctx.Fields = append(ftq.ctx.Fields, fields...)
	sbuild := &FieldTypeSelect{FieldTypeQuery: ftq}
	sbuild.label = fieldtype.Label
	sbuild.flds, sbuild.scan = &ftq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FieldTypeSelect configured with the given aggregations.
func (ftq *FieldTypeQuery) Aggregate(fns ...AggregateFunc) *FieldTypeSelect {
	return ftq.Select().Aggregate(fns...)
}

func (ftq *FieldTypeQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ftq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ftq); err != nil {
				return err
			}
		}
	}
	if ftq.path != nil {
		prev, err := ftq.path(ctx)
		if err != nil {
			return err
		}
		ftq.gremlin = prev
	}
	return nil
}

func (ftq *FieldTypeQuery) gremlinAll(ctx context.Context, hooks ...queryHook) ([]*FieldType, error) {
	res := &gremlin.Response{}
	traversal := ftq.gremlinQuery(ctx)
	if len(ftq.ctx.Fields) > 0 {
		fields := make([]any, len(ftq.ctx.Fields))
		for i, f := range ftq.ctx.Fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var fts FieldTypes
	if err := fts.FromResponse(res); err != nil {
		return nil, err
	}
	fts.config(ftq.config)
	return fts, nil
}

func (ftq *FieldTypeQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := ftq.gremlinQuery(ctx).Count().Query()
	if err := ftq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (ftq *FieldTypeQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(fieldtype.Label)
	if ftq.gremlin != nil {
		v = ftq.gremlin.Clone()
	}
	for _, p := range ftq.predicates {
		p(v)
	}
	if len(ftq.order) > 0 {
		v.Order()
		for _, p := range ftq.order {
			p(v)
		}
	}
	switch limit, offset := ftq.ctx.Limit, ftq.ctx.Offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := ftq.ctx.Unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// FieldTypeGroupBy is the group-by builder for FieldType entities.
type FieldTypeGroupBy struct {
	selector
	build *FieldTypeQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ftgb *FieldTypeGroupBy) Aggregate(fns ...AggregateFunc) *FieldTypeGroupBy {
	ftgb.fns = append(ftgb.fns, fns...)
	return ftgb
}

// Scan applies the selector query and scans the result into the given value.
func (ftgb *FieldTypeGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ftgb.build.ctx, "GroupBy")
	if err := ftgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FieldTypeQuery, *FieldTypeGroupBy](ctx, ftgb.build, ftgb, ftgb.build.inters, v)
}

func (ftgb *FieldTypeGroupBy) gremlinScan(ctx context.Context, root *FieldTypeQuery, v any) error {
	var (
		trs   []any
		names []any
	)
	for _, fn := range ftgb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range *ftgb.flds {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	query, bindings := root.gremlinQuery(ctx).Group().
		By(__.Values(*ftgb.flds...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next().
		Query()
	res := &gremlin.Response{}
	if err := ftgb.build.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(*ftgb.flds)+len(ftgb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

// FieldTypeSelect is the builder for selecting fields of FieldType entities.
type FieldTypeSelect struct {
	*FieldTypeQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fts *FieldTypeSelect) Aggregate(fns ...AggregateFunc) *FieldTypeSelect {
	fts.fns = append(fts.fns, fns...)
	return fts
}

// Scan applies the selector query and scans the result into the given value.
func (fts *FieldTypeSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fts.ctx, "Select")
	if err := fts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FieldTypeQuery, *FieldTypeSelect](ctx, fts.FieldTypeQuery, fts, fts.inters, v)
}

func (fts *FieldTypeSelect) gremlinScan(ctx context.Context, root *FieldTypeQuery, v any) error {
	var (
		res       = &gremlin.Response{}
		traversal = root.gremlinQuery(ctx)
	)
	if fields := fts.ctx.Fields; len(fields) == 1 {
		if fields[0] != fieldtype.FieldID {
			traversal = traversal.Values(fields...)
		} else {
			traversal = traversal.ID()
		}
	} else {
		fields := make([]any, len(fts.ctx.Fields))
		for i, f := range fts.ctx.Fields {
			fields[i] = f
		}
		traversal = traversal.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := fts.driver.Exec(ctx, query, bindings, res); err != nil {
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
