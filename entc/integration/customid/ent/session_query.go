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
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/customid/ent/device"
	"entgo.io/ent/entc/integration/customid/ent/predicate"
	"entgo.io/ent/entc/integration/customid/ent/schema"
	"entgo.io/ent/entc/integration/customid/ent/session"
	"entgo.io/ent/schema/field"
)

// SessionQuery is the builder for querying Session entities.
type SessionQuery struct {
	config
	ctx        *QueryContext
	order      []session.OrderOption
	inters     []Interceptor
	predicates []predicate.Session
	withDevice *DeviceQuery
	withFKs    bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the SessionQuery builder.
func (q *SessionQuery) Where(ps ...predicate.Session) *SessionQuery {
	q.predicates = append(q.predicates, ps...)
	return q
}

// Limit the number of records to be returned by this query.
func (q *SessionQuery) Limit(limit int) *SessionQuery {
	q.ctx.Limit = &limit
	return q
}

// Offset to start from.
func (q *SessionQuery) Offset(offset int) *SessionQuery {
	q.ctx.Offset = &offset
	return q
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (q *SessionQuery) Unique(unique bool) *SessionQuery {
	q.ctx.Unique = &unique
	return q
}

// Order specifies how the records should be ordered.
func (q *SessionQuery) Order(o ...session.OrderOption) *SessionQuery {
	q.order = append(q.order, o...)
	return q
}

// QueryDevice chains the current query on the "device" edge.
func (q *SessionQuery) QueryDevice() *DeviceQuery {
	query := (&DeviceClient{config: q.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := q.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := q.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(session.Table, session.FieldID, selector),
			sqlgraph.To(device.Table, device.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, session.DeviceTable, session.DeviceColumn),
		)
		fromU = sqlgraph.SetNeighbors(q.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Session entity from the query.
// Returns a *NotFoundError when no Session was found.
func (q *SessionQuery) First(ctx context.Context) (*Session, error) {
	nodes, err := q.Limit(1).All(setContextOp(ctx, q.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{session.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (q *SessionQuery) FirstX(ctx context.Context) *Session {
	node, err := q.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Session ID from the query.
// Returns a *NotFoundError when no Session ID was found.
func (q *SessionQuery) FirstID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = q.Limit(1).IDs(setContextOp(ctx, q.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{session.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (q *SessionQuery) FirstIDX(ctx context.Context) schema.ID {
	id, err := q.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Session entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Session entity is found.
// Returns a *NotFoundError when no Session entities are found.
func (q *SessionQuery) Only(ctx context.Context) (*Session, error) {
	nodes, err := q.Limit(2).All(setContextOp(ctx, q.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{session.Label}
	default:
		return nil, &NotSingularError{session.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (q *SessionQuery) OnlyX(ctx context.Context) *Session {
	node, err := q.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Session ID in the query.
// Returns a *NotSingularError when more than one Session ID is found.
// Returns a *NotFoundError when no entities are found.
func (q *SessionQuery) OnlyID(ctx context.Context) (id schema.ID, err error) {
	var ids []schema.ID
	if ids, err = q.Limit(2).IDs(setContextOp(ctx, q.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{session.Label}
	default:
		err = &NotSingularError{session.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (q *SessionQuery) OnlyIDX(ctx context.Context) schema.ID {
	id, err := q.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Sessions.
func (q *SessionQuery) All(ctx context.Context) ([]*Session, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryAll)
	if err := q.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Session, *SessionQuery]()
	return withInterceptors[[]*Session](ctx, q, qr, q.inters)
}

// AllX is like All, but panics if an error occurs.
func (q *SessionQuery) AllX(ctx context.Context) []*Session {
	nodes, err := q.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Session IDs.
func (q *SessionQuery) IDs(ctx context.Context) (ids []schema.ID, err error) {
	if q.ctx.Unique == nil && q.path != nil {
		q.Unique(true)
	}
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryIDs)
	if err = q.Select(session.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (q *SessionQuery) IDsX(ctx context.Context) []schema.ID {
	ids, err := q.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (q *SessionQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, q.ctx, ent.OpQueryCount)
	if err := q.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, q, querierCount[*SessionQuery](), q.inters)
}

// CountX is like Count, but panics if an error occurs.
func (q *SessionQuery) CountX(ctx context.Context) int {
	count, err := q.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (q *SessionQuery) Exist(ctx context.Context) (bool, error) {
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
func (q *SessionQuery) ExistX(ctx context.Context) bool {
	exist, err := q.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the SessionQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (q *SessionQuery) Clone() *SessionQuery {
	if q == nil {
		return nil
	}
	return &SessionQuery{
		config:     q.config,
		ctx:        q.ctx.Clone(),
		order:      append([]session.OrderOption{}, q.order...),
		inters:     append([]Interceptor{}, q.inters...),
		predicates: append([]predicate.Session{}, q.predicates...),
		withDevice: q.withDevice.Clone(),
		// clone intermediate query.
		sql:  q.sql.Clone(),
		path: q.path,
	}
}

// WithDevice tells the query-builder to eager-load the nodes that are connected to
// the "device" edge. The optional arguments are used to configure the query builder of the edge.
func (q *SessionQuery) WithDevice(opts ...func(*DeviceQuery)) *SessionQuery {
	query := (&DeviceClient{config: q.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	q.withDevice = query
	return q
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
func (q *SessionQuery) GroupBy(field string, fields ...string) *SessionGroupBy {
	q.ctx.Fields = append([]string{field}, fields...)
	grbuild := &SessionGroupBy{build: q}
	grbuild.flds = &q.ctx.Fields
	grbuild.label = session.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
func (q *SessionQuery) Select(fields ...string) *SessionSelect {
	q.ctx.Fields = append(q.ctx.Fields, fields...)
	sbuild := &SessionSelect{SessionQuery: q}
	sbuild.label = session.Label
	sbuild.flds, sbuild.scan = &q.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a SessionSelect configured with the given aggregations.
func (q *SessionQuery) Aggregate(fns ...AggregateFunc) *SessionSelect {
	return q.Select().Aggregate(fns...)
}

func (q *SessionQuery) prepareQuery(ctx context.Context) error {
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
		if !session.ValidColumn(f) {
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

func (q *SessionQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Session, error) {
	var (
		nodes       = []*Session{}
		withFKs     = q.withFKs
		_spec       = q.querySpec()
		loadedTypes = [1]bool{
			q.withDevice != nil,
		}
	)
	if q.withDevice != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, session.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Session).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Session{config: q.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
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
	if query := q.withDevice; query != nil {
		if err := q.loadDevice(ctx, query, nodes, nil,
			func(n *Session, e *Device) { n.Edges.Device = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (q *SessionQuery) loadDevice(ctx context.Context, query *DeviceQuery, nodes []*Session, init func(*Session), assign func(*Session, *Device)) error {
	ids := make([]schema.ID, 0, len(nodes))
	nodeids := make(map[schema.ID][]*Session)
	for i := range nodes {
		if nodes[i].device_sessions == nil {
			continue
		}
		fk := *nodes[i].device_sessions
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(device.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "device_sessions" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (q *SessionQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := q.querySpec()
	_spec.Node.Columns = q.ctx.Fields
	if len(q.ctx.Fields) > 0 {
		_spec.Unique = q.ctx.Unique != nil && *q.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, q.driver, _spec)
}

func (q *SessionQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(session.Table, session.Columns, sqlgraph.NewFieldSpec(session.FieldID, field.TypeBytes))
	_spec.From = q.sql
	if unique := q.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if q.path != nil {
		_spec.Unique = true
	}
	if fields := q.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, session.FieldID)
		for i := range fields {
			if fields[i] != session.FieldID {
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

func (q *SessionQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(q.driver.Dialect())
	t1 := builder.Table(session.Table)
	columns := q.ctx.Fields
	if len(columns) == 0 {
		columns = session.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if q.sql != nil {
		selector = q.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if q.ctx.Unique != nil && *q.ctx.Unique {
		selector.Distinct()
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

// SessionGroupBy is the group-by builder for Session entities.
type SessionGroupBy struct {
	selector
	build *SessionQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (sgb *SessionGroupBy) Aggregate(fns ...AggregateFunc) *SessionGroupBy {
	sgb.fns = append(sgb.fns, fns...)
	return sgb
}

// Scan applies the selector query and scans the result into the given value.
func (sgb *SessionGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, sgb.build.ctx, ent.OpQueryGroupBy)
	if err := sgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SessionQuery, *SessionGroupBy](ctx, sgb.build, sgb, sgb.build.inters, v)
}

func (q *SessionGroupBy) sqlScan(ctx context.Context, root *SessionQuery, v any) error {
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

// SessionSelect is the builder for selecting fields of Session entities.
type SessionSelect struct {
	*SessionQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ss *SessionSelect) Aggregate(fns ...AggregateFunc) *SessionSelect {
	ss.fns = append(ss.fns, fns...)
	return ss
}

// Scan applies the selector query and scans the result into the given value.
func (ss *SessionSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ss.ctx, ent.OpQuerySelect)
	if err := ss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*SessionQuery, *SessionSelect](ctx, ss.SessionQuery, ss, ss.inters, v)
}

func (q *SessionSelect) sqlScan(ctx context.Context, root *SessionQuery, v any) error {
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
