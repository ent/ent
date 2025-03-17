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
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/ent/exvaluescan"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// ExValueScanQuery is the builder for querying ExValueScan entities.
type ExValueScanQuery struct {
	config
	ctx        *QueryContext
	order      []exvaluescan.OrderOption
	inters     []Interceptor
	predicates []predicate.ExValueScan
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ExValueScanQuery builder.
func (q *ExValueScanQuery) Where(ps ...predicate.ExValueScan) *ExValueScanQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *ExValueScanQuery) Limit(limit int) *ExValueScanQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *ExValueScanQuery) Offset(offset int) *ExValueScanQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *ExValueScanQuery) Unique(unique bool) *ExValueScanQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *ExValueScanQuery) Order(o ...exvaluescan.OrderOption) *ExValueScanQuery {
	q.order = append(q.order, o...)
	return q
}

// First returns the first ExValueScan entity from the query.
// Returns a *NotFoundError when no ExValueScan was found.
func (q *ExValueScanQuery) First(ctx context.Context) (*ExValueScan, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{exvaluescan.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *ExValueScanQuery) FirstX(ctx context.Context) *ExValueScan {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ExValueScan ID from the query.
// Returns a *NotFoundError when no ExValueScan ID was found.
func (q *ExValueScanQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{exvaluescan.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *ExValueScanQuery) FirstIDX(ctx context.Context) int {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ExValueScan entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ExValueScan entity is found.
// Returns a *NotFoundError when no ExValueScan entities are found.
func (q *ExValueScanQuery) Only(ctx context.Context) (*ExValueScan, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{exvaluescan.Label}
	default:
		return nil, &NotSingularError{exvaluescan.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *ExValueScanQuery) OnlyX(ctx context.Context) *ExValueScan {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ExValueScan ID in the query.
// Returns a *NotSingularError when more than one ExValueScan ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *ExValueScanQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{exvaluescan.Label}
	default:
		err = &NotSingularError{exvaluescan.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *ExValueScanQuery) OnlyIDX(ctx context.Context) int {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ExValueScans.
func (q *ExValueScanQuery) All(ctx context.Context) ([]*ExValueScan, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ExValueScan, *ExValueScanQuery]()
	return withInterceptors[[]*ExValueScan](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *ExValueScanQuery) AllX(ctx context.Context) []*ExValueScan {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ExValueScan IDs.
func (q *ExValueScanQuery) IDs(ctx context.Context) (ids []int, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(exvaluescan.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *ExValueScanQuery) IDsX(ctx context.Context) []int {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *ExValueScanQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*ExValueScanQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *ExValueScanQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *ExValueScanQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *ExValueScanQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ExValueScanQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *ExValueScanQuery) Clone() *ExValueScanQuery {
	if q == nil {
		return nil
	}
	return &ExValueScanQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]exvaluescan.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.ExValueScan{}, q.predicates...),
		// clone intermediate query.
		sql:       q.sql.Clone(),
		path:      q.path,
		modifiers: append([]func(*sql.Selector){}, q.modifiers...),
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Binary *url.URL `json:"binary,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ExValueScan.Query().
//		GroupBy(exvaluescan.FieldBinary).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (q *ExValueScanQuery) GroupBy(field string, fields ...string) *ExValueScanGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ExValueScanGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = exvaluescan.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Binary *url.URL `json:"binary,omitempty"`
//	}
//
//	client.ExValueScan.Query().
//		Select(exvaluescan.FieldBinary).
//		Scan(ctx, &v)
func (q *ExValueScanQuery) Select(fields ...string) *ExValueScanSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &ExValueScanSelect{ExValueScanQuery: q}
	sbuild.label = exvaluescan.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ExValueScanSelect configured with the given aggregations.
func (q *ExValueScanQuery) Aggregate(fns ...AggregateFunc) *ExValueScanSelect {
	return q.Select().Aggregate(fns...)
}

func (q *ExValueScanQuery) prepareQuery(ctx context.Context) error {
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
	for _, f := range q.ctx.Fields {
		if !exvaluescan.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if q.path != nil {
		prev, err := q.path(ctx)
		if err != nil {
			return err
		}
		q.sql = prev
	}
	return nil
}

func (q *ExValueScanQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ExValueScan, error) {
	var (
		nodes = []*ExValueScan{}
		_spec = q.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ExValueScan).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ExValueScan{config: q.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(q.modifiers) > 0 {
		_spec.Modifiers = q.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, q.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (q *ExValueScanQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	if len(q.modifiers) > 0 {
		_spec.Modifiers = q.modifiers
	}
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *ExValueScanQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(exvaluescan.Table, exvaluescan.Columns, sqlgraph.NewFieldSpec(exvaluescan.FieldID, field.TypeInt))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, exvaluescan.FieldID)
		for i := range fields {
			if fields[i] != exvaluescan.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := q.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := q.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := q.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := q.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (q *ExValueScanQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(exvaluescan.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = exvaluescan.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if q.sql != nil {
		selector = q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if q.ctx.Unique != nil && *q.ctx.Unique {
		selector.Distinct()
	}
	for _, m := range q.modifiers {
		m(selector)
	}
	for _, p := range q.predicates {
		p(selector)
	}
	for _, p := range q.order {
		p(selector)
	}
	if offset := q.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := q.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (evsq *ExValueScanQuery) ForUpdate(opts ...sql.LockOption) *ExValueScanQuery {
	if evsq.driver.Dialect() == dialect.Postgres {
		evsq.Unique(false)
	}
	evsq.modifiers = append(evsq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return evsq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (evsq *ExValueScanQuery) ForShare(opts ...sql.LockOption) *ExValueScanQuery {
	if evsq.driver.Dialect() == dialect.Postgres {
		evsq.Unique(false)
	}
	evsq.modifiers = append(evsq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return evsq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (q *ExValueScanQuery) Modify(modifiers ...func(s *sql.Selector)) *ExValueScanSelect {
	q.modifiers = append(q.modifiers, modifiers...)
	return q.Select()
}

// ExValueScanGroupBy is the group-by builder for ExValueScan entities.
type ExValueScanGroupBy struct {
	selector
	build *ExValueScanQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (evsgb *ExValueScanGroupBy) Aggregate(fns ...AggregateFunc) *ExValueScanGroupBy {
	evsgb.fns = append(evsgb.fns, fns...)
	return evsgb
}

// Scan applies the selector query and scans the result into the given value.
func (evsgb *ExValueScanGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, evsgb.build.ctx, ent.OpQueryGroupBy)
	if err := evsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExValueScanQuery, *ExValueScanGroupBy](ctx, evsgb.build, evsgb, evsgb.build.inters, v)
}

func (q *ExValueScanGroupBy) sqlScan(ctx context.Context, root *ExValueScanQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(q.fns))
	for _, fn := range q.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*q.flds)+len(q.fns))
		for _, f := range *q.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*q.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := q.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ExValueScanSelect is the builder for selecting fields of ExValueScan entities.
type ExValueScanSelect struct {
	*ExValueScanQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (evss *ExValueScanSelect) Aggregate(fns ...AggregateFunc) *ExValueScanSelect {
	evss.fns = append(evss.fns, fns...)
	return evss
}

// Scan applies the selector query and scans the result into the given value.
func (evss *ExValueScanSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, evss.ctx, ent.OpQuerySelect)
	if err := evss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ExValueScanQuery, *ExValueScanSelect](ctx, evss.ExValueScanQuery, evss, evss.inters, v)
}

func (q *ExValueScanSelect) sqlScan(ctx context.Context, root *ExValueScanQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(q.fns))
	for _, fn := range q.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*q.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := q.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (q *ExValueScanSelect) Modify(modifiers ...func(s *sql.Selector)) *ExValueScanSelect {
	q.modifiers = append(q.modifiers, modifiers...)
	return q
}
