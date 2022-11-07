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
	"entgo.io/ent/entc/integration/ent/license"
	"entgo.io/ent/entc/integration/ent/predicate"
	"entgo.io/ent/schema/field"
)

// LicenseQuery is the builder for querying License entities.
type LicenseQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.License
	modifiers  []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the LicenseQuery builder.
func (lq *LicenseQuery) Where(ps ...predicate.License) *LicenseQuery {
	lq.predicates = append(lq.predicates, ps...)
	return lq
}

// Limit adds a limit step to the query.
func (lq *LicenseQuery) Limit(limit int) *LicenseQuery {
	lq.limit = &limit
	return lq
}

// Offset adds an offset step to the query.
func (lq *LicenseQuery) Offset(offset int) *LicenseQuery {
	lq.offset = &offset
	return lq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (lq *LicenseQuery) Unique(unique bool) *LicenseQuery {
	lq.unique = &unique
	return lq
}

// Order adds an order step to the query.
func (lq *LicenseQuery) Order(o ...OrderFunc) *LicenseQuery {
	lq.order = append(lq.order, o...)
	return lq
}

// First returns the first License entity from the query.
// Returns a *NotFoundError when no License was found.
func (lq *LicenseQuery) First(ctx context.Context) (*License, error) {
	nodes, err := lq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{license.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (lq *LicenseQuery) FirstX(ctx context.Context) *License {
	node, err := lq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first License ID from the query.
// Returns a *NotFoundError when no License ID was found.
func (lq *LicenseQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{license.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (lq *LicenseQuery) FirstIDX(ctx context.Context) int {
	id, err := lq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single License entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one License entity is found.
// Returns a *NotFoundError when no License entities are found.
func (lq *LicenseQuery) Only(ctx context.Context) (*License, error) {
	nodes, err := lq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{license.Label}
	default:
		return nil, &NotSingularError{license.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (lq *LicenseQuery) OnlyX(ctx context.Context) *License {
	node, err := lq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only License ID in the query.
// Returns a *NotSingularError when more than one License ID is found.
// Returns a *NotFoundError when no entities are found.
func (lq *LicenseQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = lq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{license.Label}
	default:
		err = &NotSingularError{license.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (lq *LicenseQuery) OnlyIDX(ctx context.Context) int {
	id, err := lq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Licenses.
func (lq *LicenseQuery) All(ctx context.Context) ([]*License, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return lq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (lq *LicenseQuery) AllX(ctx context.Context) []*License {
	nodes, err := lq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of License IDs.
func (lq *LicenseQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := lq.Select(license.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (lq *LicenseQuery) IDsX(ctx context.Context) []int {
	ids, err := lq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (lq *LicenseQuery) Count(ctx context.Context) (int, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return lq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (lq *LicenseQuery) CountX(ctx context.Context) int {
	count, err := lq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (lq *LicenseQuery) Exist(ctx context.Context) (bool, error) {
	if err := lq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return lq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (lq *LicenseQuery) ExistX(ctx context.Context) bool {
	exist, err := lq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the LicenseQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (lq *LicenseQuery) Clone() *LicenseQuery {
	if lq == nil {
		return nil
	}
	return &LicenseQuery{
		config:     lq.config,
		limit:      lq.limit,
		offset:     lq.offset,
		order:      append([]OrderFunc{}, lq.order...),
		predicates: append([]predicate.License{}, lq.predicates...),
		// clone intermediate query.
		sql:    lq.sql.Clone(),
		path:   lq.path,
		unique: lq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.License.Query().
//		GroupBy(license.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (lq *LicenseQuery) GroupBy(field string, fields ...string) *LicenseGroupBy {
	grbuild := &LicenseGroupBy{config: lq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := lq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return lq.sqlQuery(ctx), nil
	}
	grbuild.label = license.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.License.Query().
//		Select(license.FieldCreateTime).
//		Scan(ctx, &v)
func (lq *LicenseQuery) Select(fields ...string) *LicenseSelect {
	lq.fields = append(lq.fields, fields...)
	selbuild := &LicenseSelect{LicenseQuery: lq}
	selbuild.label = license.Label
	selbuild.flds, selbuild.scan = &lq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a LicenseSelect configured with the given aggregations.
func (lq *LicenseQuery) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	return lq.Select().Aggregate(fns...)
}

func (lq *LicenseQuery) prepareQuery(ctx context.Context) error {
	for _, f := range lq.fields {
		if !license.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if lq.path != nil {
		prev, err := lq.path(ctx)
		if err != nil {
			return err
		}
		lq.sql = prev
	}
	return nil
}

func (lq *LicenseQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*License, error) {
	var (
		nodes = []*License{}
		_spec = lq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*License).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &License{config: lq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, lq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (lq *LicenseQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := lq.querySpec()
	if len(lq.modifiers) > 0 {
		_spec.Modifiers = lq.modifiers
	}
	_spec.Node.Columns = lq.fields
	if len(lq.fields) > 0 {
		_spec.Unique = lq.unique != nil && *lq.unique
	}
	return sqlgraph.CountNodes(ctx, lq.driver, _spec)
}

func (lq *LicenseQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := lq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (lq *LicenseQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   license.Table,
			Columns: license.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: license.FieldID,
			},
		},
		From:   lq.sql,
		Unique: true,
	}
	if unique := lq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := lq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, license.FieldID)
		for i := range fields {
			if fields[i] != license.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := lq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := lq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := lq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := lq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (lq *LicenseQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(lq.driver.Dialect())
	t1 := builder.Table(license.Table)
	columns := lq.fields
	if len(columns) == 0 {
		columns = license.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if lq.sql != nil {
		selector = lq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if lq.unique != nil && *lq.unique {
		selector.Distinct()
	}
	for _, m := range lq.modifiers {
		m(selector)
	}
	for _, p := range lq.predicates {
		p(selector)
	}
	for _, p := range lq.order {
		p(selector)
	}
	if offset := lq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := lq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (lq *LicenseQuery) ForUpdate(opts ...sql.LockOption) *LicenseQuery {
	if lq.driver.Dialect() == dialect.Postgres {
		lq.Unique(false)
	}
	lq.modifiers = append(lq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return lq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (lq *LicenseQuery) ForShare(opts ...sql.LockOption) *LicenseQuery {
	if lq.driver.Dialect() == dialect.Postgres {
		lq.Unique(false)
	}
	lq.modifiers = append(lq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return lq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (lq *LicenseQuery) Modify(modifiers ...func(s *sql.Selector)) *LicenseSelect {
	lq.modifiers = append(lq.modifiers, modifiers...)
	return lq.Select()
}

// LicenseGroupBy is the group-by builder for License entities.
type LicenseGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (lgb *LicenseGroupBy) Aggregate(fns ...AggregateFunc) *LicenseGroupBy {
	lgb.fns = append(lgb.fns, fns...)
	return lgb
}

// Scan applies the group-by query and scans the result into the given value.
func (lgb *LicenseGroupBy) Scan(ctx context.Context, v any) error {
	query, err := lgb.path(ctx)
	if err != nil {
		return err
	}
	lgb.sql = query
	return lgb.sqlScan(ctx, v)
}

func (lgb *LicenseGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range lgb.fields {
		if !license.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := lgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := lgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (lgb *LicenseGroupBy) sqlQuery() *sql.Selector {
	selector := lgb.sql.Select()
	aggregation := make([]string, 0, len(lgb.fns))
	for _, fn := range lgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(lgb.fields)+len(lgb.fns))
		for _, f := range lgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(lgb.fields...)...)
}

// LicenseSelect is the builder for selecting fields of License entities.
type LicenseSelect struct {
	*LicenseQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ls *LicenseSelect) Aggregate(fns ...AggregateFunc) *LicenseSelect {
	ls.fns = append(ls.fns, fns...)
	return ls
}

// Scan applies the selector query and scans the result into the given value.
func (ls *LicenseSelect) Scan(ctx context.Context, v any) error {
	if err := ls.prepareQuery(ctx); err != nil {
		return err
	}
	ls.sql = ls.LicenseQuery.sqlQuery(ctx)
	return ls.sqlScan(ctx, v)
}

func (ls *LicenseSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ls.fns))
	for _, fn := range ls.fns {
		aggregation = append(aggregation, fn(ls.sql))
	}
	switch n := len(*ls.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ls.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ls.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ls.sql.Query()
	if err := ls.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (ls *LicenseSelect) Modify(modifiers ...func(s *sql.Selector)) *LicenseSelect {
	ls.modifiers = append(ls.modifiers, modifiers...)
	return ls
}
