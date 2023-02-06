// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/goods"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// GoodsQuery is the builder for querying Goods entities.
type GoodsQuery struct {
	config
	ctx        *QueryContext
	order      []OrderFunc
	inters     []Interceptor
	predicates []predicate.Goods
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GoodsQuery builder.
func (gq *GoodsQuery) Where(ps ...predicate.Goods) *GoodsQuery {
	gq.predicates = append(gq.predicates, ps...)
	return gq
}

// Limit the number of records to be returned by this query.
func (gq *GoodsQuery) Limit(limit int) *GoodsQuery {
	gq.ctx.Limit = &limit
	return gq
}

// Offset to start from.
func (gq *GoodsQuery) Offset(offset int) *GoodsQuery {
	gq.ctx.Offset = &offset
	return gq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gq *GoodsQuery) Unique(unique bool) *GoodsQuery {
	gq.ctx.Unique = &unique
	return gq
}

// Order specifies how the records should be ordered.
func (gq *GoodsQuery) Order(o ...OrderFunc) *GoodsQuery {
	gq.order = append(gq.order, o...)
	return gq
}

// First returns the first Goods entity from the query.
// Returns a *NotFoundError when no Goods was found.
func (gq *GoodsQuery) First(ctx context.Context) (*Goods, error) {
	nodes, err := gq.Limit(1).All(setContextOp(ctx, gq.ctx, "First"))
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
func (gq *GoodsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(1).IDs(setContextOp(ctx, gq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{goods.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gq *GoodsQuery) FirstIDX(ctx context.Context) int {
	id, err := gq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Goods entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Goods entity is found.
// Returns a *NotFoundError when no Goods entities are found.
func (gq *GoodsQuery) Only(ctx context.Context) (*Goods, error) {
	nodes, err := gq.Limit(2).All(setContextOp(ctx, gq.ctx, "Only"))
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
// Returns a *NotSingularError when more than one Goods ID is found.
// Returns a *NotFoundError when no entities are found.
func (gq *GoodsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = gq.Limit(2).IDs(setContextOp(ctx, gq.ctx, "OnlyID")); err != nil {
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
func (gq *GoodsQuery) OnlyIDX(ctx context.Context) int {
	id, err := gq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GoodsSlice.
func (gq *GoodsQuery) All(ctx context.Context) ([]*Goods, error) {
	ctx = setContextOp(ctx, gq.ctx, "All")
	if err := gq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Goods, *GoodsQuery]()
	return withInterceptors[[]*Goods](ctx, gq, qr, gq.inters)
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
func (gq *GoodsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if gq.ctx.Unique == nil && gq.path != nil {
		gq.Unique(true)
	}
	ctx = setContextOp(ctx, gq.ctx, "IDs")
	if err = gq.Select(goods.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gq *GoodsQuery) IDsX(ctx context.Context) []int {
	ids, err := gq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gq *GoodsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, gq.ctx, "Count")
	if err := gq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, gq, querierCount[*GoodsQuery](), gq.inters)
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
	ctx = setContextOp(ctx, gq.ctx, "Exist")
	switch _, err := gq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
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
		ctx:        gq.ctx.Clone(),
		order:      append([]OrderFunc{}, gq.order...),
		inters:     append([]Interceptor{}, gq.inters...),
		predicates: append([]predicate.Goods{}, gq.predicates...),
		// clone intermediate query.
		sql:  gq.sql.Clone(),
		path: gq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (gq *GoodsQuery) GroupBy(field string, fields ...string) *GoodsGroupBy {
	gq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &GoodsGroupBy{build: gq}
	grbuild.flds = &gq.ctx.Fields
	grbuild.label = goods.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (gq *GoodsQuery) Select(fields ...string) *GoodsSelect {
	gq.ctx.Fields = append(gq.ctx.Fields, fields...)
	sbuild := &GoodsSelect{GoodsQuery: gq}
	sbuild.label = goods.Label
	sbuild.flds, sbuild.scan = &gq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a GoodsSelect configured with the given aggregations.
func (gq *GoodsQuery) Aggregate(fns ...AggregateFunc) *GoodsSelect {
	return gq.Select().Aggregate(fns...)
}

func (gq *GoodsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range gq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, gq); err != nil {
				return err
			}
		}
	}
	for _, f := range gq.ctx.Fields {
		if !goods.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if gq.path != nil {
		prev, err := gq.path(ctx)
		if err != nil {
			return err
		}
		gq.sql = prev
	}
	return nil
}

func (gq *GoodsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Goods, error) {
	var (
		nodes = []*Goods{}
		_spec = gq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Goods).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Goods{config: gq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, gq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gq *GoodsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gq.querySpec()
	if len(gq.modifiers) > 0 {
		_spec.Modifiers = gq.modifiers
	}
	_spec.Node.Columns = gq.ctx.Fields
	if len(gq.ctx.Fields) > 0 {
		_spec.Unique = gq.ctx.Unique != nil && *gq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, gq.driver, _spec)
}

func (gq *GoodsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   goods.Table,
			Columns: goods.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: goods.FieldID,
			},
		},
		From:   gq.sql,
		Unique: true,
	}
	if unique := gq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, goods.FieldID)
		for i := range fields {
			if fields[i] != goods.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gq *GoodsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gq.driver.Dialect())
	t1 := builder.Table(goods.Table)
	columns := gq.ctx.Fields
	if len(columns) == 0 {
		columns = goods.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gq.sql != nil {
		selector = gq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if gq.ctx.Unique != nil && *gq.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range gq.modifiers {
		m(selector)
	}
	for _, p := range gq.predicates {
		p(selector)
	}
	for _, p := range gq.order {
		p(selector)
	}
	if offset := gq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (gq *GoodsQuery) ForUpdate(opts ...sql.LockOption) *GoodsQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return gq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (gq *GoodsQuery) ForShare(opts ...sql.LockOption) *GoodsQuery {
	if gq.driver.Dialect() == dialect.Postgres {
		gq.Unique(false)
	}
	gq.modifiers = append(gq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return gq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gq *GoodsQuery) Modify(modifiers ...func(s *sql.Selector)) *GoodsSelect {
	gq.modifiers = append(gq.modifiers, modifiers...)
	return gq.Select()
}

// GoodsGroupBy is the group-by builder for Goods entities.
type GoodsGroupBy struct {
	selector
	build *GoodsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ggb *GoodsGroupBy) Aggregate(fns ...AggregateFunc) *GoodsGroupBy {
	ggb.fns = append(ggb.fns, fns...)
	return ggb
}

// Scan applies the selector query and scans the result into the given value.
func (ggb *GoodsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ggb.build.ctx, "GroupBy")
	if err := ggb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodsQuery, *GoodsGroupBy](ctx, ggb.build, ggb, ggb.build.inters, v)
}

func (ggb *GoodsGroupBy) sqlScan(ctx context.Context, root *GoodsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ggb.fns))
	for _, fn := range ggb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ggb.flds)+len(ggb.fns))
		for _, f := range *ggb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ggb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ggb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// GoodsSelect is the builder for selecting fields of Goods entities.
type GoodsSelect struct {
	*GoodsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (gs *GoodsSelect) Aggregate(fns ...AggregateFunc) *GoodsSelect {
	gs.fns = append(gs.fns, fns...)
	return gs
}

// Scan applies the selector query and scans the result into the given value.
func (gs *GoodsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, gs.ctx, "Select")
	if err := gs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*GoodsQuery, *GoodsSelect](ctx, gs.GoodsQuery, gs, gs.inters, v)
}

func (gs *GoodsSelect) sqlScan(ctx context.Context, root *GoodsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(gs.fns))
	for _, fn := range gs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*gs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (gs *GoodsSelect) Modify(modifiers ...func(s *sql.Selector)) *GoodsSelect {
	gs.modifiers = append(gs.modifiers, modifiers...)
	return gs
}
