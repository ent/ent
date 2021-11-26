// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"math"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/__"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"

	"entgo.io/ent/entc/integration/gremlin/ent/goods"
	"entgo.io/ent/entc/integration/gremlin/ent/predicate"
)

// GoodsQuery is the builder for querying Goods entities.
type GoodsQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.Goods
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Where adds a new predicate for the GoodsQuery builder.
func (gq *GoodsQuery) Where(ps ...predicate.Goods) *GoodsQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit adds a limit step to the query.
func (gq *GoodsQuery) Limit(limit int) *GoodsQuery {
	gq.limit = &limit
	return gq
}

// Offset adds an offset step to the query.
func (gq *GoodsQuery) Offset(offset int) *GoodsQuery {
	gq.offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GoodsQuery) Unique(unique bool) *GoodsQuery {
	gq.unique = &unique
	return gq
}

// Order adds an order step to the query.
func (gq *GoodsQuery) Order(o ...OrderFunc) *GoodsQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Goods entity from the query.
// Returns a *NotFoundError when no Goods was found.
func (gq *GoodsQuery) First(ctx context.Context) (*Goods, error) {
	nodes, err := gq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{goods.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gq *GoodsQuery) FirstX(ctx context.Context) *Goods {
	node, err := gq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Goods ID from the query.
// Returns a *NotFoundError when no Goods ID was found.
func (gq *GoodsQuery) FirstID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goods.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GoodsQuery) FirstIDX(ctx context.Context) string {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Goods entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one Goods entity is not found.
// Returns a *NotFoundError when no Goods entities are found.
func (gq *GoodsQuery) Only(ctx context.Context) (*Goods, error) {
	nodes, err := gq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{goods.Label}
	default:
		return nil, &NotSingularError{goods.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gq *GoodsQuery) OnlyX(ctx context.Context) *Goods {
	node, err := gq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Goods ID in the query.
// Returns a *NotSingularError when exactly one Goods ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gq *GoodsQuery) OnlyID(ctx context.Context) (id string, err error) {
	var ids []string
	if ids, err = gq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = &NotSingularError{goods.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gq *GoodsQuery) OnlyIDX(ctx context.Context) string {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodsSlice.
func (gq *GoodsQuery) All(ctx context.Context) ([]*Goods, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gq.gremlinAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gq *GoodsQuery) AllX(ctx context.Context) []*Goods {
	nodes, err := gq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Goods IDs.
func (gq *GoodsQuery) IDs(ctx context.Context) ([]string, error) {
	var ids []string
	if err := gq.Select(goods.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GoodsQuery) IDsX(ctx context.Context) []string {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GoodsQuery) Count(ctx context.Context) (int, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gq.gremlinCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gq *GoodsQuery) CountX(ctx context.Context) int {
	count, err := gq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gq *GoodsQuery) Exist(ctx context.Context) (bool, error) {
	if err := gq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gq.gremlinExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gq *GoodsQuery) ExistX(ctx context.Context) bool {
	exist, err := gq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GoodsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gq *GoodsQuery) Clone() *GoodsQuery {
	if gq == nil {
		return nil
	}
	return &GoodsQuery{
		config:     gq.config,
		limit:      gq.limit,
		offset:     gq.offset,
		order:      append([]OrderFunc{}, gq.order...),
		predicates: append([]predicate.Goods{}, gq.predicates...),
		// clone intermediate query.
		gremlin: gq.gremlin.Clone(),
		path:    gq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (gq *GoodsQuery) GroupBy(field string, fields ...string) *GoodsGroupBy {
	group := &GoodsGroupBy{config: gq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *dsl.Traversal, err error) {
		if err := gq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gq.gremlinQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (gq *GoodsQuery) Select(fields ...string) *GoodsSelect {
	gq.fields = append(gq.fields, fields...)
	return &GoodsSelect{GoodsQuery: gq}
}

func (gq *GoodsQuery) prepareQuery(ctx context.Context) error {
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.gremlin = prev
	}
	return nil
}

func (gq *GoodsQuery) gremlinAll(ctx context.Context) ([]*Goods, error) {
	res := &gremlin.Response{}
	traversal := gq.gremlinQuery(ctx)
	if len(gq.fields) > 0 {
		fields := make([]interface{}, len(gq.fields))
		for i, f := range gq.fields {
			fields[i] = f
		}
		traversal.ValueMap(fields...)
	} else {
		traversal.ValueMap(true)
	}
	query, bindings := traversal.Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	var _gos GoodsSlice
	if err := _gos.FromResponse(res); err != nil {
		return nil, err
	}
	_gos.config(gq.config)
	return _gos, nil
}

func (gq *GoodsQuery) gremlinCount(ctx context.Context) (int, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery(ctx).Count().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return 0, err
	}
	return res.ReadInt()
}

func (gq *GoodsQuery) gremlinExist(ctx context.Context) (bool, error) {
	res := &gremlin.Response{}
	query, bindings := gq.gremlinQuery(ctx).HasNext().Query()
	if err := gq.driver.Exec(ctx, query, bindings, res); err != nil {
		return false, err
	}
	return res.ReadBool()
}

func (gq *GoodsQuery) gremlinQuery(context.Context) *dsl.Traversal {
	v := g.V().HasLabel(goods.Label)
	if gq.gremlin != nil {
		v = gq.gremlin.Clone()
	}
	for _, p := range gq.predicates {
		p(v)
	}
	if len(gq.order) > 0 {
		v.Order()
		for _, p := range gq.order {
			p(v)
		}
	}
	switch limit, offset := gq.limit, gq.offset; {
	case limit != nil && offset != nil:
		v.Range(*offset, *offset+*limit)
	case offset != nil:
		v.Range(*offset, math.MaxInt32)
	case limit != nil:
		v.Limit(*limit)
	}
	if unique := gq.unique; unique == nil || *unique {
		v.Dedup()
	}
	return v
}

// GoodsGroupBy is the group-by builder for Goods entities.
type GoodsGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
	path    func(context.Context) (*dsl.Traversal, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GoodsGroupBy) Aggregate(fns ...AggregateFunc) *GoodsGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the group-by query and scans the result into the given value.
func (ggb *GoodsGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ggb.path(ctx)
	if err != nil {
		return err
	}
	ggb.gremlin = query
	return ggb.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ggb *GoodsGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := ggb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GoodsGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ggb *GoodsGroupBy) StringsX(ctx context.Context) []string {
	v, err := ggb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ggb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ggb *GoodsGroupBy) StringX(ctx context.Context) string {
	v, err := ggb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GoodsGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ggb *GoodsGroupBy) IntsX(ctx context.Context) []int {
	v, err := ggb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ggb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ggb *GoodsGroupBy) IntX(ctx context.Context) int {
	v, err := ggb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GoodsGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ggb *GoodsGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := ggb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ggb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ggb *GoodsGroupBy) Float64X(ctx context.Context) float64 {
	v, err := ggb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(ggb.fields) > 1 {
		return nil, errors.New("ent: GoodsGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := ggb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ggb *GoodsGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := ggb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (ggb *GoodsGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ggb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ggb *GoodsGroupBy) BoolX(ctx context.Context) bool {
	v, err := ggb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ggb *GoodsGroupBy) gremlinScan(ctx context.Context, v interface{}) error {
	res := &gremlin.Response{}
	query, bindings := ggb.gremlinQuery().Query()
	if err := ggb.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(ggb.fields)+len(ggb.fns) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}

func (ggb *GoodsGroupBy) gremlinQuery() *dsl.Traversal {
	var (
		trs   []interface{}
		names []interface{}
	)
	for _, fn := range ggb.fns {
		name, tr := fn("p", "")
		trs = append(trs, tr)
		names = append(names, name)
	}
	for _, f := range ggb.fields {
		names = append(names, f)
		trs = append(trs, __.As("p").Unfold().Values(f).As(f))
	}
	return ggb.gremlin.Group().
		By(__.Values(ggb.fields...).Fold()).
		By(__.Fold().Match(trs...).Select(names...)).
		Select(dsl.Values).
		Next()
}

// GoodsSelect is the builder for selecting fields of Goods entities.
type GoodsSelect struct {
	*GoodsQuery
	// intermediate query (i.e. traversal path).
	gremlin *dsl.Traversal
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GoodsSelect) Scan(ctx context.Context, v interface{}) error {
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	gs.gremlin = gs.GoodsQuery.gremlinQuery(ctx)
	return gs.gremlinScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gs *GoodsSelect) ScanX(ctx context.Context, v interface{}) {
	if err := gs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Strings(ctx context.Context) ([]string, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GoodsSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gs *GoodsSelect) StringsX(ctx context.Context) []string {
	v, err := gs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gs *GoodsSelect) StringX(ctx context.Context) string {
	v, err := gs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Ints(ctx context.Context) ([]int, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GoodsSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gs *GoodsSelect) IntsX(ctx context.Context) []int {
	v, err := gs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gs *GoodsSelect) IntX(ctx context.Context) int {
	v, err := gs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GoodsSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gs *GoodsSelect) Float64sX(ctx context.Context) []float64 {
	v, err := gs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gs *GoodsSelect) Float64X(ctx context.Context) float64 {
	v, err := gs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gs.fields) > 1 {
		return nil, errors.New("ent: GoodsSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gs *GoodsSelect) BoolsX(ctx context.Context) []bool {
	v, err := gs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gs *GoodsSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{goods.Label}
	default:
		err = fmt.Errorf("ent: GoodsSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gs *GoodsSelect) BoolX(ctx context.Context) bool {
	v, err := gs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gs *GoodsSelect) gremlinScan(ctx context.Context, v interface{}) error {
	var (
		traversal *dsl.Traversal
		res       = &gremlin.Response{}
	)
	if len(gs.fields) == 1 {
		if gs.fields[0] != goods.FieldID {
			traversal = gs.gremlin.Values(gs.fields...)
		} else {
			traversal = gs.gremlin.ID()
		}
	} else {
		fields := make([]interface{}, len(gs.fields))
		for i, f := range gs.fields {
			fields[i] = f
		}
		traversal = gs.gremlin.ValueMap(fields...)
	}
	query, bindings := traversal.Query()
	if err := gs.driver.Exec(ctx, query, bindings, res); err != nil {
		return err
	}
	if len(gs.fields) == 1 {
		return res.ReadVal(v)
	}
	vm, err := res.ReadValueMap()
	if err != nil {
		return err
	}
	return vm.Decode(v)
}
